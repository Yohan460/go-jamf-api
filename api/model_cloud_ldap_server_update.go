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

// checks if the CloudLdapServerUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudLdapServerUpdate{}

// CloudLdapServerUpdate A Cloud Identity Provider LDAP server configuration for updates
type CloudLdapServerUpdate struct {
	ServerUrl string `json:"serverUrl"`
	Enabled bool `json:"enabled"`
	DomainName string `json:"domainName"`
	Port int64 `json:"port"`
	Keystore *CloudLdapKeystoreFile `json:"keystore,omitempty"`
	ConnectionTimeout int64 `json:"connectionTimeout"`
	SearchTimeout int64 `json:"searchTimeout"`
	UseWildcards bool `json:"useWildcards"`
	ConnectionType string `json:"connectionType"`
	MembershipCalculationOptimizationEnabled *bool `json:"membershipCalculationOptimizationEnabled,omitempty"`
}

type _CloudLdapServerUpdate CloudLdapServerUpdate

// NewCloudLdapServerUpdate instantiates a new CloudLdapServerUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudLdapServerUpdate(serverUrl string, enabled bool, domainName string, port int64, connectionTimeout int64, searchTimeout int64, useWildcards bool, connectionType string) *CloudLdapServerUpdate {
	this := CloudLdapServerUpdate{}
	this.ServerUrl = serverUrl
	this.Enabled = enabled
	this.DomainName = domainName
	this.Port = port
	this.ConnectionTimeout = connectionTimeout
	this.SearchTimeout = searchTimeout
	this.UseWildcards = useWildcards
	this.ConnectionType = connectionType
	return &this
}

// NewCloudLdapServerUpdateWithDefaults instantiates a new CloudLdapServerUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudLdapServerUpdateWithDefaults() *CloudLdapServerUpdate {
	this := CloudLdapServerUpdate{}
	return &this
}

// GetServerUrl returns the ServerUrl field value
func (o *CloudLdapServerUpdate) GetServerUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerUrl
}

// GetServerUrlOk returns a tuple with the ServerUrl field value
// and a boolean to check if the value has been set.
func (o *CloudLdapServerUpdate) GetServerUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerUrl, true
}

// SetServerUrl sets field value
func (o *CloudLdapServerUpdate) SetServerUrl(v string) {
	o.ServerUrl = v
}

// GetEnabled returns the Enabled field value
func (o *CloudLdapServerUpdate) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CloudLdapServerUpdate) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CloudLdapServerUpdate) SetEnabled(v bool) {
	o.Enabled = v
}

// GetDomainName returns the DomainName field value
func (o *CloudLdapServerUpdate) GetDomainName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value
// and a boolean to check if the value has been set.
func (o *CloudLdapServerUpdate) GetDomainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainName, true
}

// SetDomainName sets field value
func (o *CloudLdapServerUpdate) SetDomainName(v string) {
	o.DomainName = v
}

// GetPort returns the Port field value
func (o *CloudLdapServerUpdate) GetPort() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *CloudLdapServerUpdate) GetPortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *CloudLdapServerUpdate) SetPort(v int64) {
	o.Port = v
}

// GetKeystore returns the Keystore field value if set, zero value otherwise.
func (o *CloudLdapServerUpdate) GetKeystore() CloudLdapKeystoreFile {
	if o == nil || IsNil(o.Keystore) {
		var ret CloudLdapKeystoreFile
		return ret
	}
	return *o.Keystore
}

// GetKeystoreOk returns a tuple with the Keystore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapServerUpdate) GetKeystoreOk() (*CloudLdapKeystoreFile, bool) {
	if o == nil || IsNil(o.Keystore) {
		return nil, false
	}
	return o.Keystore, true
}

// HasKeystore returns a boolean if a field has been set.
func (o *CloudLdapServerUpdate) HasKeystore() bool {
	if o != nil && !IsNil(o.Keystore) {
		return true
	}

	return false
}

// SetKeystore gets a reference to the given CloudLdapKeystoreFile and assigns it to the Keystore field.
func (o *CloudLdapServerUpdate) SetKeystore(v CloudLdapKeystoreFile) {
	o.Keystore = &v
}

// GetConnectionTimeout returns the ConnectionTimeout field value
func (o *CloudLdapServerUpdate) GetConnectionTimeout() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ConnectionTimeout
}

// GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field value
// and a boolean to check if the value has been set.
func (o *CloudLdapServerUpdate) GetConnectionTimeoutOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionTimeout, true
}

// SetConnectionTimeout sets field value
func (o *CloudLdapServerUpdate) SetConnectionTimeout(v int64) {
	o.ConnectionTimeout = v
}

// GetSearchTimeout returns the SearchTimeout field value
func (o *CloudLdapServerUpdate) GetSearchTimeout() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SearchTimeout
}

// GetSearchTimeoutOk returns a tuple with the SearchTimeout field value
// and a boolean to check if the value has been set.
func (o *CloudLdapServerUpdate) GetSearchTimeoutOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchTimeout, true
}

// SetSearchTimeout sets field value
func (o *CloudLdapServerUpdate) SetSearchTimeout(v int64) {
	o.SearchTimeout = v
}

// GetUseWildcards returns the UseWildcards field value
func (o *CloudLdapServerUpdate) GetUseWildcards() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseWildcards
}

// GetUseWildcardsOk returns a tuple with the UseWildcards field value
// and a boolean to check if the value has been set.
func (o *CloudLdapServerUpdate) GetUseWildcardsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseWildcards, true
}

// SetUseWildcards sets field value
func (o *CloudLdapServerUpdate) SetUseWildcards(v bool) {
	o.UseWildcards = v
}

// GetConnectionType returns the ConnectionType field value
func (o *CloudLdapServerUpdate) GetConnectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *CloudLdapServerUpdate) GetConnectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *CloudLdapServerUpdate) SetConnectionType(v string) {
	o.ConnectionType = v
}

// GetMembershipCalculationOptimizationEnabled returns the MembershipCalculationOptimizationEnabled field value if set, zero value otherwise.
func (o *CloudLdapServerUpdate) GetMembershipCalculationOptimizationEnabled() bool {
	if o == nil || IsNil(o.MembershipCalculationOptimizationEnabled) {
		var ret bool
		return ret
	}
	return *o.MembershipCalculationOptimizationEnabled
}

// GetMembershipCalculationOptimizationEnabledOk returns a tuple with the MembershipCalculationOptimizationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapServerUpdate) GetMembershipCalculationOptimizationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MembershipCalculationOptimizationEnabled) {
		return nil, false
	}
	return o.MembershipCalculationOptimizationEnabled, true
}

// HasMembershipCalculationOptimizationEnabled returns a boolean if a field has been set.
func (o *CloudLdapServerUpdate) HasMembershipCalculationOptimizationEnabled() bool {
	if o != nil && !IsNil(o.MembershipCalculationOptimizationEnabled) {
		return true
	}

	return false
}

// SetMembershipCalculationOptimizationEnabled gets a reference to the given bool and assigns it to the MembershipCalculationOptimizationEnabled field.
func (o *CloudLdapServerUpdate) SetMembershipCalculationOptimizationEnabled(v bool) {
	o.MembershipCalculationOptimizationEnabled = &v
}

func (o CloudLdapServerUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudLdapServerUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serverUrl"] = o.ServerUrl
	toSerialize["enabled"] = o.Enabled
	toSerialize["domainName"] = o.DomainName
	toSerialize["port"] = o.Port
	if !IsNil(o.Keystore) {
		toSerialize["keystore"] = o.Keystore
	}
	toSerialize["connectionTimeout"] = o.ConnectionTimeout
	toSerialize["searchTimeout"] = o.SearchTimeout
	toSerialize["useWildcards"] = o.UseWildcards
	toSerialize["connectionType"] = o.ConnectionType
	if !IsNil(o.MembershipCalculationOptimizationEnabled) {
		toSerialize["membershipCalculationOptimizationEnabled"] = o.MembershipCalculationOptimizationEnabled
	}
	return toSerialize, nil
}

func (o *CloudLdapServerUpdate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"serverUrl",
		"enabled",
		"domainName",
		"port",
		"connectionTimeout",
		"searchTimeout",
		"useWildcards",
		"connectionType",
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

	varCloudLdapServerUpdate := _CloudLdapServerUpdate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudLdapServerUpdate)

	if err != nil {
		return err
	}

	*o = CloudLdapServerUpdate(varCloudLdapServerUpdate)

	return err
}

type NullableCloudLdapServerUpdate struct {
	value *CloudLdapServerUpdate
	isSet bool
}

func (v NullableCloudLdapServerUpdate) Get() *CloudLdapServerUpdate {
	return v.value
}

func (v *NullableCloudLdapServerUpdate) Set(val *CloudLdapServerUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudLdapServerUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudLdapServerUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudLdapServerUpdate(val *CloudLdapServerUpdate) *NullableCloudLdapServerUpdate {
	return &NullableCloudLdapServerUpdate{value: val, isSet: true}
}

func (v NullableCloudLdapServerUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudLdapServerUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


