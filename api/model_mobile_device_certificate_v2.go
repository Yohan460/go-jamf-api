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

// checks if the MobileDeviceCertificateV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MobileDeviceCertificateV2{}

// MobileDeviceCertificateV2 struct for MobileDeviceCertificateV2
type MobileDeviceCertificateV2 struct {
	CommonName *string `json:"commonName,omitempty"`
	Identity *bool `json:"identity,omitempty"`
	ExpirationDateEpoch *time.Time `json:"expirationDateEpoch,omitempty"`
	SubjectName *string `json:"subjectName,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty"`
	Sha1Fingerprint *string `json:"sha1Fingerprint,omitempty"`
	IssuedDateEpoch *string `json:"issuedDateEpoch,omitempty"`
	CertificateStatus *string `json:"certificateStatus,omitempty"`
	LifecycleStatus *string `json:"lifecycleStatus,omitempty"`
}

// NewMobileDeviceCertificateV2 instantiates a new MobileDeviceCertificateV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileDeviceCertificateV2() *MobileDeviceCertificateV2 {
	this := MobileDeviceCertificateV2{}
	return &this
}

// NewMobileDeviceCertificateV2WithDefaults instantiates a new MobileDeviceCertificateV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileDeviceCertificateV2WithDefaults() *MobileDeviceCertificateV2 {
	this := MobileDeviceCertificateV2{}
	return &this
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *MobileDeviceCertificateV2) GetCommonName() string {
	if o == nil || IsNil(o.CommonName) {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceCertificateV2) GetCommonNameOk() (*string, bool) {
	if o == nil || IsNil(o.CommonName) {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *MobileDeviceCertificateV2) HasCommonName() bool {
	if o != nil && !IsNil(o.CommonName) {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *MobileDeviceCertificateV2) SetCommonName(v string) {
	o.CommonName = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *MobileDeviceCertificateV2) GetIdentity() bool {
	if o == nil || IsNil(o.Identity) {
		var ret bool
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceCertificateV2) GetIdentityOk() (*bool, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *MobileDeviceCertificateV2) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given bool and assigns it to the Identity field.
func (o *MobileDeviceCertificateV2) SetIdentity(v bool) {
	o.Identity = &v
}

// GetExpirationDateEpoch returns the ExpirationDateEpoch field value if set, zero value otherwise.
func (o *MobileDeviceCertificateV2) GetExpirationDateEpoch() time.Time {
	if o == nil || IsNil(o.ExpirationDateEpoch) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDateEpoch
}

// GetExpirationDateEpochOk returns a tuple with the ExpirationDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceCertificateV2) GetExpirationDateEpochOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationDateEpoch) {
		return nil, false
	}
	return o.ExpirationDateEpoch, true
}

// HasExpirationDateEpoch returns a boolean if a field has been set.
func (o *MobileDeviceCertificateV2) HasExpirationDateEpoch() bool {
	if o != nil && !IsNil(o.ExpirationDateEpoch) {
		return true
	}

	return false
}

// SetExpirationDateEpoch gets a reference to the given time.Time and assigns it to the ExpirationDateEpoch field.
func (o *MobileDeviceCertificateV2) SetExpirationDateEpoch(v time.Time) {
	o.ExpirationDateEpoch = &v
}

// GetSubjectName returns the SubjectName field value if set, zero value otherwise.
func (o *MobileDeviceCertificateV2) GetSubjectName() string {
	if o == nil || IsNil(o.SubjectName) {
		var ret string
		return ret
	}
	return *o.SubjectName
}

// GetSubjectNameOk returns a tuple with the SubjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceCertificateV2) GetSubjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectName) {
		return nil, false
	}
	return o.SubjectName, true
}

// HasSubjectName returns a boolean if a field has been set.
func (o *MobileDeviceCertificateV2) HasSubjectName() bool {
	if o != nil && !IsNil(o.SubjectName) {
		return true
	}

	return false
}

// SetSubjectName gets a reference to the given string and assigns it to the SubjectName field.
func (o *MobileDeviceCertificateV2) SetSubjectName(v string) {
	o.SubjectName = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *MobileDeviceCertificateV2) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceCertificateV2) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *MobileDeviceCertificateV2) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *MobileDeviceCertificateV2) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSha1Fingerprint returns the Sha1Fingerprint field value if set, zero value otherwise.
func (o *MobileDeviceCertificateV2) GetSha1Fingerprint() string {
	if o == nil || IsNil(o.Sha1Fingerprint) {
		var ret string
		return ret
	}
	return *o.Sha1Fingerprint
}

// GetSha1FingerprintOk returns a tuple with the Sha1Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceCertificateV2) GetSha1FingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.Sha1Fingerprint) {
		return nil, false
	}
	return o.Sha1Fingerprint, true
}

// HasSha1Fingerprint returns a boolean if a field has been set.
func (o *MobileDeviceCertificateV2) HasSha1Fingerprint() bool {
	if o != nil && !IsNil(o.Sha1Fingerprint) {
		return true
	}

	return false
}

// SetSha1Fingerprint gets a reference to the given string and assigns it to the Sha1Fingerprint field.
func (o *MobileDeviceCertificateV2) SetSha1Fingerprint(v string) {
	o.Sha1Fingerprint = &v
}

// GetIssuedDateEpoch returns the IssuedDateEpoch field value if set, zero value otherwise.
func (o *MobileDeviceCertificateV2) GetIssuedDateEpoch() string {
	if o == nil || IsNil(o.IssuedDateEpoch) {
		var ret string
		return ret
	}
	return *o.IssuedDateEpoch
}

// GetIssuedDateEpochOk returns a tuple with the IssuedDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceCertificateV2) GetIssuedDateEpochOk() (*string, bool) {
	if o == nil || IsNil(o.IssuedDateEpoch) {
		return nil, false
	}
	return o.IssuedDateEpoch, true
}

// HasIssuedDateEpoch returns a boolean if a field has been set.
func (o *MobileDeviceCertificateV2) HasIssuedDateEpoch() bool {
	if o != nil && !IsNil(o.IssuedDateEpoch) {
		return true
	}

	return false
}

// SetIssuedDateEpoch gets a reference to the given string and assigns it to the IssuedDateEpoch field.
func (o *MobileDeviceCertificateV2) SetIssuedDateEpoch(v string) {
	o.IssuedDateEpoch = &v
}

// GetCertificateStatus returns the CertificateStatus field value if set, zero value otherwise.
func (o *MobileDeviceCertificateV2) GetCertificateStatus() string {
	if o == nil || IsNil(o.CertificateStatus) {
		var ret string
		return ret
	}
	return *o.CertificateStatus
}

// GetCertificateStatusOk returns a tuple with the CertificateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceCertificateV2) GetCertificateStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateStatus) {
		return nil, false
	}
	return o.CertificateStatus, true
}

// HasCertificateStatus returns a boolean if a field has been set.
func (o *MobileDeviceCertificateV2) HasCertificateStatus() bool {
	if o != nil && !IsNil(o.CertificateStatus) {
		return true
	}

	return false
}

// SetCertificateStatus gets a reference to the given string and assigns it to the CertificateStatus field.
func (o *MobileDeviceCertificateV2) SetCertificateStatus(v string) {
	o.CertificateStatus = &v
}

// GetLifecycleStatus returns the LifecycleStatus field value if set, zero value otherwise.
func (o *MobileDeviceCertificateV2) GetLifecycleStatus() string {
	if o == nil || IsNil(o.LifecycleStatus) {
		var ret string
		return ret
	}
	return *o.LifecycleStatus
}

// GetLifecycleStatusOk returns a tuple with the LifecycleStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceCertificateV2) GetLifecycleStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LifecycleStatus) {
		return nil, false
	}
	return o.LifecycleStatus, true
}

// HasLifecycleStatus returns a boolean if a field has been set.
func (o *MobileDeviceCertificateV2) HasLifecycleStatus() bool {
	if o != nil && !IsNil(o.LifecycleStatus) {
		return true
	}

	return false
}

// SetLifecycleStatus gets a reference to the given string and assigns it to the LifecycleStatus field.
func (o *MobileDeviceCertificateV2) SetLifecycleStatus(v string) {
	o.LifecycleStatus = &v
}

func (o MobileDeviceCertificateV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MobileDeviceCertificateV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommonName) {
		toSerialize["commonName"] = o.CommonName
	}
	if !IsNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !IsNil(o.ExpirationDateEpoch) {
		toSerialize["expirationDateEpoch"] = o.ExpirationDateEpoch
	}
	if !IsNil(o.SubjectName) {
		toSerialize["subjectName"] = o.SubjectName
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if !IsNil(o.Sha1Fingerprint) {
		toSerialize["sha1Fingerprint"] = o.Sha1Fingerprint
	}
	if !IsNil(o.IssuedDateEpoch) {
		toSerialize["issuedDateEpoch"] = o.IssuedDateEpoch
	}
	if !IsNil(o.CertificateStatus) {
		toSerialize["certificateStatus"] = o.CertificateStatus
	}
	if !IsNil(o.LifecycleStatus) {
		toSerialize["lifecycleStatus"] = o.LifecycleStatus
	}
	return toSerialize, nil
}

type NullableMobileDeviceCertificateV2 struct {
	value *MobileDeviceCertificateV2
	isSet bool
}

func (v NullableMobileDeviceCertificateV2) Get() *MobileDeviceCertificateV2 {
	return v.value
}

func (v *NullableMobileDeviceCertificateV2) Set(val *MobileDeviceCertificateV2) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileDeviceCertificateV2) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileDeviceCertificateV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileDeviceCertificateV2(val *MobileDeviceCertificateV2) *NullableMobileDeviceCertificateV2 {
	return &NullableMobileDeviceCertificateV2{value: val, isSet: true}
}

func (v NullableMobileDeviceCertificateV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileDeviceCertificateV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


