# \ManagedSoftwareUpdatesAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ManagedSoftwareUpdatesAvailableUpdatesGet**](ManagedSoftwareUpdatesAPI.md#V1ManagedSoftwareUpdatesAvailableUpdatesGet) | **Get** /v1/managed-software-updates/available-updates | Retrieve available macOS and iOS Managed Software Updates
[**V1ManagedSoftwareUpdatesPlansFeatureToggleGet**](ManagedSoftwareUpdatesAPI.md#V1ManagedSoftwareUpdatesPlansFeatureToggleGet) | **Get** /v1/managed-software-updates/plans/feature-toggle | Retrieve Status of the Feature Toggle
[**V1ManagedSoftwareUpdatesPlansFeatureTogglePut**](ManagedSoftwareUpdatesAPI.md#V1ManagedSoftwareUpdatesPlansFeatureTogglePut) | **Put** /v1/managed-software-updates/plans/feature-toggle | Updates Feature Toggle Value
[**V1ManagedSoftwareUpdatesPlansGet**](ManagedSoftwareUpdatesAPI.md#V1ManagedSoftwareUpdatesPlansGet) | **Get** /v1/managed-software-updates/plans | Retrieve Managed Software Update Plans
[**V1ManagedSoftwareUpdatesPlansGroupIdGet**](ManagedSoftwareUpdatesAPI.md#V1ManagedSoftwareUpdatesPlansGroupIdGet) | **Get** /v1/managed-software-updates/plans/group/{id} | Retrieve Managed Software Update Plans for a Group
[**V1ManagedSoftwareUpdatesPlansGroupPost**](ManagedSoftwareUpdatesAPI.md#V1ManagedSoftwareUpdatesPlansGroupPost) | **Post** /v1/managed-software-updates/plans/group | Create Managed Software Update Plans for a Group
[**V1ManagedSoftwareUpdatesPlansIdGet**](ManagedSoftwareUpdatesAPI.md#V1ManagedSoftwareUpdatesPlansIdGet) | **Get** /v1/managed-software-updates/plans/{id} | Retrieve a Managed Software Update Plan
[**V1ManagedSoftwareUpdatesPlansPost**](ManagedSoftwareUpdatesAPI.md#V1ManagedSoftwareUpdatesPlansPost) | **Post** /v1/managed-software-updates/plans | Create a Managed Software Update Plan
[**V1ManagedSoftwareUpdatesUpdateStatusesComputerGroupsIdGet**](ManagedSoftwareUpdatesAPI.md#V1ManagedSoftwareUpdatesUpdateStatusesComputerGroupsIdGet) | **Get** /v1/managed-software-updates/update-statuses/computer-groups/{id} | Retrieve Managed Software Update Statuses for Computer Groups
[**V1ManagedSoftwareUpdatesUpdateStatusesComputersIdGet**](ManagedSoftwareUpdatesAPI.md#V1ManagedSoftwareUpdatesUpdateStatusesComputersIdGet) | **Get** /v1/managed-software-updates/update-statuses/computers/{id} | Retrieve Managed Software Update Statuses for Computers
[**V1ManagedSoftwareUpdatesUpdateStatusesGet**](ManagedSoftwareUpdatesAPI.md#V1ManagedSoftwareUpdatesUpdateStatusesGet) | **Get** /v1/managed-software-updates/update-statuses | Retrieve Managed Software Update Statuses
[**V1ManagedSoftwareUpdatesUpdateStatusesMobileDeviceGroupsIdGet**](ManagedSoftwareUpdatesAPI.md#V1ManagedSoftwareUpdatesUpdateStatusesMobileDeviceGroupsIdGet) | **Get** /v1/managed-software-updates/update-statuses/mobile-device-groups/{id} | Retrieve Managed Software Update Statuses for Mobile Device Groups
[**V1ManagedSoftwareUpdatesUpdateStatusesMobileDevicesIdGet**](ManagedSoftwareUpdatesAPI.md#V1ManagedSoftwareUpdatesUpdateStatusesMobileDevicesIdGet) | **Get** /v1/managed-software-updates/update-statuses/mobile-devices/{id} | Retrieve Managed Software Update Statuses for Mobile Devices



## V1ManagedSoftwareUpdatesAvailableUpdatesGet

> AvailableOsUpdates V1ManagedSoftwareUpdatesAvailableUpdatesGet(ctx).Execute()

Retrieve available macOS and iOS Managed Software Updates



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
	resp, r, err := apiClient.ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesAvailableUpdatesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesAvailableUpdatesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ManagedSoftwareUpdatesAvailableUpdatesGet`: AvailableOsUpdates
	fmt.Fprintf(os.Stdout, "Response from `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesAvailableUpdatesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ManagedSoftwareUpdatesAvailableUpdatesGetRequest struct via the builder pattern


### Return type

[**AvailableOsUpdates**](AvailableOsUpdates.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ManagedSoftwareUpdatesPlansFeatureToggleGet

> ManagedSoftwareUpdatePlanToggle V1ManagedSoftwareUpdatesPlansFeatureToggleGet(ctx).Execute()

Retrieve Status of the Feature Toggle



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
	resp, r, err := apiClient.ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesPlansFeatureToggleGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesPlansFeatureToggleGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ManagedSoftwareUpdatesPlansFeatureToggleGet`: ManagedSoftwareUpdatePlanToggle
	fmt.Fprintf(os.Stdout, "Response from `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesPlansFeatureToggleGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ManagedSoftwareUpdatesPlansFeatureToggleGetRequest struct via the builder pattern


### Return type

[**ManagedSoftwareUpdatePlanToggle**](ManagedSoftwareUpdatePlanToggle.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ManagedSoftwareUpdatesPlansFeatureTogglePut

> ManagedSoftwareUpdatePlanToggle V1ManagedSoftwareUpdatesPlansFeatureTogglePut(ctx).ManagedSoftwareUpdatePlanToggle(managedSoftwareUpdatePlanToggle).Execute()

Updates Feature Toggle Value



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
	managedSoftwareUpdatePlanToggle := *openapiclient.NewManagedSoftwareUpdatePlanToggle(false) // ManagedSoftwareUpdatePlanToggle | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesPlansFeatureTogglePut(context.Background()).ManagedSoftwareUpdatePlanToggle(managedSoftwareUpdatePlanToggle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesPlansFeatureTogglePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ManagedSoftwareUpdatesPlansFeatureTogglePut`: ManagedSoftwareUpdatePlanToggle
	fmt.Fprintf(os.Stdout, "Response from `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesPlansFeatureTogglePut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ManagedSoftwareUpdatesPlansFeatureTogglePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **managedSoftwareUpdatePlanToggle** | [**ManagedSoftwareUpdatePlanToggle**](ManagedSoftwareUpdatePlanToggle.md) |  | 

### Return type

[**ManagedSoftwareUpdatePlanToggle**](ManagedSoftwareUpdatePlanToggle.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ManagedSoftwareUpdatesPlansGet

> ManagedSoftwareUpdatePlans V1ManagedSoftwareUpdatesPlansGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Retrieve Managed Software Update Plans



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is planUuid:asc. Multiple sort criteria are supported and must be separated with a comma. (optional) (default to ["planUuid:asc"])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter Managed Software Updates collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: planUuid, device.deviceId, device.objectType, updateAction, versionType, specificVersion, maxDeferrals, forceInstallLocalDateTime, state. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesPlansGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesPlansGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ManagedSoftwareUpdatesPlansGet`: ManagedSoftwareUpdatePlans
	fmt.Fprintf(os.Stdout, "Response from `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesPlansGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ManagedSoftwareUpdatesPlansGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is planUuid:asc. Multiple sort criteria are supported and must be separated with a comma. | [default to [&quot;planUuid:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter Managed Software Updates collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: planUuid, device.deviceId, device.objectType, updateAction, versionType, specificVersion, maxDeferrals, forceInstallLocalDateTime, state. | [default to &quot;&quot;]

### Return type

[**ManagedSoftwareUpdatePlans**](ManagedSoftwareUpdatePlans.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ManagedSoftwareUpdatesPlansGroupIdGet

> ManagedSoftwareUpdatePlans V1ManagedSoftwareUpdatesPlansGroupIdGet(ctx, id).GroupType(groupType).Execute()

Retrieve Managed Software Update Plans for a Group



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
	id := "id_example" // string | Managed Software Update Group Id
	groupType := "COMPUTER_GROUP" // string | Managed Software Update Group Type, Available options are \"COMPUTER_GROUP\" or \"MOBILE_DEVICE_GROUP\"

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesPlansGroupIdGet(context.Background(), id).GroupType(groupType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesPlansGroupIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ManagedSoftwareUpdatesPlansGroupIdGet`: ManagedSoftwareUpdatePlans
	fmt.Fprintf(os.Stdout, "Response from `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesPlansGroupIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Managed Software Update Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ManagedSoftwareUpdatesPlansGroupIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupType** | **string** | Managed Software Update Group Type, Available options are \&quot;COMPUTER_GROUP\&quot; or \&quot;MOBILE_DEVICE_GROUP\&quot; | 

### Return type

[**ManagedSoftwareUpdatePlans**](ManagedSoftwareUpdatePlans.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ManagedSoftwareUpdatesPlansGroupPost

> ManagedSoftwareUpdatePlanPostResponse V1ManagedSoftwareUpdatesPlansGroupPost(ctx).ManagedSoftwareUpdatePlanGroupPost(managedSoftwareUpdatePlanGroupPost).Execute()

Create Managed Software Update Plans for a Group



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
	managedSoftwareUpdatePlanGroupPost := *openapiclient.NewManagedSoftwareUpdatePlanGroupPost(*openapiclient.NewPlanGroupPost("1", "COMPUTER_GROUP"), *openapiclient.NewPlanConfigurationPost("DOWNLOAD_INSTALL", "SPECIFIC_VERSION")) // ManagedSoftwareUpdatePlanGroupPost | Managed Software Update Plan to create for Group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesPlansGroupPost(context.Background()).ManagedSoftwareUpdatePlanGroupPost(managedSoftwareUpdatePlanGroupPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesPlansGroupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ManagedSoftwareUpdatesPlansGroupPost`: ManagedSoftwareUpdatePlanPostResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesPlansGroupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ManagedSoftwareUpdatesPlansGroupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **managedSoftwareUpdatePlanGroupPost** | [**ManagedSoftwareUpdatePlanGroupPost**](ManagedSoftwareUpdatePlanGroupPost.md) | Managed Software Update Plan to create for Group | 

### Return type

[**ManagedSoftwareUpdatePlanPostResponse**](ManagedSoftwareUpdatePlanPostResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ManagedSoftwareUpdatesPlansIdGet

> ManagedSoftwareUpdatePlan V1ManagedSoftwareUpdatesPlansIdGet(ctx, id).Execute()

Retrieve a Managed Software Update Plan



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
	id := "id_example" // string | Managed Software Update Plan Uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesPlansIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesPlansIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ManagedSoftwareUpdatesPlansIdGet`: ManagedSoftwareUpdatePlan
	fmt.Fprintf(os.Stdout, "Response from `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesPlansIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Managed Software Update Plan Uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ManagedSoftwareUpdatesPlansIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedSoftwareUpdatePlan**](ManagedSoftwareUpdatePlan.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ManagedSoftwareUpdatesPlansPost

> ManagedSoftwareUpdatePlanPostResponse V1ManagedSoftwareUpdatesPlansPost(ctx).ManagedSoftwareUpdatePlanPost(managedSoftwareUpdatePlanPost).Execute()

Create a Managed Software Update Plan



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
	managedSoftwareUpdatePlanPost := *openapiclient.NewManagedSoftwareUpdatePlanPost([]openapiclient.PlanDevicePost{*openapiclient.NewPlanDevicePost("1", "COMPUTER")}, *openapiclient.NewPlanConfigurationPost("DOWNLOAD_INSTALL", "SPECIFIC_VERSION")) // ManagedSoftwareUpdatePlanPost | Managed Software Update Plan to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesPlansPost(context.Background()).ManagedSoftwareUpdatePlanPost(managedSoftwareUpdatePlanPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesPlansPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ManagedSoftwareUpdatesPlansPost`: ManagedSoftwareUpdatePlanPostResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesPlansPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ManagedSoftwareUpdatesPlansPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **managedSoftwareUpdatePlanPost** | [**ManagedSoftwareUpdatePlanPost**](ManagedSoftwareUpdatePlanPost.md) | Managed Software Update Plan to create | 

### Return type

[**ManagedSoftwareUpdatePlanPostResponse**](ManagedSoftwareUpdatePlanPostResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ManagedSoftwareUpdatesUpdateStatusesComputerGroupsIdGet

> ManagedSoftwareUpdateStatuses V1ManagedSoftwareUpdatesUpdateStatusesComputerGroupsIdGet(ctx, id).Execute()

Retrieve Managed Software Update Statuses for Computer Groups



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
	id := "id_example" // string | Computer Group identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesUpdateStatusesComputerGroupsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesUpdateStatusesComputerGroupsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ManagedSoftwareUpdatesUpdateStatusesComputerGroupsIdGet`: ManagedSoftwareUpdateStatuses
	fmt.Fprintf(os.Stdout, "Response from `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesUpdateStatusesComputerGroupsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Computer Group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ManagedSoftwareUpdatesUpdateStatusesComputerGroupsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedSoftwareUpdateStatuses**](ManagedSoftwareUpdateStatuses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ManagedSoftwareUpdatesUpdateStatusesComputersIdGet

> ManagedSoftwareUpdateStatuses V1ManagedSoftwareUpdatesUpdateStatusesComputersIdGet(ctx, id).Execute()

Retrieve Managed Software Update Statuses for Computers



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
	id := "id_example" // string | Computer identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesUpdateStatusesComputersIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesUpdateStatusesComputersIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ManagedSoftwareUpdatesUpdateStatusesComputersIdGet`: ManagedSoftwareUpdateStatuses
	fmt.Fprintf(os.Stdout, "Response from `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesUpdateStatusesComputersIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Computer identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ManagedSoftwareUpdatesUpdateStatusesComputersIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedSoftwareUpdateStatuses**](ManagedSoftwareUpdateStatuses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ManagedSoftwareUpdatesUpdateStatusesGet

> ManagedSoftwareUpdateStatuses V1ManagedSoftwareUpdatesUpdateStatusesGet(ctx).Filter(filter).Execute()

Retrieve Managed Software Update Statuses



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
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter Managed Software Updates collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: osUpdatesStatusId, device.deviceId, device.objectType, downloaded, downloadPercentComplete, productKey, status, deferralsRemaining, maxDeferrals, nextScheduledInstall, created and updated. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesUpdateStatusesGet(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesUpdateStatusesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ManagedSoftwareUpdatesUpdateStatusesGet`: ManagedSoftwareUpdateStatuses
	fmt.Fprintf(os.Stdout, "Response from `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesUpdateStatusesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ManagedSoftwareUpdatesUpdateStatusesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Query in the RSQL format, allowing to filter Managed Software Updates collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: osUpdatesStatusId, device.deviceId, device.objectType, downloaded, downloadPercentComplete, productKey, status, deferralsRemaining, maxDeferrals, nextScheduledInstall, created and updated. | [default to &quot;&quot;]

### Return type

[**ManagedSoftwareUpdateStatuses**](ManagedSoftwareUpdateStatuses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ManagedSoftwareUpdatesUpdateStatusesMobileDeviceGroupsIdGet

> ManagedSoftwareUpdateStatuses V1ManagedSoftwareUpdatesUpdateStatusesMobileDeviceGroupsIdGet(ctx, id).Execute()

Retrieve Managed Software Update Statuses for Mobile Device Groups



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
	id := "id_example" // string | Mobile Device Group identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesUpdateStatusesMobileDeviceGroupsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesUpdateStatusesMobileDeviceGroupsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ManagedSoftwareUpdatesUpdateStatusesMobileDeviceGroupsIdGet`: ManagedSoftwareUpdateStatuses
	fmt.Fprintf(os.Stdout, "Response from `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesUpdateStatusesMobileDeviceGroupsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Mobile Device Group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ManagedSoftwareUpdatesUpdateStatusesMobileDeviceGroupsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedSoftwareUpdateStatuses**](ManagedSoftwareUpdateStatuses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ManagedSoftwareUpdatesUpdateStatusesMobileDevicesIdGet

> ManagedSoftwareUpdateStatuses V1ManagedSoftwareUpdatesUpdateStatusesMobileDevicesIdGet(ctx, id).Execute()

Retrieve Managed Software Update Statuses for Mobile Devices



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
	id := "id_example" // string | Mobile Device identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesUpdateStatusesMobileDevicesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesUpdateStatusesMobileDevicesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ManagedSoftwareUpdatesUpdateStatusesMobileDevicesIdGet`: ManagedSoftwareUpdateStatuses
	fmt.Fprintf(os.Stdout, "Response from `ManagedSoftwareUpdatesAPI.V1ManagedSoftwareUpdatesUpdateStatusesMobileDevicesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Mobile Device identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ManagedSoftwareUpdatesUpdateStatusesMobileDevicesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManagedSoftwareUpdateStatuses**](ManagedSoftwareUpdateStatuses.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

