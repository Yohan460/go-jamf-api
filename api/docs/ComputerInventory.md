# ComputerInventory

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

### NewComputerInventory

`func NewComputerInventory() *ComputerInventory`

NewComputerInventory instantiates a new ComputerInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerInventoryWithDefaults

`func NewComputerInventoryWithDefaults() *ComputerInventory`

NewComputerInventoryWithDefaults instantiates a new ComputerInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComputerInventory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputerInventory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputerInventory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComputerInventory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUdid

`func (o *ComputerInventory) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *ComputerInventory) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *ComputerInventory) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *ComputerInventory) HasUdid() bool`

HasUdid returns a boolean if a field has been set.

### GetGeneral

`func (o *ComputerInventory) GetGeneral() ComputerGeneral`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *ComputerInventory) GetGeneralOk() (*ComputerGeneral, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *ComputerInventory) SetGeneral(v ComputerGeneral)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *ComputerInventory) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetDiskEncryption

`func (o *ComputerInventory) GetDiskEncryption() ComputerDiskEncryption`

GetDiskEncryption returns the DiskEncryption field if non-nil, zero value otherwise.

### GetDiskEncryptionOk

`func (o *ComputerInventory) GetDiskEncryptionOk() (*ComputerDiskEncryption, bool)`

GetDiskEncryptionOk returns a tuple with the DiskEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryption

`func (o *ComputerInventory) SetDiskEncryption(v ComputerDiskEncryption)`

SetDiskEncryption sets DiskEncryption field to given value.

### HasDiskEncryption

`func (o *ComputerInventory) HasDiskEncryption() bool`

HasDiskEncryption returns a boolean if a field has been set.

### GetPurchasing

`func (o *ComputerInventory) GetPurchasing() ComputerPurchase`

GetPurchasing returns the Purchasing field if non-nil, zero value otherwise.

### GetPurchasingOk

`func (o *ComputerInventory) GetPurchasingOk() (*ComputerPurchase, bool)`

GetPurchasingOk returns a tuple with the Purchasing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasing

`func (o *ComputerInventory) SetPurchasing(v ComputerPurchase)`

SetPurchasing sets Purchasing field to given value.

### HasPurchasing

`func (o *ComputerInventory) HasPurchasing() bool`

HasPurchasing returns a boolean if a field has been set.

### GetApplications

`func (o *ComputerInventory) GetApplications() []ComputerApplication`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ComputerInventory) GetApplicationsOk() (*[]ComputerApplication, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ComputerInventory) SetApplications(v []ComputerApplication)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *ComputerInventory) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetStorage

`func (o *ComputerInventory) GetStorage() ComputerStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ComputerInventory) GetStorageOk() (*ComputerStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ComputerInventory) SetStorage(v ComputerStorage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ComputerInventory) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetUserAndLocation

`func (o *ComputerInventory) GetUserAndLocation() ComputerUserAndLocation`

GetUserAndLocation returns the UserAndLocation field if non-nil, zero value otherwise.

### GetUserAndLocationOk

`func (o *ComputerInventory) GetUserAndLocationOk() (*ComputerUserAndLocation, bool)`

GetUserAndLocationOk returns a tuple with the UserAndLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAndLocation

`func (o *ComputerInventory) SetUserAndLocation(v ComputerUserAndLocation)`

SetUserAndLocation sets UserAndLocation field to given value.

### HasUserAndLocation

`func (o *ComputerInventory) HasUserAndLocation() bool`

HasUserAndLocation returns a boolean if a field has been set.

### GetConfigurationProfiles

`func (o *ComputerInventory) GetConfigurationProfiles() []ComputerConfigurationProfile`

GetConfigurationProfiles returns the ConfigurationProfiles field if non-nil, zero value otherwise.

### GetConfigurationProfilesOk

`func (o *ComputerInventory) GetConfigurationProfilesOk() (*[]ComputerConfigurationProfile, bool)`

GetConfigurationProfilesOk returns a tuple with the ConfigurationProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationProfiles

`func (o *ComputerInventory) SetConfigurationProfiles(v []ComputerConfigurationProfile)`

SetConfigurationProfiles sets ConfigurationProfiles field to given value.

### HasConfigurationProfiles

`func (o *ComputerInventory) HasConfigurationProfiles() bool`

HasConfigurationProfiles returns a boolean if a field has been set.

### GetPrinters

`func (o *ComputerInventory) GetPrinters() []ComputerPrinter`

GetPrinters returns the Printers field if non-nil, zero value otherwise.

### GetPrintersOk

`func (o *ComputerInventory) GetPrintersOk() (*[]ComputerPrinter, bool)`

GetPrintersOk returns a tuple with the Printers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinters

`func (o *ComputerInventory) SetPrinters(v []ComputerPrinter)`

SetPrinters sets Printers field to given value.

### HasPrinters

`func (o *ComputerInventory) HasPrinters() bool`

HasPrinters returns a boolean if a field has been set.

### GetServices

`func (o *ComputerInventory) GetServices() []ComputerService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ComputerInventory) GetServicesOk() (*[]ComputerService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ComputerInventory) SetServices(v []ComputerService)`

SetServices sets Services field to given value.

### HasServices

`func (o *ComputerInventory) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetHardware

`func (o *ComputerInventory) GetHardware() ComputerHardware`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *ComputerInventory) GetHardwareOk() (*ComputerHardware, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *ComputerInventory) SetHardware(v ComputerHardware)`

SetHardware sets Hardware field to given value.

### HasHardware

`func (o *ComputerInventory) HasHardware() bool`

HasHardware returns a boolean if a field has been set.

### GetLocalUserAccounts

`func (o *ComputerInventory) GetLocalUserAccounts() []ComputerLocalUserAccount`

GetLocalUserAccounts returns the LocalUserAccounts field if non-nil, zero value otherwise.

### GetLocalUserAccountsOk

`func (o *ComputerInventory) GetLocalUserAccountsOk() (*[]ComputerLocalUserAccount, bool)`

GetLocalUserAccountsOk returns a tuple with the LocalUserAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserAccounts

`func (o *ComputerInventory) SetLocalUserAccounts(v []ComputerLocalUserAccount)`

SetLocalUserAccounts sets LocalUserAccounts field to given value.

### HasLocalUserAccounts

`func (o *ComputerInventory) HasLocalUserAccounts() bool`

HasLocalUserAccounts returns a boolean if a field has been set.

### GetCertificates

`func (o *ComputerInventory) GetCertificates() []ComputerCertificate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *ComputerInventory) GetCertificatesOk() (*[]ComputerCertificate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *ComputerInventory) SetCertificates(v []ComputerCertificate)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *ComputerInventory) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetAttachments

`func (o *ComputerInventory) GetAttachments() []ComputerAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ComputerInventory) GetAttachmentsOk() (*[]ComputerAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ComputerInventory) SetAttachments(v []ComputerAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ComputerInventory) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetPlugins

`func (o *ComputerInventory) GetPlugins() []ComputerPlugin`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *ComputerInventory) GetPluginsOk() (*[]ComputerPlugin, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *ComputerInventory) SetPlugins(v []ComputerPlugin)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *ComputerInventory) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### GetPackageReceipts

`func (o *ComputerInventory) GetPackageReceipts() ComputerPackageReceipts`

GetPackageReceipts returns the PackageReceipts field if non-nil, zero value otherwise.

### GetPackageReceiptsOk

`func (o *ComputerInventory) GetPackageReceiptsOk() (*ComputerPackageReceipts, bool)`

GetPackageReceiptsOk returns a tuple with the PackageReceipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageReceipts

`func (o *ComputerInventory) SetPackageReceipts(v ComputerPackageReceipts)`

SetPackageReceipts sets PackageReceipts field to given value.

### HasPackageReceipts

`func (o *ComputerInventory) HasPackageReceipts() bool`

HasPackageReceipts returns a boolean if a field has been set.

### GetFonts

`func (o *ComputerInventory) GetFonts() []ComputerFont`

GetFonts returns the Fonts field if non-nil, zero value otherwise.

### GetFontsOk

`func (o *ComputerInventory) GetFontsOk() (*[]ComputerFont, bool)`

GetFontsOk returns a tuple with the Fonts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFonts

`func (o *ComputerInventory) SetFonts(v []ComputerFont)`

SetFonts sets Fonts field to given value.

### HasFonts

`func (o *ComputerInventory) HasFonts() bool`

HasFonts returns a boolean if a field has been set.

### GetSecurity

`func (o *ComputerInventory) GetSecurity() ComputerSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *ComputerInventory) GetSecurityOk() (*ComputerSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *ComputerInventory) SetSecurity(v ComputerSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *ComputerInventory) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *ComputerInventory) GetOperatingSystem() ComputerOperatingSystem`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *ComputerInventory) GetOperatingSystemOk() (*ComputerOperatingSystem, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *ComputerInventory) SetOperatingSystem(v ComputerOperatingSystem)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *ComputerInventory) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetLicensedSoftware

`func (o *ComputerInventory) GetLicensedSoftware() []ComputerLicensedSoftware`

GetLicensedSoftware returns the LicensedSoftware field if non-nil, zero value otherwise.

### GetLicensedSoftwareOk

`func (o *ComputerInventory) GetLicensedSoftwareOk() (*[]ComputerLicensedSoftware, bool)`

GetLicensedSoftwareOk returns a tuple with the LicensedSoftware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedSoftware

`func (o *ComputerInventory) SetLicensedSoftware(v []ComputerLicensedSoftware)`

SetLicensedSoftware sets LicensedSoftware field to given value.

### HasLicensedSoftware

`func (o *ComputerInventory) HasLicensedSoftware() bool`

HasLicensedSoftware returns a boolean if a field has been set.

### GetIbeacons

`func (o *ComputerInventory) GetIbeacons() []ComputerIbeacon`

GetIbeacons returns the Ibeacons field if non-nil, zero value otherwise.

### GetIbeaconsOk

`func (o *ComputerInventory) GetIbeaconsOk() (*[]ComputerIbeacon, bool)`

GetIbeaconsOk returns a tuple with the Ibeacons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeacons

`func (o *ComputerInventory) SetIbeacons(v []ComputerIbeacon)`

SetIbeacons sets Ibeacons field to given value.

### HasIbeacons

`func (o *ComputerInventory) HasIbeacons() bool`

HasIbeacons returns a boolean if a field has been set.

### GetSoftwareUpdates

`func (o *ComputerInventory) GetSoftwareUpdates() []ComputerSoftwareUpdate`

GetSoftwareUpdates returns the SoftwareUpdates field if non-nil, zero value otherwise.

### GetSoftwareUpdatesOk

`func (o *ComputerInventory) GetSoftwareUpdatesOk() (*[]ComputerSoftwareUpdate, bool)`

GetSoftwareUpdatesOk returns a tuple with the SoftwareUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdates

`func (o *ComputerInventory) SetSoftwareUpdates(v []ComputerSoftwareUpdate)`

SetSoftwareUpdates sets SoftwareUpdates field to given value.

### HasSoftwareUpdates

`func (o *ComputerInventory) HasSoftwareUpdates() bool`

HasSoftwareUpdates returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *ComputerInventory) GetExtensionAttributes() []ComputerExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *ComputerInventory) GetExtensionAttributesOk() (*[]ComputerExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *ComputerInventory) SetExtensionAttributes(v []ComputerExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *ComputerInventory) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.

### GetContentCaching

`func (o *ComputerInventory) GetContentCaching() ComputerContentCaching`

GetContentCaching returns the ContentCaching field if non-nil, zero value otherwise.

### GetContentCachingOk

`func (o *ComputerInventory) GetContentCachingOk() (*ComputerContentCaching, bool)`

GetContentCachingOk returns a tuple with the ContentCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentCaching

`func (o *ComputerInventory) SetContentCaching(v ComputerContentCaching)`

SetContentCaching sets ContentCaching field to given value.

### HasContentCaching

`func (o *ComputerInventory) HasContentCaching() bool`

HasContentCaching returns a boolean if a field has been set.

### GetGroupMemberships

`func (o *ComputerInventory) GetGroupMemberships() []GroupMembership`

GetGroupMemberships returns the GroupMemberships field if non-nil, zero value otherwise.

### GetGroupMembershipsOk

`func (o *ComputerInventory) GetGroupMembershipsOk() (*[]GroupMembership, bool)`

GetGroupMembershipsOk returns a tuple with the GroupMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberships

`func (o *ComputerInventory) SetGroupMemberships(v []GroupMembership)`

SetGroupMemberships sets GroupMemberships field to given value.

### HasGroupMemberships

`func (o *ComputerInventory) HasGroupMemberships() bool`

HasGroupMemberships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


