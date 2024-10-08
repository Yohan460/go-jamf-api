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

// checks if the UserTestSearchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserTestSearchRequest{}

// UserTestSearchRequest struct for UserTestSearchRequest
type UserTestSearchRequest struct {
	Username string `json:"username"`
}

type _UserTestSearchRequest UserTestSearchRequest

// NewUserTestSearchRequest instantiates a new UserTestSearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTestSearchRequest(username string) *UserTestSearchRequest {
	this := UserTestSearchRequest{}
	this.Username = username
	return &this
}

// NewUserTestSearchRequestWithDefaults instantiates a new UserTestSearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTestSearchRequestWithDefaults() *UserTestSearchRequest {
	this := UserTestSearchRequest{}
	return &this
}

// GetUsername returns the Username field value
func (o *UserTestSearchRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UserTestSearchRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UserTestSearchRequest) SetUsername(v string) {
	o.Username = v
}

func (o UserTestSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserTestSearchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

func (o *UserTestSearchRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
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

	varUserTestSearchRequest := _UserTestSearchRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserTestSearchRequest)

	if err != nil {
		return err
	}

	*o = UserTestSearchRequest(varUserTestSearchRequest)

	return err
}

type NullableUserTestSearchRequest struct {
	value *UserTestSearchRequest
	isSet bool
}

func (v NullableUserTestSearchRequest) Get() *UserTestSearchRequest {
	return v.value
}

func (v *NullableUserTestSearchRequest) Set(val *UserTestSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTestSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTestSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTestSearchRequest(val *UserTestSearchRequest) *NullableUserTestSearchRequest {
	return &NullableUserTestSearchRequest{value: val, isSet: true}
}

func (v NullableUserTestSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTestSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


