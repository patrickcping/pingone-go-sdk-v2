/*
PingOne Platform API - Credentials

Testing DigitalWalletAppsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package credentials

import (
	"context"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/credentials"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_credentials_DigitalWalletAppsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DigitalWalletAppsApiService CreateDigitalWalletApp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string

		resp, httpRes, err := apiClient.DigitalWalletAppsApi.CreateDigitalWalletApp(context.Background(), environmentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DigitalWalletAppsApiService DeleteDigitalWalletApp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string
		var digitalWalletApplicationID string

		httpRes, err := apiClient.DigitalWalletAppsApi.DeleteDigitalWalletApp(context.Background(), environmentID, digitalWalletApplicationID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DigitalWalletAppsApiService ReadAllDigitalWalletApps", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string

		resp, httpRes, err := apiClient.DigitalWalletAppsApi.ReadAllDigitalWalletApps(context.Background(), environmentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DigitalWalletAppsApiService ReadOneDigitalWalletApp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string
		var digitalWalletApplicationID string

		resp, httpRes, err := apiClient.DigitalWalletAppsApi.ReadOneDigitalWalletApp(context.Background(), environmentID, digitalWalletApplicationID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DigitalWalletAppsApiService UpdateDigitalWalletApp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string
		var digitalWalletApplicationID string

		resp, httpRes, err := apiClient.DigitalWalletAppsApi.UpdateDigitalWalletApp(context.Background(), environmentID, digitalWalletApplicationID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
