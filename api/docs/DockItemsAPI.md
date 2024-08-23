# \DockItemsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DockItemsIdDelete**](DockItemsAPI.md#V1DockItemsIdDelete) | **Delete** /v1/dock-items/{id} | Delete a DockItem at the specified id 
[**V1DockItemsIdGet**](DockItemsAPI.md#V1DockItemsIdGet) | **Get** /v1/dock-items/{id} | Retrieve a full dockItem object 
[**V1DockItemsIdPut**](DockItemsAPI.md#V1DockItemsIdPut) | **Put** /v1/dock-items/{id} | Replace the dockItem at the id with the supplied information 
[**V1DockItemsPost**](DockItemsAPI.md#V1DockItemsPost) | **Post** /v1/dock-items | Create a DockItem 



## V1DockItemsIdDelete

> V1DockItemsIdDelete(ctx, id).Execute()

Delete a DockItem at the specified id 



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
	id := "id_example" // string | DockItem object identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DockItemsAPI.V1DockItemsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockItemsAPI.V1DockItemsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DockItem object identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DockItemsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DockItemsIdGet

> DockItem V1DockItemsIdGet(ctx, id).Execute()

Retrieve a full dockItem object 



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
	id := "id_example" // string | DockItem object identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockItemsAPI.V1DockItemsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockItemsAPI.V1DockItemsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DockItemsIdGet`: DockItem
	fmt.Fprintf(os.Stdout, "Response from `DockItemsAPI.V1DockItemsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DockItem object identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DockItemsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DockItem**](DockItem.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DockItemsIdPut

> string V1DockItemsIdPut(ctx, id).DockItem(dockItem).Execute()

Replace the dockItem at the id with the supplied information 



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
	id := "id_example" // string | DockItem object identifier
	dockItem := *openapiclient.NewDockItem("DockItem Name", "FILE", "file://localhost/Applications/iTunes.app") // DockItem | new dockItem to upload to existing id. ids defined in this body will be ignored

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockItemsAPI.V1DockItemsIdPut(context.Background(), id).DockItem(dockItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockItemsAPI.V1DockItemsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DockItemsIdPut`: string
	fmt.Fprintf(os.Stdout, "Response from `DockItemsAPI.V1DockItemsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DockItem object identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DockItemsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dockItem** | [**DockItem**](DockItem.md) | new dockItem to upload to existing id. ids defined in this body will be ignored | 

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DockItemsPost

> HrefResponse V1DockItemsPost(ctx).DockItem(dockItem).Execute()

Create a DockItem 



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
	dockItem := *openapiclient.NewDockItem("DockItem Name", "FILE", "file://localhost/Applications/iTunes.app") // DockItem | new DockItem to create. ids defined in this body will be ignored

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DockItemsAPI.V1DockItemsPost(context.Background()).DockItem(dockItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DockItemsAPI.V1DockItemsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DockItemsPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `DockItemsAPI.V1DockItemsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DockItemsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dockItem** | [**DockItem**](DockItem.md) | new DockItem to create. ids defined in this body will be ignored | 

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

