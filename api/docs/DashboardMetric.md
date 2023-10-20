# DashboardMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | Usually a number associated with the tag; i.e. 23 Pending Computers | [optional] 
**Enabled** | Pointer to **bool** | Logical to decide whether metric should be enabled or disabled; i.e. Policy can be at Retrying-Disabled status | [optional] [default to true]
**Tag** | Pointer to **string** |  | [optional] 

## Methods

### NewDashboardMetric

`func NewDashboardMetric() *DashboardMetric`

NewDashboardMetric instantiates a new DashboardMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardMetricWithDefaults

`func NewDashboardMetricWithDefaults() *DashboardMetric`

NewDashboardMetricWithDefaults instantiates a new DashboardMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *DashboardMetric) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DashboardMetric) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DashboardMetric) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DashboardMetric) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetEnabled

`func (o *DashboardMetric) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DashboardMetric) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DashboardMetric) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DashboardMetric) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTag

`func (o *DashboardMetric) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DashboardMetric) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DashboardMetric) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *DashboardMetric) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


