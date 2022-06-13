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

// GetEnrollmentCustomizationPanelLdapAuth struct for GetEnrollmentCustomizationPanelLdapAuth
type GetEnrollmentCustomizationPanelLdapAuth struct {
	DisplayName string `json:"displayName"`
	Rank int32 `json:"rank"`
	UsernameLabel string `json:"usernameLabel"`
	PasswordLabel string `json:"passwordLabel"`
	Title string `json:"title"`
	BackButtonText string `json:"backButtonText"`
	ContinueButtonText string `json:"continueButtonText"`
	LdapGroupAccess []EnrollmentCustomizationLdapGroupAccess `json:"ldapGroupAccess,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewGetEnrollmentCustomizationPanelLdapAuth instantiates a new GetEnrollmentCustomizationPanelLdapAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEnrollmentCustomizationPanelLdapAuth(displayName string, rank int32, usernameLabel string, passwordLabel string, title string, backButtonText string, continueButtonText string) *GetEnrollmentCustomizationPanelLdapAuth {
	this := GetEnrollmentCustomizationPanelLdapAuth{}
	this.DisplayName = displayName
	this.Rank = rank
	this.UsernameLabel = usernameLabel
	this.PasswordLabel = passwordLabel
	this.Title = title
	this.BackButtonText = backButtonText
	this.ContinueButtonText = continueButtonText
	return &this
}

// NewGetEnrollmentCustomizationPanelLdapAuthWithDefaults instantiates a new GetEnrollmentCustomizationPanelLdapAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEnrollmentCustomizationPanelLdapAuthWithDefaults() *GetEnrollmentCustomizationPanelLdapAuth {
	this := GetEnrollmentCustomizationPanelLdapAuth{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *GetEnrollmentCustomizationPanelLdapAuth) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *GetEnrollmentCustomizationPanelLdapAuth) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *GetEnrollmentCustomizationPanelLdapAuth) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetRank returns the Rank field value
func (o *GetEnrollmentCustomizationPanelLdapAuth) GetRank() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rank
}

// GetRankOk returns a tuple with the Rank field value
// and a boolean to check if the value has been set.
func (o *GetEnrollmentCustomizationPanelLdapAuth) GetRankOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rank, true
}

// SetRank sets field value
func (o *GetEnrollmentCustomizationPanelLdapAuth) SetRank(v int32) {
	o.Rank = v
}

// GetUsernameLabel returns the UsernameLabel field value
func (o *GetEnrollmentCustomizationPanelLdapAuth) GetUsernameLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UsernameLabel
}

// GetUsernameLabelOk returns a tuple with the UsernameLabel field value
// and a boolean to check if the value has been set.
func (o *GetEnrollmentCustomizationPanelLdapAuth) GetUsernameLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsernameLabel, true
}

// SetUsernameLabel sets field value
func (o *GetEnrollmentCustomizationPanelLdapAuth) SetUsernameLabel(v string) {
	o.UsernameLabel = v
}

// GetPasswordLabel returns the PasswordLabel field value
func (o *GetEnrollmentCustomizationPanelLdapAuth) GetPasswordLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PasswordLabel
}

// GetPasswordLabelOk returns a tuple with the PasswordLabel field value
// and a boolean to check if the value has been set.
func (o *GetEnrollmentCustomizationPanelLdapAuth) GetPasswordLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordLabel, true
}

// SetPasswordLabel sets field value
func (o *GetEnrollmentCustomizationPanelLdapAuth) SetPasswordLabel(v string) {
	o.PasswordLabel = v
}

// GetTitle returns the Title field value
func (o *GetEnrollmentCustomizationPanelLdapAuth) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *GetEnrollmentCustomizationPanelLdapAuth) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *GetEnrollmentCustomizationPanelLdapAuth) SetTitle(v string) {
	o.Title = v
}

// GetBackButtonText returns the BackButtonText field value
func (o *GetEnrollmentCustomizationPanelLdapAuth) GetBackButtonText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BackButtonText
}

// GetBackButtonTextOk returns a tuple with the BackButtonText field value
// and a boolean to check if the value has been set.
func (o *GetEnrollmentCustomizationPanelLdapAuth) GetBackButtonTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackButtonText, true
}

// SetBackButtonText sets field value
func (o *GetEnrollmentCustomizationPanelLdapAuth) SetBackButtonText(v string) {
	o.BackButtonText = v
}

// GetContinueButtonText returns the ContinueButtonText field value
func (o *GetEnrollmentCustomizationPanelLdapAuth) GetContinueButtonText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContinueButtonText
}

// GetContinueButtonTextOk returns a tuple with the ContinueButtonText field value
// and a boolean to check if the value has been set.
func (o *GetEnrollmentCustomizationPanelLdapAuth) GetContinueButtonTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContinueButtonText, true
}

// SetContinueButtonText sets field value
func (o *GetEnrollmentCustomizationPanelLdapAuth) SetContinueButtonText(v string) {
	o.ContinueButtonText = v
}

// GetLdapGroupAccess returns the LdapGroupAccess field value if set, zero value otherwise.
func (o *GetEnrollmentCustomizationPanelLdapAuth) GetLdapGroupAccess() []EnrollmentCustomizationLdapGroupAccess {
	if o == nil || o.LdapGroupAccess == nil {
		var ret []EnrollmentCustomizationLdapGroupAccess
		return ret
	}
	return o.LdapGroupAccess
}

// GetLdapGroupAccessOk returns a tuple with the LdapGroupAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEnrollmentCustomizationPanelLdapAuth) GetLdapGroupAccessOk() ([]EnrollmentCustomizationLdapGroupAccess, bool) {
	if o == nil || o.LdapGroupAccess == nil {
		return nil, false
	}
	return o.LdapGroupAccess, true
}

// HasLdapGroupAccess returns a boolean if a field has been set.
func (o *GetEnrollmentCustomizationPanelLdapAuth) HasLdapGroupAccess() bool {
	if o != nil && o.LdapGroupAccess != nil {
		return true
	}

	return false
}

// SetLdapGroupAccess gets a reference to the given []EnrollmentCustomizationLdapGroupAccess and assigns it to the LdapGroupAccess field.
func (o *GetEnrollmentCustomizationPanelLdapAuth) SetLdapGroupAccess(v []EnrollmentCustomizationLdapGroupAccess) {
	o.LdapGroupAccess = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetEnrollmentCustomizationPanelLdapAuth) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEnrollmentCustomizationPanelLdapAuth) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetEnrollmentCustomizationPanelLdapAuth) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetEnrollmentCustomizationPanelLdapAuth) SetId(v int32) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetEnrollmentCustomizationPanelLdapAuth) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEnrollmentCustomizationPanelLdapAuth) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetEnrollmentCustomizationPanelLdapAuth) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetEnrollmentCustomizationPanelLdapAuth) SetType(v string) {
	o.Type = &v
}

func (o GetEnrollmentCustomizationPanelLdapAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["rank"] = o.Rank
	}
	if true {
		toSerialize["usernameLabel"] = o.UsernameLabel
	}
	if true {
		toSerialize["passwordLabel"] = o.PasswordLabel
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["backButtonText"] = o.BackButtonText
	}
	if true {
		toSerialize["continueButtonText"] = o.ContinueButtonText
	}
	if o.LdapGroupAccess != nil {
		toSerialize["ldapGroupAccess"] = o.LdapGroupAccess
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableGetEnrollmentCustomizationPanelLdapAuth struct {
	value *GetEnrollmentCustomizationPanelLdapAuth
	isSet bool
}

func (v NullableGetEnrollmentCustomizationPanelLdapAuth) Get() *GetEnrollmentCustomizationPanelLdapAuth {
	return v.value
}

func (v *NullableGetEnrollmentCustomizationPanelLdapAuth) Set(val *GetEnrollmentCustomizationPanelLdapAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEnrollmentCustomizationPanelLdapAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEnrollmentCustomizationPanelLdapAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEnrollmentCustomizationPanelLdapAuth(val *GetEnrollmentCustomizationPanelLdapAuth) *NullableGetEnrollmentCustomizationPanelLdapAuth {
	return &NullableGetEnrollmentCustomizationPanelLdapAuth{value: val, isSet: true}
}

func (v NullableGetEnrollmentCustomizationPanelLdapAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEnrollmentCustomizationPanelLdapAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


