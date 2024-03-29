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

// checks if the ClientCheckInV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientCheckInV2{}

// ClientCheckInV2 struct for ClientCheckInV2
type ClientCheckInV2 struct {
	// Suggested values are 5, 15, 30, or 60. Web interface will not display correctly if not one of those. Minimim is 5, maximum is 60.
	CheckInFrequency *int32 `json:"checkInFrequency,omitempty"`
	CreateHooks *bool `json:"createHooks,omitempty"`
	HookLog *bool `json:"hookLog,omitempty"`
	HookPolicies *bool `json:"hookPolicies,omitempty"`
	HookHideRestore *bool `json:"hookHideRestore,omitempty"`
	HookMcx *bool `json:"hookMcx,omitempty"`
	BackgroundHooks *bool `json:"backgroundHooks,omitempty"`
	HookDisplayStatus *bool `json:"hookDisplayStatus,omitempty"`
	CreateStartupScript *bool `json:"createStartupScript,omitempty"`
	StartupLog *bool `json:"startupLog,omitempty"`
	StartupPolicies *bool `json:"startupPolicies,omitempty"`
	StartupSsh *bool `json:"startupSsh,omitempty"`
	StartupMcx *bool `json:"startupMcx,omitempty"`
	EnableLocalConfigurationProfiles *bool `json:"enableLocalConfigurationProfiles,omitempty"`
}

// NewClientCheckInV2 instantiates a new ClientCheckInV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientCheckInV2() *ClientCheckInV2 {
	this := ClientCheckInV2{}
	var checkInFrequency int32 = 15
	this.CheckInFrequency = &checkInFrequency
	var createHooks bool = false
	this.CreateHooks = &createHooks
	var hookLog bool = false
	this.HookLog = &hookLog
	var hookPolicies bool = false
	this.HookPolicies = &hookPolicies
	var hookHideRestore bool = false
	this.HookHideRestore = &hookHideRestore
	var hookMcx bool = false
	this.HookMcx = &hookMcx
	var backgroundHooks bool = false
	this.BackgroundHooks = &backgroundHooks
	var hookDisplayStatus bool = false
	this.HookDisplayStatus = &hookDisplayStatus
	var createStartupScript bool = false
	this.CreateStartupScript = &createStartupScript
	var startupLog bool = false
	this.StartupLog = &startupLog
	var startupPolicies bool = false
	this.StartupPolicies = &startupPolicies
	var startupSsh bool = false
	this.StartupSsh = &startupSsh
	var startupMcx bool = false
	this.StartupMcx = &startupMcx
	var enableLocalConfigurationProfiles bool = false
	this.EnableLocalConfigurationProfiles = &enableLocalConfigurationProfiles
	return &this
}

// NewClientCheckInV2WithDefaults instantiates a new ClientCheckInV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientCheckInV2WithDefaults() *ClientCheckInV2 {
	this := ClientCheckInV2{}
	var checkInFrequency int32 = 15
	this.CheckInFrequency = &checkInFrequency
	var createHooks bool = false
	this.CreateHooks = &createHooks
	var hookLog bool = false
	this.HookLog = &hookLog
	var hookPolicies bool = false
	this.HookPolicies = &hookPolicies
	var hookHideRestore bool = false
	this.HookHideRestore = &hookHideRestore
	var hookMcx bool = false
	this.HookMcx = &hookMcx
	var backgroundHooks bool = false
	this.BackgroundHooks = &backgroundHooks
	var hookDisplayStatus bool = false
	this.HookDisplayStatus = &hookDisplayStatus
	var createStartupScript bool = false
	this.CreateStartupScript = &createStartupScript
	var startupLog bool = false
	this.StartupLog = &startupLog
	var startupPolicies bool = false
	this.StartupPolicies = &startupPolicies
	var startupSsh bool = false
	this.StartupSsh = &startupSsh
	var startupMcx bool = false
	this.StartupMcx = &startupMcx
	var enableLocalConfigurationProfiles bool = false
	this.EnableLocalConfigurationProfiles = &enableLocalConfigurationProfiles
	return &this
}

// GetCheckInFrequency returns the CheckInFrequency field value if set, zero value otherwise.
func (o *ClientCheckInV2) GetCheckInFrequency() int32 {
	if o == nil || IsNil(o.CheckInFrequency) {
		var ret int32
		return ret
	}
	return *o.CheckInFrequency
}

// GetCheckInFrequencyOk returns a tuple with the CheckInFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCheckInV2) GetCheckInFrequencyOk() (*int32, bool) {
	if o == nil || IsNil(o.CheckInFrequency) {
		return nil, false
	}
	return o.CheckInFrequency, true
}

// HasCheckInFrequency returns a boolean if a field has been set.
func (o *ClientCheckInV2) HasCheckInFrequency() bool {
	if o != nil && !IsNil(o.CheckInFrequency) {
		return true
	}

	return false
}

// SetCheckInFrequency gets a reference to the given int32 and assigns it to the CheckInFrequency field.
func (o *ClientCheckInV2) SetCheckInFrequency(v int32) {
	o.CheckInFrequency = &v
}

// GetCreateHooks returns the CreateHooks field value if set, zero value otherwise.
func (o *ClientCheckInV2) GetCreateHooks() bool {
	if o == nil || IsNil(o.CreateHooks) {
		var ret bool
		return ret
	}
	return *o.CreateHooks
}

// GetCreateHooksOk returns a tuple with the CreateHooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCheckInV2) GetCreateHooksOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateHooks) {
		return nil, false
	}
	return o.CreateHooks, true
}

// HasCreateHooks returns a boolean if a field has been set.
func (o *ClientCheckInV2) HasCreateHooks() bool {
	if o != nil && !IsNil(o.CreateHooks) {
		return true
	}

	return false
}

// SetCreateHooks gets a reference to the given bool and assigns it to the CreateHooks field.
func (o *ClientCheckInV2) SetCreateHooks(v bool) {
	o.CreateHooks = &v
}

// GetHookLog returns the HookLog field value if set, zero value otherwise.
func (o *ClientCheckInV2) GetHookLog() bool {
	if o == nil || IsNil(o.HookLog) {
		var ret bool
		return ret
	}
	return *o.HookLog
}

// GetHookLogOk returns a tuple with the HookLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCheckInV2) GetHookLogOk() (*bool, bool) {
	if o == nil || IsNil(o.HookLog) {
		return nil, false
	}
	return o.HookLog, true
}

// HasHookLog returns a boolean if a field has been set.
func (o *ClientCheckInV2) HasHookLog() bool {
	if o != nil && !IsNil(o.HookLog) {
		return true
	}

	return false
}

// SetHookLog gets a reference to the given bool and assigns it to the HookLog field.
func (o *ClientCheckInV2) SetHookLog(v bool) {
	o.HookLog = &v
}

// GetHookPolicies returns the HookPolicies field value if set, zero value otherwise.
func (o *ClientCheckInV2) GetHookPolicies() bool {
	if o == nil || IsNil(o.HookPolicies) {
		var ret bool
		return ret
	}
	return *o.HookPolicies
}

// GetHookPoliciesOk returns a tuple with the HookPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCheckInV2) GetHookPoliciesOk() (*bool, bool) {
	if o == nil || IsNil(o.HookPolicies) {
		return nil, false
	}
	return o.HookPolicies, true
}

// HasHookPolicies returns a boolean if a field has been set.
func (o *ClientCheckInV2) HasHookPolicies() bool {
	if o != nil && !IsNil(o.HookPolicies) {
		return true
	}

	return false
}

// SetHookPolicies gets a reference to the given bool and assigns it to the HookPolicies field.
func (o *ClientCheckInV2) SetHookPolicies(v bool) {
	o.HookPolicies = &v
}

// GetHookHideRestore returns the HookHideRestore field value if set, zero value otherwise.
func (o *ClientCheckInV2) GetHookHideRestore() bool {
	if o == nil || IsNil(o.HookHideRestore) {
		var ret bool
		return ret
	}
	return *o.HookHideRestore
}

// GetHookHideRestoreOk returns a tuple with the HookHideRestore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCheckInV2) GetHookHideRestoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HookHideRestore) {
		return nil, false
	}
	return o.HookHideRestore, true
}

// HasHookHideRestore returns a boolean if a field has been set.
func (o *ClientCheckInV2) HasHookHideRestore() bool {
	if o != nil && !IsNil(o.HookHideRestore) {
		return true
	}

	return false
}

// SetHookHideRestore gets a reference to the given bool and assigns it to the HookHideRestore field.
func (o *ClientCheckInV2) SetHookHideRestore(v bool) {
	o.HookHideRestore = &v
}

// GetHookMcx returns the HookMcx field value if set, zero value otherwise.
func (o *ClientCheckInV2) GetHookMcx() bool {
	if o == nil || IsNil(o.HookMcx) {
		var ret bool
		return ret
	}
	return *o.HookMcx
}

// GetHookMcxOk returns a tuple with the HookMcx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCheckInV2) GetHookMcxOk() (*bool, bool) {
	if o == nil || IsNil(o.HookMcx) {
		return nil, false
	}
	return o.HookMcx, true
}

// HasHookMcx returns a boolean if a field has been set.
func (o *ClientCheckInV2) HasHookMcx() bool {
	if o != nil && !IsNil(o.HookMcx) {
		return true
	}

	return false
}

// SetHookMcx gets a reference to the given bool and assigns it to the HookMcx field.
func (o *ClientCheckInV2) SetHookMcx(v bool) {
	o.HookMcx = &v
}

// GetBackgroundHooks returns the BackgroundHooks field value if set, zero value otherwise.
func (o *ClientCheckInV2) GetBackgroundHooks() bool {
	if o == nil || IsNil(o.BackgroundHooks) {
		var ret bool
		return ret
	}
	return *o.BackgroundHooks
}

// GetBackgroundHooksOk returns a tuple with the BackgroundHooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCheckInV2) GetBackgroundHooksOk() (*bool, bool) {
	if o == nil || IsNil(o.BackgroundHooks) {
		return nil, false
	}
	return o.BackgroundHooks, true
}

// HasBackgroundHooks returns a boolean if a field has been set.
func (o *ClientCheckInV2) HasBackgroundHooks() bool {
	if o != nil && !IsNil(o.BackgroundHooks) {
		return true
	}

	return false
}

// SetBackgroundHooks gets a reference to the given bool and assigns it to the BackgroundHooks field.
func (o *ClientCheckInV2) SetBackgroundHooks(v bool) {
	o.BackgroundHooks = &v
}

// GetHookDisplayStatus returns the HookDisplayStatus field value if set, zero value otherwise.
func (o *ClientCheckInV2) GetHookDisplayStatus() bool {
	if o == nil || IsNil(o.HookDisplayStatus) {
		var ret bool
		return ret
	}
	return *o.HookDisplayStatus
}

// GetHookDisplayStatusOk returns a tuple with the HookDisplayStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCheckInV2) GetHookDisplayStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.HookDisplayStatus) {
		return nil, false
	}
	return o.HookDisplayStatus, true
}

// HasHookDisplayStatus returns a boolean if a field has been set.
func (o *ClientCheckInV2) HasHookDisplayStatus() bool {
	if o != nil && !IsNil(o.HookDisplayStatus) {
		return true
	}

	return false
}

// SetHookDisplayStatus gets a reference to the given bool and assigns it to the HookDisplayStatus field.
func (o *ClientCheckInV2) SetHookDisplayStatus(v bool) {
	o.HookDisplayStatus = &v
}

// GetCreateStartupScript returns the CreateStartupScript field value if set, zero value otherwise.
func (o *ClientCheckInV2) GetCreateStartupScript() bool {
	if o == nil || IsNil(o.CreateStartupScript) {
		var ret bool
		return ret
	}
	return *o.CreateStartupScript
}

// GetCreateStartupScriptOk returns a tuple with the CreateStartupScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCheckInV2) GetCreateStartupScriptOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateStartupScript) {
		return nil, false
	}
	return o.CreateStartupScript, true
}

// HasCreateStartupScript returns a boolean if a field has been set.
func (o *ClientCheckInV2) HasCreateStartupScript() bool {
	if o != nil && !IsNil(o.CreateStartupScript) {
		return true
	}

	return false
}

// SetCreateStartupScript gets a reference to the given bool and assigns it to the CreateStartupScript field.
func (o *ClientCheckInV2) SetCreateStartupScript(v bool) {
	o.CreateStartupScript = &v
}

// GetStartupLog returns the StartupLog field value if set, zero value otherwise.
func (o *ClientCheckInV2) GetStartupLog() bool {
	if o == nil || IsNil(o.StartupLog) {
		var ret bool
		return ret
	}
	return *o.StartupLog
}

// GetStartupLogOk returns a tuple with the StartupLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCheckInV2) GetStartupLogOk() (*bool, bool) {
	if o == nil || IsNil(o.StartupLog) {
		return nil, false
	}
	return o.StartupLog, true
}

// HasStartupLog returns a boolean if a field has been set.
func (o *ClientCheckInV2) HasStartupLog() bool {
	if o != nil && !IsNil(o.StartupLog) {
		return true
	}

	return false
}

// SetStartupLog gets a reference to the given bool and assigns it to the StartupLog field.
func (o *ClientCheckInV2) SetStartupLog(v bool) {
	o.StartupLog = &v
}

// GetStartupPolicies returns the StartupPolicies field value if set, zero value otherwise.
func (o *ClientCheckInV2) GetStartupPolicies() bool {
	if o == nil || IsNil(o.StartupPolicies) {
		var ret bool
		return ret
	}
	return *o.StartupPolicies
}

// GetStartupPoliciesOk returns a tuple with the StartupPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCheckInV2) GetStartupPoliciesOk() (*bool, bool) {
	if o == nil || IsNil(o.StartupPolicies) {
		return nil, false
	}
	return o.StartupPolicies, true
}

// HasStartupPolicies returns a boolean if a field has been set.
func (o *ClientCheckInV2) HasStartupPolicies() bool {
	if o != nil && !IsNil(o.StartupPolicies) {
		return true
	}

	return false
}

// SetStartupPolicies gets a reference to the given bool and assigns it to the StartupPolicies field.
func (o *ClientCheckInV2) SetStartupPolicies(v bool) {
	o.StartupPolicies = &v
}

// GetStartupSsh returns the StartupSsh field value if set, zero value otherwise.
func (o *ClientCheckInV2) GetStartupSsh() bool {
	if o == nil || IsNil(o.StartupSsh) {
		var ret bool
		return ret
	}
	return *o.StartupSsh
}

// GetStartupSshOk returns a tuple with the StartupSsh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCheckInV2) GetStartupSshOk() (*bool, bool) {
	if o == nil || IsNil(o.StartupSsh) {
		return nil, false
	}
	return o.StartupSsh, true
}

// HasStartupSsh returns a boolean if a field has been set.
func (o *ClientCheckInV2) HasStartupSsh() bool {
	if o != nil && !IsNil(o.StartupSsh) {
		return true
	}

	return false
}

// SetStartupSsh gets a reference to the given bool and assigns it to the StartupSsh field.
func (o *ClientCheckInV2) SetStartupSsh(v bool) {
	o.StartupSsh = &v
}

// GetStartupMcx returns the StartupMcx field value if set, zero value otherwise.
func (o *ClientCheckInV2) GetStartupMcx() bool {
	if o == nil || IsNil(o.StartupMcx) {
		var ret bool
		return ret
	}
	return *o.StartupMcx
}

// GetStartupMcxOk returns a tuple with the StartupMcx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCheckInV2) GetStartupMcxOk() (*bool, bool) {
	if o == nil || IsNil(o.StartupMcx) {
		return nil, false
	}
	return o.StartupMcx, true
}

// HasStartupMcx returns a boolean if a field has been set.
func (o *ClientCheckInV2) HasStartupMcx() bool {
	if o != nil && !IsNil(o.StartupMcx) {
		return true
	}

	return false
}

// SetStartupMcx gets a reference to the given bool and assigns it to the StartupMcx field.
func (o *ClientCheckInV2) SetStartupMcx(v bool) {
	o.StartupMcx = &v
}

// GetEnableLocalConfigurationProfiles returns the EnableLocalConfigurationProfiles field value if set, zero value otherwise.
func (o *ClientCheckInV2) GetEnableLocalConfigurationProfiles() bool {
	if o == nil || IsNil(o.EnableLocalConfigurationProfiles) {
		var ret bool
		return ret
	}
	return *o.EnableLocalConfigurationProfiles
}

// GetEnableLocalConfigurationProfilesOk returns a tuple with the EnableLocalConfigurationProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCheckInV2) GetEnableLocalConfigurationProfilesOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableLocalConfigurationProfiles) {
		return nil, false
	}
	return o.EnableLocalConfigurationProfiles, true
}

// HasEnableLocalConfigurationProfiles returns a boolean if a field has been set.
func (o *ClientCheckInV2) HasEnableLocalConfigurationProfiles() bool {
	if o != nil && !IsNil(o.EnableLocalConfigurationProfiles) {
		return true
	}

	return false
}

// SetEnableLocalConfigurationProfiles gets a reference to the given bool and assigns it to the EnableLocalConfigurationProfiles field.
func (o *ClientCheckInV2) SetEnableLocalConfigurationProfiles(v bool) {
	o.EnableLocalConfigurationProfiles = &v
}

func (o ClientCheckInV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientCheckInV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CheckInFrequency) {
		toSerialize["checkInFrequency"] = o.CheckInFrequency
	}
	if !IsNil(o.CreateHooks) {
		toSerialize["createHooks"] = o.CreateHooks
	}
	if !IsNil(o.HookLog) {
		toSerialize["hookLog"] = o.HookLog
	}
	if !IsNil(o.HookPolicies) {
		toSerialize["hookPolicies"] = o.HookPolicies
	}
	if !IsNil(o.HookHideRestore) {
		toSerialize["hookHideRestore"] = o.HookHideRestore
	}
	if !IsNil(o.HookMcx) {
		toSerialize["hookMcx"] = o.HookMcx
	}
	if !IsNil(o.BackgroundHooks) {
		toSerialize["backgroundHooks"] = o.BackgroundHooks
	}
	if !IsNil(o.HookDisplayStatus) {
		toSerialize["hookDisplayStatus"] = o.HookDisplayStatus
	}
	if !IsNil(o.CreateStartupScript) {
		toSerialize["createStartupScript"] = o.CreateStartupScript
	}
	if !IsNil(o.StartupLog) {
		toSerialize["startupLog"] = o.StartupLog
	}
	if !IsNil(o.StartupPolicies) {
		toSerialize["startupPolicies"] = o.StartupPolicies
	}
	if !IsNil(o.StartupSsh) {
		toSerialize["startupSsh"] = o.StartupSsh
	}
	if !IsNil(o.StartupMcx) {
		toSerialize["startupMcx"] = o.StartupMcx
	}
	if !IsNil(o.EnableLocalConfigurationProfiles) {
		toSerialize["enableLocalConfigurationProfiles"] = o.EnableLocalConfigurationProfiles
	}
	return toSerialize, nil
}

type NullableClientCheckInV2 struct {
	value *ClientCheckInV2
	isSet bool
}

func (v NullableClientCheckInV2) Get() *ClientCheckInV2 {
	return v.value
}

func (v *NullableClientCheckInV2) Set(val *ClientCheckInV2) {
	v.value = val
	v.isSet = true
}

func (v NullableClientCheckInV2) IsSet() bool {
	return v.isSet
}

func (v *NullableClientCheckInV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientCheckInV2(val *ClientCheckInV2) *NullableClientCheckInV2 {
	return &NullableClientCheckInV2{value: val, isSet: true}
}

func (v NullableClientCheckInV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientCheckInV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


