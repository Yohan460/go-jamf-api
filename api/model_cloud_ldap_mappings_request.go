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

// checks if the CloudLdapMappingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudLdapMappingsRequest{}

// CloudLdapMappingsRequest Mappings configurations request for Ldap Cloud Identity Provider configuration
type CloudLdapMappingsRequest struct {
	UserMappings UserMappings `json:"userMappings"`
	GroupMappings GroupMappings `json:"groupMappings"`
	MembershipMappings MembershipMappings `json:"membershipMappings"`
}

type _CloudLdapMappingsRequest CloudLdapMappingsRequest

// NewCloudLdapMappingsRequest instantiates a new CloudLdapMappingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudLdapMappingsRequest(userMappings UserMappings, groupMappings GroupMappings, membershipMappings MembershipMappings) *CloudLdapMappingsRequest {
	this := CloudLdapMappingsRequest{}
	this.UserMappings = userMappings
	this.GroupMappings = groupMappings
	this.MembershipMappings = membershipMappings
	return &this
}

// NewCloudLdapMappingsRequestWithDefaults instantiates a new CloudLdapMappingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudLdapMappingsRequestWithDefaults() *CloudLdapMappingsRequest {
	this := CloudLdapMappingsRequest{}
	return &this
}

// GetUserMappings returns the UserMappings field value
func (o *CloudLdapMappingsRequest) GetUserMappings() UserMappings {
	if o == nil {
		var ret UserMappings
		return ret
	}

	return o.UserMappings
}

// GetUserMappingsOk returns a tuple with the UserMappings field value
// and a boolean to check if the value has been set.
func (o *CloudLdapMappingsRequest) GetUserMappingsOk() (*UserMappings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserMappings, true
}

// SetUserMappings sets field value
func (o *CloudLdapMappingsRequest) SetUserMappings(v UserMappings) {
	o.UserMappings = v
}

// GetGroupMappings returns the GroupMappings field value
func (o *CloudLdapMappingsRequest) GetGroupMappings() GroupMappings {
	if o == nil {
		var ret GroupMappings
		return ret
	}

	return o.GroupMappings
}

// GetGroupMappingsOk returns a tuple with the GroupMappings field value
// and a boolean to check if the value has been set.
func (o *CloudLdapMappingsRequest) GetGroupMappingsOk() (*GroupMappings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupMappings, true
}

// SetGroupMappings sets field value
func (o *CloudLdapMappingsRequest) SetGroupMappings(v GroupMappings) {
	o.GroupMappings = v
}

// GetMembershipMappings returns the MembershipMappings field value
func (o *CloudLdapMappingsRequest) GetMembershipMappings() MembershipMappings {
	if o == nil {
		var ret MembershipMappings
		return ret
	}

	return o.MembershipMappings
}

// GetMembershipMappingsOk returns a tuple with the MembershipMappings field value
// and a boolean to check if the value has been set.
func (o *CloudLdapMappingsRequest) GetMembershipMappingsOk() (*MembershipMappings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MembershipMappings, true
}

// SetMembershipMappings sets field value
func (o *CloudLdapMappingsRequest) SetMembershipMappings(v MembershipMappings) {
	o.MembershipMappings = v
}

func (o CloudLdapMappingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudLdapMappingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userMappings"] = o.UserMappings
	toSerialize["groupMappings"] = o.GroupMappings
	toSerialize["membershipMappings"] = o.MembershipMappings
	return toSerialize, nil
}

func (o *CloudLdapMappingsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userMappings",
		"groupMappings",
		"membershipMappings",
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

	varCloudLdapMappingsRequest := _CloudLdapMappingsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudLdapMappingsRequest)

	if err != nil {
		return err
	}

	*o = CloudLdapMappingsRequest(varCloudLdapMappingsRequest)

	return err
}

type NullableCloudLdapMappingsRequest struct {
	value *CloudLdapMappingsRequest
	isSet bool
}

func (v NullableCloudLdapMappingsRequest) Get() *CloudLdapMappingsRequest {
	return v.value
}

func (v *NullableCloudLdapMappingsRequest) Set(val *CloudLdapMappingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudLdapMappingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudLdapMappingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudLdapMappingsRequest(val *CloudLdapMappingsRequest) *NullableCloudLdapMappingsRequest {
	return &NullableCloudLdapMappingsRequest{value: val, isSet: true}
}

func (v NullableCloudLdapMappingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudLdapMappingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


