# ComputerOperatingSystemCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**Build** | Pointer to **NullableString** |  | [optional] 
**SupplementalBuildVersion** | Pointer to **NullableString** |  | [optional] 
**RapidSecurityResponse** | Pointer to **NullableString** |  | [optional] 
**ActiveDirectoryStatus** | Pointer to **NullableString** |  | [optional] 
**SoftwareUpdateDeviceId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewComputerOperatingSystemCreate

`func NewComputerOperatingSystemCreate() *ComputerOperatingSystemCreate`

NewComputerOperatingSystemCreate instantiates a new ComputerOperatingSystemCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerOperatingSystemCreateWithDefaults

`func NewComputerOperatingSystemCreateWithDefaults() *ComputerOperatingSystemCreate`

NewComputerOperatingSystemCreateWithDefaults instantiates a new ComputerOperatingSystemCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ComputerOperatingSystemCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputerOperatingSystemCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputerOperatingSystemCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputerOperatingSystemCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ComputerOperatingSystemCreate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ComputerOperatingSystemCreate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVersion

`func (o *ComputerOperatingSystemCreate) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ComputerOperatingSystemCreate) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ComputerOperatingSystemCreate) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ComputerOperatingSystemCreate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ComputerOperatingSystemCreate) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ComputerOperatingSystemCreate) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetBuild

`func (o *ComputerOperatingSystemCreate) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *ComputerOperatingSystemCreate) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *ComputerOperatingSystemCreate) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *ComputerOperatingSystemCreate) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### SetBuildNil

`func (o *ComputerOperatingSystemCreate) SetBuildNil(b bool)`

 SetBuildNil sets the value for Build to be an explicit nil

### UnsetBuild
`func (o *ComputerOperatingSystemCreate) UnsetBuild()`

UnsetBuild ensures that no value is present for Build, not even an explicit nil
### GetSupplementalBuildVersion

`func (o *ComputerOperatingSystemCreate) GetSupplementalBuildVersion() string`

GetSupplementalBuildVersion returns the SupplementalBuildVersion field if non-nil, zero value otherwise.

### GetSupplementalBuildVersionOk

`func (o *ComputerOperatingSystemCreate) GetSupplementalBuildVersionOk() (*string, bool)`

GetSupplementalBuildVersionOk returns a tuple with the SupplementalBuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementalBuildVersion

`func (o *ComputerOperatingSystemCreate) SetSupplementalBuildVersion(v string)`

SetSupplementalBuildVersion sets SupplementalBuildVersion field to given value.

### HasSupplementalBuildVersion

`func (o *ComputerOperatingSystemCreate) HasSupplementalBuildVersion() bool`

HasSupplementalBuildVersion returns a boolean if a field has been set.

### SetSupplementalBuildVersionNil

`func (o *ComputerOperatingSystemCreate) SetSupplementalBuildVersionNil(b bool)`

 SetSupplementalBuildVersionNil sets the value for SupplementalBuildVersion to be an explicit nil

### UnsetSupplementalBuildVersion
`func (o *ComputerOperatingSystemCreate) UnsetSupplementalBuildVersion()`

UnsetSupplementalBuildVersion ensures that no value is present for SupplementalBuildVersion, not even an explicit nil
### GetRapidSecurityResponse

`func (o *ComputerOperatingSystemCreate) GetRapidSecurityResponse() string`

GetRapidSecurityResponse returns the RapidSecurityResponse field if non-nil, zero value otherwise.

### GetRapidSecurityResponseOk

`func (o *ComputerOperatingSystemCreate) GetRapidSecurityResponseOk() (*string, bool)`

GetRapidSecurityResponseOk returns a tuple with the RapidSecurityResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRapidSecurityResponse

`func (o *ComputerOperatingSystemCreate) SetRapidSecurityResponse(v string)`

SetRapidSecurityResponse sets RapidSecurityResponse field to given value.

### HasRapidSecurityResponse

`func (o *ComputerOperatingSystemCreate) HasRapidSecurityResponse() bool`

HasRapidSecurityResponse returns a boolean if a field has been set.

### SetRapidSecurityResponseNil

`func (o *ComputerOperatingSystemCreate) SetRapidSecurityResponseNil(b bool)`

 SetRapidSecurityResponseNil sets the value for RapidSecurityResponse to be an explicit nil

### UnsetRapidSecurityResponse
`func (o *ComputerOperatingSystemCreate) UnsetRapidSecurityResponse()`

UnsetRapidSecurityResponse ensures that no value is present for RapidSecurityResponse, not even an explicit nil
### GetActiveDirectoryStatus

`func (o *ComputerOperatingSystemCreate) GetActiveDirectoryStatus() string`

GetActiveDirectoryStatus returns the ActiveDirectoryStatus field if non-nil, zero value otherwise.

### GetActiveDirectoryStatusOk

`func (o *ComputerOperatingSystemCreate) GetActiveDirectoryStatusOk() (*string, bool)`

GetActiveDirectoryStatusOk returns a tuple with the ActiveDirectoryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryStatus

`func (o *ComputerOperatingSystemCreate) SetActiveDirectoryStatus(v string)`

SetActiveDirectoryStatus sets ActiveDirectoryStatus field to given value.

### HasActiveDirectoryStatus

`func (o *ComputerOperatingSystemCreate) HasActiveDirectoryStatus() bool`

HasActiveDirectoryStatus returns a boolean if a field has been set.

### SetActiveDirectoryStatusNil

`func (o *ComputerOperatingSystemCreate) SetActiveDirectoryStatusNil(b bool)`

 SetActiveDirectoryStatusNil sets the value for ActiveDirectoryStatus to be an explicit nil

### UnsetActiveDirectoryStatus
`func (o *ComputerOperatingSystemCreate) UnsetActiveDirectoryStatus()`

UnsetActiveDirectoryStatus ensures that no value is present for ActiveDirectoryStatus, not even an explicit nil
### GetSoftwareUpdateDeviceId

`func (o *ComputerOperatingSystemCreate) GetSoftwareUpdateDeviceId() string`

GetSoftwareUpdateDeviceId returns the SoftwareUpdateDeviceId field if non-nil, zero value otherwise.

### GetSoftwareUpdateDeviceIdOk

`func (o *ComputerOperatingSystemCreate) GetSoftwareUpdateDeviceIdOk() (*string, bool)`

GetSoftwareUpdateDeviceIdOk returns a tuple with the SoftwareUpdateDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdateDeviceId

`func (o *ComputerOperatingSystemCreate) SetSoftwareUpdateDeviceId(v string)`

SetSoftwareUpdateDeviceId sets SoftwareUpdateDeviceId field to given value.

### HasSoftwareUpdateDeviceId

`func (o *ComputerOperatingSystemCreate) HasSoftwareUpdateDeviceId() bool`

HasSoftwareUpdateDeviceId returns a boolean if a field has been set.

### SetSoftwareUpdateDeviceIdNil

`func (o *ComputerOperatingSystemCreate) SetSoftwareUpdateDeviceIdNil(b bool)`

 SetSoftwareUpdateDeviceIdNil sets the value for SoftwareUpdateDeviceId to be an explicit nil

### UnsetSoftwareUpdateDeviceId
`func (o *ComputerOperatingSystemCreate) UnsetSoftwareUpdateDeviceId()`

UnsetSoftwareUpdateDeviceId ensures that no value is present for SoftwareUpdateDeviceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


