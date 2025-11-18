# UnlockUserAccountCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandType** | [**MdmCommandType**](MdmCommandType.md) |  | 
**UserName** | Pointer to **string** | The username of the user account to unlock | [optional] 

## Methods

### NewUnlockUserAccountCommand

`func NewUnlockUserAccountCommand(commandType MdmCommandType, ) *UnlockUserAccountCommand`

NewUnlockUserAccountCommand instantiates a new UnlockUserAccountCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnlockUserAccountCommandWithDefaults

`func NewUnlockUserAccountCommandWithDefaults() *UnlockUserAccountCommand`

NewUnlockUserAccountCommandWithDefaults instantiates a new UnlockUserAccountCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandType

`func (o *UnlockUserAccountCommand) GetCommandType() MdmCommandType`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *UnlockUserAccountCommand) GetCommandTypeOk() (*MdmCommandType, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *UnlockUserAccountCommand) SetCommandType(v MdmCommandType)`

SetCommandType sets CommandType field to given value.


### GetUserName

`func (o *UnlockUserAccountCommand) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UnlockUserAccountCommand) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UnlockUserAccountCommand) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UnlockUserAccountCommand) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


