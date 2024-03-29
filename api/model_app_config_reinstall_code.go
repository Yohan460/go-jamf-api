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

// checks if the AppConfigReinstallCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppConfigReinstallCode{}

// AppConfigReinstallCode struct for AppConfigReinstallCode
type AppConfigReinstallCode struct {
	ReinstallCode *string `json:"reinstallCode,omitempty"`
}

// NewAppConfigReinstallCode instantiates a new AppConfigReinstallCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppConfigReinstallCode() *AppConfigReinstallCode {
	this := AppConfigReinstallCode{}
	return &this
}

// NewAppConfigReinstallCodeWithDefaults instantiates a new AppConfigReinstallCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppConfigReinstallCodeWithDefaults() *AppConfigReinstallCode {
	this := AppConfigReinstallCode{}
	return &this
}

// GetReinstallCode returns the ReinstallCode field value if set, zero value otherwise.
func (o *AppConfigReinstallCode) GetReinstallCode() string {
	if o == nil || IsNil(o.ReinstallCode) {
		var ret string
		return ret
	}
	return *o.ReinstallCode
}

// GetReinstallCodeOk returns a tuple with the ReinstallCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppConfigReinstallCode) GetReinstallCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ReinstallCode) {
		return nil, false
	}
	return o.ReinstallCode, true
}

// HasReinstallCode returns a boolean if a field has been set.
func (o *AppConfigReinstallCode) HasReinstallCode() bool {
	if o != nil && !IsNil(o.ReinstallCode) {
		return true
	}

	return false
}

// SetReinstallCode gets a reference to the given string and assigns it to the ReinstallCode field.
func (o *AppConfigReinstallCode) SetReinstallCode(v string) {
	o.ReinstallCode = &v
}

func (o AppConfigReinstallCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppConfigReinstallCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReinstallCode) {
		toSerialize["reinstallCode"] = o.ReinstallCode
	}
	return toSerialize, nil
}

type NullableAppConfigReinstallCode struct {
	value *AppConfigReinstallCode
	isSet bool
}

func (v NullableAppConfigReinstallCode) Get() *AppConfigReinstallCode {
	return v.value
}

func (v *NullableAppConfigReinstallCode) Set(val *AppConfigReinstallCode) {
	v.value = val
	v.isSet = true
}

func (v NullableAppConfigReinstallCode) IsSet() bool {
	return v.isSet
}

func (v *NullableAppConfigReinstallCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppConfigReinstallCode(val *AppConfigReinstallCode) *NullableAppConfigReinstallCode {
	return &NullableAppConfigReinstallCode{value: val, isSet: true}
}

func (v NullableAppConfigReinstallCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppConfigReinstallCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


