# ClientCheckInV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckInFrequency** | Pointer to **int64** | Suggested values are 5, 15, 30, or 60. Web interface will not display correctly if not one of those. Minimim is 5, maximum is 60. | [optional] [default to 15]
**CreateHooks** | Pointer to **bool** |  | [optional] [default to false]
**HookLog** | Pointer to **bool** |  | [optional] [default to false]
**HookPolicies** | Pointer to **bool** |  | [optional] [default to false]
**HookHideRestore** | Pointer to **bool** |  | [optional] [default to false]
**HookMcx** | Pointer to **bool** |  | [optional] [default to false]
**BackgroundHooks** | Pointer to **bool** |  | [optional] [default to false]
**HookDisplayStatus** | Pointer to **bool** |  | [optional] [default to false]
**CreateStartupScript** | Pointer to **bool** |  | [optional] [default to false]
**StartupLog** | Pointer to **bool** |  | [optional] [default to false]
**StartupPolicies** | Pointer to **bool** |  | [optional] [default to false]
**StartupSsh** | Pointer to **bool** |  | [optional] [default to false]
**StartupMcx** | Pointer to **bool** |  | [optional] [default to false]
**EnableLocalConfigurationProfiles** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewClientCheckInV2

`func NewClientCheckInV2() *ClientCheckInV2`

NewClientCheckInV2 instantiates a new ClientCheckInV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCheckInV2WithDefaults

`func NewClientCheckInV2WithDefaults() *ClientCheckInV2`

NewClientCheckInV2WithDefaults instantiates a new ClientCheckInV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckInFrequency

`func (o *ClientCheckInV2) GetCheckInFrequency() int64`

GetCheckInFrequency returns the CheckInFrequency field if non-nil, zero value otherwise.

### GetCheckInFrequencyOk

`func (o *ClientCheckInV2) GetCheckInFrequencyOk() (*int64, bool)`

GetCheckInFrequencyOk returns a tuple with the CheckInFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInFrequency

`func (o *ClientCheckInV2) SetCheckInFrequency(v int64)`

SetCheckInFrequency sets CheckInFrequency field to given value.

### HasCheckInFrequency

`func (o *ClientCheckInV2) HasCheckInFrequency() bool`

HasCheckInFrequency returns a boolean if a field has been set.

### GetCreateHooks

`func (o *ClientCheckInV2) GetCreateHooks() bool`

GetCreateHooks returns the CreateHooks field if non-nil, zero value otherwise.

### GetCreateHooksOk

`func (o *ClientCheckInV2) GetCreateHooksOk() (*bool, bool)`

GetCreateHooksOk returns a tuple with the CreateHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateHooks

`func (o *ClientCheckInV2) SetCreateHooks(v bool)`

SetCreateHooks sets CreateHooks field to given value.

### HasCreateHooks

`func (o *ClientCheckInV2) HasCreateHooks() bool`

HasCreateHooks returns a boolean if a field has been set.

### GetHookLog

`func (o *ClientCheckInV2) GetHookLog() bool`

GetHookLog returns the HookLog field if non-nil, zero value otherwise.

### GetHookLogOk

`func (o *ClientCheckInV2) GetHookLogOk() (*bool, bool)`

GetHookLogOk returns a tuple with the HookLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookLog

`func (o *ClientCheckInV2) SetHookLog(v bool)`

SetHookLog sets HookLog field to given value.

### HasHookLog

`func (o *ClientCheckInV2) HasHookLog() bool`

HasHookLog returns a boolean if a field has been set.

### GetHookPolicies

`func (o *ClientCheckInV2) GetHookPolicies() bool`

GetHookPolicies returns the HookPolicies field if non-nil, zero value otherwise.

### GetHookPoliciesOk

`func (o *ClientCheckInV2) GetHookPoliciesOk() (*bool, bool)`

GetHookPoliciesOk returns a tuple with the HookPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookPolicies

`func (o *ClientCheckInV2) SetHookPolicies(v bool)`

SetHookPolicies sets HookPolicies field to given value.

### HasHookPolicies

`func (o *ClientCheckInV2) HasHookPolicies() bool`

HasHookPolicies returns a boolean if a field has been set.

### GetHookHideRestore

`func (o *ClientCheckInV2) GetHookHideRestore() bool`

GetHookHideRestore returns the HookHideRestore field if non-nil, zero value otherwise.

### GetHookHideRestoreOk

`func (o *ClientCheckInV2) GetHookHideRestoreOk() (*bool, bool)`

GetHookHideRestoreOk returns a tuple with the HookHideRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookHideRestore

`func (o *ClientCheckInV2) SetHookHideRestore(v bool)`

SetHookHideRestore sets HookHideRestore field to given value.

### HasHookHideRestore

`func (o *ClientCheckInV2) HasHookHideRestore() bool`

HasHookHideRestore returns a boolean if a field has been set.

### GetHookMcx

`func (o *ClientCheckInV2) GetHookMcx() bool`

GetHookMcx returns the HookMcx field if non-nil, zero value otherwise.

### GetHookMcxOk

`func (o *ClientCheckInV2) GetHookMcxOk() (*bool, bool)`

GetHookMcxOk returns a tuple with the HookMcx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookMcx

`func (o *ClientCheckInV2) SetHookMcx(v bool)`

SetHookMcx sets HookMcx field to given value.

### HasHookMcx

`func (o *ClientCheckInV2) HasHookMcx() bool`

HasHookMcx returns a boolean if a field has been set.

### GetBackgroundHooks

`func (o *ClientCheckInV2) GetBackgroundHooks() bool`

GetBackgroundHooks returns the BackgroundHooks field if non-nil, zero value otherwise.

### GetBackgroundHooksOk

`func (o *ClientCheckInV2) GetBackgroundHooksOk() (*bool, bool)`

GetBackgroundHooksOk returns a tuple with the BackgroundHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundHooks

`func (o *ClientCheckInV2) SetBackgroundHooks(v bool)`

SetBackgroundHooks sets BackgroundHooks field to given value.

### HasBackgroundHooks

`func (o *ClientCheckInV2) HasBackgroundHooks() bool`

HasBackgroundHooks returns a boolean if a field has been set.

### GetHookDisplayStatus

`func (o *ClientCheckInV2) GetHookDisplayStatus() bool`

GetHookDisplayStatus returns the HookDisplayStatus field if non-nil, zero value otherwise.

### GetHookDisplayStatusOk

`func (o *ClientCheckInV2) GetHookDisplayStatusOk() (*bool, bool)`

GetHookDisplayStatusOk returns a tuple with the HookDisplayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookDisplayStatus

`func (o *ClientCheckInV2) SetHookDisplayStatus(v bool)`

SetHookDisplayStatus sets HookDisplayStatus field to given value.

### HasHookDisplayStatus

`func (o *ClientCheckInV2) HasHookDisplayStatus() bool`

HasHookDisplayStatus returns a boolean if a field has been set.

### GetCreateStartupScript

`func (o *ClientCheckInV2) GetCreateStartupScript() bool`

GetCreateStartupScript returns the CreateStartupScript field if non-nil, zero value otherwise.

### GetCreateStartupScriptOk

`func (o *ClientCheckInV2) GetCreateStartupScriptOk() (*bool, bool)`

GetCreateStartupScriptOk returns a tuple with the CreateStartupScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateStartupScript

`func (o *ClientCheckInV2) SetCreateStartupScript(v bool)`

SetCreateStartupScript sets CreateStartupScript field to given value.

### HasCreateStartupScript

`func (o *ClientCheckInV2) HasCreateStartupScript() bool`

HasCreateStartupScript returns a boolean if a field has been set.

### GetStartupLog

`func (o *ClientCheckInV2) GetStartupLog() bool`

GetStartupLog returns the StartupLog field if non-nil, zero value otherwise.

### GetStartupLogOk

`func (o *ClientCheckInV2) GetStartupLogOk() (*bool, bool)`

GetStartupLogOk returns a tuple with the StartupLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupLog

`func (o *ClientCheckInV2) SetStartupLog(v bool)`

SetStartupLog sets StartupLog field to given value.

### HasStartupLog

`func (o *ClientCheckInV2) HasStartupLog() bool`

HasStartupLog returns a boolean if a field has been set.

### GetStartupPolicies

`func (o *ClientCheckInV2) GetStartupPolicies() bool`

GetStartupPolicies returns the StartupPolicies field if non-nil, zero value otherwise.

### GetStartupPoliciesOk

`func (o *ClientCheckInV2) GetStartupPoliciesOk() (*bool, bool)`

GetStartupPoliciesOk returns a tuple with the StartupPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupPolicies

`func (o *ClientCheckInV2) SetStartupPolicies(v bool)`

SetStartupPolicies sets StartupPolicies field to given value.

### HasStartupPolicies

`func (o *ClientCheckInV2) HasStartupPolicies() bool`

HasStartupPolicies returns a boolean if a field has been set.

### GetStartupSsh

`func (o *ClientCheckInV2) GetStartupSsh() bool`

GetStartupSsh returns the StartupSsh field if non-nil, zero value otherwise.

### GetStartupSshOk

`func (o *ClientCheckInV2) GetStartupSshOk() (*bool, bool)`

GetStartupSshOk returns a tuple with the StartupSsh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupSsh

`func (o *ClientCheckInV2) SetStartupSsh(v bool)`

SetStartupSsh sets StartupSsh field to given value.

### HasStartupSsh

`func (o *ClientCheckInV2) HasStartupSsh() bool`

HasStartupSsh returns a boolean if a field has been set.

### GetStartupMcx

`func (o *ClientCheckInV2) GetStartupMcx() bool`

GetStartupMcx returns the StartupMcx field if non-nil, zero value otherwise.

### GetStartupMcxOk

`func (o *ClientCheckInV2) GetStartupMcxOk() (*bool, bool)`

GetStartupMcxOk returns a tuple with the StartupMcx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupMcx

`func (o *ClientCheckInV2) SetStartupMcx(v bool)`

SetStartupMcx sets StartupMcx field to given value.

### HasStartupMcx

`func (o *ClientCheckInV2) HasStartupMcx() bool`

HasStartupMcx returns a boolean if a field has been set.

### GetEnableLocalConfigurationProfiles

`func (o *ClientCheckInV2) GetEnableLocalConfigurationProfiles() bool`

GetEnableLocalConfigurationProfiles returns the EnableLocalConfigurationProfiles field if non-nil, zero value otherwise.

### GetEnableLocalConfigurationProfilesOk

`func (o *ClientCheckInV2) GetEnableLocalConfigurationProfilesOk() (*bool, bool)`

GetEnableLocalConfigurationProfilesOk returns a tuple with the EnableLocalConfigurationProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLocalConfigurationProfiles

`func (o *ClientCheckInV2) SetEnableLocalConfigurationProfiles(v bool)`

SetEnableLocalConfigurationProfiles sets EnableLocalConfigurationProfiles field to given value.

### HasEnableLocalConfigurationProfiles

`func (o *ClientCheckInV2) HasEnableLocalConfigurationProfiles() bool`

HasEnableLocalConfigurationProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


