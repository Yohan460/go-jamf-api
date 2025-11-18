# \CloudDistributionPointAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CloudDistributionPointDelete**](CloudDistributionPointAPI.md#V1CloudDistributionPointDelete) | **Delete** /v1/cloud-distribution-point | Delete cloud distribution point.
[**V1CloudDistributionPointFilesGet**](CloudDistributionPointAPI.md#V1CloudDistributionPointFilesGet) | **Get** /v1/cloud-distribution-point/files | Get the cloud distribution point Inventory files details
[**V1CloudDistributionPointGet**](CloudDistributionPointAPI.md#V1CloudDistributionPointGet) | **Get** /v1/cloud-distribution-point | Get the cloud distribution point Details. 
[**V1CloudDistributionPointHistoryGet**](CloudDistributionPointAPI.md#V1CloudDistributionPointHistoryGet) | **Get** /v1/cloud-distribution-point/history | Get cloud distribution point history details
[**V1CloudDistributionPointHistoryPost**](CloudDistributionPointAPI.md#V1CloudDistributionPointHistoryPost) | **Post** /v1/cloud-distribution-point/history | Add specified cloud distribution point history object notes
[**V1CloudDistributionPointPatch**](CloudDistributionPointAPI.md#V1CloudDistributionPointPatch) | **Patch** /v1/cloud-distribution-point | Update specific fields on a cloud distribution point
[**V1CloudDistributionPointPost**](CloudDistributionPointAPI.md#V1CloudDistributionPointPost) | **Post** /v1/cloud-distribution-point | Create cloud distribution point
[**V1CloudDistributionPointTestConnectionGet**](CloudDistributionPointAPI.md#V1CloudDistributionPointTestConnectionGet) | **Get** /v1/cloud-distribution-point/test-connection | Get the cloud distribution point test connection details. 
[**V1CloudDistributionPointUploadCapabilityGet**](CloudDistributionPointAPI.md#V1CloudDistributionPointUploadCapabilityGet) | **Get** /v1/cloud-distribution-point/upload-capability | Finds specific information for the currently configured cloud distribution point. 



## V1CloudDistributionPointDelete

> V1CloudDistributionPointDelete(ctx).Execute()

Delete cloud distribution point.



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
	r, err := apiClient.CloudDistributionPointAPI.V1CloudDistributionPointDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudDistributionPointAPI.V1CloudDistributionPointDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudDistributionPointDeleteRequest struct via the builder pattern


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


## V1CloudDistributionPointFilesGet

> CloudDistributionPointInventoryFilesResults V1CloudDistributionPointFilesGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get the cloud distribution point Inventory files details



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
	sort := []string{"Inner_example"} // []string | Sorts results by one or more criteria, following the format property:asc/desc.<br/> Default sort is id:asc.<br/> If using multiple criteria, separate with commas. Allows sort for id, fileName and type etc. (optional) (default to {"id.asc"})
	filter := "filter_example" // string | Filters results. Use RSQL format for query. Allows for many fields, including fileName and type<br/> Can be combined with paging and sorting.<br/> Fields allowed in the query: fileName and type <br/> Default filter is an empty query and returns all results from the requested page. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudDistributionPointAPI.V1CloudDistributionPointFilesGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudDistributionPointAPI.V1CloudDistributionPointFilesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CloudDistributionPointFilesGet`: CloudDistributionPointInventoryFilesResults
	fmt.Fprintf(os.Stdout, "Response from `CloudDistributionPointAPI.V1CloudDistributionPointFilesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudDistributionPointFilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorts results by one or more criteria, following the format property:asc/desc.&lt;br/&gt; Default sort is id:asc.&lt;br/&gt; If using multiple criteria, separate with commas. Allows sort for id, fileName and type etc. | [default to {&quot;id.asc&quot;}]
 **filter** | **string** | Filters results. Use RSQL format for query. Allows for many fields, including fileName and type&lt;br/&gt; Can be combined with paging and sorting.&lt;br/&gt; Fields allowed in the query: fileName and type &lt;br/&gt; Default filter is an empty query and returns all results from the requested page. | [default to &quot;&quot;]

### Return type

[**CloudDistributionPointInventoryFilesResults**](CloudDistributionPointInventoryFilesResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudDistributionPointGet

> CloudDistributionPoint V1CloudDistributionPointGet(ctx).Execute()

Get the cloud distribution point Details. 



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
	resp, r, err := apiClient.CloudDistributionPointAPI.V1CloudDistributionPointGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudDistributionPointAPI.V1CloudDistributionPointGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CloudDistributionPointGet`: CloudDistributionPoint
	fmt.Fprintf(os.Stdout, "Response from `CloudDistributionPointAPI.V1CloudDistributionPointGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudDistributionPointGetRequest struct via the builder pattern


### Return type

[**CloudDistributionPoint**](CloudDistributionPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudDistributionPointHistoryGet

> HistorySearchResults V1CloudDistributionPointHistoryGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get cloud distribution point history details



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
	sort := []string{"Inner_example"} // []string | Sorts results by one or more criteria, following the format property:asc/desc. Default sort is ID:asc. If using multiple criteria, separate with commas. (optional) (default to {"id:asc"})
	filter := "filter_example" // string | Filters results. Use RSQL format for query. Allows for many fields, including ID, name, etc. Can be combined with paging and sorting. Default filter is an empty query and returns all results from the requested page. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudDistributionPointAPI.V1CloudDistributionPointHistoryGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudDistributionPointAPI.V1CloudDistributionPointHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CloudDistributionPointHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `CloudDistributionPointAPI.V1CloudDistributionPointHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudDistributionPointHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorts results by one or more criteria, following the format property:asc/desc. Default sort is ID:asc. If using multiple criteria, separate with commas. | [default to {&quot;id:asc&quot;}]
 **filter** | **string** | Filters results. Use RSQL format for query. Allows for many fields, including ID, name, etc. Can be combined with paging and sorting. Default filter is an empty query and returns all results from the requested page. | [default to &quot;&quot;]

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


## V1CloudDistributionPointHistoryPost

> ObjectHistory V1CloudDistributionPointHistoryPost(ctx).ObjectHistoryNote(objectHistoryNote).Execute()

Add specified cloud distribution point history object notes



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
	objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | History note to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudDistributionPointAPI.V1CloudDistributionPointHistoryPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudDistributionPointAPI.V1CloudDistributionPointHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CloudDistributionPointHistoryPost`: ObjectHistory
	fmt.Fprintf(os.Stdout, "Response from `CloudDistributionPointAPI.V1CloudDistributionPointHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudDistributionPointHistoryPostRequest struct via the builder pattern


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


## V1CloudDistributionPointPatch

> CloudDistributionPoint V1CloudDistributionPointPatch(ctx).CloudDistributionPoint(cloudDistributionPoint).Execute()

Update specific fields on a cloud distribution point



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
	cloudDistributionPoint := *openapiclient.NewCloudDistributionPoint(false, "Cannot contact JCDS", "CdnType_example", "Admin", "secretKey123") // CloudDistributionPoint | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudDistributionPointAPI.V1CloudDistributionPointPatch(context.Background()).CloudDistributionPoint(cloudDistributionPoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudDistributionPointAPI.V1CloudDistributionPointPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CloudDistributionPointPatch`: CloudDistributionPoint
	fmt.Fprintf(os.Stdout, "Response from `CloudDistributionPointAPI.V1CloudDistributionPointPatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudDistributionPointPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudDistributionPoint** | [**CloudDistributionPoint**](CloudDistributionPoint.md) |  | 

### Return type

[**CloudDistributionPoint**](CloudDistributionPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudDistributionPointPost

> CloudDistributionPoint V1CloudDistributionPointPost(ctx).CloudDistributionPoint(cloudDistributionPoint).Execute()

Create cloud distribution point



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
	cloudDistributionPoint := *openapiclient.NewCloudDistributionPoint(false, "Cannot contact JCDS", "CdnType_example", "Admin", "secretKey123") // CloudDistributionPoint | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudDistributionPointAPI.V1CloudDistributionPointPost(context.Background()).CloudDistributionPoint(cloudDistributionPoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudDistributionPointAPI.V1CloudDistributionPointPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CloudDistributionPointPost`: CloudDistributionPoint
	fmt.Fprintf(os.Stdout, "Response from `CloudDistributionPointAPI.V1CloudDistributionPointPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudDistributionPointPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudDistributionPoint** | [**CloudDistributionPoint**](CloudDistributionPoint.md) |  | 

### Return type

[**CloudDistributionPoint**](CloudDistributionPoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudDistributionPointTestConnectionGet

> CloudDistributionPointTestConnection V1CloudDistributionPointTestConnectionGet(ctx).Execute()

Get the cloud distribution point test connection details. 



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
	resp, r, err := apiClient.CloudDistributionPointAPI.V1CloudDistributionPointTestConnectionGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudDistributionPointAPI.V1CloudDistributionPointTestConnectionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CloudDistributionPointTestConnectionGet`: CloudDistributionPointTestConnection
	fmt.Fprintf(os.Stdout, "Response from `CloudDistributionPointAPI.V1CloudDistributionPointTestConnectionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudDistributionPointTestConnectionGetRequest struct via the builder pattern


### Return type

[**CloudDistributionPointTestConnection**](CloudDistributionPointTestConnection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudDistributionPointUploadCapabilityGet

> CloudDistributionPointUploadCapability V1CloudDistributionPointUploadCapabilityGet(ctx).Execute()

Finds specific information for the currently configured cloud distribution point. 



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
	resp, r, err := apiClient.CloudDistributionPointAPI.V1CloudDistributionPointUploadCapabilityGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudDistributionPointAPI.V1CloudDistributionPointUploadCapabilityGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CloudDistributionPointUploadCapabilityGet`: CloudDistributionPointUploadCapability
	fmt.Fprintf(os.Stdout, "Response from `CloudDistributionPointAPI.V1CloudDistributionPointUploadCapabilityGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudDistributionPointUploadCapabilityGetRequest struct via the builder pattern


### Return type

[**CloudDistributionPointUploadCapability**](CloudDistributionPointUploadCapability.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

