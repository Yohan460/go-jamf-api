# \PatchPoliciesAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2PatchPoliciesGet**](PatchPoliciesAPI.md#V2PatchPoliciesGet) | **Get** /v2/patch-policies | Retrieve Patch Policies
[**V2PatchPoliciesIdDashboardDelete**](PatchPoliciesAPI.md#V2PatchPoliciesIdDashboardDelete) | **Delete** /v2/patch-policies/{id}/dashboard | Remove a patch policy from the dashboard 
[**V2PatchPoliciesIdDashboardGet**](PatchPoliciesAPI.md#V2PatchPoliciesIdDashboardGet) | **Get** /v2/patch-policies/{id}/dashboard | Return whether or not the requested patch policy is on the dashboard 
[**V2PatchPoliciesIdDashboardPost**](PatchPoliciesAPI.md#V2PatchPoliciesIdDashboardPost) | **Post** /v2/patch-policies/{id}/dashboard | Add a patch policy to the dashboard 
[**V2PatchPoliciesPolicyDetailsGet**](PatchPoliciesAPI.md#V2PatchPoliciesPolicyDetailsGet) | **Get** /v2/patch-policies/policy-details | Retrieve Patch Policies



## V2PatchPoliciesGet

> PatchPolicies V2PatchPoliciesGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Retrieve Patch Policies



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
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. (optional) (default to ["id:asc"])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter Patch Policy collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, policyName, policyEnabled, policyTargetVersion, policyDeploymentMethod, softwareTitle, softwareTitleConfigurationId, pending, completed, deferred, and failed. This param can be combined with paging and sorting. (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PatchPoliciesAPI.V2PatchPoliciesGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchPoliciesAPI.V2PatchPoliciesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2PatchPoliciesGet`: PatchPolicies
    fmt.Fprintf(os.Stdout, "Response from `PatchPoliciesAPI.V2PatchPoliciesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchPoliciesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. | [default to [&quot;id:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter Patch Policy collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, policyName, policyEnabled, policyTargetVersion, policyDeploymentMethod, softwareTitle, softwareTitleConfigurationId, pending, completed, deferred, and failed. This param can be combined with paging and sorting. | [default to &quot;&quot;]

### Return type

[**PatchPolicies**](PatchPolicies.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PatchPoliciesIdDashboardDelete

> V2PatchPoliciesIdDashboardDelete(ctx, id).Execute()

Remove a patch policy from the dashboard 



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
    r, err := apiClient.PatchPoliciesAPI.V2PatchPoliciesIdDashboardDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchPoliciesAPI.V2PatchPoliciesIdDashboardDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV2PatchPoliciesIdDashboardDeleteRequest struct via the builder pattern


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


## V2PatchPoliciesIdDashboardGet

> PatchPolicyV2OnDashboard V2PatchPoliciesIdDashboardGet(ctx, id).Execute()

Return whether or not the requested patch policy is on the dashboard 



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
    resp, r, err := apiClient.PatchPoliciesAPI.V2PatchPoliciesIdDashboardGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchPoliciesAPI.V2PatchPoliciesIdDashboardGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2PatchPoliciesIdDashboardGet`: PatchPolicyV2OnDashboard
    fmt.Fprintf(os.Stdout, "Response from `PatchPoliciesAPI.V2PatchPoliciesIdDashboardGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | patch policy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchPoliciesIdDashboardGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PatchPolicyV2OnDashboard**](PatchPolicyV2OnDashboard.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PatchPoliciesIdDashboardPost

> V2PatchPoliciesIdDashboardPost(ctx, id).Execute()

Add a patch policy to the dashboard 



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
    r, err := apiClient.PatchPoliciesAPI.V2PatchPoliciesIdDashboardPost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchPoliciesAPI.V2PatchPoliciesIdDashboardPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV2PatchPoliciesIdDashboardPostRequest struct via the builder pattern


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


## V2PatchPoliciesPolicyDetailsGet

> PatchPolicyDetails V2PatchPoliciesPolicyDetailsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Retrieve Patch Policies



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
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. (optional) (default to ["id:asc"])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter Patch Policy collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name, enabled, targetPatchVersion, deploymentMethod, softwareTitleId, softwareTitleConfigurationId, killAppsDelayMinutes, killAppsMessage, isDowngrade, isPatchUnknownVersion, notificationHeader, selfServiceEnforceDeadline, selfServiceDeadline, installButtonText, selfServiceDescription, iconId, reminderFrequency, reminderEnabled. This param can be combined with paging and sorting. (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PatchPoliciesAPI.V2PatchPoliciesPolicyDetailsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchPoliciesAPI.V2PatchPoliciesPolicyDetailsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2PatchPoliciesPolicyDetailsGet`: PatchPolicyDetails
    fmt.Fprintf(os.Stdout, "Response from `PatchPoliciesAPI.V2PatchPoliciesPolicyDetailsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchPoliciesPolicyDetailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. | [default to [&quot;id:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter Patch Policy collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name, enabled, targetPatchVersion, deploymentMethod, softwareTitleId, softwareTitleConfigurationId, killAppsDelayMinutes, killAppsMessage, isDowngrade, isPatchUnknownVersion, notificationHeader, selfServiceEnforceDeadline, selfServiceDeadline, installButtonText, selfServiceDescription, iconId, reminderFrequency, reminderEnabled. This param can be combined with paging and sorting. | [default to &quot;&quot;]

### Return type

[**PatchPolicyDetails**](PatchPolicyDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

