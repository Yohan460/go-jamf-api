# \GroupsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1GroupsGet**](GroupsAPI.md#V1GroupsGet) | **Get** /v1/groups | Returns group information for all Mobile Device and Computer groups
[**V1GroupsIdGet**](GroupsAPI.md#V1GroupsIdGet) | **Get** /v1/groups/{id} | Returns group information for the given platform UUID



## V1GroupsGet

> GroupSearchResult V1GroupsGet(ctx).Page(page).PageSize(pageSize).Filter(filter).Execute()

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
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter group collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: groupName, isSmart. This param can be combined with paging. Example: filter=groupName==\"*Managed*\" and isSmart==\"true\" (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.V1GroupsGet(context.Background()).Page(page).PageSize(pageSize).Filter(filter).Execute()
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
 **filter** | **string** | Query in the RSQL format, allowing to filter group collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: groupName, isSmart. This param can be combined with paging. Example: filter&#x3D;groupName&#x3D;&#x3D;\&quot;*Managed*\&quot; and isSmart&#x3D;&#x3D;\&quot;true\&quot; | [default to &quot;&quot;]

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


## V1GroupsIdGet

> GroupV1 V1GroupsIdGet(ctx, id).Execute()

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
	// response from `V1GroupsIdGet`: GroupV1
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

[**GroupV1**](GroupV1.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

