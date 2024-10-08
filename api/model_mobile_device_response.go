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

// MobileDeviceResponse - struct for MobileDeviceResponse
type MobileDeviceResponse struct {
	MobileDeviceIosInventory *MobileDeviceIosInventory
	MobileDeviceTvOsInventory *MobileDeviceTvOsInventory
	MobileDeviceWatchOsInventory *MobileDeviceWatchOsInventory
}

// MobileDeviceIosInventoryAsMobileDeviceResponse is a convenience function that returns MobileDeviceIosInventory wrapped in MobileDeviceResponse
func MobileDeviceIosInventoryAsMobileDeviceResponse(v *MobileDeviceIosInventory) MobileDeviceResponse {
	return MobileDeviceResponse{
		MobileDeviceIosInventory: v,
	}
}

// MobileDeviceTvOsInventoryAsMobileDeviceResponse is a convenience function that returns MobileDeviceTvOsInventory wrapped in MobileDeviceResponse
func MobileDeviceTvOsInventoryAsMobileDeviceResponse(v *MobileDeviceTvOsInventory) MobileDeviceResponse {
	return MobileDeviceResponse{
		MobileDeviceTvOsInventory: v,
	}
}

// MobileDeviceWatchOsInventoryAsMobileDeviceResponse is a convenience function that returns MobileDeviceWatchOsInventory wrapped in MobileDeviceResponse
func MobileDeviceWatchOsInventoryAsMobileDeviceResponse(v *MobileDeviceWatchOsInventory) MobileDeviceResponse {
	return MobileDeviceResponse{
		MobileDeviceWatchOsInventory: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MobileDeviceResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MobileDeviceIosInventory
	err = newStrictDecoder(data).Decode(&dst.MobileDeviceIosInventory)
	if err == nil {
		jsonMobileDeviceIosInventory, _ := json.Marshal(dst.MobileDeviceIosInventory)
		if string(jsonMobileDeviceIosInventory) == "{}" { // empty struct
			dst.MobileDeviceIosInventory = nil
		} else {
			match++
		}
	} else {
		dst.MobileDeviceIosInventory = nil
	}

	// try to unmarshal data into MobileDeviceTvOsInventory
	err = newStrictDecoder(data).Decode(&dst.MobileDeviceTvOsInventory)
	if err == nil {
		jsonMobileDeviceTvOsInventory, _ := json.Marshal(dst.MobileDeviceTvOsInventory)
		if string(jsonMobileDeviceTvOsInventory) == "{}" { // empty struct
			dst.MobileDeviceTvOsInventory = nil
		} else {
			match++
		}
	} else {
		dst.MobileDeviceTvOsInventory = nil
	}

	// try to unmarshal data into MobileDeviceWatchOsInventory
	err = newStrictDecoder(data).Decode(&dst.MobileDeviceWatchOsInventory)
	if err == nil {
		jsonMobileDeviceWatchOsInventory, _ := json.Marshal(dst.MobileDeviceWatchOsInventory)
		if string(jsonMobileDeviceWatchOsInventory) == "{}" { // empty struct
			dst.MobileDeviceWatchOsInventory = nil
		} else {
			match++
		}
	} else {
		dst.MobileDeviceWatchOsInventory = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MobileDeviceIosInventory = nil
		dst.MobileDeviceTvOsInventory = nil
		dst.MobileDeviceWatchOsInventory = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MobileDeviceResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MobileDeviceResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MobileDeviceResponse) MarshalJSON() ([]byte, error) {
	if src.MobileDeviceIosInventory != nil {
		return json.Marshal(&src.MobileDeviceIosInventory)
	}

	if src.MobileDeviceTvOsInventory != nil {
		return json.Marshal(&src.MobileDeviceTvOsInventory)
	}

	if src.MobileDeviceWatchOsInventory != nil {
		return json.Marshal(&src.MobileDeviceWatchOsInventory)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MobileDeviceResponse) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MobileDeviceIosInventory != nil {
		return obj.MobileDeviceIosInventory
	}

	if obj.MobileDeviceTvOsInventory != nil {
		return obj.MobileDeviceTvOsInventory
	}

	if obj.MobileDeviceWatchOsInventory != nil {
		return obj.MobileDeviceWatchOsInventory
	}

	// all schemas are nil
	return nil
}

type NullableMobileDeviceResponse struct {
	value *MobileDeviceResponse
	isSet bool
}

func (v NullableMobileDeviceResponse) Get() *MobileDeviceResponse {
	return v.value
}

func (v *NullableMobileDeviceResponse) Set(val *MobileDeviceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileDeviceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileDeviceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileDeviceResponse(val *MobileDeviceResponse) *NullableMobileDeviceResponse {
	return &NullableMobileDeviceResponse{value: val, isSet: true}
}

func (v NullableMobileDeviceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileDeviceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


