# ClientCheckInV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckInFrequency** | Pointer to **int64** | Suggested values are 5, 15, 30, or 60. Web interface will not display correctly if not one of those. Minimim is 5, maximum is 60. | [optional] [default to 15]
**CreateHooks** | Pointer to **bool** |  | [optional] [default to false]
**HookLog** | Pointer to **bool** |  | [optional] [default to false]
**HookPolicies** | Pointer to **bool** |  | [optional] [default to false]
**CreateStartupScript** | Pointer to **bool** |  | [optional] [default to false]
**StartupLog** | Pointer to **bool** |  | [optional] [default to false]
**StartupPolicies** | Pointer to **bool** |  | [optional] [default to false]
**StartupSsh** | Pointer to **bool** |  | [optional] [default to false]
**EnableLocalConfigurationProfiles** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewClientCheckInV3

`func NewClientCheckInV3() *ClientCheckInV3`

NewClientCheckInV3 instantiates a new ClientCheckInV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCheckInV3WithDefaults

`func NewClientCheckInV3WithDefaults() *ClientCheckInV3`

NewClientCheckInV3WithDefaults instantiates a new ClientCheckInV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckInFrequency

`func (o *ClientCheckInV3) GetCheckInFrequency() int64`

GetCheckInFrequency returns the CheckInFrequency field if non-nil, zero value otherwise.

### GetCheckInFrequencyOk

`func (o *ClientCheckInV3) GetCheckInFrequencyOk() (*int64, bool)`

GetCheckInFrequencyOk returns a tuple with the CheckInFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInFrequency

`func (o *ClientCheckInV3) SetCheckInFrequency(v int64)`

SetCheckInFrequency sets CheckInFrequency field to given value.

### HasCheckInFrequency

`func (o *ClientCheckInV3) HasCheckInFrequency() bool`

HasCheckInFrequency returns a boolean if a field has been set.

### GetCreateHooks

`func (o *ClientCheckInV3) GetCreateHooks() bool`

GetCreateHooks returns the CreateHooks field if non-nil, zero value otherwise.

### GetCreateHooksOk

`func (o *ClientCheckInV3) GetCreateHooksOk() (*bool, bool)`

GetCreateHooksOk returns a tuple with the CreateHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateHooks

`func (o *ClientCheckInV3) SetCreateHooks(v bool)`

SetCreateHooks sets CreateHooks field to given value.

### HasCreateHooks

`func (o *ClientCheckInV3) HasCreateHooks() bool`

HasCreateHooks returns a boolean if a field has been set.

### GetHookLog

`func (o *ClientCheckInV3) GetHookLog() bool`

GetHookLog returns the HookLog field if non-nil, zero value otherwise.

### GetHookLogOk

`func (o *ClientCheckInV3) GetHookLogOk() (*bool, bool)`

GetHookLogOk returns a tuple with the HookLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookLog

`func (o *ClientCheckInV3) SetHookLog(v bool)`

SetHookLog sets HookLog field to given value.

### HasHookLog

`func (o *ClientCheckInV3) HasHookLog() bool`

HasHookLog returns a boolean if a field has been set.

### GetHookPolicies

`func (o *ClientCheckInV3) GetHookPolicies() bool`

GetHookPolicies returns the HookPolicies field if non-nil, zero value otherwise.

### GetHookPoliciesOk

`func (o *ClientCheckInV3) GetHookPoliciesOk() (*bool, bool)`

GetHookPoliciesOk returns a tuple with the HookPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookPolicies

`func (o *ClientCheckInV3) SetHookPolicies(v bool)`

SetHookPolicies sets HookPolicies field to given value.

### HasHookPolicies

`func (o *ClientCheckInV3) HasHookPolicies() bool`

HasHookPolicies returns a boolean if a field has been set.

### GetCreateStartupScript

`func (o *ClientCheckInV3) GetCreateStartupScript() bool`

GetCreateStartupScript returns the CreateStartupScript field if non-nil, zero value otherwise.

### GetCreateStartupScriptOk

`func (o *ClientCheckInV3) GetCreateStartupScriptOk() (*bool, bool)`

GetCreateStartupScriptOk returns a tuple with the CreateStartupScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateStartupScript

`func (o *ClientCheckInV3) SetCreateStartupScript(v bool)`

SetCreateStartupScript sets CreateStartupScript field to given value.

### HasCreateStartupScript

`func (o *ClientCheckInV3) HasCreateStartupScript() bool`

HasCreateStartupScript returns a boolean if a field has been set.

### GetStartupLog

`func (o *ClientCheckInV3) GetStartupLog() bool`

GetStartupLog returns the StartupLog field if non-nil, zero value otherwise.

### GetStartupLogOk

`func (o *ClientCheckInV3) GetStartupLogOk() (*bool, bool)`

GetStartupLogOk returns a tuple with the StartupLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupLog

`func (o *ClientCheckInV3) SetStartupLog(v bool)`

SetStartupLog sets StartupLog field to given value.

### HasStartupLog

`func (o *ClientCheckInV3) HasStartupLog() bool`

HasStartupLog returns a boolean if a field has been set.

### GetStartupPolicies

`func (o *ClientCheckInV3) GetStartupPolicies() bool`

GetStartupPolicies returns the StartupPolicies field if non-nil, zero value otherwise.

### GetStartupPoliciesOk

`func (o *ClientCheckInV3) GetStartupPoliciesOk() (*bool, bool)`

GetStartupPoliciesOk returns a tuple with the StartupPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupPolicies

`func (o *ClientCheckInV3) SetStartupPolicies(v bool)`

SetStartupPolicies sets StartupPolicies field to given value.

### HasStartupPolicies

`func (o *ClientCheckInV3) HasStartupPolicies() bool`

HasStartupPolicies returns a boolean if a field has been set.

### GetStartupSsh

`func (o *ClientCheckInV3) GetStartupSsh() bool`

GetStartupSsh returns the StartupSsh field if non-nil, zero value otherwise.

### GetStartupSshOk

`func (o *ClientCheckInV3) GetStartupSshOk() (*bool, bool)`

GetStartupSshOk returns a tuple with the StartupSsh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupSsh

`func (o *ClientCheckInV3) SetStartupSsh(v bool)`

SetStartupSsh sets StartupSsh field to given value.

### HasStartupSsh

`func (o *ClientCheckInV3) HasStartupSsh() bool`

HasStartupSsh returns a boolean if a field has been set.

### GetEnableLocalConfigurationProfiles

`func (o *ClientCheckInV3) GetEnableLocalConfigurationProfiles() bool`

GetEnableLocalConfigurationProfiles returns the EnableLocalConfigurationProfiles field if non-nil, zero value otherwise.

### GetEnableLocalConfigurationProfilesOk

`func (o *ClientCheckInV3) GetEnableLocalConfigurationProfilesOk() (*bool, bool)`

GetEnableLocalConfigurationProfilesOk returns a tuple with the EnableLocalConfigurationProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLocalConfigurationProfiles

`func (o *ClientCheckInV3) SetEnableLocalConfigurationProfiles(v bool)`

SetEnableLocalConfigurationProfiles sets EnableLocalConfigurationProfiles field to given value.

### HasEnableLocalConfigurationProfiles

`func (o *ClientCheckInV3) HasEnableLocalConfigurationProfiles() bool`

HasEnableLocalConfigurationProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


