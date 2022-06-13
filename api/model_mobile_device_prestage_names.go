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

// MobileDevicePrestageNames struct for MobileDevicePrestageNames
type MobileDevicePrestageNames struct {
	AssignNamesUsing *string `json:"assignNamesUsing,omitempty"`
	PrestageDeviceNames []MobileDevicePrestageName `json:"prestageDeviceNames,omitempty"`
	DeviceNamePrefix *string `json:"deviceNamePrefix,omitempty"`
	DeviceNameSuffix *string `json:"deviceNameSuffix,omitempty"`
	SingleDeviceName *string `json:"singleDeviceName,omitempty"`
	IsManageNames *bool `json:"isManageNames,omitempty"`
	IsDeviceNamingConfigured *bool `json:"isDeviceNamingConfigured,omitempty"`
}

// NewMobileDevicePrestageNames instantiates a new MobileDevicePrestageNames object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileDevicePrestageNames() *MobileDevicePrestageNames {
	this := MobileDevicePrestageNames{}
	return &this
}

// NewMobileDevicePrestageNamesWithDefaults instantiates a new MobileDevicePrestageNames object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileDevicePrestageNamesWithDefaults() *MobileDevicePrestageNames {
	this := MobileDevicePrestageNames{}
	return &this
}

// GetAssignNamesUsing returns the AssignNamesUsing field value if set, zero value otherwise.
func (o *MobileDevicePrestageNames) GetAssignNamesUsing() string {
	if o == nil || o.AssignNamesUsing == nil {
		var ret string
		return ret
	}
	return *o.AssignNamesUsing
}

// GetAssignNamesUsingOk returns a tuple with the AssignNamesUsing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageNames) GetAssignNamesUsingOk() (*string, bool) {
	if o == nil || o.AssignNamesUsing == nil {
		return nil, false
	}
	return o.AssignNamesUsing, true
}

// HasAssignNamesUsing returns a boolean if a field has been set.
func (o *MobileDevicePrestageNames) HasAssignNamesUsing() bool {
	if o != nil && o.AssignNamesUsing != nil {
		return true
	}

	return false
}

// SetAssignNamesUsing gets a reference to the given string and assigns it to the AssignNamesUsing field.
func (o *MobileDevicePrestageNames) SetAssignNamesUsing(v string) {
	o.AssignNamesUsing = &v
}

// GetPrestageDeviceNames returns the PrestageDeviceNames field value if set, zero value otherwise.
func (o *MobileDevicePrestageNames) GetPrestageDeviceNames() []MobileDevicePrestageName {
	if o == nil || o.PrestageDeviceNames == nil {
		var ret []MobileDevicePrestageName
		return ret
	}
	return o.PrestageDeviceNames
}

// GetPrestageDeviceNamesOk returns a tuple with the PrestageDeviceNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageNames) GetPrestageDeviceNamesOk() ([]MobileDevicePrestageName, bool) {
	if o == nil || o.PrestageDeviceNames == nil {
		return nil, false
	}
	return o.PrestageDeviceNames, true
}

// HasPrestageDeviceNames returns a boolean if a field has been set.
func (o *MobileDevicePrestageNames) HasPrestageDeviceNames() bool {
	if o != nil && o.PrestageDeviceNames != nil {
		return true
	}

	return false
}

// SetPrestageDeviceNames gets a reference to the given []MobileDevicePrestageName and assigns it to the PrestageDeviceNames field.
func (o *MobileDevicePrestageNames) SetPrestageDeviceNames(v []MobileDevicePrestageName) {
	o.PrestageDeviceNames = v
}

// GetDeviceNamePrefix returns the DeviceNamePrefix field value if set, zero value otherwise.
func (o *MobileDevicePrestageNames) GetDeviceNamePrefix() string {
	if o == nil || o.DeviceNamePrefix == nil {
		var ret string
		return ret
	}
	return *o.DeviceNamePrefix
}

// GetDeviceNamePrefixOk returns a tuple with the DeviceNamePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageNames) GetDeviceNamePrefixOk() (*string, bool) {
	if o == nil || o.DeviceNamePrefix == nil {
		return nil, false
	}
	return o.DeviceNamePrefix, true
}

// HasDeviceNamePrefix returns a boolean if a field has been set.
func (o *MobileDevicePrestageNames) HasDeviceNamePrefix() bool {
	if o != nil && o.DeviceNamePrefix != nil {
		return true
	}

	return false
}

// SetDeviceNamePrefix gets a reference to the given string and assigns it to the DeviceNamePrefix field.
func (o *MobileDevicePrestageNames) SetDeviceNamePrefix(v string) {
	o.DeviceNamePrefix = &v
}

// GetDeviceNameSuffix returns the DeviceNameSuffix field value if set, zero value otherwise.
func (o *MobileDevicePrestageNames) GetDeviceNameSuffix() string {
	if o == nil || o.DeviceNameSuffix == nil {
		var ret string
		return ret
	}
	return *o.DeviceNameSuffix
}

// GetDeviceNameSuffixOk returns a tuple with the DeviceNameSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageNames) GetDeviceNameSuffixOk() (*string, bool) {
	if o == nil || o.DeviceNameSuffix == nil {
		return nil, false
	}
	return o.DeviceNameSuffix, true
}

// HasDeviceNameSuffix returns a boolean if a field has been set.
func (o *MobileDevicePrestageNames) HasDeviceNameSuffix() bool {
	if o != nil && o.DeviceNameSuffix != nil {
		return true
	}

	return false
}

// SetDeviceNameSuffix gets a reference to the given string and assigns it to the DeviceNameSuffix field.
func (o *MobileDevicePrestageNames) SetDeviceNameSuffix(v string) {
	o.DeviceNameSuffix = &v
}

// GetSingleDeviceName returns the SingleDeviceName field value if set, zero value otherwise.
func (o *MobileDevicePrestageNames) GetSingleDeviceName() string {
	if o == nil || o.SingleDeviceName == nil {
		var ret string
		return ret
	}
	return *o.SingleDeviceName
}

// GetSingleDeviceNameOk returns a tuple with the SingleDeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageNames) GetSingleDeviceNameOk() (*string, bool) {
	if o == nil || o.SingleDeviceName == nil {
		return nil, false
	}
	return o.SingleDeviceName, true
}

// HasSingleDeviceName returns a boolean if a field has been set.
func (o *MobileDevicePrestageNames) HasSingleDeviceName() bool {
	if o != nil && o.SingleDeviceName != nil {
		return true
	}

	return false
}

// SetSingleDeviceName gets a reference to the given string and assigns it to the SingleDeviceName field.
func (o *MobileDevicePrestageNames) SetSingleDeviceName(v string) {
	o.SingleDeviceName = &v
}

// GetIsManageNames returns the IsManageNames field value if set, zero value otherwise.
func (o *MobileDevicePrestageNames) GetIsManageNames() bool {
	if o == nil || o.IsManageNames == nil {
		var ret bool
		return ret
	}
	return *o.IsManageNames
}

// GetIsManageNamesOk returns a tuple with the IsManageNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageNames) GetIsManageNamesOk() (*bool, bool) {
	if o == nil || o.IsManageNames == nil {
		return nil, false
	}
	return o.IsManageNames, true
}

// HasIsManageNames returns a boolean if a field has been set.
func (o *MobileDevicePrestageNames) HasIsManageNames() bool {
	if o != nil && o.IsManageNames != nil {
		return true
	}

	return false
}

// SetIsManageNames gets a reference to the given bool and assigns it to the IsManageNames field.
func (o *MobileDevicePrestageNames) SetIsManageNames(v bool) {
	o.IsManageNames = &v
}

// GetIsDeviceNamingConfigured returns the IsDeviceNamingConfigured field value if set, zero value otherwise.
func (o *MobileDevicePrestageNames) GetIsDeviceNamingConfigured() bool {
	if o == nil || o.IsDeviceNamingConfigured == nil {
		var ret bool
		return ret
	}
	return *o.IsDeviceNamingConfigured
}

// GetIsDeviceNamingConfiguredOk returns a tuple with the IsDeviceNamingConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestageNames) GetIsDeviceNamingConfiguredOk() (*bool, bool) {
	if o == nil || o.IsDeviceNamingConfigured == nil {
		return nil, false
	}
	return o.IsDeviceNamingConfigured, true
}

// HasIsDeviceNamingConfigured returns a boolean if a field has been set.
func (o *MobileDevicePrestageNames) HasIsDeviceNamingConfigured() bool {
	if o != nil && o.IsDeviceNamingConfigured != nil {
		return true
	}

	return false
}

// SetIsDeviceNamingConfigured gets a reference to the given bool and assigns it to the IsDeviceNamingConfigured field.
func (o *MobileDevicePrestageNames) SetIsDeviceNamingConfigured(v bool) {
	o.IsDeviceNamingConfigured = &v
}

func (o MobileDevicePrestageNames) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssignNamesUsing != nil {
		toSerialize["assignNamesUsing"] = o.AssignNamesUsing
	}
	if o.PrestageDeviceNames != nil {
		toSerialize["prestageDeviceNames"] = o.PrestageDeviceNames
	}
	if o.DeviceNamePrefix != nil {
		toSerialize["deviceNamePrefix"] = o.DeviceNamePrefix
	}
	if o.DeviceNameSuffix != nil {
		toSerialize["deviceNameSuffix"] = o.DeviceNameSuffix
	}
	if o.SingleDeviceName != nil {
		toSerialize["singleDeviceName"] = o.SingleDeviceName
	}
	if o.IsManageNames != nil {
		toSerialize["isManageNames"] = o.IsManageNames
	}
	if o.IsDeviceNamingConfigured != nil {
		toSerialize["isDeviceNamingConfigured"] = o.IsDeviceNamingConfigured
	}
	return json.Marshal(toSerialize)
}

type NullableMobileDevicePrestageNames struct {
	value *MobileDevicePrestageNames
	isSet bool
}

func (v NullableMobileDevicePrestageNames) Get() *MobileDevicePrestageNames {
	return v.value
}

func (v *NullableMobileDevicePrestageNames) Set(val *MobileDevicePrestageNames) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileDevicePrestageNames) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileDevicePrestageNames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileDevicePrestageNames(val *MobileDevicePrestageNames) *NullableMobileDevicePrestageNames {
	return &NullableMobileDevicePrestageNames{value: val, isSet: true}
}

func (v NullableMobileDevicePrestageNames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileDevicePrestageNames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


