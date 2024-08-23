# MobileDeviceResponse

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
**Ebooks** | Pointer to [**[]MobileDeviceEbookInventoryDetail**](MobileDeviceEbookInventoryDetail.md) |  | [optional] 
**Network** | Pointer to [**MobileDeviceNetwork**](MobileDeviceNetwork.md) |  | [optional] 
**ServiceSubscriptions** | Pointer to [**[]MobileDeviceServiceSubscriptions**](MobileDeviceServiceSubscriptions.md) |  | [optional] 
**ProvisioningProfiles** | Pointer to [**[]MobileDeviceProvisioningProfiles**](MobileDeviceProvisioningProfiles.md) |  | [optional] 
**SharedUsers** | Pointer to [**[]MobileDeviceSharedUser**](MobileDeviceSharedUser.md) |  | [optional] 
**Purchasing** | Pointer to [**MobileDevicePurchasing**](MobileDevicePurchasing.md) |  | [optional] 
**UserProfiles** | Pointer to [**[]MobileDeviceUserProfile**](MobileDeviceUserProfile.md) |  | [optional] 

## Methods

### NewMobileDeviceResponse

`func NewMobileDeviceResponse(deviceType string, ) *MobileDeviceResponse`

NewMobileDeviceResponse instantiates a new MobileDeviceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceResponseWithDefaults

`func NewMobileDeviceResponseWithDefaults() *MobileDeviceResponse`

NewMobileDeviceResponseWithDefaults instantiates a new MobileDeviceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMobileDeviceId

`func (o *MobileDeviceResponse) GetMobileDeviceId() string`

GetMobileDeviceId returns the MobileDeviceId field if non-nil, zero value otherwise.

### GetMobileDeviceIdOk

`func (o *MobileDeviceResponse) GetMobileDeviceIdOk() (*string, bool)`

GetMobileDeviceIdOk returns a tuple with the MobileDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileDeviceId

`func (o *MobileDeviceResponse) SetMobileDeviceId(v string)`

SetMobileDeviceId sets MobileDeviceId field to given value.

### HasMobileDeviceId

`func (o *MobileDeviceResponse) HasMobileDeviceId() bool`

HasMobileDeviceId returns a boolean if a field has been set.

### GetDeviceType

`func (o *MobileDeviceResponse) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *MobileDeviceResponse) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *MobileDeviceResponse) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.


### GetHardware

`func (o *MobileDeviceResponse) GetHardware() MobileDeviceHardware`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *MobileDeviceResponse) GetHardwareOk() (*MobileDeviceHardware, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *MobileDeviceResponse) SetHardware(v MobileDeviceHardware)`

SetHardware sets Hardware field to given value.

### HasHardware

`func (o *MobileDeviceResponse) HasHardware() bool`

HasHardware returns a boolean if a field has been set.

### GetUserAndLocation

`func (o *MobileDeviceResponse) GetUserAndLocation() MobileDeviceUserAndLocation`

GetUserAndLocation returns the UserAndLocation field if non-nil, zero value otherwise.

### GetUserAndLocationOk

`func (o *MobileDeviceResponse) GetUserAndLocationOk() (*MobileDeviceUserAndLocation, bool)`

GetUserAndLocationOk returns a tuple with the UserAndLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAndLocation

`func (o *MobileDeviceResponse) SetUserAndLocation(v MobileDeviceUserAndLocation)`

SetUserAndLocation sets UserAndLocation field to given value.

### HasUserAndLocation

`func (o *MobileDeviceResponse) HasUserAndLocation() bool`

HasUserAndLocation returns a boolean if a field has been set.

### GetApplications

`func (o *MobileDeviceResponse) GetApplications() []MobileDeviceApplicationInventoryDetail`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *MobileDeviceResponse) GetApplicationsOk() (*[]MobileDeviceApplicationInventoryDetail, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *MobileDeviceResponse) SetApplications(v []MobileDeviceApplicationInventoryDetail)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *MobileDeviceResponse) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetCertificates

`func (o *MobileDeviceResponse) GetCertificates() []MobileDeviceCertificate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *MobileDeviceResponse) GetCertificatesOk() (*[]MobileDeviceCertificate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *MobileDeviceResponse) SetCertificates(v []MobileDeviceCertificate)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *MobileDeviceResponse) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetProfiles

`func (o *MobileDeviceResponse) GetProfiles() []MobileDeviceProfile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *MobileDeviceResponse) GetProfilesOk() (*[]MobileDeviceProfile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *MobileDeviceResponse) SetProfiles(v []MobileDeviceProfile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *MobileDeviceResponse) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *MobileDeviceResponse) GetExtensionAttributes() []MobileDeviceExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *MobileDeviceResponse) GetExtensionAttributesOk() (*[]MobileDeviceExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *MobileDeviceResponse) SetExtensionAttributes(v []MobileDeviceExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *MobileDeviceResponse) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.

### GetGeneral

`func (o *MobileDeviceResponse) GetGeneral() MobileDeviceWatchOsGeneral`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *MobileDeviceResponse) GetGeneralOk() (*MobileDeviceWatchOsGeneral, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *MobileDeviceResponse) SetGeneral(v MobileDeviceWatchOsGeneral)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *MobileDeviceResponse) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetSecurity

`func (o *MobileDeviceResponse) GetSecurity() MobileDeviceSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *MobileDeviceResponse) GetSecurityOk() (*MobileDeviceSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *MobileDeviceResponse) SetSecurity(v MobileDeviceSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *MobileDeviceResponse) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetEbooks

`func (o *MobileDeviceResponse) GetEbooks() []MobileDeviceEbookInventoryDetail`

GetEbooks returns the Ebooks field if non-nil, zero value otherwise.

### GetEbooksOk

`func (o *MobileDeviceResponse) GetEbooksOk() (*[]MobileDeviceEbookInventoryDetail, bool)`

GetEbooksOk returns a tuple with the Ebooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbooks

`func (o *MobileDeviceResponse) SetEbooks(v []MobileDeviceEbookInventoryDetail)`

SetEbooks sets Ebooks field to given value.

### HasEbooks

`func (o *MobileDeviceResponse) HasEbooks() bool`

HasEbooks returns a boolean if a field has been set.

### GetNetwork

`func (o *MobileDeviceResponse) GetNetwork() MobileDeviceNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MobileDeviceResponse) GetNetworkOk() (*MobileDeviceNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MobileDeviceResponse) SetNetwork(v MobileDeviceNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *MobileDeviceResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetServiceSubscriptions

`func (o *MobileDeviceResponse) GetServiceSubscriptions() []MobileDeviceServiceSubscriptions`

GetServiceSubscriptions returns the ServiceSubscriptions field if non-nil, zero value otherwise.

### GetServiceSubscriptionsOk

`func (o *MobileDeviceResponse) GetServiceSubscriptionsOk() (*[]MobileDeviceServiceSubscriptions, bool)`

GetServiceSubscriptionsOk returns a tuple with the ServiceSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSubscriptions

`func (o *MobileDeviceResponse) SetServiceSubscriptions(v []MobileDeviceServiceSubscriptions)`

SetServiceSubscriptions sets ServiceSubscriptions field to given value.

### HasServiceSubscriptions

`func (o *MobileDeviceResponse) HasServiceSubscriptions() bool`

HasServiceSubscriptions returns a boolean if a field has been set.

### GetProvisioningProfiles

`func (o *MobileDeviceResponse) GetProvisioningProfiles() []MobileDeviceProvisioningProfiles`

GetProvisioningProfiles returns the ProvisioningProfiles field if non-nil, zero value otherwise.

### GetProvisioningProfilesOk

`func (o *MobileDeviceResponse) GetProvisioningProfilesOk() (*[]MobileDeviceProvisioningProfiles, bool)`

GetProvisioningProfilesOk returns a tuple with the ProvisioningProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningProfiles

`func (o *MobileDeviceResponse) SetProvisioningProfiles(v []MobileDeviceProvisioningProfiles)`

SetProvisioningProfiles sets ProvisioningProfiles field to given value.

### HasProvisioningProfiles

`func (o *MobileDeviceResponse) HasProvisioningProfiles() bool`

HasProvisioningProfiles returns a boolean if a field has been set.

### GetSharedUsers

`func (o *MobileDeviceResponse) GetSharedUsers() []MobileDeviceSharedUser`

GetSharedUsers returns the SharedUsers field if non-nil, zero value otherwise.

### GetSharedUsersOk

`func (o *MobileDeviceResponse) GetSharedUsersOk() (*[]MobileDeviceSharedUser, bool)`

GetSharedUsersOk returns a tuple with the SharedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedUsers

`func (o *MobileDeviceResponse) SetSharedUsers(v []MobileDeviceSharedUser)`

SetSharedUsers sets SharedUsers field to given value.

### HasSharedUsers

`func (o *MobileDeviceResponse) HasSharedUsers() bool`

HasSharedUsers returns a boolean if a field has been set.

### GetPurchasing

`func (o *MobileDeviceResponse) GetPurchasing() MobileDevicePurchasing`

GetPurchasing returns the Purchasing field if non-nil, zero value otherwise.

### GetPurchasingOk

`func (o *MobileDeviceResponse) GetPurchasingOk() (*MobileDevicePurchasing, bool)`

GetPurchasingOk returns a tuple with the Purchasing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasing

`func (o *MobileDeviceResponse) SetPurchasing(v MobileDevicePurchasing)`

SetPurchasing sets Purchasing field to given value.

### HasPurchasing

`func (o *MobileDeviceResponse) HasPurchasing() bool`

HasPurchasing returns a boolean if a field has been set.

### GetUserProfiles

`func (o *MobileDeviceResponse) GetUserProfiles() []MobileDeviceUserProfile`

GetUserProfiles returns the UserProfiles field if non-nil, zero value otherwise.

### GetUserProfilesOk

`func (o *MobileDeviceResponse) GetUserProfilesOk() (*[]MobileDeviceUserProfile, bool)`

GetUserProfilesOk returns a tuple with the UserProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfiles

`func (o *MobileDeviceResponse) SetUserProfiles(v []MobileDeviceUserProfile)`

SetUserProfiles sets UserProfiles field to given value.

### HasUserProfiles

`func (o *MobileDeviceResponse) HasUserProfiles() bool`

HasUserProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


