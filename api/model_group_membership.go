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

// checks if the GroupMembership type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupMembership{}

// GroupMembership struct for GroupMembership
type GroupMembership struct {
	GroupId *string `json:"groupId,omitempty"`
	GroupName *string `json:"groupName,omitempty"`
	// Indicates that group is smart group
	SmartGroup *bool `json:"smartGroup,omitempty"`
}

// NewGroupMembership instantiates a new GroupMembership object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMembership() *GroupMembership {
	this := GroupMembership{}
	return &this
}

// NewGroupMembershipWithDefaults instantiates a new GroupMembership object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMembershipWithDefaults() *GroupMembership {
	this := GroupMembership{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *GroupMembership) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMembership) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *GroupMembership) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *GroupMembership) SetGroupId(v string) {
	o.GroupId = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *GroupMembership) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMembership) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *GroupMembership) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *GroupMembership) SetGroupName(v string) {
	o.GroupName = &v
}

// GetSmartGroup returns the SmartGroup field value if set, zero value otherwise.
func (o *GroupMembership) GetSmartGroup() bool {
	if o == nil || IsNil(o.SmartGroup) {
		var ret bool
		return ret
	}
	return *o.SmartGroup
}

// GetSmartGroupOk returns a tuple with the SmartGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMembership) GetSmartGroupOk() (*bool, bool) {
	if o == nil || IsNil(o.SmartGroup) {
		return nil, false
	}
	return o.SmartGroup, true
}

// HasSmartGroup returns a boolean if a field has been set.
func (o *GroupMembership) HasSmartGroup() bool {
	if o != nil && !IsNil(o.SmartGroup) {
		return true
	}

	return false
}

// SetSmartGroup gets a reference to the given bool and assigns it to the SmartGroup field.
func (o *GroupMembership) SetSmartGroup(v bool) {
	o.SmartGroup = &v
}

func (o GroupMembership) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupMembership) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.GroupName) {
		toSerialize["groupName"] = o.GroupName
	}
	if !IsNil(o.SmartGroup) {
		toSerialize["smartGroup"] = o.SmartGroup
	}
	return toSerialize, nil
}

type NullableGroupMembership struct {
	value *GroupMembership
	isSet bool
}

func (v NullableGroupMembership) Get() *GroupMembership {
	return v.value
}

func (v *NullableGroupMembership) Set(val *GroupMembership) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupMembership) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupMembership) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupMembership(val *GroupMembership) *NullableGroupMembership {
	return &NullableGroupMembership{value: val, isSet: true}
}

func (v NullableGroupMembership) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupMembership) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


