# \AdvancedUserContentSearchesAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AdvancedUserContentSearchesGet**](AdvancedUserContentSearchesAPI.md#V1AdvancedUserContentSearchesGet) | **Get** /v1/advanced-user-content-searches | Get All Advanced User Content Search objects 
[**V1AdvancedUserContentSearchesIdDelete**](AdvancedUserContentSearchesAPI.md#V1AdvancedUserContentSearchesIdDelete) | **Delete** /v1/advanced-user-content-searches/{id} | Remove specified Advanced User Content Search object 
[**V1AdvancedUserContentSearchesIdGet**](AdvancedUserContentSearchesAPI.md#V1AdvancedUserContentSearchesIdGet) | **Get** /v1/advanced-user-content-searches/{id} | Get Specified Advanced User Content Search object 
[**V1AdvancedUserContentSearchesIdPut**](AdvancedUserContentSearchesAPI.md#V1AdvancedUserContentSearchesIdPut) | **Put** /v1/advanced-user-content-searches/{id} | Get Specified Advanced User Content Search object 
[**V1AdvancedUserContentSearchesPost**](AdvancedUserContentSearchesAPI.md#V1AdvancedUserContentSearchesPost) | **Post** /v1/advanced-user-content-searches | Create Advanced User Content Search object 



## V1AdvancedUserContentSearchesGet

> AdvancedUserContentSearchSearchResults V1AdvancedUserContentSearchesGet(ctx).Execute()

Get All Advanced User Content Search objects 



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
	resp, r, err := apiClient.AdvancedUserContentSearchesAPI.V1AdvancedUserContentSearchesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedUserContentSearchesAPI.V1AdvancedUserContentSearchesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AdvancedUserContentSearchesGet`: AdvancedUserContentSearchSearchResults
	fmt.Fprintf(os.Stdout, "Response from `AdvancedUserContentSearchesAPI.V1AdvancedUserContentSearchesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AdvancedUserContentSearchesGetRequest struct via the builder pattern


### Return type

[**AdvancedUserContentSearchSearchResults**](AdvancedUserContentSearchSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AdvancedUserContentSearchesIdDelete

> V1AdvancedUserContentSearchesIdDelete(ctx, id).Execute()

Remove specified Advanced User Content Search object 



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
	id := "id_example" // string | instance id of Advanced User Content Search record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdvancedUserContentSearchesAPI.V1AdvancedUserContentSearchesIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedUserContentSearchesAPI.V1AdvancedUserContentSearchesIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of Advanced User Content Search record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AdvancedUserContentSearchesIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## V1AdvancedUserContentSearchesIdGet

> AdvancedUserContentSearch V1AdvancedUserContentSearchesIdGet(ctx, id).Execute()

Get Specified Advanced User Content Search object 



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
	id := "id_example" // string | id of target Advanced User Content Search

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedUserContentSearchesAPI.V1AdvancedUserContentSearchesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedUserContentSearchesAPI.V1AdvancedUserContentSearchesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AdvancedUserContentSearchesIdGet`: AdvancedUserContentSearch
	fmt.Fprintf(os.Stdout, "Response from `AdvancedUserContentSearchesAPI.V1AdvancedUserContentSearchesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of target Advanced User Content Search | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AdvancedUserContentSearchesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdvancedUserContentSearch**](AdvancedUserContentSearch.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AdvancedUserContentSearchesIdPut

> AdvancedUserContentSearch V1AdvancedUserContentSearchesIdPut(ctx, id).AdvancedUserContentSearch(advancedUserContentSearch).Execute()

Get Specified Advanced User Content Search object 



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
	id := "id_example" // string | id of target Advanced User Content Search
	advancedUserContentSearch := *openapiclient.NewAdvancedUserContentSearch("Andy's Search") // AdvancedUserContentSearch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedUserContentSearchesAPI.V1AdvancedUserContentSearchesIdPut(context.Background(), id).AdvancedUserContentSearch(advancedUserContentSearch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedUserContentSearchesAPI.V1AdvancedUserContentSearchesIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AdvancedUserContentSearchesIdPut`: AdvancedUserContentSearch
	fmt.Fprintf(os.Stdout, "Response from `AdvancedUserContentSearchesAPI.V1AdvancedUserContentSearchesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of target Advanced User Content Search | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AdvancedUserContentSearchesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **advancedUserContentSearch** | [**AdvancedUserContentSearch**](AdvancedUserContentSearch.md) |  | 

### Return type

[**AdvancedUserContentSearch**](AdvancedUserContentSearch.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AdvancedUserContentSearchesPost

> HrefResponse V1AdvancedUserContentSearchesPost(ctx).AdvancedUserContentSearch(advancedUserContentSearch).Execute()

Create Advanced User Content Search object 



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
	advancedUserContentSearch := *openapiclient.NewAdvancedUserContentSearch("Andy's Search") // AdvancedUserContentSearch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedUserContentSearchesAPI.V1AdvancedUserContentSearchesPost(context.Background()).AdvancedUserContentSearch(advancedUserContentSearch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedUserContentSearchesAPI.V1AdvancedUserContentSearchesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AdvancedUserContentSearchesPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `AdvancedUserContentSearchesAPI.V1AdvancedUserContentSearchesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AdvancedUserContentSearchesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **advancedUserContentSearch** | [**AdvancedUserContentSearch**](AdvancedUserContentSearch.md) |  | 

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

