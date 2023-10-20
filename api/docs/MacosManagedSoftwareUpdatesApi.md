# \MacosManagedSoftwareUpdatesAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1MacosManagedSoftwareUpdatesAvailableUpdatesGet**](MacosManagedSoftwareUpdatesAPI.md#V1MacosManagedSoftwareUpdatesAvailableUpdatesGet) | **Get** /v1/macos-managed-software-updates/available-updates | Retrieve available MacOs Managed Software Updates
[**V1MacosManagedSoftwareUpdatesSendUpdatesPost**](MacosManagedSoftwareUpdatesAPI.md#V1MacosManagedSoftwareUpdatesSendUpdatesPost) | **Post** /v1/macos-managed-software-updates/send-updates | Send MacOs Managed Software Updates



## V1MacosManagedSoftwareUpdatesAvailableUpdatesGet

> AvailableUpdates V1MacosManagedSoftwareUpdatesAvailableUpdatesGet(ctx).Execute()

Retrieve available MacOs Managed Software Updates



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
    resp, r, err := apiClient.MacosManagedSoftwareUpdatesAPI.V1MacosManagedSoftwareUpdatesAvailableUpdatesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MacosManagedSoftwareUpdatesAPI.V1MacosManagedSoftwareUpdatesAvailableUpdatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1MacosManagedSoftwareUpdatesAvailableUpdatesGet`: AvailableUpdates
    fmt.Fprintf(os.Stdout, "Response from `MacosManagedSoftwareUpdatesAPI.V1MacosManagedSoftwareUpdatesAvailableUpdatesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1MacosManagedSoftwareUpdatesAvailableUpdatesGetRequest struct via the builder pattern


### Return type

[**AvailableUpdates**](AvailableUpdates.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MacosManagedSoftwareUpdatesSendUpdatesPost

> MacOsManagedSoftwareUpdateResponse V1MacosManagedSoftwareUpdatesSendUpdatesPost(ctx).MacOsManagedSoftwareUpdate(macOsManagedSoftwareUpdate).Execute()

Send MacOs Managed Software Updates



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
    macOsManagedSoftwareUpdate := *openapiclient.NewMacOsManagedSoftwareUpdate() // MacOsManagedSoftwareUpdate | MacOs Managed Software Update to send

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MacosManagedSoftwareUpdatesAPI.V1MacosManagedSoftwareUpdatesSendUpdatesPost(context.Background()).MacOsManagedSoftwareUpdate(macOsManagedSoftwareUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MacosManagedSoftwareUpdatesAPI.V1MacosManagedSoftwareUpdatesSendUpdatesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1MacosManagedSoftwareUpdatesSendUpdatesPost`: MacOsManagedSoftwareUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `MacosManagedSoftwareUpdatesAPI.V1MacosManagedSoftwareUpdatesSendUpdatesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MacosManagedSoftwareUpdatesSendUpdatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **macOsManagedSoftwareUpdate** | [**MacOsManagedSoftwareUpdate**](MacOsManagedSoftwareUpdate.md) | MacOs Managed Software Update to send | 

### Return type

[**MacOsManagedSoftwareUpdateResponse**](MacOsManagedSoftwareUpdateResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

