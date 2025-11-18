# ManagedApplicationListCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandType** | [**MdmCommandType**](MdmCommandType.md) |  | 
**Identifiers** | Pointer to **[]string** | Array of application identifiers to manage | [optional] 

## Methods

### NewManagedApplicationListCommand

`func NewManagedApplicationListCommand(commandType MdmCommandType, ) *ManagedApplicationListCommand`

NewManagedApplicationListCommand instantiates a new ManagedApplicationListCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedApplicationListCommandWithDefaults

`func NewManagedApplicationListCommandWithDefaults() *ManagedApplicationListCommand`

NewManagedApplicationListCommandWithDefaults instantiates a new ManagedApplicationListCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandType

`func (o *ManagedApplicationListCommand) GetCommandType() MdmCommandType`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *ManagedApplicationListCommand) GetCommandTypeOk() (*MdmCommandType, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *ManagedApplicationListCommand) SetCommandType(v MdmCommandType)`

SetCommandType sets CommandType field to given value.


### GetIdentifiers

`func (o *ManagedApplicationListCommand) GetIdentifiers() []string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *ManagedApplicationListCommand) GetIdentifiersOk() (*[]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *ManagedApplicationListCommand) SetIdentifiers(v []string)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *ManagedApplicationListCommand) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### SetIdentifiersNil

`func (o *ManagedApplicationListCommand) SetIdentifiersNil(b bool)`

 SetIdentifiersNil sets the value for Identifiers to be an explicit nil

### UnsetIdentifiers
`func (o *ManagedApplicationListCommand) UnsetIdentifiers()`

UnsetIdentifiers ensures that no value is present for Identifiers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


