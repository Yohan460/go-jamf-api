# PrestagePurchasingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**IsLeased** | **bool** |  | 
**IsPurchased** | **bool** |  | 
**AppleCareID** | **string** |  | 
**PoNumber** | **string** |  | 
**Vendor** | **string** |  | 
**PurchasePrice** | **string** |  | 
**LifeExpectancy** | **int64** |  | 
**PurchasingAccount** | **string** |  | 
**PurchasingContact** | **string** |  | 
**LeaseDate** | **string** |  | 
**PoDate** | **string** |  | 
**WarrantyDate** | **string** |  | 
**VersionLock** | **int64** |  | 

## Methods

### NewPrestagePurchasingInformation

`func NewPrestagePurchasingInformation(id int64, isLeased bool, isPurchased bool, appleCareID string, poNumber string, vendor string, purchasePrice string, lifeExpectancy int64, purchasingAccount string, purchasingContact string, leaseDate string, poDate string, warrantyDate string, versionLock int64, ) *PrestagePurchasingInformation`

NewPrestagePurchasingInformation instantiates a new PrestagePurchasingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrestagePurchasingInformationWithDefaults

`func NewPrestagePurchasingInformationWithDefaults() *PrestagePurchasingInformation`

NewPrestagePurchasingInformationWithDefaults instantiates a new PrestagePurchasingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrestagePurchasingInformation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrestagePurchasingInformation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrestagePurchasingInformation) SetId(v int64)`

SetId sets Id field to given value.


### GetIsLeased

`func (o *PrestagePurchasingInformation) GetIsLeased() bool`

GetIsLeased returns the IsLeased field if non-nil, zero value otherwise.

### GetIsLeasedOk

`func (o *PrestagePurchasingInformation) GetIsLeasedOk() (*bool, bool)`

GetIsLeasedOk returns a tuple with the IsLeased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLeased

`func (o *PrestagePurchasingInformation) SetIsLeased(v bool)`

SetIsLeased sets IsLeased field to given value.


### GetIsPurchased

`func (o *PrestagePurchasingInformation) GetIsPurchased() bool`

GetIsPurchased returns the IsPurchased field if non-nil, zero value otherwise.

### GetIsPurchasedOk

`func (o *PrestagePurchasingInformation) GetIsPurchasedOk() (*bool, bool)`

GetIsPurchasedOk returns a tuple with the IsPurchased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPurchased

`func (o *PrestagePurchasingInformation) SetIsPurchased(v bool)`

SetIsPurchased sets IsPurchased field to given value.


### GetAppleCareID

`func (o *PrestagePurchasingInformation) GetAppleCareID() string`

GetAppleCareID returns the AppleCareID field if non-nil, zero value otherwise.

### GetAppleCareIDOk

`func (o *PrestagePurchasingInformation) GetAppleCareIDOk() (*string, bool)`

GetAppleCareIDOk returns a tuple with the AppleCareID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleCareID

`func (o *PrestagePurchasingInformation) SetAppleCareID(v string)`

SetAppleCareID sets AppleCareID field to given value.


### GetPoNumber

`func (o *PrestagePurchasingInformation) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *PrestagePurchasingInformation) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *PrestagePurchasingInformation) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.


### GetVendor

`func (o *PrestagePurchasingInformation) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *PrestagePurchasingInformation) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *PrestagePurchasingInformation) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetPurchasePrice

`func (o *PrestagePurchasingInformation) GetPurchasePrice() string`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *PrestagePurchasingInformation) GetPurchasePriceOk() (*string, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *PrestagePurchasingInformation) SetPurchasePrice(v string)`

SetPurchasePrice sets PurchasePrice field to given value.


### GetLifeExpectancy

`func (o *PrestagePurchasingInformation) GetLifeExpectancy() int64`

GetLifeExpectancy returns the LifeExpectancy field if non-nil, zero value otherwise.

### GetLifeExpectancyOk

`func (o *PrestagePurchasingInformation) GetLifeExpectancyOk() (*int64, bool)`

GetLifeExpectancyOk returns a tuple with the LifeExpectancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeExpectancy

`func (o *PrestagePurchasingInformation) SetLifeExpectancy(v int64)`

SetLifeExpectancy sets LifeExpectancy field to given value.


### GetPurchasingAccount

`func (o *PrestagePurchasingInformation) GetPurchasingAccount() string`

GetPurchasingAccount returns the PurchasingAccount field if non-nil, zero value otherwise.

### GetPurchasingAccountOk

`func (o *PrestagePurchasingInformation) GetPurchasingAccountOk() (*string, bool)`

GetPurchasingAccountOk returns a tuple with the PurchasingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingAccount

`func (o *PrestagePurchasingInformation) SetPurchasingAccount(v string)`

SetPurchasingAccount sets PurchasingAccount field to given value.


### GetPurchasingContact

`func (o *PrestagePurchasingInformation) GetPurchasingContact() string`

GetPurchasingContact returns the PurchasingContact field if non-nil, zero value otherwise.

### GetPurchasingContactOk

`func (o *PrestagePurchasingInformation) GetPurchasingContactOk() (*string, bool)`

GetPurchasingContactOk returns a tuple with the PurchasingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingContact

`func (o *PrestagePurchasingInformation) SetPurchasingContact(v string)`

SetPurchasingContact sets PurchasingContact field to given value.


### GetLeaseDate

`func (o *PrestagePurchasingInformation) GetLeaseDate() string`

GetLeaseDate returns the LeaseDate field if non-nil, zero value otherwise.

### GetLeaseDateOk

`func (o *PrestagePurchasingInformation) GetLeaseDateOk() (*string, bool)`

GetLeaseDateOk returns a tuple with the LeaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseDate

`func (o *PrestagePurchasingInformation) SetLeaseDate(v string)`

SetLeaseDate sets LeaseDate field to given value.


### GetPoDate

`func (o *PrestagePurchasingInformation) GetPoDate() string`

GetPoDate returns the PoDate field if non-nil, zero value otherwise.

### GetPoDateOk

`func (o *PrestagePurchasingInformation) GetPoDateOk() (*string, bool)`

GetPoDateOk returns a tuple with the PoDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoDate

`func (o *PrestagePurchasingInformation) SetPoDate(v string)`

SetPoDate sets PoDate field to given value.


### GetWarrantyDate

`func (o *PrestagePurchasingInformation) GetWarrantyDate() string`

GetWarrantyDate returns the WarrantyDate field if non-nil, zero value otherwise.

### GetWarrantyDateOk

`func (o *PrestagePurchasingInformation) GetWarrantyDateOk() (*string, bool)`

GetWarrantyDateOk returns a tuple with the WarrantyDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyDate

`func (o *PrestagePurchasingInformation) SetWarrantyDate(v string)`

SetWarrantyDate sets WarrantyDate field to given value.


### GetVersionLock

`func (o *PrestagePurchasingInformation) GetVersionLock() int64`

GetVersionLock returns the VersionLock field if non-nil, zero value otherwise.

### GetVersionLockOk

`func (o *PrestagePurchasingInformation) GetVersionLockOk() (*int64, bool)`

GetVersionLockOk returns a tuple with the VersionLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionLock

`func (o *PrestagePurchasingInformation) SetVersionLock(v int64)`

SetVersionLock sets VersionLock field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


