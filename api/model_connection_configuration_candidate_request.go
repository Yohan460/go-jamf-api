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

// ConnectionConfigurationCandidateRequest Request that creates configuration and initialize connection between Jamf Pro and Team Viewer
type ConnectionConfigurationCandidateRequest struct {
	// An identifier of a site which Team Viewer Remote Administration will be configured on
	SiteId string `json:"siteId"`
	// Name for Team Viewer Connection Configuration
	DisplayName string `json:"displayName"`
	// Token which is used for connecting to Team Viewer
	ScriptToken string `json:"scriptToken"`
	// Defines the intent to enable or disable Team Viewer connection
	Enabled bool `json:"enabled"`
	// Number of minutes before the session expires
	SessionTimeout int32 `json:"sessionTimeout"`
}

// NewConnectionConfigurationCandidateRequest instantiates a new ConnectionConfigurationCandidateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionConfigurationCandidateRequest(siteId string, displayName string, scriptToken string, enabled bool, sessionTimeout int32) *ConnectionConfigurationCandidateRequest {
	this := ConnectionConfigurationCandidateRequest{}
	this.SiteId = siteId
	this.DisplayName = displayName
	this.ScriptToken = scriptToken
	this.Enabled = enabled
	this.SessionTimeout = sessionTimeout
	return &this
}

// NewConnectionConfigurationCandidateRequestWithDefaults instantiates a new ConnectionConfigurationCandidateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionConfigurationCandidateRequestWithDefaults() *ConnectionConfigurationCandidateRequest {
	this := ConnectionConfigurationCandidateRequest{}
	return &this
}

// GetSiteId returns the SiteId field value
func (o *ConnectionConfigurationCandidateRequest) GetSiteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *ConnectionConfigurationCandidateRequest) GetSiteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *ConnectionConfigurationCandidateRequest) SetSiteId(v string) {
	o.SiteId = v
}

// GetDisplayName returns the DisplayName field value
func (o *ConnectionConfigurationCandidateRequest) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ConnectionConfigurationCandidateRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ConnectionConfigurationCandidateRequest) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetScriptToken returns the ScriptToken field value
func (o *ConnectionConfigurationCandidateRequest) GetScriptToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScriptToken
}

// GetScriptTokenOk returns a tuple with the ScriptToken field value
// and a boolean to check if the value has been set.
func (o *ConnectionConfigurationCandidateRequest) GetScriptTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptToken, true
}

// SetScriptToken sets field value
func (o *ConnectionConfigurationCandidateRequest) SetScriptToken(v string) {
	o.ScriptToken = v
}

// GetEnabled returns the Enabled field value
func (o *ConnectionConfigurationCandidateRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ConnectionConfigurationCandidateRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ConnectionConfigurationCandidateRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSessionTimeout returns the SessionTimeout field value
func (o *ConnectionConfigurationCandidateRequest) GetSessionTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SessionTimeout
}

// GetSessionTimeoutOk returns a tuple with the SessionTimeout field value
// and a boolean to check if the value has been set.
func (o *ConnectionConfigurationCandidateRequest) GetSessionTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionTimeout, true
}

// SetSessionTimeout sets field value
func (o *ConnectionConfigurationCandidateRequest) SetSessionTimeout(v int32) {
	o.SessionTimeout = v
}

func (o ConnectionConfigurationCandidateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["siteId"] = o.SiteId
	}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["scriptToken"] = o.ScriptToken
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["sessionTimeout"] = o.SessionTimeout
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionConfigurationCandidateRequest struct {
	value *ConnectionConfigurationCandidateRequest
	isSet bool
}

func (v NullableConnectionConfigurationCandidateRequest) Get() *ConnectionConfigurationCandidateRequest {
	return v.value
}

func (v *NullableConnectionConfigurationCandidateRequest) Set(val *ConnectionConfigurationCandidateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionConfigurationCandidateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionConfigurationCandidateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionConfigurationCandidateRequest(val *ConnectionConfigurationCandidateRequest) *NullableConnectionConfigurationCandidateRequest {
	return &NullableConnectionConfigurationCandidateRequest{value: val, isSet: true}
}

func (v NullableConnectionConfigurationCandidateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionConfigurationCandidateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


