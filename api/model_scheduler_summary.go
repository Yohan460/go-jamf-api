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

// checks if the SchedulerSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchedulerSummary{}

// SchedulerSummary struct for SchedulerSummary
type SchedulerSummary struct {
	NumberOfPendingJobs *int64 `json:"numberOfPendingJobs,omitempty"`
	NumberOfExecutingJobs *int64 `json:"numberOfExecutingJobs,omitempty"`
	NumberOfExecutedJobs *int64 `json:"numberOfExecutedJobs,omitempty"`
	Started *bool `json:"started,omitempty"`
}

// NewSchedulerSummary instantiates a new SchedulerSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedulerSummary() *SchedulerSummary {
	this := SchedulerSummary{}
	return &this
}

// NewSchedulerSummaryWithDefaults instantiates a new SchedulerSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulerSummaryWithDefaults() *SchedulerSummary {
	this := SchedulerSummary{}
	return &this
}

// GetNumberOfPendingJobs returns the NumberOfPendingJobs field value if set, zero value otherwise.
func (o *SchedulerSummary) GetNumberOfPendingJobs() int64 {
	if o == nil || IsNil(o.NumberOfPendingJobs) {
		var ret int64
		return ret
	}
	return *o.NumberOfPendingJobs
}

// GetNumberOfPendingJobsOk returns a tuple with the NumberOfPendingJobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerSummary) GetNumberOfPendingJobsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfPendingJobs) {
		return nil, false
	}
	return o.NumberOfPendingJobs, true
}

// HasNumberOfPendingJobs returns a boolean if a field has been set.
func (o *SchedulerSummary) HasNumberOfPendingJobs() bool {
	if o != nil && !IsNil(o.NumberOfPendingJobs) {
		return true
	}

	return false
}

// SetNumberOfPendingJobs gets a reference to the given int64 and assigns it to the NumberOfPendingJobs field.
func (o *SchedulerSummary) SetNumberOfPendingJobs(v int64) {
	o.NumberOfPendingJobs = &v
}

// GetNumberOfExecutingJobs returns the NumberOfExecutingJobs field value if set, zero value otherwise.
func (o *SchedulerSummary) GetNumberOfExecutingJobs() int64 {
	if o == nil || IsNil(o.NumberOfExecutingJobs) {
		var ret int64
		return ret
	}
	return *o.NumberOfExecutingJobs
}

// GetNumberOfExecutingJobsOk returns a tuple with the NumberOfExecutingJobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerSummary) GetNumberOfExecutingJobsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfExecutingJobs) {
		return nil, false
	}
	return o.NumberOfExecutingJobs, true
}

// HasNumberOfExecutingJobs returns a boolean if a field has been set.
func (o *SchedulerSummary) HasNumberOfExecutingJobs() bool {
	if o != nil && !IsNil(o.NumberOfExecutingJobs) {
		return true
	}

	return false
}

// SetNumberOfExecutingJobs gets a reference to the given int64 and assigns it to the NumberOfExecutingJobs field.
func (o *SchedulerSummary) SetNumberOfExecutingJobs(v int64) {
	o.NumberOfExecutingJobs = &v
}

// GetNumberOfExecutedJobs returns the NumberOfExecutedJobs field value if set, zero value otherwise.
func (o *SchedulerSummary) GetNumberOfExecutedJobs() int64 {
	if o == nil || IsNil(o.NumberOfExecutedJobs) {
		var ret int64
		return ret
	}
	return *o.NumberOfExecutedJobs
}

// GetNumberOfExecutedJobsOk returns a tuple with the NumberOfExecutedJobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerSummary) GetNumberOfExecutedJobsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfExecutedJobs) {
		return nil, false
	}
	return o.NumberOfExecutedJobs, true
}

// HasNumberOfExecutedJobs returns a boolean if a field has been set.
func (o *SchedulerSummary) HasNumberOfExecutedJobs() bool {
	if o != nil && !IsNil(o.NumberOfExecutedJobs) {
		return true
	}

	return false
}

// SetNumberOfExecutedJobs gets a reference to the given int64 and assigns it to the NumberOfExecutedJobs field.
func (o *SchedulerSummary) SetNumberOfExecutedJobs(v int64) {
	o.NumberOfExecutedJobs = &v
}

// GetStarted returns the Started field value if set, zero value otherwise.
func (o *SchedulerSummary) GetStarted() bool {
	if o == nil || IsNil(o.Started) {
		var ret bool
		return ret
	}
	return *o.Started
}

// GetStartedOk returns a tuple with the Started field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerSummary) GetStartedOk() (*bool, bool) {
	if o == nil || IsNil(o.Started) {
		return nil, false
	}
	return o.Started, true
}

// HasStarted returns a boolean if a field has been set.
func (o *SchedulerSummary) HasStarted() bool {
	if o != nil && !IsNil(o.Started) {
		return true
	}

	return false
}

// SetStarted gets a reference to the given bool and assigns it to the Started field.
func (o *SchedulerSummary) SetStarted(v bool) {
	o.Started = &v
}

func (o SchedulerSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchedulerSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumberOfPendingJobs) {
		toSerialize["numberOfPendingJobs"] = o.NumberOfPendingJobs
	}
	if !IsNil(o.NumberOfExecutingJobs) {
		toSerialize["numberOfExecutingJobs"] = o.NumberOfExecutingJobs
	}
	if !IsNil(o.NumberOfExecutedJobs) {
		toSerialize["numberOfExecutedJobs"] = o.NumberOfExecutedJobs
	}
	if !IsNil(o.Started) {
		toSerialize["started"] = o.Started
	}
	return toSerialize, nil
}

type NullableSchedulerSummary struct {
	value *SchedulerSummary
	isSet bool
}

func (v NullableSchedulerSummary) Get() *SchedulerSummary {
	return v.value
}

func (v *NullableSchedulerSummary) Set(val *SchedulerSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulerSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulerSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulerSummary(val *SchedulerSummary) *NullableSchedulerSummary {
	return &NullableSchedulerSummary{value: val, isSet: true}
}

func (v NullableSchedulerSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulerSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


