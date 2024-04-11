# ComputerPurchase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Leased** | Pointer to **bool** |  | [optional] 
**Purchased** | Pointer to **bool** |  | [optional] 
**PoNumber** | Pointer to **string** |  | [optional] 
**PoDate** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**WarrantyDate** | Pointer to **string** |  | [optional] 
**AppleCareId** | Pointer to **string** |  | [optional] 
**LeaseDate** | Pointer to **string** |  | [optional] 
**PurchasePrice** | Pointer to **string** |  | [optional] 
**LifeExpectancy** | Pointer to **int64** |  | [optional] 
**PurchasingAccount** | Pointer to **string** |  | [optional] 
**PurchasingContact** | Pointer to **string** |  | [optional] 
**ExtensionAttributes** | Pointer to [**[]ComputerExtensionAttribute**](ComputerExtensionAttribute.md) |  | [optional] 

## Methods

### NewComputerPurchase

`func NewComputerPurchase() *ComputerPurchase`

NewComputerPurchase instantiates a new ComputerPurchase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerPurchaseWithDefaults

`func NewComputerPurchaseWithDefaults() *ComputerPurchase`

NewComputerPurchaseWithDefaults instantiates a new ComputerPurchase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeased

`func (o *ComputerPurchase) GetLeased() bool`

GetLeased returns the Leased field if non-nil, zero value otherwise.

### GetLeasedOk

`func (o *ComputerPurchase) GetLeasedOk() (*bool, bool)`

GetLeasedOk returns a tuple with the Leased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeased

`func (o *ComputerPurchase) SetLeased(v bool)`

SetLeased sets Leased field to given value.

### HasLeased

`func (o *ComputerPurchase) HasLeased() bool`

HasLeased returns a boolean if a field has been set.

### GetPurchased

`func (o *ComputerPurchase) GetPurchased() bool`

GetPurchased returns the Purchased field if non-nil, zero value otherwise.

### GetPurchasedOk

`func (o *ComputerPurchase) GetPurchasedOk() (*bool, bool)`

GetPurchasedOk returns a tuple with the Purchased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchased

`func (o *ComputerPurchase) SetPurchased(v bool)`

SetPurchased sets Purchased field to given value.

### HasPurchased

`func (o *ComputerPurchase) HasPurchased() bool`

HasPurchased returns a boolean if a field has been set.

### GetPoNumber

`func (o *ComputerPurchase) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *ComputerPurchase) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *ComputerPurchase) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.

### HasPoNumber

`func (o *ComputerPurchase) HasPoNumber() bool`

HasPoNumber returns a boolean if a field has been set.

### GetPoDate

`func (o *ComputerPurchase) GetPoDate() string`

GetPoDate returns the PoDate field if non-nil, zero value otherwise.

### GetPoDateOk

`func (o *ComputerPurchase) GetPoDateOk() (*string, bool)`

GetPoDateOk returns a tuple with the PoDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoDate

`func (o *ComputerPurchase) SetPoDate(v string)`

SetPoDate sets PoDate field to given value.

### HasPoDate

`func (o *ComputerPurchase) HasPoDate() bool`

HasPoDate returns a boolean if a field has been set.

### GetVendor

`func (o *ComputerPurchase) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ComputerPurchase) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ComputerPurchase) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ComputerPurchase) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetWarrantyDate

`func (o *ComputerPurchase) GetWarrantyDate() string`

GetWarrantyDate returns the WarrantyDate field if non-nil, zero value otherwise.

### GetWarrantyDateOk

`func (o *ComputerPurchase) GetWarrantyDateOk() (*string, bool)`

GetWarrantyDateOk returns a tuple with the WarrantyDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyDate

`func (o *ComputerPurchase) SetWarrantyDate(v string)`

SetWarrantyDate sets WarrantyDate field to given value.

### HasWarrantyDate

`func (o *ComputerPurchase) HasWarrantyDate() bool`

HasWarrantyDate returns a boolean if a field has been set.

### GetAppleCareId

`func (o *ComputerPurchase) GetAppleCareId() string`

GetAppleCareId returns the AppleCareId field if non-nil, zero value otherwise.

### GetAppleCareIdOk

`func (o *ComputerPurchase) GetAppleCareIdOk() (*string, bool)`

GetAppleCareIdOk returns a tuple with the AppleCareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleCareId

`func (o *ComputerPurchase) SetAppleCareId(v string)`

SetAppleCareId sets AppleCareId field to given value.

### HasAppleCareId

`func (o *ComputerPurchase) HasAppleCareId() bool`

HasAppleCareId returns a boolean if a field has been set.

### GetLeaseDate

`func (o *ComputerPurchase) GetLeaseDate() string`

GetLeaseDate returns the LeaseDate field if non-nil, zero value otherwise.

### GetLeaseDateOk

`func (o *ComputerPurchase) GetLeaseDateOk() (*string, bool)`

GetLeaseDateOk returns a tuple with the LeaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseDate

`func (o *ComputerPurchase) SetLeaseDate(v string)`

SetLeaseDate sets LeaseDate field to given value.

### HasLeaseDate

`func (o *ComputerPurchase) HasLeaseDate() bool`

HasLeaseDate returns a boolean if a field has been set.

### GetPurchasePrice

`func (o *ComputerPurchase) GetPurchasePrice() string`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *ComputerPurchase) GetPurchasePriceOk() (*string, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *ComputerPurchase) SetPurchasePrice(v string)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *ComputerPurchase) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.

### GetLifeExpectancy

`func (o *ComputerPurchase) GetLifeExpectancy() int64`

GetLifeExpectancy returns the LifeExpectancy field if non-nil, zero value otherwise.

### GetLifeExpectancyOk

`func (o *ComputerPurchase) GetLifeExpectancyOk() (*int64, bool)`

GetLifeExpectancyOk returns a tuple with the LifeExpectancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeExpectancy

`func (o *ComputerPurchase) SetLifeExpectancy(v int64)`

SetLifeExpectancy sets LifeExpectancy field to given value.

### HasLifeExpectancy

`func (o *ComputerPurchase) HasLifeExpectancy() bool`

HasLifeExpectancy returns a boolean if a field has been set.

### GetPurchasingAccount

`func (o *ComputerPurchase) GetPurchasingAccount() string`

GetPurchasingAccount returns the PurchasingAccount field if non-nil, zero value otherwise.

### GetPurchasingAccountOk

`func (o *ComputerPurchase) GetPurchasingAccountOk() (*string, bool)`

GetPurchasingAccountOk returns a tuple with the PurchasingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingAccount

`func (o *ComputerPurchase) SetPurchasingAccount(v string)`

SetPurchasingAccount sets PurchasingAccount field to given value.

### HasPurchasingAccount

`func (o *ComputerPurchase) HasPurchasingAccount() bool`

HasPurchasingAccount returns a boolean if a field has been set.

### GetPurchasingContact

`func (o *ComputerPurchase) GetPurchasingContact() string`

GetPurchasingContact returns the PurchasingContact field if non-nil, zero value otherwise.

### GetPurchasingContactOk

`func (o *ComputerPurchase) GetPurchasingContactOk() (*string, bool)`

GetPurchasingContactOk returns a tuple with the PurchasingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingContact

`func (o *ComputerPurchase) SetPurchasingContact(v string)`

SetPurchasingContact sets PurchasingContact field to given value.

### HasPurchasingContact

`func (o *ComputerPurchase) HasPurchasingContact() bool`

HasPurchasingContact returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *ComputerPurchase) GetExtensionAttributes() []ComputerExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *ComputerPurchase) GetExtensionAttributesOk() (*[]ComputerExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *ComputerPurchase) SetExtensionAttributes(v []ComputerExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *ComputerPurchase) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


