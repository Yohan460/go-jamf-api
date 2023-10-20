# \UserSessionPreviewAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserGet**](UserSessionPreviewAPI.md#UserGet) | **Get** /user | Return all Jamf Pro user acounts 
[**UserUpdateSessionPost**](UserSessionPreviewAPI.md#UserUpdateSessionPost) | **Post** /user/updateSession | Update values in the User&#39;s current session 



## UserGet

> []Account UserGet(ctx).Execute()

Return all Jamf Pro user acounts 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserSessionPreviewAPI.UserGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSessionPreviewAPI.UserGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGet`: []Account
    fmt.Fprintf(os.Stdout, "Response from `UserSessionPreviewAPI.UserGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetRequest struct via the builder pattern


### Return type

[**[]Account**](Account.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUpdateSessionPost

> Session UserUpdateSessionPost(ctx).Session(session).Execute()

Update values in the User's current session 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    session := *openapiclient.NewSession() // Session | Values to update in user's current session. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserSessionPreviewAPI.UserUpdateSessionPost(context.Background()).Session(session).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSessionPreviewAPI.UserUpdateSessionPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserUpdateSessionPost`: Session
    fmt.Fprintf(os.Stdout, "Response from `UserSessionPreviewAPI.UserUpdateSessionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserUpdateSessionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **session** | [**Session**](Session.md) | Values to update in user&#39;s current session. | 

### Return type

[**Session**](Session.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

