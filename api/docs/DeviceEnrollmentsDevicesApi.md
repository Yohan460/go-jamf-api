# \DeviceEnrollmentsDevicesAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DeviceEnrollmentsIdDevicesGet**](DeviceEnrollmentsDevicesAPI.md#V1DeviceEnrollmentsIdDevicesGet) | **Get** /v1/device-enrollments/{id}/devices | Retrieve a list of Devices assigned to the supplied id 



## V1DeviceEnrollmentsIdDevicesGet

> DeviceEnrollmentDeviceSearchResults V1DeviceEnrollmentsIdDevicesGet(ctx, id).Execute()

Retrieve a list of Devices assigned to the supplied id 



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
    id := "id_example" // string | Device Enrollment Instance identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceEnrollmentsDevicesAPI.V1DeviceEnrollmentsIdDevicesGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceEnrollmentsDevicesAPI.V1DeviceEnrollmentsIdDevicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DeviceEnrollmentsIdDevicesGet`: DeviceEnrollmentDeviceSearchResults
    fmt.Fprintf(os.Stdout, "Response from `DeviceEnrollmentsDevicesAPI.V1DeviceEnrollmentsIdDevicesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Device Enrollment Instance identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeviceEnrollmentsIdDevicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceEnrollmentDeviceSearchResults**](DeviceEnrollmentDeviceSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

