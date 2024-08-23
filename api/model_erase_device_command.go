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

// checks if the EraseDeviceCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EraseDeviceCommand{}

// EraseDeviceCommand struct for EraseDeviceCommand
type EraseDeviceCommand struct {
	CommandType MdmCommandType `json:"commandType"`
	// If true, preserve the data plan on an iPhone or iPad with eSIM functionality, if one exists. This value is available in iOS 11 and later.
	PreserveDataPlan *bool `json:"preserveDataPlan,omitempty"`
	// If true, disable Proximity Setup on the next reboot and skip the pane in Setup Assistant. This value is available in iOS 11 and later. Prior to iOS 14, don’t use this option with any other option.
	DisallowProximitySetup *bool `json:"disallowProximitySetup,omitempty"`
	// The six-character PIN for Find My. This value is available in macOS 10.8 and later.
	Pin *string `json:"pin,omitempty"`
	// This key defines the fallback behavior for erasing a device.
	ObliterationBehavior *string `json:"obliterationBehavior,omitempty"`
	ReturnToService *EraseDeviceCommandAllOfReturnToService `json:"returnToService,omitempty"`
}

type _EraseDeviceCommand EraseDeviceCommand

// NewEraseDeviceCommand instantiates a new EraseDeviceCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEraseDeviceCommand(commandType MdmCommandType) *EraseDeviceCommand {
	this := EraseDeviceCommand{}
	this.CommandType = commandType
	var preserveDataPlan bool = false
	this.PreserveDataPlan = &preserveDataPlan
	var disallowProximitySetup bool = false
	this.DisallowProximitySetup = &disallowProximitySetup
	return &this
}

// NewEraseDeviceCommandWithDefaults instantiates a new EraseDeviceCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEraseDeviceCommandWithDefaults() *EraseDeviceCommand {
	this := EraseDeviceCommand{}
	var preserveDataPlan bool = false
	this.PreserveDataPlan = &preserveDataPlan
	var disallowProximitySetup bool = false
	this.DisallowProximitySetup = &disallowProximitySetup
	return &this
}

// GetCommandType returns the CommandType field value
func (o *EraseDeviceCommand) GetCommandType() MdmCommandType {
	if o == nil {
		var ret MdmCommandType
		return ret
	}

	return o.CommandType
}

// GetCommandTypeOk returns a tuple with the CommandType field value
// and a boolean to check if the value has been set.
func (o *EraseDeviceCommand) GetCommandTypeOk() (*MdmCommandType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommandType, true
}

// SetCommandType sets field value
func (o *EraseDeviceCommand) SetCommandType(v MdmCommandType) {
	o.CommandType = v
}

// GetPreserveDataPlan returns the PreserveDataPlan field value if set, zero value otherwise.
func (o *EraseDeviceCommand) GetPreserveDataPlan() bool {
	if o == nil || IsNil(o.PreserveDataPlan) {
		var ret bool
		return ret
	}
	return *o.PreserveDataPlan
}

// GetPreserveDataPlanOk returns a tuple with the PreserveDataPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EraseDeviceCommand) GetPreserveDataPlanOk() (*bool, bool) {
	if o == nil || IsNil(o.PreserveDataPlan) {
		return nil, false
	}
	return o.PreserveDataPlan, true
}

// HasPreserveDataPlan returns a boolean if a field has been set.
func (o *EraseDeviceCommand) HasPreserveDataPlan() bool {
	if o != nil && !IsNil(o.PreserveDataPlan) {
		return true
	}

	return false
}

// SetPreserveDataPlan gets a reference to the given bool and assigns it to the PreserveDataPlan field.
func (o *EraseDeviceCommand) SetPreserveDataPlan(v bool) {
	o.PreserveDataPlan = &v
}

// GetDisallowProximitySetup returns the DisallowProximitySetup field value if set, zero value otherwise.
func (o *EraseDeviceCommand) GetDisallowProximitySetup() bool {
	if o == nil || IsNil(o.DisallowProximitySetup) {
		var ret bool
		return ret
	}
	return *o.DisallowProximitySetup
}

// GetDisallowProximitySetupOk returns a tuple with the DisallowProximitySetup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EraseDeviceCommand) GetDisallowProximitySetupOk() (*bool, bool) {
	if o == nil || IsNil(o.DisallowProximitySetup) {
		return nil, false
	}
	return o.DisallowProximitySetup, true
}

// HasDisallowProximitySetup returns a boolean if a field has been set.
func (o *EraseDeviceCommand) HasDisallowProximitySetup() bool {
	if o != nil && !IsNil(o.DisallowProximitySetup) {
		return true
	}

	return false
}

// SetDisallowProximitySetup gets a reference to the given bool and assigns it to the DisallowProximitySetup field.
func (o *EraseDeviceCommand) SetDisallowProximitySetup(v bool) {
	o.DisallowProximitySetup = &v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *EraseDeviceCommand) GetPin() string {
	if o == nil || IsNil(o.Pin) {
		var ret string
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EraseDeviceCommand) GetPinOk() (*string, bool) {
	if o == nil || IsNil(o.Pin) {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *EraseDeviceCommand) HasPin() bool {
	if o != nil && !IsNil(o.Pin) {
		return true
	}

	return false
}

// SetPin gets a reference to the given string and assigns it to the Pin field.
func (o *EraseDeviceCommand) SetPin(v string) {
	o.Pin = &v
}

// GetObliterationBehavior returns the ObliterationBehavior field value if set, zero value otherwise.
func (o *EraseDeviceCommand) GetObliterationBehavior() string {
	if o == nil || IsNil(o.ObliterationBehavior) {
		var ret string
		return ret
	}
	return *o.ObliterationBehavior
}

// GetObliterationBehaviorOk returns a tuple with the ObliterationBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EraseDeviceCommand) GetObliterationBehaviorOk() (*string, bool) {
	if o == nil || IsNil(o.ObliterationBehavior) {
		return nil, false
	}
	return o.ObliterationBehavior, true
}

// HasObliterationBehavior returns a boolean if a field has been set.
func (o *EraseDeviceCommand) HasObliterationBehavior() bool {
	if o != nil && !IsNil(o.ObliterationBehavior) {
		return true
	}

	return false
}

// SetObliterationBehavior gets a reference to the given string and assigns it to the ObliterationBehavior field.
func (o *EraseDeviceCommand) SetObliterationBehavior(v string) {
	o.ObliterationBehavior = &v
}

// GetReturnToService returns the ReturnToService field value if set, zero value otherwise.
func (o *EraseDeviceCommand) GetReturnToService() EraseDeviceCommandAllOfReturnToService {
	if o == nil || IsNil(o.ReturnToService) {
		var ret EraseDeviceCommandAllOfReturnToService
		return ret
	}
	return *o.ReturnToService
}

// GetReturnToServiceOk returns a tuple with the ReturnToService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EraseDeviceCommand) GetReturnToServiceOk() (*EraseDeviceCommandAllOfReturnToService, bool) {
	if o == nil || IsNil(o.ReturnToService) {
		return nil, false
	}
	return o.ReturnToService, true
}

// HasReturnToService returns a boolean if a field has been set.
func (o *EraseDeviceCommand) HasReturnToService() bool {
	if o != nil && !IsNil(o.ReturnToService) {
		return true
	}

	return false
}

// SetReturnToService gets a reference to the given EraseDeviceCommandAllOfReturnToService and assigns it to the ReturnToService field.
func (o *EraseDeviceCommand) SetReturnToService(v EraseDeviceCommandAllOfReturnToService) {
	o.ReturnToService = &v
}

func (o EraseDeviceCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EraseDeviceCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["commandType"] = o.CommandType
	if !IsNil(o.PreserveDataPlan) {
		toSerialize["preserveDataPlan"] = o.PreserveDataPlan
	}
	if !IsNil(o.DisallowProximitySetup) {
		toSerialize["disallowProximitySetup"] = o.DisallowProximitySetup
	}
	if !IsNil(o.Pin) {
		toSerialize["pin"] = o.Pin
	}
	if !IsNil(o.ObliterationBehavior) {
		toSerialize["obliterationBehavior"] = o.ObliterationBehavior
	}
	if !IsNil(o.ReturnToService) {
		toSerialize["returnToService"] = o.ReturnToService
	}
	return toSerialize, nil
}

func (o *EraseDeviceCommand) UnmarshalJSON(data []byte) (err error) {
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

	varEraseDeviceCommand := _EraseDeviceCommand{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEraseDeviceCommand)

	if err != nil {
		return err
	}

	*o = EraseDeviceCommand(varEraseDeviceCommand)

	return err
}

type NullableEraseDeviceCommand struct {
	value *EraseDeviceCommand
	isSet bool
}

func (v NullableEraseDeviceCommand) Get() *EraseDeviceCommand {
	return v.value
}

func (v *NullableEraseDeviceCommand) Set(val *EraseDeviceCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableEraseDeviceCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableEraseDeviceCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEraseDeviceCommand(val *EraseDeviceCommand) *NullableEraseDeviceCommand {
	return &NullableEraseDeviceCommand{value: val, isSet: true}
}

func (v NullableEraseDeviceCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEraseDeviceCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


