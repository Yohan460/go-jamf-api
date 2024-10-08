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

// checks if the TeacherSettingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeacherSettingsRequest{}

// TeacherSettingsRequest struct for TeacherSettingsRequest
type TeacherSettingsRequest struct {
	IsEnabled *bool `json:"isEnabled,omitempty"`
	TimezoneId *string `json:"timezoneId,omitempty"`
	AutoClear NullableString `json:"autoClear,omitempty"`
	MaxRestrictionLengthSeconds *int64 `json:"maxRestrictionLengthSeconds,omitempty"`
	SafelistedApps []SafelistedApp `json:"safelistedApps,omitempty"`
}

// NewTeacherSettingsRequest instantiates a new TeacherSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeacherSettingsRequest() *TeacherSettingsRequest {
	this := TeacherSettingsRequest{}
	return &this
}

// NewTeacherSettingsRequestWithDefaults instantiates a new TeacherSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeacherSettingsRequestWithDefaults() *TeacherSettingsRequest {
	this := TeacherSettingsRequest{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *TeacherSettingsRequest) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeacherSettingsRequest) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *TeacherSettingsRequest) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *TeacherSettingsRequest) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetTimezoneId returns the TimezoneId field value if set, zero value otherwise.
func (o *TeacherSettingsRequest) GetTimezoneId() string {
	if o == nil || IsNil(o.TimezoneId) {
		var ret string
		return ret
	}
	return *o.TimezoneId
}

// GetTimezoneIdOk returns a tuple with the TimezoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeacherSettingsRequest) GetTimezoneIdOk() (*string, bool) {
	if o == nil || IsNil(o.TimezoneId) {
		return nil, false
	}
	return o.TimezoneId, true
}

// HasTimezoneId returns a boolean if a field has been set.
func (o *TeacherSettingsRequest) HasTimezoneId() bool {
	if o != nil && !IsNil(o.TimezoneId) {
		return true
	}

	return false
}

// SetTimezoneId gets a reference to the given string and assigns it to the TimezoneId field.
func (o *TeacherSettingsRequest) SetTimezoneId(v string) {
	o.TimezoneId = &v
}

// GetAutoClear returns the AutoClear field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeacherSettingsRequest) GetAutoClear() string {
	if o == nil || IsNil(o.AutoClear.Get()) {
		var ret string
		return ret
	}
	return *o.AutoClear.Get()
}

// GetAutoClearOk returns a tuple with the AutoClear field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeacherSettingsRequest) GetAutoClearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoClear.Get(), o.AutoClear.IsSet()
}

// HasAutoClear returns a boolean if a field has been set.
func (o *TeacherSettingsRequest) HasAutoClear() bool {
	if o != nil && o.AutoClear.IsSet() {
		return true
	}

	return false
}

// SetAutoClear gets a reference to the given NullableString and assigns it to the AutoClear field.
func (o *TeacherSettingsRequest) SetAutoClear(v string) {
	o.AutoClear.Set(&v)
}
// SetAutoClearNil sets the value for AutoClear to be an explicit nil
func (o *TeacherSettingsRequest) SetAutoClearNil() {
	o.AutoClear.Set(nil)
}

// UnsetAutoClear ensures that no value is present for AutoClear, not even an explicit nil
func (o *TeacherSettingsRequest) UnsetAutoClear() {
	o.AutoClear.Unset()
}

// GetMaxRestrictionLengthSeconds returns the MaxRestrictionLengthSeconds field value if set, zero value otherwise.
func (o *TeacherSettingsRequest) GetMaxRestrictionLengthSeconds() int64 {
	if o == nil || IsNil(o.MaxRestrictionLengthSeconds) {
		var ret int64
		return ret
	}
	return *o.MaxRestrictionLengthSeconds
}

// GetMaxRestrictionLengthSecondsOk returns a tuple with the MaxRestrictionLengthSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeacherSettingsRequest) GetMaxRestrictionLengthSecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxRestrictionLengthSeconds) {
		return nil, false
	}
	return o.MaxRestrictionLengthSeconds, true
}

// HasMaxRestrictionLengthSeconds returns a boolean if a field has been set.
func (o *TeacherSettingsRequest) HasMaxRestrictionLengthSeconds() bool {
	if o != nil && !IsNil(o.MaxRestrictionLengthSeconds) {
		return true
	}

	return false
}

// SetMaxRestrictionLengthSeconds gets a reference to the given int64 and assigns it to the MaxRestrictionLengthSeconds field.
func (o *TeacherSettingsRequest) SetMaxRestrictionLengthSeconds(v int64) {
	o.MaxRestrictionLengthSeconds = &v
}

// GetSafelistedApps returns the SafelistedApps field value if set, zero value otherwise.
func (o *TeacherSettingsRequest) GetSafelistedApps() []SafelistedApp {
	if o == nil || IsNil(o.SafelistedApps) {
		var ret []SafelistedApp
		return ret
	}
	return o.SafelistedApps
}

// GetSafelistedAppsOk returns a tuple with the SafelistedApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeacherSettingsRequest) GetSafelistedAppsOk() ([]SafelistedApp, bool) {
	if o == nil || IsNil(o.SafelistedApps) {
		return nil, false
	}
	return o.SafelistedApps, true
}

// HasSafelistedApps returns a boolean if a field has been set.
func (o *TeacherSettingsRequest) HasSafelistedApps() bool {
	if o != nil && !IsNil(o.SafelistedApps) {
		return true
	}

	return false
}

// SetSafelistedApps gets a reference to the given []SafelistedApp and assigns it to the SafelistedApps field.
func (o *TeacherSettingsRequest) SetSafelistedApps(v []SafelistedApp) {
	o.SafelistedApps = v
}

func (o TeacherSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeacherSettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !IsNil(o.TimezoneId) {
		toSerialize["timezoneId"] = o.TimezoneId
	}
	if o.AutoClear.IsSet() {
		toSerialize["autoClear"] = o.AutoClear.Get()
	}
	if !IsNil(o.MaxRestrictionLengthSeconds) {
		toSerialize["maxRestrictionLengthSeconds"] = o.MaxRestrictionLengthSeconds
	}
	if !IsNil(o.SafelistedApps) {
		toSerialize["safelistedApps"] = o.SafelistedApps
	}
	return toSerialize, nil
}

type NullableTeacherSettingsRequest struct {
	value *TeacherSettingsRequest
	isSet bool
}

func (v NullableTeacherSettingsRequest) Get() *TeacherSettingsRequest {
	return v.value
}

func (v *NullableTeacherSettingsRequest) Set(val *TeacherSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTeacherSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTeacherSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeacherSettingsRequest(val *TeacherSettingsRequest) *NullableTeacherSettingsRequest {
	return &NullableTeacherSettingsRequest{value: val, isSet: true}
}

func (v NullableTeacherSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeacherSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


