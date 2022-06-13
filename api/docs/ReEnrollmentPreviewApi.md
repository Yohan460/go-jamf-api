# \ReEnrollmentPreviewApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ReenrollmentGet**](ReEnrollmentPreviewApi.md#V1ReenrollmentGet) | **Get** /v1/reenrollment | Get Re-enrollment object 
[**V1ReenrollmentHistoryGet**](ReEnrollmentPreviewApi.md#V1ReenrollmentHistoryGet) | **Get** /v1/reenrollment/history | Get Re-enrollment history object 
[**V1ReenrollmentHistoryPost**](ReEnrollmentPreviewApi.md#V1ReenrollmentHistoryPost) | **Post** /v1/reenrollment/history | Add specified Re-enrollment history object notes 
[**V1ReenrollmentPut**](ReEnrollmentPreviewApi.md#V1ReenrollmentPut) | **Put** /v1/reenrollment | Update the Re-enrollment object 



## V1ReenrollmentGet

> Reenrollment V1ReenrollmentGet(ctx).Execute()

Get Re-enrollment object 



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
    resp, r, err := apiClient.ReEnrollmentPreviewApi.V1ReenrollmentGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReEnrollmentPreviewApi.V1ReenrollmentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ReenrollmentGet`: Reenrollment
    fmt.Fprintf(os.Stdout, "Response from `ReEnrollmentPreviewApi.V1ReenrollmentGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ReenrollmentGetRequest struct via the builder pattern


### Return type

[**Reenrollment**](Reenrollment.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ReenrollmentHistoryGet

> HistorySearchResults V1ReenrollmentHistoryGet(ctx).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()

Get Re-enrollment history object 



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
    size := int32(56) // int32 |  (optional) (default to 100)
    pagesize := int32(56) // int32 |  (optional) (default to 100)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := "sort_example" // string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to "date:desc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReEnrollmentPreviewApi.V1ReenrollmentHistoryGet(context.Background()).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReEnrollmentPreviewApi.V1ReenrollmentHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ReenrollmentHistoryGet`: HistorySearchResults
    fmt.Fprintf(os.Stdout, "Response from `ReEnrollmentPreviewApi.V1ReenrollmentHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ReenrollmentHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]
 **pagesize** | **int32** |  | [default to 100]
 **pageSize** | **int32** |  | [default to 100]
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


## V1ReenrollmentHistoryPost

> ObjectHistory V1ReenrollmentHistoryPost(ctx).ObjectHistoryNote(objectHistoryNote).Execute()

Add specified Re-enrollment history object notes 



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
    objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | history notes to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReEnrollmentPreviewApi.V1ReenrollmentHistoryPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReEnrollmentPreviewApi.V1ReenrollmentHistoryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ReenrollmentHistoryPost`: ObjectHistory
    fmt.Fprintf(os.Stdout, "Response from `ReEnrollmentPreviewApi.V1ReenrollmentHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ReenrollmentHistoryPostRequest struct via the builder pattern


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


## V1ReenrollmentPut

> Reenrollment V1ReenrollmentPut(ctx).Reenrollment(reenrollment).Execute()

Update the Re-enrollment object 



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
    reenrollment := *openapiclient.NewReenrollment("DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED") // Reenrollment | Re-enrollment object to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReEnrollmentPreviewApi.V1ReenrollmentPut(context.Background()).Reenrollment(reenrollment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReEnrollmentPreviewApi.V1ReenrollmentPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ReenrollmentPut`: Reenrollment
    fmt.Fprintf(os.Stdout, "Response from `ReEnrollmentPreviewApi.V1ReenrollmentPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ReenrollmentPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reenrollment** | [**Reenrollment**](Reenrollment.md) | Re-enrollment object to update | 

### Return type

[**Reenrollment**](Reenrollment.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

