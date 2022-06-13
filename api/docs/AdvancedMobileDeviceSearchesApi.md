# \AdvancedMobileDeviceSearchesApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AdvancedMobileDeviceSearchesChoicesGet**](AdvancedMobileDeviceSearchesApi.md#V1AdvancedMobileDeviceSearchesChoicesGet) | **Get** /v1/advanced-mobile-device-searches/choices | Get Mobile Device Advanced Search criteria choices 
[**V1AdvancedMobileDeviceSearchesDeleteMultiplePost**](AdvancedMobileDeviceSearchesApi.md#V1AdvancedMobileDeviceSearchesDeleteMultiplePost) | **Post** /v1/advanced-mobile-device-searches/delete-multiple | Remove specified Advanced Search objects 
[**V1AdvancedMobileDeviceSearchesGet**](AdvancedMobileDeviceSearchesApi.md#V1AdvancedMobileDeviceSearchesGet) | **Get** /v1/advanced-mobile-device-searches | Get Advanced Search objects 
[**V1AdvancedMobileDeviceSearchesIdDelete**](AdvancedMobileDeviceSearchesApi.md#V1AdvancedMobileDeviceSearchesIdDelete) | **Delete** /v1/advanced-mobile-device-searches/{id} | Remove specified Advanced Search object 
[**V1AdvancedMobileDeviceSearchesIdGet**](AdvancedMobileDeviceSearchesApi.md#V1AdvancedMobileDeviceSearchesIdGet) | **Get** /v1/advanced-mobile-device-searches/{id} | Get specified Advanced Search object 
[**V1AdvancedMobileDeviceSearchesIdPut**](AdvancedMobileDeviceSearchesApi.md#V1AdvancedMobileDeviceSearchesIdPut) | **Put** /v1/advanced-mobile-device-searches/{id} | Get specified Advanced Search object 
[**V1AdvancedMobileDeviceSearchesPost**](AdvancedMobileDeviceSearchesApi.md#V1AdvancedMobileDeviceSearchesPost) | **Post** /v1/advanced-mobile-device-searches | Create Advanced Search object 



## V1AdvancedMobileDeviceSearchesChoicesGet

> AdvancedSearchCriteriaChoices V1AdvancedMobileDeviceSearchesChoicesGet(ctx).Criteria(criteria).Site(site).Contains(contains).Execute()

Get Mobile Device Advanced Search criteria choices 



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
    criteria := "criteria_example" // string | 
    site := "site_example" // string |  (optional) (default to "-1")
    contains := "contains_example" // string |  (optional) (default to "null")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdvancedMobileDeviceSearchesApi.V1AdvancedMobileDeviceSearchesChoicesGet(context.Background()).Criteria(criteria).Site(site).Contains(contains).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdvancedMobileDeviceSearchesApi.V1AdvancedMobileDeviceSearchesChoicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AdvancedMobileDeviceSearchesChoicesGet`: AdvancedSearchCriteriaChoices
    fmt.Fprintf(os.Stdout, "Response from `AdvancedMobileDeviceSearchesApi.V1AdvancedMobileDeviceSearchesChoicesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AdvancedMobileDeviceSearchesChoicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **criteria** | **string** |  | 
 **site** | **string** |  | [default to &quot;-1&quot;]
 **contains** | **string** |  | [default to &quot;null&quot;]

### Return type

[**AdvancedSearchCriteriaChoices**](AdvancedSearchCriteriaChoices.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AdvancedMobileDeviceSearchesDeleteMultiplePost

> V1AdvancedMobileDeviceSearchesDeleteMultiplePost(ctx).Ids(ids).Execute()

Remove specified Advanced Search objects 



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
    ids := *openapiclient.NewIds() // Ids | ids of the building to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdvancedMobileDeviceSearchesApi.V1AdvancedMobileDeviceSearchesDeleteMultiplePost(context.Background()).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdvancedMobileDeviceSearchesApi.V1AdvancedMobileDeviceSearchesDeleteMultiplePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AdvancedMobileDeviceSearchesDeleteMultiplePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**Ids**](Ids.md) | ids of the building to be deleted | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AdvancedMobileDeviceSearchesGet

> AdvancedSearchSearchResults V1AdvancedMobileDeviceSearchesGet(ctx).Execute()

Get Advanced Search objects 



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
    resp, r, err := apiClient.AdvancedMobileDeviceSearchesApi.V1AdvancedMobileDeviceSearchesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdvancedMobileDeviceSearchesApi.V1AdvancedMobileDeviceSearchesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AdvancedMobileDeviceSearchesGet`: AdvancedSearchSearchResults
    fmt.Fprintf(os.Stdout, "Response from `AdvancedMobileDeviceSearchesApi.V1AdvancedMobileDeviceSearchesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AdvancedMobileDeviceSearchesGetRequest struct via the builder pattern


### Return type

[**AdvancedSearchSearchResults**](AdvancedSearchSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AdvancedMobileDeviceSearchesIdDelete

> V1AdvancedMobileDeviceSearchesIdDelete(ctx, id).Execute()

Remove specified Advanced Search object 



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
    id := "id_example" // string | instance id of advanced search record

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdvancedMobileDeviceSearchesApi.V1AdvancedMobileDeviceSearchesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdvancedMobileDeviceSearchesApi.V1AdvancedMobileDeviceSearchesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of advanced search record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AdvancedMobileDeviceSearchesIdDeleteRequest struct via the builder pattern


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


## V1AdvancedMobileDeviceSearchesIdGet

> AdvancedSearch V1AdvancedMobileDeviceSearchesIdGet(ctx, id).Execute()

Get specified Advanced Search object 



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
    id := "id_example" // string | id of target Advanced Search

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdvancedMobileDeviceSearchesApi.V1AdvancedMobileDeviceSearchesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdvancedMobileDeviceSearchesApi.V1AdvancedMobileDeviceSearchesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AdvancedMobileDeviceSearchesIdGet`: AdvancedSearch
    fmt.Fprintf(os.Stdout, "Response from `AdvancedMobileDeviceSearchesApi.V1AdvancedMobileDeviceSearchesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of target Advanced Search | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AdvancedMobileDeviceSearchesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdvancedSearch**](AdvancedSearch.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AdvancedMobileDeviceSearchesIdPut

> AdvancedSearch V1AdvancedMobileDeviceSearchesIdPut(ctx, id).AdvancedSearch(advancedSearch).Execute()

Get specified Advanced Search object 



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
    id := "id_example" // string | id of target Advanced Search
    advancedSearch := *openapiclient.NewAdvancedSearch("Andy's Search") // AdvancedSearch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdvancedMobileDeviceSearchesApi.V1AdvancedMobileDeviceSearchesIdPut(context.Background(), id).AdvancedSearch(advancedSearch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdvancedMobileDeviceSearchesApi.V1AdvancedMobileDeviceSearchesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AdvancedMobileDeviceSearchesIdPut`: AdvancedSearch
    fmt.Fprintf(os.Stdout, "Response from `AdvancedMobileDeviceSearchesApi.V1AdvancedMobileDeviceSearchesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of target Advanced Search | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1AdvancedMobileDeviceSearchesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **advancedSearch** | [**AdvancedSearch**](AdvancedSearch.md) |  | 

### Return type

[**AdvancedSearch**](AdvancedSearch.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AdvancedMobileDeviceSearchesPost

> HrefResponse V1AdvancedMobileDeviceSearchesPost(ctx).AdvancedSearch(advancedSearch).Execute()

Create Advanced Search object 



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
    advancedSearch := *openapiclient.NewAdvancedSearch("Andy's Search") // AdvancedSearch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdvancedMobileDeviceSearchesApi.V1AdvancedMobileDeviceSearchesPost(context.Background()).AdvancedSearch(advancedSearch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdvancedMobileDeviceSearchesApi.V1AdvancedMobileDeviceSearchesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AdvancedMobileDeviceSearchesPost`: HrefResponse
    fmt.Fprintf(os.Stdout, "Response from `AdvancedMobileDeviceSearchesApi.V1AdvancedMobileDeviceSearchesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AdvancedMobileDeviceSearchesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **advancedSearch** | [**AdvancedSearch**](AdvancedSearch.md) |  | 

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

