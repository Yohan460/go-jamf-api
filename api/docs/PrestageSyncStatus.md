# PrestageSyncStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyncState** | Pointer to **string** |  | [optional] 
**PrestageId** | Pointer to **int64** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 

## Methods

### NewPrestageSyncStatus

`func NewPrestageSyncStatus() *PrestageSyncStatus`

NewPrestageSyncStatus instantiates a new PrestageSyncStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrestageSyncStatusWithDefaults

`func NewPrestageSyncStatusWithDefaults() *PrestageSyncStatus`

NewPrestageSyncStatusWithDefaults instantiates a new PrestageSyncStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyncState

`func (o *PrestageSyncStatus) GetSyncState() string`

GetSyncState returns the SyncState field if non-nil, zero value otherwise.

### GetSyncStateOk

`func (o *PrestageSyncStatus) GetSyncStateOk() (*string, bool)`

GetSyncStateOk returns a tuple with the SyncState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncState

`func (o *PrestageSyncStatus) SetSyncState(v string)`

SetSyncState sets SyncState field to given value.

### HasSyncState

`func (o *PrestageSyncStatus) HasSyncState() bool`

HasSyncState returns a boolean if a field has been set.

### GetPrestageId

`func (o *PrestageSyncStatus) GetPrestageId() int64`

GetPrestageId returns the PrestageId field if non-nil, zero value otherwise.

### GetPrestageIdOk

`func (o *PrestageSyncStatus) GetPrestageIdOk() (*int64, bool)`

GetPrestageIdOk returns a tuple with the PrestageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrestageId

`func (o *PrestageSyncStatus) SetPrestageId(v int64)`

SetPrestageId sets PrestageId field to given value.

### HasPrestageId

`func (o *PrestageSyncStatus) HasPrestageId() bool`

HasPrestageId returns a boolean if a field has been set.

### GetTimestamp

`func (o *PrestageSyncStatus) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PrestageSyncStatus) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PrestageSyncStatus) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PrestageSyncStatus) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


