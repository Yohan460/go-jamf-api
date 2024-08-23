# DetailsV2

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

### NewDetailsV2

`func NewDetailsV2() *DetailsV2`

NewDetailsV2 instantiates a new DetailsV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailsV2WithDefaults

`func NewDetailsV2WithDefaults() *DetailsV2`

NewDetailsV2WithDefaults instantiates a new DetailsV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *DetailsV2) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DetailsV2) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DetailsV2) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DetailsV2) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelIdentifier

`func (o *DetailsV2) GetModelIdentifier() string`

GetModelIdentifier returns the ModelIdentifier field if non-nil, zero value otherwise.

### GetModelIdentifierOk

`func (o *DetailsV2) GetModelIdentifierOk() (*string, bool)`

GetModelIdentifierOk returns a tuple with the ModelIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelIdentifier

`func (o *DetailsV2) SetModelIdentifier(v string)`

SetModelIdentifier sets ModelIdentifier field to given value.

### HasModelIdentifier

`func (o *DetailsV2) HasModelIdentifier() bool`

HasModelIdentifier returns a boolean if a field has been set.

### GetModelNumber

`func (o *DetailsV2) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *DetailsV2) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *DetailsV2) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *DetailsV2) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetSupervised

`func (o *DetailsV2) GetSupervised() bool`

GetSupervised returns the Supervised field if non-nil, zero value otherwise.

### GetSupervisedOk

`func (o *DetailsV2) GetSupervisedOk() (*bool, bool)`

GetSupervisedOk returns a tuple with the Supervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervised

`func (o *DetailsV2) SetSupervised(v bool)`

SetSupervised sets Supervised field to given value.

### HasSupervised

`func (o *DetailsV2) HasSupervised() bool`

HasSupervised returns a boolean if a field has been set.

### GetBatteryLevel

`func (o *DetailsV2) GetBatteryLevel() int64`

GetBatteryLevel returns the BatteryLevel field if non-nil, zero value otherwise.

### GetBatteryLevelOk

`func (o *DetailsV2) GetBatteryLevelOk() (*int64, bool)`

GetBatteryLevelOk returns a tuple with the BatteryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryLevel

`func (o *DetailsV2) SetBatteryLevel(v int64)`

SetBatteryLevel sets BatteryLevel field to given value.

### HasBatteryLevel

`func (o *DetailsV2) HasBatteryLevel() bool`

HasBatteryLevel returns a boolean if a field has been set.

### GetLastBackupTimestamp

`func (o *DetailsV2) GetLastBackupTimestamp() time.Time`

GetLastBackupTimestamp returns the LastBackupTimestamp field if non-nil, zero value otherwise.

### GetLastBackupTimestampOk

`func (o *DetailsV2) GetLastBackupTimestampOk() (*time.Time, bool)`

GetLastBackupTimestampOk returns a tuple with the LastBackupTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackupTimestamp

`func (o *DetailsV2) SetLastBackupTimestamp(v time.Time)`

SetLastBackupTimestamp sets LastBackupTimestamp field to given value.

### HasLastBackupTimestamp

`func (o *DetailsV2) HasLastBackupTimestamp() bool`

HasLastBackupTimestamp returns a boolean if a field has been set.

### GetCapacityMb

`func (o *DetailsV2) GetCapacityMb() int64`

GetCapacityMb returns the CapacityMb field if non-nil, zero value otherwise.

### GetCapacityMbOk

`func (o *DetailsV2) GetCapacityMbOk() (*int64, bool)`

GetCapacityMbOk returns a tuple with the CapacityMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityMb

`func (o *DetailsV2) SetCapacityMb(v int64)`

SetCapacityMb sets CapacityMb field to given value.

### HasCapacityMb

`func (o *DetailsV2) HasCapacityMb() bool`

HasCapacityMb returns a boolean if a field has been set.

### GetAvailableMb

`func (o *DetailsV2) GetAvailableMb() int64`

GetAvailableMb returns the AvailableMb field if non-nil, zero value otherwise.

### GetAvailableMbOk

`func (o *DetailsV2) GetAvailableMbOk() (*int64, bool)`

GetAvailableMbOk returns a tuple with the AvailableMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMb

`func (o *DetailsV2) SetAvailableMb(v int64)`

SetAvailableMb sets AvailableMb field to given value.

### HasAvailableMb

`func (o *DetailsV2) HasAvailableMb() bool`

HasAvailableMb returns a boolean if a field has been set.

### GetPercentageUsed

`func (o *DetailsV2) GetPercentageUsed() int64`

GetPercentageUsed returns the PercentageUsed field if non-nil, zero value otherwise.

### GetPercentageUsedOk

`func (o *DetailsV2) GetPercentageUsedOk() (*int64, bool)`

GetPercentageUsedOk returns a tuple with the PercentageUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageUsed

`func (o *DetailsV2) SetPercentageUsed(v int64)`

SetPercentageUsed sets PercentageUsed field to given value.

### HasPercentageUsed

`func (o *DetailsV2) HasPercentageUsed() bool`

HasPercentageUsed returns a boolean if a field has been set.

### GetShared

`func (o *DetailsV2) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *DetailsV2) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *DetailsV2) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *DetailsV2) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetDeviceLocatorServiceEnabled

`func (o *DetailsV2) GetDeviceLocatorServiceEnabled() bool`

GetDeviceLocatorServiceEnabled returns the DeviceLocatorServiceEnabled field if non-nil, zero value otherwise.

### GetDeviceLocatorServiceEnabledOk

`func (o *DetailsV2) GetDeviceLocatorServiceEnabledOk() (*bool, bool)`

GetDeviceLocatorServiceEnabledOk returns a tuple with the DeviceLocatorServiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLocatorServiceEnabled

`func (o *DetailsV2) SetDeviceLocatorServiceEnabled(v bool)`

SetDeviceLocatorServiceEnabled sets DeviceLocatorServiceEnabled field to given value.

### HasDeviceLocatorServiceEnabled

`func (o *DetailsV2) HasDeviceLocatorServiceEnabled() bool`

HasDeviceLocatorServiceEnabled returns a boolean if a field has been set.

### GetDoNotDisturbEnabled

`func (o *DetailsV2) GetDoNotDisturbEnabled() bool`

GetDoNotDisturbEnabled returns the DoNotDisturbEnabled field if non-nil, zero value otherwise.

### GetDoNotDisturbEnabledOk

`func (o *DetailsV2) GetDoNotDisturbEnabledOk() (*bool, bool)`

GetDoNotDisturbEnabledOk returns a tuple with the DoNotDisturbEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDisturbEnabled

`func (o *DetailsV2) SetDoNotDisturbEnabled(v bool)`

SetDoNotDisturbEnabled sets DoNotDisturbEnabled field to given value.

### HasDoNotDisturbEnabled

`func (o *DetailsV2) HasDoNotDisturbEnabled() bool`

HasDoNotDisturbEnabled returns a boolean if a field has been set.

### GetCloudBackupEnabled

`func (o *DetailsV2) GetCloudBackupEnabled() bool`

GetCloudBackupEnabled returns the CloudBackupEnabled field if non-nil, zero value otherwise.

### GetCloudBackupEnabledOk

`func (o *DetailsV2) GetCloudBackupEnabledOk() (*bool, bool)`

GetCloudBackupEnabledOk returns a tuple with the CloudBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudBackupEnabled

`func (o *DetailsV2) SetCloudBackupEnabled(v bool)`

SetCloudBackupEnabled sets CloudBackupEnabled field to given value.

### HasCloudBackupEnabled

`func (o *DetailsV2) HasCloudBackupEnabled() bool`

HasCloudBackupEnabled returns a boolean if a field has been set.

### GetLastCloudBackupTimestamp

`func (o *DetailsV2) GetLastCloudBackupTimestamp() time.Time`

GetLastCloudBackupTimestamp returns the LastCloudBackupTimestamp field if non-nil, zero value otherwise.

### GetLastCloudBackupTimestampOk

`func (o *DetailsV2) GetLastCloudBackupTimestampOk() (*time.Time, bool)`

GetLastCloudBackupTimestampOk returns a tuple with the LastCloudBackupTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCloudBackupTimestamp

`func (o *DetailsV2) SetLastCloudBackupTimestamp(v time.Time)`

SetLastCloudBackupTimestamp sets LastCloudBackupTimestamp field to given value.

### HasLastCloudBackupTimestamp

`func (o *DetailsV2) HasLastCloudBackupTimestamp() bool`

HasLastCloudBackupTimestamp returns a boolean if a field has been set.

### GetLocationServicesEnabled

`func (o *DetailsV2) GetLocationServicesEnabled() bool`

GetLocationServicesEnabled returns the LocationServicesEnabled field if non-nil, zero value otherwise.

### GetLocationServicesEnabledOk

`func (o *DetailsV2) GetLocationServicesEnabledOk() (*bool, bool)`

GetLocationServicesEnabledOk returns a tuple with the LocationServicesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationServicesEnabled

`func (o *DetailsV2) SetLocationServicesEnabled(v bool)`

SetLocationServicesEnabled sets LocationServicesEnabled field to given value.

### HasLocationServicesEnabled

`func (o *DetailsV2) HasLocationServicesEnabled() bool`

HasLocationServicesEnabled returns a boolean if a field has been set.

### GetITunesStoreAccountActive

`func (o *DetailsV2) GetITunesStoreAccountActive() bool`

GetITunesStoreAccountActive returns the ITunesStoreAccountActive field if non-nil, zero value otherwise.

### GetITunesStoreAccountActiveOk

`func (o *DetailsV2) GetITunesStoreAccountActiveOk() (*bool, bool)`

GetITunesStoreAccountActiveOk returns a tuple with the ITunesStoreAccountActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetITunesStoreAccountActive

`func (o *DetailsV2) SetITunesStoreAccountActive(v bool)`

SetITunesStoreAccountActive sets ITunesStoreAccountActive field to given value.

### HasITunesStoreAccountActive

`func (o *DetailsV2) HasITunesStoreAccountActive() bool`

HasITunesStoreAccountActive returns a boolean if a field has been set.

### GetBleCapable

`func (o *DetailsV2) GetBleCapable() bool`

GetBleCapable returns the BleCapable field if non-nil, zero value otherwise.

### GetBleCapableOk

`func (o *DetailsV2) GetBleCapableOk() (*bool, bool)`

GetBleCapableOk returns a tuple with the BleCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBleCapable

`func (o *DetailsV2) SetBleCapable(v bool)`

SetBleCapable sets BleCapable field to given value.

### HasBleCapable

`func (o *DetailsV2) HasBleCapable() bool`

HasBleCapable returns a boolean if a field has been set.

### GetComputer

`func (o *DetailsV2) GetComputer() IdAndNameV2`

GetComputer returns the Computer field if non-nil, zero value otherwise.

### GetComputerOk

`func (o *DetailsV2) GetComputerOk() (*IdAndNameV2, bool)`

GetComputerOk returns a tuple with the Computer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputer

`func (o *DetailsV2) SetComputer(v IdAndNameV2)`

SetComputer sets Computer field to given value.

### HasComputer

`func (o *DetailsV2) HasComputer() bool`

HasComputer returns a boolean if a field has been set.

### GetPurchasing

`func (o *DetailsV2) GetPurchasing() PurchasingV2`

GetPurchasing returns the Purchasing field if non-nil, zero value otherwise.

### GetPurchasingOk

`func (o *DetailsV2) GetPurchasingOk() (*PurchasingV2, bool)`

GetPurchasingOk returns a tuple with the Purchasing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasing

`func (o *DetailsV2) SetPurchasing(v PurchasingV2)`

SetPurchasing sets Purchasing field to given value.

### HasPurchasing

`func (o *DetailsV2) HasPurchasing() bool`

HasPurchasing returns a boolean if a field has been set.

### GetSecurity

`func (o *DetailsV2) GetSecurity() SecurityV2`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *DetailsV2) GetSecurityOk() (*SecurityV2, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *DetailsV2) SetSecurity(v SecurityV2)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *DetailsV2) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetNetwork

`func (o *DetailsV2) GetNetwork() NetworkV2`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *DetailsV2) GetNetworkOk() (*NetworkV2, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *DetailsV2) SetNetwork(v NetworkV2)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *DetailsV2) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetServiceSubscriptions

`func (o *DetailsV2) GetServiceSubscriptions() []MobileDeviceServiceSubscriptions`

GetServiceSubscriptions returns the ServiceSubscriptions field if non-nil, zero value otherwise.

### GetServiceSubscriptionsOk

`func (o *DetailsV2) GetServiceSubscriptionsOk() (*[]MobileDeviceServiceSubscriptions, bool)`

GetServiceSubscriptionsOk returns a tuple with the ServiceSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSubscriptions

`func (o *DetailsV2) SetServiceSubscriptions(v []MobileDeviceServiceSubscriptions)`

SetServiceSubscriptions sets ServiceSubscriptions field to given value.

### HasServiceSubscriptions

`func (o *DetailsV2) HasServiceSubscriptions() bool`

HasServiceSubscriptions returns a boolean if a field has been set.

### GetApplications

`func (o *DetailsV2) GetApplications() []MobileDeviceApplication`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *DetailsV2) GetApplicationsOk() (*[]MobileDeviceApplication, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *DetailsV2) SetApplications(v []MobileDeviceApplication)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *DetailsV2) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetCertificates

`func (o *DetailsV2) GetCertificates() []MobileDeviceCertificateV2`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *DetailsV2) GetCertificatesOk() (*[]MobileDeviceCertificateV2, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *DetailsV2) SetCertificates(v []MobileDeviceCertificateV2)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *DetailsV2) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetEbooks

`func (o *DetailsV2) GetEbooks() []MobileDeviceEbook`

GetEbooks returns the Ebooks field if non-nil, zero value otherwise.

### GetEbooksOk

`func (o *DetailsV2) GetEbooksOk() (*[]MobileDeviceEbook, bool)`

GetEbooksOk returns a tuple with the Ebooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbooks

`func (o *DetailsV2) SetEbooks(v []MobileDeviceEbook)`

SetEbooks sets Ebooks field to given value.

### HasEbooks

`func (o *DetailsV2) HasEbooks() bool`

HasEbooks returns a boolean if a field has been set.

### GetConfigurationProfiles

`func (o *DetailsV2) GetConfigurationProfiles() []ConfigurationProfile`

GetConfigurationProfiles returns the ConfigurationProfiles field if non-nil, zero value otherwise.

### GetConfigurationProfilesOk

`func (o *DetailsV2) GetConfigurationProfilesOk() (*[]ConfigurationProfile, bool)`

GetConfigurationProfilesOk returns a tuple with the ConfigurationProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationProfiles

`func (o *DetailsV2) SetConfigurationProfiles(v []ConfigurationProfile)`

SetConfigurationProfiles sets ConfigurationProfiles field to given value.

### HasConfigurationProfiles

`func (o *DetailsV2) HasConfigurationProfiles() bool`

HasConfigurationProfiles returns a boolean if a field has been set.

### GetProvisioningProfiles

`func (o *DetailsV2) GetProvisioningProfiles() []MobileDeviceProvisioningProfiles`

GetProvisioningProfiles returns the ProvisioningProfiles field if non-nil, zero value otherwise.

### GetProvisioningProfilesOk

`func (o *DetailsV2) GetProvisioningProfilesOk() (*[]MobileDeviceProvisioningProfiles, bool)`

GetProvisioningProfilesOk returns a tuple with the ProvisioningProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningProfiles

`func (o *DetailsV2) SetProvisioningProfiles(v []MobileDeviceProvisioningProfiles)`

SetProvisioningProfiles sets ProvisioningProfiles field to given value.

### HasProvisioningProfiles

`func (o *DetailsV2) HasProvisioningProfiles() bool`

HasProvisioningProfiles returns a boolean if a field has been set.

### GetAttachments

`func (o *DetailsV2) GetAttachments() []MobileDeviceAttachmentV2`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *DetailsV2) GetAttachmentsOk() (*[]MobileDeviceAttachmentV2, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *DetailsV2) SetAttachments(v []MobileDeviceAttachmentV2)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *DetailsV2) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


