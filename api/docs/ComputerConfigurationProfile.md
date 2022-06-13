# ComputerConfigurationProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Username** | Pointer to **string** |  | [optional] [readonly] 
**LastInstalled** | Pointer to **time.Time** |  | [optional] 
**Removable** | Pointer to **bool** |  | [optional] [readonly] 
**DisplayName** | Pointer to **string** |  | [optional] [readonly] 
**ProfileIdentifier** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewComputerConfigurationProfile

`func NewComputerConfigurationProfile() *ComputerConfigurationProfile`

NewComputerConfigurationProfile instantiates a new ComputerConfigurationProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerConfigurationProfileWithDefaults

`func NewComputerConfigurationProfileWithDefaults() *ComputerConfigurationProfile`

NewComputerConfigurationProfileWithDefaults instantiates a new ComputerConfigurationProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComputerConfigurationProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputerConfigurationProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputerConfigurationProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComputerConfigurationProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *ComputerConfigurationProfile) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ComputerConfigurationProfile) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ComputerConfigurationProfile) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ComputerConfigurationProfile) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetLastInstalled

`func (o *ComputerConfigurationProfile) GetLastInstalled() time.Time`

GetLastInstalled returns the LastInstalled field if non-nil, zero value otherwise.

### GetLastInstalledOk

`func (o *ComputerConfigurationProfile) GetLastInstalledOk() (*time.Time, bool)`

GetLastInstalledOk returns a tuple with the LastInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInstalled

`func (o *ComputerConfigurationProfile) SetLastInstalled(v time.Time)`

SetLastInstalled sets LastInstalled field to given value.

### HasLastInstalled

`func (o *ComputerConfigurationProfile) HasLastInstalled() bool`

HasLastInstalled returns a boolean if a field has been set.

### GetRemovable

`func (o *ComputerConfigurationProfile) GetRemovable() bool`

GetRemovable returns the Removable field if non-nil, zero value otherwise.

### GetRemovableOk

`func (o *ComputerConfigurationProfile) GetRemovableOk() (*bool, bool)`

GetRemovableOk returns a tuple with the Removable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovable

`func (o *ComputerConfigurationProfile) SetRemovable(v bool)`

SetRemovable sets Removable field to given value.

### HasRemovable

`func (o *ComputerConfigurationProfile) HasRemovable() bool`

HasRemovable returns a boolean if a field has been set.

### GetDisplayName

`func (o *ComputerConfigurationProfile) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ComputerConfigurationProfile) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ComputerConfigurationProfile) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ComputerConfigurationProfile) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetProfileIdentifier

`func (o *ComputerConfigurationProfile) GetProfileIdentifier() string`

GetProfileIdentifier returns the ProfileIdentifier field if non-nil, zero value otherwise.

### GetProfileIdentifierOk

`func (o *ComputerConfigurationProfile) GetProfileIdentifierOk() (*string, bool)`

GetProfileIdentifierOk returns a tuple with the ProfileIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIdentifier

`func (o *ComputerConfigurationProfile) SetProfileIdentifier(v string)`

SetProfileIdentifier sets ProfileIdentifier field to given value.

### HasProfileIdentifier

`func (o *ComputerConfigurationProfile) HasProfileIdentifier() bool`

HasProfileIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


