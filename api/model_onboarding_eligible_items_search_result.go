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

// checks if the OnboardingEligibleItemsSearchResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnboardingEligibleItemsSearchResult{}

// OnboardingEligibleItemsSearchResult A list of onboarding eligible items
type OnboardingEligibleItemsSearchResult struct {
	TotalCount *int64 `json:"totalCount,omitempty"`
	Results []OnboardingEligibleItem `json:"results,omitempty"`
}

// NewOnboardingEligibleItemsSearchResult instantiates a new OnboardingEligibleItemsSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnboardingEligibleItemsSearchResult() *OnboardingEligibleItemsSearchResult {
	this := OnboardingEligibleItemsSearchResult{}
	return &this
}

// NewOnboardingEligibleItemsSearchResultWithDefaults instantiates a new OnboardingEligibleItemsSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnboardingEligibleItemsSearchResultWithDefaults() *OnboardingEligibleItemsSearchResult {
	this := OnboardingEligibleItemsSearchResult{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *OnboardingEligibleItemsSearchResult) GetTotalCount() int64 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingEligibleItemsSearchResult) GetTotalCountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *OnboardingEligibleItemsSearchResult) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *OnboardingEligibleItemsSearchResult) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *OnboardingEligibleItemsSearchResult) GetResults() []OnboardingEligibleItem {
	if o == nil || IsNil(o.Results) {
		var ret []OnboardingEligibleItem
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnboardingEligibleItemsSearchResult) GetResultsOk() ([]OnboardingEligibleItem, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *OnboardingEligibleItemsSearchResult) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []OnboardingEligibleItem and assigns it to the Results field.
func (o *OnboardingEligibleItemsSearchResult) SetResults(v []OnboardingEligibleItem) {
	o.Results = v
}

func (o OnboardingEligibleItemsSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnboardingEligibleItemsSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableOnboardingEligibleItemsSearchResult struct {
	value *OnboardingEligibleItemsSearchResult
	isSet bool
}

func (v NullableOnboardingEligibleItemsSearchResult) Get() *OnboardingEligibleItemsSearchResult {
	return v.value
}

func (v *NullableOnboardingEligibleItemsSearchResult) Set(val *OnboardingEligibleItemsSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableOnboardingEligibleItemsSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableOnboardingEligibleItemsSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnboardingEligibleItemsSearchResult(val *OnboardingEligibleItemsSearchResult) *NullableOnboardingEligibleItemsSearchResult {
	return &NullableOnboardingEligibleItemsSearchResult{value: val, isSet: true}
}

func (v NullableOnboardingEligibleItemsSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnboardingEligibleItemsSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


