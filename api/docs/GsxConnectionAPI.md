# \GsxConnectionAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1GsxConnectionGet**](GsxConnectionAPI.md#V1GsxConnectionGet) | **Get** /v1/gsx-connection | Finds the Jamf Pro GSX Connection information 
[**V1GsxConnectionHistoryGet**](GsxConnectionAPI.md#V1GsxConnectionHistoryGet) | **Get** /v1/gsx-connection/history | Get specified GSX Connection History object 
[**V1GsxConnectionHistoryPost**](GsxConnectionAPI.md#V1GsxConnectionHistoryPost) | **Post** /v1/gsx-connection/history | Add specified GSX Connection history object notes 
[**V1GsxConnectionPatch**](GsxConnectionAPI.md#V1GsxConnectionPatch) | **Patch** /v1/gsx-connection | Updates Jamf Pro GSX Connection information 
[**V1GsxConnectionPut**](GsxConnectionAPI.md#V1GsxConnectionPut) | **Put** /v1/gsx-connection | Updates Jamf Pro GSX Connection information 
[**V1GsxConnectionTestPost**](GsxConnectionAPI.md#V1GsxConnectionTestPost) | **Post** /v1/gsx-connection/test | Test functionality of an GSX Connection



## V1GsxConnectionGet

> GsxConnection V1GsxConnectionGet(ctx).Execute()

Finds the Jamf Pro GSX Connection information 



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
	resp, r, err := apiClient.GsxConnectionAPI.V1GsxConnectionGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GsxConnectionAPI.V1GsxConnectionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1GsxConnectionGet`: GsxConnection
	fmt.Fprintf(os.Stdout, "Response from `GsxConnectionAPI.V1GsxConnectionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1GsxConnectionGetRequest struct via the builder pattern


### Return type

[**GsxConnection**](GsxConnection.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GsxConnectionHistoryGet

> HistorySearchResultsV1 V1GsxConnectionHistoryGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get specified GSX Connection History object 



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["date:desc"])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GsxConnectionAPI.V1GsxConnectionHistoryGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GsxConnectionAPI.V1GsxConnectionHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1GsxConnectionHistoryGet`: HistorySearchResultsV1
	fmt.Fprintf(os.Stdout, "Response from `GsxConnectionAPI.V1GsxConnectionHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1GsxConnectionHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;date:desc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15 | [default to &quot;&quot;]

### Return type

[**HistorySearchResultsV1**](HistorySearchResultsV1.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GsxConnectionHistoryPost

> HrefResponse V1GsxConnectionHistoryPost(ctx).ObjectHistoryNote(objectHistoryNote).Execute()

Add specified GSX Connection history object notes 



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
	resp, r, err := apiClient.GsxConnectionAPI.V1GsxConnectionHistoryPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GsxConnectionAPI.V1GsxConnectionHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1GsxConnectionHistoryPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `GsxConnectionAPI.V1GsxConnectionHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1GsxConnectionHistoryPostRequest struct via the builder pattern


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


## V1GsxConnectionPatch

> GsxConnection V1GsxConnectionPatch(ctx).GsxConnectionUpdate(gsxConnectionUpdate).Execute()

Updates Jamf Pro GSX Connection information 



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
	gsxConnectionUpdate := *openapiclient.NewGsxConnectionUpdate() // GsxConnectionUpdate | GSX Connection to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GsxConnectionAPI.V1GsxConnectionPatch(context.Background()).GsxConnectionUpdate(gsxConnectionUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GsxConnectionAPI.V1GsxConnectionPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1GsxConnectionPatch`: GsxConnection
	fmt.Fprintf(os.Stdout, "Response from `GsxConnectionAPI.V1GsxConnectionPatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1GsxConnectionPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gsxConnectionUpdate** | [**GsxConnectionUpdate**](GsxConnectionUpdate.md) | GSX Connection to update | 

### Return type

[**GsxConnection**](GsxConnection.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GsxConnectionPut

> GsxConnection V1GsxConnectionPut(ctx).GsxConnection(gsxConnection).Execute()

Updates Jamf Pro GSX Connection information 



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
	gsxConnection := *openapiclient.NewGsxConnection(true, "exampleEmail@example.com", "0000012345", "34dsg23-5dsgs-3sdg-4ffs-435sdgs", *openapiclient.NewGsxKeystore("certificate.p12", "test1234")) // GsxConnection | GSX Connection to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GsxConnectionAPI.V1GsxConnectionPut(context.Background()).GsxConnection(gsxConnection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GsxConnectionAPI.V1GsxConnectionPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1GsxConnectionPut`: GsxConnection
	fmt.Fprintf(os.Stdout, "Response from `GsxConnectionAPI.V1GsxConnectionPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1GsxConnectionPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gsxConnection** | [**GsxConnection**](GsxConnection.md) | GSX Connection to update | 

### Return type

[**GsxConnection**](GsxConnection.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GsxConnectionTestPost

> V1GsxConnectionTestPost(ctx).Execute()

Test functionality of an GSX Connection



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
	r, err := apiClient.GsxConnectionAPI.V1GsxConnectionTestPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GsxConnectionAPI.V1GsxConnectionTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1GsxConnectionTestPostRequest struct via the builder pattern


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

