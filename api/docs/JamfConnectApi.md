# \JamfConnectAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1JamfConnectConfigProfilesGet**](JamfConnectAPI.md#V1JamfConnectConfigProfilesGet) | **Get** /v1/jamf-connect/config-profiles | Search for config profiles linked to Jamf Connect 
[**V1JamfConnectConfigProfilesIdPut**](JamfConnectAPI.md#V1JamfConnectConfigProfilesIdPut) | **Put** /v1/jamf-connect/config-profiles/{id} | Update the way the Jamf Connect app gets updated on computers within scope of the associated configuration profile. 
[**V1JamfConnectDeploymentsIdTasksGet**](JamfConnectAPI.md#V1JamfConnectDeploymentsIdTasksGet) | **Get** /v1/jamf-connect/deployments/{id}/tasks | Search for deployment tasks for a config profile linked to Jamf Connect 
[**V1JamfConnectDeploymentsIdTasksRetryPost**](JamfConnectAPI.md#V1JamfConnectDeploymentsIdTasksRetryPost) | **Post** /v1/jamf-connect/deployments/{id}/tasks/retry | Request a retry of Connect install tasks 
[**V1JamfConnectGet**](JamfConnectAPI.md#V1JamfConnectGet) | **Get** /v1/jamf-connect | Get the Jamf Connect settings that you have access to see 
[**V1JamfConnectHistoryGet**](JamfConnectAPI.md#V1JamfConnectHistoryGet) | **Get** /v1/jamf-connect/history | Get Jamf Connect history 
[**V1JamfConnectHistoryPost**](JamfConnectAPI.md#V1JamfConnectHistoryPost) | **Post** /v1/jamf-connect/history | Add Jamf Connect history notes 



## V1JamfConnectConfigProfilesGet

> LinkedConnectProfileSearchResults V1JamfConnectConfigProfilesGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Search for config profiles linked to Jamf Connect 



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
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is not duplicated for each sort criterion, e.g., ...&sort=name:asc,date:desc. Fields that can be sorted: status, updated (optional) (default to [])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfConnectAPI.V1JamfConnectConfigProfilesGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfConnectAPI.V1JamfConnectConfigProfilesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1JamfConnectConfigProfilesGet`: LinkedConnectProfileSearchResults
	fmt.Fprintf(os.Stdout, "Response from `JamfConnectAPI.V1JamfConnectConfigProfilesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfConnectConfigProfilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the &#39;sort&#39; query param is not duplicated for each sort criterion, e.g., ...&amp;sort&#x3D;name:asc,date:desc. Fields that can be sorted: status, updated | [default to []]
 **filter** | **string** | Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15 | [default to &quot;&quot;]

### Return type

[**LinkedConnectProfileSearchResults**](LinkedConnectProfileSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JamfConnectConfigProfilesIdPut

> LinkedConnectProfile V1JamfConnectConfigProfilesIdPut(ctx, id).LinkedConnectProfile(linkedConnectProfile).Execute()

Update the way the Jamf Connect app gets updated on computers within scope of the associated configuration profile. 



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
	id := "24a7bb2a-9871-4895-9009-d1be07ed31b1" // string | the UUID of the profile to update
	linkedConnectProfile := *openapiclient.NewLinkedConnectProfile() // LinkedConnectProfile | Updatable Jamf Connect Settings (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfConnectAPI.V1JamfConnectConfigProfilesIdPut(context.Background(), id).LinkedConnectProfile(linkedConnectProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfConnectAPI.V1JamfConnectConfigProfilesIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1JamfConnectConfigProfilesIdPut`: LinkedConnectProfile
	fmt.Fprintf(os.Stdout, "Response from `JamfConnectAPI.V1JamfConnectConfigProfilesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the UUID of the profile to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfConnectConfigProfilesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **linkedConnectProfile** | [**LinkedConnectProfile**](LinkedConnectProfile.md) | Updatable Jamf Connect Settings | 

### Return type

[**LinkedConnectProfile**](LinkedConnectProfile.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JamfConnectDeploymentsIdTasksGet

> DeploymentTaskSearchResults V1JamfConnectDeploymentsIdTasksGet(ctx, id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Search for deployment tasks for a config profile linked to Jamf Connect 



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
	id := "24a7bb2a-9871-4895-9009-d1be07ed31b1" // string | the UUID of the Jamf Connect deployment
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is not duplicated for each sort criterion, e.g., ...&sort=name:asc,date:desc. Fields that can be sorted: status, updated (optional) (default to [])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfConnectAPI.V1JamfConnectDeploymentsIdTasksGet(context.Background(), id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfConnectAPI.V1JamfConnectDeploymentsIdTasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1JamfConnectDeploymentsIdTasksGet`: DeploymentTaskSearchResults
	fmt.Fprintf(os.Stdout, "Response from `JamfConnectAPI.V1JamfConnectDeploymentsIdTasksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the UUID of the Jamf Connect deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfConnectDeploymentsIdTasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
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


## V1JamfConnectDeploymentsIdTasksRetryPost

> V1JamfConnectDeploymentsIdTasksRetryPost(ctx, id).Ids(ids).Execute()

Request a retry of Connect install tasks 



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
	id := "24a7bb2a-9871-4895-9009-d1be07ed31b1" // string | the UUID of the deployment associated with the retry
	ids := *openapiclient.NewIds() // Ids | task IDs to retry

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JamfConnectAPI.V1JamfConnectDeploymentsIdTasksRetryPost(context.Background(), id).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfConnectAPI.V1JamfConnectDeploymentsIdTasksRetryPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1JamfConnectDeploymentsIdTasksRetryPostRequest struct via the builder pattern


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


## V1JamfConnectGet

> V1JamfConnectGet(ctx).Execute()

Get the Jamf Connect settings that you have access to see 



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JamfConnectAPI.V1JamfConnectGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfConnectAPI.V1JamfConnectGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfConnectGetRequest struct via the builder pattern


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


## V1JamfConnectHistoryGet

> HistorySearchResults V1JamfConnectHistoryGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get Jamf Connect history 



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
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is not duplicated for each sort criterion, e.g., ...&sort=name:asc,date:desc. Fields that can be sorted: status, updated (optional) (default to [])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfConnectAPI.V1JamfConnectHistoryGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfConnectAPI.V1JamfConnectHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1JamfConnectHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `JamfConnectAPI.V1JamfConnectHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfConnectHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
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


## V1JamfConnectHistoryPost

> HrefResponse V1JamfConnectHistoryPost(ctx).ObjectHistoryNote(objectHistoryNote).Execute()

Add Jamf Connect history notes 



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
	objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | history notes to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfConnectAPI.V1JamfConnectHistoryPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfConnectAPI.V1JamfConnectHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1JamfConnectHistoryPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `JamfConnectAPI.V1JamfConnectHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfConnectHistoryPostRequest struct via the builder pattern


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

