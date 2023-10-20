# \JamfProNotificationsPreviewAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationsAlertsGet**](JamfProNotificationsPreviewAPI.md#NotificationsAlertsGet) | **Get** /notifications/alerts | Get Notifications for user and site 
[**NotificationsAlertsIdDelete**](JamfProNotificationsPreviewAPI.md#NotificationsAlertsIdDelete) | **Delete** /notifications/alerts/{id} | DEPRECATED - USE \&quot;alerts/{type}/{id}\&quot; INSTEAD. Deletes only Patch Management notifications. 
[**NotificationsAlertsTypeIdDelete**](JamfProNotificationsPreviewAPI.md#NotificationsAlertsTypeIdDelete) | **Delete** /notifications/alerts/{type}/{id} | Delete Notifications 



## NotificationsAlertsGet

> []Notification NotificationsAlertsGet(ctx).Execute()

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
    resp, r, err := apiClient.JamfProNotificationsPreviewAPI.NotificationsAlertsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProNotificationsPreviewAPI.NotificationsAlertsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NotificationsAlertsGet`: []Notification
    fmt.Fprintf(os.Stdout, "Response from `JamfProNotificationsPreviewAPI.NotificationsAlertsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsAlertsGetRequest struct via the builder pattern


### Return type

[**[]Notification**](Notification.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationsAlertsIdDelete

> NotificationsAlertsIdDelete(ctx, id).Execute()

DEPRECATED - USE \"alerts/{type}/{id}\" INSTEAD. Deletes only Patch Management notifications. 



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
    id := int32(56) // int32 | instance ID of the notification

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.JamfProNotificationsPreviewAPI.NotificationsAlertsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProNotificationsPreviewAPI.NotificationsAlertsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | instance ID of the notification | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsAlertsIdDeleteRequest struct via the builder pattern


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


## NotificationsAlertsTypeIdDelete

> NotificationsAlertsTypeIdDelete(ctx, id, type_).Execute()

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
    id := int32(56) // int32 | instance ID of the notification
    type_ := openapiclient.NotificationType("APNS_CERT_REVOKED") // NotificationType | type of the notification

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.JamfProNotificationsPreviewAPI.NotificationsAlertsTypeIdDelete(context.Background(), id, type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JamfProNotificationsPreviewAPI.NotificationsAlertsTypeIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | instance ID of the notification | 
**type_** | [**NotificationType**](.md) | type of the notification | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsAlertsTypeIdDeleteRequest struct via the builder pattern


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

