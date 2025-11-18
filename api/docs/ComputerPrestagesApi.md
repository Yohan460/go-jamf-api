# \ComputerPrestagesAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2ComputerPrestagesIdScopeDeleteMultiplePost**](ComputerPrestagesAPI.md#V2ComputerPrestagesIdScopeDeleteMultiplePost) | **Post** /v2/computer-prestages/{id}/scope/delete-multiple | Remove device Scope for a specific Computer Prestage 
[**V2ComputerPrestagesIdScopeGet**](ComputerPrestagesAPI.md#V2ComputerPrestagesIdScopeGet) | **Get** /v2/computer-prestages/{id}/scope | Get device Scope for a specific Computer Prestage 
[**V2ComputerPrestagesIdScopePost**](ComputerPrestagesAPI.md#V2ComputerPrestagesIdScopePost) | **Post** /v2/computer-prestages/{id}/scope | Add device Scope for a specific Computer Prestage 
[**V2ComputerPrestagesIdScopePut**](ComputerPrestagesAPI.md#V2ComputerPrestagesIdScopePut) | **Put** /v2/computer-prestages/{id}/scope | Replace device Scope for a specific Computer Prestage 
[**V2ComputerPrestagesScopeGet**](ComputerPrestagesAPI.md#V2ComputerPrestagesScopeGet) | **Get** /v2/computer-prestages/scope | Get all device Scope for all Computer Prestages 
[**V3ComputerPrestagesGet**](ComputerPrestagesAPI.md#V3ComputerPrestagesGet) | **Get** /v3/computer-prestages | Get sorted and paged Computer Prestages 
[**V3ComputerPrestagesIdDelete**](ComputerPrestagesAPI.md#V3ComputerPrestagesIdDelete) | **Delete** /v3/computer-prestages/{id} | Delete a Computer Prestage with the supplied id 
[**V3ComputerPrestagesIdGet**](ComputerPrestagesAPI.md#V3ComputerPrestagesIdGet) | **Get** /v3/computer-prestages/{id} | Retrieve a Computer Prestage with the supplied id 
[**V3ComputerPrestagesIdPut**](ComputerPrestagesAPI.md#V3ComputerPrestagesIdPut) | **Put** /v3/computer-prestages/{id} | Update a Computer Prestage 
[**V3ComputerPrestagesPost**](ComputerPrestagesAPI.md#V3ComputerPrestagesPost) | **Post** /v3/computer-prestages | Create a Computer Prestage 



## V2ComputerPrestagesIdScopeDeleteMultiplePost

> PrestageScopeResponseV2 V2ComputerPrestagesIdScopeDeleteMultiplePost(ctx, id).PrestageScopeUpdate(prestageScopeUpdate).Execute()

Remove device Scope for a specific Computer Prestage 



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
	id := "id_example" // string | Computer Prestage identifier
	prestageScopeUpdate := *openapiclient.NewPrestageScopeUpdate([]string{"SerialNumbers_example"}, int64(1)) // PrestageScopeUpdate | Serial Numbers to remove from scope

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerPrestagesAPI.V2ComputerPrestagesIdScopeDeleteMultiplePost(context.Background(), id).PrestageScopeUpdate(prestageScopeUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerPrestagesAPI.V2ComputerPrestagesIdScopeDeleteMultiplePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputerPrestagesIdScopeDeleteMultiplePost`: PrestageScopeResponseV2
	fmt.Fprintf(os.Stdout, "Response from `ComputerPrestagesAPI.V2ComputerPrestagesIdScopeDeleteMultiplePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Computer Prestage identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputerPrestagesIdScopeDeleteMultiplePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prestageScopeUpdate** | [**PrestageScopeUpdate**](PrestageScopeUpdate.md) | Serial Numbers to remove from scope | 

### Return type

[**PrestageScopeResponseV2**](PrestageScopeResponseV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputerPrestagesIdScopeGet

> PrestageScopeResponseV2 V2ComputerPrestagesIdScopeGet(ctx, id).Execute()

Get device Scope for a specific Computer Prestage 



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
	id := "id_example" // string | Computer Prestage identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerPrestagesAPI.V2ComputerPrestagesIdScopeGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerPrestagesAPI.V2ComputerPrestagesIdScopeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputerPrestagesIdScopeGet`: PrestageScopeResponseV2
	fmt.Fprintf(os.Stdout, "Response from `ComputerPrestagesAPI.V2ComputerPrestagesIdScopeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Computer Prestage identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputerPrestagesIdScopeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrestageScopeResponseV2**](PrestageScopeResponseV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputerPrestagesIdScopePost

> PrestageScopeResponseV2 V2ComputerPrestagesIdScopePost(ctx, id).PrestageScopeUpdate(prestageScopeUpdate).Execute()

Add device Scope for a specific Computer Prestage 



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
	id := "id_example" // string | Computer Prestage identifier
	prestageScopeUpdate := *openapiclient.NewPrestageScopeUpdate([]string{"SerialNumbers_example"}, int64(1)) // PrestageScopeUpdate | Serial Numbers to scope

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerPrestagesAPI.V2ComputerPrestagesIdScopePost(context.Background(), id).PrestageScopeUpdate(prestageScopeUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerPrestagesAPI.V2ComputerPrestagesIdScopePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputerPrestagesIdScopePost`: PrestageScopeResponseV2
	fmt.Fprintf(os.Stdout, "Response from `ComputerPrestagesAPI.V2ComputerPrestagesIdScopePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Computer Prestage identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputerPrestagesIdScopePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prestageScopeUpdate** | [**PrestageScopeUpdate**](PrestageScopeUpdate.md) | Serial Numbers to scope | 

### Return type

[**PrestageScopeResponseV2**](PrestageScopeResponseV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputerPrestagesIdScopePut

> PrestageScopeResponseV2 V2ComputerPrestagesIdScopePut(ctx, id).PrestageScopeUpdate(prestageScopeUpdate).Execute()

Replace device Scope for a specific Computer Prestage 



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
	id := "id_example" // string | Computer Prestage identifier
	prestageScopeUpdate := *openapiclient.NewPrestageScopeUpdate([]string{"SerialNumbers_example"}, int64(1)) // PrestageScopeUpdate | Serial Numbers to scope

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerPrestagesAPI.V2ComputerPrestagesIdScopePut(context.Background(), id).PrestageScopeUpdate(prestageScopeUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerPrestagesAPI.V2ComputerPrestagesIdScopePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputerPrestagesIdScopePut`: PrestageScopeResponseV2
	fmt.Fprintf(os.Stdout, "Response from `ComputerPrestagesAPI.V2ComputerPrestagesIdScopePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Computer Prestage identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputerPrestagesIdScopePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prestageScopeUpdate** | [**PrestageScopeUpdate**](PrestageScopeUpdate.md) | Serial Numbers to scope | 

### Return type

[**PrestageScopeResponseV2**](PrestageScopeResponseV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputerPrestagesScopeGet

> PrestageScopeV2 V2ComputerPrestagesScopeGet(ctx).Execute()

Get all device Scope for all Computer Prestages 



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
	resp, r, err := apiClient.ComputerPrestagesAPI.V2ComputerPrestagesScopeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerPrestagesAPI.V2ComputerPrestagesScopeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputerPrestagesScopeGet`: PrestageScopeV2
	fmt.Fprintf(os.Stdout, "Response from `ComputerPrestagesAPI.V2ComputerPrestagesScopeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputerPrestagesScopeGetRequest struct via the builder pattern


### Return type

[**PrestageScopeV2**](PrestageScopeV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ComputerPrestagesGet

> ComputerPrestageSearchResultsV3 V3ComputerPrestagesGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Execute()

Get sorted and paged Computer Prestages 



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to {"id:desc"})

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerPrestagesAPI.V3ComputerPrestagesGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerPrestagesAPI.V3ComputerPrestagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ComputerPrestagesGet`: ComputerPrestageSearchResultsV3
	fmt.Fprintf(os.Stdout, "Response from `ComputerPrestagesAPI.V3ComputerPrestagesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ComputerPrestagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to {&quot;id:desc&quot;}]

### Return type

[**ComputerPrestageSearchResultsV3**](ComputerPrestageSearchResultsV3.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ComputerPrestagesIdDelete

> V3ComputerPrestagesIdDelete(ctx, id).Execute()

Delete a Computer Prestage with the supplied id 



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
	id := "id_example" // string | Computer Prestage identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputerPrestagesAPI.V3ComputerPrestagesIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerPrestagesAPI.V3ComputerPrestagesIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Computer Prestage identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ComputerPrestagesIdDeleteRequest struct via the builder pattern


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


## V3ComputerPrestagesIdGet

> GetComputerPrestageV3 V3ComputerPrestagesIdGet(ctx, id).Execute()

Retrieve a Computer Prestage with the supplied id 



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
	id := "id_example" // string | Computer Prestage identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerPrestagesAPI.V3ComputerPrestagesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerPrestagesAPI.V3ComputerPrestagesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ComputerPrestagesIdGet`: GetComputerPrestageV3
	fmt.Fprintf(os.Stdout, "Response from `ComputerPrestagesAPI.V3ComputerPrestagesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Computer Prestage identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ComputerPrestagesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetComputerPrestageV3**](GetComputerPrestageV3.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ComputerPrestagesIdPut

> GetComputerPrestageV3 V3ComputerPrestagesIdPut(ctx, id).PutComputerPrestageV3(putComputerPrestageV3).Execute()

Update a Computer Prestage 



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
	id := "id_example" // string | Computer Prestage identifier
	putComputerPrestageV3 := *openapiclient.NewPutComputerPrestageV3("Example Mobile Prestage Name", false, true, "5555555555", "example@example.com", "Oxbow", false, "-1", true, true, true, "LDAP authentication prompt", true, true, "5", *openapiclient.NewLocationInformationV2("name", "realName", "123-456-7890", "test@jamf.com", "room", "postion", "1", "1", "-1", int64(1)), *openapiclient.NewPrestagePurchasingInformationV2("-1", true, true, "abcd", "53-1", "Example Vendor", "$500", int64(5), "admin", "true", "2019-01-01", "2019-01-01", "2019-01-01", int64(1)), true, true, []string{"1"}, []string{"1"}, "1") // PutComputerPrestageV3 | Computer Prestage to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerPrestagesAPI.V3ComputerPrestagesIdPut(context.Background(), id).PutComputerPrestageV3(putComputerPrestageV3).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerPrestagesAPI.V3ComputerPrestagesIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ComputerPrestagesIdPut`: GetComputerPrestageV3
	fmt.Fprintf(os.Stdout, "Response from `ComputerPrestagesAPI.V3ComputerPrestagesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Computer Prestage identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ComputerPrestagesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putComputerPrestageV3** | [**PutComputerPrestageV3**](PutComputerPrestageV3.md) | Computer Prestage to update | 

### Return type

[**GetComputerPrestageV3**](GetComputerPrestageV3.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ComputerPrestagesPost

> HrefResponse V3ComputerPrestagesPost(ctx).PostComputerPrestageV3(postComputerPrestageV3).Execute()

Create a Computer Prestage 



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
	postComputerPrestageV3 := *openapiclient.NewPostComputerPrestageV3("Example Mobile Prestage Name", false, true, "5555555555", "example@example.com", "Oxbow", false, "-1", true, true, true, "LDAP authentication prompt", true, true, "5", *openapiclient.NewLocationInformationV2("name", "realName", "123-456-7890", "test@jamf.com", "room", "postion", "1", "1", "-1", int64(1)), *openapiclient.NewPrestagePurchasingInformationV2("-1", true, true, "abcd", "53-1", "Example Vendor", "$500", int64(5), "admin", "true", "2019-01-01", "2019-01-01", "2019-01-01", int64(1)), true, true, []string{"1"}, []string{"1"}, "1") // PostComputerPrestageV3 | Computer Prestage to create. ids defined in this body will be ignored

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerPrestagesAPI.V3ComputerPrestagesPost(context.Background()).PostComputerPrestageV3(postComputerPrestageV3).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerPrestagesAPI.V3ComputerPrestagesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ComputerPrestagesPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `ComputerPrestagesAPI.V3ComputerPrestagesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ComputerPrestagesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postComputerPrestageV3** | [**PostComputerPrestageV3**](PostComputerPrestageV3.md) | Computer Prestage to create. ids defined in this body will be ignored | 

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

