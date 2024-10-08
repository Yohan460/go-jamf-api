/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GroupTestSearchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupTestSearchRequest{}

// GroupTestSearchRequest struct for GroupTestSearchRequest
type GroupTestSearchRequest struct {
	Groupname string `json:"groupname"`
}

type _GroupTestSearchRequest GroupTestSearchRequest

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
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupTestSearchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groupname"] = o.Groupname
	return toSerialize, nil
}

func (o *GroupTestSearchRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"groupname",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGroupTestSearchRequest := _GroupTestSearchRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupTestSearchRequest)

	if err != nil {
		return err
	}

	*o = GroupTestSearchRequest(varGroupTestSearchRequest)

	return err
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


