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

// DeviceEnrollmentDeviceSearchResults struct for DeviceEnrollmentDeviceSearchResults
type DeviceEnrollmentDeviceSearchResults struct {
	TotalCount *int32 `json:"totalCount,omitempty"`
	Results []DeviceEnrollmentDevice `json:"results,omitempty"`
}

// NewDeviceEnrollmentDeviceSearchResults instantiates a new DeviceEnrollmentDeviceSearchResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceEnrollmentDeviceSearchResults() *DeviceEnrollmentDeviceSearchResults {
	this := DeviceEnrollmentDeviceSearchResults{}
	return &this
}

// NewDeviceEnrollmentDeviceSearchResultsWithDefaults instantiates a new DeviceEnrollmentDeviceSearchResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceEnrollmentDeviceSearchResultsWithDefaults() *DeviceEnrollmentDeviceSearchResults {
	this := DeviceEnrollmentDeviceSearchResults{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *DeviceEnrollmentDeviceSearchResults) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentDeviceSearchResults) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *DeviceEnrollmentDeviceSearchResults) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *DeviceEnrollmentDeviceSearchResults) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *DeviceEnrollmentDeviceSearchResults) GetResults() []DeviceEnrollmentDevice {
	if o == nil || o.Results == nil {
		var ret []DeviceEnrollmentDevice
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentDeviceSearchResults) GetResultsOk() ([]DeviceEnrollmentDevice, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *DeviceEnrollmentDeviceSearchResults) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DeviceEnrollmentDevice and assigns it to the Results field.
func (o *DeviceEnrollmentDeviceSearchResults) SetResults(v []DeviceEnrollmentDevice) {
	o.Results = v
}

func (o DeviceEnrollmentDeviceSearchResults) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceEnrollmentDeviceSearchResults struct {
	value *DeviceEnrollmentDeviceSearchResults
	isSet bool
}

func (v NullableDeviceEnrollmentDeviceSearchResults) Get() *DeviceEnrollmentDeviceSearchResults {
	return v.value
}

func (v *NullableDeviceEnrollmentDeviceSearchResults) Set(val *DeviceEnrollmentDeviceSearchResults) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceEnrollmentDeviceSearchResults) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceEnrollmentDeviceSearchResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceEnrollmentDeviceSearchResults(val *DeviceEnrollmentDeviceSearchResults) *NullableDeviceEnrollmentDeviceSearchResults {
	return &NullableDeviceEnrollmentDeviceSearchResults{value: val, isSet: true}
}

func (v NullableDeviceEnrollmentDeviceSearchResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceEnrollmentDeviceSearchResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


