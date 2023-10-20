# MacOsManagedSoftwareUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceIds** | Pointer to **[]string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**MaxDeferrals** | Pointer to **int32** | Allow users to defer the update the provided number of times before macOS forces the update. If a value is provided, the Software Update will use the InstallLater install action. | [optional] 
**Version** | Pointer to **string** | If no value is provided, the version will default to latest version based on device eligibility. | [optional] 
**SkipVersionVerification** | Pointer to **bool** | If no value is provided, the skipVersionVerification will default to false. If a value is provided, the specified version will be forced to complete DownloadAndInstall install action. | [optional] [default to false]
**ApplyMajorUpdate** | Pointer to **bool** | ApplyMajorUpdate setting is available only when updating to the latest version based on device eligibility. If no value is provided, the calculated latest version will only include minor version updates. If a value is provided, the calculated latest version will include minor and major version updates. | [optional] [default to false]
**UpdateAction** | Pointer to **string** | MaxDeferral is ignored if using the DownloadOnly install action. | [optional] 
**ForceRestart** | Pointer to **bool** | If not set, forceRestart will default to false. Can only be true if using the DownloadAndInstall install action and the devices the command is sent to are on macOs 11 or higher. If true, the DownloadAndInstall action is performed, a restart will be forced. MaxDeferral will be ignored if defined.  | [optional] [default to false]
**Priority** | Pointer to **string** | Priority can only be configured on macOS 12.3 and above, for minor updates only. Any version below 12.3 is always Low and cannot be changed until prerequisites are met. When qualified, if not explicitly set, priority will default to High | [optional] [default to "HIGH"]

## Methods

### NewMacOsManagedSoftwareUpdate

`func NewMacOsManagedSoftwareUpdate() *MacOsManagedSoftwareUpdate`

NewMacOsManagedSoftwareUpdate instantiates a new MacOsManagedSoftwareUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacOsManagedSoftwareUpdateWithDefaults

`func NewMacOsManagedSoftwareUpdateWithDefaults() *MacOsManagedSoftwareUpdate`

NewMacOsManagedSoftwareUpdateWithDefaults instantiates a new MacOsManagedSoftwareUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceIds

`func (o *MacOsManagedSoftwareUpdate) GetDeviceIds() []string`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *MacOsManagedSoftwareUpdate) GetDeviceIdsOk() (*[]string, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *MacOsManagedSoftwareUpdate) SetDeviceIds(v []string)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *MacOsManagedSoftwareUpdate) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetGroupId

`func (o *MacOsManagedSoftwareUpdate) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *MacOsManagedSoftwareUpdate) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *MacOsManagedSoftwareUpdate) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *MacOsManagedSoftwareUpdate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetMaxDeferrals

`func (o *MacOsManagedSoftwareUpdate) GetMaxDeferrals() int32`

GetMaxDeferrals returns the MaxDeferrals field if non-nil, zero value otherwise.

### GetMaxDeferralsOk

`func (o *MacOsManagedSoftwareUpdate) GetMaxDeferralsOk() (*int32, bool)`

GetMaxDeferralsOk returns a tuple with the MaxDeferrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeferrals

`func (o *MacOsManagedSoftwareUpdate) SetMaxDeferrals(v int32)`

SetMaxDeferrals sets MaxDeferrals field to given value.

### HasMaxDeferrals

`func (o *MacOsManagedSoftwareUpdate) HasMaxDeferrals() bool`

HasMaxDeferrals returns a boolean if a field has been set.

### GetVersion

`func (o *MacOsManagedSoftwareUpdate) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MacOsManagedSoftwareUpdate) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MacOsManagedSoftwareUpdate) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MacOsManagedSoftwareUpdate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSkipVersionVerification

`func (o *MacOsManagedSoftwareUpdate) GetSkipVersionVerification() bool`

GetSkipVersionVerification returns the SkipVersionVerification field if non-nil, zero value otherwise.

### GetSkipVersionVerificationOk

`func (o *MacOsManagedSoftwareUpdate) GetSkipVersionVerificationOk() (*bool, bool)`

GetSkipVersionVerificationOk returns a tuple with the SkipVersionVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipVersionVerification

`func (o *MacOsManagedSoftwareUpdate) SetSkipVersionVerification(v bool)`

SetSkipVersionVerification sets SkipVersionVerification field to given value.

### HasSkipVersionVerification

`func (o *MacOsManagedSoftwareUpdate) HasSkipVersionVerification() bool`

HasSkipVersionVerification returns a boolean if a field has been set.

### GetApplyMajorUpdate

`func (o *MacOsManagedSoftwareUpdate) GetApplyMajorUpdate() bool`

GetApplyMajorUpdate returns the ApplyMajorUpdate field if non-nil, zero value otherwise.

### GetApplyMajorUpdateOk

`func (o *MacOsManagedSoftwareUpdate) GetApplyMajorUpdateOk() (*bool, bool)`

GetApplyMajorUpdateOk returns a tuple with the ApplyMajorUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyMajorUpdate

`func (o *MacOsManagedSoftwareUpdate) SetApplyMajorUpdate(v bool)`

SetApplyMajorUpdate sets ApplyMajorUpdate field to given value.

### HasApplyMajorUpdate

`func (o *MacOsManagedSoftwareUpdate) HasApplyMajorUpdate() bool`

HasApplyMajorUpdate returns a boolean if a field has been set.

### GetUpdateAction

`func (o *MacOsManagedSoftwareUpdate) GetUpdateAction() string`

GetUpdateAction returns the UpdateAction field if non-nil, zero value otherwise.

### GetUpdateActionOk

`func (o *MacOsManagedSoftwareUpdate) GetUpdateActionOk() (*string, bool)`

GetUpdateActionOk returns a tuple with the UpdateAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAction

`func (o *MacOsManagedSoftwareUpdate) SetUpdateAction(v string)`

SetUpdateAction sets UpdateAction field to given value.

### HasUpdateAction

`func (o *MacOsManagedSoftwareUpdate) HasUpdateAction() bool`

HasUpdateAction returns a boolean if a field has been set.

### GetForceRestart

`func (o *MacOsManagedSoftwareUpdate) GetForceRestart() bool`

GetForceRestart returns the ForceRestart field if non-nil, zero value otherwise.

### GetForceRestartOk

`func (o *MacOsManagedSoftwareUpdate) GetForceRestartOk() (*bool, bool)`

GetForceRestartOk returns a tuple with the ForceRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRestart

`func (o *MacOsManagedSoftwareUpdate) SetForceRestart(v bool)`

SetForceRestart sets ForceRestart field to given value.

### HasForceRestart

`func (o *MacOsManagedSoftwareUpdate) HasForceRestart() bool`

HasForceRestart returns a boolean if a field has been set.

### GetPriority

`func (o *MacOsManagedSoftwareUpdate) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MacOsManagedSoftwareUpdate) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MacOsManagedSoftwareUpdate) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MacOsManagedSoftwareUpdate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


