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

// checks if the ComputerContentCachingParent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComputerContentCachingParent{}

// ComputerContentCachingParent struct for ComputerContentCachingParent
type ComputerContentCachingParent struct {
	ContentCachingParentId *string `json:"contentCachingParentId,omitempty"`
	Address *string `json:"address,omitempty"`
	Alerts *ComputerContentCachingParentAlert `json:"alerts,omitempty"`
	Details *ComputerContentCachingParentDetails `json:"details,omitempty"`
	Guid *string `json:"guid,omitempty"`
	Healthy *bool `json:"healthy,omitempty"`
	Port *int64 `json:"port,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewComputerContentCachingParent instantiates a new ComputerContentCachingParent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputerContentCachingParent() *ComputerContentCachingParent {
	this := ComputerContentCachingParent{}
	return &this
}

// NewComputerContentCachingParentWithDefaults instantiates a new ComputerContentCachingParent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputerContentCachingParentWithDefaults() *ComputerContentCachingParent {
	this := ComputerContentCachingParent{}
	return &this
}

// GetContentCachingParentId returns the ContentCachingParentId field value if set, zero value otherwise.
func (o *ComputerContentCachingParent) GetContentCachingParentId() string {
	if o == nil || IsNil(o.ContentCachingParentId) {
		var ret string
		return ret
	}
	return *o.ContentCachingParentId
}

// GetContentCachingParentIdOk returns a tuple with the ContentCachingParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerContentCachingParent) GetContentCachingParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContentCachingParentId) {
		return nil, false
	}
	return o.ContentCachingParentId, true
}

// HasContentCachingParentId returns a boolean if a field has been set.
func (o *ComputerContentCachingParent) HasContentCachingParentId() bool {
	if o != nil && !IsNil(o.ContentCachingParentId) {
		return true
	}

	return false
}

// SetContentCachingParentId gets a reference to the given string and assigns it to the ContentCachingParentId field.
func (o *ComputerContentCachingParent) SetContentCachingParentId(v string) {
	o.ContentCachingParentId = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ComputerContentCachingParent) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerContentCachingParent) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ComputerContentCachingParent) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ComputerContentCachingParent) SetAddress(v string) {
	o.Address = &v
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *ComputerContentCachingParent) GetAlerts() ComputerContentCachingParentAlert {
	if o == nil || IsNil(o.Alerts) {
		var ret ComputerContentCachingParentAlert
		return ret
	}
	return *o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerContentCachingParent) GetAlertsOk() (*ComputerContentCachingParentAlert, bool) {
	if o == nil || IsNil(o.Alerts) {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *ComputerContentCachingParent) HasAlerts() bool {
	if o != nil && !IsNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given ComputerContentCachingParentAlert and assigns it to the Alerts field.
func (o *ComputerContentCachingParent) SetAlerts(v ComputerContentCachingParentAlert) {
	o.Alerts = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ComputerContentCachingParent) GetDetails() ComputerContentCachingParentDetails {
	if o == nil || IsNil(o.Details) {
		var ret ComputerContentCachingParentDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerContentCachingParent) GetDetailsOk() (*ComputerContentCachingParentDetails, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ComputerContentCachingParent) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given ComputerContentCachingParentDetails and assigns it to the Details field.
func (o *ComputerContentCachingParent) SetDetails(v ComputerContentCachingParentDetails) {
	o.Details = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *ComputerContentCachingParent) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerContentCachingParent) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *ComputerContentCachingParent) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *ComputerContentCachingParent) SetGuid(v string) {
	o.Guid = &v
}

// GetHealthy returns the Healthy field value if set, zero value otherwise.
func (o *ComputerContentCachingParent) GetHealthy() bool {
	if o == nil || IsNil(o.Healthy) {
		var ret bool
		return ret
	}
	return *o.Healthy
}

// GetHealthyOk returns a tuple with the Healthy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerContentCachingParent) GetHealthyOk() (*bool, bool) {
	if o == nil || IsNil(o.Healthy) {
		return nil, false
	}
	return o.Healthy, true
}

// HasHealthy returns a boolean if a field has been set.
func (o *ComputerContentCachingParent) HasHealthy() bool {
	if o != nil && !IsNil(o.Healthy) {
		return true
	}

	return false
}

// SetHealthy gets a reference to the given bool and assigns it to the Healthy field.
func (o *ComputerContentCachingParent) SetHealthy(v bool) {
	o.Healthy = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ComputerContentCachingParent) GetPort() int64 {
	if o == nil || IsNil(o.Port) {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerContentCachingParent) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ComputerContentCachingParent) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *ComputerContentCachingParent) SetPort(v int64) {
	o.Port = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ComputerContentCachingParent) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerContentCachingParent) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ComputerContentCachingParent) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ComputerContentCachingParent) SetVersion(v string) {
	o.Version = &v
}

func (o ComputerContentCachingParent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComputerContentCachingParent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentCachingParentId) {
		toSerialize["contentCachingParentId"] = o.ContentCachingParentId
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Alerts) {
		toSerialize["alerts"] = o.Alerts
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Guid) {
		toSerialize["guid"] = o.Guid
	}
	if !IsNil(o.Healthy) {
		toSerialize["healthy"] = o.Healthy
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableComputerContentCachingParent struct {
	value *ComputerContentCachingParent
	isSet bool
}

func (v NullableComputerContentCachingParent) Get() *ComputerContentCachingParent {
	return v.value
}

func (v *NullableComputerContentCachingParent) Set(val *ComputerContentCachingParent) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerContentCachingParent) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerContentCachingParent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerContentCachingParent(val *ComputerContentCachingParent) *NullableComputerContentCachingParent {
	return &NullableComputerContentCachingParent{value: val, isSet: true}
}

func (v NullableComputerContentCachingParent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerContentCachingParent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


