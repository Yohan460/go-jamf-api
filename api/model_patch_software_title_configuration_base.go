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

// checks if the PatchSoftwareTitleConfigurationBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchSoftwareTitleConfigurationBase{}

// PatchSoftwareTitleConfigurationBase struct for PatchSoftwareTitleConfigurationBase
type PatchSoftwareTitleConfigurationBase struct {
	DisplayName string `json:"displayName"`
	CategoryId *string `json:"categoryId,omitempty"`
	SiteId *string `json:"siteId,omitempty"`
	UiNotifications *bool `json:"uiNotifications,omitempty"`
	EmailNotifications *bool `json:"emailNotifications,omitempty"`
	SoftwareTitleId string `json:"softwareTitleId"`
	JamfOfficial *bool `json:"jamfOfficial,omitempty"`
	ExtensionAttributes []PatchSoftwareTitleConfigurationExtensionAttributes `json:"extensionAttributes,omitempty"`
	SoftwareTitleName *string `json:"softwareTitleName,omitempty"`
	SoftwareTitleNameId *string `json:"softwareTitleNameId,omitempty"`
	SoftwareTitlePublisher *string `json:"softwareTitlePublisher,omitempty"`
	PatchSourceName *string `json:"patchSourceName,omitempty"`
	PatchSourceEnabled *bool `json:"patchSourceEnabled,omitempty"`
}

type _PatchSoftwareTitleConfigurationBase PatchSoftwareTitleConfigurationBase

// NewPatchSoftwareTitleConfigurationBase instantiates a new PatchSoftwareTitleConfigurationBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchSoftwareTitleConfigurationBase(displayName string, softwareTitleId string) *PatchSoftwareTitleConfigurationBase {
	this := PatchSoftwareTitleConfigurationBase{}
	this.DisplayName = displayName
	var categoryId string = "-1"
	this.CategoryId = &categoryId
	var siteId string = "-1"
	this.SiteId = &siteId
	var uiNotifications bool = false
	this.UiNotifications = &uiNotifications
	var emailNotifications bool = false
	this.EmailNotifications = &emailNotifications
	this.SoftwareTitleId = softwareTitleId
	return &this
}

// NewPatchSoftwareTitleConfigurationBaseWithDefaults instantiates a new PatchSoftwareTitleConfigurationBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchSoftwareTitleConfigurationBaseWithDefaults() *PatchSoftwareTitleConfigurationBase {
	this := PatchSoftwareTitleConfigurationBase{}
	var categoryId string = "-1"
	this.CategoryId = &categoryId
	var siteId string = "-1"
	this.SiteId = &siteId
	var uiNotifications bool = false
	this.UiNotifications = &uiNotifications
	var emailNotifications bool = false
	this.EmailNotifications = &emailNotifications
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *PatchSoftwareTitleConfigurationBase) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfigurationBase) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *PatchSoftwareTitleConfigurationBase) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfigurationBase) GetCategoryId() string {
	if o == nil || IsNil(o.CategoryId) {
		var ret string
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfigurationBase) GetCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfigurationBase) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *PatchSoftwareTitleConfigurationBase) SetCategoryId(v string) {
	o.CategoryId = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfigurationBase) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfigurationBase) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfigurationBase) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *PatchSoftwareTitleConfigurationBase) SetSiteId(v string) {
	o.SiteId = &v
}

// GetUiNotifications returns the UiNotifications field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfigurationBase) GetUiNotifications() bool {
	if o == nil || IsNil(o.UiNotifications) {
		var ret bool
		return ret
	}
	return *o.UiNotifications
}

// GetUiNotificationsOk returns a tuple with the UiNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfigurationBase) GetUiNotificationsOk() (*bool, bool) {
	if o == nil || IsNil(o.UiNotifications) {
		return nil, false
	}
	return o.UiNotifications, true
}

// HasUiNotifications returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfigurationBase) HasUiNotifications() bool {
	if o != nil && !IsNil(o.UiNotifications) {
		return true
	}

	return false
}

// SetUiNotifications gets a reference to the given bool and assigns it to the UiNotifications field.
func (o *PatchSoftwareTitleConfigurationBase) SetUiNotifications(v bool) {
	o.UiNotifications = &v
}

// GetEmailNotifications returns the EmailNotifications field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfigurationBase) GetEmailNotifications() bool {
	if o == nil || IsNil(o.EmailNotifications) {
		var ret bool
		return ret
	}
	return *o.EmailNotifications
}

// GetEmailNotificationsOk returns a tuple with the EmailNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfigurationBase) GetEmailNotificationsOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailNotifications) {
		return nil, false
	}
	return o.EmailNotifications, true
}

// HasEmailNotifications returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfigurationBase) HasEmailNotifications() bool {
	if o != nil && !IsNil(o.EmailNotifications) {
		return true
	}

	return false
}

// SetEmailNotifications gets a reference to the given bool and assigns it to the EmailNotifications field.
func (o *PatchSoftwareTitleConfigurationBase) SetEmailNotifications(v bool) {
	o.EmailNotifications = &v
}

// GetSoftwareTitleId returns the SoftwareTitleId field value
func (o *PatchSoftwareTitleConfigurationBase) GetSoftwareTitleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SoftwareTitleId
}

// GetSoftwareTitleIdOk returns a tuple with the SoftwareTitleId field value
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfigurationBase) GetSoftwareTitleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SoftwareTitleId, true
}

// SetSoftwareTitleId sets field value
func (o *PatchSoftwareTitleConfigurationBase) SetSoftwareTitleId(v string) {
	o.SoftwareTitleId = v
}

// GetJamfOfficial returns the JamfOfficial field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfigurationBase) GetJamfOfficial() bool {
	if o == nil || IsNil(o.JamfOfficial) {
		var ret bool
		return ret
	}
	return *o.JamfOfficial
}

// GetJamfOfficialOk returns a tuple with the JamfOfficial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfigurationBase) GetJamfOfficialOk() (*bool, bool) {
	if o == nil || IsNil(o.JamfOfficial) {
		return nil, false
	}
	return o.JamfOfficial, true
}

// HasJamfOfficial returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfigurationBase) HasJamfOfficial() bool {
	if o != nil && !IsNil(o.JamfOfficial) {
		return true
	}

	return false
}

// SetJamfOfficial gets a reference to the given bool and assigns it to the JamfOfficial field.
func (o *PatchSoftwareTitleConfigurationBase) SetJamfOfficial(v bool) {
	o.JamfOfficial = &v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfigurationBase) GetExtensionAttributes() []PatchSoftwareTitleConfigurationExtensionAttributes {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret []PatchSoftwareTitleConfigurationExtensionAttributes
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfigurationBase) GetExtensionAttributesOk() ([]PatchSoftwareTitleConfigurationExtensionAttributes, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return nil, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfigurationBase) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given []PatchSoftwareTitleConfigurationExtensionAttributes and assigns it to the ExtensionAttributes field.
func (o *PatchSoftwareTitleConfigurationBase) SetExtensionAttributes(v []PatchSoftwareTitleConfigurationExtensionAttributes) {
	o.ExtensionAttributes = v
}

// GetSoftwareTitleName returns the SoftwareTitleName field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfigurationBase) GetSoftwareTitleName() string {
	if o == nil || IsNil(o.SoftwareTitleName) {
		var ret string
		return ret
	}
	return *o.SoftwareTitleName
}

// GetSoftwareTitleNameOk returns a tuple with the SoftwareTitleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfigurationBase) GetSoftwareTitleNameOk() (*string, bool) {
	if o == nil || IsNil(o.SoftwareTitleName) {
		return nil, false
	}
	return o.SoftwareTitleName, true
}

// HasSoftwareTitleName returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfigurationBase) HasSoftwareTitleName() bool {
	if o != nil && !IsNil(o.SoftwareTitleName) {
		return true
	}

	return false
}

// SetSoftwareTitleName gets a reference to the given string and assigns it to the SoftwareTitleName field.
func (o *PatchSoftwareTitleConfigurationBase) SetSoftwareTitleName(v string) {
	o.SoftwareTitleName = &v
}

// GetSoftwareTitleNameId returns the SoftwareTitleNameId field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfigurationBase) GetSoftwareTitleNameId() string {
	if o == nil || IsNil(o.SoftwareTitleNameId) {
		var ret string
		return ret
	}
	return *o.SoftwareTitleNameId
}

// GetSoftwareTitleNameIdOk returns a tuple with the SoftwareTitleNameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfigurationBase) GetSoftwareTitleNameIdOk() (*string, bool) {
	if o == nil || IsNil(o.SoftwareTitleNameId) {
		return nil, false
	}
	return o.SoftwareTitleNameId, true
}

// HasSoftwareTitleNameId returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfigurationBase) HasSoftwareTitleNameId() bool {
	if o != nil && !IsNil(o.SoftwareTitleNameId) {
		return true
	}

	return false
}

// SetSoftwareTitleNameId gets a reference to the given string and assigns it to the SoftwareTitleNameId field.
func (o *PatchSoftwareTitleConfigurationBase) SetSoftwareTitleNameId(v string) {
	o.SoftwareTitleNameId = &v
}

// GetSoftwareTitlePublisher returns the SoftwareTitlePublisher field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfigurationBase) GetSoftwareTitlePublisher() string {
	if o == nil || IsNil(o.SoftwareTitlePublisher) {
		var ret string
		return ret
	}
	return *o.SoftwareTitlePublisher
}

// GetSoftwareTitlePublisherOk returns a tuple with the SoftwareTitlePublisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfigurationBase) GetSoftwareTitlePublisherOk() (*string, bool) {
	if o == nil || IsNil(o.SoftwareTitlePublisher) {
		return nil, false
	}
	return o.SoftwareTitlePublisher, true
}

// HasSoftwareTitlePublisher returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfigurationBase) HasSoftwareTitlePublisher() bool {
	if o != nil && !IsNil(o.SoftwareTitlePublisher) {
		return true
	}

	return false
}

// SetSoftwareTitlePublisher gets a reference to the given string and assigns it to the SoftwareTitlePublisher field.
func (o *PatchSoftwareTitleConfigurationBase) SetSoftwareTitlePublisher(v string) {
	o.SoftwareTitlePublisher = &v
}

// GetPatchSourceName returns the PatchSourceName field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfigurationBase) GetPatchSourceName() string {
	if o == nil || IsNil(o.PatchSourceName) {
		var ret string
		return ret
	}
	return *o.PatchSourceName
}

// GetPatchSourceNameOk returns a tuple with the PatchSourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfigurationBase) GetPatchSourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.PatchSourceName) {
		return nil, false
	}
	return o.PatchSourceName, true
}

// HasPatchSourceName returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfigurationBase) HasPatchSourceName() bool {
	if o != nil && !IsNil(o.PatchSourceName) {
		return true
	}

	return false
}

// SetPatchSourceName gets a reference to the given string and assigns it to the PatchSourceName field.
func (o *PatchSoftwareTitleConfigurationBase) SetPatchSourceName(v string) {
	o.PatchSourceName = &v
}

// GetPatchSourceEnabled returns the PatchSourceEnabled field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfigurationBase) GetPatchSourceEnabled() bool {
	if o == nil || IsNil(o.PatchSourceEnabled) {
		var ret bool
		return ret
	}
	return *o.PatchSourceEnabled
}

// GetPatchSourceEnabledOk returns a tuple with the PatchSourceEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfigurationBase) GetPatchSourceEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PatchSourceEnabled) {
		return nil, false
	}
	return o.PatchSourceEnabled, true
}

// HasPatchSourceEnabled returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfigurationBase) HasPatchSourceEnabled() bool {
	if o != nil && !IsNil(o.PatchSourceEnabled) {
		return true
	}

	return false
}

// SetPatchSourceEnabled gets a reference to the given bool and assigns it to the PatchSourceEnabled field.
func (o *PatchSoftwareTitleConfigurationBase) SetPatchSourceEnabled(v bool) {
	o.PatchSourceEnabled = &v
}

func (o PatchSoftwareTitleConfigurationBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchSoftwareTitleConfigurationBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["displayName"] = o.DisplayName
	if !IsNil(o.CategoryId) {
		toSerialize["categoryId"] = o.CategoryId
	}
	if !IsNil(o.SiteId) {
		toSerialize["siteId"] = o.SiteId
	}
	if !IsNil(o.UiNotifications) {
		toSerialize["uiNotifications"] = o.UiNotifications
	}
	if !IsNil(o.EmailNotifications) {
		toSerialize["emailNotifications"] = o.EmailNotifications
	}
	toSerialize["softwareTitleId"] = o.SoftwareTitleId
	if !IsNil(o.JamfOfficial) {
		toSerialize["jamfOfficial"] = o.JamfOfficial
	}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extensionAttributes"] = o.ExtensionAttributes
	}
	if !IsNil(o.SoftwareTitleName) {
		toSerialize["softwareTitleName"] = o.SoftwareTitleName
	}
	if !IsNil(o.SoftwareTitleNameId) {
		toSerialize["softwareTitleNameId"] = o.SoftwareTitleNameId
	}
	if !IsNil(o.SoftwareTitlePublisher) {
		toSerialize["softwareTitlePublisher"] = o.SoftwareTitlePublisher
	}
	if !IsNil(o.PatchSourceName) {
		toSerialize["patchSourceName"] = o.PatchSourceName
	}
	if !IsNil(o.PatchSourceEnabled) {
		toSerialize["patchSourceEnabled"] = o.PatchSourceEnabled
	}
	return toSerialize, nil
}

func (o *PatchSoftwareTitleConfigurationBase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"displayName",
		"softwareTitleId",
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

	varPatchSoftwareTitleConfigurationBase := _PatchSoftwareTitleConfigurationBase{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPatchSoftwareTitleConfigurationBase)

	if err != nil {
		return err
	}

	*o = PatchSoftwareTitleConfigurationBase(varPatchSoftwareTitleConfigurationBase)

	return err
}

type NullablePatchSoftwareTitleConfigurationBase struct {
	value *PatchSoftwareTitleConfigurationBase
	isSet bool
}

func (v NullablePatchSoftwareTitleConfigurationBase) Get() *PatchSoftwareTitleConfigurationBase {
	return v.value
}

func (v *NullablePatchSoftwareTitleConfigurationBase) Set(val *PatchSoftwareTitleConfigurationBase) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchSoftwareTitleConfigurationBase) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchSoftwareTitleConfigurationBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchSoftwareTitleConfigurationBase(val *PatchSoftwareTitleConfigurationBase) *NullablePatchSoftwareTitleConfigurationBase {
	return &NullablePatchSoftwareTitleConfigurationBase{value: val, isSet: true}
}

func (v NullablePatchSoftwareTitleConfigurationBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchSoftwareTitleConfigurationBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


