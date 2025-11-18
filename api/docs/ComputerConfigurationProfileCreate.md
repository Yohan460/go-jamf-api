# ComputerConfigurationProfileCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**LastInstalled** | Pointer to **NullableTime** |  | [optional] 
**Removable** | Pointer to **NullableBool** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**ProfileIdentifier** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewComputerConfigurationProfileCreate

`func NewComputerConfigurationProfileCreate() *ComputerConfigurationProfileCreate`

NewComputerConfigurationProfileCreate instantiates a new ComputerConfigurationProfileCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerConfigurationProfileCreateWithDefaults

`func NewComputerConfigurationProfileCreateWithDefaults() *ComputerConfigurationProfileCreate`

NewComputerConfigurationProfileCreateWithDefaults instantiates a new ComputerConfigurationProfileCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComputerConfigurationProfileCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputerConfigurationProfileCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputerConfigurationProfileCreate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComputerConfigurationProfileCreate) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ComputerConfigurationProfileCreate) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ComputerConfigurationProfileCreate) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUsername

`func (o *ComputerConfigurationProfileCreate) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ComputerConfigurationProfileCreate) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ComputerConfigurationProfileCreate) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ComputerConfigurationProfileCreate) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ComputerConfigurationProfileCreate) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ComputerConfigurationProfileCreate) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetLastInstalled

`func (o *ComputerConfigurationProfileCreate) GetLastInstalled() time.Time`

GetLastInstalled returns the LastInstalled field if non-nil, zero value otherwise.

### GetLastInstalledOk

`func (o *ComputerConfigurationProfileCreate) GetLastInstalledOk() (*time.Time, bool)`

GetLastInstalledOk returns a tuple with the LastInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInstalled

`func (o *ComputerConfigurationProfileCreate) SetLastInstalled(v time.Time)`

SetLastInstalled sets LastInstalled field to given value.

### HasLastInstalled

`func (o *ComputerConfigurationProfileCreate) HasLastInstalled() bool`

HasLastInstalled returns a boolean if a field has been set.

### SetLastInstalledNil

`func (o *ComputerConfigurationProfileCreate) SetLastInstalledNil(b bool)`

 SetLastInstalledNil sets the value for LastInstalled to be an explicit nil

### UnsetLastInstalled
`func (o *ComputerConfigurationProfileCreate) UnsetLastInstalled()`

UnsetLastInstalled ensures that no value is present for LastInstalled, not even an explicit nil
### GetRemovable

`func (o *ComputerConfigurationProfileCreate) GetRemovable() bool`

GetRemovable returns the Removable field if non-nil, zero value otherwise.

### GetRemovableOk

`func (o *ComputerConfigurationProfileCreate) GetRemovableOk() (*bool, bool)`

GetRemovableOk returns a tuple with the Removable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovable

`func (o *ComputerConfigurationProfileCreate) SetRemovable(v bool)`

SetRemovable sets Removable field to given value.

### HasRemovable

`func (o *ComputerConfigurationProfileCreate) HasRemovable() bool`

HasRemovable returns a boolean if a field has been set.

### SetRemovableNil

`func (o *ComputerConfigurationProfileCreate) SetRemovableNil(b bool)`

 SetRemovableNil sets the value for Removable to be an explicit nil

### UnsetRemovable
`func (o *ComputerConfigurationProfileCreate) UnsetRemovable()`

UnsetRemovable ensures that no value is present for Removable, not even an explicit nil
### GetDisplayName

`func (o *ComputerConfigurationProfileCreate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ComputerConfigurationProfileCreate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ComputerConfigurationProfileCreate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ComputerConfigurationProfileCreate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ComputerConfigurationProfileCreate) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ComputerConfigurationProfileCreate) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetProfileIdentifier

`func (o *ComputerConfigurationProfileCreate) GetProfileIdentifier() string`

GetProfileIdentifier returns the ProfileIdentifier field if non-nil, zero value otherwise.

### GetProfileIdentifierOk

`func (o *ComputerConfigurationProfileCreate) GetProfileIdentifierOk() (*string, bool)`

GetProfileIdentifierOk returns a tuple with the ProfileIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIdentifier

`func (o *ComputerConfigurationProfileCreate) SetProfileIdentifier(v string)`

SetProfileIdentifier sets ProfileIdentifier field to given value.

### HasProfileIdentifier

`func (o *ComputerConfigurationProfileCreate) HasProfileIdentifier() bool`

HasProfileIdentifier returns a boolean if a field has been set.

### SetProfileIdentifierNil

`func (o *ComputerConfigurationProfileCreate) SetProfileIdentifierNil(b bool)`

 SetProfileIdentifierNil sets the value for ProfileIdentifier to be an explicit nil

### UnsetProfileIdentifier
`func (o *ComputerConfigurationProfileCreate) UnsetProfileIdentifier()`

UnsetProfileIdentifier ensures that no value is present for ProfileIdentifier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


