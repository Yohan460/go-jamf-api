# \DeviceCommunicationSettingsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DeviceCommunicationSettingsGet**](DeviceCommunicationSettingsAPI.md#V1DeviceCommunicationSettingsGet) | **Get** /v1/device-communication-settings | Retrieves all settings for device communication 
[**V1DeviceCommunicationSettingsHistoryGet**](DeviceCommunicationSettingsAPI.md#V1DeviceCommunicationSettingsHistoryGet) | **Get** /v1/device-communication-settings/history | Get Device Communication settings history 
[**V1DeviceCommunicationSettingsHistoryPost**](DeviceCommunicationSettingsAPI.md#V1DeviceCommunicationSettingsHistoryPost) | **Post** /v1/device-communication-settings/history | Add Device Communication Settings history notes 
[**V1DeviceCommunicationSettingsPut**](DeviceCommunicationSettingsAPI.md#V1DeviceCommunicationSettingsPut) | **Put** /v1/device-communication-settings | Update device communication settings 



## V1DeviceCommunicationSettingsGet

> DeviceCommunicationSettings V1DeviceCommunicationSettingsGet(ctx).Execute()

Retrieves all settings for device communication 



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
	resp, r, err := apiClient.DeviceCommunicationSettingsAPI.V1DeviceCommunicationSettingsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceCommunicationSettingsAPI.V1DeviceCommunicationSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DeviceCommunicationSettingsGet`: DeviceCommunicationSettings
	fmt.Fprintf(os.Stdout, "Response from `DeviceCommunicationSettingsAPI.V1DeviceCommunicationSettingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1DeviceCommunicationSettingsGetRequest struct via the builder pattern


### Return type

[**DeviceCommunicationSettings**](DeviceCommunicationSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeviceCommunicationSettingsHistoryGet

> HistorySearchResults V1DeviceCommunicationSettingsHistoryGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get Device Communication settings history 



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
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceCommunicationSettingsAPI.V1DeviceCommunicationSettingsHistoryGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceCommunicationSettingsAPI.V1DeviceCommunicationSettingsHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DeviceCommunicationSettingsHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `DeviceCommunicationSettingsAPI.V1DeviceCommunicationSettingsHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DeviceCommunicationSettingsHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;date:desc&quot;]]
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


## V1DeviceCommunicationSettingsHistoryPost

> ObjectHistory V1DeviceCommunicationSettingsHistoryPost(ctx).ObjectHistoryNote(objectHistoryNote).Execute()

Add Device Communication Settings history notes 



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
	resp, r, err := apiClient.DeviceCommunicationSettingsAPI.V1DeviceCommunicationSettingsHistoryPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceCommunicationSettingsAPI.V1DeviceCommunicationSettingsHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DeviceCommunicationSettingsHistoryPost`: ObjectHistory
	fmt.Fprintf(os.Stdout, "Response from `DeviceCommunicationSettingsAPI.V1DeviceCommunicationSettingsHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DeviceCommunicationSettingsHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectHistoryNote** | [**ObjectHistoryNote**](ObjectHistoryNote.md) | history notes to create | 

### Return type

[**ObjectHistory**](ObjectHistory.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeviceCommunicationSettingsPut

> DeviceCommunicationSettings V1DeviceCommunicationSettingsPut(ctx).DeviceCommunicationSettings(deviceCommunicationSettings).Execute()

Update device communication settings 



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
	deviceCommunicationSettings := *openapiclient.NewDeviceCommunicationSettings() // DeviceCommunicationSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceCommunicationSettingsAPI.V1DeviceCommunicationSettingsPut(context.Background()).DeviceCommunicationSettings(deviceCommunicationSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceCommunicationSettingsAPI.V1DeviceCommunicationSettingsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DeviceCommunicationSettingsPut`: DeviceCommunicationSettings
	fmt.Fprintf(os.Stdout, "Response from `DeviceCommunicationSettingsAPI.V1DeviceCommunicationSettingsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DeviceCommunicationSettingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceCommunicationSettings** | [**DeviceCommunicationSettings**](DeviceCommunicationSettings.md) |  | 

### Return type

[**DeviceCommunicationSettings**](DeviceCommunicationSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

