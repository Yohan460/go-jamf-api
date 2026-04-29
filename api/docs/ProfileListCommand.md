# ProfileListCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandType** | [**MdmCommandType**](MdmCommandType.md) |  | 
**ManagedOnly** | Pointer to **NullableBool** | If true, only managed configuration profiles are returned. If false, all configuration profiles are returned. The default value is false. | [optional] [default to false]

## Methods

### NewProfileListCommand

`func NewProfileListCommand(commandType MdmCommandType, ) *ProfileListCommand`

NewProfileListCommand instantiates a new ProfileListCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileListCommandWithDefaults

`func NewProfileListCommandWithDefaults() *ProfileListCommand`

NewProfileListCommandWithDefaults instantiates a new ProfileListCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandType

`func (o *ProfileListCommand) GetCommandType() MdmCommandType`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *ProfileListCommand) GetCommandTypeOk() (*MdmCommandType, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *ProfileListCommand) SetCommandType(v MdmCommandType)`

SetCommandType sets CommandType field to given value.


### GetManagedOnly

`func (o *ProfileListCommand) GetManagedOnly() bool`

GetManagedOnly returns the ManagedOnly field if non-nil, zero value otherwise.

### GetManagedOnlyOk

`func (o *ProfileListCommand) GetManagedOnlyOk() (*bool, bool)`

GetManagedOnlyOk returns a tuple with the ManagedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedOnly

`func (o *ProfileListCommand) SetManagedOnly(v bool)`

SetManagedOnly sets ManagedOnly field to given value.

### HasManagedOnly

`func (o *ProfileListCommand) HasManagedOnly() bool`

HasManagedOnly returns a boolean if a field has been set.

### SetManagedOnlyNil

`func (o *ProfileListCommand) SetManagedOnlyNil(b bool)`

 SetManagedOnlyNil sets the value for ManagedOnly to be an explicit nil

### UnsetManagedOnly
`func (o *ProfileListCommand) UnsetManagedOnly()`

UnsetManagedOnly ensures that no value is present for ManagedOnly, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


