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

// ExternalRecipient struct for ExternalRecipient
type ExternalRecipient struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

// NewExternalRecipient instantiates a new ExternalRecipient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalRecipient(name string, email string) *ExternalRecipient {
	this := ExternalRecipient{}
	this.Name = name
	this.Email = email
	return &this
}

// NewExternalRecipientWithDefaults instantiates a new ExternalRecipient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalRecipientWithDefaults() *ExternalRecipient {
	this := ExternalRecipient{}
	return &this
}

// GetName returns the Name field value
func (o *ExternalRecipient) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExternalRecipient) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExternalRecipient) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *ExternalRecipient) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ExternalRecipient) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ExternalRecipient) SetEmail(v string) {
	o.Email = v
}

func (o ExternalRecipient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableExternalRecipient struct {
	value *ExternalRecipient
	isSet bool
}

func (v NullableExternalRecipient) Get() *ExternalRecipient {
	return v.value
}

func (v *NullableExternalRecipient) Set(val *ExternalRecipient) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalRecipient) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalRecipient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalRecipient(val *ExternalRecipient) *NullableExternalRecipient {
	return &NullableExternalRecipient{value: val, isSet: true}
}

func (v NullableExternalRecipient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalRecipient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

