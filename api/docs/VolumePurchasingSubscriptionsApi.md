# \VolumePurchasingSubscriptionsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1VolumePurchasingSubscriptionsGet**](VolumePurchasingSubscriptionsAPI.md#V1VolumePurchasingSubscriptionsGet) | **Get** /v1/volume-purchasing-subscriptions | Retrieve Volume Purchasing Subscriptions
[**V1VolumePurchasingSubscriptionsIdDelete**](VolumePurchasingSubscriptionsAPI.md#V1VolumePurchasingSubscriptionsIdDelete) | **Delete** /v1/volume-purchasing-subscriptions/{id} | Delete a Volume Purchasing Subscription with the supplied id
[**V1VolumePurchasingSubscriptionsIdGet**](VolumePurchasingSubscriptionsAPI.md#V1VolumePurchasingSubscriptionsIdGet) | **Get** /v1/volume-purchasing-subscriptions/{id} | Retrieve a Volume Purchasing Subscription with the supplied id
[**V1VolumePurchasingSubscriptionsIdPut**](VolumePurchasingSubscriptionsAPI.md#V1VolumePurchasingSubscriptionsIdPut) | **Put** /v1/volume-purchasing-subscriptions/{id} | Update a Volume Purchasing Subscription
[**V1VolumePurchasingSubscriptionsPost**](VolumePurchasingSubscriptionsAPI.md#V1VolumePurchasingSubscriptionsPost) | **Post** /v1/volume-purchasing-subscriptions | Create a Volume Purchasing Subscription



## V1VolumePurchasingSubscriptionsGet

> VolumePurchasingSubscriptions V1VolumePurchasingSubscriptionsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Execute()

Retrieve Volume Purchasing Subscriptions



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
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Allowable properties are id, name, and enabled. (optional) (default to ["id:asc"])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumePurchasingSubscriptionsAPI.V1VolumePurchasingSubscriptionsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumePurchasingSubscriptionsAPI.V1VolumePurchasingSubscriptionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1VolumePurchasingSubscriptionsGet`: VolumePurchasingSubscriptions
    fmt.Fprintf(os.Stdout, "Response from `VolumePurchasingSubscriptionsAPI.V1VolumePurchasingSubscriptionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1VolumePurchasingSubscriptionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Allowable properties are id, name, and enabled. | [default to [&quot;id:asc&quot;]]

### Return type

[**VolumePurchasingSubscriptions**](VolumePurchasingSubscriptions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1VolumePurchasingSubscriptionsIdDelete

> V1VolumePurchasingSubscriptionsIdDelete(ctx, id).Execute()

Delete a Volume Purchasing Subscription with the supplied id



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
    id := "id_example" // string | Volume Purchasing Subscription identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VolumePurchasingSubscriptionsAPI.V1VolumePurchasingSubscriptionsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumePurchasingSubscriptionsAPI.V1VolumePurchasingSubscriptionsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Volume Purchasing Subscription identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1VolumePurchasingSubscriptionsIdDeleteRequest struct via the builder pattern


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


## V1VolumePurchasingSubscriptionsIdGet

> VolumePurchasingSubscription V1VolumePurchasingSubscriptionsIdGet(ctx, id).Execute()

Retrieve a Volume Purchasing Subscription with the supplied id



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
    id := "id_example" // string | Volume Purchasing Subscription identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumePurchasingSubscriptionsAPI.V1VolumePurchasingSubscriptionsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumePurchasingSubscriptionsAPI.V1VolumePurchasingSubscriptionsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1VolumePurchasingSubscriptionsIdGet`: VolumePurchasingSubscription
    fmt.Fprintf(os.Stdout, "Response from `VolumePurchasingSubscriptionsAPI.V1VolumePurchasingSubscriptionsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Volume Purchasing Subscription identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1VolumePurchasingSubscriptionsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumePurchasingSubscription**](VolumePurchasingSubscription.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1VolumePurchasingSubscriptionsIdPut

> VolumePurchasingSubscription V1VolumePurchasingSubscriptionsIdPut(ctx, id).VolumePurchasingSubscriptionBase(volumePurchasingSubscriptionBase).Execute()

Update a Volume Purchasing Subscription



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
    id := "id_example" // string | Volume Purchasing Subscription identifier
    volumePurchasingSubscriptionBase := *openapiclient.NewVolumePurchasingSubscriptionBase("Example") // VolumePurchasingSubscriptionBase | Volume Purchasing Subscription to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumePurchasingSubscriptionsAPI.V1VolumePurchasingSubscriptionsIdPut(context.Background(), id).VolumePurchasingSubscriptionBase(volumePurchasingSubscriptionBase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumePurchasingSubscriptionsAPI.V1VolumePurchasingSubscriptionsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1VolumePurchasingSubscriptionsIdPut`: VolumePurchasingSubscription
    fmt.Fprintf(os.Stdout, "Response from `VolumePurchasingSubscriptionsAPI.V1VolumePurchasingSubscriptionsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Volume Purchasing Subscription identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1VolumePurchasingSubscriptionsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumePurchasingSubscriptionBase** | [**VolumePurchasingSubscriptionBase**](VolumePurchasingSubscriptionBase.md) | Volume Purchasing Subscription to update | 

### Return type

[**VolumePurchasingSubscription**](VolumePurchasingSubscription.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1VolumePurchasingSubscriptionsPost

> HrefResponse V1VolumePurchasingSubscriptionsPost(ctx).VolumePurchasingSubscriptionBase(volumePurchasingSubscriptionBase).Execute()

Create a Volume Purchasing Subscription



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
    volumePurchasingSubscriptionBase := *openapiclient.NewVolumePurchasingSubscriptionBase("Example") // VolumePurchasingSubscriptionBase | Volume Purchasing Subscription to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumePurchasingSubscriptionsAPI.V1VolumePurchasingSubscriptionsPost(context.Background()).VolumePurchasingSubscriptionBase(volumePurchasingSubscriptionBase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumePurchasingSubscriptionsAPI.V1VolumePurchasingSubscriptionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1VolumePurchasingSubscriptionsPost`: HrefResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumePurchasingSubscriptionsAPI.V1VolumePurchasingSubscriptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1VolumePurchasingSubscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumePurchasingSubscriptionBase** | [**VolumePurchasingSubscriptionBase**](VolumePurchasingSubscriptionBase.md) | Volume Purchasing Subscription to create | 

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

