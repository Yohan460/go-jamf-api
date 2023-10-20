# MobileDevicePurchasing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Purchased** | Pointer to **bool** |  | [optional] 
**Leased** | Pointer to **bool** |  | [optional] 
**PoNumber** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**AppleCareId** | Pointer to **string** |  | [optional] 
**PurchasePrice** | Pointer to **string** |  | [optional] 
**PurchasingAccount** | Pointer to **string** |  | [optional] 
**PoDate** | Pointer to **time.Time** |  | [optional] 
**WarrantyExpiresDate** | Pointer to **time.Time** |  | [optional] 
**LeaseExpiresDate** | Pointer to **time.Time** |  | [optional] 
**LifeExpectancy** | Pointer to **int32** |  | [optional] 
**PurchasingContact** | Pointer to **string** |  | [optional] 
**ExtensionAttributes** | Pointer to [**[]MobileDeviceExtensionAttribute**](MobileDeviceExtensionAttribute.md) |  | [optional] 

## Methods

### NewMobileDevicePurchasing

`func NewMobileDevicePurchasing() *MobileDevicePurchasing`

NewMobileDevicePurchasing instantiates a new MobileDevicePurchasing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDevicePurchasingWithDefaults

`func NewMobileDevicePurchasingWithDefaults() *MobileDevicePurchasing`

NewMobileDevicePurchasingWithDefaults instantiates a new MobileDevicePurchasing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurchased

`func (o *MobileDevicePurchasing) GetPurchased() bool`

GetPurchased returns the Purchased field if non-nil, zero value otherwise.

### GetPurchasedOk

`func (o *MobileDevicePurchasing) GetPurchasedOk() (*bool, bool)`

GetPurchasedOk returns a tuple with the Purchased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchased

`func (o *MobileDevicePurchasing) SetPurchased(v bool)`

SetPurchased sets Purchased field to given value.

### HasPurchased

`func (o *MobileDevicePurchasing) HasPurchased() bool`

HasPurchased returns a boolean if a field has been set.

### GetLeased

`func (o *MobileDevicePurchasing) GetLeased() bool`

GetLeased returns the Leased field if non-nil, zero value otherwise.

### GetLeasedOk

`func (o *MobileDevicePurchasing) GetLeasedOk() (*bool, bool)`

GetLeasedOk returns a tuple with the Leased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeased

`func (o *MobileDevicePurchasing) SetLeased(v bool)`

SetLeased sets Leased field to given value.

### HasLeased

`func (o *MobileDevicePurchasing) HasLeased() bool`

HasLeased returns a boolean if a field has been set.

### GetPoNumber

`func (o *MobileDevicePurchasing) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *MobileDevicePurchasing) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *MobileDevicePurchasing) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.

### HasPoNumber

`func (o *MobileDevicePurchasing) HasPoNumber() bool`

HasPoNumber returns a boolean if a field has been set.

### GetVendor

`func (o *MobileDevicePurchasing) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *MobileDevicePurchasing) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *MobileDevicePurchasing) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *MobileDevicePurchasing) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetAppleCareId

`func (o *MobileDevicePurchasing) GetAppleCareId() string`

GetAppleCareId returns the AppleCareId field if non-nil, zero value otherwise.

### GetAppleCareIdOk

`func (o *MobileDevicePurchasing) GetAppleCareIdOk() (*string, bool)`

GetAppleCareIdOk returns a tuple with the AppleCareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleCareId

`func (o *MobileDevicePurchasing) SetAppleCareId(v string)`

SetAppleCareId sets AppleCareId field to given value.

### HasAppleCareId

`func (o *MobileDevicePurchasing) HasAppleCareId() bool`

HasAppleCareId returns a boolean if a field has been set.

### GetPurchasePrice

`func (o *MobileDevicePurchasing) GetPurchasePrice() string`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *MobileDevicePurchasing) GetPurchasePriceOk() (*string, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *MobileDevicePurchasing) SetPurchasePrice(v string)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *MobileDevicePurchasing) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.

### GetPurchasingAccount

`func (o *MobileDevicePurchasing) GetPurchasingAccount() string`

GetPurchasingAccount returns the PurchasingAccount field if non-nil, zero value otherwise.

### GetPurchasingAccountOk

`func (o *MobileDevicePurchasing) GetPurchasingAccountOk() (*string, bool)`

GetPurchasingAccountOk returns a tuple with the PurchasingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingAccount

`func (o *MobileDevicePurchasing) SetPurchasingAccount(v string)`

SetPurchasingAccount sets PurchasingAccount field to given value.

### HasPurchasingAccount

`func (o *MobileDevicePurchasing) HasPurchasingAccount() bool`

HasPurchasingAccount returns a boolean if a field has been set.

### GetPoDate

`func (o *MobileDevicePurchasing) GetPoDate() time.Time`

GetPoDate returns the PoDate field if non-nil, zero value otherwise.

### GetPoDateOk

`func (o *MobileDevicePurchasing) GetPoDateOk() (*time.Time, bool)`

GetPoDateOk returns a tuple with the PoDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoDate

`func (o *MobileDevicePurchasing) SetPoDate(v time.Time)`

SetPoDate sets PoDate field to given value.

### HasPoDate

`func (o *MobileDevicePurchasing) HasPoDate() bool`

HasPoDate returns a boolean if a field has been set.

### GetWarrantyExpiresDate

`func (o *MobileDevicePurchasing) GetWarrantyExpiresDate() time.Time`

GetWarrantyExpiresDate returns the WarrantyExpiresDate field if non-nil, zero value otherwise.

### GetWarrantyExpiresDateOk

`func (o *MobileDevicePurchasing) GetWarrantyExpiresDateOk() (*time.Time, bool)`

GetWarrantyExpiresDateOk returns a tuple with the WarrantyExpiresDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyExpiresDate

`func (o *MobileDevicePurchasing) SetWarrantyExpiresDate(v time.Time)`

SetWarrantyExpiresDate sets WarrantyExpiresDate field to given value.

### HasWarrantyExpiresDate

`func (o *MobileDevicePurchasing) HasWarrantyExpiresDate() bool`

HasWarrantyExpiresDate returns a boolean if a field has been set.

### GetLeaseExpiresDate

`func (o *MobileDevicePurchasing) GetLeaseExpiresDate() time.Time`

GetLeaseExpiresDate returns the LeaseExpiresDate field if non-nil, zero value otherwise.

### GetLeaseExpiresDateOk

`func (o *MobileDevicePurchasing) GetLeaseExpiresDateOk() (*time.Time, bool)`

GetLeaseExpiresDateOk returns a tuple with the LeaseExpiresDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseExpiresDate

`func (o *MobileDevicePurchasing) SetLeaseExpiresDate(v time.Time)`

SetLeaseExpiresDate sets LeaseExpiresDate field to given value.

### HasLeaseExpiresDate

`func (o *MobileDevicePurchasing) HasLeaseExpiresDate() bool`

HasLeaseExpiresDate returns a boolean if a field has been set.

### GetLifeExpectancy

`func (o *MobileDevicePurchasing) GetLifeExpectancy() int32`

GetLifeExpectancy returns the LifeExpectancy field if non-nil, zero value otherwise.

### GetLifeExpectancyOk

`func (o *MobileDevicePurchasing) GetLifeExpectancyOk() (*int32, bool)`

GetLifeExpectancyOk returns a tuple with the LifeExpectancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeExpectancy

`func (o *MobileDevicePurchasing) SetLifeExpectancy(v int32)`

SetLifeExpectancy sets LifeExpectancy field to given value.

### HasLifeExpectancy

`func (o *MobileDevicePurchasing) HasLifeExpectancy() bool`

HasLifeExpectancy returns a boolean if a field has been set.

### GetPurchasingContact

`func (o *MobileDevicePurchasing) GetPurchasingContact() string`

GetPurchasingContact returns the PurchasingContact field if non-nil, zero value otherwise.

### GetPurchasingContactOk

`func (o *MobileDevicePurchasing) GetPurchasingContactOk() (*string, bool)`

GetPurchasingContactOk returns a tuple with the PurchasingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingContact

`func (o *MobileDevicePurchasing) SetPurchasingContact(v string)`

SetPurchasingContact sets PurchasingContact field to given value.

### HasPurchasingContact

`func (o *MobileDevicePurchasing) HasPurchasingContact() bool`

HasPurchasingContact returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *MobileDevicePurchasing) GetExtensionAttributes() []MobileDeviceExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *MobileDevicePurchasing) GetExtensionAttributesOk() (*[]MobileDeviceExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *MobileDevicePurchasing) SetExtensionAttributes(v []MobileDeviceExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *MobileDevicePurchasing) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


