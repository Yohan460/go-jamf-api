# ComputerPurchaseCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Leased** | Pointer to **NullableBool** |  | [optional] 
**Purchased** | Pointer to **NullableBool** |  | [optional] 
**PoNumber** | Pointer to **NullableString** |  | [optional] 
**PoDate** | Pointer to **NullableString** |  | [optional] 
**Vendor** | Pointer to **NullableString** |  | [optional] 
**WarrantyDate** | Pointer to **NullableString** |  | [optional] 
**AppleCareId** | Pointer to **NullableString** |  | [optional] 
**LeaseDate** | Pointer to **NullableString** |  | [optional] 
**PurchasePrice** | Pointer to **NullableString** |  | [optional] 
**LifeExpectancy** | Pointer to **NullableInt64** |  | [optional] 
**PurchasingAccount** | Pointer to **NullableString** |  | [optional] 
**PurchasingContact** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewComputerPurchaseCreate

`func NewComputerPurchaseCreate() *ComputerPurchaseCreate`

NewComputerPurchaseCreate instantiates a new ComputerPurchaseCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerPurchaseCreateWithDefaults

`func NewComputerPurchaseCreateWithDefaults() *ComputerPurchaseCreate`

NewComputerPurchaseCreateWithDefaults instantiates a new ComputerPurchaseCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeased

`func (o *ComputerPurchaseCreate) GetLeased() bool`

GetLeased returns the Leased field if non-nil, zero value otherwise.

### GetLeasedOk

`func (o *ComputerPurchaseCreate) GetLeasedOk() (*bool, bool)`

GetLeasedOk returns a tuple with the Leased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeased

`func (o *ComputerPurchaseCreate) SetLeased(v bool)`

SetLeased sets Leased field to given value.

### HasLeased

`func (o *ComputerPurchaseCreate) HasLeased() bool`

HasLeased returns a boolean if a field has been set.

### SetLeasedNil

`func (o *ComputerPurchaseCreate) SetLeasedNil(b bool)`

 SetLeasedNil sets the value for Leased to be an explicit nil

### UnsetLeased
`func (o *ComputerPurchaseCreate) UnsetLeased()`

UnsetLeased ensures that no value is present for Leased, not even an explicit nil
### GetPurchased

`func (o *ComputerPurchaseCreate) GetPurchased() bool`

GetPurchased returns the Purchased field if non-nil, zero value otherwise.

### GetPurchasedOk

`func (o *ComputerPurchaseCreate) GetPurchasedOk() (*bool, bool)`

GetPurchasedOk returns a tuple with the Purchased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchased

`func (o *ComputerPurchaseCreate) SetPurchased(v bool)`

SetPurchased sets Purchased field to given value.

### HasPurchased

`func (o *ComputerPurchaseCreate) HasPurchased() bool`

HasPurchased returns a boolean if a field has been set.

### SetPurchasedNil

`func (o *ComputerPurchaseCreate) SetPurchasedNil(b bool)`

 SetPurchasedNil sets the value for Purchased to be an explicit nil

### UnsetPurchased
`func (o *ComputerPurchaseCreate) UnsetPurchased()`

UnsetPurchased ensures that no value is present for Purchased, not even an explicit nil
### GetPoNumber

`func (o *ComputerPurchaseCreate) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *ComputerPurchaseCreate) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *ComputerPurchaseCreate) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.

### HasPoNumber

`func (o *ComputerPurchaseCreate) HasPoNumber() bool`

HasPoNumber returns a boolean if a field has been set.

### SetPoNumberNil

`func (o *ComputerPurchaseCreate) SetPoNumberNil(b bool)`

 SetPoNumberNil sets the value for PoNumber to be an explicit nil

### UnsetPoNumber
`func (o *ComputerPurchaseCreate) UnsetPoNumber()`

UnsetPoNumber ensures that no value is present for PoNumber, not even an explicit nil
### GetPoDate

`func (o *ComputerPurchaseCreate) GetPoDate() string`

GetPoDate returns the PoDate field if non-nil, zero value otherwise.

### GetPoDateOk

`func (o *ComputerPurchaseCreate) GetPoDateOk() (*string, bool)`

GetPoDateOk returns a tuple with the PoDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoDate

`func (o *ComputerPurchaseCreate) SetPoDate(v string)`

SetPoDate sets PoDate field to given value.

### HasPoDate

`func (o *ComputerPurchaseCreate) HasPoDate() bool`

HasPoDate returns a boolean if a field has been set.

### SetPoDateNil

`func (o *ComputerPurchaseCreate) SetPoDateNil(b bool)`

 SetPoDateNil sets the value for PoDate to be an explicit nil

### UnsetPoDate
`func (o *ComputerPurchaseCreate) UnsetPoDate()`

UnsetPoDate ensures that no value is present for PoDate, not even an explicit nil
### GetVendor

`func (o *ComputerPurchaseCreate) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ComputerPurchaseCreate) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ComputerPurchaseCreate) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ComputerPurchaseCreate) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### SetVendorNil

`func (o *ComputerPurchaseCreate) SetVendorNil(b bool)`

 SetVendorNil sets the value for Vendor to be an explicit nil

### UnsetVendor
`func (o *ComputerPurchaseCreate) UnsetVendor()`

UnsetVendor ensures that no value is present for Vendor, not even an explicit nil
### GetWarrantyDate

`func (o *ComputerPurchaseCreate) GetWarrantyDate() string`

GetWarrantyDate returns the WarrantyDate field if non-nil, zero value otherwise.

### GetWarrantyDateOk

`func (o *ComputerPurchaseCreate) GetWarrantyDateOk() (*string, bool)`

GetWarrantyDateOk returns a tuple with the WarrantyDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyDate

`func (o *ComputerPurchaseCreate) SetWarrantyDate(v string)`

SetWarrantyDate sets WarrantyDate field to given value.

### HasWarrantyDate

`func (o *ComputerPurchaseCreate) HasWarrantyDate() bool`

HasWarrantyDate returns a boolean if a field has been set.

### SetWarrantyDateNil

`func (o *ComputerPurchaseCreate) SetWarrantyDateNil(b bool)`

 SetWarrantyDateNil sets the value for WarrantyDate to be an explicit nil

### UnsetWarrantyDate
`func (o *ComputerPurchaseCreate) UnsetWarrantyDate()`

UnsetWarrantyDate ensures that no value is present for WarrantyDate, not even an explicit nil
### GetAppleCareId

`func (o *ComputerPurchaseCreate) GetAppleCareId() string`

GetAppleCareId returns the AppleCareId field if non-nil, zero value otherwise.

### GetAppleCareIdOk

`func (o *ComputerPurchaseCreate) GetAppleCareIdOk() (*string, bool)`

GetAppleCareIdOk returns a tuple with the AppleCareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleCareId

`func (o *ComputerPurchaseCreate) SetAppleCareId(v string)`

SetAppleCareId sets AppleCareId field to given value.

### HasAppleCareId

`func (o *ComputerPurchaseCreate) HasAppleCareId() bool`

HasAppleCareId returns a boolean if a field has been set.

### SetAppleCareIdNil

`func (o *ComputerPurchaseCreate) SetAppleCareIdNil(b bool)`

 SetAppleCareIdNil sets the value for AppleCareId to be an explicit nil

### UnsetAppleCareId
`func (o *ComputerPurchaseCreate) UnsetAppleCareId()`

UnsetAppleCareId ensures that no value is present for AppleCareId, not even an explicit nil
### GetLeaseDate

`func (o *ComputerPurchaseCreate) GetLeaseDate() string`

GetLeaseDate returns the LeaseDate field if non-nil, zero value otherwise.

### GetLeaseDateOk

`func (o *ComputerPurchaseCreate) GetLeaseDateOk() (*string, bool)`

GetLeaseDateOk returns a tuple with the LeaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseDate

`func (o *ComputerPurchaseCreate) SetLeaseDate(v string)`

SetLeaseDate sets LeaseDate field to given value.

### HasLeaseDate

`func (o *ComputerPurchaseCreate) HasLeaseDate() bool`

HasLeaseDate returns a boolean if a field has been set.

### SetLeaseDateNil

`func (o *ComputerPurchaseCreate) SetLeaseDateNil(b bool)`

 SetLeaseDateNil sets the value for LeaseDate to be an explicit nil

### UnsetLeaseDate
`func (o *ComputerPurchaseCreate) UnsetLeaseDate()`

UnsetLeaseDate ensures that no value is present for LeaseDate, not even an explicit nil
### GetPurchasePrice

`func (o *ComputerPurchaseCreate) GetPurchasePrice() string`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *ComputerPurchaseCreate) GetPurchasePriceOk() (*string, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *ComputerPurchaseCreate) SetPurchasePrice(v string)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *ComputerPurchaseCreate) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.

### SetPurchasePriceNil

`func (o *ComputerPurchaseCreate) SetPurchasePriceNil(b bool)`

 SetPurchasePriceNil sets the value for PurchasePrice to be an explicit nil

### UnsetPurchasePrice
`func (o *ComputerPurchaseCreate) UnsetPurchasePrice()`

UnsetPurchasePrice ensures that no value is present for PurchasePrice, not even an explicit nil
### GetLifeExpectancy

`func (o *ComputerPurchaseCreate) GetLifeExpectancy() int64`

GetLifeExpectancy returns the LifeExpectancy field if non-nil, zero value otherwise.

### GetLifeExpectancyOk

`func (o *ComputerPurchaseCreate) GetLifeExpectancyOk() (*int64, bool)`

GetLifeExpectancyOk returns a tuple with the LifeExpectancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeExpectancy

`func (o *ComputerPurchaseCreate) SetLifeExpectancy(v int64)`

SetLifeExpectancy sets LifeExpectancy field to given value.

### HasLifeExpectancy

`func (o *ComputerPurchaseCreate) HasLifeExpectancy() bool`

HasLifeExpectancy returns a boolean if a field has been set.

### SetLifeExpectancyNil

`func (o *ComputerPurchaseCreate) SetLifeExpectancyNil(b bool)`

 SetLifeExpectancyNil sets the value for LifeExpectancy to be an explicit nil

### UnsetLifeExpectancy
`func (o *ComputerPurchaseCreate) UnsetLifeExpectancy()`

UnsetLifeExpectancy ensures that no value is present for LifeExpectancy, not even an explicit nil
### GetPurchasingAccount

`func (o *ComputerPurchaseCreate) GetPurchasingAccount() string`

GetPurchasingAccount returns the PurchasingAccount field if non-nil, zero value otherwise.

### GetPurchasingAccountOk

`func (o *ComputerPurchaseCreate) GetPurchasingAccountOk() (*string, bool)`

GetPurchasingAccountOk returns a tuple with the PurchasingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingAccount

`func (o *ComputerPurchaseCreate) SetPurchasingAccount(v string)`

SetPurchasingAccount sets PurchasingAccount field to given value.

### HasPurchasingAccount

`func (o *ComputerPurchaseCreate) HasPurchasingAccount() bool`

HasPurchasingAccount returns a boolean if a field has been set.

### SetPurchasingAccountNil

`func (o *ComputerPurchaseCreate) SetPurchasingAccountNil(b bool)`

 SetPurchasingAccountNil sets the value for PurchasingAccount to be an explicit nil

### UnsetPurchasingAccount
`func (o *ComputerPurchaseCreate) UnsetPurchasingAccount()`

UnsetPurchasingAccount ensures that no value is present for PurchasingAccount, not even an explicit nil
### GetPurchasingContact

`func (o *ComputerPurchaseCreate) GetPurchasingContact() string`

GetPurchasingContact returns the PurchasingContact field if non-nil, zero value otherwise.

### GetPurchasingContactOk

`func (o *ComputerPurchaseCreate) GetPurchasingContactOk() (*string, bool)`

GetPurchasingContactOk returns a tuple with the PurchasingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingContact

`func (o *ComputerPurchaseCreate) SetPurchasingContact(v string)`

SetPurchasingContact sets PurchasingContact field to given value.

### HasPurchasingContact

`func (o *ComputerPurchaseCreate) HasPurchasingContact() bool`

HasPurchasingContact returns a boolean if a field has been set.

### SetPurchasingContactNil

`func (o *ComputerPurchaseCreate) SetPurchasingContactNil(b bool)`

 SetPurchasingContactNil sets the value for PurchasingContact to be an explicit nil

### UnsetPurchasingContact
`func (o *ComputerPurchaseCreate) UnsetPurchasingContact()`

UnsetPurchasingContact ensures that no value is present for PurchasingContact, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


