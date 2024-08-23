# MobileDeviceWatchOsInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MobileDeviceId** | Pointer to **string** |  | [optional] [readonly] 
**DeviceType** | **string** | Based on the value of this type either ios, appleTv, watch or visionOS objects will be populated. | 
**Hardware** | Pointer to [**MobileDeviceHardware**](MobileDeviceHardware.md) |  | [optional] 
**UserAndLocation** | Pointer to [**MobileDeviceUserAndLocation**](MobileDeviceUserAndLocation.md) |  | [optional] 
**Applications** | Pointer to [**[]MobileDeviceApplicationInventoryDetail**](MobileDeviceApplicationInventoryDetail.md) |  | [optional] 
**Certificates** | Pointer to [**[]MobileDeviceCertificate**](MobileDeviceCertificate.md) |  | [optional] 
**Profiles** | Pointer to [**[]MobileDeviceProfile**](MobileDeviceProfile.md) |  | [optional] 
**ExtensionAttributes** | Pointer to [**[]MobileDeviceExtensionAttribute**](MobileDeviceExtensionAttribute.md) |  | [optional] 
**General** | Pointer to [**MobileDeviceWatchOsGeneral**](MobileDeviceWatchOsGeneral.md) |  | [optional] 
**Security** | Pointer to [**MobileDeviceSecurity**](MobileDeviceSecurity.md) |  | [optional] 
**ProvisioningProfiles** | Pointer to [**[]MobileDeviceProvisioningProfiles**](MobileDeviceProvisioningProfiles.md) |  | [optional] 

## Methods

### NewMobileDeviceWatchOsInventory

`func NewMobileDeviceWatchOsInventory(deviceType string, ) *MobileDeviceWatchOsInventory`

NewMobileDeviceWatchOsInventory instantiates a new MobileDeviceWatchOsInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceWatchOsInventoryWithDefaults

`func NewMobileDeviceWatchOsInventoryWithDefaults() *MobileDeviceWatchOsInventory`

NewMobileDeviceWatchOsInventoryWithDefaults instantiates a new MobileDeviceWatchOsInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMobileDeviceId

`func (o *MobileDeviceWatchOsInventory) GetMobileDeviceId() string`

GetMobileDeviceId returns the MobileDeviceId field if non-nil, zero value otherwise.

### GetMobileDeviceIdOk

`func (o *MobileDeviceWatchOsInventory) GetMobileDeviceIdOk() (*string, bool)`

GetMobileDeviceIdOk returns a tuple with the MobileDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileDeviceId

`func (o *MobileDeviceWatchOsInventory) SetMobileDeviceId(v string)`

SetMobileDeviceId sets MobileDeviceId field to given value.

### HasMobileDeviceId

`func (o *MobileDeviceWatchOsInventory) HasMobileDeviceId() bool`

HasMobileDeviceId returns a boolean if a field has been set.

### GetDeviceType

`func (o *MobileDeviceWatchOsInventory) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *MobileDeviceWatchOsInventory) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *MobileDeviceWatchOsInventory) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.


### GetHardware

`func (o *MobileDeviceWatchOsInventory) GetHardware() MobileDeviceHardware`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *MobileDeviceWatchOsInventory) GetHardwareOk() (*MobileDeviceHardware, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *MobileDeviceWatchOsInventory) SetHardware(v MobileDeviceHardware)`

SetHardware sets Hardware field to given value.

### HasHardware

`func (o *MobileDeviceWatchOsInventory) HasHardware() bool`

HasHardware returns a boolean if a field has been set.

### GetUserAndLocation

`func (o *MobileDeviceWatchOsInventory) GetUserAndLocation() MobileDeviceUserAndLocation`

GetUserAndLocation returns the UserAndLocation field if non-nil, zero value otherwise.

### GetUserAndLocationOk

`func (o *MobileDeviceWatchOsInventory) GetUserAndLocationOk() (*MobileDeviceUserAndLocation, bool)`

GetUserAndLocationOk returns a tuple with the UserAndLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAndLocation

`func (o *MobileDeviceWatchOsInventory) SetUserAndLocation(v MobileDeviceUserAndLocation)`

SetUserAndLocation sets UserAndLocation field to given value.

### HasUserAndLocation

`func (o *MobileDeviceWatchOsInventory) HasUserAndLocation() bool`

HasUserAndLocation returns a boolean if a field has been set.

### GetApplications

`func (o *MobileDeviceWatchOsInventory) GetApplications() []MobileDeviceApplicationInventoryDetail`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *MobileDeviceWatchOsInventory) GetApplicationsOk() (*[]MobileDeviceApplicationInventoryDetail, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *MobileDeviceWatchOsInventory) SetApplications(v []MobileDeviceApplicationInventoryDetail)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *MobileDeviceWatchOsInventory) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetCertificates

`func (o *MobileDeviceWatchOsInventory) GetCertificates() []MobileDeviceCertificate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *MobileDeviceWatchOsInventory) GetCertificatesOk() (*[]MobileDeviceCertificate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *MobileDeviceWatchOsInventory) SetCertificates(v []MobileDeviceCertificate)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *MobileDeviceWatchOsInventory) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetProfiles

`func (o *MobileDeviceWatchOsInventory) GetProfiles() []MobileDeviceProfile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *MobileDeviceWatchOsInventory) GetProfilesOk() (*[]MobileDeviceProfile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *MobileDeviceWatchOsInventory) SetProfiles(v []MobileDeviceProfile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *MobileDeviceWatchOsInventory) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *MobileDeviceWatchOsInventory) GetExtensionAttributes() []MobileDeviceExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *MobileDeviceWatchOsInventory) GetExtensionAttributesOk() (*[]MobileDeviceExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *MobileDeviceWatchOsInventory) SetExtensionAttributes(v []MobileDeviceExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *MobileDeviceWatchOsInventory) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.

### GetGeneral

`func (o *MobileDeviceWatchOsInventory) GetGeneral() MobileDeviceWatchOsGeneral`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *MobileDeviceWatchOsInventory) GetGeneralOk() (*MobileDeviceWatchOsGeneral, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *MobileDeviceWatchOsInventory) SetGeneral(v MobileDeviceWatchOsGeneral)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *MobileDeviceWatchOsInventory) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetSecurity

`func (o *MobileDeviceWatchOsInventory) GetSecurity() MobileDeviceSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *MobileDeviceWatchOsInventory) GetSecurityOk() (*MobileDeviceSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *MobileDeviceWatchOsInventory) SetSecurity(v MobileDeviceSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *MobileDeviceWatchOsInventory) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetProvisioningProfiles

`func (o *MobileDeviceWatchOsInventory) GetProvisioningProfiles() []MobileDeviceProvisioningProfiles`

GetProvisioningProfiles returns the ProvisioningProfiles field if non-nil, zero value otherwise.

### GetProvisioningProfilesOk

`func (o *MobileDeviceWatchOsInventory) GetProvisioningProfilesOk() (*[]MobileDeviceProvisioningProfiles, bool)`

GetProvisioningProfilesOk returns a tuple with the ProvisioningProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningProfiles

`func (o *MobileDeviceWatchOsInventory) SetProvisioningProfiles(v []MobileDeviceProvisioningProfiles)`

SetProvisioningProfiles sets ProvisioningProfiles field to given value.

### HasProvisioningProfiles

`func (o *MobileDeviceWatchOsInventory) HasProvisioningProfiles() bool`

HasProvisioningProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


