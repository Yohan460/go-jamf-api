# \EnrollmentApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreviewEnrollmentAccessGroupsGet**](EnrollmentApi.md#PreviewEnrollmentAccessGroupsGet) | **Get** /preview/enrollment/access-groups | Retrieve the configured LDAP groups configured for User-Initiated Enrollment. 
[**PreviewEnrollmentAccessGroupsIdDelete**](EnrollmentApi.md#PreviewEnrollmentAccessGroupsIdDelete) | **Delete** /preview/enrollment/access-groups/{id} | Delete an LDAP group&#39;s access to user initiated Enrollment. 
[**PreviewEnrollmentAccessGroupsIdGet**](EnrollmentApi.md#PreviewEnrollmentAccessGroupsIdGet) | **Get** /preview/enrollment/access-groups/{id} | Retrieve the configured LDAP groups configured for User-Initiated Enrollment 
[**PreviewEnrollmentAccessGroupsIdPut**](EnrollmentApi.md#PreviewEnrollmentAccessGroupsIdPut) | **Put** /preview/enrollment/access-groups/{id} | Modify the configured LDAP groups configured for User-Initiated Enrollment. Only exiting Access Groups can be updated. 
[**PreviewEnrollmentAccessGroupsPost**](EnrollmentApi.md#PreviewEnrollmentAccessGroupsPost) | **Post** /preview/enrollment/access-groups | Add the configured LDAP group for User-Initiated Enrollment. 
[**V1EnrollmentAccessGroupsGet**](EnrollmentApi.md#V1EnrollmentAccessGroupsGet) | **Get** /v1/enrollment/access-groups | Retrieve the configured LDAP groups configured for User-Initiated Enrollment 
[**V1EnrollmentAccessGroupsGroupKeyDelete**](EnrollmentApi.md#V1EnrollmentAccessGroupsGroupKeyDelete) | **Delete** /v1/enrollment/access-groups/{group-key} | Delete an LDAP group&#39;s access to user initiated Enrollment 
[**V1EnrollmentAccessGroupsGroupKeyGet**](EnrollmentApi.md#V1EnrollmentAccessGroupsGroupKeyGet) | **Get** /v1/enrollment/access-groups/{group-key} | Retrieve the configured LDAP groups configured for User-Initiated Enrollment 
[**V1EnrollmentAccessGroupsGroupKeyPut**](EnrollmentApi.md#V1EnrollmentAccessGroupsGroupKeyPut) | **Put** /v1/enrollment/access-groups/{group-key} | Modify the configured LDAP groups configured for User-Initiated Enrollment 
[**V1EnrollmentFilteredLanguageCodesGet**](EnrollmentApi.md#V1EnrollmentFilteredLanguageCodesGet) | **Get** /v1/enrollment/filtered-language-codes | Retrieve the list of languages and corresponding ISO 639-1 Codes but only those not already added to Enrollment 
[**V1EnrollmentGet**](EnrollmentApi.md#V1EnrollmentGet) | **Get** /v1/enrollment | Get Enrollment object and Re-enrollment settings 
[**V1EnrollmentHistoryGet**](EnrollmentApi.md#V1EnrollmentHistoryGet) | **Get** /v1/enrollment/history | Get sorted and paged Enrollment history object 
[**V1EnrollmentHistoryPost**](EnrollmentApi.md#V1EnrollmentHistoryPost) | **Post** /v1/enrollment/history | Add Enrollment history object notes 
[**V1EnrollmentLanguageCodesGet**](EnrollmentApi.md#V1EnrollmentLanguageCodesGet) | **Get** /v1/enrollment/language-codes | Retrieve the list of languages and corresponding ISO 639-1 Codes 
[**V1EnrollmentLanguagesGet**](EnrollmentApi.md#V1EnrollmentLanguagesGet) | **Get** /v1/enrollment/languages | Get an array of the language codes that have Enrollment messaging 
[**V1EnrollmentLanguagesLanguageDelete**](EnrollmentApi.md#V1EnrollmentLanguagesLanguageDelete) | **Delete** /v1/enrollment/languages/{language} | Delete the Enrollment messaging for a language 
[**V1EnrollmentLanguagesLanguageGet**](EnrollmentApi.md#V1EnrollmentLanguagesLanguageGet) | **Get** /v1/enrollment/languages/{language} | Retrieve the Enrollment messaging for a language 
[**V1EnrollmentLanguagesLanguagePut**](EnrollmentApi.md#V1EnrollmentLanguagesLanguagePut) | **Put** /v1/enrollment/languages/{language} | Edit Enrollment messaging for a language 
[**V1EnrollmentPut**](EnrollmentApi.md#V1EnrollmentPut) | **Put** /v1/enrollment | Update Enrollment object 
[**V2EnrollmentAccessGroupsGet**](EnrollmentApi.md#V2EnrollmentAccessGroupsGet) | **Get** /v2/enrollment/access-groups | Retrieve the configured LDAP groups configured for User-Initiated Enrollment 
[**V2EnrollmentAccessGroupsPost**](EnrollmentApi.md#V2EnrollmentAccessGroupsPost) | **Post** /v2/enrollment/access-groups | Add the configured LDAP group for User-Initiated Enrollment. 
[**V2EnrollmentAccessGroupsServerIdGroupIdDelete**](EnrollmentApi.md#V2EnrollmentAccessGroupsServerIdGroupIdDelete) | **Delete** /v2/enrollment/access-groups/{serverId}/{groupId} | Delete an LDAP group&#39;s access to user initiated Enrollment 
[**V2EnrollmentAccessGroupsServerIdGroupIdGet**](EnrollmentApi.md#V2EnrollmentAccessGroupsServerIdGroupIdGet) | **Get** /v2/enrollment/access-groups/{serverId}/{groupId} | Retrieve the configured LDAP groups configured for User-Initiated Enrollment 
[**V2EnrollmentAccessGroupsServerIdGroupIdPut**](EnrollmentApi.md#V2EnrollmentAccessGroupsServerIdGroupIdPut) | **Put** /v2/enrollment/access-groups/{serverId}/{groupId} | Modify the configured LDAP groups configured for User-Initiated Enrollment 
[**V2EnrollmentFilteredLanguageCodesGet**](EnrollmentApi.md#V2EnrollmentFilteredLanguageCodesGet) | **Get** /v2/enrollment/filtered-language-codes | Retrieve the list of languages and corresponding ISO 639-1 Codes but only those not already added to Enrollment 
[**V2EnrollmentGet**](EnrollmentApi.md#V2EnrollmentGet) | **Get** /v2/enrollment | Get Enrollment object and Re-enrollment settings 
[**V2EnrollmentHistoryExportPost**](EnrollmentApi.md#V2EnrollmentHistoryExportPost) | **Post** /v2/enrollment/history/export | Export enrollment history collection 
[**V2EnrollmentHistoryGet**](EnrollmentApi.md#V2EnrollmentHistoryGet) | **Get** /v2/enrollment/history | Get sorted and paged Enrollment history object 
[**V2EnrollmentHistoryPost**](EnrollmentApi.md#V2EnrollmentHistoryPost) | **Post** /v2/enrollment/history | Add Enrollment history object notes 
[**V2EnrollmentLanguageCodesGet**](EnrollmentApi.md#V2EnrollmentLanguageCodesGet) | **Get** /v2/enrollment/language-codes | Retrieve the list of languages and corresponding ISO 639-1 Codes 
[**V2EnrollmentLanguagesGet**](EnrollmentApi.md#V2EnrollmentLanguagesGet) | **Get** /v2/enrollment/languages | Get an array of the language codes that have Enrollment messaging 
[**V2EnrollmentLanguagesLanguageIdDelete**](EnrollmentApi.md#V2EnrollmentLanguagesLanguageIdDelete) | **Delete** /v2/enrollment/languages/{languageId} | Delete the Enrollment messaging for a language 
[**V2EnrollmentLanguagesLanguageIdGet**](EnrollmentApi.md#V2EnrollmentLanguagesLanguageIdGet) | **Get** /v2/enrollment/languages/{languageId} | Retrieve the Enrollment messaging for a language 
[**V2EnrollmentLanguagesLanguageIdPut**](EnrollmentApi.md#V2EnrollmentLanguagesLanguageIdPut) | **Put** /v2/enrollment/languages/{languageId} | Edit Enrollment messaging for a language 
[**V2EnrollmentPut**](EnrollmentApi.md#V2EnrollmentPut) | **Put** /v2/enrollment | Update Enrollment object 



## PreviewEnrollmentAccessGroupsGet

> AccessGroupsPreviewSearchResults PreviewEnrollmentAccessGroupsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).AllUsersOptionFirst(allUsersOptionFirst).Execute()

Retrieve the configured LDAP groups configured for User-Initiated Enrollment. 



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
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: `property:asc/desc`. Default sort is `name:asc`. Multiple sort criteria are supported and must be separated with a comma. Example: `sort=date:desc,name:asc`.  (optional) (default to ["name:asc"])
    allUsersOptionFirst := true // bool | Return \"All LDAP Users\" option on the first position if it is present in the current page (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.PreviewEnrollmentAccessGroupsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).AllUsersOptionFirst(allUsersOptionFirst).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.PreviewEnrollmentAccessGroupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewEnrollmentAccessGroupsGet`: AccessGroupsPreviewSearchResults
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.PreviewEnrollmentAccessGroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewEnrollmentAccessGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: &#x60;property:asc/desc&#x60;. Default sort is &#x60;name:asc&#x60;. Multiple sort criteria are supported and must be separated with a comma. Example: &#x60;sort&#x3D;date:desc,name:asc&#x60;.  | [default to [&quot;name:asc&quot;]]
 **allUsersOptionFirst** | **bool** | Return \&quot;All LDAP Users\&quot; option on the first position if it is present in the current page | [default to false]

### Return type

[**AccessGroupsPreviewSearchResults**](AccessGroupsPreviewSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewEnrollmentAccessGroupsIdDelete

> PreviewEnrollmentAccessGroupsIdDelete(ctx, id).Execute()

Delete an LDAP group's access to user initiated Enrollment. 



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
    id := "id_example" // string | Autogenerated Access Group ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.PreviewEnrollmentAccessGroupsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.PreviewEnrollmentAccessGroupsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autogenerated Access Group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewEnrollmentAccessGroupsIdDeleteRequest struct via the builder pattern


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


## PreviewEnrollmentAccessGroupsIdGet

> EnrollmentAccessGroupPreview PreviewEnrollmentAccessGroupsIdGet(ctx, id).Execute()

Retrieve the configured LDAP groups configured for User-Initiated Enrollment 



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
    id := "id_example" // string | Autogenerated Access Group ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.PreviewEnrollmentAccessGroupsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.PreviewEnrollmentAccessGroupsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewEnrollmentAccessGroupsIdGet`: EnrollmentAccessGroupPreview
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.PreviewEnrollmentAccessGroupsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autogenerated Access Group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewEnrollmentAccessGroupsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnrollmentAccessGroupPreview**](EnrollmentAccessGroupPreview.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewEnrollmentAccessGroupsIdPut

> EnrollmentAccessGroupPreview PreviewEnrollmentAccessGroupsIdPut(ctx, id).EnrollmentAccessGroupPreview(enrollmentAccessGroupPreview).Execute()

Modify the configured LDAP groups configured for User-Initiated Enrollment. Only exiting Access Groups can be updated. 



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
    id := "id_example" // string | Autogenerated Access Group ID.
    enrollmentAccessGroupPreview := *openapiclient.NewEnrollmentAccessGroupPreview("1", "Grade School Pupils") // EnrollmentAccessGroupPreview | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.PreviewEnrollmentAccessGroupsIdPut(context.Background(), id).EnrollmentAccessGroupPreview(enrollmentAccessGroupPreview).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.PreviewEnrollmentAccessGroupsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewEnrollmentAccessGroupsIdPut`: EnrollmentAccessGroupPreview
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.PreviewEnrollmentAccessGroupsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autogenerated Access Group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewEnrollmentAccessGroupsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enrollmentAccessGroupPreview** | [**EnrollmentAccessGroupPreview**](EnrollmentAccessGroupPreview.md) |  | 

### Return type

[**EnrollmentAccessGroupPreview**](EnrollmentAccessGroupPreview.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewEnrollmentAccessGroupsPost

> HrefResponse PreviewEnrollmentAccessGroupsPost(ctx).EnrollmentAccessGroupPreview(enrollmentAccessGroupPreview).Execute()

Add the configured LDAP group for User-Initiated Enrollment. 



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
    enrollmentAccessGroupPreview := *openapiclient.NewEnrollmentAccessGroupPreview("1", "Grade School Pupils") // EnrollmentAccessGroupPreview | Configured LDAP group to create.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.PreviewEnrollmentAccessGroupsPost(context.Background()).EnrollmentAccessGroupPreview(enrollmentAccessGroupPreview).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.PreviewEnrollmentAccessGroupsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewEnrollmentAccessGroupsPost`: HrefResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.PreviewEnrollmentAccessGroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewEnrollmentAccessGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enrollmentAccessGroupPreview** | [**EnrollmentAccessGroupPreview**](EnrollmentAccessGroupPreview.md) | Configured LDAP group to create. | 

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


## V1EnrollmentAccessGroupsGet

> AccessGroupsSearchResults V1EnrollmentAccessGroupsGet(ctx).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()

Retrieve the configured LDAP groups configured for User-Initiated Enrollment 



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
    sort := "sort_example" // string | Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to "name:asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.V1EnrollmentAccessGroupsGet(context.Background()).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V1EnrollmentAccessGroupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentAccessGroupsGet`: AccessGroupsSearchResults
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V1EnrollmentAccessGroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentAccessGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]
 **pagesize** | **int32** |  | [default to 100]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **string** | Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to &quot;name:asc&quot;]

### Return type

[**AccessGroupsSearchResults**](AccessGroupsSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentAccessGroupsGroupKeyDelete

> V1EnrollmentAccessGroupsGroupKeyDelete(ctx, groupKey).Execute()

Delete an LDAP group's access to user initiated Enrollment 



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
    groupKey := "groupKey_example" // string | The group key is a string composed of the LDAP server ID, underscore and the LDAP group id. Example: ``1_2``

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.V1EnrollmentAccessGroupsGroupKeyDelete(context.Background(), groupKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V1EnrollmentAccessGroupsGroupKeyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupKey** | **string** | The group key is a string composed of the LDAP server ID, underscore and the LDAP group id. Example: &#x60;&#x60;1_2&#x60;&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentAccessGroupsGroupKeyDeleteRequest struct via the builder pattern


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


## V1EnrollmentAccessGroupsGroupKeyGet

> EnrollmentAccessGroup V1EnrollmentAccessGroupsGroupKeyGet(ctx, groupKey).Execute()

Retrieve the configured LDAP groups configured for User-Initiated Enrollment 



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
    groupKey := "groupKey_example" // string | The group key is a string composed of the LDAP server ID, underscore and the LDAP group id. Example: ``1_2``

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.V1EnrollmentAccessGroupsGroupKeyGet(context.Background(), groupKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V1EnrollmentAccessGroupsGroupKeyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentAccessGroupsGroupKeyGet`: EnrollmentAccessGroup
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V1EnrollmentAccessGroupsGroupKeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupKey** | **string** | The group key is a string composed of the LDAP server ID, underscore and the LDAP group id. Example: &#x60;&#x60;1_2&#x60;&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentAccessGroupsGroupKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnrollmentAccessGroup**](EnrollmentAccessGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentAccessGroupsGroupKeyPut

> EnrollmentAccessGroup V1EnrollmentAccessGroupsGroupKeyPut(ctx, groupKey).EnrollmentAccessGroup(enrollmentAccessGroup).Execute()

Modify the configured LDAP groups configured for User-Initiated Enrollment 



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
    groupKey := "groupKey_example" // string | The group key is a string composed of the LDAP server ID, underscore and the LDAP group id. Example: ``1_2``
    enrollmentAccessGroup := *openapiclient.NewEnrollmentAccessGroup() // EnrollmentAccessGroup |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.V1EnrollmentAccessGroupsGroupKeyPut(context.Background(), groupKey).EnrollmentAccessGroup(enrollmentAccessGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V1EnrollmentAccessGroupsGroupKeyPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentAccessGroupsGroupKeyPut`: EnrollmentAccessGroup
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V1EnrollmentAccessGroupsGroupKeyPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupKey** | **string** | The group key is a string composed of the LDAP server ID, underscore and the LDAP group id. Example: &#x60;&#x60;1_2&#x60;&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentAccessGroupsGroupKeyPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enrollmentAccessGroup** | [**EnrollmentAccessGroup**](EnrollmentAccessGroup.md) |  | 

### Return type

[**EnrollmentAccessGroup**](EnrollmentAccessGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentFilteredLanguageCodesGet

> []LanguageCode V1EnrollmentFilteredLanguageCodesGet(ctx).Execute()

Retrieve the list of languages and corresponding ISO 639-1 Codes but only those not already added to Enrollment 



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
    resp, r, err := apiClient.EnrollmentApi.V1EnrollmentFilteredLanguageCodesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V1EnrollmentFilteredLanguageCodesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentFilteredLanguageCodesGet`: []LanguageCode
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V1EnrollmentFilteredLanguageCodesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentFilteredLanguageCodesGetRequest struct via the builder pattern


### Return type

[**[]LanguageCode**](LanguageCode.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentGet

> EnrollmentSettings V1EnrollmentGet(ctx).Execute()

Get Enrollment object and Re-enrollment settings 



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
    resp, r, err := apiClient.EnrollmentApi.V1EnrollmentGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V1EnrollmentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentGet`: EnrollmentSettings
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V1EnrollmentGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentGetRequest struct via the builder pattern


### Return type

[**EnrollmentSettings**](EnrollmentSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentHistoryGet

> HistorySearchResults V1EnrollmentHistoryGet(ctx).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()

Get sorted and paged Enrollment history object 



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
    resp, r, err := apiClient.EnrollmentApi.V1EnrollmentHistoryGet(context.Background()).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V1EnrollmentHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentHistoryGet`: HistorySearchResults
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V1EnrollmentHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentHistoryGetRequest struct via the builder pattern


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


## V1EnrollmentHistoryPost

> ObjectHistory V1EnrollmentHistoryPost(ctx).ObjectHistoryNote(objectHistoryNote).Execute()

Add Enrollment history object notes 



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
    resp, r, err := apiClient.EnrollmentApi.V1EnrollmentHistoryPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V1EnrollmentHistoryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentHistoryPost`: ObjectHistory
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V1EnrollmentHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentHistoryPostRequest struct via the builder pattern


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


## V1EnrollmentLanguageCodesGet

> []LanguageCode V1EnrollmentLanguageCodesGet(ctx).Execute()

Retrieve the list of languages and corresponding ISO 639-1 Codes 



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
    resp, r, err := apiClient.EnrollmentApi.V1EnrollmentLanguageCodesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V1EnrollmentLanguageCodesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentLanguageCodesGet`: []LanguageCode
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V1EnrollmentLanguageCodesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentLanguageCodesGetRequest struct via the builder pattern


### Return type

[**[]LanguageCode**](LanguageCode.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentLanguagesGet

> ProcessTextsSearchResults V1EnrollmentLanguagesGet(ctx).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()

Get an array of the language codes that have Enrollment messaging 



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
    sort := "sort_example" // string | Sorting criteria in the format: property:asc/desc. Default sort is languageCode:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to "languageCode:asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.V1EnrollmentLanguagesGet(context.Background()).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V1EnrollmentLanguagesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentLanguagesGet`: ProcessTextsSearchResults
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V1EnrollmentLanguagesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentLanguagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]
 **pagesize** | **int32** |  | [default to 100]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **string** | Sorting criteria in the format: property:asc/desc. Default sort is languageCode:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to &quot;languageCode:asc&quot;]

### Return type

[**ProcessTextsSearchResults**](ProcessTextsSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentLanguagesLanguageDelete

> V1EnrollmentLanguagesLanguageDelete(ctx, language).Execute()

Delete the Enrollment messaging for a language 



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
    language := "language_example" // string | Two letter ISO 639-1 Language Code

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.V1EnrollmentLanguagesLanguageDelete(context.Background(), language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V1EnrollmentLanguagesLanguageDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | Two letter ISO 639-1 Language Code | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentLanguagesLanguageDeleteRequest struct via the builder pattern


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


## V1EnrollmentLanguagesLanguageGet

> EnrollmentProcessTextObject V1EnrollmentLanguagesLanguageGet(ctx, language).Execute()

Retrieve the Enrollment messaging for a language 



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
    language := "language_example" // string | Two letter ISO 639-1 Language Code

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.V1EnrollmentLanguagesLanguageGet(context.Background(), language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V1EnrollmentLanguagesLanguageGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentLanguagesLanguageGet`: EnrollmentProcessTextObject
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V1EnrollmentLanguagesLanguageGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | Two letter ISO 639-1 Language Code | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentLanguagesLanguageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnrollmentProcessTextObject**](EnrollmentProcessTextObject.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentLanguagesLanguagePut

> EnrollmentProcessTextObject V1EnrollmentLanguagesLanguagePut(ctx, language).EnrollmentProcessTextObject(enrollmentProcessTextObject).Execute()

Edit Enrollment messaging for a language 



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
    language := "language_example" // string | Two letter ISO 639-1 Language Code
    enrollmentProcessTextObject := *openapiclient.NewEnrollmentProcessTextObject() // EnrollmentProcessTextObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.V1EnrollmentLanguagesLanguagePut(context.Background(), language).EnrollmentProcessTextObject(enrollmentProcessTextObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V1EnrollmentLanguagesLanguagePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentLanguagesLanguagePut`: EnrollmentProcessTextObject
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V1EnrollmentLanguagesLanguagePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | Two letter ISO 639-1 Language Code | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentLanguagesLanguagePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enrollmentProcessTextObject** | [**EnrollmentProcessTextObject**](EnrollmentProcessTextObject.md) |  | 

### Return type

[**EnrollmentProcessTextObject**](EnrollmentProcessTextObject.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentPut

> EnrollmentSettings V1EnrollmentPut(ctx).EnrollmentSettings(enrollmentSettings).Execute()

Update Enrollment object 



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
    enrollmentSettings := *openapiclient.NewEnrollmentSettings("FlushMdmCommandsOnReenroll_example", false, "radmin", "PasswordType_example", "PersonalDeviceEnrollmentType_example") // EnrollmentSettings | Update enrollment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.V1EnrollmentPut(context.Background()).EnrollmentSettings(enrollmentSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V1EnrollmentPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentPut`: EnrollmentSettings
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V1EnrollmentPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enrollmentSettings** | [**EnrollmentSettings**](EnrollmentSettings.md) | Update enrollment | 

### Return type

[**EnrollmentSettings**](EnrollmentSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentAccessGroupsGet

> AccessGroupsV2SearchResults V2EnrollmentAccessGroupsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).AllUsersOptionFirst(allUsersOptionFirst).Execute()

Retrieve the configured LDAP groups configured for User-Initiated Enrollment 



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
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: `property:asc/desc`. Default sort is `name:asc`. Multiple sort criteria are supported and must be separated with a comma. Example: `sort=date:desc,name:asc`.  (optional) (default to ["name:asc"])
    allUsersOptionFirst := true // bool | Return \"All LDAP Users\" option on the first position if it is present in the current page (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.V2EnrollmentAccessGroupsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).AllUsersOptionFirst(allUsersOptionFirst).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V2EnrollmentAccessGroupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2EnrollmentAccessGroupsGet`: AccessGroupsV2SearchResults
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V2EnrollmentAccessGroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentAccessGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: &#x60;property:asc/desc&#x60;. Default sort is &#x60;name:asc&#x60;. Multiple sort criteria are supported and must be separated with a comma. Example: &#x60;sort&#x3D;date:desc,name:asc&#x60;.  | [default to [&quot;name:asc&quot;]]
 **allUsersOptionFirst** | **bool** | Return \&quot;All LDAP Users\&quot; option on the first position if it is present in the current page | [default to false]

### Return type

[**AccessGroupsV2SearchResults**](AccessGroupsV2SearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentAccessGroupsPost

> HrefResponse V2EnrollmentAccessGroupsPost(ctx).EnrollmentAccessGroupV2(enrollmentAccessGroupV2).Execute()

Add the configured LDAP group for User-Initiated Enrollment. 



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
    enrollmentAccessGroupV2 := *openapiclient.NewEnrollmentAccessGroupV2() // EnrollmentAccessGroupV2 | Configured LDAP group to create.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.V2EnrollmentAccessGroupsPost(context.Background()).EnrollmentAccessGroupV2(enrollmentAccessGroupV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V2EnrollmentAccessGroupsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2EnrollmentAccessGroupsPost`: HrefResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V2EnrollmentAccessGroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentAccessGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enrollmentAccessGroupV2** | [**EnrollmentAccessGroupV2**](EnrollmentAccessGroupV2.md) | Configured LDAP group to create. | 

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


## V2EnrollmentAccessGroupsServerIdGroupIdDelete

> V2EnrollmentAccessGroupsServerIdGroupIdDelete(ctx, serverId, groupId).Execute()

Delete an LDAP group's access to user initiated Enrollment 



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
    serverId := "serverId_example" // string | LDAP server id
    groupId := "groupId_example" // string | LDAP group id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.V2EnrollmentAccessGroupsServerIdGroupIdDelete(context.Background(), serverId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V2EnrollmentAccessGroupsServerIdGroupIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | LDAP server id | 
**groupId** | **string** | LDAP group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentAccessGroupsServerIdGroupIdDeleteRequest struct via the builder pattern


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


## V2EnrollmentAccessGroupsServerIdGroupIdGet

> EnrollmentAccessGroupV2 V2EnrollmentAccessGroupsServerIdGroupIdGet(ctx, serverId, groupId).Execute()

Retrieve the configured LDAP groups configured for User-Initiated Enrollment 



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
    serverId := "serverId_example" // string | LDAP server id.
    groupId := "groupId_example" // string | LDAP group id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.V2EnrollmentAccessGroupsServerIdGroupIdGet(context.Background(), serverId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V2EnrollmentAccessGroupsServerIdGroupIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2EnrollmentAccessGroupsServerIdGroupIdGet`: EnrollmentAccessGroupV2
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V2EnrollmentAccessGroupsServerIdGroupIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | LDAP server id. | 
**groupId** | **string** | LDAP group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentAccessGroupsServerIdGroupIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnrollmentAccessGroupV2**](EnrollmentAccessGroupV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentAccessGroupsServerIdGroupIdPut

> EnrollmentAccessGroupV2 V2EnrollmentAccessGroupsServerIdGroupIdPut(ctx, serverId, groupId).EnrollmentAccessGroupV2(enrollmentAccessGroupV2).Execute()

Modify the configured LDAP groups configured for User-Initiated Enrollment 



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
    serverId := "serverId_example" // string | LDAP server id.
    groupId := "groupId_example" // string | LDAP group id.
    enrollmentAccessGroupV2 := *openapiclient.NewEnrollmentAccessGroupV2() // EnrollmentAccessGroupV2 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.V2EnrollmentAccessGroupsServerIdGroupIdPut(context.Background(), serverId, groupId).EnrollmentAccessGroupV2(enrollmentAccessGroupV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V2EnrollmentAccessGroupsServerIdGroupIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2EnrollmentAccessGroupsServerIdGroupIdPut`: EnrollmentAccessGroupV2
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V2EnrollmentAccessGroupsServerIdGroupIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | LDAP server id. | 
**groupId** | **string** | LDAP group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentAccessGroupsServerIdGroupIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enrollmentAccessGroupV2** | [**EnrollmentAccessGroupV2**](EnrollmentAccessGroupV2.md) |  | 

### Return type

[**EnrollmentAccessGroupV2**](EnrollmentAccessGroupV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentFilteredLanguageCodesGet

> []LanguageCode V2EnrollmentFilteredLanguageCodesGet(ctx).Execute()

Retrieve the list of languages and corresponding ISO 639-1 Codes but only those not already added to Enrollment 



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
    resp, r, err := apiClient.EnrollmentApi.V2EnrollmentFilteredLanguageCodesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V2EnrollmentFilteredLanguageCodesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2EnrollmentFilteredLanguageCodesGet`: []LanguageCode
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V2EnrollmentFilteredLanguageCodesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentFilteredLanguageCodesGetRequest struct via the builder pattern


### Return type

[**[]LanguageCode**](LanguageCode.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentGet

> EnrollmentSettingsV2 V2EnrollmentGet(ctx).Execute()

Get Enrollment object and Re-enrollment settings 



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
    resp, r, err := apiClient.EnrollmentApi.V2EnrollmentGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V2EnrollmentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2EnrollmentGet`: EnrollmentSettingsV2
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V2EnrollmentGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentGetRequest struct via the builder pattern


### Return type

[**EnrollmentSettingsV2**](EnrollmentSettingsV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentHistoryExportPost

> interface{} V2EnrollmentHistoryExportPost(ctx).ExportFields(exportFields).ExportLabels(exportLabels).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).ExportParameters(exportParameters).Execute()

Export enrollment history collection 



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
    exportFields := []string{"Inner_example"} // []string | Export fields parameter, used to change default order or ignore some of the response properties. Default is empty array, which means that all fields of the response entity will be serialized. Example: export-fields=id,username (optional) (default to [])
    exportLabels := []string{"Inner_example"} // []string | Export labels parameter, used to customize fieldnames/columns in the exported file. Default is empty array, which means that response properties names will be used. Number of the provided labels must match the number of export-fields Example: export-labels=identifier,name with matching: export-fields=id,username (optional) (default to [])
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=id:desc,name:asc  (optional) (default to ["id:asc"])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name. This param can be combined with paging and sorting. Example: name==\"*script*\" (optional) (default to "")
    exportParameters := *openapiclient.NewExportParameters() // ExportParameters | Optional. Override query parameters since they can make URI exceed 2,000 character limit. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.V2EnrollmentHistoryExportPost(context.Background()).ExportFields(exportFields).ExportLabels(exportLabels).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).ExportParameters(exportParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V2EnrollmentHistoryExportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2EnrollmentHistoryExportPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V2EnrollmentHistoryExportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentHistoryExportPostRequest struct via the builder pattern


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


## V2EnrollmentHistoryGet

> HistorySearchResults V2EnrollmentHistoryGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Execute()

Get sorted and paged Enrollment history object 



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
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: `property:asc/desc`. Default sort is `date:desc`. Multiple sort criteria are supported and must be separated with a comma. Example: `sort=date:desc,name:asc`.  (optional) (default to ["date:desc"])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.V2EnrollmentHistoryGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V2EnrollmentHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2EnrollmentHistoryGet`: HistorySearchResults
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V2EnrollmentHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: &#x60;property:asc/desc&#x60;. Default sort is &#x60;date:desc&#x60;. Multiple sort criteria are supported and must be separated with a comma. Example: &#x60;sort&#x3D;date:desc,name:asc&#x60;.  | [default to [&quot;date:desc&quot;]]

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


## V2EnrollmentHistoryPost

> HrefResponse V2EnrollmentHistoryPost(ctx).ObjectHistoryNote(objectHistoryNote).Execute()

Add Enrollment history object notes 



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
    resp, r, err := apiClient.EnrollmentApi.V2EnrollmentHistoryPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V2EnrollmentHistoryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2EnrollmentHistoryPost`: HrefResponse
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V2EnrollmentHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentHistoryPostRequest struct via the builder pattern


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


## V2EnrollmentLanguageCodesGet

> []LanguageCode V2EnrollmentLanguageCodesGet(ctx).Execute()

Retrieve the list of languages and corresponding ISO 639-1 Codes 



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
    resp, r, err := apiClient.EnrollmentApi.V2EnrollmentLanguageCodesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V2EnrollmentLanguageCodesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2EnrollmentLanguageCodesGet`: []LanguageCode
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V2EnrollmentLanguageCodesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentLanguageCodesGetRequest struct via the builder pattern


### Return type

[**[]LanguageCode**](LanguageCode.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentLanguagesGet

> ProcessTextsSearchResults V2EnrollmentLanguagesGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Execute()

Get an array of the language codes that have Enrollment messaging 



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
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is `languageCode:asc`. Multiple sort criteria are supported and must be separated with a comma. Example: `sort=date:desc,name:asc`.  (optional) (default to ["languageCode:asc"])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.V2EnrollmentLanguagesGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V2EnrollmentLanguagesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2EnrollmentLanguagesGet`: ProcessTextsSearchResults
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V2EnrollmentLanguagesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentLanguagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is &#x60;languageCode:asc&#x60;. Multiple sort criteria are supported and must be separated with a comma. Example: &#x60;sort&#x3D;date:desc,name:asc&#x60;.  | [default to [&quot;languageCode:asc&quot;]]

### Return type

[**ProcessTextsSearchResults**](ProcessTextsSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentLanguagesLanguageIdDelete

> V2EnrollmentLanguagesLanguageIdDelete(ctx, languageId).Execute()

Delete the Enrollment messaging for a language 



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
    languageId := "languageId_example" // string | Two letter ISO 639-1 Language Code

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.V2EnrollmentLanguagesLanguageIdDelete(context.Background(), languageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V2EnrollmentLanguagesLanguageIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**languageId** | **string** | Two letter ISO 639-1 Language Code | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentLanguagesLanguageIdDeleteRequest struct via the builder pattern


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


## V2EnrollmentLanguagesLanguageIdGet

> EnrollmentProcessTextObject V2EnrollmentLanguagesLanguageIdGet(ctx, languageId).Execute()

Retrieve the Enrollment messaging for a language 



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
    languageId := "languageId_example" // string | Two letter ISO 639-1 Language Code

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.V2EnrollmentLanguagesLanguageIdGet(context.Background(), languageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V2EnrollmentLanguagesLanguageIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2EnrollmentLanguagesLanguageIdGet`: EnrollmentProcessTextObject
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V2EnrollmentLanguagesLanguageIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**languageId** | **string** | Two letter ISO 639-1 Language Code | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentLanguagesLanguageIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnrollmentProcessTextObject**](EnrollmentProcessTextObject.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentLanguagesLanguageIdPut

> EnrollmentProcessTextObject V2EnrollmentLanguagesLanguageIdPut(ctx, languageId).EnrollmentProcessTextObject(enrollmentProcessTextObject).Execute()

Edit Enrollment messaging for a language 



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
    languageId := "languageId_example" // string | Two letter ISO 639-1 Language Code
    enrollmentProcessTextObject := *openapiclient.NewEnrollmentProcessTextObject() // EnrollmentProcessTextObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.V2EnrollmentLanguagesLanguageIdPut(context.Background(), languageId).EnrollmentProcessTextObject(enrollmentProcessTextObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V2EnrollmentLanguagesLanguageIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2EnrollmentLanguagesLanguageIdPut`: EnrollmentProcessTextObject
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V2EnrollmentLanguagesLanguageIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**languageId** | **string** | Two letter ISO 639-1 Language Code | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentLanguagesLanguageIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enrollmentProcessTextObject** | [**EnrollmentProcessTextObject**](EnrollmentProcessTextObject.md) |  | 

### Return type

[**EnrollmentProcessTextObject**](EnrollmentProcessTextObject.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentPut

> EnrollmentSettingsV2 V2EnrollmentPut(ctx).EnrollmentSettingsV2(enrollmentSettingsV2).Execute()

Update Enrollment object 



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
    enrollmentSettingsV2 := *openapiclient.NewEnrollmentSettingsV2("radmin") // EnrollmentSettingsV2 | Update enrollment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentApi.V2EnrollmentPut(context.Background()).EnrollmentSettingsV2(enrollmentSettingsV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentApi.V2EnrollmentPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2EnrollmentPut`: EnrollmentSettingsV2
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentApi.V2EnrollmentPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enrollmentSettingsV2** | [**EnrollmentSettingsV2**](EnrollmentSettingsV2.md) | Update enrollment | 

### Return type

[**EnrollmentSettingsV2**](EnrollmentSettingsV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

