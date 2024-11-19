package pingone

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	// "github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/patrickcping/pingone-go-sdk-v2/pingone/model"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const defaultJwtExpirationTimeMins = 15

func (c *Config) getToken(ctx context.Context) error {

	if checkForValue(c.AccessToken) {
		accessTokenObject := &oauth2.Token{
			AccessToken: *c.AccessToken,
			TokenType:   "Bearer",
		}

		if accessTokenObject.Valid() {
			c.accessTokenObject = accessTokenObject
			return nil
		} else {
			return fmt.Errorf("Required parameter missing.  Access token is not valid.")
		}
	}

	if !checkForValue(c.ClientID) {
		return fmt.Errorf("Required parameter missing.  Must provide a valid Client ID.")
	}

	if !checkForValue(c.EnvironmentID) {
		return fmt.Errorf("Required parameter missing.  Must provide a valid Environment ID that contains the client application used to authenticate to the API.")
	}

	if !checkForValue(c.Region) && !checkForValue(c.RegionCode) {
		return fmt.Errorf("Required parameter missing.  Must provide a valid Region code.")
	}

	// here, do client ID/secret or JWT bearer
}

func (c *Config) getTokenClientIDClientSecret(ctx context.Context) error {

	if !checkForValue(c.ClientID) || !checkForValue(c.ClientSecret) || !checkForValue(c.EnvironmentID) || (!checkForValue(c.Region) && !checkForValue(c.RegionCode)) {
		return fmt.Errorf("Required parameter missing.  Must provide ClientID, ClientSecret, EnvironmentID and Region.")
	}

	//Get URL from SDK
	authURL := c.tokenUrl()

	//OAuth 2.0 config for client creds
	config := clientcredentials.Config{
		ClientID:     *c.ClientID,
		ClientSecret: *c.ClientSecret,
		TokenURL:     fmt.Sprintf("%s/%s/as/token", authURL, *c.EnvironmentID),
		AuthStyle:    oauth2.AuthStyleAutoDetect,
	}

	token, err := processResponse(func() (*oauth2.Token, error) {

		if v := c.ProxyURL; v != nil && *v != "" {

			// Parse the proxy URL
			proxyURLParsed, err := url.Parse(*v)
			if err != nil {
				return nil, fmt.Errorf("Failed to parse proxy URL: %s", err)
			}

			// Create a new Transport object with the proxy settings
			transport := &http.Transport{
				Proxy: http.ProxyURL(proxyURLParsed),
			}

			// Create a new HTTP client using the custom Transport
			httpClient := &http.Client{
				Transport: transport,
			}
			ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)
		}

		return config.Token(ctx)
	})
	if err != nil {
		return err
	}

	c.accessTokenObject = token

	return nil
}

func (c *Config) getTokenJWTBearer(ctx context.Context) error {

	if !checkForValue(c.ClientID) || !checkForValue(c.ClientSecret) || !checkForValue(c.EnvironmentID) || (!checkForValue(c.Region) && !checkForValue(c.RegionCode)) {
		return fmt.Errorf("Required parameter missing.  Must provide ClientID, ClientSecret, EnvironmentID and Region.")
	}

	//Generate and sign the request JWT
	requestJwt, err := c.buildJwt()
	if err != nil {
		return fmt.Errorf("Failed to generate JWT: %s", err)
	}

	token, err := processResponse(func() (*oauth2.Token, error) {

		httpClient := &http.Client{}

		if v := c.ProxyURL; v != nil && *v != "" {

			// Parse the proxy URL
			proxyURLParsed, err := url.Parse(*v)
			if err != nil {
				return nil, fmt.Errorf("Failed to parse proxy URL: %s", err)
			}

			// Create a new Transport object with the proxy settings
			transport := &http.Transport{
				Proxy: http.ProxyURL(proxyURLParsed),
			}

			// Create a new HTTP client using the custom Transport
			httpClient.Transport = transport
		}

		// invoke the http request here
		formData := url.Values{}
		formData.Add("client_id", *c.ClientID)
		formData.Add("grant_type", "client_credentials")
		formData.Add("client_assertion", *requestJwt)
		formData.Add("client_assertion_type", "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")

		httpResp, err := httpClient.Post(c.tokenUrl(), "application/x-www-form-urlencoded", strings.NewReader(formData.Encode()))
		if err != nil {
			return nil, fmt.Errorf("failed to request access token with service account: %w", err)
		}

		defer httpResp.Body.Close()
		body, err := io.ReadAll(httpResp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read access token response body: %w", err)
		}
		if httpResp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to request access token with service account. Status code: %d, body: %s", httpResp.StatusCode, body)
		}

		var token oauth2.Token
		err = json.Unmarshal(body, &token)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal access token response: %w, body: %s", err, body)
		}

		token.Expiry = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)

		return &token, nil
	})
	if err != nil {
		return err
	}

	c.accessTokenObject = token

	return nil

}

// PingOne supports RS256, RS384, or RS512 using *rsa.PrivateKey: https://golang-jwt.github.io/jwt/usage/signing_methods/#signing-methods-and-key-types
func (c *Config) buildJwt() (*string, error) {
	key, err := jwk.ParseKey(*c.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	tok, err := jwt.NewBuilder().
		Issuer(*c.ClientID).
		Subject(*c.ClientID).
		Expiration(time.Now().Add(time.Second * 899)).
		// JwtID(uuid.New().String()).
		Audience([]string{c.tokenUrl()}).
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build jwt: %w", err)
	}

	token, err := jwt.Sign(tok, jwt.WithKey(jwa.RS512, key))
	if err != nil {
		return nil, fmt.Errorf("failed to sign jwt: %w", err)
	}

	tokenString := string(token)

	return &tokenString, nil
}

func (c *Config) tokenUrl() string {

	var region model.RegionMapping
	if c.Region != "" {
		region = model.FindRegionByName(c.Region) //nolint:staticcheck
	}

	if v := c.RegionCode; v != nil {
		region = model.FindRegionByAPICode(*v)
	}

	regionSuffix := region.URLSuffix

	authURL := fmt.Sprintf("https://auth.pingone.%s", regionSuffix)
	if checkForValue(c.AuthHostnameOverride) {
		authURL = fmt.Sprintf("https://%s", *c.AuthHostnameOverride)
	}

	return fmt.Sprintf("%s/%s/as/token", authURL, *c.EnvironmentID)
}
