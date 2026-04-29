# \GroupsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1GroupsGet**](GroupsAPI.md#V1GroupsGet) | **Get** /v1/groups | Returns group information for all Mobile Device and Computer groups
[**V1GroupsIdDelete**](GroupsAPI.md#V1GroupsIdDelete) | **Delete** /v1/groups/{id} | Delete a group by platform UUID
[**V1GroupsIdGet**](GroupsAPI.md#V1GroupsIdGet) | **Get** /v1/groups/{id} | Returns group information for the given platform UUID
[**V1GroupsIdPatch**](GroupsAPI.md#V1GroupsIdPatch) | **Patch** /v1/groups/{id} | Update a group by platform UUID



## V1GroupsGet

> GroupSearchResult V1GroupsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Returns group information for all Mobile Device and Computer groups



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is groupName:asc. Multiple sort criteria are supported and must be separated with a comma. Fields allowed in sorting: groupName, groupDescription, groupType, isSmart. Example: sort=groupName:asc,groupType:desc (optional) (default to {"groupName:asc"})
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter group collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: groupPlatformId, groupName, groupDescription, groupType, isSmart. This param can be combined with paging and sorting. When using groupPlatformId in the filter, the supported operators are: =in= (match any in list), =out= (exclude all in list). When using groupType in the filter, the value must be either \"MOBILE\" or \"COMPUTER\" but not both. When using groupType in the filter, the value is case sensitive. When using groupType in the filter, it will exclude groups of the other type regardless of or/and conditionals. Example: filter=groupPlatformId=in=('uuid1','uuid2','uuid3') Example: filter=groupName==\"*Managed*\" and isSmart==\"true\" Example: filter=groupType==\"COMPUTER\" and groupDescription==\"*Admin*\" (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.V1GroupsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.V1GroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1GroupsGet`: GroupSearchResult
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.V1GroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1GroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is groupName:asc. Multiple sort criteria are supported and must be separated with a comma. Fields allowed in sorting: groupName, groupDescription, groupType, isSmart. Example: sort&#x3D;groupName:asc,groupType:desc | [default to {&quot;groupName:asc&quot;}]
 **filter** | **string** | Query in the RSQL format, allowing to filter group collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: groupPlatformId, groupName, groupDescription, groupType, isSmart. This param can be combined with paging and sorting. When using groupPlatformId in the filter, the supported operators are: &#x3D;in&#x3D; (match any in list), &#x3D;out&#x3D; (exclude all in list). When using groupType in the filter, the value must be either \&quot;MOBILE\&quot; or \&quot;COMPUTER\&quot; but not both. When using groupType in the filter, the value is case sensitive. When using groupType in the filter, it will exclude groups of the other type regardless of or/and conditionals. Example: filter&#x3D;groupPlatformId&#x3D;in&#x3D;(&#39;uuid1&#39;,&#39;uuid2&#39;,&#39;uuid3&#39;) Example: filter&#x3D;groupName&#x3D;&#x3D;\&quot;*Managed*\&quot; and isSmart&#x3D;&#x3D;\&quot;true\&quot; Example: filter&#x3D;groupType&#x3D;&#x3D;\&quot;COMPUTER\&quot; and groupDescription&#x3D;&#x3D;\&quot;*Admin*\&quot; | [default to &quot;&quot;]

### Return type

[**GroupSearchResult**](GroupSearchResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GroupsIdDelete

> V1GroupsIdDelete(ctx, id).Execute()

Delete a group by platform UUID



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
	id := "id_example" // string | The platform UUID of a group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.V1GroupsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.V1GroupsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The platform UUID of a group | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GroupsIdDeleteRequest struct via the builder pattern


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


## V1GroupsIdGet

> GroupWithCriteriaDtoV1 V1GroupsIdGet(ctx, id).Execute()

Returns group information for the given platform UUID



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
	id := "id_example" // string | The platform UUID of a group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.V1GroupsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.V1GroupsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1GroupsIdGet`: GroupWithCriteriaDtoV1
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.V1GroupsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The platform UUID of a group | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GroupsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupWithCriteriaDtoV1**](GroupWithCriteriaDtoV1.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GroupsIdPatch

> V1GroupsIdPatch(ctx, id).GroupUpdateDtoV1(groupUpdateDtoV1).Execute()

Update a group by platform UUID



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
	id := "id_example" // string | The platform UUID of a group
	groupUpdateDtoV1 := *openapiclient.NewGroupUpdateDtoV1() // GroupUpdateDtoV1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.V1GroupsIdPatch(context.Background(), id).GroupUpdateDtoV1(groupUpdateDtoV1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.V1GroupsIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The platform UUID of a group | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GroupsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupUpdateDtoV1** | [**GroupUpdateDtoV1**](GroupUpdateDtoV1.md) |  | 

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

