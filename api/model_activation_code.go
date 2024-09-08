/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ActivationCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivationCode{}

// ActivationCode struct for ActivationCode
type ActivationCode struct {
	// Activation Code for Jamf Pro. Hyphens are optional.
	ActivationCode string `json:"activationCode"`
}

type _ActivationCode ActivationCode

// NewActivationCode instantiates a new ActivationCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivationCode(activationCode string) *ActivationCode {
	this := ActivationCode{}
	this.ActivationCode = activationCode
	return &this
}

// NewActivationCodeWithDefaults instantiates a new ActivationCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivationCodeWithDefaults() *ActivationCode {
	this := ActivationCode{}
	return &this
}

// GetActivationCode returns the ActivationCode field value
func (o *ActivationCode) GetActivationCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActivationCode
}

// GetActivationCodeOk returns a tuple with the ActivationCode field value
// and a boolean to check if the value has been set.
func (o *ActivationCode) GetActivationCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActivationCode, true
}

// SetActivationCode sets field value
func (o *ActivationCode) SetActivationCode(v string) {
	o.ActivationCode = v
}

func (o ActivationCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivationCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["activationCode"] = o.ActivationCode
	return toSerialize, nil
}

func (o *ActivationCode) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"activationCode",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varActivationCode := _ActivationCode{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActivationCode)

	if err != nil {
		return err
	}

	*o = ActivationCode(varActivationCode)

	return err
}

type NullableActivationCode struct {
	value *ActivationCode
	isSet bool
}

func (v NullableActivationCode) Get() *ActivationCode {
	return v.value
}

func (v *NullableActivationCode) Set(val *ActivationCode) {
	v.value = val
	v.isSet = true
}

func (v NullableActivationCode) IsSet() bool {
	return v.isSet
}

func (v *NullableActivationCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivationCode(val *ActivationCode) *NullableActivationCode {
	return &NullableActivationCode{value: val, isSet: true}
}

func (v NullableActivationCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivationCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

