# \SmartUserGroupsPreviewAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SmartUserGroupsIdRecalculatePost**](SmartUserGroupsPreviewAPI.md#V1SmartUserGroupsIdRecalculatePost) | **Post** /v1/smart-user-groups/{id}/recalculate | Recalculate the smart group for the given id and then return the ids for the users in the smart group 
[**V1UsersIdRecalculateSmartGroupsPost**](SmartUserGroupsPreviewAPI.md#V1UsersIdRecalculateSmartGroupsPost) | **Post** /v1/users/{id}/recalculate-smart-groups | Recalculate a smart group for the given user id and then return the count of smart groups the user falls into 



## V1SmartUserGroupsIdRecalculatePost

> RecalculationResults V1SmartUserGroupsIdRecalculatePost(ctx, id).Execute()

Recalculate the smart group for the given id and then return the ids for the users in the smart group 



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
    id := int32(56) // int32 | instance id of smart group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmartUserGroupsPreviewAPI.V1SmartUserGroupsIdRecalculatePost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartUserGroupsPreviewAPI.V1SmartUserGroupsIdRecalculatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SmartUserGroupsIdRecalculatePost`: RecalculationResults
    fmt.Fprintf(os.Stdout, "Response from `SmartUserGroupsPreviewAPI.V1SmartUserGroupsIdRecalculatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | instance id of smart group | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SmartUserGroupsIdRecalculatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RecalculationResults**](RecalculationResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UsersIdRecalculateSmartGroupsPost

> RecalculationResults V1UsersIdRecalculateSmartGroupsPost(ctx, id).Execute()

Recalculate a smart group for the given user id and then return the count of smart groups the user falls into 



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
    id := int32(56) // int32 | id of user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SmartUserGroupsPreviewAPI.V1UsersIdRecalculateSmartGroupsPost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartUserGroupsPreviewAPI.V1UsersIdRecalculateSmartGroupsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1UsersIdRecalculateSmartGroupsPost`: RecalculationResults
    fmt.Fprintf(os.Stdout, "Response from `SmartUserGroupsPreviewAPI.V1UsersIdRecalculateSmartGroupsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1UsersIdRecalculateSmartGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RecalculationResults**](RecalculationResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

