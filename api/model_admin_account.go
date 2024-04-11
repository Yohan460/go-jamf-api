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

// checks if the AdminAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminAccount{}

// AdminAccount struct for AdminAccount
type AdminAccount struct {
	Id *int64 `json:"id,omitempty"`
}

// NewAdminAccount instantiates a new AdminAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminAccount() *AdminAccount {
	this := AdminAccount{}
	return &this
}

// NewAdminAccountWithDefaults instantiates a new AdminAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminAccountWithDefaults() *AdminAccount {
	this := AdminAccount{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdminAccount) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminAccount) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdminAccount) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AdminAccount) SetId(v int64) {
	o.Id = &v
}

func (o AdminAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableAdminAccount struct {
	value *AdminAccount
	isSet bool
}

func (v NullableAdminAccount) Get() *AdminAccount {
	return v.value
}

func (v *NullableAdminAccount) Set(val *AdminAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminAccount(val *AdminAccount) *NullableAdminAccount {
	return &NullableAdminAccount{value: val, isSet: true}
}

func (v NullableAdminAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


