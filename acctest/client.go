package acctest

import (
	"context"
	"fmt"
	"os"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/patrickcping/pingone-go-sdk-v2/pingone"
)

func TestClient(ctx context.Context) (*pingone.Client, error) {

	clientID := os.Getenv("PINGONE_CLIENT_ID_TESTACC")
	clientSecret := os.Getenv("PINGONE_CLIENT_SECRET_TESTACC")
	environmentID := os.Getenv("PINGONE_ENVIRONMENT_ID_TESTACC")
	regionCode := management.EnumRegionCode(os.Getenv("PINGONE_REGION_CODE_TESTACC"))

	if err := os.Setenv("PINGONE_API_ACCESS_TOKEN", ""); err != nil {
		return nil, err
	}

	if err := os.Setenv("PINGONE_API_SERVICE_HOSTNAME", ""); err != nil {
		return nil, err
	}

	if err := os.Setenv("PINGONE_AUTH_SERVICE_HOSTNAME", ""); err != nil {
		return nil, err
	}

	config := &pingone.Config{
		ClientID:      &clientID,
		ClientSecret:  &clientSecret,
		EnvironmentID: &environmentID,
		RegionCode:    &regionCode,
	}

	client, err := config.APIClient(ctx)

	if err != nil {
		return nil, err
	}

	if client.AuthorizeAPIClient == nil {
		return nil, fmt.Errorf("Authorize Client not successfully retrieved, authorize is null")
	}

	if client.CredentialsAPIClient == nil {
		return nil, fmt.Errorf("Credentials Client not successfully retrieved, credentials is null")
	}

	if client.ManagementAPIClient == nil {
		return nil, fmt.Errorf("Management Client not successfully retrieved, management is null")
	}

	if client.MFAAPIClient == nil {
		return nil, fmt.Errorf("Management Client not successfully retrieved, mfa is null")
	}

	if client.RiskAPIClient == nil {
		return nil, fmt.Errorf("Management Client not successfully retrieved, risk is null")
	}

	if client.VerifyAPIClient == nil {
		return nil, fmt.Errorf("Management Client not successfully retrieved, verify is null")
	}

	return client, nil

}
