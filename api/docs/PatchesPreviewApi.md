# \PatchesPreviewApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PatchDisclaimerAgreePost**](PatchesPreviewApi.md#PatchDisclaimerAgreePost) | **Post** /patch/disclaimerAgree | Accept Patch reporting disclaimer 
[**PatchObjIdGet**](PatchesPreviewApi.md#PatchObjIdGet) | **Get** /patch/obj/{id} | Return Active Patch Summary 
[**PatchObjIdPut**](PatchesPreviewApi.md#PatchObjIdPut) | **Put** /patch/obj/{id} | Update patch report 
[**PatchObjIdVersionsGet**](PatchesPreviewApi.md#PatchObjIdVersionsGet) | **Get** /patch/obj/{id}/versions | Return patch versions 
[**PatchObjPolicyIdLogsEligibleRetryCountGet**](PatchesPreviewApi.md#PatchObjPolicyIdLogsEligibleRetryCountGet) | **Get** /patch/obj/policy/{id}/logs/eligibleRetryCount | Return the count of the Patch Policy Logs for the policy is that are eligible for a retry attempt 
[**PatchObjPolicyIdSoftwareTitleConfigurationIdGet**](PatchesPreviewApi.md#PatchObjPolicyIdSoftwareTitleConfigurationIdGet) | **Get** /patch/obj/policy/{id}/softwareTitleConfigurationId | Return the Software Title Configuration Id for the given patch 
[**PatchObjSoftwareTitleConfigurationIdGet**](PatchesPreviewApi.md#PatchObjSoftwareTitleConfigurationIdGet) | **Get** /patch/obj/softwareTitleConfiguration/{id} | Return the Software Title Configuration 
[**PatchObjSoftwareTitleIdPoliciesGet**](PatchesPreviewApi.md#PatchObjSoftwareTitleIdPoliciesGet) | **Get** /patch/obj/softwareTitle/{id}/policies | Return the Summaries of the Patch Policies for the software title 
[**PatchObjsPolicyIdGet**](PatchesPreviewApi.md#PatchObjsPolicyIdGet) | **Get** /patch/objs/policy/{id} | Return Patch Policy Summary 
[**PatchOnDashboardGet**](PatchesPreviewApi.md#PatchOnDashboardGet) | **Get** /patch/onDashboard | Return list of Patch ids on dashboard 
[**PatchRetryPolicyPost**](PatchesPreviewApi.md#PatchRetryPolicyPost) | **Post** /patch/retryPolicy | Retry policy 
[**PatchSearchActivePatchHistoryPost**](PatchesPreviewApi.md#PatchSearchActivePatchHistoryPost) | **Post** /patch/searchActivePatchHistory | Search the history for a Specific Active Patch 
[**PatchSearchPatchPolicyLogsPost**](PatchesPreviewApi.md#PatchSearchPatchPolicyLogsPost) | **Post** /patch/searchPatchPolicyLogs | Return Patch Policy Logs 
[**PatchSvcRetryPolicyPost**](PatchesPreviewApi.md#PatchSvcRetryPolicyPost) | **Post** /patch/svc/retryPolicy | Retry policy 



## PatchDisclaimerAgreePost

> PatchDisclaimerAgreePost(ctx).Execute()

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
    resp, r, err := apiClient.PatchesPreviewApi.PatchDisclaimerAgreePost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchesPreviewApi.PatchDisclaimerAgreePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDisclaimerAgreePostRequest struct via the builder pattern


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


## PatchObjIdGet

> ActivePatchSummary PatchObjIdGet(ctx, id).Execute()

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
    resp, r, err := apiClient.PatchesPreviewApi.PatchObjIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchesPreviewApi.PatchObjIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchObjIdGet`: ActivePatchSummary
    fmt.Fprintf(os.Stdout, "Response from `PatchesPreviewApi.PatchObjIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | patch id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchObjIdGetRequest struct via the builder pattern


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


## PatchObjIdPut

> ActivePatchSummary PatchObjIdPut(ctx, id).ActivePatchSummary(activePatchSummary).Execute()

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
    resp, r, err := apiClient.PatchesPreviewApi.PatchObjIdPut(context.Background(), id).ActivePatchSummary(activePatchSummary).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchesPreviewApi.PatchObjIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchObjIdPut`: ActivePatchSummary
    fmt.Fprintf(os.Stdout, "Response from `PatchesPreviewApi.PatchObjIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | patch id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchObjIdPutRequest struct via the builder pattern


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


## PatchObjIdVersionsGet

> []PatchVersion PatchObjIdVersionsGet(ctx, id).Execute()

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
    resp, r, err := apiClient.PatchesPreviewApi.PatchObjIdVersionsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchesPreviewApi.PatchObjIdVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchObjIdVersionsGet`: []PatchVersion
    fmt.Fprintf(os.Stdout, "Response from `PatchesPreviewApi.PatchObjIdVersionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | patch id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchObjIdVersionsGetRequest struct via the builder pattern


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


## PatchObjPolicyIdLogsEligibleRetryCountGet

> int32 PatchObjPolicyIdLogsEligibleRetryCountGet(ctx, id).Execute()

Return the count of the Patch Policy Logs for the policy is that are eligible for a retry attempt 



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
    id := int32(56) // int32 | policy id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PatchesPreviewApi.PatchObjPolicyIdLogsEligibleRetryCountGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchesPreviewApi.PatchObjPolicyIdLogsEligibleRetryCountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchObjPolicyIdLogsEligibleRetryCountGet`: int32
    fmt.Fprintf(os.Stdout, "Response from `PatchesPreviewApi.PatchObjPolicyIdLogsEligibleRetryCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchObjPolicyIdLogsEligibleRetryCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**int32**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchObjPolicyIdSoftwareTitleConfigurationIdGet

> int32 PatchObjPolicyIdSoftwareTitleConfigurationIdGet(ctx, id).Execute()

Return the Software Title Configuration Id for the given patch 



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
    id := int32(56) // int32 | policy id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PatchesPreviewApi.PatchObjPolicyIdSoftwareTitleConfigurationIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchesPreviewApi.PatchObjPolicyIdSoftwareTitleConfigurationIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchObjPolicyIdSoftwareTitleConfigurationIdGet`: int32
    fmt.Fprintf(os.Stdout, "Response from `PatchesPreviewApi.PatchObjPolicyIdSoftwareTitleConfigurationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchObjPolicyIdSoftwareTitleConfigurationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**int32**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchObjSoftwareTitleConfigurationIdGet

> SoftwareTitleConfiguration PatchObjSoftwareTitleConfigurationIdGet(ctx, id).Execute()

Return the Software Title Configuration 



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
    id := int32(56) // int32 | software title configuration id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PatchesPreviewApi.PatchObjSoftwareTitleConfigurationIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchesPreviewApi.PatchObjSoftwareTitleConfigurationIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchObjSoftwareTitleConfigurationIdGet`: SoftwareTitleConfiguration
    fmt.Fprintf(os.Stdout, "Response from `PatchesPreviewApi.PatchObjSoftwareTitleConfigurationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | software title configuration id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchObjSoftwareTitleConfigurationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SoftwareTitleConfiguration**](SoftwareTitleConfiguration.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchObjSoftwareTitleIdPoliciesGet

> SoftwareTitlePatchPolicySummaries PatchObjSoftwareTitleIdPoliciesGet(ctx, id).Execute()

Return the Summaries of the Patch Policies for the software title 



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
    id := int32(56) // int32 | software title id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PatchesPreviewApi.PatchObjSoftwareTitleIdPoliciesGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchesPreviewApi.PatchObjSoftwareTitleIdPoliciesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchObjSoftwareTitleIdPoliciesGet`: SoftwareTitlePatchPolicySummaries
    fmt.Fprintf(os.Stdout, "Response from `PatchesPreviewApi.PatchObjSoftwareTitleIdPoliciesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | software title id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchObjSoftwareTitleIdPoliciesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SoftwareTitlePatchPolicySummaries**](SoftwareTitlePatchPolicySummaries.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchObjsPolicyIdGet

> PatchPolicySummary PatchObjsPolicyIdGet(ctx, id).Execute()

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
    resp, r, err := apiClient.PatchesPreviewApi.PatchObjsPolicyIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchesPreviewApi.PatchObjsPolicyIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchObjsPolicyIdGet`: PatchPolicySummary
    fmt.Fprintf(os.Stdout, "Response from `PatchesPreviewApi.PatchObjsPolicyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | patch policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchObjsPolicyIdGetRequest struct via the builder pattern


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


## PatchOnDashboardGet

> []int32 PatchOnDashboardGet(ctx).Execute()

Return list of Patch ids on dashboard 



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
    resp, r, err := apiClient.PatchesPreviewApi.PatchOnDashboardGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchesPreviewApi.PatchOnDashboardGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchOnDashboardGet`: []int32
    fmt.Fprintf(os.Stdout, "Response from `PatchesPreviewApi.PatchOnDashboardGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPatchOnDashboardGetRequest struct via the builder pattern


### Return type

**[]int32**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRetryPolicyPost

> PatchRetryPolicyPost(ctx).RetryPatchPolicyParams(retryPatchPolicyParams).Execute()

Retry policy 



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
    retryPatchPolicyParams := *openapiclient.NewRetryPatchPolicyParams() // RetryPatchPolicyParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PatchesPreviewApi.PatchRetryPolicyPost(context.Background()).RetryPatchPolicyParams(retryPatchPolicyParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchesPreviewApi.PatchRetryPolicyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchRetryPolicyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **retryPatchPolicyParams** | [**RetryPatchPolicyParams**](RetryPatchPolicyParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSearchActivePatchHistoryPost

> ActivePatchHistorySearchResults PatchSearchActivePatchHistoryPost(ctx).SearchActivePatchHistoryParams(searchActivePatchHistoryParams).Execute()

Search the history for a Specific Active Patch 



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
    searchActivePatchHistoryParams := *openapiclient.NewSearchActivePatchHistoryParams() // SearchActivePatchHistoryParams | Parameters for search (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PatchesPreviewApi.PatchSearchActivePatchHistoryPost(context.Background()).SearchActivePatchHistoryParams(searchActivePatchHistoryParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchesPreviewApi.PatchSearchActivePatchHistoryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSearchActivePatchHistoryPost`: ActivePatchHistorySearchResults
    fmt.Fprintf(os.Stdout, "Response from `PatchesPreviewApi.PatchSearchActivePatchHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchSearchActivePatchHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchActivePatchHistoryParams** | [**SearchActivePatchHistoryParams**](SearchActivePatchHistoryParams.md) | Parameters for search | 

### Return type

[**ActivePatchHistorySearchResults**](ActivePatchHistorySearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSearchPatchPolicyLogsPost

> PatchPolicyLogSearchResults PatchSearchPatchPolicyLogsPost(ctx).SearchPatchPolicyLogParams(searchPatchPolicyLogParams).Execute()

Return Patch Policy Logs 



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
    searchPatchPolicyLogParams := *openapiclient.NewSearchPatchPolicyLogParams() // SearchPatchPolicyLogParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PatchesPreviewApi.PatchSearchPatchPolicyLogsPost(context.Background()).SearchPatchPolicyLogParams(searchPatchPolicyLogParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchesPreviewApi.PatchSearchPatchPolicyLogsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSearchPatchPolicyLogsPost`: PatchPolicyLogSearchResults
    fmt.Fprintf(os.Stdout, "Response from `PatchesPreviewApi.PatchSearchPatchPolicyLogsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchSearchPatchPolicyLogsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchPatchPolicyLogParams** | [**SearchPatchPolicyLogParams**](SearchPatchPolicyLogParams.md) |  | 

### Return type

[**PatchPolicyLogSearchResults**](PatchPolicyLogSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSvcRetryPolicyPost

> PatchSvcRetryPolicyPost(ctx).RetryPatchPolicyParams(retryPatchPolicyParams).Execute()

Retry policy 



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
    retryPatchPolicyParams := *openapiclient.NewRetryPatchPolicyParams() // RetryPatchPolicyParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PatchesPreviewApi.PatchSvcRetryPolicyPost(context.Background()).RetryPatchPolicyParams(retryPatchPolicyParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchesPreviewApi.PatchSvcRetryPolicyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchSvcRetryPolicyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **retryPatchPolicyParams** | [**RetryPatchPolicyParams**](RetryPatchPolicyParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

