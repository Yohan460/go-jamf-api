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

// VolumePurchasingLocationAllOf struct for VolumePurchasingLocationAllOf
type VolumePurchasingLocationAllOf struct {
	Content []VolumePurchasingContent `json:"content,omitempty"`
}

// NewVolumePurchasingLocationAllOf instantiates a new VolumePurchasingLocationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumePurchasingLocationAllOf() *VolumePurchasingLocationAllOf {
	this := VolumePurchasingLocationAllOf{}
	return &this
}

// NewVolumePurchasingLocationAllOfWithDefaults instantiates a new VolumePurchasingLocationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumePurchasingLocationAllOfWithDefaults() *VolumePurchasingLocationAllOf {
	this := VolumePurchasingLocationAllOf{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *VolumePurchasingLocationAllOf) GetContent() []VolumePurchasingContent {
	if o == nil || o.Content == nil {
		var ret []VolumePurchasingContent
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingLocationAllOf) GetContentOk() ([]VolumePurchasingContent, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *VolumePurchasingLocationAllOf) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []VolumePurchasingContent and assigns it to the Content field.
func (o *VolumePurchasingLocationAllOf) SetContent(v []VolumePurchasingContent) {
	o.Content = v
}

func (o VolumePurchasingLocationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableVolumePurchasingLocationAllOf struct {
	value *VolumePurchasingLocationAllOf
	isSet bool
}

func (v NullableVolumePurchasingLocationAllOf) Get() *VolumePurchasingLocationAllOf {
	return v.value
}

func (v *NullableVolumePurchasingLocationAllOf) Set(val *VolumePurchasingLocationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumePurchasingLocationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumePurchasingLocationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumePurchasingLocationAllOf(val *VolumePurchasingLocationAllOf) *NullableVolumePurchasingLocationAllOf {
	return &NullableVolumePurchasingLocationAllOf{value: val, isSet: true}
}

func (v NullableVolumePurchasingLocationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumePurchasingLocationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


