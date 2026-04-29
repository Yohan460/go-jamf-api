# \AccountGroupsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountGroupsGet**](AccountGroupsAPI.md#V1AccountGroupsGet) | **Get** /v1/account-groups | Get account groups
[**V1AccountGroupsIdGet**](AccountGroupsAPI.md#V1AccountGroupsIdGet) | **Get** /v1/account-groups/{id} | Gets the account group. 



## V1AccountGroupsGet

> AccountGroupSearchResultsV1 V1AccountGroupsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get account groups



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Accepts fields: id, name, siteId, ldapServerId.  (optional) (default to {"name:asc"})
	filter := "filter_example" // string | Query in the RSQL format to filter account groups collection. An empty query returns all results for the requested page. Supported fields: id, name, siteId, ldapServerId. Multiple conditions can be combined using logical operators. This parameter can be used with paging and sorting parameters. Example: name==\"Admins\" and siteId==-1  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountGroupsAPI.V1AccountGroupsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsAPI.V1AccountGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountGroupsGet`: AccountGroupSearchResultsV1
	fmt.Fprintf(os.Stdout, "Response from `AccountGroupsAPI.V1AccountGroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Accepts fields: id, name, siteId, ldapServerId.  | [default to {&quot;name:asc&quot;}]
 **filter** | **string** | Query in the RSQL format to filter account groups collection. An empty query returns all results for the requested page. Supported fields: id, name, siteId, ldapServerId. Multiple conditions can be combined using logical operators. This parameter can be used with paging and sorting parameters. Example: name&#x3D;&#x3D;\&quot;Admins\&quot; and siteId&#x3D;&#x3D;-1  | [default to &quot;&quot;]

### Return type

[**AccountGroupSearchResultsV1**](AccountGroupSearchResultsV1.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountGroupsIdGet

> AccountGroupV1 V1AccountGroupsIdGet(ctx, id).Execute()

Gets the account group. 



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
	id := "id_example" // string | id of target account group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountGroupsAPI.V1AccountGroupsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsAPI.V1AccountGroupsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountGroupsIdGet`: AccountGroupV1
	fmt.Fprintf(os.Stdout, "Response from `AccountGroupsAPI.V1AccountGroupsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of target account group | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountGroupsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountGroupV1**](AccountGroupV1.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

