# PlanConfigurationPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateAction** | **string** |  | 
**VersionType** | **string** |  | 
**SpecificVersion** | Pointer to **string** | Optional. Indicates the specific version to update to. Only available when the version type is set to specific version or custom version, otherwise defaults to NO_SPECIFIC_VERSION. | [optional] [default to "NO_SPECIFIC_VERSION"]
**BuildVersion** | Pointer to **NullableString** | Optional. Indicates the build version to update to. Only available when the version type is set to custom version. | [optional] 
**MaxDeferrals** | Pointer to **int64** | Required when the provided updateAction is DOWNLOAD_INSTALL_ALLOW_DEFERRAL, not applicable to all managed software update plans | [optional] 
**ForceInstallLocalDateTime** | Pointer to **NullableString** | Optional. Indicates the local date and time of the device to force update by. | [optional] 

## Methods

### NewPlanConfigurationPost

`func NewPlanConfigurationPost(updateAction string, versionType string, ) *PlanConfigurationPost`

NewPlanConfigurationPost instantiates a new PlanConfigurationPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanConfigurationPostWithDefaults

`func NewPlanConfigurationPostWithDefaults() *PlanConfigurationPost`

NewPlanConfigurationPostWithDefaults instantiates a new PlanConfigurationPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateAction

`func (o *PlanConfigurationPost) GetUpdateAction() string`

GetUpdateAction returns the UpdateAction field if non-nil, zero value otherwise.

### GetUpdateActionOk

`func (o *PlanConfigurationPost) GetUpdateActionOk() (*string, bool)`

GetUpdateActionOk returns a tuple with the UpdateAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAction

`func (o *PlanConfigurationPost) SetUpdateAction(v string)`

SetUpdateAction sets UpdateAction field to given value.


### GetVersionType

`func (o *PlanConfigurationPost) GetVersionType() string`

GetVersionType returns the VersionType field if non-nil, zero value otherwise.

### GetVersionTypeOk

`func (o *PlanConfigurationPost) GetVersionTypeOk() (*string, bool)`

GetVersionTypeOk returns a tuple with the VersionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionType

`func (o *PlanConfigurationPost) SetVersionType(v string)`

SetVersionType sets VersionType field to given value.


### GetSpecificVersion

`func (o *PlanConfigurationPost) GetSpecificVersion() string`

GetSpecificVersion returns the SpecificVersion field if non-nil, zero value otherwise.

### GetSpecificVersionOk

`func (o *PlanConfigurationPost) GetSpecificVersionOk() (*string, bool)`

GetSpecificVersionOk returns a tuple with the SpecificVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificVersion

`func (o *PlanConfigurationPost) SetSpecificVersion(v string)`

SetSpecificVersion sets SpecificVersion field to given value.

### HasSpecificVersion

`func (o *PlanConfigurationPost) HasSpecificVersion() bool`

HasSpecificVersion returns a boolean if a field has been set.

### GetBuildVersion

`func (o *PlanConfigurationPost) GetBuildVersion() string`

GetBuildVersion returns the BuildVersion field if non-nil, zero value otherwise.

### GetBuildVersionOk

`func (o *PlanConfigurationPost) GetBuildVersionOk() (*string, bool)`

GetBuildVersionOk returns a tuple with the BuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVersion

`func (o *PlanConfigurationPost) SetBuildVersion(v string)`

SetBuildVersion sets BuildVersion field to given value.

### HasBuildVersion

`func (o *PlanConfigurationPost) HasBuildVersion() bool`

HasBuildVersion returns a boolean if a field has been set.

### SetBuildVersionNil

`func (o *PlanConfigurationPost) SetBuildVersionNil(b bool)`

 SetBuildVersionNil sets the value for BuildVersion to be an explicit nil

### UnsetBuildVersion
`func (o *PlanConfigurationPost) UnsetBuildVersion()`

UnsetBuildVersion ensures that no value is present for BuildVersion, not even an explicit nil
### GetMaxDeferrals

`func (o *PlanConfigurationPost) GetMaxDeferrals() int64`

GetMaxDeferrals returns the MaxDeferrals field if non-nil, zero value otherwise.

### GetMaxDeferralsOk

`func (o *PlanConfigurationPost) GetMaxDeferralsOk() (*int64, bool)`

GetMaxDeferralsOk returns a tuple with the MaxDeferrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeferrals

`func (o *PlanConfigurationPost) SetMaxDeferrals(v int64)`

SetMaxDeferrals sets MaxDeferrals field to given value.

### HasMaxDeferrals

`func (o *PlanConfigurationPost) HasMaxDeferrals() bool`

HasMaxDeferrals returns a boolean if a field has been set.

### GetForceInstallLocalDateTime

`func (o *PlanConfigurationPost) GetForceInstallLocalDateTime() string`

GetForceInstallLocalDateTime returns the ForceInstallLocalDateTime field if non-nil, zero value otherwise.

### GetForceInstallLocalDateTimeOk

`func (o *PlanConfigurationPost) GetForceInstallLocalDateTimeOk() (*string, bool)`

GetForceInstallLocalDateTimeOk returns a tuple with the ForceInstallLocalDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceInstallLocalDateTime

`func (o *PlanConfigurationPost) SetForceInstallLocalDateTime(v string)`

SetForceInstallLocalDateTime sets ForceInstallLocalDateTime field to given value.

### HasForceInstallLocalDateTime

`func (o *PlanConfigurationPost) HasForceInstallLocalDateTime() bool`

HasForceInstallLocalDateTime returns a boolean if a field has been set.

### SetForceInstallLocalDateTimeNil

`func (o *PlanConfigurationPost) SetForceInstallLocalDateTimeNil(b bool)`

 SetForceInstallLocalDateTimeNil sets the value for ForceInstallLocalDateTime to be an explicit nil

### UnsetForceInstallLocalDateTime
`func (o *PlanConfigurationPost) UnsetForceInstallLocalDateTime()`

UnsetForceInstallLocalDateTime ensures that no value is present for ForceInstallLocalDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


