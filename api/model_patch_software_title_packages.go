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

// checks if the PatchSoftwareTitlePackages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchSoftwareTitlePackages{}

// PatchSoftwareTitlePackages struct for PatchSoftwareTitlePackages
type PatchSoftwareTitlePackages struct {
	PackageId *string `json:"packageId,omitempty"`
	Version *string `json:"version,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
}

// NewPatchSoftwareTitlePackages instantiates a new PatchSoftwareTitlePackages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchSoftwareTitlePackages() *PatchSoftwareTitlePackages {
	this := PatchSoftwareTitlePackages{}
	return &this
}

// NewPatchSoftwareTitlePackagesWithDefaults instantiates a new PatchSoftwareTitlePackages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchSoftwareTitlePackagesWithDefaults() *PatchSoftwareTitlePackages {
	this := PatchSoftwareTitlePackages{}
	return &this
}

// GetPackageId returns the PackageId field value if set, zero value otherwise.
func (o *PatchSoftwareTitlePackages) GetPackageId() string {
	if o == nil || IsNil(o.PackageId) {
		var ret string
		return ret
	}
	return *o.PackageId
}

// GetPackageIdOk returns a tuple with the PackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitlePackages) GetPackageIdOk() (*string, bool) {
	if o == nil || IsNil(o.PackageId) {
		return nil, false
	}
	return o.PackageId, true
}

// HasPackageId returns a boolean if a field has been set.
func (o *PatchSoftwareTitlePackages) HasPackageId() bool {
	if o != nil && !IsNil(o.PackageId) {
		return true
	}

	return false
}

// SetPackageId gets a reference to the given string and assigns it to the PackageId field.
func (o *PatchSoftwareTitlePackages) SetPackageId(v string) {
	o.PackageId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PatchSoftwareTitlePackages) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitlePackages) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PatchSoftwareTitlePackages) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PatchSoftwareTitlePackages) SetVersion(v string) {
	o.Version = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PatchSoftwareTitlePackages) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitlePackages) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PatchSoftwareTitlePackages) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PatchSoftwareTitlePackages) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o PatchSoftwareTitlePackages) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchSoftwareTitlePackages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PackageId) {
		toSerialize["packageId"] = o.PackageId
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	return toSerialize, nil
}

type NullablePatchSoftwareTitlePackages struct {
	value *PatchSoftwareTitlePackages
	isSet bool
}

func (v NullablePatchSoftwareTitlePackages) Get() *PatchSoftwareTitlePackages {
	return v.value
}

func (v *NullablePatchSoftwareTitlePackages) Set(val *PatchSoftwareTitlePackages) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchSoftwareTitlePackages) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchSoftwareTitlePackages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchSoftwareTitlePackages(val *PatchSoftwareTitlePackages) *NullablePatchSoftwareTitlePackages {
	return &NullablePatchSoftwareTitlePackages{value: val, isSet: true}
}

func (v NullablePatchSoftwareTitlePackages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchSoftwareTitlePackages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


