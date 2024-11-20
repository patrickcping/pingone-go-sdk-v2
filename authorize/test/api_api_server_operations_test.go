/*
PingOne Platform API - Authorize

Testing APIServerOperationsApiService

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

func Test_authorize_APIServerOperationsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test APIServerOperationsApiService CreateAPIServerOperation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var apiServerID string

		resp, httpRes, err := apiClient.APIServerOperationsApi.CreateAPIServerOperation(context.Background(), environmentID, apiServerID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test APIServerOperationsApiService DeleteAPIServerOperation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var apiServerID string
		var apiServerOperationID string

		httpRes, err := apiClient.APIServerOperationsApi.DeleteAPIServerOperation(context.Background(), environmentID, apiServerID, apiServerOperationID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test APIServerOperationsApiService ReadAllAPIServerOperations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var apiServerID string

		pagedIterator := apiClient.APIServerOperationsApi.ReadAllAPIServerOperations(context.Background(), environmentID, apiServerID).Execute()

		for pageCursor, err := range pagedIterator {
			require.Nil(t, err)
			require.NotNil(t, pageCursor.Data)
			assert.Equal(t, 200, pageCursor.HTTPResponse.StatusCode)
		}

	})

	t.Run("Test APIServerOperationsApiService ReadOneAPIServerOperation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var apiServerID string
		var apiServerOperationID string

		resp, httpRes, err := apiClient.APIServerOperationsApi.ReadOneAPIServerOperation(context.Background(), environmentID, apiServerID, apiServerOperationID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test APIServerOperationsApiService UpdateAPIServerOperation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var apiServerID string
		var apiServerOperationID string

		resp, httpRes, err := apiClient.APIServerOperationsApi.UpdateAPIServerOperation(context.Background(), environmentID, apiServerID, apiServerOperationID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
