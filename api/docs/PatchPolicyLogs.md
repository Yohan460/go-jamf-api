# PatchPolicyLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]PatchPolicyLogV2**](PatchPolicyLogV2.md) |  | [optional] 

## Methods

### NewPatchPolicyLogs

`func NewPatchPolicyLogs() *PatchPolicyLogs`

NewPatchPolicyLogs instantiates a new PatchPolicyLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPolicyLogsWithDefaults

`func NewPatchPolicyLogsWithDefaults() *PatchPolicyLogs`

NewPatchPolicyLogsWithDefaults instantiates a new PatchPolicyLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *PatchPolicyLogs) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PatchPolicyLogs) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PatchPolicyLogs) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PatchPolicyLogs) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetResults

`func (o *PatchPolicyLogs) GetResults() []PatchPolicyLogV2`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PatchPolicyLogs) GetResultsOk() (*[]PatchPolicyLogV2, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PatchPolicyLogs) SetResults(v []PatchPolicyLogV2)`

SetResults sets Results field to given value.

### HasResults

`func (o *PatchPolicyLogs) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


