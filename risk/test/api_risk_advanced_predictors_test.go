/*
PingOne Platform API - PingOne Risk

Testing RiskAdvancedPredictorsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package risk

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/risk"
)

func Test_risk_RiskAdvancedPredictorsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RiskAdvancedPredictorsApiService CreateRiskPredictor", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string

		resp, httpRes, err := apiClient.RiskAdvancedPredictorsApi.CreateRiskPredictor(context.Background(), environmentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RiskAdvancedPredictorsApiService DeleteRiskAdvancedPredictor", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var riskPredictorID string

		httpRes, err := apiClient.RiskAdvancedPredictorsApi.DeleteRiskAdvancedPredictor(context.Background(), environmentID, riskPredictorID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RiskAdvancedPredictorsApiService ReadAllRiskPredictors", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string

		pagedIterator := apiClient.RiskAdvancedPredictorsApi.ReadAllRiskPredictors(context.Background(), environmentID).Execute()

		for pageCursor, err := range pagedIterator {
			require.Nil(t, err)
			require.NotNil(t, pageCursor.Data)
			assert.Equal(t, 200, pageCursor.HTTPResponse.StatusCode)
		}

	})

	t.Run("Test RiskAdvancedPredictorsApiService ReadOneRiskPredictor", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var riskPredictorID string

		resp, httpRes, err := apiClient.RiskAdvancedPredictorsApi.ReadOneRiskPredictor(context.Background(), environmentID, riskPredictorID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RiskAdvancedPredictorsApiService UpdateRiskPredictor", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var riskPredictorID string

		resp, httpRes, err := apiClient.RiskAdvancedPredictorsApi.UpdateRiskPredictor(context.Background(), environmentID, riskPredictorID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
