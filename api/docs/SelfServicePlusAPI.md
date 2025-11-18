# \SelfServicePlusAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSelfServicePlusFeatureToggleEnabled**](SelfServicePlusAPI.md#GetSelfServicePlusFeatureToggleEnabled) | **Get** /v1/self-service-plus/feature-toggle/enabled | Determines if Self Service Plus feature toggle is enabled.
[**GetSelfServicePlusSettings**](SelfServicePlusAPI.md#GetSelfServicePlusSettings) | **Get** /v1/self-service-plus/settings | Get Self Service Plus settings.
[**PutSelfServicePlusSettings**](SelfServicePlusAPI.md#PutSelfServicePlusSettings) | **Put** /v1/self-service-plus/settings | Save Self Service Plus settings.



## GetSelfServicePlusFeatureToggleEnabled

> GetSelfServicePlusFeatureToggleEnabled(ctx).Execute()

Determines if Self Service Plus feature toggle is enabled.



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
	r, err := apiClient.SelfServicePlusAPI.GetSelfServicePlusFeatureToggleEnabled(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfServicePlusAPI.GetSelfServicePlusFeatureToggleEnabled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfServicePlusFeatureToggleEnabledRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfServicePlusSettings

> SelfServicePlusSettings GetSelfServicePlusSettings(ctx).Execute()

Get Self Service Plus settings.



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
	resp, r, err := apiClient.SelfServicePlusAPI.GetSelfServicePlusSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfServicePlusAPI.GetSelfServicePlusSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSelfServicePlusSettings`: SelfServicePlusSettings
	fmt.Fprintf(os.Stdout, "Response from `SelfServicePlusAPI.GetSelfServicePlusSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfServicePlusSettingsRequest struct via the builder pattern


### Return type

[**SelfServicePlusSettings**](SelfServicePlusSettings.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSelfServicePlusSettings

> PutSelfServicePlusSettings(ctx).SelfServicePlusSettings(selfServicePlusSettings).Execute()

Save Self Service Plus settings.



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
	selfServicePlusSettings := *openapiclient.NewSelfServicePlusSettings() // SelfServicePlusSettings | Self Service Plus settings

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SelfServicePlusAPI.PutSelfServicePlusSettings(context.Background()).SelfServicePlusSettings(selfServicePlusSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfServicePlusAPI.PutSelfServicePlusSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutSelfServicePlusSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selfServicePlusSettings** | [**SelfServicePlusSettings**](SelfServicePlusSettings.md) | Self Service Plus settings | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

