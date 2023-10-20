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

// checks if the PrestageScopeResponseV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrestageScopeResponseV2{}

// PrestageScopeResponseV2 struct for PrestageScopeResponseV2
type PrestageScopeResponseV2 struct {
	PrestageId *string `json:"prestageId,omitempty"`
	Assignments []PrestageScopeAssignmentV2 `json:"assignments,omitempty"`
	VersionLock *int32 `json:"versionLock,omitempty"`
}

// NewPrestageScopeResponseV2 instantiates a new PrestageScopeResponseV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrestageScopeResponseV2() *PrestageScopeResponseV2 {
	this := PrestageScopeResponseV2{}
	return &this
}

// NewPrestageScopeResponseV2WithDefaults instantiates a new PrestageScopeResponseV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrestageScopeResponseV2WithDefaults() *PrestageScopeResponseV2 {
	this := PrestageScopeResponseV2{}
	return &this
}

// GetPrestageId returns the PrestageId field value if set, zero value otherwise.
func (o *PrestageScopeResponseV2) GetPrestageId() string {
	if o == nil || IsNil(o.PrestageId) {
		var ret string
		return ret
	}
	return *o.PrestageId
}

// GetPrestageIdOk returns a tuple with the PrestageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrestageScopeResponseV2) GetPrestageIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrestageId) {
		return nil, false
	}
	return o.PrestageId, true
}

// HasPrestageId returns a boolean if a field has been set.
func (o *PrestageScopeResponseV2) HasPrestageId() bool {
	if o != nil && !IsNil(o.PrestageId) {
		return true
	}

	return false
}

// SetPrestageId gets a reference to the given string and assigns it to the PrestageId field.
func (o *PrestageScopeResponseV2) SetPrestageId(v string) {
	o.PrestageId = &v
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *PrestageScopeResponseV2) GetAssignments() []PrestageScopeAssignmentV2 {
	if o == nil || IsNil(o.Assignments) {
		var ret []PrestageScopeAssignmentV2
		return ret
	}
	return o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrestageScopeResponseV2) GetAssignmentsOk() ([]PrestageScopeAssignmentV2, bool) {
	if o == nil || IsNil(o.Assignments) {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *PrestageScopeResponseV2) HasAssignments() bool {
	if o != nil && !IsNil(o.Assignments) {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []PrestageScopeAssignmentV2 and assigns it to the Assignments field.
func (o *PrestageScopeResponseV2) SetAssignments(v []PrestageScopeAssignmentV2) {
	o.Assignments = v
}

// GetVersionLock returns the VersionLock field value if set, zero value otherwise.
func (o *PrestageScopeResponseV2) GetVersionLock() int32 {
	if o == nil || IsNil(o.VersionLock) {
		var ret int32
		return ret
	}
	return *o.VersionLock
}

// GetVersionLockOk returns a tuple with the VersionLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrestageScopeResponseV2) GetVersionLockOk() (*int32, bool) {
	if o == nil || IsNil(o.VersionLock) {
		return nil, false
	}
	return o.VersionLock, true
}

// HasVersionLock returns a boolean if a field has been set.
func (o *PrestageScopeResponseV2) HasVersionLock() bool {
	if o != nil && !IsNil(o.VersionLock) {
		return true
	}

	return false
}

// SetVersionLock gets a reference to the given int32 and assigns it to the VersionLock field.
func (o *PrestageScopeResponseV2) SetVersionLock(v int32) {
	o.VersionLock = &v
}

func (o PrestageScopeResponseV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrestageScopeResponseV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrestageId) {
		toSerialize["prestageId"] = o.PrestageId
	}
	if !IsNil(o.Assignments) {
		toSerialize["assignments"] = o.Assignments
	}
	if !IsNil(o.VersionLock) {
		toSerialize["versionLock"] = o.VersionLock
	}
	return toSerialize, nil
}

type NullablePrestageScopeResponseV2 struct {
	value *PrestageScopeResponseV2
	isSet bool
}

func (v NullablePrestageScopeResponseV2) Get() *PrestageScopeResponseV2 {
	return v.value
}

func (v *NullablePrestageScopeResponseV2) Set(val *PrestageScopeResponseV2) {
	v.value = val
	v.isSet = true
}

func (v NullablePrestageScopeResponseV2) IsSet() bool {
	return v.isSet
}

func (v *NullablePrestageScopeResponseV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrestageScopeResponseV2(val *PrestageScopeResponseV2) *NullablePrestageScopeResponseV2 {
	return &NullablePrestageScopeResponseV2{value: val, isSet: true}
}

func (v NullablePrestageScopeResponseV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrestageScopeResponseV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


