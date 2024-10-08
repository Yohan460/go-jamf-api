/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MobileDeviceInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MobileDeviceInventory{}

// MobileDeviceInventory struct for MobileDeviceInventory
type MobileDeviceInventory struct {
	MobileDeviceId *string `json:"mobileDeviceId,omitempty"`
	// Based on the value of this type either ios, appleTv, watch or visionOS objects will be populated.
	DeviceType string `json:"deviceType"`
	Hardware *MobileDeviceHardware `json:"hardware,omitempty"`
	UserAndLocation *MobileDeviceUserAndLocation `json:"userAndLocation,omitempty"`
	Applications []MobileDeviceApplicationInventoryDetail `json:"applications,omitempty"`
	Certificates []MobileDeviceCertificate `json:"certificates,omitempty"`
	Profiles []MobileDeviceProfile `json:"profiles,omitempty"`
	ExtensionAttributes []MobileDeviceExtensionAttribute `json:"extensionAttributes,omitempty"`
}

type _MobileDeviceInventory MobileDeviceInventory

// NewMobileDeviceInventory instantiates a new MobileDeviceInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileDeviceInventory(deviceType string) *MobileDeviceInventory {
	this := MobileDeviceInventory{}
	this.DeviceType = deviceType
	return &this
}

// NewMobileDeviceInventoryWithDefaults instantiates a new MobileDeviceInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileDeviceInventoryWithDefaults() *MobileDeviceInventory {
	this := MobileDeviceInventory{}
	return &this
}

// GetMobileDeviceId returns the MobileDeviceId field value if set, zero value otherwise.
func (o *MobileDeviceInventory) GetMobileDeviceId() string {
	if o == nil || IsNil(o.MobileDeviceId) {
		var ret string
		return ret
	}
	return *o.MobileDeviceId
}

// GetMobileDeviceIdOk returns a tuple with the MobileDeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceInventory) GetMobileDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.MobileDeviceId) {
		return nil, false
	}
	return o.MobileDeviceId, true
}

// HasMobileDeviceId returns a boolean if a field has been set.
func (o *MobileDeviceInventory) HasMobileDeviceId() bool {
	if o != nil && !IsNil(o.MobileDeviceId) {
		return true
	}

	return false
}

// SetMobileDeviceId gets a reference to the given string and assigns it to the MobileDeviceId field.
func (o *MobileDeviceInventory) SetMobileDeviceId(v string) {
	o.MobileDeviceId = &v
}

// GetDeviceType returns the DeviceType field value
func (o *MobileDeviceInventory) GetDeviceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *MobileDeviceInventory) GetDeviceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *MobileDeviceInventory) SetDeviceType(v string) {
	o.DeviceType = v
}

// GetHardware returns the Hardware field value if set, zero value otherwise.
func (o *MobileDeviceInventory) GetHardware() MobileDeviceHardware {
	if o == nil || IsNil(o.Hardware) {
		var ret MobileDeviceHardware
		return ret
	}
	return *o.Hardware
}

// GetHardwareOk returns a tuple with the Hardware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceInventory) GetHardwareOk() (*MobileDeviceHardware, bool) {
	if o == nil || IsNil(o.Hardware) {
		return nil, false
	}
	return o.Hardware, true
}

// HasHardware returns a boolean if a field has been set.
func (o *MobileDeviceInventory) HasHardware() bool {
	if o != nil && !IsNil(o.Hardware) {
		return true
	}

	return false
}

// SetHardware gets a reference to the given MobileDeviceHardware and assigns it to the Hardware field.
func (o *MobileDeviceInventory) SetHardware(v MobileDeviceHardware) {
	o.Hardware = &v
}

// GetUserAndLocation returns the UserAndLocation field value if set, zero value otherwise.
func (o *MobileDeviceInventory) GetUserAndLocation() MobileDeviceUserAndLocation {
	if o == nil || IsNil(o.UserAndLocation) {
		var ret MobileDeviceUserAndLocation
		return ret
	}
	return *o.UserAndLocation
}

// GetUserAndLocationOk returns a tuple with the UserAndLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceInventory) GetUserAndLocationOk() (*MobileDeviceUserAndLocation, bool) {
	if o == nil || IsNil(o.UserAndLocation) {
		return nil, false
	}
	return o.UserAndLocation, true
}

// HasUserAndLocation returns a boolean if a field has been set.
func (o *MobileDeviceInventory) HasUserAndLocation() bool {
	if o != nil && !IsNil(o.UserAndLocation) {
		return true
	}

	return false
}

// SetUserAndLocation gets a reference to the given MobileDeviceUserAndLocation and assigns it to the UserAndLocation field.
func (o *MobileDeviceInventory) SetUserAndLocation(v MobileDeviceUserAndLocation) {
	o.UserAndLocation = &v
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *MobileDeviceInventory) GetApplications() []MobileDeviceApplicationInventoryDetail {
	if o == nil || IsNil(o.Applications) {
		var ret []MobileDeviceApplicationInventoryDetail
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceInventory) GetApplicationsOk() ([]MobileDeviceApplicationInventoryDetail, bool) {
	if o == nil || IsNil(o.Applications) {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *MobileDeviceInventory) HasApplications() bool {
	if o != nil && !IsNil(o.Applications) {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []MobileDeviceApplicationInventoryDetail and assigns it to the Applications field.
func (o *MobileDeviceInventory) SetApplications(v []MobileDeviceApplicationInventoryDetail) {
	o.Applications = v
}

// GetCertificates returns the Certificates field value if set, zero value otherwise.
func (o *MobileDeviceInventory) GetCertificates() []MobileDeviceCertificate {
	if o == nil || IsNil(o.Certificates) {
		var ret []MobileDeviceCertificate
		return ret
	}
	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceInventory) GetCertificatesOk() ([]MobileDeviceCertificate, bool) {
	if o == nil || IsNil(o.Certificates) {
		return nil, false
	}
	return o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *MobileDeviceInventory) HasCertificates() bool {
	if o != nil && !IsNil(o.Certificates) {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given []MobileDeviceCertificate and assigns it to the Certificates field.
func (o *MobileDeviceInventory) SetCertificates(v []MobileDeviceCertificate) {
	o.Certificates = v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *MobileDeviceInventory) GetProfiles() []MobileDeviceProfile {
	if o == nil || IsNil(o.Profiles) {
		var ret []MobileDeviceProfile
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceInventory) GetProfilesOk() ([]MobileDeviceProfile, bool) {
	if o == nil || IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *MobileDeviceInventory) HasProfiles() bool {
	if o != nil && !IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []MobileDeviceProfile and assigns it to the Profiles field.
func (o *MobileDeviceInventory) SetProfiles(v []MobileDeviceProfile) {
	o.Profiles = v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *MobileDeviceInventory) GetExtensionAttributes() []MobileDeviceExtensionAttribute {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret []MobileDeviceExtensionAttribute
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceInventory) GetExtensionAttributesOk() ([]MobileDeviceExtensionAttribute, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return nil, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *MobileDeviceInventory) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given []MobileDeviceExtensionAttribute and assigns it to the ExtensionAttributes field.
func (o *MobileDeviceInventory) SetExtensionAttributes(v []MobileDeviceExtensionAttribute) {
	o.ExtensionAttributes = v
}

func (o MobileDeviceInventory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MobileDeviceInventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MobileDeviceId) {
		toSerialize["mobileDeviceId"] = o.MobileDeviceId
	}
	toSerialize["deviceType"] = o.DeviceType
	if !IsNil(o.Hardware) {
		toSerialize["hardware"] = o.Hardware
	}
	if !IsNil(o.UserAndLocation) {
		toSerialize["userAndLocation"] = o.UserAndLocation
	}
	if !IsNil(o.Applications) {
		toSerialize["applications"] = o.Applications
	}
	if !IsNil(o.Certificates) {
		toSerialize["certificates"] = o.Certificates
	}
	if !IsNil(o.Profiles) {
		toSerialize["profiles"] = o.Profiles
	}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extensionAttributes"] = o.ExtensionAttributes
	}
	return toSerialize, nil
}

func (o *MobileDeviceInventory) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deviceType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMobileDeviceInventory := _MobileDeviceInventory{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMobileDeviceInventory)

	if err != nil {
		return err
	}

	*o = MobileDeviceInventory(varMobileDeviceInventory)

	return err
}

type NullableMobileDeviceInventory struct {
	value *MobileDeviceInventory
	isSet bool
}

func (v NullableMobileDeviceInventory) Get() *MobileDeviceInventory {
	return v.value
}

func (v *NullableMobileDeviceInventory) Set(val *MobileDeviceInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileDeviceInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileDeviceInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileDeviceInventory(val *MobileDeviceInventory) *NullableMobileDeviceInventory {
	return &NullableMobileDeviceInventory{value: val, isSet: true}
}

func (v NullableMobileDeviceInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileDeviceInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


