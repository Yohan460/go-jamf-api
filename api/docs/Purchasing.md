# Purchasing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPurchased** | Pointer to **bool** |  | [optional] 
**IsLeased** | Pointer to **bool** |  | [optional] 
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

### NewPurchasing

`func NewPurchasing() *Purchasing`

NewPurchasing instantiates a new Purchasing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchasingWithDefaults

`func NewPurchasingWithDefaults() *Purchasing`

NewPurchasingWithDefaults instantiates a new Purchasing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPurchased

`func (o *Purchasing) GetIsPurchased() bool`

GetIsPurchased returns the IsPurchased field if non-nil, zero value otherwise.

### GetIsPurchasedOk

`func (o *Purchasing) GetIsPurchasedOk() (*bool, bool)`

GetIsPurchasedOk returns a tuple with the IsPurchased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPurchased

`func (o *Purchasing) SetIsPurchased(v bool)`

SetIsPurchased sets IsPurchased field to given value.

### HasIsPurchased

`func (o *Purchasing) HasIsPurchased() bool`

HasIsPurchased returns a boolean if a field has been set.

### GetIsLeased

`func (o *Purchasing) GetIsLeased() bool`

GetIsLeased returns the IsLeased field if non-nil, zero value otherwise.

### GetIsLeasedOk

`func (o *Purchasing) GetIsLeasedOk() (*bool, bool)`

GetIsLeasedOk returns a tuple with the IsLeased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLeased

`func (o *Purchasing) SetIsLeased(v bool)`

SetIsLeased sets IsLeased field to given value.

### HasIsLeased

`func (o *Purchasing) HasIsLeased() bool`

HasIsLeased returns a boolean if a field has been set.

### GetPoNumber

`func (o *Purchasing) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *Purchasing) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *Purchasing) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.

### HasPoNumber

`func (o *Purchasing) HasPoNumber() bool`

HasPoNumber returns a boolean if a field has been set.

### GetVendor

`func (o *Purchasing) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *Purchasing) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *Purchasing) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *Purchasing) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetAppleCareId

`func (o *Purchasing) GetAppleCareId() string`

GetAppleCareId returns the AppleCareId field if non-nil, zero value otherwise.

### GetAppleCareIdOk

`func (o *Purchasing) GetAppleCareIdOk() (*string, bool)`

GetAppleCareIdOk returns a tuple with the AppleCareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleCareId

`func (o *Purchasing) SetAppleCareId(v string)`

SetAppleCareId sets AppleCareId field to given value.

### HasAppleCareId

`func (o *Purchasing) HasAppleCareId() bool`

HasAppleCareId returns a boolean if a field has been set.

### GetPurchasePrice

`func (o *Purchasing) GetPurchasePrice() string`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *Purchasing) GetPurchasePriceOk() (*string, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *Purchasing) SetPurchasePrice(v string)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *Purchasing) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.

### GetPurchasingAccount

`func (o *Purchasing) GetPurchasingAccount() string`

GetPurchasingAccount returns the PurchasingAccount field if non-nil, zero value otherwise.

### GetPurchasingAccountOk

`func (o *Purchasing) GetPurchasingAccountOk() (*string, bool)`

GetPurchasingAccountOk returns a tuple with the PurchasingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingAccount

`func (o *Purchasing) SetPurchasingAccount(v string)`

SetPurchasingAccount sets PurchasingAccount field to given value.

### HasPurchasingAccount

`func (o *Purchasing) HasPurchasingAccount() bool`

HasPurchasingAccount returns a boolean if a field has been set.

### GetPoDate

`func (o *Purchasing) GetPoDate() time.Time`

GetPoDate returns the PoDate field if non-nil, zero value otherwise.

### GetPoDateOk

`func (o *Purchasing) GetPoDateOk() (*time.Time, bool)`

GetPoDateOk returns a tuple with the PoDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoDate

`func (o *Purchasing) SetPoDate(v time.Time)`

SetPoDate sets PoDate field to given value.

### HasPoDate

`func (o *Purchasing) HasPoDate() bool`

HasPoDate returns a boolean if a field has been set.

### GetWarrantyExpiresDate

`func (o *Purchasing) GetWarrantyExpiresDate() time.Time`

GetWarrantyExpiresDate returns the WarrantyExpiresDate field if non-nil, zero value otherwise.

### GetWarrantyExpiresDateOk

`func (o *Purchasing) GetWarrantyExpiresDateOk() (*time.Time, bool)`

GetWarrantyExpiresDateOk returns a tuple with the WarrantyExpiresDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyExpiresDate

`func (o *Purchasing) SetWarrantyExpiresDate(v time.Time)`

SetWarrantyExpiresDate sets WarrantyExpiresDate field to given value.

### HasWarrantyExpiresDate

`func (o *Purchasing) HasWarrantyExpiresDate() bool`

HasWarrantyExpiresDate returns a boolean if a field has been set.

### GetLeaseExpiresDate

`func (o *Purchasing) GetLeaseExpiresDate() time.Time`

GetLeaseExpiresDate returns the LeaseExpiresDate field if non-nil, zero value otherwise.

### GetLeaseExpiresDateOk

`func (o *Purchasing) GetLeaseExpiresDateOk() (*time.Time, bool)`

GetLeaseExpiresDateOk returns a tuple with the LeaseExpiresDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseExpiresDate

`func (o *Purchasing) SetLeaseExpiresDate(v time.Time)`

SetLeaseExpiresDate sets LeaseExpiresDate field to given value.

### HasLeaseExpiresDate

`func (o *Purchasing) HasLeaseExpiresDate() bool`

HasLeaseExpiresDate returns a boolean if a field has been set.

### GetLifeExpectancy

`func (o *Purchasing) GetLifeExpectancy() int64`

GetLifeExpectancy returns the LifeExpectancy field if non-nil, zero value otherwise.

### GetLifeExpectancyOk

`func (o *Purchasing) GetLifeExpectancyOk() (*int64, bool)`

GetLifeExpectancyOk returns a tuple with the LifeExpectancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeExpectancy

`func (o *Purchasing) SetLifeExpectancy(v int64)`

SetLifeExpectancy sets LifeExpectancy field to given value.

### HasLifeExpectancy

`func (o *Purchasing) HasLifeExpectancy() bool`

HasLifeExpectancy returns a boolean if a field has been set.

### GetPurchasingContact

`func (o *Purchasing) GetPurchasingContact() string`

GetPurchasingContact returns the PurchasingContact field if non-nil, zero value otherwise.

### GetPurchasingContactOk

`func (o *Purchasing) GetPurchasingContactOk() (*string, bool)`

GetPurchasingContactOk returns a tuple with the PurchasingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingContact

`func (o *Purchasing) SetPurchasingContact(v string)`

SetPurchasingContact sets PurchasingContact field to given value.

### HasPurchasingContact

`func (o *Purchasing) HasPurchasingContact() bool`

HasPurchasingContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


