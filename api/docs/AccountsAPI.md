# \AccountsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AccountsGet**](AccountsAPI.md#V1AccountsGet) | **Get** /v1/accounts | Get user accounts
[**V1AccountsIdDelete**](AccountsAPI.md#V1AccountsIdDelete) | **Delete** /v1/accounts/{id} | Deletes the user account. 
[**V1AccountsIdGet**](AccountsAPI.md#V1AccountsIdGet) | **Get** /v1/accounts/{id} | Gets the user account. 
[**V1AccountsPost**](AccountsAPI.md#V1AccountsPost) | **Post** /v1/accounts | Adds new account. 



## V1AccountsGet

> UserAccountSearchResults V1AccountsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get user accounts



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is username:desc. Multiple sort criteria are supported and must be separated with a comma. Accepts fields: id, lastPasswordChange, failedLoginAttempts, username, realname, email, phone, ldapServerId, distinguishedName, siteId, privilegeLevel, changePasswordOnNextLogin, accountStatus.  If any other field is passed it will be ignored in sorting operation and/or create unpredictable results.  (optional) (default to {"username:asc"})
	filter := "filter_example" // string | Query in the RSQL format to filter user accounts collection. An empty query returns all results for the requested page. Supported fields: id, lastPasswordChange, failedLoginAttempts, username, realname, email, phone, ldapServerId, distinguishedName, siteId, privilegeLevel, changePasswordOnNextLogin, accountStatus. Multiple conditions can be combined using logical operators. This parameter can be used with paging and sorting parameters. Example: username==\"admin\" and accountStatus==Enabled and failedLoginAttempts==0  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.V1AccountsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.V1AccountsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountsGet`: UserAccountSearchResults
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.V1AccountsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is username:desc. Multiple sort criteria are supported and must be separated with a comma. Accepts fields: id, lastPasswordChange, failedLoginAttempts, username, realname, email, phone, ldapServerId, distinguishedName, siteId, privilegeLevel, changePasswordOnNextLogin, accountStatus.  If any other field is passed it will be ignored in sorting operation and/or create unpredictable results.  | [default to {&quot;username:asc&quot;}]
 **filter** | **string** | Query in the RSQL format to filter user accounts collection. An empty query returns all results for the requested page. Supported fields: id, lastPasswordChange, failedLoginAttempts, username, realname, email, phone, ldapServerId, distinguishedName, siteId, privilegeLevel, changePasswordOnNextLogin, accountStatus. Multiple conditions can be combined using logical operators. This parameter can be used with paging and sorting parameters. Example: username&#x3D;&#x3D;\&quot;admin\&quot; and accountStatus&#x3D;&#x3D;Enabled and failedLoginAttempts&#x3D;&#x3D;0  | [default to &quot;&quot;]

### Return type

[**UserAccountSearchResults**](UserAccountSearchResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountsIdDelete

> V1AccountsIdDelete(ctx, id).Execute()

Deletes the user account. 



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
	id := "id_example" // string | id of target account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountsAPI.V1AccountsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.V1AccountsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of target account | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountsIdDeleteRequest struct via the builder pattern


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


## V1AccountsIdGet

> UserAccount V1AccountsIdGet(ctx, id).Execute()

Gets the user account. 



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
	id := "id_example" // string | id of target account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.V1AccountsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.V1AccountsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountsIdGet`: UserAccount
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.V1AccountsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of target account | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserAccount**](UserAccount.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AccountsPost

> UserAccount V1AccountsPost(ctx).UserAccount(userAccount).Execute()

Adds new account. 



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
	userAccount := *openapiclient.NewUserAccount() // UserAccount | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.V1AccountsPost(context.Background()).UserAccount(userAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.V1AccountsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AccountsPost`: UserAccount
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.V1AccountsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AccountsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAccount** | [**UserAccount**](UserAccount.md) |  | 

### Return type

[**UserAccount**](UserAccount.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

