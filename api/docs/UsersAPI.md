# \UsersAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1UsersGet**](UsersAPI.md#V1UsersGet) | **Get** /v1/users | Retrieve users with pagination and filtering
[**V1UsersIdDelete**](UsersAPI.md#V1UsersIdDelete) | **Delete** /v1/users/{id} | Delete a user from inventory
[**V1UsersIdGet**](UsersAPI.md#V1UsersIdGet) | **Get** /v1/users/{id} | Retrieve a user by ID
[**V1UsersIdPut**](UsersAPI.md#V1UsersIdPut) | **Put** /v1/users/{id} | Update a user in inventory
[**V1UsersPost**](UsersAPI.md#V1UsersPost) | **Post** /v1/users | Create a new user in inventory



## V1UsersGet

> PagedUserResults V1UsersGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Platform(platform).Execute()

Retrieve users with pagination and filtering



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
	page := int64(56) // int64 | Page to display. Uses zero-based indexing. (optional) (default to 0)
	pageSize := int64(56) // int64 | Number of results per page. Maximum of 1000. (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(:asc|desc). Default sort order is ascending. Multiple sort criteria are supported.  Examples: - sort=username - sort=username:asc - sort=username:asc,realname:desc (optional) (default to {"id:asc"})
	filter := "filter_example" // string | RSQL filter to limit results. Supports all user fields.  Examples: - filter=username==\"john*\" - filter=realname==\"John Smith\" - filter=email==\"*@jamf.com\" - filter=position==\"Manager\";id!=\"1\" - filter=id=in=(123,456,789) (optional) (default to "")
	platform := true // bool | Optional. Return platform identifiers instead of internal identifiers when set to true. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.V1UsersGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Platform(platform).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.V1UsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UsersGet`: PagedUserResults
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.V1UsersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1UsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page to display. Uses zero-based indexing. | [default to 0]
 **pageSize** | **int64** | Number of results per page. Maximum of 1000. | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property(:asc|desc). Default sort order is ascending. Multiple sort criteria are supported.  Examples: - sort&#x3D;username - sort&#x3D;username:asc - sort&#x3D;username:asc,realname:desc | [default to {&quot;id:asc&quot;}]
 **filter** | **string** | RSQL filter to limit results. Supports all user fields.  Examples: - filter&#x3D;username&#x3D;&#x3D;\&quot;john*\&quot; - filter&#x3D;realname&#x3D;&#x3D;\&quot;John Smith\&quot; - filter&#x3D;email&#x3D;&#x3D;\&quot;*@jamf.com\&quot; - filter&#x3D;position&#x3D;&#x3D;\&quot;Manager\&quot;;id!&#x3D;\&quot;1\&quot; - filter&#x3D;id&#x3D;in&#x3D;(123,456,789) | [default to &quot;&quot;]
 **platform** | **bool** | Optional. Return platform identifiers instead of internal identifiers when set to true. | [default to false]

### Return type

[**PagedUserResults**](PagedUserResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UsersIdDelete

> V1UsersIdDelete(ctx, id).Execute()

Delete a user from inventory



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
	id := "1" // string | ID of the user to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.V1UsersIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.V1UsersIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the user to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UsersIdDeleteRequest struct via the builder pattern


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


## V1UsersIdGet

> User V1UsersIdGet(ctx, id).Platform(platform).Execute()

Retrieve a user by ID



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
	id := "id_example" // string | ID of the user to retrieve
	platform := true // bool | Optional. Return platform identifiers instead of internal identifiers when set to true. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.V1UsersIdGet(context.Background(), id).Platform(platform).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.V1UsersIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UsersIdGet`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.V1UsersIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the user to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UsersIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **platform** | **bool** | Optional. Return platform identifiers instead of internal identifiers when set to true. | [default to false]

### Return type

[**User**](User.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UsersIdPut

> V1UsersIdPut(ctx, id).UserInventory(userInventory).Execute()

Update a user in inventory



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
	id := "1" // string | ID of the user to update
	userInventory := *openapiclient.NewUserInventory() // UserInventory | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.V1UsersIdPut(context.Background(), id).UserInventory(userInventory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.V1UsersIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the user to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UsersIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userInventory** | [**UserInventory**](UserInventory.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UsersPost

> HrefResponse V1UsersPost(ctx).UserInventory(userInventory).Platform(platform).Execute()

Create a new user in inventory



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
	userInventory := *openapiclient.NewUserInventory() // UserInventory | 
	platform := true // bool | Internal platform request indicator (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.V1UsersPost(context.Background()).UserInventory(userInventory).Platform(platform).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.V1UsersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UsersPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.V1UsersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1UsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userInventory** | [**UserInventory**](UserInventory.md) |  | 
 **platform** | **bool** | Internal platform request indicator | [default to false]

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

