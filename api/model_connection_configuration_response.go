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

// checks if the ConnectionConfigurationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionConfigurationResponse{}

// ConnectionConfigurationResponse Response that contains information about connection configuration for Team Viewer
type ConnectionConfigurationResponse struct {
	// An identifier of connection configuration for Team Viewer Remote Administration
	Id string `json:"id"`
	// An identifier of a site which Team Viewer Remote Administration is configured on
	SiteId string `json:"siteId"`
	// Name for Team Viewer Connection Configuration
	DisplayName string `json:"displayName"`
	// Describes if Team Viewer connection is enabled or disabled
	Enabled bool `json:"enabled"`
	// Number of minutes before the session expires
	SessionTimeout NullableInt64 `json:"sessionTimeout"`
}

type _ConnectionConfigurationResponse ConnectionConfigurationResponse

// NewConnectionConfigurationResponse instantiates a new ConnectionConfigurationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionConfigurationResponse(id string, siteId string, displayName string, enabled bool, sessionTimeout NullableInt64) *ConnectionConfigurationResponse {
	this := ConnectionConfigurationResponse{}
	this.Id = id
	this.SiteId = siteId
	this.DisplayName = displayName
	this.Enabled = enabled
	this.SessionTimeout = sessionTimeout
	return &this
}

// NewConnectionConfigurationResponseWithDefaults instantiates a new ConnectionConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionConfigurationResponseWithDefaults() *ConnectionConfigurationResponse {
	this := ConnectionConfigurationResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ConnectionConfigurationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConnectionConfigurationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConnectionConfigurationResponse) SetId(v string) {
	o.Id = v
}

// GetSiteId returns the SiteId field value
func (o *ConnectionConfigurationResponse) GetSiteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *ConnectionConfigurationResponse) GetSiteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *ConnectionConfigurationResponse) SetSiteId(v string) {
	o.SiteId = v
}

// GetDisplayName returns the DisplayName field value
func (o *ConnectionConfigurationResponse) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ConnectionConfigurationResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ConnectionConfigurationResponse) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetEnabled returns the Enabled field value
func (o *ConnectionConfigurationResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ConnectionConfigurationResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ConnectionConfigurationResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSessionTimeout returns the SessionTimeout field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *ConnectionConfigurationResponse) GetSessionTimeout() int64 {
	if o == nil || o.SessionTimeout.Get() == nil {
		var ret int64
		return ret
	}

	return *o.SessionTimeout.Get()
}

// GetSessionTimeoutOk returns a tuple with the SessionTimeout field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionConfigurationResponse) GetSessionTimeoutOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionTimeout.Get(), o.SessionTimeout.IsSet()
}

// SetSessionTimeout sets field value
func (o *ConnectionConfigurationResponse) SetSessionTimeout(v int64) {
	o.SessionTimeout.Set(&v)
}

func (o ConnectionConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionConfigurationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["siteId"] = o.SiteId
	toSerialize["displayName"] = o.DisplayName
	toSerialize["enabled"] = o.Enabled
	toSerialize["sessionTimeout"] = o.SessionTimeout.Get()
	return toSerialize, nil
}

func (o *ConnectionConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"siteId",
		"displayName",
		"enabled",
		"sessionTimeout",
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

	varConnectionConfigurationResponse := _ConnectionConfigurationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConnectionConfigurationResponse)

	if err != nil {
		return err
	}

	*o = ConnectionConfigurationResponse(varConnectionConfigurationResponse)

	return err
}

type NullableConnectionConfigurationResponse struct {
	value *ConnectionConfigurationResponse
	isSet bool
}

func (v NullableConnectionConfigurationResponse) Get() *ConnectionConfigurationResponse {
	return v.value
}

func (v *NullableConnectionConfigurationResponse) Set(val *ConnectionConfigurationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionConfigurationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionConfigurationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionConfigurationResponse(val *ConnectionConfigurationResponse) *NullableConnectionConfigurationResponse {
	return &NullableConnectionConfigurationResponse{value: val, isSet: true}
}

func (v NullableConnectionConfigurationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionConfigurationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


