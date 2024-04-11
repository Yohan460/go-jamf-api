# PrestageScopeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrestageId** | Pointer to **int64** |  | [optional] 
**Assignments** | Pointer to [**[]PrestageScopeAssignment**](PrestageScopeAssignment.md) |  | [optional] 
**VersionLock** | Pointer to **int64** |  | [optional] 

## Methods

### NewPrestageScopeResponse

`func NewPrestageScopeResponse() *PrestageScopeResponse`

NewPrestageScopeResponse instantiates a new PrestageScopeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrestageScopeResponseWithDefaults

`func NewPrestageScopeResponseWithDefaults() *PrestageScopeResponse`

NewPrestageScopeResponseWithDefaults instantiates a new PrestageScopeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrestageId

`func (o *PrestageScopeResponse) GetPrestageId() int64`

GetPrestageId returns the PrestageId field if non-nil, zero value otherwise.

### GetPrestageIdOk

`func (o *PrestageScopeResponse) GetPrestageIdOk() (*int64, bool)`

GetPrestageIdOk returns a tuple with the PrestageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrestageId

`func (o *PrestageScopeResponse) SetPrestageId(v int64)`

SetPrestageId sets PrestageId field to given value.

### HasPrestageId

`func (o *PrestageScopeResponse) HasPrestageId() bool`

HasPrestageId returns a boolean if a field has been set.

### GetAssignments

`func (o *PrestageScopeResponse) GetAssignments() []PrestageScopeAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *PrestageScopeResponse) GetAssignmentsOk() (*[]PrestageScopeAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *PrestageScopeResponse) SetAssignments(v []PrestageScopeAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *PrestageScopeResponse) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetVersionLock

`func (o *PrestageScopeResponse) GetVersionLock() int64`

GetVersionLock returns the VersionLock field if non-nil, zero value otherwise.

### GetVersionLockOk

`func (o *PrestageScopeResponse) GetVersionLockOk() (*int64, bool)`

GetVersionLockOk returns a tuple with the VersionLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionLock

`func (o *PrestageScopeResponse) SetVersionLock(v int64)`

SetVersionLock sets VersionLock field to given value.

### HasVersionLock

`func (o *PrestageScopeResponse) HasVersionLock() bool`

HasVersionLock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


