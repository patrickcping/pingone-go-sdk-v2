/*
PingOne Platform API - PingOne MFA

Testing UserMFAPairingKeysApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package mfa

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func Test_mfa_UserMFAPairingKeysApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserMFAPairingKeysApiService EnvironmentsEnvironmentIDUsersUserIDPairingKeysPairingKeyIDDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var userID string
		var pairingKeyID string

		httpRes, err := apiClient.UserMFAPairingKeysApi.EnvironmentsEnvironmentIDUsersUserIDPairingKeysPairingKeyIDDelete(context.Background(), environmentID, userID, pairingKeyID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserMFAPairingKeysApiService EnvironmentsEnvironmentIDUsersUserIDPairingKeysPairingKeyIDGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var userID string
		var pairingKeyID string

		httpRes, err := apiClient.UserMFAPairingKeysApi.EnvironmentsEnvironmentIDUsersUserIDPairingKeysPairingKeyIDGet(context.Background(), environmentID, userID, pairingKeyID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserMFAPairingKeysApiService EnvironmentsEnvironmentIDUsersUserIDPairingKeysPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var userID string

		httpRes, err := apiClient.UserMFAPairingKeysApi.EnvironmentsEnvironmentIDUsersUserIDPairingKeysPost(context.Background(), environmentID, userID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
