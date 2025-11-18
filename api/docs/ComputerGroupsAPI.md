# \ComputerGroupsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ComputerGroupsGet**](ComputerGroupsAPI.md#V1ComputerGroupsGet) | **Get** /v1/computer-groups | Returns the list of all computer groups 
[**V2ComputerGroupsSmartGroupMembershipIdGet**](ComputerGroupsAPI.md#V2ComputerGroupsSmartGroupMembershipIdGet) | **Get** /v2/computer-groups/smart-group-membership/{id} | Get the membership of a Smart Computer Group 
[**V2ComputerGroupsSmartGroupsGet**](ComputerGroupsAPI.md#V2ComputerGroupsSmartGroupsGet) | **Get** /v2/computer-groups/smart-groups | Search for Smart Computer Groups 
[**V2ComputerGroupsSmartGroupsIdDelete**](ComputerGroupsAPI.md#V2ComputerGroupsSmartGroupsIdDelete) | **Delete** /v2/computer-groups/smart-groups/{id} | Remove specified Smart Computer Group 
[**V2ComputerGroupsSmartGroupsIdGet**](ComputerGroupsAPI.md#V2ComputerGroupsSmartGroupsIdGet) | **Get** /v2/computer-groups/smart-groups/{id} | Get Smart Computer Group by Id 
[**V2ComputerGroupsSmartGroupsIdPut**](ComputerGroupsAPI.md#V2ComputerGroupsSmartGroupsIdPut) | **Put** /v2/computer-groups/smart-groups/{id} | Update a Smart Computer Group 
[**V2ComputerGroupsSmartGroupsPost**](ComputerGroupsAPI.md#V2ComputerGroupsSmartGroupsPost) | **Post** /v2/computer-groups/smart-groups | Create a Smart Computer Group 
[**V2ComputerGroupsStaticGroupsGet**](ComputerGroupsAPI.md#V2ComputerGroupsStaticGroupsGet) | **Get** /v2/computer-groups/static-groups | Search for Static Computer Groups 
[**V2ComputerGroupsStaticGroupsIdDelete**](ComputerGroupsAPI.md#V2ComputerGroupsStaticGroupsIdDelete) | **Delete** /v2/computer-groups/static-groups/{id} | Remove Static Computer Group by Id 
[**V2ComputerGroupsStaticGroupsIdGet**](ComputerGroupsAPI.md#V2ComputerGroupsStaticGroupsIdGet) | **Get** /v2/computer-groups/static-groups/{id} | Get Static Computer Group by Id 
[**V2ComputerGroupsStaticGroupsIdPut**](ComputerGroupsAPI.md#V2ComputerGroupsStaticGroupsIdPut) | **Put** /v2/computer-groups/static-groups/{id} | Update membership of a static computer group. 
[**V2ComputerGroupsStaticGroupsPost**](ComputerGroupsAPI.md#V2ComputerGroupsStaticGroupsPost) | **Post** /v2/computer-groups/static-groups | Create membership of a static computer group. 



## V1ComputerGroupsGet

> []ComputerGroup V1ComputerGroupsGet(ctx).Execute()

Returns the list of all computer groups 



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
	resp, r, err := apiClient.ComputerGroupsAPI.V1ComputerGroupsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerGroupsAPI.V1ComputerGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputerGroupsGet`: []ComputerGroup
	fmt.Fprintf(os.Stdout, "Response from `ComputerGroupsAPI.V1ComputerGroupsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputerGroupsGetRequest struct via the builder pattern


### Return type

[**[]ComputerGroup**](ComputerGroup.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputerGroupsSmartGroupMembershipIdGet

> SmartGroupMembership V2ComputerGroupsSmartGroupMembershipIdGet(ctx, id).Execute()

Get the membership of a Smart Computer Group 



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
	id := "id_example" // string | id of the Smart Computer Group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerGroupsAPI.V2ComputerGroupsSmartGroupMembershipIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerGroupsAPI.V2ComputerGroupsSmartGroupMembershipIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputerGroupsSmartGroupMembershipIdGet`: SmartGroupMembership
	fmt.Fprintf(os.Stdout, "Response from `ComputerGroupsAPI.V2ComputerGroupsSmartGroupMembershipIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the Smart Computer Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputerGroupsSmartGroupMembershipIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SmartGroupMembership**](SmartGroupMembership.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputerGroupsSmartGroupsGet

> SmartGroupSearchResult V2ComputerGroupsSmartGroupsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Search for Smart Computer Groups 



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=name:asc (optional) (default to {"id:asc"})
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter smart computer group collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name, siteId. The siteId field can only be filtered by admins with full access. Any sited admin will have siteId filtered automatically. Example: name==\"*group*\" (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerGroupsAPI.V2ComputerGroupsSmartGroupsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerGroupsAPI.V2ComputerGroupsSmartGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputerGroupsSmartGroupsGet`: SmartGroupSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ComputerGroupsAPI.V2ComputerGroupsSmartGroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputerGroupsSmartGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;name:asc | [default to {&quot;id:asc&quot;}]
 **filter** | **string** | Query in the RSQL format, allowing to filter smart computer group collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name, siteId. The siteId field can only be filtered by admins with full access. Any sited admin will have siteId filtered automatically. Example: name&#x3D;&#x3D;\&quot;*group*\&quot; | [default to &quot;&quot;]

### Return type

[**SmartGroupSearchResult**](SmartGroupSearchResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputerGroupsSmartGroupsIdDelete

> V2ComputerGroupsSmartGroupsIdDelete(ctx, id).Execute()

Remove specified Smart Computer Group 



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
	id := "id_example" // string | id of target Smart Computer Group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputerGroupsAPI.V2ComputerGroupsSmartGroupsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerGroupsAPI.V2ComputerGroupsSmartGroupsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of target Smart Computer Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputerGroupsSmartGroupsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputerGroupsSmartGroupsIdGet

> SmartComputerGroupV2 V2ComputerGroupsSmartGroupsIdGet(ctx, id).Execute()

Get Smart Computer Group by Id 



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
	id := "id_example" // string | instance id of smart computer group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerGroupsAPI.V2ComputerGroupsSmartGroupsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerGroupsAPI.V2ComputerGroupsSmartGroupsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputerGroupsSmartGroupsIdGet`: SmartComputerGroupV2
	fmt.Fprintf(os.Stdout, "Response from `ComputerGroupsAPI.V2ComputerGroupsSmartGroupsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of smart computer group | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputerGroupsSmartGroupsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SmartComputerGroupV2**](SmartComputerGroupV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputerGroupsSmartGroupsIdPut

> SmartComputerGroupV2 V2ComputerGroupsSmartGroupsIdPut(ctx, id).SmartComputerGroupV2(smartComputerGroupV2).Execute()

Update a Smart Computer Group 



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
	id := "id_example" // string | id of target Smart Computer Group
	smartComputerGroupV2 := *openapiclient.NewSmartComputerGroupV2("New Group Name") // SmartComputerGroupV2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerGroupsAPI.V2ComputerGroupsSmartGroupsIdPut(context.Background(), id).SmartComputerGroupV2(smartComputerGroupV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerGroupsAPI.V2ComputerGroupsSmartGroupsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputerGroupsSmartGroupsIdPut`: SmartComputerGroupV2
	fmt.Fprintf(os.Stdout, "Response from `ComputerGroupsAPI.V2ComputerGroupsSmartGroupsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of target Smart Computer Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputerGroupsSmartGroupsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smartComputerGroupV2** | [**SmartComputerGroupV2**](SmartComputerGroupV2.md) |  | 

### Return type

[**SmartComputerGroupV2**](SmartComputerGroupV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputerGroupsSmartGroupsPost

> HrefResponse V2ComputerGroupsSmartGroupsPost(ctx).SmartComputerGroupV2(smartComputerGroupV2).Execute()

Create a Smart Computer Group 



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
	smartComputerGroupV2 := *openapiclient.NewSmartComputerGroupV2("New Group Name") // SmartComputerGroupV2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerGroupsAPI.V2ComputerGroupsSmartGroupsPost(context.Background()).SmartComputerGroupV2(smartComputerGroupV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerGroupsAPI.V2ComputerGroupsSmartGroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputerGroupsSmartGroupsPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `ComputerGroupsAPI.V2ComputerGroupsSmartGroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputerGroupsSmartGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smartComputerGroupV2** | [**SmartComputerGroupV2**](SmartComputerGroupV2.md) |  | 

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


## V2ComputerGroupsStaticGroupsGet

> StaticComputerGroupSearchResults V2ComputerGroupsStaticGroupsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Search for Static Computer Groups 



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=name:asc (optional) (default to {"id:asc"})
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter static computer group collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name, siteId. The siteId field can only be filtered by admins with full access. Any sited admin will have siteId filtered automatically. Example: name==\"*group*\" (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerGroupsAPI.V2ComputerGroupsStaticGroupsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerGroupsAPI.V2ComputerGroupsStaticGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputerGroupsStaticGroupsGet`: StaticComputerGroupSearchResults
	fmt.Fprintf(os.Stdout, "Response from `ComputerGroupsAPI.V2ComputerGroupsStaticGroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputerGroupsStaticGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;name:asc | [default to {&quot;id:asc&quot;}]
 **filter** | **string** | Query in the RSQL format, allowing to filter static computer group collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name, siteId. The siteId field can only be filtered by admins with full access. Any sited admin will have siteId filtered automatically. Example: name&#x3D;&#x3D;\&quot;*group*\&quot; | [default to &quot;&quot;]

### Return type

[**StaticComputerGroupSearchResults**](StaticComputerGroupSearchResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputerGroupsStaticGroupsIdDelete

> V2ComputerGroupsStaticGroupsIdDelete(ctx, id).Execute()

Remove Static Computer Group by Id 



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
	id := "id_example" // string | instance id of static computer group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputerGroupsAPI.V2ComputerGroupsStaticGroupsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerGroupsAPI.V2ComputerGroupsStaticGroupsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of static computer group | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputerGroupsStaticGroupsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputerGroupsStaticGroupsIdGet

> StaticComputerGroup V2ComputerGroupsStaticGroupsIdGet(ctx, id).Execute()

Get Static Computer Group by Id 



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
	id := "id_example" // string | instance id of static computer group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerGroupsAPI.V2ComputerGroupsStaticGroupsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerGroupsAPI.V2ComputerGroupsStaticGroupsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputerGroupsStaticGroupsIdGet`: StaticComputerGroup
	fmt.Fprintf(os.Stdout, "Response from `ComputerGroupsAPI.V2ComputerGroupsStaticGroupsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of static computer group | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputerGroupsStaticGroupsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StaticComputerGroup**](StaticComputerGroup.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputerGroupsStaticGroupsIdPut

> StaticComputerGroupAssignment V2ComputerGroupsStaticGroupsIdPut(ctx, id).StaticComputerGroupAssignment(staticComputerGroupAssignment).Execute()

Update membership of a static computer group. 



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
	id := "id_example" // string | instance id of a static computer group
	staticComputerGroupAssignment := *openapiclient.NewStaticComputerGroupAssignment("Test Static Computer Group") // StaticComputerGroupAssignment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerGroupsAPI.V2ComputerGroupsStaticGroupsIdPut(context.Background(), id).StaticComputerGroupAssignment(staticComputerGroupAssignment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerGroupsAPI.V2ComputerGroupsStaticGroupsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputerGroupsStaticGroupsIdPut`: StaticComputerGroupAssignment
	fmt.Fprintf(os.Stdout, "Response from `ComputerGroupsAPI.V2ComputerGroupsStaticGroupsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of a static computer group | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputerGroupsStaticGroupsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **staticComputerGroupAssignment** | [**StaticComputerGroupAssignment**](StaticComputerGroupAssignment.md) |  | 

### Return type

[**StaticComputerGroupAssignment**](StaticComputerGroupAssignment.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputerGroupsStaticGroupsPost

> HrefResponse V2ComputerGroupsStaticGroupsPost(ctx).StaticComputerGroupAssignment(staticComputerGroupAssignment).Execute()

Create membership of a static computer group. 



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
	staticComputerGroupAssignment := *openapiclient.NewStaticComputerGroupAssignment("Test Static Computer Group") // StaticComputerGroupAssignment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerGroupsAPI.V2ComputerGroupsStaticGroupsPost(context.Background()).StaticComputerGroupAssignment(staticComputerGroupAssignment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerGroupsAPI.V2ComputerGroupsStaticGroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputerGroupsStaticGroupsPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `ComputerGroupsAPI.V2ComputerGroupsStaticGroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputerGroupsStaticGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **staticComputerGroupAssignment** | [**StaticComputerGroupAssignment**](StaticComputerGroupAssignment.md) |  | 

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

