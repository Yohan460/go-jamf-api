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

// checks if the EnrollmentAccessGroupPreview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnrollmentAccessGroupPreview{}

// EnrollmentAccessGroupPreview struct for EnrollmentAccessGroupPreview
type EnrollmentAccessGroupPreview struct {
	// Autogenerated ID
	Id *string `json:"id,omitempty"`
	// LDAP Group ID
	GroupId string `json:"groupId"`
	LdapServerId string `json:"ldapServerId"`
	Name string `json:"name"`
	SiteId *string `json:"siteId,omitempty"`
	EnterpriseEnrollmentEnabled *bool `json:"enterpriseEnrollmentEnabled,omitempty"`
	PersonalEnrollmentEnabled *bool `json:"personalEnrollmentEnabled,omitempty"`
	AccountDrivenUserEnrollmentEnabled *bool `json:"accountDrivenUserEnrollmentEnabled,omitempty"`
	RequireEula *bool `json:"requireEula,omitempty"`
}

type _EnrollmentAccessGroupPreview EnrollmentAccessGroupPreview

// NewEnrollmentAccessGroupPreview instantiates a new EnrollmentAccessGroupPreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrollmentAccessGroupPreview(groupId string, ldapServerId string, name string) *EnrollmentAccessGroupPreview {
	this := EnrollmentAccessGroupPreview{}
	this.GroupId = groupId
	this.LdapServerId = ldapServerId
	this.Name = name
	return &this
}

// NewEnrollmentAccessGroupPreviewWithDefaults instantiates a new EnrollmentAccessGroupPreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrollmentAccessGroupPreviewWithDefaults() *EnrollmentAccessGroupPreview {
	this := EnrollmentAccessGroupPreview{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnrollmentAccessGroupPreview) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentAccessGroupPreview) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnrollmentAccessGroupPreview) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EnrollmentAccessGroupPreview) SetId(v string) {
	o.Id = &v
}

// GetGroupId returns the GroupId field value
func (o *EnrollmentAccessGroupPreview) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *EnrollmentAccessGroupPreview) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *EnrollmentAccessGroupPreview) SetGroupId(v string) {
	o.GroupId = v
}

// GetLdapServerId returns the LdapServerId field value
func (o *EnrollmentAccessGroupPreview) GetLdapServerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LdapServerId
}

// GetLdapServerIdOk returns a tuple with the LdapServerId field value
// and a boolean to check if the value has been set.
func (o *EnrollmentAccessGroupPreview) GetLdapServerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LdapServerId, true
}

// SetLdapServerId sets field value
func (o *EnrollmentAccessGroupPreview) SetLdapServerId(v string) {
	o.LdapServerId = v
}

// GetName returns the Name field value
func (o *EnrollmentAccessGroupPreview) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnrollmentAccessGroupPreview) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EnrollmentAccessGroupPreview) SetName(v string) {
	o.Name = v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *EnrollmentAccessGroupPreview) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentAccessGroupPreview) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *EnrollmentAccessGroupPreview) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *EnrollmentAccessGroupPreview) SetSiteId(v string) {
	o.SiteId = &v
}

// GetEnterpriseEnrollmentEnabled returns the EnterpriseEnrollmentEnabled field value if set, zero value otherwise.
func (o *EnrollmentAccessGroupPreview) GetEnterpriseEnrollmentEnabled() bool {
	if o == nil || IsNil(o.EnterpriseEnrollmentEnabled) {
		var ret bool
		return ret
	}
	return *o.EnterpriseEnrollmentEnabled
}

// GetEnterpriseEnrollmentEnabledOk returns a tuple with the EnterpriseEnrollmentEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentAccessGroupPreview) GetEnterpriseEnrollmentEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EnterpriseEnrollmentEnabled) {
		return nil, false
	}
	return o.EnterpriseEnrollmentEnabled, true
}

// HasEnterpriseEnrollmentEnabled returns a boolean if a field has been set.
func (o *EnrollmentAccessGroupPreview) HasEnterpriseEnrollmentEnabled() bool {
	if o != nil && !IsNil(o.EnterpriseEnrollmentEnabled) {
		return true
	}

	return false
}

// SetEnterpriseEnrollmentEnabled gets a reference to the given bool and assigns it to the EnterpriseEnrollmentEnabled field.
func (o *EnrollmentAccessGroupPreview) SetEnterpriseEnrollmentEnabled(v bool) {
	o.EnterpriseEnrollmentEnabled = &v
}

// GetPersonalEnrollmentEnabled returns the PersonalEnrollmentEnabled field value if set, zero value otherwise.
func (o *EnrollmentAccessGroupPreview) GetPersonalEnrollmentEnabled() bool {
	if o == nil || IsNil(o.PersonalEnrollmentEnabled) {
		var ret bool
		return ret
	}
	return *o.PersonalEnrollmentEnabled
}

// GetPersonalEnrollmentEnabledOk returns a tuple with the PersonalEnrollmentEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentAccessGroupPreview) GetPersonalEnrollmentEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PersonalEnrollmentEnabled) {
		return nil, false
	}
	return o.PersonalEnrollmentEnabled, true
}

// HasPersonalEnrollmentEnabled returns a boolean if a field has been set.
func (o *EnrollmentAccessGroupPreview) HasPersonalEnrollmentEnabled() bool {
	if o != nil && !IsNil(o.PersonalEnrollmentEnabled) {
		return true
	}

	return false
}

// SetPersonalEnrollmentEnabled gets a reference to the given bool and assigns it to the PersonalEnrollmentEnabled field.
func (o *EnrollmentAccessGroupPreview) SetPersonalEnrollmentEnabled(v bool) {
	o.PersonalEnrollmentEnabled = &v
}

// GetAccountDrivenUserEnrollmentEnabled returns the AccountDrivenUserEnrollmentEnabled field value if set, zero value otherwise.
func (o *EnrollmentAccessGroupPreview) GetAccountDrivenUserEnrollmentEnabled() bool {
	if o == nil || IsNil(o.AccountDrivenUserEnrollmentEnabled) {
		var ret bool
		return ret
	}
	return *o.AccountDrivenUserEnrollmentEnabled
}

// GetAccountDrivenUserEnrollmentEnabledOk returns a tuple with the AccountDrivenUserEnrollmentEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentAccessGroupPreview) GetAccountDrivenUserEnrollmentEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AccountDrivenUserEnrollmentEnabled) {
		return nil, false
	}
	return o.AccountDrivenUserEnrollmentEnabled, true
}

// HasAccountDrivenUserEnrollmentEnabled returns a boolean if a field has been set.
func (o *EnrollmentAccessGroupPreview) HasAccountDrivenUserEnrollmentEnabled() bool {
	if o != nil && !IsNil(o.AccountDrivenUserEnrollmentEnabled) {
		return true
	}

	return false
}

// SetAccountDrivenUserEnrollmentEnabled gets a reference to the given bool and assigns it to the AccountDrivenUserEnrollmentEnabled field.
func (o *EnrollmentAccessGroupPreview) SetAccountDrivenUserEnrollmentEnabled(v bool) {
	o.AccountDrivenUserEnrollmentEnabled = &v
}

// GetRequireEula returns the RequireEula field value if set, zero value otherwise.
func (o *EnrollmentAccessGroupPreview) GetRequireEula() bool {
	if o == nil || IsNil(o.RequireEula) {
		var ret bool
		return ret
	}
	return *o.RequireEula
}

// GetRequireEulaOk returns a tuple with the RequireEula field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentAccessGroupPreview) GetRequireEulaOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireEula) {
		return nil, false
	}
	return o.RequireEula, true
}

// HasRequireEula returns a boolean if a field has been set.
func (o *EnrollmentAccessGroupPreview) HasRequireEula() bool {
	if o != nil && !IsNil(o.RequireEula) {
		return true
	}

	return false
}

// SetRequireEula gets a reference to the given bool and assigns it to the RequireEula field.
func (o *EnrollmentAccessGroupPreview) SetRequireEula(v bool) {
	o.RequireEula = &v
}

func (o EnrollmentAccessGroupPreview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnrollmentAccessGroupPreview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["groupId"] = o.GroupId
	toSerialize["ldapServerId"] = o.LdapServerId
	toSerialize["name"] = o.Name
	if !IsNil(o.SiteId) {
		toSerialize["siteId"] = o.SiteId
	}
	if !IsNil(o.EnterpriseEnrollmentEnabled) {
		toSerialize["enterpriseEnrollmentEnabled"] = o.EnterpriseEnrollmentEnabled
	}
	if !IsNil(o.PersonalEnrollmentEnabled) {
		toSerialize["personalEnrollmentEnabled"] = o.PersonalEnrollmentEnabled
	}
	if !IsNil(o.AccountDrivenUserEnrollmentEnabled) {
		toSerialize["accountDrivenUserEnrollmentEnabled"] = o.AccountDrivenUserEnrollmentEnabled
	}
	if !IsNil(o.RequireEula) {
		toSerialize["requireEula"] = o.RequireEula
	}
	return toSerialize, nil
}

func (o *EnrollmentAccessGroupPreview) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"groupId",
		"ldapServerId",
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

	varEnrollmentAccessGroupPreview := _EnrollmentAccessGroupPreview{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEnrollmentAccessGroupPreview)

	if err != nil {
		return err
	}

	*o = EnrollmentAccessGroupPreview(varEnrollmentAccessGroupPreview)

	return err
}

type NullableEnrollmentAccessGroupPreview struct {
	value *EnrollmentAccessGroupPreview
	isSet bool
}

func (v NullableEnrollmentAccessGroupPreview) Get() *EnrollmentAccessGroupPreview {
	return v.value
}

func (v *NullableEnrollmentAccessGroupPreview) Set(val *EnrollmentAccessGroupPreview) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollmentAccessGroupPreview) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollmentAccessGroupPreview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollmentAccessGroupPreview(val *EnrollmentAccessGroupPreview) *NullableEnrollmentAccessGroupPreview {
	return &NullableEnrollmentAccessGroupPreview{value: val, isSet: true}
}

func (v NullableEnrollmentAccessGroupPreview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrollmentAccessGroupPreview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


