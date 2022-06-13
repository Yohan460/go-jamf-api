# \SmartMobileDeviceGroupsPreviewApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1MobileDevicesIdRecalculateSmartGroupsPost**](SmartMobileDeviceGroupsPreviewApi.md#V1MobileDevicesIdRecalculateSmartGroupsPost) | **Post** /v1/mobile-devices/{id}/recalculate-smart-groups | Recalculate all smart groups for the given device id and then return count of smart groups that device fall into 
[**V1SmartMobileDeviceGroupsIdRecalculatePost**](SmartMobileDeviceGroupsPreviewApi.md#V1SmartMobileDeviceGroupsIdRecalculatePost) | **Post** /v1/smart-mobile-device-groups/{id}/recalculate | Recalculate a smart group for the given id then return the ids for the devices in the smart group 



## V1MobileDevicesIdRecalculateSmartGroupsPost

> RecalculationResults V1MobileDevicesIdRecalculateSmartGroupsPost(ctx, id).Execute()

Recalculate all smart groups for the given device id and then return count of smart groups that device fall into 



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
    id := int32(56) // int32 | id of mobile device

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmartMobileDeviceGroupsPreviewApi.V1MobileDevicesIdRecalculateSmartGroupsPost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartMobileDeviceGroupsPreviewApi.V1MobileDevicesIdRecalculateSmartGroupsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1MobileDevicesIdRecalculateSmartGroupsPost`: RecalculationResults
    fmt.Fprintf(os.Stdout, "Response from `SmartMobileDeviceGroupsPreviewApi.V1MobileDevicesIdRecalculateSmartGroupsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id of mobile device | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDevicesIdRecalculateSmartGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RecalculationResults**](RecalculationResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SmartMobileDeviceGroupsIdRecalculatePost

> RecalculationResults V1SmartMobileDeviceGroupsIdRecalculatePost(ctx, id).Execute()

Recalculate a smart group for the given id then return the ids for the devices in the smart group 



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
    id := int32(56) // int32 | instance id of smart group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmartMobileDeviceGroupsPreviewApi.V1SmartMobileDeviceGroupsIdRecalculatePost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartMobileDeviceGroupsPreviewApi.V1SmartMobileDeviceGroupsIdRecalculatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SmartMobileDeviceGroupsIdRecalculatePost`: RecalculationResults
    fmt.Fprintf(os.Stdout, "Response from `SmartMobileDeviceGroupsPreviewApi.V1SmartMobileDeviceGroupsIdRecalculatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | instance id of smart group | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SmartMobileDeviceGroupsIdRecalculatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RecalculationResults**](RecalculationResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

