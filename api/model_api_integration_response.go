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

// checks if the ApiIntegrationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiIntegrationResponse{}

// ApiIntegrationResponse struct for ApiIntegrationResponse
type ApiIntegrationResponse struct {
	Id int32 `json:"id"`
	AuthorizationScopes []string `json:"authorizationScopes"`
	DisplayName string `json:"displayName"`
	Enabled bool `json:"enabled"`
	AccessTokenLifetimeSeconds *int32 `json:"accessTokenLifetimeSeconds,omitempty"`
	// Type of API Client:     * `CLIENT_CREDENTIALS` - A client ID and secret have been generated for this integration.     * `NATIVE_APP_OAUTH` - A native app (i.e., Jamf Reset) has been linked to this integration for auth code grant type via Managed App Config.     * `NONE` - No client is currently associated with this integration. 
	AppType string `json:"appType"`
	ClientId string `json:"clientId"`
}

// NewApiIntegrationResponse instantiates a new ApiIntegrationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiIntegrationResponse(id int32, authorizationScopes []string, displayName string, enabled bool, appType string, clientId string) *ApiIntegrationResponse {
	this := ApiIntegrationResponse{}
	this.Id = id
	this.AuthorizationScopes = authorizationScopes
	this.DisplayName = displayName
	this.Enabled = enabled
	this.AppType = appType
	this.ClientId = clientId
	return &this
}

// NewApiIntegrationResponseWithDefaults instantiates a new ApiIntegrationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiIntegrationResponseWithDefaults() *ApiIntegrationResponse {
	this := ApiIntegrationResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ApiIntegrationResponse) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiIntegrationResponse) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiIntegrationResponse) SetId(v int32) {
	o.Id = v
}

// GetAuthorizationScopes returns the AuthorizationScopes field value
func (o *ApiIntegrationResponse) GetAuthorizationScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AuthorizationScopes
}

// GetAuthorizationScopesOk returns a tuple with the AuthorizationScopes field value
// and a boolean to check if the value has been set.
func (o *ApiIntegrationResponse) GetAuthorizationScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizationScopes, true
}

// SetAuthorizationScopes sets field value
func (o *ApiIntegrationResponse) SetAuthorizationScopes(v []string) {
	o.AuthorizationScopes = v
}

// GetDisplayName returns the DisplayName field value
func (o *ApiIntegrationResponse) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ApiIntegrationResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ApiIntegrationResponse) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetEnabled returns the Enabled field value
func (o *ApiIntegrationResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ApiIntegrationResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ApiIntegrationResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAccessTokenLifetimeSeconds returns the AccessTokenLifetimeSeconds field value if set, zero value otherwise.
func (o *ApiIntegrationResponse) GetAccessTokenLifetimeSeconds() int32 {
	if o == nil || IsNil(o.AccessTokenLifetimeSeconds) {
		var ret int32
		return ret
	}
	return *o.AccessTokenLifetimeSeconds
}

// GetAccessTokenLifetimeSecondsOk returns a tuple with the AccessTokenLifetimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiIntegrationResponse) GetAccessTokenLifetimeSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.AccessTokenLifetimeSeconds) {
		return nil, false
	}
	return o.AccessTokenLifetimeSeconds, true
}

// HasAccessTokenLifetimeSeconds returns a boolean if a field has been set.
func (o *ApiIntegrationResponse) HasAccessTokenLifetimeSeconds() bool {
	if o != nil && !IsNil(o.AccessTokenLifetimeSeconds) {
		return true
	}

	return false
}

// SetAccessTokenLifetimeSeconds gets a reference to the given int32 and assigns it to the AccessTokenLifetimeSeconds field.
func (o *ApiIntegrationResponse) SetAccessTokenLifetimeSeconds(v int32) {
	o.AccessTokenLifetimeSeconds = &v
}

// GetAppType returns the AppType field value
func (o *ApiIntegrationResponse) GetAppType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppType
}

// GetAppTypeOk returns a tuple with the AppType field value
// and a boolean to check if the value has been set.
func (o *ApiIntegrationResponse) GetAppTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppType, true
}

// SetAppType sets field value
func (o *ApiIntegrationResponse) SetAppType(v string) {
	o.AppType = v
}

// GetClientId returns the ClientId field value
func (o *ApiIntegrationResponse) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *ApiIntegrationResponse) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *ApiIntegrationResponse) SetClientId(v string) {
	o.ClientId = v
}

func (o ApiIntegrationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiIntegrationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["authorizationScopes"] = o.AuthorizationScopes
	toSerialize["displayName"] = o.DisplayName
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.AccessTokenLifetimeSeconds) {
		toSerialize["accessTokenLifetimeSeconds"] = o.AccessTokenLifetimeSeconds
	}
	toSerialize["appType"] = o.AppType
	toSerialize["clientId"] = o.ClientId
	return toSerialize, nil
}

type NullableApiIntegrationResponse struct {
	value *ApiIntegrationResponse
	isSet bool
}

func (v NullableApiIntegrationResponse) Get() *ApiIntegrationResponse {
	return v.value
}

func (v *NullableApiIntegrationResponse) Set(val *ApiIntegrationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiIntegrationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiIntegrationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiIntegrationResponse(val *ApiIntegrationResponse) *NullableApiIntegrationResponse {
	return &NullableApiIntegrationResponse{value: val, isSet: true}
}

func (v NullableApiIntegrationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiIntegrationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


