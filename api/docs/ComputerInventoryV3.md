# ComputerInventoryV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Udid** | Pointer to **string** |  | [optional] 
**General** | Pointer to [**ComputerGeneral**](ComputerGeneral.md) |  | [optional] 
**DiskEncryption** | Pointer to [**ComputerDiskEncryption**](ComputerDiskEncryption.md) |  | [optional] 
**Purchasing** | Pointer to [**ComputerPurchase**](ComputerPurchase.md) |  | [optional] 
**Applications** | Pointer to [**[]ComputerApplicationV3**](ComputerApplicationV3.md) |  | [optional] 
**Storage** | Pointer to [**ComputerStorage**](ComputerStorage.md) |  | [optional] 
**UserAndLocation** | Pointer to [**ComputerUserAndLocation**](ComputerUserAndLocation.md) |  | [optional] 
**ConfigurationProfiles** | Pointer to [**[]ComputerConfigurationProfile**](ComputerConfigurationProfile.md) |  | [optional] 
**Printers** | Pointer to [**[]ComputerPrinter**](ComputerPrinter.md) |  | [optional] 
**Services** | Pointer to [**[]ComputerService**](ComputerService.md) |  | [optional] 
**Hardware** | Pointer to [**ComputerHardware**](ComputerHardware.md) |  | [optional] 
**LocalUserAccounts** | Pointer to [**[]ComputerLocalUserAccount**](ComputerLocalUserAccount.md) |  | [optional] 
**Certificates** | Pointer to [**[]ComputerCertificate**](ComputerCertificate.md) |  | [optional] 
**Attachments** | Pointer to [**[]ComputerAttachment**](ComputerAttachment.md) |  | [optional] 
**PackageReceipts** | Pointer to [**ComputerPackageReceipts**](ComputerPackageReceipts.md) |  | [optional] 
**Security** | Pointer to [**ComputerSecurity**](ComputerSecurity.md) |  | [optional] 
**OperatingSystem** | Pointer to [**ComputerOperatingSystem**](ComputerOperatingSystem.md) |  | [optional] 
**LicensedSoftware** | Pointer to [**[]ComputerLicensedSoftware**](ComputerLicensedSoftware.md) |  | [optional] 
**Ibeacons** | Pointer to [**[]ComputerIbeacon**](ComputerIbeacon.md) |  | [optional] 
**SoftwareUpdates** | Pointer to [**[]ComputerSoftwareUpdate**](ComputerSoftwareUpdate.md) |  | [optional] 
**ExtensionAttributes** | Pointer to [**[]ComputerExtensionAttribute**](ComputerExtensionAttribute.md) |  | [optional] 
**ContentCaching** | Pointer to [**ComputerContentCaching**](ComputerContentCaching.md) |  | [optional] 
**GroupMemberships** | Pointer to [**[]GroupMembership**](GroupMembership.md) |  | [optional] 

## Methods

### NewComputerInventoryV3

`func NewComputerInventoryV3() *ComputerInventoryV3`

NewComputerInventoryV3 instantiates a new ComputerInventoryV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerInventoryV3WithDefaults

`func NewComputerInventoryV3WithDefaults() *ComputerInventoryV3`

NewComputerInventoryV3WithDefaults instantiates a new ComputerInventoryV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComputerInventoryV3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputerInventoryV3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputerInventoryV3) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComputerInventoryV3) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUdid

`func (o *ComputerInventoryV3) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *ComputerInventoryV3) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *ComputerInventoryV3) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *ComputerInventoryV3) HasUdid() bool`

HasUdid returns a boolean if a field has been set.

### GetGeneral

`func (o *ComputerInventoryV3) GetGeneral() ComputerGeneral`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *ComputerInventoryV3) GetGeneralOk() (*ComputerGeneral, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *ComputerInventoryV3) SetGeneral(v ComputerGeneral)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *ComputerInventoryV3) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetDiskEncryption

`func (o *ComputerInventoryV3) GetDiskEncryption() ComputerDiskEncryption`

GetDiskEncryption returns the DiskEncryption field if non-nil, zero value otherwise.

### GetDiskEncryptionOk

`func (o *ComputerInventoryV3) GetDiskEncryptionOk() (*ComputerDiskEncryption, bool)`

GetDiskEncryptionOk returns a tuple with the DiskEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryption

`func (o *ComputerInventoryV3) SetDiskEncryption(v ComputerDiskEncryption)`

SetDiskEncryption sets DiskEncryption field to given value.

### HasDiskEncryption

`func (o *ComputerInventoryV3) HasDiskEncryption() bool`

HasDiskEncryption returns a boolean if a field has been set.

### GetPurchasing

`func (o *ComputerInventoryV3) GetPurchasing() ComputerPurchase`

GetPurchasing returns the Purchasing field if non-nil, zero value otherwise.

### GetPurchasingOk

`func (o *ComputerInventoryV3) GetPurchasingOk() (*ComputerPurchase, bool)`

GetPurchasingOk returns a tuple with the Purchasing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasing

`func (o *ComputerInventoryV3) SetPurchasing(v ComputerPurchase)`

SetPurchasing sets Purchasing field to given value.

### HasPurchasing

`func (o *ComputerInventoryV3) HasPurchasing() bool`

HasPurchasing returns a boolean if a field has been set.

### GetApplications

`func (o *ComputerInventoryV3) GetApplications() []ComputerApplicationV3`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ComputerInventoryV3) GetApplicationsOk() (*[]ComputerApplicationV3, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ComputerInventoryV3) SetApplications(v []ComputerApplicationV3)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *ComputerInventoryV3) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetStorage

`func (o *ComputerInventoryV3) GetStorage() ComputerStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ComputerInventoryV3) GetStorageOk() (*ComputerStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ComputerInventoryV3) SetStorage(v ComputerStorage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ComputerInventoryV3) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetUserAndLocation

`func (o *ComputerInventoryV3) GetUserAndLocation() ComputerUserAndLocation`

GetUserAndLocation returns the UserAndLocation field if non-nil, zero value otherwise.

### GetUserAndLocationOk

`func (o *ComputerInventoryV3) GetUserAndLocationOk() (*ComputerUserAndLocation, bool)`

GetUserAndLocationOk returns a tuple with the UserAndLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAndLocation

`func (o *ComputerInventoryV3) SetUserAndLocation(v ComputerUserAndLocation)`

SetUserAndLocation sets UserAndLocation field to given value.

### HasUserAndLocation

`func (o *ComputerInventoryV3) HasUserAndLocation() bool`

HasUserAndLocation returns a boolean if a field has been set.

### GetConfigurationProfiles

`func (o *ComputerInventoryV3) GetConfigurationProfiles() []ComputerConfigurationProfile`

GetConfigurationProfiles returns the ConfigurationProfiles field if non-nil, zero value otherwise.

### GetConfigurationProfilesOk

`func (o *ComputerInventoryV3) GetConfigurationProfilesOk() (*[]ComputerConfigurationProfile, bool)`

GetConfigurationProfilesOk returns a tuple with the ConfigurationProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationProfiles

`func (o *ComputerInventoryV3) SetConfigurationProfiles(v []ComputerConfigurationProfile)`

SetConfigurationProfiles sets ConfigurationProfiles field to given value.

### HasConfigurationProfiles

`func (o *ComputerInventoryV3) HasConfigurationProfiles() bool`

HasConfigurationProfiles returns a boolean if a field has been set.

### GetPrinters

`func (o *ComputerInventoryV3) GetPrinters() []ComputerPrinter`

GetPrinters returns the Printers field if non-nil, zero value otherwise.

### GetPrintersOk

`func (o *ComputerInventoryV3) GetPrintersOk() (*[]ComputerPrinter, bool)`

GetPrintersOk returns a tuple with the Printers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinters

`func (o *ComputerInventoryV3) SetPrinters(v []ComputerPrinter)`

SetPrinters sets Printers field to given value.

### HasPrinters

`func (o *ComputerInventoryV3) HasPrinters() bool`

HasPrinters returns a boolean if a field has been set.

### GetServices

`func (o *ComputerInventoryV3) GetServices() []ComputerService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ComputerInventoryV3) GetServicesOk() (*[]ComputerService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ComputerInventoryV3) SetServices(v []ComputerService)`

SetServices sets Services field to given value.

### HasServices

`func (o *ComputerInventoryV3) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetHardware

`func (o *ComputerInventoryV3) GetHardware() ComputerHardware`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *ComputerInventoryV3) GetHardwareOk() (*ComputerHardware, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *ComputerInventoryV3) SetHardware(v ComputerHardware)`

SetHardware sets Hardware field to given value.

### HasHardware

`func (o *ComputerInventoryV3) HasHardware() bool`

HasHardware returns a boolean if a field has been set.

### GetLocalUserAccounts

`func (o *ComputerInventoryV3) GetLocalUserAccounts() []ComputerLocalUserAccount`

GetLocalUserAccounts returns the LocalUserAccounts field if non-nil, zero value otherwise.

### GetLocalUserAccountsOk

`func (o *ComputerInventoryV3) GetLocalUserAccountsOk() (*[]ComputerLocalUserAccount, bool)`

GetLocalUserAccountsOk returns a tuple with the LocalUserAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserAccounts

`func (o *ComputerInventoryV3) SetLocalUserAccounts(v []ComputerLocalUserAccount)`

SetLocalUserAccounts sets LocalUserAccounts field to given value.

### HasLocalUserAccounts

`func (o *ComputerInventoryV3) HasLocalUserAccounts() bool`

HasLocalUserAccounts returns a boolean if a field has been set.

### GetCertificates

`func (o *ComputerInventoryV3) GetCertificates() []ComputerCertificate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *ComputerInventoryV3) GetCertificatesOk() (*[]ComputerCertificate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *ComputerInventoryV3) SetCertificates(v []ComputerCertificate)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *ComputerInventoryV3) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetAttachments

`func (o *ComputerInventoryV3) GetAttachments() []ComputerAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ComputerInventoryV3) GetAttachmentsOk() (*[]ComputerAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ComputerInventoryV3) SetAttachments(v []ComputerAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ComputerInventoryV3) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetPackageReceipts

`func (o *ComputerInventoryV3) GetPackageReceipts() ComputerPackageReceipts`

GetPackageReceipts returns the PackageReceipts field if non-nil, zero value otherwise.

### GetPackageReceiptsOk

`func (o *ComputerInventoryV3) GetPackageReceiptsOk() (*ComputerPackageReceipts, bool)`

GetPackageReceiptsOk returns a tuple with the PackageReceipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageReceipts

`func (o *ComputerInventoryV3) SetPackageReceipts(v ComputerPackageReceipts)`

SetPackageReceipts sets PackageReceipts field to given value.

### HasPackageReceipts

`func (o *ComputerInventoryV3) HasPackageReceipts() bool`

HasPackageReceipts returns a boolean if a field has been set.

### GetSecurity

`func (o *ComputerInventoryV3) GetSecurity() ComputerSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *ComputerInventoryV3) GetSecurityOk() (*ComputerSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *ComputerInventoryV3) SetSecurity(v ComputerSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *ComputerInventoryV3) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *ComputerInventoryV3) GetOperatingSystem() ComputerOperatingSystem`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *ComputerInventoryV3) GetOperatingSystemOk() (*ComputerOperatingSystem, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *ComputerInventoryV3) SetOperatingSystem(v ComputerOperatingSystem)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *ComputerInventoryV3) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetLicensedSoftware

`func (o *ComputerInventoryV3) GetLicensedSoftware() []ComputerLicensedSoftware`

GetLicensedSoftware returns the LicensedSoftware field if non-nil, zero value otherwise.

### GetLicensedSoftwareOk

`func (o *ComputerInventoryV3) GetLicensedSoftwareOk() (*[]ComputerLicensedSoftware, bool)`

GetLicensedSoftwareOk returns a tuple with the LicensedSoftware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedSoftware

`func (o *ComputerInventoryV3) SetLicensedSoftware(v []ComputerLicensedSoftware)`

SetLicensedSoftware sets LicensedSoftware field to given value.

### HasLicensedSoftware

`func (o *ComputerInventoryV3) HasLicensedSoftware() bool`

HasLicensedSoftware returns a boolean if a field has been set.

### GetIbeacons

`func (o *ComputerInventoryV3) GetIbeacons() []ComputerIbeacon`

GetIbeacons returns the Ibeacons field if non-nil, zero value otherwise.

### GetIbeaconsOk

`func (o *ComputerInventoryV3) GetIbeaconsOk() (*[]ComputerIbeacon, bool)`

GetIbeaconsOk returns a tuple with the Ibeacons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbeacons

`func (o *ComputerInventoryV3) SetIbeacons(v []ComputerIbeacon)`

SetIbeacons sets Ibeacons field to given value.

### HasIbeacons

`func (o *ComputerInventoryV3) HasIbeacons() bool`

HasIbeacons returns a boolean if a field has been set.

### GetSoftwareUpdates

`func (o *ComputerInventoryV3) GetSoftwareUpdates() []ComputerSoftwareUpdate`

GetSoftwareUpdates returns the SoftwareUpdates field if non-nil, zero value otherwise.

### GetSoftwareUpdatesOk

`func (o *ComputerInventoryV3) GetSoftwareUpdatesOk() (*[]ComputerSoftwareUpdate, bool)`

GetSoftwareUpdatesOk returns a tuple with the SoftwareUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdates

`func (o *ComputerInventoryV3) SetSoftwareUpdates(v []ComputerSoftwareUpdate)`

SetSoftwareUpdates sets SoftwareUpdates field to given value.

### HasSoftwareUpdates

`func (o *ComputerInventoryV3) HasSoftwareUpdates() bool`

HasSoftwareUpdates returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *ComputerInventoryV3) GetExtensionAttributes() []ComputerExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *ComputerInventoryV3) GetExtensionAttributesOk() (*[]ComputerExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *ComputerInventoryV3) SetExtensionAttributes(v []ComputerExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *ComputerInventoryV3) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.

### GetContentCaching

`func (o *ComputerInventoryV3) GetContentCaching() ComputerContentCaching`

GetContentCaching returns the ContentCaching field if non-nil, zero value otherwise.

### GetContentCachingOk

`func (o *ComputerInventoryV3) GetContentCachingOk() (*ComputerContentCaching, bool)`

GetContentCachingOk returns a tuple with the ContentCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentCaching

`func (o *ComputerInventoryV3) SetContentCaching(v ComputerContentCaching)`

SetContentCaching sets ContentCaching field to given value.

### HasContentCaching

`func (o *ComputerInventoryV3) HasContentCaching() bool`

HasContentCaching returns a boolean if a field has been set.

### GetGroupMemberships

`func (o *ComputerInventoryV3) GetGroupMemberships() []GroupMembership`

GetGroupMemberships returns the GroupMemberships field if non-nil, zero value otherwise.

### GetGroupMembershipsOk

`func (o *ComputerInventoryV3) GetGroupMembershipsOk() (*[]GroupMembership, bool)`

GetGroupMembershipsOk returns a tuple with the GroupMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberships

`func (o *ComputerInventoryV3) SetGroupMemberships(v []GroupMembership)`

SetGroupMemberships sets GroupMemberships field to given value.

### HasGroupMemberships

`func (o *ComputerInventoryV3) HasGroupMemberships() bool`

HasGroupMemberships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


