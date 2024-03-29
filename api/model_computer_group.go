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

// checks if the ComputerGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComputerGroup{}

// ComputerGroup struct for ComputerGroup
type ComputerGroup struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	SmartGroup *bool `json:"smartGroup,omitempty"`
}

// NewComputerGroup instantiates a new ComputerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputerGroup() *ComputerGroup {
	this := ComputerGroup{}
	return &this
}

// NewComputerGroupWithDefaults instantiates a new ComputerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputerGroupWithDefaults() *ComputerGroup {
	this := ComputerGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ComputerGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ComputerGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ComputerGroup) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComputerGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComputerGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComputerGroup) SetName(v string) {
	o.Name = &v
}

// GetSmartGroup returns the SmartGroup field value if set, zero value otherwise.
func (o *ComputerGroup) GetSmartGroup() bool {
	if o == nil || IsNil(o.SmartGroup) {
		var ret bool
		return ret
	}
	return *o.SmartGroup
}

// GetSmartGroupOk returns a tuple with the SmartGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerGroup) GetSmartGroupOk() (*bool, bool) {
	if o == nil || IsNil(o.SmartGroup) {
		return nil, false
	}
	return o.SmartGroup, true
}

// HasSmartGroup returns a boolean if a field has been set.
func (o *ComputerGroup) HasSmartGroup() bool {
	if o != nil && !IsNil(o.SmartGroup) {
		return true
	}

	return false
}

// SetSmartGroup gets a reference to the given bool and assigns it to the SmartGroup field.
func (o *ComputerGroup) SetSmartGroup(v bool) {
	o.SmartGroup = &v
}

func (o ComputerGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComputerGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SmartGroup) {
		toSerialize["smartGroup"] = o.SmartGroup
	}
	return toSerialize, nil
}

type NullableComputerGroup struct {
	value *ComputerGroup
	isSet bool
}

func (v NullableComputerGroup) Get() *ComputerGroup {
	return v.value
}

func (v *NullableComputerGroup) Set(val *ComputerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerGroup(val *ComputerGroup) *NullableComputerGroup {
	return &NullableComputerGroup{value: val, isSet: true}
}

func (v NullableComputerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


