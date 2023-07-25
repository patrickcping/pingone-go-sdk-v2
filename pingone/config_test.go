package pingone

import (
	"os"
	"strings"
	"testing"
)

func TestValidateAccessToken_ConfigSuccess(t *testing.T) {

	value := "testtest123"

	config := &Config{
		AccessToken: &value,
	}

	if err := config.validateAccessToken(); err != nil {
		t.Fatalf("Parameter not successfully verified: %s", err)
	}

	if *config.AccessToken != value {
		t.Fatalf("Parameter unexpectedly overwritten")
	}
}

func TestValidateAccessToken_EnvSuccess(t *testing.T) {

	value := os.Getenv("PINGONE_API_ACCESS_TOKEN")

	if value == "" {
		t.Fatalf("Required environment variable PINGONE_API_ACCESS_TOKEN not set")
	}

	config := &Config{}

	if err := config.validateAccessToken(); err != nil {
		t.Fatalf("Parameter not successfully verified: %s", err)
	}

	if *config.AccessToken != value {
		t.Fatalf("Parameter default not taken from environment")
	}
}

func TestValidateAccessToken_InvalidFormat(t *testing.T) {
	t.Skip()
}

func TestValidateAgreementMgmtHostnameOverride_ConfigSuccess(t *testing.T) {

	value := "agreement-mgmt.ping-devops.com"

	config := &Config{
		AgreementMgmtHostnameOverride: &value,
	}

	if err := config.validateAgreementMgmtHostnameOverride(); err != nil {
		t.Fatalf("Parameter not successfully verified: %s", err)
	}

	if *config.AgreementMgmtHostnameOverride != value {
		t.Fatalf("Parameter unexpectedly overwritten")
	}
}

func TestValidateAgreementMgmtHostnameOverride_EnvSuccess(t *testing.T) {

	value := os.Getenv("PINGONE_AGREEMENT_MGMT_SERVICE_HOSTNAME")

	if value == "" {
		t.Fatalf("Required environment variable PINGONE_AGREEMENT_MGMT_SERVICE_HOSTNAME not set")
	}

	config := &Config{}

	if err := config.validateAgreementMgmtHostnameOverride(); err != nil {
		t.Fatalf("Parameter not successfully verified: %s", err)
	}

	if *config.AgreementMgmtHostnameOverride != value {
		t.Fatalf("Parameter default not taken from environment")
	}
}

func TestValidateAgreementMgmtHostnameOverride_InvalidFormat(t *testing.T) {
	value := "https://dummyhostname"

	config := &Config{
		AgreementMgmtHostnameOverride: &value,
	}

	if err := config.validateAgreementMgmtHostnameOverride(); err == nil {
		t.Fatalf("Invalid Parameter format not successfully verified")
	}

	if err := config.validateAgreementMgmtHostnameOverride(); !strings.HasPrefix(err.Error(), "Invalid parameter format") {
		t.Fatalf("Invalid Parameter format not successfully verified: %s", err.Error())
	}
}

func TestValidateAPIHostnameOverride_ConfigSuccess(t *testing.T) {

	value := "api.ping-devops.com"

	config := &Config{
		APIHostnameOverride: &value,
	}

	if err := config.validateAPIHostnameOverride(); err != nil {
		t.Fatalf("Parameter not successfully verified: %s", err)
	}

	if *config.APIHostnameOverride != value {
		t.Fatalf("Parameter unexpectedly overwritten")
	}
}

func TestValidateAPIHostnameOverride_EnvSuccess(t *testing.T) {

	value := os.Getenv("PINGONE_API_SERVICE_HOSTNAME")

	if value == "" {
		t.Fatalf("Required environment variable PINGONE_API_SERVICE_HOSTNAME not set")
	}

	config := &Config{}

	if err := config.validateAPIHostnameOverride(); err != nil {
		t.Fatalf("Parameter not successfully verified: %s", err)
	}

	if *config.APIHostnameOverride != value {
		t.Fatalf("Parameter default not taken from environment")
	}
}

func TestValidateAPIHostnameOverride_InvalidFormat(t *testing.T) {
	value := "https://dummyhostname"

	config := &Config{
		APIHostnameOverride: &value,
	}

	if err := config.validateAPIHostnameOverride(); err == nil {
		t.Fatalf("Invalid Parameter format not successfully verified")
	}

	if err := config.validateAPIHostnameOverride(); !strings.HasPrefix(err.Error(), "Invalid parameter format") {
		t.Fatalf("Invalid Parameter format not successfully verified: %s", err.Error())
	}
}

func TestValidateAuthHostnameOverride_ConfigSuccess(t *testing.T) {

	value := "auth.ping-devops.com"

	config := &Config{
		AuthHostnameOverride: &value,
	}

	if err := config.validateAuthHostnameOverride(); err != nil {
		t.Fatalf("Parameter not successfully verified: %s", err)
	}

	if *config.AuthHostnameOverride != value {
		t.Fatalf("Parameter unexpectedly overwritten")
	}
}

func TestValidateAuthHostnameOverride_EnvSuccess(t *testing.T) {

	value := os.Getenv("PINGONE_AUTH_SERVICE_HOSTNAME")

	if value == "" {
		t.Fatalf("Required environment variable PINGONE_AUTH_SERVICE_HOSTNAME not set")
	}

	config := &Config{}

	if err := config.validateAuthHostnameOverride(); err != nil {
		t.Fatalf("Parameter not successfully verified: %s", err)
	}

	if *config.AuthHostnameOverride != value {
		t.Fatalf("Parameter default not taken from environment")
	}
}

func TestValidateAuthHostnameOverride_InvalidFormat(t *testing.T) {
	value := "https://dummyhostname"

	config := &Config{
		AuthHostnameOverride: &value,
	}

	if err := config.validateAuthHostnameOverride(); err == nil {
		t.Fatalf("Invalid Parameter format not successfully verified")
	}

	if err := config.validateAuthHostnameOverride(); !strings.HasPrefix(err.Error(), "Invalid parameter format") {
		t.Fatalf("Invalid Parameter format not successfully verified: %s", err.Error())
	}
}

func TestValidateClientID_ConfigSuccess(t *testing.T) {

	value := "9c052a8a-14be-44e4-8f07-2662569994ce" // dummy UUID

	config := &Config{
		ClientID: &value,
	}

	if err := config.validateClientID(); err != nil {
		t.Fatalf("Parameter not successfully verified: %s", err)
	}

	if *config.ClientID != value {
		t.Fatalf("Parameter unexpectedly overwritten")
	}
}

func TestValidateClientID_EnvSuccess(t *testing.T) {

	value := os.Getenv("PINGONE_CLIENT_ID")

	if value == "" {
		t.Fatalf("Required environment variable PINGONE_CLIENT_ID not set")
	}

	config := &Config{}

	if err := config.validateClientID(); err != nil {
		t.Fatalf("Parameter not successfully verified: %s", err)
	}

	if *config.ClientID != value {
		t.Fatalf("Parameter default not taken from environment")
	}
}

func TestValidateClientID_InvalidFormat(t *testing.T) {
	value := "dummyID"

	config := &Config{
		ClientID: &value,
	}

	if err := config.validateClientID(); !strings.HasPrefix(err.Error(), "Invalid parameter format") {
		t.Fatalf("Invalid Parameter format not successfully verified: %s", err.Error())
	}
}

func TestValidateClientSecret_ConfigSuccess(t *testing.T) {

	value := "testclientsecret"

	config := &Config{
		ClientSecret: &value,
	}

	if err := config.validateClientSecret(); err != nil {
		t.Fatalf("Parameter not successfully verified: %s", err)
	}

	if *config.ClientSecret != value {
		t.Fatalf("Parameter unexpectedly overwritten")
	}
}

func TestValidateClientSecret_EnvSuccess(t *testing.T) {

	value := os.Getenv("PINGONE_CLIENT_SECRET")

	if value == "" {
		t.Fatalf("Required environment variable PINGONE_CLIENT_SECRET not set")
	}

	config := &Config{}

	if err := config.validateClientSecret(); err != nil {
		t.Fatalf("Parameter not successfully verified: %s", err)
	}

	if *config.ClientSecret != value {
		t.Fatalf("Parameter default not taken from environment")
	}
}

func TestValidateClientSecret_InvalidFormat(t *testing.T) {
	t.Skip()
}

func TestValidateEnvironmentID_ConfigSuccess(t *testing.T) {

	value := "9c052a8a-14be-44e4-8f07-2662569994ce" // dummy UUID

	config := &Config{
		EnvironmentID: &value,
	}

	if err := config.validateEnvironmentID(); err != nil {
		t.Fatalf("Parameter not successfully verified: %s", err)
	}

	if *config.EnvironmentID != value {
		t.Fatalf("Parameter unexpectedly overwritten")
	}
}

func TestValidateEnvironmentID_EnvSuccess(t *testing.T) {

	value := os.Getenv("PINGONE_ENVIRONMENT_ID")

	if value == "" {
		t.Fatalf("Required environment variable PINGONE_ENVIRONMENT_ID not set")
	}

	config := &Config{}

	if err := config.validateEnvironmentID(); err != nil {
		t.Fatalf("Parameter not successfully verified: %s", err)
	}

	if *config.EnvironmentID != value {
		t.Fatalf("Parameter default not taken from environment")
	}
}

func TestValidateEnvironmentID_InvalidFormat(t *testing.T) {
	value := "dummyID"

	config := &Config{
		EnvironmentID: &value,
	}

	if err := config.validateEnvironmentID(); !strings.HasPrefix(err.Error(), "Invalid parameter format") {
		t.Fatalf("Invalid Parameter format not successfully verified: %s", err.Error())
	}
}

func TestValidateRegion_ConfigSuccess(t *testing.T) {

	value := "Europe"

	config := &Config{
		Region: value,
	}

	if err := config.validateRegion(); err != nil {
		t.Fatalf("Parameter not successfully verified: %s", err)
	}

	if config.Region != value {
		t.Fatalf("Parameter unexpectedly overwritten")
	}
}

func TestValidateRegion_EnvSuccess(t *testing.T) {

	value := os.Getenv("PINGONE_REGION")

	if value == "" {
		t.Fatalf("Required environment variable PINGONE_REGION not set")
	}

	config := &Config{}

	if err := config.validateRegion(); err != nil {
		t.Fatalf("Parameter not successfully verified: %s", err)
	}

	if config.Region != value {
		t.Fatalf("Parameter default not taken from environment")
	}
}

func TestValidateRegion_InvalidValue(t *testing.T) {
	value := "Tatooine"

	config := &Config{
		Region: value,
	}

	if err := config.validateRegion(); !strings.HasPrefix(err.Error(), "Invalid region value") {
		t.Fatalf("Invalid Parameter format not successfully verified: %s", err.Error())
	}
}
