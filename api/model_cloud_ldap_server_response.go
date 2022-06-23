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

// CloudLdapServerResponse A Cloud Identity Provider LDAP server configuration for responses
type CloudLdapServerResponse struct {
	Id *string `json:"id,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	ServerUrl *string `json:"serverUrl,omitempty"`
	DomainName *string `json:"domainName,omitempty"`
	Port *int32 `json:"port,omitempty"`
	Keystore *CloudLdapKeystore `json:"keystore,omitempty"`
	ConnectionTimeout *int32 `json:"connectionTimeout,omitempty"`
	SearchTimeout *int32 `json:"searchTimeout,omitempty"`
	UseWildcards *bool `json:"useWildcards,omitempty"`
	ConnectionType *string `json:"connectionType,omitempty"`
}

// NewCloudLdapServerResponse instantiates a new CloudLdapServerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudLdapServerResponse() *CloudLdapServerResponse {
	this := CloudLdapServerResponse{}
	return &this
}

// NewCloudLdapServerResponseWithDefaults instantiates a new CloudLdapServerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudLdapServerResponseWithDefaults() *CloudLdapServerResponse {
	this := CloudLdapServerResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CloudLdapServerResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapServerResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudLdapServerResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CloudLdapServerResponse) SetId(v string) {
	o.Id = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CloudLdapServerResponse) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapServerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CloudLdapServerResponse) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CloudLdapServerResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetServerUrl returns the ServerUrl field value if set, zero value otherwise.
func (o *CloudLdapServerResponse) GetServerUrl() string {
	if o == nil || o.ServerUrl == nil {
		var ret string
		return ret
	}
	return *o.ServerUrl
}

// GetServerUrlOk returns a tuple with the ServerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapServerResponse) GetServerUrlOk() (*string, bool) {
	if o == nil || o.ServerUrl == nil {
		return nil, false
	}
	return o.ServerUrl, true
}

// HasServerUrl returns a boolean if a field has been set.
func (o *CloudLdapServerResponse) HasServerUrl() bool {
	if o != nil && o.ServerUrl != nil {
		return true
	}

	return false
}

// SetServerUrl gets a reference to the given string and assigns it to the ServerUrl field.
func (o *CloudLdapServerResponse) SetServerUrl(v string) {
	o.ServerUrl = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *CloudLdapServerResponse) GetDomainName() string {
	if o == nil || o.DomainName == nil {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapServerResponse) GetDomainNameOk() (*string, bool) {
	if o == nil || o.DomainName == nil {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *CloudLdapServerResponse) HasDomainName() bool {
	if o != nil && o.DomainName != nil {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *CloudLdapServerResponse) SetDomainName(v string) {
	o.DomainName = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *CloudLdapServerResponse) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapServerResponse) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *CloudLdapServerResponse) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *CloudLdapServerResponse) SetPort(v int32) {
	o.Port = &v
}

// GetKeystore returns the Keystore field value if set, zero value otherwise.
func (o *CloudLdapServerResponse) GetKeystore() CloudLdapKeystore {
	if o == nil || o.Keystore == nil {
		var ret CloudLdapKeystore
		return ret
	}
	return *o.Keystore
}

// GetKeystoreOk returns a tuple with the Keystore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapServerResponse) GetKeystoreOk() (*CloudLdapKeystore, bool) {
	if o == nil || o.Keystore == nil {
		return nil, false
	}
	return o.Keystore, true
}

// HasKeystore returns a boolean if a field has been set.
func (o *CloudLdapServerResponse) HasKeystore() bool {
	if o != nil && o.Keystore != nil {
		return true
	}

	return false
}

// SetKeystore gets a reference to the given CloudLdapKeystore and assigns it to the Keystore field.
func (o *CloudLdapServerResponse) SetKeystore(v CloudLdapKeystore) {
	o.Keystore = &v
}

// GetConnectionTimeout returns the ConnectionTimeout field value if set, zero value otherwise.
func (o *CloudLdapServerResponse) GetConnectionTimeout() int32 {
	if o == nil || o.ConnectionTimeout == nil {
		var ret int32
		return ret
	}
	return *o.ConnectionTimeout
}

// GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapServerResponse) GetConnectionTimeoutOk() (*int32, bool) {
	if o == nil || o.ConnectionTimeout == nil {
		return nil, false
	}
	return o.ConnectionTimeout, true
}

// HasConnectionTimeout returns a boolean if a field has been set.
func (o *CloudLdapServerResponse) HasConnectionTimeout() bool {
	if o != nil && o.ConnectionTimeout != nil {
		return true
	}

	return false
}

// SetConnectionTimeout gets a reference to the given int32 and assigns it to the ConnectionTimeout field.
func (o *CloudLdapServerResponse) SetConnectionTimeout(v int32) {
	o.ConnectionTimeout = &v
}

// GetSearchTimeout returns the SearchTimeout field value if set, zero value otherwise.
func (o *CloudLdapServerResponse) GetSearchTimeout() int32 {
	if o == nil || o.SearchTimeout == nil {
		var ret int32
		return ret
	}
	return *o.SearchTimeout
}

// GetSearchTimeoutOk returns a tuple with the SearchTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapServerResponse) GetSearchTimeoutOk() (*int32, bool) {
	if o == nil || o.SearchTimeout == nil {
		return nil, false
	}
	return o.SearchTimeout, true
}

// HasSearchTimeout returns a boolean if a field has been set.
func (o *CloudLdapServerResponse) HasSearchTimeout() bool {
	if o != nil && o.SearchTimeout != nil {
		return true
	}

	return false
}

// SetSearchTimeout gets a reference to the given int32 and assigns it to the SearchTimeout field.
func (o *CloudLdapServerResponse) SetSearchTimeout(v int32) {
	o.SearchTimeout = &v
}

// GetUseWildcards returns the UseWildcards field value if set, zero value otherwise.
func (o *CloudLdapServerResponse) GetUseWildcards() bool {
	if o == nil || o.UseWildcards == nil {
		var ret bool
		return ret
	}
	return *o.UseWildcards
}

// GetUseWildcardsOk returns a tuple with the UseWildcards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapServerResponse) GetUseWildcardsOk() (*bool, bool) {
	if o == nil || o.UseWildcards == nil {
		return nil, false
	}
	return o.UseWildcards, true
}

// HasUseWildcards returns a boolean if a field has been set.
func (o *CloudLdapServerResponse) HasUseWildcards() bool {
	if o != nil && o.UseWildcards != nil {
		return true
	}

	return false
}

// SetUseWildcards gets a reference to the given bool and assigns it to the UseWildcards field.
func (o *CloudLdapServerResponse) SetUseWildcards(v bool) {
	o.UseWildcards = &v
}

// GetConnectionType returns the ConnectionType field value if set, zero value otherwise.
func (o *CloudLdapServerResponse) GetConnectionType() string {
	if o == nil || o.ConnectionType == nil {
		var ret string
		return ret
	}
	return *o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapServerResponse) GetConnectionTypeOk() (*string, bool) {
	if o == nil || o.ConnectionType == nil {
		return nil, false
	}
	return o.ConnectionType, true
}

// HasConnectionType returns a boolean if a field has been set.
func (o *CloudLdapServerResponse) HasConnectionType() bool {
	if o != nil && o.ConnectionType != nil {
		return true
	}

	return false
}

// SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.
func (o *CloudLdapServerResponse) SetConnectionType(v string) {
	o.ConnectionType = &v
}

func (o CloudLdapServerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
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

type NullableCloudLdapServerResponse struct {
	value *CloudLdapServerResponse
	isSet bool
}

func (v NullableCloudLdapServerResponse) Get() *CloudLdapServerResponse {
	return v.value
}

func (v *NullableCloudLdapServerResponse) Set(val *CloudLdapServerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudLdapServerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudLdapServerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudLdapServerResponse(val *CloudLdapServerResponse) *NullableCloudLdapServerResponse {
	return &NullableCloudLdapServerResponse{value: val, isSet: true}
}

func (v NullableCloudLdapServerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudLdapServerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

