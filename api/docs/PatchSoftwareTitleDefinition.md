# PatchSoftwareTitleDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] 
**MinimumOperatingSystem** | Pointer to **string** |  | [optional] [default to "-1"]
**ReleaseDate** | Pointer to **string** |  | [optional] [default to "-1"]
**RebootRequired** | Pointer to **bool** |  | [optional] [default to false]
**KillApps** | Pointer to [**[]PatchSoftwareTitleConfigurationDefinitionKillApp**](PatchSoftwareTitleConfigurationDefinitionKillApp.md) |  | [optional] 
**Standalone** | Pointer to **bool** |  | [optional] [default to false]
**AbsoluteOrderId** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchSoftwareTitleDefinition

`func NewPatchSoftwareTitleDefinition() *PatchSoftwareTitleDefinition`

NewPatchSoftwareTitleDefinition instantiates a new PatchSoftwareTitleDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchSoftwareTitleDefinitionWithDefaults

`func NewPatchSoftwareTitleDefinitionWithDefaults() *PatchSoftwareTitleDefinition`

NewPatchSoftwareTitleDefinitionWithDefaults instantiates a new PatchSoftwareTitleDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *PatchSoftwareTitleDefinition) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PatchSoftwareTitleDefinition) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PatchSoftwareTitleDefinition) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PatchSoftwareTitleDefinition) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetMinimumOperatingSystem

`func (o *PatchSoftwareTitleDefinition) GetMinimumOperatingSystem() string`

GetMinimumOperatingSystem returns the MinimumOperatingSystem field if non-nil, zero value otherwise.

### GetMinimumOperatingSystemOk

`func (o *PatchSoftwareTitleDefinition) GetMinimumOperatingSystemOk() (*string, bool)`

GetMinimumOperatingSystemOk returns a tuple with the MinimumOperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumOperatingSystem

`func (o *PatchSoftwareTitleDefinition) SetMinimumOperatingSystem(v string)`

SetMinimumOperatingSystem sets MinimumOperatingSystem field to given value.

### HasMinimumOperatingSystem

`func (o *PatchSoftwareTitleDefinition) HasMinimumOperatingSystem() bool`

HasMinimumOperatingSystem returns a boolean if a field has been set.

### GetReleaseDate

`func (o *PatchSoftwareTitleDefinition) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *PatchSoftwareTitleDefinition) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *PatchSoftwareTitleDefinition) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *PatchSoftwareTitleDefinition) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetRebootRequired

`func (o *PatchSoftwareTitleDefinition) GetRebootRequired() bool`

GetRebootRequired returns the RebootRequired field if non-nil, zero value otherwise.

### GetRebootRequiredOk

`func (o *PatchSoftwareTitleDefinition) GetRebootRequiredOk() (*bool, bool)`

GetRebootRequiredOk returns a tuple with the RebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootRequired

`func (o *PatchSoftwareTitleDefinition) SetRebootRequired(v bool)`

SetRebootRequired sets RebootRequired field to given value.

### HasRebootRequired

`func (o *PatchSoftwareTitleDefinition) HasRebootRequired() bool`

HasRebootRequired returns a boolean if a field has been set.

### GetKillApps

`func (o *PatchSoftwareTitleDefinition) GetKillApps() []PatchSoftwareTitleConfigurationDefinitionKillApp`

GetKillApps returns the KillApps field if non-nil, zero value otherwise.

### GetKillAppsOk

`func (o *PatchSoftwareTitleDefinition) GetKillAppsOk() (*[]PatchSoftwareTitleConfigurationDefinitionKillApp, bool)`

GetKillAppsOk returns a tuple with the KillApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillApps

`func (o *PatchSoftwareTitleDefinition) SetKillApps(v []PatchSoftwareTitleConfigurationDefinitionKillApp)`

SetKillApps sets KillApps field to given value.

### HasKillApps

`func (o *PatchSoftwareTitleDefinition) HasKillApps() bool`

HasKillApps returns a boolean if a field has been set.

### GetStandalone

`func (o *PatchSoftwareTitleDefinition) GetStandalone() bool`

GetStandalone returns the Standalone field if non-nil, zero value otherwise.

### GetStandaloneOk

`func (o *PatchSoftwareTitleDefinition) GetStandaloneOk() (*bool, bool)`

GetStandaloneOk returns a tuple with the Standalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandalone

`func (o *PatchSoftwareTitleDefinition) SetStandalone(v bool)`

SetStandalone sets Standalone field to given value.

### HasStandalone

`func (o *PatchSoftwareTitleDefinition) HasStandalone() bool`

HasStandalone returns a boolean if a field has been set.

### GetAbsoluteOrderId

`func (o *PatchSoftwareTitleDefinition) GetAbsoluteOrderId() string`

GetAbsoluteOrderId returns the AbsoluteOrderId field if non-nil, zero value otherwise.

### GetAbsoluteOrderIdOk

`func (o *PatchSoftwareTitleDefinition) GetAbsoluteOrderIdOk() (*string, bool)`

GetAbsoluteOrderIdOk returns a tuple with the AbsoluteOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteOrderId

`func (o *PatchSoftwareTitleDefinition) SetAbsoluteOrderId(v string)`

SetAbsoluteOrderId sets AbsoluteOrderId field to given value.

### HasAbsoluteOrderId

`func (o *PatchSoftwareTitleDefinition) HasAbsoluteOrderId() bool`

HasAbsoluteOrderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


