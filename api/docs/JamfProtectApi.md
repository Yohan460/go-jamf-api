# \JamfProtectApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1JamfProtectDelete**](JamfProtectApi.md#V1JamfProtectDelete) | **Delete** /v1/jamf-protect | Delete Jamf Protect API registration.
[**V1JamfProtectDeploymentsIdTasksGet**](JamfProtectApi.md#V1JamfProtectDeploymentsIdTasksGet) | **Get** /v1/jamf-protect/deployments/{id}/tasks | Search for deployment tasks for a config profile linked to Jamf Protect 
[**V1JamfProtectDeploymentsIdTasksRetryPost**](JamfProtectApi.md#V1JamfProtectDeploymentsIdTasksRetryPost) | **Post** /v1/jamf-protect/deployments/{id}/tasks/retry | Request a retry of Protect install tasks 
[**V1JamfProtectGet**](JamfProtectApi.md#V1JamfProtectGet) | **Get** /v1/jamf-protect | Jamf Protect integration settings
[**V1JamfProtectHistoryGet**](JamfProtectApi.md#V1JamfProtectHistoryGet) | **Get** /v1/jamf-protect/history | Get Jamf Protect history 
[**V1JamfProtectHistoryPost**](JamfProtectApi.md#V1JamfProtectHistoryPost) | **Post** /v1/jamf-protect/history | Add Jamf Protect history notes 
[**V1JamfProtectPlansGet**](JamfProtectApi.md#V1JamfProtectPlansGet) | **Get** /v1/jamf-protect/plans | Get all of the previously synced Jamf Protect Plans with information about their associated configuration profile
[**V1JamfProtectPlansSyncPost**](JamfProtectApi.md#V1JamfProtectPlansSyncPost) | **Post** /v1/jamf-protect/plans/sync | Sync Plans with Jamf Protect
[**V1JamfProtectPut**](JamfProtectApi.md#V1JamfProtectPut) | **Put** /v1/jamf-protect | Jamf Protect integration settings
[**V1JamfProtectRegisterPost**](JamfProtectApi.md#V1JamfProtectRegisterPost) | **Post** /v1/jamf-protect/register | Register a Jamf Protect API configuration with Jamf Pro



## V1JamfProtectDelete

> V1JamfProtectDelete(ctx).Execute()

Delete Jamf Protect API registration.



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
    resp, r, err := apiClient.JamfProtectApi.V1JamfProtectDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProtectApi.V1JamfProtectDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfProtectDeleteRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JamfProtectDeploymentsIdTasksGet

> DeploymentTaskSearchResults V1JamfProtectDeploymentsIdTasksGet(ctx, id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Search for deployment tasks for a config profile linked to Jamf Protect 



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
    id := "24a7bb2a-9871-4895-9009-d1be07ed31b1" // string | the UUID of the Jamf Protect deployment
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is not duplicated for each sort criterion, e.g., ...&sort=name:asc,date:desc. Fields that can be sorted: status, updated (optional) (default to [])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JamfProtectApi.V1JamfProtectDeploymentsIdTasksGet(context.Background(), id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProtectApi.V1JamfProtectDeploymentsIdTasksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1JamfProtectDeploymentsIdTasksGet`: DeploymentTaskSearchResults
    fmt.Fprintf(os.Stdout, "Response from `JamfProtectApi.V1JamfProtectDeploymentsIdTasksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the UUID of the Jamf Protect deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfProtectDeploymentsIdTasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the &#39;sort&#39; query param is not duplicated for each sort criterion, e.g., ...&amp;sort&#x3D;name:asc,date:desc. Fields that can be sorted: status, updated | [default to []]
 **filter** | **string** | Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15 | [default to &quot;&quot;]

### Return type

[**DeploymentTaskSearchResults**](DeploymentTaskSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JamfProtectDeploymentsIdTasksRetryPost

> V1JamfProtectDeploymentsIdTasksRetryPost(ctx, id).Ids(ids).Execute()

Request a retry of Protect install tasks 



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
    id := "24a7bb2a-9871-4895-9009-d1be07ed31b1" // string | the UUID of the deployment associated with the retry
    ids := *openapiclient.NewIds() // Ids | task IDs to retry

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JamfProtectApi.V1JamfProtectDeploymentsIdTasksRetryPost(context.Background(), id).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProtectApi.V1JamfProtectDeploymentsIdTasksRetryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the UUID of the deployment associated with the retry | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfProtectDeploymentsIdTasksRetryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ids** | [**Ids**](Ids.md) | task IDs to retry | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JamfProtectGet

> ProtectSettingsResponse V1JamfProtectGet(ctx).Execute()

Jamf Protect integration settings



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
    resp, r, err := apiClient.JamfProtectApi.V1JamfProtectGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProtectApi.V1JamfProtectGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1JamfProtectGet`: ProtectSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `JamfProtectApi.V1JamfProtectGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfProtectGetRequest struct via the builder pattern


### Return type

[**ProtectSettingsResponse**](ProtectSettingsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JamfProtectHistoryGet

> HistorySearchResults V1JamfProtectHistoryGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get Jamf Protect history 



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
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is not duplicated for each sort criterion, e.g., ...&sort=name:asc,date:desc. Fields that can be sorted: status, updated (optional) (default to [])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JamfProtectApi.V1JamfProtectHistoryGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProtectApi.V1JamfProtectHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1JamfProtectHistoryGet`: HistorySearchResults
    fmt.Fprintf(os.Stdout, "Response from `JamfProtectApi.V1JamfProtectHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfProtectHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the &#39;sort&#39; query param is not duplicated for each sort criterion, e.g., ...&amp;sort&#x3D;name:asc,date:desc. Fields that can be sorted: status, updated | [default to []]
 **filter** | **string** | Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15 | [default to &quot;&quot;]

### Return type

[**HistorySearchResults**](HistorySearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JamfProtectHistoryPost

> HrefResponse V1JamfProtectHistoryPost(ctx).ObjectHistoryNote(objectHistoryNote).Execute()

Add Jamf Protect history notes 



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
    objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | history notes to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JamfProtectApi.V1JamfProtectHistoryPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProtectApi.V1JamfProtectHistoryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1JamfProtectHistoryPost`: HrefResponse
    fmt.Fprintf(os.Stdout, "Response from `JamfProtectApi.V1JamfProtectHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfProtectHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectHistoryNote** | [**ObjectHistoryNote**](ObjectHistoryNote.md) | history notes to create | 

### Return type

[**HrefResponse**](HrefResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JamfProtectPlansGet

> PlanSearchResults V1JamfProtectPlansGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get all of the previously synced Jamf Protect Plans with information about their associated configuration profile



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
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is not duplicated for each sort criterion, e.g., ...&sort=name:asc,date:desc. Fields that can be sorted: status, updated (optional) (default to [])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JamfProtectApi.V1JamfProtectPlansGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProtectApi.V1JamfProtectPlansGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1JamfProtectPlansGet`: PlanSearchResults
    fmt.Fprintf(os.Stdout, "Response from `JamfProtectApi.V1JamfProtectPlansGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfProtectPlansGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the &#39;sort&#39; query param is not duplicated for each sort criterion, e.g., ...&amp;sort&#x3D;name:asc,date:desc. Fields that can be sorted: status, updated | [default to []]
 **filter** | **string** | Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15 | [default to &quot;&quot;]

### Return type

[**PlanSearchResults**](PlanSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JamfProtectPlansSyncPost

> V1JamfProtectPlansSyncPost(ctx).Execute()

Sync Plans with Jamf Protect



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
    resp, r, err := apiClient.JamfProtectApi.V1JamfProtectPlansSyncPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProtectApi.V1JamfProtectPlansSyncPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfProtectPlansSyncPostRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JamfProtectPut

> ProtectSettingsResponse V1JamfProtectPut(ctx).ProtectUpdatableSettingsRequest(protectUpdatableSettingsRequest).Execute()

Jamf Protect integration settings



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
    protectUpdatableSettingsRequest := *openapiclient.NewProtectUpdatableSettingsRequest() // ProtectUpdatableSettingsRequest | Updatable Jamf Protect Settings

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JamfProtectApi.V1JamfProtectPut(context.Background()).ProtectUpdatableSettingsRequest(protectUpdatableSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProtectApi.V1JamfProtectPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1JamfProtectPut`: ProtectSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `JamfProtectApi.V1JamfProtectPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfProtectPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **protectUpdatableSettingsRequest** | [**ProtectUpdatableSettingsRequest**](ProtectUpdatableSettingsRequest.md) | Updatable Jamf Protect Settings | 

### Return type

[**ProtectSettingsResponse**](ProtectSettingsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JamfProtectRegisterPost

> ProtectSettingsResponse V1JamfProtectRegisterPost(ctx).ProtectRegistrationRequest(protectRegistrationRequest).Execute()

Register a Jamf Protect API configuration with Jamf Pro



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
    protectRegistrationRequest := *openapiclient.NewProtectRegistrationRequest("https://examplejamfprotect.jamfcloud.com/graphql", "uzPJXlArmzTAmPRQtZEnQ2OFtNw8qQV", "7fyP6BphUUQ5B_zoLrkYhM5j1HTcf-4PxshettZbK0ZcnzV57gyHwF23U3F96F") // ProtectRegistrationRequest | Jamf Protect API connection information

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JamfProtectApi.V1JamfProtectRegisterPost(context.Background()).ProtectRegistrationRequest(protectRegistrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProtectApi.V1JamfProtectRegisterPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1JamfProtectRegisterPost`: ProtectSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `JamfProtectApi.V1JamfProtectRegisterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfProtectRegisterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **protectRegistrationRequest** | [**ProtectRegistrationRequest**](ProtectRegistrationRequest.md) | Jamf Protect API connection information | 

### Return type

[**ProtectSettingsResponse**](ProtectSettingsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

