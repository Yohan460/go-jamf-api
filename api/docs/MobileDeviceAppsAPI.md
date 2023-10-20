# \MobileDeviceAppsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1MobileDeviceAppsReinstallAppConfigPost**](MobileDeviceAppsAPI.md#V1MobileDeviceAppsReinstallAppConfigPost) | **Post** /v1/mobile-device-apps/reinstall-app-config | Reinstall App Config for Managed iOS Apps 



## V1MobileDeviceAppsReinstallAppConfigPost

> V1MobileDeviceAppsReinstallAppConfigPost(ctx).AppConfigReinstallCode(appConfigReinstallCode).Execute()

Reinstall App Config for Managed iOS Apps 



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
    appConfigReinstallCode := *openapiclient.NewAppConfigReinstallCode() // AppConfigReinstallCode | The $APP_CONFIG_REINSTALL_CODE variable for the specific device and app supplied by the managed iOS app's current App Config. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MobileDeviceAppsAPI.V1MobileDeviceAppsReinstallAppConfigPost(context.Background()).AppConfigReinstallCode(appConfigReinstallCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDeviceAppsAPI.V1MobileDeviceAppsReinstallAppConfigPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDeviceAppsReinstallAppConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appConfigReinstallCode** | [**AppConfigReinstallCode**](AppConfigReinstallCode.md) | The $APP_CONFIG_REINSTALL_CODE variable for the specific device and app supplied by the managed iOS app&#39;s current App Config.  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

