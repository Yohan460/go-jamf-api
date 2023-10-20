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

// checks if the EraseDeviceCommandReturnToService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EraseDeviceCommandReturnToService{}

// EraseDeviceCommandReturnToService The configuration settings for Return to Service.
type EraseDeviceCommandReturnToService struct {
	Enabled *bool `json:"enabled,omitempty"`
	// Base64 encoded mdm profile
	MdmProfileData *string `json:"mdmProfileData,omitempty"`
	// Base64 encoded wifi profile
	WifiProfileData *string `json:"wifiProfileData,omitempty"`
}

// NewEraseDeviceCommandReturnToService instantiates a new EraseDeviceCommandReturnToService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEraseDeviceCommandReturnToService() *EraseDeviceCommandReturnToService {
	this := EraseDeviceCommandReturnToService{}
	return &this
}

// NewEraseDeviceCommandReturnToServiceWithDefaults instantiates a new EraseDeviceCommandReturnToService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEraseDeviceCommandReturnToServiceWithDefaults() *EraseDeviceCommandReturnToService {
	this := EraseDeviceCommandReturnToService{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *EraseDeviceCommandReturnToService) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EraseDeviceCommandReturnToService) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *EraseDeviceCommandReturnToService) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *EraseDeviceCommandReturnToService) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMdmProfileData returns the MdmProfileData field value if set, zero value otherwise.
func (o *EraseDeviceCommandReturnToService) GetMdmProfileData() string {
	if o == nil || IsNil(o.MdmProfileData) {
		var ret string
		return ret
	}
	return *o.MdmProfileData
}

// GetMdmProfileDataOk returns a tuple with the MdmProfileData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EraseDeviceCommandReturnToService) GetMdmProfileDataOk() (*string, bool) {
	if o == nil || IsNil(o.MdmProfileData) {
		return nil, false
	}
	return o.MdmProfileData, true
}

// HasMdmProfileData returns a boolean if a field has been set.
func (o *EraseDeviceCommandReturnToService) HasMdmProfileData() bool {
	if o != nil && !IsNil(o.MdmProfileData) {
		return true
	}

	return false
}

// SetMdmProfileData gets a reference to the given string and assigns it to the MdmProfileData field.
func (o *EraseDeviceCommandReturnToService) SetMdmProfileData(v string) {
	o.MdmProfileData = &v
}

// GetWifiProfileData returns the WifiProfileData field value if set, zero value otherwise.
func (o *EraseDeviceCommandReturnToService) GetWifiProfileData() string {
	if o == nil || IsNil(o.WifiProfileData) {
		var ret string
		return ret
	}
	return *o.WifiProfileData
}

// GetWifiProfileDataOk returns a tuple with the WifiProfileData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EraseDeviceCommandReturnToService) GetWifiProfileDataOk() (*string, bool) {
	if o == nil || IsNil(o.WifiProfileData) {
		return nil, false
	}
	return o.WifiProfileData, true
}

// HasWifiProfileData returns a boolean if a field has been set.
func (o *EraseDeviceCommandReturnToService) HasWifiProfileData() bool {
	if o != nil && !IsNil(o.WifiProfileData) {
		return true
	}

	return false
}

// SetWifiProfileData gets a reference to the given string and assigns it to the WifiProfileData field.
func (o *EraseDeviceCommandReturnToService) SetWifiProfileData(v string) {
	o.WifiProfileData = &v
}

func (o EraseDeviceCommandReturnToService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EraseDeviceCommandReturnToService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.MdmProfileData) {
		toSerialize["mdmProfileData"] = o.MdmProfileData
	}
	if !IsNil(o.WifiProfileData) {
		toSerialize["wifiProfileData"] = o.WifiProfileData
	}
	return toSerialize, nil
}

type NullableEraseDeviceCommandReturnToService struct {
	value *EraseDeviceCommandReturnToService
	isSet bool
}

func (v NullableEraseDeviceCommandReturnToService) Get() *EraseDeviceCommandReturnToService {
	return v.value
}

func (v *NullableEraseDeviceCommandReturnToService) Set(val *EraseDeviceCommandReturnToService) {
	v.value = val
	v.isSet = true
}

func (v NullableEraseDeviceCommandReturnToService) IsSet() bool {
	return v.isSet
}

func (v *NullableEraseDeviceCommandReturnToService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEraseDeviceCommandReturnToService(val *EraseDeviceCommandReturnToService) *NullableEraseDeviceCommandReturnToService {
	return &NullableEraseDeviceCommandReturnToService{value: val, isSet: true}
}

func (v NullableEraseDeviceCommandReturnToService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEraseDeviceCommandReturnToService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

