# \ServiceDiscoveryEnrollmentAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ServiceDiscoveryEnrollmentWellKnownSettingsGet**](ServiceDiscoveryEnrollmentAPI.md#V1ServiceDiscoveryEnrollmentWellKnownSettingsGet) | **Get** /v1/service-discovery-enrollment/well-known-settings | Get service discovery well-known settings for all organizations
[**V1ServiceDiscoveryEnrollmentWellKnownSettingsPut**](ServiceDiscoveryEnrollmentAPI.md#V1ServiceDiscoveryEnrollmentWellKnownSettingsPut) | **Put** /v1/service-discovery-enrollment/well-known-settings | Update service discovery well-known settings



## V1ServiceDiscoveryEnrollmentWellKnownSettingsGet

> WellKnownSettingsResponse V1ServiceDiscoveryEnrollmentWellKnownSettingsGet(ctx).Execute()

Get service discovery well-known settings for all organizations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceDiscoveryEnrollmentAPI.V1ServiceDiscoveryEnrollmentWellKnownSettingsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceDiscoveryEnrollmentAPI.V1ServiceDiscoveryEnrollmentWellKnownSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ServiceDiscoveryEnrollmentWellKnownSettingsGet`: WellKnownSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceDiscoveryEnrollmentAPI.V1ServiceDiscoveryEnrollmentWellKnownSettingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ServiceDiscoveryEnrollmentWellKnownSettingsGetRequest struct via the builder pattern


### Return type

[**WellKnownSettingsResponse**](WellKnownSettingsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ServiceDiscoveryEnrollmentWellKnownSettingsPut

> V1ServiceDiscoveryEnrollmentWellKnownSettingsPut(ctx).WellKnownSettingsRequest(wellKnownSettingsRequest).Execute()

Update service discovery well-known settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {
	wellKnownSettingsRequest := *openapiclient.NewWellKnownSettingsRequest([]openapiclient.WellKnownSetting{*openapiclient.NewWellKnownSetting("a1b2c3d4-e5f6-7890-abcd-ef1234567890", openapiclient.ServiceDiscoveryVersion("none"))}) // WellKnownSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceDiscoveryEnrollmentAPI.V1ServiceDiscoveryEnrollmentWellKnownSettingsPut(context.Background()).WellKnownSettingsRequest(wellKnownSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceDiscoveryEnrollmentAPI.V1ServiceDiscoveryEnrollmentWellKnownSettingsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ServiceDiscoveryEnrollmentWellKnownSettingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wellKnownSettingsRequest** | [**WellKnownSettingsRequest**](WellKnownSettingsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

