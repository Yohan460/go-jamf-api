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

// checks if the LocationInformationV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationInformationV2{}

// LocationInformationV2 struct for LocationInformationV2
type LocationInformationV2 struct {
	Username string `json:"username"`
	Realname string `json:"realname"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Room string `json:"room"`
	Position string `json:"position"`
	DepartmentId string `json:"departmentId"`
	BuildingId string `json:"buildingId"`
	Id string `json:"id"`
	VersionLock int32 `json:"versionLock"`
}

// NewLocationInformationV2 instantiates a new LocationInformationV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationInformationV2(username string, realname string, phone string, email string, room string, position string, departmentId string, buildingId string, id string, versionLock int32) *LocationInformationV2 {
	this := LocationInformationV2{}
	this.Username = username
	this.Realname = realname
	this.Phone = phone
	this.Email = email
	this.Room = room
	this.Position = position
	this.DepartmentId = departmentId
	this.BuildingId = buildingId
	this.Id = id
	this.VersionLock = versionLock
	return &this
}

// NewLocationInformationV2WithDefaults instantiates a new LocationInformationV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationInformationV2WithDefaults() *LocationInformationV2 {
	this := LocationInformationV2{}
	return &this
}

// GetUsername returns the Username field value
func (o *LocationInformationV2) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *LocationInformationV2) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *LocationInformationV2) SetUsername(v string) {
	o.Username = v
}

// GetRealname returns the Realname field value
func (o *LocationInformationV2) GetRealname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Realname
}

// GetRealnameOk returns a tuple with the Realname field value
// and a boolean to check if the value has been set.
func (o *LocationInformationV2) GetRealnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Realname, true
}

// SetRealname sets field value
func (o *LocationInformationV2) SetRealname(v string) {
	o.Realname = v
}

// GetPhone returns the Phone field value
func (o *LocationInformationV2) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *LocationInformationV2) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *LocationInformationV2) SetPhone(v string) {
	o.Phone = v
}

// GetEmail returns the Email field value
func (o *LocationInformationV2) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *LocationInformationV2) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *LocationInformationV2) SetEmail(v string) {
	o.Email = v
}

// GetRoom returns the Room field value
func (o *LocationInformationV2) GetRoom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Room
}

// GetRoomOk returns a tuple with the Room field value
// and a boolean to check if the value has been set.
func (o *LocationInformationV2) GetRoomOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Room, true
}

// SetRoom sets field value
func (o *LocationInformationV2) SetRoom(v string) {
	o.Room = v
}

// GetPosition returns the Position field value
func (o *LocationInformationV2) GetPosition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *LocationInformationV2) GetPositionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *LocationInformationV2) SetPosition(v string) {
	o.Position = v
}

// GetDepartmentId returns the DepartmentId field value
func (o *LocationInformationV2) GetDepartmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepartmentId
}

// GetDepartmentIdOk returns a tuple with the DepartmentId field value
// and a boolean to check if the value has been set.
func (o *LocationInformationV2) GetDepartmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepartmentId, true
}

// SetDepartmentId sets field value
func (o *LocationInformationV2) SetDepartmentId(v string) {
	o.DepartmentId = v
}

// GetBuildingId returns the BuildingId field value
func (o *LocationInformationV2) GetBuildingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuildingId
}

// GetBuildingIdOk returns a tuple with the BuildingId field value
// and a boolean to check if the value has been set.
func (o *LocationInformationV2) GetBuildingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildingId, true
}

// SetBuildingId sets field value
func (o *LocationInformationV2) SetBuildingId(v string) {
	o.BuildingId = v
}

// GetId returns the Id field value
func (o *LocationInformationV2) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LocationInformationV2) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LocationInformationV2) SetId(v string) {
	o.Id = v
}

// GetVersionLock returns the VersionLock field value
func (o *LocationInformationV2) GetVersionLock() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VersionLock
}

// GetVersionLockOk returns a tuple with the VersionLock field value
// and a boolean to check if the value has been set.
func (o *LocationInformationV2) GetVersionLockOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionLock, true
}

// SetVersionLock sets field value
func (o *LocationInformationV2) SetVersionLock(v int32) {
	o.VersionLock = v
}

func (o LocationInformationV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationInformationV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	toSerialize["realname"] = o.Realname
	toSerialize["phone"] = o.Phone
	toSerialize["email"] = o.Email
	toSerialize["room"] = o.Room
	toSerialize["position"] = o.Position
	toSerialize["departmentId"] = o.DepartmentId
	toSerialize["buildingId"] = o.BuildingId
	toSerialize["id"] = o.Id
	toSerialize["versionLock"] = o.VersionLock
	return toSerialize, nil
}

type NullableLocationInformationV2 struct {
	value *LocationInformationV2
	isSet bool
}

func (v NullableLocationInformationV2) Get() *LocationInformationV2 {
	return v.value
}

func (v *NullableLocationInformationV2) Set(val *LocationInformationV2) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationInformationV2) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationInformationV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationInformationV2(val *LocationInformationV2) *NullableLocationInformationV2 {
	return &NullableLocationInformationV2{value: val, isSet: true}
}

func (v NullableLocationInformationV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationInformationV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


