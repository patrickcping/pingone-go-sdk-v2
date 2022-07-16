package pingone

import (
	"context"
	"fmt"
	"log"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/patrickcping/pingone-go-sdk-v2/mfa"
	"github.com/patrickcping/pingone-go-sdk-v2/risk"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func (c *Config) APIClient(ctx context.Context) (*Client, error) {

	token, err := getToken(ctx, c)
	if err != nil {
		return nil, err
	}

	managementClient, err := ManagementAPIClient(token)
	if err != nil {
		return nil, err
	}

	mfaClient, err := MFAAPIClient(token)
	if err != nil {
		return nil, err
	}

	riskClient, err := RiskAPIClient(token)
	if err != nil {
		return nil, err
	}

	apiClient := &Client{
		ManagementAPIClient: managementClient,
		MFAAPIClient:        mfaClient,
		RiskAPIClient:       riskClient,
		Region:              FindRegionByName(c.Region),
	}

	log.Printf("[INFO] PingOne Client configured")
	return apiClient, nil
}

func ManagementAPIClient(token *oauth2.Token) (*management.APIClient, error) {

	var client *management.APIClient

	clientcfg := management.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	client = management.NewAPIClient(clientcfg)

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise PingOne Management client")
	}

	log.Printf("[INFO] PingOne Management Client initialised")

	return client, nil

}

func MFAAPIClient(token *oauth2.Token) (*mfa.APIClient, error) {

	var client *mfa.APIClient

	clientcfg := mfa.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	client = mfa.NewAPIClient(clientcfg)

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise PingOne MFA client")
	}

	log.Printf("[INFO] PingOne MFA Client initialised")

	return client, nil

}

func RiskAPIClient(token *oauth2.Token) (*risk.APIClient, error) {

	var client *risk.APIClient

	clientcfg := risk.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	client = risk.NewAPIClient(clientcfg)

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise PingOne Risk client")
	}

	log.Printf("[INFO] PingOne Risk Client initialised")

	return client, nil

}

func getToken(ctx context.Context, c *Config) (*oauth2.Token, error) {

	regionSuffix := FindRegionByName(c.Region).URLSuffix

	//Get URL from SDK
	authURL := fmt.Sprintf("https://auth.pingone.%s", regionSuffix)
	log.Printf("[INFO] Getting token from %s", authURL)

	//OAuth 2.0 config for client creds
	config := clientcredentials.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		TokenURL:     fmt.Sprintf("%s/%s/as/token", authURL, c.EnvironmentID),
		AuthStyle:    oauth2.AuthStyleAutoDetect,
	}

	token, err := config.Token(ctx)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] Token retrieved")

	return token, nil
}
