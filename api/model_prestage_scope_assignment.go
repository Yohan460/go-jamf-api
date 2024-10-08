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

// checks if the PrestageScopeAssignment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrestageScopeAssignment{}

// PrestageScopeAssignment struct for PrestageScopeAssignment
type PrestageScopeAssignment struct {
	SerialNumber *string `json:"serialNumber,omitempty"`
	AssignmentEpoch *int64 `json:"assignmentEpoch,omitempty"`
	UserAssigned *string `json:"userAssigned,omitempty"`
}

// NewPrestageScopeAssignment instantiates a new PrestageScopeAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrestageScopeAssignment() *PrestageScopeAssignment {
	this := PrestageScopeAssignment{}
	return &this
}

// NewPrestageScopeAssignmentWithDefaults instantiates a new PrestageScopeAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrestageScopeAssignmentWithDefaults() *PrestageScopeAssignment {
	this := PrestageScopeAssignment{}
	return &this
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *PrestageScopeAssignment) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrestageScopeAssignment) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *PrestageScopeAssignment) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *PrestageScopeAssignment) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetAssignmentEpoch returns the AssignmentEpoch field value if set, zero value otherwise.
func (o *PrestageScopeAssignment) GetAssignmentEpoch() int64 {
	if o == nil || IsNil(o.AssignmentEpoch) {
		var ret int64
		return ret
	}
	return *o.AssignmentEpoch
}

// GetAssignmentEpochOk returns a tuple with the AssignmentEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrestageScopeAssignment) GetAssignmentEpochOk() (*int64, bool) {
	if o == nil || IsNil(o.AssignmentEpoch) {
		return nil, false
	}
	return o.AssignmentEpoch, true
}

// HasAssignmentEpoch returns a boolean if a field has been set.
func (o *PrestageScopeAssignment) HasAssignmentEpoch() bool {
	if o != nil && !IsNil(o.AssignmentEpoch) {
		return true
	}

	return false
}

// SetAssignmentEpoch gets a reference to the given int64 and assigns it to the AssignmentEpoch field.
func (o *PrestageScopeAssignment) SetAssignmentEpoch(v int64) {
	o.AssignmentEpoch = &v
}

// GetUserAssigned returns the UserAssigned field value if set, zero value otherwise.
func (o *PrestageScopeAssignment) GetUserAssigned() string {
	if o == nil || IsNil(o.UserAssigned) {
		var ret string
		return ret
	}
	return *o.UserAssigned
}

// GetUserAssignedOk returns a tuple with the UserAssigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrestageScopeAssignment) GetUserAssignedOk() (*string, bool) {
	if o == nil || IsNil(o.UserAssigned) {
		return nil, false
	}
	return o.UserAssigned, true
}

// HasUserAssigned returns a boolean if a field has been set.
func (o *PrestageScopeAssignment) HasUserAssigned() bool {
	if o != nil && !IsNil(o.UserAssigned) {
		return true
	}

	return false
}

// SetUserAssigned gets a reference to the given string and assigns it to the UserAssigned field.
func (o *PrestageScopeAssignment) SetUserAssigned(v string) {
	o.UserAssigned = &v
}

func (o PrestageScopeAssignment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrestageScopeAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if !IsNil(o.AssignmentEpoch) {
		toSerialize["assignmentEpoch"] = o.AssignmentEpoch
	}
	if !IsNil(o.UserAssigned) {
		toSerialize["userAssigned"] = o.UserAssigned
	}
	return toSerialize, nil
}

type NullablePrestageScopeAssignment struct {
	value *PrestageScopeAssignment
	isSet bool
}

func (v NullablePrestageScopeAssignment) Get() *PrestageScopeAssignment {
	return v.value
}

func (v *NullablePrestageScopeAssignment) Set(val *PrestageScopeAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullablePrestageScopeAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullablePrestageScopeAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrestageScopeAssignment(val *PrestageScopeAssignment) *NullablePrestageScopeAssignment {
	return &NullablePrestageScopeAssignment{value: val, isSet: true}
}

func (v NullablePrestageScopeAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrestageScopeAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


