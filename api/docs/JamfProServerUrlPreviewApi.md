# \JamfProServerUrlPreviewAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1JamfProServerUrlGet**](JamfProServerUrlPreviewAPI.md#V1JamfProServerUrlGet) | **Get** /v1/jamf-pro-server-url | Get Jamf Pro Server URL settings 
[**V1JamfProServerUrlHistoryGet**](JamfProServerUrlPreviewAPI.md#V1JamfProServerUrlHistoryGet) | **Get** /v1/jamf-pro-server-url/history | Get Jamf Pro Server URL settings history 
[**V1JamfProServerUrlHistoryPost**](JamfProServerUrlPreviewAPI.md#V1JamfProServerUrlHistoryPost) | **Post** /v1/jamf-pro-server-url/history | Add Jamf Pro Server URL settings history notes 
[**V1JamfProServerUrlPut**](JamfProServerUrlPreviewAPI.md#V1JamfProServerUrlPut) | **Put** /v1/jamf-pro-server-url | Update Jamf Pro Server URL settings 



## V1JamfProServerUrlGet

> JamfProServerUrl V1JamfProServerUrlGet(ctx).Execute()

Get Jamf Pro Server URL settings 



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
	resp, r, err := apiClient.JamfProServerUrlPreviewAPI.V1JamfProServerUrlGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfProServerUrlPreviewAPI.V1JamfProServerUrlGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1JamfProServerUrlGet`: JamfProServerUrl
	fmt.Fprintf(os.Stdout, "Response from `JamfProServerUrlPreviewAPI.V1JamfProServerUrlGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfProServerUrlGetRequest struct via the builder pattern


### Return type

[**JamfProServerUrl**](JamfProServerUrl.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JamfProServerUrlHistoryGet

> HistorySearchResults V1JamfProServerUrlHistoryGet(ctx).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()

Get Jamf Pro Server URL settings history 



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
	size := int64(56) // int64 |  (optional) (default to 100)
	pagesize := int64(56) // int64 |  (optional) (default to 100)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := "sort_example" // string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to "date:desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfProServerUrlPreviewAPI.V1JamfProServerUrlHistoryGet(context.Background()).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfProServerUrlPreviewAPI.V1JamfProServerUrlHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1JamfProServerUrlHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `JamfProServerUrlPreviewAPI.V1JamfProServerUrlHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfProServerUrlHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **size** | **int64** |  | [default to 100]
 **pagesize** | **int64** |  | [default to 100]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **string** | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to &quot;date:desc&quot;]

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


## V1JamfProServerUrlHistoryPost

> ObjectHistory V1JamfProServerUrlHistoryPost(ctx).ObjectHistoryNote(objectHistoryNote).Execute()

Add Jamf Pro Server URL settings history notes 



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
	objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | History notes to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfProServerUrlPreviewAPI.V1JamfProServerUrlHistoryPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfProServerUrlPreviewAPI.V1JamfProServerUrlHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1JamfProServerUrlHistoryPost`: ObjectHistory
	fmt.Fprintf(os.Stdout, "Response from `JamfProServerUrlPreviewAPI.V1JamfProServerUrlHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfProServerUrlHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectHistoryNote** | [**ObjectHistoryNote**](ObjectHistoryNote.md) | History notes to create | 

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


## V1JamfProServerUrlPut

> JamfProServerUrl V1JamfProServerUrlPut(ctx).JamfProServerUrl(jamfProServerUrl).Execute()

Update Jamf Pro Server URL settings 



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
	jamfProServerUrl := *openapiclient.NewJamfProServerUrl("https://example.com:8443", "http://example.com") // JamfProServerUrl | Jamf Pro Server URL settings object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfProServerUrlPreviewAPI.V1JamfProServerUrlPut(context.Background()).JamfProServerUrl(jamfProServerUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfProServerUrlPreviewAPI.V1JamfProServerUrlPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1JamfProServerUrlPut`: JamfProServerUrl
	fmt.Fprintf(os.Stdout, "Response from `JamfProServerUrlPreviewAPI.V1JamfProServerUrlPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1JamfProServerUrlPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jamfProServerUrl** | [**JamfProServerUrl**](JamfProServerUrl.md) | Jamf Pro Server URL settings object | 

### Return type

[**JamfProServerUrl**](JamfProServerUrl.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

