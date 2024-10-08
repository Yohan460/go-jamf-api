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

// checks if the ComputerPartitionEncryption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComputerPartitionEncryption{}

// ComputerPartitionEncryption struct for ComputerPartitionEncryption
type ComputerPartitionEncryption struct {
	PartitionName *string `json:"partitionName,omitempty"`
	PartitionFileVault2State *ComputerPartitionFileVault2State `json:"partitionFileVault2State,omitempty"`
	PartitionFileVault2Percent *int64 `json:"partitionFileVault2Percent,omitempty"`
}

// NewComputerPartitionEncryption instantiates a new ComputerPartitionEncryption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputerPartitionEncryption() *ComputerPartitionEncryption {
	this := ComputerPartitionEncryption{}
	return &this
}

// NewComputerPartitionEncryptionWithDefaults instantiates a new ComputerPartitionEncryption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputerPartitionEncryptionWithDefaults() *ComputerPartitionEncryption {
	this := ComputerPartitionEncryption{}
	return &this
}

// GetPartitionName returns the PartitionName field value if set, zero value otherwise.
func (o *ComputerPartitionEncryption) GetPartitionName() string {
	if o == nil || IsNil(o.PartitionName) {
		var ret string
		return ret
	}
	return *o.PartitionName
}

// GetPartitionNameOk returns a tuple with the PartitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPartitionEncryption) GetPartitionNameOk() (*string, bool) {
	if o == nil || IsNil(o.PartitionName) {
		return nil, false
	}
	return o.PartitionName, true
}

// HasPartitionName returns a boolean if a field has been set.
func (o *ComputerPartitionEncryption) HasPartitionName() bool {
	if o != nil && !IsNil(o.PartitionName) {
		return true
	}

	return false
}

// SetPartitionName gets a reference to the given string and assigns it to the PartitionName field.
func (o *ComputerPartitionEncryption) SetPartitionName(v string) {
	o.PartitionName = &v
}

// GetPartitionFileVault2State returns the PartitionFileVault2State field value if set, zero value otherwise.
func (o *ComputerPartitionEncryption) GetPartitionFileVault2State() ComputerPartitionFileVault2State {
	if o == nil || IsNil(o.PartitionFileVault2State) {
		var ret ComputerPartitionFileVault2State
		return ret
	}
	return *o.PartitionFileVault2State
}

// GetPartitionFileVault2StateOk returns a tuple with the PartitionFileVault2State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPartitionEncryption) GetPartitionFileVault2StateOk() (*ComputerPartitionFileVault2State, bool) {
	if o == nil || IsNil(o.PartitionFileVault2State) {
		return nil, false
	}
	return o.PartitionFileVault2State, true
}

// HasPartitionFileVault2State returns a boolean if a field has been set.
func (o *ComputerPartitionEncryption) HasPartitionFileVault2State() bool {
	if o != nil && !IsNil(o.PartitionFileVault2State) {
		return true
	}

	return false
}

// SetPartitionFileVault2State gets a reference to the given ComputerPartitionFileVault2State and assigns it to the PartitionFileVault2State field.
func (o *ComputerPartitionEncryption) SetPartitionFileVault2State(v ComputerPartitionFileVault2State) {
	o.PartitionFileVault2State = &v
}

// GetPartitionFileVault2Percent returns the PartitionFileVault2Percent field value if set, zero value otherwise.
func (o *ComputerPartitionEncryption) GetPartitionFileVault2Percent() int64 {
	if o == nil || IsNil(o.PartitionFileVault2Percent) {
		var ret int64
		return ret
	}
	return *o.PartitionFileVault2Percent
}

// GetPartitionFileVault2PercentOk returns a tuple with the PartitionFileVault2Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPartitionEncryption) GetPartitionFileVault2PercentOk() (*int64, bool) {
	if o == nil || IsNil(o.PartitionFileVault2Percent) {
		return nil, false
	}
	return o.PartitionFileVault2Percent, true
}

// HasPartitionFileVault2Percent returns a boolean if a field has been set.
func (o *ComputerPartitionEncryption) HasPartitionFileVault2Percent() bool {
	if o != nil && !IsNil(o.PartitionFileVault2Percent) {
		return true
	}

	return false
}

// SetPartitionFileVault2Percent gets a reference to the given int64 and assigns it to the PartitionFileVault2Percent field.
func (o *ComputerPartitionEncryption) SetPartitionFileVault2Percent(v int64) {
	o.PartitionFileVault2Percent = &v
}

func (o ComputerPartitionEncryption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComputerPartitionEncryption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PartitionName) {
		toSerialize["partitionName"] = o.PartitionName
	}
	if !IsNil(o.PartitionFileVault2State) {
		toSerialize["partitionFileVault2State"] = o.PartitionFileVault2State
	}
	if !IsNil(o.PartitionFileVault2Percent) {
		toSerialize["partitionFileVault2Percent"] = o.PartitionFileVault2Percent
	}
	return toSerialize, nil
}

type NullableComputerPartitionEncryption struct {
	value *ComputerPartitionEncryption
	isSet bool
}

func (v NullableComputerPartitionEncryption) Get() *ComputerPartitionEncryption {
	return v.value
}

func (v *NullableComputerPartitionEncryption) Set(val *ComputerPartitionEncryption) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerPartitionEncryption) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerPartitionEncryption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerPartitionEncryption(val *ComputerPartitionEncryption) *NullableComputerPartitionEncryption {
	return &NullableComputerPartitionEncryption{value: val, isSet: true}
}

func (v NullableComputerPartitionEncryption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerPartitionEncryption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


