# \SsoCertificatePreviewAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SsoCertDelete**](SsoCertificatePreviewAPI.md#V1SsoCertDelete) | **Delete** /v1/sso/cert | Delete the currently configured certificate used by SSO 
[**V1SsoCertDownloadGet**](SsoCertificatePreviewAPI.md#V1SsoCertDownloadGet) | **Get** /v1/sso/cert/download | Download the certificate currently configured for use with Jamf Pro&#39;s SSO configuration 
[**V1SsoCertGet**](SsoCertificatePreviewAPI.md#V1SsoCertGet) | **Get** /v1/sso/cert | Retrieve the certificate currently configured for use with SSO 
[**V1SsoCertParsePost**](SsoCertificatePreviewAPI.md#V1SsoCertParsePost) | **Post** /v1/sso/cert/parse | Parse the certificate to get details about certificate type and keys needed to upload certificate file 
[**V1SsoCertPost**](SsoCertificatePreviewAPI.md#V1SsoCertPost) | **Post** /v1/sso/cert | Jamf Pro will generate a new certificate and use it to sign SSO 
[**V1SsoCertPut**](SsoCertificatePreviewAPI.md#V1SsoCertPut) | **Put** /v1/sso/cert | Update the certificate used by Jamf Pro to sign SSO requests to the identify provider 



## V1SsoCertDelete

> V1SsoCertDelete(ctx).Execute()

Delete the currently configured certificate used by SSO 



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SsoCertificatePreviewAPI.V1SsoCertDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoCertificatePreviewAPI.V1SsoCertDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SsoCertDeleteRequest struct via the builder pattern


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


## V1SsoCertDownloadGet

> *os.File V1SsoCertDownloadGet(ctx).Execute()

Download the certificate currently configured for use with Jamf Pro's SSO configuration 



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SsoCertificatePreviewAPI.V1SsoCertDownloadGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoCertificatePreviewAPI.V1SsoCertDownloadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SsoCertDownloadGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SsoCertificatePreviewAPI.V1SsoCertDownloadGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SsoCertDownloadGetRequest struct via the builder pattern


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


## V1SsoCertGet

> SsoKeystoreWithDetails V1SsoCertGet(ctx).Execute()

Retrieve the certificate currently configured for use with SSO 



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SsoCertificatePreviewAPI.V1SsoCertGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoCertificatePreviewAPI.V1SsoCertGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SsoCertGet`: SsoKeystoreWithDetails
	fmt.Fprintf(os.Stdout, "Response from `SsoCertificatePreviewAPI.V1SsoCertGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SsoCertGetRequest struct via the builder pattern


### Return type

[**SsoKeystoreWithDetails**](SsoKeystoreWithDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SsoCertParsePost

> SsoKeystoreCertParseResponse V1SsoCertParsePost(ctx).SsoKeystoreParse(ssoKeystoreParse).Execute()

Parse the certificate to get details about certificate type and keys needed to upload certificate file 



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
	ssoKeystoreParse := *openapiclient.NewSsoKeystoreParse("***", string([B@138a7441), "keystore.p12") // SsoKeystoreParse | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SsoCertificatePreviewAPI.V1SsoCertParsePost(context.Background()).SsoKeystoreParse(ssoKeystoreParse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoCertificatePreviewAPI.V1SsoCertParsePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SsoCertParsePost`: SsoKeystoreCertParseResponse
	fmt.Fprintf(os.Stdout, "Response from `SsoCertificatePreviewAPI.V1SsoCertParsePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SsoCertParsePostRequest struct via the builder pattern


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


## V1SsoCertPost

> SsoKeystoreWithDetails V1SsoCertPost(ctx).Execute()

Jamf Pro will generate a new certificate and use it to sign SSO 



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SsoCertificatePreviewAPI.V1SsoCertPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoCertificatePreviewAPI.V1SsoCertPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SsoCertPost`: SsoKeystoreWithDetails
	fmt.Fprintf(os.Stdout, "Response from `SsoCertificatePreviewAPI.V1SsoCertPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SsoCertPostRequest struct via the builder pattern


### Return type

[**SsoKeystoreWithDetails**](SsoKeystoreWithDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SsoCertPut

> SsoKeystoreWithDetails V1SsoCertPut(ctx).SsoKeystore(ssoKeystore).Execute()

Update the certificate used by Jamf Pro to sign SSO requests to the identify provider 



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
	ssoKeystore := *openapiclient.NewSsoKeystore("Key_example", "***", "PKCS12", "***", string([B@138a7441), "keystore.p12") // SsoKeystore | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SsoCertificatePreviewAPI.V1SsoCertPut(context.Background()).SsoKeystore(ssoKeystore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoCertificatePreviewAPI.V1SsoCertPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SsoCertPut`: SsoKeystoreWithDetails
	fmt.Fprintf(os.Stdout, "Response from `SsoCertificatePreviewAPI.V1SsoCertPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SsoCertPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ssoKeystore** | [**SsoKeystore**](SsoKeystore.md) |  | 

### Return type

[**SsoKeystoreWithDetails**](SsoKeystoreWithDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

