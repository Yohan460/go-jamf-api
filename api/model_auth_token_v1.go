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

// AuthTokenV1 struct for AuthTokenV1
type AuthTokenV1 struct {
	Token *string `json:"token,omitempty"`
	Expires *string `json:"expires,omitempty"`
}

// NewAuthTokenV1 instantiates a new AuthTokenV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthTokenV1() *AuthTokenV1 {
	this := AuthTokenV1{}
	return &this
}

// NewAuthTokenV1WithDefaults instantiates a new AuthTokenV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthTokenV1WithDefaults() *AuthTokenV1 {
	this := AuthTokenV1{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AuthTokenV1) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenV1) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AuthTokenV1) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *AuthTokenV1) SetToken(v string) {
	o.Token = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *AuthTokenV1) GetExpires() string {
	if o == nil || o.Expires == nil {
		var ret string
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenV1) GetExpiresOk() (*string, bool) {
	if o == nil || o.Expires == nil {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *AuthTokenV1) HasExpires() bool {
	if o != nil && o.Expires != nil {
		return true
	}

	return false
}

// SetExpires gets a reference to the given string and assigns it to the Expires field.
func (o *AuthTokenV1) SetExpires(v string) {
	o.Expires = &v
}

func (o AuthTokenV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Expires != nil {
		toSerialize["expires"] = o.Expires
	}
	return json.Marshal(toSerialize)
}

type NullableAuthTokenV1 struct {
	value *AuthTokenV1
	isSet bool
}

func (v NullableAuthTokenV1) Get() *AuthTokenV1 {
	return v.value
}

func (v *NullableAuthTokenV1) Set(val *AuthTokenV1) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthTokenV1) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthTokenV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthTokenV1(val *AuthTokenV1) *NullableAuthTokenV1 {
	return &NullableAuthTokenV1{value: val, isSet: true}
}

func (v NullableAuthTokenV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthTokenV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


