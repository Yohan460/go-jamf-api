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

// checks if the DashboardMetric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardMetric{}

// DashboardMetric struct for DashboardMetric
type DashboardMetric struct {
	// Usually a number associated with the tag; i.e. 23 Pending Computers
	Value *string `json:"value,omitempty"`
	// Logical to decide whether metric should be enabled or disabled; i.e. Policy can be at Retrying-Disabled status
	Enabled *bool `json:"enabled,omitempty"`
	Tag *string `json:"tag,omitempty"`
}

// NewDashboardMetric instantiates a new DashboardMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardMetric() *DashboardMetric {
	this := DashboardMetric{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// NewDashboardMetricWithDefaults instantiates a new DashboardMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardMetricWithDefaults() *DashboardMetric {
	this := DashboardMetric{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DashboardMetric) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardMetric) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DashboardMetric) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DashboardMetric) SetValue(v string) {
	o.Value = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DashboardMetric) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardMetric) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DashboardMetric) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DashboardMetric) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *DashboardMetric) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardMetric) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *DashboardMetric) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *DashboardMetric) SetTag(v string) {
	o.Tag = &v
}

func (o DashboardMetric) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardMetric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return toSerialize, nil
}

type NullableDashboardMetric struct {
	value *DashboardMetric
	isSet bool
}

func (v NullableDashboardMetric) Get() *DashboardMetric {
	return v.value
}

func (v *NullableDashboardMetric) Set(val *DashboardMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardMetric(val *DashboardMetric) *NullableDashboardMetric {
	return &NullableDashboardMetric{value: val, isSet: true}
}

func (v NullableDashboardMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


