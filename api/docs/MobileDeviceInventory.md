# MobileDeviceInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MobileDeviceId** | Pointer to **string** |  | [optional] [readonly] 
**DeviceType** | **string** | Based on the value of this type either ios or appleTv objects will be populated. | 
**Hardware** | Pointer to [**MobileDeviceHardware**](MobileDeviceHardware.md) |  | [optional] 
**UserAndLocation** | Pointer to [**MobileDeviceUserAndLocation**](MobileDeviceUserAndLocation.md) |  | [optional] 
**Purchasing** | Pointer to [**MobileDevicePurchasing**](MobileDevicePurchasing.md) |  | [optional] 
**Applications** | Pointer to [**[]MobileDeviceApplicationInventoryDetail**](MobileDeviceApplicationInventoryDetail.md) |  | [optional] 
**Certificates** | Pointer to [**[]MobileDeviceCertificate**](MobileDeviceCertificate.md) |  | [optional] 
**Profiles** | Pointer to [**[]MobileDeviceProfile**](MobileDeviceProfile.md) |  | [optional] 
**UserProfiles** | Pointer to [**[]MobileDeviceUserProfile**](MobileDeviceUserProfile.md) |  | [optional] 
**ExtensionAttributes** | Pointer to [**[]MobileDeviceExtensionAttribute**](MobileDeviceExtensionAttribute.md) |  | [optional] 

## Methods

### NewMobileDeviceInventory

`func NewMobileDeviceInventory(deviceType string, ) *MobileDeviceInventory`

NewMobileDeviceInventory instantiates a new MobileDeviceInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceInventoryWithDefaults

`func NewMobileDeviceInventoryWithDefaults() *MobileDeviceInventory`

NewMobileDeviceInventoryWithDefaults instantiates a new MobileDeviceInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMobileDeviceId

`func (o *MobileDeviceInventory) GetMobileDeviceId() string`

GetMobileDeviceId returns the MobileDeviceId field if non-nil, zero value otherwise.

### GetMobileDeviceIdOk

`func (o *MobileDeviceInventory) GetMobileDeviceIdOk() (*string, bool)`

GetMobileDeviceIdOk returns a tuple with the MobileDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileDeviceId

`func (o *MobileDeviceInventory) SetMobileDeviceId(v string)`

SetMobileDeviceId sets MobileDeviceId field to given value.

### HasMobileDeviceId

`func (o *MobileDeviceInventory) HasMobileDeviceId() bool`

HasMobileDeviceId returns a boolean if a field has been set.

### GetDeviceType

`func (o *MobileDeviceInventory) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *MobileDeviceInventory) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *MobileDeviceInventory) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.


### GetHardware

`func (o *MobileDeviceInventory) GetHardware() MobileDeviceHardware`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *MobileDeviceInventory) GetHardwareOk() (*MobileDeviceHardware, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *MobileDeviceInventory) SetHardware(v MobileDeviceHardware)`

SetHardware sets Hardware field to given value.

### HasHardware

`func (o *MobileDeviceInventory) HasHardware() bool`

HasHardware returns a boolean if a field has been set.

### GetUserAndLocation

`func (o *MobileDeviceInventory) GetUserAndLocation() MobileDeviceUserAndLocation`

GetUserAndLocation returns the UserAndLocation field if non-nil, zero value otherwise.

### GetUserAndLocationOk

`func (o *MobileDeviceInventory) GetUserAndLocationOk() (*MobileDeviceUserAndLocation, bool)`

GetUserAndLocationOk returns a tuple with the UserAndLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAndLocation

`func (o *MobileDeviceInventory) SetUserAndLocation(v MobileDeviceUserAndLocation)`

SetUserAndLocation sets UserAndLocation field to given value.

### HasUserAndLocation

`func (o *MobileDeviceInventory) HasUserAndLocation() bool`

HasUserAndLocation returns a boolean if a field has been set.

### GetPurchasing

`func (o *MobileDeviceInventory) GetPurchasing() MobileDevicePurchasing`

GetPurchasing returns the Purchasing field if non-nil, zero value otherwise.

### GetPurchasingOk

`func (o *MobileDeviceInventory) GetPurchasingOk() (*MobileDevicePurchasing, bool)`

GetPurchasingOk returns a tuple with the Purchasing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasing

`func (o *MobileDeviceInventory) SetPurchasing(v MobileDevicePurchasing)`

SetPurchasing sets Purchasing field to given value.

### HasPurchasing

`func (o *MobileDeviceInventory) HasPurchasing() bool`

HasPurchasing returns a boolean if a field has been set.

### GetApplications

`func (o *MobileDeviceInventory) GetApplications() []MobileDeviceApplicationInventoryDetail`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *MobileDeviceInventory) GetApplicationsOk() (*[]MobileDeviceApplicationInventoryDetail, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *MobileDeviceInventory) SetApplications(v []MobileDeviceApplicationInventoryDetail)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *MobileDeviceInventory) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetCertificates

`func (o *MobileDeviceInventory) GetCertificates() []MobileDeviceCertificate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *MobileDeviceInventory) GetCertificatesOk() (*[]MobileDeviceCertificate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *MobileDeviceInventory) SetCertificates(v []MobileDeviceCertificate)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *MobileDeviceInventory) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetProfiles

`func (o *MobileDeviceInventory) GetProfiles() []MobileDeviceProfile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *MobileDeviceInventory) GetProfilesOk() (*[]MobileDeviceProfile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *MobileDeviceInventory) SetProfiles(v []MobileDeviceProfile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *MobileDeviceInventory) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetUserProfiles

`func (o *MobileDeviceInventory) GetUserProfiles() []MobileDeviceUserProfile`

GetUserProfiles returns the UserProfiles field if non-nil, zero value otherwise.

### GetUserProfilesOk

`func (o *MobileDeviceInventory) GetUserProfilesOk() (*[]MobileDeviceUserProfile, bool)`

GetUserProfilesOk returns a tuple with the UserProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfiles

`func (o *MobileDeviceInventory) SetUserProfiles(v []MobileDeviceUserProfile)`

SetUserProfiles sets UserProfiles field to given value.

### HasUserProfiles

`func (o *MobileDeviceInventory) HasUserProfiles() bool`

HasUserProfiles returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *MobileDeviceInventory) GetExtensionAttributes() []MobileDeviceExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *MobileDeviceInventory) GetExtensionAttributesOk() (*[]MobileDeviceExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *MobileDeviceInventory) SetExtensionAttributes(v []MobileDeviceExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *MobileDeviceInventory) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


