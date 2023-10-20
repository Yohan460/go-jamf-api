# \JamfProNotificationsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1NotificationsGet**](JamfProNotificationsAPI.md#V1NotificationsGet) | **Get** /v1/notifications | Get Notifications for user and site 
[**V1NotificationsTypeIdDelete**](JamfProNotificationsAPI.md#V1NotificationsTypeIdDelete) | **Delete** /v1/notifications/{type}/{id} | Delete Notifications 



## V1NotificationsGet

> []NotificationV1 V1NotificationsGet(ctx).Execute()

Get Notifications for user and site 



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
    resp, r, err := apiClient.JamfProNotificationsAPI.V1NotificationsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProNotificationsAPI.V1NotificationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1NotificationsGet`: []NotificationV1
    fmt.Fprintf(os.Stdout, "Response from `JamfProNotificationsAPI.V1NotificationsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1NotificationsGetRequest struct via the builder pattern


### Return type

[**[]NotificationV1**](NotificationV1.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1NotificationsTypeIdDelete

> V1NotificationsTypeIdDelete(ctx, id, type_).Execute()

Delete Notifications 



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
    id := "id_example" // string | instance ID of the notification
    type_ := openapiclient.NotificationType("APNS_CERT_REVOKED") // NotificationType | type of the notification

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.JamfProNotificationsAPI.V1NotificationsTypeIdDelete(context.Background(), id, type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProNotificationsAPI.V1NotificationsTypeIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance ID of the notification | 
**type_** | [**NotificationType**](.md) | type of the notification | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1NotificationsTypeIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

