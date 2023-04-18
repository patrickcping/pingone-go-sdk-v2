package pingone

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement"
	"github.com/patrickcping/pingone-go-sdk-v2/authorize"
	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/patrickcping/pingone-go-sdk-v2/mfa"
	"github.com/patrickcping/pingone-go-sdk-v2/pingone/model"
	"github.com/patrickcping/pingone-go-sdk-v2/risk"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func (c *Config) APIClient(ctx context.Context) (*Client, error) {

	token, err := getToken(ctx, c)
	if err != nil {
		return nil, err
	}

	agreementManagementClient, err := AgreementManagementAPIClient(token)
	if err != nil {
		return nil, err
	}

	authorizeClient, err := AuthorizeAPIClient(token)
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
		AgreementManagementAPIClient: agreementManagementClient,
		AuthorizeAPIClient:           authorizeClient,
		ManagementAPIClient:          managementClient,
		MFAAPIClient:                 mfaClient,
		RiskAPIClient:                riskClient,
		Region:                       model.FindRegionByName(c.Region),
	}

	log.Printf("[INFO] PingOne Client configured")
	return apiClient, nil
}

func AgreementManagementAPIClient(token *oauth2.Token) (*agreementmanagement.APIClient, error) {

	var client *agreementmanagement.APIClient

	clientcfg := agreementmanagement.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	client = agreementmanagement.NewAPIClient(clientcfg)

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise PingOne Agreement Management client")
	}

	log.Printf("[INFO] PingOne Agreement Management Client initialised")

	return client, nil

}

func AuthorizeAPIClient(token *oauth2.Token) (*authorize.APIClient, error) {

	var client *authorize.APIClient

	clientcfg := authorize.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	client = authorize.NewAPIClient(clientcfg)

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise PingOne Authorize client")
	}

	log.Printf("[INFO] PingOne Authorize Client initialised")

	return client, nil

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

	if c.AccessToken == "" {

		if c.ClientID == "" || c.ClientSecret == "" || c.EnvironmentID == "" || c.Region == "" {
			return nil, fmt.Errorf("Required parameter missing.  Must provide ClientID, ClientSecret, EnvironmentID and Region.")
		}

		regionSuffix := model.FindRegionByName(c.Region).URLSuffix

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

		token, err := processResponse(func() (*oauth2.Token, error) {
			return config.Token(ctx)
		})
		if err != nil {
			return nil, err
		}
		log.Printf("[INFO] Token retrieved")

		return token, nil

	} else {
		token := oauth2.Token{
			AccessToken: c.AccessToken,
			TokenType:   "Bearer",
		}
		return &token, nil
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
