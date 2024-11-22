/*
PingOne Platform API - PingOne MFA

Testing UserMFADevicesApiService

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

func Test_mfa_UserMFADevicesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserMFADevicesApiService EnvironmentsEnvironmentIDUsersUserIDDevicesDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var userID string

		httpRes, err := apiClient.UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesDelete(context.Background(), environmentID, userID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserMFADevicesApiService EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var userID string
		var deviceID string

		httpRes, err := apiClient.UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDDelete(context.Background(), environmentID, userID, deviceID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserMFADevicesApiService EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var userID string
		var deviceID string

		httpRes, err := apiClient.UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDGet(context.Background(), environmentID, userID, deviceID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserMFADevicesApiService EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDLogsPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var userID string
		var deviceID string

		httpRes, err := apiClient.UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDLogsPut(context.Background(), environmentID, userID, deviceID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserMFADevicesApiService EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDNicknamePut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var userID string
		var deviceID string

		httpRes, err := apiClient.UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDNicknamePut(context.Background(), environmentID, userID, deviceID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserMFADevicesApiService EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var userID string
		var deviceID string

		httpRes, err := apiClient.UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDPost(context.Background(), environmentID, userID, deviceID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserMFADevicesApiService EnvironmentsEnvironmentIDUsersUserIDDevicesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var userID string

		httpRes, err := apiClient.UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesGet(context.Background(), environmentID, userID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserMFADevicesApiService EnvironmentsEnvironmentIDUsersUserIDDevicesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var userID string

		httpRes, err := apiClient.UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesPost(context.Background(), environmentID, userID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}