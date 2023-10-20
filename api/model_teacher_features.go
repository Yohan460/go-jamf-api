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

// checks if the TeacherFeatures type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeacherFeatures{}

// TeacherFeatures struct for TeacherFeatures
type TeacherFeatures struct {
	IsAllowAppLock *bool `json:"isAllowAppLock,omitempty"`
	IsAllowWebLock *bool `json:"isAllowWebLock,omitempty"`
	IsAllowRestrictions *bool `json:"isAllowRestrictions,omitempty"`
	IsAllowAttentionScreen *bool `json:"isAllowAttentionScreen,omitempty"`
	IsAllowClearPasscode *bool `json:"isAllowClearPasscode,omitempty"`
}

// NewTeacherFeatures instantiates a new TeacherFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeacherFeatures() *TeacherFeatures {
	this := TeacherFeatures{}
	return &this
}

// NewTeacherFeaturesWithDefaults instantiates a new TeacherFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeacherFeaturesWithDefaults() *TeacherFeatures {
	this := TeacherFeatures{}
	return &this
}

// GetIsAllowAppLock returns the IsAllowAppLock field value if set, zero value otherwise.
func (o *TeacherFeatures) GetIsAllowAppLock() bool {
	if o == nil || IsNil(o.IsAllowAppLock) {
		var ret bool
		return ret
	}
	return *o.IsAllowAppLock
}

// GetIsAllowAppLockOk returns a tuple with the IsAllowAppLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeacherFeatures) GetIsAllowAppLockOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAllowAppLock) {
		return nil, false
	}
	return o.IsAllowAppLock, true
}

// HasIsAllowAppLock returns a boolean if a field has been set.
func (o *TeacherFeatures) HasIsAllowAppLock() bool {
	if o != nil && !IsNil(o.IsAllowAppLock) {
		return true
	}

	return false
}

// SetIsAllowAppLock gets a reference to the given bool and assigns it to the IsAllowAppLock field.
func (o *TeacherFeatures) SetIsAllowAppLock(v bool) {
	o.IsAllowAppLock = &v
}

// GetIsAllowWebLock returns the IsAllowWebLock field value if set, zero value otherwise.
func (o *TeacherFeatures) GetIsAllowWebLock() bool {
	if o == nil || IsNil(o.IsAllowWebLock) {
		var ret bool
		return ret
	}
	return *o.IsAllowWebLock
}

// GetIsAllowWebLockOk returns a tuple with the IsAllowWebLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeacherFeatures) GetIsAllowWebLockOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAllowWebLock) {
		return nil, false
	}
	return o.IsAllowWebLock, true
}

// HasIsAllowWebLock returns a boolean if a field has been set.
func (o *TeacherFeatures) HasIsAllowWebLock() bool {
	if o != nil && !IsNil(o.IsAllowWebLock) {
		return true
	}

	return false
}

// SetIsAllowWebLock gets a reference to the given bool and assigns it to the IsAllowWebLock field.
func (o *TeacherFeatures) SetIsAllowWebLock(v bool) {
	o.IsAllowWebLock = &v
}

// GetIsAllowRestrictions returns the IsAllowRestrictions field value if set, zero value otherwise.
func (o *TeacherFeatures) GetIsAllowRestrictions() bool {
	if o == nil || IsNil(o.IsAllowRestrictions) {
		var ret bool
		return ret
	}
	return *o.IsAllowRestrictions
}

// GetIsAllowRestrictionsOk returns a tuple with the IsAllowRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeacherFeatures) GetIsAllowRestrictionsOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAllowRestrictions) {
		return nil, false
	}
	return o.IsAllowRestrictions, true
}

// HasIsAllowRestrictions returns a boolean if a field has been set.
func (o *TeacherFeatures) HasIsAllowRestrictions() bool {
	if o != nil && !IsNil(o.IsAllowRestrictions) {
		return true
	}

	return false
}

// SetIsAllowRestrictions gets a reference to the given bool and assigns it to the IsAllowRestrictions field.
func (o *TeacherFeatures) SetIsAllowRestrictions(v bool) {
	o.IsAllowRestrictions = &v
}

// GetIsAllowAttentionScreen returns the IsAllowAttentionScreen field value if set, zero value otherwise.
func (o *TeacherFeatures) GetIsAllowAttentionScreen() bool {
	if o == nil || IsNil(o.IsAllowAttentionScreen) {
		var ret bool
		return ret
	}
	return *o.IsAllowAttentionScreen
}

// GetIsAllowAttentionScreenOk returns a tuple with the IsAllowAttentionScreen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeacherFeatures) GetIsAllowAttentionScreenOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAllowAttentionScreen) {
		return nil, false
	}
	return o.IsAllowAttentionScreen, true
}

// HasIsAllowAttentionScreen returns a boolean if a field has been set.
func (o *TeacherFeatures) HasIsAllowAttentionScreen() bool {
	if o != nil && !IsNil(o.IsAllowAttentionScreen) {
		return true
	}

	return false
}

// SetIsAllowAttentionScreen gets a reference to the given bool and assigns it to the IsAllowAttentionScreen field.
func (o *TeacherFeatures) SetIsAllowAttentionScreen(v bool) {
	o.IsAllowAttentionScreen = &v
}

// GetIsAllowClearPasscode returns the IsAllowClearPasscode field value if set, zero value otherwise.
func (o *TeacherFeatures) GetIsAllowClearPasscode() bool {
	if o == nil || IsNil(o.IsAllowClearPasscode) {
		var ret bool
		return ret
	}
	return *o.IsAllowClearPasscode
}

// GetIsAllowClearPasscodeOk returns a tuple with the IsAllowClearPasscode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeacherFeatures) GetIsAllowClearPasscodeOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAllowClearPasscode) {
		return nil, false
	}
	return o.IsAllowClearPasscode, true
}

// HasIsAllowClearPasscode returns a boolean if a field has been set.
func (o *TeacherFeatures) HasIsAllowClearPasscode() bool {
	if o != nil && !IsNil(o.IsAllowClearPasscode) {
		return true
	}

	return false
}

// SetIsAllowClearPasscode gets a reference to the given bool and assigns it to the IsAllowClearPasscode field.
func (o *TeacherFeatures) SetIsAllowClearPasscode(v bool) {
	o.IsAllowClearPasscode = &v
}

func (o TeacherFeatures) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeacherFeatures) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsAllowAppLock) {
		toSerialize["isAllowAppLock"] = o.IsAllowAppLock
	}
	if !IsNil(o.IsAllowWebLock) {
		toSerialize["isAllowWebLock"] = o.IsAllowWebLock
	}
	if !IsNil(o.IsAllowRestrictions) {
		toSerialize["isAllowRestrictions"] = o.IsAllowRestrictions
	}
	if !IsNil(o.IsAllowAttentionScreen) {
		toSerialize["isAllowAttentionScreen"] = o.IsAllowAttentionScreen
	}
	if !IsNil(o.IsAllowClearPasscode) {
		toSerialize["isAllowClearPasscode"] = o.IsAllowClearPasscode
	}
	return toSerialize, nil
}

type NullableTeacherFeatures struct {
	value *TeacherFeatures
	isSet bool
}

func (v NullableTeacherFeatures) Get() *TeacherFeatures {
	return v.value
}

func (v *NullableTeacherFeatures) Set(val *TeacherFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableTeacherFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableTeacherFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeacherFeatures(val *TeacherFeatures) *NullableTeacherFeatures {
	return &NullableTeacherFeatures{value: val, isSet: true}
}

func (v NullableTeacherFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeacherFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


