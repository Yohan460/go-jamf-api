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

// checks if the PrestageScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrestageScope{}

// PrestageScope struct for PrestageScope
type PrestageScope struct {
	SerialsByPrestageId *map[string]int64 `json:"serialsByPrestageId,omitempty"`
}

// NewPrestageScope instantiates a new PrestageScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrestageScope() *PrestageScope {
	this := PrestageScope{}
	return &this
}

// NewPrestageScopeWithDefaults instantiates a new PrestageScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrestageScopeWithDefaults() *PrestageScope {
	this := PrestageScope{}
	return &this
}

// GetSerialsByPrestageId returns the SerialsByPrestageId field value if set, zero value otherwise.
func (o *PrestageScope) GetSerialsByPrestageId() map[string]int64 {
	if o == nil || IsNil(o.SerialsByPrestageId) {
		var ret map[string]int64
		return ret
	}
	return *o.SerialsByPrestageId
}

// GetSerialsByPrestageIdOk returns a tuple with the SerialsByPrestageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrestageScope) GetSerialsByPrestageIdOk() (*map[string]int64, bool) {
	if o == nil || IsNil(o.SerialsByPrestageId) {
		return nil, false
	}
	return o.SerialsByPrestageId, true
}

// HasSerialsByPrestageId returns a boolean if a field has been set.
func (o *PrestageScope) HasSerialsByPrestageId() bool {
	if o != nil && !IsNil(o.SerialsByPrestageId) {
		return true
	}

	return false
}

// SetSerialsByPrestageId gets a reference to the given map[string]int64 and assigns it to the SerialsByPrestageId field.
func (o *PrestageScope) SetSerialsByPrestageId(v map[string]int64) {
	o.SerialsByPrestageId = &v
}

func (o PrestageScope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrestageScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SerialsByPrestageId) {
		toSerialize["serialsByPrestageId"] = o.SerialsByPrestageId
	}
	return toSerialize, nil
}

type NullablePrestageScope struct {
	value *PrestageScope
	isSet bool
}

func (v NullablePrestageScope) Get() *PrestageScope {
	return v.value
}

func (v *NullablePrestageScope) Set(val *PrestageScope) {
	v.value = val
	v.isSet = true
}

func (v NullablePrestageScope) IsSet() bool {
	return v.isSet
}

func (v *NullablePrestageScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrestageScope(val *PrestageScope) *NullablePrestageScope {
	return &NullablePrestageScope{value: val, isSet: true}
}

func (v NullablePrestageScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrestageScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


