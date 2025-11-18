# LogFlushingV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetentionPolicies** | Pointer to [**[]RetentionPolicyV1**](RetentionPolicyV1.md) |  | [optional] 
**HourOfDay** | Pointer to **int64** |  | [optional] [default to 0]

## Methods

### NewLogFlushingV1

`func NewLogFlushingV1() *LogFlushingV1`

NewLogFlushingV1 instantiates a new LogFlushingV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogFlushingV1WithDefaults

`func NewLogFlushingV1WithDefaults() *LogFlushingV1`

NewLogFlushingV1WithDefaults instantiates a new LogFlushingV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetentionPolicies

`func (o *LogFlushingV1) GetRetentionPolicies() []RetentionPolicyV1`

GetRetentionPolicies returns the RetentionPolicies field if non-nil, zero value otherwise.

### GetRetentionPoliciesOk

`func (o *LogFlushingV1) GetRetentionPoliciesOk() (*[]RetentionPolicyV1, bool)`

GetRetentionPoliciesOk returns a tuple with the RetentionPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicies

`func (o *LogFlushingV1) SetRetentionPolicies(v []RetentionPolicyV1)`

SetRetentionPolicies sets RetentionPolicies field to given value.

### HasRetentionPolicies

`func (o *LogFlushingV1) HasRetentionPolicies() bool`

HasRetentionPolicies returns a boolean if a field has been set.

### GetHourOfDay

`func (o *LogFlushingV1) GetHourOfDay() int64`

GetHourOfDay returns the HourOfDay field if non-nil, zero value otherwise.

### GetHourOfDayOk

`func (o *LogFlushingV1) GetHourOfDayOk() (*int64, bool)`

GetHourOfDayOk returns a tuple with the HourOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourOfDay

`func (o *LogFlushingV1) SetHourOfDay(v int64)`

SetHourOfDay sets HourOfDay field to given value.

### HasHourOfDay

`func (o *LogFlushingV1) HasHourOfDay() bool`

HasHourOfDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


