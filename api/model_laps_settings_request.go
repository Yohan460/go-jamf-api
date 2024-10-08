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

// checks if the LapsSettingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LapsSettingsRequest{}

// LapsSettingsRequest struct for LapsSettingsRequest
type LapsSettingsRequest struct {
	// When enabled, all appropriate computers will have the SetAutoAdminPassword command sent to them automatically.
	AutoDeployEnabled bool `json:"autoDeployEnabled"`
	// The amount of time in seconds that the local admin password will be rotated after viewing.
	PasswordRotationTime int64 `json:"passwordRotationTime"`
	// The amount of time in seconds that the local admin password will be rotated automatically if it is never viewed.
	AutoExpirationTime int64 `json:"autoExpirationTime"`
}

type _LapsSettingsRequest LapsSettingsRequest

// NewLapsSettingsRequest instantiates a new LapsSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLapsSettingsRequest(autoDeployEnabled bool, passwordRotationTime int64, autoExpirationTime int64) *LapsSettingsRequest {
	this := LapsSettingsRequest{}
	this.AutoDeployEnabled = autoDeployEnabled
	this.PasswordRotationTime = passwordRotationTime
	this.AutoExpirationTime = autoExpirationTime
	return &this
}

// NewLapsSettingsRequestWithDefaults instantiates a new LapsSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLapsSettingsRequestWithDefaults() *LapsSettingsRequest {
	this := LapsSettingsRequest{}
	return &this
}

// GetAutoDeployEnabled returns the AutoDeployEnabled field value
func (o *LapsSettingsRequest) GetAutoDeployEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoDeployEnabled
}

// GetAutoDeployEnabledOk returns a tuple with the AutoDeployEnabled field value
// and a boolean to check if the value has been set.
func (o *LapsSettingsRequest) GetAutoDeployEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoDeployEnabled, true
}

// SetAutoDeployEnabled sets field value
func (o *LapsSettingsRequest) SetAutoDeployEnabled(v bool) {
	o.AutoDeployEnabled = v
}

// GetPasswordRotationTime returns the PasswordRotationTime field value
func (o *LapsSettingsRequest) GetPasswordRotationTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PasswordRotationTime
}

// GetPasswordRotationTimeOk returns a tuple with the PasswordRotationTime field value
// and a boolean to check if the value has been set.
func (o *LapsSettingsRequest) GetPasswordRotationTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordRotationTime, true
}

// SetPasswordRotationTime sets field value
func (o *LapsSettingsRequest) SetPasswordRotationTime(v int64) {
	o.PasswordRotationTime = v
}

// GetAutoExpirationTime returns the AutoExpirationTime field value
func (o *LapsSettingsRequest) GetAutoExpirationTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AutoExpirationTime
}

// GetAutoExpirationTimeOk returns a tuple with the AutoExpirationTime field value
// and a boolean to check if the value has been set.
func (o *LapsSettingsRequest) GetAutoExpirationTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoExpirationTime, true
}

// SetAutoExpirationTime sets field value
func (o *LapsSettingsRequest) SetAutoExpirationTime(v int64) {
	o.AutoExpirationTime = v
}

func (o LapsSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LapsSettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["autoDeployEnabled"] = o.AutoDeployEnabled
	toSerialize["passwordRotationTime"] = o.PasswordRotationTime
	toSerialize["autoExpirationTime"] = o.AutoExpirationTime
	return toSerialize, nil
}

func (o *LapsSettingsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"autoDeployEnabled",
		"passwordRotationTime",
		"autoExpirationTime",
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

	varLapsSettingsRequest := _LapsSettingsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLapsSettingsRequest)

	if err != nil {
		return err
	}

	*o = LapsSettingsRequest(varLapsSettingsRequest)

	return err
}

type NullableLapsSettingsRequest struct {
	value *LapsSettingsRequest
	isSet bool
}

func (v NullableLapsSettingsRequest) Get() *LapsSettingsRequest {
	return v.value
}

func (v *NullableLapsSettingsRequest) Set(val *LapsSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLapsSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLapsSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLapsSettingsRequest(val *LapsSettingsRequest) *NullableLapsSettingsRequest {
	return &NullableLapsSettingsRequest{value: val, isSet: true}
}

func (v NullableLapsSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLapsSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


