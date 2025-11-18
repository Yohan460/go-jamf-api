# \ActivationCodeAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ActivationCodeHistoryExportPost**](ActivationCodeAPI.md#V1ActivationCodeHistoryExportPost) | **Post** /v1/activation-code/history/export | Export history object collection in specified format for Activation Code 
[**V1ActivationCodeHistoryGet**](ActivationCodeAPI.md#V1ActivationCodeHistoryGet) | **Get** /v1/activation-code/history | Get Activation Code history object
[**V1ActivationCodeHistoryPost**](ActivationCodeAPI.md#V1ActivationCodeHistoryPost) | **Post** /v1/activation-code/history | Add Activation Code object note 
[**V1ActivationCodeOrganizationNamePatch**](ActivationCodeAPI.md#V1ActivationCodeOrganizationNamePatch) | **Patch** /v1/activation-code/organization-name | Updates Organization Name
[**V1ActivationCodePut**](ActivationCodeAPI.md#V1ActivationCodePut) | **Put** /v1/activation-code | Updates Activation Code



## V1ActivationCodeHistoryExportPost

> interface{} V1ActivationCodeHistoryExportPost(ctx).ExportFields(exportFields).ExportLabels(exportLabels).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).ExportParameters(exportParameters).Execute()

Export history object collection in specified format for Activation Code 



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
	exportFields := []string{"Inner_example"} // []string | Export fields parameter, used to change default order or ignore some of the response properties. Default is empty array, which means that all fields of the response entity will be serialized. Example: export-fields=id,username (optional) (default to {})
	exportLabels := []string{"Inner_example"} // []string | Export labels parameter, used to customize fieldnames/columns in the exported file. Default is empty array, which means that response properties names will be used. Number of the provided labels must match the number of export-fields Example: export-labels=identifier,name with matching: export-fields=id,username (optional) (default to {})
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to {"date:desc"})
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")
	exportParameters := *openapiclient.NewExportParameters() // ExportParameters | Optional. Override query parameters since they can make URI exceed 2,000 character limit. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivationCodeAPI.V1ActivationCodeHistoryExportPost(context.Background()).ExportFields(exportFields).ExportLabels(exportLabels).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).ExportParameters(exportParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivationCodeAPI.V1ActivationCodeHistoryExportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ActivationCodeHistoryExportPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ActivationCodeAPI.V1ActivationCodeHistoryExportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ActivationCodeHistoryExportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportFields** | **[]string** | Export fields parameter, used to change default order or ignore some of the response properties. Default is empty array, which means that all fields of the response entity will be serialized. Example: export-fields&#x3D;id,username | [default to {}]
 **exportLabels** | **[]string** | Export labels parameter, used to customize fieldnames/columns in the exported file. Default is empty array, which means that response properties names will be used. Number of the provided labels must match the number of export-fields Example: export-labels&#x3D;identifier,name with matching: export-fields&#x3D;id,username | [default to {}]
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to {&quot;date:desc&quot;}]
 **filter** | **string** | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15 | [default to &quot;&quot;]
 **exportParameters** | [**ExportParameters**](ExportParameters.md) | Optional. Override query parameters since they can make URI exceed 2,000 character limit. | 

### Return type

**interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ActivationCodeHistoryGet

> HistorySearchResults V1ActivationCodeHistoryGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get Activation Code history object



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Fields allowed in the query: id, username, date, note, details Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,note:asc  (optional) (default to {"date:desc"})
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivationCodeAPI.V1ActivationCodeHistoryGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivationCodeAPI.V1ActivationCodeHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ActivationCodeHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `ActivationCodeAPI.V1ActivationCodeHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ActivationCodeHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Fields allowed in the query: id, username, date, note, details Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,note:asc  | [default to {&quot;date:desc&quot;}]
 **filter** | **string** | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, username, date, note, details. This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15 | [default to &quot;&quot;]

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


## V1ActivationCodeHistoryPost

> ObjectHistory V1ActivationCodeHistoryPost(ctx).ObjectHistoryNote(objectHistoryNote).Execute()

Add Activation Code object note 



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
	objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | Activation Code history notes to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivationCodeAPI.V1ActivationCodeHistoryPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivationCodeAPI.V1ActivationCodeHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ActivationCodeHistoryPost`: ObjectHistory
	fmt.Fprintf(os.Stdout, "Response from `ActivationCodeAPI.V1ActivationCodeHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ActivationCodeHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectHistoryNote** | [**ObjectHistoryNote**](ObjectHistoryNote.md) | Activation Code history notes to create. | 

### Return type

[**ObjectHistory**](ObjectHistory.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ActivationCodeOrganizationNamePatch

> V1ActivationCodeOrganizationNamePatch(ctx).OrganizationName(organizationName).Execute()

Updates Organization Name



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
	organizationName := *openapiclient.NewOrganizationName("Your Organization Name") // OrganizationName |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ActivationCodeAPI.V1ActivationCodeOrganizationNamePatch(context.Background()).OrganizationName(organizationName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivationCodeAPI.V1ActivationCodeOrganizationNamePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ActivationCodeOrganizationNamePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationName** | [**OrganizationName**](OrganizationName.md) |  | 

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


## V1ActivationCodePut

> V1ActivationCodePut(ctx).ActivationCode(activationCode).Execute()

Updates Activation Code



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
	activationCode := *openapiclient.NewActivationCode("A1A1-B2B2-C3C3-D4D4-E5E5-F6F6-G7G7-H8H8") // ActivationCode |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ActivationCodeAPI.V1ActivationCodePut(context.Background()).ActivationCode(activationCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivationCodeAPI.V1ActivationCodePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ActivationCodePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activationCode** | [**ActivationCode**](ActivationCode.md) |  | 

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

