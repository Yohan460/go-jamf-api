# \SsoSettingsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SsoDependenciesGet**](SsoSettingsAPI.md#V1SsoDependenciesGet) | **Get** /v1/sso/dependencies | Retrieve the list of Enrollment Customizations using SSO 
[**V1SsoDisablePost**](SsoSettingsAPI.md#V1SsoDisablePost) | **Post** /v1/sso/disable | Disable SSO 
[**V1SsoGet**](SsoSettingsAPI.md#V1SsoGet) | **Get** /v1/sso | Retrieve the current Single Sign On configuration settings 
[**V1SsoHistoryGet**](SsoSettingsAPI.md#V1SsoHistoryGet) | **Get** /v1/sso/history | Get SSO history object 
[**V1SsoHistoryPost**](SsoSettingsAPI.md#V1SsoHistoryPost) | **Post** /v1/sso/history | Add SSO history object notes 
[**V1SsoMetadataDownloadGet**](SsoSettingsAPI.md#V1SsoMetadataDownloadGet) | **Get** /v1/sso/metadata/download | Download the Jamf Pro SAML metadata file 
[**V1SsoPut**](SsoSettingsAPI.md#V1SsoPut) | **Put** /v1/sso | Updates the current Single Sign On configuration settings 
[**V1SsoValidatePost**](SsoSettingsAPI.md#V1SsoValidatePost) | **Post** /v1/sso/validate | Endpoint for validation of a saml metadata url 
[**V2SsoDependenciesGet**](SsoSettingsAPI.md#V2SsoDependenciesGet) | **Get** /v2/sso/dependencies | Retrieve the list of Enrollment Customizations using SSO 
[**V2SsoDisablePost**](SsoSettingsAPI.md#V2SsoDisablePost) | **Post** /v2/sso/disable | Disable SSO 
[**V2SsoGet**](SsoSettingsAPI.md#V2SsoGet) | **Get** /v2/sso | Retrieve the current Single Sign On configuration settings 
[**V2SsoHistoryGet**](SsoSettingsAPI.md#V2SsoHistoryGet) | **Get** /v2/sso/history | Get SSO history object 
[**V2SsoHistoryPost**](SsoSettingsAPI.md#V2SsoHistoryPost) | **Post** /v2/sso/history | Add SSO history object notes 
[**V2SsoMetadataDownloadGet**](SsoSettingsAPI.md#V2SsoMetadataDownloadGet) | **Get** /v2/sso/metadata/download | Download the Jamf Pro SAML metadata file 
[**V2SsoPut**](SsoSettingsAPI.md#V2SsoPut) | **Put** /v2/sso | Updates the current Single Sign On configuration settings 
[**V2SsoValidatePost**](SsoSettingsAPI.md#V2SsoValidatePost) | **Post** /v2/sso/validate | Endpoint for validation of a saml metadata url 



## V1SsoDependenciesGet

> EnrollmentCustomizationDependencies V1SsoDependenciesGet(ctx).Execute()

Retrieve the list of Enrollment Customizations using SSO 



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
	resp, r, err := apiClient.SsoSettingsAPI.V1SsoDependenciesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoSettingsAPI.V1SsoDependenciesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SsoDependenciesGet`: EnrollmentCustomizationDependencies
	fmt.Fprintf(os.Stdout, "Response from `SsoSettingsAPI.V1SsoDependenciesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SsoDependenciesGetRequest struct via the builder pattern


### Return type

[**EnrollmentCustomizationDependencies**](EnrollmentCustomizationDependencies.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SsoDisablePost

> V1SsoDisablePost(ctx).Execute()

Disable SSO 



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
	r, err := apiClient.SsoSettingsAPI.V1SsoDisablePost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoSettingsAPI.V1SsoDisablePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SsoDisablePostRequest struct via the builder pattern


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


## V1SsoGet

> SsoSettingsV1 V1SsoGet(ctx).Execute()

Retrieve the current Single Sign On configuration settings 



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
	resp, r, err := apiClient.SsoSettingsAPI.V1SsoGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoSettingsAPI.V1SsoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SsoGet`: SsoSettingsV1
	fmt.Fprintf(os.Stdout, "Response from `SsoSettingsAPI.V1SsoGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SsoGetRequest struct via the builder pattern


### Return type

[**SsoSettingsV1**](SsoSettingsV1.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SsoHistoryGet

> HistorySearchResults V1SsoHistoryGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get SSO history object 



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
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["id:desc"])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SsoSettingsAPI.V1SsoHistoryGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoSettingsAPI.V1SsoHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SsoHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `SsoSettingsAPI.V1SsoHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SsoHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;id:desc&quot;]]
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


## V1SsoHistoryPost

> HrefResponse V1SsoHistoryPost(ctx).ObjectHistoryNote(objectHistoryNote).Execute()

Add SSO history object notes 



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
	objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | history notes to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SsoSettingsAPI.V1SsoHistoryPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoSettingsAPI.V1SsoHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SsoHistoryPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `SsoSettingsAPI.V1SsoHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SsoHistoryPostRequest struct via the builder pattern


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


## V1SsoMetadataDownloadGet

> *os.File V1SsoMetadataDownloadGet(ctx).Execute()

Download the Jamf Pro SAML metadata file 



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
	resp, r, err := apiClient.SsoSettingsAPI.V1SsoMetadataDownloadGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoSettingsAPI.V1SsoMetadataDownloadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SsoMetadataDownloadGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SsoSettingsAPI.V1SsoMetadataDownloadGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SsoMetadataDownloadGetRequest struct via the builder pattern


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


## V1SsoPut

> SsoSettingsV1 V1SsoPut(ctx).SsoSettingsV1(ssoSettingsV1).Execute()

Updates the current Single Sign On configuration settings 



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
	ssoSettingsV1 := *openapiclient.NewSsoSettingsV1(false, false, false, false, false, false, "USERNAME", false, false, "GroupAttributeName_example", "GroupRdnKey_example", "ADFS", "saml/metadata", "URL") // SsoSettingsV1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SsoSettingsAPI.V1SsoPut(context.Background()).SsoSettingsV1(ssoSettingsV1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoSettingsAPI.V1SsoPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SsoPut`: SsoSettingsV1
	fmt.Fprintf(os.Stdout, "Response from `SsoSettingsAPI.V1SsoPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SsoPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ssoSettingsV1** | [**SsoSettingsV1**](SsoSettingsV1.md) |  | 

### Return type

[**SsoSettingsV1**](SsoSettingsV1.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SsoValidatePost

> V1SsoValidatePost(ctx).SsoMetadataUrl(ssoMetadataUrl).Execute()

Endpoint for validation of a saml metadata url 



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
	ssoMetadataUrl := *openapiclient.NewSsoMetadataUrl("https://your-idp.com/saml/metadata") // SsoMetadataUrl | url to validate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SsoSettingsAPI.V1SsoValidatePost(context.Background()).SsoMetadataUrl(ssoMetadataUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoSettingsAPI.V1SsoValidatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SsoValidatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ssoMetadataUrl** | [**SsoMetadataUrl**](SsoMetadataUrl.md) | url to validate | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SsoDependenciesGet

> EnrollmentCustomizationDependencies V2SsoDependenciesGet(ctx).Execute()

Retrieve the list of Enrollment Customizations using SSO 



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
	resp, r, err := apiClient.SsoSettingsAPI.V2SsoDependenciesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoSettingsAPI.V2SsoDependenciesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2SsoDependenciesGet`: EnrollmentCustomizationDependencies
	fmt.Fprintf(os.Stdout, "Response from `SsoSettingsAPI.V2SsoDependenciesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2SsoDependenciesGetRequest struct via the builder pattern


### Return type

[**EnrollmentCustomizationDependencies**](EnrollmentCustomizationDependencies.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SsoDisablePost

> V2SsoDisablePost(ctx).Execute()

Disable SSO 



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
	r, err := apiClient.SsoSettingsAPI.V2SsoDisablePost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoSettingsAPI.V2SsoDisablePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2SsoDisablePostRequest struct via the builder pattern


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


## V2SsoGet

> SsoSettingsV2 V2SsoGet(ctx).Execute()

Retrieve the current Single Sign On configuration settings 



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
	resp, r, err := apiClient.SsoSettingsAPI.V2SsoGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoSettingsAPI.V2SsoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2SsoGet`: SsoSettingsV2
	fmt.Fprintf(os.Stdout, "Response from `SsoSettingsAPI.V2SsoGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2SsoGetRequest struct via the builder pattern


### Return type

[**SsoSettingsV2**](SsoSettingsV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SsoHistoryGet

> HistorySearchResults V2SsoHistoryGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get SSO history object 



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
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["id:desc"])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SsoSettingsAPI.V2SsoHistoryGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoSettingsAPI.V2SsoHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2SsoHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `SsoSettingsAPI.V2SsoHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2SsoHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;id:desc&quot;]]
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


## V2SsoHistoryPost

> HrefResponse V2SsoHistoryPost(ctx).ObjectHistoryNote(objectHistoryNote).Execute()

Add SSO history object notes 



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
	objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | history notes to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SsoSettingsAPI.V2SsoHistoryPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoSettingsAPI.V2SsoHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2SsoHistoryPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `SsoSettingsAPI.V2SsoHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2SsoHistoryPostRequest struct via the builder pattern


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


## V2SsoMetadataDownloadGet

> *os.File V2SsoMetadataDownloadGet(ctx).Execute()

Download the Jamf Pro SAML metadata file 



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
	resp, r, err := apiClient.SsoSettingsAPI.V2SsoMetadataDownloadGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoSettingsAPI.V2SsoMetadataDownloadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2SsoMetadataDownloadGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SsoSettingsAPI.V2SsoMetadataDownloadGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2SsoMetadataDownloadGetRequest struct via the builder pattern


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


## V2SsoPut

> SsoSettingsV2 V2SsoPut(ctx).SsoSettingsV2(ssoSettingsV2).Execute()

Updates the current Single Sign On configuration settings 



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
	ssoSettingsV2 := *openapiclient.NewSsoSettingsV2(false, false, false, false, false, false, "USERNAME", false, false, "GroupAttributeName_example", "GroupRdnKey_example", "ADFS", "saml/metadata", "URL") // SsoSettingsV2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SsoSettingsAPI.V2SsoPut(context.Background()).SsoSettingsV2(ssoSettingsV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoSettingsAPI.V2SsoPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2SsoPut`: SsoSettingsV2
	fmt.Fprintf(os.Stdout, "Response from `SsoSettingsAPI.V2SsoPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2SsoPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ssoSettingsV2** | [**SsoSettingsV2**](SsoSettingsV2.md) |  | 

### Return type

[**SsoSettingsV2**](SsoSettingsV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SsoValidatePost

> V2SsoValidatePost(ctx).SsoMetadataUrl(ssoMetadataUrl).Execute()

Endpoint for validation of a saml metadata url 



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
	ssoMetadataUrl := *openapiclient.NewSsoMetadataUrl("https://your-idp.com/saml/metadata") // SsoMetadataUrl | url to validate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SsoSettingsAPI.V2SsoValidatePost(context.Background()).SsoMetadataUrl(ssoMetadataUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoSettingsAPI.V2SsoValidatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2SsoValidatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ssoMetadataUrl** | [**SsoMetadataUrl**](SsoMetadataUrl.md) | url to validate | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

