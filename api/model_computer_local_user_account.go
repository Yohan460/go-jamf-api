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

// checks if the ComputerLocalUserAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComputerLocalUserAccount{}

// ComputerLocalUserAccount struct for ComputerLocalUserAccount
type ComputerLocalUserAccount struct {
	Uid *string `json:"uid,omitempty"`
	UserGuid *string `json:"userGuid,omitempty"`
	Username *string `json:"username,omitempty"`
	FullName *string `json:"fullName,omitempty"`
	Admin *bool `json:"admin,omitempty"`
	HomeDirectory *string `json:"homeDirectory,omitempty"`
	// Home directory size in MB.
	HomeDirectorySizeMb *int64 `json:"homeDirectorySizeMb,omitempty"`
	FileVault2Enabled *bool `json:"fileVault2Enabled,omitempty"`
	UserAccountType *string `json:"userAccountType,omitempty"`
	PasswordMinLength *int64 `json:"passwordMinLength,omitempty"`
	PasswordMaxAge *int64 `json:"passwordMaxAge,omitempty"`
	PasswordMinComplexCharacters *int64 `json:"passwordMinComplexCharacters,omitempty"`
	PasswordHistoryDepth *int64 `json:"passwordHistoryDepth,omitempty"`
	PasswordRequireAlphanumeric *bool `json:"passwordRequireAlphanumeric,omitempty"`
	ComputerAzureActiveDirectoryId *string `json:"computerAzureActiveDirectoryId,omitempty"`
	UserAzureActiveDirectoryId *string `json:"userAzureActiveDirectoryId,omitempty"`
	AzureActiveDirectoryId *string `json:"azureActiveDirectoryId,omitempty"`
}

// NewComputerLocalUserAccount instantiates a new ComputerLocalUserAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputerLocalUserAccount() *ComputerLocalUserAccount {
	this := ComputerLocalUserAccount{}
	return &this
}

// NewComputerLocalUserAccountWithDefaults instantiates a new ComputerLocalUserAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputerLocalUserAccountWithDefaults() *ComputerLocalUserAccount {
	this := ComputerLocalUserAccount{}
	return &this
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *ComputerLocalUserAccount) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerLocalUserAccount) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *ComputerLocalUserAccount) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *ComputerLocalUserAccount) SetUid(v string) {
	o.Uid = &v
}

// GetUserGuid returns the UserGuid field value if set, zero value otherwise.
func (o *ComputerLocalUserAccount) GetUserGuid() string {
	if o == nil || IsNil(o.UserGuid) {
		var ret string
		return ret
	}
	return *o.UserGuid
}

// GetUserGuidOk returns a tuple with the UserGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerLocalUserAccount) GetUserGuidOk() (*string, bool) {
	if o == nil || IsNil(o.UserGuid) {
		return nil, false
	}
	return o.UserGuid, true
}

// HasUserGuid returns a boolean if a field has been set.
func (o *ComputerLocalUserAccount) HasUserGuid() bool {
	if o != nil && !IsNil(o.UserGuid) {
		return true
	}

	return false
}

// SetUserGuid gets a reference to the given string and assigns it to the UserGuid field.
func (o *ComputerLocalUserAccount) SetUserGuid(v string) {
	o.UserGuid = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ComputerLocalUserAccount) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerLocalUserAccount) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ComputerLocalUserAccount) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ComputerLocalUserAccount) SetUsername(v string) {
	o.Username = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *ComputerLocalUserAccount) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerLocalUserAccount) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *ComputerLocalUserAccount) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *ComputerLocalUserAccount) SetFullName(v string) {
	o.FullName = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *ComputerLocalUserAccount) GetAdmin() bool {
	if o == nil || IsNil(o.Admin) {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerLocalUserAccount) GetAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.Admin) {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *ComputerLocalUserAccount) HasAdmin() bool {
	if o != nil && !IsNil(o.Admin) {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *ComputerLocalUserAccount) SetAdmin(v bool) {
	o.Admin = &v
}

// GetHomeDirectory returns the HomeDirectory field value if set, zero value otherwise.
func (o *ComputerLocalUserAccount) GetHomeDirectory() string {
	if o == nil || IsNil(o.HomeDirectory) {
		var ret string
		return ret
	}
	return *o.HomeDirectory
}

// GetHomeDirectoryOk returns a tuple with the HomeDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerLocalUserAccount) GetHomeDirectoryOk() (*string, bool) {
	if o == nil || IsNil(o.HomeDirectory) {
		return nil, false
	}
	return o.HomeDirectory, true
}

// HasHomeDirectory returns a boolean if a field has been set.
func (o *ComputerLocalUserAccount) HasHomeDirectory() bool {
	if o != nil && !IsNil(o.HomeDirectory) {
		return true
	}

	return false
}

// SetHomeDirectory gets a reference to the given string and assigns it to the HomeDirectory field.
func (o *ComputerLocalUserAccount) SetHomeDirectory(v string) {
	o.HomeDirectory = &v
}

// GetHomeDirectorySizeMb returns the HomeDirectorySizeMb field value if set, zero value otherwise.
func (o *ComputerLocalUserAccount) GetHomeDirectorySizeMb() int64 {
	if o == nil || IsNil(o.HomeDirectorySizeMb) {
		var ret int64
		return ret
	}
	return *o.HomeDirectorySizeMb
}

// GetHomeDirectorySizeMbOk returns a tuple with the HomeDirectorySizeMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerLocalUserAccount) GetHomeDirectorySizeMbOk() (*int64, bool) {
	if o == nil || IsNil(o.HomeDirectorySizeMb) {
		return nil, false
	}
	return o.HomeDirectorySizeMb, true
}

// HasHomeDirectorySizeMb returns a boolean if a field has been set.
func (o *ComputerLocalUserAccount) HasHomeDirectorySizeMb() bool {
	if o != nil && !IsNil(o.HomeDirectorySizeMb) {
		return true
	}

	return false
}

// SetHomeDirectorySizeMb gets a reference to the given int64 and assigns it to the HomeDirectorySizeMb field.
func (o *ComputerLocalUserAccount) SetHomeDirectorySizeMb(v int64) {
	o.HomeDirectorySizeMb = &v
}

// GetFileVault2Enabled returns the FileVault2Enabled field value if set, zero value otherwise.
func (o *ComputerLocalUserAccount) GetFileVault2Enabled() bool {
	if o == nil || IsNil(o.FileVault2Enabled) {
		var ret bool
		return ret
	}
	return *o.FileVault2Enabled
}

// GetFileVault2EnabledOk returns a tuple with the FileVault2Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerLocalUserAccount) GetFileVault2EnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FileVault2Enabled) {
		return nil, false
	}
	return o.FileVault2Enabled, true
}

// HasFileVault2Enabled returns a boolean if a field has been set.
func (o *ComputerLocalUserAccount) HasFileVault2Enabled() bool {
	if o != nil && !IsNil(o.FileVault2Enabled) {
		return true
	}

	return false
}

// SetFileVault2Enabled gets a reference to the given bool and assigns it to the FileVault2Enabled field.
func (o *ComputerLocalUserAccount) SetFileVault2Enabled(v bool) {
	o.FileVault2Enabled = &v
}

// GetUserAccountType returns the UserAccountType field value if set, zero value otherwise.
func (o *ComputerLocalUserAccount) GetUserAccountType() string {
	if o == nil || IsNil(o.UserAccountType) {
		var ret string
		return ret
	}
	return *o.UserAccountType
}

// GetUserAccountTypeOk returns a tuple with the UserAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerLocalUserAccount) GetUserAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UserAccountType) {
		return nil, false
	}
	return o.UserAccountType, true
}

// HasUserAccountType returns a boolean if a field has been set.
func (o *ComputerLocalUserAccount) HasUserAccountType() bool {
	if o != nil && !IsNil(o.UserAccountType) {
		return true
	}

	return false
}

// SetUserAccountType gets a reference to the given string and assigns it to the UserAccountType field.
func (o *ComputerLocalUserAccount) SetUserAccountType(v string) {
	o.UserAccountType = &v
}

// GetPasswordMinLength returns the PasswordMinLength field value if set, zero value otherwise.
func (o *ComputerLocalUserAccount) GetPasswordMinLength() int64 {
	if o == nil || IsNil(o.PasswordMinLength) {
		var ret int64
		return ret
	}
	return *o.PasswordMinLength
}

// GetPasswordMinLengthOk returns a tuple with the PasswordMinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerLocalUserAccount) GetPasswordMinLengthOk() (*int64, bool) {
	if o == nil || IsNil(o.PasswordMinLength) {
		return nil, false
	}
	return o.PasswordMinLength, true
}

// HasPasswordMinLength returns a boolean if a field has been set.
func (o *ComputerLocalUserAccount) HasPasswordMinLength() bool {
	if o != nil && !IsNil(o.PasswordMinLength) {
		return true
	}

	return false
}

// SetPasswordMinLength gets a reference to the given int64 and assigns it to the PasswordMinLength field.
func (o *ComputerLocalUserAccount) SetPasswordMinLength(v int64) {
	o.PasswordMinLength = &v
}

// GetPasswordMaxAge returns the PasswordMaxAge field value if set, zero value otherwise.
func (o *ComputerLocalUserAccount) GetPasswordMaxAge() int64 {
	if o == nil || IsNil(o.PasswordMaxAge) {
		var ret int64
		return ret
	}
	return *o.PasswordMaxAge
}

// GetPasswordMaxAgeOk returns a tuple with the PasswordMaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerLocalUserAccount) GetPasswordMaxAgeOk() (*int64, bool) {
	if o == nil || IsNil(o.PasswordMaxAge) {
		return nil, false
	}
	return o.PasswordMaxAge, true
}

// HasPasswordMaxAge returns a boolean if a field has been set.
func (o *ComputerLocalUserAccount) HasPasswordMaxAge() bool {
	if o != nil && !IsNil(o.PasswordMaxAge) {
		return true
	}

	return false
}

// SetPasswordMaxAge gets a reference to the given int64 and assigns it to the PasswordMaxAge field.
func (o *ComputerLocalUserAccount) SetPasswordMaxAge(v int64) {
	o.PasswordMaxAge = &v
}

// GetPasswordMinComplexCharacters returns the PasswordMinComplexCharacters field value if set, zero value otherwise.
func (o *ComputerLocalUserAccount) GetPasswordMinComplexCharacters() int64 {
	if o == nil || IsNil(o.PasswordMinComplexCharacters) {
		var ret int64
		return ret
	}
	return *o.PasswordMinComplexCharacters
}

// GetPasswordMinComplexCharactersOk returns a tuple with the PasswordMinComplexCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerLocalUserAccount) GetPasswordMinComplexCharactersOk() (*int64, bool) {
	if o == nil || IsNil(o.PasswordMinComplexCharacters) {
		return nil, false
	}
	return o.PasswordMinComplexCharacters, true
}

// HasPasswordMinComplexCharacters returns a boolean if a field has been set.
func (o *ComputerLocalUserAccount) HasPasswordMinComplexCharacters() bool {
	if o != nil && !IsNil(o.PasswordMinComplexCharacters) {
		return true
	}

	return false
}

// SetPasswordMinComplexCharacters gets a reference to the given int64 and assigns it to the PasswordMinComplexCharacters field.
func (o *ComputerLocalUserAccount) SetPasswordMinComplexCharacters(v int64) {
	o.PasswordMinComplexCharacters = &v
}

// GetPasswordHistoryDepth returns the PasswordHistoryDepth field value if set, zero value otherwise.
func (o *ComputerLocalUserAccount) GetPasswordHistoryDepth() int64 {
	if o == nil || IsNil(o.PasswordHistoryDepth) {
		var ret int64
		return ret
	}
	return *o.PasswordHistoryDepth
}

// GetPasswordHistoryDepthOk returns a tuple with the PasswordHistoryDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerLocalUserAccount) GetPasswordHistoryDepthOk() (*int64, bool) {
	if o == nil || IsNil(o.PasswordHistoryDepth) {
		return nil, false
	}
	return o.PasswordHistoryDepth, true
}

// HasPasswordHistoryDepth returns a boolean if a field has been set.
func (o *ComputerLocalUserAccount) HasPasswordHistoryDepth() bool {
	if o != nil && !IsNil(o.PasswordHistoryDepth) {
		return true
	}

	return false
}

// SetPasswordHistoryDepth gets a reference to the given int64 and assigns it to the PasswordHistoryDepth field.
func (o *ComputerLocalUserAccount) SetPasswordHistoryDepth(v int64) {
	o.PasswordHistoryDepth = &v
}

// GetPasswordRequireAlphanumeric returns the PasswordRequireAlphanumeric field value if set, zero value otherwise.
func (o *ComputerLocalUserAccount) GetPasswordRequireAlphanumeric() bool {
	if o == nil || IsNil(o.PasswordRequireAlphanumeric) {
		var ret bool
		return ret
	}
	return *o.PasswordRequireAlphanumeric
}

// GetPasswordRequireAlphanumericOk returns a tuple with the PasswordRequireAlphanumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerLocalUserAccount) GetPasswordRequireAlphanumericOk() (*bool, bool) {
	if o == nil || IsNil(o.PasswordRequireAlphanumeric) {
		return nil, false
	}
	return o.PasswordRequireAlphanumeric, true
}

// HasPasswordRequireAlphanumeric returns a boolean if a field has been set.
func (o *ComputerLocalUserAccount) HasPasswordRequireAlphanumeric() bool {
	if o != nil && !IsNil(o.PasswordRequireAlphanumeric) {
		return true
	}

	return false
}

// SetPasswordRequireAlphanumeric gets a reference to the given bool and assigns it to the PasswordRequireAlphanumeric field.
func (o *ComputerLocalUserAccount) SetPasswordRequireAlphanumeric(v bool) {
	o.PasswordRequireAlphanumeric = &v
}

// GetComputerAzureActiveDirectoryId returns the ComputerAzureActiveDirectoryId field value if set, zero value otherwise.
func (o *ComputerLocalUserAccount) GetComputerAzureActiveDirectoryId() string {
	if o == nil || IsNil(o.ComputerAzureActiveDirectoryId) {
		var ret string
		return ret
	}
	return *o.ComputerAzureActiveDirectoryId
}

// GetComputerAzureActiveDirectoryIdOk returns a tuple with the ComputerAzureActiveDirectoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerLocalUserAccount) GetComputerAzureActiveDirectoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.ComputerAzureActiveDirectoryId) {
		return nil, false
	}
	return o.ComputerAzureActiveDirectoryId, true
}

// HasComputerAzureActiveDirectoryId returns a boolean if a field has been set.
func (o *ComputerLocalUserAccount) HasComputerAzureActiveDirectoryId() bool {
	if o != nil && !IsNil(o.ComputerAzureActiveDirectoryId) {
		return true
	}

	return false
}

// SetComputerAzureActiveDirectoryId gets a reference to the given string and assigns it to the ComputerAzureActiveDirectoryId field.
func (o *ComputerLocalUserAccount) SetComputerAzureActiveDirectoryId(v string) {
	o.ComputerAzureActiveDirectoryId = &v
}

// GetUserAzureActiveDirectoryId returns the UserAzureActiveDirectoryId field value if set, zero value otherwise.
func (o *ComputerLocalUserAccount) GetUserAzureActiveDirectoryId() string {
	if o == nil || IsNil(o.UserAzureActiveDirectoryId) {
		var ret string
		return ret
	}
	return *o.UserAzureActiveDirectoryId
}

// GetUserAzureActiveDirectoryIdOk returns a tuple with the UserAzureActiveDirectoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerLocalUserAccount) GetUserAzureActiveDirectoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserAzureActiveDirectoryId) {
		return nil, false
	}
	return o.UserAzureActiveDirectoryId, true
}

// HasUserAzureActiveDirectoryId returns a boolean if a field has been set.
func (o *ComputerLocalUserAccount) HasUserAzureActiveDirectoryId() bool {
	if o != nil && !IsNil(o.UserAzureActiveDirectoryId) {
		return true
	}

	return false
}

// SetUserAzureActiveDirectoryId gets a reference to the given string and assigns it to the UserAzureActiveDirectoryId field.
func (o *ComputerLocalUserAccount) SetUserAzureActiveDirectoryId(v string) {
	o.UserAzureActiveDirectoryId = &v
}

// GetAzureActiveDirectoryId returns the AzureActiveDirectoryId field value if set, zero value otherwise.
func (o *ComputerLocalUserAccount) GetAzureActiveDirectoryId() string {
	if o == nil || IsNil(o.AzureActiveDirectoryId) {
		var ret string
		return ret
	}
	return *o.AzureActiveDirectoryId
}

// GetAzureActiveDirectoryIdOk returns a tuple with the AzureActiveDirectoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerLocalUserAccount) GetAzureActiveDirectoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.AzureActiveDirectoryId) {
		return nil, false
	}
	return o.AzureActiveDirectoryId, true
}

// HasAzureActiveDirectoryId returns a boolean if a field has been set.
func (o *ComputerLocalUserAccount) HasAzureActiveDirectoryId() bool {
	if o != nil && !IsNil(o.AzureActiveDirectoryId) {
		return true
	}

	return false
}

// SetAzureActiveDirectoryId gets a reference to the given string and assigns it to the AzureActiveDirectoryId field.
func (o *ComputerLocalUserAccount) SetAzureActiveDirectoryId(v string) {
	o.AzureActiveDirectoryId = &v
}

func (o ComputerLocalUserAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComputerLocalUserAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	if !IsNil(o.UserGuid) {
		toSerialize["userGuid"] = o.UserGuid
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.FullName) {
		toSerialize["fullName"] = o.FullName
	}
	if !IsNil(o.Admin) {
		toSerialize["admin"] = o.Admin
	}
	if !IsNil(o.HomeDirectory) {
		toSerialize["homeDirectory"] = o.HomeDirectory
	}
	if !IsNil(o.HomeDirectorySizeMb) {
		toSerialize["homeDirectorySizeMb"] = o.HomeDirectorySizeMb
	}
	if !IsNil(o.FileVault2Enabled) {
		toSerialize["fileVault2Enabled"] = o.FileVault2Enabled
	}
	if !IsNil(o.UserAccountType) {
		toSerialize["userAccountType"] = o.UserAccountType
	}
	if !IsNil(o.PasswordMinLength) {
		toSerialize["passwordMinLength"] = o.PasswordMinLength
	}
	if !IsNil(o.PasswordMaxAge) {
		toSerialize["passwordMaxAge"] = o.PasswordMaxAge
	}
	if !IsNil(o.PasswordMinComplexCharacters) {
		toSerialize["passwordMinComplexCharacters"] = o.PasswordMinComplexCharacters
	}
	if !IsNil(o.PasswordHistoryDepth) {
		toSerialize["passwordHistoryDepth"] = o.PasswordHistoryDepth
	}
	if !IsNil(o.PasswordRequireAlphanumeric) {
		toSerialize["passwordRequireAlphanumeric"] = o.PasswordRequireAlphanumeric
	}
	if !IsNil(o.ComputerAzureActiveDirectoryId) {
		toSerialize["computerAzureActiveDirectoryId"] = o.ComputerAzureActiveDirectoryId
	}
	if !IsNil(o.UserAzureActiveDirectoryId) {
		toSerialize["userAzureActiveDirectoryId"] = o.UserAzureActiveDirectoryId
	}
	if !IsNil(o.AzureActiveDirectoryId) {
		toSerialize["azureActiveDirectoryId"] = o.AzureActiveDirectoryId
	}
	return toSerialize, nil
}

type NullableComputerLocalUserAccount struct {
	value *ComputerLocalUserAccount
	isSet bool
}

func (v NullableComputerLocalUserAccount) Get() *ComputerLocalUserAccount {
	return v.value
}

func (v *NullableComputerLocalUserAccount) Set(val *ComputerLocalUserAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerLocalUserAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerLocalUserAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerLocalUserAccount(val *ComputerLocalUserAccount) *NullableComputerLocalUserAccount {
	return &NullableComputerLocalUserAccount{value: val, isSet: true}
}

func (v NullableComputerLocalUserAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerLocalUserAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


