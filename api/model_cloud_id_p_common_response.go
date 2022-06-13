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

// CloudIdPCommonResponse A Cloud Identity Provider information for responses
type CloudIdPCommonResponse struct {
	Id *string `json:"id,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	ProviderName *string `json:"providerName,omitempty"`
}

// NewCloudIdPCommonResponse instantiates a new CloudIdPCommonResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudIdPCommonResponse() *CloudIdPCommonResponse {
	this := CloudIdPCommonResponse{}
	return &this
}

// NewCloudIdPCommonResponseWithDefaults instantiates a new CloudIdPCommonResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudIdPCommonResponseWithDefaults() *CloudIdPCommonResponse {
	this := CloudIdPCommonResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CloudIdPCommonResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudIdPCommonResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudIdPCommonResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CloudIdPCommonResponse) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CloudIdPCommonResponse) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudIdPCommonResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CloudIdPCommonResponse) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CloudIdPCommonResponse) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CloudIdPCommonResponse) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudIdPCommonResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CloudIdPCommonResponse) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CloudIdPCommonResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *CloudIdPCommonResponse) GetProviderName() string {
	if o == nil || o.ProviderName == nil {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudIdPCommonResponse) GetProviderNameOk() (*string, bool) {
	if o == nil || o.ProviderName == nil {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *CloudIdPCommonResponse) HasProviderName() bool {
	if o != nil && o.ProviderName != nil {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *CloudIdPCommonResponse) SetProviderName(v string) {
	o.ProviderName = &v
}

func (o CloudIdPCommonResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ProviderName != nil {
		toSerialize["providerName"] = o.ProviderName
	}
	return json.Marshal(toSerialize)
}

type NullableCloudIdPCommonResponse struct {
	value *CloudIdPCommonResponse
	isSet bool
}

func (v NullableCloudIdPCommonResponse) Get() *CloudIdPCommonResponse {
	return v.value
}

func (v *NullableCloudIdPCommonResponse) Set(val *CloudIdPCommonResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudIdPCommonResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudIdPCommonResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudIdPCommonResponse(val *CloudIdPCommonResponse) *NullableCloudIdPCommonResponse {
	return &NullableCloudIdPCommonResponse{value: val, isSet: true}
}

func (v NullableCloudIdPCommonResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudIdPCommonResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


