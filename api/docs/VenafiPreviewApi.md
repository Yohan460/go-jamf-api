# \VenafiPreviewApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PkiVenafiIdConnectionStatusGet**](VenafiPreviewApi.md#V1PkiVenafiIdConnectionStatusGet) | **Get** /v1/pki/venafi/{id}/connection-status | Tests the communication between Jamf Pro and a Jamf Pro PKI Proxy Server 
[**V1PkiVenafiIdDelete**](VenafiPreviewApi.md#V1PkiVenafiIdDelete) | **Delete** /v1/pki/venafi/{id} | Delete a Venafi PKI configuration from Jamf Pro 
[**V1PkiVenafiIdDependentProfilesGet**](VenafiPreviewApi.md#V1PkiVenafiIdDependentProfilesGet) | **Get** /v1/pki/venafi/{id}/dependent-profiles | Get configuration profile data using specified Venafi CA object 
[**V1PkiVenafiIdGet**](VenafiPreviewApi.md#V1PkiVenafiIdGet) | **Get** /v1/pki/venafi/{id} | Retrieve a Venafi PKI configuration from Jamf Pro 
[**V1PkiVenafiIdHistoryGet**](VenafiPreviewApi.md#V1PkiVenafiIdHistoryGet) | **Get** /v1/pki/venafi/{id}/history | Get specified Venafi CA history object 
[**V1PkiVenafiIdHistoryPost**](VenafiPreviewApi.md#V1PkiVenafiIdHistoryPost) | **Post** /v1/pki/venafi/{id}/history | Add specified Venafi CA Object Note 
[**V1PkiVenafiIdJamfPublicKeyGet**](VenafiPreviewApi.md#V1PkiVenafiIdJamfPublicKeyGet) | **Get** /v1/pki/venafi/{id}/jamf-public-key | Downloads a certificate used to secure communication between Jamf Pro and a Jamf Pro PKI Proxy Server 
[**V1PkiVenafiIdJamfPublicKeyRegeneratePost**](VenafiPreviewApi.md#V1PkiVenafiIdJamfPublicKeyRegeneratePost) | **Post** /v1/pki/venafi/{id}/jamf-public-key/regenerate | Regenerates a certificate used to secure communication between Jamf Pro and a Jamf Pro PKI Proxy Server 
[**V1PkiVenafiIdPatch**](VenafiPreviewApi.md#V1PkiVenafiIdPatch) | **Patch** /v1/pki/venafi/{id} | Update a Venafi PKI configuration in Jamf Pro 
[**V1PkiVenafiIdProxyTrustStoreDelete**](VenafiPreviewApi.md#V1PkiVenafiIdProxyTrustStoreDelete) | **Delete** /v1/pki/venafi/{id}/proxy-trust-store | Removes the PKI Proxy Server public key used to secure communication between Jamf Pro and a Jamf Pro PKI Proxy Server 
[**V1PkiVenafiIdProxyTrustStoreGet**](VenafiPreviewApi.md#V1PkiVenafiIdProxyTrustStoreGet) | **Get** /v1/pki/venafi/{id}/proxy-trust-store | Downloads the PKI Proxy Server public key to secure communication between Jamf Pro and a Jamf Pro PKI Proxy Server 
[**V1PkiVenafiIdProxyTrustStorePost**](VenafiPreviewApi.md#V1PkiVenafiIdProxyTrustStorePost) | **Post** /v1/pki/venafi/{id}/proxy-trust-store | Uploads the PKI Proxy Server public key to secure communication between Jamf Pro and a Jamf Pro PKI Proxy Server 
[**V1PkiVenafiPost**](VenafiPreviewApi.md#V1PkiVenafiPost) | **Post** /v1/pki/venafi | Create a PKI configuration in Jamf Pro for Venafi 



## V1PkiVenafiIdConnectionStatusGet

> VenafiServiceStatus V1PkiVenafiIdConnectionStatusGet(ctx, id).Execute()

Tests the communication between Jamf Pro and a Jamf Pro PKI Proxy Server 



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
    id := "id_example" // string | ID of the Venafi configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VenafiPreviewApi.V1PkiVenafiIdConnectionStatusGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VenafiPreviewApi.V1PkiVenafiIdConnectionStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PkiVenafiIdConnectionStatusGet`: VenafiServiceStatus
    fmt.Fprintf(os.Stdout, "Response from `VenafiPreviewApi.V1PkiVenafiIdConnectionStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Venafi configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiVenafiIdConnectionStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VenafiServiceStatus**](VenafiServiceStatus.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PkiVenafiIdDelete

> V1PkiVenafiIdDelete(ctx, id).Execute()

Delete a Venafi PKI configuration from Jamf Pro 



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
    id := "id_example" // string | ID of the Venafi configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VenafiPreviewApi.V1PkiVenafiIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VenafiPreviewApi.V1PkiVenafiIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Venafi configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiVenafiIdDeleteRequest struct via the builder pattern


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


## V1PkiVenafiIdDependentProfilesGet

> VenafiPkiPayloadRecordSearchResults V1PkiVenafiIdDependentProfilesGet(ctx, id).Execute()

Get configuration profile data using specified Venafi CA object 



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
    id := "id_example" // string | ID of the Venafi configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VenafiPreviewApi.V1PkiVenafiIdDependentProfilesGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VenafiPreviewApi.V1PkiVenafiIdDependentProfilesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PkiVenafiIdDependentProfilesGet`: VenafiPkiPayloadRecordSearchResults
    fmt.Fprintf(os.Stdout, "Response from `VenafiPreviewApi.V1PkiVenafiIdDependentProfilesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Venafi configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiVenafiIdDependentProfilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VenafiPkiPayloadRecordSearchResults**](VenafiPkiPayloadRecordSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PkiVenafiIdGet

> VenafiCaRecord V1PkiVenafiIdGet(ctx, id).Execute()

Retrieve a Venafi PKI configuration from Jamf Pro 



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
    id := "id_example" // string | ID of the Venafi configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VenafiPreviewApi.V1PkiVenafiIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VenafiPreviewApi.V1PkiVenafiIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PkiVenafiIdGet`: VenafiCaRecord
    fmt.Fprintf(os.Stdout, "Response from `VenafiPreviewApi.V1PkiVenafiIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Venafi configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiVenafiIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VenafiCaRecord**](VenafiCaRecord.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PkiVenafiIdHistoryGet

> HistorySearchResults V1PkiVenafiIdHistoryGet(ctx, id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get specified Venafi CA history object 



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
    id := "id_example" // string | ID of the Venafi configuration
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["date:desc"])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VenafiPreviewApi.V1PkiVenafiIdHistoryGet(context.Background(), id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VenafiPreviewApi.V1PkiVenafiIdHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PkiVenafiIdHistoryGet`: HistorySearchResults
    fmt.Fprintf(os.Stdout, "Response from `VenafiPreviewApi.V1PkiVenafiIdHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Venafi configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiVenafiIdHistoryGetRequest struct via the builder pattern


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


## V1PkiVenafiIdHistoryPost

> HrefResponse V1PkiVenafiIdHistoryPost(ctx, id).ObjectHistoryNote(objectHistoryNote).Execute()

Add specified Venafi CA Object Note 



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
    id := "id_example" // string | instance id of department history record
    objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | venafi ca history notes to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VenafiPreviewApi.V1PkiVenafiIdHistoryPost(context.Background(), id).ObjectHistoryNote(objectHistoryNote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VenafiPreviewApi.V1PkiVenafiIdHistoryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PkiVenafiIdHistoryPost`: HrefResponse
    fmt.Fprintf(os.Stdout, "Response from `VenafiPreviewApi.V1PkiVenafiIdHistoryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of department history record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiVenafiIdHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectHistoryNote** | [**ObjectHistoryNote**](ObjectHistoryNote.md) | venafi ca history notes to create | 

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


## V1PkiVenafiIdJamfPublicKeyGet

> *os.File V1PkiVenafiIdJamfPublicKeyGet(ctx, id).Execute()

Downloads a certificate used to secure communication between Jamf Pro and a Jamf Pro PKI Proxy Server 



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
    id := "id_example" // string | ID of the Venafi configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VenafiPreviewApi.V1PkiVenafiIdJamfPublicKeyGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VenafiPreviewApi.V1PkiVenafiIdJamfPublicKeyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PkiVenafiIdJamfPublicKeyGet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `VenafiPreviewApi.V1PkiVenafiIdJamfPublicKeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Venafi configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiVenafiIdJamfPublicKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pem-certificate-chain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PkiVenafiIdJamfPublicKeyRegeneratePost

> V1PkiVenafiIdJamfPublicKeyRegeneratePost(ctx, id).Execute()

Regenerates a certificate used to secure communication between Jamf Pro and a Jamf Pro PKI Proxy Server 



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
    id := "id_example" // string | ID of the Venafi configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VenafiPreviewApi.V1PkiVenafiIdJamfPublicKeyRegeneratePost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VenafiPreviewApi.V1PkiVenafiIdJamfPublicKeyRegeneratePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Venafi configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiVenafiIdJamfPublicKeyRegeneratePostRequest struct via the builder pattern


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


## V1PkiVenafiIdPatch

> VenafiCaRecord V1PkiVenafiIdPatch(ctx, id).VenafiCaRecord(venafiCaRecord).Execute()

Update a Venafi PKI configuration in Jamf Pro 



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
    id := "id_example" // string | ID of the Venafi configuration
    venafiCaRecord := *openapiclient.NewVenafiCaRecord("Venafi Certificate Authority") // VenafiCaRecord | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VenafiPreviewApi.V1PkiVenafiIdPatch(context.Background(), id).VenafiCaRecord(venafiCaRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VenafiPreviewApi.V1PkiVenafiIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PkiVenafiIdPatch`: VenafiCaRecord
    fmt.Fprintf(os.Stdout, "Response from `VenafiPreviewApi.V1PkiVenafiIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Venafi configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiVenafiIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **venafiCaRecord** | [**VenafiCaRecord**](VenafiCaRecord.md) |  | 

### Return type

[**VenafiCaRecord**](VenafiCaRecord.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PkiVenafiIdProxyTrustStoreDelete

> V1PkiVenafiIdProxyTrustStoreDelete(ctx, id).Execute()

Removes the PKI Proxy Server public key used to secure communication between Jamf Pro and a Jamf Pro PKI Proxy Server 



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
    id := "id_example" // string | ID of the Venafi configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VenafiPreviewApi.V1PkiVenafiIdProxyTrustStoreDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VenafiPreviewApi.V1PkiVenafiIdProxyTrustStoreDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Venafi configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiVenafiIdProxyTrustStoreDeleteRequest struct via the builder pattern


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


## V1PkiVenafiIdProxyTrustStoreGet

> *os.File V1PkiVenafiIdProxyTrustStoreGet(ctx, id).Execute()

Downloads the PKI Proxy Server public key to secure communication between Jamf Pro and a Jamf Pro PKI Proxy Server 



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
    id := "id_example" // string | ID of the Venafi configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VenafiPreviewApi.V1PkiVenafiIdProxyTrustStoreGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VenafiPreviewApi.V1PkiVenafiIdProxyTrustStoreGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PkiVenafiIdProxyTrustStoreGet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `VenafiPreviewApi.V1PkiVenafiIdProxyTrustStoreGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Venafi configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiVenafiIdProxyTrustStoreGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pem-certificate-chain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PkiVenafiIdProxyTrustStorePost

> V1PkiVenafiIdProxyTrustStorePost(ctx, id).Body(body).Execute()

Uploads the PKI Proxy Server public key to secure communication between Jamf Pro and a Jamf Pro PKI Proxy Server 



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
    id := "id_example" // string | ID of the Venafi configuration
    body := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VenafiPreviewApi.V1PkiVenafiIdProxyTrustStorePost(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VenafiPreviewApi.V1PkiVenafiIdProxyTrustStorePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Venafi configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiVenafiIdProxyTrustStorePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/pem-certificate-chain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PkiVenafiPost

> HrefResponse V1PkiVenafiPost(ctx).VenafiCaRecord(venafiCaRecord).Execute()

Create a PKI configuration in Jamf Pro for Venafi 



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
    venafiCaRecord := *openapiclient.NewVenafiCaRecord("Venafi Certificate Authority") // VenafiCaRecord | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VenafiPreviewApi.V1PkiVenafiPost(context.Background()).VenafiCaRecord(venafiCaRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VenafiPreviewApi.V1PkiVenafiPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PkiVenafiPost`: HrefResponse
    fmt.Fprintf(os.Stdout, "Response from `VenafiPreviewApi.V1PkiVenafiPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiVenafiPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **venafiCaRecord** | [**VenafiCaRecord**](VenafiCaRecord.md) |  | 

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

