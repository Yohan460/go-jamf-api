# \AdcsSettingsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PkiAdcsSettingsIdDelete**](AdcsSettingsAPI.md#V1PkiAdcsSettingsIdDelete) | **Delete** /v1/pki/adcs-settings/{id} | Delete AD CS Settings configuration by ID 
[**V1PkiAdcsSettingsIdDependenciesGet**](AdcsSettingsAPI.md#V1PkiAdcsSettingsIdDependenciesGet) | **Get** /v1/pki/adcs-settings/{id}/dependencies | Retrieve list of AD CS Settings dependencies
[**V1PkiAdcsSettingsIdGet**](AdcsSettingsAPI.md#V1PkiAdcsSettingsIdGet) | **Get** /v1/pki/adcs-settings/{id} | Get AD CS Settings configuration for the ID value 
[**V1PkiAdcsSettingsIdHistoryGet**](AdcsSettingsAPI.md#V1PkiAdcsSettingsIdHistoryGet) | **Get** /v1/pki/adcs-settings/{id}/history | Get specified AD CS Settings history object 
[**V1PkiAdcsSettingsIdHistoryPost**](AdcsSettingsAPI.md#V1PkiAdcsSettingsIdHistoryPost) | **Post** /v1/pki/adcs-settings/{id}/history | Add specified AD CS Settings object note 
[**V1PkiAdcsSettingsIdPatch**](AdcsSettingsAPI.md#V1PkiAdcsSettingsIdPatch) | **Patch** /v1/pki/adcs-settings/{id} | Update AD CS Settings configuration 
[**V1PkiAdcsSettingsPost**](AdcsSettingsAPI.md#V1PkiAdcsSettingsPost) | **Post** /v1/pki/adcs-settings | Create AD CS Settings configuration for either inbound or outbound mode 
[**V1PkiAdcsSettingsValidateCertificatePost**](AdcsSettingsAPI.md#V1PkiAdcsSettingsValidateCertificatePost) | **Post** /v1/pki/adcs-settings/validate-certificate | Validate AD CS Settings server certificate 
[**V1PkiAdcsSettingsValidateClientCertificatePost**](AdcsSettingsAPI.md#V1PkiAdcsSettingsValidateClientCertificatePost) | **Post** /v1/pki/adcs-settings/validate-client-certificate | Validate AD CS Settings client certificate 



## V1PkiAdcsSettingsIdDelete

> V1PkiAdcsSettingsIdDelete(ctx, id).Execute()

Delete AD CS Settings configuration by ID 



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
	id := "id_example" // string | ID of the AD CS Settings configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdcsSettingsAPI.V1PkiAdcsSettingsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdcsSettingsAPI.V1PkiAdcsSettingsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the AD CS Settings configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiAdcsSettingsIdDeleteRequest struct via the builder pattern


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


## V1PkiAdcsSettingsIdDependenciesGet

> AdcsDependencies V1PkiAdcsSettingsIdDependenciesGet(ctx, id).Execute()

Retrieve list of AD CS Settings dependencies



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
	id := "id_example" // string | AD CS Settings ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdcsSettingsAPI.V1PkiAdcsSettingsIdDependenciesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdcsSettingsAPI.V1PkiAdcsSettingsIdDependenciesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PkiAdcsSettingsIdDependenciesGet`: AdcsDependencies
	fmt.Fprintf(os.Stdout, "Response from `AdcsSettingsAPI.V1PkiAdcsSettingsIdDependenciesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | AD CS Settings ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiAdcsSettingsIdDependenciesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdcsDependencies**](AdcsDependencies.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PkiAdcsSettingsIdGet

> AdcsSettingsResponse V1PkiAdcsSettingsIdGet(ctx, id).Execute()

Get AD CS Settings configuration for the ID value 



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
	id := "id_example" // string | ID of the AD CS Settings configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdcsSettingsAPI.V1PkiAdcsSettingsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdcsSettingsAPI.V1PkiAdcsSettingsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PkiAdcsSettingsIdGet`: AdcsSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `AdcsSettingsAPI.V1PkiAdcsSettingsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the AD CS Settings configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiAdcsSettingsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdcsSettingsResponse**](AdcsSettingsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PkiAdcsSettingsIdHistoryGet

> HistorySearchResults V1PkiAdcsSettingsIdHistoryGet(ctx, id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get specified AD CS Settings history object 



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
	id := "id_example" // string | ID of the AD CS Settings configuration.
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to {"date:desc"})
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdcsSettingsAPI.V1PkiAdcsSettingsIdHistoryGet(context.Background(), id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdcsSettingsAPI.V1PkiAdcsSettingsIdHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PkiAdcsSettingsIdHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `AdcsSettingsAPI.V1PkiAdcsSettingsIdHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the AD CS Settings configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiAdcsSettingsIdHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to {&quot;date:desc&quot;}]
 **filter** | **string** | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15 | [default to &quot;&quot;]

### Return type

[**HistorySearchResults**](HistorySearchResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PkiAdcsSettingsIdHistoryPost

> HrefResponse V1PkiAdcsSettingsIdHistoryPost(ctx, id).ObjectHistoryNote(objectHistoryNote).Execute()

Add specified AD CS Settings object note 



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
	id := "id_example" // string | Instance ID of AD CS Settings history record.
	objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | AD CS Settings history notes to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdcsSettingsAPI.V1PkiAdcsSettingsIdHistoryPost(context.Background(), id).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdcsSettingsAPI.V1PkiAdcsSettingsIdHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PkiAdcsSettingsIdHistoryPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `AdcsSettingsAPI.V1PkiAdcsSettingsIdHistoryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Instance ID of AD CS Settings history record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiAdcsSettingsIdHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectHistoryNote** | [**ObjectHistoryNote**](ObjectHistoryNote.md) | AD CS Settings history notes to create. | 

### Return type

[**HrefResponse**](HrefResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PkiAdcsSettingsIdPatch

> V1PkiAdcsSettingsIdPatch(ctx, id).AdcsSettings(adcsSettings).Execute()

Update AD CS Settings configuration 



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
	id := "id_example" // string | ID of the AD CS Settings configuration.
	adcsSettings := *openapiclient.NewAdcsSettings() // AdcsSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdcsSettingsAPI.V1PkiAdcsSettingsIdPatch(context.Background(), id).AdcsSettings(adcsSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdcsSettingsAPI.V1PkiAdcsSettingsIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the AD CS Settings configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiAdcsSettingsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adcsSettings** | [**AdcsSettings**](AdcsSettings.md) |  | 

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


## V1PkiAdcsSettingsPost

> HrefResponse V1PkiAdcsSettingsPost(ctx).AdcsSettings(adcsSettings).Execute()

Create AD CS Settings configuration for either inbound or outbound mode 



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
	adcsSettings := *openapiclient.NewAdcsSettings() // AdcsSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdcsSettingsAPI.V1PkiAdcsSettingsPost(context.Background()).AdcsSettings(adcsSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdcsSettingsAPI.V1PkiAdcsSettingsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PkiAdcsSettingsPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `AdcsSettingsAPI.V1PkiAdcsSettingsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiAdcsSettingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adcsSettings** | [**AdcsSettings**](AdcsSettings.md) |  | 

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


## V1PkiAdcsSettingsValidateCertificatePost

> V1PkiAdcsSettingsValidateCertificatePost(ctx).AdcsCertificate(adcsCertificate).Execute()

Validate AD CS Settings server certificate 



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
	adcsCertificate := *openapiclient.NewAdcsCertificate("example.cer", []string{string(123)}) // AdcsCertificate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdcsSettingsAPI.V1PkiAdcsSettingsValidateCertificatePost(context.Background()).AdcsCertificate(adcsCertificate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdcsSettingsAPI.V1PkiAdcsSettingsValidateCertificatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiAdcsSettingsValidateCertificatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adcsCertificate** | [**AdcsCertificate**](AdcsCertificate.md) |  | 

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


## V1PkiAdcsSettingsValidateClientCertificatePost

> V1PkiAdcsSettingsValidateClientCertificatePost(ctx).AdcsCertificate(adcsCertificate).Execute()

Validate AD CS Settings client certificate 



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
	adcsCertificate := *openapiclient.NewAdcsCertificate("example.cer", []string{string(123)}) // AdcsCertificate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdcsSettingsAPI.V1PkiAdcsSettingsValidateClientCertificatePost(context.Background()).AdcsCertificate(adcsCertificate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdcsSettingsAPI.V1PkiAdcsSettingsValidateClientCertificatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1PkiAdcsSettingsValidateClientCertificatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adcsCertificate** | [**AdcsCertificate**](AdcsCertificate.md) |  | 

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

