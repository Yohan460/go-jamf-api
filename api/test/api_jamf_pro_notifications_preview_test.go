/*
Jamf Pro API

Testing JamfProNotificationsPreviewAPIService

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

func Test_api_JamfProNotificationsPreviewAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test JamfProNotificationsPreviewAPIService NotificationsAlertsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.JamfProNotificationsPreviewAPI.NotificationsAlertsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JamfProNotificationsPreviewAPIService NotificationsAlertsIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		httpRes, err := apiClient.JamfProNotificationsPreviewAPI.NotificationsAlertsIdDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JamfProNotificationsPreviewAPIService NotificationsAlertsTypeIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32
		var type_ NotificationType

		httpRes, err := apiClient.JamfProNotificationsPreviewAPI.NotificationsAlertsTypeIdDelete(context.Background(), id, type_).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
