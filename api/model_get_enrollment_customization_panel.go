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

// checks if the GetEnrollmentCustomizationPanel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEnrollmentCustomizationPanel{}

// GetEnrollmentCustomizationPanel struct for GetEnrollmentCustomizationPanel
type GetEnrollmentCustomizationPanel struct {
	DisplayName string `json:"displayName"`
	Rank int64 `json:"rank"`
	Id *int64 `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
}

type _GetEnrollmentCustomizationPanel GetEnrollmentCustomizationPanel

// NewGetEnrollmentCustomizationPanel instantiates a new GetEnrollmentCustomizationPanel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEnrollmentCustomizationPanel(displayName string, rank int64) *GetEnrollmentCustomizationPanel {
	this := GetEnrollmentCustomizationPanel{}
	this.DisplayName = displayName
	this.Rank = rank
	return &this
}

// NewGetEnrollmentCustomizationPanelWithDefaults instantiates a new GetEnrollmentCustomizationPanel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEnrollmentCustomizationPanelWithDefaults() *GetEnrollmentCustomizationPanel {
	this := GetEnrollmentCustomizationPanel{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *GetEnrollmentCustomizationPanel) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *GetEnrollmentCustomizationPanel) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *GetEnrollmentCustomizationPanel) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetRank returns the Rank field value
func (o *GetEnrollmentCustomizationPanel) GetRank() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Rank
}

// GetRankOk returns a tuple with the Rank field value
// and a boolean to check if the value has been set.
func (o *GetEnrollmentCustomizationPanel) GetRankOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rank, true
}

// SetRank sets field value
func (o *GetEnrollmentCustomizationPanel) SetRank(v int64) {
	o.Rank = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetEnrollmentCustomizationPanel) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEnrollmentCustomizationPanel) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetEnrollmentCustomizationPanel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetEnrollmentCustomizationPanel) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetEnrollmentCustomizationPanel) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEnrollmentCustomizationPanel) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetEnrollmentCustomizationPanel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetEnrollmentCustomizationPanel) SetType(v string) {
	o.Type = &v
}

func (o GetEnrollmentCustomizationPanel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEnrollmentCustomizationPanel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["rank"] = o.Rank
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *GetEnrollmentCustomizationPanel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"displayName",
		"rank",
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

	varGetEnrollmentCustomizationPanel := _GetEnrollmentCustomizationPanel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetEnrollmentCustomizationPanel)

	if err != nil {
		return err
	}

	*o = GetEnrollmentCustomizationPanel(varGetEnrollmentCustomizationPanel)

	return err
}

type NullableGetEnrollmentCustomizationPanel struct {
	value *GetEnrollmentCustomizationPanel
	isSet bool
}

func (v NullableGetEnrollmentCustomizationPanel) Get() *GetEnrollmentCustomizationPanel {
	return v.value
}

func (v *NullableGetEnrollmentCustomizationPanel) Set(val *GetEnrollmentCustomizationPanel) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEnrollmentCustomizationPanel) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEnrollmentCustomizationPanel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEnrollmentCustomizationPanel(val *GetEnrollmentCustomizationPanel) *NullableGetEnrollmentCustomizationPanel {
	return &NullableGetEnrollmentCustomizationPanel{value: val, isSet: true}
}

func (v NullableGetEnrollmentCustomizationPanel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEnrollmentCustomizationPanel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


