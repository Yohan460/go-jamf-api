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

// checks if the SelfServiceInstallSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelfServiceInstallSettings{}

// SelfServiceInstallSettings object representation of Self Service settings regarding installation 
type SelfServiceInstallSettings struct {
	// true if Self Service is insalled automatically, false if not 
	InstallAutomatically *bool `json:"installAutomatically,omitempty"`
	// path at which Self Service is installed 
	InstallLocation string `json:"installLocation"`
}

// NewSelfServiceInstallSettings instantiates a new SelfServiceInstallSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceInstallSettings(installLocation string) *SelfServiceInstallSettings {
	this := SelfServiceInstallSettings{}
	var installAutomatically bool = false
	this.InstallAutomatically = &installAutomatically
	this.InstallLocation = installLocation
	return &this
}

// NewSelfServiceInstallSettingsWithDefaults instantiates a new SelfServiceInstallSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceInstallSettingsWithDefaults() *SelfServiceInstallSettings {
	this := SelfServiceInstallSettings{}
	var installAutomatically bool = false
	this.InstallAutomatically = &installAutomatically
	return &this
}

// GetInstallAutomatically returns the InstallAutomatically field value if set, zero value otherwise.
func (o *SelfServiceInstallSettings) GetInstallAutomatically() bool {
	if o == nil || IsNil(o.InstallAutomatically) {
		var ret bool
		return ret
	}
	return *o.InstallAutomatically
}

// GetInstallAutomaticallyOk returns a tuple with the InstallAutomatically field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceInstallSettings) GetInstallAutomaticallyOk() (*bool, bool) {
	if o == nil || IsNil(o.InstallAutomatically) {
		return nil, false
	}
	return o.InstallAutomatically, true
}

// HasInstallAutomatically returns a boolean if a field has been set.
func (o *SelfServiceInstallSettings) HasInstallAutomatically() bool {
	if o != nil && !IsNil(o.InstallAutomatically) {
		return true
	}

	return false
}

// SetInstallAutomatically gets a reference to the given bool and assigns it to the InstallAutomatically field.
func (o *SelfServiceInstallSettings) SetInstallAutomatically(v bool) {
	o.InstallAutomatically = &v
}

// GetInstallLocation returns the InstallLocation field value
func (o *SelfServiceInstallSettings) GetInstallLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstallLocation
}

// GetInstallLocationOk returns a tuple with the InstallLocation field value
// and a boolean to check if the value has been set.
func (o *SelfServiceInstallSettings) GetInstallLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstallLocation, true
}

// SetInstallLocation sets field value
func (o *SelfServiceInstallSettings) SetInstallLocation(v string) {
	o.InstallLocation = v
}

func (o SelfServiceInstallSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelfServiceInstallSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstallAutomatically) {
		toSerialize["installAutomatically"] = o.InstallAutomatically
	}
	toSerialize["installLocation"] = o.InstallLocation
	return toSerialize, nil
}

type NullableSelfServiceInstallSettings struct {
	value *SelfServiceInstallSettings
	isSet bool
}

func (v NullableSelfServiceInstallSettings) Get() *SelfServiceInstallSettings {
	return v.value
}

func (v *NullableSelfServiceInstallSettings) Set(val *SelfServiceInstallSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceInstallSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceInstallSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceInstallSettings(val *SelfServiceInstallSettings) *NullableSelfServiceInstallSettings {
	return &NullableSelfServiceInstallSettings{value: val, isSet: true}
}

func (v NullableSelfServiceInstallSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceInstallSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


