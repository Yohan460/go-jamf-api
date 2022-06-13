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

// DeprecatedConfigurationResponse An old Cloud Identity Provider configuration for responses
type DeprecatedConfigurationResponse struct {
	Id string `json:"id"`
	Server DeprecatedServerResponse `json:"server"`
	Mappings *CloudLdapMappingsResponse `json:"mappings,omitempty"`
}

// NewDeprecatedConfigurationResponse instantiates a new DeprecatedConfigurationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeprecatedConfigurationResponse(id string, server DeprecatedServerResponse) *DeprecatedConfigurationResponse {
	this := DeprecatedConfigurationResponse{}
	this.Id = id
	this.Server = server
	return &this
}

// NewDeprecatedConfigurationResponseWithDefaults instantiates a new DeprecatedConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeprecatedConfigurationResponseWithDefaults() *DeprecatedConfigurationResponse {
	this := DeprecatedConfigurationResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DeprecatedConfigurationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeprecatedConfigurationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeprecatedConfigurationResponse) SetId(v string) {
	o.Id = v
}

// GetServer returns the Server field value
func (o *DeprecatedConfigurationResponse) GetServer() DeprecatedServerResponse {
	if o == nil {
		var ret DeprecatedServerResponse
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *DeprecatedConfigurationResponse) GetServerOk() (*DeprecatedServerResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *DeprecatedConfigurationResponse) SetServer(v DeprecatedServerResponse) {
	o.Server = v
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *DeprecatedConfigurationResponse) GetMappings() CloudLdapMappingsResponse {
	if o == nil || o.Mappings == nil {
		var ret CloudLdapMappingsResponse
		return ret
	}
	return *o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedConfigurationResponse) GetMappingsOk() (*CloudLdapMappingsResponse, bool) {
	if o == nil || o.Mappings == nil {
		return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *DeprecatedConfigurationResponse) HasMappings() bool {
	if o != nil && o.Mappings != nil {
		return true
	}

	return false
}

// SetMappings gets a reference to the given CloudLdapMappingsResponse and assigns it to the Mappings field.
func (o *DeprecatedConfigurationResponse) SetMappings(v CloudLdapMappingsResponse) {
	o.Mappings = &v
}

func (o DeprecatedConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["server"] = o.Server
	}
	if o.Mappings != nil {
		toSerialize["mappings"] = o.Mappings
	}
	return json.Marshal(toSerialize)
}

type NullableDeprecatedConfigurationResponse struct {
	value *DeprecatedConfigurationResponse
	isSet bool
}

func (v NullableDeprecatedConfigurationResponse) Get() *DeprecatedConfigurationResponse {
	return v.value
}

func (v *NullableDeprecatedConfigurationResponse) Set(val *DeprecatedConfigurationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeprecatedConfigurationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeprecatedConfigurationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeprecatedConfigurationResponse(val *DeprecatedConfigurationResponse) *NullableDeprecatedConfigurationResponse {
	return &NullableDeprecatedConfigurationResponse{value: val, isSet: true}
}

func (v NullableDeprecatedConfigurationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeprecatedConfigurationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


