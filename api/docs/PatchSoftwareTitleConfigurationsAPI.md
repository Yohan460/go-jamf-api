# \PatchSoftwareTitleConfigurationsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2PatchSoftwareTitleConfigurationsGet**](PatchSoftwareTitleConfigurationsAPI.md#V2PatchSoftwareTitleConfigurationsGet) | **Get** /v2/patch-software-title-configurations | Retrieve Patch Software Title Configurations
[**V2PatchSoftwareTitleConfigurationsIdDashboardDelete**](PatchSoftwareTitleConfigurationsAPI.md#V2PatchSoftwareTitleConfigurationsIdDashboardDelete) | **Delete** /v2/patch-software-title-configurations/{id}/dashboard | Remove a software title configuration from the dashboard 
[**V2PatchSoftwareTitleConfigurationsIdDashboardGet**](PatchSoftwareTitleConfigurationsAPI.md#V2PatchSoftwareTitleConfigurationsIdDashboardGet) | **Get** /v2/patch-software-title-configurations/{id}/dashboard | Return whether or not the requested software title configuration is on the dashboard 
[**V2PatchSoftwareTitleConfigurationsIdDashboardPost**](PatchSoftwareTitleConfigurationsAPI.md#V2PatchSoftwareTitleConfigurationsIdDashboardPost) | **Post** /v2/patch-software-title-configurations/{id}/dashboard | Add a software title configuration to the dashboard 
[**V2PatchSoftwareTitleConfigurationsIdDefinitionsGet**](PatchSoftwareTitleConfigurationsAPI.md#V2PatchSoftwareTitleConfigurationsIdDefinitionsGet) | **Get** /v2/patch-software-title-configurations/{id}/definitions | Retrieve Patch Software Title Definitions with the supplied id
[**V2PatchSoftwareTitleConfigurationsIdDelete**](PatchSoftwareTitleConfigurationsAPI.md#V2PatchSoftwareTitleConfigurationsIdDelete) | **Delete** /v2/patch-software-title-configurations/{id} | Delete Patch Software Title Configurations with the supplied id
[**V2PatchSoftwareTitleConfigurationsIdDependenciesGet**](PatchSoftwareTitleConfigurationsAPI.md#V2PatchSoftwareTitleConfigurationsIdDependenciesGet) | **Get** /v2/patch-software-title-configurations/{id}/dependencies | Retrieve list of Patch Software Title Configuration Dependencies
[**V2PatchSoftwareTitleConfigurationsIdExportReportGet**](PatchSoftwareTitleConfigurationsAPI.md#V2PatchSoftwareTitleConfigurationsIdExportReportGet) | **Get** /v2/patch-software-title-configurations/{id}/export-report | Export Patch Reporting Data
[**V2PatchSoftwareTitleConfigurationsIdExtensionAttributesGet**](PatchSoftwareTitleConfigurationsAPI.md#V2PatchSoftwareTitleConfigurationsIdExtensionAttributesGet) | **Get** /v2/patch-software-title-configurations/{id}/extension-attributes | Retrieve Software Title Extension Attributes with the supplied id
[**V2PatchSoftwareTitleConfigurationsIdGet**](PatchSoftwareTitleConfigurationsAPI.md#V2PatchSoftwareTitleConfigurationsIdGet) | **Get** /v2/patch-software-title-configurations/{id} | Retrieve Patch Software Title Configurations with the supplied id
[**V2PatchSoftwareTitleConfigurationsIdHistoryGet**](PatchSoftwareTitleConfigurationsAPI.md#V2PatchSoftwareTitleConfigurationsIdHistoryGet) | **Get** /v2/patch-software-title-configurations/{id}/history | Get specified Patch Software Title Configuration history object 
[**V2PatchSoftwareTitleConfigurationsIdHistoryPost**](PatchSoftwareTitleConfigurationsAPI.md#V2PatchSoftwareTitleConfigurationsIdHistoryPost) | **Post** /v2/patch-software-title-configurations/{id}/history | Add Patch Software Title Configuration history object notes 
[**V2PatchSoftwareTitleConfigurationsIdPatch**](PatchSoftwareTitleConfigurationsAPI.md#V2PatchSoftwareTitleConfigurationsIdPatch) | **Patch** /v2/patch-software-title-configurations/{id} | Update Patch Software Title Configurations
[**V2PatchSoftwareTitleConfigurationsIdPatchReportGet**](PatchSoftwareTitleConfigurationsAPI.md#V2PatchSoftwareTitleConfigurationsIdPatchReportGet) | **Get** /v2/patch-software-title-configurations/{id}/patch-report | Retrieve Patch Software Title Configuration Patch Report
[**V2PatchSoftwareTitleConfigurationsIdPatchSummaryGet**](PatchSoftwareTitleConfigurationsAPI.md#V2PatchSoftwareTitleConfigurationsIdPatchSummaryGet) | **Get** /v2/patch-software-title-configurations/{id}/patch-summary | Return Active Patch Summary
[**V2PatchSoftwareTitleConfigurationsIdPatchSummaryVersionsGet**](PatchSoftwareTitleConfigurationsAPI.md#V2PatchSoftwareTitleConfigurationsIdPatchSummaryVersionsGet) | **Get** /v2/patch-software-title-configurations/{id}/patch-summary/versions | Returns patch versions
[**V2PatchSoftwareTitleConfigurationsPost**](PatchSoftwareTitleConfigurationsAPI.md#V2PatchSoftwareTitleConfigurationsPost) | **Post** /v2/patch-software-title-configurations | Create Patch Software Title Configurations



## V2PatchSoftwareTitleConfigurationsGet

> []PatchSoftwareTitleConfiguration V2PatchSoftwareTitleConfigurationsGet(ctx).Execute()

Retrieve Patch Software Title Configurations



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
	resp, r, err := apiClient.PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2PatchSoftwareTitleConfigurationsGet`: []PatchSoftwareTitleConfiguration
	fmt.Fprintf(os.Stdout, "Response from `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchSoftwareTitleConfigurationsGetRequest struct via the builder pattern


### Return type

[**[]PatchSoftwareTitleConfiguration**](PatchSoftwareTitleConfiguration.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PatchSoftwareTitleConfigurationsIdDashboardDelete

> V2PatchSoftwareTitleConfigurationsIdDashboardDelete(ctx, id).Execute()

Remove a software title configuration from the dashboard 



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
	id := "id_example" // string | software title configuration id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdDashboardDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdDashboardDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | software title configuration id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchSoftwareTitleConfigurationsIdDashboardDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PatchSoftwareTitleConfigurationsIdDashboardGet

> SoftwareTitleConfigurationOnDashboard V2PatchSoftwareTitleConfigurationsIdDashboardGet(ctx, id).Execute()

Return whether or not the requested software title configuration is on the dashboard 



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
	id := "id_example" // string | software title configuration id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdDashboardGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdDashboardGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2PatchSoftwareTitleConfigurationsIdDashboardGet`: SoftwareTitleConfigurationOnDashboard
	fmt.Fprintf(os.Stdout, "Response from `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdDashboardGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | software title configuration id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchSoftwareTitleConfigurationsIdDashboardGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SoftwareTitleConfigurationOnDashboard**](SoftwareTitleConfigurationOnDashboard.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PatchSoftwareTitleConfigurationsIdDashboardPost

> V2PatchSoftwareTitleConfigurationsIdDashboardPost(ctx, id).Execute()

Add a software title configuration to the dashboard 



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
	id := "id_example" // string | software title configuration id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdDashboardPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdDashboardPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | software title configuration id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchSoftwareTitleConfigurationsIdDashboardPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PatchSoftwareTitleConfigurationsIdDefinitionsGet

> PatchSoftwareTitleDefinitions V2PatchSoftwareTitleConfigurationsIdDefinitionsGet(ctx, id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Retrieve Patch Software Title Definitions with the supplied id



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
	id := "id_example" // string | Patch Software Title identifier
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is absoluteOrderId:asc. Multiple sort criteria are supported and must be separated with a comma. (optional) (default to ["absoluteOrderId:asc"])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter Patch Software Title Definition collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, version, minimumOperatingSystem, releaseDate, reboot, standalone and absoluteOrderId. This param can be combined with paging and sorting. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdDefinitionsGet(context.Background(), id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdDefinitionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2PatchSoftwareTitleConfigurationsIdDefinitionsGet`: PatchSoftwareTitleDefinitions
	fmt.Fprintf(os.Stdout, "Response from `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdDefinitionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Patch Software Title identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchSoftwareTitleConfigurationsIdDefinitionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is absoluteOrderId:asc. Multiple sort criteria are supported and must be separated with a comma. | [default to [&quot;absoluteOrderId:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter Patch Software Title Definition collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, version, minimumOperatingSystem, releaseDate, reboot, standalone and absoluteOrderId. This param can be combined with paging and sorting. | [default to &quot;&quot;]

### Return type

[**PatchSoftwareTitleDefinitions**](PatchSoftwareTitleDefinitions.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PatchSoftwareTitleConfigurationsIdDelete

> V2PatchSoftwareTitleConfigurationsIdDelete(ctx, id).Execute()

Delete Patch Software Title Configurations with the supplied id



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
	id := "id_example" // string | Patch Software Title Configurations identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Patch Software Title Configurations identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchSoftwareTitleConfigurationsIdDeleteRequest struct via the builder pattern


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


## V2PatchSoftwareTitleConfigurationsIdDependenciesGet

> PatchSoftwareTitleConfigurationDependencies V2PatchSoftwareTitleConfigurationsIdDependenciesGet(ctx, id).Execute()

Retrieve list of Patch Software Title Configuration Dependencies



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
	id := "id_example" // string | Patch Software Title Configuration Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdDependenciesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdDependenciesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2PatchSoftwareTitleConfigurationsIdDependenciesGet`: PatchSoftwareTitleConfigurationDependencies
	fmt.Fprintf(os.Stdout, "Response from `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdDependenciesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Patch Software Title Configuration Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchSoftwareTitleConfigurationsIdDependenciesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PatchSoftwareTitleConfigurationDependencies**](PatchSoftwareTitleConfigurationDependencies.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PatchSoftwareTitleConfigurationsIdExportReportGet

> interface{} V2PatchSoftwareTitleConfigurationsIdExportReportGet(ctx, id).ColumnsToExport(columnsToExport).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Accept(accept).Execute()

Export Patch Reporting Data



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
	id := "id_example" // string | Patch Software Title Configurations identifier
	columnsToExport := []string{"Inner_example"} // []string | List of column names to export (default to ["computerName","deviceId","username","operatingSystemVersion","lastContactTime","buildingName","departmentName","siteName","version"])
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 | Leave blank for full export (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is computerName:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=id:desc,name:asc Supported fields: computerName, deviceId, username, operatingSystemVersion, lastContactTime, buildingName, departmentName, siteName, version (optional) (default to ["computerName:asc"])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter Patch Report collection on version equality only. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: version. Comparators allowed in the query: ==, != This param can be combined with paging and sorting. (optional) (default to "")
	accept := "accept_example" // string | File (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdExportReportGet(context.Background(), id).ColumnsToExport(columnsToExport).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Accept(accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdExportReportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2PatchSoftwareTitleConfigurationsIdExportReportGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdExportReportGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Patch Software Title Configurations identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchSoftwareTitleConfigurationsIdExportReportGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **columnsToExport** | **[]string** | List of column names to export | [default to [&quot;computerName&quot;,&quot;deviceId&quot;,&quot;username&quot;,&quot;operatingSystemVersion&quot;,&quot;lastContactTime&quot;,&quot;buildingName&quot;,&quot;departmentName&quot;,&quot;siteName&quot;,&quot;version&quot;]]
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** | Leave blank for full export | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is computerName:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;id:desc,name:asc Supported fields: computerName, deviceId, username, operatingSystemVersion, lastContactTime, buildingName, departmentName, siteName, version | [default to [&quot;computerName:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter Patch Report collection on version equality only. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: version. Comparators allowed in the query: &#x3D;&#x3D;, !&#x3D; This param can be combined with paging and sorting. | [default to &quot;&quot;]
 **accept** | **string** | File | 

### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, text/tab, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PatchSoftwareTitleConfigurationsIdExtensionAttributesGet

> []PatchSoftwareTitleExtensionAttributes V2PatchSoftwareTitleConfigurationsIdExtensionAttributesGet(ctx, id).Execute()

Retrieve Software Title Extension Attributes with the supplied id



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
	id := "id_example" // string | Patch Software Title identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdExtensionAttributesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdExtensionAttributesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2PatchSoftwareTitleConfigurationsIdExtensionAttributesGet`: []PatchSoftwareTitleExtensionAttributes
	fmt.Fprintf(os.Stdout, "Response from `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdExtensionAttributesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Patch Software Title identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchSoftwareTitleConfigurationsIdExtensionAttributesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PatchSoftwareTitleExtensionAttributes**](PatchSoftwareTitleExtensionAttributes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PatchSoftwareTitleConfigurationsIdGet

> PatchSoftwareTitleConfiguration V2PatchSoftwareTitleConfigurationsIdGet(ctx, id).Execute()

Retrieve Patch Software Title Configurations with the supplied id



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
	id := "id_example" // string | Patch Software Title Configurations identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2PatchSoftwareTitleConfigurationsIdGet`: PatchSoftwareTitleConfiguration
	fmt.Fprintf(os.Stdout, "Response from `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Patch Software Title Configurations identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchSoftwareTitleConfigurationsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PatchSoftwareTitleConfiguration**](PatchSoftwareTitleConfiguration.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PatchSoftwareTitleConfigurationsIdHistoryGet

> HistorySearchResults V2PatchSoftwareTitleConfigurationsIdHistoryGet(ctx, id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get specified Patch Software Title Configuration history object 



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
	id := "id_example" // string | Patch Software Title Configuration Id
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma.  (optional) (default to ["date:desc"])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdHistoryGet(context.Background(), id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2PatchSoftwareTitleConfigurationsIdHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Patch Software Title Configuration Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchSoftwareTitleConfigurationsIdHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma.  | [default to [&quot;date:desc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15 | [default to &quot;&quot;]

### Return type

[**HistorySearchResults**](HistorySearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PatchSoftwareTitleConfigurationsIdHistoryPost

> HrefResponse V2PatchSoftwareTitleConfigurationsIdHistoryPost(ctx, id).ObjectHistoryNote(objectHistoryNote).Execute()

Add Patch Software Title Configuration history object notes 



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
	id := "id_example" // string | Patch Software Title Configuration Id
	objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | History notes to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdHistoryPost(context.Background(), id).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2PatchSoftwareTitleConfigurationsIdHistoryPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdHistoryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Patch Software Title Configuration Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchSoftwareTitleConfigurationsIdHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectHistoryNote** | [**ObjectHistoryNote**](ObjectHistoryNote.md) | History notes to create | 

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


## V2PatchSoftwareTitleConfigurationsIdPatch

> PatchSoftwareTitleConfiguration V2PatchSoftwareTitleConfigurationsIdPatch(ctx, id).PatchSoftwareTitleConfigurationPatch(patchSoftwareTitleConfigurationPatch).Execute()

Update Patch Software Title Configurations



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
	id := "id_example" // string | Patch Software Title Configurations identifier
	patchSoftwareTitleConfigurationPatch := *openapiclient.NewPatchSoftwareTitleConfigurationPatch() // PatchSoftwareTitleConfigurationPatch | Patch Software Title Configurations to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdPatch(context.Background(), id).PatchSoftwareTitleConfigurationPatch(patchSoftwareTitleConfigurationPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2PatchSoftwareTitleConfigurationsIdPatch`: PatchSoftwareTitleConfiguration
	fmt.Fprintf(os.Stdout, "Response from `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Patch Software Title Configurations identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchSoftwareTitleConfigurationsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchSoftwareTitleConfigurationPatch** | [**PatchSoftwareTitleConfigurationPatch**](PatchSoftwareTitleConfigurationPatch.md) | Patch Software Title Configurations to update | 

### Return type

[**PatchSoftwareTitleConfiguration**](PatchSoftwareTitleConfiguration.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PatchSoftwareTitleConfigurationsIdPatchReportGet

> PatchSoftwareTitleReportSearchResult V2PatchSoftwareTitleConfigurationsIdPatchReportGet(ctx, id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Retrieve Patch Software Title Configuration Patch Report



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
	id := "id_example" // string | Patch Software Title Configurations identifier
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is computerName:asc. Multiple sort criteria are supported and must be separated with a comma. Supported fields: computerName, deviceId, username, operatingSystemVersion, lastContactTime, buildingName, departmentName, siteName, version (optional) (default to ["computerName:asc"])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter Patch Report collection on version equality only. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: version. Comparators allowed in the query: ==, != This param can be combined with paging and sorting. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdPatchReportGet(context.Background(), id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdPatchReportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2PatchSoftwareTitleConfigurationsIdPatchReportGet`: PatchSoftwareTitleReportSearchResult
	fmt.Fprintf(os.Stdout, "Response from `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdPatchReportGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Patch Software Title Configurations identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchSoftwareTitleConfigurationsIdPatchReportGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is computerName:asc. Multiple sort criteria are supported and must be separated with a comma. Supported fields: computerName, deviceId, username, operatingSystemVersion, lastContactTime, buildingName, departmentName, siteName, version | [default to [&quot;computerName:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter Patch Report collection on version equality only. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: version. Comparators allowed in the query: &#x3D;&#x3D;, !&#x3D; This param can be combined with paging and sorting. | [default to &quot;&quot;]

### Return type

[**PatchSoftwareTitleReportSearchResult**](PatchSoftwareTitleReportSearchResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PatchSoftwareTitleConfigurationsIdPatchSummaryGet

> PatchSummary V2PatchSoftwareTitleConfigurationsIdPatchSummaryGet(ctx, id).Execute()

Return Active Patch Summary



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
	id := "id_example" // string | Patch id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdPatchSummaryGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdPatchSummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2PatchSoftwareTitleConfigurationsIdPatchSummaryGet`: PatchSummary
	fmt.Fprintf(os.Stdout, "Response from `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdPatchSummaryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Patch id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchSoftwareTitleConfigurationsIdPatchSummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PatchSummary**](PatchSummary.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PatchSoftwareTitleConfigurationsIdPatchSummaryVersionsGet

> []PatchSummaryVersion V2PatchSoftwareTitleConfigurationsIdPatchSummaryVersionsGet(ctx, id).Execute()

Returns patch versions



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
	id := "id_example" // string | Patch id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdPatchSummaryVersionsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdPatchSummaryVersionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2PatchSoftwareTitleConfigurationsIdPatchSummaryVersionsGet`: []PatchSummaryVersion
	fmt.Fprintf(os.Stdout, "Response from `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsIdPatchSummaryVersionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Patch id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchSoftwareTitleConfigurationsIdPatchSummaryVersionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PatchSummaryVersion**](PatchSummaryVersion.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2PatchSoftwareTitleConfigurationsPost

> HrefResponse V2PatchSoftwareTitleConfigurationsPost(ctx).PatchSoftwareTitleConfigurationBase(patchSoftwareTitleConfigurationBase).Execute()

Create Patch Software Title Configurations



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
	patchSoftwareTitleConfigurationBase := *openapiclient.NewPatchSoftwareTitleConfigurationBase("Google Chrome", "1") // PatchSoftwareTitleConfigurationBase | Software title configurations to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsPost(context.Background()).PatchSoftwareTitleConfigurationBase(patchSoftwareTitleConfigurationBase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2PatchSoftwareTitleConfigurationsPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `PatchSoftwareTitleConfigurationsAPI.V2PatchSoftwareTitleConfigurationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchSoftwareTitleConfigurationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchSoftwareTitleConfigurationBase** | [**PatchSoftwareTitleConfigurationBase**](PatchSoftwareTitleConfigurationBase.md) | Software title configurations to create | 

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

