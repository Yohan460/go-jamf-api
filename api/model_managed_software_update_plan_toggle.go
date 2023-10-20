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

// checks if the ManagedSoftwareUpdatePlanToggle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedSoftwareUpdatePlanToggle{}

// ManagedSoftwareUpdatePlanToggle struct for ManagedSoftwareUpdatePlanToggle
type ManagedSoftwareUpdatePlanToggle struct {
	Toggle bool `json:"toggle"`
	ForceInstallLocalDateEnabled *bool `json:"forceInstallLocalDateEnabled,omitempty"`
	DssEnabled *bool `json:"dssEnabled,omitempty"`
}

// NewManagedSoftwareUpdatePlanToggle instantiates a new ManagedSoftwareUpdatePlanToggle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedSoftwareUpdatePlanToggle(toggle bool) *ManagedSoftwareUpdatePlanToggle {
	this := ManagedSoftwareUpdatePlanToggle{}
	this.Toggle = toggle
	return &this
}

// NewManagedSoftwareUpdatePlanToggleWithDefaults instantiates a new ManagedSoftwareUpdatePlanToggle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedSoftwareUpdatePlanToggleWithDefaults() *ManagedSoftwareUpdatePlanToggle {
	this := ManagedSoftwareUpdatePlanToggle{}
	return &this
}

// GetToggle returns the Toggle field value
func (o *ManagedSoftwareUpdatePlanToggle) GetToggle() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Toggle
}

// GetToggleOk returns a tuple with the Toggle field value
// and a boolean to check if the value has been set.
func (o *ManagedSoftwareUpdatePlanToggle) GetToggleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Toggle, true
}

// SetToggle sets field value
func (o *ManagedSoftwareUpdatePlanToggle) SetToggle(v bool) {
	o.Toggle = v
}

// GetForceInstallLocalDateEnabled returns the ForceInstallLocalDateEnabled field value if set, zero value otherwise.
func (o *ManagedSoftwareUpdatePlanToggle) GetForceInstallLocalDateEnabled() bool {
	if o == nil || IsNil(o.ForceInstallLocalDateEnabled) {
		var ret bool
		return ret
	}
	return *o.ForceInstallLocalDateEnabled
}

// GetForceInstallLocalDateEnabledOk returns a tuple with the ForceInstallLocalDateEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedSoftwareUpdatePlanToggle) GetForceInstallLocalDateEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceInstallLocalDateEnabled) {
		return nil, false
	}
	return o.ForceInstallLocalDateEnabled, true
}

// HasForceInstallLocalDateEnabled returns a boolean if a field has been set.
func (o *ManagedSoftwareUpdatePlanToggle) HasForceInstallLocalDateEnabled() bool {
	if o != nil && !IsNil(o.ForceInstallLocalDateEnabled) {
		return true
	}

	return false
}

// SetForceInstallLocalDateEnabled gets a reference to the given bool and assigns it to the ForceInstallLocalDateEnabled field.
func (o *ManagedSoftwareUpdatePlanToggle) SetForceInstallLocalDateEnabled(v bool) {
	o.ForceInstallLocalDateEnabled = &v
}

// GetDssEnabled returns the DssEnabled field value if set, zero value otherwise.
func (o *ManagedSoftwareUpdatePlanToggle) GetDssEnabled() bool {
	if o == nil || IsNil(o.DssEnabled) {
		var ret bool
		return ret
	}
	return *o.DssEnabled
}

// GetDssEnabledOk returns a tuple with the DssEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedSoftwareUpdatePlanToggle) GetDssEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DssEnabled) {
		return nil, false
	}
	return o.DssEnabled, true
}

// HasDssEnabled returns a boolean if a field has been set.
func (o *ManagedSoftwareUpdatePlanToggle) HasDssEnabled() bool {
	if o != nil && !IsNil(o.DssEnabled) {
		return true
	}

	return false
}

// SetDssEnabled gets a reference to the given bool and assigns it to the DssEnabled field.
func (o *ManagedSoftwareUpdatePlanToggle) SetDssEnabled(v bool) {
	o.DssEnabled = &v
}

func (o ManagedSoftwareUpdatePlanToggle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedSoftwareUpdatePlanToggle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["toggle"] = o.Toggle
	if !IsNil(o.ForceInstallLocalDateEnabled) {
		toSerialize["forceInstallLocalDateEnabled"] = o.ForceInstallLocalDateEnabled
	}
	if !IsNil(o.DssEnabled) {
		toSerialize["dssEnabled"] = o.DssEnabled
	}
	return toSerialize, nil
}

type NullableManagedSoftwareUpdatePlanToggle struct {
	value *ManagedSoftwareUpdatePlanToggle
	isSet bool
}

func (v NullableManagedSoftwareUpdatePlanToggle) Get() *ManagedSoftwareUpdatePlanToggle {
	return v.value
}

func (v *NullableManagedSoftwareUpdatePlanToggle) Set(val *ManagedSoftwareUpdatePlanToggle) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedSoftwareUpdatePlanToggle) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedSoftwareUpdatePlanToggle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedSoftwareUpdatePlanToggle(val *ManagedSoftwareUpdatePlanToggle) *NullableManagedSoftwareUpdatePlanToggle {
	return &NullableManagedSoftwareUpdatePlanToggle{value: val, isSet: true}
}

func (v NullableManagedSoftwareUpdatePlanToggle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedSoftwareUpdatePlanToggle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


