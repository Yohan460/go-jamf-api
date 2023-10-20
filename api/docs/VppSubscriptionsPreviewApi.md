# \VppSubscriptionsPreviewAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VppSubscriptionsGet**](VppSubscriptionsPreviewAPI.md#VppSubscriptionsGet) | **Get** /vpp/subscriptions | Found all VPP - subscriptions 
[**VppSubscriptionsIdGet**](VppSubscriptionsPreviewAPI.md#VppSubscriptionsIdGet) | **Get** /vpp/subscriptions/{id} | Found VPP subscription by id 



## VppSubscriptionsGet

> []VppTokenSubscription VppSubscriptionsGet(ctx).Execute()

Found all VPP - subscriptions 



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
    resp, r, err := apiClient.VppSubscriptionsPreviewAPI.VppSubscriptionsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VppSubscriptionsPreviewAPI.VppSubscriptionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VppSubscriptionsGet`: []VppTokenSubscription
    fmt.Fprintf(os.Stdout, "Response from `VppSubscriptionsPreviewAPI.VppSubscriptionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVppSubscriptionsGetRequest struct via the builder pattern


### Return type

[**[]VppTokenSubscription**](VppTokenSubscription.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VppSubscriptionsIdGet

> VppTokenSubscription VppSubscriptionsIdGet(ctx, id).Execute()

Found VPP subscription by id 



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
    id := int32(56) // int32 | id of vpp subscription to be retrieved

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VppSubscriptionsPreviewAPI.VppSubscriptionsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VppSubscriptionsPreviewAPI.VppSubscriptionsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VppSubscriptionsIdGet`: VppTokenSubscription
    fmt.Fprintf(os.Stdout, "Response from `VppSubscriptionsPreviewAPI.VppSubscriptionsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id of vpp subscription to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiVppSubscriptionsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VppTokenSubscription**](VppTokenSubscription.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

