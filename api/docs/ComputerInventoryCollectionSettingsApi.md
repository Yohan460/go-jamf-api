# \ComputerInventoryCollectionSettingsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ComputerInventoryCollectionSettingsCustomPathIdDelete**](ComputerInventoryCollectionSettingsAPI.md#V1ComputerInventoryCollectionSettingsCustomPathIdDelete) | **Delete** /v1/computer-inventory-collection-settings/custom-path/{id} | Delete Custom Path from Computer Inventory Collection Settings
[**V1ComputerInventoryCollectionSettingsCustomPathPost**](ComputerInventoryCollectionSettingsAPI.md#V1ComputerInventoryCollectionSettingsCustomPathPost) | **Post** /v1/computer-inventory-collection-settings/custom-path | Create Computer Inventory Collection Settings Custom Path
[**V1ComputerInventoryCollectionSettingsGet**](ComputerInventoryCollectionSettingsAPI.md#V1ComputerInventoryCollectionSettingsGet) | **Get** /v1/computer-inventory-collection-settings | Returns computer inventory settings
[**V1ComputerInventoryCollectionSettingsPatch**](ComputerInventoryCollectionSettingsAPI.md#V1ComputerInventoryCollectionSettingsPatch) | **Patch** /v1/computer-inventory-collection-settings | Update computer inventory settings



## V1ComputerInventoryCollectionSettingsCustomPathIdDelete

> V1ComputerInventoryCollectionSettingsCustomPathIdDelete(ctx, id).Execute()

Delete Custom Path from Computer Inventory Collection Settings



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
	id := "id_example" // string | id of Custom Path

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputerInventoryCollectionSettingsAPI.V1ComputerInventoryCollectionSettingsCustomPathIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryCollectionSettingsAPI.V1ComputerInventoryCollectionSettingsCustomPathIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of Custom Path | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputerInventoryCollectionSettingsCustomPathIdDeleteRequest struct via the builder pattern


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


## V1ComputerInventoryCollectionSettingsCustomPathPost

> HrefResponse V1ComputerInventoryCollectionSettingsCustomPathPost(ctx).CreatePath(createPath).Execute()

Create Computer Inventory Collection Settings Custom Path



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
	createPath := *openapiclient.NewCreatePath("APP", "/Example/Path/") // CreatePath | Computer inventory settings to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryCollectionSettingsAPI.V1ComputerInventoryCollectionSettingsCustomPathPost(context.Background()).CreatePath(createPath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryCollectionSettingsAPI.V1ComputerInventoryCollectionSettingsCustomPathPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputerInventoryCollectionSettingsCustomPathPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryCollectionSettingsAPI.V1ComputerInventoryCollectionSettingsCustomPathPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputerInventoryCollectionSettingsCustomPathPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPath** | [**CreatePath**](CreatePath.md) | Computer inventory settings to update | 

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


## V1ComputerInventoryCollectionSettingsGet

> ComputerInventoryCollectionSettings V1ComputerInventoryCollectionSettingsGet(ctx).Execute()

Returns computer inventory settings



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
	resp, r, err := apiClient.ComputerInventoryCollectionSettingsAPI.V1ComputerInventoryCollectionSettingsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryCollectionSettingsAPI.V1ComputerInventoryCollectionSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputerInventoryCollectionSettingsGet`: ComputerInventoryCollectionSettings
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryCollectionSettingsAPI.V1ComputerInventoryCollectionSettingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputerInventoryCollectionSettingsGetRequest struct via the builder pattern


### Return type

[**ComputerInventoryCollectionSettings**](ComputerInventoryCollectionSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ComputerInventoryCollectionSettingsPatch

> ComputerInventoryCollectionSettings V1ComputerInventoryCollectionSettingsPatch(ctx).ComputerInventoryCollectionSettings(computerInventoryCollectionSettings).Execute()

Update computer inventory settings



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
	computerInventoryCollectionSettings := *openapiclient.NewComputerInventoryCollectionSettings() // ComputerInventoryCollectionSettings | Computer inventory settings to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryCollectionSettingsAPI.V1ComputerInventoryCollectionSettingsPatch(context.Background()).ComputerInventoryCollectionSettings(computerInventoryCollectionSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryCollectionSettingsAPI.V1ComputerInventoryCollectionSettingsPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputerInventoryCollectionSettingsPatch`: ComputerInventoryCollectionSettings
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryCollectionSettingsAPI.V1ComputerInventoryCollectionSettingsPatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputerInventoryCollectionSettingsPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **computerInventoryCollectionSettings** | [**ComputerInventoryCollectionSettings**](ComputerInventoryCollectionSettings.md) | Computer inventory settings to update | 

### Return type

[**ComputerInventoryCollectionSettings**](ComputerInventoryCollectionSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

