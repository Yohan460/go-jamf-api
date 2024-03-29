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

// MdmCommandRequestCommandData - struct for MdmCommandRequestCommandData
type MdmCommandRequestCommandData struct {
	DeclarativeManagementCommand *DeclarativeManagementCommand
	DeleteUserCommand *DeleteUserCommand
	EnableLostModeCommand *EnableLostModeCommand
	EraseDeviceCommand *EraseDeviceCommand
	RestartDeviceCommand *RestartDeviceCommand
	SetAutoAdminPasswordCommand *SetAutoAdminPasswordCommand
	SetRecoveryLockCommand *SetRecoveryLockCommand
	SettingsCommand *SettingsCommand
	Generic *map[string]interface{}
}

// DeclarativeManagementCommandAsMdmCommandRequestCommandData is a convenience function that returns DeclarativeManagementCommand wrapped in MdmCommandRequestCommandData
func DeclarativeManagementCommandAsMdmCommandRequestCommandData(v *DeclarativeManagementCommand) MdmCommandRequestCommandData {
	return MdmCommandRequestCommandData{
		DeclarativeManagementCommand: v,
	}
}

// DeleteUserCommandAsMdmCommandRequestCommandData is a convenience function that returns DeleteUserCommand wrapped in MdmCommandRequestCommandData
func DeleteUserCommandAsMdmCommandRequestCommandData(v *DeleteUserCommand) MdmCommandRequestCommandData {
	return MdmCommandRequestCommandData{
		DeleteUserCommand: v,
	}
}

// EnableLostModeCommandAsMdmCommandRequestCommandData is a convenience function that returns EnableLostModeCommand wrapped in MdmCommandRequestCommandData
func EnableLostModeCommandAsMdmCommandRequestCommandData(v *EnableLostModeCommand) MdmCommandRequestCommandData {
	return MdmCommandRequestCommandData{
		EnableLostModeCommand: v,
	}
}

// EraseDeviceCommandAsMdmCommandRequestCommandData is a convenience function that returns EraseDeviceCommand wrapped in MdmCommandRequestCommandData
func EraseDeviceCommandAsMdmCommandRequestCommandData(v *EraseDeviceCommand) MdmCommandRequestCommandData {
	return MdmCommandRequestCommandData{
		EraseDeviceCommand: v,
	}
}

// RestartDeviceCommandAsMdmCommandRequestCommandData is a convenience function that returns RestartDeviceCommand wrapped in MdmCommandRequestCommandData
func RestartDeviceCommandAsMdmCommandRequestCommandData(v *RestartDeviceCommand) MdmCommandRequestCommandData {
	return MdmCommandRequestCommandData{
		RestartDeviceCommand: v,
	}
}

// SetAutoAdminPasswordCommandAsMdmCommandRequestCommandData is a convenience function that returns SetAutoAdminPasswordCommand wrapped in MdmCommandRequestCommandData
func SetAutoAdminPasswordCommandAsMdmCommandRequestCommandData(v *SetAutoAdminPasswordCommand) MdmCommandRequestCommandData {
	return MdmCommandRequestCommandData{
		SetAutoAdminPasswordCommand: v,
	}
}

// SetRecoveryLockCommandAsMdmCommandRequestCommandData is a convenience function that returns SetRecoveryLockCommand wrapped in MdmCommandRequestCommandData
func SetRecoveryLockCommandAsMdmCommandRequestCommandData(v *SetRecoveryLockCommand) MdmCommandRequestCommandData {
	return MdmCommandRequestCommandData{
		SetRecoveryLockCommand: v,
	}
}

// SettingsCommandAsMdmCommandRequestCommandData is a convenience function that returns SettingsCommand wrapped in MdmCommandRequestCommandData
func SettingsCommandAsMdmCommandRequestCommandData(v *SettingsCommand) MdmCommandRequestCommandData {
	return MdmCommandRequestCommandData{
		SettingsCommand: v,
	}
}

// map[string]interface{}AsMdmCommandRequestCommandData is a convenience function that returns map[string]interface{} wrapped in MdmCommandRequestCommandData
func GenericAsMdmCommandRequestCommandData(v *map[string]interface{}) MdmCommandRequestCommandData {
	return MdmCommandRequestCommandData{
		Generic: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MdmCommandRequestCommandData) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DeclarativeManagementCommand
	err = newStrictDecoder(data).Decode(&dst.DeclarativeManagementCommand)
	if err == nil {
		jsonDeclarativeManagementCommand, _ := json.Marshal(dst.DeclarativeManagementCommand)
		if string(jsonDeclarativeManagementCommand) == "{}" { // empty struct
			dst.DeclarativeManagementCommand = nil
		} else {
			match++
		}
	} else {
		dst.DeclarativeManagementCommand = nil
	}

	// try to unmarshal data into DeleteUserCommand
	err = newStrictDecoder(data).Decode(&dst.DeleteUserCommand)
	if err == nil {
		jsonDeleteUserCommand, _ := json.Marshal(dst.DeleteUserCommand)
		if string(jsonDeleteUserCommand) == "{}" { // empty struct
			dst.DeleteUserCommand = nil
		} else {
			match++
		}
	} else {
		dst.DeleteUserCommand = nil
	}

	// try to unmarshal data into EnableLostModeCommand
	err = newStrictDecoder(data).Decode(&dst.EnableLostModeCommand)
	if err == nil {
		jsonEnableLostModeCommand, _ := json.Marshal(dst.EnableLostModeCommand)
		if string(jsonEnableLostModeCommand) == "{}" { // empty struct
			dst.EnableLostModeCommand = nil
		} else {
			match++
		}
	} else {
		dst.EnableLostModeCommand = nil
	}

	// try to unmarshal data into EraseDeviceCommand
	err = newStrictDecoder(data).Decode(&dst.EraseDeviceCommand)
	if err == nil {
		jsonEraseDeviceCommand, _ := json.Marshal(dst.EraseDeviceCommand)
		if string(jsonEraseDeviceCommand) == "{}" { // empty struct
			dst.EraseDeviceCommand = nil
		} else {
			match++
		}
	} else {
		dst.EraseDeviceCommand = nil
	}

	// try to unmarshal data into RestartDeviceCommand
	err = newStrictDecoder(data).Decode(&dst.RestartDeviceCommand)
	if err == nil {
		jsonRestartDeviceCommand, _ := json.Marshal(dst.RestartDeviceCommand)
		if string(jsonRestartDeviceCommand) == "{}" { // empty struct
			dst.RestartDeviceCommand = nil
		} else {
			match++
		}
	} else {
		dst.RestartDeviceCommand = nil
	}

	// try to unmarshal data into SetAutoAdminPasswordCommand
	err = newStrictDecoder(data).Decode(&dst.SetAutoAdminPasswordCommand)
	if err == nil {
		jsonSetAutoAdminPasswordCommand, _ := json.Marshal(dst.SetAutoAdminPasswordCommand)
		if string(jsonSetAutoAdminPasswordCommand) == "{}" { // empty struct
			dst.SetAutoAdminPasswordCommand = nil
		} else {
			match++
		}
	} else {
		dst.SetAutoAdminPasswordCommand = nil
	}

	// try to unmarshal data into SetRecoveryLockCommand
	err = newStrictDecoder(data).Decode(&dst.SetRecoveryLockCommand)
	if err == nil {
		jsonSetRecoveryLockCommand, _ := json.Marshal(dst.SetRecoveryLockCommand)
		if string(jsonSetRecoveryLockCommand) == "{}" { // empty struct
			dst.SetRecoveryLockCommand = nil
		} else {
			match++
		}
	} else {
		dst.SetRecoveryLockCommand = nil
	}

	// try to unmarshal data into SettingsCommand
	err = newStrictDecoder(data).Decode(&dst.SettingsCommand)
	if err == nil {
		jsonSettingsCommand, _ := json.Marshal(dst.SettingsCommand)
		if string(jsonSettingsCommand) == "{}" { // empty struct
			dst.SettingsCommand = nil
		} else {
			match++
		}
	} else {
		dst.SettingsCommand = nil
	}

	// try to unmarshal data into Generic
	err = newStrictDecoder(data).Decode(&dst.Generic)
	if err == nil {
		jsonGeneric, _ := json.Marshal(dst.Generic)
		if string(jsonGeneric) == "{}" { // empty struct
			dst.Generic = nil
		} else {
			match++
		}
	} else {
		dst.Generic = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DeclarativeManagementCommand = nil
		dst.DeleteUserCommand = nil
		dst.EnableLostModeCommand = nil
		dst.EraseDeviceCommand = nil
		dst.RestartDeviceCommand = nil
		dst.SetAutoAdminPasswordCommand = nil
		dst.SetRecoveryLockCommand = nil
		dst.SettingsCommand = nil
		dst.Generic = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MdmCommandRequestCommandData)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MdmCommandRequestCommandData)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MdmCommandRequestCommandData) MarshalJSON() ([]byte, error) {
	if src.DeclarativeManagementCommand != nil {
		return json.Marshal(&src.DeclarativeManagementCommand)
	}

	if src.DeleteUserCommand != nil {
		return json.Marshal(&src.DeleteUserCommand)
	}

	if src.EnableLostModeCommand != nil {
		return json.Marshal(&src.EnableLostModeCommand)
	}

	if src.EraseDeviceCommand != nil {
		return json.Marshal(&src.EraseDeviceCommand)
	}

	if src.RestartDeviceCommand != nil {
		return json.Marshal(&src.RestartDeviceCommand)
	}

	if src.SetAutoAdminPasswordCommand != nil {
		return json.Marshal(&src.SetAutoAdminPasswordCommand)
	}

	if src.SetRecoveryLockCommand != nil {
		return json.Marshal(&src.SetRecoveryLockCommand)
	}

	if src.SettingsCommand != nil {
		return json.Marshal(&src.SettingsCommand)
	}

	if src.Generic != nil {
		return json.Marshal(&src.Generic)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MdmCommandRequestCommandData) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DeclarativeManagementCommand != nil {
		return obj.DeclarativeManagementCommand
	}

	if obj.DeleteUserCommand != nil {
		return obj.DeleteUserCommand
	}

	if obj.EnableLostModeCommand != nil {
		return obj.EnableLostModeCommand
	}

	if obj.EraseDeviceCommand != nil {
		return obj.EraseDeviceCommand
	}

	if obj.RestartDeviceCommand != nil {
		return obj.RestartDeviceCommand
	}

	if obj.SetAutoAdminPasswordCommand != nil {
		return obj.SetAutoAdminPasswordCommand
	}

	if obj.SetRecoveryLockCommand != nil {
		return obj.SetRecoveryLockCommand
	}

	if obj.SettingsCommand != nil {
		return obj.SettingsCommand
	}

	if obj.Generic != nil {
		return obj.Generic
	}

	// all schemas are nil
	return nil
}

type NullableMdmCommandRequestCommandData struct {
	value *MdmCommandRequestCommandData
	isSet bool
}

func (v NullableMdmCommandRequestCommandData) Get() *MdmCommandRequestCommandData {
	return v.value
}

func (v *NullableMdmCommandRequestCommandData) Set(val *MdmCommandRequestCommandData) {
	v.value = val
	v.isSet = true
}

func (v NullableMdmCommandRequestCommandData) IsSet() bool {
	return v.isSet
}

func (v *NullableMdmCommandRequestCommandData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMdmCommandRequestCommandData(val *MdmCommandRequestCommandData) *NullableMdmCommandRequestCommandData {
	return &NullableMdmCommandRequestCommandData{value: val, isSet: true}
}

func (v NullableMdmCommandRequestCommandData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMdmCommandRequestCommandData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


