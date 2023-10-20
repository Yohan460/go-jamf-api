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

// checks if the IosBrandingConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IosBrandingConfiguration{}

// IosBrandingConfiguration struct for IosBrandingConfiguration
type IosBrandingConfiguration struct {
	Id *string `json:"id,omitempty"`
	BrandingName string `json:"brandingName"`
	IconId *int32 `json:"iconId,omitempty"`
	HeaderBackgroundColorCode string `json:"headerBackgroundColorCode"`
	MenuIconColorCode string `json:"menuIconColorCode"`
	BrandingNameColorCode string `json:"brandingNameColorCode"`
	StatusBarTextColor string `json:"statusBarTextColor"`
}

// NewIosBrandingConfiguration instantiates a new IosBrandingConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIosBrandingConfiguration(brandingName string, headerBackgroundColorCode string, menuIconColorCode string, brandingNameColorCode string, statusBarTextColor string) *IosBrandingConfiguration {
	this := IosBrandingConfiguration{}
	this.BrandingName = brandingName
	this.HeaderBackgroundColorCode = headerBackgroundColorCode
	this.MenuIconColorCode = menuIconColorCode
	this.BrandingNameColorCode = brandingNameColorCode
	this.StatusBarTextColor = statusBarTextColor
	return &this
}

// NewIosBrandingConfigurationWithDefaults instantiates a new IosBrandingConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIosBrandingConfigurationWithDefaults() *IosBrandingConfiguration {
	this := IosBrandingConfiguration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IosBrandingConfiguration) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosBrandingConfiguration) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IosBrandingConfiguration) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IosBrandingConfiguration) SetId(v string) {
	o.Id = &v
}

// GetBrandingName returns the BrandingName field value
func (o *IosBrandingConfiguration) GetBrandingName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandingName
}

// GetBrandingNameOk returns a tuple with the BrandingName field value
// and a boolean to check if the value has been set.
func (o *IosBrandingConfiguration) GetBrandingNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandingName, true
}

// SetBrandingName sets field value
func (o *IosBrandingConfiguration) SetBrandingName(v string) {
	o.BrandingName = v
}

// GetIconId returns the IconId field value if set, zero value otherwise.
func (o *IosBrandingConfiguration) GetIconId() int32 {
	if o == nil || IsNil(o.IconId) {
		var ret int32
		return ret
	}
	return *o.IconId
}

// GetIconIdOk returns a tuple with the IconId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosBrandingConfiguration) GetIconIdOk() (*int32, bool) {
	if o == nil || IsNil(o.IconId) {
		return nil, false
	}
	return o.IconId, true
}

// HasIconId returns a boolean if a field has been set.
func (o *IosBrandingConfiguration) HasIconId() bool {
	if o != nil && !IsNil(o.IconId) {
		return true
	}

	return false
}

// SetIconId gets a reference to the given int32 and assigns it to the IconId field.
func (o *IosBrandingConfiguration) SetIconId(v int32) {
	o.IconId = &v
}

// GetHeaderBackgroundColorCode returns the HeaderBackgroundColorCode field value
func (o *IosBrandingConfiguration) GetHeaderBackgroundColorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HeaderBackgroundColorCode
}

// GetHeaderBackgroundColorCodeOk returns a tuple with the HeaderBackgroundColorCode field value
// and a boolean to check if the value has been set.
func (o *IosBrandingConfiguration) GetHeaderBackgroundColorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HeaderBackgroundColorCode, true
}

// SetHeaderBackgroundColorCode sets field value
func (o *IosBrandingConfiguration) SetHeaderBackgroundColorCode(v string) {
	o.HeaderBackgroundColorCode = v
}

// GetMenuIconColorCode returns the MenuIconColorCode field value
func (o *IosBrandingConfiguration) GetMenuIconColorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MenuIconColorCode
}

// GetMenuIconColorCodeOk returns a tuple with the MenuIconColorCode field value
// and a boolean to check if the value has been set.
func (o *IosBrandingConfiguration) GetMenuIconColorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MenuIconColorCode, true
}

// SetMenuIconColorCode sets field value
func (o *IosBrandingConfiguration) SetMenuIconColorCode(v string) {
	o.MenuIconColorCode = v
}

// GetBrandingNameColorCode returns the BrandingNameColorCode field value
func (o *IosBrandingConfiguration) GetBrandingNameColorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandingNameColorCode
}

// GetBrandingNameColorCodeOk returns a tuple with the BrandingNameColorCode field value
// and a boolean to check if the value has been set.
func (o *IosBrandingConfiguration) GetBrandingNameColorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandingNameColorCode, true
}

// SetBrandingNameColorCode sets field value
func (o *IosBrandingConfiguration) SetBrandingNameColorCode(v string) {
	o.BrandingNameColorCode = v
}

// GetStatusBarTextColor returns the StatusBarTextColor field value
func (o *IosBrandingConfiguration) GetStatusBarTextColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusBarTextColor
}

// GetStatusBarTextColorOk returns a tuple with the StatusBarTextColor field value
// and a boolean to check if the value has been set.
func (o *IosBrandingConfiguration) GetStatusBarTextColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusBarTextColor, true
}

// SetStatusBarTextColor sets field value
func (o *IosBrandingConfiguration) SetStatusBarTextColor(v string) {
	o.StatusBarTextColor = v
}

func (o IosBrandingConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IosBrandingConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["brandingName"] = o.BrandingName
	if !IsNil(o.IconId) {
		toSerialize["iconId"] = o.IconId
	}
	toSerialize["headerBackgroundColorCode"] = o.HeaderBackgroundColorCode
	toSerialize["menuIconColorCode"] = o.MenuIconColorCode
	toSerialize["brandingNameColorCode"] = o.BrandingNameColorCode
	toSerialize["statusBarTextColor"] = o.StatusBarTextColor
	return toSerialize, nil
}

type NullableIosBrandingConfiguration struct {
	value *IosBrandingConfiguration
	isSet bool
}

func (v NullableIosBrandingConfiguration) Get() *IosBrandingConfiguration {
	return v.value
}

func (v *NullableIosBrandingConfiguration) Set(val *IosBrandingConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableIosBrandingConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableIosBrandingConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIosBrandingConfiguration(val *IosBrandingConfiguration) *NullableIosBrandingConfiguration {
	return &NullableIosBrandingConfiguration{value: val, isSet: true}
}

func (v NullableIosBrandingConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIosBrandingConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


