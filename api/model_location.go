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

// Location struct for Location
type Location struct {
	Username *string `json:"username,omitempty"`
	RealName *string `json:"realName,omitempty"`
	EmailAddress *string `json:"emailAddress,omitempty"`
	Position *string `json:"position,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Department *IdAndName `json:"department,omitempty"`
	Building *IdAndName `json:"building,omitempty"`
	Room *string `json:"room,omitempty"`
}

// NewLocation instantiates a new Location object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocation() *Location {
	this := Location{}
	return &this
}

// NewLocationWithDefaults instantiates a new Location object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationWithDefaults() *Location {
	this := Location{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *Location) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *Location) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *Location) SetUsername(v string) {
	o.Username = &v
}

// GetRealName returns the RealName field value if set, zero value otherwise.
func (o *Location) GetRealName() string {
	if o == nil || o.RealName == nil {
		var ret string
		return ret
	}
	return *o.RealName
}

// GetRealNameOk returns a tuple with the RealName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetRealNameOk() (*string, bool) {
	if o == nil || o.RealName == nil {
		return nil, false
	}
	return o.RealName, true
}

// HasRealName returns a boolean if a field has been set.
func (o *Location) HasRealName() bool {
	if o != nil && o.RealName != nil {
		return true
	}

	return false
}

// SetRealName gets a reference to the given string and assigns it to the RealName field.
func (o *Location) SetRealName(v string) {
	o.RealName = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *Location) GetEmailAddress() string {
	if o == nil || o.EmailAddress == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetEmailAddressOk() (*string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *Location) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *Location) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *Location) GetPosition() string {
	if o == nil || o.Position == nil {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetPositionOk() (*string, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *Location) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *Location) SetPosition(v string) {
	o.Position = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *Location) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *Location) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *Location) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *Location) GetDepartment() IdAndName {
	if o == nil || o.Department == nil {
		var ret IdAndName
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetDepartmentOk() (*IdAndName, bool) {
	if o == nil || o.Department == nil {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *Location) HasDepartment() bool {
	if o != nil && o.Department != nil {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given IdAndName and assigns it to the Department field.
func (o *Location) SetDepartment(v IdAndName) {
	o.Department = &v
}

// GetBuilding returns the Building field value if set, zero value otherwise.
func (o *Location) GetBuilding() IdAndName {
	if o == nil || o.Building == nil {
		var ret IdAndName
		return ret
	}
	return *o.Building
}

// GetBuildingOk returns a tuple with the Building field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetBuildingOk() (*IdAndName, bool) {
	if o == nil || o.Building == nil {
		return nil, false
	}
	return o.Building, true
}

// HasBuilding returns a boolean if a field has been set.
func (o *Location) HasBuilding() bool {
	if o != nil && o.Building != nil {
		return true
	}

	return false
}

// SetBuilding gets a reference to the given IdAndName and assigns it to the Building field.
func (o *Location) SetBuilding(v IdAndName) {
	o.Building = &v
}

// GetRoom returns the Room field value if set, zero value otherwise.
func (o *Location) GetRoom() string {
	if o == nil || o.Room == nil {
		var ret string
		return ret
	}
	return *o.Room
}

// GetRoomOk returns a tuple with the Room field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetRoomOk() (*string, bool) {
	if o == nil || o.Room == nil {
		return nil, false
	}
	return o.Room, true
}

// HasRoom returns a boolean if a field has been set.
func (o *Location) HasRoom() bool {
	if o != nil && o.Room != nil {
		return true
	}

	return false
}

// SetRoom gets a reference to the given string and assigns it to the Room field.
func (o *Location) SetRoom(v string) {
	o.Room = &v
}

func (o Location) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.RealName != nil {
		toSerialize["realName"] = o.RealName
	}
	if o.EmailAddress != nil {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.PhoneNumber != nil {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if o.Department != nil {
		toSerialize["department"] = o.Department
	}
	if o.Building != nil {
		toSerialize["building"] = o.Building
	}
	if o.Room != nil {
		toSerialize["room"] = o.Room
	}
	return json.Marshal(toSerialize)
}

type NullableLocation struct {
	value *Location
	isSet bool
}

func (v NullableLocation) Get() *Location {
	return v.value
}

func (v *NullableLocation) Set(val *Location) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation(val *Location) *NullableLocation {
	return &NullableLocation{value: val, isSet: true}
}

func (v NullableLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

