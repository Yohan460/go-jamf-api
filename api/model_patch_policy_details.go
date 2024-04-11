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

// checks if the PatchPolicyDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchPolicyDetails{}

// PatchPolicyDetails struct for PatchPolicyDetails
type PatchPolicyDetails struct {
	TotalCount *int64 `json:"totalCount,omitempty"`
	Results []PatchPolicyDetail `json:"results,omitempty"`
}

// NewPatchPolicyDetails instantiates a new PatchPolicyDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchPolicyDetails() *PatchPolicyDetails {
	this := PatchPolicyDetails{}
	return &this
}

// NewPatchPolicyDetailsWithDefaults instantiates a new PatchPolicyDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchPolicyDetailsWithDefaults() *PatchPolicyDetails {
	this := PatchPolicyDetails{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PatchPolicyDetails) GetTotalCount() int64 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPolicyDetails) GetTotalCountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PatchPolicyDetails) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *PatchPolicyDetails) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PatchPolicyDetails) GetResults() []PatchPolicyDetail {
	if o == nil || IsNil(o.Results) {
		var ret []PatchPolicyDetail
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPolicyDetails) GetResultsOk() ([]PatchPolicyDetail, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PatchPolicyDetails) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []PatchPolicyDetail and assigns it to the Results field.
func (o *PatchPolicyDetails) SetResults(v []PatchPolicyDetail) {
	o.Results = v
}

func (o PatchPolicyDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchPolicyDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullablePatchPolicyDetails struct {
	value *PatchPolicyDetails
	isSet bool
}

func (v NullablePatchPolicyDetails) Get() *PatchPolicyDetails {
	return v.value
}

func (v *NullablePatchPolicyDetails) Set(val *PatchPolicyDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchPolicyDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchPolicyDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchPolicyDetails(val *PatchPolicyDetails) *NullablePatchPolicyDetails {
	return &NullablePatchPolicyDetails{value: val, isSet: true}
}

func (v NullablePatchPolicyDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchPolicyDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


