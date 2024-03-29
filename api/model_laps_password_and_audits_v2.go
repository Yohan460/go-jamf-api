/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the LapsPasswordAndAuditsV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LapsPasswordAndAuditsV2{}

// LapsPasswordAndAuditsV2 struct for LapsPasswordAndAuditsV2
type LapsPasswordAndAuditsV2 struct {
	Password *string `json:"password,omitempty"`
	DateLastSeen NullableTime `json:"dateLastSeen,omitempty"`
	ExpirationTime NullableTime `json:"expirationTime,omitempty"`
	Audits []LapsAuditV2 `json:"audits,omitempty"`
}

// NewLapsPasswordAndAuditsV2 instantiates a new LapsPasswordAndAuditsV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLapsPasswordAndAuditsV2() *LapsPasswordAndAuditsV2 {
	this := LapsPasswordAndAuditsV2{}
	return &this
}

// NewLapsPasswordAndAuditsV2WithDefaults instantiates a new LapsPasswordAndAuditsV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLapsPasswordAndAuditsV2WithDefaults() *LapsPasswordAndAuditsV2 {
	this := LapsPasswordAndAuditsV2{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *LapsPasswordAndAuditsV2) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LapsPasswordAndAuditsV2) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *LapsPasswordAndAuditsV2) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *LapsPasswordAndAuditsV2) SetPassword(v string) {
	o.Password = &v
}

// GetDateLastSeen returns the DateLastSeen field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LapsPasswordAndAuditsV2) GetDateLastSeen() time.Time {
	if o == nil || IsNil(o.DateLastSeen.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DateLastSeen.Get()
}

// GetDateLastSeenOk returns a tuple with the DateLastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LapsPasswordAndAuditsV2) GetDateLastSeenOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateLastSeen.Get(), o.DateLastSeen.IsSet()
}

// HasDateLastSeen returns a boolean if a field has been set.
func (o *LapsPasswordAndAuditsV2) HasDateLastSeen() bool {
	if o != nil && o.DateLastSeen.IsSet() {
		return true
	}

	return false
}

// SetDateLastSeen gets a reference to the given NullableTime and assigns it to the DateLastSeen field.
func (o *LapsPasswordAndAuditsV2) SetDateLastSeen(v time.Time) {
	o.DateLastSeen.Set(&v)
}
// SetDateLastSeenNil sets the value for DateLastSeen to be an explicit nil
func (o *LapsPasswordAndAuditsV2) SetDateLastSeenNil() {
	o.DateLastSeen.Set(nil)
}

// UnsetDateLastSeen ensures that no value is present for DateLastSeen, not even an explicit nil
func (o *LapsPasswordAndAuditsV2) UnsetDateLastSeen() {
	o.DateLastSeen.Unset()
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LapsPasswordAndAuditsV2) GetExpirationTime() time.Time {
	if o == nil || IsNil(o.ExpirationTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationTime.Get()
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LapsPasswordAndAuditsV2) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationTime.Get(), o.ExpirationTime.IsSet()
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *LapsPasswordAndAuditsV2) HasExpirationTime() bool {
	if o != nil && o.ExpirationTime.IsSet() {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given NullableTime and assigns it to the ExpirationTime field.
func (o *LapsPasswordAndAuditsV2) SetExpirationTime(v time.Time) {
	o.ExpirationTime.Set(&v)
}
// SetExpirationTimeNil sets the value for ExpirationTime to be an explicit nil
func (o *LapsPasswordAndAuditsV2) SetExpirationTimeNil() {
	o.ExpirationTime.Set(nil)
}

// UnsetExpirationTime ensures that no value is present for ExpirationTime, not even an explicit nil
func (o *LapsPasswordAndAuditsV2) UnsetExpirationTime() {
	o.ExpirationTime.Unset()
}

// GetAudits returns the Audits field value if set, zero value otherwise.
func (o *LapsPasswordAndAuditsV2) GetAudits() []LapsAuditV2 {
	if o == nil || IsNil(o.Audits) {
		var ret []LapsAuditV2
		return ret
	}
	return o.Audits
}

// GetAuditsOk returns a tuple with the Audits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LapsPasswordAndAuditsV2) GetAuditsOk() ([]LapsAuditV2, bool) {
	if o == nil || IsNil(o.Audits) {
		return nil, false
	}
	return o.Audits, true
}

// HasAudits returns a boolean if a field has been set.
func (o *LapsPasswordAndAuditsV2) HasAudits() bool {
	if o != nil && !IsNil(o.Audits) {
		return true
	}

	return false
}

// SetAudits gets a reference to the given []LapsAuditV2 and assigns it to the Audits field.
func (o *LapsPasswordAndAuditsV2) SetAudits(v []LapsAuditV2) {
	o.Audits = v
}

func (o LapsPasswordAndAuditsV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LapsPasswordAndAuditsV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if o.DateLastSeen.IsSet() {
		toSerialize["dateLastSeen"] = o.DateLastSeen.Get()
	}
	if o.ExpirationTime.IsSet() {
		toSerialize["expirationTime"] = o.ExpirationTime.Get()
	}
	if !IsNil(o.Audits) {
		toSerialize["audits"] = o.Audits
	}
	return toSerialize, nil
}

type NullableLapsPasswordAndAuditsV2 struct {
	value *LapsPasswordAndAuditsV2
	isSet bool
}

func (v NullableLapsPasswordAndAuditsV2) Get() *LapsPasswordAndAuditsV2 {
	return v.value
}

func (v *NullableLapsPasswordAndAuditsV2) Set(val *LapsPasswordAndAuditsV2) {
	v.value = val
	v.isSet = true
}

func (v NullableLapsPasswordAndAuditsV2) IsSet() bool {
	return v.isSet
}

func (v *NullableLapsPasswordAndAuditsV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLapsPasswordAndAuditsV2(val *LapsPasswordAndAuditsV2) *NullableLapsPasswordAndAuditsV2 {
	return &NullableLapsPasswordAndAuditsV2{value: val, isSet: true}
}

func (v NullableLapsPasswordAndAuditsV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLapsPasswordAndAuditsV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


