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

// checks if the ShutDownDeviceCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShutDownDeviceCommand{}

// ShutDownDeviceCommand struct for ShutDownDeviceCommand
type ShutDownDeviceCommand struct {
	CommandType MdmCommandType `json:"commandType"`
}

type _ShutDownDeviceCommand ShutDownDeviceCommand

// NewShutDownDeviceCommand instantiates a new ShutDownDeviceCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShutDownDeviceCommand(commandType MdmCommandType) *ShutDownDeviceCommand {
	this := ShutDownDeviceCommand{}
	this.CommandType = commandType
	return &this
}

// NewShutDownDeviceCommandWithDefaults instantiates a new ShutDownDeviceCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShutDownDeviceCommandWithDefaults() *ShutDownDeviceCommand {
	this := ShutDownDeviceCommand{}
	return &this
}

// GetCommandType returns the CommandType field value
func (o *ShutDownDeviceCommand) GetCommandType() MdmCommandType {
	if o == nil {
		var ret MdmCommandType
		return ret
	}

	return o.CommandType
}

// GetCommandTypeOk returns a tuple with the CommandType field value
// and a boolean to check if the value has been set.
func (o *ShutDownDeviceCommand) GetCommandTypeOk() (*MdmCommandType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommandType, true
}

// SetCommandType sets field value
func (o *ShutDownDeviceCommand) SetCommandType(v MdmCommandType) {
	o.CommandType = v
}

func (o ShutDownDeviceCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShutDownDeviceCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["commandType"] = o.CommandType
	return toSerialize, nil
}

func (o *ShutDownDeviceCommand) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"commandType",
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

	varShutDownDeviceCommand := _ShutDownDeviceCommand{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShutDownDeviceCommand)

	if err != nil {
		return err
	}

	*o = ShutDownDeviceCommand(varShutDownDeviceCommand)

	return err
}

type NullableShutDownDeviceCommand struct {
	value *ShutDownDeviceCommand
	isSet bool
}

func (v NullableShutDownDeviceCommand) Get() *ShutDownDeviceCommand {
	return v.value
}

func (v *NullableShutDownDeviceCommand) Set(val *ShutDownDeviceCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableShutDownDeviceCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableShutDownDeviceCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShutDownDeviceCommand(val *ShutDownDeviceCommand) *NullableShutDownDeviceCommand {
	return &NullableShutDownDeviceCommand{value: val, isSet: true}
}

func (v NullableShutDownDeviceCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShutDownDeviceCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


