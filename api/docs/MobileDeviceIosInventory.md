# MobileDeviceIosInventory

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
**General** | Pointer to [**MobileDeviceIosGeneral**](MobileDeviceIosGeneral.md) |  | [optional] 
**Security** | Pointer to [**MobileDeviceSecurity**](MobileDeviceSecurity.md) |  | [optional] 
**Ebooks** | Pointer to [**[]MobileDeviceEbookInventoryDetail**](MobileDeviceEbookInventoryDetail.md) |  | [optional] 
**Network** | Pointer to [**MobileDeviceNetwork**](MobileDeviceNetwork.md) |  | [optional] 
**ServiceSubscriptions** | Pointer to [**[]MobileDeviceServiceSubscriptions**](MobileDeviceServiceSubscriptions.md) |  | [optional] 
**ProvisioningProfiles** | Pointer to [**[]MobileDeviceProvisioningProfiles**](MobileDeviceProvisioningProfiles.md) |  | [optional] 
**SharedUsers** | Pointer to [**[]MobileDeviceSharedUser**](MobileDeviceSharedUser.md) |  | [optional] 

## Methods

### NewMobileDeviceIosInventory

`func NewMobileDeviceIosInventory(deviceType string, ) *MobileDeviceIosInventory`

NewMobileDeviceIosInventory instantiates a new MobileDeviceIosInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceIosInventoryWithDefaults

`func NewMobileDeviceIosInventoryWithDefaults() *MobileDeviceIosInventory`

NewMobileDeviceIosInventoryWithDefaults instantiates a new MobileDeviceIosInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMobileDeviceId

`func (o *MobileDeviceIosInventory) GetMobileDeviceId() string`

GetMobileDeviceId returns the MobileDeviceId field if non-nil, zero value otherwise.

### GetMobileDeviceIdOk

`func (o *MobileDeviceIosInventory) GetMobileDeviceIdOk() (*string, bool)`

GetMobileDeviceIdOk returns a tuple with the MobileDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileDeviceId

`func (o *MobileDeviceIosInventory) SetMobileDeviceId(v string)`

SetMobileDeviceId sets MobileDeviceId field to given value.

### HasMobileDeviceId

`func (o *MobileDeviceIosInventory) HasMobileDeviceId() bool`

HasMobileDeviceId returns a boolean if a field has been set.

### GetDeviceType

`func (o *MobileDeviceIosInventory) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *MobileDeviceIosInventory) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *MobileDeviceIosInventory) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.


### GetHardware

`func (o *MobileDeviceIosInventory) GetHardware() MobileDeviceHardware`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *MobileDeviceIosInventory) GetHardwareOk() (*MobileDeviceHardware, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *MobileDeviceIosInventory) SetHardware(v MobileDeviceHardware)`

SetHardware sets Hardware field to given value.

### HasHardware

`func (o *MobileDeviceIosInventory) HasHardware() bool`

HasHardware returns a boolean if a field has been set.

### GetUserAndLocation

`func (o *MobileDeviceIosInventory) GetUserAndLocation() MobileDeviceUserAndLocation`

GetUserAndLocation returns the UserAndLocation field if non-nil, zero value otherwise.

### GetUserAndLocationOk

`func (o *MobileDeviceIosInventory) GetUserAndLocationOk() (*MobileDeviceUserAndLocation, bool)`

GetUserAndLocationOk returns a tuple with the UserAndLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAndLocation

`func (o *MobileDeviceIosInventory) SetUserAndLocation(v MobileDeviceUserAndLocation)`

SetUserAndLocation sets UserAndLocation field to given value.

### HasUserAndLocation

`func (o *MobileDeviceIosInventory) HasUserAndLocation() bool`

HasUserAndLocation returns a boolean if a field has been set.

### GetPurchasing

`func (o *MobileDeviceIosInventory) GetPurchasing() MobileDevicePurchasing`

GetPurchasing returns the Purchasing field if non-nil, zero value otherwise.

### GetPurchasingOk

`func (o *MobileDeviceIosInventory) GetPurchasingOk() (*MobileDevicePurchasing, bool)`

GetPurchasingOk returns a tuple with the Purchasing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasing

`func (o *MobileDeviceIosInventory) SetPurchasing(v MobileDevicePurchasing)`

SetPurchasing sets Purchasing field to given value.

### HasPurchasing

`func (o *MobileDeviceIosInventory) HasPurchasing() bool`

HasPurchasing returns a boolean if a field has been set.

### GetApplications

`func (o *MobileDeviceIosInventory) GetApplications() []MobileDeviceApplicationInventoryDetail`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *MobileDeviceIosInventory) GetApplicationsOk() (*[]MobileDeviceApplicationInventoryDetail, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *MobileDeviceIosInventory) SetApplications(v []MobileDeviceApplicationInventoryDetail)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *MobileDeviceIosInventory) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetCertificates

`func (o *MobileDeviceIosInventory) GetCertificates() []MobileDeviceCertificate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *MobileDeviceIosInventory) GetCertificatesOk() (*[]MobileDeviceCertificate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *MobileDeviceIosInventory) SetCertificates(v []MobileDeviceCertificate)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *MobileDeviceIosInventory) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetProfiles

`func (o *MobileDeviceIosInventory) GetProfiles() []MobileDeviceProfile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *MobileDeviceIosInventory) GetProfilesOk() (*[]MobileDeviceProfile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *MobileDeviceIosInventory) SetProfiles(v []MobileDeviceProfile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *MobileDeviceIosInventory) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetUserProfiles

`func (o *MobileDeviceIosInventory) GetUserProfiles() []MobileDeviceUserProfile`

GetUserProfiles returns the UserProfiles field if non-nil, zero value otherwise.

### GetUserProfilesOk

`func (o *MobileDeviceIosInventory) GetUserProfilesOk() (*[]MobileDeviceUserProfile, bool)`

GetUserProfilesOk returns a tuple with the UserProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfiles

`func (o *MobileDeviceIosInventory) SetUserProfiles(v []MobileDeviceUserProfile)`

SetUserProfiles sets UserProfiles field to given value.

### HasUserProfiles

`func (o *MobileDeviceIosInventory) HasUserProfiles() bool`

HasUserProfiles returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *MobileDeviceIosInventory) GetExtensionAttributes() []MobileDeviceExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *MobileDeviceIosInventory) GetExtensionAttributesOk() (*[]MobileDeviceExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *MobileDeviceIosInventory) SetExtensionAttributes(v []MobileDeviceExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *MobileDeviceIosInventory) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.

### GetGeneral

`func (o *MobileDeviceIosInventory) GetGeneral() MobileDeviceIosGeneral`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *MobileDeviceIosInventory) GetGeneralOk() (*MobileDeviceIosGeneral, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *MobileDeviceIosInventory) SetGeneral(v MobileDeviceIosGeneral)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *MobileDeviceIosInventory) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetSecurity

`func (o *MobileDeviceIosInventory) GetSecurity() MobileDeviceSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *MobileDeviceIosInventory) GetSecurityOk() (*MobileDeviceSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *MobileDeviceIosInventory) SetSecurity(v MobileDeviceSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *MobileDeviceIosInventory) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetEbooks

`func (o *MobileDeviceIosInventory) GetEbooks() []MobileDeviceEbookInventoryDetail`

GetEbooks returns the Ebooks field if non-nil, zero value otherwise.

### GetEbooksOk

`func (o *MobileDeviceIosInventory) GetEbooksOk() (*[]MobileDeviceEbookInventoryDetail, bool)`

GetEbooksOk returns a tuple with the Ebooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbooks

`func (o *MobileDeviceIosInventory) SetEbooks(v []MobileDeviceEbookInventoryDetail)`

SetEbooks sets Ebooks field to given value.

### HasEbooks

`func (o *MobileDeviceIosInventory) HasEbooks() bool`

HasEbooks returns a boolean if a field has been set.

### GetNetwork

`func (o *MobileDeviceIosInventory) GetNetwork() MobileDeviceNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MobileDeviceIosInventory) GetNetworkOk() (*MobileDeviceNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MobileDeviceIosInventory) SetNetwork(v MobileDeviceNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *MobileDeviceIosInventory) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetServiceSubscriptions

`func (o *MobileDeviceIosInventory) GetServiceSubscriptions() []MobileDeviceServiceSubscriptions`

GetServiceSubscriptions returns the ServiceSubscriptions field if non-nil, zero value otherwise.

### GetServiceSubscriptionsOk

`func (o *MobileDeviceIosInventory) GetServiceSubscriptionsOk() (*[]MobileDeviceServiceSubscriptions, bool)`

GetServiceSubscriptionsOk returns a tuple with the ServiceSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSubscriptions

`func (o *MobileDeviceIosInventory) SetServiceSubscriptions(v []MobileDeviceServiceSubscriptions)`

SetServiceSubscriptions sets ServiceSubscriptions field to given value.

### HasServiceSubscriptions

`func (o *MobileDeviceIosInventory) HasServiceSubscriptions() bool`

HasServiceSubscriptions returns a boolean if a field has been set.

### GetProvisioningProfiles

`func (o *MobileDeviceIosInventory) GetProvisioningProfiles() []MobileDeviceProvisioningProfiles`

GetProvisioningProfiles returns the ProvisioningProfiles field if non-nil, zero value otherwise.

### GetProvisioningProfilesOk

`func (o *MobileDeviceIosInventory) GetProvisioningProfilesOk() (*[]MobileDeviceProvisioningProfiles, bool)`

GetProvisioningProfilesOk returns a tuple with the ProvisioningProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningProfiles

`func (o *MobileDeviceIosInventory) SetProvisioningProfiles(v []MobileDeviceProvisioningProfiles)`

SetProvisioningProfiles sets ProvisioningProfiles field to given value.

### HasProvisioningProfiles

`func (o *MobileDeviceIosInventory) HasProvisioningProfiles() bool`

HasProvisioningProfiles returns a boolean if a field has been set.

### GetSharedUsers

`func (o *MobileDeviceIosInventory) GetSharedUsers() []MobileDeviceSharedUser`

GetSharedUsers returns the SharedUsers field if non-nil, zero value otherwise.

### GetSharedUsersOk

`func (o *MobileDeviceIosInventory) GetSharedUsersOk() (*[]MobileDeviceSharedUser, bool)`

GetSharedUsersOk returns a tuple with the SharedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedUsers

`func (o *MobileDeviceIosInventory) SetSharedUsers(v []MobileDeviceSharedUser)`

SetSharedUsers sets SharedUsers field to given value.

### HasSharedUsers

`func (o *MobileDeviceIosInventory) HasSharedUsers() bool`

HasSharedUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


