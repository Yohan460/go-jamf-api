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

// checks if the AccountSettingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountSettingsResponse{}

// AccountSettingsResponse struct for AccountSettingsResponse
type AccountSettingsResponse struct {
	// id of Account Settings
	Id *string `json:"id,omitempty"`
	PayloadConfigured *bool `json:"payloadConfigured,omitempty"`
	LocalAdminAccountEnabled *bool `json:"localAdminAccountEnabled,omitempty"`
	AdminUsername *string `json:"adminUsername,omitempty"`
	HiddenAdminAccount *bool `json:"hiddenAdminAccount,omitempty"`
	LocalUserManaged *bool `json:"localUserManaged,omitempty"`
	UserAccountType *string `json:"userAccountType,omitempty"`
	VersionLock *int64 `json:"versionLock,omitempty"`
	PrefillPrimaryAccountInfoFeatureEnabled *bool `json:"prefillPrimaryAccountInfoFeatureEnabled,omitempty"`
	// Values accepted are only CUSTOM and DEVICE_OWNER
	PrefillType *string `json:"prefillType,omitempty"`
	PrefillAccountFullName *string `json:"prefillAccountFullName,omitempty"`
	PrefillAccountUserName *string `json:"prefillAccountUserName,omitempty"`
	PreventPrefillInfoFromModification *bool `json:"preventPrefillInfoFromModification,omitempty"`
}

// NewAccountSettingsResponse instantiates a new AccountSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountSettingsResponse() *AccountSettingsResponse {
	this := AccountSettingsResponse{}
	var payloadConfigured bool = false
	this.PayloadConfigured = &payloadConfigured
	var localAdminAccountEnabled bool = false
	this.LocalAdminAccountEnabled = &localAdminAccountEnabled
	var adminUsername string = ""
	this.AdminUsername = &adminUsername
	var hiddenAdminAccount bool = false
	this.HiddenAdminAccount = &hiddenAdminAccount
	var localUserManaged bool = false
	this.LocalUserManaged = &localUserManaged
	var userAccountType string = "STANDARD"
	this.UserAccountType = &userAccountType
	var versionLock int64 = 0
	this.VersionLock = &versionLock
	var prefillPrimaryAccountInfoFeatureEnabled bool = false
	this.PrefillPrimaryAccountInfoFeatureEnabled = &prefillPrimaryAccountInfoFeatureEnabled
	var prefillType string = "CUSTOM"
	this.PrefillType = &prefillType
	var prefillAccountFullName string = ""
	this.PrefillAccountFullName = &prefillAccountFullName
	var prefillAccountUserName string = ""
	this.PrefillAccountUserName = &prefillAccountUserName
	var preventPrefillInfoFromModification bool = false
	this.PreventPrefillInfoFromModification = &preventPrefillInfoFromModification
	return &this
}

// NewAccountSettingsResponseWithDefaults instantiates a new AccountSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountSettingsResponseWithDefaults() *AccountSettingsResponse {
	this := AccountSettingsResponse{}
	var payloadConfigured bool = false
	this.PayloadConfigured = &payloadConfigured
	var localAdminAccountEnabled bool = false
	this.LocalAdminAccountEnabled = &localAdminAccountEnabled
	var adminUsername string = ""
	this.AdminUsername = &adminUsername
	var hiddenAdminAccount bool = false
	this.HiddenAdminAccount = &hiddenAdminAccount
	var localUserManaged bool = false
	this.LocalUserManaged = &localUserManaged
	var userAccountType string = "STANDARD"
	this.UserAccountType = &userAccountType
	var versionLock int64 = 0
	this.VersionLock = &versionLock
	var prefillPrimaryAccountInfoFeatureEnabled bool = false
	this.PrefillPrimaryAccountInfoFeatureEnabled = &prefillPrimaryAccountInfoFeatureEnabled
	var prefillType string = "CUSTOM"
	this.PrefillType = &prefillType
	var prefillAccountFullName string = ""
	this.PrefillAccountFullName = &prefillAccountFullName
	var prefillAccountUserName string = ""
	this.PrefillAccountUserName = &prefillAccountUserName
	var preventPrefillInfoFromModification bool = false
	this.PreventPrefillInfoFromModification = &preventPrefillInfoFromModification
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountSettingsResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSettingsResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountSettingsResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountSettingsResponse) SetId(v string) {
	o.Id = &v
}

// GetPayloadConfigured returns the PayloadConfigured field value if set, zero value otherwise.
func (o *AccountSettingsResponse) GetPayloadConfigured() bool {
	if o == nil || IsNil(o.PayloadConfigured) {
		var ret bool
		return ret
	}
	return *o.PayloadConfigured
}

// GetPayloadConfiguredOk returns a tuple with the PayloadConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSettingsResponse) GetPayloadConfiguredOk() (*bool, bool) {
	if o == nil || IsNil(o.PayloadConfigured) {
		return nil, false
	}
	return o.PayloadConfigured, true
}

// HasPayloadConfigured returns a boolean if a field has been set.
func (o *AccountSettingsResponse) HasPayloadConfigured() bool {
	if o != nil && !IsNil(o.PayloadConfigured) {
		return true
	}

	return false
}

// SetPayloadConfigured gets a reference to the given bool and assigns it to the PayloadConfigured field.
func (o *AccountSettingsResponse) SetPayloadConfigured(v bool) {
	o.PayloadConfigured = &v
}

// GetLocalAdminAccountEnabled returns the LocalAdminAccountEnabled field value if set, zero value otherwise.
func (o *AccountSettingsResponse) GetLocalAdminAccountEnabled() bool {
	if o == nil || IsNil(o.LocalAdminAccountEnabled) {
		var ret bool
		return ret
	}
	return *o.LocalAdminAccountEnabled
}

// GetLocalAdminAccountEnabledOk returns a tuple with the LocalAdminAccountEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSettingsResponse) GetLocalAdminAccountEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LocalAdminAccountEnabled) {
		return nil, false
	}
	return o.LocalAdminAccountEnabled, true
}

// HasLocalAdminAccountEnabled returns a boolean if a field has been set.
func (o *AccountSettingsResponse) HasLocalAdminAccountEnabled() bool {
	if o != nil && !IsNil(o.LocalAdminAccountEnabled) {
		return true
	}

	return false
}

// SetLocalAdminAccountEnabled gets a reference to the given bool and assigns it to the LocalAdminAccountEnabled field.
func (o *AccountSettingsResponse) SetLocalAdminAccountEnabled(v bool) {
	o.LocalAdminAccountEnabled = &v
}

// GetAdminUsername returns the AdminUsername field value if set, zero value otherwise.
func (o *AccountSettingsResponse) GetAdminUsername() string {
	if o == nil || IsNil(o.AdminUsername) {
		var ret string
		return ret
	}
	return *o.AdminUsername
}

// GetAdminUsernameOk returns a tuple with the AdminUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSettingsResponse) GetAdminUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.AdminUsername) {
		return nil, false
	}
	return o.AdminUsername, true
}

// HasAdminUsername returns a boolean if a field has been set.
func (o *AccountSettingsResponse) HasAdminUsername() bool {
	if o != nil && !IsNil(o.AdminUsername) {
		return true
	}

	return false
}

// SetAdminUsername gets a reference to the given string and assigns it to the AdminUsername field.
func (o *AccountSettingsResponse) SetAdminUsername(v string) {
	o.AdminUsername = &v
}

// GetHiddenAdminAccount returns the HiddenAdminAccount field value if set, zero value otherwise.
func (o *AccountSettingsResponse) GetHiddenAdminAccount() bool {
	if o == nil || IsNil(o.HiddenAdminAccount) {
		var ret bool
		return ret
	}
	return *o.HiddenAdminAccount
}

// GetHiddenAdminAccountOk returns a tuple with the HiddenAdminAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSettingsResponse) GetHiddenAdminAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.HiddenAdminAccount) {
		return nil, false
	}
	return o.HiddenAdminAccount, true
}

// HasHiddenAdminAccount returns a boolean if a field has been set.
func (o *AccountSettingsResponse) HasHiddenAdminAccount() bool {
	if o != nil && !IsNil(o.HiddenAdminAccount) {
		return true
	}

	return false
}

// SetHiddenAdminAccount gets a reference to the given bool and assigns it to the HiddenAdminAccount field.
func (o *AccountSettingsResponse) SetHiddenAdminAccount(v bool) {
	o.HiddenAdminAccount = &v
}

// GetLocalUserManaged returns the LocalUserManaged field value if set, zero value otherwise.
func (o *AccountSettingsResponse) GetLocalUserManaged() bool {
	if o == nil || IsNil(o.LocalUserManaged) {
		var ret bool
		return ret
	}
	return *o.LocalUserManaged
}

// GetLocalUserManagedOk returns a tuple with the LocalUserManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSettingsResponse) GetLocalUserManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.LocalUserManaged) {
		return nil, false
	}
	return o.LocalUserManaged, true
}

// HasLocalUserManaged returns a boolean if a field has been set.
func (o *AccountSettingsResponse) HasLocalUserManaged() bool {
	if o != nil && !IsNil(o.LocalUserManaged) {
		return true
	}

	return false
}

// SetLocalUserManaged gets a reference to the given bool and assigns it to the LocalUserManaged field.
func (o *AccountSettingsResponse) SetLocalUserManaged(v bool) {
	o.LocalUserManaged = &v
}

// GetUserAccountType returns the UserAccountType field value if set, zero value otherwise.
func (o *AccountSettingsResponse) GetUserAccountType() string {
	if o == nil || IsNil(o.UserAccountType) {
		var ret string
		return ret
	}
	return *o.UserAccountType
}

// GetUserAccountTypeOk returns a tuple with the UserAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSettingsResponse) GetUserAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UserAccountType) {
		return nil, false
	}
	return o.UserAccountType, true
}

// HasUserAccountType returns a boolean if a field has been set.
func (o *AccountSettingsResponse) HasUserAccountType() bool {
	if o != nil && !IsNil(o.UserAccountType) {
		return true
	}

	return false
}

// SetUserAccountType gets a reference to the given string and assigns it to the UserAccountType field.
func (o *AccountSettingsResponse) SetUserAccountType(v string) {
	o.UserAccountType = &v
}

// GetVersionLock returns the VersionLock field value if set, zero value otherwise.
func (o *AccountSettingsResponse) GetVersionLock() int64 {
	if o == nil || IsNil(o.VersionLock) {
		var ret int64
		return ret
	}
	return *o.VersionLock
}

// GetVersionLockOk returns a tuple with the VersionLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSettingsResponse) GetVersionLockOk() (*int64, bool) {
	if o == nil || IsNil(o.VersionLock) {
		return nil, false
	}
	return o.VersionLock, true
}

// HasVersionLock returns a boolean if a field has been set.
func (o *AccountSettingsResponse) HasVersionLock() bool {
	if o != nil && !IsNil(o.VersionLock) {
		return true
	}

	return false
}

// SetVersionLock gets a reference to the given int64 and assigns it to the VersionLock field.
func (o *AccountSettingsResponse) SetVersionLock(v int64) {
	o.VersionLock = &v
}

// GetPrefillPrimaryAccountInfoFeatureEnabled returns the PrefillPrimaryAccountInfoFeatureEnabled field value if set, zero value otherwise.
func (o *AccountSettingsResponse) GetPrefillPrimaryAccountInfoFeatureEnabled() bool {
	if o == nil || IsNil(o.PrefillPrimaryAccountInfoFeatureEnabled) {
		var ret bool
		return ret
	}
	return *o.PrefillPrimaryAccountInfoFeatureEnabled
}

// GetPrefillPrimaryAccountInfoFeatureEnabledOk returns a tuple with the PrefillPrimaryAccountInfoFeatureEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSettingsResponse) GetPrefillPrimaryAccountInfoFeatureEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PrefillPrimaryAccountInfoFeatureEnabled) {
		return nil, false
	}
	return o.PrefillPrimaryAccountInfoFeatureEnabled, true
}

// HasPrefillPrimaryAccountInfoFeatureEnabled returns a boolean if a field has been set.
func (o *AccountSettingsResponse) HasPrefillPrimaryAccountInfoFeatureEnabled() bool {
	if o != nil && !IsNil(o.PrefillPrimaryAccountInfoFeatureEnabled) {
		return true
	}

	return false
}

// SetPrefillPrimaryAccountInfoFeatureEnabled gets a reference to the given bool and assigns it to the PrefillPrimaryAccountInfoFeatureEnabled field.
func (o *AccountSettingsResponse) SetPrefillPrimaryAccountInfoFeatureEnabled(v bool) {
	o.PrefillPrimaryAccountInfoFeatureEnabled = &v
}

// GetPrefillType returns the PrefillType field value if set, zero value otherwise.
func (o *AccountSettingsResponse) GetPrefillType() string {
	if o == nil || IsNil(o.PrefillType) {
		var ret string
		return ret
	}
	return *o.PrefillType
}

// GetPrefillTypeOk returns a tuple with the PrefillType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSettingsResponse) GetPrefillTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PrefillType) {
		return nil, false
	}
	return o.PrefillType, true
}

// HasPrefillType returns a boolean if a field has been set.
func (o *AccountSettingsResponse) HasPrefillType() bool {
	if o != nil && !IsNil(o.PrefillType) {
		return true
	}

	return false
}

// SetPrefillType gets a reference to the given string and assigns it to the PrefillType field.
func (o *AccountSettingsResponse) SetPrefillType(v string) {
	o.PrefillType = &v
}

// GetPrefillAccountFullName returns the PrefillAccountFullName field value if set, zero value otherwise.
func (o *AccountSettingsResponse) GetPrefillAccountFullName() string {
	if o == nil || IsNil(o.PrefillAccountFullName) {
		var ret string
		return ret
	}
	return *o.PrefillAccountFullName
}

// GetPrefillAccountFullNameOk returns a tuple with the PrefillAccountFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSettingsResponse) GetPrefillAccountFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.PrefillAccountFullName) {
		return nil, false
	}
	return o.PrefillAccountFullName, true
}

// HasPrefillAccountFullName returns a boolean if a field has been set.
func (o *AccountSettingsResponse) HasPrefillAccountFullName() bool {
	if o != nil && !IsNil(o.PrefillAccountFullName) {
		return true
	}

	return false
}

// SetPrefillAccountFullName gets a reference to the given string and assigns it to the PrefillAccountFullName field.
func (o *AccountSettingsResponse) SetPrefillAccountFullName(v string) {
	o.PrefillAccountFullName = &v
}

// GetPrefillAccountUserName returns the PrefillAccountUserName field value if set, zero value otherwise.
func (o *AccountSettingsResponse) GetPrefillAccountUserName() string {
	if o == nil || IsNil(o.PrefillAccountUserName) {
		var ret string
		return ret
	}
	return *o.PrefillAccountUserName
}

// GetPrefillAccountUserNameOk returns a tuple with the PrefillAccountUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSettingsResponse) GetPrefillAccountUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.PrefillAccountUserName) {
		return nil, false
	}
	return o.PrefillAccountUserName, true
}

// HasPrefillAccountUserName returns a boolean if a field has been set.
func (o *AccountSettingsResponse) HasPrefillAccountUserName() bool {
	if o != nil && !IsNil(o.PrefillAccountUserName) {
		return true
	}

	return false
}

// SetPrefillAccountUserName gets a reference to the given string and assigns it to the PrefillAccountUserName field.
func (o *AccountSettingsResponse) SetPrefillAccountUserName(v string) {
	o.PrefillAccountUserName = &v
}

// GetPreventPrefillInfoFromModification returns the PreventPrefillInfoFromModification field value if set, zero value otherwise.
func (o *AccountSettingsResponse) GetPreventPrefillInfoFromModification() bool {
	if o == nil || IsNil(o.PreventPrefillInfoFromModification) {
		var ret bool
		return ret
	}
	return *o.PreventPrefillInfoFromModification
}

// GetPreventPrefillInfoFromModificationOk returns a tuple with the PreventPrefillInfoFromModification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSettingsResponse) GetPreventPrefillInfoFromModificationOk() (*bool, bool) {
	if o == nil || IsNil(o.PreventPrefillInfoFromModification) {
		return nil, false
	}
	return o.PreventPrefillInfoFromModification, true
}

// HasPreventPrefillInfoFromModification returns a boolean if a field has been set.
func (o *AccountSettingsResponse) HasPreventPrefillInfoFromModification() bool {
	if o != nil && !IsNil(o.PreventPrefillInfoFromModification) {
		return true
	}

	return false
}

// SetPreventPrefillInfoFromModification gets a reference to the given bool and assigns it to the PreventPrefillInfoFromModification field.
func (o *AccountSettingsResponse) SetPreventPrefillInfoFromModification(v bool) {
	o.PreventPrefillInfoFromModification = &v
}

func (o AccountSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountSettingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PayloadConfigured) {
		toSerialize["payloadConfigured"] = o.PayloadConfigured
	}
	if !IsNil(o.LocalAdminAccountEnabled) {
		toSerialize["localAdminAccountEnabled"] = o.LocalAdminAccountEnabled
	}
	if !IsNil(o.AdminUsername) {
		toSerialize["adminUsername"] = o.AdminUsername
	}
	if !IsNil(o.HiddenAdminAccount) {
		toSerialize["hiddenAdminAccount"] = o.HiddenAdminAccount
	}
	if !IsNil(o.LocalUserManaged) {
		toSerialize["localUserManaged"] = o.LocalUserManaged
	}
	if !IsNil(o.UserAccountType) {
		toSerialize["userAccountType"] = o.UserAccountType
	}
	if !IsNil(o.VersionLock) {
		toSerialize["versionLock"] = o.VersionLock
	}
	if !IsNil(o.PrefillPrimaryAccountInfoFeatureEnabled) {
		toSerialize["prefillPrimaryAccountInfoFeatureEnabled"] = o.PrefillPrimaryAccountInfoFeatureEnabled
	}
	if !IsNil(o.PrefillType) {
		toSerialize["prefillType"] = o.PrefillType
	}
	if !IsNil(o.PrefillAccountFullName) {
		toSerialize["prefillAccountFullName"] = o.PrefillAccountFullName
	}
	if !IsNil(o.PrefillAccountUserName) {
		toSerialize["prefillAccountUserName"] = o.PrefillAccountUserName
	}
	if !IsNil(o.PreventPrefillInfoFromModification) {
		toSerialize["preventPrefillInfoFromModification"] = o.PreventPrefillInfoFromModification
	}
	return toSerialize, nil
}

type NullableAccountSettingsResponse struct {
	value *AccountSettingsResponse
	isSet bool
}

func (v NullableAccountSettingsResponse) Get() *AccountSettingsResponse {
	return v.value
}

func (v *NullableAccountSettingsResponse) Set(val *AccountSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountSettingsResponse(val *AccountSettingsResponse) *NullableAccountSettingsResponse {
	return &NullableAccountSettingsResponse{value: val, isSet: true}
}

func (v NullableAccountSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


