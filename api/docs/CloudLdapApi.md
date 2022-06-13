# \CloudLdapApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CloudLdapsDefaultsMappingsGet**](CloudLdapApi.md#V1CloudLdapsDefaultsMappingsGet) | **Get** /v1/cloud-ldaps/defaults/mappings | Get default mappings
[**V1CloudLdapsDefaultsServerConfigurationGet**](CloudLdapApi.md#V1CloudLdapsDefaultsServerConfigurationGet) | **Get** /v1/cloud-ldaps/defaults/server-configuration | Get default server configuration
[**V1CloudLdapsGet**](CloudLdapApi.md#V1CloudLdapsGet) | **Get** /v1/cloud-ldaps | Get all Cloud Identity Providers configurations.
[**V1CloudLdapsIdConnectionBindGet**](CloudLdapApi.md#V1CloudLdapsIdConnectionBindGet) | **Get** /v1/cloud-ldaps/{id}/connection/bind | Get bind connection pool statistics
[**V1CloudLdapsIdConnectionSearchGet**](CloudLdapApi.md#V1CloudLdapsIdConnectionSearchGet) | **Get** /v1/cloud-ldaps/{id}/connection/search | Get search connection pool statistics
[**V1CloudLdapsIdDelete**](CloudLdapApi.md#V1CloudLdapsIdDelete) | **Delete** /v1/cloud-ldaps/{id} | Delete Cloud Identity Provider configuration.
[**V1CloudLdapsIdGet**](CloudLdapApi.md#V1CloudLdapsIdGet) | **Get** /v1/cloud-ldaps/{id} | Get Cloud Identity Provider configuration with given id.
[**V1CloudLdapsIdHistoryGet**](CloudLdapApi.md#V1CloudLdapsIdHistoryGet) | **Get** /v1/cloud-ldaps/{id}/history | Get Cloud Identity Provider history
[**V1CloudLdapsIdHistoryPost**](CloudLdapApi.md#V1CloudLdapsIdHistoryPost) | **Post** /v1/cloud-ldaps/{id}/history | Add Cloud Identity Provider history note
[**V1CloudLdapsIdMappingsGet**](CloudLdapApi.md#V1CloudLdapsIdMappingsGet) | **Get** /v1/cloud-ldaps/{id}/mappings | Get mappings configurations for Cloud Identity Providers server configuration.
[**V1CloudLdapsIdMappingsPut**](CloudLdapApi.md#V1CloudLdapsIdMappingsPut) | **Put** /v1/cloud-ldaps/{id}/mappings | Update Cloud Identity Provider mappings configuration.
[**V1CloudLdapsIdPut**](CloudLdapApi.md#V1CloudLdapsIdPut) | **Put** /v1/cloud-ldaps/{id} | Old update Cloud Identity Provider configuration
[**V1CloudLdapsIdTestGroupPost**](CloudLdapApi.md#V1CloudLdapsIdTestGroupPost) | **Post** /v1/cloud-ldaps/{id}/test-group | Get group test search
[**V1CloudLdapsIdTestUserMembershipPost**](CloudLdapApi.md#V1CloudLdapsIdTestUserMembershipPost) | **Post** /v1/cloud-ldaps/{id}/test-user-membership | Get membership test search
[**V1CloudLdapsIdTestUserPost**](CloudLdapApi.md#V1CloudLdapsIdTestUserPost) | **Post** /v1/cloud-ldaps/{id}/test-user | Get user test search
[**V1CloudLdapsPost**](CloudLdapApi.md#V1CloudLdapsPost) | **Post** /v1/cloud-ldaps | Create Cloud Identity Provider configuration
[**V1LdapKeystoreVerifyPost**](CloudLdapApi.md#V1LdapKeystoreVerifyPost) | **Post** /v1/ldap-keystore/verify | Validate keystore for Cloud Identity Provider secure connection
[**V2CloudLdapsDefaultsProviderMappingsGet**](CloudLdapApi.md#V2CloudLdapsDefaultsProviderMappingsGet) | **Get** /v2/cloud-ldaps/defaults/{provider}/mappings | Get default mappings
[**V2CloudLdapsDefaultsProviderServerConfigurationGet**](CloudLdapApi.md#V2CloudLdapsDefaultsProviderServerConfigurationGet) | **Get** /v2/cloud-ldaps/defaults/{provider}/server-configuration | Get default server configuration
[**V2CloudLdapsIdConnectionBindGet**](CloudLdapApi.md#V2CloudLdapsIdConnectionBindGet) | **Get** /v2/cloud-ldaps/{id}/connection/bind | Get bind connection pool statistics
[**V2CloudLdapsIdConnectionSearchGet**](CloudLdapApi.md#V2CloudLdapsIdConnectionSearchGet) | **Get** /v2/cloud-ldaps/{id}/connection/search | Get search connection pool statistics
[**V2CloudLdapsIdConnectionStatusGet**](CloudLdapApi.md#V2CloudLdapsIdConnectionStatusGet) | **Get** /v2/cloud-ldaps/{id}/connection/status | Tests the communication with the specified cloud connection 
[**V2CloudLdapsIdDelete**](CloudLdapApi.md#V2CloudLdapsIdDelete) | **Delete** /v2/cloud-ldaps/{id} | Delete Cloud Identity Provider configuration.
[**V2CloudLdapsIdGet**](CloudLdapApi.md#V2CloudLdapsIdGet) | **Get** /v2/cloud-ldaps/{id} | Get Cloud Identity Provider configuration with given id.
[**V2CloudLdapsIdMappingsGet**](CloudLdapApi.md#V2CloudLdapsIdMappingsGet) | **Get** /v2/cloud-ldaps/{id}/mappings | Get mappings configurations for Cloud Identity Providers server configuration.
[**V2CloudLdapsIdMappingsPut**](CloudLdapApi.md#V2CloudLdapsIdMappingsPut) | **Put** /v2/cloud-ldaps/{id}/mappings | Update Cloud Identity Provider mappings configuration.
[**V2CloudLdapsIdPut**](CloudLdapApi.md#V2CloudLdapsIdPut) | **Put** /v2/cloud-ldaps/{id} | Update Cloud Identity Provider configuration
[**V2CloudLdapsPost**](CloudLdapApi.md#V2CloudLdapsPost) | **Post** /v2/cloud-ldaps | Create Cloud Identity Provider configuration



## V1CloudLdapsDefaultsMappingsGet

> CloudLdapMappingsResponse V1CloudLdapsDefaultsMappingsGet(ctx).Execute()

Get default mappings



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
    resp, r, err := apiClient.CloudLdapApi.V1CloudLdapsDefaultsMappingsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V1CloudLdapsDefaultsMappingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudLdapsDefaultsMappingsGet`: CloudLdapMappingsResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V1CloudLdapsDefaultsMappingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudLdapsDefaultsMappingsGetRequest struct via the builder pattern


### Return type

[**CloudLdapMappingsResponse**](CloudLdapMappingsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudLdapsDefaultsServerConfigurationGet

> DeprecatedServerResponse V1CloudLdapsDefaultsServerConfigurationGet(ctx).Execute()

Get default server configuration



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
    resp, r, err := apiClient.CloudLdapApi.V1CloudLdapsDefaultsServerConfigurationGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V1CloudLdapsDefaultsServerConfigurationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudLdapsDefaultsServerConfigurationGet`: DeprecatedServerResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V1CloudLdapsDefaultsServerConfigurationGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudLdapsDefaultsServerConfigurationGetRequest struct via the builder pattern


### Return type

[**DeprecatedServerResponse**](DeprecatedServerResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudLdapsGet

> DeprecatedConfigurationSearchResults V1CloudLdapsGet(ctx).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()

Get all Cloud Identity Providers configurations.



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
    page := int32(56) // int32 |  (optional) (default to 0)
    size := int32(56) // int32 |  (optional) (default to 100)
    pagesize := int32(56) // int32 |  (optional) (default to 100)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["id:desc"])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V1CloudLdapsGet(context.Background()).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V1CloudLdapsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudLdapsGet`: DeprecatedConfigurationSearchResults
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V1CloudLdapsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudLdapsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]
 **pagesize** | **int32** |  | [default to 100]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;id:desc&quot;]]

### Return type

[**DeprecatedConfigurationSearchResults**](DeprecatedConfigurationSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudLdapsIdConnectionBindGet

> CloudLdapConnectionPoolStatistics V1CloudLdapsIdConnectionBindGet(ctx, id).Execute()

Get bind connection pool statistics



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
    id := "id_example" // string | Cloud Identity Provider identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V1CloudLdapsIdConnectionBindGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V1CloudLdapsIdConnectionBindGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudLdapsIdConnectionBindGet`: CloudLdapConnectionPoolStatistics
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V1CloudLdapsIdConnectionBindGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudLdapsIdConnectionBindGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudLdapConnectionPoolStatistics**](CloudLdapConnectionPoolStatistics.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudLdapsIdConnectionSearchGet

> CloudLdapConnectionPoolStatistics V1CloudLdapsIdConnectionSearchGet(ctx, id).Execute()

Get search connection pool statistics



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
    id := "id_example" // string | Cloud Identity Provider identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V1CloudLdapsIdConnectionSearchGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V1CloudLdapsIdConnectionSearchGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudLdapsIdConnectionSearchGet`: CloudLdapConnectionPoolStatistics
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V1CloudLdapsIdConnectionSearchGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudLdapsIdConnectionSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudLdapConnectionPoolStatistics**](CloudLdapConnectionPoolStatistics.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudLdapsIdDelete

> V1CloudLdapsIdDelete(ctx, id).Execute()

Delete Cloud Identity Provider configuration.



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
    id := "id_example" // string | Cloud Identity Provider identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V1CloudLdapsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V1CloudLdapsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudLdapsIdDeleteRequest struct via the builder pattern


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


## V1CloudLdapsIdGet

> DeprecatedConfigurationResponse V1CloudLdapsIdGet(ctx, id).Execute()

Get Cloud Identity Provider configuration with given id.



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
    id := "id_example" // string | Cloud Identity Provider identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V1CloudLdapsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V1CloudLdapsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudLdapsIdGet`: DeprecatedConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V1CloudLdapsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudLdapsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeprecatedConfigurationResponse**](DeprecatedConfigurationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudLdapsIdHistoryGet

> HistorySearchResults V1CloudLdapsIdHistoryGet(ctx, id).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()

Get Cloud Identity Provider history



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
    id := "id_example" // string | Cloud Identity Provider identifier
    page := int32(56) // int32 |  (optional) (default to 0)
    size := int32(56) // int32 |  (optional) (default to 100)
    pagesize := int32(56) // int32 |  (optional) (default to 100)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["date:desc"])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V1CloudLdapsIdHistoryGet(context.Background(), id).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V1CloudLdapsIdHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudLdapsIdHistoryGet`: HistorySearchResults
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V1CloudLdapsIdHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudLdapsIdHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]
 **pagesize** | **int32** |  | [default to 100]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;date:desc&quot;]]

### Return type

[**HistorySearchResults**](HistorySearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudLdapsIdHistoryPost

> ObjectHistory V1CloudLdapsIdHistoryPost(ctx, id).ObjectHistoryNote(objectHistoryNote).Execute()

Add Cloud Identity Provider history note



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
    id := "id_example" // string | Cloud Identity Provider identifier
    objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | history notes to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V1CloudLdapsIdHistoryPost(context.Background(), id).ObjectHistoryNote(objectHistoryNote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V1CloudLdapsIdHistoryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudLdapsIdHistoryPost`: ObjectHistory
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V1CloudLdapsIdHistoryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudLdapsIdHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectHistoryNote** | [**ObjectHistoryNote**](ObjectHistoryNote.md) | history notes to create | 

### Return type

[**ObjectHistory**](ObjectHistory.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudLdapsIdMappingsGet

> CloudLdapMappingsResponse V1CloudLdapsIdMappingsGet(ctx, id).Execute()

Get mappings configurations for Cloud Identity Providers server configuration.



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
    id := "id_example" // string | Cloud Identity Provider identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V1CloudLdapsIdMappingsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V1CloudLdapsIdMappingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudLdapsIdMappingsGet`: CloudLdapMappingsResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V1CloudLdapsIdMappingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudLdapsIdMappingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudLdapMappingsResponse**](CloudLdapMappingsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudLdapsIdMappingsPut

> CloudLdapMappingsResponse V1CloudLdapsIdMappingsPut(ctx, id).CloudLdapMappingsRequest(cloudLdapMappingsRequest).Execute()

Update Cloud Identity Provider mappings configuration.



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
    id := "id_example" // string | Cloud Identity Provider identifier
    cloudLdapMappingsRequest := *openapiclient.NewCloudLdapMappingsRequest(*openapiclient.NewUserMappings("ANY_OBJECT_CLASSES", "inetOrgPerson", "ou=Users", "ALL_SUBTREES", "mail", "uid", "displayName", "mail", "departmentNumber", "Building_example", "Room_example", "Phone_example", "title", "uid"), *openapiclient.NewGroupMappings("ANY_OBJECT_CLASSES", "groupOfNames", "ou=Groups", "ALL_SUBTREES", "cn", "cn", "gidNumber"), *openapiclient.NewMembershipMappings("memberOf")) // CloudLdapMappingsRequest | Cloud Identity Provider mappings to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V1CloudLdapsIdMappingsPut(context.Background(), id).CloudLdapMappingsRequest(cloudLdapMappingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V1CloudLdapsIdMappingsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudLdapsIdMappingsPut`: CloudLdapMappingsResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V1CloudLdapsIdMappingsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudLdapsIdMappingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudLdapMappingsRequest** | [**CloudLdapMappingsRequest**](CloudLdapMappingsRequest.md) | Cloud Identity Provider mappings to update. | 

### Return type

[**CloudLdapMappingsResponse**](CloudLdapMappingsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudLdapsIdPut

> DeprecatedConfigurationResponse V1CloudLdapsIdPut(ctx, id).DeprecatedConfigurationUpdate(deprecatedConfigurationUpdate).Execute()

Old update Cloud Identity Provider configuration



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
    id := "id_example" // string | Cloud Identity Provider identifier
    deprecatedConfigurationUpdate := *openapiclient.NewDeprecatedConfigurationUpdate(*openapiclient.NewDeprecatedServerUpdate("1001", false, "Google", "Google Secure LDAP", "ldap.google.com", "jamf.com", int32(636), int32(15), int32(60), true, "LDAPS")) // DeprecatedConfigurationUpdate | Cloud Identity Provider configuration to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V1CloudLdapsIdPut(context.Background(), id).DeprecatedConfigurationUpdate(deprecatedConfigurationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V1CloudLdapsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudLdapsIdPut`: DeprecatedConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V1CloudLdapsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudLdapsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deprecatedConfigurationUpdate** | [**DeprecatedConfigurationUpdate**](DeprecatedConfigurationUpdate.md) | Cloud Identity Provider configuration to update | 

### Return type

[**DeprecatedConfigurationResponse**](DeprecatedConfigurationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudLdapsIdTestGroupPost

> GroupTestSearchResponse V1CloudLdapsIdTestGroupPost(ctx, id).GroupTestSearchRequest(groupTestSearchRequest).Execute()

Get group test search



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
    id := "id_example" // string | Cloud Identity Provider identifier
    groupTestSearchRequest := *openapiclient.NewGroupTestSearchRequest("users") // GroupTestSearchRequest | Search request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V1CloudLdapsIdTestGroupPost(context.Background(), id).GroupTestSearchRequest(groupTestSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V1CloudLdapsIdTestGroupPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudLdapsIdTestGroupPost`: GroupTestSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V1CloudLdapsIdTestGroupPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudLdapsIdTestGroupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupTestSearchRequest** | [**GroupTestSearchRequest**](GroupTestSearchRequest.md) | Search request | 

### Return type

[**GroupTestSearchResponse**](GroupTestSearchResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudLdapsIdTestUserMembershipPost

> MembershipTestSearchResponse V1CloudLdapsIdTestUserMembershipPost(ctx, id).MembershipTestSearchRequest(membershipTestSearchRequest).Execute()

Get membership test search



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
    id := "id_example" // string | Cloud Identity Provider identifier
    membershipTestSearchRequest := *openapiclient.NewMembershipTestSearchRequest("admin", "users") // MembershipTestSearchRequest | Search request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V1CloudLdapsIdTestUserMembershipPost(context.Background(), id).MembershipTestSearchRequest(membershipTestSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V1CloudLdapsIdTestUserMembershipPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudLdapsIdTestUserMembershipPost`: MembershipTestSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V1CloudLdapsIdTestUserMembershipPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudLdapsIdTestUserMembershipPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **membershipTestSearchRequest** | [**MembershipTestSearchRequest**](MembershipTestSearchRequest.md) | Search request | 

### Return type

[**MembershipTestSearchResponse**](MembershipTestSearchResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudLdapsIdTestUserPost

> UserTestSearchResponse V1CloudLdapsIdTestUserPost(ctx, id).UserTestSearchRequest(userTestSearchRequest).Execute()

Get user test search



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
    id := "id_example" // string | Cloud Identity Provider identifier
    userTestSearchRequest := *openapiclient.NewUserTestSearchRequest("admin") // UserTestSearchRequest | Search request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V1CloudLdapsIdTestUserPost(context.Background(), id).UserTestSearchRequest(userTestSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V1CloudLdapsIdTestUserPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudLdapsIdTestUserPost`: UserTestSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V1CloudLdapsIdTestUserPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudLdapsIdTestUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userTestSearchRequest** | [**UserTestSearchRequest**](UserTestSearchRequest.md) | Search request | 

### Return type

[**UserTestSearchResponse**](UserTestSearchResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudLdapsPost

> DeprecatedConfigurationResponse V1CloudLdapsPost(ctx).DeprecatedConfigurationRequest(deprecatedConfigurationRequest).Execute()

Create Cloud Identity Provider configuration



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
    deprecatedConfigurationRequest := *openapiclient.NewDeprecatedConfigurationRequest(*openapiclient.NewDeprecatedServerRequest(true, "Google", "Google Secure LDAP", "ldap.google.com", "jamf.com", int32(636), *openapiclient.NewCloudLdapKeystoreFile("***", string(123), "keystore.p12"), int32(15), int32(60), true, "LDAPS")) // DeprecatedConfigurationRequest | Cloud Identity Provider configuration to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V1CloudLdapsPost(context.Background()).DeprecatedConfigurationRequest(deprecatedConfigurationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V1CloudLdapsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudLdapsPost`: DeprecatedConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V1CloudLdapsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudLdapsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deprecatedConfigurationRequest** | [**DeprecatedConfigurationRequest**](DeprecatedConfigurationRequest.md) | Cloud Identity Provider configuration to create | 

### Return type

[**DeprecatedConfigurationResponse**](DeprecatedConfigurationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1LdapKeystoreVerifyPost

> CloudLdapKeystore V1LdapKeystoreVerifyPost(ctx).CloudLdapKeystoreFile(cloudLdapKeystoreFile).Execute()

Validate keystore for Cloud Identity Provider secure connection



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
    cloudLdapKeystoreFile := *openapiclient.NewCloudLdapKeystoreFile("***", string(123), "keystore.p12") // CloudLdapKeystoreFile | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V1LdapKeystoreVerifyPost(context.Background()).CloudLdapKeystoreFile(cloudLdapKeystoreFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V1LdapKeystoreVerifyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1LdapKeystoreVerifyPost`: CloudLdapKeystore
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V1LdapKeystoreVerifyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1LdapKeystoreVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudLdapKeystoreFile** | [**CloudLdapKeystoreFile**](CloudLdapKeystoreFile.md) |  | 

### Return type

[**CloudLdapKeystore**](CloudLdapKeystore.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CloudLdapsDefaultsProviderMappingsGet

> CloudLdapMappingsResponse V2CloudLdapsDefaultsProviderMappingsGet(ctx, provider).Execute()

Get default mappings



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
    provider := "google" // string | Cloud Identity Provider name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V2CloudLdapsDefaultsProviderMappingsGet(context.Background(), provider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V2CloudLdapsDefaultsProviderMappingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2CloudLdapsDefaultsProviderMappingsGet`: CloudLdapMappingsResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V2CloudLdapsDefaultsProviderMappingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Cloud Identity Provider name | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2CloudLdapsDefaultsProviderMappingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudLdapMappingsResponse**](CloudLdapMappingsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CloudLdapsDefaultsProviderServerConfigurationGet

> CloudLdapServerResponse V2CloudLdapsDefaultsProviderServerConfigurationGet(ctx, provider).Execute()

Get default server configuration



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
    provider := "google" // string | Cloud Identity Provider name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V2CloudLdapsDefaultsProviderServerConfigurationGet(context.Background(), provider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V2CloudLdapsDefaultsProviderServerConfigurationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2CloudLdapsDefaultsProviderServerConfigurationGet`: CloudLdapServerResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V2CloudLdapsDefaultsProviderServerConfigurationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Cloud Identity Provider name | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2CloudLdapsDefaultsProviderServerConfigurationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudLdapServerResponse**](CloudLdapServerResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CloudLdapsIdConnectionBindGet

> CloudLdapConnectionPoolStatistics V2CloudLdapsIdConnectionBindGet(ctx, id).Execute()

Get bind connection pool statistics



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
    id := "id_example" // string | Cloud Identity Provider identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V2CloudLdapsIdConnectionBindGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V2CloudLdapsIdConnectionBindGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2CloudLdapsIdConnectionBindGet`: CloudLdapConnectionPoolStatistics
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V2CloudLdapsIdConnectionBindGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2CloudLdapsIdConnectionBindGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudLdapConnectionPoolStatistics**](CloudLdapConnectionPoolStatistics.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CloudLdapsIdConnectionSearchGet

> CloudLdapConnectionPoolStatistics V2CloudLdapsIdConnectionSearchGet(ctx, id).Execute()

Get search connection pool statistics



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
    id := "id_example" // string | Cloud Identity Provider identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V2CloudLdapsIdConnectionSearchGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V2CloudLdapsIdConnectionSearchGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2CloudLdapsIdConnectionSearchGet`: CloudLdapConnectionPoolStatistics
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V2CloudLdapsIdConnectionSearchGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2CloudLdapsIdConnectionSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudLdapConnectionPoolStatistics**](CloudLdapConnectionPoolStatistics.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CloudLdapsIdConnectionStatusGet

> CloudLdapConnectionStatus V2CloudLdapsIdConnectionStatusGet(ctx, id).Execute()

Tests the communication with the specified cloud connection 



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
    id := "id_example" // string | Cloud Identity Provider identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V2CloudLdapsIdConnectionStatusGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V2CloudLdapsIdConnectionStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2CloudLdapsIdConnectionStatusGet`: CloudLdapConnectionStatus
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V2CloudLdapsIdConnectionStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2CloudLdapsIdConnectionStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudLdapConnectionStatus**](CloudLdapConnectionStatus.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CloudLdapsIdDelete

> V2CloudLdapsIdDelete(ctx, id).Execute()

Delete Cloud Identity Provider configuration.



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
    id := "id_example" // string | Cloud Identity Provider identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V2CloudLdapsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V2CloudLdapsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2CloudLdapsIdDeleteRequest struct via the builder pattern


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


## V2CloudLdapsIdGet

> LdapConfigurationResponse V2CloudLdapsIdGet(ctx, id).Execute()

Get Cloud Identity Provider configuration with given id.



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
    id := "id_example" // string | Cloud Identity Provider identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V2CloudLdapsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V2CloudLdapsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2CloudLdapsIdGet`: LdapConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V2CloudLdapsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2CloudLdapsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LdapConfigurationResponse**](LdapConfigurationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CloudLdapsIdMappingsGet

> CloudLdapMappingsResponse V2CloudLdapsIdMappingsGet(ctx, id).Execute()

Get mappings configurations for Cloud Identity Providers server configuration.



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
    id := "id_example" // string | Cloud Identity Provider identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V2CloudLdapsIdMappingsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V2CloudLdapsIdMappingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2CloudLdapsIdMappingsGet`: CloudLdapMappingsResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V2CloudLdapsIdMappingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2CloudLdapsIdMappingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudLdapMappingsResponse**](CloudLdapMappingsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CloudLdapsIdMappingsPut

> CloudLdapMappingsResponse V2CloudLdapsIdMappingsPut(ctx, id).CloudLdapMappingsRequest(cloudLdapMappingsRequest).Execute()

Update Cloud Identity Provider mappings configuration.



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
    id := "id_example" // string | Cloud Identity Provider identifier
    cloudLdapMappingsRequest := *openapiclient.NewCloudLdapMappingsRequest(*openapiclient.NewUserMappings("ANY_OBJECT_CLASSES", "inetOrgPerson", "ou=Users", "ALL_SUBTREES", "mail", "uid", "displayName", "mail", "departmentNumber", "Building_example", "Room_example", "Phone_example", "title", "uid"), *openapiclient.NewGroupMappings("ANY_OBJECT_CLASSES", "groupOfNames", "ou=Groups", "ALL_SUBTREES", "cn", "cn", "gidNumber"), *openapiclient.NewMembershipMappings("memberOf")) // CloudLdapMappingsRequest | Cloud Identity Provider mappings to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V2CloudLdapsIdMappingsPut(context.Background(), id).CloudLdapMappingsRequest(cloudLdapMappingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V2CloudLdapsIdMappingsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2CloudLdapsIdMappingsPut`: CloudLdapMappingsResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V2CloudLdapsIdMappingsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2CloudLdapsIdMappingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudLdapMappingsRequest** | [**CloudLdapMappingsRequest**](CloudLdapMappingsRequest.md) | Cloud Identity Provider mappings to update. | 

### Return type

[**CloudLdapMappingsResponse**](CloudLdapMappingsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CloudLdapsIdPut

> LdapConfigurationResponse V2CloudLdapsIdPut(ctx, id).LdapConfigurationUpdate(ldapConfigurationUpdate).Execute()

Update Cloud Identity Provider configuration



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
    id := "id_example" // string | Cloud Identity Provider identifier
    ldapConfigurationUpdate := *openapiclient.NewLdapConfigurationUpdate(*openapiclient.NewCloudIdPCommon("1001", "Google Secure LDAP", "GOOGLE"), *openapiclient.NewCloudLdapServerUpdate("ldap.google.com", true, "jamf.com", int32(636), int32(15), int32(60), true, "LDAPS")) // LdapConfigurationUpdate | Cloud Identity Provider configuration to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V2CloudLdapsIdPut(context.Background(), id).LdapConfigurationUpdate(ldapConfigurationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V2CloudLdapsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2CloudLdapsIdPut`: LdapConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V2CloudLdapsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2CloudLdapsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ldapConfigurationUpdate** | [**LdapConfigurationUpdate**](LdapConfigurationUpdate.md) | Cloud Identity Provider configuration to update | 

### Return type

[**LdapConfigurationResponse**](LdapConfigurationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CloudLdapsPost

> HrefResponse V2CloudLdapsPost(ctx).LdapConfigurationRequest(ldapConfigurationRequest).Execute()

Create Cloud Identity Provider configuration



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
    ldapConfigurationRequest := *openapiclient.NewLdapConfigurationRequest(*openapiclient.NewCloudIdPCommonRequest("Google Secure LDAP", "GOOGLE"), *openapiclient.NewCloudLdapServerRequest("ldap.google.com", true, "jamf.com", int32(636), *openapiclient.NewCloudLdapKeystoreFile("***", string(123), "keystore.p12"), int32(15), int32(60), true, "LDAPS")) // LdapConfigurationRequest | Cloud Identity Provider configuration to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudLdapApi.V2CloudLdapsPost(context.Background()).LdapConfigurationRequest(ldapConfigurationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudLdapApi.V2CloudLdapsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2CloudLdapsPost`: HrefResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudLdapApi.V2CloudLdapsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2CloudLdapsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapConfigurationRequest** | [**LdapConfigurationRequest**](LdapConfigurationRequest.md) | Cloud Identity Provider configuration to create | 

### Return type

[**HrefResponse**](HrefResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

