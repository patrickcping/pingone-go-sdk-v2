package fsconfig

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReadFile_Sample_Success(t *testing.T) {
	t.Parallel()

	expectedData := map[string]string{
		"PINGONE_API_URL":                  "https://api.pingone.eu/v1",
		"PINGONE_AUTH_URL":                 "https://auth.pingone.eu",
		"PINGONE_ENVIRONMENT_ID":           "da0583be-61af-40ba-acbf-randomstring",
		"PINGONE_WORKER_APP_CLIENT_ID":     "d74589c2-301e-4def-8f5f-randomstring",
		"PINGONE_WORKER_APP_GRANT_TYPE":    "client_credentials",
		"PINGONE_WORKER_APP_REDIRECT_URI":  "http://localhost:8000",
		"PINGONE_WORKER_APP_CLIENT_SECRET": "randomstringdoesnotexistbd7f18c6-06b3-420b-83e5-ca8a0d46b308",
		"PINGONE_USERNAME":                 "jsmith@example.com",
	}

	filedata, err := ReadFile("./testfiles/config_sample.env")

	if err != nil {
		t.Fatalf("Filedata not successfully retrieved: %v", err)
	}

	if !reflect.DeepEqual(filedata, expectedData) {
		t.Fatalf("ReadFile resulted in %v, expected %v", filedata, expectedData)
	}
}

func TestReadFile_Sample_UnsupportedGrantType_Success(t *testing.T) {
	t.Parallel()

	expectedData := map[string]string{
		"PINGONE_API_URL":                 "https://api.pingone.eu/v1",
		"PINGONE_AUTH_URL":                "https://auth.pingone.eu",
		"PINGONE_ENVIRONMENT_ID":          "da0583be-61af-40ba-acbf-randomstring",
		"PINGONE_WORKER_APP_CLIENT_ID":    "d74589c2-301e-4def-8f5f-randomstring",
		"PINGONE_WORKER_APP_GRANT_TYPE":   "authorization_code",
		"PINGONE_WORKER_APP_REDIRECT_URI": "http://localhost:8000",
		"PINGONE_USERNAME":                "jsmith@example.com",
	}

	filedata, err := ReadFile("./testfiles/config_sample_unsupportedgranttype.env")

	if err != nil {
		t.Fatalf("Filedata not successfully retrieved: %v", err)
	}

	if !reflect.DeepEqual(filedata, expectedData) {
		t.Fatalf("ReadFile resulted in %v, expected %v", filedata, expectedData)
	}
}

func TestReadFile_Sample_NoVals_Success(t *testing.T) {
	t.Parallel()

	expectedData := map[string]string{}

	filedata, err := ReadFile("./testfiles/config_sample_novals.env")

	if err != nil {
		t.Fatalf("Filedata not successfully retrieved: %v", err)
	}

	if !reflect.DeepEqual(filedata, expectedData) {
		t.Fatalf("ReadFile resulted in %v, expected %v", filedata, expectedData)
	}
}

func TestReadFile_DoesNotExist(t *testing.T) {
	t.Parallel()

	_, err := ReadFile("./testfiles/config_sample_doesnotexist.env")

	if err == nil {
		t.Fatalf("File not expected to be found but no error exists")
	}

	if fmt.Sprint(err) != "open ./testfiles/config_sample_doesnotexist.env: no such file or directory" {
		t.Fatalf("File not found, but error message inconsistent")
	}
}
