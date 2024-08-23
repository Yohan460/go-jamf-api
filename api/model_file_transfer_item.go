/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the FileTransferItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileTransferItem{}

// FileTransferItem struct for FileTransferItem
type FileTransferItem struct {
	FilePath *string `json:"filePath,omitempty"`
	TransferTimestamp *time.Time `json:"transferTimestamp,omitempty"`
	FileTransferType *string `json:"fileTransferType,omitempty"`
}

// NewFileTransferItem instantiates a new FileTransferItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileTransferItem() *FileTransferItem {
	this := FileTransferItem{}
	return &this
}

// NewFileTransferItemWithDefaults instantiates a new FileTransferItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileTransferItemWithDefaults() *FileTransferItem {
	this := FileTransferItem{}
	return &this
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *FileTransferItem) GetFilePath() string {
	if o == nil || IsNil(o.FilePath) {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileTransferItem) GetFilePathOk() (*string, bool) {
	if o == nil || IsNil(o.FilePath) {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *FileTransferItem) HasFilePath() bool {
	if o != nil && !IsNil(o.FilePath) {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *FileTransferItem) SetFilePath(v string) {
	o.FilePath = &v
}

// GetTransferTimestamp returns the TransferTimestamp field value if set, zero value otherwise.
func (o *FileTransferItem) GetTransferTimestamp() time.Time {
	if o == nil || IsNil(o.TransferTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.TransferTimestamp
}

// GetTransferTimestampOk returns a tuple with the TransferTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileTransferItem) GetTransferTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TransferTimestamp) {
		return nil, false
	}
	return o.TransferTimestamp, true
}

// HasTransferTimestamp returns a boolean if a field has been set.
func (o *FileTransferItem) HasTransferTimestamp() bool {
	if o != nil && !IsNil(o.TransferTimestamp) {
		return true
	}

	return false
}

// SetTransferTimestamp gets a reference to the given time.Time and assigns it to the TransferTimestamp field.
func (o *FileTransferItem) SetTransferTimestamp(v time.Time) {
	o.TransferTimestamp = &v
}

// GetFileTransferType returns the FileTransferType field value if set, zero value otherwise.
func (o *FileTransferItem) GetFileTransferType() string {
	if o == nil || IsNil(o.FileTransferType) {
		var ret string
		return ret
	}
	return *o.FileTransferType
}

// GetFileTransferTypeOk returns a tuple with the FileTransferType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileTransferItem) GetFileTransferTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FileTransferType) {
		return nil, false
	}
	return o.FileTransferType, true
}

// HasFileTransferType returns a boolean if a field has been set.
func (o *FileTransferItem) HasFileTransferType() bool {
	if o != nil && !IsNil(o.FileTransferType) {
		return true
	}

	return false
}

// SetFileTransferType gets a reference to the given string and assigns it to the FileTransferType field.
func (o *FileTransferItem) SetFileTransferType(v string) {
	o.FileTransferType = &v
}

func (o FileTransferItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileTransferItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FilePath) {
		toSerialize["filePath"] = o.FilePath
	}
	if !IsNil(o.TransferTimestamp) {
		toSerialize["transferTimestamp"] = o.TransferTimestamp
	}
	if !IsNil(o.FileTransferType) {
		toSerialize["fileTransferType"] = o.FileTransferType
	}
	return toSerialize, nil
}

type NullableFileTransferItem struct {
	value *FileTransferItem
	isSet bool
}

func (v NullableFileTransferItem) Get() *FileTransferItem {
	return v.value
}

func (v *NullableFileTransferItem) Set(val *FileTransferItem) {
	v.value = val
	v.isSet = true
}

func (v NullableFileTransferItem) IsSet() bool {
	return v.isSet
}

func (v *NullableFileTransferItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileTransferItem(val *FileTransferItem) *NullableFileTransferItem {
	return &NullableFileTransferItem{value: val, isSet: true}
}

func (v NullableFileTransferItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileTransferItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


