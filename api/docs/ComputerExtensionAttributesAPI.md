# \ComputerExtensionAttributesAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ComputerExtensionAttributesDeleteMultiplePost**](ComputerExtensionAttributesAPI.md#V1ComputerExtensionAttributesDeleteMultiplePost) | **Post** /v1/computer-extension-attributes/delete-multiple | Delete multiple Computer Extension Attribute at once.
[**V1ComputerExtensionAttributesGet**](ComputerExtensionAttributesAPI.md#V1ComputerExtensionAttributesGet) | **Get** /v1/computer-extension-attributes | Retrieve Computer Extension Attributes.
[**V1ComputerExtensionAttributesIdDataDependencyGet**](ComputerExtensionAttributesAPI.md#V1ComputerExtensionAttributesIdDataDependencyGet) | **Get** /v1/computer-extension-attributes/{id}/data-dependency | Get smart group/advance search dependent objects for a specified computer extension attribute
[**V1ComputerExtensionAttributesIdDelete**](ComputerExtensionAttributesAPI.md#V1ComputerExtensionAttributesIdDelete) | **Delete** /v1/computer-extension-attributes/{id} | Remove specified Computer Extension Attribute.
[**V1ComputerExtensionAttributesIdDownloadGet**](ComputerExtensionAttributesAPI.md#V1ComputerExtensionAttributesIdDownloadGet) | **Get** /v1/computer-extension-attributes/{id}/download | Download the specified Computer Extension Attribute.
[**V1ComputerExtensionAttributesIdGet**](ComputerExtensionAttributesAPI.md#V1ComputerExtensionAttributesIdGet) | **Get** /v1/computer-extension-attributes/{id} | Get specified Computer Extension Attribute object.
[**V1ComputerExtensionAttributesIdHistoryGet**](ComputerExtensionAttributesAPI.md#V1ComputerExtensionAttributesIdHistoryGet) | **Get** /v1/computer-extension-attributes/{id}/history | Get specified Computer Extension Attribute History object
[**V1ComputerExtensionAttributesIdHistoryPost**](ComputerExtensionAttributesAPI.md#V1ComputerExtensionAttributesIdHistoryPost) | **Post** /v1/computer-extension-attributes/{id}/history | Add specified Computer Extension Attribute history object notes
[**V1ComputerExtensionAttributesIdPut**](ComputerExtensionAttributesAPI.md#V1ComputerExtensionAttributesIdPut) | **Put** /v1/computer-extension-attributes/{id} | Update specified Computer Extension Attribute object.
[**V1ComputerExtensionAttributesPost**](ComputerExtensionAttributesAPI.md#V1ComputerExtensionAttributesPost) | **Post** /v1/computer-extension-attributes | Create Computer Extension Attribute.
[**V1ComputerExtensionAttributesTemplatesGet**](ComputerExtensionAttributesAPI.md#V1ComputerExtensionAttributesTemplatesGet) | **Get** /v1/computer-extension-attributes/templates | Retrieve All Computer Extension Attributes Templates.
[**V1ComputerExtensionAttributesTemplatesIdGet**](ComputerExtensionAttributesAPI.md#V1ComputerExtensionAttributesTemplatesIdGet) | **Get** /v1/computer-extension-attributes/templates/{id} | Get specified Computer Extension Attribute Template object.
[**V1ComputerExtensionAttributesUploadPost**](ComputerExtensionAttributesAPI.md#V1ComputerExtensionAttributesUploadPost) | **Post** /v1/computer-extension-attributes/upload | Upload Computer Extension Attribute.



## V1ComputerExtensionAttributesDeleteMultiplePost

> V1ComputerExtensionAttributesDeleteMultiplePost(ctx).Ids(ids).Execute()

Delete multiple Computer Extension Attribute at once.



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
	ids := *openapiclient.NewIds() // Ids | IDs of the Computer Extension Attribute to be deleted

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesDeleteMultiplePost(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesDeleteMultiplePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputerExtensionAttributesDeleteMultiplePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**Ids**](Ids.md) | IDs of the Computer Extension Attribute to be deleted | 

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


## V1ComputerExtensionAttributesGet

> ComputerExtensionAttributeSearchResults V1ComputerExtensionAttributesGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Retrieve Computer Extension Attributes.



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
	resp, r, err := apiClient.ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputerExtensionAttributesGet`: ComputerExtensionAttributeSearchResults
	fmt.Fprintf(os.Stdout, "Response from `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputerExtensionAttributesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorts results by one or more criteria, following the format property:asc/desc.&lt;br/&gt; Default sort is name:asc.&lt;br/&gt; If using multiple criteria, separate with commas. Allows sort for id and name. | [default to {&quot;name.asc&quot;}]
 **filter** | **string** | Filters results. Use RSQL format for query. Allows for many fields, including ID, name, etc.&lt;br/&gt; Can be combined with paging and sorting.&lt;br/&gt; Fields allowed in the query: id, name &lt;br/&gt; Default filter is an empty query and returns all results from the requested page. | [default to &quot;&quot;]

### Return type

[**ComputerExtensionAttributeSearchResults**](ComputerExtensionAttributeSearchResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ComputerExtensionAttributesIdDataDependencyGet

> DependencyObjectResults V1ComputerExtensionAttributesIdDataDependencyGet(ctx, id).Execute()

Get smart group/advance search dependent objects for a specified computer extension attribute



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
	id := "id_example" // string | Unique ID of computer extension attribute.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesIdDataDependencyGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesIdDataDependencyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputerExtensionAttributesIdDataDependencyGet`: DependencyObjectResults
	fmt.Fprintf(os.Stdout, "Response from `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesIdDataDependencyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique ID of computer extension attribute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputerExtensionAttributesIdDataDependencyGetRequest struct via the builder pattern


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


## V1ComputerExtensionAttributesIdDelete

> V1ComputerExtensionAttributesIdDelete(ctx, id).Execute()

Remove specified Computer Extension Attribute.



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
	id := "id_example" // string | Unique ID of Computer Extension Attribute.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique ID of Computer Extension Attribute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputerExtensionAttributesIdDeleteRequest struct via the builder pattern


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


## V1ComputerExtensionAttributesIdDownloadGet

> string V1ComputerExtensionAttributesIdDownloadGet(ctx, id).Execute()

Download the specified Computer Extension Attribute.



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
	id := "id_example" // string | The unique ID of the Computer Extension Attribute to be downloaded.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesIdDownloadGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesIdDownloadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputerExtensionAttributesIdDownloadGet`: string
	fmt.Fprintf(os.Stdout, "Response from `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesIdDownloadGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique ID of the Computer Extension Attribute to be downloaded. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputerExtensionAttributesIdDownloadGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ComputerExtensionAttributesIdGet

> ComputerExtensionAttributes V1ComputerExtensionAttributesIdGet(ctx, id).Execute()

Get specified Computer Extension Attribute object.



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
	id := "id_example" // string | Unique ID of Computer Extension Attribute.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputerExtensionAttributesIdGet`: ComputerExtensionAttributes
	fmt.Fprintf(os.Stdout, "Response from `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique ID of Computer Extension Attribute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputerExtensionAttributesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComputerExtensionAttributes**](ComputerExtensionAttributes.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ComputerExtensionAttributesIdHistoryGet

> HistorySearchResults V1ComputerExtensionAttributesIdHistoryGet(ctx, id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get specified Computer Extension Attribute History object



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
	id := "id_example" // string | Instance ID of Computer Extension Attribute history
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorts results by one or more criteria, following the format property:asc/desc. Default sort is ID:asc. If using multiple criteria, separate with commas. (optional) (default to {"id:asc"})
	filter := "filter_example" // string | Filters results. Use RSQL format for query. Allows for many fields, including ID, name, etc. Can be combined with paging and sorting. Default filter is an empty query and returns all results from the requested page. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesIdHistoryGet(context.Background(), id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesIdHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputerExtensionAttributesIdHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesIdHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Instance ID of Computer Extension Attribute history | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputerExtensionAttributesIdHistoryGetRequest struct via the builder pattern


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


## V1ComputerExtensionAttributesIdHistoryPost

> ObjectHistory V1ComputerExtensionAttributesIdHistoryPost(ctx, id).ObjectHistoryNote(objectHistoryNote).Execute()

Add specified Computer Extension Attribute history object notes



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
	id := "id_example" // string | Instance ID of Computer Extension Attribute history
	objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | History note to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesIdHistoryPost(context.Background(), id).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesIdHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputerExtensionAttributesIdHistoryPost`: ObjectHistory
	fmt.Fprintf(os.Stdout, "Response from `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesIdHistoryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Instance ID of Computer Extension Attribute history | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputerExtensionAttributesIdHistoryPostRequest struct via the builder pattern


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


## V1ComputerExtensionAttributesIdPut

> ComputerExtensionAttributes V1ComputerExtensionAttributesIdPut(ctx, id).ComputerExtensionAttributes(computerExtensionAttributes).Execute()

Update specified Computer Extension Attribute object.



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
	id := "id_example" // string | Unique ID of Computer Extension Attribute.
	computerExtensionAttributes := *openapiclient.NewComputerExtensionAttributes("MobileDeviceExtensionAttribute", "DataType_example", "GENERAL", "InputType_example") // ComputerExtensionAttributes | Computer Extension Attribute object to be updated. IDs defined in this body will be ignored.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesIdPut(context.Background(), id).ComputerExtensionAttributes(computerExtensionAttributes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputerExtensionAttributesIdPut`: ComputerExtensionAttributes
	fmt.Fprintf(os.Stdout, "Response from `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique ID of Computer Extension Attribute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputerExtensionAttributesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **computerExtensionAttributes** | [**ComputerExtensionAttributes**](ComputerExtensionAttributes.md) | Computer Extension Attribute object to be updated. IDs defined in this body will be ignored. | 

### Return type

[**ComputerExtensionAttributes**](ComputerExtensionAttributes.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ComputerExtensionAttributesPost

> HrefResponse V1ComputerExtensionAttributesPost(ctx).ComputerExtensionAttributes(computerExtensionAttributes).Execute()

Create Computer Extension Attribute.



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
	computerExtensionAttributes := *openapiclient.NewComputerExtensionAttributes("MobileDeviceExtensionAttribute", "DataType_example", "GENERAL", "InputType_example") // ComputerExtensionAttributes | Computer Extension Attribute to be created.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesPost(context.Background()).ComputerExtensionAttributes(computerExtensionAttributes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputerExtensionAttributesPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputerExtensionAttributesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **computerExtensionAttributes** | [**ComputerExtensionAttributes**](ComputerExtensionAttributes.md) | Computer Extension Attribute to be created. | 

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


## V1ComputerExtensionAttributesTemplatesGet

> CeaTemplatesResults V1ComputerExtensionAttributesTemplatesGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Retrieve All Computer Extension Attributes Templates.



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
	sort := []string{"Inner_example"} // []string | Sorts results by one or more criteria, following the format property:asc/desc.<br/> Default sort is templateName:asc.<br/> If using multiple criteria, separate with commas. Allows sort for templateName and templateCategory. (optional) (default to {"templateName.asc"})
	filter := "filter_example" // string | Filters results. Use RSQL format for queries. which allows filtering by multiple fields such as templateName, templateCategoryName.<br/> Can be combined with paging and sorting.<br/> Fields allowed in the query: templateName, templateCategoryName <br/> Default filter is an empty query and returns all results from the requested page. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesTemplatesGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesTemplatesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputerExtensionAttributesTemplatesGet`: CeaTemplatesResults
	fmt.Fprintf(os.Stdout, "Response from `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesTemplatesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputerExtensionAttributesTemplatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorts results by one or more criteria, following the format property:asc/desc.&lt;br/&gt; Default sort is templateName:asc.&lt;br/&gt; If using multiple criteria, separate with commas. Allows sort for templateName and templateCategory. | [default to {&quot;templateName.asc&quot;}]
 **filter** | **string** | Filters results. Use RSQL format for queries. which allows filtering by multiple fields such as templateName, templateCategoryName.&lt;br/&gt; Can be combined with paging and sorting.&lt;br/&gt; Fields allowed in the query: templateName, templateCategoryName &lt;br/&gt; Default filter is an empty query and returns all results from the requested page. | [default to &quot;&quot;]

### Return type

[**CeaTemplatesResults**](CeaTemplatesResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ComputerExtensionAttributesTemplatesIdGet

> ComputerExtensionAttributes V1ComputerExtensionAttributesTemplatesIdGet(ctx, id).Execute()

Get specified Computer Extension Attribute Template object.



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
	id := "id_example" // string | Unique Id of the Template.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesTemplatesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesTemplatesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputerExtensionAttributesTemplatesIdGet`: ComputerExtensionAttributes
	fmt.Fprintf(os.Stdout, "Response from `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesTemplatesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique Id of the Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputerExtensionAttributesTemplatesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComputerExtensionAttributes**](ComputerExtensionAttributes.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ComputerExtensionAttributesUploadPost

> ComputerExtensionAttributes V1ComputerExtensionAttributesUploadPost(ctx).File(file).Execute()

Upload Computer Extension Attribute.



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
	file := os.NewFile(1234, "some_file") // *os.File | The file to upload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesUploadPost(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesUploadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputerExtensionAttributesUploadPost`: ComputerExtensionAttributes
	fmt.Fprintf(os.Stdout, "Response from `ComputerExtensionAttributesAPI.V1ComputerExtensionAttributesUploadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputerExtensionAttributesUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The file to upload | 

### Return type

[**ComputerExtensionAttributes**](ComputerExtensionAttributes.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

