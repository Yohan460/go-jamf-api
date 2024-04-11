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

// checks if the AdvancedUserContentSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvancedUserContentSearch{}

// AdvancedUserContentSearch struct for AdvancedUserContentSearch
type AdvancedUserContentSearch struct {
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	Criteria []SmartSearchCriterion `json:"criteria,omitempty"`
	DisplayFields []string `json:"displayFields,omitempty"`
	SiteId NullableString `json:"siteId,omitempty"`
}

type _AdvancedUserContentSearch AdvancedUserContentSearch

// NewAdvancedUserContentSearch instantiates a new AdvancedUserContentSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvancedUserContentSearch(name string) *AdvancedUserContentSearch {
	this := AdvancedUserContentSearch{}
	this.Name = name
	return &this
}

// NewAdvancedUserContentSearchWithDefaults instantiates a new AdvancedUserContentSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvancedUserContentSearchWithDefaults() *AdvancedUserContentSearch {
	this := AdvancedUserContentSearch{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdvancedUserContentSearch) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedUserContentSearch) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdvancedUserContentSearch) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AdvancedUserContentSearch) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *AdvancedUserContentSearch) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AdvancedUserContentSearch) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AdvancedUserContentSearch) SetName(v string) {
	o.Name = v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *AdvancedUserContentSearch) GetCriteria() []SmartSearchCriterion {
	if o == nil || IsNil(o.Criteria) {
		var ret []SmartSearchCriterion
		return ret
	}
	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedUserContentSearch) GetCriteriaOk() ([]SmartSearchCriterion, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *AdvancedUserContentSearch) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given []SmartSearchCriterion and assigns it to the Criteria field.
func (o *AdvancedUserContentSearch) SetCriteria(v []SmartSearchCriterion) {
	o.Criteria = v
}

// GetDisplayFields returns the DisplayFields field value if set, zero value otherwise.
func (o *AdvancedUserContentSearch) GetDisplayFields() []string {
	if o == nil || IsNil(o.DisplayFields) {
		var ret []string
		return ret
	}
	return o.DisplayFields
}

// GetDisplayFieldsOk returns a tuple with the DisplayFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedUserContentSearch) GetDisplayFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.DisplayFields) {
		return nil, false
	}
	return o.DisplayFields, true
}

// HasDisplayFields returns a boolean if a field has been set.
func (o *AdvancedUserContentSearch) HasDisplayFields() bool {
	if o != nil && !IsNil(o.DisplayFields) {
		return true
	}

	return false
}

// SetDisplayFields gets a reference to the given []string and assigns it to the DisplayFields field.
func (o *AdvancedUserContentSearch) SetDisplayFields(v []string) {
	o.DisplayFields = v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdvancedUserContentSearch) GetSiteId() string {
	if o == nil || IsNil(o.SiteId.Get()) {
		var ret string
		return ret
	}
	return *o.SiteId.Get()
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdvancedUserContentSearch) GetSiteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SiteId.Get(), o.SiteId.IsSet()
}

// HasSiteId returns a boolean if a field has been set.
func (o *AdvancedUserContentSearch) HasSiteId() bool {
	if o != nil && o.SiteId.IsSet() {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given NullableString and assigns it to the SiteId field.
func (o *AdvancedUserContentSearch) SetSiteId(v string) {
	o.SiteId.Set(&v)
}
// SetSiteIdNil sets the value for SiteId to be an explicit nil
func (o *AdvancedUserContentSearch) SetSiteIdNil() {
	o.SiteId.Set(nil)
}

// UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
func (o *AdvancedUserContentSearch) UnsetSiteId() {
	o.SiteId.Unset()
}

func (o AdvancedUserContentSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvancedUserContentSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.DisplayFields) {
		toSerialize["displayFields"] = o.DisplayFields
	}
	if o.SiteId.IsSet() {
		toSerialize["siteId"] = o.SiteId.Get()
	}
	return toSerialize, nil
}

func (o *AdvancedUserContentSearch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varAdvancedUserContentSearch := _AdvancedUserContentSearch{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAdvancedUserContentSearch)

	if err != nil {
		return err
	}

	*o = AdvancedUserContentSearch(varAdvancedUserContentSearch)

	return err
}

type NullableAdvancedUserContentSearch struct {
	value *AdvancedUserContentSearch
	isSet bool
}

func (v NullableAdvancedUserContentSearch) Get() *AdvancedUserContentSearch {
	return v.value
}

func (v *NullableAdvancedUserContentSearch) Set(val *AdvancedUserContentSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvancedUserContentSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvancedUserContentSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvancedUserContentSearch(val *AdvancedUserContentSearch) *NullableAdvancedUserContentSearch {
	return &NullableAdvancedUserContentSearch{value: val, isSet: true}
}

func (v NullableAdvancedUserContentSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvancedUserContentSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


