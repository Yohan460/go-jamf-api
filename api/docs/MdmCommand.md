# MdmCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**DateSent** | Pointer to **time.Time** |  | [optional] 
**DateCompleted** | Pointer to **time.Time** |  | [optional] 
**Client** | Pointer to [**MdmCommandClient**](MdmCommandClient.md) |  | [optional] 
**CommandState** | Pointer to [**MdmCommandState**](MdmCommandState.md) |  | [optional] 
**CommandType** | Pointer to [**MdmCommandType**](MdmCommandType.md) |  | [optional] 
**CommandError** | Pointer to [**MdmCommandError**](MdmCommandError.md) |  | [optional] 
**ProfileId** | Pointer to **int64** |  | [optional] 

## Methods

### NewMdmCommand

`func NewMdmCommand() *MdmCommand`

NewMdmCommand instantiates a new MdmCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdmCommandWithDefaults

`func NewMdmCommandWithDefaults() *MdmCommand`

NewMdmCommandWithDefaults instantiates a new MdmCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *MdmCommand) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MdmCommand) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MdmCommand) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *MdmCommand) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDateSent

`func (o *MdmCommand) GetDateSent() time.Time`

GetDateSent returns the DateSent field if non-nil, zero value otherwise.

### GetDateSentOk

`func (o *MdmCommand) GetDateSentOk() (*time.Time, bool)`

GetDateSentOk returns a tuple with the DateSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateSent

`func (o *MdmCommand) SetDateSent(v time.Time)`

SetDateSent sets DateSent field to given value.

### HasDateSent

`func (o *MdmCommand) HasDateSent() bool`

HasDateSent returns a boolean if a field has been set.

### GetDateCompleted

`func (o *MdmCommand) GetDateCompleted() time.Time`

GetDateCompleted returns the DateCompleted field if non-nil, zero value otherwise.

### GetDateCompletedOk

`func (o *MdmCommand) GetDateCompletedOk() (*time.Time, bool)`

GetDateCompletedOk returns a tuple with the DateCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCompleted

`func (o *MdmCommand) SetDateCompleted(v time.Time)`

SetDateCompleted sets DateCompleted field to given value.

### HasDateCompleted

`func (o *MdmCommand) HasDateCompleted() bool`

HasDateCompleted returns a boolean if a field has been set.

### GetClient

`func (o *MdmCommand) GetClient() MdmCommandClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *MdmCommand) GetClientOk() (*MdmCommandClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *MdmCommand) SetClient(v MdmCommandClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *MdmCommand) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetCommandState

`func (o *MdmCommand) GetCommandState() MdmCommandState`

GetCommandState returns the CommandState field if non-nil, zero value otherwise.

### GetCommandStateOk

`func (o *MdmCommand) GetCommandStateOk() (*MdmCommandState, bool)`

GetCommandStateOk returns a tuple with the CommandState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandState

`func (o *MdmCommand) SetCommandState(v MdmCommandState)`

SetCommandState sets CommandState field to given value.

### HasCommandState

`func (o *MdmCommand) HasCommandState() bool`

HasCommandState returns a boolean if a field has been set.

### GetCommandType

`func (o *MdmCommand) GetCommandType() MdmCommandType`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *MdmCommand) GetCommandTypeOk() (*MdmCommandType, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *MdmCommand) SetCommandType(v MdmCommandType)`

SetCommandType sets CommandType field to given value.

### HasCommandType

`func (o *MdmCommand) HasCommandType() bool`

HasCommandType returns a boolean if a field has been set.

### GetCommandError

`func (o *MdmCommand) GetCommandError() MdmCommandError`

GetCommandError returns the CommandError field if non-nil, zero value otherwise.

### GetCommandErrorOk

`func (o *MdmCommand) GetCommandErrorOk() (*MdmCommandError, bool)`

GetCommandErrorOk returns a tuple with the CommandError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandError

`func (o *MdmCommand) SetCommandError(v MdmCommandError)`

SetCommandError sets CommandError field to given value.

### HasCommandError

`func (o *MdmCommand) HasCommandError() bool`

HasCommandError returns a boolean if a field has been set.

### GetProfileId

`func (o *MdmCommand) GetProfileId() int64`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *MdmCommand) GetProfileIdOk() (*int64, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *MdmCommand) SetProfileId(v int64)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *MdmCommand) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


