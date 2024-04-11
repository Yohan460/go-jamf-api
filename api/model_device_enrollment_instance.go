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

// checks if the DeviceEnrollmentInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceEnrollmentInstance{}

// DeviceEnrollmentInstance struct for DeviceEnrollmentInstance
type DeviceEnrollmentInstance struct {
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	SupervisionIdentityId *string `json:"supervisionIdentityId,omitempty"`
	SiteId *string `json:"siteId,omitempty"`
	ServerName *string `json:"serverName,omitempty"`
	ServerUuid *string `json:"serverUuid,omitempty"`
	AdminId *string `json:"adminId,omitempty"`
	OrgName *string `json:"orgName,omitempty"`
	OrgEmail *string `json:"orgEmail,omitempty"`
	OrgPhone *string `json:"orgPhone,omitempty"`
	OrgAddress *string `json:"orgAddress,omitempty"`
	TokenExpirationDate *string `json:"tokenExpirationDate,omitempty"`
}

type _DeviceEnrollmentInstance DeviceEnrollmentInstance

// NewDeviceEnrollmentInstance instantiates a new DeviceEnrollmentInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceEnrollmentInstance(name string) *DeviceEnrollmentInstance {
	this := DeviceEnrollmentInstance{}
	this.Name = name
	return &this
}

// NewDeviceEnrollmentInstanceWithDefaults instantiates a new DeviceEnrollmentInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceEnrollmentInstanceWithDefaults() *DeviceEnrollmentInstance {
	this := DeviceEnrollmentInstance{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceEnrollmentInstance) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentInstance) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceEnrollmentInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeviceEnrollmentInstance) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *DeviceEnrollmentInstance) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentInstance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeviceEnrollmentInstance) SetName(v string) {
	o.Name = v
}

// GetSupervisionIdentityId returns the SupervisionIdentityId field value if set, zero value otherwise.
func (o *DeviceEnrollmentInstance) GetSupervisionIdentityId() string {
	if o == nil || IsNil(o.SupervisionIdentityId) {
		var ret string
		return ret
	}
	return *o.SupervisionIdentityId
}

// GetSupervisionIdentityIdOk returns a tuple with the SupervisionIdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentInstance) GetSupervisionIdentityIdOk() (*string, bool) {
	if o == nil || IsNil(o.SupervisionIdentityId) {
		return nil, false
	}
	return o.SupervisionIdentityId, true
}

// HasSupervisionIdentityId returns a boolean if a field has been set.
func (o *DeviceEnrollmentInstance) HasSupervisionIdentityId() bool {
	if o != nil && !IsNil(o.SupervisionIdentityId) {
		return true
	}

	return false
}

// SetSupervisionIdentityId gets a reference to the given string and assigns it to the SupervisionIdentityId field.
func (o *DeviceEnrollmentInstance) SetSupervisionIdentityId(v string) {
	o.SupervisionIdentityId = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *DeviceEnrollmentInstance) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentInstance) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *DeviceEnrollmentInstance) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *DeviceEnrollmentInstance) SetSiteId(v string) {
	o.SiteId = &v
}

// GetServerName returns the ServerName field value if set, zero value otherwise.
func (o *DeviceEnrollmentInstance) GetServerName() string {
	if o == nil || IsNil(o.ServerName) {
		var ret string
		return ret
	}
	return *o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentInstance) GetServerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServerName) {
		return nil, false
	}
	return o.ServerName, true
}

// HasServerName returns a boolean if a field has been set.
func (o *DeviceEnrollmentInstance) HasServerName() bool {
	if o != nil && !IsNil(o.ServerName) {
		return true
	}

	return false
}

// SetServerName gets a reference to the given string and assigns it to the ServerName field.
func (o *DeviceEnrollmentInstance) SetServerName(v string) {
	o.ServerName = &v
}

// GetServerUuid returns the ServerUuid field value if set, zero value otherwise.
func (o *DeviceEnrollmentInstance) GetServerUuid() string {
	if o == nil || IsNil(o.ServerUuid) {
		var ret string
		return ret
	}
	return *o.ServerUuid
}

// GetServerUuidOk returns a tuple with the ServerUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentInstance) GetServerUuidOk() (*string, bool) {
	if o == nil || IsNil(o.ServerUuid) {
		return nil, false
	}
	return o.ServerUuid, true
}

// HasServerUuid returns a boolean if a field has been set.
func (o *DeviceEnrollmentInstance) HasServerUuid() bool {
	if o != nil && !IsNil(o.ServerUuid) {
		return true
	}

	return false
}

// SetServerUuid gets a reference to the given string and assigns it to the ServerUuid field.
func (o *DeviceEnrollmentInstance) SetServerUuid(v string) {
	o.ServerUuid = &v
}

// GetAdminId returns the AdminId field value if set, zero value otherwise.
func (o *DeviceEnrollmentInstance) GetAdminId() string {
	if o == nil || IsNil(o.AdminId) {
		var ret string
		return ret
	}
	return *o.AdminId
}

// GetAdminIdOk returns a tuple with the AdminId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentInstance) GetAdminIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdminId) {
		return nil, false
	}
	return o.AdminId, true
}

// HasAdminId returns a boolean if a field has been set.
func (o *DeviceEnrollmentInstance) HasAdminId() bool {
	if o != nil && !IsNil(o.AdminId) {
		return true
	}

	return false
}

// SetAdminId gets a reference to the given string and assigns it to the AdminId field.
func (o *DeviceEnrollmentInstance) SetAdminId(v string) {
	o.AdminId = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *DeviceEnrollmentInstance) GetOrgName() string {
	if o == nil || IsNil(o.OrgName) {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentInstance) GetOrgNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrgName) {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *DeviceEnrollmentInstance) HasOrgName() bool {
	if o != nil && !IsNil(o.OrgName) {
		return true
	}

	return false
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *DeviceEnrollmentInstance) SetOrgName(v string) {
	o.OrgName = &v
}

// GetOrgEmail returns the OrgEmail field value if set, zero value otherwise.
func (o *DeviceEnrollmentInstance) GetOrgEmail() string {
	if o == nil || IsNil(o.OrgEmail) {
		var ret string
		return ret
	}
	return *o.OrgEmail
}

// GetOrgEmailOk returns a tuple with the OrgEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentInstance) GetOrgEmailOk() (*string, bool) {
	if o == nil || IsNil(o.OrgEmail) {
		return nil, false
	}
	return o.OrgEmail, true
}

// HasOrgEmail returns a boolean if a field has been set.
func (o *DeviceEnrollmentInstance) HasOrgEmail() bool {
	if o != nil && !IsNil(o.OrgEmail) {
		return true
	}

	return false
}

// SetOrgEmail gets a reference to the given string and assigns it to the OrgEmail field.
func (o *DeviceEnrollmentInstance) SetOrgEmail(v string) {
	o.OrgEmail = &v
}

// GetOrgPhone returns the OrgPhone field value if set, zero value otherwise.
func (o *DeviceEnrollmentInstance) GetOrgPhone() string {
	if o == nil || IsNil(o.OrgPhone) {
		var ret string
		return ret
	}
	return *o.OrgPhone
}

// GetOrgPhoneOk returns a tuple with the OrgPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentInstance) GetOrgPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.OrgPhone) {
		return nil, false
	}
	return o.OrgPhone, true
}

// HasOrgPhone returns a boolean if a field has been set.
func (o *DeviceEnrollmentInstance) HasOrgPhone() bool {
	if o != nil && !IsNil(o.OrgPhone) {
		return true
	}

	return false
}

// SetOrgPhone gets a reference to the given string and assigns it to the OrgPhone field.
func (o *DeviceEnrollmentInstance) SetOrgPhone(v string) {
	o.OrgPhone = &v
}

// GetOrgAddress returns the OrgAddress field value if set, zero value otherwise.
func (o *DeviceEnrollmentInstance) GetOrgAddress() string {
	if o == nil || IsNil(o.OrgAddress) {
		var ret string
		return ret
	}
	return *o.OrgAddress
}

// GetOrgAddressOk returns a tuple with the OrgAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentInstance) GetOrgAddressOk() (*string, bool) {
	if o == nil || IsNil(o.OrgAddress) {
		return nil, false
	}
	return o.OrgAddress, true
}

// HasOrgAddress returns a boolean if a field has been set.
func (o *DeviceEnrollmentInstance) HasOrgAddress() bool {
	if o != nil && !IsNil(o.OrgAddress) {
		return true
	}

	return false
}

// SetOrgAddress gets a reference to the given string and assigns it to the OrgAddress field.
func (o *DeviceEnrollmentInstance) SetOrgAddress(v string) {
	o.OrgAddress = &v
}

// GetTokenExpirationDate returns the TokenExpirationDate field value if set, zero value otherwise.
func (o *DeviceEnrollmentInstance) GetTokenExpirationDate() string {
	if o == nil || IsNil(o.TokenExpirationDate) {
		var ret string
		return ret
	}
	return *o.TokenExpirationDate
}

// GetTokenExpirationDateOk returns a tuple with the TokenExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentInstance) GetTokenExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.TokenExpirationDate) {
		return nil, false
	}
	return o.TokenExpirationDate, true
}

// HasTokenExpirationDate returns a boolean if a field has been set.
func (o *DeviceEnrollmentInstance) HasTokenExpirationDate() bool {
	if o != nil && !IsNil(o.TokenExpirationDate) {
		return true
	}

	return false
}

// SetTokenExpirationDate gets a reference to the given string and assigns it to the TokenExpirationDate field.
func (o *DeviceEnrollmentInstance) SetTokenExpirationDate(v string) {
	o.TokenExpirationDate = &v
}

func (o DeviceEnrollmentInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceEnrollmentInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.SupervisionIdentityId) {
		toSerialize["supervisionIdentityId"] = o.SupervisionIdentityId
	}
	if !IsNil(o.SiteId) {
		toSerialize["siteId"] = o.SiteId
	}
	if !IsNil(o.ServerName) {
		toSerialize["serverName"] = o.ServerName
	}
	if !IsNil(o.ServerUuid) {
		toSerialize["serverUuid"] = o.ServerUuid
	}
	if !IsNil(o.AdminId) {
		toSerialize["adminId"] = o.AdminId
	}
	if !IsNil(o.OrgName) {
		toSerialize["orgName"] = o.OrgName
	}
	if !IsNil(o.OrgEmail) {
		toSerialize["orgEmail"] = o.OrgEmail
	}
	if !IsNil(o.OrgPhone) {
		toSerialize["orgPhone"] = o.OrgPhone
	}
	if !IsNil(o.OrgAddress) {
		toSerialize["orgAddress"] = o.OrgAddress
	}
	if !IsNil(o.TokenExpirationDate) {
		toSerialize["tokenExpirationDate"] = o.TokenExpirationDate
	}
	return toSerialize, nil
}

func (o *DeviceEnrollmentInstance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varDeviceEnrollmentInstance := _DeviceEnrollmentInstance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeviceEnrollmentInstance)

	if err != nil {
		return err
	}

	*o = DeviceEnrollmentInstance(varDeviceEnrollmentInstance)

	return err
}

type NullableDeviceEnrollmentInstance struct {
	value *DeviceEnrollmentInstance
	isSet bool
}

func (v NullableDeviceEnrollmentInstance) Get() *DeviceEnrollmentInstance {
	return v.value
}

func (v *NullableDeviceEnrollmentInstance) Set(val *DeviceEnrollmentInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceEnrollmentInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceEnrollmentInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceEnrollmentInstance(val *DeviceEnrollmentInstance) *NullableDeviceEnrollmentInstance {
	return &NullableDeviceEnrollmentInstance{value: val, isSet: true}
}

func (v NullableDeviceEnrollmentInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceEnrollmentInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


