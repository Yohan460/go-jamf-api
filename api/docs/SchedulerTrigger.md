# SchedulerTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TriggerKey** | Pointer to **string** |  | [optional] 
**PreviousFireTime** | Pointer to **time.Time** |  | [optional] 
**NextFireTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSchedulerTrigger

`func NewSchedulerTrigger() *SchedulerTrigger`

NewSchedulerTrigger instantiates a new SchedulerTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerTriggerWithDefaults

`func NewSchedulerTriggerWithDefaults() *SchedulerTrigger`

NewSchedulerTriggerWithDefaults instantiates a new SchedulerTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggerKey

`func (o *SchedulerTrigger) GetTriggerKey() string`

GetTriggerKey returns the TriggerKey field if non-nil, zero value otherwise.

### GetTriggerKeyOk

`func (o *SchedulerTrigger) GetTriggerKeyOk() (*string, bool)`

GetTriggerKeyOk returns a tuple with the TriggerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerKey

`func (o *SchedulerTrigger) SetTriggerKey(v string)`

SetTriggerKey sets TriggerKey field to given value.

### HasTriggerKey

`func (o *SchedulerTrigger) HasTriggerKey() bool`

HasTriggerKey returns a boolean if a field has been set.

### GetPreviousFireTime

`func (o *SchedulerTrigger) GetPreviousFireTime() time.Time`

GetPreviousFireTime returns the PreviousFireTime field if non-nil, zero value otherwise.

### GetPreviousFireTimeOk

`func (o *SchedulerTrigger) GetPreviousFireTimeOk() (*time.Time, bool)`

GetPreviousFireTimeOk returns a tuple with the PreviousFireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousFireTime

`func (o *SchedulerTrigger) SetPreviousFireTime(v time.Time)`

SetPreviousFireTime sets PreviousFireTime field to given value.

### HasPreviousFireTime

`func (o *SchedulerTrigger) HasPreviousFireTime() bool`

HasPreviousFireTime returns a boolean if a field has been set.

### GetNextFireTime

`func (o *SchedulerTrigger) GetNextFireTime() time.Time`

GetNextFireTime returns the NextFireTime field if non-nil, zero value otherwise.

### GetNextFireTimeOk

`func (o *SchedulerTrigger) GetNextFireTimeOk() (*time.Time, bool)`

GetNextFireTimeOk returns a tuple with the NextFireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFireTime

`func (o *SchedulerTrigger) SetNextFireTime(v time.Time)`

SetNextFireTime sets NextFireTime field to given value.

### HasNextFireTime

`func (o *SchedulerTrigger) HasNextFireTime() bool`

HasNextFireTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


