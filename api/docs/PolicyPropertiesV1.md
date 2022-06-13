# PolicyPropertiesV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoliciesRequireNetworkStateChange** | Pointer to **bool** | This field always returns false. | [optional] [default to false]
**AllowNetworkStateChangeTriggers** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewPolicyPropertiesV1

`func NewPolicyPropertiesV1() *PolicyPropertiesV1`

NewPolicyPropertiesV1 instantiates a new PolicyPropertiesV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyPropertiesV1WithDefaults

`func NewPolicyPropertiesV1WithDefaults() *PolicyPropertiesV1`

NewPolicyPropertiesV1WithDefaults instantiates a new PolicyPropertiesV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoliciesRequireNetworkStateChange

`func (o *PolicyPropertiesV1) GetPoliciesRequireNetworkStateChange() bool`

GetPoliciesRequireNetworkStateChange returns the PoliciesRequireNetworkStateChange field if non-nil, zero value otherwise.

### GetPoliciesRequireNetworkStateChangeOk

`func (o *PolicyPropertiesV1) GetPoliciesRequireNetworkStateChangeOk() (*bool, bool)`

GetPoliciesRequireNetworkStateChangeOk returns a tuple with the PoliciesRequireNetworkStateChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoliciesRequireNetworkStateChange

`func (o *PolicyPropertiesV1) SetPoliciesRequireNetworkStateChange(v bool)`

SetPoliciesRequireNetworkStateChange sets PoliciesRequireNetworkStateChange field to given value.

### HasPoliciesRequireNetworkStateChange

`func (o *PolicyPropertiesV1) HasPoliciesRequireNetworkStateChange() bool`

HasPoliciesRequireNetworkStateChange returns a boolean if a field has been set.

### GetAllowNetworkStateChangeTriggers

`func (o *PolicyPropertiesV1) GetAllowNetworkStateChangeTriggers() bool`

GetAllowNetworkStateChangeTriggers returns the AllowNetworkStateChangeTriggers field if non-nil, zero value otherwise.

### GetAllowNetworkStateChangeTriggersOk

`func (o *PolicyPropertiesV1) GetAllowNetworkStateChangeTriggersOk() (*bool, bool)`

GetAllowNetworkStateChangeTriggersOk returns a tuple with the AllowNetworkStateChangeTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNetworkStateChangeTriggers

`func (o *PolicyPropertiesV1) SetAllowNetworkStateChangeTriggers(v bool)`

SetAllowNetworkStateChangeTriggers sets AllowNetworkStateChangeTriggers field to given value.

### HasAllowNetworkStateChangeTriggers

`func (o *PolicyPropertiesV1) HasAllowNetworkStateChangeTriggers() bool`

HasAllowNetworkStateChangeTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


