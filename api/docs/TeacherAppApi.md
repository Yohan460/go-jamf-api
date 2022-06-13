# \TeacherAppApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TeacherAppGet**](TeacherAppApi.md#V1TeacherAppGet) | **Get** /v1/teacher-app | Get the Jamf Teacher settings that you have access to see 
[**V1TeacherAppHistoryGet**](TeacherAppApi.md#V1TeacherAppHistoryGet) | **Get** /v1/teacher-app/history | Get Jamf Teacher app settings history 
[**V1TeacherAppHistoryPost**](TeacherAppApi.md#V1TeacherAppHistoryPost) | **Post** /v1/teacher-app/history | Add Jamf Teacher app settings history notes 
[**V1TeacherAppPut**](TeacherAppApi.md#V1TeacherAppPut) | **Put** /v1/teacher-app | Update a Jamf Teacher settings object 



## V1TeacherAppGet

> TeacherSettingsResponse V1TeacherAppGet(ctx).Execute()

Get the Jamf Teacher settings that you have access to see 



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
    resp, r, err := apiClient.TeacherAppApi.V1TeacherAppGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeacherAppApi.V1TeacherAppGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TeacherAppGet`: TeacherSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `TeacherAppApi.V1TeacherAppGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1TeacherAppGetRequest struct via the builder pattern


### Return type

[**TeacherSettingsResponse**](TeacherSettingsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TeacherAppHistoryGet

> HistorySearchResults V1TeacherAppHistoryGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get Jamf Teacher app settings history 



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
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is not duplicated for each sort criterion, e.g., ...&sort=name:asc,date:desc. Fields that can be sorted: status, updated (optional) (default to [])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeacherAppApi.V1TeacherAppHistoryGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeacherAppApi.V1TeacherAppHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TeacherAppHistoryGet`: HistorySearchResults
    fmt.Fprintf(os.Stdout, "Response from `TeacherAppApi.V1TeacherAppHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TeacherAppHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the &#39;sort&#39; query param is not duplicated for each sort criterion, e.g., ...&amp;sort&#x3D;name:asc,date:desc. Fields that can be sorted: status, updated | [default to []]
 **filter** | **string** | Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15 | [default to &quot;&quot;]

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


## V1TeacherAppHistoryPost

> HrefResponse V1TeacherAppHistoryPost(ctx).ObjectHistoryNote(objectHistoryNote).Execute()

Add Jamf Teacher app settings history notes 



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
    resp, r, err := apiClient.TeacherAppApi.V1TeacherAppHistoryPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeacherAppApi.V1TeacherAppHistoryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TeacherAppHistoryPost`: HrefResponse
    fmt.Fprintf(os.Stdout, "Response from `TeacherAppApi.V1TeacherAppHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TeacherAppHistoryPostRequest struct via the builder pattern


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


## V1TeacherAppPut

> TeacherSettingsResponse V1TeacherAppPut(ctx).TeacherSettingsRequest(teacherSettingsRequest).Execute()

Update a Jamf Teacher settings object 



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
    teacherSettingsRequest := *openapiclient.NewTeacherSettingsRequest() // TeacherSettingsRequest | Teacher settings to create.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeacherAppApi.V1TeacherAppPut(context.Background()).TeacherSettingsRequest(teacherSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeacherAppApi.V1TeacherAppPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TeacherAppPut`: TeacherSettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `TeacherAppApi.V1TeacherAppPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TeacherAppPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teacherSettingsRequest** | [**TeacherSettingsRequest**](TeacherSettingsRequest.md) | Teacher settings to create. | 

### Return type

[**TeacherSettingsResponse**](TeacherSettingsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

