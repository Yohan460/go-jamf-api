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

// FileAttachmentDelete struct for FileAttachmentDelete
type FileAttachmentDelete struct {
	Ids []int32 `json:"ids,omitempty"`
}

// NewFileAttachmentDelete instantiates a new FileAttachmentDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileAttachmentDelete() *FileAttachmentDelete {
	this := FileAttachmentDelete{}
	return &this
}

// NewFileAttachmentDeleteWithDefaults instantiates a new FileAttachmentDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileAttachmentDeleteWithDefaults() *FileAttachmentDelete {
	this := FileAttachmentDelete{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *FileAttachmentDelete) GetIds() []int32 {
	if o == nil || o.Ids == nil {
		var ret []int32
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileAttachmentDelete) GetIdsOk() ([]int32, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *FileAttachmentDelete) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int32 and assigns it to the Ids field.
func (o *FileAttachmentDelete) SetIds(v []int32) {
	o.Ids = v
}

func (o FileAttachmentDelete) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	return json.Marshal(toSerialize)
}

type NullableFileAttachmentDelete struct {
	value *FileAttachmentDelete
	isSet bool
}

func (v NullableFileAttachmentDelete) Get() *FileAttachmentDelete {
	return v.value
}

func (v *NullableFileAttachmentDelete) Set(val *FileAttachmentDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableFileAttachmentDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableFileAttachmentDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileAttachmentDelete(val *FileAttachmentDelete) *NullableFileAttachmentDelete {
	return &NullableFileAttachmentDelete{value: val, isSet: true}
}

func (v NullableFileAttachmentDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileAttachmentDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


