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

// checks if the ObjectHistoryV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectHistoryV1{}

// ObjectHistoryV1 struct for ObjectHistoryV1
type ObjectHistoryV1 struct {
	Id *string `json:"id,omitempty"`
	Username *string `json:"username,omitempty"`
	Date *string `json:"date,omitempty"`
	Note *string `json:"note,omitempty"`
	Details NullableString `json:"details,omitempty"`
}

// NewObjectHistoryV1 instantiates a new ObjectHistoryV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectHistoryV1() *ObjectHistoryV1 {
	this := ObjectHistoryV1{}
	return &this
}

// NewObjectHistoryV1WithDefaults instantiates a new ObjectHistoryV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectHistoryV1WithDefaults() *ObjectHistoryV1 {
	this := ObjectHistoryV1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ObjectHistoryV1) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectHistoryV1) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ObjectHistoryV1) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ObjectHistoryV1) SetId(v string) {
	o.Id = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ObjectHistoryV1) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectHistoryV1) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ObjectHistoryV1) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ObjectHistoryV1) SetUsername(v string) {
	o.Username = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *ObjectHistoryV1) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectHistoryV1) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *ObjectHistoryV1) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *ObjectHistoryV1) SetDate(v string) {
	o.Date = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *ObjectHistoryV1) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectHistoryV1) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *ObjectHistoryV1) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *ObjectHistoryV1) SetNote(v string) {
	o.Note = &v
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectHistoryV1) GetDetails() string {
	if o == nil || IsNil(o.Details.Get()) {
		var ret string
		return ret
	}
	return *o.Details.Get()
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectHistoryV1) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Details.Get(), o.Details.IsSet()
}

// HasDetails returns a boolean if a field has been set.
func (o *ObjectHistoryV1) HasDetails() bool {
	if o != nil && o.Details.IsSet() {
		return true
	}

	return false
}

// SetDetails gets a reference to the given NullableString and assigns it to the Details field.
func (o *ObjectHistoryV1) SetDetails(v string) {
	o.Details.Set(&v)
}
// SetDetailsNil sets the value for Details to be an explicit nil
func (o *ObjectHistoryV1) SetDetailsNil() {
	o.Details.Set(nil)
}

// UnsetDetails ensures that no value is present for Details, not even an explicit nil
func (o *ObjectHistoryV1) UnsetDetails() {
	o.Details.Unset()
}

func (o ObjectHistoryV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectHistoryV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if o.Details.IsSet() {
		toSerialize["details"] = o.Details.Get()
	}
	return toSerialize, nil
}

type NullableObjectHistoryV1 struct {
	value *ObjectHistoryV1
	isSet bool
}

func (v NullableObjectHistoryV1) Get() *ObjectHistoryV1 {
	return v.value
}

func (v *NullableObjectHistoryV1) Set(val *ObjectHistoryV1) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectHistoryV1) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectHistoryV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectHistoryV1(val *ObjectHistoryV1) *NullableObjectHistoryV1 {
	return &NullableObjectHistoryV1{value: val, isSet: true}
}

func (v NullableObjectHistoryV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectHistoryV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


