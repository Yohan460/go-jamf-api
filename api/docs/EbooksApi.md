# \EbooksAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EbooksGet**](EbooksAPI.md#V1EbooksGet) | **Get** /v1/ebooks | Get Ebook object 
[**V1EbooksIdGet**](EbooksAPI.md#V1EbooksIdGet) | **Get** /v1/ebooks/{id} | Get specified Ebook object 
[**V1EbooksIdScopeGet**](EbooksAPI.md#V1EbooksIdScopeGet) | **Get** /v1/ebooks/{id}/scope | Get specified scope of Ebook object 



## V1EbooksGet

> EbookSearchResults V1EbooksGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Execute()

Get Ebook object 



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["name:asc"])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EbooksAPI.V1EbooksGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EbooksAPI.V1EbooksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EbooksGet`: EbookSearchResults
	fmt.Fprintf(os.Stdout, "Response from `EbooksAPI.V1EbooksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EbooksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;name:asc&quot;]]

### Return type

[**EbookSearchResults**](EbookSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EbooksIdGet

> Ebook V1EbooksIdGet(ctx, id).Execute()

Get specified Ebook object 



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
	id := "id_example" // string | instance id of ebook record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EbooksAPI.V1EbooksIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EbooksAPI.V1EbooksIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EbooksIdGet`: Ebook
	fmt.Fprintf(os.Stdout, "Response from `EbooksAPI.V1EbooksIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of ebook record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EbooksIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Ebook**](Ebook.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EbooksIdScopeGet

> EbookScope V1EbooksIdScopeGet(ctx, id).Execute()

Get specified scope of Ebook object 



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
	id := "id_example" // string | instance id of ebook record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EbooksAPI.V1EbooksIdScopeGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EbooksAPI.V1EbooksIdScopeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EbooksIdScopeGet`: EbookScope
	fmt.Fprintf(os.Stdout, "Response from `EbooksAPI.V1EbooksIdScopeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of ebook record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EbooksIdScopeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EbookScope**](EbookScope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

