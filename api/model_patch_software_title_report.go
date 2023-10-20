/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the PatchSoftwareTitleReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchSoftwareTitleReport{}

// PatchSoftwareTitleReport struct for PatchSoftwareTitleReport
type PatchSoftwareTitleReport struct {
	ComputerName *string `json:"computerName,omitempty"`
	DeviceId *string `json:"deviceId,omitempty"`
	Username *string `json:"username,omitempty"`
	OperatingSystemVersion *string `json:"operatingSystemVersion,omitempty"`
	LastContactTime *time.Time `json:"lastContactTime,omitempty"`
	BuildingName *string `json:"buildingName,omitempty"`
	DepartmentName *string `json:"departmentName,omitempty"`
	SiteName *string `json:"siteName,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewPatchSoftwareTitleReport instantiates a new PatchSoftwareTitleReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchSoftwareTitleReport() *PatchSoftwareTitleReport {
	this := PatchSoftwareTitleReport{}
	return &this
}

// NewPatchSoftwareTitleReportWithDefaults instantiates a new PatchSoftwareTitleReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchSoftwareTitleReportWithDefaults() *PatchSoftwareTitleReport {
	this := PatchSoftwareTitleReport{}
	return &this
}

// GetComputerName returns the ComputerName field value if set, zero value otherwise.
func (o *PatchSoftwareTitleReport) GetComputerName() string {
	if o == nil || IsNil(o.ComputerName) {
		var ret string
		return ret
	}
	return *o.ComputerName
}

// GetComputerNameOk returns a tuple with the ComputerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleReport) GetComputerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ComputerName) {
		return nil, false
	}
	return o.ComputerName, true
}

// HasComputerName returns a boolean if a field has been set.
func (o *PatchSoftwareTitleReport) HasComputerName() bool {
	if o != nil && !IsNil(o.ComputerName) {
		return true
	}

	return false
}

// SetComputerName gets a reference to the given string and assigns it to the ComputerName field.
func (o *PatchSoftwareTitleReport) SetComputerName(v string) {
	o.ComputerName = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *PatchSoftwareTitleReport) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleReport) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *PatchSoftwareTitleReport) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *PatchSoftwareTitleReport) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *PatchSoftwareTitleReport) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleReport) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *PatchSoftwareTitleReport) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *PatchSoftwareTitleReport) SetUsername(v string) {
	o.Username = &v
}

// GetOperatingSystemVersion returns the OperatingSystemVersion field value if set, zero value otherwise.
func (o *PatchSoftwareTitleReport) GetOperatingSystemVersion() string {
	if o == nil || IsNil(o.OperatingSystemVersion) {
		var ret string
		return ret
	}
	return *o.OperatingSystemVersion
}

// GetOperatingSystemVersionOk returns a tuple with the OperatingSystemVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleReport) GetOperatingSystemVersionOk() (*string, bool) {
	if o == nil || IsNil(o.OperatingSystemVersion) {
		return nil, false
	}
	return o.OperatingSystemVersion, true
}

// HasOperatingSystemVersion returns a boolean if a field has been set.
func (o *PatchSoftwareTitleReport) HasOperatingSystemVersion() bool {
	if o != nil && !IsNil(o.OperatingSystemVersion) {
		return true
	}

	return false
}

// SetOperatingSystemVersion gets a reference to the given string and assigns it to the OperatingSystemVersion field.
func (o *PatchSoftwareTitleReport) SetOperatingSystemVersion(v string) {
	o.OperatingSystemVersion = &v
}

// GetLastContactTime returns the LastContactTime field value if set, zero value otherwise.
func (o *PatchSoftwareTitleReport) GetLastContactTime() time.Time {
	if o == nil || IsNil(o.LastContactTime) {
		var ret time.Time
		return ret
	}
	return *o.LastContactTime
}

// GetLastContactTimeOk returns a tuple with the LastContactTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleReport) GetLastContactTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastContactTime) {
		return nil, false
	}
	return o.LastContactTime, true
}

// HasLastContactTime returns a boolean if a field has been set.
func (o *PatchSoftwareTitleReport) HasLastContactTime() bool {
	if o != nil && !IsNil(o.LastContactTime) {
		return true
	}

	return false
}

// SetLastContactTime gets a reference to the given time.Time and assigns it to the LastContactTime field.
func (o *PatchSoftwareTitleReport) SetLastContactTime(v time.Time) {
	o.LastContactTime = &v
}

// GetBuildingName returns the BuildingName field value if set, zero value otherwise.
func (o *PatchSoftwareTitleReport) GetBuildingName() string {
	if o == nil || IsNil(o.BuildingName) {
		var ret string
		return ret
	}
	return *o.BuildingName
}

// GetBuildingNameOk returns a tuple with the BuildingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleReport) GetBuildingNameOk() (*string, bool) {
	if o == nil || IsNil(o.BuildingName) {
		return nil, false
	}
	return o.BuildingName, true
}

// HasBuildingName returns a boolean if a field has been set.
func (o *PatchSoftwareTitleReport) HasBuildingName() bool {
	if o != nil && !IsNil(o.BuildingName) {
		return true
	}

	return false
}

// SetBuildingName gets a reference to the given string and assigns it to the BuildingName field.
func (o *PatchSoftwareTitleReport) SetBuildingName(v string) {
	o.BuildingName = &v
}

// GetDepartmentName returns the DepartmentName field value if set, zero value otherwise.
func (o *PatchSoftwareTitleReport) GetDepartmentName() string {
	if o == nil || IsNil(o.DepartmentName) {
		var ret string
		return ret
	}
	return *o.DepartmentName
}

// GetDepartmentNameOk returns a tuple with the DepartmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleReport) GetDepartmentNameOk() (*string, bool) {
	if o == nil || IsNil(o.DepartmentName) {
		return nil, false
	}
	return o.DepartmentName, true
}

// HasDepartmentName returns a boolean if a field has been set.
func (o *PatchSoftwareTitleReport) HasDepartmentName() bool {
	if o != nil && !IsNil(o.DepartmentName) {
		return true
	}

	return false
}

// SetDepartmentName gets a reference to the given string and assigns it to the DepartmentName field.
func (o *PatchSoftwareTitleReport) SetDepartmentName(v string) {
	o.DepartmentName = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *PatchSoftwareTitleReport) GetSiteName() string {
	if o == nil || IsNil(o.SiteName) {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleReport) GetSiteNameOk() (*string, bool) {
	if o == nil || IsNil(o.SiteName) {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *PatchSoftwareTitleReport) HasSiteName() bool {
	if o != nil && !IsNil(o.SiteName) {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *PatchSoftwareTitleReport) SetSiteName(v string) {
	o.SiteName = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PatchSoftwareTitleReport) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleReport) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PatchSoftwareTitleReport) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PatchSoftwareTitleReport) SetVersion(v string) {
	o.Version = &v
}

func (o PatchSoftwareTitleReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchSoftwareTitleReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComputerName) {
		toSerialize["computerName"] = o.ComputerName
	}
	if !IsNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.OperatingSystemVersion) {
		toSerialize["operatingSystemVersion"] = o.OperatingSystemVersion
	}
	if !IsNil(o.LastContactTime) {
		toSerialize["lastContactTime"] = o.LastContactTime
	}
	if !IsNil(o.BuildingName) {
		toSerialize["buildingName"] = o.BuildingName
	}
	if !IsNil(o.DepartmentName) {
		toSerialize["departmentName"] = o.DepartmentName
	}
	if !IsNil(o.SiteName) {
		toSerialize["siteName"] = o.SiteName
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullablePatchSoftwareTitleReport struct {
	value *PatchSoftwareTitleReport
	isSet bool
}

func (v NullablePatchSoftwareTitleReport) Get() *PatchSoftwareTitleReport {
	return v.value
}

func (v *NullablePatchSoftwareTitleReport) Set(val *PatchSoftwareTitleReport) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchSoftwareTitleReport) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchSoftwareTitleReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchSoftwareTitleReport(val *PatchSoftwareTitleReport) *NullablePatchSoftwareTitleReport {
	return &NullablePatchSoftwareTitleReport{value: val, isSet: true}
}

func (v NullablePatchSoftwareTitleReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchSoftwareTitleReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

