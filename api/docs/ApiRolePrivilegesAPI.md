# \ApiRolePrivilegesAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ApiRolePrivilegesGet**](ApiRolePrivilegesAPI.md#V1ApiRolePrivilegesGet) | **Get** /v1/api-role-privileges | Get the current Jamf API Role Privileges
[**V1ApiRolePrivilegesSearchGet**](ApiRolePrivilegesAPI.md#V1ApiRolePrivilegesSearchGet) | **Get** /v1/api-role-privileges/search | Search the current Jamf API Role Privileges



## V1ApiRolePrivilegesGet

> ApiRolePrivileges V1ApiRolePrivilegesGet(ctx).Execute()

Get the current Jamf API Role Privileges



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
    resp, r, err := apiClient.ApiRolePrivilegesAPI.V1ApiRolePrivilegesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiRolePrivilegesAPI.V1ApiRolePrivilegesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ApiRolePrivilegesGet`: ApiRolePrivileges
    fmt.Fprintf(os.Stdout, "Response from `ApiRolePrivilegesAPI.V1ApiRolePrivilegesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ApiRolePrivilegesGetRequest struct via the builder pattern


### Return type

[**ApiRolePrivileges**](ApiRolePrivileges.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ApiRolePrivilegesSearchGet

> ApiRolePrivileges V1ApiRolePrivilegesSearchGet(ctx).Name(name).Limit(limit).Execute()

Search the current Jamf API Role Privileges



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
    name := "Flush" // string | The partial or complete privilege name we are searching for
    limit := "50" // string | Limit the query results, defaults to 15 (optional) (default to "15")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiRolePrivilegesAPI.V1ApiRolePrivilegesSearchGet(context.Background()).Name(name).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiRolePrivilegesAPI.V1ApiRolePrivilegesSearchGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ApiRolePrivilegesSearchGet`: ApiRolePrivileges
    fmt.Fprintf(os.Stdout, "Response from `ApiRolePrivilegesAPI.V1ApiRolePrivilegesSearchGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ApiRolePrivilegesSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The partial or complete privilege name we are searching for | 
 **limit** | **string** | Limit the query results, defaults to 15 | [default to &quot;15&quot;]

### Return type

[**ApiRolePrivileges**](ApiRolePrivileges.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

