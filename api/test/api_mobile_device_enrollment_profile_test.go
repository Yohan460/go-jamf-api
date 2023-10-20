/*
Jamf Pro API

Testing MobileDeviceEnrollmentProfileAPIService

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

func Test_api_MobileDeviceEnrollmentProfileAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MobileDeviceEnrollmentProfileAPIService V1MobileDeviceEnrollmentProfileIdDownloadProfileGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.MobileDeviceEnrollmentProfileAPI.V1MobileDeviceEnrollmentProfileIdDownloadProfileGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}