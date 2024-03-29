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

// checks if the NotificationV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationV1{}

// NotificationV1 Jamf Pro notification used for important alerts.
type NotificationV1 struct {
	Type *NotificationType `json:"type,omitempty"`
	Id *string `json:"id,omitempty"`
	Params map[string]map[string]interface{} `json:"params,omitempty"`
}

// NewNotificationV1 instantiates a new NotificationV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationV1() *NotificationV1 {
	this := NotificationV1{}
	return &this
}

// NewNotificationV1WithDefaults instantiates a new NotificationV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationV1WithDefaults() *NotificationV1 {
	this := NotificationV1{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NotificationV1) GetType() NotificationType {
	if o == nil || IsNil(o.Type) {
		var ret NotificationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationV1) GetTypeOk() (*NotificationType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NotificationV1) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given NotificationType and assigns it to the Type field.
func (o *NotificationV1) SetType(v NotificationType) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NotificationV1) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationV1) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NotificationV1) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NotificationV1) SetId(v string) {
	o.Id = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *NotificationV1) GetParams() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Params) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationV1) GetParamsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Params) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *NotificationV1) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string]map[string]interface{} and assigns it to the Params field.
func (o *NotificationV1) SetParams(v map[string]map[string]interface{}) {
	o.Params = v
}

func (o NotificationV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	return toSerialize, nil
}

type NullableNotificationV1 struct {
	value *NotificationV1
	isSet bool
}

func (v NullableNotificationV1) Get() *NotificationV1 {
	return v.value
}

func (v *NullableNotificationV1) Set(val *NotificationV1) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationV1) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationV1(val *NotificationV1) *NullableNotificationV1 {
	return &NullableNotificationV1{value: val, isSet: true}
}

func (v NullableNotificationV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


