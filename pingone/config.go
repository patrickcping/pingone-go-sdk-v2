package pingone

import (
	"github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement"
	"github.com/patrickcping/pingone-go-sdk-v2/authorize"
	"github.com/patrickcping/pingone-go-sdk-v2/credentials"
	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/patrickcping/pingone-go-sdk-v2/mfa"
	"github.com/patrickcping/pingone-go-sdk-v2/pingone/model"
	"github.com/patrickcping/pingone-go-sdk-v2/risk"
	"github.com/patrickcping/pingone-go-sdk-v2/verify"
)

type Config struct {
	ClientID      string
	ClientSecret  string
	EnvironmentID string
	AccessToken   string
	Region        string
}

type Client struct {
	AgreementManagementAPIClient *agreementmanagement.APIClient
	AuthorizeAPIClient           *authorize.APIClient
	CredentialsAPIClient         *credentials.APIClient
	ManagementAPIClient          *management.APIClient
	MFAAPIClient                 *mfa.APIClient
	RiskAPIClient                *risk.APIClient
	VerifyAPIClient              *verify.APIClient
	Region                       model.RegionMapping
}
