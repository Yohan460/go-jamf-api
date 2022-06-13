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

// AppDynamicsConfig Configuration parameters needed for AppDynamics script initialization.
type AppDynamicsConfig struct {
	EnableEum bool `json:"enableEum"`
	AppKey string `json:"appKey"`
	AdrumExtUrlHttp string `json:"adrumExtUrlHttp"`
	AdrumExtUrlHttps string `json:"adrumExtUrlHttps"`
	AdrumScriptHttp string `json:"adrumScriptHttp"`
	AdrumScriptHttps string `json:"adrumScriptHttps"`
	BeaconUrlHttp string `json:"beaconUrlHttp"`
	BeaconUrlHttps string `json:"beaconUrlHttps"`
}

// NewAppDynamicsConfig instantiates a new AppDynamicsConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppDynamicsConfig(enableEum bool, appKey string, adrumExtUrlHttp string, adrumExtUrlHttps string, adrumScriptHttp string, adrumScriptHttps string, beaconUrlHttp string, beaconUrlHttps string) *AppDynamicsConfig {
	this := AppDynamicsConfig{}
	this.EnableEum = enableEum
	this.AppKey = appKey
	this.AdrumExtUrlHttp = adrumExtUrlHttp
	this.AdrumExtUrlHttps = adrumExtUrlHttps
	this.AdrumScriptHttp = adrumScriptHttp
	this.AdrumScriptHttps = adrumScriptHttps
	this.BeaconUrlHttp = beaconUrlHttp
	this.BeaconUrlHttps = beaconUrlHttps
	return &this
}

// NewAppDynamicsConfigWithDefaults instantiates a new AppDynamicsConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppDynamicsConfigWithDefaults() *AppDynamicsConfig {
	this := AppDynamicsConfig{}
	var enableEum bool = false
	this.EnableEum = enableEum
	return &this
}

// GetEnableEum returns the EnableEum field value
func (o *AppDynamicsConfig) GetEnableEum() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableEum
}

// GetEnableEumOk returns a tuple with the EnableEum field value
// and a boolean to check if the value has been set.
func (o *AppDynamicsConfig) GetEnableEumOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableEum, true
}

// SetEnableEum sets field value
func (o *AppDynamicsConfig) SetEnableEum(v bool) {
	o.EnableEum = v
}

// GetAppKey returns the AppKey field value
func (o *AppDynamicsConfig) GetAppKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppKey
}

// GetAppKeyOk returns a tuple with the AppKey field value
// and a boolean to check if the value has been set.
func (o *AppDynamicsConfig) GetAppKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppKey, true
}

// SetAppKey sets field value
func (o *AppDynamicsConfig) SetAppKey(v string) {
	o.AppKey = v
}

// GetAdrumExtUrlHttp returns the AdrumExtUrlHttp field value
func (o *AppDynamicsConfig) GetAdrumExtUrlHttp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdrumExtUrlHttp
}

// GetAdrumExtUrlHttpOk returns a tuple with the AdrumExtUrlHttp field value
// and a boolean to check if the value has been set.
func (o *AppDynamicsConfig) GetAdrumExtUrlHttpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdrumExtUrlHttp, true
}

// SetAdrumExtUrlHttp sets field value
func (o *AppDynamicsConfig) SetAdrumExtUrlHttp(v string) {
	o.AdrumExtUrlHttp = v
}

// GetAdrumExtUrlHttps returns the AdrumExtUrlHttps field value
func (o *AppDynamicsConfig) GetAdrumExtUrlHttps() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdrumExtUrlHttps
}

// GetAdrumExtUrlHttpsOk returns a tuple with the AdrumExtUrlHttps field value
// and a boolean to check if the value has been set.
func (o *AppDynamicsConfig) GetAdrumExtUrlHttpsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdrumExtUrlHttps, true
}

// SetAdrumExtUrlHttps sets field value
func (o *AppDynamicsConfig) SetAdrumExtUrlHttps(v string) {
	o.AdrumExtUrlHttps = v
}

// GetAdrumScriptHttp returns the AdrumScriptHttp field value
func (o *AppDynamicsConfig) GetAdrumScriptHttp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdrumScriptHttp
}

// GetAdrumScriptHttpOk returns a tuple with the AdrumScriptHttp field value
// and a boolean to check if the value has been set.
func (o *AppDynamicsConfig) GetAdrumScriptHttpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdrumScriptHttp, true
}

// SetAdrumScriptHttp sets field value
func (o *AppDynamicsConfig) SetAdrumScriptHttp(v string) {
	o.AdrumScriptHttp = v
}

// GetAdrumScriptHttps returns the AdrumScriptHttps field value
func (o *AppDynamicsConfig) GetAdrumScriptHttps() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdrumScriptHttps
}

// GetAdrumScriptHttpsOk returns a tuple with the AdrumScriptHttps field value
// and a boolean to check if the value has been set.
func (o *AppDynamicsConfig) GetAdrumScriptHttpsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdrumScriptHttps, true
}

// SetAdrumScriptHttps sets field value
func (o *AppDynamicsConfig) SetAdrumScriptHttps(v string) {
	o.AdrumScriptHttps = v
}

// GetBeaconUrlHttp returns the BeaconUrlHttp field value
func (o *AppDynamicsConfig) GetBeaconUrlHttp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BeaconUrlHttp
}

// GetBeaconUrlHttpOk returns a tuple with the BeaconUrlHttp field value
// and a boolean to check if the value has been set.
func (o *AppDynamicsConfig) GetBeaconUrlHttpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BeaconUrlHttp, true
}

// SetBeaconUrlHttp sets field value
func (o *AppDynamicsConfig) SetBeaconUrlHttp(v string) {
	o.BeaconUrlHttp = v
}

// GetBeaconUrlHttps returns the BeaconUrlHttps field value
func (o *AppDynamicsConfig) GetBeaconUrlHttps() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BeaconUrlHttps
}

// GetBeaconUrlHttpsOk returns a tuple with the BeaconUrlHttps field value
// and a boolean to check if the value has been set.
func (o *AppDynamicsConfig) GetBeaconUrlHttpsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BeaconUrlHttps, true
}

// SetBeaconUrlHttps sets field value
func (o *AppDynamicsConfig) SetBeaconUrlHttps(v string) {
	o.BeaconUrlHttps = v
}

func (o AppDynamicsConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enableEum"] = o.EnableEum
	}
	if true {
		toSerialize["appKey"] = o.AppKey
	}
	if true {
		toSerialize["adrumExtUrlHttp"] = o.AdrumExtUrlHttp
	}
	if true {
		toSerialize["adrumExtUrlHttps"] = o.AdrumExtUrlHttps
	}
	if true {
		toSerialize["adrumScriptHttp"] = o.AdrumScriptHttp
	}
	if true {
		toSerialize["adrumScriptHttps"] = o.AdrumScriptHttps
	}
	if true {
		toSerialize["beaconUrlHttp"] = o.BeaconUrlHttp
	}
	if true {
		toSerialize["beaconUrlHttps"] = o.BeaconUrlHttps
	}
	return json.Marshal(toSerialize)
}

type NullableAppDynamicsConfig struct {
	value *AppDynamicsConfig
	isSet bool
}

func (v NullableAppDynamicsConfig) Get() *AppDynamicsConfig {
	return v.value
}

func (v *NullableAppDynamicsConfig) Set(val *AppDynamicsConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAppDynamicsConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAppDynamicsConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppDynamicsConfig(val *AppDynamicsConfig) *NullableAppDynamicsConfig {
	return &NullableAppDynamicsConfig{value: val, isSet: true}
}

func (v NullableAppDynamicsConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppDynamicsConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


