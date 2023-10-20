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

// checks if the FileAttachmentV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileAttachmentV2{}

// FileAttachmentV2 struct for FileAttachmentV2
type FileAttachmentV2 struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewFileAttachmentV2 instantiates a new FileAttachmentV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileAttachmentV2() *FileAttachmentV2 {
	this := FileAttachmentV2{}
	return &this
}

// NewFileAttachmentV2WithDefaults instantiates a new FileAttachmentV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileAttachmentV2WithDefaults() *FileAttachmentV2 {
	this := FileAttachmentV2{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FileAttachmentV2) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileAttachmentV2) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FileAttachmentV2) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FileAttachmentV2) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FileAttachmentV2) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileAttachmentV2) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FileAttachmentV2) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FileAttachmentV2) SetName(v string) {
	o.Name = &v
}

func (o FileAttachmentV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileAttachmentV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableFileAttachmentV2 struct {
	value *FileAttachmentV2
	isSet bool
}

func (v NullableFileAttachmentV2) Get() *FileAttachmentV2 {
	return v.value
}

func (v *NullableFileAttachmentV2) Set(val *FileAttachmentV2) {
	v.value = val
	v.isSet = true
}

func (v NullableFileAttachmentV2) IsSet() bool {
	return v.isSet
}

func (v *NullableFileAttachmentV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileAttachmentV2(val *FileAttachmentV2) *NullableFileAttachmentV2 {
	return &NullableFileAttachmentV2{value: val, isSet: true}
}

func (v NullableFileAttachmentV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileAttachmentV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


