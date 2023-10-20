/*
Jamf Pro API

Testing SmartMobileDeviceGroupsPreviewAPIService

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

func Test_api_SmartMobileDeviceGroupsPreviewAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SmartMobileDeviceGroupsPreviewAPIService V1MobileDevicesIdRecalculateSmartGroupsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.SmartMobileDeviceGroupsPreviewAPI.V1MobileDevicesIdRecalculateSmartGroupsPost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SmartMobileDeviceGroupsPreviewAPIService V1SmartMobileDeviceGroupsIdRecalculatePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.SmartMobileDeviceGroupsPreviewAPI.V1SmartMobileDeviceGroupsIdRecalculatePost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}