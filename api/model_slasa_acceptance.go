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

// checks if the SlasaAcceptance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlasaAcceptance{}

// SlasaAcceptance struct for SlasaAcceptance
type SlasaAcceptance struct {
	SlasaAcceptanceStatus *string `json:"slasaAcceptanceStatus,omitempty"`
}

// NewSlasaAcceptance instantiates a new SlasaAcceptance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlasaAcceptance() *SlasaAcceptance {
	this := SlasaAcceptance{}
	return &this
}

// NewSlasaAcceptanceWithDefaults instantiates a new SlasaAcceptance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlasaAcceptanceWithDefaults() *SlasaAcceptance {
	this := SlasaAcceptance{}
	return &this
}

// GetSlasaAcceptanceStatus returns the SlasaAcceptanceStatus field value if set, zero value otherwise.
func (o *SlasaAcceptance) GetSlasaAcceptanceStatus() string {
	if o == nil || IsNil(o.SlasaAcceptanceStatus) {
		var ret string
		return ret
	}
	return *o.SlasaAcceptanceStatus
}

// GetSlasaAcceptanceStatusOk returns a tuple with the SlasaAcceptanceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlasaAcceptance) GetSlasaAcceptanceStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SlasaAcceptanceStatus) {
		return nil, false
	}
	return o.SlasaAcceptanceStatus, true
}

// HasSlasaAcceptanceStatus returns a boolean if a field has been set.
func (o *SlasaAcceptance) HasSlasaAcceptanceStatus() bool {
	if o != nil && !IsNil(o.SlasaAcceptanceStatus) {
		return true
	}

	return false
}

// SetSlasaAcceptanceStatus gets a reference to the given string and assigns it to the SlasaAcceptanceStatus field.
func (o *SlasaAcceptance) SetSlasaAcceptanceStatus(v string) {
	o.SlasaAcceptanceStatus = &v
}

func (o SlasaAcceptance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlasaAcceptance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SlasaAcceptanceStatus) {
		toSerialize["slasaAcceptanceStatus"] = o.SlasaAcceptanceStatus
	}
	return toSerialize, nil
}

type NullableSlasaAcceptance struct {
	value *SlasaAcceptance
	isSet bool
}

func (v NullableSlasaAcceptance) Get() *SlasaAcceptance {
	return v.value
}

func (v *NullableSlasaAcceptance) Set(val *SlasaAcceptance) {
	v.value = val
	v.isSet = true
}

func (v NullableSlasaAcceptance) IsSet() bool {
	return v.isSet
}

func (v *NullableSlasaAcceptance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlasaAcceptance(val *SlasaAcceptance) *NullableSlasaAcceptance {
	return &NullableSlasaAcceptance{value: val, isSet: true}
}

func (v NullableSlasaAcceptance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlasaAcceptance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


