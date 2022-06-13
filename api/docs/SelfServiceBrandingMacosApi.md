# \SelfServiceBrandingMacosApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SelfServiceBrandingMacosGet**](SelfServiceBrandingMacosApi.md#V1SelfServiceBrandingMacosGet) | **Get** /v1/self-service/branding/macos | Search for sorted and paged macOS branding configurations 
[**V1SelfServiceBrandingMacosIdDelete**](SelfServiceBrandingMacosApi.md#V1SelfServiceBrandingMacosIdDelete) | **Delete** /v1/self-service/branding/macos/{id} | Delete the Self Service macOS branding configuration indicated by the provided id 
[**V1SelfServiceBrandingMacosIdGet**](SelfServiceBrandingMacosApi.md#V1SelfServiceBrandingMacosIdGet) | **Get** /v1/self-service/branding/macos/{id} | Read a single Self Service macOS branding configuration indicated by the provided id 
[**V1SelfServiceBrandingMacosIdPut**](SelfServiceBrandingMacosApi.md#V1SelfServiceBrandingMacosIdPut) | **Put** /v1/self-service/branding/macos/{id} | Update a Self Service macOS branding configuration with the supplied details 
[**V1SelfServiceBrandingMacosPost**](SelfServiceBrandingMacosApi.md#V1SelfServiceBrandingMacosPost) | **Post** /v1/self-service/branding/macos | Create a Self Service macOS branding configuration with the supplied 



## V1SelfServiceBrandingMacosGet

> MacOsBrandingSearchResults V1SelfServiceBrandingMacosGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Execute()

Search for sorted and paged macOS branding configurations 



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
    resp, r, err := apiClient.SelfServiceBrandingMacosApi.V1SelfServiceBrandingMacosGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceBrandingMacosApi.V1SelfServiceBrandingMacosGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SelfServiceBrandingMacosGet`: MacOsBrandingSearchResults
    fmt.Fprintf(os.Stdout, "Response from `SelfServiceBrandingMacosApi.V1SelfServiceBrandingMacosGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SelfServiceBrandingMacosGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;id:desc,brandingName:asc  | [default to [&quot;id:asc&quot;]]

### Return type

[**MacOsBrandingSearchResults**](MacOsBrandingSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SelfServiceBrandingMacosIdDelete

> V1SelfServiceBrandingMacosIdDelete(ctx, id).Execute()

Delete the Self Service macOS branding configuration indicated by the provided id 



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
    id := "id_example" // string | id of macOS branding configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SelfServiceBrandingMacosApi.V1SelfServiceBrandingMacosIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceBrandingMacosApi.V1SelfServiceBrandingMacosIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of macOS branding configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SelfServiceBrandingMacosIdDeleteRequest struct via the builder pattern


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


## V1SelfServiceBrandingMacosIdGet

> MacOsBrandingConfiguration V1SelfServiceBrandingMacosIdGet(ctx, id).Execute()

Read a single Self Service macOS branding configuration indicated by the provided id 



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
    id := "id_example" // string | id of macOS branding configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SelfServiceBrandingMacosApi.V1SelfServiceBrandingMacosIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceBrandingMacosApi.V1SelfServiceBrandingMacosIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SelfServiceBrandingMacosIdGet`: MacOsBrandingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `SelfServiceBrandingMacosApi.V1SelfServiceBrandingMacosIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of macOS branding configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SelfServiceBrandingMacosIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MacOsBrandingConfiguration**](MacOsBrandingConfiguration.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SelfServiceBrandingMacosIdPut

> MacOsBrandingConfiguration V1SelfServiceBrandingMacosIdPut(ctx, id).MacOsBrandingConfiguration(macOsBrandingConfiguration).Execute()

Update a Self Service macOS branding configuration with the supplied details 



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
    id := "id_example" // string | id of macOS branding configuration
    macOsBrandingConfiguration := *openapiclient.NewMacOsBrandingConfiguration() // MacOsBrandingConfiguration | The macOS branding configuration values to update (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SelfServiceBrandingMacosApi.V1SelfServiceBrandingMacosIdPut(context.Background(), id).MacOsBrandingConfiguration(macOsBrandingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceBrandingMacosApi.V1SelfServiceBrandingMacosIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SelfServiceBrandingMacosIdPut`: MacOsBrandingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `SelfServiceBrandingMacosApi.V1SelfServiceBrandingMacosIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of macOS branding configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SelfServiceBrandingMacosIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **macOsBrandingConfiguration** | [**MacOsBrandingConfiguration**](MacOsBrandingConfiguration.md) | The macOS branding configuration values to update | 

### Return type

[**MacOsBrandingConfiguration**](MacOsBrandingConfiguration.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SelfServiceBrandingMacosPost

> HrefResponse V1SelfServiceBrandingMacosPost(ctx).MacOsBrandingConfiguration(macOsBrandingConfiguration).Execute()

Create a Self Service macOS branding configuration with the supplied 



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
    macOsBrandingConfiguration := *openapiclient.NewMacOsBrandingConfiguration() // MacOsBrandingConfiguration | The macOS branding configuration to create (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SelfServiceBrandingMacosApi.V1SelfServiceBrandingMacosPost(context.Background()).MacOsBrandingConfiguration(macOsBrandingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceBrandingMacosApi.V1SelfServiceBrandingMacosPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SelfServiceBrandingMacosPost`: HrefResponse
    fmt.Fprintf(os.Stdout, "Response from `SelfServiceBrandingMacosApi.V1SelfServiceBrandingMacosPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SelfServiceBrandingMacosPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **macOsBrandingConfiguration** | [**MacOsBrandingConfiguration**](MacOsBrandingConfiguration.md) | The macOS branding configuration to create | 

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

