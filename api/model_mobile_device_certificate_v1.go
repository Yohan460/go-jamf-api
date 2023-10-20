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

// checks if the MobileDeviceCertificateV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MobileDeviceCertificateV1{}

// MobileDeviceCertificateV1 struct for MobileDeviceCertificateV1
type MobileDeviceCertificateV1 struct {
	CommonName *string `json:"commonName,omitempty"`
	IsIdentity *bool `json:"isIdentity,omitempty"`
}

// NewMobileDeviceCertificateV1 instantiates a new MobileDeviceCertificateV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileDeviceCertificateV1() *MobileDeviceCertificateV1 {
	this := MobileDeviceCertificateV1{}
	return &this
}

// NewMobileDeviceCertificateV1WithDefaults instantiates a new MobileDeviceCertificateV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileDeviceCertificateV1WithDefaults() *MobileDeviceCertificateV1 {
	this := MobileDeviceCertificateV1{}
	return &this
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *MobileDeviceCertificateV1) GetCommonName() string {
	if o == nil || IsNil(o.CommonName) {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceCertificateV1) GetCommonNameOk() (*string, bool) {
	if o == nil || IsNil(o.CommonName) {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *MobileDeviceCertificateV1) HasCommonName() bool {
	if o != nil && !IsNil(o.CommonName) {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *MobileDeviceCertificateV1) SetCommonName(v string) {
	o.CommonName = &v
}

// GetIsIdentity returns the IsIdentity field value if set, zero value otherwise.
func (o *MobileDeviceCertificateV1) GetIsIdentity() bool {
	if o == nil || IsNil(o.IsIdentity) {
		var ret bool
		return ret
	}
	return *o.IsIdentity
}

// GetIsIdentityOk returns a tuple with the IsIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceCertificateV1) GetIsIdentityOk() (*bool, bool) {
	if o == nil || IsNil(o.IsIdentity) {
		return nil, false
	}
	return o.IsIdentity, true
}

// HasIsIdentity returns a boolean if a field has been set.
func (o *MobileDeviceCertificateV1) HasIsIdentity() bool {
	if o != nil && !IsNil(o.IsIdentity) {
		return true
	}

	return false
}

// SetIsIdentity gets a reference to the given bool and assigns it to the IsIdentity field.
func (o *MobileDeviceCertificateV1) SetIsIdentity(v bool) {
	o.IsIdentity = &v
}

func (o MobileDeviceCertificateV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MobileDeviceCertificateV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommonName) {
		toSerialize["commonName"] = o.CommonName
	}
	if !IsNil(o.IsIdentity) {
		toSerialize["isIdentity"] = o.IsIdentity
	}
	return toSerialize, nil
}

type NullableMobileDeviceCertificateV1 struct {
	value *MobileDeviceCertificateV1
	isSet bool
}

func (v NullableMobileDeviceCertificateV1) Get() *MobileDeviceCertificateV1 {
	return v.value
}

func (v *NullableMobileDeviceCertificateV1) Set(val *MobileDeviceCertificateV1) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileDeviceCertificateV1) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileDeviceCertificateV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileDeviceCertificateV1(val *MobileDeviceCertificateV1) *NullableMobileDeviceCertificateV1 {
	return &NullableMobileDeviceCertificateV1{value: val, isSet: true}
}

func (v NullableMobileDeviceCertificateV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileDeviceCertificateV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


