# \SelfServiceAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SelfServiceSettingsGet**](SelfServiceAPI.md#V1SelfServiceSettingsGet) | **Get** /v1/self-service/settings | Get an object representation of Self Service settings 
[**V1SelfServiceSettingsHistoryGet**](SelfServiceAPI.md#V1SelfServiceSettingsHistoryGet) | **Get** /v1/self-service/settings/history | Get a page of Self Service settings history 
[**V1SelfServiceSettingsHistoryPost**](SelfServiceAPI.md#V1SelfServiceSettingsHistoryPost) | **Post** /v1/self-service/settings/history | Add Self Service settings history notes 
[**V1SelfServiceSettingsPut**](SelfServiceAPI.md#V1SelfServiceSettingsPut) | **Put** /v1/self-service/settings | Put an object representation of Self Service settings 



## V1SelfServiceSettingsGet

> SelfServiceSettings V1SelfServiceSettingsGet(ctx).Execute()

Get an object representation of Self Service settings 



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
	resp, r, err := apiClient.SelfServiceAPI.V1SelfServiceSettingsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceAPI.V1SelfServiceSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SelfServiceSettingsGet`: SelfServiceSettings
	fmt.Fprintf(os.Stdout, "Response from `SelfServiceAPI.V1SelfServiceSettingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SelfServiceSettingsGetRequest struct via the builder pattern


### Return type

[**SelfServiceSettings**](SelfServiceSettings.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SelfServiceSettingsHistoryGet

> HistorySearchResults V1SelfServiceSettingsHistoryGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get a page of Self Service settings history 



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is not duplicated for each sort criterion, e.g., ...&sort=name:asc,date:desc. Fields that can be sorted: status, updated (optional) (default to {})
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SelfServiceAPI.V1SelfServiceSettingsHistoryGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceAPI.V1SelfServiceSettingsHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SelfServiceSettingsHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `SelfServiceAPI.V1SelfServiceSettingsHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SelfServiceSettingsHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the &#39;sort&#39; query param is not duplicated for each sort criterion, e.g., ...&amp;sort&#x3D;name:asc,date:desc. Fields that can be sorted: status, updated | [default to {}]
 **filter** | **string** | Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15 | [default to &quot;&quot;]

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


## V1SelfServiceSettingsHistoryPost

> HrefResponse V1SelfServiceSettingsHistoryPost(ctx).ObjectHistoryNote(objectHistoryNote).Execute()

Add Self Service settings history notes 



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
	resp, r, err := apiClient.SelfServiceAPI.V1SelfServiceSettingsHistoryPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceAPI.V1SelfServiceSettingsHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SelfServiceSettingsHistoryPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `SelfServiceAPI.V1SelfServiceSettingsHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SelfServiceSettingsHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectHistoryNote** | [**ObjectHistoryNote**](ObjectHistoryNote.md) | history notes to create | 

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


## V1SelfServiceSettingsPut

> SelfServiceSettings V1SelfServiceSettingsPut(ctx).SelfServiceSettings(selfServiceSettings).Execute()

Put an object representation of Self Service settings 



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
	selfServiceSettings := *openapiclient.NewSelfServiceSettings() // SelfServiceSettings | object that contains all editable global fields to alter Self Service settings 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SelfServiceAPI.V1SelfServiceSettingsPut(context.Background()).SelfServiceSettings(selfServiceSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceAPI.V1SelfServiceSettingsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SelfServiceSettingsPut`: SelfServiceSettings
	fmt.Fprintf(os.Stdout, "Response from `SelfServiceAPI.V1SelfServiceSettingsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SelfServiceSettingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selfServiceSettings** | [**SelfServiceSettings**](SelfServiceSettings.md) | object that contains all editable global fields to alter Self Service settings  | 

### Return type

[**SelfServiceSettings**](SelfServiceSettings.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

