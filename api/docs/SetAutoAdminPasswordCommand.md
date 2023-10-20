# SetAutoAdminPasswordCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | Pointer to **string** | The unique identifier of the local administrator account. Must match the GUID of an administrator account that MDM created during Device Enrollment Program (DEP) enrollment. | [optional] 
**Password** | Pointer to **string** | The new password for the local administrator account. | [optional] 

## Methods

### NewSetAutoAdminPasswordCommand

`func NewSetAutoAdminPasswordCommand() *SetAutoAdminPasswordCommand`

NewSetAutoAdminPasswordCommand instantiates a new SetAutoAdminPasswordCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetAutoAdminPasswordCommandWithDefaults

`func NewSetAutoAdminPasswordCommandWithDefaults() *SetAutoAdminPasswordCommand`

NewSetAutoAdminPasswordCommandWithDefaults instantiates a new SetAutoAdminPasswordCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuid

`func (o *SetAutoAdminPasswordCommand) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *SetAutoAdminPasswordCommand) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *SetAutoAdminPasswordCommand) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *SetAutoAdminPasswordCommand) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetPassword

`func (o *SetAutoAdminPasswordCommand) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SetAutoAdminPasswordCommand) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SetAutoAdminPasswordCommand) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SetAutoAdminPasswordCommand) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


