# \DigicertAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PkiDigicertTrustLifecycleManagerIdConnectionStatusGet**](DigicertAPI.md#V1PkiDigicertTrustLifecycleManagerIdConnectionStatusGet) | **Get** /v1/pki/digicert/trust-lifecycle-manager/{id}/connection-status | Get connection status of DigiCert Trust Lifecycle Manager for a given ID
[**V1PkiDigicertTrustLifecycleManagerIdDelete**](DigicertAPI.md#V1PkiDigicertTrustLifecycleManagerIdDelete) | **Delete** /v1/pki/digicert/trust-lifecycle-manager/{id} | Delete DigiCert Trust Lifecycle Manager configuration
[**V1PkiDigicertTrustLifecycleManagerIdDependenciesGet**](DigicertAPI.md#V1PkiDigicertTrustLifecycleManagerIdDependenciesGet) | **Get** /v1/pki/digicert/trust-lifecycle-manager/{id}/dependencies | Retrieve list of DigiCert Trust Lifecycle Manager Settings dependencies
[**V1PkiDigicertTrustLifecycleManagerIdGet**](DigicertAPI.md#V1PkiDigicertTrustLifecycleManagerIdGet) | **Get** /v1/pki/digicert/trust-lifecycle-manager/{id} | Retrieve DigiCert Trust Lifecycle Manager configuration
[**V1PkiDigicertTrustLifecycleManagerIdPatch**](DigicertAPI.md#V1PkiDigicertTrustLifecycleManagerIdPatch) | **Patch** /v1/pki/digicert/trust-lifecycle-manager/{id} | Update DigiCert Trust Lifecycle Manager configuration 
[**V1PkiDigicertTrustLifecycleManagerPost**](DigicertAPI.md#V1PkiDigicertTrustLifecycleManagerPost) | **Post** /v1/pki/digicert/trust-lifecycle-manager | Create DigiCert Trust Lifecycle Manager configuration with client authentication via client certificate. 
[**V1PkiDigicertTrustLifecycleManagerValidateClientCertificatePost**](DigicertAPI.md#V1PkiDigicertTrustLifecycleManagerValidateClientCertificatePost) | **Post** /v1/pki/digicert/trust-lifecycle-manager/validate-client-certificate | Validate DigiCert Trust Lifecycle Manager client certificate 



## V1PkiDigicertTrustLifecycleManagerIdConnectionStatusGet

> DigicertConnectionStatus V1PkiDigicertTrustLifecycleManagerIdConnectionStatusGet(ctx, id).Execute()

Get connection status of DigiCert Trust Lifecycle Manager for a given ID



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
	id := "12" // string | ID of the DigiCert Trust Lifecycle Manager settings.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigicertAPI.V1PkiDigicertTrustLifecycleManagerIdConnectionStatusGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigicertAPI.V1PkiDigicertTrustLifecycleManagerIdConnectionStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PkiDigicertTrustLifecycleManagerIdConnectionStatusGet`: DigicertConnectionStatus
	fmt.Fprintf(os.Stdout, "Response from `DigicertAPI.V1PkiDigicertTrustLifecycleManagerIdConnectionStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the DigiCert Trust Lifecycle Manager settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiDigicertTrustLifecycleManagerIdConnectionStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DigicertConnectionStatus**](DigicertConnectionStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PkiDigicertTrustLifecycleManagerIdDelete

> V1PkiDigicertTrustLifecycleManagerIdDelete(ctx, id).Execute()

Delete DigiCert Trust Lifecycle Manager configuration



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
	id := "id_example" // string | ID of the DigiCert Trust Lifecycle Manager configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DigicertAPI.V1PkiDigicertTrustLifecycleManagerIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigicertAPI.V1PkiDigicertTrustLifecycleManagerIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the DigiCert Trust Lifecycle Manager configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiDigicertTrustLifecycleManagerIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## V1PkiDigicertTrustLifecycleManagerIdDependenciesGet

> DigicertDependencies V1PkiDigicertTrustLifecycleManagerIdDependenciesGet(ctx, id).Execute()

Retrieve list of DigiCert Trust Lifecycle Manager Settings dependencies



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
	id := "id_example" // string | ID of the DigiCert Trust Lifecycle Manager configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigicertAPI.V1PkiDigicertTrustLifecycleManagerIdDependenciesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigicertAPI.V1PkiDigicertTrustLifecycleManagerIdDependenciesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PkiDigicertTrustLifecycleManagerIdDependenciesGet`: DigicertDependencies
	fmt.Fprintf(os.Stdout, "Response from `DigicertAPI.V1PkiDigicertTrustLifecycleManagerIdDependenciesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the DigiCert Trust Lifecycle Manager configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiDigicertTrustLifecycleManagerIdDependenciesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DigicertDependencies**](DigicertDependencies.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PkiDigicertTrustLifecycleManagerIdGet

> DigiCertSettingResponse V1PkiDigicertTrustLifecycleManagerIdGet(ctx, id).Execute()

Retrieve DigiCert Trust Lifecycle Manager configuration



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
	id := "id_example" // string | ID of the DigiCert Trust Lifecycle Manager configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigicertAPI.V1PkiDigicertTrustLifecycleManagerIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigicertAPI.V1PkiDigicertTrustLifecycleManagerIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PkiDigicertTrustLifecycleManagerIdGet`: DigiCertSettingResponse
	fmt.Fprintf(os.Stdout, "Response from `DigicertAPI.V1PkiDigicertTrustLifecycleManagerIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the DigiCert Trust Lifecycle Manager configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiDigicertTrustLifecycleManagerIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DigiCertSettingResponse**](DigiCertSettingResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PkiDigicertTrustLifecycleManagerIdPatch

> V1PkiDigicertTrustLifecycleManagerIdPatch(ctx, id).DigiCertSetting(digiCertSetting).Execute()

Update DigiCert Trust Lifecycle Manager configuration 



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
	id := "id_example" // string | ID of the DigiCert Trust Lifecycle Manager configuration.
	digiCertSetting := *openapiclient.NewDigiCertSetting() // DigiCertSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DigicertAPI.V1PkiDigicertTrustLifecycleManagerIdPatch(context.Background(), id).DigiCertSetting(digiCertSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigicertAPI.V1PkiDigicertTrustLifecycleManagerIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the DigiCert Trust Lifecycle Manager configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiDigicertTrustLifecycleManagerIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **digiCertSetting** | [**DigiCertSetting**](DigiCertSetting.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PkiDigicertTrustLifecycleManagerPost

> HrefResponse V1PkiDigicertTrustLifecycleManagerPost(ctx).DigiCertSetting(digiCertSetting).Execute()

Create DigiCert Trust Lifecycle Manager configuration with client authentication via client certificate. 



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
	digiCertSetting := *openapiclient.NewDigiCertSetting() // DigiCertSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigicertAPI.V1PkiDigicertTrustLifecycleManagerPost(context.Background()).DigiCertSetting(digiCertSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigicertAPI.V1PkiDigicertTrustLifecycleManagerPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PkiDigicertTrustLifecycleManagerPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `DigicertAPI.V1PkiDigicertTrustLifecycleManagerPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiDigicertTrustLifecycleManagerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **digiCertSetting** | [**DigiCertSetting**](DigiCertSetting.md) |  | 

### Return type

[**HrefResponse**](HrefResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PkiDigicertTrustLifecycleManagerValidateClientCertificatePost

> V1PkiDigicertTrustLifecycleManagerValidateClientCertificatePost(ctx).Certificate(certificate).Execute()

Validate DigiCert Trust Lifecycle Manager client certificate 



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
	certificate := *openapiclient.NewCertificate("example.p12", []string{string(123)}) // Certificate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DigicertAPI.V1PkiDigicertTrustLifecycleManagerValidateClientCertificatePost(context.Background()).Certificate(certificate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigicertAPI.V1PkiDigicertTrustLifecycleManagerValidateClientCertificatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiDigicertTrustLifecycleManagerValidateClientCertificatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificate** | [**Certificate**](Certificate.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

