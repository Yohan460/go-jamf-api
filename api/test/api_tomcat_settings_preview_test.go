/*
Jamf Pro API

Testing TomcatSettingsPreviewAPIService

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

func Test_api_TomcatSettingsPreviewAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TomcatSettingsPreviewAPIService SettingsIssueTomcatSslCertificatePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.TomcatSettingsPreviewAPI.SettingsIssueTomcatSslCertificatePost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
