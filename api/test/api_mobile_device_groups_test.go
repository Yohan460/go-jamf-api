/*
Jamf Pro API

Testing MobileDeviceGroupsAPIService

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

func Test_api_MobileDeviceGroupsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MobileDeviceGroupsAPIService V1MobileDeviceGroupsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MobileDeviceGroupsAPI.V1MobileDeviceGroupsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MobileDeviceGroupsAPIService V1MobileDeviceGroupsStaticGroupMembershipIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupMembershipIdGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MobileDeviceGroupsAPIService V1MobileDeviceGroupsStaticGroupsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MobileDeviceGroupsAPIService V1MobileDeviceGroupsStaticGroupsIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupsIdDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MobileDeviceGroupsAPIService V1MobileDeviceGroupsStaticGroupsIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupsIdGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MobileDeviceGroupsAPIService V1MobileDeviceGroupsStaticGroupsIdPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupsIdPatch(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MobileDeviceGroupsAPIService V1MobileDeviceGroupsStaticGroupsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
