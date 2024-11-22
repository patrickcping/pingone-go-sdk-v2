/*
PingOne Platform API - PingOne Verify

Testing VoicePhrasesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package verify

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/verify"
)

func Test_verify_VoicePhrasesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VoicePhrasesApiService CreateVoicePhrase", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string

		resp, httpRes, err := apiClient.VoicePhrasesApi.CreateVoicePhrase(context.Background(), environmentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicePhrasesApiService DeleteVoicePhrase", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var voicePhraseID string

		httpRes, err := apiClient.VoicePhrasesApi.DeleteVoicePhrase(context.Background(), environmentID, voicePhraseID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicePhrasesApiService ReadAllVoicePhrases", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string

		pagedIterator := apiClient.VoicePhrasesApi.ReadAllVoicePhrases(context.Background(), environmentID).Execute()

		for pageCursor, err := range pagedIterator {
			require.Nil(t, err)
			require.NotNil(t, pageCursor.Data)
			assert.Equal(t, 200, pageCursor.HTTPResponse.StatusCode)
		}

	})

	t.Run("Test VoicePhrasesApiService ReadOneVoicePhrase", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var voicePhraseID string

		resp, httpRes, err := apiClient.VoicePhrasesApi.ReadOneVoicePhrase(context.Background(), environmentID, voicePhraseID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicePhrasesApiService UpdateVoicePhrase", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var voicePhraseID string

		resp, httpRes, err := apiClient.VoicePhrasesApi.UpdateVoicePhrase(context.Background(), environmentID, voicePhraseID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
