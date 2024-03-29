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

// checks if the ComputerContentCachingDataMigrationErrorUserInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComputerContentCachingDataMigrationErrorUserInfo{}

// ComputerContentCachingDataMigrationErrorUserInfo struct for ComputerContentCachingDataMigrationErrorUserInfo
type ComputerContentCachingDataMigrationErrorUserInfo struct {
	Key *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewComputerContentCachingDataMigrationErrorUserInfo instantiates a new ComputerContentCachingDataMigrationErrorUserInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputerContentCachingDataMigrationErrorUserInfo() *ComputerContentCachingDataMigrationErrorUserInfo {
	this := ComputerContentCachingDataMigrationErrorUserInfo{}
	return &this
}

// NewComputerContentCachingDataMigrationErrorUserInfoWithDefaults instantiates a new ComputerContentCachingDataMigrationErrorUserInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputerContentCachingDataMigrationErrorUserInfoWithDefaults() *ComputerContentCachingDataMigrationErrorUserInfo {
	this := ComputerContentCachingDataMigrationErrorUserInfo{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ComputerContentCachingDataMigrationErrorUserInfo) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerContentCachingDataMigrationErrorUserInfo) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ComputerContentCachingDataMigrationErrorUserInfo) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ComputerContentCachingDataMigrationErrorUserInfo) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ComputerContentCachingDataMigrationErrorUserInfo) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerContentCachingDataMigrationErrorUserInfo) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ComputerContentCachingDataMigrationErrorUserInfo) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ComputerContentCachingDataMigrationErrorUserInfo) SetValue(v string) {
	o.Value = &v
}

func (o ComputerContentCachingDataMigrationErrorUserInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComputerContentCachingDataMigrationErrorUserInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableComputerContentCachingDataMigrationErrorUserInfo struct {
	value *ComputerContentCachingDataMigrationErrorUserInfo
	isSet bool
}

func (v NullableComputerContentCachingDataMigrationErrorUserInfo) Get() *ComputerContentCachingDataMigrationErrorUserInfo {
	return v.value
}

func (v *NullableComputerContentCachingDataMigrationErrorUserInfo) Set(val *ComputerContentCachingDataMigrationErrorUserInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerContentCachingDataMigrationErrorUserInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerContentCachingDataMigrationErrorUserInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerContentCachingDataMigrationErrorUserInfo(val *ComputerContentCachingDataMigrationErrorUserInfo) *NullableComputerContentCachingDataMigrationErrorUserInfo {
	return &NullableComputerContentCachingDataMigrationErrorUserInfo{value: val, isSet: true}
}

func (v NullableComputerContentCachingDataMigrationErrorUserInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerContentCachingDataMigrationErrorUserInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


