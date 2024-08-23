# SetRecoveryLockCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandType** | [**MdmCommandType**](MdmCommandType.md) |  | 
**NewPassword** | Pointer to **string** | The new password for Recovery Lock. Set as an empty string to clear the Recovery Lock password. | [optional] 

## Methods

### NewSetRecoveryLockCommand

`func NewSetRecoveryLockCommand(commandType MdmCommandType, ) *SetRecoveryLockCommand`

NewSetRecoveryLockCommand instantiates a new SetRecoveryLockCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetRecoveryLockCommandWithDefaults

`func NewSetRecoveryLockCommandWithDefaults() *SetRecoveryLockCommand`

NewSetRecoveryLockCommandWithDefaults instantiates a new SetRecoveryLockCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandType

`func (o *SetRecoveryLockCommand) GetCommandType() MdmCommandType`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *SetRecoveryLockCommand) GetCommandTypeOk() (*MdmCommandType, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *SetRecoveryLockCommand) SetCommandType(v MdmCommandType)`

SetCommandType sets CommandType field to given value.


### GetNewPassword

`func (o *SetRecoveryLockCommand) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *SetRecoveryLockCommand) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *SetRecoveryLockCommand) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *SetRecoveryLockCommand) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


