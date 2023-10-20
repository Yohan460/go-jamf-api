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

// checks if the MobileDeviceCertificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MobileDeviceCertificate{}

// MobileDeviceCertificate struct for MobileDeviceCertificate
type MobileDeviceCertificate struct {
	CommonName *string `json:"commonName,omitempty"`
	Identity *bool `json:"identity,omitempty"`
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
}

// NewMobileDeviceCertificate instantiates a new MobileDeviceCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileDeviceCertificate() *MobileDeviceCertificate {
	this := MobileDeviceCertificate{}
	return &this
}

// NewMobileDeviceCertificateWithDefaults instantiates a new MobileDeviceCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileDeviceCertificateWithDefaults() *MobileDeviceCertificate {
	this := MobileDeviceCertificate{}
	return &this
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *MobileDeviceCertificate) GetCommonName() string {
	if o == nil || IsNil(o.CommonName) {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceCertificate) GetCommonNameOk() (*string, bool) {
	if o == nil || IsNil(o.CommonName) {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *MobileDeviceCertificate) HasCommonName() bool {
	if o != nil && !IsNil(o.CommonName) {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *MobileDeviceCertificate) SetCommonName(v string) {
	o.CommonName = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *MobileDeviceCertificate) GetIdentity() bool {
	if o == nil || IsNil(o.Identity) {
		var ret bool
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceCertificate) GetIdentityOk() (*bool, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *MobileDeviceCertificate) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given bool and assigns it to the Identity field.
func (o *MobileDeviceCertificate) SetIdentity(v bool) {
	o.Identity = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *MobileDeviceCertificate) GetExpirationDate() time.Time {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceCertificate) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *MobileDeviceCertificate) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *MobileDeviceCertificate) SetExpirationDate(v time.Time) {
	o.ExpirationDate = &v
}

func (o MobileDeviceCertificate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MobileDeviceCertificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommonName) {
		toSerialize["commonName"] = o.CommonName
	}
	if !IsNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	return toSerialize, nil
}

type NullableMobileDeviceCertificate struct {
	value *MobileDeviceCertificate
	isSet bool
}

func (v NullableMobileDeviceCertificate) Get() *MobileDeviceCertificate {
	return v.value
}

func (v *NullableMobileDeviceCertificate) Set(val *MobileDeviceCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileDeviceCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileDeviceCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileDeviceCertificate(val *MobileDeviceCertificate) *NullableMobileDeviceCertificate {
	return &NullableMobileDeviceCertificate{value: val, isSet: true}
}

func (v NullableMobileDeviceCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileDeviceCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


