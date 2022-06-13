# \JamfManagementFrameworkApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1JamfManagementFrameworkRedeployIdPost**](JamfManagementFrameworkApi.md#V1JamfManagementFrameworkRedeployIdPost) | **Post** /v1/jamf-management-framework/redeploy/{id} | Redeploy Jamf Management Framework 



## V1JamfManagementFrameworkRedeployIdPost

> RedeployJamfManagementFrameworkResponse V1JamfManagementFrameworkRedeployIdPost(ctx, id).Execute()

Redeploy Jamf Management Framework 



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
    id := "id_example" // string | instance id of computer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JamfManagementFrameworkApi.V1JamfManagementFrameworkRedeployIdPost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfManagementFrameworkApi.V1JamfManagementFrameworkRedeployIdPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1JamfManagementFrameworkRedeployIdPost`: RedeployJamfManagementFrameworkResponse
    fmt.Fprintf(os.Stdout, "Response from `JamfManagementFrameworkApi.V1JamfManagementFrameworkRedeployIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of computer | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfManagementFrameworkRedeployIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RedeployJamfManagementFrameworkResponse**](RedeployJamfManagementFrameworkResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

