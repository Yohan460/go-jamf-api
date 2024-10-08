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

// checks if the LdapConfigurationUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdapConfigurationUpdate{}

// LdapConfigurationUpdate A Cloud Identity Provider LDAP configuration for updates
type LdapConfigurationUpdate struct {
	CloudIdPCommon CloudIdPCommon `json:"cloudIdPCommon"`
	Server CloudLdapServerUpdate `json:"server"`
	Mappings *CloudLdapMappingsRequest `json:"mappings,omitempty"`
}

type _LdapConfigurationUpdate LdapConfigurationUpdate

// NewLdapConfigurationUpdate instantiates a new LdapConfigurationUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapConfigurationUpdate(cloudIdPCommon CloudIdPCommon, server CloudLdapServerUpdate) *LdapConfigurationUpdate {
	this := LdapConfigurationUpdate{}
	this.CloudIdPCommon = cloudIdPCommon
	this.Server = server
	return &this
}

// NewLdapConfigurationUpdateWithDefaults instantiates a new LdapConfigurationUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapConfigurationUpdateWithDefaults() *LdapConfigurationUpdate {
	this := LdapConfigurationUpdate{}
	return &this
}

// GetCloudIdPCommon returns the CloudIdPCommon field value
func (o *LdapConfigurationUpdate) GetCloudIdPCommon() CloudIdPCommon {
	if o == nil {
		var ret CloudIdPCommon
		return ret
	}

	return o.CloudIdPCommon
}

// GetCloudIdPCommonOk returns a tuple with the CloudIdPCommon field value
// and a boolean to check if the value has been set.
func (o *LdapConfigurationUpdate) GetCloudIdPCommonOk() (*CloudIdPCommon, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudIdPCommon, true
}

// SetCloudIdPCommon sets field value
func (o *LdapConfigurationUpdate) SetCloudIdPCommon(v CloudIdPCommon) {
	o.CloudIdPCommon = v
}

// GetServer returns the Server field value
func (o *LdapConfigurationUpdate) GetServer() CloudLdapServerUpdate {
	if o == nil {
		var ret CloudLdapServerUpdate
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *LdapConfigurationUpdate) GetServerOk() (*CloudLdapServerUpdate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *LdapConfigurationUpdate) SetServer(v CloudLdapServerUpdate) {
	o.Server = v
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *LdapConfigurationUpdate) GetMappings() CloudLdapMappingsRequest {
	if o == nil || IsNil(o.Mappings) {
		var ret CloudLdapMappingsRequest
		return ret
	}
	return *o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapConfigurationUpdate) GetMappingsOk() (*CloudLdapMappingsRequest, bool) {
	if o == nil || IsNil(o.Mappings) {
		return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *LdapConfigurationUpdate) HasMappings() bool {
	if o != nil && !IsNil(o.Mappings) {
		return true
	}

	return false
}

// SetMappings gets a reference to the given CloudLdapMappingsRequest and assigns it to the Mappings field.
func (o *LdapConfigurationUpdate) SetMappings(v CloudLdapMappingsRequest) {
	o.Mappings = &v
}

func (o LdapConfigurationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdapConfigurationUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudIdPCommon"] = o.CloudIdPCommon
	toSerialize["server"] = o.Server
	if !IsNil(o.Mappings) {
		toSerialize["mappings"] = o.Mappings
	}
	return toSerialize, nil
}

func (o *LdapConfigurationUpdate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloudIdPCommon",
		"server",
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

	varLdapConfigurationUpdate := _LdapConfigurationUpdate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLdapConfigurationUpdate)

	if err != nil {
		return err
	}

	*o = LdapConfigurationUpdate(varLdapConfigurationUpdate)

	return err
}

type NullableLdapConfigurationUpdate struct {
	value *LdapConfigurationUpdate
	isSet bool
}

func (v NullableLdapConfigurationUpdate) Get() *LdapConfigurationUpdate {
	return v.value
}

func (v *NullableLdapConfigurationUpdate) Set(val *LdapConfigurationUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapConfigurationUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapConfigurationUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapConfigurationUpdate(val *LdapConfigurationUpdate) *NullableLdapConfigurationUpdate {
	return &NullableLdapConfigurationUpdate{value: val, isSet: true}
}

func (v NullableLdapConfigurationUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapConfigurationUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


