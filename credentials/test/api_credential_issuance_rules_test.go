/*
PingOne Platform API - Credentials

Testing CredentialIssuanceRulesApiService

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

func Test_credentials_CredentialIssuanceRulesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CredentialIssuanceRulesApiService ApplyCredentialIssuanceRuleStagedChanges", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string
		var credentialTypeID string
		var credentialIssuanceRuleID string

		resp, httpRes, err := apiClient.CredentialIssuanceRulesApi.ApplyCredentialIssuanceRuleStagedChanges(context.Background(), environmentID, credentialTypeID, credentialIssuanceRuleID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CredentialIssuanceRulesApiService CreateCredentialIssuanceRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string
		var credentialTypeID string

		resp, httpRes, err := apiClient.CredentialIssuanceRulesApi.CreateCredentialIssuanceRule(context.Background(), environmentID, credentialTypeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CredentialIssuanceRulesApiService DeleteCredentialIssuanceRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string
		var credentialTypeID string
		var credentialIssuanceRuleID string

		httpRes, err := apiClient.CredentialIssuanceRulesApi.DeleteCredentialIssuanceRule(context.Background(), environmentID, credentialTypeID, credentialIssuanceRuleID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CredentialIssuanceRulesApiService ReadAllCredentialIssuanceRules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string
		var credentialTypeID string

		resp, httpRes, err := apiClient.CredentialIssuanceRulesApi.ReadAllCredentialIssuanceRules(context.Background(), environmentID, credentialTypeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CredentialIssuanceRulesApiService ReadCredentialIssuanceRuleStagedChanges", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string
		var credentialTypeID string
		var credentialIssuanceRuleID string

		resp, httpRes, err := apiClient.CredentialIssuanceRulesApi.ReadCredentialIssuanceRuleStagedChanges(context.Background(), environmentID, credentialTypeID, credentialIssuanceRuleID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CredentialIssuanceRulesApiService ReadCredentialIssuanceRuleUsageCounts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string
		var credentialTypeID string
		var credentialIssuanceRuleID string

		resp, httpRes, err := apiClient.CredentialIssuanceRulesApi.ReadCredentialIssuanceRuleUsageCounts(context.Background(), environmentID, credentialTypeID, credentialIssuanceRuleID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CredentialIssuanceRulesApiService ReadCredentialIssuanceRuleUsageDetails", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string
		var credentialTypeID string
		var credentialIssuanceRuleID string

		resp, httpRes, err := apiClient.CredentialIssuanceRulesApi.ReadCredentialIssuanceRuleUsageDetails(context.Background(), environmentID, credentialTypeID, credentialIssuanceRuleID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CredentialIssuanceRulesApiService ReadOneCredentialIssuanceRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string
		var credentialTypeID string
		var credentialIssuanceRuleID string

		resp, httpRes, err := apiClient.CredentialIssuanceRulesApi.ReadOneCredentialIssuanceRule(context.Background(), environmentID, credentialTypeID, credentialIssuanceRuleID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CredentialIssuanceRulesApiService UpdateCredentialIssuanceRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string
		var credentialTypeID string
		var credentialIssuanceRuleID string

		resp, httpRes, err := apiClient.CredentialIssuanceRulesApi.UpdateCredentialIssuanceRule(context.Background(), environmentID, credentialTypeID, credentialIssuanceRuleID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
