# ComputerOperatingSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 
**Build** | Pointer to **string** |  | [optional] [readonly] 
**SupplementalBuildVersion** | Pointer to **string** | Collected for macOS 13.0 or later | [optional] [readonly] 
**RapidSecurityResponse** | Pointer to **string** | Collected for macOS 13.0 or later | [optional] [readonly] 
**ActiveDirectoryStatus** | Pointer to **string** |  | [optional] [readonly] 
**FileVault2Status** | Pointer to **string** |  | [optional] 
**SoftwareUpdateDeviceId** | Pointer to **string** |  | [optional] [readonly] 
**ExtensionAttributes** | Pointer to [**[]ComputerExtensionAttribute**](ComputerExtensionAttribute.md) |  | [optional] 

## Methods

### NewComputerOperatingSystem

`func NewComputerOperatingSystem() *ComputerOperatingSystem`

NewComputerOperatingSystem instantiates a new ComputerOperatingSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerOperatingSystemWithDefaults

`func NewComputerOperatingSystemWithDefaults() *ComputerOperatingSystem`

NewComputerOperatingSystemWithDefaults instantiates a new ComputerOperatingSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ComputerOperatingSystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputerOperatingSystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputerOperatingSystem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputerOperatingSystem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *ComputerOperatingSystem) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ComputerOperatingSystem) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ComputerOperatingSystem) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ComputerOperatingSystem) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBuild

`func (o *ComputerOperatingSystem) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *ComputerOperatingSystem) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *ComputerOperatingSystem) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *ComputerOperatingSystem) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetSupplementalBuildVersion

`func (o *ComputerOperatingSystem) GetSupplementalBuildVersion() string`

GetSupplementalBuildVersion returns the SupplementalBuildVersion field if non-nil, zero value otherwise.

### GetSupplementalBuildVersionOk

`func (o *ComputerOperatingSystem) GetSupplementalBuildVersionOk() (*string, bool)`

GetSupplementalBuildVersionOk returns a tuple with the SupplementalBuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementalBuildVersion

`func (o *ComputerOperatingSystem) SetSupplementalBuildVersion(v string)`

SetSupplementalBuildVersion sets SupplementalBuildVersion field to given value.

### HasSupplementalBuildVersion

`func (o *ComputerOperatingSystem) HasSupplementalBuildVersion() bool`

HasSupplementalBuildVersion returns a boolean if a field has been set.

### GetRapidSecurityResponse

`func (o *ComputerOperatingSystem) GetRapidSecurityResponse() string`

GetRapidSecurityResponse returns the RapidSecurityResponse field if non-nil, zero value otherwise.

### GetRapidSecurityResponseOk

`func (o *ComputerOperatingSystem) GetRapidSecurityResponseOk() (*string, bool)`

GetRapidSecurityResponseOk returns a tuple with the RapidSecurityResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRapidSecurityResponse

`func (o *ComputerOperatingSystem) SetRapidSecurityResponse(v string)`

SetRapidSecurityResponse sets RapidSecurityResponse field to given value.

### HasRapidSecurityResponse

`func (o *ComputerOperatingSystem) HasRapidSecurityResponse() bool`

HasRapidSecurityResponse returns a boolean if a field has been set.

### GetActiveDirectoryStatus

`func (o *ComputerOperatingSystem) GetActiveDirectoryStatus() string`

GetActiveDirectoryStatus returns the ActiveDirectoryStatus field if non-nil, zero value otherwise.

### GetActiveDirectoryStatusOk

`func (o *ComputerOperatingSystem) GetActiveDirectoryStatusOk() (*string, bool)`

GetActiveDirectoryStatusOk returns a tuple with the ActiveDirectoryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryStatus

`func (o *ComputerOperatingSystem) SetActiveDirectoryStatus(v string)`

SetActiveDirectoryStatus sets ActiveDirectoryStatus field to given value.

### HasActiveDirectoryStatus

`func (o *ComputerOperatingSystem) HasActiveDirectoryStatus() bool`

HasActiveDirectoryStatus returns a boolean if a field has been set.

### GetFileVault2Status

`func (o *ComputerOperatingSystem) GetFileVault2Status() string`

GetFileVault2Status returns the FileVault2Status field if non-nil, zero value otherwise.

### GetFileVault2StatusOk

`func (o *ComputerOperatingSystem) GetFileVault2StatusOk() (*string, bool)`

GetFileVault2StatusOk returns a tuple with the FileVault2Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileVault2Status

`func (o *ComputerOperatingSystem) SetFileVault2Status(v string)`

SetFileVault2Status sets FileVault2Status field to given value.

### HasFileVault2Status

`func (o *ComputerOperatingSystem) HasFileVault2Status() bool`

HasFileVault2Status returns a boolean if a field has been set.

### GetSoftwareUpdateDeviceId

`func (o *ComputerOperatingSystem) GetSoftwareUpdateDeviceId() string`

GetSoftwareUpdateDeviceId returns the SoftwareUpdateDeviceId field if non-nil, zero value otherwise.

### GetSoftwareUpdateDeviceIdOk

`func (o *ComputerOperatingSystem) GetSoftwareUpdateDeviceIdOk() (*string, bool)`

GetSoftwareUpdateDeviceIdOk returns a tuple with the SoftwareUpdateDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdateDeviceId

`func (o *ComputerOperatingSystem) SetSoftwareUpdateDeviceId(v string)`

SetSoftwareUpdateDeviceId sets SoftwareUpdateDeviceId field to given value.

### HasSoftwareUpdateDeviceId

`func (o *ComputerOperatingSystem) HasSoftwareUpdateDeviceId() bool`

HasSoftwareUpdateDeviceId returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *ComputerOperatingSystem) GetExtensionAttributes() []ComputerExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *ComputerOperatingSystem) GetExtensionAttributesOk() (*[]ComputerExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *ComputerOperatingSystem) SetExtensionAttributes(v []ComputerExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *ComputerOperatingSystem) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


