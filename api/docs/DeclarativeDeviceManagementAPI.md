# \DeclarativeDeviceManagementAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DdmClientManagementIdStatusItemsGet**](DeclarativeDeviceManagementAPI.md#V1DdmClientManagementIdStatusItemsGet) | **Get** /v1/ddm/{clientManagementId}/status-items | Retrieve the Status Items from the latest Status Report for a device
[**V1DdmClientManagementIdStatusItemsKeyGet**](DeclarativeDeviceManagementAPI.md#V1DdmClientManagementIdStatusItemsKeyGet) | **Get** /v1/ddm/{clientManagementId}/status-items/{key} | Retrieve a Status Item from the latest Status Report for a device
[**V1DdmClientManagementIdSyncPost**](DeclarativeDeviceManagementAPI.md#V1DdmClientManagementIdSyncPost) | **Post** /v1/ddm/{clientManagementId}/sync | Force a device DDM sync
[**V1DssDeclarationsDeclarationIdGet**](DeclarativeDeviceManagementAPI.md#V1DssDeclarationsDeclarationIdGet) | **Get** /v1/dss-declarations/{declarationId} | Retrieve an existing declaration



## V1DdmClientManagementIdStatusItemsGet

> StatusItems V1DdmClientManagementIdStatusItemsGet(ctx, clientManagementId).Execute()

Retrieve the Status Items from the latest Status Report for a device



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
	clientManagementId := "clientManagementId_example" // string | client management id of the target device.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeclarativeDeviceManagementAPI.V1DdmClientManagementIdStatusItemsGet(context.Background(), clientManagementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeclarativeDeviceManagementAPI.V1DdmClientManagementIdStatusItemsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DdmClientManagementIdStatusItemsGet`: StatusItems
	fmt.Fprintf(os.Stdout, "Response from `DeclarativeDeviceManagementAPI.V1DdmClientManagementIdStatusItemsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientManagementId** | **string** | client management id of the target device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DdmClientManagementIdStatusItemsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StatusItems**](StatusItems.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DdmClientManagementIdStatusItemsKeyGet

> StatusItem V1DdmClientManagementIdStatusItemsKeyGet(ctx, clientManagementId, key).Execute()

Retrieve a Status Item from the latest Status Report for a device



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
	clientManagementId := "clientManagementId_example" // string | client management id of the target device.
	key := "key_example" // string | the status item key to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeclarativeDeviceManagementAPI.V1DdmClientManagementIdStatusItemsKeyGet(context.Background(), clientManagementId, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeclarativeDeviceManagementAPI.V1DdmClientManagementIdStatusItemsKeyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DdmClientManagementIdStatusItemsKeyGet`: StatusItem
	fmt.Fprintf(os.Stdout, "Response from `DeclarativeDeviceManagementAPI.V1DdmClientManagementIdStatusItemsKeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientManagementId** | **string** | client management id of the target device. | 
**key** | **string** | the status item key to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DdmClientManagementIdStatusItemsKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StatusItem**](StatusItem.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DdmClientManagementIdSyncPost

> V1DdmClientManagementIdSyncPost(ctx, clientManagementId).Execute()

Force a device DDM sync



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
	clientManagementId := "clientManagementId_example" // string | The client management id of the target device.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeclarativeDeviceManagementAPI.V1DdmClientManagementIdSyncPost(context.Background(), clientManagementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeclarativeDeviceManagementAPI.V1DdmClientManagementIdSyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientManagementId** | **string** | The client management id of the target device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DdmClientManagementIdSyncPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DssDeclarationsDeclarationIdGet

> DssDeclarations V1DssDeclarationsDeclarationIdGet(ctx, declarationId).Execute()

Retrieve an existing declaration



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
	declarationId := "538F90D7-9383-4A2B-B1C8-81E845A9CFD7" // string | Declaration UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeclarativeDeviceManagementAPI.V1DssDeclarationsDeclarationIdGet(context.Background(), declarationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeclarativeDeviceManagementAPI.V1DssDeclarationsDeclarationIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DssDeclarationsDeclarationIdGet`: DssDeclarations
	fmt.Fprintf(os.Stdout, "Response from `DeclarativeDeviceManagementAPI.V1DssDeclarationsDeclarationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**declarationId** | **string** | Declaration UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DssDeclarationsDeclarationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DssDeclarations**](DssDeclarations.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

