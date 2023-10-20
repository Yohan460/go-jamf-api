# \ParentAppPreviewAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ParentAppGet**](ParentAppPreviewAPI.md#V1ParentAppGet) | **Get** /v1/parent-app | Get the current Jamf Parent app settings 
[**V1ParentAppHistoryGet**](ParentAppPreviewAPI.md#V1ParentAppHistoryGet) | **Get** /v1/parent-app/history | Get Jamf Parent app settings history 
[**V1ParentAppHistoryPost**](ParentAppPreviewAPI.md#V1ParentAppHistoryPost) | **Post** /v1/parent-app/history | Add Jamf Parent app settings history notes 
[**V1ParentAppPut**](ParentAppPreviewAPI.md#V1ParentAppPut) | **Put** /v1/parent-app | Update Jamf Parent app settings 



## V1ParentAppGet

> ParentApp V1ParentAppGet(ctx).Execute()

Get the current Jamf Parent app settings 



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
    resp, r, err := apiClient.ParentAppPreviewAPI.V1ParentAppGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParentAppPreviewAPI.V1ParentAppGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ParentAppGet`: ParentApp
    fmt.Fprintf(os.Stdout, "Response from `ParentAppPreviewAPI.V1ParentAppGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ParentAppGetRequest struct via the builder pattern


### Return type

[**ParentApp**](ParentApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ParentAppHistoryGet

> HistorySearchResults V1ParentAppHistoryGet(ctx).Page(page).PageSize(pageSize).Filter(filter).Sort(sort).Execute()

Get Jamf Parent app settings history 



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
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")
    sort := "sort_example" // string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to "date:desc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ParentAppPreviewAPI.V1ParentAppHistoryGet(context.Background()).Page(page).PageSize(pageSize).Filter(filter).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParentAppPreviewAPI.V1ParentAppHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ParentAppHistoryGet`: HistorySearchResults
    fmt.Fprintf(os.Stdout, "Response from `ParentAppPreviewAPI.V1ParentAppHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ParentAppHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **filter** | **string** | Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15 | [default to &quot;&quot;]
 **sort** | **string** | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to &quot;date:desc&quot;]

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


## V1ParentAppHistoryPost

> ObjectHistory V1ParentAppHistoryPost(ctx).ObjectHistoryNote(objectHistoryNote).Execute()

Add Jamf Parent app settings history notes 



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
    objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | history notes to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ParentAppPreviewAPI.V1ParentAppHistoryPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParentAppPreviewAPI.V1ParentAppHistoryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ParentAppHistoryPost`: ObjectHistory
    fmt.Fprintf(os.Stdout, "Response from `ParentAppPreviewAPI.V1ParentAppHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ParentAppHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectHistoryNote** | [**ObjectHistoryNote**](ObjectHistoryNote.md) | history notes to create | 

### Return type

[**ObjectHistory**](ObjectHistory.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ParentAppPut

> ParentApp V1ParentAppPut(ctx).ParentApp(parentApp).Execute()

Update Jamf Parent app settings 



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
    parentApp := *openapiclient.NewParentApp("Europe/Paris", *openapiclient.NewParentAppRestrictedTimes(), int32(1), true) // ParentApp | Jamf Parent app settings to save.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ParentAppPreviewAPI.V1ParentAppPut(context.Background()).ParentApp(parentApp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParentAppPreviewAPI.V1ParentAppPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ParentAppPut`: ParentApp
    fmt.Fprintf(os.Stdout, "Response from `ParentAppPreviewAPI.V1ParentAppPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ParentAppPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentApp** | [**ParentApp**](ParentApp.md) | Jamf Parent app settings to save. | 

### Return type

[**ParentApp**](ParentApp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

