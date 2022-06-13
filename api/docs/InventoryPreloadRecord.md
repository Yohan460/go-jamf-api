# InventoryPreloadRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**SerialNumber** | **string** |  | 
**DeviceType** | **string** |  | 
**Username** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**EmailAddress** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **string** |  | [optional] 
**Department** | Pointer to **string** |  | [optional] 
**Building** | Pointer to **string** |  | [optional] 
**Room** | Pointer to **string** |  | [optional] 
**PoNumber** | Pointer to **string** |  | [optional] 
**PoDate** | Pointer to **string** |  | [optional] 
**WarrantyExpiration** | Pointer to **string** |  | [optional] 
**AppleCareId** | Pointer to **string** |  | [optional] 
**LifeExpectancy** | Pointer to **string** |  | [optional] 
**PurchasePrice** | Pointer to **string** |  | [optional] 
**PurchasingContact** | Pointer to **string** |  | [optional] 
**PurchasingAccount** | Pointer to **string** |  | [optional] 
**LeaseExpiration** | Pointer to **string** |  | [optional] 
**BarCode1** | Pointer to **string** |  | [optional] 
**BarCode2** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**ExtensionAttributes** | Pointer to [**[]InventoryPreloadExtensionAttribute**](InventoryPreloadExtensionAttribute.md) |  | [optional] 

## Methods

### NewInventoryPreloadRecord

`func NewInventoryPreloadRecord(serialNumber string, deviceType string, ) *InventoryPreloadRecord`

NewInventoryPreloadRecord instantiates a new InventoryPreloadRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryPreloadRecordWithDefaults

`func NewInventoryPreloadRecordWithDefaults() *InventoryPreloadRecord`

NewInventoryPreloadRecordWithDefaults instantiates a new InventoryPreloadRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InventoryPreloadRecord) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InventoryPreloadRecord) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InventoryPreloadRecord) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InventoryPreloadRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *InventoryPreloadRecord) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *InventoryPreloadRecord) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *InventoryPreloadRecord) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetDeviceType

`func (o *InventoryPreloadRecord) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *InventoryPreloadRecord) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *InventoryPreloadRecord) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.


### GetUsername

`func (o *InventoryPreloadRecord) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InventoryPreloadRecord) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InventoryPreloadRecord) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *InventoryPreloadRecord) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFullName

`func (o *InventoryPreloadRecord) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *InventoryPreloadRecord) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *InventoryPreloadRecord) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *InventoryPreloadRecord) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *InventoryPreloadRecord) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *InventoryPreloadRecord) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *InventoryPreloadRecord) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *InventoryPreloadRecord) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *InventoryPreloadRecord) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *InventoryPreloadRecord) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *InventoryPreloadRecord) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *InventoryPreloadRecord) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPosition

`func (o *InventoryPreloadRecord) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *InventoryPreloadRecord) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *InventoryPreloadRecord) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *InventoryPreloadRecord) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetDepartment

`func (o *InventoryPreloadRecord) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *InventoryPreloadRecord) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *InventoryPreloadRecord) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *InventoryPreloadRecord) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetBuilding

`func (o *InventoryPreloadRecord) GetBuilding() string`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *InventoryPreloadRecord) GetBuildingOk() (*string, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *InventoryPreloadRecord) SetBuilding(v string)`

SetBuilding sets Building field to given value.

### HasBuilding

`func (o *InventoryPreloadRecord) HasBuilding() bool`

HasBuilding returns a boolean if a field has been set.

### GetRoom

`func (o *InventoryPreloadRecord) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *InventoryPreloadRecord) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *InventoryPreloadRecord) SetRoom(v string)`

SetRoom sets Room field to given value.

### HasRoom

`func (o *InventoryPreloadRecord) HasRoom() bool`

HasRoom returns a boolean if a field has been set.

### GetPoNumber

`func (o *InventoryPreloadRecord) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *InventoryPreloadRecord) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *InventoryPreloadRecord) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.

### HasPoNumber

`func (o *InventoryPreloadRecord) HasPoNumber() bool`

HasPoNumber returns a boolean if a field has been set.

### GetPoDate

`func (o *InventoryPreloadRecord) GetPoDate() string`

GetPoDate returns the PoDate field if non-nil, zero value otherwise.

### GetPoDateOk

`func (o *InventoryPreloadRecord) GetPoDateOk() (*string, bool)`

GetPoDateOk returns a tuple with the PoDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoDate

`func (o *InventoryPreloadRecord) SetPoDate(v string)`

SetPoDate sets PoDate field to given value.

### HasPoDate

`func (o *InventoryPreloadRecord) HasPoDate() bool`

HasPoDate returns a boolean if a field has been set.

### GetWarrantyExpiration

`func (o *InventoryPreloadRecord) GetWarrantyExpiration() string`

GetWarrantyExpiration returns the WarrantyExpiration field if non-nil, zero value otherwise.

### GetWarrantyExpirationOk

`func (o *InventoryPreloadRecord) GetWarrantyExpirationOk() (*string, bool)`

GetWarrantyExpirationOk returns a tuple with the WarrantyExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyExpiration

`func (o *InventoryPreloadRecord) SetWarrantyExpiration(v string)`

SetWarrantyExpiration sets WarrantyExpiration field to given value.

### HasWarrantyExpiration

`func (o *InventoryPreloadRecord) HasWarrantyExpiration() bool`

HasWarrantyExpiration returns a boolean if a field has been set.

### GetAppleCareId

`func (o *InventoryPreloadRecord) GetAppleCareId() string`

GetAppleCareId returns the AppleCareId field if non-nil, zero value otherwise.

### GetAppleCareIdOk

`func (o *InventoryPreloadRecord) GetAppleCareIdOk() (*string, bool)`

GetAppleCareIdOk returns a tuple with the AppleCareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleCareId

`func (o *InventoryPreloadRecord) SetAppleCareId(v string)`

SetAppleCareId sets AppleCareId field to given value.

### HasAppleCareId

`func (o *InventoryPreloadRecord) HasAppleCareId() bool`

HasAppleCareId returns a boolean if a field has been set.

### GetLifeExpectancy

`func (o *InventoryPreloadRecord) GetLifeExpectancy() string`

GetLifeExpectancy returns the LifeExpectancy field if non-nil, zero value otherwise.

### GetLifeExpectancyOk

`func (o *InventoryPreloadRecord) GetLifeExpectancyOk() (*string, bool)`

GetLifeExpectancyOk returns a tuple with the LifeExpectancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeExpectancy

`func (o *InventoryPreloadRecord) SetLifeExpectancy(v string)`

SetLifeExpectancy sets LifeExpectancy field to given value.

### HasLifeExpectancy

`func (o *InventoryPreloadRecord) HasLifeExpectancy() bool`

HasLifeExpectancy returns a boolean if a field has been set.

### GetPurchasePrice

`func (o *InventoryPreloadRecord) GetPurchasePrice() string`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *InventoryPreloadRecord) GetPurchasePriceOk() (*string, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *InventoryPreloadRecord) SetPurchasePrice(v string)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *InventoryPreloadRecord) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.

### GetPurchasingContact

`func (o *InventoryPreloadRecord) GetPurchasingContact() string`

GetPurchasingContact returns the PurchasingContact field if non-nil, zero value otherwise.

### GetPurchasingContactOk

`func (o *InventoryPreloadRecord) GetPurchasingContactOk() (*string, bool)`

GetPurchasingContactOk returns a tuple with the PurchasingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingContact

`func (o *InventoryPreloadRecord) SetPurchasingContact(v string)`

SetPurchasingContact sets PurchasingContact field to given value.

### HasPurchasingContact

`func (o *InventoryPreloadRecord) HasPurchasingContact() bool`

HasPurchasingContact returns a boolean if a field has been set.

### GetPurchasingAccount

`func (o *InventoryPreloadRecord) GetPurchasingAccount() string`

GetPurchasingAccount returns the PurchasingAccount field if non-nil, zero value otherwise.

### GetPurchasingAccountOk

`func (o *InventoryPreloadRecord) GetPurchasingAccountOk() (*string, bool)`

GetPurchasingAccountOk returns a tuple with the PurchasingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingAccount

`func (o *InventoryPreloadRecord) SetPurchasingAccount(v string)`

SetPurchasingAccount sets PurchasingAccount field to given value.

### HasPurchasingAccount

`func (o *InventoryPreloadRecord) HasPurchasingAccount() bool`

HasPurchasingAccount returns a boolean if a field has been set.

### GetLeaseExpiration

`func (o *InventoryPreloadRecord) GetLeaseExpiration() string`

GetLeaseExpiration returns the LeaseExpiration field if non-nil, zero value otherwise.

### GetLeaseExpirationOk

`func (o *InventoryPreloadRecord) GetLeaseExpirationOk() (*string, bool)`

GetLeaseExpirationOk returns a tuple with the LeaseExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseExpiration

`func (o *InventoryPreloadRecord) SetLeaseExpiration(v string)`

SetLeaseExpiration sets LeaseExpiration field to given value.

### HasLeaseExpiration

`func (o *InventoryPreloadRecord) HasLeaseExpiration() bool`

HasLeaseExpiration returns a boolean if a field has been set.

### GetBarCode1

`func (o *InventoryPreloadRecord) GetBarCode1() string`

GetBarCode1 returns the BarCode1 field if non-nil, zero value otherwise.

### GetBarCode1Ok

`func (o *InventoryPreloadRecord) GetBarCode1Ok() (*string, bool)`

GetBarCode1Ok returns a tuple with the BarCode1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarCode1

`func (o *InventoryPreloadRecord) SetBarCode1(v string)`

SetBarCode1 sets BarCode1 field to given value.

### HasBarCode1

`func (o *InventoryPreloadRecord) HasBarCode1() bool`

HasBarCode1 returns a boolean if a field has been set.

### GetBarCode2

`func (o *InventoryPreloadRecord) GetBarCode2() string`

GetBarCode2 returns the BarCode2 field if non-nil, zero value otherwise.

### GetBarCode2Ok

`func (o *InventoryPreloadRecord) GetBarCode2Ok() (*string, bool)`

GetBarCode2Ok returns a tuple with the BarCode2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarCode2

`func (o *InventoryPreloadRecord) SetBarCode2(v string)`

SetBarCode2 sets BarCode2 field to given value.

### HasBarCode2

`func (o *InventoryPreloadRecord) HasBarCode2() bool`

HasBarCode2 returns a boolean if a field has been set.

### GetAssetTag

`func (o *InventoryPreloadRecord) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *InventoryPreloadRecord) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *InventoryPreloadRecord) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *InventoryPreloadRecord) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetVendor

`func (o *InventoryPreloadRecord) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *InventoryPreloadRecord) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *InventoryPreloadRecord) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *InventoryPreloadRecord) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *InventoryPreloadRecord) GetExtensionAttributes() []InventoryPreloadExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *InventoryPreloadRecord) GetExtensionAttributesOk() (*[]InventoryPreloadExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *InventoryPreloadRecord) SetExtensionAttributes(v []InventoryPreloadExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *InventoryPreloadRecord) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


