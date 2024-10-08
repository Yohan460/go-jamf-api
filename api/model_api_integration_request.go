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

// checks if the ApiIntegrationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiIntegrationRequest{}

// ApiIntegrationRequest struct for ApiIntegrationRequest
type ApiIntegrationRequest struct {
	// API Role display names. 
	AuthorizationScopes []string `json:"authorizationScopes"`
	DisplayName string `json:"displayName"`
	Enabled *bool `json:"enabled,omitempty"`
	AccessTokenLifetimeSeconds *int64 `json:"accessTokenLifetimeSeconds,omitempty"`
}

type _ApiIntegrationRequest ApiIntegrationRequest

// NewApiIntegrationRequest instantiates a new ApiIntegrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiIntegrationRequest(authorizationScopes []string, displayName string) *ApiIntegrationRequest {
	this := ApiIntegrationRequest{}
	this.AuthorizationScopes = authorizationScopes
	this.DisplayName = displayName
	return &this
}

// NewApiIntegrationRequestWithDefaults instantiates a new ApiIntegrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiIntegrationRequestWithDefaults() *ApiIntegrationRequest {
	this := ApiIntegrationRequest{}
	return &this
}

// GetAuthorizationScopes returns the AuthorizationScopes field value
func (o *ApiIntegrationRequest) GetAuthorizationScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AuthorizationScopes
}

// GetAuthorizationScopesOk returns a tuple with the AuthorizationScopes field value
// and a boolean to check if the value has been set.
func (o *ApiIntegrationRequest) GetAuthorizationScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizationScopes, true
}

// SetAuthorizationScopes sets field value
func (o *ApiIntegrationRequest) SetAuthorizationScopes(v []string) {
	o.AuthorizationScopes = v
}

// GetDisplayName returns the DisplayName field value
func (o *ApiIntegrationRequest) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ApiIntegrationRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ApiIntegrationRequest) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApiIntegrationRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiIntegrationRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApiIntegrationRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApiIntegrationRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAccessTokenLifetimeSeconds returns the AccessTokenLifetimeSeconds field value if set, zero value otherwise.
func (o *ApiIntegrationRequest) GetAccessTokenLifetimeSeconds() int64 {
	if o == nil || IsNil(o.AccessTokenLifetimeSeconds) {
		var ret int64
		return ret
	}
	return *o.AccessTokenLifetimeSeconds
}

// GetAccessTokenLifetimeSecondsOk returns a tuple with the AccessTokenLifetimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiIntegrationRequest) GetAccessTokenLifetimeSecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.AccessTokenLifetimeSeconds) {
		return nil, false
	}
	return o.AccessTokenLifetimeSeconds, true
}

// HasAccessTokenLifetimeSeconds returns a boolean if a field has been set.
func (o *ApiIntegrationRequest) HasAccessTokenLifetimeSeconds() bool {
	if o != nil && !IsNil(o.AccessTokenLifetimeSeconds) {
		return true
	}

	return false
}

// SetAccessTokenLifetimeSeconds gets a reference to the given int64 and assigns it to the AccessTokenLifetimeSeconds field.
func (o *ApiIntegrationRequest) SetAccessTokenLifetimeSeconds(v int64) {
	o.AccessTokenLifetimeSeconds = &v
}

func (o ApiIntegrationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiIntegrationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authorizationScopes"] = o.AuthorizationScopes
	toSerialize["displayName"] = o.DisplayName
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.AccessTokenLifetimeSeconds) {
		toSerialize["accessTokenLifetimeSeconds"] = o.AccessTokenLifetimeSeconds
	}
	return toSerialize, nil
}

func (o *ApiIntegrationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authorizationScopes",
		"displayName",
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

	varApiIntegrationRequest := _ApiIntegrationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiIntegrationRequest)

	if err != nil {
		return err
	}

	*o = ApiIntegrationRequest(varApiIntegrationRequest)

	return err
}

type NullableApiIntegrationRequest struct {
	value *ApiIntegrationRequest
	isSet bool
}

func (v NullableApiIntegrationRequest) Get() *ApiIntegrationRequest {
	return v.value
}

func (v *NullableApiIntegrationRequest) Set(val *ApiIntegrationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiIntegrationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiIntegrationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiIntegrationRequest(val *ApiIntegrationRequest) *NullableApiIntegrationRequest {
	return &NullableApiIntegrationRequest{value: val, isSet: true}
}

func (v NullableApiIntegrationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiIntegrationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


