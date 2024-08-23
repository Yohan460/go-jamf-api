# \OnboardingAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1OnboardingEligibleAppsGet**](OnboardingAPI.md#V1OnboardingEligibleAppsGet) | **Get** /v1/onboarding/eligible-apps | Retrieves a list of applications that are eligible to be used in an onboarding configuration
[**V1OnboardingEligibleConfigurationProfilesGet**](OnboardingAPI.md#V1OnboardingEligibleConfigurationProfilesGet) | **Get** /v1/onboarding/eligible-configuration-profiles | Retrieves a list of configuration profiles that are eligible to be used in an onboarding configuration
[**V1OnboardingEligiblePoliciesGet**](OnboardingAPI.md#V1OnboardingEligiblePoliciesGet) | **Get** /v1/onboarding/eligible-policies | Retrieves a list of policies that are eligible to be used in an onboarding configuration
[**V1OnboardingGet**](OnboardingAPI.md#V1OnboardingGet) | **Get** /v1/onboarding | Get the current onboarding settings configuration.
[**V1OnboardingHistoryExportPost**](OnboardingAPI.md#V1OnboardingHistoryExportPost) | **Post** /v1/onboarding/history/export | Export history object collection in specified format for Onboarding 
[**V1OnboardingHistoryGet**](OnboardingAPI.md#V1OnboardingHistoryGet) | **Get** /v1/onboarding/history | Get Onboarding history object 
[**V1OnboardingHistoryPost**](OnboardingAPI.md#V1OnboardingHistoryPost) | **Post** /v1/onboarding/history | Add Onboarding history object notes 
[**V1OnboardingPut**](OnboardingAPI.md#V1OnboardingPut) | **Put** /v1/onboarding | Update the onboarding configuration.



## V1OnboardingEligibleAppsGet

> OnboardingEligibleItemsSearchResult V1OnboardingEligibleAppsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Execute()

Retrieves a list of applications that are eligible to be used in an onboarding configuration



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["id:asc"])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnboardingAPI.V1OnboardingEligibleAppsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnboardingAPI.V1OnboardingEligibleAppsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OnboardingEligibleAppsGet`: OnboardingEligibleItemsSearchResult
	fmt.Fprintf(os.Stdout, "Response from `OnboardingAPI.V1OnboardingEligibleAppsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1OnboardingEligibleAppsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;id:asc&quot;]]

### Return type

[**OnboardingEligibleItemsSearchResult**](OnboardingEligibleItemsSearchResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OnboardingEligibleConfigurationProfilesGet

> OnboardingEligibleItemsSearchResult V1OnboardingEligibleConfigurationProfilesGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Execute()

Retrieves a list of configuration profiles that are eligible to be used in an onboarding configuration



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["id:asc"])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnboardingAPI.V1OnboardingEligibleConfigurationProfilesGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnboardingAPI.V1OnboardingEligibleConfigurationProfilesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OnboardingEligibleConfigurationProfilesGet`: OnboardingEligibleItemsSearchResult
	fmt.Fprintf(os.Stdout, "Response from `OnboardingAPI.V1OnboardingEligibleConfigurationProfilesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1OnboardingEligibleConfigurationProfilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;id:asc&quot;]]

### Return type

[**OnboardingEligibleItemsSearchResult**](OnboardingEligibleItemsSearchResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OnboardingEligiblePoliciesGet

> OnboardingEligibleItemsSearchResult V1OnboardingEligiblePoliciesGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Execute()

Retrieves a list of policies that are eligible to be used in an onboarding configuration



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["id:asc"])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnboardingAPI.V1OnboardingEligiblePoliciesGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnboardingAPI.V1OnboardingEligiblePoliciesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OnboardingEligiblePoliciesGet`: OnboardingEligibleItemsSearchResult
	fmt.Fprintf(os.Stdout, "Response from `OnboardingAPI.V1OnboardingEligiblePoliciesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1OnboardingEligiblePoliciesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;id:asc&quot;]]

### Return type

[**OnboardingEligibleItemsSearchResult**](OnboardingEligibleItemsSearchResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OnboardingGet

> OnboardingConfiguration V1OnboardingGet(ctx).Execute()

Get the current onboarding settings configuration.



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
	resp, r, err := apiClient.OnboardingAPI.V1OnboardingGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnboardingAPI.V1OnboardingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OnboardingGet`: OnboardingConfiguration
	fmt.Fprintf(os.Stdout, "Response from `OnboardingAPI.V1OnboardingGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1OnboardingGetRequest struct via the builder pattern


### Return type

[**OnboardingConfiguration**](OnboardingConfiguration.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OnboardingHistoryExportPost

> interface{} V1OnboardingHistoryExportPost(ctx).ExportFields(exportFields).ExportLabels(exportLabels).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).ExportParameters(exportParameters).Execute()

Export history object collection in specified format for Onboarding 



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
	exportFields := []string{"Inner_example"} // []string | Export fields parameter, used to change default order or ignore some of the response properties. Default is empty array, which means that all fields of the response entity will be serialized. Example: export-fields=id,username (optional) (default to [])
	exportLabels := []string{"Inner_example"} // []string | Export labels parameter, used to customize fieldnames/columns in the exported file. Default is empty array, which means that response properties names will be used. Number of the provided labels must match the number of export-fields Example: export-labels=identifier,name with matching: export-fields=id,username (optional) (default to [])
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["date:desc"])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and date<2019-12-15 (optional) (default to "")
	exportParameters := *openapiclient.NewExportParameters() // ExportParameters | Optional. Override query parameters since they can make URI exceed 2,000 character limit. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnboardingAPI.V1OnboardingHistoryExportPost(context.Background()).ExportFields(exportFields).ExportLabels(exportLabels).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).ExportParameters(exportParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnboardingAPI.V1OnboardingHistoryExportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OnboardingHistoryExportPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `OnboardingAPI.V1OnboardingHistoryExportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1OnboardingHistoryExportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportFields** | **[]string** | Export fields parameter, used to change default order or ignore some of the response properties. Default is empty array, which means that all fields of the response entity will be serialized. Example: export-fields&#x3D;id,username | [default to []]
 **exportLabels** | **[]string** | Export labels parameter, used to customize fieldnames/columns in the exported file. Default is empty array, which means that response properties names will be used. Number of the provided labels must match the number of export-fields Example: export-labels&#x3D;identifier,name with matching: export-fields&#x3D;id,username | [default to []]
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;date:desc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and date&lt;2019-12-15 | [default to &quot;&quot;]
 **exportParameters** | [**ExportParameters**](ExportParameters.md) | Optional. Override query parameters since they can make URI exceed 2,000 character limit. | 

### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OnboardingHistoryGet

> HistorySearchResults V1OnboardingHistoryGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get Onboarding history object 



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["date:desc"])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and date<2019-12-15 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnboardingAPI.V1OnboardingHistoryGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnboardingAPI.V1OnboardingHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OnboardingHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `OnboardingAPI.V1OnboardingHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1OnboardingHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;date:desc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and date&lt;2019-12-15 | [default to &quot;&quot;]

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


## V1OnboardingHistoryPost

> HrefResponse V1OnboardingHistoryPost(ctx).ObjectHistoryNote(objectHistoryNote).Execute()

Add Onboarding history object notes 



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
	resp, r, err := apiClient.OnboardingAPI.V1OnboardingHistoryPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnboardingAPI.V1OnboardingHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OnboardingHistoryPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `OnboardingAPI.V1OnboardingHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1OnboardingHistoryPostRequest struct via the builder pattern


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


## V1OnboardingPut

> OnboardingConfiguration V1OnboardingPut(ctx).OnboardingConfiguration(onboardingConfiguration).Execute()

Update the onboarding configuration.



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
	onboardingConfiguration := *openapiclient.NewOnboardingConfiguration(true, []openapiclient.OnboardingItem{*openapiclient.NewOnboardingItem("1", "OS_X_POLICY", int64(35))}) // OnboardingConfiguration | Onboarding settings to save.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnboardingAPI.V1OnboardingPut(context.Background()).OnboardingConfiguration(onboardingConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnboardingAPI.V1OnboardingPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1OnboardingPut`: OnboardingConfiguration
	fmt.Fprintf(os.Stdout, "Response from `OnboardingAPI.V1OnboardingPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1OnboardingPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onboardingConfiguration** | [**OnboardingConfiguration**](OnboardingConfiguration.md) | Onboarding settings to save. | 

### Return type

[**OnboardingConfiguration**](OnboardingConfiguration.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

