# \SupervisionIdentitiesPreviewAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SupervisionIdentitiesGet**](SupervisionIdentitiesPreviewAPI.md#V1SupervisionIdentitiesGet) | **Get** /v1/supervision-identities | Search for sorted and paged Supervision Identities 
[**V1SupervisionIdentitiesIdDelete**](SupervisionIdentitiesPreviewAPI.md#V1SupervisionIdentitiesIdDelete) | **Delete** /v1/supervision-identities/{id} | Delete a Supervision Identity with the supplied id 
[**V1SupervisionIdentitiesIdDownloadGet**](SupervisionIdentitiesPreviewAPI.md#V1SupervisionIdentitiesIdDownloadGet) | **Get** /v1/supervision-identities/{id}/download | Download the Supervision Identity .p12 file 
[**V1SupervisionIdentitiesIdGet**](SupervisionIdentitiesPreviewAPI.md#V1SupervisionIdentitiesIdGet) | **Get** /v1/supervision-identities/{id} | Retrieve a Supervision Identity with the supplied id 
[**V1SupervisionIdentitiesIdPut**](SupervisionIdentitiesPreviewAPI.md#V1SupervisionIdentitiesIdPut) | **Put** /v1/supervision-identities/{id} | Update a Supervision Identity with the supplied information 
[**V1SupervisionIdentitiesPost**](SupervisionIdentitiesPreviewAPI.md#V1SupervisionIdentitiesPost) | **Post** /v1/supervision-identities | Create a Supervision Identity for the supplied information 
[**V1SupervisionIdentitiesUploadPost**](SupervisionIdentitiesPreviewAPI.md#V1SupervisionIdentitiesUploadPost) | **Post** /v1/supervision-identities/upload | Upload the Supervision Identity .p12 file 



## V1SupervisionIdentitiesGet

> SupervisionIdentitySearchResults V1SupervisionIdentitiesGet(ctx).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()

Search for sorted and paged Supervision Identities 



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
    sort := "sort_example" // string | Sorting criteria in the format: property:asc/desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to "id:asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SupervisionIdentitiesPreviewAPI.V1SupervisionIdentitiesGet(context.Background()).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupervisionIdentitiesPreviewAPI.V1SupervisionIdentitiesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SupervisionIdentitiesGet`: SupervisionIdentitySearchResults
    fmt.Fprintf(os.Stdout, "Response from `SupervisionIdentitiesPreviewAPI.V1SupervisionIdentitiesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SupervisionIdentitiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]
 **pagesize** | **int32** |  | [default to 100]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **string** | Sorting criteria in the format: property:asc/desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to &quot;id:asc&quot;]

### Return type

[**SupervisionIdentitySearchResults**](SupervisionIdentitySearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SupervisionIdentitiesIdDelete

> V1SupervisionIdentitiesIdDelete(ctx, id).Execute()

Delete a Supervision Identity with the supplied id 



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
    id := int32(56) // int32 | Supervision Identity identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SupervisionIdentitiesPreviewAPI.V1SupervisionIdentitiesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupervisionIdentitiesPreviewAPI.V1SupervisionIdentitiesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Supervision Identity identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SupervisionIdentitiesIdDeleteRequest struct via the builder pattern


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


## V1SupervisionIdentitiesIdDownloadGet

> *os.File V1SupervisionIdentitiesIdDownloadGet(ctx, id).Execute()

Download the Supervision Identity .p12 file 



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
    id := int32(56) // int32 | Supervision Identity identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SupervisionIdentitiesPreviewAPI.V1SupervisionIdentitiesIdDownloadGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupervisionIdentitiesPreviewAPI.V1SupervisionIdentitiesIdDownloadGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SupervisionIdentitiesIdDownloadGet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SupervisionIdentitiesPreviewAPI.V1SupervisionIdentitiesIdDownloadGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Supervision Identity identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SupervisionIdentitiesIdDownloadGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SupervisionIdentitiesIdGet

> SupervisionIdentity V1SupervisionIdentitiesIdGet(ctx, id).Execute()

Retrieve a Supervision Identity with the supplied id 



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
    id := int32(56) // int32 | Supervision Identity identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SupervisionIdentitiesPreviewAPI.V1SupervisionIdentitiesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupervisionIdentitiesPreviewAPI.V1SupervisionIdentitiesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SupervisionIdentitiesIdGet`: SupervisionIdentity
    fmt.Fprintf(os.Stdout, "Response from `SupervisionIdentitiesPreviewAPI.V1SupervisionIdentitiesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Supervision Identity identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SupervisionIdentitiesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SupervisionIdentity**](SupervisionIdentity.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SupervisionIdentitiesIdPut

> SupervisionIdentity V1SupervisionIdentitiesIdPut(ctx, id).SupervisionIdentityUpdate(supervisionIdentityUpdate).Execute()

Update a Supervision Identity with the supplied information 



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
    id := int32(56) // int32 | Supervision Identity identifier
    supervisionIdentityUpdate := *openapiclient.NewSupervisionIdentityUpdate("Supervision Identity") // SupervisionIdentityUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SupervisionIdentitiesPreviewAPI.V1SupervisionIdentitiesIdPut(context.Background(), id).SupervisionIdentityUpdate(supervisionIdentityUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupervisionIdentitiesPreviewAPI.V1SupervisionIdentitiesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SupervisionIdentitiesIdPut`: SupervisionIdentity
    fmt.Fprintf(os.Stdout, "Response from `SupervisionIdentitiesPreviewAPI.V1SupervisionIdentitiesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Supervision Identity identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SupervisionIdentitiesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supervisionIdentityUpdate** | [**SupervisionIdentityUpdate**](SupervisionIdentityUpdate.md) |  | 

### Return type

[**SupervisionIdentity**](SupervisionIdentity.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SupervisionIdentitiesPost

> SupervisionIdentity V1SupervisionIdentitiesPost(ctx).SupervisionIdentityCreate(supervisionIdentityCreate).Execute()

Create a Supervision Identity for the supplied information 



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
    supervisionIdentityCreate := *openapiclient.NewSupervisionIdentityCreate("Supervision Identity", "jamf1234") // SupervisionIdentityCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SupervisionIdentitiesPreviewAPI.V1SupervisionIdentitiesPost(context.Background()).SupervisionIdentityCreate(supervisionIdentityCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupervisionIdentitiesPreviewAPI.V1SupervisionIdentitiesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SupervisionIdentitiesPost`: SupervisionIdentity
    fmt.Fprintf(os.Stdout, "Response from `SupervisionIdentitiesPreviewAPI.V1SupervisionIdentitiesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SupervisionIdentitiesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supervisionIdentityCreate** | [**SupervisionIdentityCreate**](SupervisionIdentityCreate.md) |  | 

### Return type

[**SupervisionIdentity**](SupervisionIdentity.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SupervisionIdentitiesUploadPost

> SupervisionIdentity V1SupervisionIdentitiesUploadPost(ctx).SupervisionIdentityCertificateUpload(supervisionIdentityCertificateUpload).Execute()

Upload the Supervision Identity .p12 file 



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
    supervisionIdentityCertificateUpload := *openapiclient.NewSupervisionIdentityCertificateUpload("Supervision Identity", "jamf1234") // SupervisionIdentityCertificateUpload | The base 64 encoded .p12 file alone with other needed information

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SupervisionIdentitiesPreviewAPI.V1SupervisionIdentitiesUploadPost(context.Background()).SupervisionIdentityCertificateUpload(supervisionIdentityCertificateUpload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupervisionIdentitiesPreviewAPI.V1SupervisionIdentitiesUploadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SupervisionIdentitiesUploadPost`: SupervisionIdentity
    fmt.Fprintf(os.Stdout, "Response from `SupervisionIdentitiesPreviewAPI.V1SupervisionIdentitiesUploadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SupervisionIdentitiesUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supervisionIdentityCertificateUpload** | [**SupervisionIdentityCertificateUpload**](SupervisionIdentityCertificateUpload.md) | The base 64 encoded .p12 file alone with other needed information | 

### Return type

[**SupervisionIdentity**](SupervisionIdentity.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

