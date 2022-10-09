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

	config := &Config{
		ClientID:      os.Getenv("PINGONE_CLIENT_ID"),
		ClientSecret:  os.Getenv("PINGONE_CLIENT_SECRET"),
		EnvironmentID: os.Getenv("PINGONE_ENVIRONMENT_ID"),
		Region:        os.Getenv("PINGONE_REGION"),
	}

	client, err := config.APIClient(ctx)

	if err != nil {
		t.Fatalf("Client not successfully retrieved")
	}

	if client.AuthorizeAPIClient == nil {
		t.Fatalf("Authorize Client not successfully retrieved")
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

	if client.Region != model.FindRegionByName(os.Getenv("PINGONE_REGION")) {
		t.Fatalf("Unexpected region.  Expected %s, got %v", os.Getenv("PINGONE_REGION"), client.Region)
	}
}

func TestAccAPIClient_MissingClientID(t *testing.T) {
	t.Parallel()

	var ctx = context.Background()

	config := &Config{
		ClientID:      "",
		ClientSecret:  os.Getenv("PINGONE_CLIENT_SECRET"),
		EnvironmentID: os.Getenv("PINGONE_ENVIRONMENT_ID"),
		Region:        os.Getenv("PINGONE_REGION"),
	}

	client, err := config.APIClient(ctx)

	if err == nil || client != nil {
		t.Fatalf("Client not expected to be successfully retrieved")
	}
}

func TestAccAPIClient_MissingClientSecret(t *testing.T) {
	t.Parallel()

	var ctx = context.Background()

	config := &Config{
		ClientID:      os.Getenv("PINGONE_CLIENT_ID"),
		ClientSecret:  "",
		EnvironmentID: os.Getenv("PINGONE_ENVIRONMENT_ID"),
		Region:        os.Getenv("PINGONE_REGION"),
	}

	client, err := config.APIClient(ctx)

	if err == nil || client != nil {
		t.Fatalf("Client not expected to be successfully retrieved")
	}
}

func TestAccAPIClient_MissingClientEnvironment(t *testing.T) {
	t.Parallel()

	var ctx = context.Background()

	config := &Config{
		ClientID:      os.Getenv("PINGONE_CLIENT_ID"),
		ClientSecret:  os.Getenv("PINGONE_CLIENT_SECRET"),
		EnvironmentID: "",
		Region:        os.Getenv("PINGONE_REGION"),
	}

	client, err := config.APIClient(ctx)

	if err == nil || client != nil {
		t.Fatalf("Client not expected to be successfully retrieved")
	}
}

func TestAccAPIClient_MissingClientRegion(t *testing.T) {
	t.Parallel()

	var ctx = context.Background()

	config := &Config{
		ClientID:      os.Getenv("PINGONE_CLIENT_ID"),
		ClientSecret:  os.Getenv("PINGONE_CLIENT_SECRET"),
		EnvironmentID: os.Getenv("PINGONE_ENVIRONMENT_ID"),
		Region:        "",
	}

	client, err := config.APIClient(ctx)

	if err == nil || client != nil {
		t.Fatalf("Client not expected to be successfully retrieved")
	}
}

func TestAccAPIClient_FailedAuth(t *testing.T) {
	t.Parallel()

	var ctx = context.Background()

	config := &Config{
		ClientID:      os.Getenv("PINGONE_CLIENT_ID"),
		ClientSecret:  os.Getenv("PINGONE_CLIENT_ID"),
		EnvironmentID: os.Getenv("PINGONE_ENVIRONMENT_ID"),
		Region:        os.Getenv("PINGONE_REGION"),
	}

	client, err := config.APIClient(ctx)

	if err == nil || client != nil {
		t.Fatalf("Client not expected to be successfully retrieved")
	}
}

func TestAccAPIClient_BadRegion(t *testing.T) {
	t.Parallel()

	var ctx = context.Background()

	config := &Config{
		ClientID:      os.Getenv("PINGONE_CLIENT_ID"),
		ClientSecret:  os.Getenv("PINGONE_CLIENT_ID"),
		EnvironmentID: os.Getenv("PINGONE_ENVIRONMENT_ID"),
		Region:        "NZ",
	}

	client, err := config.APIClient(ctx)

	if err == nil || client != nil {
		t.Fatalf("Client not expected to be successfully retrieved")
	}
}
