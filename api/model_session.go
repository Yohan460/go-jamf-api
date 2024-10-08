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

// checks if the Session type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Session{}

// Session struct for Session
type Session struct {
	CurrentSiteId *int64 `json:"currentSiteId,omitempty"`
}

// NewSession instantiates a new Session object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSession() *Session {
	this := Session{}
	return &this
}

// NewSessionWithDefaults instantiates a new Session object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionWithDefaults() *Session {
	this := Session{}
	return &this
}

// GetCurrentSiteId returns the CurrentSiteId field value if set, zero value otherwise.
func (o *Session) GetCurrentSiteId() int64 {
	if o == nil || IsNil(o.CurrentSiteId) {
		var ret int64
		return ret
	}
	return *o.CurrentSiteId
}

// GetCurrentSiteIdOk returns a tuple with the CurrentSiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetCurrentSiteIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CurrentSiteId) {
		return nil, false
	}
	return o.CurrentSiteId, true
}

// HasCurrentSiteId returns a boolean if a field has been set.
func (o *Session) HasCurrentSiteId() bool {
	if o != nil && !IsNil(o.CurrentSiteId) {
		return true
	}

	return false
}

// SetCurrentSiteId gets a reference to the given int64 and assigns it to the CurrentSiteId field.
func (o *Session) SetCurrentSiteId(v int64) {
	o.CurrentSiteId = &v
}

func (o Session) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Session) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentSiteId) {
		toSerialize["currentSiteId"] = o.CurrentSiteId
	}
	return toSerialize, nil
}

type NullableSession struct {
	value *Session
	isSet bool
}

func (v NullableSession) Get() *Session {
	return v.value
}

func (v *NullableSession) Set(val *Session) {
	v.value = val
	v.isSet = true
}

func (v NullableSession) IsSet() bool {
	return v.isSet
}

func (v *NullableSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSession(val *Session) *NullableSession {
	return &NullableSession{value: val, isSet: true}
}

func (v NullableSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


