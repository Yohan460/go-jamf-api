# \LdapApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LdapGroupsGet**](LdapApi.md#LdapGroupsGet) | **Get** /ldap/groups | Retrieve the configured access groups that contain the text in the search param 
[**LdapServersGet**](LdapApi.md#LdapServersGet) | **Get** /ldap/servers | Retrieve all Servers including LDAP and Cloud Identity Providers. 
[**V1LdapGroupsGet**](LdapApi.md#V1LdapGroupsGet) | **Get** /v1/ldap/groups | Retrieve the configured access groups that contain the text in the search param 
[**V1LdapLdapServersGet**](LdapApi.md#V1LdapLdapServersGet) | **Get** /v1/ldap/ldap-servers | Retrieve all LDAP Servers. 
[**V1LdapServersGet**](LdapApi.md#V1LdapServersGet) | **Get** /v1/ldap/servers | Retrieve all Servers including LDAP and Cloud Identity Providers. 



## LdapGroupsGet

> LdapGroupSearchResults LdapGroupsGet(ctx).Q(q).Execute()

Retrieve the configured access groups that contain the text in the search param 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    q := "q_example" // string | Will perform a \"contains\" search on the names of access groups (optional) (default to "null")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LdapApi.LdapGroupsGet(context.Background()).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LdapApi.LdapGroupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LdapGroupsGet`: LdapGroupSearchResults
    fmt.Fprintf(os.Stdout, "Response from `LdapApi.LdapGroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLdapGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Will perform a \&quot;contains\&quot; search on the names of access groups | [default to &quot;null&quot;]

### Return type

[**LdapGroupSearchResults**](LdapGroupSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LdapServersGet

> []LdapServer LdapServersGet(ctx).Execute()

Retrieve all Servers including LDAP and Cloud Identity Providers. 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LdapApi.LdapServersGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LdapApi.LdapServersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LdapServersGet`: []LdapServer
    fmt.Fprintf(os.Stdout, "Response from `LdapApi.LdapServersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLdapServersGetRequest struct via the builder pattern


### Return type

[**[]LdapServer**](LdapServer.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1LdapGroupsGet

> LdapGroupSearchResults V1LdapGroupsGet(ctx).Q(q).Execute()

Retrieve the configured access groups that contain the text in the search param 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    q := "q_example" // string | Will perform a \"contains\" search on the names of access groups (optional) (default to "null")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LdapApi.V1LdapGroupsGet(context.Background()).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LdapApi.V1LdapGroupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1LdapGroupsGet`: LdapGroupSearchResults
    fmt.Fprintf(os.Stdout, "Response from `LdapApi.V1LdapGroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1LdapGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Will perform a \&quot;contains\&quot; search on the names of access groups | [default to &quot;null&quot;]

### Return type

[**LdapGroupSearchResults**](LdapGroupSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1LdapLdapServersGet

> []LdapServer V1LdapLdapServersGet(ctx).Execute()

Retrieve all LDAP Servers. 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LdapApi.V1LdapLdapServersGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LdapApi.V1LdapLdapServersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1LdapLdapServersGet`: []LdapServer
    fmt.Fprintf(os.Stdout, "Response from `LdapApi.V1LdapLdapServersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1LdapLdapServersGetRequest struct via the builder pattern


### Return type

[**[]LdapServer**](LdapServer.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1LdapServersGet

> []LdapServer V1LdapServersGet(ctx).Execute()

Retrieve all Servers including LDAP and Cloud Identity Providers. 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LdapApi.V1LdapServersGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LdapApi.V1LdapServersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1LdapServersGet`: []LdapServer
    fmt.Fprintf(os.Stdout, "Response from `LdapApi.V1LdapServersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1LdapServersGetRequest struct via the builder pattern


### Return type

[**[]LdapServer**](LdapServer.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

