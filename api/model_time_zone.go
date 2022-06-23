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

// TimeZone struct for TimeZone
type TimeZone struct {
	ZoneId *string `json:"zoneId,omitempty"`
	Region *string `json:"region,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
}

// NewTimeZone instantiates a new TimeZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeZone() *TimeZone {
	this := TimeZone{}
	return &this
}

// NewTimeZoneWithDefaults instantiates a new TimeZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeZoneWithDefaults() *TimeZone {
	this := TimeZone{}
	return &this
}

// GetZoneId returns the ZoneId field value if set, zero value otherwise.
func (o *TimeZone) GetZoneId() string {
	if o == nil || o.ZoneId == nil {
		var ret string
		return ret
	}
	return *o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeZone) GetZoneIdOk() (*string, bool) {
	if o == nil || o.ZoneId == nil {
		return nil, false
	}
	return o.ZoneId, true
}

// HasZoneId returns a boolean if a field has been set.
func (o *TimeZone) HasZoneId() bool {
	if o != nil && o.ZoneId != nil {
		return true
	}

	return false
}

// SetZoneId gets a reference to the given string and assigns it to the ZoneId field.
func (o *TimeZone) SetZoneId(v string) {
	o.ZoneId = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *TimeZone) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeZone) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *TimeZone) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *TimeZone) SetRegion(v string) {
	o.Region = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *TimeZone) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeZone) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *TimeZone) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *TimeZone) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o TimeZone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ZoneId != nil {
		toSerialize["zoneId"] = o.ZoneId
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	return json.Marshal(toSerialize)
}

type NullableTimeZone struct {
	value *TimeZone
	isSet bool
}

func (v NullableTimeZone) Get() *TimeZone {
	return v.value
}

func (v *NullableTimeZone) Set(val *TimeZone) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeZone) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeZone(val *TimeZone) *NullableTimeZone {
	return &NullableTimeZone{value: val, isSet: true}
}

func (v NullableTimeZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

