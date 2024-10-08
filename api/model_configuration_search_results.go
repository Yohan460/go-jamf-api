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

// checks if the ConfigurationSearchResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationSearchResults{}

// ConfigurationSearchResults A list with Cloud Identity Providers informations about configurations
type ConfigurationSearchResults struct {
	TotalCount *int64 `json:"totalCount,omitempty"`
	Results []CloudIdPCommonResponse `json:"results,omitempty"`
}

// NewConfigurationSearchResults instantiates a new ConfigurationSearchResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationSearchResults() *ConfigurationSearchResults {
	this := ConfigurationSearchResults{}
	return &this
}

// NewConfigurationSearchResultsWithDefaults instantiates a new ConfigurationSearchResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationSearchResultsWithDefaults() *ConfigurationSearchResults {
	this := ConfigurationSearchResults{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ConfigurationSearchResults) GetTotalCount() int64 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationSearchResults) GetTotalCountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ConfigurationSearchResults) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *ConfigurationSearchResults) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ConfigurationSearchResults) GetResults() []CloudIdPCommonResponse {
	if o == nil || IsNil(o.Results) {
		var ret []CloudIdPCommonResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationSearchResults) GetResultsOk() ([]CloudIdPCommonResponse, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ConfigurationSearchResults) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []CloudIdPCommonResponse and assigns it to the Results field.
func (o *ConfigurationSearchResults) SetResults(v []CloudIdPCommonResponse) {
	o.Results = v
}

func (o ConfigurationSearchResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationSearchResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableConfigurationSearchResults struct {
	value *ConfigurationSearchResults
	isSet bool
}

func (v NullableConfigurationSearchResults) Get() *ConfigurationSearchResults {
	return v.value
}

func (v *NullableConfigurationSearchResults) Set(val *ConfigurationSearchResults) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationSearchResults) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationSearchResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationSearchResults(val *ConfigurationSearchResults) *NullableConfigurationSearchResults {
	return &NullableConfigurationSearchResults{value: val, isSet: true}
}

func (v NullableConfigurationSearchResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationSearchResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


