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

// checks if the StartupStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartupStatus{}

// StartupStatus struct for StartupStatus
type StartupStatus struct {
	Step *string `json:"step,omitempty"`
	StepCode *string `json:"stepCode,omitempty"`
	StepParam NullableString `json:"stepParam,omitempty"`
	Percentage *int64 `json:"percentage,omitempty"`
	Warning NullableString `json:"warning,omitempty"`
	WarningCode NullableString `json:"warningCode,omitempty"`
	WarningParam NullableString `json:"warningParam,omitempty"`
	Error NullableString `json:"error,omitempty"`
	ErrorCode NullableString `json:"errorCode,omitempty"`
	SetupAssistantNecessary *bool `json:"setupAssistantNecessary,omitempty"`
}

// NewStartupStatus instantiates a new StartupStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartupStatus() *StartupStatus {
	this := StartupStatus{}
	return &this
}

// NewStartupStatusWithDefaults instantiates a new StartupStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartupStatusWithDefaults() *StartupStatus {
	this := StartupStatus{}
	return &this
}

// GetStep returns the Step field value if set, zero value otherwise.
func (o *StartupStatus) GetStep() string {
	if o == nil || IsNil(o.Step) {
		var ret string
		return ret
	}
	return *o.Step
}

// GetStepOk returns a tuple with the Step field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupStatus) GetStepOk() (*string, bool) {
	if o == nil || IsNil(o.Step) {
		return nil, false
	}
	return o.Step, true
}

// HasStep returns a boolean if a field has been set.
func (o *StartupStatus) HasStep() bool {
	if o != nil && !IsNil(o.Step) {
		return true
	}

	return false
}

// SetStep gets a reference to the given string and assigns it to the Step field.
func (o *StartupStatus) SetStep(v string) {
	o.Step = &v
}

// GetStepCode returns the StepCode field value if set, zero value otherwise.
func (o *StartupStatus) GetStepCode() string {
	if o == nil || IsNil(o.StepCode) {
		var ret string
		return ret
	}
	return *o.StepCode
}

// GetStepCodeOk returns a tuple with the StepCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupStatus) GetStepCodeOk() (*string, bool) {
	if o == nil || IsNil(o.StepCode) {
		return nil, false
	}
	return o.StepCode, true
}

// HasStepCode returns a boolean if a field has been set.
func (o *StartupStatus) HasStepCode() bool {
	if o != nil && !IsNil(o.StepCode) {
		return true
	}

	return false
}

// SetStepCode gets a reference to the given string and assigns it to the StepCode field.
func (o *StartupStatus) SetStepCode(v string) {
	o.StepCode = &v
}

// GetStepParam returns the StepParam field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StartupStatus) GetStepParam() string {
	if o == nil || IsNil(o.StepParam.Get()) {
		var ret string
		return ret
	}
	return *o.StepParam.Get()
}

// GetStepParamOk returns a tuple with the StepParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StartupStatus) GetStepParamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StepParam.Get(), o.StepParam.IsSet()
}

// HasStepParam returns a boolean if a field has been set.
func (o *StartupStatus) HasStepParam() bool {
	if o != nil && o.StepParam.IsSet() {
		return true
	}

	return false
}

// SetStepParam gets a reference to the given NullableString and assigns it to the StepParam field.
func (o *StartupStatus) SetStepParam(v string) {
	o.StepParam.Set(&v)
}
// SetStepParamNil sets the value for StepParam to be an explicit nil
func (o *StartupStatus) SetStepParamNil() {
	o.StepParam.Set(nil)
}

// UnsetStepParam ensures that no value is present for StepParam, not even an explicit nil
func (o *StartupStatus) UnsetStepParam() {
	o.StepParam.Unset()
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *StartupStatus) GetPercentage() int64 {
	if o == nil || IsNil(o.Percentage) {
		var ret int64
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupStatus) GetPercentageOk() (*int64, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *StartupStatus) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given int64 and assigns it to the Percentage field.
func (o *StartupStatus) SetPercentage(v int64) {
	o.Percentage = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StartupStatus) GetWarning() string {
	if o == nil || IsNil(o.Warning.Get()) {
		var ret string
		return ret
	}
	return *o.Warning.Get()
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StartupStatus) GetWarningOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Warning.Get(), o.Warning.IsSet()
}

// HasWarning returns a boolean if a field has been set.
func (o *StartupStatus) HasWarning() bool {
	if o != nil && o.Warning.IsSet() {
		return true
	}

	return false
}

// SetWarning gets a reference to the given NullableString and assigns it to the Warning field.
func (o *StartupStatus) SetWarning(v string) {
	o.Warning.Set(&v)
}
// SetWarningNil sets the value for Warning to be an explicit nil
func (o *StartupStatus) SetWarningNil() {
	o.Warning.Set(nil)
}

// UnsetWarning ensures that no value is present for Warning, not even an explicit nil
func (o *StartupStatus) UnsetWarning() {
	o.Warning.Unset()
}

// GetWarningCode returns the WarningCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StartupStatus) GetWarningCode() string {
	if o == nil || IsNil(o.WarningCode.Get()) {
		var ret string
		return ret
	}
	return *o.WarningCode.Get()
}

// GetWarningCodeOk returns a tuple with the WarningCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StartupStatus) GetWarningCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WarningCode.Get(), o.WarningCode.IsSet()
}

// HasWarningCode returns a boolean if a field has been set.
func (o *StartupStatus) HasWarningCode() bool {
	if o != nil && o.WarningCode.IsSet() {
		return true
	}

	return false
}

// SetWarningCode gets a reference to the given NullableString and assigns it to the WarningCode field.
func (o *StartupStatus) SetWarningCode(v string) {
	o.WarningCode.Set(&v)
}
// SetWarningCodeNil sets the value for WarningCode to be an explicit nil
func (o *StartupStatus) SetWarningCodeNil() {
	o.WarningCode.Set(nil)
}

// UnsetWarningCode ensures that no value is present for WarningCode, not even an explicit nil
func (o *StartupStatus) UnsetWarningCode() {
	o.WarningCode.Unset()
}

// GetWarningParam returns the WarningParam field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StartupStatus) GetWarningParam() string {
	if o == nil || IsNil(o.WarningParam.Get()) {
		var ret string
		return ret
	}
	return *o.WarningParam.Get()
}

// GetWarningParamOk returns a tuple with the WarningParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StartupStatus) GetWarningParamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WarningParam.Get(), o.WarningParam.IsSet()
}

// HasWarningParam returns a boolean if a field has been set.
func (o *StartupStatus) HasWarningParam() bool {
	if o != nil && o.WarningParam.IsSet() {
		return true
	}

	return false
}

// SetWarningParam gets a reference to the given NullableString and assigns it to the WarningParam field.
func (o *StartupStatus) SetWarningParam(v string) {
	o.WarningParam.Set(&v)
}
// SetWarningParamNil sets the value for WarningParam to be an explicit nil
func (o *StartupStatus) SetWarningParamNil() {
	o.WarningParam.Set(nil)
}

// UnsetWarningParam ensures that no value is present for WarningParam, not even an explicit nil
func (o *StartupStatus) UnsetWarningParam() {
	o.WarningParam.Unset()
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StartupStatus) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StartupStatus) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *StartupStatus) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *StartupStatus) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *StartupStatus) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *StartupStatus) UnsetError() {
	o.Error.Unset()
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StartupStatus) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StartupStatus) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *StartupStatus) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableString and assigns it to the ErrorCode field.
func (o *StartupStatus) SetErrorCode(v string) {
	o.ErrorCode.Set(&v)
}
// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *StartupStatus) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *StartupStatus) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

// GetSetupAssistantNecessary returns the SetupAssistantNecessary field value if set, zero value otherwise.
func (o *StartupStatus) GetSetupAssistantNecessary() bool {
	if o == nil || IsNil(o.SetupAssistantNecessary) {
		var ret bool
		return ret
	}
	return *o.SetupAssistantNecessary
}

// GetSetupAssistantNecessaryOk returns a tuple with the SetupAssistantNecessary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartupStatus) GetSetupAssistantNecessaryOk() (*bool, bool) {
	if o == nil || IsNil(o.SetupAssistantNecessary) {
		return nil, false
	}
	return o.SetupAssistantNecessary, true
}

// HasSetupAssistantNecessary returns a boolean if a field has been set.
func (o *StartupStatus) HasSetupAssistantNecessary() bool {
	if o != nil && !IsNil(o.SetupAssistantNecessary) {
		return true
	}

	return false
}

// SetSetupAssistantNecessary gets a reference to the given bool and assigns it to the SetupAssistantNecessary field.
func (o *StartupStatus) SetSetupAssistantNecessary(v bool) {
	o.SetupAssistantNecessary = &v
}

func (o StartupStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartupStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Step) {
		toSerialize["step"] = o.Step
	}
	if !IsNil(o.StepCode) {
		toSerialize["stepCode"] = o.StepCode
	}
	if o.StepParam.IsSet() {
		toSerialize["stepParam"] = o.StepParam.Get()
	}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	if o.Warning.IsSet() {
		toSerialize["warning"] = o.Warning.Get()
	}
	if o.WarningCode.IsSet() {
		toSerialize["warningCode"] = o.WarningCode.Get()
	}
	if o.WarningParam.IsSet() {
		toSerialize["warningParam"] = o.WarningParam.Get()
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if o.ErrorCode.IsSet() {
		toSerialize["errorCode"] = o.ErrorCode.Get()
	}
	if !IsNil(o.SetupAssistantNecessary) {
		toSerialize["setupAssistantNecessary"] = o.SetupAssistantNecessary
	}
	return toSerialize, nil
}

type NullableStartupStatus struct {
	value *StartupStatus
	isSet bool
}

func (v NullableStartupStatus) Get() *StartupStatus {
	return v.value
}

func (v *NullableStartupStatus) Set(val *StartupStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableStartupStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStartupStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartupStatus(val *StartupStatus) *NullableStartupStatus {
	return &NullableStartupStatus{value: val, isSet: true}
}

func (v NullableStartupStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartupStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


