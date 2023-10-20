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

// checks if the PlanDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanDevice{}

// PlanDevice struct for PlanDevice
type PlanDevice struct {
	DeviceId *string `json:"deviceId,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
	Href *string `json:"href,omitempty"`
}

// NewPlanDevice instantiates a new PlanDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanDevice() *PlanDevice {
	this := PlanDevice{}
	return &this
}

// NewPlanDeviceWithDefaults instantiates a new PlanDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanDeviceWithDefaults() *PlanDevice {
	this := PlanDevice{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *PlanDevice) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanDevice) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *PlanDevice) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *PlanDevice) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *PlanDevice) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanDevice) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *PlanDevice) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *PlanDevice) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *PlanDevice) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanDevice) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *PlanDevice) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *PlanDevice) SetHref(v string) {
	o.Href = &v
}

func (o PlanDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlanDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !IsNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return toSerialize, nil
}

type NullablePlanDevice struct {
	value *PlanDevice
	isSet bool
}

func (v NullablePlanDevice) Get() *PlanDevice {
	return v.value
}

func (v *NullablePlanDevice) Set(val *PlanDevice) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanDevice) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanDevice(val *PlanDevice) *NullablePlanDevice {
	return &NullablePlanDevice{value: val, isSet: true}
}

func (v NullablePlanDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

