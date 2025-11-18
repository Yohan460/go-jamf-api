# \ImpactAlertNotificationSettingsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ImpactAlertNotificationSettingsGet**](ImpactAlertNotificationSettingsAPI.md#V1ImpactAlertNotificationSettingsGet) | **Get** /v1/impact-alert-notification-settings | Get Impact Alert Notification Settings
[**V1ImpactAlertNotificationSettingsPut**](ImpactAlertNotificationSettingsAPI.md#V1ImpactAlertNotificationSettingsPut) | **Put** /v1/impact-alert-notification-settings | Update Impact Alert Notification Settings



## V1ImpactAlertNotificationSettingsGet

> ImpactAlertNotificationSettingsV1 V1ImpactAlertNotificationSettingsGet(ctx).Execute()

Get Impact Alert Notification Settings



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
	resp, r, err := apiClient.ImpactAlertNotificationSettingsAPI.V1ImpactAlertNotificationSettingsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImpactAlertNotificationSettingsAPI.V1ImpactAlertNotificationSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ImpactAlertNotificationSettingsGet`: ImpactAlertNotificationSettingsV1
	fmt.Fprintf(os.Stdout, "Response from `ImpactAlertNotificationSettingsAPI.V1ImpactAlertNotificationSettingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ImpactAlertNotificationSettingsGetRequest struct via the builder pattern


### Return type

[**ImpactAlertNotificationSettingsV1**](ImpactAlertNotificationSettingsV1.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ImpactAlertNotificationSettingsPut

> V1ImpactAlertNotificationSettingsPut(ctx).ImpactAlertNotificationSettingsV1(impactAlertNotificationSettingsV1).Execute()

Update Impact Alert Notification Settings



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
	impactAlertNotificationSettingsV1 := *openapiclient.NewImpactAlertNotificationSettingsV1(false, false, false, false) // ImpactAlertNotificationSettingsV1 | Configure Access Management settings

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImpactAlertNotificationSettingsAPI.V1ImpactAlertNotificationSettingsPut(context.Background()).ImpactAlertNotificationSettingsV1(impactAlertNotificationSettingsV1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImpactAlertNotificationSettingsAPI.V1ImpactAlertNotificationSettingsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ImpactAlertNotificationSettingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **impactAlertNotificationSettingsV1** | [**ImpactAlertNotificationSettingsV1**](ImpactAlertNotificationSettingsV1.md) | Configure Access Management settings | 

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

