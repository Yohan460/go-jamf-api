# \MobileDeviceExtensionAttributesAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1MobileDeviceExtensionAttributesGet**](MobileDeviceExtensionAttributesAPI.md#V1MobileDeviceExtensionAttributesGet) | **Get** /v1/mobile-device-extension-attributes | Retrieve Mobile Device Extension Attributes.
[**V1MobileDeviceExtensionAttributesIdDataDependencyGet**](MobileDeviceExtensionAttributesAPI.md#V1MobileDeviceExtensionAttributesIdDataDependencyGet) | **Get** /v1/mobile-device-extension-attributes/{id}/data-dependency | Get smart group dependent object for a specified mobile device extension attribute
[**V1MobileDeviceExtensionAttributesIdDelete**](MobileDeviceExtensionAttributesAPI.md#V1MobileDeviceExtensionAttributesIdDelete) | **Delete** /v1/mobile-device-extension-attributes/{id} | Delete a Mobile Device Extension Attribute by ID.
[**V1MobileDeviceExtensionAttributesIdGet**](MobileDeviceExtensionAttributesAPI.md#V1MobileDeviceExtensionAttributesIdGet) | **Get** /v1/mobile-device-extension-attributes/{id} | Get specified Mobile Device Extension Attribute object.
[**V1MobileDeviceExtensionAttributesIdHistoryGet**](MobileDeviceExtensionAttributesAPI.md#V1MobileDeviceExtensionAttributesIdHistoryGet) | **Get** /v1/mobile-device-extension-attributes/{id}/history | Get specified Mobile Device Extension Attribute History object
[**V1MobileDeviceExtensionAttributesIdHistoryPost**](MobileDeviceExtensionAttributesAPI.md#V1MobileDeviceExtensionAttributesIdHistoryPost) | **Post** /v1/mobile-device-extension-attributes/{id}/history | Add specified Mobile Device Extension Attribute history object notes
[**V1MobileDeviceExtensionAttributesIdPut**](MobileDeviceExtensionAttributesAPI.md#V1MobileDeviceExtensionAttributesIdPut) | **Put** /v1/mobile-device-extension-attributes/{id} | Update specified Mobile Device Extension Attribute object.
[**V1MobileDeviceExtensionAttributesPost**](MobileDeviceExtensionAttributesAPI.md#V1MobileDeviceExtensionAttributesPost) | **Post** /v1/mobile-device-extension-attributes | Create Mobile Device Extension Attribute.



## V1MobileDeviceExtensionAttributesGet

> MobileDeviceExtensionAttributeSearchResults V1MobileDeviceExtensionAttributesGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Retrieve Mobile Device Extension Attributes.



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
	sort := []string{"Inner_example"} // []string | Sorts results by one or more criteria, following the format property:asc/desc.<br/> Default sort is name:asc.<br/> If using multiple criteria, separate with commas. Allows sort for id and name. (optional) (default to {"name.asc"})
	filter := "filter_example" // string | Filters results. Use RSQL format for query. Allows for many fields, including ID, name, etc.<br/> Can be combined with paging and sorting.<br/> Fields allowed in the query: id, name <br/> Default filter is an empty query and returns all results from the requested page. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MobileDeviceExtensionAttributesGet`: MobileDeviceExtensionAttributeSearchResults
	fmt.Fprintf(os.Stdout, "Response from `MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDeviceExtensionAttributesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorts results by one or more criteria, following the format property:asc/desc.&lt;br/&gt; Default sort is name:asc.&lt;br/&gt; If using multiple criteria, separate with commas. Allows sort for id and name. | [default to {&quot;name.asc&quot;}]
 **filter** | **string** | Filters results. Use RSQL format for query. Allows for many fields, including ID, name, etc.&lt;br/&gt; Can be combined with paging and sorting.&lt;br/&gt; Fields allowed in the query: id, name &lt;br/&gt; Default filter is an empty query and returns all results from the requested page. | [default to &quot;&quot;]

### Return type

[**MobileDeviceExtensionAttributeSearchResults**](MobileDeviceExtensionAttributeSearchResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MobileDeviceExtensionAttributesIdDataDependencyGet

> DependencyObjectResults V1MobileDeviceExtensionAttributesIdDataDependencyGet(ctx, id).Execute()

Get smart group dependent object for a specified mobile device extension attribute



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
	id := "id_example" // string | Unique ID of mobile device extension attribute.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesIdDataDependencyGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesIdDataDependencyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MobileDeviceExtensionAttributesIdDataDependencyGet`: DependencyObjectResults
	fmt.Fprintf(os.Stdout, "Response from `MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesIdDataDependencyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique ID of mobile device extension attribute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDeviceExtensionAttributesIdDataDependencyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DependencyObjectResults**](DependencyObjectResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MobileDeviceExtensionAttributesIdDelete

> V1MobileDeviceExtensionAttributesIdDelete(ctx, id).Execute()

Delete a Mobile Device Extension Attribute by ID.



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
	id := "id_example" // string | Unique ID of Mobile Device Extension Attribute.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique ID of Mobile Device Extension Attribute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDeviceExtensionAttributesIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## V1MobileDeviceExtensionAttributesIdGet

> MobileDeviceExtensionAttributes V1MobileDeviceExtensionAttributesIdGet(ctx, id).Execute()

Get specified Mobile Device Extension Attribute object.



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
	id := "id_example" // string | Unique ID of Mobile Device Extension Attribute.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MobileDeviceExtensionAttributesIdGet`: MobileDeviceExtensionAttributes
	fmt.Fprintf(os.Stdout, "Response from `MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique ID of Mobile Device Extension Attribute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDeviceExtensionAttributesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MobileDeviceExtensionAttributes**](MobileDeviceExtensionAttributes.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MobileDeviceExtensionAttributesIdHistoryGet

> HistorySearchResults V1MobileDeviceExtensionAttributesIdHistoryGet(ctx, id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get specified Mobile Device Extension Attribute History object



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
	id := "id_example" // string | Instance ID of Mobile Device Extension Attribute
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorts results by one or more criteria, following the format property:asc/desc. Default sort is ID:asc. If using multiple criteria, separate with commas. (optional) (default to {"id:asc"})
	filter := "filter_example" // string | Filters results. Use RSQL format for query. Allows for many fields, including ID, name, etc. Can be combined with paging and sorting. Default filter is an empty query and returns all results from the requested page. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesIdHistoryGet(context.Background(), id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesIdHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MobileDeviceExtensionAttributesIdHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesIdHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Instance ID of Mobile Device Extension Attribute | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDeviceExtensionAttributesIdHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorts results by one or more criteria, following the format property:asc/desc. Default sort is ID:asc. If using multiple criteria, separate with commas. | [default to {&quot;id:asc&quot;}]
 **filter** | **string** | Filters results. Use RSQL format for query. Allows for many fields, including ID, name, etc. Can be combined with paging and sorting. Default filter is an empty query and returns all results from the requested page. | [default to &quot;&quot;]

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


## V1MobileDeviceExtensionAttributesIdHistoryPost

> ObjectHistory V1MobileDeviceExtensionAttributesIdHistoryPost(ctx, id).ObjectHistoryNote(objectHistoryNote).Execute()

Add specified Mobile Device Extension Attribute history object notes



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
	id := "id_example" // string | Instance ID of Mobile Device Extension Attribute
	objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | History note to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesIdHistoryPost(context.Background(), id).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesIdHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MobileDeviceExtensionAttributesIdHistoryPost`: ObjectHistory
	fmt.Fprintf(os.Stdout, "Response from `MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesIdHistoryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Instance ID of Mobile Device Extension Attribute | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDeviceExtensionAttributesIdHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectHistoryNote** | [**ObjectHistoryNote**](ObjectHistoryNote.md) | History note to be created | 

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


## V1MobileDeviceExtensionAttributesIdPut

> MobileDeviceExtensionAttributes V1MobileDeviceExtensionAttributesIdPut(ctx, id).MobileDeviceExtensionAttributes(mobileDeviceExtensionAttributes).Execute()

Update specified Mobile Device Extension Attribute object.



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
	id := "id_example" // string | Unique ID of Mobile Device Extension Attribute.
	mobileDeviceExtensionAttributes := *openapiclient.NewMobileDeviceExtensionAttributes("MobileDeviceExtensionAttribute", "DataType_example", "GENERAL", "InputType_example") // MobileDeviceExtensionAttributes | Mobile Device Extension Attribute object to be updated. IDs defined in this body will be ignored.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesIdPut(context.Background(), id).MobileDeviceExtensionAttributes(mobileDeviceExtensionAttributes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MobileDeviceExtensionAttributesIdPut`: MobileDeviceExtensionAttributes
	fmt.Fprintf(os.Stdout, "Response from `MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique ID of Mobile Device Extension Attribute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDeviceExtensionAttributesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mobileDeviceExtensionAttributes** | [**MobileDeviceExtensionAttributes**](MobileDeviceExtensionAttributes.md) | Mobile Device Extension Attribute object to be updated. IDs defined in this body will be ignored. | 

### Return type

[**MobileDeviceExtensionAttributes**](MobileDeviceExtensionAttributes.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MobileDeviceExtensionAttributesPost

> HrefResponse V1MobileDeviceExtensionAttributesPost(ctx).MobileDeviceExtensionAttributes(mobileDeviceExtensionAttributes).Execute()

Create Mobile Device Extension Attribute.



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
	mobileDeviceExtensionAttributes := *openapiclient.NewMobileDeviceExtensionAttributes("MobileDeviceExtensionAttribute", "DataType_example", "GENERAL", "InputType_example") // MobileDeviceExtensionAttributes | Mobile Device Extension Attribute to be created.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesPost(context.Background()).MobileDeviceExtensionAttributes(mobileDeviceExtensionAttributes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MobileDeviceExtensionAttributesPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `MobileDeviceExtensionAttributesAPI.V1MobileDeviceExtensionAttributesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDeviceExtensionAttributesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mobileDeviceExtensionAttributes** | [**MobileDeviceExtensionAttributes**](MobileDeviceExtensionAttributes.md) | Mobile Device Extension Attribute to be created. | 

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

