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

// DeprecatedServerResponse An old Cloud Identity Provider LDAP server configuration for responses
type DeprecatedServerResponse struct {
	Id *string `json:"id,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	ProviderName *string `json:"providerName,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ServerUrl *string `json:"serverUrl,omitempty"`
	DomainName *string `json:"domainName,omitempty"`
	Port *int32 `json:"port,omitempty"`
	Keystore *CloudLdapKeystore `json:"keystore,omitempty"`
	ConnectionTimeout *int32 `json:"connectionTimeout,omitempty"`
	SearchTimeout *int32 `json:"searchTimeout,omitempty"`
	UseWildcards *bool `json:"useWildcards,omitempty"`
	ConnectionType *string `json:"connectionType,omitempty"`
}

// NewDeprecatedServerResponse instantiates a new DeprecatedServerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeprecatedServerResponse() *DeprecatedServerResponse {
	this := DeprecatedServerResponse{}
	return &this
}

// NewDeprecatedServerResponseWithDefaults instantiates a new DeprecatedServerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeprecatedServerResponseWithDefaults() *DeprecatedServerResponse {
	this := DeprecatedServerResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeprecatedServerResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedServerResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeprecatedServerResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeprecatedServerResponse) SetId(v string) {
	o.Id = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DeprecatedServerResponse) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedServerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DeprecatedServerResponse) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DeprecatedServerResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *DeprecatedServerResponse) GetProviderName() string {
	if o == nil || o.ProviderName == nil {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedServerResponse) GetProviderNameOk() (*string, bool) {
	if o == nil || o.ProviderName == nil {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *DeprecatedServerResponse) HasProviderName() bool {
	if o != nil && o.ProviderName != nil {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *DeprecatedServerResponse) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *DeprecatedServerResponse) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedServerResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *DeprecatedServerResponse) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *DeprecatedServerResponse) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetServerUrl returns the ServerUrl field value if set, zero value otherwise.
func (o *DeprecatedServerResponse) GetServerUrl() string {
	if o == nil || o.ServerUrl == nil {
		var ret string
		return ret
	}
	return *o.ServerUrl
}

// GetServerUrlOk returns a tuple with the ServerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedServerResponse) GetServerUrlOk() (*string, bool) {
	if o == nil || o.ServerUrl == nil {
		return nil, false
	}
	return o.ServerUrl, true
}

// HasServerUrl returns a boolean if a field has been set.
func (o *DeprecatedServerResponse) HasServerUrl() bool {
	if o != nil && o.ServerUrl != nil {
		return true
	}

	return false
}

// SetServerUrl gets a reference to the given string and assigns it to the ServerUrl field.
func (o *DeprecatedServerResponse) SetServerUrl(v string) {
	o.ServerUrl = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *DeprecatedServerResponse) GetDomainName() string {
	if o == nil || o.DomainName == nil {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedServerResponse) GetDomainNameOk() (*string, bool) {
	if o == nil || o.DomainName == nil {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *DeprecatedServerResponse) HasDomainName() bool {
	if o != nil && o.DomainName != nil {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *DeprecatedServerResponse) SetDomainName(v string) {
	o.DomainName = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *DeprecatedServerResponse) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedServerResponse) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *DeprecatedServerResponse) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *DeprecatedServerResponse) SetPort(v int32) {
	o.Port = &v
}

// GetKeystore returns the Keystore field value if set, zero value otherwise.
func (o *DeprecatedServerResponse) GetKeystore() CloudLdapKeystore {
	if o == nil || o.Keystore == nil {
		var ret CloudLdapKeystore
		return ret
	}
	return *o.Keystore
}

// GetKeystoreOk returns a tuple with the Keystore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedServerResponse) GetKeystoreOk() (*CloudLdapKeystore, bool) {
	if o == nil || o.Keystore == nil {
		return nil, false
	}
	return o.Keystore, true
}

// HasKeystore returns a boolean if a field has been set.
func (o *DeprecatedServerResponse) HasKeystore() bool {
	if o != nil && o.Keystore != nil {
		return true
	}

	return false
}

// SetKeystore gets a reference to the given CloudLdapKeystore and assigns it to the Keystore field.
func (o *DeprecatedServerResponse) SetKeystore(v CloudLdapKeystore) {
	o.Keystore = &v
}

// GetConnectionTimeout returns the ConnectionTimeout field value if set, zero value otherwise.
func (o *DeprecatedServerResponse) GetConnectionTimeout() int32 {
	if o == nil || o.ConnectionTimeout == nil {
		var ret int32
		return ret
	}
	return *o.ConnectionTimeout
}

// GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedServerResponse) GetConnectionTimeoutOk() (*int32, bool) {
	if o == nil || o.ConnectionTimeout == nil {
		return nil, false
	}
	return o.ConnectionTimeout, true
}

// HasConnectionTimeout returns a boolean if a field has been set.
func (o *DeprecatedServerResponse) HasConnectionTimeout() bool {
	if o != nil && o.ConnectionTimeout != nil {
		return true
	}

	return false
}

// SetConnectionTimeout gets a reference to the given int32 and assigns it to the ConnectionTimeout field.
func (o *DeprecatedServerResponse) SetConnectionTimeout(v int32) {
	o.ConnectionTimeout = &v
}

// GetSearchTimeout returns the SearchTimeout field value if set, zero value otherwise.
func (o *DeprecatedServerResponse) GetSearchTimeout() int32 {
	if o == nil || o.SearchTimeout == nil {
		var ret int32
		return ret
	}
	return *o.SearchTimeout
}

// GetSearchTimeoutOk returns a tuple with the SearchTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedServerResponse) GetSearchTimeoutOk() (*int32, bool) {
	if o == nil || o.SearchTimeout == nil {
		return nil, false
	}
	return o.SearchTimeout, true
}

// HasSearchTimeout returns a boolean if a field has been set.
func (o *DeprecatedServerResponse) HasSearchTimeout() bool {
	if o != nil && o.SearchTimeout != nil {
		return true
	}

	return false
}

// SetSearchTimeout gets a reference to the given int32 and assigns it to the SearchTimeout field.
func (o *DeprecatedServerResponse) SetSearchTimeout(v int32) {
	o.SearchTimeout = &v
}

// GetUseWildcards returns the UseWildcards field value if set, zero value otherwise.
func (o *DeprecatedServerResponse) GetUseWildcards() bool {
	if o == nil || o.UseWildcards == nil {
		var ret bool
		return ret
	}
	return *o.UseWildcards
}

// GetUseWildcardsOk returns a tuple with the UseWildcards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedServerResponse) GetUseWildcardsOk() (*bool, bool) {
	if o == nil || o.UseWildcards == nil {
		return nil, false
	}
	return o.UseWildcards, true
}

// HasUseWildcards returns a boolean if a field has been set.
func (o *DeprecatedServerResponse) HasUseWildcards() bool {
	if o != nil && o.UseWildcards != nil {
		return true
	}

	return false
}

// SetUseWildcards gets a reference to the given bool and assigns it to the UseWildcards field.
func (o *DeprecatedServerResponse) SetUseWildcards(v bool) {
	o.UseWildcards = &v
}

// GetConnectionType returns the ConnectionType field value if set, zero value otherwise.
func (o *DeprecatedServerResponse) GetConnectionType() string {
	if o == nil || o.ConnectionType == nil {
		var ret string
		return ret
	}
	return *o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedServerResponse) GetConnectionTypeOk() (*string, bool) {
	if o == nil || o.ConnectionType == nil {
		return nil, false
	}
	return o.ConnectionType, true
}

// HasConnectionType returns a boolean if a field has been set.
func (o *DeprecatedServerResponse) HasConnectionType() bool {
	if o != nil && o.ConnectionType != nil {
		return true
	}

	return false
}

// SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.
func (o *DeprecatedServerResponse) SetConnectionType(v string) {
	o.ConnectionType = &v
}

func (o DeprecatedServerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ProviderName != nil {
		toSerialize["providerName"] = o.ProviderName
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.ServerUrl != nil {
		toSerialize["serverUrl"] = o.ServerUrl
	}
	if o.DomainName != nil {
		toSerialize["domainName"] = o.DomainName
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Keystore != nil {
		toSerialize["keystore"] = o.Keystore
	}
	if o.ConnectionTimeout != nil {
		toSerialize["connectionTimeout"] = o.ConnectionTimeout
	}
	if o.SearchTimeout != nil {
		toSerialize["searchTimeout"] = o.SearchTimeout
	}
	if o.UseWildcards != nil {
		toSerialize["useWildcards"] = o.UseWildcards
	}
	if o.ConnectionType != nil {
		toSerialize["connectionType"] = o.ConnectionType
	}
	return json.Marshal(toSerialize)
}

type NullableDeprecatedServerResponse struct {
	value *DeprecatedServerResponse
	isSet bool
}

func (v NullableDeprecatedServerResponse) Get() *DeprecatedServerResponse {
	return v.value
}

func (v *NullableDeprecatedServerResponse) Set(val *DeprecatedServerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeprecatedServerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeprecatedServerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeprecatedServerResponse(val *DeprecatedServerResponse) *NullableDeprecatedServerResponse {
	return &NullableDeprecatedServerResponse{value: val, isSet: true}
}

func (v NullableDeprecatedServerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeprecatedServerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

