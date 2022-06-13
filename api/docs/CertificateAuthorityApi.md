# \CertificateAuthorityApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PkiCertificateAuthorityActiveDerGet**](CertificateAuthorityApi.md#V1PkiCertificateAuthorityActiveDerGet) | **Get** /v1/pki/certificate-authority/active/der | Returns X.509 of active Certificate Authority (CA) in DER format
[**V1PkiCertificateAuthorityActiveGet**](CertificateAuthorityApi.md#V1PkiCertificateAuthorityActiveGet) | **Get** /v1/pki/certificate-authority/active | Returns X.509 details of the active Certificate Authority (CA)
[**V1PkiCertificateAuthorityActivePemGet**](CertificateAuthorityApi.md#V1PkiCertificateAuthorityActivePemGet) | **Get** /v1/pki/certificate-authority/active/pem | Returns active Certificate Authority (CA) in PEM format
[**V1PkiCertificateAuthorityIdDerGet**](CertificateAuthorityApi.md#V1PkiCertificateAuthorityIdDerGet) | **Get** /v1/pki/certificate-authority/{id}/der | Returns X.509 current Certificate Authority (CA) with provided ID in DER format
[**V1PkiCertificateAuthorityIdGet**](CertificateAuthorityApi.md#V1PkiCertificateAuthorityIdGet) | **Get** /v1/pki/certificate-authority/{id} | Returns X.509 details of Certificate Authority (CA) with provided ID
[**V1PkiCertificateAuthorityIdPemGet**](CertificateAuthorityApi.md#V1PkiCertificateAuthorityIdPemGet) | **Get** /v1/pki/certificate-authority/{id}/pem | Returns current Certificate Authority (CA) with provided ID in PEM format



## V1PkiCertificateAuthorityActiveDerGet

> *os.File V1PkiCertificateAuthorityActiveDerGet(ctx).Execute()

Returns X.509 of active Certificate Authority (CA) in DER format



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
    resp, r, err := apiClient.CertificateAuthorityApi.V1PkiCertificateAuthorityActiveDerGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.V1PkiCertificateAuthorityActiveDerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PkiCertificateAuthorityActiveDerGet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.V1PkiCertificateAuthorityActiveDerGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiCertificateAuthorityActiveDerGetRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pkix-cert

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PkiCertificateAuthorityActiveGet

> CertificateRecord V1PkiCertificateAuthorityActiveGet(ctx).Execute()

Returns X.509 details of the active Certificate Authority (CA)



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
    resp, r, err := apiClient.CertificateAuthorityApi.V1PkiCertificateAuthorityActiveGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.V1PkiCertificateAuthorityActiveGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PkiCertificateAuthorityActiveGet`: CertificateRecord
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.V1PkiCertificateAuthorityActiveGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiCertificateAuthorityActiveGetRequest struct via the builder pattern


### Return type

[**CertificateRecord**](CertificateRecord.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PkiCertificateAuthorityActivePemGet

> *os.File V1PkiCertificateAuthorityActivePemGet(ctx).Execute()

Returns active Certificate Authority (CA) in PEM format



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
    resp, r, err := apiClient.CertificateAuthorityApi.V1PkiCertificateAuthorityActivePemGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.V1PkiCertificateAuthorityActivePemGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PkiCertificateAuthorityActivePemGet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.V1PkiCertificateAuthorityActivePemGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiCertificateAuthorityActivePemGetRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pem-certificate-chain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PkiCertificateAuthorityIdDerGet

> *os.File V1PkiCertificateAuthorityIdDerGet(ctx, id).Execute()

Returns X.509 current Certificate Authority (CA) with provided ID in DER format



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
    id := "id_example" // string | UUID of the Certificate Authority (CA)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.V1PkiCertificateAuthorityIdDerGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.V1PkiCertificateAuthorityIdDerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PkiCertificateAuthorityIdDerGet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.V1PkiCertificateAuthorityIdDerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UUID of the Certificate Authority (CA) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiCertificateAuthorityIdDerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pkix-cert, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PkiCertificateAuthorityIdGet

> CertificateRecord V1PkiCertificateAuthorityIdGet(ctx, id).Execute()

Returns X.509 details of Certificate Authority (CA) with provided ID



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
    id := "id_example" // string | UUID of the Certificate Authority (CA)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.V1PkiCertificateAuthorityIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.V1PkiCertificateAuthorityIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PkiCertificateAuthorityIdGet`: CertificateRecord
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.V1PkiCertificateAuthorityIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UUID of the Certificate Authority (CA) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiCertificateAuthorityIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CertificateRecord**](CertificateRecord.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PkiCertificateAuthorityIdPemGet

> *os.File V1PkiCertificateAuthorityIdPemGet(ctx, id).Execute()

Returns current Certificate Authority (CA) with provided ID in PEM format



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
    id := "id_example" // string | UUID of the Certificate Authority (CA)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateAuthorityApi.V1PkiCertificateAuthorityIdPemGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthorityApi.V1PkiCertificateAuthorityIdPemGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PkiCertificateAuthorityIdPemGet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthorityApi.V1PkiCertificateAuthorityIdPemGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UUID of the Certificate Authority (CA) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiCertificateAuthorityIdPemGetRequest struct via the builder pattern


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

