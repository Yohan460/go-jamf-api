# \JamfRemoteAssistAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1JamfRemoteAssistSessionGet**](JamfRemoteAssistAPI.md#V1JamfRemoteAssistSessionGet) | **Get** /v1/jamf-remote-assist/session | Gets session history items. 
[**V1JamfRemoteAssistSessionIdGet**](JamfRemoteAssistAPI.md#V1JamfRemoteAssistSessionIdGet) | **Get** /v1/jamf-remote-assist/session/{id} | Gets single session history item. 
[**V2JamfRemoteAssistSessionExportPost**](JamfRemoteAssistAPI.md#V2JamfRemoteAssistSessionExportPost) | **Post** /v2/jamf-remote-assist/session/export | Export Jamf Remote Assist sessions history 
[**V2JamfRemoteAssistSessionGet**](JamfRemoteAssistAPI.md#V2JamfRemoteAssistSessionGet) | **Get** /v2/jamf-remote-assist/session | Gets session history items. 
[**V2JamfRemoteAssistSessionIdGet**](JamfRemoteAssistAPI.md#V2JamfRemoteAssistSessionIdGet) | **Get** /v2/jamf-remote-assist/session/{id} | Gets single session history item. 



## V1JamfRemoteAssistSessionGet

> []SessionHistoryItem V1JamfRemoteAssistSessionGet(ctx).Execute()

Gets session history items. 



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
	resp, r, err := apiClient.JamfRemoteAssistAPI.V1JamfRemoteAssistSessionGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfRemoteAssistAPI.V1JamfRemoteAssistSessionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1JamfRemoteAssistSessionGet`: []SessionHistoryItem
	fmt.Fprintf(os.Stdout, "Response from `JamfRemoteAssistAPI.V1JamfRemoteAssistSessionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfRemoteAssistSessionGetRequest struct via the builder pattern


### Return type

[**[]SessionHistoryItem**](SessionHistoryItem.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JamfRemoteAssistSessionIdGet

> SessionHistoryItemWithDetails V1JamfRemoteAssistSessionIdGet(ctx, id).Execute()

Gets single session history item. 



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
	id := "id_example" // string | instance id of session

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfRemoteAssistAPI.V1JamfRemoteAssistSessionIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfRemoteAssistAPI.V1JamfRemoteAssistSessionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1JamfRemoteAssistSessionIdGet`: SessionHistoryItemWithDetails
	fmt.Fprintf(os.Stdout, "Response from `JamfRemoteAssistAPI.V1JamfRemoteAssistSessionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of session | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfRemoteAssistSessionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SessionHistoryItemWithDetails**](SessionHistoryItemWithDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2JamfRemoteAssistSessionExportPost

> interface{} V2JamfRemoteAssistSessionExportPost(ctx).ExportParameters(exportParameters).Execute()

Export Jamf Remote Assist sessions history 



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
	exportParameters := *openapiclient.NewExportParameters() // ExportParameters | Optional. Override query parameters since they can make URI exceed 2,000 character limit. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfRemoteAssistAPI.V2JamfRemoteAssistSessionExportPost(context.Background()).ExportParameters(exportParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfRemoteAssistAPI.V2JamfRemoteAssistSessionExportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2JamfRemoteAssistSessionExportPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `JamfRemoteAssistAPI.V2JamfRemoteAssistSessionExportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2JamfRemoteAssistSessionExportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## V2JamfRemoteAssistSessionGet

> SessionHistorySearchResults V2JamfRemoteAssistSessionGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Gets session history items. 



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is sessionId:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=sessionId:desc,deviceId:asc  (optional) (default to ["sessionId:desc"])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter session history items collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: sessionId, deviceId, sessionAdminId. This param can be combined with paging and sorting. Example: sessionAdminId==\"*Andrzej*\" (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfRemoteAssistAPI.V2JamfRemoteAssistSessionGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfRemoteAssistAPI.V2JamfRemoteAssistSessionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2JamfRemoteAssistSessionGet`: SessionHistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `JamfRemoteAssistAPI.V2JamfRemoteAssistSessionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2JamfRemoteAssistSessionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is sessionId:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;sessionId:desc,deviceId:asc  | [default to [&quot;sessionId:desc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter session history items collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: sessionId, deviceId, sessionAdminId. This param can be combined with paging and sorting. Example: sessionAdminId&#x3D;&#x3D;\&quot;*Andrzej*\&quot; | [default to &quot;&quot;]

### Return type

[**SessionHistorySearchResults**](SessionHistorySearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2JamfRemoteAssistSessionIdGet

> SessionHistoryItemWithDetails V2JamfRemoteAssistSessionIdGet(ctx, id).Execute()

Gets single session history item. 



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
	id := "id_example" // string | instance id of session

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfRemoteAssistAPI.V2JamfRemoteAssistSessionIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfRemoteAssistAPI.V2JamfRemoteAssistSessionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2JamfRemoteAssistSessionIdGet`: SessionHistoryItemWithDetails
	fmt.Fprintf(os.Stdout, "Response from `JamfRemoteAssistAPI.V2JamfRemoteAssistSessionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of session | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2JamfRemoteAssistSessionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SessionHistoryItemWithDetails**](SessionHistoryItemWithDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

