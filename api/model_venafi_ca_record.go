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

// VenafiCaRecord struct for VenafiCaRecord
type VenafiCaRecord struct {
	Id *int32 `json:"id,omitempty"`
	Name string `json:"name"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	PasswordConfigured *bool `json:"passwordConfigured,omitempty"`
	ProxyAddress *string `json:"proxyAddress,omitempty"`
	RevocationEnabled *bool `json:"revocationEnabled,omitempty"`
	ClientId *string `json:"clientId,omitempty"`
	RefreshToken *string `json:"refreshToken,omitempty"`
	RefreshTokenConfigured *bool `json:"refreshTokenConfigured,omitempty"`
}

// NewVenafiCaRecord instantiates a new VenafiCaRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVenafiCaRecord(name string) *VenafiCaRecord {
	this := VenafiCaRecord{}
	this.Name = name
	return &this
}

// NewVenafiCaRecordWithDefaults instantiates a new VenafiCaRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVenafiCaRecordWithDefaults() *VenafiCaRecord {
	this := VenafiCaRecord{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VenafiCaRecord) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VenafiCaRecord) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VenafiCaRecord) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *VenafiCaRecord) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *VenafiCaRecord) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VenafiCaRecord) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VenafiCaRecord) SetName(v string) {
	o.Name = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *VenafiCaRecord) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VenafiCaRecord) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *VenafiCaRecord) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *VenafiCaRecord) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *VenafiCaRecord) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VenafiCaRecord) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *VenafiCaRecord) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *VenafiCaRecord) SetPassword(v string) {
	o.Password = &v
}

// GetPasswordConfigured returns the PasswordConfigured field value if set, zero value otherwise.
func (o *VenafiCaRecord) GetPasswordConfigured() bool {
	if o == nil || o.PasswordConfigured == nil {
		var ret bool
		return ret
	}
	return *o.PasswordConfigured
}

// GetPasswordConfiguredOk returns a tuple with the PasswordConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VenafiCaRecord) GetPasswordConfiguredOk() (*bool, bool) {
	if o == nil || o.PasswordConfigured == nil {
		return nil, false
	}
	return o.PasswordConfigured, true
}

// HasPasswordConfigured returns a boolean if a field has been set.
func (o *VenafiCaRecord) HasPasswordConfigured() bool {
	if o != nil && o.PasswordConfigured != nil {
		return true
	}

	return false
}

// SetPasswordConfigured gets a reference to the given bool and assigns it to the PasswordConfigured field.
func (o *VenafiCaRecord) SetPasswordConfigured(v bool) {
	o.PasswordConfigured = &v
}

// GetProxyAddress returns the ProxyAddress field value if set, zero value otherwise.
func (o *VenafiCaRecord) GetProxyAddress() string {
	if o == nil || o.ProxyAddress == nil {
		var ret string
		return ret
	}
	return *o.ProxyAddress
}

// GetProxyAddressOk returns a tuple with the ProxyAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VenafiCaRecord) GetProxyAddressOk() (*string, bool) {
	if o == nil || o.ProxyAddress == nil {
		return nil, false
	}
	return o.ProxyAddress, true
}

// HasProxyAddress returns a boolean if a field has been set.
func (o *VenafiCaRecord) HasProxyAddress() bool {
	if o != nil && o.ProxyAddress != nil {
		return true
	}

	return false
}

// SetProxyAddress gets a reference to the given string and assigns it to the ProxyAddress field.
func (o *VenafiCaRecord) SetProxyAddress(v string) {
	o.ProxyAddress = &v
}

// GetRevocationEnabled returns the RevocationEnabled field value if set, zero value otherwise.
func (o *VenafiCaRecord) GetRevocationEnabled() bool {
	if o == nil || o.RevocationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RevocationEnabled
}

// GetRevocationEnabledOk returns a tuple with the RevocationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VenafiCaRecord) GetRevocationEnabledOk() (*bool, bool) {
	if o == nil || o.RevocationEnabled == nil {
		return nil, false
	}
	return o.RevocationEnabled, true
}

// HasRevocationEnabled returns a boolean if a field has been set.
func (o *VenafiCaRecord) HasRevocationEnabled() bool {
	if o != nil && o.RevocationEnabled != nil {
		return true
	}

	return false
}

// SetRevocationEnabled gets a reference to the given bool and assigns it to the RevocationEnabled field.
func (o *VenafiCaRecord) SetRevocationEnabled(v bool) {
	o.RevocationEnabled = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *VenafiCaRecord) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VenafiCaRecord) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *VenafiCaRecord) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *VenafiCaRecord) SetClientId(v string) {
	o.ClientId = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *VenafiCaRecord) GetRefreshToken() string {
	if o == nil || o.RefreshToken == nil {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VenafiCaRecord) GetRefreshTokenOk() (*string, bool) {
	if o == nil || o.RefreshToken == nil {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *VenafiCaRecord) HasRefreshToken() bool {
	if o != nil && o.RefreshToken != nil {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *VenafiCaRecord) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetRefreshTokenConfigured returns the RefreshTokenConfigured field value if set, zero value otherwise.
func (o *VenafiCaRecord) GetRefreshTokenConfigured() bool {
	if o == nil || o.RefreshTokenConfigured == nil {
		var ret bool
		return ret
	}
	return *o.RefreshTokenConfigured
}

// GetRefreshTokenConfiguredOk returns a tuple with the RefreshTokenConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VenafiCaRecord) GetRefreshTokenConfiguredOk() (*bool, bool) {
	if o == nil || o.RefreshTokenConfigured == nil {
		return nil, false
	}
	return o.RefreshTokenConfigured, true
}

// HasRefreshTokenConfigured returns a boolean if a field has been set.
func (o *VenafiCaRecord) HasRefreshTokenConfigured() bool {
	if o != nil && o.RefreshTokenConfigured != nil {
		return true
	}

	return false
}

// SetRefreshTokenConfigured gets a reference to the given bool and assigns it to the RefreshTokenConfigured field.
func (o *VenafiCaRecord) SetRefreshTokenConfigured(v bool) {
	o.RefreshTokenConfigured = &v
}

func (o VenafiCaRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.PasswordConfigured != nil {
		toSerialize["passwordConfigured"] = o.PasswordConfigured
	}
	if o.ProxyAddress != nil {
		toSerialize["proxyAddress"] = o.ProxyAddress
	}
	if o.RevocationEnabled != nil {
		toSerialize["revocationEnabled"] = o.RevocationEnabled
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.RefreshToken != nil {
		toSerialize["refreshToken"] = o.RefreshToken
	}
	if o.RefreshTokenConfigured != nil {
		toSerialize["refreshTokenConfigured"] = o.RefreshTokenConfigured
	}
	return json.Marshal(toSerialize)
}

type NullableVenafiCaRecord struct {
	value *VenafiCaRecord
	isSet bool
}

func (v NullableVenafiCaRecord) Get() *VenafiCaRecord {
	return v.value
}

func (v *NullableVenafiCaRecord) Set(val *VenafiCaRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableVenafiCaRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableVenafiCaRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVenafiCaRecord(val *VenafiCaRecord) *NullableVenafiCaRecord {
	return &NullableVenafiCaRecord{value: val, isSet: true}
}

func (v NullableVenafiCaRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVenafiCaRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


