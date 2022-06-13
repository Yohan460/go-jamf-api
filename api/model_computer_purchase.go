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

// ComputerPurchase struct for ComputerPurchase
type ComputerPurchase struct {
	Leased *bool `json:"leased,omitempty"`
	Purchased *bool `json:"purchased,omitempty"`
	PoNumber *string `json:"poNumber,omitempty"`
	PoDate *string `json:"poDate,omitempty"`
	Vendor *string `json:"vendor,omitempty"`
	WarrantyDate *string `json:"warrantyDate,omitempty"`
	AppleCareId *string `json:"appleCareId,omitempty"`
	LeaseDate *string `json:"leaseDate,omitempty"`
	PurchasePrice *string `json:"purchasePrice,omitempty"`
	LifeExpectancy *int32 `json:"lifeExpectancy,omitempty"`
	PurchasingAccount *string `json:"purchasingAccount,omitempty"`
	PurchasingContact *string `json:"purchasingContact,omitempty"`
	ExtensionAttributes []ComputerExtensionAttribute `json:"extensionAttributes,omitempty"`
}

// NewComputerPurchase instantiates a new ComputerPurchase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputerPurchase() *ComputerPurchase {
	this := ComputerPurchase{}
	return &this
}

// NewComputerPurchaseWithDefaults instantiates a new ComputerPurchase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputerPurchaseWithDefaults() *ComputerPurchase {
	this := ComputerPurchase{}
	return &this
}

// GetLeased returns the Leased field value if set, zero value otherwise.
func (o *ComputerPurchase) GetLeased() bool {
	if o == nil || o.Leased == nil {
		var ret bool
		return ret
	}
	return *o.Leased
}

// GetLeasedOk returns a tuple with the Leased field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPurchase) GetLeasedOk() (*bool, bool) {
	if o == nil || o.Leased == nil {
		return nil, false
	}
	return o.Leased, true
}

// HasLeased returns a boolean if a field has been set.
func (o *ComputerPurchase) HasLeased() bool {
	if o != nil && o.Leased != nil {
		return true
	}

	return false
}

// SetLeased gets a reference to the given bool and assigns it to the Leased field.
func (o *ComputerPurchase) SetLeased(v bool) {
	o.Leased = &v
}

// GetPurchased returns the Purchased field value if set, zero value otherwise.
func (o *ComputerPurchase) GetPurchased() bool {
	if o == nil || o.Purchased == nil {
		var ret bool
		return ret
	}
	return *o.Purchased
}

// GetPurchasedOk returns a tuple with the Purchased field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPurchase) GetPurchasedOk() (*bool, bool) {
	if o == nil || o.Purchased == nil {
		return nil, false
	}
	return o.Purchased, true
}

// HasPurchased returns a boolean if a field has been set.
func (o *ComputerPurchase) HasPurchased() bool {
	if o != nil && o.Purchased != nil {
		return true
	}

	return false
}

// SetPurchased gets a reference to the given bool and assigns it to the Purchased field.
func (o *ComputerPurchase) SetPurchased(v bool) {
	o.Purchased = &v
}

// GetPoNumber returns the PoNumber field value if set, zero value otherwise.
func (o *ComputerPurchase) GetPoNumber() string {
	if o == nil || o.PoNumber == nil {
		var ret string
		return ret
	}
	return *o.PoNumber
}

// GetPoNumberOk returns a tuple with the PoNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPurchase) GetPoNumberOk() (*string, bool) {
	if o == nil || o.PoNumber == nil {
		return nil, false
	}
	return o.PoNumber, true
}

// HasPoNumber returns a boolean if a field has been set.
func (o *ComputerPurchase) HasPoNumber() bool {
	if o != nil && o.PoNumber != nil {
		return true
	}

	return false
}

// SetPoNumber gets a reference to the given string and assigns it to the PoNumber field.
func (o *ComputerPurchase) SetPoNumber(v string) {
	o.PoNumber = &v
}

// GetPoDate returns the PoDate field value if set, zero value otherwise.
func (o *ComputerPurchase) GetPoDate() string {
	if o == nil || o.PoDate == nil {
		var ret string
		return ret
	}
	return *o.PoDate
}

// GetPoDateOk returns a tuple with the PoDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPurchase) GetPoDateOk() (*string, bool) {
	if o == nil || o.PoDate == nil {
		return nil, false
	}
	return o.PoDate, true
}

// HasPoDate returns a boolean if a field has been set.
func (o *ComputerPurchase) HasPoDate() bool {
	if o != nil && o.PoDate != nil {
		return true
	}

	return false
}

// SetPoDate gets a reference to the given string and assigns it to the PoDate field.
func (o *ComputerPurchase) SetPoDate(v string) {
	o.PoDate = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *ComputerPurchase) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPurchase) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *ComputerPurchase) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *ComputerPurchase) SetVendor(v string) {
	o.Vendor = &v
}

// GetWarrantyDate returns the WarrantyDate field value if set, zero value otherwise.
func (o *ComputerPurchase) GetWarrantyDate() string {
	if o == nil || o.WarrantyDate == nil {
		var ret string
		return ret
	}
	return *o.WarrantyDate
}

// GetWarrantyDateOk returns a tuple with the WarrantyDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPurchase) GetWarrantyDateOk() (*string, bool) {
	if o == nil || o.WarrantyDate == nil {
		return nil, false
	}
	return o.WarrantyDate, true
}

// HasWarrantyDate returns a boolean if a field has been set.
func (o *ComputerPurchase) HasWarrantyDate() bool {
	if o != nil && o.WarrantyDate != nil {
		return true
	}

	return false
}

// SetWarrantyDate gets a reference to the given string and assigns it to the WarrantyDate field.
func (o *ComputerPurchase) SetWarrantyDate(v string) {
	o.WarrantyDate = &v
}

// GetAppleCareId returns the AppleCareId field value if set, zero value otherwise.
func (o *ComputerPurchase) GetAppleCareId() string {
	if o == nil || o.AppleCareId == nil {
		var ret string
		return ret
	}
	return *o.AppleCareId
}

// GetAppleCareIdOk returns a tuple with the AppleCareId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPurchase) GetAppleCareIdOk() (*string, bool) {
	if o == nil || o.AppleCareId == nil {
		return nil, false
	}
	return o.AppleCareId, true
}

// HasAppleCareId returns a boolean if a field has been set.
func (o *ComputerPurchase) HasAppleCareId() bool {
	if o != nil && o.AppleCareId != nil {
		return true
	}

	return false
}

// SetAppleCareId gets a reference to the given string and assigns it to the AppleCareId field.
func (o *ComputerPurchase) SetAppleCareId(v string) {
	o.AppleCareId = &v
}

// GetLeaseDate returns the LeaseDate field value if set, zero value otherwise.
func (o *ComputerPurchase) GetLeaseDate() string {
	if o == nil || o.LeaseDate == nil {
		var ret string
		return ret
	}
	return *o.LeaseDate
}

// GetLeaseDateOk returns a tuple with the LeaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPurchase) GetLeaseDateOk() (*string, bool) {
	if o == nil || o.LeaseDate == nil {
		return nil, false
	}
	return o.LeaseDate, true
}

// HasLeaseDate returns a boolean if a field has been set.
func (o *ComputerPurchase) HasLeaseDate() bool {
	if o != nil && o.LeaseDate != nil {
		return true
	}

	return false
}

// SetLeaseDate gets a reference to the given string and assigns it to the LeaseDate field.
func (o *ComputerPurchase) SetLeaseDate(v string) {
	o.LeaseDate = &v
}

// GetPurchasePrice returns the PurchasePrice field value if set, zero value otherwise.
func (o *ComputerPurchase) GetPurchasePrice() string {
	if o == nil || o.PurchasePrice == nil {
		var ret string
		return ret
	}
	return *o.PurchasePrice
}

// GetPurchasePriceOk returns a tuple with the PurchasePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPurchase) GetPurchasePriceOk() (*string, bool) {
	if o == nil || o.PurchasePrice == nil {
		return nil, false
	}
	return o.PurchasePrice, true
}

// HasPurchasePrice returns a boolean if a field has been set.
func (o *ComputerPurchase) HasPurchasePrice() bool {
	if o != nil && o.PurchasePrice != nil {
		return true
	}

	return false
}

// SetPurchasePrice gets a reference to the given string and assigns it to the PurchasePrice field.
func (o *ComputerPurchase) SetPurchasePrice(v string) {
	o.PurchasePrice = &v
}

// GetLifeExpectancy returns the LifeExpectancy field value if set, zero value otherwise.
func (o *ComputerPurchase) GetLifeExpectancy() int32 {
	if o == nil || o.LifeExpectancy == nil {
		var ret int32
		return ret
	}
	return *o.LifeExpectancy
}

// GetLifeExpectancyOk returns a tuple with the LifeExpectancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPurchase) GetLifeExpectancyOk() (*int32, bool) {
	if o == nil || o.LifeExpectancy == nil {
		return nil, false
	}
	return o.LifeExpectancy, true
}

// HasLifeExpectancy returns a boolean if a field has been set.
func (o *ComputerPurchase) HasLifeExpectancy() bool {
	if o != nil && o.LifeExpectancy != nil {
		return true
	}

	return false
}

// SetLifeExpectancy gets a reference to the given int32 and assigns it to the LifeExpectancy field.
func (o *ComputerPurchase) SetLifeExpectancy(v int32) {
	o.LifeExpectancy = &v
}

// GetPurchasingAccount returns the PurchasingAccount field value if set, zero value otherwise.
func (o *ComputerPurchase) GetPurchasingAccount() string {
	if o == nil || o.PurchasingAccount == nil {
		var ret string
		return ret
	}
	return *o.PurchasingAccount
}

// GetPurchasingAccountOk returns a tuple with the PurchasingAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPurchase) GetPurchasingAccountOk() (*string, bool) {
	if o == nil || o.PurchasingAccount == nil {
		return nil, false
	}
	return o.PurchasingAccount, true
}

// HasPurchasingAccount returns a boolean if a field has been set.
func (o *ComputerPurchase) HasPurchasingAccount() bool {
	if o != nil && o.PurchasingAccount != nil {
		return true
	}

	return false
}

// SetPurchasingAccount gets a reference to the given string and assigns it to the PurchasingAccount field.
func (o *ComputerPurchase) SetPurchasingAccount(v string) {
	o.PurchasingAccount = &v
}

// GetPurchasingContact returns the PurchasingContact field value if set, zero value otherwise.
func (o *ComputerPurchase) GetPurchasingContact() string {
	if o == nil || o.PurchasingContact == nil {
		var ret string
		return ret
	}
	return *o.PurchasingContact
}

// GetPurchasingContactOk returns a tuple with the PurchasingContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPurchase) GetPurchasingContactOk() (*string, bool) {
	if o == nil || o.PurchasingContact == nil {
		return nil, false
	}
	return o.PurchasingContact, true
}

// HasPurchasingContact returns a boolean if a field has been set.
func (o *ComputerPurchase) HasPurchasingContact() bool {
	if o != nil && o.PurchasingContact != nil {
		return true
	}

	return false
}

// SetPurchasingContact gets a reference to the given string and assigns it to the PurchasingContact field.
func (o *ComputerPurchase) SetPurchasingContact(v string) {
	o.PurchasingContact = &v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *ComputerPurchase) GetExtensionAttributes() []ComputerExtensionAttribute {
	if o == nil || o.ExtensionAttributes == nil {
		var ret []ComputerExtensionAttribute
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputerPurchase) GetExtensionAttributesOk() ([]ComputerExtensionAttribute, bool) {
	if o == nil || o.ExtensionAttributes == nil {
		return nil, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *ComputerPurchase) HasExtensionAttributes() bool {
	if o != nil && o.ExtensionAttributes != nil {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given []ComputerExtensionAttribute and assigns it to the ExtensionAttributes field.
func (o *ComputerPurchase) SetExtensionAttributes(v []ComputerExtensionAttribute) {
	o.ExtensionAttributes = v
}

func (o ComputerPurchase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Leased != nil {
		toSerialize["leased"] = o.Leased
	}
	if o.Purchased != nil {
		toSerialize["purchased"] = o.Purchased
	}
	if o.PoNumber != nil {
		toSerialize["poNumber"] = o.PoNumber
	}
	if o.PoDate != nil {
		toSerialize["poDate"] = o.PoDate
	}
	if o.Vendor != nil {
		toSerialize["vendor"] = o.Vendor
	}
	if o.WarrantyDate != nil {
		toSerialize["warrantyDate"] = o.WarrantyDate
	}
	if o.AppleCareId != nil {
		toSerialize["appleCareId"] = o.AppleCareId
	}
	if o.LeaseDate != nil {
		toSerialize["leaseDate"] = o.LeaseDate
	}
	if o.PurchasePrice != nil {
		toSerialize["purchasePrice"] = o.PurchasePrice
	}
	if o.LifeExpectancy != nil {
		toSerialize["lifeExpectancy"] = o.LifeExpectancy
	}
	if o.PurchasingAccount != nil {
		toSerialize["purchasingAccount"] = o.PurchasingAccount
	}
	if o.PurchasingContact != nil {
		toSerialize["purchasingContact"] = o.PurchasingContact
	}
	if o.ExtensionAttributes != nil {
		toSerialize["extensionAttributes"] = o.ExtensionAttributes
	}
	return json.Marshal(toSerialize)
}

type NullableComputerPurchase struct {
	value *ComputerPurchase
	isSet bool
}

func (v NullableComputerPurchase) Get() *ComputerPurchase {
	return v.value
}

func (v *NullableComputerPurchase) Set(val *ComputerPurchase) {
	v.value = val
	v.isSet = true
}

func (v NullableComputerPurchase) IsSet() bool {
	return v.isSet
}

func (v *NullableComputerPurchase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputerPurchase(val *ComputerPurchase) *NullableComputerPurchase {
	return &NullableComputerPurchase{value: val, isSet: true}
}

func (v NullableComputerPurchase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputerPurchase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


