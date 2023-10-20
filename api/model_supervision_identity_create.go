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

// checks if the SupervisionIdentityCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupervisionIdentityCreate{}

// SupervisionIdentityCreate struct for SupervisionIdentityCreate
type SupervisionIdentityCreate struct {
	DisplayName string `json:"displayName"`
	Password string `json:"password"`
}

// NewSupervisionIdentityCreate instantiates a new SupervisionIdentityCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupervisionIdentityCreate(displayName string, password string) *SupervisionIdentityCreate {
	this := SupervisionIdentityCreate{}
	this.DisplayName = displayName
	this.Password = password
	return &this
}

// NewSupervisionIdentityCreateWithDefaults instantiates a new SupervisionIdentityCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupervisionIdentityCreateWithDefaults() *SupervisionIdentityCreate {
	this := SupervisionIdentityCreate{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *SupervisionIdentityCreate) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *SupervisionIdentityCreate) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *SupervisionIdentityCreate) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetPassword returns the Password field value
func (o *SupervisionIdentityCreate) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SupervisionIdentityCreate) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *SupervisionIdentityCreate) SetPassword(v string) {
	o.Password = v
}

func (o SupervisionIdentityCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupervisionIdentityCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

type NullableSupervisionIdentityCreate struct {
	value *SupervisionIdentityCreate
	isSet bool
}

func (v NullableSupervisionIdentityCreate) Get() *SupervisionIdentityCreate {
	return v.value
}

func (v *NullableSupervisionIdentityCreate) Set(val *SupervisionIdentityCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSupervisionIdentityCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSupervisionIdentityCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupervisionIdentityCreate(val *SupervisionIdentityCreate) *NullableSupervisionIdentityCreate {
	return &NullableSupervisionIdentityCreate{value: val, isSet: true}
}

func (v NullableSupervisionIdentityCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupervisionIdentityCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


