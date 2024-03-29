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

// checks if the PlanDeviceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanDeviceResponse{}

// PlanDeviceResponse struct for PlanDeviceResponse
type PlanDeviceResponse struct {
	Device *PlanDevice `json:"device,omitempty"`
	PlanId *string `json:"planId,omitempty"`
	Href *string `json:"href,omitempty"`
}

// NewPlanDeviceResponse instantiates a new PlanDeviceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanDeviceResponse() *PlanDeviceResponse {
	this := PlanDeviceResponse{}
	return &this
}

// NewPlanDeviceResponseWithDefaults instantiates a new PlanDeviceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanDeviceResponseWithDefaults() *PlanDeviceResponse {
	this := PlanDeviceResponse{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *PlanDeviceResponse) GetDevice() PlanDevice {
	if o == nil || IsNil(o.Device) {
		var ret PlanDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanDeviceResponse) GetDeviceOk() (*PlanDevice, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *PlanDeviceResponse) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given PlanDevice and assigns it to the Device field.
func (o *PlanDeviceResponse) SetDevice(v PlanDevice) {
	o.Device = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *PlanDeviceResponse) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanDeviceResponse) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *PlanDeviceResponse) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *PlanDeviceResponse) SetPlanId(v string) {
	o.PlanId = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *PlanDeviceResponse) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanDeviceResponse) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *PlanDeviceResponse) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *PlanDeviceResponse) SetHref(v string) {
	o.Href = &v
}

func (o PlanDeviceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlanDeviceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.PlanId) {
		toSerialize["planId"] = o.PlanId
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return toSerialize, nil
}

type NullablePlanDeviceResponse struct {
	value *PlanDeviceResponse
	isSet bool
}

func (v NullablePlanDeviceResponse) Get() *PlanDeviceResponse {
	return v.value
}

func (v *NullablePlanDeviceResponse) Set(val *PlanDeviceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanDeviceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanDeviceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanDeviceResponse(val *PlanDeviceResponse) *NullablePlanDeviceResponse {
	return &NullablePlanDeviceResponse{value: val, isSet: true}
}

func (v NullablePlanDeviceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanDeviceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


