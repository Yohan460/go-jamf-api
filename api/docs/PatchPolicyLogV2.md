# PatchPolicyLogV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PatchPolicyId** | Pointer to **string** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**StatusEnum** | Pointer to **string** |  | [optional] 
**AttemptNumber** | Pointer to **int32** |  | [optional] 
**IgnoredForPatchPolicyId** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchPolicyLogV2

`func NewPatchPolicyLogV2() *PatchPolicyLogV2`

NewPatchPolicyLogV2 instantiates a new PatchPolicyLogV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPolicyLogV2WithDefaults

`func NewPatchPolicyLogV2WithDefaults() *PatchPolicyLogV2`

NewPatchPolicyLogV2WithDefaults instantiates a new PatchPolicyLogV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatchPolicyId

`func (o *PatchPolicyLogV2) GetPatchPolicyId() string`

GetPatchPolicyId returns the PatchPolicyId field if non-nil, zero value otherwise.

### GetPatchPolicyIdOk

`func (o *PatchPolicyLogV2) GetPatchPolicyIdOk() (*string, bool)`

GetPatchPolicyIdOk returns a tuple with the PatchPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPolicyId

`func (o *PatchPolicyLogV2) SetPatchPolicyId(v string)`

SetPatchPolicyId sets PatchPolicyId field to given value.

### HasPatchPolicyId

`func (o *PatchPolicyLogV2) HasPatchPolicyId() bool`

HasPatchPolicyId returns a boolean if a field has been set.

### GetDeviceName

`func (o *PatchPolicyLogV2) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *PatchPolicyLogV2) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *PatchPolicyLogV2) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *PatchPolicyLogV2) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceId

`func (o *PatchPolicyLogV2) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *PatchPolicyLogV2) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *PatchPolicyLogV2) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *PatchPolicyLogV2) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetStatusCode

`func (o *PatchPolicyLogV2) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *PatchPolicyLogV2) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *PatchPolicyLogV2) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *PatchPolicyLogV2) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetStatusDate

`func (o *PatchPolicyLogV2) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *PatchPolicyLogV2) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *PatchPolicyLogV2) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *PatchPolicyLogV2) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusEnum

`func (o *PatchPolicyLogV2) GetStatusEnum() string`

GetStatusEnum returns the StatusEnum field if non-nil, zero value otherwise.

### GetStatusEnumOk

`func (o *PatchPolicyLogV2) GetStatusEnumOk() (*string, bool)`

GetStatusEnumOk returns a tuple with the StatusEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEnum

`func (o *PatchPolicyLogV2) SetStatusEnum(v string)`

SetStatusEnum sets StatusEnum field to given value.

### HasStatusEnum

`func (o *PatchPolicyLogV2) HasStatusEnum() bool`

HasStatusEnum returns a boolean if a field has been set.

### GetAttemptNumber

`func (o *PatchPolicyLogV2) GetAttemptNumber() int32`

GetAttemptNumber returns the AttemptNumber field if non-nil, zero value otherwise.

### GetAttemptNumberOk

`func (o *PatchPolicyLogV2) GetAttemptNumberOk() (*int32, bool)`

GetAttemptNumberOk returns a tuple with the AttemptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptNumber

`func (o *PatchPolicyLogV2) SetAttemptNumber(v int32)`

SetAttemptNumber sets AttemptNumber field to given value.

### HasAttemptNumber

`func (o *PatchPolicyLogV2) HasAttemptNumber() bool`

HasAttemptNumber returns a boolean if a field has been set.

### GetIgnoredForPatchPolicyId

`func (o *PatchPolicyLogV2) GetIgnoredForPatchPolicyId() string`

GetIgnoredForPatchPolicyId returns the IgnoredForPatchPolicyId field if non-nil, zero value otherwise.

### GetIgnoredForPatchPolicyIdOk

`func (o *PatchPolicyLogV2) GetIgnoredForPatchPolicyIdOk() (*string, bool)`

GetIgnoredForPatchPolicyIdOk returns a tuple with the IgnoredForPatchPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredForPatchPolicyId

`func (o *PatchPolicyLogV2) SetIgnoredForPatchPolicyId(v string)`

SetIgnoredForPatchPolicyId sets IgnoredForPatchPolicyId field to given value.

### HasIgnoredForPatchPolicyId

`func (o *PatchPolicyLogV2) HasIgnoredForPatchPolicyId() bool`

HasIgnoredForPatchPolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


