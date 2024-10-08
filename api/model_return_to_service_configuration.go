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

// checks if the ReturnToServiceConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnToServiceConfiguration{}

// ReturnToServiceConfiguration struct for ReturnToServiceConfiguration
type ReturnToServiceConfiguration struct {
	// Id of the Return to Service Configuration.
	Id *string `json:"id,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	// Id of the wifi profile that is associated with the return to service configuration.
	WifiProfileId *string `json:"wifiProfileId,omitempty"`
}

// NewReturnToServiceConfiguration instantiates a new ReturnToServiceConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnToServiceConfiguration() *ReturnToServiceConfiguration {
	this := ReturnToServiceConfiguration{}
	var displayName string = "false"
	this.DisplayName = &displayName
	return &this
}

// NewReturnToServiceConfigurationWithDefaults instantiates a new ReturnToServiceConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnToServiceConfigurationWithDefaults() *ReturnToServiceConfiguration {
	this := ReturnToServiceConfiguration{}
	var displayName string = "false"
	this.DisplayName = &displayName
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReturnToServiceConfiguration) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnToServiceConfiguration) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReturnToServiceConfiguration) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReturnToServiceConfiguration) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ReturnToServiceConfiguration) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnToServiceConfiguration) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ReturnToServiceConfiguration) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ReturnToServiceConfiguration) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetWifiProfileId returns the WifiProfileId field value if set, zero value otherwise.
func (o *ReturnToServiceConfiguration) GetWifiProfileId() string {
	if o == nil || IsNil(o.WifiProfileId) {
		var ret string
		return ret
	}
	return *o.WifiProfileId
}

// GetWifiProfileIdOk returns a tuple with the WifiProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnToServiceConfiguration) GetWifiProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.WifiProfileId) {
		return nil, false
	}
	return o.WifiProfileId, true
}

// HasWifiProfileId returns a boolean if a field has been set.
func (o *ReturnToServiceConfiguration) HasWifiProfileId() bool {
	if o != nil && !IsNil(o.WifiProfileId) {
		return true
	}

	return false
}

// SetWifiProfileId gets a reference to the given string and assigns it to the WifiProfileId field.
func (o *ReturnToServiceConfiguration) SetWifiProfileId(v string) {
	o.WifiProfileId = &v
}

func (o ReturnToServiceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnToServiceConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.WifiProfileId) {
		toSerialize["wifiProfileId"] = o.WifiProfileId
	}
	return toSerialize, nil
}

type NullableReturnToServiceConfiguration struct {
	value *ReturnToServiceConfiguration
	isSet bool
}

func (v NullableReturnToServiceConfiguration) Get() *ReturnToServiceConfiguration {
	return v.value
}

func (v *NullableReturnToServiceConfiguration) Set(val *ReturnToServiceConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnToServiceConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnToServiceConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnToServiceConfiguration(val *ReturnToServiceConfiguration) *NullableReturnToServiceConfiguration {
	return &NullableReturnToServiceConfiguration{value: val, isSet: true}
}

func (v NullableReturnToServiceConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnToServiceConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


