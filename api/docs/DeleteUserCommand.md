# DeleteUserCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandType** | [**MdmCommandType**](MdmCommandType.md) |  | 
**UserName** | Pointer to **string** |  | [optional] 
**ForceDeletion** | Pointer to **bool** |  | [optional] 
**DeleteAllUsers** | Pointer to **bool** |  | [optional] 

## Methods

### NewDeleteUserCommand

`func NewDeleteUserCommand(commandType MdmCommandType, ) *DeleteUserCommand`

NewDeleteUserCommand instantiates a new DeleteUserCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteUserCommandWithDefaults

`func NewDeleteUserCommandWithDefaults() *DeleteUserCommand`

NewDeleteUserCommandWithDefaults instantiates a new DeleteUserCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandType

`func (o *DeleteUserCommand) GetCommandType() MdmCommandType`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *DeleteUserCommand) GetCommandTypeOk() (*MdmCommandType, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *DeleteUserCommand) SetCommandType(v MdmCommandType)`

SetCommandType sets CommandType field to given value.


### GetUserName

`func (o *DeleteUserCommand) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *DeleteUserCommand) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *DeleteUserCommand) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *DeleteUserCommand) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetForceDeletion

`func (o *DeleteUserCommand) GetForceDeletion() bool`

GetForceDeletion returns the ForceDeletion field if non-nil, zero value otherwise.

### GetForceDeletionOk

`func (o *DeleteUserCommand) GetForceDeletionOk() (*bool, bool)`

GetForceDeletionOk returns a tuple with the ForceDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceDeletion

`func (o *DeleteUserCommand) SetForceDeletion(v bool)`

SetForceDeletion sets ForceDeletion field to given value.

### HasForceDeletion

`func (o *DeleteUserCommand) HasForceDeletion() bool`

HasForceDeletion returns a boolean if a field has been set.

### GetDeleteAllUsers

`func (o *DeleteUserCommand) GetDeleteAllUsers() bool`

GetDeleteAllUsers returns the DeleteAllUsers field if non-nil, zero value otherwise.

### GetDeleteAllUsersOk

`func (o *DeleteUserCommand) GetDeleteAllUsersOk() (*bool, bool)`

GetDeleteAllUsersOk returns a tuple with the DeleteAllUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAllUsers

`func (o *DeleteUserCommand) SetDeleteAllUsers(v bool)`

SetDeleteAllUsers sets DeleteAllUsers field to given value.

### HasDeleteAllUsers

`func (o *DeleteUserCommand) HasDeleteAllUsers() bool`

HasDeleteAllUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


