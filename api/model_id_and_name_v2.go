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

// IdAndNameV2 struct for IdAndNameV2
type IdAndNameV2 struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewIdAndNameV2 instantiates a new IdAndNameV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdAndNameV2() *IdAndNameV2 {
	this := IdAndNameV2{}
	return &this
}

// NewIdAndNameV2WithDefaults instantiates a new IdAndNameV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdAndNameV2WithDefaults() *IdAndNameV2 {
	this := IdAndNameV2{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdAndNameV2) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdAndNameV2) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdAndNameV2) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdAndNameV2) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdAndNameV2) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdAndNameV2) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdAndNameV2) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdAndNameV2) SetName(v string) {
	o.Name = &v
}

func (o IdAndNameV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableIdAndNameV2 struct {
	value *IdAndNameV2
	isSet bool
}

func (v NullableIdAndNameV2) Get() *IdAndNameV2 {
	return v.value
}

func (v *NullableIdAndNameV2) Set(val *IdAndNameV2) {
	v.value = val
	v.isSet = true
}

func (v NullableIdAndNameV2) IsSet() bool {
	return v.isSet
}

func (v *NullableIdAndNameV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdAndNameV2(val *IdAndNameV2) *NullableIdAndNameV2 {
	return &NullableIdAndNameV2{value: val, isSet: true}
}

func (v NullableIdAndNameV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdAndNameV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

