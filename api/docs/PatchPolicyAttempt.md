# PatchPolicyAttempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AttemptNo** | Pointer to **int32** |  | [optional] 
**DeviceId** | Pointer to **int32** |  | [optional] 
**Actions** | Pointer to [**[]PatchPolicyAttemptAction**](PatchPolicyAttemptAction.md) |  | [optional] 

## Methods

### NewPatchPolicyAttempt

`func NewPatchPolicyAttempt() *PatchPolicyAttempt`

NewPatchPolicyAttempt instantiates a new PatchPolicyAttempt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPolicyAttemptWithDefaults

`func NewPatchPolicyAttemptWithDefaults() *PatchPolicyAttempt`

NewPatchPolicyAttemptWithDefaults instantiates a new PatchPolicyAttempt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchPolicyAttempt) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchPolicyAttempt) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchPolicyAttempt) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchPolicyAttempt) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAttemptNo

`func (o *PatchPolicyAttempt) GetAttemptNo() int32`

GetAttemptNo returns the AttemptNo field if non-nil, zero value otherwise.

### GetAttemptNoOk

`func (o *PatchPolicyAttempt) GetAttemptNoOk() (*int32, bool)`

GetAttemptNoOk returns a tuple with the AttemptNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptNo

`func (o *PatchPolicyAttempt) SetAttemptNo(v int32)`

SetAttemptNo sets AttemptNo field to given value.

### HasAttemptNo

`func (o *PatchPolicyAttempt) HasAttemptNo() bool`

HasAttemptNo returns a boolean if a field has been set.

### GetDeviceId

`func (o *PatchPolicyAttempt) GetDeviceId() int32`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *PatchPolicyAttempt) GetDeviceIdOk() (*int32, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *PatchPolicyAttempt) SetDeviceId(v int32)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *PatchPolicyAttempt) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetActions

`func (o *PatchPolicyAttempt) GetActions() []PatchPolicyAttemptAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PatchPolicyAttempt) GetActionsOk() (*[]PatchPolicyAttemptAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PatchPolicyAttempt) SetActions(v []PatchPolicyAttemptAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PatchPolicyAttempt) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


