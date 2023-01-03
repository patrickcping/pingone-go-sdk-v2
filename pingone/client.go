package pingone

import (
	"context"
	"fmt"
	"log"

	"github.com/patrickcping/pingone-go-sdk-v2/authorize"
	"github.com/patrickcping/pingone-go-sdk-v2/internal/fsconfig"
	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/patrickcping/pingone-go-sdk-v2/mfa"
	"github.com/patrickcping/pingone-go-sdk-v2/pingone/model"
	"github.com/patrickcping/pingone-go-sdk-v2/risk"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func (c *Config) APIClient(ctx context.Context) (*Client, error) {

	var token *oauth2.Token
	var err error

	if c.AccessToken != "" {

		if c.Region == "" {
			return nil, fmt.Errorf("Region is required to be set when using the Access Token credential configuration method")
		}

		if c.Profile != "" || c.ClientID != "" || c.ClientSecret != "" || c.EnvironmentID != "" {
			log.Println("An access token has been set in credential configuration along with a profile or worker client details (client ID, client secret, environment ID).  The access token takes priority.  All other credential configuration parameters are ignored.")
		}

		token = &oauth2.Token{
			AccessToken: c.AccessToken,
			TokenType:   "Bearer",
		}

	} else {

		if c.UseDefaultProfile || c.Profile != "" {

			if c.ClientID != "" || c.ClientSecret != "" || c.EnvironmentID != "" || c.Region != "" {
				log.Println("A profile has been set in credential configuration along with worker client details (client ID, client secret, environment ID, region).  The configured profile takes priority, but behaviour may be unpredictable.")
			}

			profile, err := fetchProfile(ctx, c)
			if err != nil {
				return nil, err
			}

			c.ClientID = profile.ClientID
			c.ClientSecret = profile.ClientSecret
			c.EnvironmentID = profile.EnvironmentID
			c.Region = profile.Region
			c.AuthEndpoint = profile.AuthEndpoint
			c.ApiEndpoint = profile.ApiEndpoint
		}

		token, err = getToken(ctx, c)
		if err != nil {
			return nil, err
		}

	}

	authorizeClient, err := APIClientFactory(token, authorize.NewConfiguration())
	if err != nil {
		return nil, err
	}

	managementClient, err := APIClientFactory(token, management.NewConfiguration())
	if err != nil {
		return nil, err
	}

	mfaClient, err := APIClientFactory(token, mfa.NewConfiguration())
	if err != nil {
		return nil, err
	}

	riskClient, err := APIClientFactory(token, risk.NewConfiguration())
	if err != nil {
		return nil, err
	}

	apiClient := &Client{
		AuthorizeAPIClient:  authorizeClient.(*authorize.APIClient),
		ManagementAPIClient: managementClient.(*management.APIClient),
		MFAAPIClient:        mfaClient.(*mfa.APIClient),
		RiskAPIClient:       riskClient.(*risk.APIClient),
	}

	if c.Region != "" {
		apiClient.Region = model.FindRegionByName(c.Region)
	}

	return apiClient, nil
}

func APIClientFactory(token *oauth2.Token, c interface{}) (interface{}, error) {

	var client interface{}

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token.AccessToken),
	}

	switch v := c.(type) {
	case *authorize.Configuration:
		clientcfg := v
		for k, v := range headers {
			clientcfg.AddDefaultHeader(k, v)
		}
		client = authorize.NewAPIClient(clientcfg)
	case *management.Configuration:
		clientcfg := v
		for k, v := range headers {
			clientcfg.AddDefaultHeader(k, v)
		}
		client = management.NewAPIClient(clientcfg)
	case *mfa.Configuration:
		clientcfg := v
		for k, v := range headers {
			clientcfg.AddDefaultHeader(k, v)
		}
		client = mfa.NewAPIClient(clientcfg)
	case *risk.Configuration:
		clientcfg := v
		for k, v := range headers {
			clientcfg.AddDefaultHeader(k, v)
		}
		client = risk.NewAPIClient(clientcfg)
	}

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise client")
	}

	return client, nil

}

func AuthorizeAPIClient(token *oauth2.Token) (*authorize.APIClient, error) {

	client, err := APIClientFactory(token, authorize.NewConfiguration())
	if err != nil {
		return nil, fmt.Errorf("Cannot initialise PingOne Authorize client")
	}

	return client.(*authorize.APIClient), nil

}

func ManagementAPIClient(token *oauth2.Token) (*management.APIClient, error) {

	client, err := APIClientFactory(token, management.NewConfiguration())
	if err != nil {
		return nil, fmt.Errorf("Cannot initialise PingOne Management client")
	}

	return client.(*management.APIClient), nil

}

func MFAAPIClient(token *oauth2.Token) (*mfa.APIClient, error) {

	client, err := APIClientFactory(token, mfa.NewConfiguration())
	if err != nil {
		return nil, fmt.Errorf("Cannot initialise PingOne MFA client")
	}

	return client.(*mfa.APIClient), nil

}

func RiskAPIClient(token *oauth2.Token) (*risk.APIClient, error) {

	client, err := APIClientFactory(token, risk.NewConfiguration())
	if err != nil {
		return nil, fmt.Errorf("Cannot initialise PingOne Risk client")
	}

	return client.(*risk.APIClient), nil

}

func fetchProfile(ctx context.Context, c *Config) (*ProfileConfig, error) {

	filename := "$HOME/.pingidentity/config"

	if c.ProfileCredentialsFilePath != "" {
		filename = c.ProfileCredentialsFilePath
	}

	filedata, err := fsconfig.ReadFile(filename)

	profilePrefix := "PINGONE"
	if c.Profile != "" {
		profilePrefix += fmt.Sprintf("_%s", c.Profile)
	}
	config := ProfileConfig{
		ClientID:      filedata[fmt.Sprintf("%s_WORKER_APP_CLIENT_ID", profilePrefix)],
		ClientSecret:  filedata[fmt.Sprintf("%s_WORKER_APP_CLIENT_SECRET", profilePrefix)],
		EnvironmentID: filedata[fmt.Sprintf("%s_ENVIRONMENT_ID", profilePrefix)],
		Region:        filedata[fmt.Sprintf("%s_REGION", profilePrefix)],
		AuthEndpoint:  filedata[fmt.Sprintf("%s_AUTH_URL", profilePrefix)],
		ApiEndpoint:   filedata[fmt.Sprintf("%s_API_URL", profilePrefix)],
	}

	if err != nil {
		return nil, err
	}

	return &config, nil

}

func getToken(ctx context.Context, c *Config) (*oauth2.Token, error) {

	var config clientcredentials.Config

	if c.ClientID != "" && c.ClientSecret != "" && c.EnvironmentID != "" && (c.Region != "" || c.AuthEndpoint != "") {

		var authURL string

		if c.AuthEndpoint != "" {
			authURL = c.AuthEndpoint
		} else {

			if c.Region != "" {
				regionSuffix := model.FindRegionByName(c.Region).URLSuffix

				//Get URL from SDK
				authURL = fmt.Sprintf("https://auth.pingone.%s", regionSuffix)
			} else {
				return nil, fmt.Errorf("Required parameter missing.  Must provide Region or AuthEndpoint.")
			}
		}

		//OAuth 2.0 config for client creds
		config = clientcredentials.Config{
			ClientID:     c.ClientID,
			ClientSecret: c.ClientSecret,
			TokenURL:     fmt.Sprintf("%s/%s/as/token", authURL, c.EnvironmentID),
			AuthStyle:    oauth2.AuthStyleAutoDetect,
		}

	} else {
		return nil, fmt.Errorf("Required parameter missing.  Must provide ClientID, ClientSecret, EnvironmentID and Region/AuthEndpoint.")
	}

	token, err := config.Token(ctx)
	if err != nil {
		return nil, err
	}

	return token, nil

}
