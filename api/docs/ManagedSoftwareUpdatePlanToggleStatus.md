# ManagedSoftwareUpdatePlanToggleStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **NullableString** | The local server time when the toggle was initiated. Null if state is NEVER_RAN | [optional] 
**EndTime** | Pointer to **NullableString** | The local server time when the toggle was completed. Null if state is NEVER_RAN | [optional] 
**ElapsedTime** | Pointer to **NullableInt64** | Duration in seconds between the start time and end time. \&quot;Now\&quot; is used when end time is null. Null if state is NEVER_RAN | [optional] 
**State** | Pointer to **string** | The current state of the toggle | [optional] 
**TotalRecords** | Pointer to **int64** | The total number of records that will be deleted | [optional] 
**ProcessedRecords** | Pointer to **int64** | The total number of records that have been deleted | [optional] 
**PercentComplete** | Pointer to **float64** | The percentage between total and completed records. | [optional] 
**FormattedPercentComplete** | Pointer to **string** | Pretty print of total, processed, and percentage complete | [optional] 
**ExitState** | Pointer to **string** | Troubleshooting - The exit status code from the toggle processing job. \&quot;Unknown\&quot; will return when the toggle is running. | [optional] 
**ExitMessage** | Pointer to **string** | Troubleshooting - The exit message of the toggle job if it encounters an exception while running. Nominal return is an empty string | [optional] 

## Methods

### NewManagedSoftwareUpdatePlanToggleStatus

`func NewManagedSoftwareUpdatePlanToggleStatus() *ManagedSoftwareUpdatePlanToggleStatus`

NewManagedSoftwareUpdatePlanToggleStatus instantiates a new ManagedSoftwareUpdatePlanToggleStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedSoftwareUpdatePlanToggleStatusWithDefaults

`func NewManagedSoftwareUpdatePlanToggleStatusWithDefaults() *ManagedSoftwareUpdatePlanToggleStatus`

NewManagedSoftwareUpdatePlanToggleStatusWithDefaults instantiates a new ManagedSoftwareUpdatePlanToggleStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *ManagedSoftwareUpdatePlanToggleStatus) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ManagedSoftwareUpdatePlanToggleStatus) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ManagedSoftwareUpdatePlanToggleStatus) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ManagedSoftwareUpdatePlanToggleStatus) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *ManagedSoftwareUpdatePlanToggleStatus) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *ManagedSoftwareUpdatePlanToggleStatus) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *ManagedSoftwareUpdatePlanToggleStatus) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ManagedSoftwareUpdatePlanToggleStatus) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ManagedSoftwareUpdatePlanToggleStatus) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ManagedSoftwareUpdatePlanToggleStatus) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *ManagedSoftwareUpdatePlanToggleStatus) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *ManagedSoftwareUpdatePlanToggleStatus) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetElapsedTime

`func (o *ManagedSoftwareUpdatePlanToggleStatus) GetElapsedTime() int64`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *ManagedSoftwareUpdatePlanToggleStatus) GetElapsedTimeOk() (*int64, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *ManagedSoftwareUpdatePlanToggleStatus) SetElapsedTime(v int64)`

SetElapsedTime sets ElapsedTime field to given value.

### HasElapsedTime

`func (o *ManagedSoftwareUpdatePlanToggleStatus) HasElapsedTime() bool`

HasElapsedTime returns a boolean if a field has been set.

### SetElapsedTimeNil

`func (o *ManagedSoftwareUpdatePlanToggleStatus) SetElapsedTimeNil(b bool)`

 SetElapsedTimeNil sets the value for ElapsedTime to be an explicit nil

### UnsetElapsedTime
`func (o *ManagedSoftwareUpdatePlanToggleStatus) UnsetElapsedTime()`

UnsetElapsedTime ensures that no value is present for ElapsedTime, not even an explicit nil
### GetState

`func (o *ManagedSoftwareUpdatePlanToggleStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ManagedSoftwareUpdatePlanToggleStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ManagedSoftwareUpdatePlanToggleStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ManagedSoftwareUpdatePlanToggleStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTotalRecords

`func (o *ManagedSoftwareUpdatePlanToggleStatus) GetTotalRecords() int64`

GetTotalRecords returns the TotalRecords field if non-nil, zero value otherwise.

### GetTotalRecordsOk

`func (o *ManagedSoftwareUpdatePlanToggleStatus) GetTotalRecordsOk() (*int64, bool)`

GetTotalRecordsOk returns a tuple with the TotalRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecords

`func (o *ManagedSoftwareUpdatePlanToggleStatus) SetTotalRecords(v int64)`

SetTotalRecords sets TotalRecords field to given value.

### HasTotalRecords

`func (o *ManagedSoftwareUpdatePlanToggleStatus) HasTotalRecords() bool`

HasTotalRecords returns a boolean if a field has been set.

### GetProcessedRecords

`func (o *ManagedSoftwareUpdatePlanToggleStatus) GetProcessedRecords() int64`

GetProcessedRecords returns the ProcessedRecords field if non-nil, zero value otherwise.

### GetProcessedRecordsOk

`func (o *ManagedSoftwareUpdatePlanToggleStatus) GetProcessedRecordsOk() (*int64, bool)`

GetProcessedRecordsOk returns a tuple with the ProcessedRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedRecords

`func (o *ManagedSoftwareUpdatePlanToggleStatus) SetProcessedRecords(v int64)`

SetProcessedRecords sets ProcessedRecords field to given value.

### HasProcessedRecords

`func (o *ManagedSoftwareUpdatePlanToggleStatus) HasProcessedRecords() bool`

HasProcessedRecords returns a boolean if a field has been set.

### GetPercentComplete

`func (o *ManagedSoftwareUpdatePlanToggleStatus) GetPercentComplete() float64`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *ManagedSoftwareUpdatePlanToggleStatus) GetPercentCompleteOk() (*float64, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *ManagedSoftwareUpdatePlanToggleStatus) SetPercentComplete(v float64)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *ManagedSoftwareUpdatePlanToggleStatus) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### GetFormattedPercentComplete

`func (o *ManagedSoftwareUpdatePlanToggleStatus) GetFormattedPercentComplete() string`

GetFormattedPercentComplete returns the FormattedPercentComplete field if non-nil, zero value otherwise.

### GetFormattedPercentCompleteOk

`func (o *ManagedSoftwareUpdatePlanToggleStatus) GetFormattedPercentCompleteOk() (*string, bool)`

GetFormattedPercentCompleteOk returns a tuple with the FormattedPercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPercentComplete

`func (o *ManagedSoftwareUpdatePlanToggleStatus) SetFormattedPercentComplete(v string)`

SetFormattedPercentComplete sets FormattedPercentComplete field to given value.

### HasFormattedPercentComplete

`func (o *ManagedSoftwareUpdatePlanToggleStatus) HasFormattedPercentComplete() bool`

HasFormattedPercentComplete returns a boolean if a field has been set.

### GetExitState

`func (o *ManagedSoftwareUpdatePlanToggleStatus) GetExitState() string`

GetExitState returns the ExitState field if non-nil, zero value otherwise.

### GetExitStateOk

`func (o *ManagedSoftwareUpdatePlanToggleStatus) GetExitStateOk() (*string, bool)`

GetExitStateOk returns a tuple with the ExitState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitState

`func (o *ManagedSoftwareUpdatePlanToggleStatus) SetExitState(v string)`

SetExitState sets ExitState field to given value.

### HasExitState

`func (o *ManagedSoftwareUpdatePlanToggleStatus) HasExitState() bool`

HasExitState returns a boolean if a field has been set.

### GetExitMessage

`func (o *ManagedSoftwareUpdatePlanToggleStatus) GetExitMessage() string`

GetExitMessage returns the ExitMessage field if non-nil, zero value otherwise.

### GetExitMessageOk

`func (o *ManagedSoftwareUpdatePlanToggleStatus) GetExitMessageOk() (*string, bool)`

GetExitMessageOk returns a tuple with the ExitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitMessage

`func (o *ManagedSoftwareUpdatePlanToggleStatus) SetExitMessage(v string)`

SetExitMessage sets ExitMessage field to given value.

### HasExitMessage

`func (o *ManagedSoftwareUpdatePlanToggleStatus) HasExitMessage() bool`

HasExitMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


