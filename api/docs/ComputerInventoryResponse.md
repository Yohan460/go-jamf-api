# ComputerInventoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Udid** | Pointer to **string** |  | [optional] 
**General** | Pointer to [**ComputerGeneral**](ComputerGeneral.md) |  | [optional] 
**DiskEncryption** | Pointer to [**ComputerDiskEncryption**](ComputerDiskEncryption.md) |  | [optional] 
**Purchasing** | Pointer to [**ComputerPurchase**](ComputerPurchase.md) |  | [optional] 
**Applications** | Pointer to [**[]ComputerApplication**](ComputerApplication.md) |  | [optional] 
**Storage** | Pointer to [**ComputerStorage**](ComputerStorage.md) |  | [optional] 
**UserAndLocation** | Pointer to [**ComputerUserAndLocation**](ComputerUserAndLocation.md) |  | [optional] 
**ConfigurationProfiles** | Pointer to [**[]ComputerConfigurationProfile**](ComputerConfigurationProfile.md) |  | [optional] 
**Printers** | Pointer to [**[]ComputerPrinter**](ComputerPrinter.md) |  | [optional] 
**Services** | Pointer to [**[]ComputerService**](ComputerService.md) |  | [optional] 
**Hardware** | Pointer to [**ComputerHardware**](ComputerHardware.md) |  | [optional] 
**LocalUserAccounts** | Pointer to [**[]ComputerLocalUserAccount**](ComputerLocalUserAccount.md) |  | [optional] 
**Certificates** | Pointer to [**[]ComputerCertificate**](ComputerCertificate.md) |  | [optional] 
**Attachments** | Pointer to [**[]ComputerAttachment**](ComputerAttachment.md) |  | [optional] 
**Plugins** | Pointer to [**[]ComputerPlugin**](ComputerPlugin.md) |  | [optional] 
**PackageReceipts** | Pointer to [**ComputerPackageReceipts**](ComputerPackageReceipts.md) |  | [optional] 
**Fonts** | Pointer to [**[]ComputerFont**](ComputerFont.md) |  | [optional] 
**Security** | Pointer to [**ComputerSecurity**](ComputerSecurity.md) |  | [optional] 
**OperatingSystem** | Pointer to [**ComputerOperatingSystem**](ComputerOperatingSystem.md) |  | [optional] 
**LicensedSoftware** | Pointer to [**[]ComputerLicensedSoftware**](ComputerLicensedSoftware.md) |  | [optional] 
**Ibeacons** | Pointer to [**[]ComputerIbeacon**](ComputerIbeacon.md) |  | [optional] 
**SoftwareUpdates** | Pointer to [**[]ComputerSoftwareUpdate**](ComputerSoftwareUpdate.md) |  | [optional] 
**ExtensionAttributes** | Pointer to [**[]ComputerExtensionAttribute**](ComputerExtensionAttribute.md) |  | [optional] 
**ContentCaching** | Pointer to [**ComputerContentCaching**](ComputerContentCaching.md) |  | [optional] 
**GroupMemberships** | Pointer to [**[]GroupMembership**](GroupMembership.md) |  | [optional] 

## Methods

### NewComputerInventoryResponse

`func NewComputerInventoryResponse() *ComputerInventoryResponse`

NewComputerInventoryResponse instantiates a new ComputerInventoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerInventoryResponseWithDefaults

`func NewComputerInventoryResponseWithDefaults() *ComputerInventoryResponse`

NewComputerInventoryResponseWithDefaults instantiates a new ComputerInventoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComputerInventoryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputerInventoryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputerInventoryResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComputerInventoryResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUdid

`func (o *ComputerInventoryResponse) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *ComputerInventoryResponse) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *ComputerInventoryResponse) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *ComputerInventoryResponse) HasUdid() bool`

HasUdid returns a boolean if a field has been set.

### GetGeneral

`func (o *ComputerInventoryResponse) GetGeneral() ComputerGeneral`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *ComputerInventoryResponse) GetGeneralOk() (*ComputerGeneral, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *ComputerInventoryResponse) SetGeneral(v ComputerGeneral)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *ComputerInventoryResponse) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetDiskEncryption

`func (o *ComputerInventoryResponse) GetDiskEncryption() ComputerDiskEncryption`

GetDiskEncryption returns the DiskEncryption field if non-nil, zero value otherwise.

### GetDiskEncryptionOk

`func (o *ComputerInventoryResponse) GetDiskEncryptionOk() (*ComputerDiskEncryption, bool)`

GetDiskEncryptionOk returns a tuple with the DiskEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryption

`func (o *ComputerInventoryResponse) SetDiskEncryption(v ComputerDiskEncryption)`

SetDiskEncryption sets DiskEncryption field to given value.

### HasDiskEncryption

`func (o *ComputerInventoryResponse) HasDiskEncryption() bool`

HasDiskEncryption returns a boolean if a field has been set.

### GetPurchasing

`func (o *ComputerInventoryResponse) GetPurchasing() ComputerPurchase`

GetPurchasing returns the Purchasing field if non-nil, zero value otherwise.

### GetPurchasingOk

`func (o *ComputerInventoryResponse) GetPurchasingOk() (*ComputerPurchase, bool)`

GetPurchasingOk returns a tuple with the Purchasing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasing

`func (o *ComputerInventoryResponse) SetPurchasing(v ComputerPurchase)`

SetPurchasing sets Purchasing field to given value.

### HasPurchasing

`func (o *ComputerInventoryResponse) HasPurchasing() bool`

HasPurchasing returns a boolean if a field has been set.

### GetApplications

`func (o *ComputerInventoryResponse) GetApplications() []ComputerApplication`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ComputerInventoryResponse) GetApplicationsOk() (*[]ComputerApplication, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ComputerInventoryResponse) SetApplications(v []ComputerApplication)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *ComputerInventoryResponse) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetStorage

`func (o *ComputerInventoryResponse) GetStorage() ComputerStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ComputerInventoryResponse) GetStorageOk() (*ComputerStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ComputerInventoryResponse) SetStorage(v ComputerStorage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ComputerInventoryResponse) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetUserAndLocation

`func (o *ComputerInventoryResponse) GetUserAndLocation() ComputerUserAndLocation`

GetUserAndLocation returns the UserAndLocation field if non-nil, zero value otherwise.

### GetUserAndLocationOk

`func (o *ComputerInventoryResponse) GetUserAndLocationOk() (*ComputerUserAndLocation, bool)`

GetUserAndLocationOk returns a tuple with the UserAndLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAndLocation

`func (o *ComputerInventoryResponse) SetUserAndLocation(v ComputerUserAndLocation)`

SetUserAndLocation sets UserAndLocation field to given value.

### HasUserAndLocation

`func (o *ComputerInventoryResponse) HasUserAndLocation() bool`

HasUserAndLocation returns a boolean if a field has been set.

### GetConfigurationProfiles

`func (o *ComputerInventoryResponse) GetConfigurationProfiles() []ComputerConfigurationProfile`

GetConfigurationProfiles returns the ConfigurationProfiles field if non-nil, zero value otherwise.

### GetConfigurationProfilesOk

`func (o *ComputerInventoryResponse) GetConfigurationProfilesOk() (*[]ComputerConfigurationProfile, bool)`

GetConfigurationProfilesOk returns a tuple with the ConfigurationProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationProfiles

`func (o *ComputerInventoryResponse) SetConfigurationProfiles(v []ComputerConfigurationProfile)`

SetConfigurationProfiles sets ConfigurationProfiles field to given value.

### HasConfigurationProfiles

`func (o *ComputerInventoryResponse) HasConfigurationProfiles() bool`

HasConfigurationProfiles returns a boolean if a field has been set.

### GetPrinters

`func (o *ComputerInventoryResponse) GetPrinters() []ComputerPrinter`

GetPrinters returns the Printers field if non-nil, zero value otherwise.

### GetPrintersOk

`func (o *ComputerInventoryResponse) GetPrintersOk() (*[]ComputerPrinter, bool)`

GetPrintersOk returns a tuple with the Printers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinters

`func (o *ComputerInventoryResponse) SetPrinters(v []ComputerPrinter)`

SetPrinters sets Printers field to given value.

### HasPrinters

`func (o *ComputerInventoryResponse) HasPrinters() bool`

HasPrinters returns a boolean if a field has been set.

### GetServices

`func (o *ComputerInventoryResponse) GetServices() []ComputerService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ComputerInventoryResponse) GetServicesOk() (*[]ComputerService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ComputerInventoryResponse) SetServices(v []ComputerService)`

SetServices sets Services field to given value.

### HasServices

`func (o *ComputerInventoryResponse) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetHardware

`func (o *ComputerInventoryResponse) GetHardware() ComputerHardware`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *ComputerInventoryResponse) GetHardwareOk() (*ComputerHardware, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *ComputerInventoryResponse) SetHardware(v ComputerHardware)`

SetHardware sets Hardware field to given value.

### HasHardware

`func (o *ComputerInventoryResponse) HasHardware() bool`

HasHardware returns a boolean if a field has been set.

### GetLocalUserAccounts

`func (o *ComputerInventoryResponse) GetLocalUserAccounts() []ComputerLocalUserAccount`

GetLocalUserAccounts returns the LocalUserAccounts field if non-nil, zero value otherwise.

### GetLocalUserAccountsOk

`func (o *ComputerInventoryResponse) GetLocalUserAccountsOk() (*[]ComputerLocalUserAccount, bool)`

GetLocalUserAccountsOk returns a tuple with the LocalUserAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserAccounts

`func (o *ComputerInventoryResponse) SetLocalUserAccounts(v []ComputerLocalUserAccount)`

SetLocalUserAccounts sets LocalUserAccounts field to given value.

### HasLocalUserAccounts

`func (o *ComputerInventoryResponse) HasLocalUserAccounts() bool`

HasLocalUserAccounts returns a boolean if a field has been set.

### GetCertificates

`func (o *ComputerInventoryResponse) GetCertificates() []ComputerCertificate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *ComputerInventoryResponse) GetCertificatesOk() (*[]ComputerCertificate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *ComputerInventoryResponse) SetCertificates(v []ComputerCertificate)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *ComputerInventoryResponse) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetAttachments

`func (o *ComputerInventoryResponse) GetAttachments() []ComputerAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ComputerInventoryResponse) GetAttachmentsOk() (*[]ComputerAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ComputerInventoryResponse) SetAttachments(v []ComputerAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ComputerInventoryResponse) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetPlugins

`func (o *ComputerInventoryResponse) GetPlugins() []ComputerPlugin`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *ComputerInventoryResponse) GetPluginsOk() (*[]ComputerPlugin, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *ComputerInventoryResponse) SetPlugins(v []ComputerPlugin)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *ComputerInventoryResponse) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### GetPackageReceipts

`func (o *ComputerInventoryResponse) GetPackageReceipts() ComputerPackageReceipts`

GetPackageReceipts returns the PackageReceipts field if non-nil, zero value otherwise.

### GetPackageReceiptsOk

`func (o *ComputerInventoryResponse) GetPackageReceiptsOk() (*ComputerPackageReceipts, bool)`

GetPackageReceiptsOk returns a tuple with the PackageReceipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageReceipts

`func (o *ComputerInventoryResponse) SetPackageReceipts(v ComputerPackageReceipts)`

SetPackageReceipts sets PackageReceipts field to given value.

### HasPackageReceipts

`func (o *ComputerInventoryResponse) HasPackageReceipts() bool`

HasPackageReceipts returns a boolean if a field has been set.

### GetFonts

`func (o *ComputerInventoryResponse) GetFonts() []ComputerFont`

GetFonts returns the Fonts field if non-nil, zero value otherwise.

### GetFontsOk

`func (o *ComputerInventoryResponse) GetFontsOk() (*[]ComputerFont, bool)`

GetFontsOk returns a tuple with the Fonts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFonts

`func (o *ComputerInventoryResponse) SetFonts(v []ComputerFont)`

SetFonts sets Fonts field to given value.

### HasFonts

`func (o *ComputerInventoryResponse) HasFonts() bool`

HasFonts returns a boolean if a field has been set.

### GetSecurity

`func (o *ComputerInventoryResponse) GetSecurity() ComputerSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *ComputerInventoryResponse) GetSecurityOk() (*ComputerSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *ComputerInventoryResponse) SetSecurity(v ComputerSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *ComputerInventoryResponse) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *ComputerInventoryResponse) GetOperatingSystem() ComputerOperatingSystem`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *ComputerInventoryResponse) GetOperatingSystemOk() (*ComputerOperatingSystem, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *ComputerInventoryResponse) SetOperatingSystem(v ComputerOperatingSystem)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *ComputerInventoryResponse) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetLicensedSoftware

`func (o *ComputerInventoryResponse) GetLicensedSoftware() []ComputerLicensedSoftware`

GetLicensedSoftware returns the LicensedSoftware field if non-nil, zero value otherwise.

### GetLicensedSoftwareOk

`func (o *ComputerInventoryResponse) GetLicensedSoftwareOk() (*[]ComputerLicensedSoftware, bool)`

GetLicensedSoftwareOk returns a tuple with the LicensedSoftware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedSoftware

`func (o *ComputerInventoryResponse) SetLicensedSoftware(v []ComputerLicensedSoftware)`

SetLicensedSoftware sets LicensedSoftware field to given value.

### HasLicensedSoftware

`func (o *ComputerInventoryResponse) HasLicensedSoftware() bool`

HasLicensedSoftware returns a boolean if a field has been set.

### GetIbeacons

`func (o *ComputerInventoryResponse) GetIbeacons() []ComputerIbeacon`

GetIbeacons returns the Ibeacons field if non-nil, zero value otherwise.

### GetIbeaconsOk

`func (o *ComputerInventoryResponse) GetIbeaconsOk() (*[]ComputerIbeacon, bool)`

GetIbeaconsOk returns a tuple with the Ibeacons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeacons

`func (o *ComputerInventoryResponse) SetIbeacons(v []ComputerIbeacon)`

SetIbeacons sets Ibeacons field to given value.

### HasIbeacons

`func (o *ComputerInventoryResponse) HasIbeacons() bool`

HasIbeacons returns a boolean if a field has been set.

### GetSoftwareUpdates

`func (o *ComputerInventoryResponse) GetSoftwareUpdates() []ComputerSoftwareUpdate`

GetSoftwareUpdates returns the SoftwareUpdates field if non-nil, zero value otherwise.

### GetSoftwareUpdatesOk

`func (o *ComputerInventoryResponse) GetSoftwareUpdatesOk() (*[]ComputerSoftwareUpdate, bool)`

GetSoftwareUpdatesOk returns a tuple with the SoftwareUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdates

`func (o *ComputerInventoryResponse) SetSoftwareUpdates(v []ComputerSoftwareUpdate)`

SetSoftwareUpdates sets SoftwareUpdates field to given value.

### HasSoftwareUpdates

`func (o *ComputerInventoryResponse) HasSoftwareUpdates() bool`

HasSoftwareUpdates returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *ComputerInventoryResponse) GetExtensionAttributes() []ComputerExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *ComputerInventoryResponse) GetExtensionAttributesOk() (*[]ComputerExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *ComputerInventoryResponse) SetExtensionAttributes(v []ComputerExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *ComputerInventoryResponse) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.

### GetContentCaching

`func (o *ComputerInventoryResponse) GetContentCaching() ComputerContentCaching`

GetContentCaching returns the ContentCaching field if non-nil, zero value otherwise.

### GetContentCachingOk

`func (o *ComputerInventoryResponse) GetContentCachingOk() (*ComputerContentCaching, bool)`

GetContentCachingOk returns a tuple with the ContentCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentCaching

`func (o *ComputerInventoryResponse) SetContentCaching(v ComputerContentCaching)`

SetContentCaching sets ContentCaching field to given value.

### HasContentCaching

`func (o *ComputerInventoryResponse) HasContentCaching() bool`

HasContentCaching returns a boolean if a field has been set.

### GetGroupMemberships

`func (o *ComputerInventoryResponse) GetGroupMemberships() []GroupMembership`

GetGroupMemberships returns the GroupMemberships field if non-nil, zero value otherwise.

### GetGroupMembershipsOk

`func (o *ComputerInventoryResponse) GetGroupMembershipsOk() (*[]GroupMembership, bool)`

GetGroupMembershipsOk returns a tuple with the GroupMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberships

`func (o *ComputerInventoryResponse) SetGroupMemberships(v []GroupMembership)`

SetGroupMemberships sets GroupMemberships field to given value.

### HasGroupMemberships

`func (o *ComputerInventoryResponse) HasGroupMemberships() bool`

HasGroupMemberships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


