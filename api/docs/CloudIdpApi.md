# \CloudIdpAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CloudIdpExportPost**](CloudIdpAPI.md#V1CloudIdpExportPost) | **Post** /v1/cloud-idp/export | Export Cloud Identity Providers collection 
[**V1CloudIdpGet**](CloudIdpAPI.md#V1CloudIdpGet) | **Get** /v1/cloud-idp | Get information about all Cloud Identity Providers configurations.
[**V1CloudIdpIdGet**](CloudIdpAPI.md#V1CloudIdpIdGet) | **Get** /v1/cloud-idp/{id} | Get Cloud Identity Provider configuration with given ID.
[**V1CloudIdpIdHistoryGet**](CloudIdpAPI.md#V1CloudIdpIdHistoryGet) | **Get** /v1/cloud-idp/{id}/history | Get Cloud Identity Provider history
[**V1CloudIdpIdHistoryPost**](CloudIdpAPI.md#V1CloudIdpIdHistoryPost) | **Post** /v1/cloud-idp/{id}/history | Add Cloud Identity Provider history note
[**V1CloudIdpIdTestGroupPost**](CloudIdpAPI.md#V1CloudIdpIdTestGroupPost) | **Post** /v1/cloud-idp/{id}/test-group | Get group test search
[**V1CloudIdpIdTestUserMembershipPost**](CloudIdpAPI.md#V1CloudIdpIdTestUserMembershipPost) | **Post** /v1/cloud-idp/{id}/test-user-membership | Get membership test search
[**V1CloudIdpIdTestUserPost**](CloudIdpAPI.md#V1CloudIdpIdTestUserPost) | **Post** /v1/cloud-idp/{id}/test-user | Get user test search



## V1CloudIdpExportPost

> interface{} V1CloudIdpExportPost(ctx).ExportFields(exportFields).ExportLabels(exportLabels).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).ExportParameters(exportParameters).Execute()

Export Cloud Identity Providers collection 



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
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:desc. Multiple sort criteria are supported and must be seperated with a comma. Example: sort=id:desc,name:asc (optional) (default to ["id:asc"])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name. This param can be combined with paging and sorting. Example: name==\"*department*\" (optional) (default to "")
    exportParameters := *openapiclient.NewExportParameters() // ExportParameters | Optional. Override query parameters since they can make URI exceed 2,000 character limit. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudIdpAPI.V1CloudIdpExportPost(context.Background()).ExportFields(exportFields).ExportLabels(exportLabels).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).ExportParameters(exportParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudIdpAPI.V1CloudIdpExportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudIdpExportPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudIdpAPI.V1CloudIdpExportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudIdpExportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportFields** | **[]string** | Export fields parameter, used to change default order or ignore some of the response properties. Default is empty array, which means that all fields of the response entity will be serialized. Example: export-fields&#x3D;id,username | [default to []]
 **exportLabels** | **[]string** | Export labels parameter, used to customize fieldnames/columns in the exported file. Default is empty array, which means that response properties names will be used. Number of the provided labels must match the number of export-fields Example: export-labels&#x3D;identifier,name with matching: export-fields&#x3D;id,username | [default to []]
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:desc. Multiple sort criteria are supported and must be seperated with a comma. Example: sort&#x3D;id:desc,name:asc | [default to [&quot;id:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name. This param can be combined with paging and sorting. Example: name&#x3D;&#x3D;\&quot;*department*\&quot; | [default to &quot;&quot;]
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


## V1CloudIdpGet

> ConfigurationSearchResults V1CloudIdpGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Execute()

Get information about all Cloud Identity Providers configurations.



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
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["id:desc"])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudIdpAPI.V1CloudIdpGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudIdpAPI.V1CloudIdpGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudIdpGet`: ConfigurationSearchResults
    fmt.Fprintf(os.Stdout, "Response from `CloudIdpAPI.V1CloudIdpGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudIdpGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;id:desc&quot;]]

### Return type

[**ConfigurationSearchResults**](ConfigurationSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudIdpIdGet

> CloudIdPCommon V1CloudIdpIdGet(ctx, id).Execute()

Get Cloud Identity Provider configuration with given ID.



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
    id := "id_example" // string | Cloud Identity Provider identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudIdpAPI.V1CloudIdpIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudIdpAPI.V1CloudIdpIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudIdpIdGet`: CloudIdPCommon
    fmt.Fprintf(os.Stdout, "Response from `CloudIdpAPI.V1CloudIdpIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudIdpIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudIdPCommon**](CloudIdPCommon.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudIdpIdHistoryGet

> HistorySearchResults V1CloudIdpIdHistoryGet(ctx, id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get Cloud Identity Provider history



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
    id := "id_example" // string | Cloud Identity Provider identifier
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["date:desc"])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudIdpAPI.V1CloudIdpIdHistoryGet(context.Background(), id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudIdpAPI.V1CloudIdpIdHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudIdpIdHistoryGet`: HistorySearchResults
    fmt.Fprintf(os.Stdout, "Response from `CloudIdpAPI.V1CloudIdpIdHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudIdpIdHistoryGetRequest struct via the builder pattern


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


## V1CloudIdpIdHistoryPost

> ObjectHistory V1CloudIdpIdHistoryPost(ctx, id).ObjectHistoryNote(objectHistoryNote).Execute()

Add Cloud Identity Provider history note



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
    id := "id_example" // string | Cloud Identity Provider identifier
    objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | history notes to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudIdpAPI.V1CloudIdpIdHistoryPost(context.Background(), id).ObjectHistoryNote(objectHistoryNote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudIdpAPI.V1CloudIdpIdHistoryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudIdpIdHistoryPost`: ObjectHistory
    fmt.Fprintf(os.Stdout, "Response from `CloudIdpAPI.V1CloudIdpIdHistoryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudIdpIdHistoryPostRequest struct via the builder pattern


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


## V1CloudIdpIdTestGroupPost

> GroupTestSearchResponse V1CloudIdpIdTestGroupPost(ctx, id).GroupTestSearchRequest(groupTestSearchRequest).Execute()

Get group test search



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
    id := "id_example" // string | Cloud Identity Provider identifier
    groupTestSearchRequest := *openapiclient.NewGroupTestSearchRequest("users") // GroupTestSearchRequest | Search request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudIdpAPI.V1CloudIdpIdTestGroupPost(context.Background(), id).GroupTestSearchRequest(groupTestSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudIdpAPI.V1CloudIdpIdTestGroupPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudIdpIdTestGroupPost`: GroupTestSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudIdpAPI.V1CloudIdpIdTestGroupPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudIdpIdTestGroupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupTestSearchRequest** | [**GroupTestSearchRequest**](GroupTestSearchRequest.md) | Search request | 

### Return type

[**GroupTestSearchResponse**](GroupTestSearchResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudIdpIdTestUserMembershipPost

> MembershipTestSearchResponse V1CloudIdpIdTestUserMembershipPost(ctx, id).MembershipTestSearchRequest(membershipTestSearchRequest).Execute()

Get membership test search



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
    id := "id_example" // string | Cloud Identity Provider identifier
    membershipTestSearchRequest := *openapiclient.NewMembershipTestSearchRequest("admin", "users") // MembershipTestSearchRequest | Search request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudIdpAPI.V1CloudIdpIdTestUserMembershipPost(context.Background(), id).MembershipTestSearchRequest(membershipTestSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudIdpAPI.V1CloudIdpIdTestUserMembershipPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudIdpIdTestUserMembershipPost`: MembershipTestSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudIdpAPI.V1CloudIdpIdTestUserMembershipPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudIdpIdTestUserMembershipPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **membershipTestSearchRequest** | [**MembershipTestSearchRequest**](MembershipTestSearchRequest.md) | Search request | 

### Return type

[**MembershipTestSearchResponse**](MembershipTestSearchResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CloudIdpIdTestUserPost

> UserTestSearchResponse V1CloudIdpIdTestUserPost(ctx, id).UserTestSearchRequest(userTestSearchRequest).Execute()

Get user test search



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
    id := "id_example" // string | Cloud Identity Provider identifier
    userTestSearchRequest := *openapiclient.NewUserTestSearchRequest("admin") // UserTestSearchRequest | Search request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudIdpAPI.V1CloudIdpIdTestUserPost(context.Background(), id).UserTestSearchRequest(userTestSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudIdpAPI.V1CloudIdpIdTestUserPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1CloudIdpIdTestUserPost`: UserTestSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudIdpAPI.V1CloudIdpIdTestUserPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cloud Identity Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudIdpIdTestUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userTestSearchRequest** | [**UserTestSearchRequest**](UserTestSearchRequest.md) | Search request | 

### Return type

[**UserTestSearchResponse**](UserTestSearchResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

