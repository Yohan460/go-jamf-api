# \SelfServiceAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SelfServiceSettingsGet**](SelfServiceAPI.md#V1SelfServiceSettingsGet) | **Get** /v1/self-service/settings | Get an object representation of Self Service settings 
[**V1SelfServiceSettingsPut**](SelfServiceAPI.md#V1SelfServiceSettingsPut) | **Put** /v1/self-service/settings | Put an object representation of Self Service settings 



## V1SelfServiceSettingsGet

> SelfServiceSettings V1SelfServiceSettingsGet(ctx).Execute()

Get an object representation of Self Service settings 



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
	resp, r, err := apiClient.SelfServiceAPI.V1SelfServiceSettingsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceAPI.V1SelfServiceSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SelfServiceSettingsGet`: SelfServiceSettings
	fmt.Fprintf(os.Stdout, "Response from `SelfServiceAPI.V1SelfServiceSettingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SelfServiceSettingsGetRequest struct via the builder pattern


### Return type

[**SelfServiceSettings**](SelfServiceSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SelfServiceSettingsPut

> SelfServiceSettings V1SelfServiceSettingsPut(ctx).SelfServiceSettings(selfServiceSettings).Execute()

Put an object representation of Self Service settings 



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
	selfServiceSettings := *openapiclient.NewSelfServiceSettings() // SelfServiceSettings | object that contains all editable global fields to alter Self Service settings 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SelfServiceAPI.V1SelfServiceSettingsPut(context.Background()).SelfServiceSettings(selfServiceSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceAPI.V1SelfServiceSettingsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SelfServiceSettingsPut`: SelfServiceSettings
	fmt.Fprintf(os.Stdout, "Response from `SelfServiceAPI.V1SelfServiceSettingsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SelfServiceSettingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selfServiceSettings** | [**SelfServiceSettings**](SelfServiceSettings.md) | object that contains all editable global fields to alter Self Service settings  | 

### Return type

[**SelfServiceSettings**](SelfServiceSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

