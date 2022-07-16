package pingone

import (
	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/patrickcping/pingone-go-sdk-v2/mfa"
	"github.com/patrickcping/pingone-go-sdk-v2/risk"
)

type Config struct {
	ClientID      string
	ClientSecret  string
	EnvironmentID string
	Region        string
}

type Client struct {
	ManagementAPIClient *management.APIClient
	MFAAPIClient        *mfa.APIClient
	RiskAPIClient       *risk.APIClient
	RegionSuffix        string
}
