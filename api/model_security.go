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

// checks if the Security type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Security{}

// Security struct for Security
type Security struct {
	IsDataProtected *bool `json:"isDataProtected,omitempty"`
	IsBlockLevelEncryptionCapable *bool `json:"isBlockLevelEncryptionCapable,omitempty"`
	IsFileLevelEncryptionCapable *bool `json:"isFileLevelEncryptionCapable,omitempty"`
	IsPasscodePresent *bool `json:"isPasscodePresent,omitempty"`
	IsPasscodeCompliant *bool `json:"isPasscodeCompliant,omitempty"`
	IsPasscodeCompliantWithProfile *bool `json:"isPasscodeCompliantWithProfile,omitempty"`
	HardwareEncryption *int32 `json:"hardwareEncryption,omitempty"`
	IsActivationLockEnabled *bool `json:"isActivationLockEnabled,omitempty"`
	IsJailBreakDetected *bool `json:"isJailBreakDetected,omitempty"`
}

// NewSecurity instantiates a new Security object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurity() *Security {
	this := Security{}
	return &this
}

// NewSecurityWithDefaults instantiates a new Security object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityWithDefaults() *Security {
	this := Security{}
	return &this
}

// GetIsDataProtected returns the IsDataProtected field value if set, zero value otherwise.
func (o *Security) GetIsDataProtected() bool {
	if o == nil || IsNil(o.IsDataProtected) {
		var ret bool
		return ret
	}
	return *o.IsDataProtected
}

// GetIsDataProtectedOk returns a tuple with the IsDataProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Security) GetIsDataProtectedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDataProtected) {
		return nil, false
	}
	return o.IsDataProtected, true
}

// HasIsDataProtected returns a boolean if a field has been set.
func (o *Security) HasIsDataProtected() bool {
	if o != nil && !IsNil(o.IsDataProtected) {
		return true
	}

	return false
}

// SetIsDataProtected gets a reference to the given bool and assigns it to the IsDataProtected field.
func (o *Security) SetIsDataProtected(v bool) {
	o.IsDataProtected = &v
}

// GetIsBlockLevelEncryptionCapable returns the IsBlockLevelEncryptionCapable field value if set, zero value otherwise.
func (o *Security) GetIsBlockLevelEncryptionCapable() bool {
	if o == nil || IsNil(o.IsBlockLevelEncryptionCapable) {
		var ret bool
		return ret
	}
	return *o.IsBlockLevelEncryptionCapable
}

// GetIsBlockLevelEncryptionCapableOk returns a tuple with the IsBlockLevelEncryptionCapable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Security) GetIsBlockLevelEncryptionCapableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBlockLevelEncryptionCapable) {
		return nil, false
	}
	return o.IsBlockLevelEncryptionCapable, true
}

// HasIsBlockLevelEncryptionCapable returns a boolean if a field has been set.
func (o *Security) HasIsBlockLevelEncryptionCapable() bool {
	if o != nil && !IsNil(o.IsBlockLevelEncryptionCapable) {
		return true
	}

	return false
}

// SetIsBlockLevelEncryptionCapable gets a reference to the given bool and assigns it to the IsBlockLevelEncryptionCapable field.
func (o *Security) SetIsBlockLevelEncryptionCapable(v bool) {
	o.IsBlockLevelEncryptionCapable = &v
}

// GetIsFileLevelEncryptionCapable returns the IsFileLevelEncryptionCapable field value if set, zero value otherwise.
func (o *Security) GetIsFileLevelEncryptionCapable() bool {
	if o == nil || IsNil(o.IsFileLevelEncryptionCapable) {
		var ret bool
		return ret
	}
	return *o.IsFileLevelEncryptionCapable
}

// GetIsFileLevelEncryptionCapableOk returns a tuple with the IsFileLevelEncryptionCapable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Security) GetIsFileLevelEncryptionCapableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFileLevelEncryptionCapable) {
		return nil, false
	}
	return o.IsFileLevelEncryptionCapable, true
}

// HasIsFileLevelEncryptionCapable returns a boolean if a field has been set.
func (o *Security) HasIsFileLevelEncryptionCapable() bool {
	if o != nil && !IsNil(o.IsFileLevelEncryptionCapable) {
		return true
	}

	return false
}

// SetIsFileLevelEncryptionCapable gets a reference to the given bool and assigns it to the IsFileLevelEncryptionCapable field.
func (o *Security) SetIsFileLevelEncryptionCapable(v bool) {
	o.IsFileLevelEncryptionCapable = &v
}

// GetIsPasscodePresent returns the IsPasscodePresent field value if set, zero value otherwise.
func (o *Security) GetIsPasscodePresent() bool {
	if o == nil || IsNil(o.IsPasscodePresent) {
		var ret bool
		return ret
	}
	return *o.IsPasscodePresent
}

// GetIsPasscodePresentOk returns a tuple with the IsPasscodePresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Security) GetIsPasscodePresentOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPasscodePresent) {
		return nil, false
	}
	return o.IsPasscodePresent, true
}

// HasIsPasscodePresent returns a boolean if a field has been set.
func (o *Security) HasIsPasscodePresent() bool {
	if o != nil && !IsNil(o.IsPasscodePresent) {
		return true
	}

	return false
}

// SetIsPasscodePresent gets a reference to the given bool and assigns it to the IsPasscodePresent field.
func (o *Security) SetIsPasscodePresent(v bool) {
	o.IsPasscodePresent = &v
}

// GetIsPasscodeCompliant returns the IsPasscodeCompliant field value if set, zero value otherwise.
func (o *Security) GetIsPasscodeCompliant() bool {
	if o == nil || IsNil(o.IsPasscodeCompliant) {
		var ret bool
		return ret
	}
	return *o.IsPasscodeCompliant
}

// GetIsPasscodeCompliantOk returns a tuple with the IsPasscodeCompliant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Security) GetIsPasscodeCompliantOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPasscodeCompliant) {
		return nil, false
	}
	return o.IsPasscodeCompliant, true
}

// HasIsPasscodeCompliant returns a boolean if a field has been set.
func (o *Security) HasIsPasscodeCompliant() bool {
	if o != nil && !IsNil(o.IsPasscodeCompliant) {
		return true
	}

	return false
}

// SetIsPasscodeCompliant gets a reference to the given bool and assigns it to the IsPasscodeCompliant field.
func (o *Security) SetIsPasscodeCompliant(v bool) {
	o.IsPasscodeCompliant = &v
}

// GetIsPasscodeCompliantWithProfile returns the IsPasscodeCompliantWithProfile field value if set, zero value otherwise.
func (o *Security) GetIsPasscodeCompliantWithProfile() bool {
	if o == nil || IsNil(o.IsPasscodeCompliantWithProfile) {
		var ret bool
		return ret
	}
	return *o.IsPasscodeCompliantWithProfile
}

// GetIsPasscodeCompliantWithProfileOk returns a tuple with the IsPasscodeCompliantWithProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Security) GetIsPasscodeCompliantWithProfileOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPasscodeCompliantWithProfile) {
		return nil, false
	}
	return o.IsPasscodeCompliantWithProfile, true
}

// HasIsPasscodeCompliantWithProfile returns a boolean if a field has been set.
func (o *Security) HasIsPasscodeCompliantWithProfile() bool {
	if o != nil && !IsNil(o.IsPasscodeCompliantWithProfile) {
		return true
	}

	return false
}

// SetIsPasscodeCompliantWithProfile gets a reference to the given bool and assigns it to the IsPasscodeCompliantWithProfile field.
func (o *Security) SetIsPasscodeCompliantWithProfile(v bool) {
	o.IsPasscodeCompliantWithProfile = &v
}

// GetHardwareEncryption returns the HardwareEncryption field value if set, zero value otherwise.
func (o *Security) GetHardwareEncryption() int32 {
	if o == nil || IsNil(o.HardwareEncryption) {
		var ret int32
		return ret
	}
	return *o.HardwareEncryption
}

// GetHardwareEncryptionOk returns a tuple with the HardwareEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Security) GetHardwareEncryptionOk() (*int32, bool) {
	if o == nil || IsNil(o.HardwareEncryption) {
		return nil, false
	}
	return o.HardwareEncryption, true
}

// HasHardwareEncryption returns a boolean if a field has been set.
func (o *Security) HasHardwareEncryption() bool {
	if o != nil && !IsNil(o.HardwareEncryption) {
		return true
	}

	return false
}

// SetHardwareEncryption gets a reference to the given int32 and assigns it to the HardwareEncryption field.
func (o *Security) SetHardwareEncryption(v int32) {
	o.HardwareEncryption = &v
}

// GetIsActivationLockEnabled returns the IsActivationLockEnabled field value if set, zero value otherwise.
func (o *Security) GetIsActivationLockEnabled() bool {
	if o == nil || IsNil(o.IsActivationLockEnabled) {
		var ret bool
		return ret
	}
	return *o.IsActivationLockEnabled
}

// GetIsActivationLockEnabledOk returns a tuple with the IsActivationLockEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Security) GetIsActivationLockEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActivationLockEnabled) {
		return nil, false
	}
	return o.IsActivationLockEnabled, true
}

// HasIsActivationLockEnabled returns a boolean if a field has been set.
func (o *Security) HasIsActivationLockEnabled() bool {
	if o != nil && !IsNil(o.IsActivationLockEnabled) {
		return true
	}

	return false
}

// SetIsActivationLockEnabled gets a reference to the given bool and assigns it to the IsActivationLockEnabled field.
func (o *Security) SetIsActivationLockEnabled(v bool) {
	o.IsActivationLockEnabled = &v
}

// GetIsJailBreakDetected returns the IsJailBreakDetected field value if set, zero value otherwise.
func (o *Security) GetIsJailBreakDetected() bool {
	if o == nil || IsNil(o.IsJailBreakDetected) {
		var ret bool
		return ret
	}
	return *o.IsJailBreakDetected
}

// GetIsJailBreakDetectedOk returns a tuple with the IsJailBreakDetected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Security) GetIsJailBreakDetectedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsJailBreakDetected) {
		return nil, false
	}
	return o.IsJailBreakDetected, true
}

// HasIsJailBreakDetected returns a boolean if a field has been set.
func (o *Security) HasIsJailBreakDetected() bool {
	if o != nil && !IsNil(o.IsJailBreakDetected) {
		return true
	}

	return false
}

// SetIsJailBreakDetected gets a reference to the given bool and assigns it to the IsJailBreakDetected field.
func (o *Security) SetIsJailBreakDetected(v bool) {
	o.IsJailBreakDetected = &v
}

func (o Security) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Security) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsDataProtected) {
		toSerialize["isDataProtected"] = o.IsDataProtected
	}
	if !IsNil(o.IsBlockLevelEncryptionCapable) {
		toSerialize["isBlockLevelEncryptionCapable"] = o.IsBlockLevelEncryptionCapable
	}
	if !IsNil(o.IsFileLevelEncryptionCapable) {
		toSerialize["isFileLevelEncryptionCapable"] = o.IsFileLevelEncryptionCapable
	}
	if !IsNil(o.IsPasscodePresent) {
		toSerialize["isPasscodePresent"] = o.IsPasscodePresent
	}
	if !IsNil(o.IsPasscodeCompliant) {
		toSerialize["isPasscodeCompliant"] = o.IsPasscodeCompliant
	}
	if !IsNil(o.IsPasscodeCompliantWithProfile) {
		toSerialize["isPasscodeCompliantWithProfile"] = o.IsPasscodeCompliantWithProfile
	}
	if !IsNil(o.HardwareEncryption) {
		toSerialize["hardwareEncryption"] = o.HardwareEncryption
	}
	if !IsNil(o.IsActivationLockEnabled) {
		toSerialize["isActivationLockEnabled"] = o.IsActivationLockEnabled
	}
	if !IsNil(o.IsJailBreakDetected) {
		toSerialize["isJailBreakDetected"] = o.IsJailBreakDetected
	}
	return toSerialize, nil
}

type NullableSecurity struct {
	value *Security
	isSet bool
}

func (v NullableSecurity) Get() *Security {
	return v.value
}

func (v *NullableSecurity) Set(val *Security) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurity(val *Security) *NullableSecurity {
	return &NullableSecurity{value: val, isSet: true}
}

func (v NullableSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


