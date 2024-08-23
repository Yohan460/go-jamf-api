# \MobileDeviceGroupsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1MobileDeviceGroupsGet**](MobileDeviceGroupsAPI.md#V1MobileDeviceGroupsGet) | **Get** /v1/mobile-device-groups | Return the list of all Mobile Device Groups 
[**V1MobileDeviceGroupsStaticGroupMembershipIdGet**](MobileDeviceGroupsAPI.md#V1MobileDeviceGroupsStaticGroupMembershipIdGet) | **Get** /v1/mobile-device-groups/static-group-membership/{id} | Get Static Group Membership by Id 
[**V1MobileDeviceGroupsStaticGroupsGet**](MobileDeviceGroupsAPI.md#V1MobileDeviceGroupsStaticGroupsGet) | **Get** /v1/mobile-device-groups/static-groups | Get Static Groups 
[**V1MobileDeviceGroupsStaticGroupsIdDelete**](MobileDeviceGroupsAPI.md#V1MobileDeviceGroupsStaticGroupsIdDelete) | **Delete** /v1/mobile-device-groups/static-groups/{id} | Remove Static Group by Id 
[**V1MobileDeviceGroupsStaticGroupsIdGet**](MobileDeviceGroupsAPI.md#V1MobileDeviceGroupsStaticGroupsIdGet) | **Get** /v1/mobile-device-groups/static-groups/{id} | Get Static Group by Id 
[**V1MobileDeviceGroupsStaticGroupsIdPatch**](MobileDeviceGroupsAPI.md#V1MobileDeviceGroupsStaticGroupsIdPatch) | **Patch** /v1/mobile-device-groups/static-groups/{id} | Update membership of a static group. 
[**V1MobileDeviceGroupsStaticGroupsPost**](MobileDeviceGroupsAPI.md#V1MobileDeviceGroupsStaticGroupsPost) | **Post** /v1/mobile-device-groups/static-groups | Create membership of a static group. 



## V1MobileDeviceGroupsGet

> []MobileDeviceGroup V1MobileDeviceGroupsGet(ctx).Execute()

Return the list of all Mobile Device Groups 



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
	resp, r, err := apiClient.MobileDeviceGroupsAPI.V1MobileDeviceGroupsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileDeviceGroupsAPI.V1MobileDeviceGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MobileDeviceGroupsGet`: []MobileDeviceGroup
	fmt.Fprintf(os.Stdout, "Response from `MobileDeviceGroupsAPI.V1MobileDeviceGroupsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDeviceGroupsGetRequest struct via the builder pattern


### Return type

[**[]MobileDeviceGroup**](MobileDeviceGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MobileDeviceGroupsStaticGroupMembershipIdGet

> InventoryListMobileDeviceSearchResults V1MobileDeviceGroupsStaticGroupMembershipIdGet(ctx, id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get Static Group Membership by Id 



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
	id := "id_example" // string | instance id of static-group
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is mobileDeviceId:asc. Multiple sort criteria are supported and must be separated with a comma.   Fields allowed in the sort: `airPlayPassword`, `appAnalyticsEnabled`, `assetTag`, `availableSpaceMb`,  `batteryLevel`, `bluetoothLowEnergyCapable`, `bluetoothMacAddress`, `capacityMb`,  `lostModeEnabledDate`, `declarativeDeviceManagementEnabled`, `deviceId`, `deviceLocatorServiceEnabled`, `devicePhoneNumber`, `diagnosticAndUsageReportingEnabled`, `displayName`, `doNotDisturbEnabled`,  `enrollmentSessionTokenValid`, `exchangeDeviceId`, `cloudBackupEnabled`, `osBuild`, `osRapidSecurityResponse`, `osSupplementalBuildVersion`, `osVersion`, `ipAddress`, `itunesStoreAccountActive`, `mobileDeviceId`, `languages`, `lastBackupDate`, `lastEnrolledDate`, `lastCloudBackupDate`, `lastInventoryUpdateDate`, `locales`, `locationServicesForSelfServiceMobileEnabled`, `lostModeEnabled`, `managed`, `mdmProfileExpirationDate`, `model`, `modelIdentifier`, `modelNumber`, `modemFirmwareVersion`, `quotaSize`, `residentUsers`, `serialNumber`, `sharedIpad`, `supervised`, `tethered`, `timeZone`, `udid`, `usedSpacePercentage`, `wifiMacAddress`, `deviceOwnershipType`, `building`, `department`, `emailAddress`, `fullName`, `userPhoneNumber`, `position`, `room`, `username`, `appleCareId`, `leaseExpirationDate`,`lifeExpectancyYears`, `poDate`, `poNumber`, `purchasePrice`, `purchasedOrLeased`, `purchasingAccount`, `purchasingContact`, `vendor`, `warrantyExpirationDate`, `activationLockEnabled`, `blockEncryptionCapable`, `dataProtection`, `fileEncryptionCapable`, `hardwareEncryptionSupported`, `jailbreakStatus`, `passcodeCompliant`, `passcodeCompliantWithProfile`, `passcodeLockGracePeriodEnforcedSeconds`, `passcodePresent`, `personalDeviceProfileCurrent`, `carrierSettingsVersion`, `cellularTechnology`, `currentCarrierNetwork`, `currentMobileCountryCode`, `currentMobileNetworkCode`,  `dataRoamingEnabled`, `eid`, `network`, `homeMobileCountryCode`,  `homeMobileNetworkCode`, `iccid`, `imei`, `imei2`, `meid`, `personalHotspotEnabled`, `voiceRoamingEnabled`, `roaming`  Example: `sort=displayName:desc,username:asc`  (optional) (default to ["displayName:asc"])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter mobile device collection. Default filter is empty query - returning all results for the requested page.  Fields allowed in the query: `airPlayPassword`, `appAnalyticsEnabled`, `assetTag`, `availableSpaceMb`,  `batteryLevel`, `bluetoothLowEnergyCapable`, `bluetoothMacAddress`, `capacityMb`,  `declarativeDeviceManagementEnabled`, `deviceId`, `deviceLocatorServiceEnabled`, `devicePhoneNumber`, `diagnosticAndUsageReportingEnabled`, `displayName`, `doNotDisturbEnabled`, `exchangeDeviceId`,  `cloudBackupEnabled`, `osBuild`, `osSupplementalBuildVersion`, `osVersion`, `osRapidSecurityResponse`, `ipAddress`,  `itunesStoreAccountActive`, `mobileDeviceId`, `languages`, `lastInventoryUpdateDate`, `locales`, `locationServicesForSelfServiceMobileEnabled`, `lostModeEnabled`, `managed`, `model`,  `modelIdentifier`, `modelNumber`, `modemFirmwareVersion`, `quotaSize`,  `residentUsers`, `serialNumber`, `sharedIpad`, `supervised`, `tethered`, `timeZone`, `udid`, `usedSpacePercentage`,  `wifiMacAddress`, `building`, `department`, `emailAddress`, `fullName`, `userPhoneNumber`, `position`, `room`, `username`, `appleCareId`, `lifeExpectancyYears`, `poNumber`,  `purchasePrice`, `purchasedOrLeased`, `purchasingAccount`, `purchasingContact`, `vendor`, `activationLockEnabled`, `blockEncryptionCapable`, `dataProtection`,  `fileEncryptionCapable`, `passcodeCompliant`, `passcodeCompliantWithProfile`, `passcodeLockGracePeriodEnforcedSeconds`, `passcodePresent`, `personalDeviceProfileCurrent`, `carrierSettingsVersion`, `currentCarrierNetwork`, `currentMobileCountryCode`, `currentMobileNetworkCode`, `dataRoamingEnabled`, `eid`, `network`, `homeMobileCountryCode`, `homeMobileNetworkCode`, `iccid`, `imei`, `imei2`, `meid`, `personalHotspotEnabled`,  `roaming`  This param can be combined with paging and sorting. Example: `filter=displayName==\"iPad\"`  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupMembershipIdGet(context.Background(), id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupMembershipIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MobileDeviceGroupsStaticGroupMembershipIdGet`: InventoryListMobileDeviceSearchResults
	fmt.Fprintf(os.Stdout, "Response from `MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupMembershipIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of static-group | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDeviceGroupsStaticGroupMembershipIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is mobileDeviceId:asc. Multiple sort criteria are supported and must be separated with a comma.   Fields allowed in the sort: &#x60;airPlayPassword&#x60;, &#x60;appAnalyticsEnabled&#x60;, &#x60;assetTag&#x60;, &#x60;availableSpaceMb&#x60;,  &#x60;batteryLevel&#x60;, &#x60;bluetoothLowEnergyCapable&#x60;, &#x60;bluetoothMacAddress&#x60;, &#x60;capacityMb&#x60;,  &#x60;lostModeEnabledDate&#x60;, &#x60;declarativeDeviceManagementEnabled&#x60;, &#x60;deviceId&#x60;, &#x60;deviceLocatorServiceEnabled&#x60;, &#x60;devicePhoneNumber&#x60;, &#x60;diagnosticAndUsageReportingEnabled&#x60;, &#x60;displayName&#x60;, &#x60;doNotDisturbEnabled&#x60;,  &#x60;enrollmentSessionTokenValid&#x60;, &#x60;exchangeDeviceId&#x60;, &#x60;cloudBackupEnabled&#x60;, &#x60;osBuild&#x60;, &#x60;osRapidSecurityResponse&#x60;, &#x60;osSupplementalBuildVersion&#x60;, &#x60;osVersion&#x60;, &#x60;ipAddress&#x60;, &#x60;itunesStoreAccountActive&#x60;, &#x60;mobileDeviceId&#x60;, &#x60;languages&#x60;, &#x60;lastBackupDate&#x60;, &#x60;lastEnrolledDate&#x60;, &#x60;lastCloudBackupDate&#x60;, &#x60;lastInventoryUpdateDate&#x60;, &#x60;locales&#x60;, &#x60;locationServicesForSelfServiceMobileEnabled&#x60;, &#x60;lostModeEnabled&#x60;, &#x60;managed&#x60;, &#x60;mdmProfileExpirationDate&#x60;, &#x60;model&#x60;, &#x60;modelIdentifier&#x60;, &#x60;modelNumber&#x60;, &#x60;modemFirmwareVersion&#x60;, &#x60;quotaSize&#x60;, &#x60;residentUsers&#x60;, &#x60;serialNumber&#x60;, &#x60;sharedIpad&#x60;, &#x60;supervised&#x60;, &#x60;tethered&#x60;, &#x60;timeZone&#x60;, &#x60;udid&#x60;, &#x60;usedSpacePercentage&#x60;, &#x60;wifiMacAddress&#x60;, &#x60;deviceOwnershipType&#x60;, &#x60;building&#x60;, &#x60;department&#x60;, &#x60;emailAddress&#x60;, &#x60;fullName&#x60;, &#x60;userPhoneNumber&#x60;, &#x60;position&#x60;, &#x60;room&#x60;, &#x60;username&#x60;, &#x60;appleCareId&#x60;, &#x60;leaseExpirationDate&#x60;,&#x60;lifeExpectancyYears&#x60;, &#x60;poDate&#x60;, &#x60;poNumber&#x60;, &#x60;purchasePrice&#x60;, &#x60;purchasedOrLeased&#x60;, &#x60;purchasingAccount&#x60;, &#x60;purchasingContact&#x60;, &#x60;vendor&#x60;, &#x60;warrantyExpirationDate&#x60;, &#x60;activationLockEnabled&#x60;, &#x60;blockEncryptionCapable&#x60;, &#x60;dataProtection&#x60;, &#x60;fileEncryptionCapable&#x60;, &#x60;hardwareEncryptionSupported&#x60;, &#x60;jailbreakStatus&#x60;, &#x60;passcodeCompliant&#x60;, &#x60;passcodeCompliantWithProfile&#x60;, &#x60;passcodeLockGracePeriodEnforcedSeconds&#x60;, &#x60;passcodePresent&#x60;, &#x60;personalDeviceProfileCurrent&#x60;, &#x60;carrierSettingsVersion&#x60;, &#x60;cellularTechnology&#x60;, &#x60;currentCarrierNetwork&#x60;, &#x60;currentMobileCountryCode&#x60;, &#x60;currentMobileNetworkCode&#x60;,  &#x60;dataRoamingEnabled&#x60;, &#x60;eid&#x60;, &#x60;network&#x60;, &#x60;homeMobileCountryCode&#x60;,  &#x60;homeMobileNetworkCode&#x60;, &#x60;iccid&#x60;, &#x60;imei&#x60;, &#x60;imei2&#x60;, &#x60;meid&#x60;, &#x60;personalHotspotEnabled&#x60;, &#x60;voiceRoamingEnabled&#x60;, &#x60;roaming&#x60;  Example: &#x60;sort&#x3D;displayName:desc,username:asc&#x60;  | [default to [&quot;displayName:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter mobile device collection. Default filter is empty query - returning all results for the requested page.  Fields allowed in the query: &#x60;airPlayPassword&#x60;, &#x60;appAnalyticsEnabled&#x60;, &#x60;assetTag&#x60;, &#x60;availableSpaceMb&#x60;,  &#x60;batteryLevel&#x60;, &#x60;bluetoothLowEnergyCapable&#x60;, &#x60;bluetoothMacAddress&#x60;, &#x60;capacityMb&#x60;,  &#x60;declarativeDeviceManagementEnabled&#x60;, &#x60;deviceId&#x60;, &#x60;deviceLocatorServiceEnabled&#x60;, &#x60;devicePhoneNumber&#x60;, &#x60;diagnosticAndUsageReportingEnabled&#x60;, &#x60;displayName&#x60;, &#x60;doNotDisturbEnabled&#x60;, &#x60;exchangeDeviceId&#x60;,  &#x60;cloudBackupEnabled&#x60;, &#x60;osBuild&#x60;, &#x60;osSupplementalBuildVersion&#x60;, &#x60;osVersion&#x60;, &#x60;osRapidSecurityResponse&#x60;, &#x60;ipAddress&#x60;,  &#x60;itunesStoreAccountActive&#x60;, &#x60;mobileDeviceId&#x60;, &#x60;languages&#x60;, &#x60;lastInventoryUpdateDate&#x60;, &#x60;locales&#x60;, &#x60;locationServicesForSelfServiceMobileEnabled&#x60;, &#x60;lostModeEnabled&#x60;, &#x60;managed&#x60;, &#x60;model&#x60;,  &#x60;modelIdentifier&#x60;, &#x60;modelNumber&#x60;, &#x60;modemFirmwareVersion&#x60;, &#x60;quotaSize&#x60;,  &#x60;residentUsers&#x60;, &#x60;serialNumber&#x60;, &#x60;sharedIpad&#x60;, &#x60;supervised&#x60;, &#x60;tethered&#x60;, &#x60;timeZone&#x60;, &#x60;udid&#x60;, &#x60;usedSpacePercentage&#x60;,  &#x60;wifiMacAddress&#x60;, &#x60;building&#x60;, &#x60;department&#x60;, &#x60;emailAddress&#x60;, &#x60;fullName&#x60;, &#x60;userPhoneNumber&#x60;, &#x60;position&#x60;, &#x60;room&#x60;, &#x60;username&#x60;, &#x60;appleCareId&#x60;, &#x60;lifeExpectancyYears&#x60;, &#x60;poNumber&#x60;,  &#x60;purchasePrice&#x60;, &#x60;purchasedOrLeased&#x60;, &#x60;purchasingAccount&#x60;, &#x60;purchasingContact&#x60;, &#x60;vendor&#x60;, &#x60;activationLockEnabled&#x60;, &#x60;blockEncryptionCapable&#x60;, &#x60;dataProtection&#x60;,  &#x60;fileEncryptionCapable&#x60;, &#x60;passcodeCompliant&#x60;, &#x60;passcodeCompliantWithProfile&#x60;, &#x60;passcodeLockGracePeriodEnforcedSeconds&#x60;, &#x60;passcodePresent&#x60;, &#x60;personalDeviceProfileCurrent&#x60;, &#x60;carrierSettingsVersion&#x60;, &#x60;currentCarrierNetwork&#x60;, &#x60;currentMobileCountryCode&#x60;, &#x60;currentMobileNetworkCode&#x60;, &#x60;dataRoamingEnabled&#x60;, &#x60;eid&#x60;, &#x60;network&#x60;, &#x60;homeMobileCountryCode&#x60;, &#x60;homeMobileNetworkCode&#x60;, &#x60;iccid&#x60;, &#x60;imei&#x60;, &#x60;imei2&#x60;, &#x60;meid&#x60;, &#x60;personalHotspotEnabled&#x60;,  &#x60;roaming&#x60;  This param can be combined with paging and sorting. Example: &#x60;filter&#x3D;displayName&#x3D;&#x3D;\&quot;iPad\&quot;&#x60;  | [default to &quot;&quot;]

### Return type

[**InventoryListMobileDeviceSearchResults**](InventoryListMobileDeviceSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MobileDeviceGroupsStaticGroupsGet

> StaticGroupSearchResults V1MobileDeviceGroupsStaticGroupsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get Static Groups 



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Available criteria to sort on: groupId, groupName, siteId. (optional) (default to ["groupId:asc"])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter department collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: groupId, groupName, siteId. The siteId field can only be filtered by admins with full access. Any sited admin will have siteId filtered automatically. This param can be combined with paging and sorting. Example: groupName==\"staticGroup1\" (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MobileDeviceGroupsStaticGroupsGet`: StaticGroupSearchResults
	fmt.Fprintf(os.Stdout, "Response from `MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDeviceGroupsStaticGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Available criteria to sort on: groupId, groupName, siteId. | [default to [&quot;groupId:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter department collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: groupId, groupName, siteId. The siteId field can only be filtered by admins with full access. Any sited admin will have siteId filtered automatically. This param can be combined with paging and sorting. Example: groupName&#x3D;&#x3D;\&quot;staticGroup1\&quot; | [default to &quot;&quot;]

### Return type

[**StaticGroupSearchResults**](StaticGroupSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MobileDeviceGroupsStaticGroupsIdDelete

> V1MobileDeviceGroupsStaticGroupsIdDelete(ctx, id).Execute()

Remove Static Group by Id 



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
	id := "id_example" // string | instance id of static-group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of static-group | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDeviceGroupsStaticGroupsIdDeleteRequest struct via the builder pattern


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


## V1MobileDeviceGroupsStaticGroupsIdGet

> StaticGroup V1MobileDeviceGroupsStaticGroupsIdGet(ctx, id).Execute()

Get Static Group by Id 



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
	id := "id_example" // string | instance id of static-group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MobileDeviceGroupsStaticGroupsIdGet`: StaticGroup
	fmt.Fprintf(os.Stdout, "Response from `MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of static-group | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDeviceGroupsStaticGroupsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StaticGroup**](StaticGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MobileDeviceGroupsStaticGroupsIdPatch

> StaticGroupAssignment V1MobileDeviceGroupsStaticGroupsIdPatch(ctx, id).StaticGroupAssignment(staticGroupAssignment).Execute()

Update membership of a static group. 



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
	id := "id_example" // string | instance id of a static group
	staticGroupAssignment := *openapiclient.NewStaticGroupAssignment() // StaticGroupAssignment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupsIdPatch(context.Background(), id).StaticGroupAssignment(staticGroupAssignment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupsIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MobileDeviceGroupsStaticGroupsIdPatch`: StaticGroupAssignment
	fmt.Fprintf(os.Stdout, "Response from `MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of a static group | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDeviceGroupsStaticGroupsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **staticGroupAssignment** | [**StaticGroupAssignment**](StaticGroupAssignment.md) |  | 

### Return type

[**StaticGroupAssignment**](StaticGroupAssignment.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MobileDeviceGroupsStaticGroupsPost

> HrefResponse V1MobileDeviceGroupsStaticGroupsPost(ctx).StaticGroupAssignment(staticGroupAssignment).Execute()

Create membership of a static group. 



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
	staticGroupAssignment := *openapiclient.NewStaticGroupAssignment() // StaticGroupAssignment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupsPost(context.Background()).StaticGroupAssignment(staticGroupAssignment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MobileDeviceGroupsStaticGroupsPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `MobileDeviceGroupsAPI.V1MobileDeviceGroupsStaticGroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDeviceGroupsStaticGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **staticGroupAssignment** | [**StaticGroupAssignment**](StaticGroupAssignment.md) |  | 

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

