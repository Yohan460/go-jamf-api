# \SelfServiceBrandingIosApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SelfServiceBrandingIosGet**](SelfServiceBrandingIosApi.md#V1SelfServiceBrandingIosGet) | **Get** /v1/self-service/branding/ios | Search for sorted and paged iOS branding configurations 
[**V1SelfServiceBrandingIosIdDelete**](SelfServiceBrandingIosApi.md#V1SelfServiceBrandingIosIdDelete) | **Delete** /v1/self-service/branding/ios/{id} | Delete the Self Service iOS branding configuration indicated by the provided id 
[**V1SelfServiceBrandingIosIdGet**](SelfServiceBrandingIosApi.md#V1SelfServiceBrandingIosIdGet) | **Get** /v1/self-service/branding/ios/{id} | Read a single Self Service iOS branding configuration indicated by the provided id 
[**V1SelfServiceBrandingIosIdPut**](SelfServiceBrandingIosApi.md#V1SelfServiceBrandingIosIdPut) | **Put** /v1/self-service/branding/ios/{id} | Update a Self Service iOS branding configuration with the supplied details 
[**V1SelfServiceBrandingIosPost**](SelfServiceBrandingIosApi.md#V1SelfServiceBrandingIosPost) | **Post** /v1/self-service/branding/ios | Create a Self Service iOS branding configuration with the supplied 



## V1SelfServiceBrandingIosGet

> IosBrandingSearchResults V1SelfServiceBrandingIosGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Execute()

Search for sorted and paged iOS branding configurations 



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
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=id:desc,brandingName:asc  (optional) (default to ["id:asc"])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SelfServiceBrandingIosApi.V1SelfServiceBrandingIosGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceBrandingIosApi.V1SelfServiceBrandingIosGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SelfServiceBrandingIosGet`: IosBrandingSearchResults
    fmt.Fprintf(os.Stdout, "Response from `SelfServiceBrandingIosApi.V1SelfServiceBrandingIosGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SelfServiceBrandingIosGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;id:desc,brandingName:asc  | [default to [&quot;id:asc&quot;]]

### Return type

[**IosBrandingSearchResults**](IosBrandingSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SelfServiceBrandingIosIdDelete

> V1SelfServiceBrandingIosIdDelete(ctx, id).Execute()

Delete the Self Service iOS branding configuration indicated by the provided id 



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
    id := "id_example" // string | id of iOS branding configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SelfServiceBrandingIosApi.V1SelfServiceBrandingIosIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceBrandingIosApi.V1SelfServiceBrandingIosIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of iOS branding configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SelfServiceBrandingIosIdDeleteRequest struct via the builder pattern


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


## V1SelfServiceBrandingIosIdGet

> IosBrandingConfiguration V1SelfServiceBrandingIosIdGet(ctx, id).Execute()

Read a single Self Service iOS branding configuration indicated by the provided id 



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
    id := "id_example" // string | id of iOS branding configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SelfServiceBrandingIosApi.V1SelfServiceBrandingIosIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceBrandingIosApi.V1SelfServiceBrandingIosIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SelfServiceBrandingIosIdGet`: IosBrandingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `SelfServiceBrandingIosApi.V1SelfServiceBrandingIosIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of iOS branding configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SelfServiceBrandingIosIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IosBrandingConfiguration**](IosBrandingConfiguration.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SelfServiceBrandingIosIdPut

> IosBrandingConfiguration V1SelfServiceBrandingIosIdPut(ctx, id).IosBrandingConfiguration(iosBrandingConfiguration).Execute()

Update a Self Service iOS branding configuration with the supplied details 



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
    id := "id_example" // string | id of iOS branding configuration
    iosBrandingConfiguration := *openapiclient.NewIosBrandingConfiguration("Self Service", "FFFFFF", "000001", "000000", "dark") // IosBrandingConfiguration | The iOS branding configuration values to update (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SelfServiceBrandingIosApi.V1SelfServiceBrandingIosIdPut(context.Background(), id).IosBrandingConfiguration(iosBrandingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceBrandingIosApi.V1SelfServiceBrandingIosIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SelfServiceBrandingIosIdPut`: IosBrandingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `SelfServiceBrandingIosApi.V1SelfServiceBrandingIosIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of iOS branding configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SelfServiceBrandingIosIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iosBrandingConfiguration** | [**IosBrandingConfiguration**](IosBrandingConfiguration.md) | The iOS branding configuration values to update | 

### Return type

[**IosBrandingConfiguration**](IosBrandingConfiguration.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SelfServiceBrandingIosPost

> HrefResponse V1SelfServiceBrandingIosPost(ctx).IosBrandingConfiguration(iosBrandingConfiguration).Execute()

Create a Self Service iOS branding configuration with the supplied 



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
    iosBrandingConfiguration := *openapiclient.NewIosBrandingConfiguration("Self Service", "FFFFFF", "000001", "000000", "dark") // IosBrandingConfiguration | The iOS branding configuration to create (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SelfServiceBrandingIosApi.V1SelfServiceBrandingIosPost(context.Background()).IosBrandingConfiguration(iosBrandingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceBrandingIosApi.V1SelfServiceBrandingIosPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SelfServiceBrandingIosPost`: HrefResponse
    fmt.Fprintf(os.Stdout, "Response from `SelfServiceBrandingIosApi.V1SelfServiceBrandingIosPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SelfServiceBrandingIosPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iosBrandingConfiguration** | [**IosBrandingConfiguration**](IosBrandingConfiguration.md) | The iOS branding configuration to create | 

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

