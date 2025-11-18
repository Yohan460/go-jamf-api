# ClearPasscodeCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandType** | [**MdmCommandType**](MdmCommandType.md) |  | 
**UnlockToken** | **string** |  | 

## Methods

### NewClearPasscodeCommand

`func NewClearPasscodeCommand(commandType MdmCommandType, unlockToken string, ) *ClearPasscodeCommand`

NewClearPasscodeCommand instantiates a new ClearPasscodeCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearPasscodeCommandWithDefaults

`func NewClearPasscodeCommandWithDefaults() *ClearPasscodeCommand`

NewClearPasscodeCommandWithDefaults instantiates a new ClearPasscodeCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandType

`func (o *ClearPasscodeCommand) GetCommandType() MdmCommandType`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *ClearPasscodeCommand) GetCommandTypeOk() (*MdmCommandType, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *ClearPasscodeCommand) SetCommandType(v MdmCommandType)`

SetCommandType sets CommandType field to given value.


### GetUnlockToken

`func (o *ClearPasscodeCommand) GetUnlockToken() string`

GetUnlockToken returns the UnlockToken field if non-nil, zero value otherwise.

### GetUnlockTokenOk

`func (o *ClearPasscodeCommand) GetUnlockTokenOk() (*string, bool)`

GetUnlockTokenOk returns a tuple with the UnlockToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockToken

`func (o *ClearPasscodeCommand) SetUnlockToken(v string)`

SetUnlockToken sets UnlockToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


