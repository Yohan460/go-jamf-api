# \MobileDevicesAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1MobileDevicesGet**](MobileDevicesAPI.md#V1MobileDevicesGet) | **Get** /v1/mobile-devices | Get Mobile Device objects 
[**V1MobileDevicesIdDetailGet**](MobileDevicesAPI.md#V1MobileDevicesIdDetailGet) | **Get** /v1/mobile-devices/{id}/detail | Get Mobile Device 
[**V1MobileDevicesIdGet**](MobileDevicesAPI.md#V1MobileDevicesIdGet) | **Get** /v1/mobile-devices/{id} | Get Mobile Device 
[**V1MobileDevicesIdPatch**](MobileDevicesAPI.md#V1MobileDevicesIdPatch) | **Patch** /v1/mobile-devices/{id} | Update fields on a mobile device that are allowed to be modified by users 
[**V1SearchMobileDevicesPost**](MobileDevicesAPI.md#V1SearchMobileDevicesPost) | **Post** /v1/search-mobile-devices | Search Mobile Devices 
[**V2MobileDevicesDetailGet**](MobileDevicesAPI.md#V2MobileDevicesDetailGet) | **Get** /v2/mobile-devices/detail | Return paginated Mobile Device Inventory records
[**V2MobileDevicesGet**](MobileDevicesAPI.md#V2MobileDevicesGet) | **Get** /v2/mobile-devices | Get Mobile Device objects 
[**V2MobileDevicesIdDetailGet**](MobileDevicesAPI.md#V2MobileDevicesIdDetailGet) | **Get** /v2/mobile-devices/{id}/detail | Get Mobile Device 
[**V2MobileDevicesIdGet**](MobileDevicesAPI.md#V2MobileDevicesIdGet) | **Get** /v2/mobile-devices/{id} | Get Mobile Device 
[**V2MobileDevicesIdPatch**](MobileDevicesAPI.md#V2MobileDevicesIdPatch) | **Patch** /v2/mobile-devices/{id} | Update fields on a mobile device that are allowed to be modified by users 



## V1MobileDevicesGet

> []MobileDevice V1MobileDevicesGet(ctx).Execute()

Get Mobile Device objects 



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
    resp, r, err := apiClient.MobileDevicesAPI.V1MobileDevicesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDevicesAPI.V1MobileDevicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1MobileDevicesGet`: []MobileDevice
    fmt.Fprintf(os.Stdout, "Response from `MobileDevicesAPI.V1MobileDevicesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDevicesGetRequest struct via the builder pattern


### Return type

[**[]MobileDevice**](MobileDevice.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MobileDevicesIdDetailGet

> MobileDeviceDetails V1MobileDevicesIdDetailGet(ctx, id).Execute()

Get Mobile Device 



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
    id := int32(56) // int32 | instance id of mobile device record

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileDevicesAPI.V1MobileDevicesIdDetailGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDevicesAPI.V1MobileDevicesIdDetailGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1MobileDevicesIdDetailGet`: MobileDeviceDetails
    fmt.Fprintf(os.Stdout, "Response from `MobileDevicesAPI.V1MobileDevicesIdDetailGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | instance id of mobile device record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDevicesIdDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MobileDeviceDetails**](MobileDeviceDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MobileDevicesIdGet

> MobileDevice V1MobileDevicesIdGet(ctx, id).Execute()

Get Mobile Device 



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
    id := int32(56) // int32 | instance id of mobile device record

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileDevicesAPI.V1MobileDevicesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDevicesAPI.V1MobileDevicesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1MobileDevicesIdGet`: MobileDevice
    fmt.Fprintf(os.Stdout, "Response from `MobileDevicesAPI.V1MobileDevicesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | instance id of mobile device record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDevicesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MobileDevice**](MobileDevice.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MobileDevicesIdPatch

> MobileDeviceDetails V1MobileDevicesIdPatch(ctx, id).UpdateMobileDevice(updateMobileDevice).Execute()

Update fields on a mobile device that are allowed to be modified by users 



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
    id := int32(56) // int32 | instance id of mobile device record
    updateMobileDevice := *openapiclient.NewUpdateMobileDevice() // UpdateMobileDevice | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileDevicesAPI.V1MobileDevicesIdPatch(context.Background(), id).UpdateMobileDevice(updateMobileDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDevicesAPI.V1MobileDevicesIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1MobileDevicesIdPatch`: MobileDeviceDetails
    fmt.Fprintf(os.Stdout, "Response from `MobileDevicesAPI.V1MobileDevicesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | instance id of mobile device record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MobileDevicesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMobileDevice** | [**UpdateMobileDevice**](UpdateMobileDevice.md) |  | 

### Return type

[**MobileDeviceDetails**](MobileDeviceDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SearchMobileDevicesPost

> MobileDeviceSearchResults V1SearchMobileDevicesPost(ctx).MobileDeviceSearchParams(mobileDeviceSearchParams).Execute()

Search Mobile Devices 



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
    mobileDeviceSearchParams := *openapiclient.NewMobileDeviceSearchParams() // MobileDeviceSearchParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileDevicesAPI.V1SearchMobileDevicesPost(context.Background()).MobileDeviceSearchParams(mobileDeviceSearchParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDevicesAPI.V1SearchMobileDevicesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SearchMobileDevicesPost`: MobileDeviceSearchResults
    fmt.Fprintf(os.Stdout, "Response from `MobileDevicesAPI.V1SearchMobileDevicesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SearchMobileDevicesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mobileDeviceSearchParams** | [**MobileDeviceSearchParams**](MobileDeviceSearchParams.md) |  | 

### Return type

[**MobileDeviceSearchResults**](MobileDeviceSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2MobileDevicesDetailGet

> MobileDeviceInventorySearchResults V2MobileDevicesDetailGet(ctx).Section(section).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Return paginated Mobile Device Inventory records



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
    section := []openapiclient.MobileDeviceSection{openapiclient.MobileDeviceSection("GENERAL")} // []MobileDeviceSection | section of mobile device details, if not specified, General section data is returned. Multiple section parameters are supported, e.g. section=GENERAL&section=HARDWARE (optional) (default to ["GENERAL"])
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is displayName:asc. Multiple sort criteria are supported and must be separated with a comma.   Fields allowed in the sort: `airPlayPassword`, `appAnalyticsEnabled`, `assetTag`, `availableSpaceMb`,  `batteryLevel`, `bluetoothLowEnergyCapable`, `bluetoothMacAddress`, `capacityMb`,  `lostModeEnabledDate`, `declarativeDeviceManagementEnabled`, `deviceId`, `deviceLocatorServiceEnabled`, `devicePhoneNumber`, `diagnosticAndUsageReportingEnabled`, `displayName`, `doNotDisturbEnabled`,  `enrollmentSessionTokenValid`, `exchangeDeviceId`, `cloudBackupEnabled`, `osBuild`, `osSupplementalBuildVersion`, `osVersion`, `osRapidSecurityResponse`, `ipAddress`, `itunesStoreAccountActive`, `mobileDeviceId`, `languages`, `lastBackupDate`, `lastEnrolledDate`, `lastCloudBackupDate`, `lastInventoryUpdateDate`, `locales`, `locationServicesForSelfServiceMobileEnabled`, `lostModeEnabled`, `managed`, `mdmProfileExpirationDate`, `model`, `modelIdentifier`, `modelNumber`, `modemFirmwareVersion`, `quotaSize`, `residentUsers`, `serialNumber`, `sharedIpad`, `supervised`, `tethered`, `timeZone`, `udid`, `usedSpacePercentage`, `wifiMacAddress`, `deviceOwnershipType`, `building`, `department`, `emailAddress`, `fullName`, `userPhoneNumber`, `position`, `room`, `username`, `appleCareId`, `leaseExpirationDate`,`lifeExpectancyYears`, `poDate`, `poNumber`, `purchasePrice`, `purchasedOrLeased`, `purchasingAccount`, `purchasingContact`, `vendor`, `warrantyExpirationDate`, `activationLockEnabled`, `blockEncryptionCapable`, `dataProtection`, `fileEncryptionCapable`, `hardwareEncryptionSupported`, `jailbreakStatus`, `passcodeCompliant`, `passcodeCompliantWithProfile`, `passcodeLockGracePeriodEnforcedSeconds`, `passcodePresent`, `personalDeviceProfileCurrent`, `carrierSettingsVersion`, `cellularTechnology`, `currentCarrierNetwork`, `currentMobileCountryCode`, `currentMobileNetworkCode`,  `dataRoamingEnabled`, `eid`, `network`, `homeMobileCountryCode`,  `homeMobileNetworkCode`, `iccid`, `imei`, `imei2`, `meid`, `personalHotspotEnabled`, `voiceRoamingEnabled`, `roaming`  Example: `sort=displayName:desc,username:asc`  (optional) (default to ["displayName:asc"])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter mobile device collection. Default filter is empty query - returning all results for the requested page.  Fields allowed in the query: `airPlayPassword`, `appAnalyticsEnabled`, `assetTag`, `availableSpaceMb`,  `batteryLevel`, `bluetoothLowEnergyCapable`, `bluetoothMacAddress`, `capacityMb`,  `declarativeDeviceManagementEnabled`, `deviceId`, `deviceLocatorServiceEnabled`, `devicePhoneNumber`, `diagnosticAndUsageReportingEnabled`, `displayName`, `doNotDisturbEnabled`, `exchangeDeviceId`,  `cloudBackupEnabled`, `osBuild`, `osSupplementalBuildVersion`, `osVersion`, `osRapidSecurityResponse`, `ipAddress`, `itunesStoreAccountActive`, `mobileDeviceId`, `languages`, `locales`, `locationServicesForSelfServiceMobileEnabled`, `lostModeEnabled`, `managed`, `model`,  `modelIdentifier`, `modelNumber`, `modemFirmwareVersion`, `quotaSize`,  `residentUsers`, `serialNumber`, `sharedIpad`, `supervised`, `tethered`, `timeZone`, `udid`, `usedSpacePercentage`,  `wifiMacAddress`, `building`, `department`, `emailAddress`, `fullName`, `userPhoneNumber`, `position`, `room`, `username`, `appleCareId`, `lifeExpectancyYears`, `poNumber`,  `purchasePrice`, `purchasedOrLeased`, `purchasingAccount`, `purchasingContact`, `vendor`, `activationLockEnabled`, `blockEncryptionCapable`, `dataProtection`,  `fileEncryptionCapable`, `passcodeCompliant`, `passcodeCompliantWithProfile`, `passcodeLockGracePeriodEnforcedSeconds`, `passcodePresent`, `personalDeviceProfileCurrent`, `carrierSettingsVersion`, `currentCarrierNetwork`, `currentMobileCountryCode`, `currentMobileNetworkCode`, `dataRoamingEnabled`, `eid`, `network`, `homeMobileCountryCode`, `homeMobileNetworkCode`, `iccid`, `imei`, `imei2`, `meid`, `personalHotspotEnabled`,  `roaming`  This param can be combined with paging and sorting. Example: `filter=displayName==\"iPad\"`  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileDevicesAPI.V2MobileDevicesDetailGet(context.Background()).Section(section).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDevicesAPI.V2MobileDevicesDetailGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2MobileDevicesDetailGet`: MobileDeviceInventorySearchResults
    fmt.Fprintf(os.Stdout, "Response from `MobileDevicesAPI.V2MobileDevicesDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2MobileDevicesDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **section** | [**[]MobileDeviceSection**](MobileDeviceSection.md) | section of mobile device details, if not specified, General section data is returned. Multiple section parameters are supported, e.g. section&#x3D;GENERAL&amp;section&#x3D;HARDWARE | [default to [&quot;GENERAL&quot;]]
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is displayName:asc. Multiple sort criteria are supported and must be separated with a comma.   Fields allowed in the sort: &#x60;airPlayPassword&#x60;, &#x60;appAnalyticsEnabled&#x60;, &#x60;assetTag&#x60;, &#x60;availableSpaceMb&#x60;,  &#x60;batteryLevel&#x60;, &#x60;bluetoothLowEnergyCapable&#x60;, &#x60;bluetoothMacAddress&#x60;, &#x60;capacityMb&#x60;,  &#x60;lostModeEnabledDate&#x60;, &#x60;declarativeDeviceManagementEnabled&#x60;, &#x60;deviceId&#x60;, &#x60;deviceLocatorServiceEnabled&#x60;, &#x60;devicePhoneNumber&#x60;, &#x60;diagnosticAndUsageReportingEnabled&#x60;, &#x60;displayName&#x60;, &#x60;doNotDisturbEnabled&#x60;,  &#x60;enrollmentSessionTokenValid&#x60;, &#x60;exchangeDeviceId&#x60;, &#x60;cloudBackupEnabled&#x60;, &#x60;osBuild&#x60;, &#x60;osSupplementalBuildVersion&#x60;, &#x60;osVersion&#x60;, &#x60;osRapidSecurityResponse&#x60;, &#x60;ipAddress&#x60;, &#x60;itunesStoreAccountActive&#x60;, &#x60;mobileDeviceId&#x60;, &#x60;languages&#x60;, &#x60;lastBackupDate&#x60;, &#x60;lastEnrolledDate&#x60;, &#x60;lastCloudBackupDate&#x60;, &#x60;lastInventoryUpdateDate&#x60;, &#x60;locales&#x60;, &#x60;locationServicesForSelfServiceMobileEnabled&#x60;, &#x60;lostModeEnabled&#x60;, &#x60;managed&#x60;, &#x60;mdmProfileExpirationDate&#x60;, &#x60;model&#x60;, &#x60;modelIdentifier&#x60;, &#x60;modelNumber&#x60;, &#x60;modemFirmwareVersion&#x60;, &#x60;quotaSize&#x60;, &#x60;residentUsers&#x60;, &#x60;serialNumber&#x60;, &#x60;sharedIpad&#x60;, &#x60;supervised&#x60;, &#x60;tethered&#x60;, &#x60;timeZone&#x60;, &#x60;udid&#x60;, &#x60;usedSpacePercentage&#x60;, &#x60;wifiMacAddress&#x60;, &#x60;deviceOwnershipType&#x60;, &#x60;building&#x60;, &#x60;department&#x60;, &#x60;emailAddress&#x60;, &#x60;fullName&#x60;, &#x60;userPhoneNumber&#x60;, &#x60;position&#x60;, &#x60;room&#x60;, &#x60;username&#x60;, &#x60;appleCareId&#x60;, &#x60;leaseExpirationDate&#x60;,&#x60;lifeExpectancyYears&#x60;, &#x60;poDate&#x60;, &#x60;poNumber&#x60;, &#x60;purchasePrice&#x60;, &#x60;purchasedOrLeased&#x60;, &#x60;purchasingAccount&#x60;, &#x60;purchasingContact&#x60;, &#x60;vendor&#x60;, &#x60;warrantyExpirationDate&#x60;, &#x60;activationLockEnabled&#x60;, &#x60;blockEncryptionCapable&#x60;, &#x60;dataProtection&#x60;, &#x60;fileEncryptionCapable&#x60;, &#x60;hardwareEncryptionSupported&#x60;, &#x60;jailbreakStatus&#x60;, &#x60;passcodeCompliant&#x60;, &#x60;passcodeCompliantWithProfile&#x60;, &#x60;passcodeLockGracePeriodEnforcedSeconds&#x60;, &#x60;passcodePresent&#x60;, &#x60;personalDeviceProfileCurrent&#x60;, &#x60;carrierSettingsVersion&#x60;, &#x60;cellularTechnology&#x60;, &#x60;currentCarrierNetwork&#x60;, &#x60;currentMobileCountryCode&#x60;, &#x60;currentMobileNetworkCode&#x60;,  &#x60;dataRoamingEnabled&#x60;, &#x60;eid&#x60;, &#x60;network&#x60;, &#x60;homeMobileCountryCode&#x60;,  &#x60;homeMobileNetworkCode&#x60;, &#x60;iccid&#x60;, &#x60;imei&#x60;, &#x60;imei2&#x60;, &#x60;meid&#x60;, &#x60;personalHotspotEnabled&#x60;, &#x60;voiceRoamingEnabled&#x60;, &#x60;roaming&#x60;  Example: &#x60;sort&#x3D;displayName:desc,username:asc&#x60;  | [default to [&quot;displayName:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter mobile device collection. Default filter is empty query - returning all results for the requested page.  Fields allowed in the query: &#x60;airPlayPassword&#x60;, &#x60;appAnalyticsEnabled&#x60;, &#x60;assetTag&#x60;, &#x60;availableSpaceMb&#x60;,  &#x60;batteryLevel&#x60;, &#x60;bluetoothLowEnergyCapable&#x60;, &#x60;bluetoothMacAddress&#x60;, &#x60;capacityMb&#x60;,  &#x60;declarativeDeviceManagementEnabled&#x60;, &#x60;deviceId&#x60;, &#x60;deviceLocatorServiceEnabled&#x60;, &#x60;devicePhoneNumber&#x60;, &#x60;diagnosticAndUsageReportingEnabled&#x60;, &#x60;displayName&#x60;, &#x60;doNotDisturbEnabled&#x60;, &#x60;exchangeDeviceId&#x60;,  &#x60;cloudBackupEnabled&#x60;, &#x60;osBuild&#x60;, &#x60;osSupplementalBuildVersion&#x60;, &#x60;osVersion&#x60;, &#x60;osRapidSecurityResponse&#x60;, &#x60;ipAddress&#x60;, &#x60;itunesStoreAccountActive&#x60;, &#x60;mobileDeviceId&#x60;, &#x60;languages&#x60;, &#x60;locales&#x60;, &#x60;locationServicesForSelfServiceMobileEnabled&#x60;, &#x60;lostModeEnabled&#x60;, &#x60;managed&#x60;, &#x60;model&#x60;,  &#x60;modelIdentifier&#x60;, &#x60;modelNumber&#x60;, &#x60;modemFirmwareVersion&#x60;, &#x60;quotaSize&#x60;,  &#x60;residentUsers&#x60;, &#x60;serialNumber&#x60;, &#x60;sharedIpad&#x60;, &#x60;supervised&#x60;, &#x60;tethered&#x60;, &#x60;timeZone&#x60;, &#x60;udid&#x60;, &#x60;usedSpacePercentage&#x60;,  &#x60;wifiMacAddress&#x60;, &#x60;building&#x60;, &#x60;department&#x60;, &#x60;emailAddress&#x60;, &#x60;fullName&#x60;, &#x60;userPhoneNumber&#x60;, &#x60;position&#x60;, &#x60;room&#x60;, &#x60;username&#x60;, &#x60;appleCareId&#x60;, &#x60;lifeExpectancyYears&#x60;, &#x60;poNumber&#x60;,  &#x60;purchasePrice&#x60;, &#x60;purchasedOrLeased&#x60;, &#x60;purchasingAccount&#x60;, &#x60;purchasingContact&#x60;, &#x60;vendor&#x60;, &#x60;activationLockEnabled&#x60;, &#x60;blockEncryptionCapable&#x60;, &#x60;dataProtection&#x60;,  &#x60;fileEncryptionCapable&#x60;, &#x60;passcodeCompliant&#x60;, &#x60;passcodeCompliantWithProfile&#x60;, &#x60;passcodeLockGracePeriodEnforcedSeconds&#x60;, &#x60;passcodePresent&#x60;, &#x60;personalDeviceProfileCurrent&#x60;, &#x60;carrierSettingsVersion&#x60;, &#x60;currentCarrierNetwork&#x60;, &#x60;currentMobileCountryCode&#x60;, &#x60;currentMobileNetworkCode&#x60;, &#x60;dataRoamingEnabled&#x60;, &#x60;eid&#x60;, &#x60;network&#x60;, &#x60;homeMobileCountryCode&#x60;, &#x60;homeMobileNetworkCode&#x60;, &#x60;iccid&#x60;, &#x60;imei&#x60;, &#x60;imei2&#x60;, &#x60;meid&#x60;, &#x60;personalHotspotEnabled&#x60;,  &#x60;roaming&#x60;  This param can be combined with paging and sorting. Example: &#x60;filter&#x3D;displayName&#x3D;&#x3D;\&quot;iPad\&quot;&#x60;  | [default to &quot;&quot;]

### Return type

[**MobileDeviceInventorySearchResults**](MobileDeviceInventorySearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2MobileDevicesGet

> MobileDeviceSearchResultsV2 V2MobileDevicesGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Execute()

Get Mobile Device objects 



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
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["id:asc"])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileDevicesAPI.V2MobileDevicesGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDevicesAPI.V2MobileDevicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2MobileDevicesGet`: MobileDeviceSearchResultsV2
    fmt.Fprintf(os.Stdout, "Response from `MobileDevicesAPI.V2MobileDevicesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2MobileDevicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;id:asc&quot;]]

### Return type

[**MobileDeviceSearchResultsV2**](MobileDeviceSearchResultsV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2MobileDevicesIdDetailGet

> MobileDeviceDetailsGetV2 V2MobileDevicesIdDetailGet(ctx, id).Execute()

Get Mobile Device 



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
    id := "id_example" // string | instance id of mobile device record

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileDevicesAPI.V2MobileDevicesIdDetailGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDevicesAPI.V2MobileDevicesIdDetailGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2MobileDevicesIdDetailGet`: MobileDeviceDetailsGetV2
    fmt.Fprintf(os.Stdout, "Response from `MobileDevicesAPI.V2MobileDevicesIdDetailGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of mobile device record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2MobileDevicesIdDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MobileDeviceDetailsGetV2**](MobileDeviceDetailsGetV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2MobileDevicesIdGet

> MobileDeviceV2 V2MobileDevicesIdGet(ctx, id).Execute()

Get Mobile Device 



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
    id := "id_example" // string | instance id of mobile device record

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileDevicesAPI.V2MobileDevicesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDevicesAPI.V2MobileDevicesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2MobileDevicesIdGet`: MobileDeviceV2
    fmt.Fprintf(os.Stdout, "Response from `MobileDevicesAPI.V2MobileDevicesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of mobile device record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2MobileDevicesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MobileDeviceV2**](MobileDeviceV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2MobileDevicesIdPatch

> MobileDeviceDetailsV2 V2MobileDevicesIdPatch(ctx, id).UpdateMobileDeviceV2(updateMobileDeviceV2).Execute()

Update fields on a mobile device that are allowed to be modified by users 



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
    id := "id_example" // string | instance id of mobile device record
    updateMobileDeviceV2 := *openapiclient.NewUpdateMobileDeviceV2() // UpdateMobileDeviceV2 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileDevicesAPI.V2MobileDevicesIdPatch(context.Background(), id).UpdateMobileDeviceV2(updateMobileDeviceV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileDevicesAPI.V2MobileDevicesIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2MobileDevicesIdPatch`: MobileDeviceDetailsV2
    fmt.Fprintf(os.Stdout, "Response from `MobileDevicesAPI.V2MobileDevicesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of mobile device record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2MobileDevicesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMobileDeviceV2** | [**UpdateMobileDeviceV2**](UpdateMobileDeviceV2.md) |  | 

### Return type

[**MobileDeviceDetailsV2**](MobileDeviceDetailsV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

