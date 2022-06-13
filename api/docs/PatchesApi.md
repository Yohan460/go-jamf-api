# \PatchesApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PatchIdGet**](PatchesApi.md#PatchIdGet) | **Get** /patch/{id} | Return Active Patch Summary 
[**PatchIdPut**](PatchesApi.md#PatchIdPut) | **Put** /patch/{id} | Update patch report 
[**PatchIdVersionsGet**](PatchesApi.md#PatchIdVersionsGet) | **Get** /patch/{id}/versions | Return patch versions 
[**PatchObjPolicyIdGet**](PatchesApi.md#PatchObjPolicyIdGet) | **Get** /patch/obj/policy/{id} | Return Patch Policy Summary 
[**PatchSvcDisclaimerAgreePost**](PatchesApi.md#PatchSvcDisclaimerAgreePost) | **Post** /patch/svc/disclaimerAgree | Accept Patch reporting disclaimer 



## PatchIdGet

> ActivePatchSummary PatchIdGet(ctx, id).Execute()

Return Active Patch Summary 



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
    resp, r, err := apiClient.PatchesApi.PatchIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchesApi.PatchIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchIdGet`: ActivePatchSummary
    fmt.Fprintf(os.Stdout, "Response from `PatchesApi.PatchIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | patch id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActivePatchSummary**](ActivePatchSummary.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchIdPut

> ActivePatchSummary PatchIdPut(ctx, id).ActivePatchSummary(activePatchSummary).Execute()

Update patch report 



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
    activePatchSummary := *openapiclient.NewActivePatchSummary() // ActivePatchSummary | Active patch summary.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PatchesApi.PatchIdPut(context.Background(), id).ActivePatchSummary(activePatchSummary).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchesApi.PatchIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchIdPut`: ActivePatchSummary
    fmt.Fprintf(os.Stdout, "Response from `PatchesApi.PatchIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | patch id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activePatchSummary** | [**ActivePatchSummary**](ActivePatchSummary.md) | Active patch summary. | 

### Return type

[**ActivePatchSummary**](ActivePatchSummary.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchIdVersionsGet

> []PatchVersion PatchIdVersionsGet(ctx, id).Execute()

Return patch versions 



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
    resp, r, err := apiClient.PatchesApi.PatchIdVersionsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchesApi.PatchIdVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchIdVersionsGet`: []PatchVersion
    fmt.Fprintf(os.Stdout, "Response from `PatchesApi.PatchIdVersionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | patch id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchIdVersionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PatchVersion**](PatchVersion.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchObjPolicyIdGet

> PatchPolicySummary PatchObjPolicyIdGet(ctx, id).Execute()

Return Patch Policy Summary 



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
    resp, r, err := apiClient.PatchesApi.PatchObjPolicyIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchesApi.PatchObjPolicyIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchObjPolicyIdGet`: PatchPolicySummary
    fmt.Fprintf(os.Stdout, "Response from `PatchesApi.PatchObjPolicyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | patch policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchObjPolicyIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PatchPolicySummary**](PatchPolicySummary.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSvcDisclaimerAgreePost

> PatchSvcDisclaimerAgreePost(ctx).Execute()

Accept Patch reporting disclaimer 



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
    resp, r, err := apiClient.PatchesApi.PatchSvcDisclaimerAgreePost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchesApi.PatchSvcDisclaimerAgreePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSvcDisclaimerAgreePostRequest struct via the builder pattern


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

