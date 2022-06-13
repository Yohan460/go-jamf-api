# InventoryPreloadRecordV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**SerialNumber** | **string** |  | 
**DeviceType** | **string** |  | 
**Username** | Pointer to **NullableString** |  | [optional] 
**FullName** | Pointer to **NullableString** |  | [optional] 
**EmailAddress** | Pointer to **NullableString** |  | [optional] 
**PhoneNumber** | Pointer to **NullableString** |  | [optional] 
**Position** | Pointer to **NullableString** |  | [optional] 
**Department** | Pointer to **NullableString** |  | [optional] 
**Building** | Pointer to **NullableString** |  | [optional] 
**Room** | Pointer to **NullableString** |  | [optional] 
**PoNumber** | Pointer to **NullableString** |  | [optional] 
**PoDate** | Pointer to **NullableString** |  | [optional] 
**WarrantyExpiration** | Pointer to **NullableString** |  | [optional] 
**AppleCareId** | Pointer to **NullableString** |  | [optional] 
**LifeExpectancy** | Pointer to **NullableString** |  | [optional] 
**PurchasePrice** | Pointer to **NullableString** |  | [optional] 
**PurchasingContact** | Pointer to **NullableString** |  | [optional] 
**PurchasingAccount** | Pointer to **NullableString** |  | [optional] 
**LeaseExpiration** | Pointer to **NullableString** |  | [optional] 
**BarCode1** | Pointer to **NullableString** |  | [optional] 
**BarCode2** | Pointer to **NullableString** |  | [optional] 
**AssetTag** | Pointer to **NullableString** |  | [optional] 
**Vendor** | Pointer to **NullableString** |  | [optional] 
**ExtensionAttributes** | Pointer to [**[]InventoryPreloadExtensionAttribute**](InventoryPreloadExtensionAttribute.md) |  | [optional] 

## Methods

### NewInventoryPreloadRecordV2

`func NewInventoryPreloadRecordV2(serialNumber string, deviceType string, ) *InventoryPreloadRecordV2`

NewInventoryPreloadRecordV2 instantiates a new InventoryPreloadRecordV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryPreloadRecordV2WithDefaults

`func NewInventoryPreloadRecordV2WithDefaults() *InventoryPreloadRecordV2`

NewInventoryPreloadRecordV2WithDefaults instantiates a new InventoryPreloadRecordV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InventoryPreloadRecordV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InventoryPreloadRecordV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InventoryPreloadRecordV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InventoryPreloadRecordV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *InventoryPreloadRecordV2) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *InventoryPreloadRecordV2) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *InventoryPreloadRecordV2) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetDeviceType

`func (o *InventoryPreloadRecordV2) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *InventoryPreloadRecordV2) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *InventoryPreloadRecordV2) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.


### GetUsername

`func (o *InventoryPreloadRecordV2) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InventoryPreloadRecordV2) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InventoryPreloadRecordV2) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *InventoryPreloadRecordV2) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *InventoryPreloadRecordV2) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *InventoryPreloadRecordV2) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetFullName

`func (o *InventoryPreloadRecordV2) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *InventoryPreloadRecordV2) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *InventoryPreloadRecordV2) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *InventoryPreloadRecordV2) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *InventoryPreloadRecordV2) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *InventoryPreloadRecordV2) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetEmailAddress

`func (o *InventoryPreloadRecordV2) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *InventoryPreloadRecordV2) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *InventoryPreloadRecordV2) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *InventoryPreloadRecordV2) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *InventoryPreloadRecordV2) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *InventoryPreloadRecordV2) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetPhoneNumber

`func (o *InventoryPreloadRecordV2) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *InventoryPreloadRecordV2) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *InventoryPreloadRecordV2) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *InventoryPreloadRecordV2) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *InventoryPreloadRecordV2) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *InventoryPreloadRecordV2) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetPosition

`func (o *InventoryPreloadRecordV2) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *InventoryPreloadRecordV2) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *InventoryPreloadRecordV2) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *InventoryPreloadRecordV2) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *InventoryPreloadRecordV2) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *InventoryPreloadRecordV2) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetDepartment

`func (o *InventoryPreloadRecordV2) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *InventoryPreloadRecordV2) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *InventoryPreloadRecordV2) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *InventoryPreloadRecordV2) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *InventoryPreloadRecordV2) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *InventoryPreloadRecordV2) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetBuilding

`func (o *InventoryPreloadRecordV2) GetBuilding() string`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *InventoryPreloadRecordV2) GetBuildingOk() (*string, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *InventoryPreloadRecordV2) SetBuilding(v string)`

SetBuilding sets Building field to given value.

### HasBuilding

`func (o *InventoryPreloadRecordV2) HasBuilding() bool`

HasBuilding returns a boolean if a field has been set.

### SetBuildingNil

`func (o *InventoryPreloadRecordV2) SetBuildingNil(b bool)`

 SetBuildingNil sets the value for Building to be an explicit nil

### UnsetBuilding
`func (o *InventoryPreloadRecordV2) UnsetBuilding()`

UnsetBuilding ensures that no value is present for Building, not even an explicit nil
### GetRoom

`func (o *InventoryPreloadRecordV2) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *InventoryPreloadRecordV2) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *InventoryPreloadRecordV2) SetRoom(v string)`

SetRoom sets Room field to given value.

### HasRoom

`func (o *InventoryPreloadRecordV2) HasRoom() bool`

HasRoom returns a boolean if a field has been set.

### SetRoomNil

`func (o *InventoryPreloadRecordV2) SetRoomNil(b bool)`

 SetRoomNil sets the value for Room to be an explicit nil

### UnsetRoom
`func (o *InventoryPreloadRecordV2) UnsetRoom()`

UnsetRoom ensures that no value is present for Room, not even an explicit nil
### GetPoNumber

`func (o *InventoryPreloadRecordV2) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *InventoryPreloadRecordV2) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *InventoryPreloadRecordV2) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.

### HasPoNumber

`func (o *InventoryPreloadRecordV2) HasPoNumber() bool`

HasPoNumber returns a boolean if a field has been set.

### SetPoNumberNil

`func (o *InventoryPreloadRecordV2) SetPoNumberNil(b bool)`

 SetPoNumberNil sets the value for PoNumber to be an explicit nil

### UnsetPoNumber
`func (o *InventoryPreloadRecordV2) UnsetPoNumber()`

UnsetPoNumber ensures that no value is present for PoNumber, not even an explicit nil
### GetPoDate

`func (o *InventoryPreloadRecordV2) GetPoDate() string`

GetPoDate returns the PoDate field if non-nil, zero value otherwise.

### GetPoDateOk

`func (o *InventoryPreloadRecordV2) GetPoDateOk() (*string, bool)`

GetPoDateOk returns a tuple with the PoDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoDate

`func (o *InventoryPreloadRecordV2) SetPoDate(v string)`

SetPoDate sets PoDate field to given value.

### HasPoDate

`func (o *InventoryPreloadRecordV2) HasPoDate() bool`

HasPoDate returns a boolean if a field has been set.

### SetPoDateNil

`func (o *InventoryPreloadRecordV2) SetPoDateNil(b bool)`

 SetPoDateNil sets the value for PoDate to be an explicit nil

### UnsetPoDate
`func (o *InventoryPreloadRecordV2) UnsetPoDate()`

UnsetPoDate ensures that no value is present for PoDate, not even an explicit nil
### GetWarrantyExpiration

`func (o *InventoryPreloadRecordV2) GetWarrantyExpiration() string`

GetWarrantyExpiration returns the WarrantyExpiration field if non-nil, zero value otherwise.

### GetWarrantyExpirationOk

`func (o *InventoryPreloadRecordV2) GetWarrantyExpirationOk() (*string, bool)`

GetWarrantyExpirationOk returns a tuple with the WarrantyExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyExpiration

`func (o *InventoryPreloadRecordV2) SetWarrantyExpiration(v string)`

SetWarrantyExpiration sets WarrantyExpiration field to given value.

### HasWarrantyExpiration

`func (o *InventoryPreloadRecordV2) HasWarrantyExpiration() bool`

HasWarrantyExpiration returns a boolean if a field has been set.

### SetWarrantyExpirationNil

`func (o *InventoryPreloadRecordV2) SetWarrantyExpirationNil(b bool)`

 SetWarrantyExpirationNil sets the value for WarrantyExpiration to be an explicit nil

### UnsetWarrantyExpiration
`func (o *InventoryPreloadRecordV2) UnsetWarrantyExpiration()`

UnsetWarrantyExpiration ensures that no value is present for WarrantyExpiration, not even an explicit nil
### GetAppleCareId

`func (o *InventoryPreloadRecordV2) GetAppleCareId() string`

GetAppleCareId returns the AppleCareId field if non-nil, zero value otherwise.

### GetAppleCareIdOk

`func (o *InventoryPreloadRecordV2) GetAppleCareIdOk() (*string, bool)`

GetAppleCareIdOk returns a tuple with the AppleCareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleCareId

`func (o *InventoryPreloadRecordV2) SetAppleCareId(v string)`

SetAppleCareId sets AppleCareId field to given value.

### HasAppleCareId

`func (o *InventoryPreloadRecordV2) HasAppleCareId() bool`

HasAppleCareId returns a boolean if a field has been set.

### SetAppleCareIdNil

`func (o *InventoryPreloadRecordV2) SetAppleCareIdNil(b bool)`

 SetAppleCareIdNil sets the value for AppleCareId to be an explicit nil

### UnsetAppleCareId
`func (o *InventoryPreloadRecordV2) UnsetAppleCareId()`

UnsetAppleCareId ensures that no value is present for AppleCareId, not even an explicit nil
### GetLifeExpectancy

`func (o *InventoryPreloadRecordV2) GetLifeExpectancy() string`

GetLifeExpectancy returns the LifeExpectancy field if non-nil, zero value otherwise.

### GetLifeExpectancyOk

`func (o *InventoryPreloadRecordV2) GetLifeExpectancyOk() (*string, bool)`

GetLifeExpectancyOk returns a tuple with the LifeExpectancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeExpectancy

`func (o *InventoryPreloadRecordV2) SetLifeExpectancy(v string)`

SetLifeExpectancy sets LifeExpectancy field to given value.

### HasLifeExpectancy

`func (o *InventoryPreloadRecordV2) HasLifeExpectancy() bool`

HasLifeExpectancy returns a boolean if a field has been set.

### SetLifeExpectancyNil

`func (o *InventoryPreloadRecordV2) SetLifeExpectancyNil(b bool)`

 SetLifeExpectancyNil sets the value for LifeExpectancy to be an explicit nil

### UnsetLifeExpectancy
`func (o *InventoryPreloadRecordV2) UnsetLifeExpectancy()`

UnsetLifeExpectancy ensures that no value is present for LifeExpectancy, not even an explicit nil
### GetPurchasePrice

`func (o *InventoryPreloadRecordV2) GetPurchasePrice() string`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *InventoryPreloadRecordV2) GetPurchasePriceOk() (*string, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *InventoryPreloadRecordV2) SetPurchasePrice(v string)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *InventoryPreloadRecordV2) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.

### SetPurchasePriceNil

`func (o *InventoryPreloadRecordV2) SetPurchasePriceNil(b bool)`

 SetPurchasePriceNil sets the value for PurchasePrice to be an explicit nil

### UnsetPurchasePrice
`func (o *InventoryPreloadRecordV2) UnsetPurchasePrice()`

UnsetPurchasePrice ensures that no value is present for PurchasePrice, not even an explicit nil
### GetPurchasingContact

`func (o *InventoryPreloadRecordV2) GetPurchasingContact() string`

GetPurchasingContact returns the PurchasingContact field if non-nil, zero value otherwise.

### GetPurchasingContactOk

`func (o *InventoryPreloadRecordV2) GetPurchasingContactOk() (*string, bool)`

GetPurchasingContactOk returns a tuple with the PurchasingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingContact

`func (o *InventoryPreloadRecordV2) SetPurchasingContact(v string)`

SetPurchasingContact sets PurchasingContact field to given value.

### HasPurchasingContact

`func (o *InventoryPreloadRecordV2) HasPurchasingContact() bool`

HasPurchasingContact returns a boolean if a field has been set.

### SetPurchasingContactNil

`func (o *InventoryPreloadRecordV2) SetPurchasingContactNil(b bool)`

 SetPurchasingContactNil sets the value for PurchasingContact to be an explicit nil

### UnsetPurchasingContact
`func (o *InventoryPreloadRecordV2) UnsetPurchasingContact()`

UnsetPurchasingContact ensures that no value is present for PurchasingContact, not even an explicit nil
### GetPurchasingAccount

`func (o *InventoryPreloadRecordV2) GetPurchasingAccount() string`

GetPurchasingAccount returns the PurchasingAccount field if non-nil, zero value otherwise.

### GetPurchasingAccountOk

`func (o *InventoryPreloadRecordV2) GetPurchasingAccountOk() (*string, bool)`

GetPurchasingAccountOk returns a tuple with the PurchasingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingAccount

`func (o *InventoryPreloadRecordV2) SetPurchasingAccount(v string)`

SetPurchasingAccount sets PurchasingAccount field to given value.

### HasPurchasingAccount

`func (o *InventoryPreloadRecordV2) HasPurchasingAccount() bool`

HasPurchasingAccount returns a boolean if a field has been set.

### SetPurchasingAccountNil

`func (o *InventoryPreloadRecordV2) SetPurchasingAccountNil(b bool)`

 SetPurchasingAccountNil sets the value for PurchasingAccount to be an explicit nil

### UnsetPurchasingAccount
`func (o *InventoryPreloadRecordV2) UnsetPurchasingAccount()`

UnsetPurchasingAccount ensures that no value is present for PurchasingAccount, not even an explicit nil
### GetLeaseExpiration

`func (o *InventoryPreloadRecordV2) GetLeaseExpiration() string`

GetLeaseExpiration returns the LeaseExpiration field if non-nil, zero value otherwise.

### GetLeaseExpirationOk

`func (o *InventoryPreloadRecordV2) GetLeaseExpirationOk() (*string, bool)`

GetLeaseExpirationOk returns a tuple with the LeaseExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseExpiration

`func (o *InventoryPreloadRecordV2) SetLeaseExpiration(v string)`

SetLeaseExpiration sets LeaseExpiration field to given value.

### HasLeaseExpiration

`func (o *InventoryPreloadRecordV2) HasLeaseExpiration() bool`

HasLeaseExpiration returns a boolean if a field has been set.

### SetLeaseExpirationNil

`func (o *InventoryPreloadRecordV2) SetLeaseExpirationNil(b bool)`

 SetLeaseExpirationNil sets the value for LeaseExpiration to be an explicit nil

### UnsetLeaseExpiration
`func (o *InventoryPreloadRecordV2) UnsetLeaseExpiration()`

UnsetLeaseExpiration ensures that no value is present for LeaseExpiration, not even an explicit nil
### GetBarCode1

`func (o *InventoryPreloadRecordV2) GetBarCode1() string`

GetBarCode1 returns the BarCode1 field if non-nil, zero value otherwise.

### GetBarCode1Ok

`func (o *InventoryPreloadRecordV2) GetBarCode1Ok() (*string, bool)`

GetBarCode1Ok returns a tuple with the BarCode1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarCode1

`func (o *InventoryPreloadRecordV2) SetBarCode1(v string)`

SetBarCode1 sets BarCode1 field to given value.

### HasBarCode1

`func (o *InventoryPreloadRecordV2) HasBarCode1() bool`

HasBarCode1 returns a boolean if a field has been set.

### SetBarCode1Nil

`func (o *InventoryPreloadRecordV2) SetBarCode1Nil(b bool)`

 SetBarCode1Nil sets the value for BarCode1 to be an explicit nil

### UnsetBarCode1
`func (o *InventoryPreloadRecordV2) UnsetBarCode1()`

UnsetBarCode1 ensures that no value is present for BarCode1, not even an explicit nil
### GetBarCode2

`func (o *InventoryPreloadRecordV2) GetBarCode2() string`

GetBarCode2 returns the BarCode2 field if non-nil, zero value otherwise.

### GetBarCode2Ok

`func (o *InventoryPreloadRecordV2) GetBarCode2Ok() (*string, bool)`

GetBarCode2Ok returns a tuple with the BarCode2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarCode2

`func (o *InventoryPreloadRecordV2) SetBarCode2(v string)`

SetBarCode2 sets BarCode2 field to given value.

### HasBarCode2

`func (o *InventoryPreloadRecordV2) HasBarCode2() bool`

HasBarCode2 returns a boolean if a field has been set.

### SetBarCode2Nil

`func (o *InventoryPreloadRecordV2) SetBarCode2Nil(b bool)`

 SetBarCode2Nil sets the value for BarCode2 to be an explicit nil

### UnsetBarCode2
`func (o *InventoryPreloadRecordV2) UnsetBarCode2()`

UnsetBarCode2 ensures that no value is present for BarCode2, not even an explicit nil
### GetAssetTag

`func (o *InventoryPreloadRecordV2) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *InventoryPreloadRecordV2) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *InventoryPreloadRecordV2) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *InventoryPreloadRecordV2) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *InventoryPreloadRecordV2) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *InventoryPreloadRecordV2) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetVendor

`func (o *InventoryPreloadRecordV2) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *InventoryPreloadRecordV2) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *InventoryPreloadRecordV2) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *InventoryPreloadRecordV2) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### SetVendorNil

`func (o *InventoryPreloadRecordV2) SetVendorNil(b bool)`

 SetVendorNil sets the value for Vendor to be an explicit nil

### UnsetVendor
`func (o *InventoryPreloadRecordV2) UnsetVendor()`

UnsetVendor ensures that no value is present for Vendor, not even an explicit nil
### GetExtensionAttributes

`func (o *InventoryPreloadRecordV2) GetExtensionAttributes() []InventoryPreloadExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *InventoryPreloadRecordV2) GetExtensionAttributesOk() (*[]InventoryPreloadExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *InventoryPreloadRecordV2) SetExtensionAttributes(v []InventoryPreloadExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *InventoryPreloadRecordV2) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


