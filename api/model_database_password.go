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

// DatabasePassword struct for DatabasePassword
type DatabasePassword struct {
	Password string `json:"password"`
}

// NewDatabasePassword instantiates a new DatabasePassword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabasePassword(password string) *DatabasePassword {
	this := DatabasePassword{}
	this.Password = password
	return &this
}

// NewDatabasePasswordWithDefaults instantiates a new DatabasePassword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabasePasswordWithDefaults() *DatabasePassword {
	this := DatabasePassword{}
	return &this
}

// GetPassword returns the Password field value
func (o *DatabasePassword) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *DatabasePassword) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *DatabasePassword) SetPassword(v string) {
	o.Password = v
}

func (o DatabasePassword) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableDatabasePassword struct {
	value *DatabasePassword
	isSet bool
}

func (v NullableDatabasePassword) Get() *DatabasePassword {
	return v.value
}

func (v *NullableDatabasePassword) Set(val *DatabasePassword) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabasePassword) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabasePassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabasePassword(val *DatabasePassword) *NullableDatabasePassword {
	return &NullableDatabasePassword{value: val, isSet: true}
}

func (v NullableDatabasePassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabasePassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


