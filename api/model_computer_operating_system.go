/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ComputerOperatingSystem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComputerOperatingSystem{}

// ComputerOperatingSystem struct for ComputerOperatingSystem
type ComputerOperatingSystem struct {
	Name *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
	Build *string `json:"build,omitempty"`
	// Collected for macOS 13.0 or later
	SupplementalBuildVersion *string `json:"supplementalBuildVersion,omitempty"`
	// Collected for macOS 13.0 or later
	RapidSecurityResponse *string `json:"rapidSecurityResponse,omitempty"`
	ActiveDirectoryStatus *string `json:"activeDirectoryStatus,omitempty"`
	FileVault2Status *string `json:"fileVault2Status,omitempty"`
	SoftwareUpdateDeviceId *string `json:"softwareUpdateDeviceId,omitempty"`
	ExtensionAttributes []ComputerExtensionAttribute `json:"extensionAttributes,omitempty"`
}

// NewComputerOperatingSystem instantiates a new ComputerOperatingSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputerOperatingSystem() *ComputerOperatingSystem {
	this := ComputerOperatingSystem{}
	return &this
}

// NewComputerOperatingSystemWithDefaults instantiates a new ComputerOperatingSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputerOperatingSystemWithDefaults() *ComputerOperatingSystem {
	this := ComputerOperatingSystem{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComputerOperatingSystem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOperatingSystem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComputerOperatingSystem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComputerOperatingSystem) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ComputerOperatingSystem) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOperatingSystem) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ComputerOperatingSystem) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ComputerOperatingSystem) SetVersion(v string) {
	o.Version = &v
}

// GetBuild returns the Build field value if set, zero value otherwise.
func (o *ComputerOperatingSystem) GetBuild() string {
	if o == nil || IsNil(o.Build) {
		var ret string
		return ret
	}
	return *o.Build
}

// GetBuildOk returns a tuple with the Build field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOperatingSystem) GetBuildOk() (*string, bool) {
	if o == nil || IsNil(o.Build) {
		return nil, false
	}
	return o.Build, true
}

// HasBuild returns a boolean if a field has been set.
func (o *ComputerOperatingSystem) HasBuild() bool {
	if o != nil && !IsNil(o.Build) {
		return true
	}

	return false
}

// SetBuild gets a reference to the given string and assigns it to the Build field.
func (o *ComputerOperatingSystem) SetBuild(v string) {
	o.Build = &v
}

// GetSupplementalBuildVersion returns the SupplementalBuildVersion field value if set, zero value otherwise.
func (o *ComputerOperatingSystem) GetSupplementalBuildVersion() string {
	if o == nil || IsNil(o.SupplementalBuildVersion) {
		var ret string
		return ret
	}
	return *o.SupplementalBuildVersion
}

// GetSupplementalBuildVersionOk returns a tuple with the SupplementalBuildVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOperatingSystem) GetSupplementalBuildVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SupplementalBuildVersion) {
		return nil, false
	}
	return o.SupplementalBuildVersion, true
}

// HasSupplementalBuildVersion returns a boolean if a field has been set.
func (o *ComputerOperatingSystem) HasSupplementalBuildVersion() bool {
	if o != nil && !IsNil(o.SupplementalBuildVersion) {
		return true
	}

	return false
}

// SetSupplementalBuildVersion gets a reference to the given string and assigns it to the SupplementalBuildVersion field.
func (o *ComputerOperatingSystem) SetSupplementalBuildVersion(v string) {
	o.SupplementalBuildVersion = &v
}

// GetRapidSecurityResponse returns the RapidSecurityResponse field value if set, zero value otherwise.
func (o *ComputerOperatingSystem) GetRapidSecurityResponse() string {
	if o == nil || IsNil(o.RapidSecurityResponse) {
		var ret string
		return ret
	}
	return *o.RapidSecurityResponse
}

// GetRapidSecurityResponseOk returns a tuple with the RapidSecurityResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOperatingSystem) GetRapidSecurityResponseOk() (*string, bool) {
	if o == nil || IsNil(o.RapidSecurityResponse) {
		return nil, false
	}
	return o.RapidSecurityResponse, true
}

// HasRapidSecurityResponse returns a boolean if a field has been set.
func (o *ComputerOperatingSystem) HasRapidSecurityResponse() bool {
	if o != nil && !IsNil(o.RapidSecurityResponse) {
		return true
	}

	return false
}

// SetRapidSecurityResponse gets a reference to the given string and assigns it to the RapidSecurityResponse field.
func (o *ComputerOperatingSystem) SetRapidSecurityResponse(v string) {
	o.RapidSecurityResponse = &v
}

// GetActiveDirectoryStatus returns the ActiveDirectoryStatus field value if set, zero value otherwise.
func (o *ComputerOperatingSystem) GetActiveDirectoryStatus() string {
	if o == nil || IsNil(o.ActiveDirectoryStatus) {
		var ret string
		return ret
	}
	return *o.ActiveDirectoryStatus
}

// GetActiveDirectoryStatusOk returns a tuple with the ActiveDirectoryStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOperatingSystem) GetActiveDirectoryStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ActiveDirectoryStatus) {
		return nil, false
	}
	return o.ActiveDirectoryStatus, true
}

// HasActiveDirectoryStatus returns a boolean if a field has been set.
func (o *ComputerOperatingSystem) HasActiveDirectoryStatus() bool {
	if o != nil && !IsNil(o.ActiveDirectoryStatus) {
		return true
	}

	return false
}

// SetActiveDirectoryStatus gets a reference to the given string and assigns it to the ActiveDirectoryStatus field.
func (o *ComputerOperatingSystem) SetActiveDirectoryStatus(v string) {
	o.ActiveDirectoryStatus = &v
}

// GetFileVault2Status returns the FileVault2Status field value if set, zero value otherwise.
func (o *ComputerOperatingSystem) GetFileVault2Status() string {
	if o == nil || IsNil(o.FileVault2Status) {
		var ret string
		return ret
	}
	return *o.FileVault2Status
}

// GetFileVault2StatusOk returns a tuple with the FileVault2Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOperatingSystem) GetFileVault2StatusOk() (*string, bool) {
	if o == nil || IsNil(o.FileVault2Status) {
		return nil, false
	}
	return o.FileVault2Status, true
}

// HasFileVault2Status returns a boolean if a field has been set.
func (o *ComputerOperatingSystem) HasFileVault2Status() bool {
	if o != nil && !IsNil(o.FileVault2Status) {
		return true
	}

	return false
}

// SetFileVault2Status gets a reference to the given string and assigns it to the FileVault2Status field.
func (o *ComputerOperatingSystem) SetFileVault2Status(v string) {
	o.FileVault2Status = &v
}

// GetSoftwareUpdateDeviceId returns the SoftwareUpdateDeviceId field value if set, zero value otherwise.
func (o *ComputerOperatingSystem) GetSoftwareUpdateDeviceId() string {
	if o == nil || IsNil(o.SoftwareUpdateDeviceId) {
		var ret string
		return ret
	}
	return *o.SoftwareUpdateDeviceId
}

// GetSoftwareUpdateDeviceIdOk returns a tuple with the SoftwareUpdateDeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOperatingSystem) GetSoftwareUpdateDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SoftwareUpdateDeviceId) {
		return nil, false
	}
	return o.SoftwareUpdateDeviceId, true
}

// HasSoftwareUpdateDeviceId returns a boolean if a field has been set.
func (o *ComputerOperatingSystem) HasSoftwareUpdateDeviceId() bool {
	if o != nil && !IsNil(o.SoftwareUpdateDeviceId) {
		return true
	}

	return false
}

// SetSoftwareUpdateDeviceId gets a reference to the given string and assigns it to the SoftwareUpdateDeviceId field.
func (o *ComputerOperatingSystem) SetSoftwareUpdateDeviceId(v string) {
	o.SoftwareUpdateDeviceId = &v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *ComputerOperatingSystem) GetExtensionAttributes() []ComputerExtensionAttribute {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret []ComputerExtensionAttribute
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOperatingSystem) GetExtensionAttributesOk() ([]ComputerExtensionAttribute, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return nil, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *ComputerOperatingSystem) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given []ComputerExtensionAttribute and assigns it to the ExtensionAttributes field.
func (o *ComputerOperatingSystem) SetExtensionAttributes(v []ComputerExtensionAttribute) {
	o.ExtensionAttributes = v
}

func (o ComputerOperatingSystem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComputerOperatingSystem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Build) {
		toSerialize["build"] = o.Build
	}
	if !IsNil(o.SupplementalBuildVersion) {
		toSerialize["supplementalBuildVersion"] = o.SupplementalBuildVersion
	}
	if !IsNil(o.RapidSecurityResponse) {
		toSerialize["rapidSecurityResponse"] = o.RapidSecurityResponse
	}
	if !IsNil(o.ActiveDirectoryStatus) {
		toSerialize["activeDirectoryStatus"] = o.ActiveDirectoryStatus
	}
	if !IsNil(o.FileVault2Status) {
		toSerialize["fileVault2Status"] = o.FileVault2Status
	}
	if !IsNil(o.SoftwareUpdateDeviceId) {
		toSerialize["softwareUpdateDeviceId"] = o.SoftwareUpdateDeviceId
	}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extensionAttributes"] = o.ExtensionAttributes
	}
	return toSerialize, nil
}

type NullableComputerOperatingSystem struct {
	value *ComputerOperatingSystem
	isSet bool
}

func (v NullableComputerOperatingSystem) Get() *ComputerOperatingSystem {
	return v.value
}

func (v *NullableComputerOperatingSystem) Set(val *ComputerOperatingSystem) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerOperatingSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerOperatingSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerOperatingSystem(val *ComputerOperatingSystem) *NullableComputerOperatingSystem {
	return &NullableComputerOperatingSystem{value: val, isSet: true}
}

func (v NullableComputerOperatingSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerOperatingSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


