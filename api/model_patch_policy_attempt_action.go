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

// PatchPolicyAttemptAction struct for PatchPolicyAttemptAction
type PatchPolicyAttemptAction struct {
	Id *int32 `json:"id,omitempty"`
	ActionOrder *int32 `json:"actionOrder,omitempty"`
	Action *string `json:"action,omitempty"`
}

// NewPatchPolicyAttemptAction instantiates a new PatchPolicyAttemptAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchPolicyAttemptAction() *PatchPolicyAttemptAction {
	this := PatchPolicyAttemptAction{}
	return &this
}

// NewPatchPolicyAttemptActionWithDefaults instantiates a new PatchPolicyAttemptAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchPolicyAttemptActionWithDefaults() *PatchPolicyAttemptAction {
	this := PatchPolicyAttemptAction{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchPolicyAttemptAction) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPolicyAttemptAction) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchPolicyAttemptAction) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PatchPolicyAttemptAction) SetId(v int32) {
	o.Id = &v
}

// GetActionOrder returns the ActionOrder field value if set, zero value otherwise.
func (o *PatchPolicyAttemptAction) GetActionOrder() int32 {
	if o == nil || o.ActionOrder == nil {
		var ret int32
		return ret
	}
	return *o.ActionOrder
}

// GetActionOrderOk returns a tuple with the ActionOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPolicyAttemptAction) GetActionOrderOk() (*int32, bool) {
	if o == nil || o.ActionOrder == nil {
		return nil, false
	}
	return o.ActionOrder, true
}

// HasActionOrder returns a boolean if a field has been set.
func (o *PatchPolicyAttemptAction) HasActionOrder() bool {
	if o != nil && o.ActionOrder != nil {
		return true
	}

	return false
}

// SetActionOrder gets a reference to the given int32 and assigns it to the ActionOrder field.
func (o *PatchPolicyAttemptAction) SetActionOrder(v int32) {
	o.ActionOrder = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PatchPolicyAttemptAction) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPolicyAttemptAction) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PatchPolicyAttemptAction) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *PatchPolicyAttemptAction) SetAction(v string) {
	o.Action = &v
}

func (o PatchPolicyAttemptAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ActionOrder != nil {
		toSerialize["actionOrder"] = o.ActionOrder
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	return json.Marshal(toSerialize)
}

type NullablePatchPolicyAttemptAction struct {
	value *PatchPolicyAttemptAction
	isSet bool
}

func (v NullablePatchPolicyAttemptAction) Get() *PatchPolicyAttemptAction {
	return v.value
}

func (v *NullablePatchPolicyAttemptAction) Set(val *PatchPolicyAttemptAction) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchPolicyAttemptAction) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchPolicyAttemptAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchPolicyAttemptAction(val *PatchPolicyAttemptAction) *NullablePatchPolicyAttemptAction {
	return &NullablePatchPolicyAttemptAction{value: val, isSet: true}
}

func (v NullablePatchPolicyAttemptAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchPolicyAttemptAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


