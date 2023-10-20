/*
Jamf Pro API

Testing ComputerInventoryCollectionSettingsAPIService

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

func Test_api_ComputerInventoryCollectionSettingsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ComputerInventoryCollectionSettingsAPIService V1ComputerInventoryCollectionSettingsCustomPathIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ComputerInventoryCollectionSettingsAPI.V1ComputerInventoryCollectionSettingsCustomPathIdDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ComputerInventoryCollectionSettingsAPIService V1ComputerInventoryCollectionSettingsCustomPathPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ComputerInventoryCollectionSettingsAPI.V1ComputerInventoryCollectionSettingsCustomPathPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ComputerInventoryCollectionSettingsAPIService V1ComputerInventoryCollectionSettingsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ComputerInventoryCollectionSettingsAPI.V1ComputerInventoryCollectionSettingsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ComputerInventoryCollectionSettingsAPIService V1ComputerInventoryCollectionSettingsPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ComputerInventoryCollectionSettingsAPI.V1ComputerInventoryCollectionSettingsPatch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}