# \MobileDeviceExtensionAttributesPreviewAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesExtensionAttributesGet**](MobileDeviceExtensionAttributesPreviewAPI.md#DevicesExtensionAttributesGet) | **Get** /devices/extensionAttributes | Get Mobile Device Extension Attribute values placed in select paramter 



## DevicesExtensionAttributesGet

> MobileDeviceExtensionAttributeResults DevicesExtensionAttributesGet(ctx).Select_(select_).Execute()

Get Mobile Device Extension Attribute values placed in select paramter 



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
    select_ := "select__example" // string | Acceptable values currently include: * name  (optional) (default to "name")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileDeviceExtensionAttributesPreviewAPI.DevicesExtensionAttributesGet(context.Background()).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDeviceExtensionAttributesPreviewAPI.DevicesExtensionAttributesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesExtensionAttributesGet`: MobileDeviceExtensionAttributeResults
    fmt.Fprintf(os.Stdout, "Response from `MobileDeviceExtensionAttributesPreviewAPI.DevicesExtensionAttributesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicesExtensionAttributesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **string** | Acceptable values currently include: * name  | [default to &quot;name&quot;]

### Return type

[**MobileDeviceExtensionAttributeResults**](MobileDeviceExtensionAttributeResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

