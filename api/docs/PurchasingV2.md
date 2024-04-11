# PurchasingV2

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
**LifeExpectancy** | Pointer to **int64** |  | [optional] 
**PurchasingContact** | Pointer to **string** |  | [optional] 

## Methods

### NewPurchasingV2

`func NewPurchasingV2() *PurchasingV2`

NewPurchasingV2 instantiates a new PurchasingV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchasingV2WithDefaults

`func NewPurchasingV2WithDefaults() *PurchasingV2`

NewPurchasingV2WithDefaults instantiates a new PurchasingV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurchased

`func (o *PurchasingV2) GetPurchased() bool`

GetPurchased returns the Purchased field if non-nil, zero value otherwise.

### GetPurchasedOk

`func (o *PurchasingV2) GetPurchasedOk() (*bool, bool)`

GetPurchasedOk returns a tuple with the Purchased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchased

`func (o *PurchasingV2) SetPurchased(v bool)`

SetPurchased sets Purchased field to given value.

### HasPurchased

`func (o *PurchasingV2) HasPurchased() bool`

HasPurchased returns a boolean if a field has been set.

### GetLeased

`func (o *PurchasingV2) GetLeased() bool`

GetLeased returns the Leased field if non-nil, zero value otherwise.

### GetLeasedOk

`func (o *PurchasingV2) GetLeasedOk() (*bool, bool)`

GetLeasedOk returns a tuple with the Leased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeased

`func (o *PurchasingV2) SetLeased(v bool)`

SetLeased sets Leased field to given value.

### HasLeased

`func (o *PurchasingV2) HasLeased() bool`

HasLeased returns a boolean if a field has been set.

### GetPoNumber

`func (o *PurchasingV2) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *PurchasingV2) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *PurchasingV2) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.

### HasPoNumber

`func (o *PurchasingV2) HasPoNumber() bool`

HasPoNumber returns a boolean if a field has been set.

### GetVendor

`func (o *PurchasingV2) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *PurchasingV2) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *PurchasingV2) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *PurchasingV2) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetAppleCareId

`func (o *PurchasingV2) GetAppleCareId() string`

GetAppleCareId returns the AppleCareId field if non-nil, zero value otherwise.

### GetAppleCareIdOk

`func (o *PurchasingV2) GetAppleCareIdOk() (*string, bool)`

GetAppleCareIdOk returns a tuple with the AppleCareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleCareId

`func (o *PurchasingV2) SetAppleCareId(v string)`

SetAppleCareId sets AppleCareId field to given value.

### HasAppleCareId

`func (o *PurchasingV2) HasAppleCareId() bool`

HasAppleCareId returns a boolean if a field has been set.

### GetPurchasePrice

`func (o *PurchasingV2) GetPurchasePrice() string`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *PurchasingV2) GetPurchasePriceOk() (*string, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *PurchasingV2) SetPurchasePrice(v string)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *PurchasingV2) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.

### GetPurchasingAccount

`func (o *PurchasingV2) GetPurchasingAccount() string`

GetPurchasingAccount returns the PurchasingAccount field if non-nil, zero value otherwise.

### GetPurchasingAccountOk

`func (o *PurchasingV2) GetPurchasingAccountOk() (*string, bool)`

GetPurchasingAccountOk returns a tuple with the PurchasingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingAccount

`func (o *PurchasingV2) SetPurchasingAccount(v string)`

SetPurchasingAccount sets PurchasingAccount field to given value.

### HasPurchasingAccount

`func (o *PurchasingV2) HasPurchasingAccount() bool`

HasPurchasingAccount returns a boolean if a field has been set.

### GetPoDate

`func (o *PurchasingV2) GetPoDate() time.Time`

GetPoDate returns the PoDate field if non-nil, zero value otherwise.

### GetPoDateOk

`func (o *PurchasingV2) GetPoDateOk() (*time.Time, bool)`

GetPoDateOk returns a tuple with the PoDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoDate

`func (o *PurchasingV2) SetPoDate(v time.Time)`

SetPoDate sets PoDate field to given value.

### HasPoDate

`func (o *PurchasingV2) HasPoDate() bool`

HasPoDate returns a boolean if a field has been set.

### GetWarrantyExpiresDate

`func (o *PurchasingV2) GetWarrantyExpiresDate() time.Time`

GetWarrantyExpiresDate returns the WarrantyExpiresDate field if non-nil, zero value otherwise.

### GetWarrantyExpiresDateOk

`func (o *PurchasingV2) GetWarrantyExpiresDateOk() (*time.Time, bool)`

GetWarrantyExpiresDateOk returns a tuple with the WarrantyExpiresDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyExpiresDate

`func (o *PurchasingV2) SetWarrantyExpiresDate(v time.Time)`

SetWarrantyExpiresDate sets WarrantyExpiresDate field to given value.

### HasWarrantyExpiresDate

`func (o *PurchasingV2) HasWarrantyExpiresDate() bool`

HasWarrantyExpiresDate returns a boolean if a field has been set.

### GetLeaseExpiresDate

`func (o *PurchasingV2) GetLeaseExpiresDate() time.Time`

GetLeaseExpiresDate returns the LeaseExpiresDate field if non-nil, zero value otherwise.

### GetLeaseExpiresDateOk

`func (o *PurchasingV2) GetLeaseExpiresDateOk() (*time.Time, bool)`

GetLeaseExpiresDateOk returns a tuple with the LeaseExpiresDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseExpiresDate

`func (o *PurchasingV2) SetLeaseExpiresDate(v time.Time)`

SetLeaseExpiresDate sets LeaseExpiresDate field to given value.

### HasLeaseExpiresDate

`func (o *PurchasingV2) HasLeaseExpiresDate() bool`

HasLeaseExpiresDate returns a boolean if a field has been set.

### GetLifeExpectancy

`func (o *PurchasingV2) GetLifeExpectancy() int64`

GetLifeExpectancy returns the LifeExpectancy field if non-nil, zero value otherwise.

### GetLifeExpectancyOk

`func (o *PurchasingV2) GetLifeExpectancyOk() (*int64, bool)`

GetLifeExpectancyOk returns a tuple with the LifeExpectancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeExpectancy

`func (o *PurchasingV2) SetLifeExpectancy(v int64)`

SetLifeExpectancy sets LifeExpectancy field to given value.

### HasLifeExpectancy

`func (o *PurchasingV2) HasLifeExpectancy() bool`

HasLifeExpectancy returns a boolean if a field has been set.

### GetPurchasingContact

`func (o *PurchasingV2) GetPurchasingContact() string`

GetPurchasingContact returns the PurchasingContact field if non-nil, zero value otherwise.

### GetPurchasingContactOk

`func (o *PurchasingV2) GetPurchasingContactOk() (*string, bool)`

GetPurchasingContactOk returns a tuple with the PurchasingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingContact

`func (o *PurchasingV2) SetPurchasingContact(v string)`

SetPurchasingContact sets PurchasingContact field to given value.

### HasPurchasingContact

`func (o *PurchasingV2) HasPurchasingContact() bool`

HasPurchasingContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


