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

// checks if the AccountPreferencesV4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountPreferencesV4{}

// AccountPreferencesV4 struct for AccountPreferencesV4
type AccountPreferencesV4 struct {
	// Language codes supported by Jamf Pro
	Language string `json:"language"`
	DateFormat string `json:"dateFormat"`
	Timezone string `json:"timezone"`
	DisableRelativeDates bool `json:"disableRelativeDates"`
	DisablePageLeaveCheck bool `json:"disablePageLeaveCheck"`
	DisableTablePagination bool `json:"disableTablePagination"`
	DisableShortcutsTooltips bool `json:"disableShortcutsTooltips"`
	ConfigProfilesSortingMethod string `json:"configProfilesSortingMethod"`
}

type _AccountPreferencesV4 AccountPreferencesV4

// NewAccountPreferencesV4 instantiates a new AccountPreferencesV4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountPreferencesV4(language string, dateFormat string, timezone string, disableRelativeDates bool, disablePageLeaveCheck bool, disableTablePagination bool, disableShortcutsTooltips bool, configProfilesSortingMethod string) *AccountPreferencesV4 {
	this := AccountPreferencesV4{}
	this.Language = language
	this.DateFormat = dateFormat
	this.Timezone = timezone
	this.DisableRelativeDates = disableRelativeDates
	this.DisablePageLeaveCheck = disablePageLeaveCheck
	this.DisableTablePagination = disableTablePagination
	this.DisableShortcutsTooltips = disableShortcutsTooltips
	this.ConfigProfilesSortingMethod = configProfilesSortingMethod
	return &this
}

// NewAccountPreferencesV4WithDefaults instantiates a new AccountPreferencesV4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountPreferencesV4WithDefaults() *AccountPreferencesV4 {
	this := AccountPreferencesV4{}
	var language string = "en"
	this.Language = language
	return &this
}

// GetLanguage returns the Language field value
func (o *AccountPreferencesV4) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV4) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *AccountPreferencesV4) SetLanguage(v string) {
	o.Language = v
}

// GetDateFormat returns the DateFormat field value
func (o *AccountPreferencesV4) GetDateFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateFormat
}

// GetDateFormatOk returns a tuple with the DateFormat field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV4) GetDateFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateFormat, true
}

// SetDateFormat sets field value
func (o *AccountPreferencesV4) SetDateFormat(v string) {
	o.DateFormat = v
}

// GetTimezone returns the Timezone field value
func (o *AccountPreferencesV4) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV4) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *AccountPreferencesV4) SetTimezone(v string) {
	o.Timezone = v
}

// GetDisableRelativeDates returns the DisableRelativeDates field value
func (o *AccountPreferencesV4) GetDisableRelativeDates() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableRelativeDates
}

// GetDisableRelativeDatesOk returns a tuple with the DisableRelativeDates field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV4) GetDisableRelativeDatesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableRelativeDates, true
}

// SetDisableRelativeDates sets field value
func (o *AccountPreferencesV4) SetDisableRelativeDates(v bool) {
	o.DisableRelativeDates = v
}

// GetDisablePageLeaveCheck returns the DisablePageLeaveCheck field value
func (o *AccountPreferencesV4) GetDisablePageLeaveCheck() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisablePageLeaveCheck
}

// GetDisablePageLeaveCheckOk returns a tuple with the DisablePageLeaveCheck field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV4) GetDisablePageLeaveCheckOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisablePageLeaveCheck, true
}

// SetDisablePageLeaveCheck sets field value
func (o *AccountPreferencesV4) SetDisablePageLeaveCheck(v bool) {
	o.DisablePageLeaveCheck = v
}

// GetDisableTablePagination returns the DisableTablePagination field value
func (o *AccountPreferencesV4) GetDisableTablePagination() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableTablePagination
}

// GetDisableTablePaginationOk returns a tuple with the DisableTablePagination field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV4) GetDisableTablePaginationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableTablePagination, true
}

// SetDisableTablePagination sets field value
func (o *AccountPreferencesV4) SetDisableTablePagination(v bool) {
	o.DisableTablePagination = v
}

// GetDisableShortcutsTooltips returns the DisableShortcutsTooltips field value
func (o *AccountPreferencesV4) GetDisableShortcutsTooltips() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableShortcutsTooltips
}

// GetDisableShortcutsTooltipsOk returns a tuple with the DisableShortcutsTooltips field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV4) GetDisableShortcutsTooltipsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableShortcutsTooltips, true
}

// SetDisableShortcutsTooltips sets field value
func (o *AccountPreferencesV4) SetDisableShortcutsTooltips(v bool) {
	o.DisableShortcutsTooltips = v
}

// GetConfigProfilesSortingMethod returns the ConfigProfilesSortingMethod field value
func (o *AccountPreferencesV4) GetConfigProfilesSortingMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigProfilesSortingMethod
}

// GetConfigProfilesSortingMethodOk returns a tuple with the ConfigProfilesSortingMethod field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV4) GetConfigProfilesSortingMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigProfilesSortingMethod, true
}

// SetConfigProfilesSortingMethod sets field value
func (o *AccountPreferencesV4) SetConfigProfilesSortingMethod(v string) {
	o.ConfigProfilesSortingMethod = v
}

func (o AccountPreferencesV4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountPreferencesV4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["language"] = o.Language
	toSerialize["dateFormat"] = o.DateFormat
	toSerialize["timezone"] = o.Timezone
	toSerialize["disableRelativeDates"] = o.DisableRelativeDates
	toSerialize["disablePageLeaveCheck"] = o.DisablePageLeaveCheck
	toSerialize["disableTablePagination"] = o.DisableTablePagination
	toSerialize["disableShortcutsTooltips"] = o.DisableShortcutsTooltips
	toSerialize["configProfilesSortingMethod"] = o.ConfigProfilesSortingMethod
	return toSerialize, nil
}

func (o *AccountPreferencesV4) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"language",
		"dateFormat",
		"timezone",
		"disableRelativeDates",
		"disablePageLeaveCheck",
		"disableTablePagination",
		"disableShortcutsTooltips",
		"configProfilesSortingMethod",
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

	varAccountPreferencesV4 := _AccountPreferencesV4{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountPreferencesV4)

	if err != nil {
		return err
	}

	*o = AccountPreferencesV4(varAccountPreferencesV4)

	return err
}

type NullableAccountPreferencesV4 struct {
	value *AccountPreferencesV4
	isSet bool
}

func (v NullableAccountPreferencesV4) Get() *AccountPreferencesV4 {
	return v.value
}

func (v *NullableAccountPreferencesV4) Set(val *AccountPreferencesV4) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountPreferencesV4) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountPreferencesV4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountPreferencesV4(val *AccountPreferencesV4) *NullableAccountPreferencesV4 {
	return &NullableAccountPreferencesV4{value: val, isSet: true}
}

func (v NullableAccountPreferencesV4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountPreferencesV4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


