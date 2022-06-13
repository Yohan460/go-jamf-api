# PatchPolicyLogSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]PatchPolicyLog**](PatchPolicyLog.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewPatchPolicyLogSearchResults

`func NewPatchPolicyLogSearchResults() *PatchPolicyLogSearchResults`

NewPatchPolicyLogSearchResults instantiates a new PatchPolicyLogSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPolicyLogSearchResultsWithDefaults

`func NewPatchPolicyLogSearchResultsWithDefaults() *PatchPolicyLogSearchResults`

NewPatchPolicyLogSearchResultsWithDefaults instantiates a new PatchPolicyLogSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *PatchPolicyLogSearchResults) GetResults() []PatchPolicyLog`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PatchPolicyLogSearchResults) GetResultsOk() (*[]PatchPolicyLog, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PatchPolicyLogSearchResults) SetResults(v []PatchPolicyLog)`

SetResults sets Results field to given value.

### HasResults

`func (o *PatchPolicyLogSearchResults) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalCount

`func (o *PatchPolicyLogSearchResults) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PatchPolicyLogSearchResults) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PatchPolicyLogSearchResults) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PatchPolicyLogSearchResults) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


