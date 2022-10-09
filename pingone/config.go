package pingone

import (
	"github.com/patrickcping/pingone-go-sdk-v2/authorize"
	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/patrickcping/pingone-go-sdk-v2/mfa"
	"github.com/patrickcping/pingone-go-sdk-v2/pingone/model"
	"github.com/patrickcping/pingone-go-sdk-v2/risk"
)

type Config struct {
	ClientID      string
	ClientSecret  string
	EnvironmentID string
	Region        string
}

type Client struct {
	AuthorizeAPIClient  *authorize.APIClient
	ManagementAPIClient *management.APIClient
	MFAAPIClient        *mfa.APIClient
	RiskAPIClient       *risk.APIClient
	Region              model.RegionMapping
}
