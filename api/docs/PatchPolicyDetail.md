# PatchPolicyDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**TargetPatchVersion** | Pointer to **string** |  | [optional] 
**DeploymentMethod** | Pointer to **string** |  | [optional] 
**SoftwareTitleId** | Pointer to **string** |  | [optional] 
**SoftwareTitleConfigurationId** | Pointer to **string** |  | [optional] 
**KillAppsDelayMinutes** | Pointer to **int32** |  | [optional] 
**KillAppsMessage** | Pointer to **string** |  | [optional] 
**Downgrade** | Pointer to **bool** |  | [optional] 
**PatchUnknownVersion** | Pointer to **bool** |  | [optional] 
**NotificationHeader** | Pointer to **string** |  | [optional] 
**SelfServiceEnforceDeadline** | Pointer to **bool** |  | [optional] 
**SelfServiceDeadline** | Pointer to **int32** |  | [optional] 
**InstallButtonText** | Pointer to **string** |  | [optional] 
**SelfServiceDescription** | Pointer to **string** |  | [optional] 
**IconId** | Pointer to **string** |  | [optional] 
**ReminderFrequency** | Pointer to **int32** |  | [optional] 
**ReminderEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchPolicyDetail

`func NewPatchPolicyDetail() *PatchPolicyDetail`

NewPatchPolicyDetail instantiates a new PatchPolicyDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPolicyDetailWithDefaults

`func NewPatchPolicyDetailWithDefaults() *PatchPolicyDetail`

NewPatchPolicyDetailWithDefaults instantiates a new PatchPolicyDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchPolicyDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchPolicyDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchPolicyDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchPolicyDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchPolicyDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchPolicyDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchPolicyDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchPolicyDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchPolicyDetail) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchPolicyDetail) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchPolicyDetail) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchPolicyDetail) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTargetPatchVersion

`func (o *PatchPolicyDetail) GetTargetPatchVersion() string`

GetTargetPatchVersion returns the TargetPatchVersion field if non-nil, zero value otherwise.

### GetTargetPatchVersionOk

`func (o *PatchPolicyDetail) GetTargetPatchVersionOk() (*string, bool)`

GetTargetPatchVersionOk returns a tuple with the TargetPatchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPatchVersion

`func (o *PatchPolicyDetail) SetTargetPatchVersion(v string)`

SetTargetPatchVersion sets TargetPatchVersion field to given value.

### HasTargetPatchVersion

`func (o *PatchPolicyDetail) HasTargetPatchVersion() bool`

HasTargetPatchVersion returns a boolean if a field has been set.

### GetDeploymentMethod

`func (o *PatchPolicyDetail) GetDeploymentMethod() string`

GetDeploymentMethod returns the DeploymentMethod field if non-nil, zero value otherwise.

### GetDeploymentMethodOk

`func (o *PatchPolicyDetail) GetDeploymentMethodOk() (*string, bool)`

GetDeploymentMethodOk returns a tuple with the DeploymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentMethod

`func (o *PatchPolicyDetail) SetDeploymentMethod(v string)`

SetDeploymentMethod sets DeploymentMethod field to given value.

### HasDeploymentMethod

`func (o *PatchPolicyDetail) HasDeploymentMethod() bool`

HasDeploymentMethod returns a boolean if a field has been set.

### GetSoftwareTitleId

`func (o *PatchPolicyDetail) GetSoftwareTitleId() string`

GetSoftwareTitleId returns the SoftwareTitleId field if non-nil, zero value otherwise.

### GetSoftwareTitleIdOk

`func (o *PatchPolicyDetail) GetSoftwareTitleIdOk() (*string, bool)`

GetSoftwareTitleIdOk returns a tuple with the SoftwareTitleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTitleId

`func (o *PatchPolicyDetail) SetSoftwareTitleId(v string)`

SetSoftwareTitleId sets SoftwareTitleId field to given value.

### HasSoftwareTitleId

`func (o *PatchPolicyDetail) HasSoftwareTitleId() bool`

HasSoftwareTitleId returns a boolean if a field has been set.

### GetSoftwareTitleConfigurationId

`func (o *PatchPolicyDetail) GetSoftwareTitleConfigurationId() string`

GetSoftwareTitleConfigurationId returns the SoftwareTitleConfigurationId field if non-nil, zero value otherwise.

### GetSoftwareTitleConfigurationIdOk

`func (o *PatchPolicyDetail) GetSoftwareTitleConfigurationIdOk() (*string, bool)`

GetSoftwareTitleConfigurationIdOk returns a tuple with the SoftwareTitleConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTitleConfigurationId

`func (o *PatchPolicyDetail) SetSoftwareTitleConfigurationId(v string)`

SetSoftwareTitleConfigurationId sets SoftwareTitleConfigurationId field to given value.

### HasSoftwareTitleConfigurationId

`func (o *PatchPolicyDetail) HasSoftwareTitleConfigurationId() bool`

HasSoftwareTitleConfigurationId returns a boolean if a field has been set.

### GetKillAppsDelayMinutes

`func (o *PatchPolicyDetail) GetKillAppsDelayMinutes() int32`

GetKillAppsDelayMinutes returns the KillAppsDelayMinutes field if non-nil, zero value otherwise.

### GetKillAppsDelayMinutesOk

`func (o *PatchPolicyDetail) GetKillAppsDelayMinutesOk() (*int32, bool)`

GetKillAppsDelayMinutesOk returns a tuple with the KillAppsDelayMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillAppsDelayMinutes

`func (o *PatchPolicyDetail) SetKillAppsDelayMinutes(v int32)`

SetKillAppsDelayMinutes sets KillAppsDelayMinutes field to given value.

### HasKillAppsDelayMinutes

`func (o *PatchPolicyDetail) HasKillAppsDelayMinutes() bool`

HasKillAppsDelayMinutes returns a boolean if a field has been set.

### GetKillAppsMessage

`func (o *PatchPolicyDetail) GetKillAppsMessage() string`

GetKillAppsMessage returns the KillAppsMessage field if non-nil, zero value otherwise.

### GetKillAppsMessageOk

`func (o *PatchPolicyDetail) GetKillAppsMessageOk() (*string, bool)`

GetKillAppsMessageOk returns a tuple with the KillAppsMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillAppsMessage

`func (o *PatchPolicyDetail) SetKillAppsMessage(v string)`

SetKillAppsMessage sets KillAppsMessage field to given value.

### HasKillAppsMessage

`func (o *PatchPolicyDetail) HasKillAppsMessage() bool`

HasKillAppsMessage returns a boolean if a field has been set.

### GetDowngrade

`func (o *PatchPolicyDetail) GetDowngrade() bool`

GetDowngrade returns the Downgrade field if non-nil, zero value otherwise.

### GetDowngradeOk

`func (o *PatchPolicyDetail) GetDowngradeOk() (*bool, bool)`

GetDowngradeOk returns a tuple with the Downgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowngrade

`func (o *PatchPolicyDetail) SetDowngrade(v bool)`

SetDowngrade sets Downgrade field to given value.

### HasDowngrade

`func (o *PatchPolicyDetail) HasDowngrade() bool`

HasDowngrade returns a boolean if a field has been set.

### GetPatchUnknownVersion

`func (o *PatchPolicyDetail) GetPatchUnknownVersion() bool`

GetPatchUnknownVersion returns the PatchUnknownVersion field if non-nil, zero value otherwise.

### GetPatchUnknownVersionOk

`func (o *PatchPolicyDetail) GetPatchUnknownVersionOk() (*bool, bool)`

GetPatchUnknownVersionOk returns a tuple with the PatchUnknownVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchUnknownVersion

`func (o *PatchPolicyDetail) SetPatchUnknownVersion(v bool)`

SetPatchUnknownVersion sets PatchUnknownVersion field to given value.

### HasPatchUnknownVersion

`func (o *PatchPolicyDetail) HasPatchUnknownVersion() bool`

HasPatchUnknownVersion returns a boolean if a field has been set.

### GetNotificationHeader

`func (o *PatchPolicyDetail) GetNotificationHeader() string`

GetNotificationHeader returns the NotificationHeader field if non-nil, zero value otherwise.

### GetNotificationHeaderOk

`func (o *PatchPolicyDetail) GetNotificationHeaderOk() (*string, bool)`

GetNotificationHeaderOk returns a tuple with the NotificationHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationHeader

`func (o *PatchPolicyDetail) SetNotificationHeader(v string)`

SetNotificationHeader sets NotificationHeader field to given value.

### HasNotificationHeader

`func (o *PatchPolicyDetail) HasNotificationHeader() bool`

HasNotificationHeader returns a boolean if a field has been set.

### GetSelfServiceEnforceDeadline

`func (o *PatchPolicyDetail) GetSelfServiceEnforceDeadline() bool`

GetSelfServiceEnforceDeadline returns the SelfServiceEnforceDeadline field if non-nil, zero value otherwise.

### GetSelfServiceEnforceDeadlineOk

`func (o *PatchPolicyDetail) GetSelfServiceEnforceDeadlineOk() (*bool, bool)`

GetSelfServiceEnforceDeadlineOk returns a tuple with the SelfServiceEnforceDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServiceEnforceDeadline

`func (o *PatchPolicyDetail) SetSelfServiceEnforceDeadline(v bool)`

SetSelfServiceEnforceDeadline sets SelfServiceEnforceDeadline field to given value.

### HasSelfServiceEnforceDeadline

`func (o *PatchPolicyDetail) HasSelfServiceEnforceDeadline() bool`

HasSelfServiceEnforceDeadline returns a boolean if a field has been set.

### GetSelfServiceDeadline

`func (o *PatchPolicyDetail) GetSelfServiceDeadline() int32`

GetSelfServiceDeadline returns the SelfServiceDeadline field if non-nil, zero value otherwise.

### GetSelfServiceDeadlineOk

`func (o *PatchPolicyDetail) GetSelfServiceDeadlineOk() (*int32, bool)`

GetSelfServiceDeadlineOk returns a tuple with the SelfServiceDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServiceDeadline

`func (o *PatchPolicyDetail) SetSelfServiceDeadline(v int32)`

SetSelfServiceDeadline sets SelfServiceDeadline field to given value.

### HasSelfServiceDeadline

`func (o *PatchPolicyDetail) HasSelfServiceDeadline() bool`

HasSelfServiceDeadline returns a boolean if a field has been set.

### GetInstallButtonText

`func (o *PatchPolicyDetail) GetInstallButtonText() string`

GetInstallButtonText returns the InstallButtonText field if non-nil, zero value otherwise.

### GetInstallButtonTextOk

`func (o *PatchPolicyDetail) GetInstallButtonTextOk() (*string, bool)`

GetInstallButtonTextOk returns a tuple with the InstallButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallButtonText

`func (o *PatchPolicyDetail) SetInstallButtonText(v string)`

SetInstallButtonText sets InstallButtonText field to given value.

### HasInstallButtonText

`func (o *PatchPolicyDetail) HasInstallButtonText() bool`

HasInstallButtonText returns a boolean if a field has been set.

### GetSelfServiceDescription

`func (o *PatchPolicyDetail) GetSelfServiceDescription() string`

GetSelfServiceDescription returns the SelfServiceDescription field if non-nil, zero value otherwise.

### GetSelfServiceDescriptionOk

`func (o *PatchPolicyDetail) GetSelfServiceDescriptionOk() (*string, bool)`

GetSelfServiceDescriptionOk returns a tuple with the SelfServiceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServiceDescription

`func (o *PatchPolicyDetail) SetSelfServiceDescription(v string)`

SetSelfServiceDescription sets SelfServiceDescription field to given value.

### HasSelfServiceDescription

`func (o *PatchPolicyDetail) HasSelfServiceDescription() bool`

HasSelfServiceDescription returns a boolean if a field has been set.

### GetIconId

`func (o *PatchPolicyDetail) GetIconId() string`

GetIconId returns the IconId field if non-nil, zero value otherwise.

### GetIconIdOk

`func (o *PatchPolicyDetail) GetIconIdOk() (*string, bool)`

GetIconIdOk returns a tuple with the IconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconId

`func (o *PatchPolicyDetail) SetIconId(v string)`

SetIconId sets IconId field to given value.

### HasIconId

`func (o *PatchPolicyDetail) HasIconId() bool`

HasIconId returns a boolean if a field has been set.

### GetReminderFrequency

`func (o *PatchPolicyDetail) GetReminderFrequency() int32`

GetReminderFrequency returns the ReminderFrequency field if non-nil, zero value otherwise.

### GetReminderFrequencyOk

`func (o *PatchPolicyDetail) GetReminderFrequencyOk() (*int32, bool)`

GetReminderFrequencyOk returns a tuple with the ReminderFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderFrequency

`func (o *PatchPolicyDetail) SetReminderFrequency(v int32)`

SetReminderFrequency sets ReminderFrequency field to given value.

### HasReminderFrequency

`func (o *PatchPolicyDetail) HasReminderFrequency() bool`

HasReminderFrequency returns a boolean if a field has been set.

### GetReminderEnabled

`func (o *PatchPolicyDetail) GetReminderEnabled() bool`

GetReminderEnabled returns the ReminderEnabled field if non-nil, zero value otherwise.

### GetReminderEnabledOk

`func (o *PatchPolicyDetail) GetReminderEnabledOk() (*bool, bool)`

GetReminderEnabledOk returns a tuple with the ReminderEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderEnabled

`func (o *PatchPolicyDetail) SetReminderEnabled(v bool)`

SetReminderEnabled sets ReminderEnabled field to given value.

### HasReminderEnabled

`func (o *PatchPolicyDetail) HasReminderEnabled() bool`

HasReminderEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


