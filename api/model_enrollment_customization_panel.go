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

// checks if the EnrollmentCustomizationPanel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnrollmentCustomizationPanel{}

// EnrollmentCustomizationPanel struct for EnrollmentCustomizationPanel
type EnrollmentCustomizationPanel struct {
	DisplayName string `json:"displayName"`
	Rank int32 `json:"rank"`
}

// NewEnrollmentCustomizationPanel instantiates a new EnrollmentCustomizationPanel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrollmentCustomizationPanel(displayName string, rank int32) *EnrollmentCustomizationPanel {
	this := EnrollmentCustomizationPanel{}
	this.DisplayName = displayName
	this.Rank = rank
	return &this
}

// NewEnrollmentCustomizationPanelWithDefaults instantiates a new EnrollmentCustomizationPanel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrollmentCustomizationPanelWithDefaults() *EnrollmentCustomizationPanel {
	this := EnrollmentCustomizationPanel{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *EnrollmentCustomizationPanel) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationPanel) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *EnrollmentCustomizationPanel) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetRank returns the Rank field value
func (o *EnrollmentCustomizationPanel) GetRank() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rank
}

// GetRankOk returns a tuple with the Rank field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationPanel) GetRankOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rank, true
}

// SetRank sets field value
func (o *EnrollmentCustomizationPanel) SetRank(v int32) {
	o.Rank = v
}

func (o EnrollmentCustomizationPanel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnrollmentCustomizationPanel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["rank"] = o.Rank
	return toSerialize, nil
}

type NullableEnrollmentCustomizationPanel struct {
	value *EnrollmentCustomizationPanel
	isSet bool
}

func (v NullableEnrollmentCustomizationPanel) Get() *EnrollmentCustomizationPanel {
	return v.value
}

func (v *NullableEnrollmentCustomizationPanel) Set(val *EnrollmentCustomizationPanel) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollmentCustomizationPanel) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollmentCustomizationPanel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollmentCustomizationPanel(val *EnrollmentCustomizationPanel) *NullableEnrollmentCustomizationPanel {
	return &NullableEnrollmentCustomizationPanel{value: val, isSet: true}
}

func (v NullableEnrollmentCustomizationPanel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrollmentCustomizationPanel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


