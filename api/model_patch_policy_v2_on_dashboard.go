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

// checks if the PatchPolicyV2OnDashboard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchPolicyV2OnDashboard{}

// PatchPolicyV2OnDashboard struct for PatchPolicyV2OnDashboard
type PatchPolicyV2OnDashboard struct {
	OnDashboard *bool `json:"onDashboard,omitempty"`
}

// NewPatchPolicyV2OnDashboard instantiates a new PatchPolicyV2OnDashboard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchPolicyV2OnDashboard() *PatchPolicyV2OnDashboard {
	this := PatchPolicyV2OnDashboard{}
	return &this
}

// NewPatchPolicyV2OnDashboardWithDefaults instantiates a new PatchPolicyV2OnDashboard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchPolicyV2OnDashboardWithDefaults() *PatchPolicyV2OnDashboard {
	this := PatchPolicyV2OnDashboard{}
	return &this
}

// GetOnDashboard returns the OnDashboard field value if set, zero value otherwise.
func (o *PatchPolicyV2OnDashboard) GetOnDashboard() bool {
	if o == nil || IsNil(o.OnDashboard) {
		var ret bool
		return ret
	}
	return *o.OnDashboard
}

// GetOnDashboardOk returns a tuple with the OnDashboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPolicyV2OnDashboard) GetOnDashboardOk() (*bool, bool) {
	if o == nil || IsNil(o.OnDashboard) {
		return nil, false
	}
	return o.OnDashboard, true
}

// HasOnDashboard returns a boolean if a field has been set.
func (o *PatchPolicyV2OnDashboard) HasOnDashboard() bool {
	if o != nil && !IsNil(o.OnDashboard) {
		return true
	}

	return false
}

// SetOnDashboard gets a reference to the given bool and assigns it to the OnDashboard field.
func (o *PatchPolicyV2OnDashboard) SetOnDashboard(v bool) {
	o.OnDashboard = &v
}

func (o PatchPolicyV2OnDashboard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchPolicyV2OnDashboard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OnDashboard) {
		toSerialize["onDashboard"] = o.OnDashboard
	}
	return toSerialize, nil
}

type NullablePatchPolicyV2OnDashboard struct {
	value *PatchPolicyV2OnDashboard
	isSet bool
}

func (v NullablePatchPolicyV2OnDashboard) Get() *PatchPolicyV2OnDashboard {
	return v.value
}

func (v *NullablePatchPolicyV2OnDashboard) Set(val *PatchPolicyV2OnDashboard) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchPolicyV2OnDashboard) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchPolicyV2OnDashboard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchPolicyV2OnDashboard(val *PatchPolicyV2OnDashboard) *NullablePatchPolicyV2OnDashboard {
	return &NullablePatchPolicyV2OnDashboard{value: val, isSet: true}
}

func (v NullablePatchPolicyV2OnDashboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchPolicyV2OnDashboard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


