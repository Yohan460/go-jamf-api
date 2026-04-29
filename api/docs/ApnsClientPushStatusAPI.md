# \ApnsClientPushStatusAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ApnsClientPushStatusEnableAllClientsPost**](ApnsClientPushStatusAPI.md#V1ApnsClientPushStatusEnableAllClientsPost) | **Post** /v1/apns-client-push-status/enable-all-clients | Enable push notifications for all clients 
[**V1ApnsClientPushStatusEnableAllClientsStatusGet**](ApnsClientPushStatusAPI.md#V1ApnsClientPushStatusEnableAllClientsStatusGet) | **Get** /v1/apns-client-push-status/enable-all-clients/status | Get status of enable all clients request 
[**V1ApnsClientPushStatusEnableClientPost**](ApnsClientPushStatusAPI.md#V1ApnsClientPushStatusEnableClientPost) | **Post** /v1/apns-client-push-status/enable-client | Enable push notifications for a single client 
[**V1ApnsClientPushStatusGet**](ApnsClientPushStatusAPI.md#V1ApnsClientPushStatusGet) | **Get** /v1/apns-client-push-status | Search for clients with push notifications disabled 



## V1ApnsClientPushStatusEnableAllClientsPost

> V1ApnsClientPushStatusEnableAllClientsPost(ctx).Execute()

Enable push notifications for all clients 



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
	r, err := apiClient.ApnsClientPushStatusAPI.V1ApnsClientPushStatusEnableAllClientsPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApnsClientPushStatusAPI.V1ApnsClientPushStatusEnableAllClientsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ApnsClientPushStatusEnableAllClientsPostRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ApnsClientPushStatusEnableAllClientsStatusGet

> ApnsPushEnableRequest V1ApnsClientPushStatusEnableAllClientsStatusGet(ctx).Execute()

Get status of enable all clients request 



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
	resp, r, err := apiClient.ApnsClientPushStatusAPI.V1ApnsClientPushStatusEnableAllClientsStatusGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApnsClientPushStatusAPI.V1ApnsClientPushStatusEnableAllClientsStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ApnsClientPushStatusEnableAllClientsStatusGet`: ApnsPushEnableRequest
	fmt.Fprintf(os.Stdout, "Response from `ApnsClientPushStatusAPI.V1ApnsClientPushStatusEnableAllClientsStatusGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ApnsClientPushStatusEnableAllClientsStatusGetRequest struct via the builder pattern


### Return type

[**ApnsPushEnableRequest**](ApnsPushEnableRequest.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ApnsClientPushStatusEnableClientPost

> V1ApnsClientPushStatusEnableClientPost(ctx).EnablePushRequest(enablePushRequest).Execute()

Enable push notifications for a single client 



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
	enablePushRequest := *openapiclient.NewEnablePushRequest("a1b2c3d4-e5f6-7890-abcd-ef1234567890") // EnablePushRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApnsClientPushStatusAPI.V1ApnsClientPushStatusEnableClientPost(context.Background()).EnablePushRequest(enablePushRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApnsClientPushStatusAPI.V1ApnsClientPushStatusEnableClientPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ApnsClientPushStatusEnableClientPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enablePushRequest** | [**EnablePushRequest**](EnablePushRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ApnsClientPushStatusGet

> ApnsClientPushStatusSearchResults V1ApnsClientPushStatusGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Search for clients with push notifications disabled 



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. Sortable fields: pushDisabledTime, deviceType, managementId  (optional) (default to {"pushDisabledTime:asc"})
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter results. Fields allowed in the query: deviceType, disabledAt, managementId. This param can be combined with paging and sorting.  Example: filter=deviceType==\"MOBILE_DEVICE\" Example: filter=disabledAt>2024-11-01T00:00:00Z Example: filter=deviceType==\"COMPUTER\";disabledAt>2024-01-01T00:00:00Z  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApnsClientPushStatusAPI.V1ApnsClientPushStatusGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApnsClientPushStatusAPI.V1ApnsClientPushStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ApnsClientPushStatusGet`: ApnsClientPushStatusSearchResults
	fmt.Fprintf(os.Stdout, "Response from `ApnsClientPushStatusAPI.V1ApnsClientPushStatusGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ApnsClientPushStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. Sortable fields: pushDisabledTime, deviceType, managementId  | [default to {&quot;pushDisabledTime:asc&quot;}]
 **filter** | **string** | Query in the RSQL format, allowing to filter results. Fields allowed in the query: deviceType, disabledAt, managementId. This param can be combined with paging and sorting.  Example: filter&#x3D;deviceType&#x3D;&#x3D;\&quot;MOBILE_DEVICE\&quot; Example: filter&#x3D;disabledAt&gt;2024-11-01T00:00:00Z Example: filter&#x3D;deviceType&#x3D;&#x3D;\&quot;COMPUTER\&quot;;disabledAt&gt;2024-01-01T00:00:00Z  | [default to &quot;&quot;]

### Return type

[**ApnsClientPushStatusSearchResults**](ApnsClientPushStatusSearchResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

