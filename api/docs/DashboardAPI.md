# \DashboardAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DashboardGet**](DashboardAPI.md#V1DashboardGet) | **Get** /v1/dashboard | Get all the dashboard setup information 
[**V1DashboardTogglePost**](DashboardAPI.md#V1DashboardTogglePost) | **Post** /v1/dashboard/toggle | Add or remove an object to the Jamf Pro dashboard 



## V1DashboardGet

> DashboardSetup V1DashboardGet(ctx).Execute()

Get all the dashboard setup information 



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
	resp, r, err := apiClient.DashboardAPI.V1DashboardGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.V1DashboardGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DashboardGet`: DashboardSetup
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.V1DashboardGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1DashboardGetRequest struct via the builder pattern


### Return type

[**DashboardSetup**](DashboardSetup.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DashboardTogglePost

> HrefResponse V1DashboardTogglePost(ctx).DashboardObject(dashboardObject).Execute()

Add or remove an object to the Jamf Pro dashboard 



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
	dashboardObject := *openapiclient.NewDashboardObject("1", "TYPE_IOS_CONFIGURATION_PROFILE", true) // DashboardObject | Dashboard object with the associated type and ID with a toggle to add or remove the object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.V1DashboardTogglePost(context.Background()).DashboardObject(dashboardObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.V1DashboardTogglePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DashboardTogglePost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.V1DashboardTogglePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DashboardTogglePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboardObject** | [**DashboardObject**](DashboardObject.md) | Dashboard object with the associated type and ID with a toggle to add or remove the object | 

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

