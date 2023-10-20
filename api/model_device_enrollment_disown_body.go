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

// checks if the DeviceEnrollmentDisownBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceEnrollmentDisownBody{}

// DeviceEnrollmentDisownBody struct for DeviceEnrollmentDisownBody
type DeviceEnrollmentDisownBody struct {
	Devices []string `json:"devices,omitempty"`
}

// NewDeviceEnrollmentDisownBody instantiates a new DeviceEnrollmentDisownBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceEnrollmentDisownBody() *DeviceEnrollmentDisownBody {
	this := DeviceEnrollmentDisownBody{}
	return &this
}

// NewDeviceEnrollmentDisownBodyWithDefaults instantiates a new DeviceEnrollmentDisownBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceEnrollmentDisownBodyWithDefaults() *DeviceEnrollmentDisownBody {
	this := DeviceEnrollmentDisownBody{}
	return &this
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *DeviceEnrollmentDisownBody) GetDevices() []string {
	if o == nil || IsNil(o.Devices) {
		var ret []string
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentDisownBody) GetDevicesOk() ([]string, bool) {
	if o == nil || IsNil(o.Devices) {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *DeviceEnrollmentDisownBody) HasDevices() bool {
	if o != nil && !IsNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []string and assigns it to the Devices field.
func (o *DeviceEnrollmentDisownBody) SetDevices(v []string) {
	o.Devices = v
}

func (o DeviceEnrollmentDisownBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceEnrollmentDisownBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	return toSerialize, nil
}

type NullableDeviceEnrollmentDisownBody struct {
	value *DeviceEnrollmentDisownBody
	isSet bool
}

func (v NullableDeviceEnrollmentDisownBody) Get() *DeviceEnrollmentDisownBody {
	return v.value
}

func (v *NullableDeviceEnrollmentDisownBody) Set(val *DeviceEnrollmentDisownBody) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceEnrollmentDisownBody) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceEnrollmentDisownBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceEnrollmentDisownBody(val *DeviceEnrollmentDisownBody) *NullableDeviceEnrollmentDisownBody {
	return &NullableDeviceEnrollmentDisownBody{value: val, isSet: true}
}

func (v NullableDeviceEnrollmentDisownBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceEnrollmentDisownBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


