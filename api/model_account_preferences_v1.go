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

// checks if the AccountPreferencesV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountPreferencesV1{}

// AccountPreferencesV1 struct for AccountPreferencesV1
type AccountPreferencesV1 struct {
	Language *string `json:"language,omitempty"`
	DateFormat *string `json:"dateFormat,omitempty"`
	Region *string `json:"region,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	DisableRelativeDates *bool `json:"disableRelativeDates,omitempty"`
}

// NewAccountPreferencesV1 instantiates a new AccountPreferencesV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountPreferencesV1() *AccountPreferencesV1 {
	this := AccountPreferencesV1{}
	return &this
}

// NewAccountPreferencesV1WithDefaults instantiates a new AccountPreferencesV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountPreferencesV1WithDefaults() *AccountPreferencesV1 {
	this := AccountPreferencesV1{}
	return &this
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *AccountPreferencesV1) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV1) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *AccountPreferencesV1) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *AccountPreferencesV1) SetLanguage(v string) {
	o.Language = &v
}

// GetDateFormat returns the DateFormat field value if set, zero value otherwise.
func (o *AccountPreferencesV1) GetDateFormat() string {
	if o == nil || IsNil(o.DateFormat) {
		var ret string
		return ret
	}
	return *o.DateFormat
}

// GetDateFormatOk returns a tuple with the DateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV1) GetDateFormatOk() (*string, bool) {
	if o == nil || IsNil(o.DateFormat) {
		return nil, false
	}
	return o.DateFormat, true
}

// HasDateFormat returns a boolean if a field has been set.
func (o *AccountPreferencesV1) HasDateFormat() bool {
	if o != nil && !IsNil(o.DateFormat) {
		return true
	}

	return false
}

// SetDateFormat gets a reference to the given string and assigns it to the DateFormat field.
func (o *AccountPreferencesV1) SetDateFormat(v string) {
	o.DateFormat = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AccountPreferencesV1) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV1) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AccountPreferencesV1) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AccountPreferencesV1) SetRegion(v string) {
	o.Region = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *AccountPreferencesV1) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV1) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *AccountPreferencesV1) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *AccountPreferencesV1) SetTimezone(v string) {
	o.Timezone = &v
}

// GetDisableRelativeDates returns the DisableRelativeDates field value if set, zero value otherwise.
func (o *AccountPreferencesV1) GetDisableRelativeDates() bool {
	if o == nil || IsNil(o.DisableRelativeDates) {
		var ret bool
		return ret
	}
	return *o.DisableRelativeDates
}

// GetDisableRelativeDatesOk returns a tuple with the DisableRelativeDates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV1) GetDisableRelativeDatesOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableRelativeDates) {
		return nil, false
	}
	return o.DisableRelativeDates, true
}

// HasDisableRelativeDates returns a boolean if a field has been set.
func (o *AccountPreferencesV1) HasDisableRelativeDates() bool {
	if o != nil && !IsNil(o.DisableRelativeDates) {
		return true
	}

	return false
}

// SetDisableRelativeDates gets a reference to the given bool and assigns it to the DisableRelativeDates field.
func (o *AccountPreferencesV1) SetDisableRelativeDates(v bool) {
	o.DisableRelativeDates = &v
}

func (o AccountPreferencesV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountPreferencesV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.DateFormat) {
		toSerialize["dateFormat"] = o.DateFormat
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.DisableRelativeDates) {
		toSerialize["disableRelativeDates"] = o.DisableRelativeDates
	}
	return toSerialize, nil
}

type NullableAccountPreferencesV1 struct {
	value *AccountPreferencesV1
	isSet bool
}

func (v NullableAccountPreferencesV1) Get() *AccountPreferencesV1 {
	return v.value
}

func (v *NullableAccountPreferencesV1) Set(val *AccountPreferencesV1) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountPreferencesV1) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountPreferencesV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountPreferencesV1(val *AccountPreferencesV1) *NullableAccountPreferencesV1 {
	return &NullableAccountPreferencesV1{value: val, isSet: true}
}

func (v NullableAccountPreferencesV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountPreferencesV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


