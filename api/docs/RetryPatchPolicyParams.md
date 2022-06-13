# RetryPatchPolicyParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PatchPolicyId** | Pointer to **int32** |  | [optional] 
**DeviceIds** | Pointer to **[]int32** |  | [optional] 
**IsRetryAllFailed** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewRetryPatchPolicyParams

`func NewRetryPatchPolicyParams() *RetryPatchPolicyParams`

NewRetryPatchPolicyParams instantiates a new RetryPatchPolicyParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryPatchPolicyParamsWithDefaults

`func NewRetryPatchPolicyParamsWithDefaults() *RetryPatchPolicyParams`

NewRetryPatchPolicyParamsWithDefaults instantiates a new RetryPatchPolicyParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatchPolicyId

`func (o *RetryPatchPolicyParams) GetPatchPolicyId() int32`

GetPatchPolicyId returns the PatchPolicyId field if non-nil, zero value otherwise.

### GetPatchPolicyIdOk

`func (o *RetryPatchPolicyParams) GetPatchPolicyIdOk() (*int32, bool)`

GetPatchPolicyIdOk returns a tuple with the PatchPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPolicyId

`func (o *RetryPatchPolicyParams) SetPatchPolicyId(v int32)`

SetPatchPolicyId sets PatchPolicyId field to given value.

### HasPatchPolicyId

`func (o *RetryPatchPolicyParams) HasPatchPolicyId() bool`

HasPatchPolicyId returns a boolean if a field has been set.

### GetDeviceIds

`func (o *RetryPatchPolicyParams) GetDeviceIds() []int32`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *RetryPatchPolicyParams) GetDeviceIdsOk() (*[]int32, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *RetryPatchPolicyParams) SetDeviceIds(v []int32)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *RetryPatchPolicyParams) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetIsRetryAllFailed

`func (o *RetryPatchPolicyParams) GetIsRetryAllFailed() bool`

GetIsRetryAllFailed returns the IsRetryAllFailed field if non-nil, zero value otherwise.

### GetIsRetryAllFailedOk

`func (o *RetryPatchPolicyParams) GetIsRetryAllFailedOk() (*bool, bool)`

GetIsRetryAllFailedOk returns a tuple with the IsRetryAllFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRetryAllFailed

`func (o *RetryPatchPolicyParams) SetIsRetryAllFailed(v bool)`

SetIsRetryAllFailed sets IsRetryAllFailed field to given value.

### HasIsRetryAllFailed

`func (o *RetryPatchPolicyParams) HasIsRetryAllFailed() bool`

HasIsRetryAllFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


