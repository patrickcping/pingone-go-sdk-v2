package pingone

import (
	"context"
	"os"
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/pingone/model"
)

func TestAccAPIClient_Success(t *testing.T) {
	t.Parallel()

	var ctx = context.Background()

	clientID := os.Getenv("PINGONE_CLIENT_ID_TESTACC")
	clientSecret := os.Getenv("PINGONE_CLIENT_SECRET_TESTACC")
	environmentID := os.Getenv("PINGONE_ENVIRONMENT_ID_TESTACC")
	region := os.Getenv("PINGONE_REGION_TESTACC")

	if err := os.Setenv("PINGONE_API_ACCESS_TOKEN", ""); err != nil {
		t.Fatalf("Client not successfully retrieved: %s", err)
	}

	if err := os.Setenv("PINGONE_API_SERVICE_HOSTNAME", ""); err != nil {
		t.Fatalf("Client not successfully retrieved: %s", err)
	}

	if err := os.Setenv("PINGONE_AUTH_SERVICE_HOSTNAME", ""); err != nil {
		t.Fatalf("Client not successfully retrieved: %s", err)
	}

	config := &Config{
		ClientID:      &clientID,
		ClientSecret:  &clientSecret,
		EnvironmentID: &environmentID,
		Region:        region,
	}

	client, err := config.APIClient(ctx)

	if err != nil {
		t.Fatalf("Client not successfully retrieved: %s", err)
	}

	if client.AuthorizeAPIClient == nil {
		t.Fatalf("Authorize Client not successfully retrieved")
	}

	if client.CredentialsAPIClient == nil {
		t.Fatalf("Credentials Client not successfully retrieved")
	}

	if client.ManagementAPIClient == nil {
		t.Fatalf("Management Client not successfully retrieved")
	}

	if client.MFAAPIClient == nil {
		t.Fatalf("MFA Client not successfully retrieved")
	}

	if client.RiskAPIClient == nil {
		t.Fatalf("Risk Client not successfully retrieved")
	}

	if client.VerifyAPIClient == nil {
		t.Fatalf("Verify Client not successfully retrieved")
	}

	if client.Region != model.FindRegionByName(region) {
		t.Fatalf("Unexpected region.  Expected %s, got %v", region, client.Region)
	}
}

func TestAccAPIClient_MissingClientID(t *testing.T) {
	t.Parallel()

	var ctx = context.Background()

	clientID := ""
	clientSecret := os.Getenv("PINGONE_CLIENT_SECRET_TESTACC")
	environmentID := os.Getenv("PINGONE_ENVIRONMENT_ID_TESTACC")
	region := os.Getenv("PINGONE_REGION_TESTACC")

	config := &Config{
		ClientID:      &clientID,
		ClientSecret:  &clientSecret,
		EnvironmentID: &environmentID,
		Region:        region,
	}

	client, err := config.APIClient(ctx)

	if err == nil || client != nil {
		t.Fatalf("Client not expected to be successfully retrieved")
	}
}

func TestAccAPIClient_MissingClientSecret(t *testing.T) {
	t.Parallel()

	var ctx = context.Background()

	clientID := os.Getenv("PINGONE_CLIENT_ID_TESTACC")
	clientSecret := ""
	environmentID := os.Getenv("PINGONE_ENVIRONMENT_ID_TESTACC")
	region := os.Getenv("PINGONE_REGION_TESTACC")

	config := &Config{
		ClientID:      &clientID,
		ClientSecret:  &clientSecret,
		EnvironmentID: &environmentID,
		Region:        region,
	}

	client, err := config.APIClient(ctx)

	if err == nil || client != nil {
		t.Fatalf("Client not expected to be successfully retrieved")
	}
}

func TestAccAPIClient_MissingClientEnvironment(t *testing.T) {
	t.Parallel()

	var ctx = context.Background()

	clientID := os.Getenv("PINGONE_CLIENT_ID_TESTACC")
	clientSecret := os.Getenv("PINGONE_CLIENT_SECRET_TESTACC")
	environmentID := ""
	region := os.Getenv("PINGONE_REGION_TESTACC")

	config := &Config{
		ClientID:      &clientID,
		ClientSecret:  &clientSecret,
		EnvironmentID: &environmentID,
		Region:        region,
	}

	client, err := config.APIClient(ctx)

	if err == nil || client != nil {
		t.Fatalf("Client not expected to be successfully retrieved")
	}
}

func TestAccAPIClient_MissingClientRegion(t *testing.T) {
	t.Parallel()

	var ctx = context.Background()

	clientID := os.Getenv("PINGONE_CLIENT_ID_TESTACC")
	clientSecret := os.Getenv("PINGONE_CLIENT_SECRET_TESTACC")
	environmentID := os.Getenv("PINGONE_ENVIRONMENT_ID_TESTACC")
	region := ""

	config := &Config{
		ClientID:      &clientID,
		ClientSecret:  &clientSecret,
		EnvironmentID: &environmentID,
		Region:        region,
	}

	client, err := config.APIClient(ctx)

	if err == nil || client != nil {
		t.Fatalf("Client not expected to be successfully retrieved")
	}
}

func TestAccAPIClient_FailedAuth(t *testing.T) {
	t.Parallel()

	var ctx = context.Background()

	clientID := os.Getenv("PINGONE_CLIENT_ID_TESTACC")
	clientSecret := os.Getenv("PINGONE_CLIENT_ID_TESTACC")
	environmentID := os.Getenv("PINGONE_ENVIRONMENT_ID_TESTACC")
	region := os.Getenv("PINGONE_REGION_TESTACC")

	config := &Config{
		ClientID:      &clientID,
		ClientSecret:  &clientSecret,
		EnvironmentID: &environmentID,
		Region:        region,
	}

	client, err := config.APIClient(ctx)

	if err == nil || client != nil {
		t.Fatalf("Client not expected to be successfully retrieved")
	}
}

func TestAccAPIClient_BadRegion(t *testing.T) {
	t.Parallel()

	var ctx = context.Background()

	clientID := os.Getenv("PINGONE_CLIENT_ID_TESTACC")
	clientSecret := os.Getenv("PINGONE_CLIENT_SECRET_TESTACC")
	environmentID := os.Getenv("PINGONE_ENVIRONMENT_ID_TESTACC")
	region := "NZ"

	config := &Config{
		ClientID:      &clientID,
		ClientSecret:  &clientSecret,
		EnvironmentID: &environmentID,
		Region:        region,
	}

	client, err := config.APIClient(ctx)

	if err == nil || client != nil {
		t.Fatalf("Client not expected to be successfully retrieved")
	}
}
