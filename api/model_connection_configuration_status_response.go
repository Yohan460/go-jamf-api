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

// checks if the ConnectionConfigurationStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionConfigurationStatusResponse{}

// ConnectionConfigurationStatusResponse Response that contains connection configuration status for Team Viewer
type ConnectionConfigurationStatusResponse struct {
	// connection configuration status for Team Viewer
	ConnectionVerificationResult string `json:"connectionVerificationResult"`
}

type _ConnectionConfigurationStatusResponse ConnectionConfigurationStatusResponse

// NewConnectionConfigurationStatusResponse instantiates a new ConnectionConfigurationStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionConfigurationStatusResponse(connectionVerificationResult string) *ConnectionConfigurationStatusResponse {
	this := ConnectionConfigurationStatusResponse{}
	this.ConnectionVerificationResult = connectionVerificationResult
	return &this
}

// NewConnectionConfigurationStatusResponseWithDefaults instantiates a new ConnectionConfigurationStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionConfigurationStatusResponseWithDefaults() *ConnectionConfigurationStatusResponse {
	this := ConnectionConfigurationStatusResponse{}
	return &this
}

// GetConnectionVerificationResult returns the ConnectionVerificationResult field value
func (o *ConnectionConfigurationStatusResponse) GetConnectionVerificationResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionVerificationResult
}

// GetConnectionVerificationResultOk returns a tuple with the ConnectionVerificationResult field value
// and a boolean to check if the value has been set.
func (o *ConnectionConfigurationStatusResponse) GetConnectionVerificationResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionVerificationResult, true
}

// SetConnectionVerificationResult sets field value
func (o *ConnectionConfigurationStatusResponse) SetConnectionVerificationResult(v string) {
	o.ConnectionVerificationResult = v
}

func (o ConnectionConfigurationStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionConfigurationStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connectionVerificationResult"] = o.ConnectionVerificationResult
	return toSerialize, nil
}

func (o *ConnectionConfigurationStatusResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connectionVerificationResult",
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

	varConnectionConfigurationStatusResponse := _ConnectionConfigurationStatusResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConnectionConfigurationStatusResponse)

	if err != nil {
		return err
	}

	*o = ConnectionConfigurationStatusResponse(varConnectionConfigurationStatusResponse)

	return err
}

type NullableConnectionConfigurationStatusResponse struct {
	value *ConnectionConfigurationStatusResponse
	isSet bool
}

func (v NullableConnectionConfigurationStatusResponse) Get() *ConnectionConfigurationStatusResponse {
	return v.value
}

func (v *NullableConnectionConfigurationStatusResponse) Set(val *ConnectionConfigurationStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionConfigurationStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionConfigurationStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionConfigurationStatusResponse(val *ConnectionConfigurationStatusResponse) *NullableConnectionConfigurationStatusResponse {
	return &NullableConnectionConfigurationStatusResponse{value: val, isSet: true}
}

func (v NullableConnectionConfigurationStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionConfigurationStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


