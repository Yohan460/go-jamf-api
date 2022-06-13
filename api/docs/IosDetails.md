# IosDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** |  | [optional] 
**ModelIdentifier** | Pointer to **string** |  | [optional] 
**ModelNumber** | Pointer to **string** |  | [optional] 
**IsSupervised** | Pointer to **bool** |  | [optional] 
**BatteryLevel** | Pointer to **int32** |  | [optional] 
**LastBackupTimestamp** | Pointer to **time.Time** |  | [optional] 
**CapacityMb** | Pointer to **int32** |  | [optional] 
**AvailableMb** | Pointer to **int32** |  | [optional] 
**PercentageUsed** | Pointer to **int32** |  | [optional] 
**IsShared** | Pointer to **bool** |  | [optional] 
**IsDeviceLocatorServiceEnabled** | Pointer to **bool** |  | [optional] 
**IsDoNotDisturbEnabled** | Pointer to **bool** |  | [optional] 
**IsCloudBackupEnabled** | Pointer to **bool** |  | [optional] 
**LastCloudBackupTimestamp** | Pointer to **time.Time** |  | [optional] 
**IsLocationServicesEnabled** | Pointer to **bool** |  | [optional] 
**IsITunesStoreAccountActive** | Pointer to **bool** |  | [optional] 
**IsBleCapable** | Pointer to **bool** |  | [optional] 
**Computer** | Pointer to [**IdAndName**](IdAndName.md) |  | [optional] 
**Purchasing** | Pointer to [**Purchasing**](Purchasing.md) |  | [optional] 
**Security** | Pointer to [**Security**](Security.md) |  | [optional] 
**Network** | Pointer to [**Network**](Network.md) |  | [optional] 
**Applications** | Pointer to [**[]MobileDeviceApplication**](MobileDeviceApplication.md) |  | [optional] 
**Certificates** | Pointer to [**[]MobileDeviceCertificateV1**](MobileDeviceCertificateV1.md) |  | [optional] 
**Ebooks** | Pointer to [**[]MobileDeviceEbook**](MobileDeviceEbook.md) |  | [optional] 
**ConfigurationProfiles** | Pointer to [**[]ConfigurationProfile**](ConfigurationProfile.md) |  | [optional] 
**ProvisioningProfiles** | Pointer to [**[]ProvisioningProfile**](ProvisioningProfile.md) |  | [optional] 
**Attachments** | Pointer to [**[]MobileDeviceAttachment**](MobileDeviceAttachment.md) |  | [optional] 

## Methods

### NewIosDetails

`func NewIosDetails() *IosDetails`

NewIosDetails instantiates a new IosDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIosDetailsWithDefaults

`func NewIosDetailsWithDefaults() *IosDetails`

NewIosDetailsWithDefaults instantiates a new IosDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *IosDetails) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *IosDetails) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *IosDetails) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *IosDetails) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelIdentifier

`func (o *IosDetails) GetModelIdentifier() string`

GetModelIdentifier returns the ModelIdentifier field if non-nil, zero value otherwise.

### GetModelIdentifierOk

`func (o *IosDetails) GetModelIdentifierOk() (*string, bool)`

GetModelIdentifierOk returns a tuple with the ModelIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelIdentifier

`func (o *IosDetails) SetModelIdentifier(v string)`

SetModelIdentifier sets ModelIdentifier field to given value.

### HasModelIdentifier

`func (o *IosDetails) HasModelIdentifier() bool`

HasModelIdentifier returns a boolean if a field has been set.

### GetModelNumber

`func (o *IosDetails) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *IosDetails) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *IosDetails) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *IosDetails) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetIsSupervised

`func (o *IosDetails) GetIsSupervised() bool`

GetIsSupervised returns the IsSupervised field if non-nil, zero value otherwise.

### GetIsSupervisedOk

`func (o *IosDetails) GetIsSupervisedOk() (*bool, bool)`

GetIsSupervisedOk returns a tuple with the IsSupervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupervised

`func (o *IosDetails) SetIsSupervised(v bool)`

SetIsSupervised sets IsSupervised field to given value.

### HasIsSupervised

`func (o *IosDetails) HasIsSupervised() bool`

HasIsSupervised returns a boolean if a field has been set.

### GetBatteryLevel

`func (o *IosDetails) GetBatteryLevel() int32`

GetBatteryLevel returns the BatteryLevel field if non-nil, zero value otherwise.

### GetBatteryLevelOk

`func (o *IosDetails) GetBatteryLevelOk() (*int32, bool)`

GetBatteryLevelOk returns a tuple with the BatteryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryLevel

`func (o *IosDetails) SetBatteryLevel(v int32)`

SetBatteryLevel sets BatteryLevel field to given value.

### HasBatteryLevel

`func (o *IosDetails) HasBatteryLevel() bool`

HasBatteryLevel returns a boolean if a field has been set.

### GetLastBackupTimestamp

`func (o *IosDetails) GetLastBackupTimestamp() time.Time`

GetLastBackupTimestamp returns the LastBackupTimestamp field if non-nil, zero value otherwise.

### GetLastBackupTimestampOk

`func (o *IosDetails) GetLastBackupTimestampOk() (*time.Time, bool)`

GetLastBackupTimestampOk returns a tuple with the LastBackupTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackupTimestamp

`func (o *IosDetails) SetLastBackupTimestamp(v time.Time)`

SetLastBackupTimestamp sets LastBackupTimestamp field to given value.

### HasLastBackupTimestamp

`func (o *IosDetails) HasLastBackupTimestamp() bool`

HasLastBackupTimestamp returns a boolean if a field has been set.

### GetCapacityMb

`func (o *IosDetails) GetCapacityMb() int32`

GetCapacityMb returns the CapacityMb field if non-nil, zero value otherwise.

### GetCapacityMbOk

`func (o *IosDetails) GetCapacityMbOk() (*int32, bool)`

GetCapacityMbOk returns a tuple with the CapacityMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityMb

`func (o *IosDetails) SetCapacityMb(v int32)`

SetCapacityMb sets CapacityMb field to given value.

### HasCapacityMb

`func (o *IosDetails) HasCapacityMb() bool`

HasCapacityMb returns a boolean if a field has been set.

### GetAvailableMb

`func (o *IosDetails) GetAvailableMb() int32`

GetAvailableMb returns the AvailableMb field if non-nil, zero value otherwise.

### GetAvailableMbOk

`func (o *IosDetails) GetAvailableMbOk() (*int32, bool)`

GetAvailableMbOk returns a tuple with the AvailableMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMb

`func (o *IosDetails) SetAvailableMb(v int32)`

SetAvailableMb sets AvailableMb field to given value.

### HasAvailableMb

`func (o *IosDetails) HasAvailableMb() bool`

HasAvailableMb returns a boolean if a field has been set.

### GetPercentageUsed

`func (o *IosDetails) GetPercentageUsed() int32`

GetPercentageUsed returns the PercentageUsed field if non-nil, zero value otherwise.

### GetPercentageUsedOk

`func (o *IosDetails) GetPercentageUsedOk() (*int32, bool)`

GetPercentageUsedOk returns a tuple with the PercentageUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageUsed

`func (o *IosDetails) SetPercentageUsed(v int32)`

SetPercentageUsed sets PercentageUsed field to given value.

### HasPercentageUsed

`func (o *IosDetails) HasPercentageUsed() bool`

HasPercentageUsed returns a boolean if a field has been set.

### GetIsShared

`func (o *IosDetails) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *IosDetails) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *IosDetails) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *IosDetails) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetIsDeviceLocatorServiceEnabled

`func (o *IosDetails) GetIsDeviceLocatorServiceEnabled() bool`

GetIsDeviceLocatorServiceEnabled returns the IsDeviceLocatorServiceEnabled field if non-nil, zero value otherwise.

### GetIsDeviceLocatorServiceEnabledOk

`func (o *IosDetails) GetIsDeviceLocatorServiceEnabledOk() (*bool, bool)`

GetIsDeviceLocatorServiceEnabledOk returns a tuple with the IsDeviceLocatorServiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeviceLocatorServiceEnabled

`func (o *IosDetails) SetIsDeviceLocatorServiceEnabled(v bool)`

SetIsDeviceLocatorServiceEnabled sets IsDeviceLocatorServiceEnabled field to given value.

### HasIsDeviceLocatorServiceEnabled

`func (o *IosDetails) HasIsDeviceLocatorServiceEnabled() bool`

HasIsDeviceLocatorServiceEnabled returns a boolean if a field has been set.

### GetIsDoNotDisturbEnabled

`func (o *IosDetails) GetIsDoNotDisturbEnabled() bool`

GetIsDoNotDisturbEnabled returns the IsDoNotDisturbEnabled field if non-nil, zero value otherwise.

### GetIsDoNotDisturbEnabledOk

`func (o *IosDetails) GetIsDoNotDisturbEnabledOk() (*bool, bool)`

GetIsDoNotDisturbEnabledOk returns a tuple with the IsDoNotDisturbEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDoNotDisturbEnabled

`func (o *IosDetails) SetIsDoNotDisturbEnabled(v bool)`

SetIsDoNotDisturbEnabled sets IsDoNotDisturbEnabled field to given value.

### HasIsDoNotDisturbEnabled

`func (o *IosDetails) HasIsDoNotDisturbEnabled() bool`

HasIsDoNotDisturbEnabled returns a boolean if a field has been set.

### GetIsCloudBackupEnabled

`func (o *IosDetails) GetIsCloudBackupEnabled() bool`

GetIsCloudBackupEnabled returns the IsCloudBackupEnabled field if non-nil, zero value otherwise.

### GetIsCloudBackupEnabledOk

`func (o *IosDetails) GetIsCloudBackupEnabledOk() (*bool, bool)`

GetIsCloudBackupEnabledOk returns a tuple with the IsCloudBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudBackupEnabled

`func (o *IosDetails) SetIsCloudBackupEnabled(v bool)`

SetIsCloudBackupEnabled sets IsCloudBackupEnabled field to given value.

### HasIsCloudBackupEnabled

`func (o *IosDetails) HasIsCloudBackupEnabled() bool`

HasIsCloudBackupEnabled returns a boolean if a field has been set.

### GetLastCloudBackupTimestamp

`func (o *IosDetails) GetLastCloudBackupTimestamp() time.Time`

GetLastCloudBackupTimestamp returns the LastCloudBackupTimestamp field if non-nil, zero value otherwise.

### GetLastCloudBackupTimestampOk

`func (o *IosDetails) GetLastCloudBackupTimestampOk() (*time.Time, bool)`

GetLastCloudBackupTimestampOk returns a tuple with the LastCloudBackupTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCloudBackupTimestamp

`func (o *IosDetails) SetLastCloudBackupTimestamp(v time.Time)`

SetLastCloudBackupTimestamp sets LastCloudBackupTimestamp field to given value.

### HasLastCloudBackupTimestamp

`func (o *IosDetails) HasLastCloudBackupTimestamp() bool`

HasLastCloudBackupTimestamp returns a boolean if a field has been set.

### GetIsLocationServicesEnabled

`func (o *IosDetails) GetIsLocationServicesEnabled() bool`

GetIsLocationServicesEnabled returns the IsLocationServicesEnabled field if non-nil, zero value otherwise.

### GetIsLocationServicesEnabledOk

`func (o *IosDetails) GetIsLocationServicesEnabledOk() (*bool, bool)`

GetIsLocationServicesEnabledOk returns a tuple with the IsLocationServicesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocationServicesEnabled

`func (o *IosDetails) SetIsLocationServicesEnabled(v bool)`

SetIsLocationServicesEnabled sets IsLocationServicesEnabled field to given value.

### HasIsLocationServicesEnabled

`func (o *IosDetails) HasIsLocationServicesEnabled() bool`

HasIsLocationServicesEnabled returns a boolean if a field has been set.

### GetIsITunesStoreAccountActive

`func (o *IosDetails) GetIsITunesStoreAccountActive() bool`

GetIsITunesStoreAccountActive returns the IsITunesStoreAccountActive field if non-nil, zero value otherwise.

### GetIsITunesStoreAccountActiveOk

`func (o *IosDetails) GetIsITunesStoreAccountActiveOk() (*bool, bool)`

GetIsITunesStoreAccountActiveOk returns a tuple with the IsITunesStoreAccountActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsITunesStoreAccountActive

`func (o *IosDetails) SetIsITunesStoreAccountActive(v bool)`

SetIsITunesStoreAccountActive sets IsITunesStoreAccountActive field to given value.

### HasIsITunesStoreAccountActive

`func (o *IosDetails) HasIsITunesStoreAccountActive() bool`

HasIsITunesStoreAccountActive returns a boolean if a field has been set.

### GetIsBleCapable

`func (o *IosDetails) GetIsBleCapable() bool`

GetIsBleCapable returns the IsBleCapable field if non-nil, zero value otherwise.

### GetIsBleCapableOk

`func (o *IosDetails) GetIsBleCapableOk() (*bool, bool)`

GetIsBleCapableOk returns a tuple with the IsBleCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBleCapable

`func (o *IosDetails) SetIsBleCapable(v bool)`

SetIsBleCapable sets IsBleCapable field to given value.

### HasIsBleCapable

`func (o *IosDetails) HasIsBleCapable() bool`

HasIsBleCapable returns a boolean if a field has been set.

### GetComputer

`func (o *IosDetails) GetComputer() IdAndName`

GetComputer returns the Computer field if non-nil, zero value otherwise.

### GetComputerOk

`func (o *IosDetails) GetComputerOk() (*IdAndName, bool)`

GetComputerOk returns a tuple with the Computer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputer

`func (o *IosDetails) SetComputer(v IdAndName)`

SetComputer sets Computer field to given value.

### HasComputer

`func (o *IosDetails) HasComputer() bool`

HasComputer returns a boolean if a field has been set.

### GetPurchasing

`func (o *IosDetails) GetPurchasing() Purchasing`

GetPurchasing returns the Purchasing field if non-nil, zero value otherwise.

### GetPurchasingOk

`func (o *IosDetails) GetPurchasingOk() (*Purchasing, bool)`

GetPurchasingOk returns a tuple with the Purchasing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasing

`func (o *IosDetails) SetPurchasing(v Purchasing)`

SetPurchasing sets Purchasing field to given value.

### HasPurchasing

`func (o *IosDetails) HasPurchasing() bool`

HasPurchasing returns a boolean if a field has been set.

### GetSecurity

`func (o *IosDetails) GetSecurity() Security`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *IosDetails) GetSecurityOk() (*Security, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *IosDetails) SetSecurity(v Security)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *IosDetails) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetNetwork

`func (o *IosDetails) GetNetwork() Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *IosDetails) GetNetworkOk() (*Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *IosDetails) SetNetwork(v Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *IosDetails) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetApplications

`func (o *IosDetails) GetApplications() []MobileDeviceApplication`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *IosDetails) GetApplicationsOk() (*[]MobileDeviceApplication, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *IosDetails) SetApplications(v []MobileDeviceApplication)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *IosDetails) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetCertificates

`func (o *IosDetails) GetCertificates() []MobileDeviceCertificateV1`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *IosDetails) GetCertificatesOk() (*[]MobileDeviceCertificateV1, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *IosDetails) SetCertificates(v []MobileDeviceCertificateV1)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *IosDetails) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetEbooks

`func (o *IosDetails) GetEbooks() []MobileDeviceEbook`

GetEbooks returns the Ebooks field if non-nil, zero value otherwise.

### GetEbooksOk

`func (o *IosDetails) GetEbooksOk() (*[]MobileDeviceEbook, bool)`

GetEbooksOk returns a tuple with the Ebooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbooks

`func (o *IosDetails) SetEbooks(v []MobileDeviceEbook)`

SetEbooks sets Ebooks field to given value.

### HasEbooks

`func (o *IosDetails) HasEbooks() bool`

HasEbooks returns a boolean if a field has been set.

### GetConfigurationProfiles

`func (o *IosDetails) GetConfigurationProfiles() []ConfigurationProfile`

GetConfigurationProfiles returns the ConfigurationProfiles field if non-nil, zero value otherwise.

### GetConfigurationProfilesOk

`func (o *IosDetails) GetConfigurationProfilesOk() (*[]ConfigurationProfile, bool)`

GetConfigurationProfilesOk returns a tuple with the ConfigurationProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationProfiles

`func (o *IosDetails) SetConfigurationProfiles(v []ConfigurationProfile)`

SetConfigurationProfiles sets ConfigurationProfiles field to given value.

### HasConfigurationProfiles

`func (o *IosDetails) HasConfigurationProfiles() bool`

HasConfigurationProfiles returns a boolean if a field has been set.

### GetProvisioningProfiles

`func (o *IosDetails) GetProvisioningProfiles() []ProvisioningProfile`

GetProvisioningProfiles returns the ProvisioningProfiles field if non-nil, zero value otherwise.

### GetProvisioningProfilesOk

`func (o *IosDetails) GetProvisioningProfilesOk() (*[]ProvisioningProfile, bool)`

GetProvisioningProfilesOk returns a tuple with the ProvisioningProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningProfiles

`func (o *IosDetails) SetProvisioningProfiles(v []ProvisioningProfile)`

SetProvisioningProfiles sets ProvisioningProfiles field to given value.

### HasProvisioningProfiles

`func (o *IosDetails) HasProvisioningProfiles() bool`

HasProvisioningProfiles returns a boolean if a field has been set.

### GetAttachments

`func (o *IosDetails) GetAttachments() []MobileDeviceAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *IosDetails) GetAttachmentsOk() (*[]MobileDeviceAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *IosDetails) SetAttachments(v []MobileDeviceAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *IosDetails) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


