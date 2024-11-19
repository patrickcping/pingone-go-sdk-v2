package pingone

import (
	"context"
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/authorize"
	"github.com/patrickcping/pingone-go-sdk-v2/credentials"
	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/patrickcping/pingone-go-sdk-v2/mfa"
	"github.com/patrickcping/pingone-go-sdk-v2/pingone/model"
	"github.com/patrickcping/pingone-go-sdk-v2/risk"
	"github.com/patrickcping/pingone-go-sdk-v2/verify"
	"golang.org/x/oauth2"
)

type Client struct {
	AuthorizeAPIClient   *authorize.APIClient
	CredentialsAPIClient *credentials.APIClient
	ManagementAPIClient  *management.APIClient
	MFAAPIClient         *mfa.APIClient
	RiskAPIClient        *risk.APIClient
	VerifyAPIClient      *verify.APIClient
	Region               model.RegionMapping
}

var version = "0.12.4"

func (c *Config) APIClient(ctx context.Context) (*Client, error) {

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

	var region model.RegionMapping
	if c.Region != "" {
		region = model.FindRegionByName(c.Region) //nolint:staticcheck
	}

	if v := c.RegionCode; v != nil {
		region = model.FindRegionByAPICode(*v)
	}

	apiClient := &Client{
		AuthorizeAPIClient:   authorizeClient,
		CredentialsAPIClient: credentialsClient,
		ManagementAPIClient:  managementClient,
		MFAAPIClient:         mfaClient,
		RiskAPIClient:        riskClient,
		VerifyAPIClient:      verifyClient,
		Region:               region,
	}

	return apiClient, nil
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

	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("Client validation error: %s", err)
	}

	if !c.accessTokenObject.Valid() {
		err := c.getToken(ctx)
		if err != nil {
			return nil, err
		}
	}

	var client *authorize.APIClient

	clientcfg := authorize.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", c.accessTokenObject.AccessToken))

	if checkForValue(c.APIHostnameOverride) {
		clientcfg.SetDefaultServerIndex(1)
		err := clientcfg.SetDefaultServerVariableDefaultValue("baseHostname", *c.APIHostnameOverride)
		if err != nil {
			return nil, err
		}
	} else {
		clientcfg.SetDefaultServerIndex(0)

		var region model.RegionMapping
		if c.Region != "" {
			region = model.FindRegionByName(c.Region) //nolint:staticcheck
		}

		if v := c.RegionCode; v != nil {
			region = model.FindRegionByAPICode(*v)
		}

		err := clientcfg.SetDefaultServerVariableDefaultValue("suffix", region.URLSuffix)
		if err != nil {
			return nil, err
		}
	}

	if checkForValue(c.UserAgentOverride) {
		clientcfg.SetUserAgent(*c.UserAgentOverride)
	} else {
		clientcfg.AppendUserAgent(fmt.Sprintf("PingOne-GOLANG-SDK/%s", version))
	}

	if checkForValue(c.UserAgentSuffix) {
		clientcfg.AppendUserAgent(*c.UserAgentSuffix)
	}

	if checkForValue(c.ProxyURL) {
		clientcfg.ProxyURL = c.ProxyURL
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

	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("Client validation error: %s", err)
	}

	if !c.accessTokenObject.Valid() {
		err := c.getToken(ctx)
		if err != nil {
			return nil, err
		}
	}

	var client *credentials.APIClient

	clientcfg := credentials.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", c.accessTokenObject.AccessToken))

	if checkForValue(c.APIHostnameOverride) {
		clientcfg.SetDefaultServerIndex(1)
		err := clientcfg.SetDefaultServerVariableDefaultValue("baseHostname", *c.APIHostnameOverride)
		if err != nil {
			return nil, err
		}
	} else {
		clientcfg.SetDefaultServerIndex(0)

		var region model.RegionMapping
		if c.Region != "" {
			region = model.FindRegionByName(c.Region) //nolint:staticcheck
		}

		if v := c.RegionCode; v != nil {
			region = model.FindRegionByAPICode(*v)
		}

		err := clientcfg.SetDefaultServerVariableDefaultValue("suffix", region.URLSuffix)
		if err != nil {
			return nil, err
		}
	}

	if checkForValue(c.UserAgentOverride) {
		clientcfg.SetUserAgent(*c.UserAgentOverride)
	} else {
		clientcfg.AppendUserAgent(fmt.Sprintf("PingOne-GOLANG-SDK/%s", version))
	}

	if checkForValue(c.UserAgentSuffix) {
		clientcfg.AppendUserAgent(*c.UserAgentSuffix)
	}

	if checkForValue(c.ProxyURL) {
		clientcfg.ProxyURL = c.ProxyURL
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

	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("Client validation error: %s", err)
	}

	if !c.accessTokenObject.Valid() {
		err := c.getToken(ctx)
		if err != nil {
			return nil, err
		}
	}

	var client *management.APIClient

	clientcfg := management.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", c.accessTokenObject.AccessToken))

	if checkForValue(c.APIHostnameOverride) {
		clientcfg.SetDefaultServerIndex(1)
		err := clientcfg.SetDefaultServerVariableDefaultValue("baseHostname", *c.APIHostnameOverride)
		if err != nil {
			return nil, err
		}
	} else {
		clientcfg.SetDefaultServerIndex(0)

		var region model.RegionMapping
		if c.Region != "" {
			region = model.FindRegionByName(c.Region) //nolint:staticcheck
		}

		if v := c.RegionCode; v != nil {
			region = model.FindRegionByAPICode(*v)
		}

		err := clientcfg.SetDefaultServerVariableDefaultValue("suffix", region.URLSuffix)
		if err != nil {
			return nil, err
		}
	}

	if checkForValue(c.UserAgentOverride) {
		clientcfg.SetUserAgent(*c.UserAgentOverride)
	} else {
		clientcfg.AppendUserAgent(fmt.Sprintf("PingOne-GOLANG-SDK/%s", version))
	}

	if checkForValue(c.UserAgentSuffix) {
		clientcfg.AppendUserAgent(*c.UserAgentSuffix)
	}

	if checkForValue(c.ProxyURL) {
		clientcfg.ProxyURL = c.ProxyURL
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

	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("Client validation error: %s", err)
	}

	if !c.accessTokenObject.Valid() {
		err := c.getToken(ctx)
		if err != nil {
			return nil, err
		}
	}

	var client *mfa.APIClient

	clientcfg := mfa.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", c.accessTokenObject.AccessToken))

	if checkForValue(c.APIHostnameOverride) {
		clientcfg.SetDefaultServerIndex(1)
		err := clientcfg.SetDefaultServerVariableDefaultValue("baseHostname", *c.APIHostnameOverride)
		if err != nil {
			return nil, err
		}
	} else {
		clientcfg.SetDefaultServerIndex(0)

		var region model.RegionMapping
		if c.Region != "" {
			region = model.FindRegionByName(c.Region) //nolint:staticcheck
		}

		if v := c.RegionCode; v != nil {
			region = model.FindRegionByAPICode(*v)
		}

		err := clientcfg.SetDefaultServerVariableDefaultValue("suffix", region.URLSuffix)
		if err != nil {
			return nil, err
		}
	}

	if checkForValue(c.UserAgentOverride) {
		clientcfg.SetUserAgent(*c.UserAgentOverride)
	} else {
		clientcfg.AppendUserAgent(fmt.Sprintf("PingOne-GOLANG-SDK/%s", version))
	}

	if checkForValue(c.UserAgentSuffix) {
		clientcfg.AppendUserAgent(*c.UserAgentSuffix)
	}

	if checkForValue(c.ProxyURL) {
		clientcfg.ProxyURL = c.ProxyURL
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

	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("Client validation error: %s", err)
	}

	if !c.accessTokenObject.Valid() {
		err := c.getToken(ctx)
		if err != nil {
			return nil, err
		}
	}

	var client *risk.APIClient

	clientcfg := risk.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", c.accessTokenObject.AccessToken))

	if checkForValue(c.APIHostnameOverride) {
		clientcfg.SetDefaultServerIndex(1)
		err := clientcfg.SetDefaultServerVariableDefaultValue("baseHostname", *c.APIHostnameOverride)
		if err != nil {
			return nil, err
		}
	} else {
		clientcfg.SetDefaultServerIndex(0)

		var region model.RegionMapping
		if c.Region != "" {
			region = model.FindRegionByName(c.Region) //nolint:staticcheck
		}

		if v := c.RegionCode; v != nil {
			region = model.FindRegionByAPICode(*v)
		}

		err := clientcfg.SetDefaultServerVariableDefaultValue("suffix", region.URLSuffix)
		if err != nil {
			return nil, err
		}
	}

	if checkForValue(c.UserAgentOverride) {
		clientcfg.SetUserAgent(*c.UserAgentOverride)
	} else {
		clientcfg.AppendUserAgent(fmt.Sprintf("PingOne-GOLANG-SDK/%s", version))
	}

	if checkForValue(c.UserAgentSuffix) {
		clientcfg.AppendUserAgent(*c.UserAgentSuffix)
	}

	if checkForValue(c.ProxyURL) {
		clientcfg.ProxyURL = c.ProxyURL
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

	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("Client validation error: %s", err)
	}

	if !c.accessTokenObject.Valid() {
		err := c.getToken(ctx)
		if err != nil {
			return nil, err
		}
	}

	var client *verify.APIClient

	clientcfg := verify.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", c.accessTokenObject.AccessToken))

	if checkForValue(c.APIHostnameOverride) {
		clientcfg.SetDefaultServerIndex(1)
		err := clientcfg.SetDefaultServerVariableDefaultValue("baseHostname", *c.APIHostnameOverride)
		if err != nil {
			return nil, err
		}
	} else {
		clientcfg.SetDefaultServerIndex(0)

		var region model.RegionMapping
		if c.Region != "" {
			region = model.FindRegionByName(c.Region) //nolint:staticcheck
		}

		if v := c.RegionCode; v != nil {
			region = model.FindRegionByAPICode(*v)
		}

		err := clientcfg.SetDefaultServerVariableDefaultValue("suffix", region.URLSuffix)
		if err != nil {
			return nil, err
		}
	}

	if checkForValue(c.UserAgentOverride) {
		clientcfg.SetUserAgent(*c.UserAgentOverride)
	} else {
		clientcfg.AppendUserAgent(fmt.Sprintf("PingOne-GOLANG-SDK/%s", version))
	}

	if checkForValue(c.UserAgentSuffix) {
		clientcfg.AppendUserAgent(*c.UserAgentSuffix)
	}

	if checkForValue(c.ProxyURL) {
		clientcfg.ProxyURL = c.ProxyURL
	}

	client = verify.NewAPIClient(clientcfg)

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise PingOne Verify client")
	}

	return client, nil
}
