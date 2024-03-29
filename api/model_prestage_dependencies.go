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

// checks if the PrestageDependencies type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrestageDependencies{}

// PrestageDependencies struct for PrestageDependencies
type PrestageDependencies struct {
	Dependencies []PrestageDependency `json:"dependencies,omitempty"`
}

// NewPrestageDependencies instantiates a new PrestageDependencies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrestageDependencies() *PrestageDependencies {
	this := PrestageDependencies{}
	return &this
}

// NewPrestageDependenciesWithDefaults instantiates a new PrestageDependencies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrestageDependenciesWithDefaults() *PrestageDependencies {
	this := PrestageDependencies{}
	return &this
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *PrestageDependencies) GetDependencies() []PrestageDependency {
	if o == nil || IsNil(o.Dependencies) {
		var ret []PrestageDependency
		return ret
	}
	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrestageDependencies) GetDependenciesOk() ([]PrestageDependency, bool) {
	if o == nil || IsNil(o.Dependencies) {
		return nil, false
	}
	return o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *PrestageDependencies) HasDependencies() bool {
	if o != nil && !IsNil(o.Dependencies) {
		return true
	}

	return false
}

// SetDependencies gets a reference to the given []PrestageDependency and assigns it to the Dependencies field.
func (o *PrestageDependencies) SetDependencies(v []PrestageDependency) {
	o.Dependencies = v
}

func (o PrestageDependencies) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrestageDependencies) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dependencies) {
		toSerialize["dependencies"] = o.Dependencies
	}
	return toSerialize, nil
}

type NullablePrestageDependencies struct {
	value *PrestageDependencies
	isSet bool
}

func (v NullablePrestageDependencies) Get() *PrestageDependencies {
	return v.value
}

func (v *NullablePrestageDependencies) Set(val *PrestageDependencies) {
	v.value = val
	v.isSet = true
}

func (v NullablePrestageDependencies) IsSet() bool {
	return v.isSet
}

func (v *NullablePrestageDependencies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrestageDependencies(val *PrestageDependencies) *NullablePrestageDependencies {
	return &NullablePrestageDependencies{value: val, isSet: true}
}

func (v NullablePrestageDependencies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrestageDependencies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


