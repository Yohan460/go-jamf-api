# PrestageScopeResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrestageId** | Pointer to **string** |  | [optional] 
**Assignments** | Pointer to [**[]PrestageScopeAssignmentV2**](PrestageScopeAssignmentV2.md) |  | [optional] 
**VersionLock** | Pointer to **int64** |  | [optional] 

## Methods

### NewPrestageScopeResponseV2

`func NewPrestageScopeResponseV2() *PrestageScopeResponseV2`

NewPrestageScopeResponseV2 instantiates a new PrestageScopeResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrestageScopeResponseV2WithDefaults

`func NewPrestageScopeResponseV2WithDefaults() *PrestageScopeResponseV2`

NewPrestageScopeResponseV2WithDefaults instantiates a new PrestageScopeResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrestageId

`func (o *PrestageScopeResponseV2) GetPrestageId() string`

GetPrestageId returns the PrestageId field if non-nil, zero value otherwise.

### GetPrestageIdOk

`func (o *PrestageScopeResponseV2) GetPrestageIdOk() (*string, bool)`

GetPrestageIdOk returns a tuple with the PrestageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrestageId

`func (o *PrestageScopeResponseV2) SetPrestageId(v string)`

SetPrestageId sets PrestageId field to given value.

### HasPrestageId

`func (o *PrestageScopeResponseV2) HasPrestageId() bool`

HasPrestageId returns a boolean if a field has been set.

### GetAssignments

`func (o *PrestageScopeResponseV2) GetAssignments() []PrestageScopeAssignmentV2`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *PrestageScopeResponseV2) GetAssignmentsOk() (*[]PrestageScopeAssignmentV2, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *PrestageScopeResponseV2) SetAssignments(v []PrestageScopeAssignmentV2)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *PrestageScopeResponseV2) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetVersionLock

`func (o *PrestageScopeResponseV2) GetVersionLock() int64`

GetVersionLock returns the VersionLock field if non-nil, zero value otherwise.

### GetVersionLockOk

`func (o *PrestageScopeResponseV2) GetVersionLockOk() (*int64, bool)`

GetVersionLockOk returns a tuple with the VersionLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionLock

`func (o *PrestageScopeResponseV2) SetVersionLock(v int64)`

SetVersionLock sets VersionLock field to given value.

### HasVersionLock

`func (o *PrestageScopeResponseV2) HasVersionLock() bool`

HasVersionLock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


