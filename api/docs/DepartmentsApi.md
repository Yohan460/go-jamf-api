# \DepartmentsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DepartmentsDeleteMultiplePost**](DepartmentsAPI.md#V1DepartmentsDeleteMultiplePost) | **Post** /v1/departments/delete-multiple | Deletes all departments by ids passed in body 
[**V1DepartmentsGet**](DepartmentsAPI.md#V1DepartmentsGet) | **Get** /v1/departments | Search for Departments 
[**V1DepartmentsIdDelete**](DepartmentsAPI.md#V1DepartmentsIdDelete) | **Delete** /v1/departments/{id} | Remove specified department record 
[**V1DepartmentsIdGet**](DepartmentsAPI.md#V1DepartmentsIdGet) | **Get** /v1/departments/{id} | Get specified Department object 
[**V1DepartmentsIdHistoryGet**](DepartmentsAPI.md#V1DepartmentsIdHistoryGet) | **Get** /v1/departments/{id}/history | Get specified Department history object 
[**V1DepartmentsIdHistoryPost**](DepartmentsAPI.md#V1DepartmentsIdHistoryPost) | **Post** /v1/departments/{id}/history | Add specified Department history object notes 
[**V1DepartmentsIdPut**](DepartmentsAPI.md#V1DepartmentsIdPut) | **Put** /v1/departments/{id} | Update specified department object 
[**V1DepartmentsPost**](DepartmentsAPI.md#V1DepartmentsPost) | **Post** /v1/departments | Create department record 



## V1DepartmentsDeleteMultiplePost

> V1DepartmentsDeleteMultiplePost(ctx).Ids(ids).Execute()

Deletes all departments by ids passed in body 



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
    ids := *openapiclient.NewIds() // Ids | ids of departments to be deleted. pass in an array of ids

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DepartmentsAPI.V1DepartmentsDeleteMultiplePost(context.Background()).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.V1DepartmentsDeleteMultiplePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DepartmentsDeleteMultiplePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**Ids**](Ids.md) | ids of departments to be deleted. pass in an array of ids | 

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


## V1DepartmentsGet

> DepartmentsSearchResults V1DepartmentsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Search for Departments 



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
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["id:asc"])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter department collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name. Example: name==\"*department*\" (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DepartmentsAPI.V1DepartmentsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.V1DepartmentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DepartmentsGet`: DepartmentsSearchResults
    fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.V1DepartmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DepartmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;id:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter department collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name. Example: name&#x3D;&#x3D;\&quot;*department*\&quot; | [default to &quot;&quot;]

### Return type

[**DepartmentsSearchResults**](DepartmentsSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DepartmentsIdDelete

> V1DepartmentsIdDelete(ctx, id).Execute()

Remove specified department record 



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
    id := "id_example" // string | instance id of department record

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DepartmentsAPI.V1DepartmentsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.V1DepartmentsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of department record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DepartmentsIdDeleteRequest struct via the builder pattern


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


## V1DepartmentsIdGet

> Department V1DepartmentsIdGet(ctx, id).Execute()

Get specified Department object 



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
    id := "id_example" // string | instance id of department record

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DepartmentsAPI.V1DepartmentsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.V1DepartmentsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DepartmentsIdGet`: Department
    fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.V1DepartmentsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of department record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DepartmentsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Department**](Department.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DepartmentsIdHistoryGet

> HistorySearchResults V1DepartmentsIdHistoryGet(ctx, id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get specified Department history object 



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
    id := "id_example" // string | instance id of department history record
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["date:desc"])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DepartmentsAPI.V1DepartmentsIdHistoryGet(context.Background(), id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.V1DepartmentsIdHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DepartmentsIdHistoryGet`: HistorySearchResults
    fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.V1DepartmentsIdHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of department history record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DepartmentsIdHistoryGetRequest struct via the builder pattern


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


## V1DepartmentsIdHistoryPost

> HrefResponse V1DepartmentsIdHistoryPost(ctx, id).ObjectHistoryNote(objectHistoryNote).Execute()

Add specified Department history object notes 



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
    id := "id_example" // string | instance id of department history record
    objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | history notes to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DepartmentsAPI.V1DepartmentsIdHistoryPost(context.Background(), id).ObjectHistoryNote(objectHistoryNote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.V1DepartmentsIdHistoryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DepartmentsIdHistoryPost`: HrefResponse
    fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.V1DepartmentsIdHistoryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of department history record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DepartmentsIdHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectHistoryNote** | [**ObjectHistoryNote**](ObjectHistoryNote.md) | history notes to create | 

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


## V1DepartmentsIdPut

> Department V1DepartmentsIdPut(ctx, id).Department(department).Execute()

Update specified department object 



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
    id := "id_example" // string | instance id of department record
    department := *openapiclient.NewDepartment("Department of Redundancy Department") // Department | department object to create. ids defined in this body will be ignored

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DepartmentsAPI.V1DepartmentsIdPut(context.Background(), id).Department(department).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.V1DepartmentsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DepartmentsIdPut`: Department
    fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.V1DepartmentsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of department record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DepartmentsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **department** | [**Department**](Department.md) | department object to create. ids defined in this body will be ignored | 

### Return type

[**Department**](Department.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DepartmentsPost

> HrefResponse V1DepartmentsPost(ctx).Department(department).Execute()

Create department record 



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
    department := *openapiclient.NewDepartment("Department of Redundancy Department") // Department | department object to create. ids defined in this body will be ignored

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DepartmentsAPI.V1DepartmentsPost(context.Background()).Department(department).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.V1DepartmentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DepartmentsPost`: HrefResponse
    fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.V1DepartmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DepartmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **department** | [**Department**](Department.md) | department object to create. ids defined in this body will be ignored | 

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

