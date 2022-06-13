# \JamfProUserAccountSettingsPreviewApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserObjPreferenceKeyDelete**](JamfProUserAccountSettingsPreviewApi.md#UserObjPreferenceKeyDelete) | **Delete** /user/obj/preference/{key} | Remove specified setting for authenticated user 
[**UserObjPreferenceKeyGet**](JamfProUserAccountSettingsPreviewApi.md#UserObjPreferenceKeyGet) | **Get** /user/obj/preference/{key} | Get the user setting for the authenticated user and key 
[**UserObjPreferenceKeyPut**](JamfProUserAccountSettingsPreviewApi.md#UserObjPreferenceKeyPut) | **Put** /user/obj/preference/{key} | Persist the user setting 



## UserObjPreferenceKeyDelete

> UserObjPreferenceKeyDelete(ctx, key).Execute()

Remove specified setting for authenticated user 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | key of user setting to be persisted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JamfProUserAccountSettingsPreviewApi.UserObjPreferenceKeyDelete(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProUserAccountSettingsPreviewApi.UserObjPreferenceKeyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | key of user setting to be persisted | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserObjPreferenceKeyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserObjPreferenceKeyGet

> map[string]interface{} UserObjPreferenceKeyGet(ctx, key).Execute()

Get the user setting for the authenticated user and key 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | user setting to be retrieved

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JamfProUserAccountSettingsPreviewApi.UserObjPreferenceKeyGet(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProUserAccountSettingsPreviewApi.UserObjPreferenceKeyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserObjPreferenceKeyGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `JamfProUserAccountSettingsPreviewApi.UserObjPreferenceKeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | user setting to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserObjPreferenceKeyGetRequest struct via the builder pattern


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


## UserObjPreferenceKeyPut

> map[string]interface{} UserObjPreferenceKeyPut(ctx, key).Body(body).Execute()

Persist the user setting 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    key := "key_example" // string | key of user setting to be persisted
    body := "body_example" // string | user setting value to be persisted (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JamfProUserAccountSettingsPreviewApi.UserObjPreferenceKeyPut(context.Background(), key).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProUserAccountSettingsPreviewApi.UserObjPreferenceKeyPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserObjPreferenceKeyPut`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `JamfProUserAccountSettingsPreviewApi.UserObjPreferenceKeyPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | key of user setting to be persisted | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserObjPreferenceKeyPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | user setting value to be persisted | 

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

