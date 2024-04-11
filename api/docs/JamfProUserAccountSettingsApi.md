# \JamfProUserAccountSettingsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1UserPreferencesKeyIdDelete**](JamfProUserAccountSettingsAPI.md#V1UserPreferencesKeyIdDelete) | **Delete** /v1/user/preferences/{keyId} | Remove specified setting for authenticated user 
[**V1UserPreferencesKeyIdGet**](JamfProUserAccountSettingsAPI.md#V1UserPreferencesKeyIdGet) | **Get** /v1/user/preferences/{keyId} | Get the user setting for the authenticated user and key 
[**V1UserPreferencesKeyIdPut**](JamfProUserAccountSettingsAPI.md#V1UserPreferencesKeyIdPut) | **Put** /v1/user/preferences/{keyId} | Persist the user setting 
[**V1UserPreferencesSettingsKeyIdGet**](JamfProUserAccountSettingsAPI.md#V1UserPreferencesSettingsKeyIdGet) | **Get** /v1/user/preferences/settings/{keyId} | Get the user preferences for the authenticated user and key. 



## V1UserPreferencesKeyIdDelete

> V1UserPreferencesKeyIdDelete(ctx, keyId).Execute()

Remove specified setting for authenticated user 



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
	keyId := "keyId_example" // string | unique key of user setting to be persisted

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JamfProUserAccountSettingsAPI.V1UserPreferencesKeyIdDelete(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfProUserAccountSettingsAPI.V1UserPreferencesKeyIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | unique key of user setting to be persisted | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UserPreferencesKeyIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UserPreferencesKeyIdGet

> map[string]interface{} V1UserPreferencesKeyIdGet(ctx, keyId).Execute()

Get the user setting for the authenticated user and key 



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
	keyId := "keyId_example" // string | user setting to be retrieved

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfProUserAccountSettingsAPI.V1UserPreferencesKeyIdGet(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfProUserAccountSettingsAPI.V1UserPreferencesKeyIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UserPreferencesKeyIdGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `JamfProUserAccountSettingsAPI.V1UserPreferencesKeyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | user setting to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UserPreferencesKeyIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UserPreferencesKeyIdPut

> map[string]interface{} V1UserPreferencesKeyIdPut(ctx, keyId).Body(body).Execute()

Persist the user setting 



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
	keyId := "keyId_example" // string | unique key of user setting to be persisted
	body := map[string]interface{}{ ... } // map[string]interface{} | user setting value to be persisted (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfProUserAccountSettingsAPI.V1UserPreferencesKeyIdPut(context.Background(), keyId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfProUserAccountSettingsAPI.V1UserPreferencesKeyIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UserPreferencesKeyIdPut`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `JamfProUserAccountSettingsAPI.V1UserPreferencesKeyIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | unique key of user setting to be persisted | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UserPreferencesKeyIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | user setting value to be persisted | 

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UserPreferencesSettingsKeyIdGet

> UserPreferencesSettings V1UserPreferencesSettingsKeyIdGet(ctx, keyId).Execute()

Get the user preferences for the authenticated user and key. 



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
	keyId := "keyId_example" // string | user setting to be retrieved

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfProUserAccountSettingsAPI.V1UserPreferencesSettingsKeyIdGet(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfProUserAccountSettingsAPI.V1UserPreferencesSettingsKeyIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UserPreferencesSettingsKeyIdGet`: UserPreferencesSettings
	fmt.Fprintf(os.Stdout, "Response from `JamfProUserAccountSettingsAPI.V1UserPreferencesSettingsKeyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | user setting to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UserPreferencesSettingsKeyIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserPreferencesSettings**](UserPreferencesSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

