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

// Udids struct for Udids
type Udids struct {
	Udids []string `json:"udids,omitempty"`
}

// NewUdids instantiates a new Udids object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUdids() *Udids {
	this := Udids{}
	return &this
}

// NewUdidsWithDefaults instantiates a new Udids object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUdidsWithDefaults() *Udids {
	this := Udids{}
	return &this
}

// GetUdids returns the Udids field value if set, zero value otherwise.
func (o *Udids) GetUdids() []string {
	if o == nil || o.Udids == nil {
		var ret []string
		return ret
	}
	return o.Udids
}

// GetUdidsOk returns a tuple with the Udids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Udids) GetUdidsOk() ([]string, bool) {
	if o == nil || o.Udids == nil {
		return nil, false
	}
	return o.Udids, true
}

// HasUdids returns a boolean if a field has been set.
func (o *Udids) HasUdids() bool {
	if o != nil && o.Udids != nil {
		return true
	}

	return false
}

// SetUdids gets a reference to the given []string and assigns it to the Udids field.
func (o *Udids) SetUdids(v []string) {
	o.Udids = v
}

func (o Udids) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Udids != nil {
		toSerialize["udids"] = o.Udids
	}
	return json.Marshal(toSerialize)
}

type NullableUdids struct {
	value *Udids
	isSet bool
}

func (v NullableUdids) Get() *Udids {
	return v.value
}

func (v *NullableUdids) Set(val *Udids) {
	v.value = val
	v.isSet = true
}

func (v NullableUdids) IsSet() bool {
	return v.isSet
}

func (v *NullableUdids) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUdids(val *Udids) *NullableUdids {
	return &NullableUdids{value: val, isSet: true}
}

func (v NullableUdids) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUdids) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


