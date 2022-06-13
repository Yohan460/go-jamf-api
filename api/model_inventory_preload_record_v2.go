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

// InventoryPreloadRecordV2 struct for InventoryPreloadRecordV2
type InventoryPreloadRecordV2 struct {
	Id *string `json:"id,omitempty"`
	SerialNumber string `json:"serialNumber"`
	DeviceType string `json:"deviceType"`
	Username NullableString `json:"username,omitempty"`
	FullName NullableString `json:"fullName,omitempty"`
	EmailAddress NullableString `json:"emailAddress,omitempty"`
	PhoneNumber NullableString `json:"phoneNumber,omitempty"`
	Position NullableString `json:"position,omitempty"`
	Department NullableString `json:"department,omitempty"`
	Building NullableString `json:"building,omitempty"`
	Room NullableString `json:"room,omitempty"`
	PoNumber NullableString `json:"poNumber,omitempty"`
	PoDate NullableString `json:"poDate,omitempty"`
	WarrantyExpiration NullableString `json:"warrantyExpiration,omitempty"`
	AppleCareId NullableString `json:"appleCareId,omitempty"`
	LifeExpectancy NullableString `json:"lifeExpectancy,omitempty"`
	PurchasePrice NullableString `json:"purchasePrice,omitempty"`
	PurchasingContact NullableString `json:"purchasingContact,omitempty"`
	PurchasingAccount NullableString `json:"purchasingAccount,omitempty"`
	LeaseExpiration NullableString `json:"leaseExpiration,omitempty"`
	BarCode1 NullableString `json:"barCode1,omitempty"`
	BarCode2 NullableString `json:"barCode2,omitempty"`
	AssetTag NullableString `json:"assetTag,omitempty"`
	Vendor NullableString `json:"vendor,omitempty"`
	ExtensionAttributes []InventoryPreloadExtensionAttribute `json:"extensionAttributes,omitempty"`
}

// NewInventoryPreloadRecordV2 instantiates a new InventoryPreloadRecordV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryPreloadRecordV2(serialNumber string, deviceType string) *InventoryPreloadRecordV2 {
	this := InventoryPreloadRecordV2{}
	this.SerialNumber = serialNumber
	this.DeviceType = deviceType
	return &this
}

// NewInventoryPreloadRecordV2WithDefaults instantiates a new InventoryPreloadRecordV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryPreloadRecordV2WithDefaults() *InventoryPreloadRecordV2 {
	this := InventoryPreloadRecordV2{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InventoryPreloadRecordV2) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecordV2) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InventoryPreloadRecordV2) SetId(v string) {
	o.Id = &v
}

// GetSerialNumber returns the SerialNumber field value
func (o *InventoryPreloadRecordV2) GetSerialNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecordV2) GetSerialNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SerialNumber, true
}

// SetSerialNumber sets field value
func (o *InventoryPreloadRecordV2) SetSerialNumber(v string) {
	o.SerialNumber = v
}

// GetDeviceType returns the DeviceType field value
func (o *InventoryPreloadRecordV2) GetDeviceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecordV2) GetDeviceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *InventoryPreloadRecordV2) SetDeviceType(v string) {
	o.DeviceType = v
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryPreloadRecordV2) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryPreloadRecordV2) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *InventoryPreloadRecordV2) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *InventoryPreloadRecordV2) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *InventoryPreloadRecordV2) UnsetUsername() {
	o.Username.Unset()
}

// GetFullName returns the FullName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryPreloadRecordV2) GetFullName() string {
	if o == nil || o.FullName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FullName.Get()
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryPreloadRecordV2) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FullName.Get(), o.FullName.IsSet()
}

// HasFullName returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasFullName() bool {
	if o != nil && o.FullName.IsSet() {
		return true
	}

	return false
}

// SetFullName gets a reference to the given NullableString and assigns it to the FullName field.
func (o *InventoryPreloadRecordV2) SetFullName(v string) {
	o.FullName.Set(&v)
}
// SetFullNameNil sets the value for FullName to be an explicit nil
func (o *InventoryPreloadRecordV2) SetFullNameNil() {
	o.FullName.Set(nil)
}

// UnsetFullName ensures that no value is present for FullName, not even an explicit nil
func (o *InventoryPreloadRecordV2) UnsetFullName() {
	o.FullName.Unset()
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryPreloadRecordV2) GetEmailAddress() string {
	if o == nil || o.EmailAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress.Get()
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryPreloadRecordV2) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailAddress.Get(), o.EmailAddress.IsSet()
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasEmailAddress() bool {
	if o != nil && o.EmailAddress.IsSet() {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given NullableString and assigns it to the EmailAddress field.
func (o *InventoryPreloadRecordV2) SetEmailAddress(v string) {
	o.EmailAddress.Set(&v)
}
// SetEmailAddressNil sets the value for EmailAddress to be an explicit nil
func (o *InventoryPreloadRecordV2) SetEmailAddressNil() {
	o.EmailAddress.Set(nil)
}

// UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
func (o *InventoryPreloadRecordV2) UnsetEmailAddress() {
	o.EmailAddress.Unset()
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryPreloadRecordV2) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryPreloadRecordV2) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *InventoryPreloadRecordV2) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *InventoryPreloadRecordV2) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *InventoryPreloadRecordV2) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetPosition returns the Position field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryPreloadRecordV2) GetPosition() string {
	if o == nil || o.Position.Get() == nil {
		var ret string
		return ret
	}
	return *o.Position.Get()
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryPreloadRecordV2) GetPositionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Position.Get(), o.Position.IsSet()
}

// HasPosition returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasPosition() bool {
	if o != nil && o.Position.IsSet() {
		return true
	}

	return false
}

// SetPosition gets a reference to the given NullableString and assigns it to the Position field.
func (o *InventoryPreloadRecordV2) SetPosition(v string) {
	o.Position.Set(&v)
}
// SetPositionNil sets the value for Position to be an explicit nil
func (o *InventoryPreloadRecordV2) SetPositionNil() {
	o.Position.Set(nil)
}

// UnsetPosition ensures that no value is present for Position, not even an explicit nil
func (o *InventoryPreloadRecordV2) UnsetPosition() {
	o.Position.Unset()
}

// GetDepartment returns the Department field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryPreloadRecordV2) GetDepartment() string {
	if o == nil || o.Department.Get() == nil {
		var ret string
		return ret
	}
	return *o.Department.Get()
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryPreloadRecordV2) GetDepartmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Department.Get(), o.Department.IsSet()
}

// HasDepartment returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasDepartment() bool {
	if o != nil && o.Department.IsSet() {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given NullableString and assigns it to the Department field.
func (o *InventoryPreloadRecordV2) SetDepartment(v string) {
	o.Department.Set(&v)
}
// SetDepartmentNil sets the value for Department to be an explicit nil
func (o *InventoryPreloadRecordV2) SetDepartmentNil() {
	o.Department.Set(nil)
}

// UnsetDepartment ensures that no value is present for Department, not even an explicit nil
func (o *InventoryPreloadRecordV2) UnsetDepartment() {
	o.Department.Unset()
}

// GetBuilding returns the Building field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryPreloadRecordV2) GetBuilding() string {
	if o == nil || o.Building.Get() == nil {
		var ret string
		return ret
	}
	return *o.Building.Get()
}

// GetBuildingOk returns a tuple with the Building field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryPreloadRecordV2) GetBuildingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Building.Get(), o.Building.IsSet()
}

// HasBuilding returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasBuilding() bool {
	if o != nil && o.Building.IsSet() {
		return true
	}

	return false
}

// SetBuilding gets a reference to the given NullableString and assigns it to the Building field.
func (o *InventoryPreloadRecordV2) SetBuilding(v string) {
	o.Building.Set(&v)
}
// SetBuildingNil sets the value for Building to be an explicit nil
func (o *InventoryPreloadRecordV2) SetBuildingNil() {
	o.Building.Set(nil)
}

// UnsetBuilding ensures that no value is present for Building, not even an explicit nil
func (o *InventoryPreloadRecordV2) UnsetBuilding() {
	o.Building.Unset()
}

// GetRoom returns the Room field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryPreloadRecordV2) GetRoom() string {
	if o == nil || o.Room.Get() == nil {
		var ret string
		return ret
	}
	return *o.Room.Get()
}

// GetRoomOk returns a tuple with the Room field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryPreloadRecordV2) GetRoomOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Room.Get(), o.Room.IsSet()
}

// HasRoom returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasRoom() bool {
	if o != nil && o.Room.IsSet() {
		return true
	}

	return false
}

// SetRoom gets a reference to the given NullableString and assigns it to the Room field.
func (o *InventoryPreloadRecordV2) SetRoom(v string) {
	o.Room.Set(&v)
}
// SetRoomNil sets the value for Room to be an explicit nil
func (o *InventoryPreloadRecordV2) SetRoomNil() {
	o.Room.Set(nil)
}

// UnsetRoom ensures that no value is present for Room, not even an explicit nil
func (o *InventoryPreloadRecordV2) UnsetRoom() {
	o.Room.Unset()
}

// GetPoNumber returns the PoNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryPreloadRecordV2) GetPoNumber() string {
	if o == nil || o.PoNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.PoNumber.Get()
}

// GetPoNumberOk returns a tuple with the PoNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryPreloadRecordV2) GetPoNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PoNumber.Get(), o.PoNumber.IsSet()
}

// HasPoNumber returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasPoNumber() bool {
	if o != nil && o.PoNumber.IsSet() {
		return true
	}

	return false
}

// SetPoNumber gets a reference to the given NullableString and assigns it to the PoNumber field.
func (o *InventoryPreloadRecordV2) SetPoNumber(v string) {
	o.PoNumber.Set(&v)
}
// SetPoNumberNil sets the value for PoNumber to be an explicit nil
func (o *InventoryPreloadRecordV2) SetPoNumberNil() {
	o.PoNumber.Set(nil)
}

// UnsetPoNumber ensures that no value is present for PoNumber, not even an explicit nil
func (o *InventoryPreloadRecordV2) UnsetPoNumber() {
	o.PoNumber.Unset()
}

// GetPoDate returns the PoDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryPreloadRecordV2) GetPoDate() string {
	if o == nil || o.PoDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.PoDate.Get()
}

// GetPoDateOk returns a tuple with the PoDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryPreloadRecordV2) GetPoDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PoDate.Get(), o.PoDate.IsSet()
}

// HasPoDate returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasPoDate() bool {
	if o != nil && o.PoDate.IsSet() {
		return true
	}

	return false
}

// SetPoDate gets a reference to the given NullableString and assigns it to the PoDate field.
func (o *InventoryPreloadRecordV2) SetPoDate(v string) {
	o.PoDate.Set(&v)
}
// SetPoDateNil sets the value for PoDate to be an explicit nil
func (o *InventoryPreloadRecordV2) SetPoDateNil() {
	o.PoDate.Set(nil)
}

// UnsetPoDate ensures that no value is present for PoDate, not even an explicit nil
func (o *InventoryPreloadRecordV2) UnsetPoDate() {
	o.PoDate.Unset()
}

// GetWarrantyExpiration returns the WarrantyExpiration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryPreloadRecordV2) GetWarrantyExpiration() string {
	if o == nil || o.WarrantyExpiration.Get() == nil {
		var ret string
		return ret
	}
	return *o.WarrantyExpiration.Get()
}

// GetWarrantyExpirationOk returns a tuple with the WarrantyExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryPreloadRecordV2) GetWarrantyExpirationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WarrantyExpiration.Get(), o.WarrantyExpiration.IsSet()
}

// HasWarrantyExpiration returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasWarrantyExpiration() bool {
	if o != nil && o.WarrantyExpiration.IsSet() {
		return true
	}

	return false
}

// SetWarrantyExpiration gets a reference to the given NullableString and assigns it to the WarrantyExpiration field.
func (o *InventoryPreloadRecordV2) SetWarrantyExpiration(v string) {
	o.WarrantyExpiration.Set(&v)
}
// SetWarrantyExpirationNil sets the value for WarrantyExpiration to be an explicit nil
func (o *InventoryPreloadRecordV2) SetWarrantyExpirationNil() {
	o.WarrantyExpiration.Set(nil)
}

// UnsetWarrantyExpiration ensures that no value is present for WarrantyExpiration, not even an explicit nil
func (o *InventoryPreloadRecordV2) UnsetWarrantyExpiration() {
	o.WarrantyExpiration.Unset()
}

// GetAppleCareId returns the AppleCareId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryPreloadRecordV2) GetAppleCareId() string {
	if o == nil || o.AppleCareId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AppleCareId.Get()
}

// GetAppleCareIdOk returns a tuple with the AppleCareId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryPreloadRecordV2) GetAppleCareIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppleCareId.Get(), o.AppleCareId.IsSet()
}

// HasAppleCareId returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasAppleCareId() bool {
	if o != nil && o.AppleCareId.IsSet() {
		return true
	}

	return false
}

// SetAppleCareId gets a reference to the given NullableString and assigns it to the AppleCareId field.
func (o *InventoryPreloadRecordV2) SetAppleCareId(v string) {
	o.AppleCareId.Set(&v)
}
// SetAppleCareIdNil sets the value for AppleCareId to be an explicit nil
func (o *InventoryPreloadRecordV2) SetAppleCareIdNil() {
	o.AppleCareId.Set(nil)
}

// UnsetAppleCareId ensures that no value is present for AppleCareId, not even an explicit nil
func (o *InventoryPreloadRecordV2) UnsetAppleCareId() {
	o.AppleCareId.Unset()
}

// GetLifeExpectancy returns the LifeExpectancy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryPreloadRecordV2) GetLifeExpectancy() string {
	if o == nil || o.LifeExpectancy.Get() == nil {
		var ret string
		return ret
	}
	return *o.LifeExpectancy.Get()
}

// GetLifeExpectancyOk returns a tuple with the LifeExpectancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryPreloadRecordV2) GetLifeExpectancyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LifeExpectancy.Get(), o.LifeExpectancy.IsSet()
}

// HasLifeExpectancy returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasLifeExpectancy() bool {
	if o != nil && o.LifeExpectancy.IsSet() {
		return true
	}

	return false
}

// SetLifeExpectancy gets a reference to the given NullableString and assigns it to the LifeExpectancy field.
func (o *InventoryPreloadRecordV2) SetLifeExpectancy(v string) {
	o.LifeExpectancy.Set(&v)
}
// SetLifeExpectancyNil sets the value for LifeExpectancy to be an explicit nil
func (o *InventoryPreloadRecordV2) SetLifeExpectancyNil() {
	o.LifeExpectancy.Set(nil)
}

// UnsetLifeExpectancy ensures that no value is present for LifeExpectancy, not even an explicit nil
func (o *InventoryPreloadRecordV2) UnsetLifeExpectancy() {
	o.LifeExpectancy.Unset()
}

// GetPurchasePrice returns the PurchasePrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryPreloadRecordV2) GetPurchasePrice() string {
	if o == nil || o.PurchasePrice.Get() == nil {
		var ret string
		return ret
	}
	return *o.PurchasePrice.Get()
}

// GetPurchasePriceOk returns a tuple with the PurchasePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryPreloadRecordV2) GetPurchasePriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PurchasePrice.Get(), o.PurchasePrice.IsSet()
}

// HasPurchasePrice returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasPurchasePrice() bool {
	if o != nil && o.PurchasePrice.IsSet() {
		return true
	}

	return false
}

// SetPurchasePrice gets a reference to the given NullableString and assigns it to the PurchasePrice field.
func (o *InventoryPreloadRecordV2) SetPurchasePrice(v string) {
	o.PurchasePrice.Set(&v)
}
// SetPurchasePriceNil sets the value for PurchasePrice to be an explicit nil
func (o *InventoryPreloadRecordV2) SetPurchasePriceNil() {
	o.PurchasePrice.Set(nil)
}

// UnsetPurchasePrice ensures that no value is present for PurchasePrice, not even an explicit nil
func (o *InventoryPreloadRecordV2) UnsetPurchasePrice() {
	o.PurchasePrice.Unset()
}

// GetPurchasingContact returns the PurchasingContact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryPreloadRecordV2) GetPurchasingContact() string {
	if o == nil || o.PurchasingContact.Get() == nil {
		var ret string
		return ret
	}
	return *o.PurchasingContact.Get()
}

// GetPurchasingContactOk returns a tuple with the PurchasingContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryPreloadRecordV2) GetPurchasingContactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PurchasingContact.Get(), o.PurchasingContact.IsSet()
}

// HasPurchasingContact returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasPurchasingContact() bool {
	if o != nil && o.PurchasingContact.IsSet() {
		return true
	}

	return false
}

// SetPurchasingContact gets a reference to the given NullableString and assigns it to the PurchasingContact field.
func (o *InventoryPreloadRecordV2) SetPurchasingContact(v string) {
	o.PurchasingContact.Set(&v)
}
// SetPurchasingContactNil sets the value for PurchasingContact to be an explicit nil
func (o *InventoryPreloadRecordV2) SetPurchasingContactNil() {
	o.PurchasingContact.Set(nil)
}

// UnsetPurchasingContact ensures that no value is present for PurchasingContact, not even an explicit nil
func (o *InventoryPreloadRecordV2) UnsetPurchasingContact() {
	o.PurchasingContact.Unset()
}

// GetPurchasingAccount returns the PurchasingAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryPreloadRecordV2) GetPurchasingAccount() string {
	if o == nil || o.PurchasingAccount.Get() == nil {
		var ret string
		return ret
	}
	return *o.PurchasingAccount.Get()
}

// GetPurchasingAccountOk returns a tuple with the PurchasingAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryPreloadRecordV2) GetPurchasingAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PurchasingAccount.Get(), o.PurchasingAccount.IsSet()
}

// HasPurchasingAccount returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasPurchasingAccount() bool {
	if o != nil && o.PurchasingAccount.IsSet() {
		return true
	}

	return false
}

// SetPurchasingAccount gets a reference to the given NullableString and assigns it to the PurchasingAccount field.
func (o *InventoryPreloadRecordV2) SetPurchasingAccount(v string) {
	o.PurchasingAccount.Set(&v)
}
// SetPurchasingAccountNil sets the value for PurchasingAccount to be an explicit nil
func (o *InventoryPreloadRecordV2) SetPurchasingAccountNil() {
	o.PurchasingAccount.Set(nil)
}

// UnsetPurchasingAccount ensures that no value is present for PurchasingAccount, not even an explicit nil
func (o *InventoryPreloadRecordV2) UnsetPurchasingAccount() {
	o.PurchasingAccount.Unset()
}

// GetLeaseExpiration returns the LeaseExpiration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryPreloadRecordV2) GetLeaseExpiration() string {
	if o == nil || o.LeaseExpiration.Get() == nil {
		var ret string
		return ret
	}
	return *o.LeaseExpiration.Get()
}

// GetLeaseExpirationOk returns a tuple with the LeaseExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryPreloadRecordV2) GetLeaseExpirationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LeaseExpiration.Get(), o.LeaseExpiration.IsSet()
}

// HasLeaseExpiration returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasLeaseExpiration() bool {
	if o != nil && o.LeaseExpiration.IsSet() {
		return true
	}

	return false
}

// SetLeaseExpiration gets a reference to the given NullableString and assigns it to the LeaseExpiration field.
func (o *InventoryPreloadRecordV2) SetLeaseExpiration(v string) {
	o.LeaseExpiration.Set(&v)
}
// SetLeaseExpirationNil sets the value for LeaseExpiration to be an explicit nil
func (o *InventoryPreloadRecordV2) SetLeaseExpirationNil() {
	o.LeaseExpiration.Set(nil)
}

// UnsetLeaseExpiration ensures that no value is present for LeaseExpiration, not even an explicit nil
func (o *InventoryPreloadRecordV2) UnsetLeaseExpiration() {
	o.LeaseExpiration.Unset()
}

// GetBarCode1 returns the BarCode1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryPreloadRecordV2) GetBarCode1() string {
	if o == nil || o.BarCode1.Get() == nil {
		var ret string
		return ret
	}
	return *o.BarCode1.Get()
}

// GetBarCode1Ok returns a tuple with the BarCode1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryPreloadRecordV2) GetBarCode1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BarCode1.Get(), o.BarCode1.IsSet()
}

// HasBarCode1 returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasBarCode1() bool {
	if o != nil && o.BarCode1.IsSet() {
		return true
	}

	return false
}

// SetBarCode1 gets a reference to the given NullableString and assigns it to the BarCode1 field.
func (o *InventoryPreloadRecordV2) SetBarCode1(v string) {
	o.BarCode1.Set(&v)
}
// SetBarCode1Nil sets the value for BarCode1 to be an explicit nil
func (o *InventoryPreloadRecordV2) SetBarCode1Nil() {
	o.BarCode1.Set(nil)
}

// UnsetBarCode1 ensures that no value is present for BarCode1, not even an explicit nil
func (o *InventoryPreloadRecordV2) UnsetBarCode1() {
	o.BarCode1.Unset()
}

// GetBarCode2 returns the BarCode2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryPreloadRecordV2) GetBarCode2() string {
	if o == nil || o.BarCode2.Get() == nil {
		var ret string
		return ret
	}
	return *o.BarCode2.Get()
}

// GetBarCode2Ok returns a tuple with the BarCode2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryPreloadRecordV2) GetBarCode2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BarCode2.Get(), o.BarCode2.IsSet()
}

// HasBarCode2 returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasBarCode2() bool {
	if o != nil && o.BarCode2.IsSet() {
		return true
	}

	return false
}

// SetBarCode2 gets a reference to the given NullableString and assigns it to the BarCode2 field.
func (o *InventoryPreloadRecordV2) SetBarCode2(v string) {
	o.BarCode2.Set(&v)
}
// SetBarCode2Nil sets the value for BarCode2 to be an explicit nil
func (o *InventoryPreloadRecordV2) SetBarCode2Nil() {
	o.BarCode2.Set(nil)
}

// UnsetBarCode2 ensures that no value is present for BarCode2, not even an explicit nil
func (o *InventoryPreloadRecordV2) UnsetBarCode2() {
	o.BarCode2.Unset()
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryPreloadRecordV2) GetAssetTag() string {
	if o == nil || o.AssetTag.Get() == nil {
		var ret string
		return ret
	}
	return *o.AssetTag.Get()
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryPreloadRecordV2) GetAssetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTag.Get(), o.AssetTag.IsSet()
}

// HasAssetTag returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasAssetTag() bool {
	if o != nil && o.AssetTag.IsSet() {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given NullableString and assigns it to the AssetTag field.
func (o *InventoryPreloadRecordV2) SetAssetTag(v string) {
	o.AssetTag.Set(&v)
}
// SetAssetTagNil sets the value for AssetTag to be an explicit nil
func (o *InventoryPreloadRecordV2) SetAssetTagNil() {
	o.AssetTag.Set(nil)
}

// UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
func (o *InventoryPreloadRecordV2) UnsetAssetTag() {
	o.AssetTag.Unset()
}

// GetVendor returns the Vendor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryPreloadRecordV2) GetVendor() string {
	if o == nil || o.Vendor.Get() == nil {
		var ret string
		return ret
	}
	return *o.Vendor.Get()
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryPreloadRecordV2) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vendor.Get(), o.Vendor.IsSet()
}

// HasVendor returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasVendor() bool {
	if o != nil && o.Vendor.IsSet() {
		return true
	}

	return false
}

// SetVendor gets a reference to the given NullableString and assigns it to the Vendor field.
func (o *InventoryPreloadRecordV2) SetVendor(v string) {
	o.Vendor.Set(&v)
}
// SetVendorNil sets the value for Vendor to be an explicit nil
func (o *InventoryPreloadRecordV2) SetVendorNil() {
	o.Vendor.Set(nil)
}

// UnsetVendor ensures that no value is present for Vendor, not even an explicit nil
func (o *InventoryPreloadRecordV2) UnsetVendor() {
	o.Vendor.Unset()
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *InventoryPreloadRecordV2) GetExtensionAttributes() []InventoryPreloadExtensionAttribute {
	if o == nil || o.ExtensionAttributes == nil {
		var ret []InventoryPreloadExtensionAttribute
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecordV2) GetExtensionAttributesOk() ([]InventoryPreloadExtensionAttribute, bool) {
	if o == nil || o.ExtensionAttributes == nil {
		return nil, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *InventoryPreloadRecordV2) HasExtensionAttributes() bool {
	if o != nil && o.ExtensionAttributes != nil {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given []InventoryPreloadExtensionAttribute and assigns it to the ExtensionAttributes field.
func (o *InventoryPreloadRecordV2) SetExtensionAttributes(v []InventoryPreloadExtensionAttribute) {
	o.ExtensionAttributes = v
}

func (o InventoryPreloadRecordV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if true {
		toSerialize["deviceType"] = o.DeviceType
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.FullName.IsSet() {
		toSerialize["fullName"] = o.FullName.Get()
	}
	if o.EmailAddress.IsSet() {
		toSerialize["emailAddress"] = o.EmailAddress.Get()
	}
	if o.PhoneNumber.IsSet() {
		toSerialize["phoneNumber"] = o.PhoneNumber.Get()
	}
	if o.Position.IsSet() {
		toSerialize["position"] = o.Position.Get()
	}
	if o.Department.IsSet() {
		toSerialize["department"] = o.Department.Get()
	}
	if o.Building.IsSet() {
		toSerialize["building"] = o.Building.Get()
	}
	if o.Room.IsSet() {
		toSerialize["room"] = o.Room.Get()
	}
	if o.PoNumber.IsSet() {
		toSerialize["poNumber"] = o.PoNumber.Get()
	}
	if o.PoDate.IsSet() {
		toSerialize["poDate"] = o.PoDate.Get()
	}
	if o.WarrantyExpiration.IsSet() {
		toSerialize["warrantyExpiration"] = o.WarrantyExpiration.Get()
	}
	if o.AppleCareId.IsSet() {
		toSerialize["appleCareId"] = o.AppleCareId.Get()
	}
	if o.LifeExpectancy.IsSet() {
		toSerialize["lifeExpectancy"] = o.LifeExpectancy.Get()
	}
	if o.PurchasePrice.IsSet() {
		toSerialize["purchasePrice"] = o.PurchasePrice.Get()
	}
	if o.PurchasingContact.IsSet() {
		toSerialize["purchasingContact"] = o.PurchasingContact.Get()
	}
	if o.PurchasingAccount.IsSet() {
		toSerialize["purchasingAccount"] = o.PurchasingAccount.Get()
	}
	if o.LeaseExpiration.IsSet() {
		toSerialize["leaseExpiration"] = o.LeaseExpiration.Get()
	}
	if o.BarCode1.IsSet() {
		toSerialize["barCode1"] = o.BarCode1.Get()
	}
	if o.BarCode2.IsSet() {
		toSerialize["barCode2"] = o.BarCode2.Get()
	}
	if o.AssetTag.IsSet() {
		toSerialize["assetTag"] = o.AssetTag.Get()
	}
	if o.Vendor.IsSet() {
		toSerialize["vendor"] = o.Vendor.Get()
	}
	if o.ExtensionAttributes != nil {
		toSerialize["extensionAttributes"] = o.ExtensionAttributes
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryPreloadRecordV2 struct {
	value *InventoryPreloadRecordV2
	isSet bool
}

func (v NullableInventoryPreloadRecordV2) Get() *InventoryPreloadRecordV2 {
	return v.value
}

func (v *NullableInventoryPreloadRecordV2) Set(val *InventoryPreloadRecordV2) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryPreloadRecordV2) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryPreloadRecordV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryPreloadRecordV2(val *InventoryPreloadRecordV2) *NullableInventoryPreloadRecordV2 {
	return &NullableInventoryPreloadRecordV2{value: val, isSet: true}
}

func (v NullableInventoryPreloadRecordV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryPreloadRecordV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


