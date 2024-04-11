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

// checks if the VolumePurchasingContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumePurchasingContent{}

// VolumePurchasingContent struct for VolumePurchasingContent
type VolumePurchasingContent struct {
	Name *string `json:"name,omitempty"`
	LicenseCountTotal *int64 `json:"licenseCountTotal,omitempty"`
	LicenseCountInUse *int64 `json:"licenseCountInUse,omitempty"`
	LicenseCountReported *int64 `json:"licenseCountReported,omitempty"`
	IconUrl *string `json:"iconUrl,omitempty"`
	DeviceTypes []string `json:"deviceTypes,omitempty"`
	ContentType *string `json:"contentType,omitempty"`
	PricingParam *string `json:"pricingParam,omitempty"`
	AdamId *string `json:"adamId,omitempty"`
}

// NewVolumePurchasingContent instantiates a new VolumePurchasingContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumePurchasingContent() *VolumePurchasingContent {
	this := VolumePurchasingContent{}
	return &this
}

// NewVolumePurchasingContentWithDefaults instantiates a new VolumePurchasingContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumePurchasingContentWithDefaults() *VolumePurchasingContent {
	this := VolumePurchasingContent{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VolumePurchasingContent) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingContent) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VolumePurchasingContent) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VolumePurchasingContent) SetName(v string) {
	o.Name = &v
}

// GetLicenseCountTotal returns the LicenseCountTotal field value if set, zero value otherwise.
func (o *VolumePurchasingContent) GetLicenseCountTotal() int64 {
	if o == nil || IsNil(o.LicenseCountTotal) {
		var ret int64
		return ret
	}
	return *o.LicenseCountTotal
}

// GetLicenseCountTotalOk returns a tuple with the LicenseCountTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingContent) GetLicenseCountTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.LicenseCountTotal) {
		return nil, false
	}
	return o.LicenseCountTotal, true
}

// HasLicenseCountTotal returns a boolean if a field has been set.
func (o *VolumePurchasingContent) HasLicenseCountTotal() bool {
	if o != nil && !IsNil(o.LicenseCountTotal) {
		return true
	}

	return false
}

// SetLicenseCountTotal gets a reference to the given int64 and assigns it to the LicenseCountTotal field.
func (o *VolumePurchasingContent) SetLicenseCountTotal(v int64) {
	o.LicenseCountTotal = &v
}

// GetLicenseCountInUse returns the LicenseCountInUse field value if set, zero value otherwise.
func (o *VolumePurchasingContent) GetLicenseCountInUse() int64 {
	if o == nil || IsNil(o.LicenseCountInUse) {
		var ret int64
		return ret
	}
	return *o.LicenseCountInUse
}

// GetLicenseCountInUseOk returns a tuple with the LicenseCountInUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingContent) GetLicenseCountInUseOk() (*int64, bool) {
	if o == nil || IsNil(o.LicenseCountInUse) {
		return nil, false
	}
	return o.LicenseCountInUse, true
}

// HasLicenseCountInUse returns a boolean if a field has been set.
func (o *VolumePurchasingContent) HasLicenseCountInUse() bool {
	if o != nil && !IsNil(o.LicenseCountInUse) {
		return true
	}

	return false
}

// SetLicenseCountInUse gets a reference to the given int64 and assigns it to the LicenseCountInUse field.
func (o *VolumePurchasingContent) SetLicenseCountInUse(v int64) {
	o.LicenseCountInUse = &v
}

// GetLicenseCountReported returns the LicenseCountReported field value if set, zero value otherwise.
func (o *VolumePurchasingContent) GetLicenseCountReported() int64 {
	if o == nil || IsNil(o.LicenseCountReported) {
		var ret int64
		return ret
	}
	return *o.LicenseCountReported
}

// GetLicenseCountReportedOk returns a tuple with the LicenseCountReported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingContent) GetLicenseCountReportedOk() (*int64, bool) {
	if o == nil || IsNil(o.LicenseCountReported) {
		return nil, false
	}
	return o.LicenseCountReported, true
}

// HasLicenseCountReported returns a boolean if a field has been set.
func (o *VolumePurchasingContent) HasLicenseCountReported() bool {
	if o != nil && !IsNil(o.LicenseCountReported) {
		return true
	}

	return false
}

// SetLicenseCountReported gets a reference to the given int64 and assigns it to the LicenseCountReported field.
func (o *VolumePurchasingContent) SetLicenseCountReported(v int64) {
	o.LicenseCountReported = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *VolumePurchasingContent) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl) {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingContent) GetIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IconUrl) {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *VolumePurchasingContent) HasIconUrl() bool {
	if o != nil && !IsNil(o.IconUrl) {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *VolumePurchasingContent) SetIconUrl(v string) {
	o.IconUrl = &v
}

// GetDeviceTypes returns the DeviceTypes field value if set, zero value otherwise.
func (o *VolumePurchasingContent) GetDeviceTypes() []string {
	if o == nil || IsNil(o.DeviceTypes) {
		var ret []string
		return ret
	}
	return o.DeviceTypes
}

// GetDeviceTypesOk returns a tuple with the DeviceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingContent) GetDeviceTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.DeviceTypes) {
		return nil, false
	}
	return o.DeviceTypes, true
}

// HasDeviceTypes returns a boolean if a field has been set.
func (o *VolumePurchasingContent) HasDeviceTypes() bool {
	if o != nil && !IsNil(o.DeviceTypes) {
		return true
	}

	return false
}

// SetDeviceTypes gets a reference to the given []string and assigns it to the DeviceTypes field.
func (o *VolumePurchasingContent) SetDeviceTypes(v []string) {
	o.DeviceTypes = v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *VolumePurchasingContent) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingContent) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *VolumePurchasingContent) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *VolumePurchasingContent) SetContentType(v string) {
	o.ContentType = &v
}

// GetPricingParam returns the PricingParam field value if set, zero value otherwise.
func (o *VolumePurchasingContent) GetPricingParam() string {
	if o == nil || IsNil(o.PricingParam) {
		var ret string
		return ret
	}
	return *o.PricingParam
}

// GetPricingParamOk returns a tuple with the PricingParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingContent) GetPricingParamOk() (*string, bool) {
	if o == nil || IsNil(o.PricingParam) {
		return nil, false
	}
	return o.PricingParam, true
}

// HasPricingParam returns a boolean if a field has been set.
func (o *VolumePurchasingContent) HasPricingParam() bool {
	if o != nil && !IsNil(o.PricingParam) {
		return true
	}

	return false
}

// SetPricingParam gets a reference to the given string and assigns it to the PricingParam field.
func (o *VolumePurchasingContent) SetPricingParam(v string) {
	o.PricingParam = &v
}

// GetAdamId returns the AdamId field value if set, zero value otherwise.
func (o *VolumePurchasingContent) GetAdamId() string {
	if o == nil || IsNil(o.AdamId) {
		var ret string
		return ret
	}
	return *o.AdamId
}

// GetAdamIdOk returns a tuple with the AdamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumePurchasingContent) GetAdamIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdamId) {
		return nil, false
	}
	return o.AdamId, true
}

// HasAdamId returns a boolean if a field has been set.
func (o *VolumePurchasingContent) HasAdamId() bool {
	if o != nil && !IsNil(o.AdamId) {
		return true
	}

	return false
}

// SetAdamId gets a reference to the given string and assigns it to the AdamId field.
func (o *VolumePurchasingContent) SetAdamId(v string) {
	o.AdamId = &v
}

func (o VolumePurchasingContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumePurchasingContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.LicenseCountTotal) {
		toSerialize["licenseCountTotal"] = o.LicenseCountTotal
	}
	if !IsNil(o.LicenseCountInUse) {
		toSerialize["licenseCountInUse"] = o.LicenseCountInUse
	}
	if !IsNil(o.LicenseCountReported) {
		toSerialize["licenseCountReported"] = o.LicenseCountReported
	}
	if !IsNil(o.IconUrl) {
		toSerialize["iconUrl"] = o.IconUrl
	}
	if !IsNil(o.DeviceTypes) {
		toSerialize["deviceTypes"] = o.DeviceTypes
	}
	if !IsNil(o.ContentType) {
		toSerialize["contentType"] = o.ContentType
	}
	if !IsNil(o.PricingParam) {
		toSerialize["pricingParam"] = o.PricingParam
	}
	if !IsNil(o.AdamId) {
		toSerialize["adamId"] = o.AdamId
	}
	return toSerialize, nil
}

type NullableVolumePurchasingContent struct {
	value *VolumePurchasingContent
	isSet bool
}

func (v NullableVolumePurchasingContent) Get() *VolumePurchasingContent {
	return v.value
}

func (v *NullableVolumePurchasingContent) Set(val *VolumePurchasingContent) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumePurchasingContent) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumePurchasingContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumePurchasingContent(val *VolumePurchasingContent) *NullableVolumePurchasingContent {
	return &NullableVolumePurchasingContent{value: val, isSet: true}
}

func (v NullableVolumePurchasingContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumePurchasingContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


