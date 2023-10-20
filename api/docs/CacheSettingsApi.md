# \CacheSettingsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CacheSettingsGet**](CacheSettingsAPI.md#V1CacheSettingsGet) | **Get** /v1/cache-settings | Get Cache Settings 
[**V1CacheSettingsPut**](CacheSettingsAPI.md#V1CacheSettingsPut) | **Put** /v1/cache-settings | Update Cache Settings 



## V1CacheSettingsGet

> CacheSettings V1CacheSettingsGet(ctx).Execute()

Get Cache Settings 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CacheSettingsAPI.V1CacheSettingsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CacheSettingsAPI.V1CacheSettingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CacheSettingsGet`: CacheSettings
    fmt.Fprintf(os.Stdout, "Response from `CacheSettingsAPI.V1CacheSettingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CacheSettingsGetRequest struct via the builder pattern


### Return type

[**CacheSettings**](CacheSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CacheSettingsPut

> CacheSettings V1CacheSettingsPut(ctx).CacheSettings(cacheSettings).Execute()

Update Cache Settings 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    cacheSettings := *openapiclient.NewCacheSettings("ehcache", int32(120), "24864549-94ea-4cc1-bb80-d7fb392c6556", []openapiclient.MemcachedEndpoints{*openapiclient.NewMemcachedEndpoints()}) // CacheSettings | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CacheSettingsAPI.V1CacheSettingsPut(context.Background()).CacheSettings(cacheSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CacheSettingsAPI.V1CacheSettingsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CacheSettingsPut`: CacheSettings
    fmt.Fprintf(os.Stdout, "Response from `CacheSettingsAPI.V1CacheSettingsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CacheSettingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cacheSettings** | [**CacheSettings**](CacheSettings.md) |  | 

### Return type

[**CacheSettings**](CacheSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

