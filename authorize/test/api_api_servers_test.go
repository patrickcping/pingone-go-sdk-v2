/*
PingOne Platform API - Authorize

Testing APIServersApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package authorize

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/authorize"
)

func Test_authorize_APIServersApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test APIServersApiService CreateAPIServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string

		resp, httpRes, err := apiClient.APIServersApi.CreateAPIServer(context.Background(), environmentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test APIServersApiService DeleteAPIServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var apiServerID string

		httpRes, err := apiClient.APIServersApi.DeleteAPIServer(context.Background(), environmentID, apiServerID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test APIServersApiService ReadAllAPIServers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string

		resp, httpRes, err := apiClient.APIServersApi.ReadAllAPIServers(context.Background(), environmentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test APIServersApiService ReadOneAPIServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var apiServerID string

		resp, httpRes, err := apiClient.APIServersApi.ReadOneAPIServer(context.Background(), environmentID, apiServerID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test APIServersApiService UpdateAPIServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var apiServerID string

		resp, httpRes, err := apiClient.APIServersApi.UpdateAPIServer(context.Background(), environmentID, apiServerID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
