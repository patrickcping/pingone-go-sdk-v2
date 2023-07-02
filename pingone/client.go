package pingone

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement"
	"github.com/patrickcping/pingone-go-sdk-v2/authorize"
	"github.com/patrickcping/pingone-go-sdk-v2/credentials"
	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/patrickcping/pingone-go-sdk-v2/mfa"
	"github.com/patrickcping/pingone-go-sdk-v2/pingone/model"
	"github.com/patrickcping/pingone-go-sdk-v2/risk"
	"github.com/patrickcping/pingone-go-sdk-v2/verify"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func (c *Config) APIClient(ctx context.Context) (*Client, error) {

	agreementManagementClient, err := c.AgreementManagementAPIClient(ctx)
	if err != nil {
		return nil, err
	}

	authorizeClient, err := c.AuthorizeAPIClient(ctx)
	if err != nil {
		return nil, err
	}

	credentialsClient, err := c.CredentialsAPIClient(ctx)
	if err != nil {
		return nil, err
	}

	managementClient, err := c.ManagementAPIClient(ctx)
	if err != nil {
		return nil, err
	}

	mfaClient, err := c.MFAAPIClient(ctx)
	if err != nil {
		return nil, err
	}

	riskClient, err := c.RiskAPIClient(ctx)
	if err != nil {
		return nil, err
	}

	verifyClient, err := c.VerifyAPIClient(ctx)
	if err != nil {
		return nil, err
	}

	apiClient := &Client{
		AgreementManagementAPIClient: agreementManagementClient,
		AuthorizeAPIClient:           authorizeClient,
		CredentialsAPIClient:         credentialsClient,
		ManagementAPIClient:          managementClient,
		MFAAPIClient:                 mfaClient,
		RiskAPIClient:                riskClient,
		VerifyAPIClient:              verifyClient,
		Region:                       model.FindRegionByName(c.Region),
	}

	return apiClient, nil
}

// Deprecated: Use (c *Config).AgreementManagementAPIClient() instead
func AgreementManagementAPIClient(token *oauth2.Token) (*agreementmanagement.APIClient, error) {

	var client *agreementmanagement.APIClient

	clientcfg := agreementmanagement.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	client = agreementmanagement.NewAPIClient(clientcfg)

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise PingOne Agreement Management client")
	}

	return client, nil

}

func (c *Config) AgreementManagementAPIClient(ctx context.Context) (*agreementmanagement.APIClient, error) {

	if !c.accessTokenObject.Valid() {
		err := c.getToken(ctx)
		if err != nil {
			return nil, err
		}
	}

	var client *agreementmanagement.APIClient

	clientcfg := agreementmanagement.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", c.accessTokenObject.AccessToken))

	if c.AgreementMgmtHostnameOverride != nil && *c.AgreementMgmtHostnameOverride != "" {
		clientcfg.SetDefaultServerIndex(1)
		err := clientcfg.SetDefaultServerVariableDefaultValue("baseHostname", *c.AgreementMgmtHostnameOverride)
		if err != nil {
			return nil, err
		}
	} else {
		clientcfg.SetDefaultServerIndex(0)
		err := clientcfg.SetDefaultServerVariableDefaultValue("suffix", model.FindRegionByName(c.Region).URLSuffix)
		if err != nil {
			return nil, err
		}
	}
	client = agreementmanagement.NewAPIClient(clientcfg)

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise PingOne Agreement Management client")
	}

	return client, nil
}

// Deprecated: Use (c *Config).AuthorizeAPIClient() instead
func AuthorizeAPIClient(token *oauth2.Token) (*authorize.APIClient, error) {

	var client *authorize.APIClient

	clientcfg := authorize.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	client = authorize.NewAPIClient(clientcfg)

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise PingOne Authorize client")
	}

	return client, nil

}

func (c *Config) AuthorizeAPIClient(ctx context.Context) (*authorize.APIClient, error) {

	if !c.accessTokenObject.Valid() {
		err := c.getToken(ctx)
		if err != nil {
			return nil, err
		}
	}

	var client *authorize.APIClient

	clientcfg := authorize.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", c.accessTokenObject.AccessToken))

	if c.APIHostnameOverride != nil && *c.APIHostnameOverride != "" {
		clientcfg.SetDefaultServerIndex(1)
		err := clientcfg.SetDefaultServerVariableDefaultValue("baseHostname", *c.APIHostnameOverride)
		if err != nil {
			return nil, err
		}
	} else {
		clientcfg.SetDefaultServerIndex(0)
		err := clientcfg.SetDefaultServerVariableDefaultValue("suffix", model.FindRegionByName(c.Region).URLSuffix)
		if err != nil {
			return nil, err
		}
	}
	client = authorize.NewAPIClient(clientcfg)

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise PingOne Authorize client")
	}

	return client, nil
}

// Deprecated: Use (c *Config).CredentialsAPIClient() instead
func CredentialsAPIClient(token *oauth2.Token) (*credentials.APIClient, error) {

	var client *credentials.APIClient

	clientcfg := credentials.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	client = credentials.NewAPIClient(clientcfg)

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise PingOne Credentials client")
	}

	return client, nil

}

func (c *Config) CredentialsAPIClient(ctx context.Context) (*credentials.APIClient, error) {

	if !c.accessTokenObject.Valid() {
		err := c.getToken(ctx)
		if err != nil {
			return nil, err
		}
	}

	var client *credentials.APIClient

	clientcfg := credentials.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", c.accessTokenObject.AccessToken))

	if c.APIHostnameOverride != nil && *c.APIHostnameOverride != "" {
		clientcfg.SetDefaultServerIndex(1)
		err := clientcfg.SetDefaultServerVariableDefaultValue("baseHostname", *c.APIHostnameOverride)
		if err != nil {
			return nil, err
		}
	} else {
		clientcfg.SetDefaultServerIndex(0)
		err := clientcfg.SetDefaultServerVariableDefaultValue("suffix", model.FindRegionByName(c.Region).URLSuffix)
		if err != nil {
			return nil, err
		}
	}
	client = credentials.NewAPIClient(clientcfg)

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise PingOne Credentials client")
	}

	return client, nil
}

// Deprecated: Use (c *Config).ManagementAPIClient() instead
func ManagementAPIClient(token *oauth2.Token) (*management.APIClient, error) {

	var client *management.APIClient

	clientcfg := management.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	client = management.NewAPIClient(clientcfg)

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise PingOne Management client")
	}

	return client, nil

}

func (c *Config) ManagementAPIClient(ctx context.Context) (*management.APIClient, error) {

	if !c.accessTokenObject.Valid() {
		err := c.getToken(ctx)
		if err != nil {
			return nil, err
		}
	}

	var client *management.APIClient

	clientcfg := management.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", c.accessTokenObject.AccessToken))

	if c.APIHostnameOverride != nil && *c.APIHostnameOverride != "" {
		clientcfg.SetDefaultServerIndex(1)
		err := clientcfg.SetDefaultServerVariableDefaultValue("baseHostname", *c.APIHostnameOverride)
		if err != nil {
			return nil, err
		}
	} else {
		clientcfg.SetDefaultServerIndex(0)
		err := clientcfg.SetDefaultServerVariableDefaultValue("suffix", model.FindRegionByName(c.Region).URLSuffix)
		if err != nil {
			return nil, err
		}
	}
	client = management.NewAPIClient(clientcfg)

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise PingOne Management client")
	}

	return client, nil
}

// Deprecated: Use (c *Config).MFAAPIClient() instead
func MFAAPIClient(token *oauth2.Token) (*mfa.APIClient, error) {

	var client *mfa.APIClient

	clientcfg := mfa.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	client = mfa.NewAPIClient(clientcfg)

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise PingOne MFA client")
	}

	return client, nil

}

func (c *Config) MFAAPIClient(ctx context.Context) (*mfa.APIClient, error) {

	if !c.accessTokenObject.Valid() {
		err := c.getToken(ctx)
		if err != nil {
			return nil, err
		}
	}

	var client *mfa.APIClient

	clientcfg := mfa.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", c.accessTokenObject.AccessToken))

	if c.APIHostnameOverride != nil && *c.APIHostnameOverride != "" {
		clientcfg.SetDefaultServerIndex(1)
		err := clientcfg.SetDefaultServerVariableDefaultValue("baseHostname", *c.APIHostnameOverride)
		if err != nil {
			return nil, err
		}
	} else {
		clientcfg.SetDefaultServerIndex(0)
		err := clientcfg.SetDefaultServerVariableDefaultValue("suffix", model.FindRegionByName(c.Region).URLSuffix)
		if err != nil {
			return nil, err
		}
	}
	client = mfa.NewAPIClient(clientcfg)

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise PingOne MFA client")
	}

	return client, nil
}

// Deprecated: Use (c *Config).RiskAPIClient() instead
func RiskAPIClient(token *oauth2.Token) (*risk.APIClient, error) {

	var client *risk.APIClient

	clientcfg := risk.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	client = risk.NewAPIClient(clientcfg)

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise PingOne Risk client")
	}

	return client, nil

}

func (c *Config) RiskAPIClient(ctx context.Context) (*risk.APIClient, error) {

	if !c.accessTokenObject.Valid() {
		err := c.getToken(ctx)
		if err != nil {
			return nil, err
		}
	}

	var client *risk.APIClient

	clientcfg := risk.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", c.accessTokenObject.AccessToken))

	if c.APIHostnameOverride != nil && *c.APIHostnameOverride != "" {
		clientcfg.SetDefaultServerIndex(1)
		err := clientcfg.SetDefaultServerVariableDefaultValue("baseHostname", *c.APIHostnameOverride)
		if err != nil {
			return nil, err
		}
	} else {
		clientcfg.SetDefaultServerIndex(0)
		err := clientcfg.SetDefaultServerVariableDefaultValue("suffix", model.FindRegionByName(c.Region).URLSuffix)
		if err != nil {
			return nil, err
		}
	}
	client = risk.NewAPIClient(clientcfg)

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise PingOne Risk client")
	}

	return client, nil
}

// Deprecated: Use (c *Config).VerifyAPIClient() instead
func VerifyAPIClient(token *oauth2.Token) (*verify.APIClient, error) {

	var client *verify.APIClient

	clientcfg := verify.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	client = verify.NewAPIClient(clientcfg)

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise PingOne Verify client")
	}

	return client, nil

}

func (c *Config) VerifyAPIClient(ctx context.Context) (*verify.APIClient, error) {

	if !c.accessTokenObject.Valid() {
		err := c.getToken(ctx)
		if err != nil {
			return nil, err
		}
	}

	var client *verify.APIClient

	clientcfg := verify.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", c.accessTokenObject.AccessToken))

	if c.APIHostnameOverride != nil && *c.APIHostnameOverride != "" {
		clientcfg.SetDefaultServerIndex(1)
		err := clientcfg.SetDefaultServerVariableDefaultValue("baseHostname", *c.APIHostnameOverride)
		if err != nil {
			return nil, err
		}
	} else {
		clientcfg.SetDefaultServerIndex(0)
		err := clientcfg.SetDefaultServerVariableDefaultValue("suffix", model.FindRegionByName(c.Region).URLSuffix)
		if err != nil {
			return nil, err
		}
	}
	client = verify.NewAPIClient(clientcfg)

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise PingOne Verify client")
	}

	return client, nil
}

func (c *Config) getToken(ctx context.Context) error {

	if c.AccessToken == "" {

		if c.ClientID == "" || c.ClientSecret == "" || c.EnvironmentID == "" || c.Region == "" {
			return fmt.Errorf("Required parameter missing.  Must provide ClientID, ClientSecret, EnvironmentID and Region.")
		}

		regionSuffix := model.FindRegionByName(c.Region).URLSuffix

		//Get URL from SDK
		authURL := fmt.Sprintf("https://auth.pingone.%s", regionSuffix)
		if c.AuthHostnameOverride != nil && *c.AuthHostnameOverride != "" {
			authURL = fmt.Sprintf("https://%s", *c.AuthHostnameOverride)
		}

		//OAuth 2.0 config for client creds
		config := clientcredentials.Config{
			ClientID:     c.ClientID,
			ClientSecret: c.ClientSecret,
			TokenURL:     fmt.Sprintf("%s/%s/as/token", authURL, c.EnvironmentID),
			AuthStyle:    oauth2.AuthStyleAutoDetect,
		}

		token, err := processResponse(func() (*oauth2.Token, error) {
			return config.Token(ctx)
		})
		if err != nil {
			return err
		}

		c.accessTokenObject = token

		return nil

	} else {
		c.accessTokenObject = &oauth2.Token{
			AccessToken: c.AccessToken,
			TokenType:   "Bearer",
		}
		return nil
	}
}

var (
	maxRetries = 5
)

func processResponse(f func() (*oauth2.Token, error)) (*oauth2.Token, error) {
	return exponentialBackOffRetry(f)
}

func exponentialBackOffRetry(f func() (*oauth2.Token, error)) (*oauth2.Token, error) {
	var token *oauth2.Token
	var err error
	backOffTime := time.Second
	var isRetryable bool

	for i := 0; i < maxRetries; i++ {
		token, err = f()

		if err != nil {
			backOffTime, isRetryable = testForRetryable(err, backOffTime)

			if isRetryable {
				log.Printf("Attempt %d failed: %v, backing off by %s.", i+1, err, backOffTime.String())
				time.Sleep(backOffTime)
				continue
			}
		}

		return token, err
	}

	log.Printf("Request failed after %d attempts", maxRetries)

	return token, err // output the final error
}

func testForRetryable(err error, currentBackoff time.Duration) (time.Duration, bool) {

	backoff := currentBackoff * 2

	retryAbleCodes := []int{
		429,
		500,
		501,
		502,
		503,
		504,
	}

	for _, v := range retryAbleCodes {
		if m, mErr := regexp.MatchString(fmt.Sprintf("%d ", v), err.Error()); mErr == nil && m {
			log.Printf("HTTP status code %d detected, available for retry", v)
			return backoff, true
		}
	}

	return backoff, false
}
