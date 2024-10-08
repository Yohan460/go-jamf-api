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

// checks if the JamfProServerUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JamfProServerUrl{}

// JamfProServerUrl struct for JamfProServerUrl
type JamfProServerUrl struct {
	Url string `json:"url"`
}

type _JamfProServerUrl JamfProServerUrl

// NewJamfProServerUrl instantiates a new JamfProServerUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJamfProServerUrl(url string) *JamfProServerUrl {
	this := JamfProServerUrl{}
	this.Url = url
	return &this
}

// NewJamfProServerUrlWithDefaults instantiates a new JamfProServerUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJamfProServerUrlWithDefaults() *JamfProServerUrl {
	this := JamfProServerUrl{}
	return &this
}

// GetUrl returns the Url field value
func (o *JamfProServerUrl) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *JamfProServerUrl) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *JamfProServerUrl) SetUrl(v string) {
	o.Url = v
}

func (o JamfProServerUrl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JamfProServerUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *JamfProServerUrl) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varJamfProServerUrl := _JamfProServerUrl{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJamfProServerUrl)

	if err != nil {
		return err
	}

	*o = JamfProServerUrl(varJamfProServerUrl)

	return err
}

type NullableJamfProServerUrl struct {
	value *JamfProServerUrl
	isSet bool
}

func (v NullableJamfProServerUrl) Get() *JamfProServerUrl {
	return v.value
}

func (v *NullableJamfProServerUrl) Set(val *JamfProServerUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableJamfProServerUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableJamfProServerUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJamfProServerUrl(val *JamfProServerUrl) *NullableJamfProServerUrl {
	return &NullableJamfProServerUrl{value: val, isSet: true}
}

func (v NullableJamfProServerUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJamfProServerUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


