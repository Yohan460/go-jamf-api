# RetentionPolicyV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] [readonly] 
**Qualifier** | Pointer to **string** |  | [optional] [readonly] 
**RetentionPeriod** | Pointer to **int64** |  | [optional] 
**RetentionPeriodUnit** | Pointer to **string** | The unit of the retention period (eg: DAY, WEEK, MONTH, YEAR) | [optional] 

## Methods

### NewRetentionPolicyV1

`func NewRetentionPolicyV1() *RetentionPolicyV1`

NewRetentionPolicyV1 instantiates a new RetentionPolicyV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetentionPolicyV1WithDefaults

`func NewRetentionPolicyV1WithDefaults() *RetentionPolicyV1`

NewRetentionPolicyV1WithDefaults instantiates a new RetentionPolicyV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *RetentionPolicyV1) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RetentionPolicyV1) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RetentionPolicyV1) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RetentionPolicyV1) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetQualifier

`func (o *RetentionPolicyV1) GetQualifier() string`

GetQualifier returns the Qualifier field if non-nil, zero value otherwise.

### GetQualifierOk

`func (o *RetentionPolicyV1) GetQualifierOk() (*string, bool)`

GetQualifierOk returns a tuple with the Qualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifier

`func (o *RetentionPolicyV1) SetQualifier(v string)`

SetQualifier sets Qualifier field to given value.

### HasQualifier

`func (o *RetentionPolicyV1) HasQualifier() bool`

HasQualifier returns a boolean if a field has been set.

### GetRetentionPeriod

`func (o *RetentionPolicyV1) GetRetentionPeriod() int64`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *RetentionPolicyV1) GetRetentionPeriodOk() (*int64, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *RetentionPolicyV1) SetRetentionPeriod(v int64)`

SetRetentionPeriod sets RetentionPeriod field to given value.

### HasRetentionPeriod

`func (o *RetentionPolicyV1) HasRetentionPeriod() bool`

HasRetentionPeriod returns a boolean if a field has been set.

### GetRetentionPeriodUnit

`func (o *RetentionPolicyV1) GetRetentionPeriodUnit() string`

GetRetentionPeriodUnit returns the RetentionPeriodUnit field if non-nil, zero value otherwise.

### GetRetentionPeriodUnitOk

`func (o *RetentionPolicyV1) GetRetentionPeriodUnitOk() (*string, bool)`

GetRetentionPeriodUnitOk returns a tuple with the RetentionPeriodUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodUnit

`func (o *RetentionPolicyV1) SetRetentionPeriodUnit(v string)`

SetRetentionPeriodUnit sets RetentionPeriodUnit field to given value.

### HasRetentionPeriodUnit

`func (o *RetentionPolicyV1) HasRetentionPeriodUnit() bool`

HasRetentionPeriodUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


