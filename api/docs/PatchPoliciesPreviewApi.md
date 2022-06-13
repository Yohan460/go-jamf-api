# \PatchPoliciesPreviewApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PatchPatchPoliciesGet**](PatchPoliciesPreviewApi.md#PatchPatchPoliciesGet) | **Get** /patch/patch-policies | Return a list of patch policies 
[**PatchPatchPoliciesIdDashboardDelete**](PatchPoliciesPreviewApi.md#PatchPatchPoliciesIdDashboardDelete) | **Delete** /patch/patch-policies/{id}/dashboard | Remove a patch policy from the dashboard 
[**PatchPatchPoliciesIdDashboardGet**](PatchPoliciesPreviewApi.md#PatchPatchPoliciesIdDashboardGet) | **Get** /patch/patch-policies/{id}/dashboard | Return whether or not the requested patch policy is on the dashboard 
[**PatchPatchPoliciesIdDashboardPost**](PatchPoliciesPreviewApi.md#PatchPatchPoliciesIdDashboardPost) | **Post** /patch/patch-policies/{id}/dashboard | Add a patch policy to the dashboard 



## PatchPatchPoliciesGet

> []PatchPolicySummary PatchPatchPoliciesGet(ctx).OnDashboard(onDashboard).Enabled(enabled).Execute()

Return a list of patch policies 



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
    onDashboard := true // bool | Filters whether or not the patch policies are on the dashboard. (optional) (default to false)
    enabled := true // bool | Filters whether or not the patch policies are enabled. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PatchPoliciesPreviewApi.PatchPatchPoliciesGet(context.Background()).OnDashboard(onDashboard).Enabled(enabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchPoliciesPreviewApi.PatchPatchPoliciesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPatchPoliciesGet`: []PatchPolicySummary
    fmt.Fprintf(os.Stdout, "Response from `PatchPoliciesPreviewApi.PatchPatchPoliciesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchPatchPoliciesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onDashboard** | **bool** | Filters whether or not the patch policies are on the dashboard. | [default to false]
 **enabled** | **bool** | Filters whether or not the patch policies are enabled. | [default to false]

### Return type

[**[]PatchPolicySummary**](PatchPolicySummary.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPatchPoliciesIdDashboardDelete

> PatchPatchPoliciesIdDashboardDelete(ctx, id).Execute()

Remove a patch policy from the dashboard 



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
    id := int32(56) // int32 | patch id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PatchPoliciesPreviewApi.PatchPatchPoliciesIdDashboardDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchPoliciesPreviewApi.PatchPatchPoliciesIdDashboardDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | patch id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPatchPoliciesIdDashboardDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPatchPoliciesIdDashboardGet

> PatchPolicyOnDashboard PatchPatchPoliciesIdDashboardGet(ctx, id).Execute()

Return whether or not the requested patch policy is on the dashboard 



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PatchPoliciesPreviewApi.PatchPatchPoliciesIdDashboardGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchPoliciesPreviewApi.PatchPatchPoliciesIdDashboardGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPatchPoliciesIdDashboardGet`: PatchPolicyOnDashboard
    fmt.Fprintf(os.Stdout, "Response from `PatchPoliciesPreviewApi.PatchPatchPoliciesIdDashboardGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | patch policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPatchPoliciesIdDashboardGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PatchPolicyOnDashboard**](PatchPolicyOnDashboard.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPatchPoliciesIdDashboardPost

> PatchPatchPoliciesIdDashboardPost(ctx, id).Execute()

Add a patch policy to the dashboard 



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PatchPoliciesPreviewApi.PatchPatchPoliciesIdDashboardPost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchPoliciesPreviewApi.PatchPatchPoliciesIdDashboardPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | patch policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPatchPoliciesIdDashboardPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

