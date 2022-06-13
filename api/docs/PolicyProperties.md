# PolicyProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPoliciesRequireNetworkStateChange** | Pointer to **bool** | This field always returns false. | [optional] [default to false]
**IsAllowNetworkStateChangeTriggers** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewPolicyProperties

`func NewPolicyProperties() *PolicyProperties`

NewPolicyProperties instantiates a new PolicyProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyPropertiesWithDefaults

`func NewPolicyPropertiesWithDefaults() *PolicyProperties`

NewPolicyPropertiesWithDefaults instantiates a new PolicyProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPoliciesRequireNetworkStateChange

`func (o *PolicyProperties) GetIsPoliciesRequireNetworkStateChange() bool`

GetIsPoliciesRequireNetworkStateChange returns the IsPoliciesRequireNetworkStateChange field if non-nil, zero value otherwise.

### GetIsPoliciesRequireNetworkStateChangeOk

`func (o *PolicyProperties) GetIsPoliciesRequireNetworkStateChangeOk() (*bool, bool)`

GetIsPoliciesRequireNetworkStateChangeOk returns a tuple with the IsPoliciesRequireNetworkStateChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPoliciesRequireNetworkStateChange

`func (o *PolicyProperties) SetIsPoliciesRequireNetworkStateChange(v bool)`

SetIsPoliciesRequireNetworkStateChange sets IsPoliciesRequireNetworkStateChange field to given value.

### HasIsPoliciesRequireNetworkStateChange

`func (o *PolicyProperties) HasIsPoliciesRequireNetworkStateChange() bool`

HasIsPoliciesRequireNetworkStateChange returns a boolean if a field has been set.

### GetIsAllowNetworkStateChangeTriggers

`func (o *PolicyProperties) GetIsAllowNetworkStateChangeTriggers() bool`

GetIsAllowNetworkStateChangeTriggers returns the IsAllowNetworkStateChangeTriggers field if non-nil, zero value otherwise.

### GetIsAllowNetworkStateChangeTriggersOk

`func (o *PolicyProperties) GetIsAllowNetworkStateChangeTriggersOk() (*bool, bool)`

GetIsAllowNetworkStateChangeTriggersOk returns a tuple with the IsAllowNetworkStateChangeTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowNetworkStateChangeTriggers

`func (o *PolicyProperties) SetIsAllowNetworkStateChangeTriggers(v bool)`

SetIsAllowNetworkStateChangeTriggers sets IsAllowNetworkStateChangeTriggers field to given value.

### HasIsAllowNetworkStateChangeTriggers

`func (o *PolicyProperties) HasIsAllowNetworkStateChangeTriggers() bool`

HasIsAllowNetworkStateChangeTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


