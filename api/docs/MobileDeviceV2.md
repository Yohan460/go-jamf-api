# MobileDeviceV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**WifiMacAddress** | Pointer to **string** |  | [optional] 
**Udid** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**ModelIdentifier** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ManagementId** | Pointer to **string** |  | [optional] [readonly] 
**SoftwareUpdateDeviceId** | Pointer to **string** |  | [optional] 

## Methods

### NewMobileDeviceV2

`func NewMobileDeviceV2() *MobileDeviceV2`

NewMobileDeviceV2 instantiates a new MobileDeviceV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceV2WithDefaults

`func NewMobileDeviceV2WithDefaults() *MobileDeviceV2`

NewMobileDeviceV2WithDefaults instantiates a new MobileDeviceV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MobileDeviceV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MobileDeviceV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MobileDeviceV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MobileDeviceV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MobileDeviceV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MobileDeviceV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MobileDeviceV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MobileDeviceV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerialNumber

`func (o *MobileDeviceV2) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *MobileDeviceV2) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *MobileDeviceV2) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *MobileDeviceV2) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetWifiMacAddress

`func (o *MobileDeviceV2) GetWifiMacAddress() string`

GetWifiMacAddress returns the WifiMacAddress field if non-nil, zero value otherwise.

### GetWifiMacAddressOk

`func (o *MobileDeviceV2) GetWifiMacAddressOk() (*string, bool)`

GetWifiMacAddressOk returns a tuple with the WifiMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMacAddress

`func (o *MobileDeviceV2) SetWifiMacAddress(v string)`

SetWifiMacAddress sets WifiMacAddress field to given value.

### HasWifiMacAddress

`func (o *MobileDeviceV2) HasWifiMacAddress() bool`

HasWifiMacAddress returns a boolean if a field has been set.

### GetUdid

`func (o *MobileDeviceV2) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *MobileDeviceV2) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *MobileDeviceV2) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *MobileDeviceV2) HasUdid() bool`

HasUdid returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *MobileDeviceV2) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *MobileDeviceV2) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *MobileDeviceV2) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *MobileDeviceV2) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetModel

`func (o *MobileDeviceV2) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MobileDeviceV2) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MobileDeviceV2) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *MobileDeviceV2) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelIdentifier

`func (o *MobileDeviceV2) GetModelIdentifier() string`

GetModelIdentifier returns the ModelIdentifier field if non-nil, zero value otherwise.

### GetModelIdentifierOk

`func (o *MobileDeviceV2) GetModelIdentifierOk() (*string, bool)`

GetModelIdentifierOk returns a tuple with the ModelIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelIdentifier

`func (o *MobileDeviceV2) SetModelIdentifier(v string)`

SetModelIdentifier sets ModelIdentifier field to given value.

### HasModelIdentifier

`func (o *MobileDeviceV2) HasModelIdentifier() bool`

HasModelIdentifier returns a boolean if a field has been set.

### GetUsername

`func (o *MobileDeviceV2) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MobileDeviceV2) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MobileDeviceV2) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *MobileDeviceV2) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetType

`func (o *MobileDeviceV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MobileDeviceV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MobileDeviceV2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MobileDeviceV2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetManagementId

`func (o *MobileDeviceV2) GetManagementId() string`

GetManagementId returns the ManagementId field if non-nil, zero value otherwise.

### GetManagementIdOk

`func (o *MobileDeviceV2) GetManagementIdOk() (*string, bool)`

GetManagementIdOk returns a tuple with the ManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementId

`func (o *MobileDeviceV2) SetManagementId(v string)`

SetManagementId sets ManagementId field to given value.

### HasManagementId

`func (o *MobileDeviceV2) HasManagementId() bool`

HasManagementId returns a boolean if a field has been set.

### GetSoftwareUpdateDeviceId

`func (o *MobileDeviceV2) GetSoftwareUpdateDeviceId() string`

GetSoftwareUpdateDeviceId returns the SoftwareUpdateDeviceId field if non-nil, zero value otherwise.

### GetSoftwareUpdateDeviceIdOk

`func (o *MobileDeviceV2) GetSoftwareUpdateDeviceIdOk() (*string, bool)`

GetSoftwareUpdateDeviceIdOk returns a tuple with the SoftwareUpdateDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdateDeviceId

`func (o *MobileDeviceV2) SetSoftwareUpdateDeviceId(v string)`

SetSoftwareUpdateDeviceId sets SoftwareUpdateDeviceId field to given value.

### HasSoftwareUpdateDeviceId

`func (o *MobileDeviceV2) HasSoftwareUpdateDeviceId() bool`

HasSoftwareUpdateDeviceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


