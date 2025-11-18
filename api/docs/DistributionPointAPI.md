# \DistributionPointAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DistributionPointsDeleteMultiplePost**](DistributionPointAPI.md#V1DistributionPointsDeleteMultiplePost) | **Post** /v1/distribution-points/delete-multiple | Delete multiple distribution points at once
[**V1DistributionPointsGet**](DistributionPointAPI.md#V1DistributionPointsGet) | **Get** /v1/distribution-points | Finds all Distribution Points
[**V1DistributionPointsIdDelete**](DistributionPointAPI.md#V1DistributionPointsIdDelete) | **Delete** /v1/distribution-points/{id} | Remove specified distribution point 
[**V1DistributionPointsIdGet**](DistributionPointAPI.md#V1DistributionPointsIdGet) | **Get** /v1/distribution-points/{id} | Get specified distribution point 
[**V1DistributionPointsIdHistoryGet**](DistributionPointAPI.md#V1DistributionPointsIdHistoryGet) | **Get** /v1/distribution-points/{id}/history | Get specified distribution point History object 
[**V1DistributionPointsIdHistoryPost**](DistributionPointAPI.md#V1DistributionPointsIdHistoryPost) | **Post** /v1/distribution-points/{id}/history | Add specified distribution point History object notes 
[**V1DistributionPointsIdPatch**](DistributionPointAPI.md#V1DistributionPointsIdPatch) | **Patch** /v1/distribution-points/{id} | Update specified distribution point object 
[**V1DistributionPointsIdPut**](DistributionPointAPI.md#V1DistributionPointsIdPut) | **Put** /v1/distribution-points/{id} | Update specified distribution point object 
[**V1DistributionPointsPost**](DistributionPointAPI.md#V1DistributionPointsPost) | **Post** /v1/distribution-points | Create distribution point



## V1DistributionPointsDeleteMultiplePost

> V1DistributionPointsDeleteMultiplePost(ctx).Ids(ids).Execute()

Delete multiple distribution points at once



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
	ids := *openapiclient.NewIds() // Ids | ids of the distribution points to be deleted

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DistributionPointAPI.V1DistributionPointsDeleteMultiplePost(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionPointAPI.V1DistributionPointsDeleteMultiplePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DistributionPointsDeleteMultiplePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**Ids**](Ids.md) | ids of the distribution points to be deleted | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DistributionPointsGet

> DistributionPointSearchResults V1DistributionPointsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Finds all Distribution Points



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
	sort := []string{"Inner_example"} // []string | Sorts results by one or more criteria, following the format property:asc/desc. Default sort is id:asc. If using multiple criteria, separate with commas. Allows fields such as - name, serverName (optional) (default to {"id:asc"})
	filter := "filter_example" // string | Filters results. Use RSQL format for query. Allows fields such as - name, serverName, principal, fileSharingConnectionType, and httpsEnabled Can be combined with paging and sorting. Default filter is an empty query and returns all results from the requested page. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionPointAPI.V1DistributionPointsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionPointAPI.V1DistributionPointsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DistributionPointsGet`: DistributionPointSearchResults
	fmt.Fprintf(os.Stdout, "Response from `DistributionPointAPI.V1DistributionPointsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DistributionPointsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorts results by one or more criteria, following the format property:asc/desc. Default sort is id:asc. If using multiple criteria, separate with commas. Allows fields such as - name, serverName | [default to {&quot;id:asc&quot;}]
 **filter** | **string** | Filters results. Use RSQL format for query. Allows fields such as - name, serverName, principal, fileSharingConnectionType, and httpsEnabled Can be combined with paging and sorting. Default filter is an empty query and returns all results from the requested page. | [default to &quot;&quot;]

### Return type

[**DistributionPointSearchResults**](DistributionPointSearchResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DistributionPointsIdDelete

> V1DistributionPointsIdDelete(ctx, id).Execute()

Remove specified distribution point 



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
	id := "id_example" // string | Instance id of distribution point

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DistributionPointAPI.V1DistributionPointsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionPointAPI.V1DistributionPointsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Instance id of distribution point | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DistributionPointsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DistributionPointsIdGet

> DistributionPoint V1DistributionPointsIdGet(ctx, id).Execute()

Get specified distribution point 



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
	id := "id_example" // string | instance id of distribution point

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionPointAPI.V1DistributionPointsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionPointAPI.V1DistributionPointsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DistributionPointsIdGet`: DistributionPoint
	fmt.Fprintf(os.Stdout, "Response from `DistributionPointAPI.V1DistributionPointsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of distribution point | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DistributionPointsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DistributionPoint**](DistributionPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DistributionPointsIdHistoryGet

> HistorySearchResults V1DistributionPointsIdHistoryGet(ctx, id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get specified distribution point History object 



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
	id := "id_example" // string | Instance id of distribution point history
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorts results by one or more criteria, following the format property:asc/desc. Default sort is id:asc. If using multiple criteria, separate with commas. (optional) (default to {"date:desc"})
	filter := "filter_example" // string | Filters results. Use RSQL format for query. Allows for many fields, including id, name, etc. Can be combined with paging and sorting. Default filter is an empty query and returns all results from the requested page. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionPointAPI.V1DistributionPointsIdHistoryGet(context.Background(), id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionPointAPI.V1DistributionPointsIdHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DistributionPointsIdHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `DistributionPointAPI.V1DistributionPointsIdHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Instance id of distribution point history | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DistributionPointsIdHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorts results by one or more criteria, following the format property:asc/desc. Default sort is id:asc. If using multiple criteria, separate with commas. | [default to {&quot;date:desc&quot;}]
 **filter** | **string** | Filters results. Use RSQL format for query. Allows for many fields, including id, name, etc. Can be combined with paging and sorting. Default filter is an empty query and returns all results from the requested page. | [default to &quot;&quot;]

### Return type

[**HistorySearchResults**](HistorySearchResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DistributionPointsIdHistoryPost

> ObjectHistory V1DistributionPointsIdHistoryPost(ctx, id).ObjectHistoryNote(objectHistoryNote).Execute()

Add specified distribution point History object notes 



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
	id := "id_example" // string | Instance id of distribution point history
	objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | History note to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionPointAPI.V1DistributionPointsIdHistoryPost(context.Background(), id).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionPointAPI.V1DistributionPointsIdHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DistributionPointsIdHistoryPost`: ObjectHistory
	fmt.Fprintf(os.Stdout, "Response from `DistributionPointAPI.V1DistributionPointsIdHistoryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Instance id of distribution point history | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DistributionPointsIdHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectHistoryNote** | [**ObjectHistoryNote**](ObjectHistoryNote.md) | History note to be created | 

### Return type

[**ObjectHistory**](ObjectHistory.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DistributionPointsIdPatch

> DistributionPoint V1DistributionPointsIdPatch(ctx, id).DistributionPoint(distributionPoint).Execute()

Update specified distribution point object 



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
	id := "id_example" // string | Instance id of distribution point
	distributionPoint := *openapiclient.NewDistributionPoint("My distribution point", "My Server", "AFP") // DistributionPoint | distribution point object to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionPointAPI.V1DistributionPointsIdPatch(context.Background(), id).DistributionPoint(distributionPoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionPointAPI.V1DistributionPointsIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DistributionPointsIdPatch`: DistributionPoint
	fmt.Fprintf(os.Stdout, "Response from `DistributionPointAPI.V1DistributionPointsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Instance id of distribution point | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DistributionPointsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distributionPoint** | [**DistributionPoint**](DistributionPoint.md) | distribution point object to update. | 

### Return type

[**DistributionPoint**](DistributionPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DistributionPointsIdPut

> DistributionPoint V1DistributionPointsIdPut(ctx, id).DistributionPoint(distributionPoint).Execute()

Update specified distribution point object 



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
	id := "id_example" // string | Instance id of distribution point
	distributionPoint := *openapiclient.NewDistributionPoint("My distribution point", "My Server", "AFP") // DistributionPoint | distribution point object to update. ids defined in this body will be ignored

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionPointAPI.V1DistributionPointsIdPut(context.Background(), id).DistributionPoint(distributionPoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionPointAPI.V1DistributionPointsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DistributionPointsIdPut`: DistributionPoint
	fmt.Fprintf(os.Stdout, "Response from `DistributionPointAPI.V1DistributionPointsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Instance id of distribution point | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DistributionPointsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distributionPoint** | [**DistributionPoint**](DistributionPoint.md) | distribution point object to update. ids defined in this body will be ignored | 

### Return type

[**DistributionPoint**](DistributionPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DistributionPointsPost

> HrefResponse V1DistributionPointsPost(ctx).DistributionPoint(distributionPoint).Execute()

Create distribution point



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
	distributionPoint := *openapiclient.NewDistributionPoint("My distribution point", "My Server", "AFP") // DistributionPoint | distribution point to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistributionPointAPI.V1DistributionPointsPost(context.Background()).DistributionPoint(distributionPoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistributionPointAPI.V1DistributionPointsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DistributionPointsPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `DistributionPointAPI.V1DistributionPointsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DistributionPointsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **distributionPoint** | [**DistributionPoint**](DistributionPoint.md) | distribution point to be created | 

### Return type

[**HrefResponse**](HrefResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

