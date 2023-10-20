# \PatchPolicyLogsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2PatchPoliciesIdLogsDeviceIdDetailsGet**](PatchPolicyLogsAPI.md#V2PatchPoliciesIdLogsDeviceIdDetailsGet) | **Get** /v2/patch-policies/{id}/logs/{deviceId}/details | Return attempt details for a specific log
[**V2PatchPoliciesIdLogsDeviceIdGet**](PatchPolicyLogsAPI.md#V2PatchPoliciesIdLogsDeviceIdGet) | **Get** /v2/patch-policies/{id}/logs/{deviceId} | Retrieves a single Patch Policy Log 
[**V2PatchPoliciesIdLogsEligibleRetryCountGet**](PatchPolicyLogsAPI.md#V2PatchPoliciesIdLogsEligibleRetryCountGet) | **Get** /v2/patch-policies/{id}/logs/eligible-retry-count | Return the count of the Patch Policy Logs for the patch policy id that are eligible for a retry attempt 
[**V2PatchPoliciesIdLogsGet**](PatchPolicyLogsAPI.md#V2PatchPoliciesIdLogsGet) | **Get** /v2/patch-policies/{id}/logs | Retrieve Patch Policy Logs 
[**V2PatchPoliciesIdLogsRetryAllPost**](PatchPolicyLogsAPI.md#V2PatchPoliciesIdLogsRetryAllPost) | **Post** /v2/patch-policies/{id}/logs/retry-all | Send retry attempts for all devices
[**V2PatchPoliciesIdLogsRetryPost**](PatchPolicyLogsAPI.md#V2PatchPoliciesIdLogsRetryPost) | **Post** /v2/patch-policies/{id}/logs/retry | Send retry attempts for specific devices



## V2PatchPoliciesIdLogsDeviceIdDetailsGet

> []PatchPolicyLogDetail V2PatchPoliciesIdLogsDeviceIdDetailsGet(ctx, id, deviceId).Execute()

Return attempt details for a specific log



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {
    id := "id_example" // string | patch policy id
    deviceId := "deviceId_example" // string | device id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PatchPolicyLogsAPI.V2PatchPoliciesIdLogsDeviceIdDetailsGet(context.Background(), id, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchPolicyLogsAPI.V2PatchPoliciesIdLogsDeviceIdDetailsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2PatchPoliciesIdLogsDeviceIdDetailsGet`: []PatchPolicyLogDetail
    fmt.Fprintf(os.Stdout, "Response from `PatchPolicyLogsAPI.V2PatchPoliciesIdLogsDeviceIdDetailsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | patch policy id | 
**deviceId** | **string** | device id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchPoliciesIdLogsDeviceIdDetailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]PatchPolicyLogDetail**](PatchPolicyLogDetail.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PatchPoliciesIdLogsDeviceIdGet

> PatchPolicyLogV2 V2PatchPoliciesIdLogsDeviceIdGet(ctx, id, deviceId).Execute()

Retrieves a single Patch Policy Log 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {
    id := "id_example" // string | patch policy id
    deviceId := "deviceId_example" // string | device id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PatchPolicyLogsAPI.V2PatchPoliciesIdLogsDeviceIdGet(context.Background(), id, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchPolicyLogsAPI.V2PatchPoliciesIdLogsDeviceIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2PatchPoliciesIdLogsDeviceIdGet`: PatchPolicyLogV2
    fmt.Fprintf(os.Stdout, "Response from `PatchPolicyLogsAPI.V2PatchPoliciesIdLogsDeviceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | patch policy id | 
**deviceId** | **string** | device id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchPoliciesIdLogsDeviceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PatchPolicyLogV2**](PatchPolicyLogV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PatchPoliciesIdLogsEligibleRetryCountGet

> PatchPolicyLogEligibleRetryCount V2PatchPoliciesIdLogsEligibleRetryCountGet(ctx, id).Execute()

Return the count of the Patch Policy Logs for the patch policy id that are eligible for a retry attempt 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {
    id := "id_example" // string | patch policy id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PatchPolicyLogsAPI.V2PatchPoliciesIdLogsEligibleRetryCountGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchPolicyLogsAPI.V2PatchPoliciesIdLogsEligibleRetryCountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2PatchPoliciesIdLogsEligibleRetryCountGet`: PatchPolicyLogEligibleRetryCount
    fmt.Fprintf(os.Stdout, "Response from `PatchPolicyLogsAPI.V2PatchPoliciesIdLogsEligibleRetryCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | patch policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchPoliciesIdLogsEligibleRetryCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PatchPolicyLogEligibleRetryCount**](PatchPolicyLogEligibleRetryCount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PatchPoliciesIdLogsGet

> PatchPolicyLogs V2PatchPoliciesIdLogsGet(ctx, id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Retrieve Patch Policy Logs 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {
    id := "id_example" // string | patch policy id
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is deviceName:asc. Multiple sort criteria are supported and must be separated with a comma. (optional) (default to ["deviceName:asc"])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter Patch Policy Logs collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: deviceId, deviceName, statusCode, statusDate, attemptNumber, ignoredForPatchPolicyId. This param can be combined with paging and sorting. (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PatchPolicyLogsAPI.V2PatchPoliciesIdLogsGet(context.Background(), id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchPolicyLogsAPI.V2PatchPoliciesIdLogsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2PatchPoliciesIdLogsGet`: PatchPolicyLogs
    fmt.Fprintf(os.Stdout, "Response from `PatchPolicyLogsAPI.V2PatchPoliciesIdLogsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | patch policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchPoliciesIdLogsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is deviceName:asc. Multiple sort criteria are supported and must be separated with a comma. | [default to [&quot;deviceName:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter Patch Policy Logs collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: deviceId, deviceName, statusCode, statusDate, attemptNumber, ignoredForPatchPolicyId. This param can be combined with paging and sorting. | [default to &quot;&quot;]

### Return type

[**PatchPolicyLogs**](PatchPolicyLogs.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PatchPoliciesIdLogsRetryAllPost

> V2PatchPoliciesIdLogsRetryAllPost(ctx, id).Execute()

Send retry attempts for all devices



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {
    id := "id_example" // string | patch policy id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PatchPolicyLogsAPI.V2PatchPoliciesIdLogsRetryAllPost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchPolicyLogsAPI.V2PatchPoliciesIdLogsRetryAllPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | patch policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchPoliciesIdLogsRetryAllPostRequest struct via the builder pattern


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


## V2PatchPoliciesIdLogsRetryPost

> V2PatchPoliciesIdLogsRetryPost(ctx, id).PatchPolicyLogRetry(patchPolicyLogRetry).Execute()

Send retry attempts for specific devices



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {
    id := "id_example" // string | patch policy id
    patchPolicyLogRetry := *openapiclient.NewPatchPolicyLogRetry() // PatchPolicyLogRetry | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PatchPolicyLogsAPI.V2PatchPoliciesIdLogsRetryPost(context.Background(), id).PatchPolicyLogRetry(patchPolicyLogRetry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchPolicyLogsAPI.V2PatchPoliciesIdLogsRetryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | patch policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchPoliciesIdLogsRetryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchPolicyLogRetry** | [**PatchPolicyLogRetry**](PatchPolicyLogRetry.md) |  | 

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

