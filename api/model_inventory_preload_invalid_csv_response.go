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

// checks if the InventoryPreloadInvalidCsvResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryPreloadInvalidCsvResponse{}

// InventoryPreloadInvalidCsvResponse struct for InventoryPreloadInvalidCsvResponse
type InventoryPreloadInvalidCsvResponse struct {
	HttpsStatus *int32 `json:"httpsStatus,omitempty"`
	Errors []InventoryPreloadCsvError `json:"errors,omitempty"`
}

// NewInventoryPreloadInvalidCsvResponse instantiates a new InventoryPreloadInvalidCsvResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryPreloadInvalidCsvResponse() *InventoryPreloadInvalidCsvResponse {
	this := InventoryPreloadInvalidCsvResponse{}
	return &this
}

// NewInventoryPreloadInvalidCsvResponseWithDefaults instantiates a new InventoryPreloadInvalidCsvResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryPreloadInvalidCsvResponseWithDefaults() *InventoryPreloadInvalidCsvResponse {
	this := InventoryPreloadInvalidCsvResponse{}
	return &this
}

// GetHttpsStatus returns the HttpsStatus field value if set, zero value otherwise.
func (o *InventoryPreloadInvalidCsvResponse) GetHttpsStatus() int32 {
	if o == nil || IsNil(o.HttpsStatus) {
		var ret int32
		return ret
	}
	return *o.HttpsStatus
}

// GetHttpsStatusOk returns a tuple with the HttpsStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadInvalidCsvResponse) GetHttpsStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.HttpsStatus) {
		return nil, false
	}
	return o.HttpsStatus, true
}

// HasHttpsStatus returns a boolean if a field has been set.
func (o *InventoryPreloadInvalidCsvResponse) HasHttpsStatus() bool {
	if o != nil && !IsNil(o.HttpsStatus) {
		return true
	}

	return false
}

// SetHttpsStatus gets a reference to the given int32 and assigns it to the HttpsStatus field.
func (o *InventoryPreloadInvalidCsvResponse) SetHttpsStatus(v int32) {
	o.HttpsStatus = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *InventoryPreloadInvalidCsvResponse) GetErrors() []InventoryPreloadCsvError {
	if o == nil || IsNil(o.Errors) {
		var ret []InventoryPreloadCsvError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadInvalidCsvResponse) GetErrorsOk() ([]InventoryPreloadCsvError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *InventoryPreloadInvalidCsvResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []InventoryPreloadCsvError and assigns it to the Errors field.
func (o *InventoryPreloadInvalidCsvResponse) SetErrors(v []InventoryPreloadCsvError) {
	o.Errors = v
}

func (o InventoryPreloadInvalidCsvResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryPreloadInvalidCsvResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HttpsStatus) {
		toSerialize["httpsStatus"] = o.HttpsStatus
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableInventoryPreloadInvalidCsvResponse struct {
	value *InventoryPreloadInvalidCsvResponse
	isSet bool
}

func (v NullableInventoryPreloadInvalidCsvResponse) Get() *InventoryPreloadInvalidCsvResponse {
	return v.value
}

func (v *NullableInventoryPreloadInvalidCsvResponse) Set(val *InventoryPreloadInvalidCsvResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryPreloadInvalidCsvResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryPreloadInvalidCsvResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryPreloadInvalidCsvResponse(val *InventoryPreloadInvalidCsvResponse) *NullableInventoryPreloadInvalidCsvResponse {
	return &NullableInventoryPreloadInvalidCsvResponse{value: val, isSet: true}
}

func (v NullableInventoryPreloadInvalidCsvResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryPreloadInvalidCsvResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


