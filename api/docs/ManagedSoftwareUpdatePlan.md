# ManagedSoftwareUpdatePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanUuid** | **string** |  | [readonly] 
**Device** | [**PlanDevice**](PlanDevice.md) |  | 
**UpdateAction** | **string** |  | [readonly] 
**VersionType** | **string** |  | [readonly] 
**SpecificVersion** | Pointer to **string** | Optional. Indicates the specific version to update to. Only available when the version type is set to specific version or custom version, otherwise defaults to NO_SPECIFIC_VERSION. | [optional] [readonly] [default to "NO_SPECIFIC_VERSION"]
**BuildVersion** | Pointer to **NullableString** | Optional. Indicates the build version to update to. Only available when the version type is set to custom version. | [optional] [readonly] 
**MaxDeferrals** | **int64** | Not applicable to all managed software update plans | [readonly] 
**ForceInstallLocalDateTime** | Pointer to **NullableString** | Optional. Indicates the local date and time of the device to force update by. | [optional] 
**RecipeId** | Pointer to **string** | The id of the recipe that was used to generate the plan. | [optional] [readonly] [default to "-1"]
**Status** | [**PlanStatus**](PlanStatus.md) |  | 

## Methods

### NewManagedSoftwareUpdatePlan

`func NewManagedSoftwareUpdatePlan(planUuid string, device PlanDevice, updateAction string, versionType string, maxDeferrals int64, status PlanStatus, ) *ManagedSoftwareUpdatePlan`

NewManagedSoftwareUpdatePlan instantiates a new ManagedSoftwareUpdatePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedSoftwareUpdatePlanWithDefaults

`func NewManagedSoftwareUpdatePlanWithDefaults() *ManagedSoftwareUpdatePlan`

NewManagedSoftwareUpdatePlanWithDefaults instantiates a new ManagedSoftwareUpdatePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanUuid

`func (o *ManagedSoftwareUpdatePlan) GetPlanUuid() string`

GetPlanUuid returns the PlanUuid field if non-nil, zero value otherwise.

### GetPlanUuidOk

`func (o *ManagedSoftwareUpdatePlan) GetPlanUuidOk() (*string, bool)`

GetPlanUuidOk returns a tuple with the PlanUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanUuid

`func (o *ManagedSoftwareUpdatePlan) SetPlanUuid(v string)`

SetPlanUuid sets PlanUuid field to given value.


### GetDevice

`func (o *ManagedSoftwareUpdatePlan) GetDevice() PlanDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ManagedSoftwareUpdatePlan) GetDeviceOk() (*PlanDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ManagedSoftwareUpdatePlan) SetDevice(v PlanDevice)`

SetDevice sets Device field to given value.


### GetUpdateAction

`func (o *ManagedSoftwareUpdatePlan) GetUpdateAction() string`

GetUpdateAction returns the UpdateAction field if non-nil, zero value otherwise.

### GetUpdateActionOk

`func (o *ManagedSoftwareUpdatePlan) GetUpdateActionOk() (*string, bool)`

GetUpdateActionOk returns a tuple with the UpdateAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAction

`func (o *ManagedSoftwareUpdatePlan) SetUpdateAction(v string)`

SetUpdateAction sets UpdateAction field to given value.


### GetVersionType

`func (o *ManagedSoftwareUpdatePlan) GetVersionType() string`

GetVersionType returns the VersionType field if non-nil, zero value otherwise.

### GetVersionTypeOk

`func (o *ManagedSoftwareUpdatePlan) GetVersionTypeOk() (*string, bool)`

GetVersionTypeOk returns a tuple with the VersionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionType

`func (o *ManagedSoftwareUpdatePlan) SetVersionType(v string)`

SetVersionType sets VersionType field to given value.


### GetSpecificVersion

`func (o *ManagedSoftwareUpdatePlan) GetSpecificVersion() string`

GetSpecificVersion returns the SpecificVersion field if non-nil, zero value otherwise.

### GetSpecificVersionOk

`func (o *ManagedSoftwareUpdatePlan) GetSpecificVersionOk() (*string, bool)`

GetSpecificVersionOk returns a tuple with the SpecificVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificVersion

`func (o *ManagedSoftwareUpdatePlan) SetSpecificVersion(v string)`

SetSpecificVersion sets SpecificVersion field to given value.

### HasSpecificVersion

`func (o *ManagedSoftwareUpdatePlan) HasSpecificVersion() bool`

HasSpecificVersion returns a boolean if a field has been set.

### GetBuildVersion

`func (o *ManagedSoftwareUpdatePlan) GetBuildVersion() string`

GetBuildVersion returns the BuildVersion field if non-nil, zero value otherwise.

### GetBuildVersionOk

`func (o *ManagedSoftwareUpdatePlan) GetBuildVersionOk() (*string, bool)`

GetBuildVersionOk returns a tuple with the BuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVersion

`func (o *ManagedSoftwareUpdatePlan) SetBuildVersion(v string)`

SetBuildVersion sets BuildVersion field to given value.

### HasBuildVersion

`func (o *ManagedSoftwareUpdatePlan) HasBuildVersion() bool`

HasBuildVersion returns a boolean if a field has been set.

### SetBuildVersionNil

`func (o *ManagedSoftwareUpdatePlan) SetBuildVersionNil(b bool)`

 SetBuildVersionNil sets the value for BuildVersion to be an explicit nil

### UnsetBuildVersion
`func (o *ManagedSoftwareUpdatePlan) UnsetBuildVersion()`

UnsetBuildVersion ensures that no value is present for BuildVersion, not even an explicit nil
### GetMaxDeferrals

`func (o *ManagedSoftwareUpdatePlan) GetMaxDeferrals() int64`

GetMaxDeferrals returns the MaxDeferrals field if non-nil, zero value otherwise.

### GetMaxDeferralsOk

`func (o *ManagedSoftwareUpdatePlan) GetMaxDeferralsOk() (*int64, bool)`

GetMaxDeferralsOk returns a tuple with the MaxDeferrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeferrals

`func (o *ManagedSoftwareUpdatePlan) SetMaxDeferrals(v int64)`

SetMaxDeferrals sets MaxDeferrals field to given value.


### GetForceInstallLocalDateTime

`func (o *ManagedSoftwareUpdatePlan) GetForceInstallLocalDateTime() string`

GetForceInstallLocalDateTime returns the ForceInstallLocalDateTime field if non-nil, zero value otherwise.

### GetForceInstallLocalDateTimeOk

`func (o *ManagedSoftwareUpdatePlan) GetForceInstallLocalDateTimeOk() (*string, bool)`

GetForceInstallLocalDateTimeOk returns a tuple with the ForceInstallLocalDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceInstallLocalDateTime

`func (o *ManagedSoftwareUpdatePlan) SetForceInstallLocalDateTime(v string)`

SetForceInstallLocalDateTime sets ForceInstallLocalDateTime field to given value.

### HasForceInstallLocalDateTime

`func (o *ManagedSoftwareUpdatePlan) HasForceInstallLocalDateTime() bool`

HasForceInstallLocalDateTime returns a boolean if a field has been set.

### SetForceInstallLocalDateTimeNil

`func (o *ManagedSoftwareUpdatePlan) SetForceInstallLocalDateTimeNil(b bool)`

 SetForceInstallLocalDateTimeNil sets the value for ForceInstallLocalDateTime to be an explicit nil

### UnsetForceInstallLocalDateTime
`func (o *ManagedSoftwareUpdatePlan) UnsetForceInstallLocalDateTime()`

UnsetForceInstallLocalDateTime ensures that no value is present for ForceInstallLocalDateTime, not even an explicit nil
### GetRecipeId

`func (o *ManagedSoftwareUpdatePlan) GetRecipeId() string`

GetRecipeId returns the RecipeId field if non-nil, zero value otherwise.

### GetRecipeIdOk

`func (o *ManagedSoftwareUpdatePlan) GetRecipeIdOk() (*string, bool)`

GetRecipeIdOk returns a tuple with the RecipeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipeId

`func (o *ManagedSoftwareUpdatePlan) SetRecipeId(v string)`

SetRecipeId sets RecipeId field to given value.

### HasRecipeId

`func (o *ManagedSoftwareUpdatePlan) HasRecipeId() bool`

HasRecipeId returns a boolean if a field has been set.

### GetStatus

`func (o *ManagedSoftwareUpdatePlan) GetStatus() PlanStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManagedSoftwareUpdatePlan) GetStatusOk() (*PlanStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManagedSoftwareUpdatePlan) SetStatus(v PlanStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


