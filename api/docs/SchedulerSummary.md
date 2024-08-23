# SchedulerSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfPendingJobs** | Pointer to **int64** |  | [optional] 
**NumberOfExecutingJobs** | Pointer to **int64** |  | [optional] 
**NumberOfExecutedJobs** | Pointer to **int64** |  | [optional] 
**Started** | Pointer to **bool** |  | [optional] 

## Methods

### NewSchedulerSummary

`func NewSchedulerSummary() *SchedulerSummary`

NewSchedulerSummary instantiates a new SchedulerSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerSummaryWithDefaults

`func NewSchedulerSummaryWithDefaults() *SchedulerSummary`

NewSchedulerSummaryWithDefaults instantiates a new SchedulerSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfPendingJobs

`func (o *SchedulerSummary) GetNumberOfPendingJobs() int64`

GetNumberOfPendingJobs returns the NumberOfPendingJobs field if non-nil, zero value otherwise.

### GetNumberOfPendingJobsOk

`func (o *SchedulerSummary) GetNumberOfPendingJobsOk() (*int64, bool)`

GetNumberOfPendingJobsOk returns a tuple with the NumberOfPendingJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPendingJobs

`func (o *SchedulerSummary) SetNumberOfPendingJobs(v int64)`

SetNumberOfPendingJobs sets NumberOfPendingJobs field to given value.

### HasNumberOfPendingJobs

`func (o *SchedulerSummary) HasNumberOfPendingJobs() bool`

HasNumberOfPendingJobs returns a boolean if a field has been set.

### GetNumberOfExecutingJobs

`func (o *SchedulerSummary) GetNumberOfExecutingJobs() int64`

GetNumberOfExecutingJobs returns the NumberOfExecutingJobs field if non-nil, zero value otherwise.

### GetNumberOfExecutingJobsOk

`func (o *SchedulerSummary) GetNumberOfExecutingJobsOk() (*int64, bool)`

GetNumberOfExecutingJobsOk returns a tuple with the NumberOfExecutingJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfExecutingJobs

`func (o *SchedulerSummary) SetNumberOfExecutingJobs(v int64)`

SetNumberOfExecutingJobs sets NumberOfExecutingJobs field to given value.

### HasNumberOfExecutingJobs

`func (o *SchedulerSummary) HasNumberOfExecutingJobs() bool`

HasNumberOfExecutingJobs returns a boolean if a field has been set.

### GetNumberOfExecutedJobs

`func (o *SchedulerSummary) GetNumberOfExecutedJobs() int64`

GetNumberOfExecutedJobs returns the NumberOfExecutedJobs field if non-nil, zero value otherwise.

### GetNumberOfExecutedJobsOk

`func (o *SchedulerSummary) GetNumberOfExecutedJobsOk() (*int64, bool)`

GetNumberOfExecutedJobsOk returns a tuple with the NumberOfExecutedJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfExecutedJobs

`func (o *SchedulerSummary) SetNumberOfExecutedJobs(v int64)`

SetNumberOfExecutedJobs sets NumberOfExecutedJobs field to given value.

### HasNumberOfExecutedJobs

`func (o *SchedulerSummary) HasNumberOfExecutedJobs() bool`

HasNumberOfExecutedJobs returns a boolean if a field has been set.

### GetStarted

`func (o *SchedulerSummary) GetStarted() bool`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *SchedulerSummary) GetStartedOk() (*bool, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *SchedulerSummary) SetStarted(v bool)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *SchedulerSummary) HasStarted() bool`

HasStarted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


