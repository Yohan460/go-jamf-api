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

// ComputerPrinter struct for ComputerPrinter
type ComputerPrinter struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Uri *string `json:"uri,omitempty"`
	Location *string `json:"location,omitempty"`
}

// NewComputerPrinter instantiates a new ComputerPrinter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputerPrinter() *ComputerPrinter {
	this := ComputerPrinter{}
	return &this
}

// NewComputerPrinterWithDefaults instantiates a new ComputerPrinter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputerPrinterWithDefaults() *ComputerPrinter {
	this := ComputerPrinter{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComputerPrinter) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPrinter) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComputerPrinter) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComputerPrinter) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ComputerPrinter) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPrinter) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ComputerPrinter) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ComputerPrinter) SetType(v string) {
	o.Type = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ComputerPrinter) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPrinter) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ComputerPrinter) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ComputerPrinter) SetUri(v string) {
	o.Uri = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ComputerPrinter) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPrinter) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ComputerPrinter) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *ComputerPrinter) SetLocation(v string) {
	o.Location = &v
}

func (o ComputerPrinter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	return json.Marshal(toSerialize)
}

type NullableComputerPrinter struct {
	value *ComputerPrinter
	isSet bool
}

func (v NullableComputerPrinter) Get() *ComputerPrinter {
	return v.value
}

func (v *NullableComputerPrinter) Set(val *ComputerPrinter) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerPrinter) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerPrinter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerPrinter(val *ComputerPrinter) *NullableComputerPrinter {
	return &NullableComputerPrinter{value: val, isSet: true}
}

func (v NullableComputerPrinter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerPrinter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


