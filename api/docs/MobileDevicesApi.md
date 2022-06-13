# \MobileDevicesApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1MobileDevicesGet**](MobileDevicesApi.md#V1MobileDevicesGet) | **Get** /v1/mobile-devices | Get Mobile Device objects 
[**V1MobileDevicesIdDetailGet**](MobileDevicesApi.md#V1MobileDevicesIdDetailGet) | **Get** /v1/mobile-devices/{id}/detail | Get Mobile Device 
[**V1MobileDevicesIdGet**](MobileDevicesApi.md#V1MobileDevicesIdGet) | **Get** /v1/mobile-devices/{id} | Get Mobile Device 
[**V1MobileDevicesIdPatch**](MobileDevicesApi.md#V1MobileDevicesIdPatch) | **Patch** /v1/mobile-devices/{id} | Update fields on a mobile device that are allowed to be modified by users 
[**V1SearchMobileDevicesPost**](MobileDevicesApi.md#V1SearchMobileDevicesPost) | **Post** /v1/search-mobile-devices | Search Mobile Devices 
[**V2MobileDevicesGet**](MobileDevicesApi.md#V2MobileDevicesGet) | **Get** /v2/mobile-devices | Get Mobile Device objects 
[**V2MobileDevicesIdDetailGet**](MobileDevicesApi.md#V2MobileDevicesIdDetailGet) | **Get** /v2/mobile-devices/{id}/detail | Get Mobile Device 
[**V2MobileDevicesIdGet**](MobileDevicesApi.md#V2MobileDevicesIdGet) | **Get** /v2/mobile-devices/{id} | Get Mobile Device 
[**V2MobileDevicesIdPatch**](MobileDevicesApi.md#V2MobileDevicesIdPatch) | **Patch** /v2/mobile-devices/{id} | Update fields on a mobile device that are allowed to be modified by users 



## V1MobileDevicesGet

> []MobileDevice V1MobileDevicesGet(ctx).Execute()

Get Mobile Device objects 



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
    resp, r, err := apiClient.MobileDevicesApi.V1MobileDevicesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDevicesApi.V1MobileDevicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1MobileDevicesGet`: []MobileDevice
    fmt.Fprintf(os.Stdout, "Response from `MobileDevicesApi.V1MobileDevicesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDevicesGetRequest struct via the builder pattern


### Return type

[**[]MobileDevice**](MobileDevice.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MobileDevicesIdDetailGet

> MobileDeviceDetails V1MobileDevicesIdDetailGet(ctx, id).Execute()

Get Mobile Device 



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
    id := int32(56) // int32 | instance id of mobile device record

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileDevicesApi.V1MobileDevicesIdDetailGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDevicesApi.V1MobileDevicesIdDetailGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1MobileDevicesIdDetailGet`: MobileDeviceDetails
    fmt.Fprintf(os.Stdout, "Response from `MobileDevicesApi.V1MobileDevicesIdDetailGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | instance id of mobile device record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDevicesIdDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MobileDeviceDetails**](MobileDeviceDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MobileDevicesIdGet

> MobileDevice V1MobileDevicesIdGet(ctx, id).Execute()

Get Mobile Device 



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
    id := int32(56) // int32 | instance id of mobile device record

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileDevicesApi.V1MobileDevicesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDevicesApi.V1MobileDevicesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1MobileDevicesIdGet`: MobileDevice
    fmt.Fprintf(os.Stdout, "Response from `MobileDevicesApi.V1MobileDevicesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | instance id of mobile device record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDevicesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MobileDevice**](MobileDevice.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MobileDevicesIdPatch

> MobileDeviceDetails V1MobileDevicesIdPatch(ctx, id).UpdateMobileDevice(updateMobileDevice).Execute()

Update fields on a mobile device that are allowed to be modified by users 



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
    id := int32(56) // int32 | instance id of mobile device record
    updateMobileDevice := *openapiclient.NewUpdateMobileDevice() // UpdateMobileDevice | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileDevicesApi.V1MobileDevicesIdPatch(context.Background(), id).UpdateMobileDevice(updateMobileDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDevicesApi.V1MobileDevicesIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1MobileDevicesIdPatch`: MobileDeviceDetails
    fmt.Fprintf(os.Stdout, "Response from `MobileDevicesApi.V1MobileDevicesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | instance id of mobile device record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDevicesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMobileDevice** | [**UpdateMobileDevice**](UpdateMobileDevice.md) |  | 

### Return type

[**MobileDeviceDetails**](MobileDeviceDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SearchMobileDevicesPost

> MobileDeviceSearchResults V1SearchMobileDevicesPost(ctx).MobileDeviceSearchParams(mobileDeviceSearchParams).Execute()

Search Mobile Devices 



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
    mobileDeviceSearchParams := *openapiclient.NewMobileDeviceSearchParams() // MobileDeviceSearchParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileDevicesApi.V1SearchMobileDevicesPost(context.Background()).MobileDeviceSearchParams(mobileDeviceSearchParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDevicesApi.V1SearchMobileDevicesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SearchMobileDevicesPost`: MobileDeviceSearchResults
    fmt.Fprintf(os.Stdout, "Response from `MobileDevicesApi.V1SearchMobileDevicesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SearchMobileDevicesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mobileDeviceSearchParams** | [**MobileDeviceSearchParams**](MobileDeviceSearchParams.md) |  | 

### Return type

[**MobileDeviceSearchResults**](MobileDeviceSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2MobileDevicesGet

> MobileDeviceSearchResultsV2 V2MobileDevicesGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Execute()

Get Mobile Device objects 



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
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["id:asc"])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileDevicesApi.V2MobileDevicesGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDevicesApi.V2MobileDevicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2MobileDevicesGet`: MobileDeviceSearchResultsV2
    fmt.Fprintf(os.Stdout, "Response from `MobileDevicesApi.V2MobileDevicesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2MobileDevicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;id:asc&quot;]]

### Return type

[**MobileDeviceSearchResultsV2**](MobileDeviceSearchResultsV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2MobileDevicesIdDetailGet

> MobileDeviceDetailsGetV2 V2MobileDevicesIdDetailGet(ctx, id).Execute()

Get Mobile Device 



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
    id := "id_example" // string | instance id of mobile device record

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileDevicesApi.V2MobileDevicesIdDetailGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDevicesApi.V2MobileDevicesIdDetailGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2MobileDevicesIdDetailGet`: MobileDeviceDetailsGetV2
    fmt.Fprintf(os.Stdout, "Response from `MobileDevicesApi.V2MobileDevicesIdDetailGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of mobile device record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2MobileDevicesIdDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MobileDeviceDetailsGetV2**](MobileDeviceDetailsGetV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2MobileDevicesIdGet

> MobileDeviceV2 V2MobileDevicesIdGet(ctx, id).Execute()

Get Mobile Device 



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
    id := "id_example" // string | instance id of mobile device record

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileDevicesApi.V2MobileDevicesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDevicesApi.V2MobileDevicesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2MobileDevicesIdGet`: MobileDeviceV2
    fmt.Fprintf(os.Stdout, "Response from `MobileDevicesApi.V2MobileDevicesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of mobile device record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2MobileDevicesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MobileDeviceV2**](MobileDeviceV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2MobileDevicesIdPatch

> MobileDeviceDetailsV2 V2MobileDevicesIdPatch(ctx, id).UpdateMobileDeviceV2(updateMobileDeviceV2).Execute()

Update fields on a mobile device that are allowed to be modified by users 



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
    id := "id_example" // string | instance id of mobile device record
    updateMobileDeviceV2 := *openapiclient.NewUpdateMobileDeviceV2() // UpdateMobileDeviceV2 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileDevicesApi.V2MobileDevicesIdPatch(context.Background(), id).UpdateMobileDeviceV2(updateMobileDeviceV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDevicesApi.V2MobileDevicesIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2MobileDevicesIdPatch`: MobileDeviceDetailsV2
    fmt.Fprintf(os.Stdout, "Response from `MobileDevicesApi.V2MobileDevicesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of mobile device record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2MobileDevicesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMobileDeviceV2** | [**UpdateMobileDeviceV2**](UpdateMobileDeviceV2.md) |  | 

### Return type

[**MobileDeviceDetailsV2**](MobileDeviceDetailsV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

