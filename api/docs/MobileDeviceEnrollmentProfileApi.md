# \MobileDeviceEnrollmentProfileAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1MobileDeviceEnrollmentProfileIdDownloadProfileGet**](MobileDeviceEnrollmentProfileAPI.md#V1MobileDeviceEnrollmentProfileIdDownloadProfileGet) | **Get** /v1/mobile-device-enrollment-profile/{id}/download-profile | Retrieve the MDM Enrollment Profile 



## V1MobileDeviceEnrollmentProfileIdDownloadProfileGet

> *os.File V1MobileDeviceEnrollmentProfileIdDownloadProfileGet(ctx, id).Execute()

Retrieve the MDM Enrollment Profile 



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
    id := "id_example" // string | MDM Enrollment Profile identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileDeviceEnrollmentProfileAPI.V1MobileDeviceEnrollmentProfileIdDownloadProfileGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDeviceEnrollmentProfileAPI.V1MobileDeviceEnrollmentProfileIdDownloadProfileGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1MobileDeviceEnrollmentProfileIdDownloadProfileGet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `MobileDeviceEnrollmentProfileAPI.V1MobileDeviceEnrollmentProfileIdDownloadProfileGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | MDM Enrollment Profile identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDeviceEnrollmentProfileIdDownloadProfileGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-apple-aspen-config

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

