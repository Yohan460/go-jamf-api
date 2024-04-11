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

// checks if the InventoryPreloadExtensionAttributeColumnResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryPreloadExtensionAttributeColumnResult{}

// InventoryPreloadExtensionAttributeColumnResult struct for InventoryPreloadExtensionAttributeColumnResult
type InventoryPreloadExtensionAttributeColumnResult struct {
	TotalCount *int64 `json:"totalCount,omitempty"`
	Results []InventoryPreloadExtensionAttributeColumn `json:"results,omitempty"`
}

// NewInventoryPreloadExtensionAttributeColumnResult instantiates a new InventoryPreloadExtensionAttributeColumnResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryPreloadExtensionAttributeColumnResult() *InventoryPreloadExtensionAttributeColumnResult {
	this := InventoryPreloadExtensionAttributeColumnResult{}
	return &this
}

// NewInventoryPreloadExtensionAttributeColumnResultWithDefaults instantiates a new InventoryPreloadExtensionAttributeColumnResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryPreloadExtensionAttributeColumnResultWithDefaults() *InventoryPreloadExtensionAttributeColumnResult {
	this := InventoryPreloadExtensionAttributeColumnResult{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *InventoryPreloadExtensionAttributeColumnResult) GetTotalCount() int64 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadExtensionAttributeColumnResult) GetTotalCountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *InventoryPreloadExtensionAttributeColumnResult) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *InventoryPreloadExtensionAttributeColumnResult) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *InventoryPreloadExtensionAttributeColumnResult) GetResults() []InventoryPreloadExtensionAttributeColumn {
	if o == nil || IsNil(o.Results) {
		var ret []InventoryPreloadExtensionAttributeColumn
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadExtensionAttributeColumnResult) GetResultsOk() ([]InventoryPreloadExtensionAttributeColumn, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *InventoryPreloadExtensionAttributeColumnResult) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []InventoryPreloadExtensionAttributeColumn and assigns it to the Results field.
func (o *InventoryPreloadExtensionAttributeColumnResult) SetResults(v []InventoryPreloadExtensionAttributeColumn) {
	o.Results = v
}

func (o InventoryPreloadExtensionAttributeColumnResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryPreloadExtensionAttributeColumnResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableInventoryPreloadExtensionAttributeColumnResult struct {
	value *InventoryPreloadExtensionAttributeColumnResult
	isSet bool
}

func (v NullableInventoryPreloadExtensionAttributeColumnResult) Get() *InventoryPreloadExtensionAttributeColumnResult {
	return v.value
}

func (v *NullableInventoryPreloadExtensionAttributeColumnResult) Set(val *InventoryPreloadExtensionAttributeColumnResult) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryPreloadExtensionAttributeColumnResult) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryPreloadExtensionAttributeColumnResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryPreloadExtensionAttributeColumnResult(val *InventoryPreloadExtensionAttributeColumnResult) *NullableInventoryPreloadExtensionAttributeColumnResult {
	return &NullableInventoryPreloadExtensionAttributeColumnResult{value: val, isSet: true}
}

func (v NullableInventoryPreloadExtensionAttributeColumnResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryPreloadExtensionAttributeColumnResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


