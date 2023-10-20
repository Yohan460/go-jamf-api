# MobileDeviceProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Removable** | Pointer to **bool** |  | [optional] 
**LastInstalled** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewMobileDeviceProfile

`func NewMobileDeviceProfile() *MobileDeviceProfile`

NewMobileDeviceProfile instantiates a new MobileDeviceProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceProfileWithDefaults

`func NewMobileDeviceProfileWithDefaults() *MobileDeviceProfile`

NewMobileDeviceProfileWithDefaults instantiates a new MobileDeviceProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *MobileDeviceProfile) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MobileDeviceProfile) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MobileDeviceProfile) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MobileDeviceProfile) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetVersion

`func (o *MobileDeviceProfile) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MobileDeviceProfile) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MobileDeviceProfile) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MobileDeviceProfile) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUuid

`func (o *MobileDeviceProfile) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MobileDeviceProfile) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MobileDeviceProfile) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *MobileDeviceProfile) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetIdentifier

`func (o *MobileDeviceProfile) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *MobileDeviceProfile) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *MobileDeviceProfile) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *MobileDeviceProfile) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetRemovable

`func (o *MobileDeviceProfile) GetRemovable() bool`

GetRemovable returns the Removable field if non-nil, zero value otherwise.

### GetRemovableOk

`func (o *MobileDeviceProfile) GetRemovableOk() (*bool, bool)`

GetRemovableOk returns a tuple with the Removable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovable

`func (o *MobileDeviceProfile) SetRemovable(v bool)`

SetRemovable sets Removable field to given value.

### HasRemovable

`func (o *MobileDeviceProfile) HasRemovable() bool`

HasRemovable returns a boolean if a field has been set.

### GetLastInstalled

`func (o *MobileDeviceProfile) GetLastInstalled() time.Time`

GetLastInstalled returns the LastInstalled field if non-nil, zero value otherwise.

### GetLastInstalledOk

`func (o *MobileDeviceProfile) GetLastInstalledOk() (*time.Time, bool)`

GetLastInstalledOk returns a tuple with the LastInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInstalled

`func (o *MobileDeviceProfile) SetLastInstalled(v time.Time)`

SetLastInstalled sets LastInstalled field to given value.

### HasLastInstalled

`func (o *MobileDeviceProfile) HasLastInstalled() bool`

HasLastInstalled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


