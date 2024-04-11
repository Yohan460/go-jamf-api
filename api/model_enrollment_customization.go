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

// checks if the EnrollmentCustomization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnrollmentCustomization{}

// EnrollmentCustomization struct for EnrollmentCustomization
type EnrollmentCustomization struct {
	SiteId int64 `json:"siteId"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	EnrollmentCustomizationBrandingSettings EnrollmentCustomizationBrandingSettings `json:"enrollmentCustomizationBrandingSettings"`
}

type _EnrollmentCustomization EnrollmentCustomization

// NewEnrollmentCustomization instantiates a new EnrollmentCustomization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrollmentCustomization(siteId int64, displayName string, description string, enrollmentCustomizationBrandingSettings EnrollmentCustomizationBrandingSettings) *EnrollmentCustomization {
	this := EnrollmentCustomization{}
	this.SiteId = siteId
	this.DisplayName = displayName
	this.Description = description
	this.EnrollmentCustomizationBrandingSettings = enrollmentCustomizationBrandingSettings
	return &this
}

// NewEnrollmentCustomizationWithDefaults instantiates a new EnrollmentCustomization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrollmentCustomizationWithDefaults() *EnrollmentCustomization {
	this := EnrollmentCustomization{}
	return &this
}

// GetSiteId returns the SiteId field value
func (o *EnrollmentCustomization) GetSiteId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomization) GetSiteIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *EnrollmentCustomization) SetSiteId(v int64) {
	o.SiteId = v
}

// GetDisplayName returns the DisplayName field value
func (o *EnrollmentCustomization) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomization) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *EnrollmentCustomization) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDescription returns the Description field value
func (o *EnrollmentCustomization) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomization) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *EnrollmentCustomization) SetDescription(v string) {
	o.Description = v
}

// GetEnrollmentCustomizationBrandingSettings returns the EnrollmentCustomizationBrandingSettings field value
func (o *EnrollmentCustomization) GetEnrollmentCustomizationBrandingSettings() EnrollmentCustomizationBrandingSettings {
	if o == nil {
		var ret EnrollmentCustomizationBrandingSettings
		return ret
	}

	return o.EnrollmentCustomizationBrandingSettings
}

// GetEnrollmentCustomizationBrandingSettingsOk returns a tuple with the EnrollmentCustomizationBrandingSettings field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomization) GetEnrollmentCustomizationBrandingSettingsOk() (*EnrollmentCustomizationBrandingSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnrollmentCustomizationBrandingSettings, true
}

// SetEnrollmentCustomizationBrandingSettings sets field value
func (o *EnrollmentCustomization) SetEnrollmentCustomizationBrandingSettings(v EnrollmentCustomizationBrandingSettings) {
	o.EnrollmentCustomizationBrandingSettings = v
}

func (o EnrollmentCustomization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnrollmentCustomization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["siteId"] = o.SiteId
	toSerialize["displayName"] = o.DisplayName
	toSerialize["description"] = o.Description
	toSerialize["enrollmentCustomizationBrandingSettings"] = o.EnrollmentCustomizationBrandingSettings
	return toSerialize, nil
}

func (o *EnrollmentCustomization) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"siteId",
		"displayName",
		"description",
		"enrollmentCustomizationBrandingSettings",
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

	varEnrollmentCustomization := _EnrollmentCustomization{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEnrollmentCustomization)

	if err != nil {
		return err
	}

	*o = EnrollmentCustomization(varEnrollmentCustomization)

	return err
}

type NullableEnrollmentCustomization struct {
	value *EnrollmentCustomization
	isSet bool
}

func (v NullableEnrollmentCustomization) Get() *EnrollmentCustomization {
	return v.value
}

func (v *NullableEnrollmentCustomization) Set(val *EnrollmentCustomization) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollmentCustomization) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollmentCustomization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollmentCustomization(val *EnrollmentCustomization) *NullableEnrollmentCustomization {
	return &NullableEnrollmentCustomization{value: val, isSet: true}
}

func (v NullableEnrollmentCustomization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrollmentCustomization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


