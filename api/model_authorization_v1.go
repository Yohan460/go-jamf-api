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

// checks if the AuthorizationV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationV1{}

// AuthorizationV1 struct for AuthorizationV1
type AuthorizationV1 struct {
	Account *AuthAccountV1 `json:"account,omitempty"`
	AccountGroups []AccountGroup `json:"accountGroups,omitempty"`
	Sites []V1Site `json:"sites,omitempty"`
	AuthenticationType *AuthenticationType `json:"authenticationType,omitempty"`
}

// NewAuthorizationV1 instantiates a new AuthorizationV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationV1() *AuthorizationV1 {
	this := AuthorizationV1{}
	return &this
}

// NewAuthorizationV1WithDefaults instantiates a new AuthorizationV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationV1WithDefaults() *AuthorizationV1 {
	this := AuthorizationV1{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AuthorizationV1) GetAccount() AuthAccountV1 {
	if o == nil || IsNil(o.Account) {
		var ret AuthAccountV1
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationV1) GetAccountOk() (*AuthAccountV1, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AuthorizationV1) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given AuthAccountV1 and assigns it to the Account field.
func (o *AuthorizationV1) SetAccount(v AuthAccountV1) {
	o.Account = &v
}

// GetAccountGroups returns the AccountGroups field value if set, zero value otherwise.
func (o *AuthorizationV1) GetAccountGroups() []AccountGroup {
	if o == nil || IsNil(o.AccountGroups) {
		var ret []AccountGroup
		return ret
	}
	return o.AccountGroups
}

// GetAccountGroupsOk returns a tuple with the AccountGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationV1) GetAccountGroupsOk() ([]AccountGroup, bool) {
	if o == nil || IsNil(o.AccountGroups) {
		return nil, false
	}
	return o.AccountGroups, true
}

// HasAccountGroups returns a boolean if a field has been set.
func (o *AuthorizationV1) HasAccountGroups() bool {
	if o != nil && !IsNil(o.AccountGroups) {
		return true
	}

	return false
}

// SetAccountGroups gets a reference to the given []AccountGroup and assigns it to the AccountGroups field.
func (o *AuthorizationV1) SetAccountGroups(v []AccountGroup) {
	o.AccountGroups = v
}

// GetSites returns the Sites field value if set, zero value otherwise.
func (o *AuthorizationV1) GetSites() []V1Site {
	if o == nil || IsNil(o.Sites) {
		var ret []V1Site
		return ret
	}
	return o.Sites
}

// GetSitesOk returns a tuple with the Sites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationV1) GetSitesOk() ([]V1Site, bool) {
	if o == nil || IsNil(o.Sites) {
		return nil, false
	}
	return o.Sites, true
}

// HasSites returns a boolean if a field has been set.
func (o *AuthorizationV1) HasSites() bool {
	if o != nil && !IsNil(o.Sites) {
		return true
	}

	return false
}

// SetSites gets a reference to the given []V1Site and assigns it to the Sites field.
func (o *AuthorizationV1) SetSites(v []V1Site) {
	o.Sites = v
}

// GetAuthenticationType returns the AuthenticationType field value if set, zero value otherwise.
func (o *AuthorizationV1) GetAuthenticationType() AuthenticationType {
	if o == nil || IsNil(o.AuthenticationType) {
		var ret AuthenticationType
		return ret
	}
	return *o.AuthenticationType
}

// GetAuthenticationTypeOk returns a tuple with the AuthenticationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationV1) GetAuthenticationTypeOk() (*AuthenticationType, bool) {
	if o == nil || IsNil(o.AuthenticationType) {
		return nil, false
	}
	return o.AuthenticationType, true
}

// HasAuthenticationType returns a boolean if a field has been set.
func (o *AuthorizationV1) HasAuthenticationType() bool {
	if o != nil && !IsNil(o.AuthenticationType) {
		return true
	}

	return false
}

// SetAuthenticationType gets a reference to the given AuthenticationType and assigns it to the AuthenticationType field.
func (o *AuthorizationV1) SetAuthenticationType(v AuthenticationType) {
	o.AuthenticationType = &v
}

func (o AuthorizationV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.AccountGroups) {
		toSerialize["accountGroups"] = o.AccountGroups
	}
	if !IsNil(o.Sites) {
		toSerialize["sites"] = o.Sites
	}
	if !IsNil(o.AuthenticationType) {
		toSerialize["authenticationType"] = o.AuthenticationType
	}
	return toSerialize, nil
}

type NullableAuthorizationV1 struct {
	value *AuthorizationV1
	isSet bool
}

func (v NullableAuthorizationV1) Get() *AuthorizationV1 {
	return v.value
}

func (v *NullableAuthorizationV1) Set(val *AuthorizationV1) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationV1) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationV1(val *AuthorizationV1) *NullableAuthorizationV1 {
	return &NullableAuthorizationV1{value: val, isSet: true}
}

func (v NullableAuthorizationV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


