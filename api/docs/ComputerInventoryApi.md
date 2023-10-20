# \ComputerInventoryAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
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
[**V1ComputersInventoryIdViewRecoveryLockPasswordGet**](ComputerInventoryAPI.md#V1ComputersInventoryIdViewRecoveryLockPasswordGet) | **Get** /v1/computers-inventory/{id}/view-recovery-lock-password | Return a Computers Recovery Lock Password



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
    openapiclient "github.com/yohan460/go-jamf-api"
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

[Bearer](../README.md#Bearer)

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
    openapiclient "github.com/yohan460/go-jamf-api"
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

[Bearer](../README.md#Bearer)

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
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)

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
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]

### Return type

[**ComputerInventoryFileVaultSearchResults**](ComputerInventoryFileVaultSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

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
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    section := []openapiclient.ComputerSection{openapiclient.ComputerSection("GENERAL")} // []ComputerSection | section of computer details, if not specified, General section data is returned. Multiple section parameters are supported, e.g. section=GENERAL&section=HARDWARE (optional) (default to ["GENERAL"])
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: `property:asc/desc`. Default sort is `general.name:asc`. Multiple sort criteria are supported and must be separated with a comma.  Fields allowed in the sort: `general.name`, `udid`, `id`, `general.assetTag`, `general.jamfBinaryVersion`, `general.lastContactTime`, `general.lastEnrolledDate`, `general.lastCloudBackupDate`, `general.reportDate`, `general.remoteManagement.managementUsername`, `general.mdmCertificateExpiration`, `general.platform`, `hardware.make`, `hardware.model`, `operatingSystem.build`, `operatingSystem.supplementalBuildVersion`, `operatingSystem.rapidSecurityResponse`, `operatingSystem.name`, `operatingSystem.version`, `userAndLocation.realname`, `purchasing.lifeExpectancy`, `purchasing.warrantyDate`  Example: `sort=udid:desc,general.name:asc`.  (optional) (default to ["general.name:asc"])
    filter := "filter_example" // string | Query in the RSQL format, allowing to filter computer inventory collection. Default filter is empty query - returning all results for the requested page.  Fields allowed in the query: `general.name`, `udid`, `id`, `general.assetTag`, `general.barcode1`, `general.barcode2`, `general.enrolledViaAutomatedDeviceEnrollment`, `general.lastIpAddress`, `general.itunesStoreAccountActive`, `general.jamfBinaryVersion`, `general.lastContactTime`, `general.lastEnrolledDate`, `general.lastCloudBackupDate`, `general.reportDate`, `general.lastReportedIp`, `general.remoteManagement.managed`, `general.remoteManagement.managementUsername`, `general.mdmCapable.capable`, `general.mdmCertificateExpiration`, `general.platform`, `general.supervised`, `general.userApprovedMdm`, `general.declarativeDeviceManagementEnabled`,  `hardware.bleCapable`, `hardware.macAddress`, `hardware.make`, `hardware.model`, `hardware.modelIdentifier`, `hardware.serialNumber`, `hardware.supportsIosAppInstalls`,`hardware.appleSilicon`, `operatingSystem.activeDirectoryStatus`, `operatingSystem.fileVault2Status`, `operatingSystem.build`, `operatingSystem.supplementalBuildVersion`, `operatingSystem.rapidSecurityResponse`, `operatingSystem.name`, `operatingSystem.version`, `security.activationLockEnabled`, `security.recoveryLockEnabled`,`security.firewallEnabled`,`userAndLocation.buildingId`, `userAndLocation.departmentId`, `userAndLocation.email`, `userAndLocation.realname`, `userAndLocation.phone`, `userAndLocation.position`,`userAndLocation.room`, `userAndLocation.username`, `purchasing.appleCareId`, `purchasing.lifeExpectancy`, `purchasing.purchased`, `purchasing.leased`, `purchasing.vendor`, `purchasing.warrantyDate`,  This param can be combined with paging and sorting. Example: `filter=general.name==\"Orchard\"`  (optional) (default to "")

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
 **section** | [**[]ComputerSection**](ComputerSection.md) | section of computer details, if not specified, General section data is returned. Multiple section parameters are supported, e.g. section&#x3D;GENERAL&amp;section&#x3D;HARDWARE | [default to [&quot;GENERAL&quot;]]
 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: &#x60;property:asc/desc&#x60;. Default sort is &#x60;general.name:asc&#x60;. Multiple sort criteria are supported and must be separated with a comma.  Fields allowed in the sort: &#x60;general.name&#x60;, &#x60;udid&#x60;, &#x60;id&#x60;, &#x60;general.assetTag&#x60;, &#x60;general.jamfBinaryVersion&#x60;, &#x60;general.lastContactTime&#x60;, &#x60;general.lastEnrolledDate&#x60;, &#x60;general.lastCloudBackupDate&#x60;, &#x60;general.reportDate&#x60;, &#x60;general.remoteManagement.managementUsername&#x60;, &#x60;general.mdmCertificateExpiration&#x60;, &#x60;general.platform&#x60;, &#x60;hardware.make&#x60;, &#x60;hardware.model&#x60;, &#x60;operatingSystem.build&#x60;, &#x60;operatingSystem.supplementalBuildVersion&#x60;, &#x60;operatingSystem.rapidSecurityResponse&#x60;, &#x60;operatingSystem.name&#x60;, &#x60;operatingSystem.version&#x60;, &#x60;userAndLocation.realname&#x60;, &#x60;purchasing.lifeExpectancy&#x60;, &#x60;purchasing.warrantyDate&#x60;  Example: &#x60;sort&#x3D;udid:desc,general.name:asc&#x60;.  | [default to [&quot;general.name:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter computer inventory collection. Default filter is empty query - returning all results for the requested page.  Fields allowed in the query: &#x60;general.name&#x60;, &#x60;udid&#x60;, &#x60;id&#x60;, &#x60;general.assetTag&#x60;, &#x60;general.barcode1&#x60;, &#x60;general.barcode2&#x60;, &#x60;general.enrolledViaAutomatedDeviceEnrollment&#x60;, &#x60;general.lastIpAddress&#x60;, &#x60;general.itunesStoreAccountActive&#x60;, &#x60;general.jamfBinaryVersion&#x60;, &#x60;general.lastContactTime&#x60;, &#x60;general.lastEnrolledDate&#x60;, &#x60;general.lastCloudBackupDate&#x60;, &#x60;general.reportDate&#x60;, &#x60;general.lastReportedIp&#x60;, &#x60;general.remoteManagement.managed&#x60;, &#x60;general.remoteManagement.managementUsername&#x60;, &#x60;general.mdmCapable.capable&#x60;, &#x60;general.mdmCertificateExpiration&#x60;, &#x60;general.platform&#x60;, &#x60;general.supervised&#x60;, &#x60;general.userApprovedMdm&#x60;, &#x60;general.declarativeDeviceManagementEnabled&#x60;,  &#x60;hardware.bleCapable&#x60;, &#x60;hardware.macAddress&#x60;, &#x60;hardware.make&#x60;, &#x60;hardware.model&#x60;, &#x60;hardware.modelIdentifier&#x60;, &#x60;hardware.serialNumber&#x60;, &#x60;hardware.supportsIosAppInstalls&#x60;,&#x60;hardware.appleSilicon&#x60;, &#x60;operatingSystem.activeDirectoryStatus&#x60;, &#x60;operatingSystem.fileVault2Status&#x60;, &#x60;operatingSystem.build&#x60;, &#x60;operatingSystem.supplementalBuildVersion&#x60;, &#x60;operatingSystem.rapidSecurityResponse&#x60;, &#x60;operatingSystem.name&#x60;, &#x60;operatingSystem.version&#x60;, &#x60;security.activationLockEnabled&#x60;, &#x60;security.recoveryLockEnabled&#x60;,&#x60;security.firewallEnabled&#x60;,&#x60;userAndLocation.buildingId&#x60;, &#x60;userAndLocation.departmentId&#x60;, &#x60;userAndLocation.email&#x60;, &#x60;userAndLocation.realname&#x60;, &#x60;userAndLocation.phone&#x60;, &#x60;userAndLocation.position&#x60;,&#x60;userAndLocation.room&#x60;, &#x60;userAndLocation.username&#x60;, &#x60;purchasing.appleCareId&#x60;, &#x60;purchasing.lifeExpectancy&#x60;, &#x60;purchasing.purchased&#x60;, &#x60;purchasing.leased&#x60;, &#x60;purchasing.vendor&#x60;, &#x60;purchasing.warrantyDate&#x60;,  This param can be combined with paging and sorting. Example: &#x60;filter&#x3D;general.name&#x3D;&#x3D;\&quot;Orchard\&quot;&#x60;  | [default to &quot;&quot;]

### Return type

[**ComputerInventorySearchResults**](ComputerInventorySearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

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
    openapiclient "github.com/yohan460/go-jamf-api"
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

[Bearer](../README.md#Bearer)

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
    openapiclient "github.com/yohan460/go-jamf-api"
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

[Bearer](../README.md#Bearer)

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
    openapiclient "github.com/yohan460/go-jamf-api"
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

[Bearer](../README.md#Bearer)

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
    openapiclient "github.com/yohan460/go-jamf-api"
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

[Bearer](../README.md#Bearer)

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
    openapiclient "github.com/yohan460/go-jamf-api"
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

[Bearer](../README.md#Bearer)

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
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    id := "id_example" // string | instance id of computer record
    section := []openapiclient.ComputerSection{openapiclient.ComputerSection("GENERAL")} // []ComputerSection | section of computer details, if not specified, General section data is returned. Multiple section parameters are supported, e.g. section=general&section=hardware (optional) (default to ["GENERAL"])

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

 **section** | [**[]ComputerSection**](ComputerSection.md) | section of computer details, if not specified, General section data is returned. Multiple section parameters are supported, e.g. section&#x3D;general&amp;section&#x3D;hardware | [default to [&quot;GENERAL&quot;]]

### Return type

[**ComputerInventory**](ComputerInventory.md)

### Authorization

[Bearer](../README.md#Bearer)

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
    openapiclient "github.com/yohan460/go-jamf-api"
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

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

