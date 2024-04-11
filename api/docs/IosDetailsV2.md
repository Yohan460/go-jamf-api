# IosDetailsV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** |  | [optional] 
**ModelIdentifier** | Pointer to **string** |  | [optional] 
**ModelNumber** | Pointer to **string** |  | [optional] 
**Supervised** | Pointer to **bool** |  | [optional] 
**BatteryLevel** | Pointer to **int64** |  | [optional] 
**LastBackupTimestamp** | Pointer to **time.Time** |  | [optional] 
**CapacityMb** | Pointer to **int64** |  | [optional] 
**AvailableMb** | Pointer to **int64** |  | [optional] 
**PercentageUsed** | Pointer to **int64** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**DeviceLocatorServiceEnabled** | Pointer to **bool** |  | [optional] 
**DoNotDisturbEnabled** | Pointer to **bool** |  | [optional] 
**CloudBackupEnabled** | Pointer to **bool** |  | [optional] 
**LastCloudBackupTimestamp** | Pointer to **time.Time** |  | [optional] 
**LocationServicesEnabled** | Pointer to **bool** |  | [optional] 
**ITunesStoreAccountActive** | Pointer to **bool** |  | [optional] 
**BleCapable** | Pointer to **bool** |  | [optional] 
**Computer** | Pointer to [**IdAndNameV2**](IdAndNameV2.md) |  | [optional] 
**Purchasing** | Pointer to [**PurchasingV2**](PurchasingV2.md) |  | [optional] 
**Security** | Pointer to [**SecurityV2**](SecurityV2.md) |  | [optional] 
**Network** | Pointer to [**NetworkV2**](NetworkV2.md) |  | [optional] 
**ServiceSubscriptions** | Pointer to [**[]MobileDeviceServiceSubscriptions**](MobileDeviceServiceSubscriptions.md) |  | [optional] 
**Applications** | Pointer to [**[]MobileDeviceApplication**](MobileDeviceApplication.md) |  | [optional] 
**Certificates** | Pointer to [**[]MobileDeviceCertificateV2**](MobileDeviceCertificateV2.md) |  | [optional] 
**Ebooks** | Pointer to [**[]MobileDeviceEbook**](MobileDeviceEbook.md) |  | [optional] 
**ConfigurationProfiles** | Pointer to [**[]ConfigurationProfile**](ConfigurationProfile.md) |  | [optional] 
**ProvisioningProfiles** | Pointer to [**[]MobileDeviceProvisioningProfiles**](MobileDeviceProvisioningProfiles.md) |  | [optional] 
**Attachments** | Pointer to [**[]MobileDeviceAttachmentV2**](MobileDeviceAttachmentV2.md) |  | [optional] 

## Methods

### NewIosDetailsV2

`func NewIosDetailsV2() *IosDetailsV2`

NewIosDetailsV2 instantiates a new IosDetailsV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIosDetailsV2WithDefaults

`func NewIosDetailsV2WithDefaults() *IosDetailsV2`

NewIosDetailsV2WithDefaults instantiates a new IosDetailsV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *IosDetailsV2) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *IosDetailsV2) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *IosDetailsV2) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *IosDetailsV2) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelIdentifier

`func (o *IosDetailsV2) GetModelIdentifier() string`

GetModelIdentifier returns the ModelIdentifier field if non-nil, zero value otherwise.

### GetModelIdentifierOk

`func (o *IosDetailsV2) GetModelIdentifierOk() (*string, bool)`

GetModelIdentifierOk returns a tuple with the ModelIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelIdentifier

`func (o *IosDetailsV2) SetModelIdentifier(v string)`

SetModelIdentifier sets ModelIdentifier field to given value.

### HasModelIdentifier

`func (o *IosDetailsV2) HasModelIdentifier() bool`

HasModelIdentifier returns a boolean if a field has been set.

### GetModelNumber

`func (o *IosDetailsV2) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *IosDetailsV2) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *IosDetailsV2) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *IosDetailsV2) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetSupervised

`func (o *IosDetailsV2) GetSupervised() bool`

GetSupervised returns the Supervised field if non-nil, zero value otherwise.

### GetSupervisedOk

`func (o *IosDetailsV2) GetSupervisedOk() (*bool, bool)`

GetSupervisedOk returns a tuple with the Supervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervised

`func (o *IosDetailsV2) SetSupervised(v bool)`

SetSupervised sets Supervised field to given value.

### HasSupervised

`func (o *IosDetailsV2) HasSupervised() bool`

HasSupervised returns a boolean if a field has been set.

### GetBatteryLevel

`func (o *IosDetailsV2) GetBatteryLevel() int64`

GetBatteryLevel returns the BatteryLevel field if non-nil, zero value otherwise.

### GetBatteryLevelOk

`func (o *IosDetailsV2) GetBatteryLevelOk() (*int64, bool)`

GetBatteryLevelOk returns a tuple with the BatteryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryLevel

`func (o *IosDetailsV2) SetBatteryLevel(v int64)`

SetBatteryLevel sets BatteryLevel field to given value.

### HasBatteryLevel

`func (o *IosDetailsV2) HasBatteryLevel() bool`

HasBatteryLevel returns a boolean if a field has been set.

### GetLastBackupTimestamp

`func (o *IosDetailsV2) GetLastBackupTimestamp() time.Time`

GetLastBackupTimestamp returns the LastBackupTimestamp field if non-nil, zero value otherwise.

### GetLastBackupTimestampOk

`func (o *IosDetailsV2) GetLastBackupTimestampOk() (*time.Time, bool)`

GetLastBackupTimestampOk returns a tuple with the LastBackupTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackupTimestamp

`func (o *IosDetailsV2) SetLastBackupTimestamp(v time.Time)`

SetLastBackupTimestamp sets LastBackupTimestamp field to given value.

### HasLastBackupTimestamp

`func (o *IosDetailsV2) HasLastBackupTimestamp() bool`

HasLastBackupTimestamp returns a boolean if a field has been set.

### GetCapacityMb

`func (o *IosDetailsV2) GetCapacityMb() int64`

GetCapacityMb returns the CapacityMb field if non-nil, zero value otherwise.

### GetCapacityMbOk

`func (o *IosDetailsV2) GetCapacityMbOk() (*int64, bool)`

GetCapacityMbOk returns a tuple with the CapacityMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityMb

`func (o *IosDetailsV2) SetCapacityMb(v int64)`

SetCapacityMb sets CapacityMb field to given value.

### HasCapacityMb

`func (o *IosDetailsV2) HasCapacityMb() bool`

HasCapacityMb returns a boolean if a field has been set.

### GetAvailableMb

`func (o *IosDetailsV2) GetAvailableMb() int64`

GetAvailableMb returns the AvailableMb field if non-nil, zero value otherwise.

### GetAvailableMbOk

`func (o *IosDetailsV2) GetAvailableMbOk() (*int64, bool)`

GetAvailableMbOk returns a tuple with the AvailableMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMb

`func (o *IosDetailsV2) SetAvailableMb(v int64)`

SetAvailableMb sets AvailableMb field to given value.

### HasAvailableMb

`func (o *IosDetailsV2) HasAvailableMb() bool`

HasAvailableMb returns a boolean if a field has been set.

### GetPercentageUsed

`func (o *IosDetailsV2) GetPercentageUsed() int64`

GetPercentageUsed returns the PercentageUsed field if non-nil, zero value otherwise.

### GetPercentageUsedOk

`func (o *IosDetailsV2) GetPercentageUsedOk() (*int64, bool)`

GetPercentageUsedOk returns a tuple with the PercentageUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageUsed

`func (o *IosDetailsV2) SetPercentageUsed(v int64)`

SetPercentageUsed sets PercentageUsed field to given value.

### HasPercentageUsed

`func (o *IosDetailsV2) HasPercentageUsed() bool`

HasPercentageUsed returns a boolean if a field has been set.

### GetShared

`func (o *IosDetailsV2) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *IosDetailsV2) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *IosDetailsV2) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *IosDetailsV2) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetDeviceLocatorServiceEnabled

`func (o *IosDetailsV2) GetDeviceLocatorServiceEnabled() bool`

GetDeviceLocatorServiceEnabled returns the DeviceLocatorServiceEnabled field if non-nil, zero value otherwise.

### GetDeviceLocatorServiceEnabledOk

`func (o *IosDetailsV2) GetDeviceLocatorServiceEnabledOk() (*bool, bool)`

GetDeviceLocatorServiceEnabledOk returns a tuple with the DeviceLocatorServiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLocatorServiceEnabled

`func (o *IosDetailsV2) SetDeviceLocatorServiceEnabled(v bool)`

SetDeviceLocatorServiceEnabled sets DeviceLocatorServiceEnabled field to given value.

### HasDeviceLocatorServiceEnabled

`func (o *IosDetailsV2) HasDeviceLocatorServiceEnabled() bool`

HasDeviceLocatorServiceEnabled returns a boolean if a field has been set.

### GetDoNotDisturbEnabled

`func (o *IosDetailsV2) GetDoNotDisturbEnabled() bool`

GetDoNotDisturbEnabled returns the DoNotDisturbEnabled field if non-nil, zero value otherwise.

### GetDoNotDisturbEnabledOk

`func (o *IosDetailsV2) GetDoNotDisturbEnabledOk() (*bool, bool)`

GetDoNotDisturbEnabledOk returns a tuple with the DoNotDisturbEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDisturbEnabled

`func (o *IosDetailsV2) SetDoNotDisturbEnabled(v bool)`

SetDoNotDisturbEnabled sets DoNotDisturbEnabled field to given value.

### HasDoNotDisturbEnabled

`func (o *IosDetailsV2) HasDoNotDisturbEnabled() bool`

HasDoNotDisturbEnabled returns a boolean if a field has been set.

### GetCloudBackupEnabled

`func (o *IosDetailsV2) GetCloudBackupEnabled() bool`

GetCloudBackupEnabled returns the CloudBackupEnabled field if non-nil, zero value otherwise.

### GetCloudBackupEnabledOk

`func (o *IosDetailsV2) GetCloudBackupEnabledOk() (*bool, bool)`

GetCloudBackupEnabledOk returns a tuple with the CloudBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudBackupEnabled

`func (o *IosDetailsV2) SetCloudBackupEnabled(v bool)`

SetCloudBackupEnabled sets CloudBackupEnabled field to given value.

### HasCloudBackupEnabled

`func (o *IosDetailsV2) HasCloudBackupEnabled() bool`

HasCloudBackupEnabled returns a boolean if a field has been set.

### GetLastCloudBackupTimestamp

`func (o *IosDetailsV2) GetLastCloudBackupTimestamp() time.Time`

GetLastCloudBackupTimestamp returns the LastCloudBackupTimestamp field if non-nil, zero value otherwise.

### GetLastCloudBackupTimestampOk

`func (o *IosDetailsV2) GetLastCloudBackupTimestampOk() (*time.Time, bool)`

GetLastCloudBackupTimestampOk returns a tuple with the LastCloudBackupTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCloudBackupTimestamp

`func (o *IosDetailsV2) SetLastCloudBackupTimestamp(v time.Time)`

SetLastCloudBackupTimestamp sets LastCloudBackupTimestamp field to given value.

### HasLastCloudBackupTimestamp

`func (o *IosDetailsV2) HasLastCloudBackupTimestamp() bool`

HasLastCloudBackupTimestamp returns a boolean if a field has been set.

### GetLocationServicesEnabled

`func (o *IosDetailsV2) GetLocationServicesEnabled() bool`

GetLocationServicesEnabled returns the LocationServicesEnabled field if non-nil, zero value otherwise.

### GetLocationServicesEnabledOk

`func (o *IosDetailsV2) GetLocationServicesEnabledOk() (*bool, bool)`

GetLocationServicesEnabledOk returns a tuple with the LocationServicesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationServicesEnabled

`func (o *IosDetailsV2) SetLocationServicesEnabled(v bool)`

SetLocationServicesEnabled sets LocationServicesEnabled field to given value.

### HasLocationServicesEnabled

`func (o *IosDetailsV2) HasLocationServicesEnabled() bool`

HasLocationServicesEnabled returns a boolean if a field has been set.

### GetITunesStoreAccountActive

`func (o *IosDetailsV2) GetITunesStoreAccountActive() bool`

GetITunesStoreAccountActive returns the ITunesStoreAccountActive field if non-nil, zero value otherwise.

### GetITunesStoreAccountActiveOk

`func (o *IosDetailsV2) GetITunesStoreAccountActiveOk() (*bool, bool)`

GetITunesStoreAccountActiveOk returns a tuple with the ITunesStoreAccountActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetITunesStoreAccountActive

`func (o *IosDetailsV2) SetITunesStoreAccountActive(v bool)`

SetITunesStoreAccountActive sets ITunesStoreAccountActive field to given value.

### HasITunesStoreAccountActive

`func (o *IosDetailsV2) HasITunesStoreAccountActive() bool`

HasITunesStoreAccountActive returns a boolean if a field has been set.

### GetBleCapable

`func (o *IosDetailsV2) GetBleCapable() bool`

GetBleCapable returns the BleCapable field if non-nil, zero value otherwise.

### GetBleCapableOk

`func (o *IosDetailsV2) GetBleCapableOk() (*bool, bool)`

GetBleCapableOk returns a tuple with the BleCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBleCapable

`func (o *IosDetailsV2) SetBleCapable(v bool)`

SetBleCapable sets BleCapable field to given value.

### HasBleCapable

`func (o *IosDetailsV2) HasBleCapable() bool`

HasBleCapable returns a boolean if a field has been set.

### GetComputer

`func (o *IosDetailsV2) GetComputer() IdAndNameV2`

GetComputer returns the Computer field if non-nil, zero value otherwise.

### GetComputerOk

`func (o *IosDetailsV2) GetComputerOk() (*IdAndNameV2, bool)`

GetComputerOk returns a tuple with the Computer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputer

`func (o *IosDetailsV2) SetComputer(v IdAndNameV2)`

SetComputer sets Computer field to given value.

### HasComputer

`func (o *IosDetailsV2) HasComputer() bool`

HasComputer returns a boolean if a field has been set.

### GetPurchasing

`func (o *IosDetailsV2) GetPurchasing() PurchasingV2`

GetPurchasing returns the Purchasing field if non-nil, zero value otherwise.

### GetPurchasingOk

`func (o *IosDetailsV2) GetPurchasingOk() (*PurchasingV2, bool)`

GetPurchasingOk returns a tuple with the Purchasing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasing

`func (o *IosDetailsV2) SetPurchasing(v PurchasingV2)`

SetPurchasing sets Purchasing field to given value.

### HasPurchasing

`func (o *IosDetailsV2) HasPurchasing() bool`

HasPurchasing returns a boolean if a field has been set.

### GetSecurity

`func (o *IosDetailsV2) GetSecurity() SecurityV2`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *IosDetailsV2) GetSecurityOk() (*SecurityV2, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *IosDetailsV2) SetSecurity(v SecurityV2)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *IosDetailsV2) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetNetwork

`func (o *IosDetailsV2) GetNetwork() NetworkV2`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *IosDetailsV2) GetNetworkOk() (*NetworkV2, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *IosDetailsV2) SetNetwork(v NetworkV2)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *IosDetailsV2) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetServiceSubscriptions

`func (o *IosDetailsV2) GetServiceSubscriptions() []MobileDeviceServiceSubscriptions`

GetServiceSubscriptions returns the ServiceSubscriptions field if non-nil, zero value otherwise.

### GetServiceSubscriptionsOk

`func (o *IosDetailsV2) GetServiceSubscriptionsOk() (*[]MobileDeviceServiceSubscriptions, bool)`

GetServiceSubscriptionsOk returns a tuple with the ServiceSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSubscriptions

`func (o *IosDetailsV2) SetServiceSubscriptions(v []MobileDeviceServiceSubscriptions)`

SetServiceSubscriptions sets ServiceSubscriptions field to given value.

### HasServiceSubscriptions

`func (o *IosDetailsV2) HasServiceSubscriptions() bool`

HasServiceSubscriptions returns a boolean if a field has been set.

### GetApplications

`func (o *IosDetailsV2) GetApplications() []MobileDeviceApplication`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *IosDetailsV2) GetApplicationsOk() (*[]MobileDeviceApplication, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *IosDetailsV2) SetApplications(v []MobileDeviceApplication)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *IosDetailsV2) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetCertificates

`func (o *IosDetailsV2) GetCertificates() []MobileDeviceCertificateV2`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *IosDetailsV2) GetCertificatesOk() (*[]MobileDeviceCertificateV2, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *IosDetailsV2) SetCertificates(v []MobileDeviceCertificateV2)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *IosDetailsV2) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetEbooks

`func (o *IosDetailsV2) GetEbooks() []MobileDeviceEbook`

GetEbooks returns the Ebooks field if non-nil, zero value otherwise.

### GetEbooksOk

`func (o *IosDetailsV2) GetEbooksOk() (*[]MobileDeviceEbook, bool)`

GetEbooksOk returns a tuple with the Ebooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbooks

`func (o *IosDetailsV2) SetEbooks(v []MobileDeviceEbook)`

SetEbooks sets Ebooks field to given value.

### HasEbooks

`func (o *IosDetailsV2) HasEbooks() bool`

HasEbooks returns a boolean if a field has been set.

### GetConfigurationProfiles

`func (o *IosDetailsV2) GetConfigurationProfiles() []ConfigurationProfile`

GetConfigurationProfiles returns the ConfigurationProfiles field if non-nil, zero value otherwise.

### GetConfigurationProfilesOk

`func (o *IosDetailsV2) GetConfigurationProfilesOk() (*[]ConfigurationProfile, bool)`

GetConfigurationProfilesOk returns a tuple with the ConfigurationProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationProfiles

`func (o *IosDetailsV2) SetConfigurationProfiles(v []ConfigurationProfile)`

SetConfigurationProfiles sets ConfigurationProfiles field to given value.

### HasConfigurationProfiles

`func (o *IosDetailsV2) HasConfigurationProfiles() bool`

HasConfigurationProfiles returns a boolean if a field has been set.

### GetProvisioningProfiles

`func (o *IosDetailsV2) GetProvisioningProfiles() []MobileDeviceProvisioningProfiles`

GetProvisioningProfiles returns the ProvisioningProfiles field if non-nil, zero value otherwise.

### GetProvisioningProfilesOk

`func (o *IosDetailsV2) GetProvisioningProfilesOk() (*[]MobileDeviceProvisioningProfiles, bool)`

GetProvisioningProfilesOk returns a tuple with the ProvisioningProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningProfiles

`func (o *IosDetailsV2) SetProvisioningProfiles(v []MobileDeviceProvisioningProfiles)`

SetProvisioningProfiles sets ProvisioningProfiles field to given value.

### HasProvisioningProfiles

`func (o *IosDetailsV2) HasProvisioningProfiles() bool`

HasProvisioningProfiles returns a boolean if a field has been set.

### GetAttachments

`func (o *IosDetailsV2) GetAttachments() []MobileDeviceAttachmentV2`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *IosDetailsV2) GetAttachmentsOk() (*[]MobileDeviceAttachmentV2, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *IosDetailsV2) SetAttachments(v []MobileDeviceAttachmentV2)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *IosDetailsV2) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


