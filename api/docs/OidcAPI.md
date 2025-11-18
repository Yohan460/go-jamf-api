# \OidcAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1OidcDirectIdpLoginUrlGet**](OidcAPI.md#V1OidcDirectIdpLoginUrlGet) | **Get** /v1/oidc/direct-idp-login-url | Retrieve the URL to directly login to the IdP 
[**V1OidcDispatchPost**](OidcAPI.md#V1OidcDispatchPost) | **Post** /v1/oidc/dispatch | Provide the url to redirect for OIDC login
[**V1OidcGenerateCertificatePost**](OidcAPI.md#V1OidcGenerateCertificatePost) | **Post** /v1/oidc/generate-certificate | Generate a new keystore used for signing OIDC messages 
[**V1OidcPublicFeaturesGet**](OidcAPI.md#V1OidcPublicFeaturesGet) | **Get** /v1/oidc/public-features | Get the public features of the OIDC configuration
[**V1OidcPublicKeyGet**](OidcAPI.md#V1OidcPublicKeyGet) | **Get** /v1/oidc/public-key | Get the public key of the keystore used for signing OIDC messages as a JWT 



## V1OidcDirectIdpLoginUrlGet

> OidcDirectIdpLoginSkipUrl V1OidcDirectIdpLoginUrlGet(ctx).Execute()

Retrieve the URL to directly login to the IdP 



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
	resp, r, err := apiClient.OidcAPI.V1OidcDirectIdpLoginUrlGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OidcAPI.V1OidcDirectIdpLoginUrlGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OidcDirectIdpLoginUrlGet`: OidcDirectIdpLoginSkipUrl
	fmt.Fprintf(os.Stdout, "Response from `OidcAPI.V1OidcDirectIdpLoginUrlGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1OidcDirectIdpLoginUrlGetRequest struct via the builder pattern


### Return type

[**OidcDirectIdpLoginSkipUrl**](OidcDirectIdpLoginSkipUrl.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OidcDispatchPost

> OidcLoginDispatchResponse V1OidcDispatchPost(ctx).OidcLoginDispatchRequest(oidcLoginDispatchRequest).Execute()

Provide the url to redirect for OIDC login



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
	oidcLoginDispatchRequest := *openapiclient.NewOidcLoginDispatchRequest("aHR0cHM6Ly9qYW1mLXByby11cmwuY29tL2xvZ2dpbmcuaHRtbA==", "admin@domain.name") // OidcLoginDispatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OidcAPI.V1OidcDispatchPost(context.Background()).OidcLoginDispatchRequest(oidcLoginDispatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OidcAPI.V1OidcDispatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OidcDispatchPost`: OidcLoginDispatchResponse
	fmt.Fprintf(os.Stdout, "Response from `OidcAPI.V1OidcDispatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1OidcDispatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oidcLoginDispatchRequest** | [**OidcLoginDispatchRequest**](OidcLoginDispatchRequest.md) |  | 

### Return type

[**OidcLoginDispatchResponse**](OidcLoginDispatchResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OidcGenerateCertificatePost

> V1OidcGenerateCertificatePost(ctx).Execute()

Generate a new keystore used for signing OIDC messages 



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
	r, err := apiClient.OidcAPI.V1OidcGenerateCertificatePost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OidcAPI.V1OidcGenerateCertificatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1OidcGenerateCertificatePostRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OidcPublicFeaturesGet

> OidcPublicFeaturesResponse V1OidcPublicFeaturesGet(ctx).Execute()

Get the public features of the OIDC configuration



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
	resp, r, err := apiClient.OidcAPI.V1OidcPublicFeaturesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OidcAPI.V1OidcPublicFeaturesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OidcPublicFeaturesGet`: OidcPublicFeaturesResponse
	fmt.Fprintf(os.Stdout, "Response from `OidcAPI.V1OidcPublicFeaturesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1OidcPublicFeaturesGetRequest struct via the builder pattern


### Return type

[**OidcPublicFeaturesResponse**](OidcPublicFeaturesResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OidcPublicKeyGet

> OidcJwksResponse V1OidcPublicKeyGet(ctx).Execute()

Get the public key of the keystore used for signing OIDC messages as a JWT 



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
	resp, r, err := apiClient.OidcAPI.V1OidcPublicKeyGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OidcAPI.V1OidcPublicKeyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OidcPublicKeyGet`: OidcJwksResponse
	fmt.Fprintf(os.Stdout, "Response from `OidcAPI.V1OidcPublicKeyGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1OidcPublicKeyGetRequest struct via the builder pattern


### Return type

[**OidcJwksResponse**](OidcJwksResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

