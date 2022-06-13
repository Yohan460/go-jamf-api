# \SmartComputerGroupsPreviewApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ComputersIdRecalculateSmartGroupsPost**](SmartComputerGroupsPreviewApi.md#V1ComputersIdRecalculateSmartGroupsPost) | **Post** /v1/computers/{id}/recalculate-smart-groups | Recalculate a smart group for the given id 
[**V1SmartComputerGroupsIdRecalculatePost**](SmartComputerGroupsPreviewApi.md#V1SmartComputerGroupsIdRecalculatePost) | **Post** /v1/smart-computer-groups/{id}/recalculate | Recalculate the smart group for the given id 



## V1ComputersIdRecalculateSmartGroupsPost

> RecalculationResults V1ComputersIdRecalculateSmartGroupsPost(ctx, id).Execute()

Recalculate a smart group for the given id 



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
    id := int32(56) // int32 | id of computer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmartComputerGroupsPreviewApi.V1ComputersIdRecalculateSmartGroupsPost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartComputerGroupsPreviewApi.V1ComputersIdRecalculateSmartGroupsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ComputersIdRecalculateSmartGroupsPost`: RecalculationResults
    fmt.Fprintf(os.Stdout, "Response from `SmartComputerGroupsPreviewApi.V1ComputersIdRecalculateSmartGroupsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id of computer | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputersIdRecalculateSmartGroupsPostRequest struct via the builder pattern


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


## V1SmartComputerGroupsIdRecalculatePost

> RecalculationResults V1SmartComputerGroupsIdRecalculatePost(ctx, id).Execute()

Recalculate the smart group for the given id 



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
    resp, r, err := apiClient.SmartComputerGroupsPreviewApi.V1SmartComputerGroupsIdRecalculatePost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartComputerGroupsPreviewApi.V1SmartComputerGroupsIdRecalculatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SmartComputerGroupsIdRecalculatePost`: RecalculationResults
    fmt.Fprintf(os.Stdout, "Response from `SmartComputerGroupsPreviewApi.V1SmartComputerGroupsIdRecalculatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | instance id of smart group | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SmartComputerGroupsIdRecalculatePostRequest struct via the builder pattern


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

