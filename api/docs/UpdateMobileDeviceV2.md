# UpdateMobileDeviceV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Mobile Device Name. When updated, Jamf Pro sends an MDM settings command to the device (device must be supervised). | [optional] 
**EnforceName** | Pointer to **bool** | Enforce the mobile device name. Device must be supervised. If set to true, Jamf Pro will revert the Mobile Device Name to the ‘name’ value each time the device checks in. | [optional] 
**AssetTag** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**TimeZone** | Pointer to **string** | IANA time zone database name | [optional] 
**Location** | Pointer to [**LocationV2**](LocationV2.md) |  | [optional] 
**UpdatedExtensionAttributes** | Pointer to [**[]ExtensionAttributeV2**](ExtensionAttributeV2.md) |  | [optional] 
**Ios** | Pointer to [**UpdateIosV2**](UpdateIosV2.md) |  | [optional] 
**Tvos** | Pointer to [**UpdateTvOs**](UpdateTvOs.md) |  | [optional] 

## Methods

### NewUpdateMobileDeviceV2

`func NewUpdateMobileDeviceV2() *UpdateMobileDeviceV2`

NewUpdateMobileDeviceV2 instantiates a new UpdateMobileDeviceV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMobileDeviceV2WithDefaults

`func NewUpdateMobileDeviceV2WithDefaults() *UpdateMobileDeviceV2`

NewUpdateMobileDeviceV2WithDefaults instantiates a new UpdateMobileDeviceV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateMobileDeviceV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMobileDeviceV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMobileDeviceV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMobileDeviceV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnforceName

`func (o *UpdateMobileDeviceV2) GetEnforceName() bool`

GetEnforceName returns the EnforceName field if non-nil, zero value otherwise.

### GetEnforceNameOk

`func (o *UpdateMobileDeviceV2) GetEnforceNameOk() (*bool, bool)`

GetEnforceNameOk returns a tuple with the EnforceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceName

`func (o *UpdateMobileDeviceV2) SetEnforceName(v bool)`

SetEnforceName sets EnforceName field to given value.

### HasEnforceName

`func (o *UpdateMobileDeviceV2) HasEnforceName() bool`

HasEnforceName returns a boolean if a field has been set.

### GetAssetTag

`func (o *UpdateMobileDeviceV2) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *UpdateMobileDeviceV2) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *UpdateMobileDeviceV2) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *UpdateMobileDeviceV2) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetSiteId

`func (o *UpdateMobileDeviceV2) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *UpdateMobileDeviceV2) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *UpdateMobileDeviceV2) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *UpdateMobileDeviceV2) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetTimeZone

`func (o *UpdateMobileDeviceV2) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *UpdateMobileDeviceV2) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *UpdateMobileDeviceV2) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *UpdateMobileDeviceV2) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetLocation

`func (o *UpdateMobileDeviceV2) GetLocation() LocationV2`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UpdateMobileDeviceV2) GetLocationOk() (*LocationV2, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UpdateMobileDeviceV2) SetLocation(v LocationV2)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UpdateMobileDeviceV2) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetUpdatedExtensionAttributes

`func (o *UpdateMobileDeviceV2) GetUpdatedExtensionAttributes() []ExtensionAttributeV2`

GetUpdatedExtensionAttributes returns the UpdatedExtensionAttributes field if non-nil, zero value otherwise.

### GetUpdatedExtensionAttributesOk

`func (o *UpdateMobileDeviceV2) GetUpdatedExtensionAttributesOk() (*[]ExtensionAttributeV2, bool)`

GetUpdatedExtensionAttributesOk returns a tuple with the UpdatedExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedExtensionAttributes

`func (o *UpdateMobileDeviceV2) SetUpdatedExtensionAttributes(v []ExtensionAttributeV2)`

SetUpdatedExtensionAttributes sets UpdatedExtensionAttributes field to given value.

### HasUpdatedExtensionAttributes

`func (o *UpdateMobileDeviceV2) HasUpdatedExtensionAttributes() bool`

HasUpdatedExtensionAttributes returns a boolean if a field has been set.

### GetIos

`func (o *UpdateMobileDeviceV2) GetIos() UpdateIosV2`

GetIos returns the Ios field if non-nil, zero value otherwise.

### GetIosOk

`func (o *UpdateMobileDeviceV2) GetIosOk() (*UpdateIosV2, bool)`

GetIosOk returns a tuple with the Ios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIos

`func (o *UpdateMobileDeviceV2) SetIos(v UpdateIosV2)`

SetIos sets Ios field to given value.

### HasIos

`func (o *UpdateMobileDeviceV2) HasIos() bool`

HasIos returns a boolean if a field has been set.

### GetTvos

`func (o *UpdateMobileDeviceV2) GetTvos() UpdateTvOs`

GetTvos returns the Tvos field if non-nil, zero value otherwise.

### GetTvosOk

`func (o *UpdateMobileDeviceV2) GetTvosOk() (*UpdateTvOs, bool)`

GetTvosOk returns a tuple with the Tvos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvos

`func (o *UpdateMobileDeviceV2) SetTvos(v UpdateTvOs)`

SetTvos sets Tvos field to given value.

### HasTvos

`func (o *UpdateMobileDeviceV2) HasTvos() bool`

HasTvos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


