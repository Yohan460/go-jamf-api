# \PatchPolicyLogsPreviewApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PatchPatchPoliciesIdLogsGet**](PatchPolicyLogsPreviewApi.md#PatchPatchPoliciesIdLogsGet) | **Get** /patch/patch-policies/{id}/logs | Return the Patch Policy Attempt details 



## PatchPatchPoliciesIdLogsGet

> []PatchPolicyAttempt PatchPatchPoliciesIdLogsGet(ctx, id).DeviceId(deviceId).Execute()

Return the Patch Policy Attempt details 



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
    id := int32(56) // int32 | patch policy id
    deviceId := int32(56) // int32 | device id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PatchPolicyLogsPreviewApi.PatchPatchPoliciesIdLogsGet(context.Background(), id).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchPolicyLogsPreviewApi.PatchPatchPoliciesIdLogsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPatchPoliciesIdLogsGet`: []PatchPolicyAttempt
    fmt.Fprintf(os.Stdout, "Response from `PatchPolicyLogsPreviewApi.PatchPatchPoliciesIdLogsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | patch policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPatchPoliciesIdLogsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceId** | **int32** | device id | 

### Return type

[**[]PatchPolicyAttempt**](PatchPolicyAttempt.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

