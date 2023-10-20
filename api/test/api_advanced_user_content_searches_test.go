/*
Jamf Pro API

Testing AdvancedUserContentSearchesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/yohan460/go-jamf-api"
)

func Test_api_AdvancedUserContentSearchesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AdvancedUserContentSearchesAPIService V1AdvancedUserContentSearchesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AdvancedUserContentSearchesAPI.V1AdvancedUserContentSearchesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdvancedUserContentSearchesAPIService V1AdvancedUserContentSearchesIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.AdvancedUserContentSearchesAPI.V1AdvancedUserContentSearchesIdDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdvancedUserContentSearchesAPIService V1AdvancedUserContentSearchesIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AdvancedUserContentSearchesAPI.V1AdvancedUserContentSearchesIdGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdvancedUserContentSearchesAPIService V1AdvancedUserContentSearchesIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AdvancedUserContentSearchesAPI.V1AdvancedUserContentSearchesIdPut(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdvancedUserContentSearchesAPIService V1AdvancedUserContentSearchesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AdvancedUserContentSearchesAPI.V1AdvancedUserContentSearchesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}