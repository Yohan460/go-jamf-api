# WatchOsDetailsV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** |  | [optional] 
**ModelIdentifier** | Pointer to **string** |  | [optional] 
**ModelNumber** | Pointer to **string** |  | [optional] 
**Supervised** | Pointer to **bool** |  | [optional] 
**BatteryLevel** | Pointer to **int64** |  | [optional] 
**CapacityMb** | Pointer to **int64** |  | [optional] 
**AvailableMb** | Pointer to **int64** |  | [optional] 
**PercentageUsed** | Pointer to **int64** |  | [optional] 
**DeviceLocatorServiceEnabled** | Pointer to **bool** |  | [optional] 
**DoNotDisturbEnabled** | Pointer to **bool** |  | [optional] 
**LastCloudBackupTimestamp** | Pointer to **time.Time** |  | [optional] 
**ITunesStoreAccountActive** | Pointer to **bool** |  | [optional] 
**BleCapable** | Pointer to **bool** |  | [optional] 
**Security** | Pointer to [**SecurityV2**](SecurityV2.md) |  | [optional] 
**Applications** | Pointer to [**[]MobileDeviceApplication**](MobileDeviceApplication.md) |  | [optional] 
**Certificates** | Pointer to [**[]MobileDeviceCertificateV2**](MobileDeviceCertificateV2.md) |  | [optional] 
**ConfigurationProfiles** | Pointer to [**[]ConfigurationProfile**](ConfigurationProfile.md) |  | [optional] 
**ProvisioningProfiles** | Pointer to [**[]MobileDeviceProvisioningProfiles**](MobileDeviceProvisioningProfiles.md) |  | [optional] 
**Attachments** | Pointer to [**[]MobileDeviceAttachmentV2**](MobileDeviceAttachmentV2.md) |  | [optional] 

## Methods

### NewWatchOsDetailsV2

`func NewWatchOsDetailsV2() *WatchOsDetailsV2`

NewWatchOsDetailsV2 instantiates a new WatchOsDetailsV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatchOsDetailsV2WithDefaults

`func NewWatchOsDetailsV2WithDefaults() *WatchOsDetailsV2`

NewWatchOsDetailsV2WithDefaults instantiates a new WatchOsDetailsV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *WatchOsDetailsV2) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *WatchOsDetailsV2) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *WatchOsDetailsV2) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *WatchOsDetailsV2) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelIdentifier

`func (o *WatchOsDetailsV2) GetModelIdentifier() string`

GetModelIdentifier returns the ModelIdentifier field if non-nil, zero value otherwise.

### GetModelIdentifierOk

`func (o *WatchOsDetailsV2) GetModelIdentifierOk() (*string, bool)`

GetModelIdentifierOk returns a tuple with the ModelIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelIdentifier

`func (o *WatchOsDetailsV2) SetModelIdentifier(v string)`

SetModelIdentifier sets ModelIdentifier field to given value.

### HasModelIdentifier

`func (o *WatchOsDetailsV2) HasModelIdentifier() bool`

HasModelIdentifier returns a boolean if a field has been set.

### GetModelNumber

`func (o *WatchOsDetailsV2) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *WatchOsDetailsV2) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *WatchOsDetailsV2) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *WatchOsDetailsV2) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetSupervised

`func (o *WatchOsDetailsV2) GetSupervised() bool`

GetSupervised returns the Supervised field if non-nil, zero value otherwise.

### GetSupervisedOk

`func (o *WatchOsDetailsV2) GetSupervisedOk() (*bool, bool)`

GetSupervisedOk returns a tuple with the Supervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervised

`func (o *WatchOsDetailsV2) SetSupervised(v bool)`

SetSupervised sets Supervised field to given value.

### HasSupervised

`func (o *WatchOsDetailsV2) HasSupervised() bool`

HasSupervised returns a boolean if a field has been set.

### GetBatteryLevel

`func (o *WatchOsDetailsV2) GetBatteryLevel() int64`

GetBatteryLevel returns the BatteryLevel field if non-nil, zero value otherwise.

### GetBatteryLevelOk

`func (o *WatchOsDetailsV2) GetBatteryLevelOk() (*int64, bool)`

GetBatteryLevelOk returns a tuple with the BatteryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryLevel

`func (o *WatchOsDetailsV2) SetBatteryLevel(v int64)`

SetBatteryLevel sets BatteryLevel field to given value.

### HasBatteryLevel

`func (o *WatchOsDetailsV2) HasBatteryLevel() bool`

HasBatteryLevel returns a boolean if a field has been set.

### GetCapacityMb

`func (o *WatchOsDetailsV2) GetCapacityMb() int64`

GetCapacityMb returns the CapacityMb field if non-nil, zero value otherwise.

### GetCapacityMbOk

`func (o *WatchOsDetailsV2) GetCapacityMbOk() (*int64, bool)`

GetCapacityMbOk returns a tuple with the CapacityMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityMb

`func (o *WatchOsDetailsV2) SetCapacityMb(v int64)`

SetCapacityMb sets CapacityMb field to given value.

### HasCapacityMb

`func (o *WatchOsDetailsV2) HasCapacityMb() bool`

HasCapacityMb returns a boolean if a field has been set.

### GetAvailableMb

`func (o *WatchOsDetailsV2) GetAvailableMb() int64`

GetAvailableMb returns the AvailableMb field if non-nil, zero value otherwise.

### GetAvailableMbOk

`func (o *WatchOsDetailsV2) GetAvailableMbOk() (*int64, bool)`

GetAvailableMbOk returns a tuple with the AvailableMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMb

`func (o *WatchOsDetailsV2) SetAvailableMb(v int64)`

SetAvailableMb sets AvailableMb field to given value.

### HasAvailableMb

`func (o *WatchOsDetailsV2) HasAvailableMb() bool`

HasAvailableMb returns a boolean if a field has been set.

### GetPercentageUsed

`func (o *WatchOsDetailsV2) GetPercentageUsed() int64`

GetPercentageUsed returns the PercentageUsed field if non-nil, zero value otherwise.

### GetPercentageUsedOk

`func (o *WatchOsDetailsV2) GetPercentageUsedOk() (*int64, bool)`

GetPercentageUsedOk returns a tuple with the PercentageUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageUsed

`func (o *WatchOsDetailsV2) SetPercentageUsed(v int64)`

SetPercentageUsed sets PercentageUsed field to given value.

### HasPercentageUsed

`func (o *WatchOsDetailsV2) HasPercentageUsed() bool`

HasPercentageUsed returns a boolean if a field has been set.

### GetDeviceLocatorServiceEnabled

`func (o *WatchOsDetailsV2) GetDeviceLocatorServiceEnabled() bool`

GetDeviceLocatorServiceEnabled returns the DeviceLocatorServiceEnabled field if non-nil, zero value otherwise.

### GetDeviceLocatorServiceEnabledOk

`func (o *WatchOsDetailsV2) GetDeviceLocatorServiceEnabledOk() (*bool, bool)`

GetDeviceLocatorServiceEnabledOk returns a tuple with the DeviceLocatorServiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLocatorServiceEnabled

`func (o *WatchOsDetailsV2) SetDeviceLocatorServiceEnabled(v bool)`

SetDeviceLocatorServiceEnabled sets DeviceLocatorServiceEnabled field to given value.

### HasDeviceLocatorServiceEnabled

`func (o *WatchOsDetailsV2) HasDeviceLocatorServiceEnabled() bool`

HasDeviceLocatorServiceEnabled returns a boolean if a field has been set.

### GetDoNotDisturbEnabled

`func (o *WatchOsDetailsV2) GetDoNotDisturbEnabled() bool`

GetDoNotDisturbEnabled returns the DoNotDisturbEnabled field if non-nil, zero value otherwise.

### GetDoNotDisturbEnabledOk

`func (o *WatchOsDetailsV2) GetDoNotDisturbEnabledOk() (*bool, bool)`

GetDoNotDisturbEnabledOk returns a tuple with the DoNotDisturbEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDisturbEnabled

`func (o *WatchOsDetailsV2) SetDoNotDisturbEnabled(v bool)`

SetDoNotDisturbEnabled sets DoNotDisturbEnabled field to given value.

### HasDoNotDisturbEnabled

`func (o *WatchOsDetailsV2) HasDoNotDisturbEnabled() bool`

HasDoNotDisturbEnabled returns a boolean if a field has been set.

### GetLastCloudBackupTimestamp

`func (o *WatchOsDetailsV2) GetLastCloudBackupTimestamp() time.Time`

GetLastCloudBackupTimestamp returns the LastCloudBackupTimestamp field if non-nil, zero value otherwise.

### GetLastCloudBackupTimestampOk

`func (o *WatchOsDetailsV2) GetLastCloudBackupTimestampOk() (*time.Time, bool)`

GetLastCloudBackupTimestampOk returns a tuple with the LastCloudBackupTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCloudBackupTimestamp

`func (o *WatchOsDetailsV2) SetLastCloudBackupTimestamp(v time.Time)`

SetLastCloudBackupTimestamp sets LastCloudBackupTimestamp field to given value.

### HasLastCloudBackupTimestamp

`func (o *WatchOsDetailsV2) HasLastCloudBackupTimestamp() bool`

HasLastCloudBackupTimestamp returns a boolean if a field has been set.

### GetITunesStoreAccountActive

`func (o *WatchOsDetailsV2) GetITunesStoreAccountActive() bool`

GetITunesStoreAccountActive returns the ITunesStoreAccountActive field if non-nil, zero value otherwise.

### GetITunesStoreAccountActiveOk

`func (o *WatchOsDetailsV2) GetITunesStoreAccountActiveOk() (*bool, bool)`

GetITunesStoreAccountActiveOk returns a tuple with the ITunesStoreAccountActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetITunesStoreAccountActive

`func (o *WatchOsDetailsV2) SetITunesStoreAccountActive(v bool)`

SetITunesStoreAccountActive sets ITunesStoreAccountActive field to given value.

### HasITunesStoreAccountActive

`func (o *WatchOsDetailsV2) HasITunesStoreAccountActive() bool`

HasITunesStoreAccountActive returns a boolean if a field has been set.

### GetBleCapable

`func (o *WatchOsDetailsV2) GetBleCapable() bool`

GetBleCapable returns the BleCapable field if non-nil, zero value otherwise.

### GetBleCapableOk

`func (o *WatchOsDetailsV2) GetBleCapableOk() (*bool, bool)`

GetBleCapableOk returns a tuple with the BleCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBleCapable

`func (o *WatchOsDetailsV2) SetBleCapable(v bool)`

SetBleCapable sets BleCapable field to given value.

### HasBleCapable

`func (o *WatchOsDetailsV2) HasBleCapable() bool`

HasBleCapable returns a boolean if a field has been set.

### GetSecurity

`func (o *WatchOsDetailsV2) GetSecurity() SecurityV2`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *WatchOsDetailsV2) GetSecurityOk() (*SecurityV2, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *WatchOsDetailsV2) SetSecurity(v SecurityV2)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *WatchOsDetailsV2) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetApplications

`func (o *WatchOsDetailsV2) GetApplications() []MobileDeviceApplication`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *WatchOsDetailsV2) GetApplicationsOk() (*[]MobileDeviceApplication, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *WatchOsDetailsV2) SetApplications(v []MobileDeviceApplication)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *WatchOsDetailsV2) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetCertificates

`func (o *WatchOsDetailsV2) GetCertificates() []MobileDeviceCertificateV2`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *WatchOsDetailsV2) GetCertificatesOk() (*[]MobileDeviceCertificateV2, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *WatchOsDetailsV2) SetCertificates(v []MobileDeviceCertificateV2)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *WatchOsDetailsV2) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetConfigurationProfiles

`func (o *WatchOsDetailsV2) GetConfigurationProfiles() []ConfigurationProfile`

GetConfigurationProfiles returns the ConfigurationProfiles field if non-nil, zero value otherwise.

### GetConfigurationProfilesOk

`func (o *WatchOsDetailsV2) GetConfigurationProfilesOk() (*[]ConfigurationProfile, bool)`

GetConfigurationProfilesOk returns a tuple with the ConfigurationProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationProfiles

`func (o *WatchOsDetailsV2) SetConfigurationProfiles(v []ConfigurationProfile)`

SetConfigurationProfiles sets ConfigurationProfiles field to given value.

### HasConfigurationProfiles

`func (o *WatchOsDetailsV2) HasConfigurationProfiles() bool`

HasConfigurationProfiles returns a boolean if a field has been set.

### GetProvisioningProfiles

`func (o *WatchOsDetailsV2) GetProvisioningProfiles() []MobileDeviceProvisioningProfiles`

GetProvisioningProfiles returns the ProvisioningProfiles field if non-nil, zero value otherwise.

### GetProvisioningProfilesOk

`func (o *WatchOsDetailsV2) GetProvisioningProfilesOk() (*[]MobileDeviceProvisioningProfiles, bool)`

GetProvisioningProfilesOk returns a tuple with the ProvisioningProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningProfiles

`func (o *WatchOsDetailsV2) SetProvisioningProfiles(v []MobileDeviceProvisioningProfiles)`

SetProvisioningProfiles sets ProvisioningProfiles field to given value.

### HasProvisioningProfiles

`func (o *WatchOsDetailsV2) HasProvisioningProfiles() bool`

HasProvisioningProfiles returns a boolean if a field has been set.

### GetAttachments

`func (o *WatchOsDetailsV2) GetAttachments() []MobileDeviceAttachmentV2`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *WatchOsDetailsV2) GetAttachmentsOk() (*[]MobileDeviceAttachmentV2, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *WatchOsDetailsV2) SetAttachments(v []MobileDeviceAttachmentV2)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *WatchOsDetailsV2) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


