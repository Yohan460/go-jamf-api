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

// checks if the SoftwareUpdateSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoftwareUpdateSettings{}

// SoftwareUpdateSettings struct for SoftwareUpdateSettings
type SoftwareUpdateSettings struct {
	RecommendationCadence *string `json:"recommendationCadence,omitempty"`
}

// NewSoftwareUpdateSettings instantiates a new SoftwareUpdateSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareUpdateSettings() *SoftwareUpdateSettings {
	this := SoftwareUpdateSettings{}
	return &this
}

// NewSoftwareUpdateSettingsWithDefaults instantiates a new SoftwareUpdateSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareUpdateSettingsWithDefaults() *SoftwareUpdateSettings {
	this := SoftwareUpdateSettings{}
	return &this
}

// GetRecommendationCadence returns the RecommendationCadence field value if set, zero value otherwise.
func (o *SoftwareUpdateSettings) GetRecommendationCadence() string {
	if o == nil || IsNil(o.RecommendationCadence) {
		var ret string
		return ret
	}
	return *o.RecommendationCadence
}

// GetRecommendationCadenceOk returns a tuple with the RecommendationCadence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareUpdateSettings) GetRecommendationCadenceOk() (*string, bool) {
	if o == nil || IsNil(o.RecommendationCadence) {
		return nil, false
	}
	return o.RecommendationCadence, true
}

// HasRecommendationCadence returns a boolean if a field has been set.
func (o *SoftwareUpdateSettings) HasRecommendationCadence() bool {
	if o != nil && !IsNil(o.RecommendationCadence) {
		return true
	}

	return false
}

// SetRecommendationCadence gets a reference to the given string and assigns it to the RecommendationCadence field.
func (o *SoftwareUpdateSettings) SetRecommendationCadence(v string) {
	o.RecommendationCadence = &v
}

func (o SoftwareUpdateSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoftwareUpdateSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecommendationCadence) {
		toSerialize["recommendationCadence"] = o.RecommendationCadence
	}
	return toSerialize, nil
}

type NullableSoftwareUpdateSettings struct {
	value *SoftwareUpdateSettings
	isSet bool
}

func (v NullableSoftwareUpdateSettings) Get() *SoftwareUpdateSettings {
	return v.value
}

func (v *NullableSoftwareUpdateSettings) Set(val *SoftwareUpdateSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareUpdateSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareUpdateSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareUpdateSettings(val *SoftwareUpdateSettings) *NullableSoftwareUpdateSettings {
	return &NullableSoftwareUpdateSettings{value: val, isSet: true}
}

func (v NullableSoftwareUpdateSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareUpdateSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


