# \SsoCertificateAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2SsoCertDelete**](SsoCertificateAPI.md#V2SsoCertDelete) | **Delete** /v2/sso/cert | Delete the currently configured certificate used by SSO 
[**V2SsoCertDownloadGet**](SsoCertificateAPI.md#V2SsoCertDownloadGet) | **Get** /v2/sso/cert/download | Download the certificate currently configured for use with Jamf Pro&#39;s SSO configuration 
[**V2SsoCertGet**](SsoCertificateAPI.md#V2SsoCertGet) | **Get** /v2/sso/cert | Retrieve the certificate currently configured for use with SSO 
[**V2SsoCertParsePost**](SsoCertificateAPI.md#V2SsoCertParsePost) | **Post** /v2/sso/cert/parse | Parse the certificate to get details about certificate type and keys needed to upload certificate file 
[**V2SsoCertPost**](SsoCertificateAPI.md#V2SsoCertPost) | **Post** /v2/sso/cert | Jamf Pro will generate a new certificate and use it to sign SSO 
[**V2SsoCertPut**](SsoCertificateAPI.md#V2SsoCertPut) | **Put** /v2/sso/cert | Update the certificate used by Jamf Pro to sign SSO requests to the identify provider 



## V2SsoCertDelete

> V2SsoCertDelete(ctx).Execute()

Delete the currently configured certificate used by SSO 



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
    r, err := apiClient.SsoCertificateAPI.V2SsoCertDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsoCertificateAPI.V2SsoCertDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2SsoCertDeleteRequest struct via the builder pattern


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


## V2SsoCertDownloadGet

> *os.File V2SsoCertDownloadGet(ctx).Execute()

Download the certificate currently configured for use with Jamf Pro's SSO configuration 



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
    resp, r, err := apiClient.SsoCertificateAPI.V2SsoCertDownloadGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsoCertificateAPI.V2SsoCertDownloadGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2SsoCertDownloadGet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SsoCertificateAPI.V2SsoCertDownloadGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2SsoCertDownloadGetRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SsoCertGet

> SsoKeystoreResponseWithDetails V2SsoCertGet(ctx).Execute()

Retrieve the certificate currently configured for use with SSO 



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
    resp, r, err := apiClient.SsoCertificateAPI.V2SsoCertGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsoCertificateAPI.V2SsoCertGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2SsoCertGet`: SsoKeystoreResponseWithDetails
    fmt.Fprintf(os.Stdout, "Response from `SsoCertificateAPI.V2SsoCertGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2SsoCertGetRequest struct via the builder pattern


### Return type

[**SsoKeystoreResponseWithDetails**](SsoKeystoreResponseWithDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SsoCertParsePost

> SsoKeystoreCertParseResponse V2SsoCertParsePost(ctx).SsoKeystoreParse(ssoKeystoreParse).Execute()

Parse the certificate to get details about certificate type and keys needed to upload certificate file 



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
    ssoKeystoreParse := *openapiclient.NewSsoKeystoreParse("***", string([B@ca66933), "keystore.p12") // SsoKeystoreParse | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsoCertificateAPI.V2SsoCertParsePost(context.Background()).SsoKeystoreParse(ssoKeystoreParse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsoCertificateAPI.V2SsoCertParsePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2SsoCertParsePost`: SsoKeystoreCertParseResponse
    fmt.Fprintf(os.Stdout, "Response from `SsoCertificateAPI.V2SsoCertParsePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2SsoCertParsePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ssoKeystoreParse** | [**SsoKeystoreParse**](SsoKeystoreParse.md) |  | 

### Return type

[**SsoKeystoreCertParseResponse**](SsoKeystoreCertParseResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SsoCertPost

> SsoKeystoreResponseWithDetails V2SsoCertPost(ctx).Execute()

Jamf Pro will generate a new certificate and use it to sign SSO 



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
    resp, r, err := apiClient.SsoCertificateAPI.V2SsoCertPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsoCertificateAPI.V2SsoCertPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2SsoCertPost`: SsoKeystoreResponseWithDetails
    fmt.Fprintf(os.Stdout, "Response from `SsoCertificateAPI.V2SsoCertPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2SsoCertPostRequest struct via the builder pattern


### Return type

[**SsoKeystoreResponseWithDetails**](SsoKeystoreResponseWithDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SsoCertPut

> SsoKeystoreResponseWithDetails V2SsoCertPut(ctx).SsoKeystore(ssoKeystore).Execute()

Update the certificate used by Jamf Pro to sign SSO requests to the identify provider 



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
    ssoKeystore := *openapiclient.NewSsoKeystore("Key_example", "***", "PKCS12", "***", string([B@ca66933), "keystore.p12") // SsoKeystore | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsoCertificateAPI.V2SsoCertPut(context.Background()).SsoKeystore(ssoKeystore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsoCertificateAPI.V2SsoCertPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2SsoCertPut`: SsoKeystoreResponseWithDetails
    fmt.Fprintf(os.Stdout, "Response from `SsoCertificateAPI.V2SsoCertPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2SsoCertPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ssoKeystore** | [**SsoKeystore**](SsoKeystore.md) |  | 

### Return type

[**SsoKeystoreResponseWithDetails**](SsoKeystoreResponseWithDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

