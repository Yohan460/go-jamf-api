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

// checks if the ComputerOverview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComputerOverview{}

// ComputerOverview struct for ComputerOverview
type ComputerOverview struct {
	Id *string `json:"id,omitempty"`
	Location *ComputerLocation `json:"location,omitempty"`
	Name *string `json:"name,omitempty"`
	Udid *string `json:"udid,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty"`
	LastContactDate *string `json:"lastContactDate,omitempty"`
	LastReportDate *string `json:"lastReportDate,omitempty"`
	LastEnrolledDate *string `json:"lastEnrolledDate,omitempty"`
	OperatingSystemVersion *string `json:"operatingSystemVersion,omitempty"`
	OperatingSystemBuild *string `json:"operatingSystemBuild,omitempty"`
	// Collected for macOS 13.0 or later
	OperatingSystemSupplementalBuildVersion *string `json:"operatingSystemSupplementalBuildVersion,omitempty"`
	// Collected for macOS 13.0 or later
	OperatingSystemRapidSecurityResponse *string `json:"operatingSystemRapidSecurityResponse,omitempty"`
	IpAddress *string `json:"ipAddress,omitempty"`
	MacAddress *string `json:"macAddress,omitempty"`
	AssetTag *string `json:"assetTag,omitempty"`
	ModelIdentifier *string `json:"modelIdentifier,omitempty"`
	MdmAccessRights *int64 `json:"mdmAccessRights,omitempty"`
	IsManaged *bool `json:"isManaged,omitempty"`
	ManagementId *string `json:"managementId,omitempty"`
}

// NewComputerOverview instantiates a new ComputerOverview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputerOverview() *ComputerOverview {
	this := ComputerOverview{}
	return &this
}

// NewComputerOverviewWithDefaults instantiates a new ComputerOverview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputerOverviewWithDefaults() *ComputerOverview {
	this := ComputerOverview{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ComputerOverview) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOverview) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ComputerOverview) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ComputerOverview) SetId(v string) {
	o.Id = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ComputerOverview) GetLocation() ComputerLocation {
	if o == nil || IsNil(o.Location) {
		var ret ComputerLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOverview) GetLocationOk() (*ComputerLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ComputerOverview) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given ComputerLocation and assigns it to the Location field.
func (o *ComputerOverview) SetLocation(v ComputerLocation) {
	o.Location = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComputerOverview) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOverview) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComputerOverview) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComputerOverview) SetName(v string) {
	o.Name = &v
}

// GetUdid returns the Udid field value if set, zero value otherwise.
func (o *ComputerOverview) GetUdid() string {
	if o == nil || IsNil(o.Udid) {
		var ret string
		return ret
	}
	return *o.Udid
}

// GetUdidOk returns a tuple with the Udid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOverview) GetUdidOk() (*string, bool) {
	if o == nil || IsNil(o.Udid) {
		return nil, false
	}
	return o.Udid, true
}

// HasUdid returns a boolean if a field has been set.
func (o *ComputerOverview) HasUdid() bool {
	if o != nil && !IsNil(o.Udid) {
		return true
	}

	return false
}

// SetUdid gets a reference to the given string and assigns it to the Udid field.
func (o *ComputerOverview) SetUdid(v string) {
	o.Udid = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *ComputerOverview) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOverview) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *ComputerOverview) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *ComputerOverview) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetLastContactDate returns the LastContactDate field value if set, zero value otherwise.
func (o *ComputerOverview) GetLastContactDate() string {
	if o == nil || IsNil(o.LastContactDate) {
		var ret string
		return ret
	}
	return *o.LastContactDate
}

// GetLastContactDateOk returns a tuple with the LastContactDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOverview) GetLastContactDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastContactDate) {
		return nil, false
	}
	return o.LastContactDate, true
}

// HasLastContactDate returns a boolean if a field has been set.
func (o *ComputerOverview) HasLastContactDate() bool {
	if o != nil && !IsNil(o.LastContactDate) {
		return true
	}

	return false
}

// SetLastContactDate gets a reference to the given string and assigns it to the LastContactDate field.
func (o *ComputerOverview) SetLastContactDate(v string) {
	o.LastContactDate = &v
}

// GetLastReportDate returns the LastReportDate field value if set, zero value otherwise.
func (o *ComputerOverview) GetLastReportDate() string {
	if o == nil || IsNil(o.LastReportDate) {
		var ret string
		return ret
	}
	return *o.LastReportDate
}

// GetLastReportDateOk returns a tuple with the LastReportDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOverview) GetLastReportDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastReportDate) {
		return nil, false
	}
	return o.LastReportDate, true
}

// HasLastReportDate returns a boolean if a field has been set.
func (o *ComputerOverview) HasLastReportDate() bool {
	if o != nil && !IsNil(o.LastReportDate) {
		return true
	}

	return false
}

// SetLastReportDate gets a reference to the given string and assigns it to the LastReportDate field.
func (o *ComputerOverview) SetLastReportDate(v string) {
	o.LastReportDate = &v
}

// GetLastEnrolledDate returns the LastEnrolledDate field value if set, zero value otherwise.
func (o *ComputerOverview) GetLastEnrolledDate() string {
	if o == nil || IsNil(o.LastEnrolledDate) {
		var ret string
		return ret
	}
	return *o.LastEnrolledDate
}

// GetLastEnrolledDateOk returns a tuple with the LastEnrolledDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOverview) GetLastEnrolledDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastEnrolledDate) {
		return nil, false
	}
	return o.LastEnrolledDate, true
}

// HasLastEnrolledDate returns a boolean if a field has been set.
func (o *ComputerOverview) HasLastEnrolledDate() bool {
	if o != nil && !IsNil(o.LastEnrolledDate) {
		return true
	}

	return false
}

// SetLastEnrolledDate gets a reference to the given string and assigns it to the LastEnrolledDate field.
func (o *ComputerOverview) SetLastEnrolledDate(v string) {
	o.LastEnrolledDate = &v
}

// GetOperatingSystemVersion returns the OperatingSystemVersion field value if set, zero value otherwise.
func (o *ComputerOverview) GetOperatingSystemVersion() string {
	if o == nil || IsNil(o.OperatingSystemVersion) {
		var ret string
		return ret
	}
	return *o.OperatingSystemVersion
}

// GetOperatingSystemVersionOk returns a tuple with the OperatingSystemVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOverview) GetOperatingSystemVersionOk() (*string, bool) {
	if o == nil || IsNil(o.OperatingSystemVersion) {
		return nil, false
	}
	return o.OperatingSystemVersion, true
}

// HasOperatingSystemVersion returns a boolean if a field has been set.
func (o *ComputerOverview) HasOperatingSystemVersion() bool {
	if o != nil && !IsNil(o.OperatingSystemVersion) {
		return true
	}

	return false
}

// SetOperatingSystemVersion gets a reference to the given string and assigns it to the OperatingSystemVersion field.
func (o *ComputerOverview) SetOperatingSystemVersion(v string) {
	o.OperatingSystemVersion = &v
}

// GetOperatingSystemBuild returns the OperatingSystemBuild field value if set, zero value otherwise.
func (o *ComputerOverview) GetOperatingSystemBuild() string {
	if o == nil || IsNil(o.OperatingSystemBuild) {
		var ret string
		return ret
	}
	return *o.OperatingSystemBuild
}

// GetOperatingSystemBuildOk returns a tuple with the OperatingSystemBuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOverview) GetOperatingSystemBuildOk() (*string, bool) {
	if o == nil || IsNil(o.OperatingSystemBuild) {
		return nil, false
	}
	return o.OperatingSystemBuild, true
}

// HasOperatingSystemBuild returns a boolean if a field has been set.
func (o *ComputerOverview) HasOperatingSystemBuild() bool {
	if o != nil && !IsNil(o.OperatingSystemBuild) {
		return true
	}

	return false
}

// SetOperatingSystemBuild gets a reference to the given string and assigns it to the OperatingSystemBuild field.
func (o *ComputerOverview) SetOperatingSystemBuild(v string) {
	o.OperatingSystemBuild = &v
}

// GetOperatingSystemSupplementalBuildVersion returns the OperatingSystemSupplementalBuildVersion field value if set, zero value otherwise.
func (o *ComputerOverview) GetOperatingSystemSupplementalBuildVersion() string {
	if o == nil || IsNil(o.OperatingSystemSupplementalBuildVersion) {
		var ret string
		return ret
	}
	return *o.OperatingSystemSupplementalBuildVersion
}

// GetOperatingSystemSupplementalBuildVersionOk returns a tuple with the OperatingSystemSupplementalBuildVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOverview) GetOperatingSystemSupplementalBuildVersionOk() (*string, bool) {
	if o == nil || IsNil(o.OperatingSystemSupplementalBuildVersion) {
		return nil, false
	}
	return o.OperatingSystemSupplementalBuildVersion, true
}

// HasOperatingSystemSupplementalBuildVersion returns a boolean if a field has been set.
func (o *ComputerOverview) HasOperatingSystemSupplementalBuildVersion() bool {
	if o != nil && !IsNil(o.OperatingSystemSupplementalBuildVersion) {
		return true
	}

	return false
}

// SetOperatingSystemSupplementalBuildVersion gets a reference to the given string and assigns it to the OperatingSystemSupplementalBuildVersion field.
func (o *ComputerOverview) SetOperatingSystemSupplementalBuildVersion(v string) {
	o.OperatingSystemSupplementalBuildVersion = &v
}

// GetOperatingSystemRapidSecurityResponse returns the OperatingSystemRapidSecurityResponse field value if set, zero value otherwise.
func (o *ComputerOverview) GetOperatingSystemRapidSecurityResponse() string {
	if o == nil || IsNil(o.OperatingSystemRapidSecurityResponse) {
		var ret string
		return ret
	}
	return *o.OperatingSystemRapidSecurityResponse
}

// GetOperatingSystemRapidSecurityResponseOk returns a tuple with the OperatingSystemRapidSecurityResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOverview) GetOperatingSystemRapidSecurityResponseOk() (*string, bool) {
	if o == nil || IsNil(o.OperatingSystemRapidSecurityResponse) {
		return nil, false
	}
	return o.OperatingSystemRapidSecurityResponse, true
}

// HasOperatingSystemRapidSecurityResponse returns a boolean if a field has been set.
func (o *ComputerOverview) HasOperatingSystemRapidSecurityResponse() bool {
	if o != nil && !IsNil(o.OperatingSystemRapidSecurityResponse) {
		return true
	}

	return false
}

// SetOperatingSystemRapidSecurityResponse gets a reference to the given string and assigns it to the OperatingSystemRapidSecurityResponse field.
func (o *ComputerOverview) SetOperatingSystemRapidSecurityResponse(v string) {
	o.OperatingSystemRapidSecurityResponse = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *ComputerOverview) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOverview) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *ComputerOverview) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *ComputerOverview) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *ComputerOverview) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOverview) GetMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *ComputerOverview) HasMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *ComputerOverview) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise.
func (o *ComputerOverview) GetAssetTag() string {
	if o == nil || IsNil(o.AssetTag) {
		var ret string
		return ret
	}
	return *o.AssetTag
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOverview) GetAssetTagOk() (*string, bool) {
	if o == nil || IsNil(o.AssetTag) {
		return nil, false
	}
	return o.AssetTag, true
}

// HasAssetTag returns a boolean if a field has been set.
func (o *ComputerOverview) HasAssetTag() bool {
	if o != nil && !IsNil(o.AssetTag) {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given string and assigns it to the AssetTag field.
func (o *ComputerOverview) SetAssetTag(v string) {
	o.AssetTag = &v
}

// GetModelIdentifier returns the ModelIdentifier field value if set, zero value otherwise.
func (o *ComputerOverview) GetModelIdentifier() string {
	if o == nil || IsNil(o.ModelIdentifier) {
		var ret string
		return ret
	}
	return *o.ModelIdentifier
}

// GetModelIdentifierOk returns a tuple with the ModelIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOverview) GetModelIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.ModelIdentifier) {
		return nil, false
	}
	return o.ModelIdentifier, true
}

// HasModelIdentifier returns a boolean if a field has been set.
func (o *ComputerOverview) HasModelIdentifier() bool {
	if o != nil && !IsNil(o.ModelIdentifier) {
		return true
	}

	return false
}

// SetModelIdentifier gets a reference to the given string and assigns it to the ModelIdentifier field.
func (o *ComputerOverview) SetModelIdentifier(v string) {
	o.ModelIdentifier = &v
}

// GetMdmAccessRights returns the MdmAccessRights field value if set, zero value otherwise.
func (o *ComputerOverview) GetMdmAccessRights() int64 {
	if o == nil || IsNil(o.MdmAccessRights) {
		var ret int64
		return ret
	}
	return *o.MdmAccessRights
}

// GetMdmAccessRightsOk returns a tuple with the MdmAccessRights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOverview) GetMdmAccessRightsOk() (*int64, bool) {
	if o == nil || IsNil(o.MdmAccessRights) {
		return nil, false
	}
	return o.MdmAccessRights, true
}

// HasMdmAccessRights returns a boolean if a field has been set.
func (o *ComputerOverview) HasMdmAccessRights() bool {
	if o != nil && !IsNil(o.MdmAccessRights) {
		return true
	}

	return false
}

// SetMdmAccessRights gets a reference to the given int64 and assigns it to the MdmAccessRights field.
func (o *ComputerOverview) SetMdmAccessRights(v int64) {
	o.MdmAccessRights = &v
}

// GetIsManaged returns the IsManaged field value if set, zero value otherwise.
func (o *ComputerOverview) GetIsManaged() bool {
	if o == nil || IsNil(o.IsManaged) {
		var ret bool
		return ret
	}
	return *o.IsManaged
}

// GetIsManagedOk returns a tuple with the IsManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOverview) GetIsManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsManaged) {
		return nil, false
	}
	return o.IsManaged, true
}

// HasIsManaged returns a boolean if a field has been set.
func (o *ComputerOverview) HasIsManaged() bool {
	if o != nil && !IsNil(o.IsManaged) {
		return true
	}

	return false
}

// SetIsManaged gets a reference to the given bool and assigns it to the IsManaged field.
func (o *ComputerOverview) SetIsManaged(v bool) {
	o.IsManaged = &v
}

// GetManagementId returns the ManagementId field value if set, zero value otherwise.
func (o *ComputerOverview) GetManagementId() string {
	if o == nil || IsNil(o.ManagementId) {
		var ret string
		return ret
	}
	return *o.ManagementId
}

// GetManagementIdOk returns a tuple with the ManagementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerOverview) GetManagementIdOk() (*string, bool) {
	if o == nil || IsNil(o.ManagementId) {
		return nil, false
	}
	return o.ManagementId, true
}

// HasManagementId returns a boolean if a field has been set.
func (o *ComputerOverview) HasManagementId() bool {
	if o != nil && !IsNil(o.ManagementId) {
		return true
	}

	return false
}

// SetManagementId gets a reference to the given string and assigns it to the ManagementId field.
func (o *ComputerOverview) SetManagementId(v string) {
	o.ManagementId = &v
}

func (o ComputerOverview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComputerOverview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Udid) {
		toSerialize["udid"] = o.Udid
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if !IsNil(o.LastContactDate) {
		toSerialize["lastContactDate"] = o.LastContactDate
	}
	if !IsNil(o.LastReportDate) {
		toSerialize["lastReportDate"] = o.LastReportDate
	}
	if !IsNil(o.LastEnrolledDate) {
		toSerialize["lastEnrolledDate"] = o.LastEnrolledDate
	}
	if !IsNil(o.OperatingSystemVersion) {
		toSerialize["operatingSystemVersion"] = o.OperatingSystemVersion
	}
	if !IsNil(o.OperatingSystemBuild) {
		toSerialize["operatingSystemBuild"] = o.OperatingSystemBuild
	}
	if !IsNil(o.OperatingSystemSupplementalBuildVersion) {
		toSerialize["operatingSystemSupplementalBuildVersion"] = o.OperatingSystemSupplementalBuildVersion
	}
	if !IsNil(o.OperatingSystemRapidSecurityResponse) {
		toSerialize["operatingSystemRapidSecurityResponse"] = o.OperatingSystemRapidSecurityResponse
	}
	if !IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !IsNil(o.MacAddress) {
		toSerialize["macAddress"] = o.MacAddress
	}
	if !IsNil(o.AssetTag) {
		toSerialize["assetTag"] = o.AssetTag
	}
	if !IsNil(o.ModelIdentifier) {
		toSerialize["modelIdentifier"] = o.ModelIdentifier
	}
	if !IsNil(o.MdmAccessRights) {
		toSerialize["mdmAccessRights"] = o.MdmAccessRights
	}
	if !IsNil(o.IsManaged) {
		toSerialize["isManaged"] = o.IsManaged
	}
	if !IsNil(o.ManagementId) {
		toSerialize["managementId"] = o.ManagementId
	}
	return toSerialize, nil
}

type NullableComputerOverview struct {
	value *ComputerOverview
	isSet bool
}

func (v NullableComputerOverview) Get() *ComputerOverview {
	return v.value
}

func (v *NullableComputerOverview) Set(val *ComputerOverview) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerOverview) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerOverview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerOverview(val *ComputerOverview) *NullableComputerOverview {
	return &NullableComputerOverview{value: val, isSet: true}
}

func (v NullableComputerOverview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerOverview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


