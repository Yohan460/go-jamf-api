# MobileDeviceTvOsInventory

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
**General** | Pointer to [**MobileDeviceTvOsGeneral**](MobileDeviceTvOsGeneral.md) |  | [optional] 
**Purchasing** | Pointer to [**MobileDevicePurchasing**](MobileDevicePurchasing.md) |  | [optional] 
**UserProfiles** | Pointer to [**[]MobileDeviceUserProfile**](MobileDeviceUserProfile.md) |  | [optional] 

## Methods

### NewMobileDeviceTvOsInventory

`func NewMobileDeviceTvOsInventory(deviceType string, ) *MobileDeviceTvOsInventory`

NewMobileDeviceTvOsInventory instantiates a new MobileDeviceTvOsInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceTvOsInventoryWithDefaults

`func NewMobileDeviceTvOsInventoryWithDefaults() *MobileDeviceTvOsInventory`

NewMobileDeviceTvOsInventoryWithDefaults instantiates a new MobileDeviceTvOsInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMobileDeviceId

`func (o *MobileDeviceTvOsInventory) GetMobileDeviceId() string`

GetMobileDeviceId returns the MobileDeviceId field if non-nil, zero value otherwise.

### GetMobileDeviceIdOk

`func (o *MobileDeviceTvOsInventory) GetMobileDeviceIdOk() (*string, bool)`

GetMobileDeviceIdOk returns a tuple with the MobileDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileDeviceId

`func (o *MobileDeviceTvOsInventory) SetMobileDeviceId(v string)`

SetMobileDeviceId sets MobileDeviceId field to given value.

### HasMobileDeviceId

`func (o *MobileDeviceTvOsInventory) HasMobileDeviceId() bool`

HasMobileDeviceId returns a boolean if a field has been set.

### GetDeviceType

`func (o *MobileDeviceTvOsInventory) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *MobileDeviceTvOsInventory) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *MobileDeviceTvOsInventory) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.


### GetHardware

`func (o *MobileDeviceTvOsInventory) GetHardware() MobileDeviceHardware`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *MobileDeviceTvOsInventory) GetHardwareOk() (*MobileDeviceHardware, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *MobileDeviceTvOsInventory) SetHardware(v MobileDeviceHardware)`

SetHardware sets Hardware field to given value.

### HasHardware

`func (o *MobileDeviceTvOsInventory) HasHardware() bool`

HasHardware returns a boolean if a field has been set.

### GetUserAndLocation

`func (o *MobileDeviceTvOsInventory) GetUserAndLocation() MobileDeviceUserAndLocation`

GetUserAndLocation returns the UserAndLocation field if non-nil, zero value otherwise.

### GetUserAndLocationOk

`func (o *MobileDeviceTvOsInventory) GetUserAndLocationOk() (*MobileDeviceUserAndLocation, bool)`

GetUserAndLocationOk returns a tuple with the UserAndLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAndLocation

`func (o *MobileDeviceTvOsInventory) SetUserAndLocation(v MobileDeviceUserAndLocation)`

SetUserAndLocation sets UserAndLocation field to given value.

### HasUserAndLocation

`func (o *MobileDeviceTvOsInventory) HasUserAndLocation() bool`

HasUserAndLocation returns a boolean if a field has been set.

### GetApplications

`func (o *MobileDeviceTvOsInventory) GetApplications() []MobileDeviceApplicationInventoryDetail`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *MobileDeviceTvOsInventory) GetApplicationsOk() (*[]MobileDeviceApplicationInventoryDetail, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *MobileDeviceTvOsInventory) SetApplications(v []MobileDeviceApplicationInventoryDetail)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *MobileDeviceTvOsInventory) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetCertificates

`func (o *MobileDeviceTvOsInventory) GetCertificates() []MobileDeviceCertificate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *MobileDeviceTvOsInventory) GetCertificatesOk() (*[]MobileDeviceCertificate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *MobileDeviceTvOsInventory) SetCertificates(v []MobileDeviceCertificate)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *MobileDeviceTvOsInventory) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetProfiles

`func (o *MobileDeviceTvOsInventory) GetProfiles() []MobileDeviceProfile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *MobileDeviceTvOsInventory) GetProfilesOk() (*[]MobileDeviceProfile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *MobileDeviceTvOsInventory) SetProfiles(v []MobileDeviceProfile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *MobileDeviceTvOsInventory) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *MobileDeviceTvOsInventory) GetExtensionAttributes() []MobileDeviceExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *MobileDeviceTvOsInventory) GetExtensionAttributesOk() (*[]MobileDeviceExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *MobileDeviceTvOsInventory) SetExtensionAttributes(v []MobileDeviceExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *MobileDeviceTvOsInventory) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.

### GetGeneral

`func (o *MobileDeviceTvOsInventory) GetGeneral() MobileDeviceTvOsGeneral`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *MobileDeviceTvOsInventory) GetGeneralOk() (*MobileDeviceTvOsGeneral, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *MobileDeviceTvOsInventory) SetGeneral(v MobileDeviceTvOsGeneral)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *MobileDeviceTvOsInventory) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetPurchasing

`func (o *MobileDeviceTvOsInventory) GetPurchasing() MobileDevicePurchasing`

GetPurchasing returns the Purchasing field if non-nil, zero value otherwise.

### GetPurchasingOk

`func (o *MobileDeviceTvOsInventory) GetPurchasingOk() (*MobileDevicePurchasing, bool)`

GetPurchasingOk returns a tuple with the Purchasing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasing

`func (o *MobileDeviceTvOsInventory) SetPurchasing(v MobileDevicePurchasing)`

SetPurchasing sets Purchasing field to given value.

### HasPurchasing

`func (o *MobileDeviceTvOsInventory) HasPurchasing() bool`

HasPurchasing returns a boolean if a field has been set.

### GetUserProfiles

`func (o *MobileDeviceTvOsInventory) GetUserProfiles() []MobileDeviceUserProfile`

GetUserProfiles returns the UserProfiles field if non-nil, zero value otherwise.

### GetUserProfilesOk

`func (o *MobileDeviceTvOsInventory) GetUserProfilesOk() (*[]MobileDeviceUserProfile, bool)`

GetUserProfilesOk returns a tuple with the UserProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfiles

`func (o *MobileDeviceTvOsInventory) SetUserProfiles(v []MobileDeviceUserProfile)`

SetUserProfiles sets UserProfiles field to given value.

### HasUserProfiles

`func (o *MobileDeviceTvOsInventory) HasUserProfiles() bool`

HasUserProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


