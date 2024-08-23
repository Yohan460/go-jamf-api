# \SitesAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SitesGet**](SitesAPI.md#V1SitesGet) | **Get** /v1/sites | Find all sites 
[**V1SitesIdObjectsGet**](SitesAPI.md#V1SitesIdObjectsGet) | **Get** /v1/sites/{id}/objects | Find and filter site objects for a site ID 



## V1SitesGet

> []V1Site V1SitesGet(ctx).Execute()

Find all sites 



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
	resp, r, err := apiClient.SitesAPI.V1SitesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.V1SitesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SitesGet`: []V1Site
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.V1SitesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SitesGetRequest struct via the builder pattern


### Return type

[**[]V1Site**](V1Site.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SitesIdObjectsGet

> []SiteObject V1SitesIdObjectsGet(ctx, id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Find and filter site objects for a site ID 



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
	id := "1" // string | Site ID to get objects for 
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: `property:asc/desc`. Default sort is `objectType:asc`. Multiple sort criteria are supported and must be separated with a comma.  Example: `sort=objectId:asc,objectType:desc`.  (optional) (default to ["objectType:asc"])
	filter := "filter_example" // string | Query in the RSQL format, allowing filter of site object information. Default filter returns all objects for the site ID.  Fields allowed in the query: `objectType`, `objectId`  Example: `filter=objectType==\"User\"`      List of `objectType` options (case-insensitive) [\"Computer\", \"Peripheral\", \"Licensed Software\", \"Licensed Software Template\", \"Policy\", \"macOS Configuration Profile\", \"Restricted Software\", \"Managed Preference Profile\", \"Computer Group\", \"Mobile Device\", \"Apple TV\", \"Android Device\", \"User Group\", \"iOS Configuration Profile\", \"Mobile Device App\", \"E-book\", \"Mobile Device Group\", \"Classroom\", \"Advanced Computer Search\", \"Advanced Mobile Search\", \"Advanced User Search\", \"Advanced User Content Search\", \"Computer Invitation\", \"Mobile Device Invitation\", \"Mobile Device Enrollment Profile\", \"Device Enrollment Program Instance\", \"Mobile Device Prestage\", \"Computer DEP Prestage\", \"Enrollment Customization\", \"VPP Location\", \"VPP Subscription\", \"VPP Invitation\", \"VPP Assignment\", \"User\", \"Network Integration\", \"Mac App\", \"App Installer\", \"BYO Profile\", \"Self Service Plugin\", \"Software Title\", \"Patch Software Title Summary\", \"Patch Policy\", \"Patch Software Title Configuration\", \"Change Password\", \"Mobile Device Inventory\", \"Computer Inventory\", \"Change Management\", \"Licensed Software License\"]  (optional) (default to "objectType==\"User\"")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAPI.V1SitesIdObjectsGet(context.Background(), id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.V1SitesIdObjectsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SitesIdObjectsGet`: []SiteObject
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.V1SitesIdObjectsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Site ID to get objects for  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SitesIdObjectsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: &#x60;property:asc/desc&#x60;. Default sort is &#x60;objectType:asc&#x60;. Multiple sort criteria are supported and must be separated with a comma.  Example: &#x60;sort&#x3D;objectId:asc,objectType:desc&#x60;.  | [default to [&quot;objectType:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing filter of site object information. Default filter returns all objects for the site ID.  Fields allowed in the query: &#x60;objectType&#x60;, &#x60;objectId&#x60;  Example: &#x60;filter&#x3D;objectType&#x3D;&#x3D;\&quot;User\&quot;&#x60;      List of &#x60;objectType&#x60; options (case-insensitive) [\&quot;Computer\&quot;, \&quot;Peripheral\&quot;, \&quot;Licensed Software\&quot;, \&quot;Licensed Software Template\&quot;, \&quot;Policy\&quot;, \&quot;macOS Configuration Profile\&quot;, \&quot;Restricted Software\&quot;, \&quot;Managed Preference Profile\&quot;, \&quot;Computer Group\&quot;, \&quot;Mobile Device\&quot;, \&quot;Apple TV\&quot;, \&quot;Android Device\&quot;, \&quot;User Group\&quot;, \&quot;iOS Configuration Profile\&quot;, \&quot;Mobile Device App\&quot;, \&quot;E-book\&quot;, \&quot;Mobile Device Group\&quot;, \&quot;Classroom\&quot;, \&quot;Advanced Computer Search\&quot;, \&quot;Advanced Mobile Search\&quot;, \&quot;Advanced User Search\&quot;, \&quot;Advanced User Content Search\&quot;, \&quot;Computer Invitation\&quot;, \&quot;Mobile Device Invitation\&quot;, \&quot;Mobile Device Enrollment Profile\&quot;, \&quot;Device Enrollment Program Instance\&quot;, \&quot;Mobile Device Prestage\&quot;, \&quot;Computer DEP Prestage\&quot;, \&quot;Enrollment Customization\&quot;, \&quot;VPP Location\&quot;, \&quot;VPP Subscription\&quot;, \&quot;VPP Invitation\&quot;, \&quot;VPP Assignment\&quot;, \&quot;User\&quot;, \&quot;Network Integration\&quot;, \&quot;Mac App\&quot;, \&quot;App Installer\&quot;, \&quot;BYO Profile\&quot;, \&quot;Self Service Plugin\&quot;, \&quot;Software Title\&quot;, \&quot;Patch Software Title Summary\&quot;, \&quot;Patch Policy\&quot;, \&quot;Patch Software Title Configuration\&quot;, \&quot;Change Password\&quot;, \&quot;Mobile Device Inventory\&quot;, \&quot;Computer Inventory\&quot;, \&quot;Change Management\&quot;, \&quot;Licensed Software License\&quot;]  | [default to &quot;objectType&#x3D;&#x3D;\&quot;User\&quot;&quot;]

### Return type

[**[]SiteObject**](SiteObject.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

