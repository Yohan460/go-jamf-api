# VerifyRecoveryLockCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandType** | [**MdmCommandType**](MdmCommandType.md) |  | 
**Password** | Pointer to **string** | The password to verify. | [optional] 

## Methods

### NewVerifyRecoveryLockCommand

`func NewVerifyRecoveryLockCommand(commandType MdmCommandType, ) *VerifyRecoveryLockCommand`

NewVerifyRecoveryLockCommand instantiates a new VerifyRecoveryLockCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyRecoveryLockCommandWithDefaults

`func NewVerifyRecoveryLockCommandWithDefaults() *VerifyRecoveryLockCommand`

NewVerifyRecoveryLockCommandWithDefaults instantiates a new VerifyRecoveryLockCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandType

`func (o *VerifyRecoveryLockCommand) GetCommandType() MdmCommandType`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *VerifyRecoveryLockCommand) GetCommandTypeOk() (*MdmCommandType, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *VerifyRecoveryLockCommand) SetCommandType(v MdmCommandType)`

SetCommandType sets CommandType field to given value.


### GetPassword

`func (o *VerifyRecoveryLockCommand) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VerifyRecoveryLockCommand) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VerifyRecoveryLockCommand) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VerifyRecoveryLockCommand) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


