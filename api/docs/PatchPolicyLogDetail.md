# PatchPolicyLogDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AttemptNumber** | Pointer to **int64** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**Actions** | Pointer to [**[]PatchPolicyLogDetailAction**](PatchPolicyLogDetailAction.md) |  | [optional] 

## Methods

### NewPatchPolicyLogDetail

`func NewPatchPolicyLogDetail() *PatchPolicyLogDetail`

NewPatchPolicyLogDetail instantiates a new PatchPolicyLogDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPolicyLogDetailWithDefaults

`func NewPatchPolicyLogDetailWithDefaults() *PatchPolicyLogDetail`

NewPatchPolicyLogDetailWithDefaults instantiates a new PatchPolicyLogDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchPolicyLogDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchPolicyLogDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchPolicyLogDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchPolicyLogDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAttemptNumber

`func (o *PatchPolicyLogDetail) GetAttemptNumber() int64`

GetAttemptNumber returns the AttemptNumber field if non-nil, zero value otherwise.

### GetAttemptNumberOk

`func (o *PatchPolicyLogDetail) GetAttemptNumberOk() (*int64, bool)`

GetAttemptNumberOk returns a tuple with the AttemptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptNumber

`func (o *PatchPolicyLogDetail) SetAttemptNumber(v int64)`

SetAttemptNumber sets AttemptNumber field to given value.

### HasAttemptNumber

`func (o *PatchPolicyLogDetail) HasAttemptNumber() bool`

HasAttemptNumber returns a boolean if a field has been set.

### GetDeviceId

`func (o *PatchPolicyLogDetail) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *PatchPolicyLogDetail) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *PatchPolicyLogDetail) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *PatchPolicyLogDetail) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetActions

`func (o *PatchPolicyLogDetail) GetActions() []PatchPolicyLogDetailAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PatchPolicyLogDetail) GetActionsOk() (*[]PatchPolicyLogDetailAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PatchPolicyLogDetail) SetActions(v []PatchPolicyLogDetailAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PatchPolicyLogDetail) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


