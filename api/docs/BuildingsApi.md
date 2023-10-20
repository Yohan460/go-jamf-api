# \BuildingsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1BuildingsDeleteMultiplePost**](BuildingsAPI.md#V1BuildingsDeleteMultiplePost) | **Post** /v1/buildings/delete-multiple | Delete multiple Buildings by their ids 
[**V1BuildingsExportPost**](BuildingsAPI.md#V1BuildingsExportPost) | **Post** /v1/buildings/export | Export Buildings collection 
[**V1BuildingsGet**](BuildingsAPI.md#V1BuildingsGet) | **Get** /v1/buildings | Search for sorted and paged Buildings 
[**V1BuildingsIdDelete**](BuildingsAPI.md#V1BuildingsIdDelete) | **Delete** /v1/buildings/{id} | Remove specified Building record 
[**V1BuildingsIdGet**](BuildingsAPI.md#V1BuildingsIdGet) | **Get** /v1/buildings/{id} | Get specified Building object 
[**V1BuildingsIdHistoryExportPost**](BuildingsAPI.md#V1BuildingsIdHistoryExportPost) | **Post** /v1/buildings/{id}/history/export | Export history object collection in specified format for specified Buildings 
[**V1BuildingsIdHistoryGet**](BuildingsAPI.md#V1BuildingsIdHistoryGet) | **Get** /v1/buildings/{id}/history | Get specified Building History object 
[**V1BuildingsIdHistoryPost**](BuildingsAPI.md#V1BuildingsIdHistoryPost) | **Post** /v1/buildings/{id}/history | Add specified Building history object notes 
[**V1BuildingsIdPut**](BuildingsAPI.md#V1BuildingsIdPut) | **Put** /v1/buildings/{id} | Update specified Building object 
[**V1BuildingsPost**](BuildingsAPI.md#V1BuildingsPost) | **Post** /v1/buildings | Create Building record 



## V1BuildingsDeleteMultiplePost

> V1BuildingsDeleteMultiplePost(ctx).Ids(ids).Execute()

Delete multiple Buildings by their ids 



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
    ids := *openapiclient.NewIds() // Ids | ids of the building to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BuildingsAPI.V1BuildingsDeleteMultiplePost(context.Background()).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildingsAPI.V1BuildingsDeleteMultiplePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1BuildingsDeleteMultiplePostRequest struct via the builder pattern


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


## V1BuildingsExportPost

> interface{} V1BuildingsExportPost(ctx).ExportFields(exportFields).ExportLabels(exportLabels).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).ExportParameters(exportParameters).Execute()

Export Buildings collection 



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
    exportFields := []string{"Inner_example"} // []string | Export fields parameter, used to change default order or ignore some of the response properties. Default is empty array, which means that all fields of the response entity will be serialized. Example: export-fields=id,username (optional) (default to [])
    exportLabels := []string{"Inner_example"} // []string | Export labels parameter, used to customize fieldnames/columns in the exported file. Default is empty array, which means that response properties names will be used. Number of the provided labels must match the number of export-fields Example: export-labels=identifier,name with matching: export-fields=id,username (optional) (default to [])
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=id:desc,name:asc  (optional) (default to ["id:asc"])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name. This param can be combined with paging and sorting. Example: name==\"*buildings*\" (optional) (default to "")
    exportParameters := *openapiclient.NewExportParameters() // ExportParameters | Optional. Override query parameters since they can make URI exceed 2,000 character limit. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildingsAPI.V1BuildingsExportPost(context.Background()).ExportFields(exportFields).ExportLabels(exportLabels).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).ExportParameters(exportParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildingsAPI.V1BuildingsExportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BuildingsExportPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `BuildingsAPI.V1BuildingsExportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1BuildingsExportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportFields** | **[]string** | Export fields parameter, used to change default order or ignore some of the response properties. Default is empty array, which means that all fields of the response entity will be serialized. Example: export-fields&#x3D;id,username | [default to []]
 **exportLabels** | **[]string** | Export labels parameter, used to customize fieldnames/columns in the exported file. Default is empty array, which means that response properties names will be used. Number of the provided labels must match the number of export-fields Example: export-labels&#x3D;identifier,name with matching: export-fields&#x3D;id,username | [default to []]
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;id:desc,name:asc  | [default to [&quot;id:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name. This param can be combined with paging and sorting. Example: name&#x3D;&#x3D;\&quot;*buildings*\&quot; | [default to &quot;&quot;]
 **exportParameters** | [**ExportParameters**](ExportParameters.md) | Optional. Override query parameters since they can make URI exceed 2,000 character limit. | 

### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BuildingsGet

> BuildingSearchResults V1BuildingsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Search for sorted and paged Buildings 



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
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["id:asc"])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter buildings collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: name, streetAddress1, streetAddress2, city, stateProvince, zipPostalCode, country. This param can be combined with paging and sorting. Example: filter=city==\"Chicago\" and name==\"*build*\" (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildingsAPI.V1BuildingsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildingsAPI.V1BuildingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BuildingsGet`: BuildingSearchResults
    fmt.Fprintf(os.Stdout, "Response from `BuildingsAPI.V1BuildingsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1BuildingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;id:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter buildings collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: name, streetAddress1, streetAddress2, city, stateProvince, zipPostalCode, country. This param can be combined with paging and sorting. Example: filter&#x3D;city&#x3D;&#x3D;\&quot;Chicago\&quot; and name&#x3D;&#x3D;\&quot;*build*\&quot; | [default to &quot;&quot;]

### Return type

[**BuildingSearchResults**](BuildingSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BuildingsIdDelete

> V1BuildingsIdDelete(ctx, id).Execute()

Remove specified Building record 



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
    id := "id_example" // string | instance id of building record

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BuildingsAPI.V1BuildingsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildingsAPI.V1BuildingsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of building record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BuildingsIdDeleteRequest struct via the builder pattern


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


## V1BuildingsIdGet

> Building V1BuildingsIdGet(ctx, id).Execute()

Get specified Building object 



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
    id := "id_example" // string | instance id of building record

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildingsAPI.V1BuildingsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildingsAPI.V1BuildingsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BuildingsIdGet`: Building
    fmt.Fprintf(os.Stdout, "Response from `BuildingsAPI.V1BuildingsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of building record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BuildingsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Building**](Building.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BuildingsIdHistoryExportPost

> interface{} V1BuildingsIdHistoryExportPost(ctx, id).ExportFields(exportFields).ExportLabels(exportLabels).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).ExportParameters(exportParameters).Execute()

Export history object collection in specified format for specified Buildings 



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
    id := "id_example" // string | instance id of buildings
    exportFields := []string{"Inner_example"} // []string | Export fields parameter, used to change default order or ignore some of the response properties. Default is empty array, which means that all fields of the response entity will be serialized. Example: export-fields=id,username (optional) (default to [])
    exportLabels := []string{"Inner_example"} // []string | Export labels parameter, used to customize fieldnames/columns in the exported file. Default is empty array, which means that response properties names will be used. Number of the provided labels must match the number of export-fields Example: export-labels=identifier,name with matching: export-fields=id,username (optional) (default to [])
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["date:desc"])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")
    exportParameters := *openapiclient.NewExportParameters() // ExportParameters | Optional. Override query parameters since they can make URI exceed 2,000 character limit. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildingsAPI.V1BuildingsIdHistoryExportPost(context.Background(), id).ExportFields(exportFields).ExportLabels(exportLabels).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).ExportParameters(exportParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildingsAPI.V1BuildingsIdHistoryExportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BuildingsIdHistoryExportPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `BuildingsAPI.V1BuildingsIdHistoryExportPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of buildings | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BuildingsIdHistoryExportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportFields** | **[]string** | Export fields parameter, used to change default order or ignore some of the response properties. Default is empty array, which means that all fields of the response entity will be serialized. Example: export-fields&#x3D;id,username | [default to []]
 **exportLabels** | **[]string** | Export labels parameter, used to customize fieldnames/columns in the exported file. Default is empty array, which means that response properties names will be used. Number of the provided labels must match the number of export-fields Example: export-labels&#x3D;identifier,name with matching: export-fields&#x3D;id,username | [default to []]
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;date:desc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15 | [default to &quot;&quot;]
 **exportParameters** | [**ExportParameters**](ExportParameters.md) | Optional. Override query parameters since they can make URI exceed 2,000 character limit. | 

### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BuildingsIdHistoryGet

> HistorySearchResults V1BuildingsIdHistoryGet(ctx, id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get specified Building History object 



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
    id := "id_example" // string | instance id of building history record
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["date:desc"])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildingsAPI.V1BuildingsIdHistoryGet(context.Background(), id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildingsAPI.V1BuildingsIdHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BuildingsIdHistoryGet`: HistorySearchResults
    fmt.Fprintf(os.Stdout, "Response from `BuildingsAPI.V1BuildingsIdHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of building history record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BuildingsIdHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;date:desc&quot;]]
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


## V1BuildingsIdHistoryPost

> ObjectHistory V1BuildingsIdHistoryPost(ctx, id).ObjectHistoryNote(objectHistoryNote).Execute()

Add specified Building history object notes 



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
    id := "id_example" // string | instance id of building history record
    objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | history notes to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildingsAPI.V1BuildingsIdHistoryPost(context.Background(), id).ObjectHistoryNote(objectHistoryNote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildingsAPI.V1BuildingsIdHistoryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BuildingsIdHistoryPost`: ObjectHistory
    fmt.Fprintf(os.Stdout, "Response from `BuildingsAPI.V1BuildingsIdHistoryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of building history record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BuildingsIdHistoryPostRequest struct via the builder pattern


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


## V1BuildingsIdPut

> Building V1BuildingsIdPut(ctx, id).Building(building).Execute()

Update specified Building object 



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
    id := "id_example" // string | instance id of building record
    building := *openapiclient.NewBuilding("Apple Park") // Building | building object to update. ids defined in this body will be ignored

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildingsAPI.V1BuildingsIdPut(context.Background(), id).Building(building).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildingsAPI.V1BuildingsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BuildingsIdPut`: Building
    fmt.Fprintf(os.Stdout, "Response from `BuildingsAPI.V1BuildingsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of building record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BuildingsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **building** | [**Building**](Building.md) | building object to update. ids defined in this body will be ignored | 

### Return type

[**Building**](Building.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1BuildingsPost

> HrefResponse V1BuildingsPost(ctx).Building(building).Execute()

Create Building record 



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
    building := *openapiclient.NewBuilding("Apple Park") // Building | building object to create. ids defined in this body will be ignored

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildingsAPI.V1BuildingsPost(context.Background()).Building(building).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildingsAPI.V1BuildingsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BuildingsPost`: HrefResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildingsAPI.V1BuildingsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1BuildingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **building** | [**Building**](Building.md) | building object to create. ids defined in this body will be ignored | 

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

