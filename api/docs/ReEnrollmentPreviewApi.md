# \ReEnrollmentPreviewAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ReenrollmentGet**](ReEnrollmentPreviewAPI.md#V1ReenrollmentGet) | **Get** /v1/reenrollment | Get Re-enrollment object 
[**V1ReenrollmentHistoryExportPost**](ReEnrollmentPreviewAPI.md#V1ReenrollmentHistoryExportPost) | **Post** /v1/reenrollment/history/export | Export reenrollment history collection 
[**V1ReenrollmentHistoryGet**](ReEnrollmentPreviewAPI.md#V1ReenrollmentHistoryGet) | **Get** /v1/reenrollment/history | Get Re-enrollment history object 
[**V1ReenrollmentHistoryPost**](ReEnrollmentPreviewAPI.md#V1ReenrollmentHistoryPost) | **Post** /v1/reenrollment/history | Add specified Re-enrollment history object notes 
[**V1ReenrollmentPut**](ReEnrollmentPreviewAPI.md#V1ReenrollmentPut) | **Put** /v1/reenrollment | Update the Re-enrollment object 



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
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReEnrollmentPreviewAPI.V1ReenrollmentGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReEnrollmentPreviewAPI.V1ReenrollmentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ReenrollmentGet`: Reenrollment
    fmt.Fprintf(os.Stdout, "Response from `ReEnrollmentPreviewAPI.V1ReenrollmentGet`: %v\n", resp)
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


## V1ReenrollmentHistoryExportPost

> interface{} V1ReenrollmentHistoryExportPost(ctx).ExportFields(exportFields).ExportLabels(exportLabels).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).ExportParameters(exportParameters).Execute()

Export reenrollment history collection 



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
    exportFields := []string{"Inner_example"} // []string | Export fields parameter, used to change default order or ignore some of the response properties. Default is empty array, which means that all fields of the response entity will be serialized. Example: export-fields=id,username (optional) (default to [])
    exportLabels := []string{"Inner_example"} // []string | Export labels parameter, used to customize fieldnames/columns in the exported file. Default is empty array, which means that response properties names will be used. Number of the provided labels must match the number of export-fields Example: export-labels=identifier,name with matching: export-fields=id,username (optional) (default to [])
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=id:desc,name:asc  (optional) (default to ["id:asc"])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name. This param can be combined with paging and sorting. Example: name==\"*script*\" (optional) (default to "")
    exportParameters := *openapiclient.NewExportParameters() // ExportParameters | Optional. Override query parameters since they can make URI exceed 2,000 character limit. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReEnrollmentPreviewAPI.V1ReenrollmentHistoryExportPost(context.Background()).ExportFields(exportFields).ExportLabels(exportLabels).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).ExportParameters(exportParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReEnrollmentPreviewAPI.V1ReenrollmentHistoryExportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ReenrollmentHistoryExportPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `ReEnrollmentPreviewAPI.V1ReenrollmentHistoryExportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ReenrollmentHistoryExportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportFields** | **[]string** | Export fields parameter, used to change default order or ignore some of the response properties. Default is empty array, which means that all fields of the response entity will be serialized. Example: export-fields&#x3D;id,username | [default to []]
 **exportLabels** | **[]string** | Export labels parameter, used to customize fieldnames/columns in the exported file. Default is empty array, which means that response properties names will be used. Number of the provided labels must match the number of export-fields Example: export-labels&#x3D;identifier,name with matching: export-fields&#x3D;id,username | [default to []]
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;id:desc,name:asc  | [default to [&quot;id:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name. This param can be combined with paging and sorting. Example: name&#x3D;&#x3D;\&quot;*script*\&quot; | [default to &quot;&quot;]
 **exportParameters** | [**ExportParameters**](ExportParameters.md) | Optional. Override query parameters since they can make URI exceed 2,000 character limit. | 

### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/csv,application/json, application/json

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
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    page := int32(56) // int32 |  (optional) (default to 0)
    size := int32(56) // int32 |  (optional) (default to 100)
    pagesize := int32(56) // int32 |  (optional) (default to 100)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := "sort_example" // string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to "date:desc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReEnrollmentPreviewAPI.V1ReenrollmentHistoryGet(context.Background()).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReEnrollmentPreviewAPI.V1ReenrollmentHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ReenrollmentHistoryGet`: HistorySearchResults
    fmt.Fprintf(os.Stdout, "Response from `ReEnrollmentPreviewAPI.V1ReenrollmentHistoryGet`: %v\n", resp)
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
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | history notes to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReEnrollmentPreviewAPI.V1ReenrollmentHistoryPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReEnrollmentPreviewAPI.V1ReenrollmentHistoryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ReenrollmentHistoryPost`: ObjectHistory
    fmt.Fprintf(os.Stdout, "Response from `ReEnrollmentPreviewAPI.V1ReenrollmentHistoryPost`: %v\n", resp)
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
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    reenrollment := *openapiclient.NewReenrollment("DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED") // Reenrollment | Re-enrollment object to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReEnrollmentPreviewAPI.V1ReenrollmentPut(context.Background()).Reenrollment(reenrollment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReEnrollmentPreviewAPI.V1ReenrollmentPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ReenrollmentPut`: Reenrollment
    fmt.Fprintf(os.Stdout, "Response from `ReEnrollmentPreviewAPI.V1ReenrollmentPut`: %v\n", resp)
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

