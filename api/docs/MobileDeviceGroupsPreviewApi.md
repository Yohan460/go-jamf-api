# \MobileDeviceGroupsPreviewApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1MobileDeviceGroupsGet**](MobileDeviceGroupsPreviewApi.md#V1MobileDeviceGroupsGet) | **Get** /v1/mobile-device-groups | Return the list of all Mobile Device Groups 



## V1MobileDeviceGroupsGet

> []MobileDeviceGroup V1MobileDeviceGroupsGet(ctx).Execute()

Return the list of all Mobile Device Groups 



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileDeviceGroupsPreviewApi.V1MobileDeviceGroupsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDeviceGroupsPreviewApi.V1MobileDeviceGroupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1MobileDeviceGroupsGet`: []MobileDeviceGroup
    fmt.Fprintf(os.Stdout, "Response from `MobileDeviceGroupsPreviewApi.V1MobileDeviceGroupsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDeviceGroupsGetRequest struct via the builder pattern


### Return type

[**[]MobileDeviceGroup**](MobileDeviceGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

