# \VolumePurchasingLocationsApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1VolumePurchasingLocationsGet**](VolumePurchasingLocationsApi.md#V1VolumePurchasingLocationsGet) | **Get** /v1/volume-purchasing-locations | Retrieve Volume Purchasing Locations
[**V1VolumePurchasingLocationsIdDelete**](VolumePurchasingLocationsApi.md#V1VolumePurchasingLocationsIdDelete) | **Delete** /v1/volume-purchasing-locations/{id} | Delete a Volume Purchasing Location with the supplied id
[**V1VolumePurchasingLocationsIdGet**](VolumePurchasingLocationsApi.md#V1VolumePurchasingLocationsIdGet) | **Get** /v1/volume-purchasing-locations/{id} | Retrieve a Volume Purchasing Location with the supplied id
[**V1VolumePurchasingLocationsIdHistoryGet**](VolumePurchasingLocationsApi.md#V1VolumePurchasingLocationsIdHistoryGet) | **Get** /v1/volume-purchasing-locations/{id}/history | Get specified Volume Purchasing Location history object 
[**V1VolumePurchasingLocationsIdPatch**](VolumePurchasingLocationsApi.md#V1VolumePurchasingLocationsIdPatch) | **Patch** /v1/volume-purchasing-locations/{id} | Update a Volume Purchasing Location
[**V1VolumePurchasingLocationsIdReclaimPost**](VolumePurchasingLocationsApi.md#V1VolumePurchasingLocationsIdReclaimPost) | **Post** /v1/volume-purchasing-locations/{id}/reclaim | Reclaim a Volume Purchasing Location with the supplied id
[**V1VolumePurchasingLocationsIdRevokeLicensesPost**](VolumePurchasingLocationsApi.md#V1VolumePurchasingLocationsIdRevokeLicensesPost) | **Post** /v1/volume-purchasing-locations/{id}/revoke-licenses | Revoke licenses for a Volume Purchasing Location with the supplied id
[**V1VolumePurchasingLocationsPost**](VolumePurchasingLocationsApi.md#V1VolumePurchasingLocationsPost) | **Post** /v1/volume-purchasing-locations | Create a Volume Purchasing Location



## V1VolumePurchasingLocationsGet

> VolumePurchasingLocations V1VolumePurchasingLocationsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Retrieve Volume Purchasing Locations



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
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. (optional) (default to ["id:asc"])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter Volume Purchasing Location collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name, appleId, organizationName, tokenExpiration, countryCode, locationName, automaticallyPopulatePurchasedContent, and sendNotificationWhenNoLongerAssigned. This param can be combined with paging and sorting. (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumePurchasingLocationsApi.V1VolumePurchasingLocationsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumePurchasingLocationsApi.V1VolumePurchasingLocationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1VolumePurchasingLocationsGet`: VolumePurchasingLocations
    fmt.Fprintf(os.Stdout, "Response from `VolumePurchasingLocationsApi.V1VolumePurchasingLocationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1VolumePurchasingLocationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. | [default to [&quot;id:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter Volume Purchasing Location collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name, appleId, organizationName, tokenExpiration, countryCode, locationName, automaticallyPopulatePurchasedContent, and sendNotificationWhenNoLongerAssigned. This param can be combined with paging and sorting. | [default to &quot;&quot;]

### Return type

[**VolumePurchasingLocations**](VolumePurchasingLocations.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1VolumePurchasingLocationsIdDelete

> V1VolumePurchasingLocationsIdDelete(ctx, id).Execute()

Delete a Volume Purchasing Location with the supplied id



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
    id := "id_example" // string | Volume Purchasing Location identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumePurchasingLocationsApi.V1VolumePurchasingLocationsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumePurchasingLocationsApi.V1VolumePurchasingLocationsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Volume Purchasing Location identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1VolumePurchasingLocationsIdDeleteRequest struct via the builder pattern


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


## V1VolumePurchasingLocationsIdGet

> VolumePurchasingLocation V1VolumePurchasingLocationsIdGet(ctx, id).Execute()

Retrieve a Volume Purchasing Location with the supplied id



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
    id := "id_example" // string | Volume Purchasing Location identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumePurchasingLocationsApi.V1VolumePurchasingLocationsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumePurchasingLocationsApi.V1VolumePurchasingLocationsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1VolumePurchasingLocationsIdGet`: VolumePurchasingLocation
    fmt.Fprintf(os.Stdout, "Response from `VolumePurchasingLocationsApi.V1VolumePurchasingLocationsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Volume Purchasing Location identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1VolumePurchasingLocationsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumePurchasingLocation**](VolumePurchasingLocation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1VolumePurchasingLocationsIdHistoryGet

> HistorySearchResults V1VolumePurchasingLocationsIdHistoryGet(ctx, id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get specified Volume Purchasing Location history object 



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
    id := "id_example" // string | instance id of Volume Purchasing Location history record
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma.  (optional) (default to ["date:desc"])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumePurchasingLocationsApi.V1VolumePurchasingLocationsIdHistoryGet(context.Background(), id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumePurchasingLocationsApi.V1VolumePurchasingLocationsIdHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1VolumePurchasingLocationsIdHistoryGet`: HistorySearchResults
    fmt.Fprintf(os.Stdout, "Response from `VolumePurchasingLocationsApi.V1VolumePurchasingLocationsIdHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of Volume Purchasing Location history record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1VolumePurchasingLocationsIdHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma.  | [default to [&quot;date:desc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15 | [default to &quot;&quot;]

### Return type

[**HistorySearchResults**](HistorySearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1VolumePurchasingLocationsIdPatch

> VolumePurchasingLocation V1VolumePurchasingLocationsIdPatch(ctx, id).VolumePurchasingLocationPatch(volumePurchasingLocationPatch).Execute()

Update a Volume Purchasing Location



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
    id := "id_example" // string | Volume Purchasing Location identifier
    volumePurchasingLocationPatch := *openapiclient.NewVolumePurchasingLocationPatch() // VolumePurchasingLocationPatch | Volume Purchasing Location to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumePurchasingLocationsApi.V1VolumePurchasingLocationsIdPatch(context.Background(), id).VolumePurchasingLocationPatch(volumePurchasingLocationPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumePurchasingLocationsApi.V1VolumePurchasingLocationsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1VolumePurchasingLocationsIdPatch`: VolumePurchasingLocation
    fmt.Fprintf(os.Stdout, "Response from `VolumePurchasingLocationsApi.V1VolumePurchasingLocationsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Volume Purchasing Location identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1VolumePurchasingLocationsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumePurchasingLocationPatch** | [**VolumePurchasingLocationPatch**](VolumePurchasingLocationPatch.md) | Volume Purchasing Location to update | 

### Return type

[**VolumePurchasingLocation**](VolumePurchasingLocation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1VolumePurchasingLocationsIdReclaimPost

> V1VolumePurchasingLocationsIdReclaimPost(ctx, id).Execute()

Reclaim a Volume Purchasing Location with the supplied id



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
    id := "id_example" // string | Volume Purchasing Location identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumePurchasingLocationsApi.V1VolumePurchasingLocationsIdReclaimPost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumePurchasingLocationsApi.V1VolumePurchasingLocationsIdReclaimPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Volume Purchasing Location identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1VolumePurchasingLocationsIdReclaimPostRequest struct via the builder pattern


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


## V1VolumePurchasingLocationsIdRevokeLicensesPost

> V1VolumePurchasingLocationsIdRevokeLicensesPost(ctx, id).Execute()

Revoke licenses for a Volume Purchasing Location with the supplied id



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
    id := "id_example" // string | Volume Purchasing Location identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumePurchasingLocationsApi.V1VolumePurchasingLocationsIdRevokeLicensesPost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumePurchasingLocationsApi.V1VolumePurchasingLocationsIdRevokeLicensesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Volume Purchasing Location identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1VolumePurchasingLocationsIdRevokeLicensesPostRequest struct via the builder pattern


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


## V1VolumePurchasingLocationsPost

> HrefResponse V1VolumePurchasingLocationsPost(ctx).VolumePurchasingLocationPost(volumePurchasingLocationPost).Execute()

Create a Volume Purchasing Location



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
    volumePurchasingLocationPost := *openapiclient.NewVolumePurchasingLocationPost("eyJleHBEYXRlIjoiMjAyMi0wMy0yOVQxNTozNjoyNiswMDAwIiwidG9rZW4iOiJWR2hwY3lCcGN5QnViM1FnWVNCMGIydGxiaTRnU0c5d1pXWjFiR3g1SUdsMElHeHZiMnR6SUd4cGEyVWdZU0IwYjJ0bGJpd2dZblYwSUdsMEozTWdibTkwTGc9PSIsIm9yZ05hbWUiOiJFeGFtcGxlIE9yZyJ9") // VolumePurchasingLocationPost | Volume Purchasing Location to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumePurchasingLocationsApi.V1VolumePurchasingLocationsPost(context.Background()).VolumePurchasingLocationPost(volumePurchasingLocationPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumePurchasingLocationsApi.V1VolumePurchasingLocationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1VolumePurchasingLocationsPost`: HrefResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumePurchasingLocationsApi.V1VolumePurchasingLocationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1VolumePurchasingLocationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumePurchasingLocationPost** | [**VolumePurchasingLocationPost**](VolumePurchasingLocationPost.md) | Volume Purchasing Location to create | 

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

