# MobileDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**WifiMacAddress** | Pointer to **string** |  | [optional] 
**Udid** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**ModelIdentifier** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**SoftwareUpdateDeviceId** | Pointer to **string** |  | [optional] 

## Methods

### NewMobileDevice

`func NewMobileDevice() *MobileDevice`

NewMobileDevice instantiates a new MobileDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceWithDefaults

`func NewMobileDeviceWithDefaults() *MobileDevice`

NewMobileDeviceWithDefaults instantiates a new MobileDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MobileDevice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MobileDevice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MobileDevice) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MobileDevice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MobileDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MobileDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MobileDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MobileDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerialNumber

`func (o *MobileDevice) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *MobileDevice) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *MobileDevice) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *MobileDevice) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetWifiMacAddress

`func (o *MobileDevice) GetWifiMacAddress() string`

GetWifiMacAddress returns the WifiMacAddress field if non-nil, zero value otherwise.

### GetWifiMacAddressOk

`func (o *MobileDevice) GetWifiMacAddressOk() (*string, bool)`

GetWifiMacAddressOk returns a tuple with the WifiMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMacAddress

`func (o *MobileDevice) SetWifiMacAddress(v string)`

SetWifiMacAddress sets WifiMacAddress field to given value.

### HasWifiMacAddress

`func (o *MobileDevice) HasWifiMacAddress() bool`

HasWifiMacAddress returns a boolean if a field has been set.

### GetUdid

`func (o *MobileDevice) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *MobileDevice) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *MobileDevice) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *MobileDevice) HasUdid() bool`

HasUdid returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *MobileDevice) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *MobileDevice) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *MobileDevice) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *MobileDevice) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetModel

`func (o *MobileDevice) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MobileDevice) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MobileDevice) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *MobileDevice) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelIdentifier

`func (o *MobileDevice) GetModelIdentifier() string`

GetModelIdentifier returns the ModelIdentifier field if non-nil, zero value otherwise.

### GetModelIdentifierOk

`func (o *MobileDevice) GetModelIdentifierOk() (*string, bool)`

GetModelIdentifierOk returns a tuple with the ModelIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelIdentifier

`func (o *MobileDevice) SetModelIdentifier(v string)`

SetModelIdentifier sets ModelIdentifier field to given value.

### HasModelIdentifier

`func (o *MobileDevice) HasModelIdentifier() bool`

HasModelIdentifier returns a boolean if a field has been set.

### GetUsername

`func (o *MobileDevice) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MobileDevice) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MobileDevice) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *MobileDevice) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetType

`func (o *MobileDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MobileDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MobileDevice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MobileDevice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSoftwareUpdateDeviceId

`func (o *MobileDevice) GetSoftwareUpdateDeviceId() string`

GetSoftwareUpdateDeviceId returns the SoftwareUpdateDeviceId field if non-nil, zero value otherwise.

### GetSoftwareUpdateDeviceIdOk

`func (o *MobileDevice) GetSoftwareUpdateDeviceIdOk() (*string, bool)`

GetSoftwareUpdateDeviceIdOk returns a tuple with the SoftwareUpdateDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdateDeviceId

`func (o *MobileDevice) SetSoftwareUpdateDeviceId(v string)`

SetSoftwareUpdateDeviceId sets SoftwareUpdateDeviceId field to given value.

### HasSoftwareUpdateDeviceId

`func (o *MobileDevice) HasSoftwareUpdateDeviceId() bool`

HasSoftwareUpdateDeviceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


