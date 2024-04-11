# UpdateMobileDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**UpdatedExtensionAttributes** | Pointer to [**[]ExtensionAttribute**](ExtensionAttribute.md) |  | [optional] 
**Ios** | Pointer to [**UpdateIos**](UpdateIos.md) |  | [optional] 
**AppleTv** | Pointer to [**UpdateAppleTv**](UpdateAppleTv.md) |  | [optional] 

## Methods

### NewUpdateMobileDevice

`func NewUpdateMobileDevice() *UpdateMobileDevice`

NewUpdateMobileDevice instantiates a new UpdateMobileDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMobileDeviceWithDefaults

`func NewUpdateMobileDeviceWithDefaults() *UpdateMobileDevice`

NewUpdateMobileDeviceWithDefaults instantiates a new UpdateMobileDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateMobileDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMobileDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMobileDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMobileDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAssetTag

`func (o *UpdateMobileDevice) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *UpdateMobileDevice) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *UpdateMobileDevice) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *UpdateMobileDevice) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetSiteId

`func (o *UpdateMobileDevice) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *UpdateMobileDevice) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *UpdateMobileDevice) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *UpdateMobileDevice) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetLocation

`func (o *UpdateMobileDevice) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UpdateMobileDevice) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UpdateMobileDevice) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UpdateMobileDevice) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetUpdatedExtensionAttributes

`func (o *UpdateMobileDevice) GetUpdatedExtensionAttributes() []ExtensionAttribute`

GetUpdatedExtensionAttributes returns the UpdatedExtensionAttributes field if non-nil, zero value otherwise.

### GetUpdatedExtensionAttributesOk

`func (o *UpdateMobileDevice) GetUpdatedExtensionAttributesOk() (*[]ExtensionAttribute, bool)`

GetUpdatedExtensionAttributesOk returns a tuple with the UpdatedExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedExtensionAttributes

`func (o *UpdateMobileDevice) SetUpdatedExtensionAttributes(v []ExtensionAttribute)`

SetUpdatedExtensionAttributes sets UpdatedExtensionAttributes field to given value.

### HasUpdatedExtensionAttributes

`func (o *UpdateMobileDevice) HasUpdatedExtensionAttributes() bool`

HasUpdatedExtensionAttributes returns a boolean if a field has been set.

### GetIos

`func (o *UpdateMobileDevice) GetIos() UpdateIos`

GetIos returns the Ios field if non-nil, zero value otherwise.

### GetIosOk

`func (o *UpdateMobileDevice) GetIosOk() (*UpdateIos, bool)`

GetIosOk returns a tuple with the Ios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIos

`func (o *UpdateMobileDevice) SetIos(v UpdateIos)`

SetIos sets Ios field to given value.

### HasIos

`func (o *UpdateMobileDevice) HasIos() bool`

HasIos returns a boolean if a field has been set.

### GetAppleTv

`func (o *UpdateMobileDevice) GetAppleTv() UpdateAppleTv`

GetAppleTv returns the AppleTv field if non-nil, zero value otherwise.

### GetAppleTvOk

`func (o *UpdateMobileDevice) GetAppleTvOk() (*UpdateAppleTv, bool)`

GetAppleTvOk returns a tuple with the AppleTv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleTv

`func (o *UpdateMobileDevice) SetAppleTv(v UpdateAppleTv)`

SetAppleTv sets AppleTv field to given value.

### HasAppleTv

`func (o *UpdateMobileDevice) HasAppleTv() bool`

HasAppleTv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


