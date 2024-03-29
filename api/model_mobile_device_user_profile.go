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

// checks if the MobileDeviceUserProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MobileDeviceUserProfile{}

// MobileDeviceUserProfile struct for MobileDeviceUserProfile
type MobileDeviceUserProfile struct {
	DisplayName *string `json:"displayName,omitempty"`
	Version *string `json:"version,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
	Removable *bool `json:"removable,omitempty"`
	LastInstalled *time.Time `json:"lastInstalled,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewMobileDeviceUserProfile instantiates a new MobileDeviceUserProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileDeviceUserProfile() *MobileDeviceUserProfile {
	this := MobileDeviceUserProfile{}
	return &this
}

// NewMobileDeviceUserProfileWithDefaults instantiates a new MobileDeviceUserProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileDeviceUserProfileWithDefaults() *MobileDeviceUserProfile {
	this := MobileDeviceUserProfile{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *MobileDeviceUserProfile) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceUserProfile) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MobileDeviceUserProfile) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MobileDeviceUserProfile) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *MobileDeviceUserProfile) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceUserProfile) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MobileDeviceUserProfile) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *MobileDeviceUserProfile) SetVersion(v string) {
	o.Version = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *MobileDeviceUserProfile) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceUserProfile) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *MobileDeviceUserProfile) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *MobileDeviceUserProfile) SetUuid(v string) {
	o.Uuid = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *MobileDeviceUserProfile) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceUserProfile) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *MobileDeviceUserProfile) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *MobileDeviceUserProfile) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetRemovable returns the Removable field value if set, zero value otherwise.
func (o *MobileDeviceUserProfile) GetRemovable() bool {
	if o == nil || IsNil(o.Removable) {
		var ret bool
		return ret
	}
	return *o.Removable
}

// GetRemovableOk returns a tuple with the Removable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceUserProfile) GetRemovableOk() (*bool, bool) {
	if o == nil || IsNil(o.Removable) {
		return nil, false
	}
	return o.Removable, true
}

// HasRemovable returns a boolean if a field has been set.
func (o *MobileDeviceUserProfile) HasRemovable() bool {
	if o != nil && !IsNil(o.Removable) {
		return true
	}

	return false
}

// SetRemovable gets a reference to the given bool and assigns it to the Removable field.
func (o *MobileDeviceUserProfile) SetRemovable(v bool) {
	o.Removable = &v
}

// GetLastInstalled returns the LastInstalled field value if set, zero value otherwise.
func (o *MobileDeviceUserProfile) GetLastInstalled() time.Time {
	if o == nil || IsNil(o.LastInstalled) {
		var ret time.Time
		return ret
	}
	return *o.LastInstalled
}

// GetLastInstalledOk returns a tuple with the LastInstalled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceUserProfile) GetLastInstalledOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastInstalled) {
		return nil, false
	}
	return o.LastInstalled, true
}

// HasLastInstalled returns a boolean if a field has been set.
func (o *MobileDeviceUserProfile) HasLastInstalled() bool {
	if o != nil && !IsNil(o.LastInstalled) {
		return true
	}

	return false
}

// SetLastInstalled gets a reference to the given time.Time and assigns it to the LastInstalled field.
func (o *MobileDeviceUserProfile) SetLastInstalled(v time.Time) {
	o.LastInstalled = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *MobileDeviceUserProfile) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceUserProfile) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *MobileDeviceUserProfile) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *MobileDeviceUserProfile) SetUsername(v string) {
	o.Username = &v
}

func (o MobileDeviceUserProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MobileDeviceUserProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Removable) {
		toSerialize["removable"] = o.Removable
	}
	if !IsNil(o.LastInstalled) {
		toSerialize["lastInstalled"] = o.LastInstalled
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableMobileDeviceUserProfile struct {
	value *MobileDeviceUserProfile
	isSet bool
}

func (v NullableMobileDeviceUserProfile) Get() *MobileDeviceUserProfile {
	return v.value
}

func (v *NullableMobileDeviceUserProfile) Set(val *MobileDeviceUserProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileDeviceUserProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileDeviceUserProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileDeviceUserProfile(val *MobileDeviceUserProfile) *NullableMobileDeviceUserProfile {
	return &NullableMobileDeviceUserProfile{value: val, isSet: true}
}

func (v NullableMobileDeviceUserProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileDeviceUserProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


