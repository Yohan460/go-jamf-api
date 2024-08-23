# \EnrollmentAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AdueSessionTokenSettingsGet**](EnrollmentAPI.md#V1AdueSessionTokenSettingsGet) | **Get** /v1/adue-session-token-settings | Retrieve the Account Driven User Enrollment Session Token Settings 
[**V1AdueSessionTokenSettingsPut**](EnrollmentAPI.md#V1AdueSessionTokenSettingsPut) | **Put** /v1/adue-session-token-settings | Update Account Driven User Enrollment Session Token Settings. 
[**V2EnrollmentAccessGroupsGet**](EnrollmentAPI.md#V2EnrollmentAccessGroupsGet) | **Get** /v2/enrollment/access-groups | Retrieve the configured LDAP groups configured for User-Initiated Enrollment 
[**V2EnrollmentAccessGroupsPost**](EnrollmentAPI.md#V2EnrollmentAccessGroupsPost) | **Post** /v2/enrollment/access-groups | Add the configured LDAP group for User-Initiated Enrollment. 
[**V2EnrollmentAccessGroupsServerIdGroupIdDelete**](EnrollmentAPI.md#V2EnrollmentAccessGroupsServerIdGroupIdDelete) | **Delete** /v2/enrollment/access-groups/{serverId}/{groupId} | Delete an LDAP group&#39;s access to user initiated Enrollment 
[**V2EnrollmentAccessGroupsServerIdGroupIdGet**](EnrollmentAPI.md#V2EnrollmentAccessGroupsServerIdGroupIdGet) | **Get** /v2/enrollment/access-groups/{serverId}/{groupId} | Retrieve the configured LDAP groups configured for User-Initiated Enrollment 
[**V2EnrollmentAccessGroupsServerIdGroupIdPut**](EnrollmentAPI.md#V2EnrollmentAccessGroupsServerIdGroupIdPut) | **Put** /v2/enrollment/access-groups/{serverId}/{groupId} | Modify the configured LDAP groups configured for User-Initiated Enrollment 
[**V2EnrollmentFilteredLanguageCodesGet**](EnrollmentAPI.md#V2EnrollmentFilteredLanguageCodesGet) | **Get** /v2/enrollment/filtered-language-codes | Retrieve the list of languages and corresponding ISO 639-1 Codes but only those not already added to Enrollment 
[**V2EnrollmentGet**](EnrollmentAPI.md#V2EnrollmentGet) | **Get** /v2/enrollment | Get Enrollment object and Re-enrollment settings 
[**V2EnrollmentHistoryExportPost**](EnrollmentAPI.md#V2EnrollmentHistoryExportPost) | **Post** /v2/enrollment/history/export | Export enrollment history collection 
[**V2EnrollmentHistoryGet**](EnrollmentAPI.md#V2EnrollmentHistoryGet) | **Get** /v2/enrollment/history | Get sorted and paged Enrollment history object 
[**V2EnrollmentHistoryPost**](EnrollmentAPI.md#V2EnrollmentHistoryPost) | **Post** /v2/enrollment/history | Add Enrollment history object notes 
[**V2EnrollmentLanguageCodesGet**](EnrollmentAPI.md#V2EnrollmentLanguageCodesGet) | **Get** /v2/enrollment/language-codes | Retrieve the list of languages and corresponding ISO 639-1 Codes 
[**V2EnrollmentLanguagesDeleteMultiplePost**](EnrollmentAPI.md#V2EnrollmentLanguagesDeleteMultiplePost) | **Post** /v2/enrollment/languages/delete-multiple | Delete multiple configured languages from User-Initiated Enrollment settings 
[**V2EnrollmentLanguagesGet**](EnrollmentAPI.md#V2EnrollmentLanguagesGet) | **Get** /v2/enrollment/languages | Get an array of the language codes that have Enrollment messaging 
[**V2EnrollmentLanguagesLanguageIdDelete**](EnrollmentAPI.md#V2EnrollmentLanguagesLanguageIdDelete) | **Delete** /v2/enrollment/languages/{languageId} | Delete the Enrollment messaging for a language 
[**V2EnrollmentLanguagesLanguageIdGet**](EnrollmentAPI.md#V2EnrollmentLanguagesLanguageIdGet) | **Get** /v2/enrollment/languages/{languageId} | Retrieve the Enrollment messaging for a language 
[**V2EnrollmentLanguagesLanguageIdPut**](EnrollmentAPI.md#V2EnrollmentLanguagesLanguageIdPut) | **Put** /v2/enrollment/languages/{languageId} | Edit Enrollment messaging for a language 
[**V2EnrollmentPut**](EnrollmentAPI.md#V2EnrollmentPut) | **Put** /v2/enrollment | Update Enrollment object 
[**V3EnrollmentAccessGroupsGet**](EnrollmentAPI.md#V3EnrollmentAccessGroupsGet) | **Get** /v3/enrollment/access-groups | Retrieve the configured LDAP groups configured for User-Initiated Enrollment. 
[**V3EnrollmentAccessGroupsIdDelete**](EnrollmentAPI.md#V3EnrollmentAccessGroupsIdDelete) | **Delete** /v3/enrollment/access-groups/{id} | Delete an LDAP group&#39;s access to user initiated Enrollment. 
[**V3EnrollmentAccessGroupsIdGet**](EnrollmentAPI.md#V3EnrollmentAccessGroupsIdGet) | **Get** /v3/enrollment/access-groups/{id} | Retrieve the configured LDAP groups configured for User-Initiated Enrollment 
[**V3EnrollmentAccessGroupsIdPut**](EnrollmentAPI.md#V3EnrollmentAccessGroupsIdPut) | **Put** /v3/enrollment/access-groups/{id} | Modify the configured LDAP groups configured for User-Initiated Enrollment. Only exiting Access Groups can be updated. 
[**V3EnrollmentAccessGroupsPost**](EnrollmentAPI.md#V3EnrollmentAccessGroupsPost) | **Post** /v3/enrollment/access-groups | Add the configured LDAP group for User-Initiated Enrollment. 
[**V3EnrollmentFilteredLanguageCodesGet**](EnrollmentAPI.md#V3EnrollmentFilteredLanguageCodesGet) | **Get** /v3/enrollment/filtered-language-codes | Retrieve the list of languages and corresponding ISO 639-1 Codes but only those not already added to Enrollment 
[**V3EnrollmentGet**](EnrollmentAPI.md#V3EnrollmentGet) | **Get** /v3/enrollment | Get Enrollment object and Re-enrollment settings 
[**V3EnrollmentLanguageCodesGet**](EnrollmentAPI.md#V3EnrollmentLanguageCodesGet) | **Get** /v3/enrollment/language-codes | Retrieve the list of languages and corresponding ISO 639-1 Codes 
[**V3EnrollmentLanguagesDeleteMultiplePost**](EnrollmentAPI.md#V3EnrollmentLanguagesDeleteMultiplePost) | **Post** /v3/enrollment/languages/delete-multiple | Delete multiple configured languages from User-Initiated Enrollment settings 
[**V3EnrollmentLanguagesGet**](EnrollmentAPI.md#V3EnrollmentLanguagesGet) | **Get** /v3/enrollment/languages | Get an array of the language codes that have Enrollment messaging 
[**V3EnrollmentLanguagesLanguageIdDelete**](EnrollmentAPI.md#V3EnrollmentLanguagesLanguageIdDelete) | **Delete** /v3/enrollment/languages/{languageId} | Delete the Enrollment messaging for a language 
[**V3EnrollmentLanguagesLanguageIdGet**](EnrollmentAPI.md#V3EnrollmentLanguagesLanguageIdGet) | **Get** /v3/enrollment/languages/{languageId} | Retrieve the Enrollment messaging for a language 
[**V3EnrollmentLanguagesLanguageIdPut**](EnrollmentAPI.md#V3EnrollmentLanguagesLanguageIdPut) | **Put** /v3/enrollment/languages/{languageId} | Edit Enrollment messaging for a language 
[**V3EnrollmentPut**](EnrollmentAPI.md#V3EnrollmentPut) | **Put** /v3/enrollment | Update Enrollment object 
[**V4EnrollmentGet**](EnrollmentAPI.md#V4EnrollmentGet) | **Get** /v4/enrollment | Get Enrollment object and Re-enrollment settings 
[**V4EnrollmentPut**](EnrollmentAPI.md#V4EnrollmentPut) | **Put** /v4/enrollment | Update Enrollment object 



## V1AdueSessionTokenSettingsGet

> AccountDrivenUserEnrollmentSessionTokenSettings V1AdueSessionTokenSettingsGet(ctx).Execute()

Retrieve the Account Driven User Enrollment Session Token Settings 



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
	resp, r, err := apiClient.EnrollmentAPI.V1AdueSessionTokenSettingsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V1AdueSessionTokenSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AdueSessionTokenSettingsGet`: AccountDrivenUserEnrollmentSessionTokenSettings
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V1AdueSessionTokenSettingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AdueSessionTokenSettingsGetRequest struct via the builder pattern


### Return type

[**AccountDrivenUserEnrollmentSessionTokenSettings**](AccountDrivenUserEnrollmentSessionTokenSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AdueSessionTokenSettingsPut

> AccountDrivenUserEnrollmentSessionTokenSettings V1AdueSessionTokenSettingsPut(ctx).AccountDrivenUserEnrollmentSessionTokenSettings(accountDrivenUserEnrollmentSessionTokenSettings).Execute()

Update Account Driven User Enrollment Session Token Settings. 



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
	accountDrivenUserEnrollmentSessionTokenSettings := *openapiclient.NewAccountDrivenUserEnrollmentSessionTokenSettings() // AccountDrivenUserEnrollmentSessionTokenSettings | Update Account Driven User Enrollment Session Token Settings.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V1AdueSessionTokenSettingsPut(context.Background()).AccountDrivenUserEnrollmentSessionTokenSettings(accountDrivenUserEnrollmentSessionTokenSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V1AdueSessionTokenSettingsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AdueSessionTokenSettingsPut`: AccountDrivenUserEnrollmentSessionTokenSettings
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V1AdueSessionTokenSettingsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AdueSessionTokenSettingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountDrivenUserEnrollmentSessionTokenSettings** | [**AccountDrivenUserEnrollmentSessionTokenSettings**](AccountDrivenUserEnrollmentSessionTokenSettings.md) | Update Account Driven User Enrollment Session Token Settings. | 

### Return type

[**AccountDrivenUserEnrollmentSessionTokenSettings**](AccountDrivenUserEnrollmentSessionTokenSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentAccessGroupsGet

> AccessGroupsV2SearchResults V2EnrollmentAccessGroupsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).AllUsersOptionFirst(allUsersOptionFirst).Execute()

Retrieve the configured LDAP groups configured for User-Initiated Enrollment 



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: `property:asc/desc`. Default sort is `name:asc`. Multiple sort criteria are supported and must be separated with a comma. Example: `sort=date:desc,name:asc`.  (optional) (default to ["name:asc"])
	allUsersOptionFirst := true // bool | Return \"All LDAP Users\" option on the first position if it is present in the current page (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V2EnrollmentAccessGroupsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).AllUsersOptionFirst(allUsersOptionFirst).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V2EnrollmentAccessGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentAccessGroupsGet`: AccessGroupsV2SearchResults
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V2EnrollmentAccessGroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentAccessGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: &#x60;property:asc/desc&#x60;. Default sort is &#x60;name:asc&#x60;. Multiple sort criteria are supported and must be separated with a comma. Example: &#x60;sort&#x3D;date:desc,name:asc&#x60;.  | [default to [&quot;name:asc&quot;]]
 **allUsersOptionFirst** | **bool** | Return \&quot;All LDAP Users\&quot; option on the first position if it is present in the current page | [default to false]

### Return type

[**AccessGroupsV2SearchResults**](AccessGroupsV2SearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentAccessGroupsPost

> HrefResponse V2EnrollmentAccessGroupsPost(ctx).EnrollmentAccessGroupV2(enrollmentAccessGroupV2).Execute()

Add the configured LDAP group for User-Initiated Enrollment. 



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
	enrollmentAccessGroupV2 := *openapiclient.NewEnrollmentAccessGroupV2() // EnrollmentAccessGroupV2 | Configured LDAP group to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V2EnrollmentAccessGroupsPost(context.Background()).EnrollmentAccessGroupV2(enrollmentAccessGroupV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V2EnrollmentAccessGroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentAccessGroupsPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V2EnrollmentAccessGroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentAccessGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enrollmentAccessGroupV2** | [**EnrollmentAccessGroupV2**](EnrollmentAccessGroupV2.md) | Configured LDAP group to create. | 

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


## V2EnrollmentAccessGroupsServerIdGroupIdDelete

> V2EnrollmentAccessGroupsServerIdGroupIdDelete(ctx, serverId, groupId).Execute()

Delete an LDAP group's access to user initiated Enrollment 



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
	serverId := "serverId_example" // string | LDAP server id
	groupId := "groupId_example" // string | LDAP group id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnrollmentAPI.V2EnrollmentAccessGroupsServerIdGroupIdDelete(context.Background(), serverId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V2EnrollmentAccessGroupsServerIdGroupIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | LDAP server id | 
**groupId** | **string** | LDAP group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentAccessGroupsServerIdGroupIdDeleteRequest struct via the builder pattern


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


## V2EnrollmentAccessGroupsServerIdGroupIdGet

> EnrollmentAccessGroupV2 V2EnrollmentAccessGroupsServerIdGroupIdGet(ctx, serverId, groupId).Execute()

Retrieve the configured LDAP groups configured for User-Initiated Enrollment 



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
	serverId := "serverId_example" // string | LDAP server id.
	groupId := "groupId_example" // string | LDAP group id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V2EnrollmentAccessGroupsServerIdGroupIdGet(context.Background(), serverId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V2EnrollmentAccessGroupsServerIdGroupIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentAccessGroupsServerIdGroupIdGet`: EnrollmentAccessGroupV2
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V2EnrollmentAccessGroupsServerIdGroupIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | LDAP server id. | 
**groupId** | **string** | LDAP group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentAccessGroupsServerIdGroupIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnrollmentAccessGroupV2**](EnrollmentAccessGroupV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentAccessGroupsServerIdGroupIdPut

> EnrollmentAccessGroupV2 V2EnrollmentAccessGroupsServerIdGroupIdPut(ctx, serverId, groupId).EnrollmentAccessGroupV2(enrollmentAccessGroupV2).Execute()

Modify the configured LDAP groups configured for User-Initiated Enrollment 



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
	serverId := "serverId_example" // string | LDAP server id.
	groupId := "groupId_example" // string | LDAP group id.
	enrollmentAccessGroupV2 := *openapiclient.NewEnrollmentAccessGroupV2() // EnrollmentAccessGroupV2 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V2EnrollmentAccessGroupsServerIdGroupIdPut(context.Background(), serverId, groupId).EnrollmentAccessGroupV2(enrollmentAccessGroupV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V2EnrollmentAccessGroupsServerIdGroupIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentAccessGroupsServerIdGroupIdPut`: EnrollmentAccessGroupV2
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V2EnrollmentAccessGroupsServerIdGroupIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | LDAP server id. | 
**groupId** | **string** | LDAP group id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentAccessGroupsServerIdGroupIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enrollmentAccessGroupV2** | [**EnrollmentAccessGroupV2**](EnrollmentAccessGroupV2.md) |  | 

### Return type

[**EnrollmentAccessGroupV2**](EnrollmentAccessGroupV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentFilteredLanguageCodesGet

> []LanguageCode V2EnrollmentFilteredLanguageCodesGet(ctx).Execute()

Retrieve the list of languages and corresponding ISO 639-1 Codes but only those not already added to Enrollment 



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
	resp, r, err := apiClient.EnrollmentAPI.V2EnrollmentFilteredLanguageCodesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V2EnrollmentFilteredLanguageCodesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentFilteredLanguageCodesGet`: []LanguageCode
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V2EnrollmentFilteredLanguageCodesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentFilteredLanguageCodesGetRequest struct via the builder pattern


### Return type

[**[]LanguageCode**](LanguageCode.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentGet

> EnrollmentSettingsV2 V2EnrollmentGet(ctx).Execute()

Get Enrollment object and Re-enrollment settings 



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
	resp, r, err := apiClient.EnrollmentAPI.V2EnrollmentGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V2EnrollmentGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentGet`: EnrollmentSettingsV2
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V2EnrollmentGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentGetRequest struct via the builder pattern


### Return type

[**EnrollmentSettingsV2**](EnrollmentSettingsV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentHistoryExportPost

> interface{} V2EnrollmentHistoryExportPost(ctx).ExportFields(exportFields).ExportLabels(exportLabels).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).ExportParameters(exportParameters).Execute()

Export enrollment history collection 



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
	exportFields := []string{"Inner_example"} // []string | Export fields parameter, used to change default order or ignore some of the response properties. Default is empty array, which means that all fields of the response entity will be serialized. Example: export-fields=id,username (optional) (default to [])
	exportLabels := []string{"Inner_example"} // []string | Export labels parameter, used to customize fieldnames/columns in the exported file. Default is empty array, which means that response properties names will be used. Number of the provided labels must match the number of export-fields Example: export-labels=identifier,name with matching: export-fields=id,username (optional) (default to [])
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=id:desc,name:asc  (optional) (default to ["id:asc"])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name. This param can be combined with paging and sorting. Example: name==\"*script*\" (optional) (default to "")
	exportParameters := *openapiclient.NewExportParameters() // ExportParameters | Optional. Override query parameters since they can make URI exceed 2,000 character limit. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V2EnrollmentHistoryExportPost(context.Background()).ExportFields(exportFields).ExportLabels(exportLabels).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).ExportParameters(exportParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V2EnrollmentHistoryExportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentHistoryExportPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V2EnrollmentHistoryExportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentHistoryExportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportFields** | **[]string** | Export fields parameter, used to change default order or ignore some of the response properties. Default is empty array, which means that all fields of the response entity will be serialized. Example: export-fields&#x3D;id,username | [default to []]
 **exportLabels** | **[]string** | Export labels parameter, used to customize fieldnames/columns in the exported file. Default is empty array, which means that response properties names will be used. Number of the provided labels must match the number of export-fields Example: export-labels&#x3D;identifier,name with matching: export-fields&#x3D;id,username | [default to []]
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;id:desc,name:asc  | [default to [&quot;id:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name. This param can be combined with paging and sorting. Example: name&#x3D;&#x3D;\&quot;*script*\&quot; | [default to &quot;&quot;]
 **exportParameters** | [**ExportParameters**](ExportParameters.md) | Optional. Override query parameters since they can make URI exceed 2,000 character limit. | 

### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/csv,application/json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentHistoryGet

> HistorySearchResults V2EnrollmentHistoryGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Execute()

Get sorted and paged Enrollment history object 



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: `property:asc/desc`. Default sort is `date:desc`. Multiple sort criteria are supported and must be separated with a comma. Example: `sort=date:desc,name:asc`.  (optional) (default to ["date:desc"])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V2EnrollmentHistoryGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V2EnrollmentHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V2EnrollmentHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: &#x60;property:asc/desc&#x60;. Default sort is &#x60;date:desc&#x60;. Multiple sort criteria are supported and must be separated with a comma. Example: &#x60;sort&#x3D;date:desc,name:asc&#x60;.  | [default to [&quot;date:desc&quot;]]

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


## V2EnrollmentHistoryPost

> HrefResponse V2EnrollmentHistoryPost(ctx).ObjectHistoryNote(objectHistoryNote).Execute()

Add Enrollment history object notes 



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
	objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | history notes to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V2EnrollmentHistoryPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V2EnrollmentHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentHistoryPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V2EnrollmentHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectHistoryNote** | [**ObjectHistoryNote**](ObjectHistoryNote.md) | history notes to create | 

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


## V2EnrollmentLanguageCodesGet

> []LanguageCode V2EnrollmentLanguageCodesGet(ctx).Execute()

Retrieve the list of languages and corresponding ISO 639-1 Codes 



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
	resp, r, err := apiClient.EnrollmentAPI.V2EnrollmentLanguageCodesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V2EnrollmentLanguageCodesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentLanguageCodesGet`: []LanguageCode
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V2EnrollmentLanguageCodesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentLanguageCodesGetRequest struct via the builder pattern


### Return type

[**[]LanguageCode**](LanguageCode.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentLanguagesDeleteMultiplePost

> V2EnrollmentLanguagesDeleteMultiplePost(ctx).Ids(ids).Execute()

Delete multiple configured languages from User-Initiated Enrollment settings 



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
	ids := *openapiclient.NewIds() // Ids | ids of each language to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnrollmentAPI.V2EnrollmentLanguagesDeleteMultiplePost(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V2EnrollmentLanguagesDeleteMultiplePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentLanguagesDeleteMultiplePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**Ids**](Ids.md) | ids of each language to delete | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentLanguagesGet

> ProcessTextsSearchResults V2EnrollmentLanguagesGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Execute()

Get an array of the language codes that have Enrollment messaging 



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is `languageCode:asc`. Multiple sort criteria are supported and must be separated with a comma. Example: `sort=date:desc,name:asc`.  (optional) (default to ["languageCode:asc"])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V2EnrollmentLanguagesGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V2EnrollmentLanguagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentLanguagesGet`: ProcessTextsSearchResults
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V2EnrollmentLanguagesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentLanguagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is &#x60;languageCode:asc&#x60;. Multiple sort criteria are supported and must be separated with a comma. Example: &#x60;sort&#x3D;date:desc,name:asc&#x60;.  | [default to [&quot;languageCode:asc&quot;]]

### Return type

[**ProcessTextsSearchResults**](ProcessTextsSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentLanguagesLanguageIdDelete

> V2EnrollmentLanguagesLanguageIdDelete(ctx, languageId).Execute()

Delete the Enrollment messaging for a language 



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
	languageId := "languageId_example" // string | Two letter ISO 639-1 Language Code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnrollmentAPI.V2EnrollmentLanguagesLanguageIdDelete(context.Background(), languageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V2EnrollmentLanguagesLanguageIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**languageId** | **string** | Two letter ISO 639-1 Language Code | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentLanguagesLanguageIdDeleteRequest struct via the builder pattern


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


## V2EnrollmentLanguagesLanguageIdGet

> EnrollmentProcessTextObject V2EnrollmentLanguagesLanguageIdGet(ctx, languageId).Execute()

Retrieve the Enrollment messaging for a language 



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
	languageId := "languageId_example" // string | Two letter ISO 639-1 Language Code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V2EnrollmentLanguagesLanguageIdGet(context.Background(), languageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V2EnrollmentLanguagesLanguageIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentLanguagesLanguageIdGet`: EnrollmentProcessTextObject
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V2EnrollmentLanguagesLanguageIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**languageId** | **string** | Two letter ISO 639-1 Language Code | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentLanguagesLanguageIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnrollmentProcessTextObject**](EnrollmentProcessTextObject.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentLanguagesLanguageIdPut

> EnrollmentProcessTextObject V2EnrollmentLanguagesLanguageIdPut(ctx, languageId).EnrollmentProcessTextObject(enrollmentProcessTextObject).Execute()

Edit Enrollment messaging for a language 



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
	languageId := "languageId_example" // string | Two letter ISO 639-1 Language Code
	enrollmentProcessTextObject := *openapiclient.NewEnrollmentProcessTextObject() // EnrollmentProcessTextObject |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V2EnrollmentLanguagesLanguageIdPut(context.Background(), languageId).EnrollmentProcessTextObject(enrollmentProcessTextObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V2EnrollmentLanguagesLanguageIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentLanguagesLanguageIdPut`: EnrollmentProcessTextObject
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V2EnrollmentLanguagesLanguageIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**languageId** | **string** | Two letter ISO 639-1 Language Code | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentLanguagesLanguageIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enrollmentProcessTextObject** | [**EnrollmentProcessTextObject**](EnrollmentProcessTextObject.md) |  | 

### Return type

[**EnrollmentProcessTextObject**](EnrollmentProcessTextObject.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2EnrollmentPut

> EnrollmentSettingsV2 V2EnrollmentPut(ctx).EnrollmentSettingsV2(enrollmentSettingsV2).Execute()

Update Enrollment object 



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
	enrollmentSettingsV2 := *openapiclient.NewEnrollmentSettingsV2("radmin") // EnrollmentSettingsV2 | Update enrollment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V2EnrollmentPut(context.Background()).EnrollmentSettingsV2(enrollmentSettingsV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V2EnrollmentPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2EnrollmentPut`: EnrollmentSettingsV2
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V2EnrollmentPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2EnrollmentPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enrollmentSettingsV2** | [**EnrollmentSettingsV2**](EnrollmentSettingsV2.md) | Update enrollment | 

### Return type

[**EnrollmentSettingsV2**](EnrollmentSettingsV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3EnrollmentAccessGroupsGet

> AccessGroupsPreviewSearchResults V3EnrollmentAccessGroupsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).AllUsersOptionFirst(allUsersOptionFirst).Execute()

Retrieve the configured LDAP groups configured for User-Initiated Enrollment. 



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: `property:asc/desc`. Default sort is `name:asc`. Multiple sort criteria are supported and must be separated with a comma. Example: `sort=date:desc,name:asc`.  (optional) (default to ["name:asc"])
	allUsersOptionFirst := true // bool | Return \"All LDAP Users\" option on the first position if it is present in the current page (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V3EnrollmentAccessGroupsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).AllUsersOptionFirst(allUsersOptionFirst).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V3EnrollmentAccessGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3EnrollmentAccessGroupsGet`: AccessGroupsPreviewSearchResults
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V3EnrollmentAccessGroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3EnrollmentAccessGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: &#x60;property:asc/desc&#x60;. Default sort is &#x60;name:asc&#x60;. Multiple sort criteria are supported and must be separated with a comma. Example: &#x60;sort&#x3D;date:desc,name:asc&#x60;.  | [default to [&quot;name:asc&quot;]]
 **allUsersOptionFirst** | **bool** | Return \&quot;All LDAP Users\&quot; option on the first position if it is present in the current page | [default to false]

### Return type

[**AccessGroupsPreviewSearchResults**](AccessGroupsPreviewSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3EnrollmentAccessGroupsIdDelete

> V3EnrollmentAccessGroupsIdDelete(ctx, id).Execute()

Delete an LDAP group's access to user initiated Enrollment. 



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
	id := "id_example" // string | Autogenerated Access Group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnrollmentAPI.V3EnrollmentAccessGroupsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V3EnrollmentAccessGroupsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autogenerated Access Group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3EnrollmentAccessGroupsIdDeleteRequest struct via the builder pattern


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


## V3EnrollmentAccessGroupsIdGet

> EnrollmentAccessGroupPreview V3EnrollmentAccessGroupsIdGet(ctx, id).Execute()

Retrieve the configured LDAP groups configured for User-Initiated Enrollment 



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
	id := "id_example" // string | Autogenerated Access Group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V3EnrollmentAccessGroupsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V3EnrollmentAccessGroupsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3EnrollmentAccessGroupsIdGet`: EnrollmentAccessGroupPreview
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V3EnrollmentAccessGroupsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autogenerated Access Group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3EnrollmentAccessGroupsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnrollmentAccessGroupPreview**](EnrollmentAccessGroupPreview.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3EnrollmentAccessGroupsIdPut

> EnrollmentAccessGroupPreview V3EnrollmentAccessGroupsIdPut(ctx, id).EnrollmentAccessGroupPreview(enrollmentAccessGroupPreview).Execute()

Modify the configured LDAP groups configured for User-Initiated Enrollment. Only exiting Access Groups can be updated. 



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
	id := "id_example" // string | Autogenerated Access Group ID.
	enrollmentAccessGroupPreview := *openapiclient.NewEnrollmentAccessGroupPreview("1", "1", "Grade School Pupils") // EnrollmentAccessGroupPreview | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V3EnrollmentAccessGroupsIdPut(context.Background(), id).EnrollmentAccessGroupPreview(enrollmentAccessGroupPreview).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V3EnrollmentAccessGroupsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3EnrollmentAccessGroupsIdPut`: EnrollmentAccessGroupPreview
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V3EnrollmentAccessGroupsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autogenerated Access Group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3EnrollmentAccessGroupsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enrollmentAccessGroupPreview** | [**EnrollmentAccessGroupPreview**](EnrollmentAccessGroupPreview.md) |  | 

### Return type

[**EnrollmentAccessGroupPreview**](EnrollmentAccessGroupPreview.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3EnrollmentAccessGroupsPost

> HrefResponse V3EnrollmentAccessGroupsPost(ctx).EnrollmentAccessGroupPreview(enrollmentAccessGroupPreview).Execute()

Add the configured LDAP group for User-Initiated Enrollment. 



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
	enrollmentAccessGroupPreview := *openapiclient.NewEnrollmentAccessGroupPreview("1", "1", "Grade School Pupils") // EnrollmentAccessGroupPreview | Configured LDAP group to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V3EnrollmentAccessGroupsPost(context.Background()).EnrollmentAccessGroupPreview(enrollmentAccessGroupPreview).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V3EnrollmentAccessGroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3EnrollmentAccessGroupsPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V3EnrollmentAccessGroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3EnrollmentAccessGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enrollmentAccessGroupPreview** | [**EnrollmentAccessGroupPreview**](EnrollmentAccessGroupPreview.md) | Configured LDAP group to create. | 

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


## V3EnrollmentFilteredLanguageCodesGet

> []LanguageCode V3EnrollmentFilteredLanguageCodesGet(ctx).Execute()

Retrieve the list of languages and corresponding ISO 639-1 Codes but only those not already added to Enrollment 



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
	resp, r, err := apiClient.EnrollmentAPI.V3EnrollmentFilteredLanguageCodesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V3EnrollmentFilteredLanguageCodesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3EnrollmentFilteredLanguageCodesGet`: []LanguageCode
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V3EnrollmentFilteredLanguageCodesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV3EnrollmentFilteredLanguageCodesGetRequest struct via the builder pattern


### Return type

[**[]LanguageCode**](LanguageCode.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3EnrollmentGet

> EnrollmentSettingsV3 V3EnrollmentGet(ctx).Execute()

Get Enrollment object and Re-enrollment settings 



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
	resp, r, err := apiClient.EnrollmentAPI.V3EnrollmentGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V3EnrollmentGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3EnrollmentGet`: EnrollmentSettingsV3
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V3EnrollmentGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV3EnrollmentGetRequest struct via the builder pattern


### Return type

[**EnrollmentSettingsV3**](EnrollmentSettingsV3.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3EnrollmentLanguageCodesGet

> []LanguageCode V3EnrollmentLanguageCodesGet(ctx).Execute()

Retrieve the list of languages and corresponding ISO 639-1 Codes 



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
	resp, r, err := apiClient.EnrollmentAPI.V3EnrollmentLanguageCodesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V3EnrollmentLanguageCodesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3EnrollmentLanguageCodesGet`: []LanguageCode
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V3EnrollmentLanguageCodesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV3EnrollmentLanguageCodesGetRequest struct via the builder pattern


### Return type

[**[]LanguageCode**](LanguageCode.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3EnrollmentLanguagesDeleteMultiplePost

> V3EnrollmentLanguagesDeleteMultiplePost(ctx).Ids(ids).Execute()

Delete multiple configured languages from User-Initiated Enrollment settings 



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
	ids := *openapiclient.NewIds() // Ids | ids of each language to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnrollmentAPI.V3EnrollmentLanguagesDeleteMultiplePost(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V3EnrollmentLanguagesDeleteMultiplePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3EnrollmentLanguagesDeleteMultiplePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**Ids**](Ids.md) | ids of each language to delete | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3EnrollmentLanguagesGet

> ProcessTextsSearchResults V3EnrollmentLanguagesGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Execute()

Get an array of the language codes that have Enrollment messaging 



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is `languageCode:asc`. Multiple sort criteria are supported and must be separated with a comma. Example: `sort=date:desc,name:asc`.  (optional) (default to ["languageCode:asc"])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V3EnrollmentLanguagesGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V3EnrollmentLanguagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3EnrollmentLanguagesGet`: ProcessTextsSearchResults
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V3EnrollmentLanguagesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3EnrollmentLanguagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is &#x60;languageCode:asc&#x60;. Multiple sort criteria are supported and must be separated with a comma. Example: &#x60;sort&#x3D;date:desc,name:asc&#x60;.  | [default to [&quot;languageCode:asc&quot;]]

### Return type

[**ProcessTextsSearchResults**](ProcessTextsSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3EnrollmentLanguagesLanguageIdDelete

> V3EnrollmentLanguagesLanguageIdDelete(ctx, languageId).Execute()

Delete the Enrollment messaging for a language 



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
	languageId := "languageId_example" // string | Two letter ISO 639-1 Language Code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnrollmentAPI.V3EnrollmentLanguagesLanguageIdDelete(context.Background(), languageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V3EnrollmentLanguagesLanguageIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**languageId** | **string** | Two letter ISO 639-1 Language Code | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3EnrollmentLanguagesLanguageIdDeleteRequest struct via the builder pattern


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


## V3EnrollmentLanguagesLanguageIdGet

> EnrollmentProcessTextObject V3EnrollmentLanguagesLanguageIdGet(ctx, languageId).Execute()

Retrieve the Enrollment messaging for a language 



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
	languageId := "languageId_example" // string | Two letter ISO 639-1 Language Code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V3EnrollmentLanguagesLanguageIdGet(context.Background(), languageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V3EnrollmentLanguagesLanguageIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3EnrollmentLanguagesLanguageIdGet`: EnrollmentProcessTextObject
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V3EnrollmentLanguagesLanguageIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**languageId** | **string** | Two letter ISO 639-1 Language Code | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3EnrollmentLanguagesLanguageIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnrollmentProcessTextObject**](EnrollmentProcessTextObject.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3EnrollmentLanguagesLanguageIdPut

> EnrollmentProcessTextObject V3EnrollmentLanguagesLanguageIdPut(ctx, languageId).EnrollmentProcessTextObject(enrollmentProcessTextObject).Execute()

Edit Enrollment messaging for a language 



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
	languageId := "languageId_example" // string | Two letter ISO 639-1 Language Code
	enrollmentProcessTextObject := *openapiclient.NewEnrollmentProcessTextObject() // EnrollmentProcessTextObject |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V3EnrollmentLanguagesLanguageIdPut(context.Background(), languageId).EnrollmentProcessTextObject(enrollmentProcessTextObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V3EnrollmentLanguagesLanguageIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3EnrollmentLanguagesLanguageIdPut`: EnrollmentProcessTextObject
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V3EnrollmentLanguagesLanguageIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**languageId** | **string** | Two letter ISO 639-1 Language Code | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3EnrollmentLanguagesLanguageIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enrollmentProcessTextObject** | [**EnrollmentProcessTextObject**](EnrollmentProcessTextObject.md) |  | 

### Return type

[**EnrollmentProcessTextObject**](EnrollmentProcessTextObject.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3EnrollmentPut

> EnrollmentSettingsV3 V3EnrollmentPut(ctx).EnrollmentSettingsV3(enrollmentSettingsV3).Execute()

Update Enrollment object 



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
	enrollmentSettingsV3 := *openapiclient.NewEnrollmentSettingsV3("radmin") // EnrollmentSettingsV3 | Update enrollment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V3EnrollmentPut(context.Background()).EnrollmentSettingsV3(enrollmentSettingsV3).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V3EnrollmentPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3EnrollmentPut`: EnrollmentSettingsV3
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V3EnrollmentPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3EnrollmentPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enrollmentSettingsV3** | [**EnrollmentSettingsV3**](EnrollmentSettingsV3.md) | Update enrollment | 

### Return type

[**EnrollmentSettingsV3**](EnrollmentSettingsV3.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4EnrollmentGet

> EnrollmentSettingsV4 V4EnrollmentGet(ctx).Execute()

Get Enrollment object and Re-enrollment settings 



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
	resp, r, err := apiClient.EnrollmentAPI.V4EnrollmentGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V4EnrollmentGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V4EnrollmentGet`: EnrollmentSettingsV4
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V4EnrollmentGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV4EnrollmentGetRequest struct via the builder pattern


### Return type

[**EnrollmentSettingsV4**](EnrollmentSettingsV4.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4EnrollmentPut

> EnrollmentSettingsV4 V4EnrollmentPut(ctx).EnrollmentSettingsV4(enrollmentSettingsV4).Execute()

Update Enrollment object 



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
	enrollmentSettingsV4 := *openapiclient.NewEnrollmentSettingsV4("radmin") // EnrollmentSettingsV4 | Update enrollment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V4EnrollmentPut(context.Background()).EnrollmentSettingsV4(enrollmentSettingsV4).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V4EnrollmentPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V4EnrollmentPut`: EnrollmentSettingsV4
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V4EnrollmentPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV4EnrollmentPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enrollmentSettingsV4** | [**EnrollmentSettingsV4**](EnrollmentSettingsV4.md) | Update enrollment | 

### Return type

[**EnrollmentSettingsV4**](EnrollmentSettingsV4.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

