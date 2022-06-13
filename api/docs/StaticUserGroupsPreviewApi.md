# \StaticUserGroupsPreviewApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1StaticUserGroupsGet**](StaticUserGroupsPreviewApi.md#V1StaticUserGroupsGet) | **Get** /v1/static-user-groups | Return a list of all Static User Groups 
[**V1StaticUserGroupsIdGet**](StaticUserGroupsPreviewApi.md#V1StaticUserGroupsIdGet) | **Get** /v1/static-user-groups/{id} | Return a specific Static User Group by id 



## V1StaticUserGroupsGet

> []StaticUserGroup V1StaticUserGroupsGet(ctx).Execute()

Return a list of all Static User Groups 



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
    resp, r, err := apiClient.StaticUserGroupsPreviewApi.V1StaticUserGroupsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StaticUserGroupsPreviewApi.V1StaticUserGroupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1StaticUserGroupsGet`: []StaticUserGroup
    fmt.Fprintf(os.Stdout, "Response from `StaticUserGroupsPreviewApi.V1StaticUserGroupsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1StaticUserGroupsGetRequest struct via the builder pattern


### Return type

[**[]StaticUserGroup**](StaticUserGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1StaticUserGroupsIdGet

> StaticUserGroup V1StaticUserGroupsIdGet(ctx, id).Execute()

Return a specific Static User Group by id 



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
    id := int32(56) // int32 | Instance id of static user group record

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StaticUserGroupsPreviewApi.V1StaticUserGroupsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StaticUserGroupsPreviewApi.V1StaticUserGroupsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1StaticUserGroupsIdGet`: StaticUserGroup
    fmt.Fprintf(os.Stdout, "Response from `StaticUserGroupsPreviewApi.V1StaticUserGroupsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Instance id of static user group record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1StaticUserGroupsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StaticUserGroup**](StaticUserGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

