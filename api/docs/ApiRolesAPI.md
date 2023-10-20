# \ApiRolesAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiRole**](ApiRolesAPI.md#DeleteApiRole) | **Delete** /v1/api-roles/{id} | Delete API Integrations Role
[**GetAllApiRoles**](ApiRolesAPI.md#GetAllApiRoles) | **Get** /v1/api-roles | Get the current Jamf API Roles
[**GetOneApiRole**](ApiRolesAPI.md#GetOneApiRole) | **Get** /v1/api-roles/{id} | Get the specific Jamf API Role
[**PostCreateApiRole**](ApiRolesAPI.md#PostCreateApiRole) | **Post** /v1/api-roles | Create a new API role
[**PutUpdateApiRole**](ApiRolesAPI.md#PutUpdateApiRole) | **Put** /v1/api-roles/{id} | Update API Integrations Role



## DeleteApiRole

> DeleteApiRole(ctx, id).Execute()

Delete API Integrations Role



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
    id := "id_example" // string | instance id of API role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApiRolesAPI.DeleteApiRole(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiRolesAPI.DeleteApiRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of API role | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiRoleRequest struct via the builder pattern


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


## GetAllApiRoles

> ApiRoleResult GetAllApiRoles(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get the current Jamf API Roles



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
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Fields allowed in the query: id, displayName. Example: sort=displayName:desc (optional) (default to ["id:asc"])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter app titles collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, displayName. Example: displayName==\"*myRole*\" (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiRolesAPI.GetAllApiRoles(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiRolesAPI.GetAllApiRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllApiRoles`: ApiRoleResult
    fmt.Fprintf(os.Stdout, "Response from `ApiRolesAPI.GetAllApiRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllApiRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Fields allowed in the query: id, displayName. Example: sort&#x3D;displayName:desc | [default to [&quot;id:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter app titles collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, displayName. Example: displayName&#x3D;&#x3D;\&quot;*myRole*\&quot; | [default to &quot;&quot;]

### Return type

[**ApiRoleResult**](ApiRoleResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOneApiRole

> ApiRole GetOneApiRole(ctx, id).Execute()

Get the specific Jamf API Role



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
    id := "id_example" // string | instance id of API role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiRolesAPI.GetOneApiRole(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiRolesAPI.GetOneApiRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOneApiRole`: ApiRole
    fmt.Fprintf(os.Stdout, "Response from `ApiRolesAPI.GetOneApiRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of API role | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOneApiRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiRole**](ApiRole.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateApiRole

> ApiRole PostCreateApiRole(ctx).ApiRoleRequest(apiRoleRequest).Execute()

Create a new API role



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
    apiRoleRequest := *openapiclient.NewApiRoleRequest("One Role to Rule them all", []string{"View License Serial Numbers"}) // ApiRoleRequest | API Integrations Role to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiRolesAPI.PostCreateApiRole(context.Background()).ApiRoleRequest(apiRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiRolesAPI.PostCreateApiRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateApiRole`: ApiRole
    fmt.Fprintf(os.Stdout, "Response from `ApiRolesAPI.PostCreateApiRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateApiRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiRoleRequest** | [**ApiRoleRequest**](ApiRoleRequest.md) | API Integrations Role to create | 

### Return type

[**ApiRole**](ApiRole.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUpdateApiRole

> ApiRole PutUpdateApiRole(ctx, id).ApiRoleRequest(apiRoleRequest).Execute()

Update API Integrations Role



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
    id := "id_example" // string | instance id of API role
    apiRoleRequest := *openapiclient.NewApiRoleRequest("One Role to Rule them all", []string{"View License Serial Numbers"}) // ApiRoleRequest | API Integrations Role to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiRolesAPI.PutUpdateApiRole(context.Background(), id).ApiRoleRequest(apiRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiRolesAPI.PutUpdateApiRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateApiRole`: ApiRole
    fmt.Fprintf(os.Stdout, "Response from `ApiRolesAPI.PutUpdateApiRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of API role | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateApiRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiRoleRequest** | [**ApiRoleRequest**](ApiRoleRequest.md) | API Integrations Role to update | 

### Return type

[**ApiRole**](ApiRole.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

