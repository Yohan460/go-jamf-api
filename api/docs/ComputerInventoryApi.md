# \ComputerInventoryAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ComputerInventoryIdErasePost**](ComputerInventoryAPI.md#V1ComputerInventoryIdErasePost) | **Post** /v1/computer-inventory/{id}/erase | Erase a computer 
[**V1ComputerInventoryIdRemoveMdmProfilePost**](ComputerInventoryAPI.md#V1ComputerInventoryIdRemoveMdmProfilePost) | **Post** /v1/computer-inventory/{id}/remove-mdm-profile | Remove a computer&#39;s MDM profile 
[**V1ComputersInventoryDetailIdGet**](ComputerInventoryAPI.md#V1ComputersInventoryDetailIdGet) | **Get** /v1/computers-inventory-detail/{id} | Return all sections of a computer
[**V1ComputersInventoryDetailIdPatch**](ComputerInventoryAPI.md#V1ComputersInventoryDetailIdPatch) | **Patch** /v1/computers-inventory-detail/{id} | Update specific fields on a computer
[**V1ComputersInventoryFilevaultGet**](ComputerInventoryAPI.md#V1ComputersInventoryFilevaultGet) | **Get** /v1/computers-inventory/filevault | Return paginated FileVault information for all computers
[**V1ComputersInventoryGet**](ComputerInventoryAPI.md#V1ComputersInventoryGet) | **Get** /v1/computers-inventory | Return paginated Computer Inventory records
[**V1ComputersInventoryIdAttachmentsAttachmentIdDelete**](ComputerInventoryAPI.md#V1ComputersInventoryIdAttachmentsAttachmentIdDelete) | **Delete** /v1/computers-inventory/{id}/attachments/{attachmentId} | Remove attachment
[**V1ComputersInventoryIdAttachmentsAttachmentIdGet**](ComputerInventoryAPI.md#V1ComputersInventoryIdAttachmentsAttachmentIdGet) | **Get** /v1/computers-inventory/{id}/attachments/{attachmentId} | Download attachment file
[**V1ComputersInventoryIdAttachmentsPost**](ComputerInventoryAPI.md#V1ComputersInventoryIdAttachmentsPost) | **Post** /v1/computers-inventory/{id}/attachments | Upload attachment and assign to computer
[**V1ComputersInventoryIdDelete**](ComputerInventoryAPI.md#V1ComputersInventoryIdDelete) | **Delete** /v1/computers-inventory/{id} | Remove specified Computer record
[**V1ComputersInventoryIdFilevaultGet**](ComputerInventoryAPI.md#V1ComputersInventoryIdFilevaultGet) | **Get** /v1/computers-inventory/{id}/filevault | Return FileVault information for a specific computer
[**V1ComputersInventoryIdGet**](ComputerInventoryAPI.md#V1ComputersInventoryIdGet) | **Get** /v1/computers-inventory/{id} | Return General section of a Computer
[**V1ComputersInventoryIdViewDeviceLockPinGet**](ComputerInventoryAPI.md#V1ComputersInventoryIdViewDeviceLockPinGet) | **Get** /v1/computers-inventory/{id}/view-device-lock-pin | Return a computer&#39;s Device Lock PIN
[**V1ComputersInventoryIdViewRecoveryLockPasswordGet**](ComputerInventoryAPI.md#V1ComputersInventoryIdViewRecoveryLockPasswordGet) | **Get** /v1/computers-inventory/{id}/view-recovery-lock-password | Return a Computers Recovery Lock Password
[**V1ComputersInventoryPost**](ComputerInventoryAPI.md#V1ComputersInventoryPost) | **Post** /v1/computers-inventory | Create Computer Inventory record
[**V2ComputersInventoryDetailIdGet**](ComputerInventoryAPI.md#V2ComputersInventoryDetailIdGet) | **Get** /v2/computers-inventory-detail/{id} | Return all sections of a computer
[**V2ComputersInventoryDetailIdPatch**](ComputerInventoryAPI.md#V2ComputersInventoryDetailIdPatch) | **Patch** /v2/computers-inventory-detail/{id} | Update specific fields on a computer
[**V2ComputersInventoryFilevaultGet**](ComputerInventoryAPI.md#V2ComputersInventoryFilevaultGet) | **Get** /v2/computers-inventory/filevault | Return paginated FileVault information for all computers
[**V2ComputersInventoryGet**](ComputerInventoryAPI.md#V2ComputersInventoryGet) | **Get** /v2/computers-inventory | Return paginated Computer Inventory records
[**V2ComputersInventoryIdAttachmentsAttachmentIdDelete**](ComputerInventoryAPI.md#V2ComputersInventoryIdAttachmentsAttachmentIdDelete) | **Delete** /v2/computers-inventory/{id}/attachments/{attachmentId} | Remove attachment
[**V2ComputersInventoryIdAttachmentsAttachmentIdGet**](ComputerInventoryAPI.md#V2ComputersInventoryIdAttachmentsAttachmentIdGet) | **Get** /v2/computers-inventory/{id}/attachments/{attachmentId} | Download attachment file
[**V2ComputersInventoryIdAttachmentsPost**](ComputerInventoryAPI.md#V2ComputersInventoryIdAttachmentsPost) | **Post** /v2/computers-inventory/{id}/attachments | Upload attachment and assign to computer
[**V2ComputersInventoryIdDelete**](ComputerInventoryAPI.md#V2ComputersInventoryIdDelete) | **Delete** /v2/computers-inventory/{id} | Remove specified Computer record
[**V2ComputersInventoryIdFilevaultGet**](ComputerInventoryAPI.md#V2ComputersInventoryIdFilevaultGet) | **Get** /v2/computers-inventory/{id}/filevault | Return FileVault information for a specific computer
[**V2ComputersInventoryIdGet**](ComputerInventoryAPI.md#V2ComputersInventoryIdGet) | **Get** /v2/computers-inventory/{id} | Return General section of a Computer
[**V2ComputersInventoryIdViewDeviceLockPinGet**](ComputerInventoryAPI.md#V2ComputersInventoryIdViewDeviceLockPinGet) | **Get** /v2/computers-inventory/{id}/view-device-lock-pin | Return a computer&#39;s Device Lock PIN
[**V2ComputersInventoryIdViewRecoveryLockPasswordGet**](ComputerInventoryAPI.md#V2ComputersInventoryIdViewRecoveryLockPasswordGet) | **Get** /v2/computers-inventory/{id}/view-recovery-lock-password | Return a Computers Recovery Lock Password
[**V2ComputersInventoryPost**](ComputerInventoryAPI.md#V2ComputersInventoryPost) | **Post** /v2/computers-inventory | Create Computer Inventory record



## V1ComputerInventoryIdErasePost

> EraseDeviceComputerResponse V1ComputerInventoryIdErasePost(ctx, id).EraseDeviceComputerRequest(eraseDeviceComputerRequest).Execute()

Erase a computer 



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
	id := "id_example" // string | Id of the computer to erase
	eraseDeviceComputerRequest := *openapiclient.NewEraseDeviceComputerRequest() // EraseDeviceComputerRequest | Options for eraseDevice command (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V1ComputerInventoryIdErasePost(context.Background(), id).EraseDeviceComputerRequest(eraseDeviceComputerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V1ComputerInventoryIdErasePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputerInventoryIdErasePost`: EraseDeviceComputerResponse
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V1ComputerInventoryIdErasePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the computer to erase | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputerInventoryIdErasePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eraseDeviceComputerRequest** | [**EraseDeviceComputerRequest**](EraseDeviceComputerRequest.md) | Options for eraseDevice command | 

### Return type

[**EraseDeviceComputerResponse**](EraseDeviceComputerResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ComputerInventoryIdRemoveMdmProfilePost

> RemoveComputerMdmProfileResponse V1ComputerInventoryIdRemoveMdmProfilePost(ctx, id).Execute()

Remove a computer's MDM profile 



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
	id := "id_example" // string | Id of the computer to remove the MDM profile from

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V1ComputerInventoryIdRemoveMdmProfilePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V1ComputerInventoryIdRemoveMdmProfilePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputerInventoryIdRemoveMdmProfilePost`: RemoveComputerMdmProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V1ComputerInventoryIdRemoveMdmProfilePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the computer to remove the MDM profile from | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputerInventoryIdRemoveMdmProfilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RemoveComputerMdmProfileResponse**](RemoveComputerMdmProfileResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ComputersInventoryDetailIdGet

> ComputerInventory V1ComputersInventoryDetailIdGet(ctx, id).Execute()

Return all sections of a computer



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
	id := "id_example" // string | instance id of computer record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V1ComputersInventoryDetailIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V1ComputersInventoryDetailIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputersInventoryDetailIdGet`: ComputerInventory
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V1ComputersInventoryDetailIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of computer record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputersInventoryDetailIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComputerInventory**](ComputerInventory.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ComputersInventoryDetailIdPatch

> ComputerInventory V1ComputersInventoryDetailIdPatch(ctx, id).ComputerInventoryUpdateRequest(computerInventoryUpdateRequest).Execute()

Update specific fields on a computer



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
	id := "id_example" // string | instance id of computer record
	computerInventoryUpdateRequest := *openapiclient.NewComputerInventoryUpdateRequest() // ComputerInventoryUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V1ComputersInventoryDetailIdPatch(context.Background(), id).ComputerInventoryUpdateRequest(computerInventoryUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V1ComputersInventoryDetailIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputersInventoryDetailIdPatch`: ComputerInventory
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V1ComputersInventoryDetailIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of computer record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputersInventoryDetailIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **computerInventoryUpdateRequest** | [**ComputerInventoryUpdateRequest**](ComputerInventoryUpdateRequest.md) |  | 

### Return type

[**ComputerInventory**](ComputerInventory.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ComputersInventoryFilevaultGet

> ComputerInventoryFileVaultSearchResults V1ComputersInventoryFilevaultGet(ctx).Page(page).PageSize(pageSize).Execute()

Return paginated FileVault information for all computers



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V1ComputersInventoryFilevaultGet(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V1ComputersInventoryFilevaultGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputersInventoryFilevaultGet`: ComputerInventoryFileVaultSearchResults
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V1ComputersInventoryFilevaultGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputersInventoryFilevaultGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]

### Return type

[**ComputerInventoryFileVaultSearchResults**](ComputerInventoryFileVaultSearchResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ComputersInventoryGet

> ComputerInventorySearchResults V1ComputersInventoryGet(ctx).Section(section).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Return paginated Computer Inventory records



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
	section := []openapiclient.ComputerSection{openapiclient.ComputerSection("GENERAL")} // []ComputerSection | section of computer details, if not specified, General section data is returned. Multiple section parameters are supported, e.g. section=GENERAL&section=HARDWARE (optional) (default to {"GENERAL"})
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: `property:asc/desc`. Default sort is `general.name:asc`. Multiple sort criteria are supported and must be separated with a comma.  Fields allowed in the sort: `general.name`, `udid`, `id`, `general.assetTag`, `general.jamfBinaryVersion`, `general.lastContactTime`, `general.lastEnrolledDate`, `general.lastCloudBackupDate`, `general.reportDate`, `general.mdmCertificateExpiration`, `general.platform`, `general.lastLoggedInUsernameSelfService`, `general.lastLoggedInUsernameSelfServiceTimestamp`, `general.mdmCertificateExpiration`, `general.platform`, `general.lastLoggedInUsernameBinary`, `general.lastLoggedInUsernameBinaryTimestamp` `hardware.make`, `hardware.model`, `operatingSystem.build`, `operatingSystem.supplementalBuildVersion`, `operatingSystem.rapidSecurityResponse`, `operatingSystem.name`, `operatingSystem.version`, `userAndLocation.realname`, `purchasing.lifeExpectancy`, `purchasing.warrantyDate`  Example: `sort=udid:desc,general.name:asc`.  (optional) (default to {"general.name:asc"})
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter computer inventory collection. Default filter is empty query - returning all results for the requested page.  Fields allowed in the query: `general.name`, `udid`, `id`, `general.assetTag`, `general.barcode1`, `general.barcode2`, `general.enrolledViaAutomatedDeviceEnrollment`, `general.lastIpAddress`, `general.itunesStoreAccountActive`, `general.jamfBinaryVersion`, `general.lastContactTime`, `general.lastEnrolledDate`, `general.lastCloudBackupDate`, `general.reportDate`, `general.lastReportedIp`, `general.lastReportedIpV4`, `general.lastReportedIpV6`, `general.managementId`, `general.remoteManagement.managed`, `general.mdmCapable.capable`, `general.mdmCertificateExpiration`, `general.platform`, `general.supervised`, `general.userApprovedMdm`, `general.declarativeDeviceManagementEnabled`,  `general.lastLoggedInUsernameSelfService`, `general.lastLoggedInUsernameSelfServiceTimestamp`,  `general.mdmCapable.capable`, `general.mdmCertificateExpiration`, `general.platform`, `general.supervised`, `general.userApprovedMdm`, `general.declarativeDeviceManagementEnabled`, `general.lastLoggedInUsernameBinary`, `general.lastLoggedInUsernameBinaryTimestamp`, `hardware.bleCapable`, `hardware.macAddress`, `hardware.make`, `hardware.model`, `hardware.modelIdentifier`, `hardware.serialNumber`, `hardware.supportsIosAppInstalls`,`hardware.appleSilicon`, `operatingSystem.activeDirectoryStatus`, `operatingSystem.fileVault2Status`, `operatingSystem.build`, `operatingSystem.supplementalBuildVersion`, `operatingSystem.rapidSecurityResponse`, `operatingSystem.name`, `operatingSystem.version`, `security.activationLockEnabled`, `security.recoveryLockEnabled`,`security.firewallEnabled`,`userAndLocation.buildingId`, `userAndLocation.departmentId`, `userAndLocation.email`, `userAndLocation.realname`, `userAndLocation.phone`, `userAndLocation.position`,`userAndLocation.room`, `userAndLocation.username`, `diskEncryption.fileVault2Enabled`, `purchasing.appleCareId`, `purchasing.lifeExpectancy`, `purchasing.purchased`, `purchasing.leased`, `purchasing.vendor`, `purchasing.warrantyDate`,  This param can be combined with paging and sorting. Example: `filter=general.name==\"Orchard\"`  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V1ComputersInventoryGet(context.Background()).Section(section).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V1ComputersInventoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputersInventoryGet`: ComputerInventorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V1ComputersInventoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputersInventoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **section** | [**[]ComputerSection**](ComputerSection.md) | section of computer details, if not specified, General section data is returned. Multiple section parameters are supported, e.g. section&#x3D;GENERAL&amp;section&#x3D;HARDWARE | [default to {&quot;GENERAL&quot;}]
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: &#x60;property:asc/desc&#x60;. Default sort is &#x60;general.name:asc&#x60;. Multiple sort criteria are supported and must be separated with a comma.  Fields allowed in the sort: &#x60;general.name&#x60;, &#x60;udid&#x60;, &#x60;id&#x60;, &#x60;general.assetTag&#x60;, &#x60;general.jamfBinaryVersion&#x60;, &#x60;general.lastContactTime&#x60;, &#x60;general.lastEnrolledDate&#x60;, &#x60;general.lastCloudBackupDate&#x60;, &#x60;general.reportDate&#x60;, &#x60;general.mdmCertificateExpiration&#x60;, &#x60;general.platform&#x60;, &#x60;general.lastLoggedInUsernameSelfService&#x60;, &#x60;general.lastLoggedInUsernameSelfServiceTimestamp&#x60;, &#x60;general.mdmCertificateExpiration&#x60;, &#x60;general.platform&#x60;, &#x60;general.lastLoggedInUsernameBinary&#x60;, &#x60;general.lastLoggedInUsernameBinaryTimestamp&#x60; &#x60;hardware.make&#x60;, &#x60;hardware.model&#x60;, &#x60;operatingSystem.build&#x60;, &#x60;operatingSystem.supplementalBuildVersion&#x60;, &#x60;operatingSystem.rapidSecurityResponse&#x60;, &#x60;operatingSystem.name&#x60;, &#x60;operatingSystem.version&#x60;, &#x60;userAndLocation.realname&#x60;, &#x60;purchasing.lifeExpectancy&#x60;, &#x60;purchasing.warrantyDate&#x60;  Example: &#x60;sort&#x3D;udid:desc,general.name:asc&#x60;.  | [default to {&quot;general.name:asc&quot;}]
 **filter** | **string** | Query in the RSQL format, allowing to filter computer inventory collection. Default filter is empty query - returning all results for the requested page.  Fields allowed in the query: &#x60;general.name&#x60;, &#x60;udid&#x60;, &#x60;id&#x60;, &#x60;general.assetTag&#x60;, &#x60;general.barcode1&#x60;, &#x60;general.barcode2&#x60;, &#x60;general.enrolledViaAutomatedDeviceEnrollment&#x60;, &#x60;general.lastIpAddress&#x60;, &#x60;general.itunesStoreAccountActive&#x60;, &#x60;general.jamfBinaryVersion&#x60;, &#x60;general.lastContactTime&#x60;, &#x60;general.lastEnrolledDate&#x60;, &#x60;general.lastCloudBackupDate&#x60;, &#x60;general.reportDate&#x60;, &#x60;general.lastReportedIp&#x60;, &#x60;general.lastReportedIpV4&#x60;, &#x60;general.lastReportedIpV6&#x60;, &#x60;general.managementId&#x60;, &#x60;general.remoteManagement.managed&#x60;, &#x60;general.mdmCapable.capable&#x60;, &#x60;general.mdmCertificateExpiration&#x60;, &#x60;general.platform&#x60;, &#x60;general.supervised&#x60;, &#x60;general.userApprovedMdm&#x60;, &#x60;general.declarativeDeviceManagementEnabled&#x60;,  &#x60;general.lastLoggedInUsernameSelfService&#x60;, &#x60;general.lastLoggedInUsernameSelfServiceTimestamp&#x60;,  &#x60;general.mdmCapable.capable&#x60;, &#x60;general.mdmCertificateExpiration&#x60;, &#x60;general.platform&#x60;, &#x60;general.supervised&#x60;, &#x60;general.userApprovedMdm&#x60;, &#x60;general.declarativeDeviceManagementEnabled&#x60;, &#x60;general.lastLoggedInUsernameBinary&#x60;, &#x60;general.lastLoggedInUsernameBinaryTimestamp&#x60;, &#x60;hardware.bleCapable&#x60;, &#x60;hardware.macAddress&#x60;, &#x60;hardware.make&#x60;, &#x60;hardware.model&#x60;, &#x60;hardware.modelIdentifier&#x60;, &#x60;hardware.serialNumber&#x60;, &#x60;hardware.supportsIosAppInstalls&#x60;,&#x60;hardware.appleSilicon&#x60;, &#x60;operatingSystem.activeDirectoryStatus&#x60;, &#x60;operatingSystem.fileVault2Status&#x60;, &#x60;operatingSystem.build&#x60;, &#x60;operatingSystem.supplementalBuildVersion&#x60;, &#x60;operatingSystem.rapidSecurityResponse&#x60;, &#x60;operatingSystem.name&#x60;, &#x60;operatingSystem.version&#x60;, &#x60;security.activationLockEnabled&#x60;, &#x60;security.recoveryLockEnabled&#x60;,&#x60;security.firewallEnabled&#x60;,&#x60;userAndLocation.buildingId&#x60;, &#x60;userAndLocation.departmentId&#x60;, &#x60;userAndLocation.email&#x60;, &#x60;userAndLocation.realname&#x60;, &#x60;userAndLocation.phone&#x60;, &#x60;userAndLocation.position&#x60;,&#x60;userAndLocation.room&#x60;, &#x60;userAndLocation.username&#x60;, &#x60;diskEncryption.fileVault2Enabled&#x60;, &#x60;purchasing.appleCareId&#x60;, &#x60;purchasing.lifeExpectancy&#x60;, &#x60;purchasing.purchased&#x60;, &#x60;purchasing.leased&#x60;, &#x60;purchasing.vendor&#x60;, &#x60;purchasing.warrantyDate&#x60;,  This param can be combined with paging and sorting. Example: &#x60;filter&#x3D;general.name&#x3D;&#x3D;\&quot;Orchard\&quot;&#x60;  | [default to &quot;&quot;]

### Return type

[**ComputerInventorySearchResults**](ComputerInventorySearchResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ComputersInventoryIdAttachmentsAttachmentIdDelete

> V1ComputersInventoryIdAttachmentsAttachmentIdDelete(ctx, id, attachmentId).Execute()

Remove attachment



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
	id := "id_example" // string | instance id of computer record
	attachmentId := "attachmentId_example" // string | instance id of attachment object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputerInventoryAPI.V1ComputersInventoryIdAttachmentsAttachmentIdDelete(context.Background(), id, attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V1ComputersInventoryIdAttachmentsAttachmentIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of computer record | 
**attachmentId** | **string** | instance id of attachment object | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputersInventoryIdAttachmentsAttachmentIdDeleteRequest struct via the builder pattern


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


## V1ComputersInventoryIdAttachmentsAttachmentIdGet

> *os.File V1ComputersInventoryIdAttachmentsAttachmentIdGet(ctx, id, attachmentId).Execute()

Download attachment file



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
	id := "id_example" // string | instance id of computer record
	attachmentId := "attachmentId_example" // string | instance id of attachment object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V1ComputersInventoryIdAttachmentsAttachmentIdGet(context.Background(), id, attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V1ComputersInventoryIdAttachmentsAttachmentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputersInventoryIdAttachmentsAttachmentIdGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V1ComputersInventoryIdAttachmentsAttachmentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of computer record | 
**attachmentId** | **string** | instance id of attachment object | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputersInventoryIdAttachmentsAttachmentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ComputersInventoryIdAttachmentsPost

> HrefResponse V1ComputersInventoryIdAttachmentsPost(ctx, id).File(file).Execute()

Upload attachment and assign to computer



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
	id := "id_example" // string | instance id of computer record
	file := os.NewFile(1234, "some_file") // *os.File | The file to upload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V1ComputersInventoryIdAttachmentsPost(context.Background(), id).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V1ComputersInventoryIdAttachmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputersInventoryIdAttachmentsPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V1ComputersInventoryIdAttachmentsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of computer record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputersInventoryIdAttachmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The file to upload | 

### Return type

[**HrefResponse**](HrefResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ComputersInventoryIdDelete

> V1ComputersInventoryIdDelete(ctx, id).Execute()

Remove specified Computer record



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
	id := "id_example" // string | instance id of computer record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputerInventoryAPI.V1ComputersInventoryIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V1ComputersInventoryIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of computer record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputersInventoryIdDeleteRequest struct via the builder pattern


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


## V1ComputersInventoryIdFilevaultGet

> ComputerInventoryFileVault V1ComputersInventoryIdFilevaultGet(ctx, id).Execute()

Return FileVault information for a specific computer



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
	id := "id_example" // string | instance id of computer record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V1ComputersInventoryIdFilevaultGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V1ComputersInventoryIdFilevaultGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputersInventoryIdFilevaultGet`: ComputerInventoryFileVault
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V1ComputersInventoryIdFilevaultGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of computer record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputersInventoryIdFilevaultGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComputerInventoryFileVault**](ComputerInventoryFileVault.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ComputersInventoryIdGet

> ComputerInventory V1ComputersInventoryIdGet(ctx, id).Section(section).Execute()

Return General section of a Computer



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
	id := "id_example" // string | instance id of computer record
	section := []openapiclient.ComputerSection{openapiclient.ComputerSection("GENERAL")} // []ComputerSection | section of computer details, if not specified, General section data is returned. Multiple section parameters are supported, e.g. section=general&section=hardware (optional) (default to {"GENERAL"})

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V1ComputersInventoryIdGet(context.Background(), id).Section(section).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V1ComputersInventoryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputersInventoryIdGet`: ComputerInventory
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V1ComputersInventoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of computer record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputersInventoryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **section** | [**[]ComputerSection**](ComputerSection.md) | section of computer details, if not specified, General section data is returned. Multiple section parameters are supported, e.g. section&#x3D;general&amp;section&#x3D;hardware | [default to {&quot;GENERAL&quot;}]

### Return type

[**ComputerInventory**](ComputerInventory.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ComputersInventoryIdViewDeviceLockPinGet

> ComputerInventoryDeviceLockPinResponse V1ComputersInventoryIdViewDeviceLockPinGet(ctx, id).Execute()

Return a computer's Device Lock PIN



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
	id := "id_example" // string | instance id of computer record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V1ComputersInventoryIdViewDeviceLockPinGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V1ComputersInventoryIdViewDeviceLockPinGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputersInventoryIdViewDeviceLockPinGet`: ComputerInventoryDeviceLockPinResponse
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V1ComputersInventoryIdViewDeviceLockPinGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of computer record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputersInventoryIdViewDeviceLockPinGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComputerInventoryDeviceLockPinResponse**](ComputerInventoryDeviceLockPinResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ComputersInventoryIdViewRecoveryLockPasswordGet

> ComputerInventoryRecoveryLockPasswordResponse V1ComputersInventoryIdViewRecoveryLockPasswordGet(ctx, id).Execute()

Return a Computers Recovery Lock Password



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
	id := "id_example" // string | instance id of computer record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V1ComputersInventoryIdViewRecoveryLockPasswordGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V1ComputersInventoryIdViewRecoveryLockPasswordGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputersInventoryIdViewRecoveryLockPasswordGet`: ComputerInventoryRecoveryLockPasswordResponse
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V1ComputersInventoryIdViewRecoveryLockPasswordGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of computer record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputersInventoryIdViewRecoveryLockPasswordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComputerInventoryRecoveryLockPasswordResponse**](ComputerInventoryRecoveryLockPasswordResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ComputersInventoryPost

> HrefResponse V1ComputersInventoryPost(ctx).ComputerInventoryCreateRequest(computerInventoryCreateRequest).Execute()

Create Computer Inventory record



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
	computerInventoryCreateRequest := *openapiclient.NewComputerInventoryCreateRequest() // ComputerInventoryCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V1ComputersInventoryPost(context.Background()).ComputerInventoryCreateRequest(computerInventoryCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V1ComputersInventoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ComputersInventoryPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V1ComputersInventoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputersInventoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **computerInventoryCreateRequest** | [**ComputerInventoryCreateRequest**](ComputerInventoryCreateRequest.md) |  | 

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


## V2ComputersInventoryDetailIdGet

> ComputerInventoryV2 V2ComputersInventoryDetailIdGet(ctx, id).Execute()

Return all sections of a computer



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
	id := "id_example" // string | instance id of computer record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V2ComputersInventoryDetailIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V2ComputersInventoryDetailIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputersInventoryDetailIdGet`: ComputerInventoryV2
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V2ComputersInventoryDetailIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of computer record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputersInventoryDetailIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComputerInventoryV2**](ComputerInventoryV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputersInventoryDetailIdPatch

> V2ComputersInventoryDetailIdPatch(ctx, id).ComputerInventoryUpdateRequest(computerInventoryUpdateRequest).Execute()

Update specific fields on a computer



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
	id := "id_example" // string | instance id of computer record
	computerInventoryUpdateRequest := *openapiclient.NewComputerInventoryUpdateRequest() // ComputerInventoryUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputerInventoryAPI.V2ComputersInventoryDetailIdPatch(context.Background(), id).ComputerInventoryUpdateRequest(computerInventoryUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V2ComputersInventoryDetailIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of computer record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputersInventoryDetailIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **computerInventoryUpdateRequest** | [**ComputerInventoryUpdateRequest**](ComputerInventoryUpdateRequest.md) |  | 

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


## V2ComputersInventoryFilevaultGet

> ComputerInventoryFileVaultSearchResults V2ComputersInventoryFilevaultGet(ctx).Page(page).PageSize(pageSize).Execute()

Return paginated FileVault information for all computers



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V2ComputersInventoryFilevaultGet(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V2ComputersInventoryFilevaultGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputersInventoryFilevaultGet`: ComputerInventoryFileVaultSearchResults
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V2ComputersInventoryFilevaultGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputersInventoryFilevaultGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]

### Return type

[**ComputerInventoryFileVaultSearchResults**](ComputerInventoryFileVaultSearchResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputersInventoryGet

> ComputerInventorySearchResultsV2 V2ComputersInventoryGet(ctx).Section(section).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Return paginated Computer Inventory records



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
	section := []openapiclient.ComputerSectionV2{openapiclient.ComputerSectionV2("GENERAL")} // []ComputerSectionV2 | section of computer details, if not specified, General section data is returned. Multiple section parameters are supported, e.g. section=GENERAL&section=HARDWARE (optional) (default to {"GENERAL"})
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: `property:asc/desc`. Default sort is `general.name:asc`. Multiple sort criteria are supported and must be separated with a comma.  Fields allowed in the sort: `general.name`, `udid`, `id`, `general.assetTag`, `general.jamfBinaryVersion`, `general.lastContactTime`, `general.lastEnrolledDate`, `general.lastCloudBackupDate`, `general.reportDate`, `general.mdmCertificateExpiration`, `general.platform`, `general.lastLoggedInUsernameSelfService`, `general.lastLoggedInUsernameSelfServiceTimestamp`, `general.mdmCertificateExpiration`, `general.platform`, `general.lastLoggedInUsernameBinary`, `general.lastLoggedInUsernameBinaryTimestamp` `hardware.make`, `hardware.model`, `operatingSystem.build`, `operatingSystem.supplementalBuildVersion`, `operatingSystem.rapidSecurityResponse`, `operatingSystem.name`, `operatingSystem.version`, `userAndLocation.realname`, `purchasing.lifeExpectancy`, `purchasing.warrantyDate`  Example: `sort=udid:desc,general.name:asc`.  (optional) (default to {"general.name:asc"})
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter computer inventory collection. Default filter is empty query - returning all results for the requested page.  Fields allowed in the query: `general.name`, `udid`, `id`, `general.assetTag`, `general.barcode1`, `general.barcode2`, `general.enrolledViaAutomatedDeviceEnrollment`, `general.lastIpAddress`, `general.itunesStoreAccountActive`, `general.jamfBinaryVersion`, `general.lastContactTime`, `general.lastEnrolledDate`, `general.lastCloudBackupDate`, `general.reportDate`, `general.lastReportedIp`, `general.lastReportedIpV4`, `general.lastReportedIpV6`, `general.managementId`, `general.remoteManagement.managed`, `general.mdmCapable.capable`, `general.mdmCertificateExpiration`, `general.platform`, `general.supervised`, `general.userApprovedMdm`, `general.declarativeDeviceManagementEnabled`,  `general.lastLoggedInUsernameSelfService`, `general.lastLoggedInUsernameSelfServiceTimestamp`,  `general.mdmCapable.capable`, `general.mdmCertificateExpiration`, `general.platform`, `general.supervised`, `general.userApprovedMdm`, `general.declarativeDeviceManagementEnabled`, `general.lastLoggedInUsernameBinary`, `general.lastLoggedInUsernameBinaryTimestamp`, `hardware.bleCapable`, `hardware.macAddress`, `hardware.make`, `hardware.model`, `hardware.modelIdentifier`, `hardware.serialNumber`, `hardware.supportsIosAppInstalls`,`hardware.appleSilicon`, `operatingSystem.activeDirectoryStatus`, `operatingSystem.fileVault2Status`, `operatingSystem.build`, `operatingSystem.supplementalBuildVersion`, `operatingSystem.rapidSecurityResponse`, `operatingSystem.name`, `operatingSystem.version`, `security.activationLockEnabled`, `security.recoveryLockEnabled`,`security.firewallEnabled`,`userAndLocation.buildingId`, `userAndLocation.departmentId`, `userAndLocation.email`, `userAndLocation.realname`, `userAndLocation.phone`, `userAndLocation.position`,`userAndLocation.room`, `userAndLocation.username`, `diskEncryption.fileVault2Enabled`, `purchasing.appleCareId`, `purchasing.lifeExpectancy`, `purchasing.purchased`, `purchasing.leased`, `purchasing.vendor`, `purchasing.warrantyDate`,  This param can be combined with paging and sorting. Example: `filter=general.name==\"Orchard\"`  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V2ComputersInventoryGet(context.Background()).Section(section).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V2ComputersInventoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputersInventoryGet`: ComputerInventorySearchResultsV2
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V2ComputersInventoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputersInventoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **section** | [**[]ComputerSectionV2**](ComputerSectionV2.md) | section of computer details, if not specified, General section data is returned. Multiple section parameters are supported, e.g. section&#x3D;GENERAL&amp;section&#x3D;HARDWARE | [default to {&quot;GENERAL&quot;}]
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: &#x60;property:asc/desc&#x60;. Default sort is &#x60;general.name:asc&#x60;. Multiple sort criteria are supported and must be separated with a comma.  Fields allowed in the sort: &#x60;general.name&#x60;, &#x60;udid&#x60;, &#x60;id&#x60;, &#x60;general.assetTag&#x60;, &#x60;general.jamfBinaryVersion&#x60;, &#x60;general.lastContactTime&#x60;, &#x60;general.lastEnrolledDate&#x60;, &#x60;general.lastCloudBackupDate&#x60;, &#x60;general.reportDate&#x60;, &#x60;general.mdmCertificateExpiration&#x60;, &#x60;general.platform&#x60;, &#x60;general.lastLoggedInUsernameSelfService&#x60;, &#x60;general.lastLoggedInUsernameSelfServiceTimestamp&#x60;, &#x60;general.mdmCertificateExpiration&#x60;, &#x60;general.platform&#x60;, &#x60;general.lastLoggedInUsernameBinary&#x60;, &#x60;general.lastLoggedInUsernameBinaryTimestamp&#x60; &#x60;hardware.make&#x60;, &#x60;hardware.model&#x60;, &#x60;operatingSystem.build&#x60;, &#x60;operatingSystem.supplementalBuildVersion&#x60;, &#x60;operatingSystem.rapidSecurityResponse&#x60;, &#x60;operatingSystem.name&#x60;, &#x60;operatingSystem.version&#x60;, &#x60;userAndLocation.realname&#x60;, &#x60;purchasing.lifeExpectancy&#x60;, &#x60;purchasing.warrantyDate&#x60;  Example: &#x60;sort&#x3D;udid:desc,general.name:asc&#x60;.  | [default to {&quot;general.name:asc&quot;}]
 **filter** | **string** | Query in the RSQL format, allowing to filter computer inventory collection. Default filter is empty query - returning all results for the requested page.  Fields allowed in the query: &#x60;general.name&#x60;, &#x60;udid&#x60;, &#x60;id&#x60;, &#x60;general.assetTag&#x60;, &#x60;general.barcode1&#x60;, &#x60;general.barcode2&#x60;, &#x60;general.enrolledViaAutomatedDeviceEnrollment&#x60;, &#x60;general.lastIpAddress&#x60;, &#x60;general.itunesStoreAccountActive&#x60;, &#x60;general.jamfBinaryVersion&#x60;, &#x60;general.lastContactTime&#x60;, &#x60;general.lastEnrolledDate&#x60;, &#x60;general.lastCloudBackupDate&#x60;, &#x60;general.reportDate&#x60;, &#x60;general.lastReportedIp&#x60;, &#x60;general.lastReportedIpV4&#x60;, &#x60;general.lastReportedIpV6&#x60;, &#x60;general.managementId&#x60;, &#x60;general.remoteManagement.managed&#x60;, &#x60;general.mdmCapable.capable&#x60;, &#x60;general.mdmCertificateExpiration&#x60;, &#x60;general.platform&#x60;, &#x60;general.supervised&#x60;, &#x60;general.userApprovedMdm&#x60;, &#x60;general.declarativeDeviceManagementEnabled&#x60;,  &#x60;general.lastLoggedInUsernameSelfService&#x60;, &#x60;general.lastLoggedInUsernameSelfServiceTimestamp&#x60;,  &#x60;general.mdmCapable.capable&#x60;, &#x60;general.mdmCertificateExpiration&#x60;, &#x60;general.platform&#x60;, &#x60;general.supervised&#x60;, &#x60;general.userApprovedMdm&#x60;, &#x60;general.declarativeDeviceManagementEnabled&#x60;, &#x60;general.lastLoggedInUsernameBinary&#x60;, &#x60;general.lastLoggedInUsernameBinaryTimestamp&#x60;, &#x60;hardware.bleCapable&#x60;, &#x60;hardware.macAddress&#x60;, &#x60;hardware.make&#x60;, &#x60;hardware.model&#x60;, &#x60;hardware.modelIdentifier&#x60;, &#x60;hardware.serialNumber&#x60;, &#x60;hardware.supportsIosAppInstalls&#x60;,&#x60;hardware.appleSilicon&#x60;, &#x60;operatingSystem.activeDirectoryStatus&#x60;, &#x60;operatingSystem.fileVault2Status&#x60;, &#x60;operatingSystem.build&#x60;, &#x60;operatingSystem.supplementalBuildVersion&#x60;, &#x60;operatingSystem.rapidSecurityResponse&#x60;, &#x60;operatingSystem.name&#x60;, &#x60;operatingSystem.version&#x60;, &#x60;security.activationLockEnabled&#x60;, &#x60;security.recoveryLockEnabled&#x60;,&#x60;security.firewallEnabled&#x60;,&#x60;userAndLocation.buildingId&#x60;, &#x60;userAndLocation.departmentId&#x60;, &#x60;userAndLocation.email&#x60;, &#x60;userAndLocation.realname&#x60;, &#x60;userAndLocation.phone&#x60;, &#x60;userAndLocation.position&#x60;,&#x60;userAndLocation.room&#x60;, &#x60;userAndLocation.username&#x60;, &#x60;diskEncryption.fileVault2Enabled&#x60;, &#x60;purchasing.appleCareId&#x60;, &#x60;purchasing.lifeExpectancy&#x60;, &#x60;purchasing.purchased&#x60;, &#x60;purchasing.leased&#x60;, &#x60;purchasing.vendor&#x60;, &#x60;purchasing.warrantyDate&#x60;,  This param can be combined with paging and sorting. Example: &#x60;filter&#x3D;general.name&#x3D;&#x3D;\&quot;Orchard\&quot;&#x60;  | [default to &quot;&quot;]

### Return type

[**ComputerInventorySearchResultsV2**](ComputerInventorySearchResultsV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputersInventoryIdAttachmentsAttachmentIdDelete

> V2ComputersInventoryIdAttachmentsAttachmentIdDelete(ctx, id, attachmentId).Execute()

Remove attachment



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
	id := "id_example" // string | instance id of computer record
	attachmentId := "attachmentId_example" // string | instance id of attachment object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputerInventoryAPI.V2ComputersInventoryIdAttachmentsAttachmentIdDelete(context.Background(), id, attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V2ComputersInventoryIdAttachmentsAttachmentIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of computer record | 
**attachmentId** | **string** | instance id of attachment object | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputersInventoryIdAttachmentsAttachmentIdDeleteRequest struct via the builder pattern


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


## V2ComputersInventoryIdAttachmentsAttachmentIdGet

> *os.File V2ComputersInventoryIdAttachmentsAttachmentIdGet(ctx, id, attachmentId).Execute()

Download attachment file



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
	id := "id_example" // string | instance id of computer record
	attachmentId := "attachmentId_example" // string | instance id of attachment object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V2ComputersInventoryIdAttachmentsAttachmentIdGet(context.Background(), id, attachmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V2ComputersInventoryIdAttachmentsAttachmentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputersInventoryIdAttachmentsAttachmentIdGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V2ComputersInventoryIdAttachmentsAttachmentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of computer record | 
**attachmentId** | **string** | instance id of attachment object | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputersInventoryIdAttachmentsAttachmentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputersInventoryIdAttachmentsPost

> HrefResponse V2ComputersInventoryIdAttachmentsPost(ctx, id).File(file).Execute()

Upload attachment and assign to computer



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
	id := "id_example" // string | instance id of computer record
	file := os.NewFile(1234, "some_file") // *os.File | The file to upload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V2ComputersInventoryIdAttachmentsPost(context.Background(), id).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V2ComputersInventoryIdAttachmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputersInventoryIdAttachmentsPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V2ComputersInventoryIdAttachmentsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of computer record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputersInventoryIdAttachmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The file to upload | 

### Return type

[**HrefResponse**](HrefResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputersInventoryIdDelete

> V2ComputersInventoryIdDelete(ctx, id).Execute()

Remove specified Computer record



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
	id := "id_example" // string | instance id of computer record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputerInventoryAPI.V2ComputersInventoryIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V2ComputersInventoryIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of computer record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputersInventoryIdDeleteRequest struct via the builder pattern


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


## V2ComputersInventoryIdFilevaultGet

> ComputerInventoryFileVault V2ComputersInventoryIdFilevaultGet(ctx, id).Execute()

Return FileVault information for a specific computer



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
	id := "id_example" // string | instance id of computer record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V2ComputersInventoryIdFilevaultGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V2ComputersInventoryIdFilevaultGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputersInventoryIdFilevaultGet`: ComputerInventoryFileVault
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V2ComputersInventoryIdFilevaultGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of computer record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputersInventoryIdFilevaultGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComputerInventoryFileVault**](ComputerInventoryFileVault.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputersInventoryIdGet

> ComputerInventoryV2 V2ComputersInventoryIdGet(ctx, id).Section(section).Execute()

Return General section of a Computer



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
	id := "id_example" // string | instance id of computer record
	section := []openapiclient.ComputerSectionV2{openapiclient.ComputerSectionV2("GENERAL")} // []ComputerSectionV2 | section of computer details, if not specified, General section data is returned. Multiple section parameters are supported, e.g. section=general&section=hardware (optional) (default to {"GENERAL"})

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V2ComputersInventoryIdGet(context.Background(), id).Section(section).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V2ComputersInventoryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputersInventoryIdGet`: ComputerInventoryV2
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V2ComputersInventoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of computer record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputersInventoryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **section** | [**[]ComputerSectionV2**](ComputerSectionV2.md) | section of computer details, if not specified, General section data is returned. Multiple section parameters are supported, e.g. section&#x3D;general&amp;section&#x3D;hardware | [default to {&quot;GENERAL&quot;}]

### Return type

[**ComputerInventoryV2**](ComputerInventoryV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputersInventoryIdViewDeviceLockPinGet

> ComputerInventoryDeviceLockPinResponse V2ComputersInventoryIdViewDeviceLockPinGet(ctx, id).Execute()

Return a computer's Device Lock PIN



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
	id := "id_example" // string | instance id of computer record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V2ComputersInventoryIdViewDeviceLockPinGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V2ComputersInventoryIdViewDeviceLockPinGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputersInventoryIdViewDeviceLockPinGet`: ComputerInventoryDeviceLockPinResponse
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V2ComputersInventoryIdViewDeviceLockPinGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of computer record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputersInventoryIdViewDeviceLockPinGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComputerInventoryDeviceLockPinResponse**](ComputerInventoryDeviceLockPinResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputersInventoryIdViewRecoveryLockPasswordGet

> ComputerInventoryRecoveryLockPasswordResponse V2ComputersInventoryIdViewRecoveryLockPasswordGet(ctx, id).Execute()

Return a Computers Recovery Lock Password



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
	id := "id_example" // string | instance id of computer record

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V2ComputersInventoryIdViewRecoveryLockPasswordGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V2ComputersInventoryIdViewRecoveryLockPasswordGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputersInventoryIdViewRecoveryLockPasswordGet`: ComputerInventoryRecoveryLockPasswordResponse
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V2ComputersInventoryIdViewRecoveryLockPasswordGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of computer record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputersInventoryIdViewRecoveryLockPasswordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComputerInventoryRecoveryLockPasswordResponse**](ComputerInventoryRecoveryLockPasswordResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ComputersInventoryPost

> HrefResponse V2ComputersInventoryPost(ctx).ComputerInventoryCreateRequestV2(computerInventoryCreateRequestV2).Execute()

Create Computer Inventory record



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
	computerInventoryCreateRequestV2 := *openapiclient.NewComputerInventoryCreateRequestV2() // ComputerInventoryCreateRequestV2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputerInventoryAPI.V2ComputersInventoryPost(context.Background()).ComputerInventoryCreateRequestV2(computerInventoryCreateRequestV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputerInventoryAPI.V2ComputersInventoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2ComputersInventoryPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `ComputerInventoryAPI.V2ComputersInventoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2ComputersInventoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **computerInventoryCreateRequestV2** | [**ComputerInventoryCreateRequestV2**](ComputerInventoryCreateRequestV2.md) |  | 

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

