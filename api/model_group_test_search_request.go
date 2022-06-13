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

// GroupTestSearchRequest struct for GroupTestSearchRequest
type GroupTestSearchRequest struct {
	Groupname string `json:"groupname"`
}

// NewGroupTestSearchRequest instantiates a new GroupTestSearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupTestSearchRequest(groupname string) *GroupTestSearchRequest {
	this := GroupTestSearchRequest{}
	this.Groupname = groupname
	return &this
}

// NewGroupTestSearchRequestWithDefaults instantiates a new GroupTestSearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupTestSearchRequestWithDefaults() *GroupTestSearchRequest {
	this := GroupTestSearchRequest{}
	return &this
}

// GetGroupname returns the Groupname field value
func (o *GroupTestSearchRequest) GetGroupname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Groupname
}

// GetGroupnameOk returns a tuple with the Groupname field value
// and a boolean to check if the value has been set.
func (o *GroupTestSearchRequest) GetGroupnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Groupname, true
}

// SetGroupname sets field value
func (o *GroupTestSearchRequest) SetGroupname(v string) {
	o.Groupname = v
}

func (o GroupTestSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["groupname"] = o.Groupname
	}
	return json.Marshal(toSerialize)
}

type NullableGroupTestSearchRequest struct {
	value *GroupTestSearchRequest
	isSet bool
}

func (v NullableGroupTestSearchRequest) Get() *GroupTestSearchRequest {
	return v.value
}

func (v *NullableGroupTestSearchRequest) Set(val *GroupTestSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupTestSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupTestSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupTestSearchRequest(val *GroupTestSearchRequest) *NullableGroupTestSearchRequest {
	return &NullableGroupTestSearchRequest{value: val, isSet: true}
}

func (v NullableGroupTestSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupTestSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


