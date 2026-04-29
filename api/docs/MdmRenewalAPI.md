# \MdmRenewalAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1MdmRenewalDeviceCommonDetailsClientManagementIdGet**](MdmRenewalAPI.md#V1MdmRenewalDeviceCommonDetailsClientManagementIdGet) | **Get** /v1/mdm-renewal/device-common-details/{clientManagementId} | Get device common details for a client management ID 
[**V1MdmRenewalDeviceCommonDetailsPatch**](MdmRenewalAPI.md#V1MdmRenewalDeviceCommonDetailsPatch) | **Patch** /v1/mdm-renewal/device-common-details | Update device common details (partial update) 
[**V1MdmRenewalRenewalStrategiesClientManagementIdDelete**](MdmRenewalAPI.md#V1MdmRenewalRenewalStrategiesClientManagementIdDelete) | **Delete** /v1/mdm-renewal/renewal-strategies/{clientManagementId} | Delete MDM renewal strategies for a client management ID 
[**V1MdmRenewalRenewalStrategiesClientManagementIdGet**](MdmRenewalAPI.md#V1MdmRenewalRenewalStrategiesClientManagementIdGet) | **Get** /v1/mdm-renewal/renewal-strategies/{clientManagementId} | Get MDM renewal errors and strategies for a client management ID 



## V1MdmRenewalDeviceCommonDetailsClientManagementIdGet

> DeviceCommonDetails V1MdmRenewalDeviceCommonDetailsClientManagementIdGet(ctx, clientManagementId).Execute()

Get device common details for a client management ID 



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
	clientManagementId := "550e8400-e29b-41d4-a716-446655440000" // string | The client management ID to retrieve device common details for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MdmRenewalAPI.V1MdmRenewalDeviceCommonDetailsClientManagementIdGet(context.Background(), clientManagementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MdmRenewalAPI.V1MdmRenewalDeviceCommonDetailsClientManagementIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MdmRenewalDeviceCommonDetailsClientManagementIdGet`: DeviceCommonDetails
	fmt.Fprintf(os.Stdout, "Response from `MdmRenewalAPI.V1MdmRenewalDeviceCommonDetailsClientManagementIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientManagementId** | **string** | The client management ID to retrieve device common details for | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MdmRenewalDeviceCommonDetailsClientManagementIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceCommonDetails**](DeviceCommonDetails.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MdmRenewalDeviceCommonDetailsPatch

> V1MdmRenewalDeviceCommonDetailsPatch(ctx).DeviceCommonDetailsRequest(deviceCommonDetailsRequest).Execute()

Update device common details (partial update) 



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
	deviceCommonDetailsRequest := *openapiclient.NewDeviceCommonDetailsRequest("550e8400-e29b-41d4-a716-446655440000") // DeviceCommonDetailsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MdmRenewalAPI.V1MdmRenewalDeviceCommonDetailsPatch(context.Background()).DeviceCommonDetailsRequest(deviceCommonDetailsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MdmRenewalAPI.V1MdmRenewalDeviceCommonDetailsPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MdmRenewalDeviceCommonDetailsPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceCommonDetailsRequest** | [**DeviceCommonDetailsRequest**](DeviceCommonDetailsRequest.md) |  | 

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


## V1MdmRenewalRenewalStrategiesClientManagementIdDelete

> V1MdmRenewalRenewalStrategiesClientManagementIdDelete(ctx, clientManagementId).Execute()

Delete MDM renewal strategies for a client management ID 



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
	clientManagementId := "550e8400-e29b-41d4-a716-446655440000" // string | The client management ID to delete renewal strategies for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MdmRenewalAPI.V1MdmRenewalRenewalStrategiesClientManagementIdDelete(context.Background(), clientManagementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MdmRenewalAPI.V1MdmRenewalRenewalStrategiesClientManagementIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientManagementId** | **string** | The client management ID to delete renewal strategies for | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MdmRenewalRenewalStrategiesClientManagementIdDeleteRequest struct via the builder pattern


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


## V1MdmRenewalRenewalStrategiesClientManagementIdGet

> []MdmRenewalErrorStrategiesResponse V1MdmRenewalRenewalStrategiesClientManagementIdGet(ctx, clientManagementId).Execute()

Get MDM renewal errors and strategies for a client management ID 



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
	clientManagementId := "550e8400-e29b-41d4-a716-446655440000" // string | The client management ID to retrieve renewal strategies for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MdmRenewalAPI.V1MdmRenewalRenewalStrategiesClientManagementIdGet(context.Background(), clientManagementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MdmRenewalAPI.V1MdmRenewalRenewalStrategiesClientManagementIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MdmRenewalRenewalStrategiesClientManagementIdGet`: []MdmRenewalErrorStrategiesResponse
	fmt.Fprintf(os.Stdout, "Response from `MdmRenewalAPI.V1MdmRenewalRenewalStrategiesClientManagementIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientManagementId** | **string** | The client management ID to retrieve renewal strategies for | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MdmRenewalRenewalStrategiesClientManagementIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MdmRenewalErrorStrategiesResponse**](MdmRenewalErrorStrategiesResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

