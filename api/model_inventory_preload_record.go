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

// checks if the InventoryPreloadRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryPreloadRecord{}

// InventoryPreloadRecord struct for InventoryPreloadRecord
type InventoryPreloadRecord struct {
	Id *int64 `json:"id,omitempty"`
	SerialNumber string `json:"serialNumber"`
	DeviceType string `json:"deviceType"`
	Username *string `json:"username,omitempty"`
	FullName *string `json:"fullName,omitempty"`
	EmailAddress *string `json:"emailAddress,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Position *string `json:"position,omitempty"`
	Department *string `json:"department,omitempty"`
	Building *string `json:"building,omitempty"`
	Room *string `json:"room,omitempty"`
	PoNumber *string `json:"poNumber,omitempty"`
	PoDate *string `json:"poDate,omitempty"`
	WarrantyExpiration *string `json:"warrantyExpiration,omitempty"`
	AppleCareId *string `json:"appleCareId,omitempty"`
	LifeExpectancy *string `json:"lifeExpectancy,omitempty"`
	PurchasePrice *string `json:"purchasePrice,omitempty"`
	PurchasingContact *string `json:"purchasingContact,omitempty"`
	PurchasingAccount *string `json:"purchasingAccount,omitempty"`
	LeaseExpiration *string `json:"leaseExpiration,omitempty"`
	BarCode1 *string `json:"barCode1,omitempty"`
	BarCode2 *string `json:"barCode2,omitempty"`
	AssetTag *string `json:"assetTag,omitempty"`
	Vendor *string `json:"vendor,omitempty"`
	ExtensionAttributes []InventoryPreloadExtensionAttribute `json:"extensionAttributes,omitempty"`
}

type _InventoryPreloadRecord InventoryPreloadRecord

// NewInventoryPreloadRecord instantiates a new InventoryPreloadRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryPreloadRecord(serialNumber string, deviceType string) *InventoryPreloadRecord {
	this := InventoryPreloadRecord{}
	this.SerialNumber = serialNumber
	this.DeviceType = deviceType
	return &this
}

// NewInventoryPreloadRecordWithDefaults instantiates a new InventoryPreloadRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryPreloadRecordWithDefaults() *InventoryPreloadRecord {
	this := InventoryPreloadRecord{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *InventoryPreloadRecord) SetId(v int64) {
	o.Id = &v
}

// GetSerialNumber returns the SerialNumber field value
func (o *InventoryPreloadRecord) GetSerialNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetSerialNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SerialNumber, true
}

// SetSerialNumber sets field value
func (o *InventoryPreloadRecord) SetSerialNumber(v string) {
	o.SerialNumber = v
}

// GetDeviceType returns the DeviceType field value
func (o *InventoryPreloadRecord) GetDeviceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetDeviceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *InventoryPreloadRecord) SetDeviceType(v string) {
	o.DeviceType = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *InventoryPreloadRecord) SetUsername(v string) {
	o.Username = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *InventoryPreloadRecord) SetFullName(v string) {
	o.FullName = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *InventoryPreloadRecord) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *InventoryPreloadRecord) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetPosition() string {
	if o == nil || IsNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetPositionOk() (*string, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *InventoryPreloadRecord) SetPosition(v string) {
	o.Position = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetDepartment() string {
	if o == nil || IsNil(o.Department) {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetDepartmentOk() (*string, bool) {
	if o == nil || IsNil(o.Department) {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasDepartment() bool {
	if o != nil && !IsNil(o.Department) {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *InventoryPreloadRecord) SetDepartment(v string) {
	o.Department = &v
}

// GetBuilding returns the Building field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetBuilding() string {
	if o == nil || IsNil(o.Building) {
		var ret string
		return ret
	}
	return *o.Building
}

// GetBuildingOk returns a tuple with the Building field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetBuildingOk() (*string, bool) {
	if o == nil || IsNil(o.Building) {
		return nil, false
	}
	return o.Building, true
}

// HasBuilding returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasBuilding() bool {
	if o != nil && !IsNil(o.Building) {
		return true
	}

	return false
}

// SetBuilding gets a reference to the given string and assigns it to the Building field.
func (o *InventoryPreloadRecord) SetBuilding(v string) {
	o.Building = &v
}

// GetRoom returns the Room field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetRoom() string {
	if o == nil || IsNil(o.Room) {
		var ret string
		return ret
	}
	return *o.Room
}

// GetRoomOk returns a tuple with the Room field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetRoomOk() (*string, bool) {
	if o == nil || IsNil(o.Room) {
		return nil, false
	}
	return o.Room, true
}

// HasRoom returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasRoom() bool {
	if o != nil && !IsNil(o.Room) {
		return true
	}

	return false
}

// SetRoom gets a reference to the given string and assigns it to the Room field.
func (o *InventoryPreloadRecord) SetRoom(v string) {
	o.Room = &v
}

// GetPoNumber returns the PoNumber field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetPoNumber() string {
	if o == nil || IsNil(o.PoNumber) {
		var ret string
		return ret
	}
	return *o.PoNumber
}

// GetPoNumberOk returns a tuple with the PoNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetPoNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PoNumber) {
		return nil, false
	}
	return o.PoNumber, true
}

// HasPoNumber returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasPoNumber() bool {
	if o != nil && !IsNil(o.PoNumber) {
		return true
	}

	return false
}

// SetPoNumber gets a reference to the given string and assigns it to the PoNumber field.
func (o *InventoryPreloadRecord) SetPoNumber(v string) {
	o.PoNumber = &v
}

// GetPoDate returns the PoDate field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetPoDate() string {
	if o == nil || IsNil(o.PoDate) {
		var ret string
		return ret
	}
	return *o.PoDate
}

// GetPoDateOk returns a tuple with the PoDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetPoDateOk() (*string, bool) {
	if o == nil || IsNil(o.PoDate) {
		return nil, false
	}
	return o.PoDate, true
}

// HasPoDate returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasPoDate() bool {
	if o != nil && !IsNil(o.PoDate) {
		return true
	}

	return false
}

// SetPoDate gets a reference to the given string and assigns it to the PoDate field.
func (o *InventoryPreloadRecord) SetPoDate(v string) {
	o.PoDate = &v
}

// GetWarrantyExpiration returns the WarrantyExpiration field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetWarrantyExpiration() string {
	if o == nil || IsNil(o.WarrantyExpiration) {
		var ret string
		return ret
	}
	return *o.WarrantyExpiration
}

// GetWarrantyExpirationOk returns a tuple with the WarrantyExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetWarrantyExpirationOk() (*string, bool) {
	if o == nil || IsNil(o.WarrantyExpiration) {
		return nil, false
	}
	return o.WarrantyExpiration, true
}

// HasWarrantyExpiration returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasWarrantyExpiration() bool {
	if o != nil && !IsNil(o.WarrantyExpiration) {
		return true
	}

	return false
}

// SetWarrantyExpiration gets a reference to the given string and assigns it to the WarrantyExpiration field.
func (o *InventoryPreloadRecord) SetWarrantyExpiration(v string) {
	o.WarrantyExpiration = &v
}

// GetAppleCareId returns the AppleCareId field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetAppleCareId() string {
	if o == nil || IsNil(o.AppleCareId) {
		var ret string
		return ret
	}
	return *o.AppleCareId
}

// GetAppleCareIdOk returns a tuple with the AppleCareId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetAppleCareIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppleCareId) {
		return nil, false
	}
	return o.AppleCareId, true
}

// HasAppleCareId returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasAppleCareId() bool {
	if o != nil && !IsNil(o.AppleCareId) {
		return true
	}

	return false
}

// SetAppleCareId gets a reference to the given string and assigns it to the AppleCareId field.
func (o *InventoryPreloadRecord) SetAppleCareId(v string) {
	o.AppleCareId = &v
}

// GetLifeExpectancy returns the LifeExpectancy field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetLifeExpectancy() string {
	if o == nil || IsNil(o.LifeExpectancy) {
		var ret string
		return ret
	}
	return *o.LifeExpectancy
}

// GetLifeExpectancyOk returns a tuple with the LifeExpectancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetLifeExpectancyOk() (*string, bool) {
	if o == nil || IsNil(o.LifeExpectancy) {
		return nil, false
	}
	return o.LifeExpectancy, true
}

// HasLifeExpectancy returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasLifeExpectancy() bool {
	if o != nil && !IsNil(o.LifeExpectancy) {
		return true
	}

	return false
}

// SetLifeExpectancy gets a reference to the given string and assigns it to the LifeExpectancy field.
func (o *InventoryPreloadRecord) SetLifeExpectancy(v string) {
	o.LifeExpectancy = &v
}

// GetPurchasePrice returns the PurchasePrice field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetPurchasePrice() string {
	if o == nil || IsNil(o.PurchasePrice) {
		var ret string
		return ret
	}
	return *o.PurchasePrice
}

// GetPurchasePriceOk returns a tuple with the PurchasePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetPurchasePriceOk() (*string, bool) {
	if o == nil || IsNil(o.PurchasePrice) {
		return nil, false
	}
	return o.PurchasePrice, true
}

// HasPurchasePrice returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasPurchasePrice() bool {
	if o != nil && !IsNil(o.PurchasePrice) {
		return true
	}

	return false
}

// SetPurchasePrice gets a reference to the given string and assigns it to the PurchasePrice field.
func (o *InventoryPreloadRecord) SetPurchasePrice(v string) {
	o.PurchasePrice = &v
}

// GetPurchasingContact returns the PurchasingContact field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetPurchasingContact() string {
	if o == nil || IsNil(o.PurchasingContact) {
		var ret string
		return ret
	}
	return *o.PurchasingContact
}

// GetPurchasingContactOk returns a tuple with the PurchasingContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetPurchasingContactOk() (*string, bool) {
	if o == nil || IsNil(o.PurchasingContact) {
		return nil, false
	}
	return o.PurchasingContact, true
}

// HasPurchasingContact returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasPurchasingContact() bool {
	if o != nil && !IsNil(o.PurchasingContact) {
		return true
	}

	return false
}

// SetPurchasingContact gets a reference to the given string and assigns it to the PurchasingContact field.
func (o *InventoryPreloadRecord) SetPurchasingContact(v string) {
	o.PurchasingContact = &v
}

// GetPurchasingAccount returns the PurchasingAccount field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetPurchasingAccount() string {
	if o == nil || IsNil(o.PurchasingAccount) {
		var ret string
		return ret
	}
	return *o.PurchasingAccount
}

// GetPurchasingAccountOk returns a tuple with the PurchasingAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetPurchasingAccountOk() (*string, bool) {
	if o == nil || IsNil(o.PurchasingAccount) {
		return nil, false
	}
	return o.PurchasingAccount, true
}

// HasPurchasingAccount returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasPurchasingAccount() bool {
	if o != nil && !IsNil(o.PurchasingAccount) {
		return true
	}

	return false
}

// SetPurchasingAccount gets a reference to the given string and assigns it to the PurchasingAccount field.
func (o *InventoryPreloadRecord) SetPurchasingAccount(v string) {
	o.PurchasingAccount = &v
}

// GetLeaseExpiration returns the LeaseExpiration field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetLeaseExpiration() string {
	if o == nil || IsNil(o.LeaseExpiration) {
		var ret string
		return ret
	}
	return *o.LeaseExpiration
}

// GetLeaseExpirationOk returns a tuple with the LeaseExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetLeaseExpirationOk() (*string, bool) {
	if o == nil || IsNil(o.LeaseExpiration) {
		return nil, false
	}
	return o.LeaseExpiration, true
}

// HasLeaseExpiration returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasLeaseExpiration() bool {
	if o != nil && !IsNil(o.LeaseExpiration) {
		return true
	}

	return false
}

// SetLeaseExpiration gets a reference to the given string and assigns it to the LeaseExpiration field.
func (o *InventoryPreloadRecord) SetLeaseExpiration(v string) {
	o.LeaseExpiration = &v
}

// GetBarCode1 returns the BarCode1 field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetBarCode1() string {
	if o == nil || IsNil(o.BarCode1) {
		var ret string
		return ret
	}
	return *o.BarCode1
}

// GetBarCode1Ok returns a tuple with the BarCode1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetBarCode1Ok() (*string, bool) {
	if o == nil || IsNil(o.BarCode1) {
		return nil, false
	}
	return o.BarCode1, true
}

// HasBarCode1 returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasBarCode1() bool {
	if o != nil && !IsNil(o.BarCode1) {
		return true
	}

	return false
}

// SetBarCode1 gets a reference to the given string and assigns it to the BarCode1 field.
func (o *InventoryPreloadRecord) SetBarCode1(v string) {
	o.BarCode1 = &v
}

// GetBarCode2 returns the BarCode2 field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetBarCode2() string {
	if o == nil || IsNil(o.BarCode2) {
		var ret string
		return ret
	}
	return *o.BarCode2
}

// GetBarCode2Ok returns a tuple with the BarCode2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetBarCode2Ok() (*string, bool) {
	if o == nil || IsNil(o.BarCode2) {
		return nil, false
	}
	return o.BarCode2, true
}

// HasBarCode2 returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasBarCode2() bool {
	if o != nil && !IsNil(o.BarCode2) {
		return true
	}

	return false
}

// SetBarCode2 gets a reference to the given string and assigns it to the BarCode2 field.
func (o *InventoryPreloadRecord) SetBarCode2(v string) {
	o.BarCode2 = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetAssetTag() string {
	if o == nil || IsNil(o.AssetTag) {
		var ret string
		return ret
	}
	return *o.AssetTag
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetAssetTagOk() (*string, bool) {
	if o == nil || IsNil(o.AssetTag) {
		return nil, false
	}
	return o.AssetTag, true
}

// HasAssetTag returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasAssetTag() bool {
	if o != nil && !IsNil(o.AssetTag) {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given string and assigns it to the AssetTag field.
func (o *InventoryPreloadRecord) SetAssetTag(v string) {
	o.AssetTag = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *InventoryPreloadRecord) SetVendor(v string) {
	o.Vendor = &v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *InventoryPreloadRecord) GetExtensionAttributes() []InventoryPreloadExtensionAttribute {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret []InventoryPreloadExtensionAttribute
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryPreloadRecord) GetExtensionAttributesOk() ([]InventoryPreloadExtensionAttribute, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return nil, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *InventoryPreloadRecord) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given []InventoryPreloadExtensionAttribute and assigns it to the ExtensionAttributes field.
func (o *InventoryPreloadRecord) SetExtensionAttributes(v []InventoryPreloadExtensionAttribute) {
	o.ExtensionAttributes = v
}

func (o InventoryPreloadRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryPreloadRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["serialNumber"] = o.SerialNumber
	toSerialize["deviceType"] = o.DeviceType
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.FullName) {
		toSerialize["fullName"] = o.FullName
	}
	if !IsNil(o.EmailAddress) {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Department) {
		toSerialize["department"] = o.Department
	}
	if !IsNil(o.Building) {
		toSerialize["building"] = o.Building
	}
	if !IsNil(o.Room) {
		toSerialize["room"] = o.Room
	}
	if !IsNil(o.PoNumber) {
		toSerialize["poNumber"] = o.PoNumber
	}
	if !IsNil(o.PoDate) {
		toSerialize["poDate"] = o.PoDate
	}
	if !IsNil(o.WarrantyExpiration) {
		toSerialize["warrantyExpiration"] = o.WarrantyExpiration
	}
	if !IsNil(o.AppleCareId) {
		toSerialize["appleCareId"] = o.AppleCareId
	}
	if !IsNil(o.LifeExpectancy) {
		toSerialize["lifeExpectancy"] = o.LifeExpectancy
	}
	if !IsNil(o.PurchasePrice) {
		toSerialize["purchasePrice"] = o.PurchasePrice
	}
	if !IsNil(o.PurchasingContact) {
		toSerialize["purchasingContact"] = o.PurchasingContact
	}
	if !IsNil(o.PurchasingAccount) {
		toSerialize["purchasingAccount"] = o.PurchasingAccount
	}
	if !IsNil(o.LeaseExpiration) {
		toSerialize["leaseExpiration"] = o.LeaseExpiration
	}
	if !IsNil(o.BarCode1) {
		toSerialize["barCode1"] = o.BarCode1
	}
	if !IsNil(o.BarCode2) {
		toSerialize["barCode2"] = o.BarCode2
	}
	if !IsNil(o.AssetTag) {
		toSerialize["assetTag"] = o.AssetTag
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extensionAttributes"] = o.ExtensionAttributes
	}
	return toSerialize, nil
}

func (o *InventoryPreloadRecord) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"serialNumber",
		"deviceType",
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

	varInventoryPreloadRecord := _InventoryPreloadRecord{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInventoryPreloadRecord)

	if err != nil {
		return err
	}

	*o = InventoryPreloadRecord(varInventoryPreloadRecord)

	return err
}

type NullableInventoryPreloadRecord struct {
	value *InventoryPreloadRecord
	isSet bool
}

func (v NullableInventoryPreloadRecord) Get() *InventoryPreloadRecord {
	return v.value
}

func (v *NullableInventoryPreloadRecord) Set(val *InventoryPreloadRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryPreloadRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryPreloadRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryPreloadRecord(val *InventoryPreloadRecord) *NullableInventoryPreloadRecord {
	return &NullableInventoryPreloadRecord{value: val, isSet: true}
}

func (v NullableInventoryPreloadRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryPreloadRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


