# \ConditionalAccessAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ConditionalAccessDeviceComplianceFeatureToggleGet**](ConditionalAccessAPI.md#V1ConditionalAccessDeviceComplianceFeatureToggleGet) | **Get** /v1/conditional-access/device-compliance/feature-toggle | Retrieves Status of the Feature Toggle
[**V1ConditionalAccessDeviceComplianceInformationComputerDeviceIdGet**](ConditionalAccessAPI.md#V1ConditionalAccessDeviceComplianceInformationComputerDeviceIdGet) | **Get** /v1/conditional-access/device-compliance-information/computer/{deviceId} | Get compliance information for a single computer device
[**V1ConditionalAccessDeviceComplianceInformationMobileDeviceIdGet**](ConditionalAccessAPI.md#V1ConditionalAccessDeviceComplianceInformationMobileDeviceIdGet) | **Get** /v1/conditional-access/device-compliance-information/mobile/{deviceId} | Get compliance information for a single mobile device



## V1ConditionalAccessDeviceComplianceFeatureToggleGet

> SharedDeviceComplianceFeatureToggle V1ConditionalAccessDeviceComplianceFeatureToggleGet(ctx).Execute()

Retrieves Status of the Feature Toggle



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
	resp, r, err := apiClient.ConditionalAccessAPI.V1ConditionalAccessDeviceComplianceFeatureToggleGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConditionalAccessAPI.V1ConditionalAccessDeviceComplianceFeatureToggleGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ConditionalAccessDeviceComplianceFeatureToggleGet`: SharedDeviceComplianceFeatureToggle
	fmt.Fprintf(os.Stdout, "Response from `ConditionalAccessAPI.V1ConditionalAccessDeviceComplianceFeatureToggleGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConditionalAccessDeviceComplianceFeatureToggleGetRequest struct via the builder pattern


### Return type

[**SharedDeviceComplianceFeatureToggle**](SharedDeviceComplianceFeatureToggle.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ConditionalAccessDeviceComplianceInformationComputerDeviceIdGet

> []DeviceComplianceInformation V1ConditionalAccessDeviceComplianceInformationComputerDeviceIdGet(ctx, deviceId).Execute()

Get compliance information for a single computer device



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
	deviceId := "deviceId_example" // string | ID of the device the query pertains

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConditionalAccessAPI.V1ConditionalAccessDeviceComplianceInformationComputerDeviceIdGet(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConditionalAccessAPI.V1ConditionalAccessDeviceComplianceInformationComputerDeviceIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ConditionalAccessDeviceComplianceInformationComputerDeviceIdGet`: []DeviceComplianceInformation
	fmt.Fprintf(os.Stdout, "Response from `ConditionalAccessAPI.V1ConditionalAccessDeviceComplianceInformationComputerDeviceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | ID of the device the query pertains | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConditionalAccessDeviceComplianceInformationComputerDeviceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DeviceComplianceInformation**](DeviceComplianceInformation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ConditionalAccessDeviceComplianceInformationMobileDeviceIdGet

> []DeviceComplianceInformation V1ConditionalAccessDeviceComplianceInformationMobileDeviceIdGet(ctx, deviceId).Execute()

Get compliance information for a single mobile device



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
	deviceId := "deviceId_example" // string | ID of the device the query pertains

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConditionalAccessAPI.V1ConditionalAccessDeviceComplianceInformationMobileDeviceIdGet(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConditionalAccessAPI.V1ConditionalAccessDeviceComplianceInformationMobileDeviceIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ConditionalAccessDeviceComplianceInformationMobileDeviceIdGet`: []DeviceComplianceInformation
	fmt.Fprintf(os.Stdout, "Response from `ConditionalAccessAPI.V1ConditionalAccessDeviceComplianceInformationMobileDeviceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | ID of the device the query pertains | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ConditionalAccessDeviceComplianceInformationMobileDeviceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DeviceComplianceInformation**](DeviceComplianceInformation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

