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

// UserMappings Cloud Identity Provider user mappings configuration
type UserMappings struct {
	ObjectClassLimitation string `json:"objectClassLimitation"`
	ObjectClasses string `json:"objectClasses"`
	SearchBase string `json:"searchBase"`
	SearchScope string `json:"searchScope"`
	AdditionalSearchBase *string `json:"additionalSearchBase,omitempty"`
	UserID string `json:"userID"`
	Username string `json:"username"`
	RealName string `json:"realName"`
	EmailAddress string `json:"emailAddress"`
	Department string `json:"department"`
	Building string `json:"building"`
	Room string `json:"room"`
	Phone string `json:"phone"`
	Position string `json:"position"`
	UserUuid string `json:"userUuid"`
}

// NewUserMappings instantiates a new UserMappings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMappings(objectClassLimitation string, objectClasses string, searchBase string, searchScope string, userID string, username string, realName string, emailAddress string, department string, building string, room string, phone string, position string, userUuid string) *UserMappings {
	this := UserMappings{}
	this.ObjectClassLimitation = objectClassLimitation
	this.ObjectClasses = objectClasses
	this.SearchBase = searchBase
	this.SearchScope = searchScope
	var additionalSearchBase string = ""
	this.AdditionalSearchBase = &additionalSearchBase
	this.UserID = userID
	this.Username = username
	this.RealName = realName
	this.EmailAddress = emailAddress
	this.Department = department
	this.Building = building
	this.Room = room
	this.Phone = phone
	this.Position = position
	this.UserUuid = userUuid
	return &this
}

// NewUserMappingsWithDefaults instantiates a new UserMappings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMappingsWithDefaults() *UserMappings {
	this := UserMappings{}
	var additionalSearchBase string = ""
	this.AdditionalSearchBase = &additionalSearchBase
	var building string = ""
	this.Building = building
	var room string = ""
	this.Room = room
	var phone string = ""
	this.Phone = phone
	return &this
}

// GetObjectClassLimitation returns the ObjectClassLimitation field value
func (o *UserMappings) GetObjectClassLimitation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectClassLimitation
}

// GetObjectClassLimitationOk returns a tuple with the ObjectClassLimitation field value
// and a boolean to check if the value has been set.
func (o *UserMappings) GetObjectClassLimitationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectClassLimitation, true
}

// SetObjectClassLimitation sets field value
func (o *UserMappings) SetObjectClassLimitation(v string) {
	o.ObjectClassLimitation = v
}

// GetObjectClasses returns the ObjectClasses field value
func (o *UserMappings) GetObjectClasses() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectClasses
}

// GetObjectClassesOk returns a tuple with the ObjectClasses field value
// and a boolean to check if the value has been set.
func (o *UserMappings) GetObjectClassesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectClasses, true
}

// SetObjectClasses sets field value
func (o *UserMappings) SetObjectClasses(v string) {
	o.ObjectClasses = v
}

// GetSearchBase returns the SearchBase field value
func (o *UserMappings) GetSearchBase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SearchBase
}

// GetSearchBaseOk returns a tuple with the SearchBase field value
// and a boolean to check if the value has been set.
func (o *UserMappings) GetSearchBaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchBase, true
}

// SetSearchBase sets field value
func (o *UserMappings) SetSearchBase(v string) {
	o.SearchBase = v
}

// GetSearchScope returns the SearchScope field value
func (o *UserMappings) GetSearchScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SearchScope
}

// GetSearchScopeOk returns a tuple with the SearchScope field value
// and a boolean to check if the value has been set.
func (o *UserMappings) GetSearchScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchScope, true
}

// SetSearchScope sets field value
func (o *UserMappings) SetSearchScope(v string) {
	o.SearchScope = v
}

// GetAdditionalSearchBase returns the AdditionalSearchBase field value if set, zero value otherwise.
func (o *UserMappings) GetAdditionalSearchBase() string {
	if o == nil || o.AdditionalSearchBase == nil {
		var ret string
		return ret
	}
	return *o.AdditionalSearchBase
}

// GetAdditionalSearchBaseOk returns a tuple with the AdditionalSearchBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMappings) GetAdditionalSearchBaseOk() (*string, bool) {
	if o == nil || o.AdditionalSearchBase == nil {
		return nil, false
	}
	return o.AdditionalSearchBase, true
}

// HasAdditionalSearchBase returns a boolean if a field has been set.
func (o *UserMappings) HasAdditionalSearchBase() bool {
	if o != nil && o.AdditionalSearchBase != nil {
		return true
	}

	return false
}

// SetAdditionalSearchBase gets a reference to the given string and assigns it to the AdditionalSearchBase field.
func (o *UserMappings) SetAdditionalSearchBase(v string) {
	o.AdditionalSearchBase = &v
}

// GetUserID returns the UserID field value
func (o *UserMappings) GetUserID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value
// and a boolean to check if the value has been set.
func (o *UserMappings) GetUserIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserID, true
}

// SetUserID sets field value
func (o *UserMappings) SetUserID(v string) {
	o.UserID = v
}

// GetUsername returns the Username field value
func (o *UserMappings) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UserMappings) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UserMappings) SetUsername(v string) {
	o.Username = v
}

// GetRealName returns the RealName field value
func (o *UserMappings) GetRealName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RealName
}

// GetRealNameOk returns a tuple with the RealName field value
// and a boolean to check if the value has been set.
func (o *UserMappings) GetRealNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RealName, true
}

// SetRealName sets field value
func (o *UserMappings) SetRealName(v string) {
	o.RealName = v
}

// GetEmailAddress returns the EmailAddress field value
func (o *UserMappings) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *UserMappings) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *UserMappings) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetDepartment returns the Department field value
func (o *UserMappings) GetDepartment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Department
}

// GetDepartmentOk returns a tuple with the Department field value
// and a boolean to check if the value has been set.
func (o *UserMappings) GetDepartmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Department, true
}

// SetDepartment sets field value
func (o *UserMappings) SetDepartment(v string) {
	o.Department = v
}

// GetBuilding returns the Building field value
func (o *UserMappings) GetBuilding() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Building
}

// GetBuildingOk returns a tuple with the Building field value
// and a boolean to check if the value has been set.
func (o *UserMappings) GetBuildingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Building, true
}

// SetBuilding sets field value
func (o *UserMappings) SetBuilding(v string) {
	o.Building = v
}

// GetRoom returns the Room field value
func (o *UserMappings) GetRoom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Room
}

// GetRoomOk returns a tuple with the Room field value
// and a boolean to check if the value has been set.
func (o *UserMappings) GetRoomOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Room, true
}

// SetRoom sets field value
func (o *UserMappings) SetRoom(v string) {
	o.Room = v
}

// GetPhone returns the Phone field value
func (o *UserMappings) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *UserMappings) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *UserMappings) SetPhone(v string) {
	o.Phone = v
}

// GetPosition returns the Position field value
func (o *UserMappings) GetPosition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *UserMappings) GetPositionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *UserMappings) SetPosition(v string) {
	o.Position = v
}

// GetUserUuid returns the UserUuid field value
func (o *UserMappings) GetUserUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value
// and a boolean to check if the value has been set.
func (o *UserMappings) GetUserUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserUuid, true
}

// SetUserUuid sets field value
func (o *UserMappings) SetUserUuid(v string) {
	o.UserUuid = v
}

func (o UserMappings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objectClassLimitation"] = o.ObjectClassLimitation
	}
	if true {
		toSerialize["objectClasses"] = o.ObjectClasses
	}
	if true {
		toSerialize["searchBase"] = o.SearchBase
	}
	if true {
		toSerialize["searchScope"] = o.SearchScope
	}
	if o.AdditionalSearchBase != nil {
		toSerialize["additionalSearchBase"] = o.AdditionalSearchBase
	}
	if true {
		toSerialize["userID"] = o.UserID
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["realName"] = o.RealName
	}
	if true {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if true {
		toSerialize["department"] = o.Department
	}
	if true {
		toSerialize["building"] = o.Building
	}
	if true {
		toSerialize["room"] = o.Room
	}
	if true {
		toSerialize["phone"] = o.Phone
	}
	if true {
		toSerialize["position"] = o.Position
	}
	if true {
		toSerialize["userUuid"] = o.UserUuid
	}
	return json.Marshal(toSerialize)
}

type NullableUserMappings struct {
	value *UserMappings
	isSet bool
}

func (v NullableUserMappings) Get() *UserMappings {
	return v.value
}

func (v *NullableUserMappings) Set(val *UserMappings) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMappings) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMappings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMappings(val *UserMappings) *NullableUserMappings {
	return &NullableUserMappings{value: val, isSet: true}
}

func (v NullableUserMappings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMappings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


