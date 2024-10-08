/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// MdmCommandType the model 'MdmCommandType'
type MdmCommandType string

// List of MdmCommandType
const (
	MDMCOMMANDTYPE_CLEAR_RESTRICTIONS_PASSWORD MdmCommandType = "CLEAR_RESTRICTIONS_PASSWORD"
	MDMCOMMANDTYPE_DECLARATIVE_MANAGEMENT MdmCommandType = "DECLARATIVE_MANAGEMENT"
	MDMCOMMANDTYPE_DELETE_USER MdmCommandType = "DELETE_USER"
	MDMCOMMANDTYPE_DEVICE_LOCK MdmCommandType = "DEVICE_LOCK"
	MDMCOMMANDTYPE_ENABLE_LOST_MODE MdmCommandType = "ENABLE_LOST_MODE"
	MDMCOMMANDTYPE_ERASE_DEVICE MdmCommandType = "ERASE_DEVICE"
	MDMCOMMANDTYPE_LOG_OUT_USER MdmCommandType = "LOG_OUT_USER"
	MDMCOMMANDTYPE_RESTART_DEVICE MdmCommandType = "RESTART_DEVICE"
	MDMCOMMANDTYPE_SETTINGS MdmCommandType = "SETTINGS"
	MDMCOMMANDTYPE_SET_RECOVERY_LOCK MdmCommandType = "SET_RECOVERY_LOCK"
	MDMCOMMANDTYPE_SET_AUTO_ADMIN_PASSWORD MdmCommandType = "SET_AUTO_ADMIN_PASSWORD"
	MDMCOMMANDTYPE_SHUT_DOWN_DEVICE MdmCommandType = "SHUT_DOWN_DEVICE"
	MDMCOMMANDTYPE_DEVICE_INFORMATION MdmCommandType = "DEVICE_INFORMATION"
)

// All allowed values of MdmCommandType enum
var AllowedMdmCommandTypeEnumValues = []MdmCommandType{
	"CLEAR_RESTRICTIONS_PASSWORD",
	"DECLARATIVE_MANAGEMENT",
	"DELETE_USER",
	"DEVICE_LOCK",
	"ENABLE_LOST_MODE",
	"ERASE_DEVICE",
	"LOG_OUT_USER",
	"RESTART_DEVICE",
	"SETTINGS",
	"SET_RECOVERY_LOCK",
	"SET_AUTO_ADMIN_PASSWORD",
	"SHUT_DOWN_DEVICE",
	"DEVICE_INFORMATION",
}

func (v *MdmCommandType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MdmCommandType(value)
	for _, existing := range AllowedMdmCommandTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MdmCommandType", value)
}

// NewMdmCommandTypeFromValue returns a pointer to a valid MdmCommandType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMdmCommandTypeFromValue(v string) (*MdmCommandType, error) {
	ev := MdmCommandType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MdmCommandType: valid values are %v", v, AllowedMdmCommandTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MdmCommandType) IsValid() bool {
	for _, existing := range AllowedMdmCommandTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MdmCommandType value
func (v MdmCommandType) Ptr() *MdmCommandType {
	return &v
}

type NullableMdmCommandType struct {
	value *MdmCommandType
	isSet bool
}

func (v NullableMdmCommandType) Get() *MdmCommandType {
	return v.value
}

func (v *NullableMdmCommandType) Set(val *MdmCommandType) {
	v.value = val
	v.isSet = true
}

func (v NullableMdmCommandType) IsSet() bool {
	return v.isSet
}

func (v *NullableMdmCommandType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMdmCommandType(val *MdmCommandType) *NullableMdmCommandType {
	return &NullableMdmCommandType{value: val, isSet: true}
}

func (v NullableMdmCommandType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMdmCommandType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

