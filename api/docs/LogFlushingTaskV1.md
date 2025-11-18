# LogFlushingTaskV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the log flushing task | [optional] [readonly] 
**Qualifier** | **string** | The qualifier of the retention policy | 
**RetentionPeriod** | **int64** | The period beyond which data will be flushed | 
**RetentionPeriodUnit** | **string** | The unit of the retention period (eg: DAY, WEEK, MONTH, YEAR) | 
**State** | Pointer to **string** | The state of the task (eg: RUNNING, SUCCESS, FAILED, CANCELLED) | [optional] [readonly] 

## Methods

### NewLogFlushingTaskV1

`func NewLogFlushingTaskV1(qualifier string, retentionPeriod int64, retentionPeriodUnit string, ) *LogFlushingTaskV1`

NewLogFlushingTaskV1 instantiates a new LogFlushingTaskV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogFlushingTaskV1WithDefaults

`func NewLogFlushingTaskV1WithDefaults() *LogFlushingTaskV1`

NewLogFlushingTaskV1WithDefaults instantiates a new LogFlushingTaskV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogFlushingTaskV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogFlushingTaskV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogFlushingTaskV1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LogFlushingTaskV1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQualifier

`func (o *LogFlushingTaskV1) GetQualifier() string`

GetQualifier returns the Qualifier field if non-nil, zero value otherwise.

### GetQualifierOk

`func (o *LogFlushingTaskV1) GetQualifierOk() (*string, bool)`

GetQualifierOk returns a tuple with the Qualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifier

`func (o *LogFlushingTaskV1) SetQualifier(v string)`

SetQualifier sets Qualifier field to given value.


### GetRetentionPeriod

`func (o *LogFlushingTaskV1) GetRetentionPeriod() int64`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *LogFlushingTaskV1) GetRetentionPeriodOk() (*int64, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *LogFlushingTaskV1) SetRetentionPeriod(v int64)`

SetRetentionPeriod sets RetentionPeriod field to given value.


### GetRetentionPeriodUnit

`func (o *LogFlushingTaskV1) GetRetentionPeriodUnit() string`

GetRetentionPeriodUnit returns the RetentionPeriodUnit field if non-nil, zero value otherwise.

### GetRetentionPeriodUnitOk

`func (o *LogFlushingTaskV1) GetRetentionPeriodUnitOk() (*string, bool)`

GetRetentionPeriodUnitOk returns a tuple with the RetentionPeriodUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodUnit

`func (o *LogFlushingTaskV1) SetRetentionPeriodUnit(v string)`

SetRetentionPeriodUnit sets RetentionPeriodUnit field to given value.


### GetState

`func (o *LogFlushingTaskV1) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LogFlushingTaskV1) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LogFlushingTaskV1) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *LogFlushingTaskV1) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


