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

// checks if the CsaTenantIdInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CsaTenantIdInfo{}

// CsaTenantIdInfo struct for CsaTenantIdInfo
type CsaTenantIdInfo struct {
	// The tenant ID
	TenantId NullableString `json:"tenantId,omitempty"`
}

// NewCsaTenantIdInfo instantiates a new CsaTenantIdInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCsaTenantIdInfo() *CsaTenantIdInfo {
	this := CsaTenantIdInfo{}
	return &this
}

// NewCsaTenantIdInfoWithDefaults instantiates a new CsaTenantIdInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCsaTenantIdInfoWithDefaults() *CsaTenantIdInfo {
	this := CsaTenantIdInfo{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CsaTenantIdInfo) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CsaTenantIdInfo) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *CsaTenantIdInfo) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *CsaTenantIdInfo) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *CsaTenantIdInfo) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *CsaTenantIdInfo) UnsetTenantId() {
	o.TenantId.Unset()
}

func (o CsaTenantIdInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CsaTenantIdInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	return toSerialize, nil
}

type NullableCsaTenantIdInfo struct {
	value *CsaTenantIdInfo
	isSet bool
}

func (v NullableCsaTenantIdInfo) Get() *CsaTenantIdInfo {
	return v.value
}

func (v *NullableCsaTenantIdInfo) Set(val *CsaTenantIdInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCsaTenantIdInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCsaTenantIdInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCsaTenantIdInfo(val *CsaTenantIdInfo) *NullableCsaTenantIdInfo {
	return &NullableCsaTenantIdInfo{value: val, isSet: true}
}

func (v NullableCsaTenantIdInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCsaTenantIdInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


