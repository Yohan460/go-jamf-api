# MobileDeviceHardware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityMb** | Pointer to **int32** |  | [optional] 
**AvailableSpaceMb** | Pointer to **int32** |  | [optional] 
**UsedSpacePercentage** | Pointer to **int32** |  | [optional] 
**BatteryLevel** | Pointer to **int32** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**WifiMacAddress** | Pointer to **string** |  | [optional] 
**BluetoothMacAddress** | Pointer to **string** |  | [optional] 
**ModemFirmwareVersion** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**ModelIdentifier** | Pointer to **string** |  | [optional] 
**ModelNumber** | Pointer to **string** |  | [optional] 
**BluetoothLowEnergyCapable** | Pointer to **bool** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**ExtensionAttributes** | Pointer to [**[]MobileDeviceExtensionAttribute**](MobileDeviceExtensionAttribute.md) |  | [optional] 

## Methods

### NewMobileDeviceHardware

`func NewMobileDeviceHardware() *MobileDeviceHardware`

NewMobileDeviceHardware instantiates a new MobileDeviceHardware object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceHardwareWithDefaults

`func NewMobileDeviceHardwareWithDefaults() *MobileDeviceHardware`

NewMobileDeviceHardwareWithDefaults instantiates a new MobileDeviceHardware object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityMb

`func (o *MobileDeviceHardware) GetCapacityMb() int32`

GetCapacityMb returns the CapacityMb field if non-nil, zero value otherwise.

### GetCapacityMbOk

`func (o *MobileDeviceHardware) GetCapacityMbOk() (*int32, bool)`

GetCapacityMbOk returns a tuple with the CapacityMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityMb

`func (o *MobileDeviceHardware) SetCapacityMb(v int32)`

SetCapacityMb sets CapacityMb field to given value.

### HasCapacityMb

`func (o *MobileDeviceHardware) HasCapacityMb() bool`

HasCapacityMb returns a boolean if a field has been set.

### GetAvailableSpaceMb

`func (o *MobileDeviceHardware) GetAvailableSpaceMb() int32`

GetAvailableSpaceMb returns the AvailableSpaceMb field if non-nil, zero value otherwise.

### GetAvailableSpaceMbOk

`func (o *MobileDeviceHardware) GetAvailableSpaceMbOk() (*int32, bool)`

GetAvailableSpaceMbOk returns a tuple with the AvailableSpaceMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableSpaceMb

`func (o *MobileDeviceHardware) SetAvailableSpaceMb(v int32)`

SetAvailableSpaceMb sets AvailableSpaceMb field to given value.

### HasAvailableSpaceMb

`func (o *MobileDeviceHardware) HasAvailableSpaceMb() bool`

HasAvailableSpaceMb returns a boolean if a field has been set.

### GetUsedSpacePercentage

`func (o *MobileDeviceHardware) GetUsedSpacePercentage() int32`

GetUsedSpacePercentage returns the UsedSpacePercentage field if non-nil, zero value otherwise.

### GetUsedSpacePercentageOk

`func (o *MobileDeviceHardware) GetUsedSpacePercentageOk() (*int32, bool)`

GetUsedSpacePercentageOk returns a tuple with the UsedSpacePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSpacePercentage

`func (o *MobileDeviceHardware) SetUsedSpacePercentage(v int32)`

SetUsedSpacePercentage sets UsedSpacePercentage field to given value.

### HasUsedSpacePercentage

`func (o *MobileDeviceHardware) HasUsedSpacePercentage() bool`

HasUsedSpacePercentage returns a boolean if a field has been set.

### GetBatteryLevel

`func (o *MobileDeviceHardware) GetBatteryLevel() int32`

GetBatteryLevel returns the BatteryLevel field if non-nil, zero value otherwise.

### GetBatteryLevelOk

`func (o *MobileDeviceHardware) GetBatteryLevelOk() (*int32, bool)`

GetBatteryLevelOk returns a tuple with the BatteryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryLevel

`func (o *MobileDeviceHardware) SetBatteryLevel(v int32)`

SetBatteryLevel sets BatteryLevel field to given value.

### HasBatteryLevel

`func (o *MobileDeviceHardware) HasBatteryLevel() bool`

HasBatteryLevel returns a boolean if a field has been set.

### GetSerialNumber

`func (o *MobileDeviceHardware) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *MobileDeviceHardware) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *MobileDeviceHardware) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *MobileDeviceHardware) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetWifiMacAddress

`func (o *MobileDeviceHardware) GetWifiMacAddress() string`

GetWifiMacAddress returns the WifiMacAddress field if non-nil, zero value otherwise.

### GetWifiMacAddressOk

`func (o *MobileDeviceHardware) GetWifiMacAddressOk() (*string, bool)`

GetWifiMacAddressOk returns a tuple with the WifiMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMacAddress

`func (o *MobileDeviceHardware) SetWifiMacAddress(v string)`

SetWifiMacAddress sets WifiMacAddress field to given value.

### HasWifiMacAddress

`func (o *MobileDeviceHardware) HasWifiMacAddress() bool`

HasWifiMacAddress returns a boolean if a field has been set.

### GetBluetoothMacAddress

`func (o *MobileDeviceHardware) GetBluetoothMacAddress() string`

GetBluetoothMacAddress returns the BluetoothMacAddress field if non-nil, zero value otherwise.

### GetBluetoothMacAddressOk

`func (o *MobileDeviceHardware) GetBluetoothMacAddressOk() (*string, bool)`

GetBluetoothMacAddressOk returns a tuple with the BluetoothMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBluetoothMacAddress

`func (o *MobileDeviceHardware) SetBluetoothMacAddress(v string)`

SetBluetoothMacAddress sets BluetoothMacAddress field to given value.

### HasBluetoothMacAddress

`func (o *MobileDeviceHardware) HasBluetoothMacAddress() bool`

HasBluetoothMacAddress returns a boolean if a field has been set.

### GetModemFirmwareVersion

`func (o *MobileDeviceHardware) GetModemFirmwareVersion() string`

GetModemFirmwareVersion returns the ModemFirmwareVersion field if non-nil, zero value otherwise.

### GetModemFirmwareVersionOk

`func (o *MobileDeviceHardware) GetModemFirmwareVersionOk() (*string, bool)`

GetModemFirmwareVersionOk returns a tuple with the ModemFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModemFirmwareVersion

`func (o *MobileDeviceHardware) SetModemFirmwareVersion(v string)`

SetModemFirmwareVersion sets ModemFirmwareVersion field to given value.

### HasModemFirmwareVersion

`func (o *MobileDeviceHardware) HasModemFirmwareVersion() bool`

HasModemFirmwareVersion returns a boolean if a field has been set.

### GetModel

`func (o *MobileDeviceHardware) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MobileDeviceHardware) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MobileDeviceHardware) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *MobileDeviceHardware) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelIdentifier

`func (o *MobileDeviceHardware) GetModelIdentifier() string`

GetModelIdentifier returns the ModelIdentifier field if non-nil, zero value otherwise.

### GetModelIdentifierOk

`func (o *MobileDeviceHardware) GetModelIdentifierOk() (*string, bool)`

GetModelIdentifierOk returns a tuple with the ModelIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelIdentifier

`func (o *MobileDeviceHardware) SetModelIdentifier(v string)`

SetModelIdentifier sets ModelIdentifier field to given value.

### HasModelIdentifier

`func (o *MobileDeviceHardware) HasModelIdentifier() bool`

HasModelIdentifier returns a boolean if a field has been set.

### GetModelNumber

`func (o *MobileDeviceHardware) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *MobileDeviceHardware) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *MobileDeviceHardware) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *MobileDeviceHardware) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetBluetoothLowEnergyCapable

`func (o *MobileDeviceHardware) GetBluetoothLowEnergyCapable() bool`

GetBluetoothLowEnergyCapable returns the BluetoothLowEnergyCapable field if non-nil, zero value otherwise.

### GetBluetoothLowEnergyCapableOk

`func (o *MobileDeviceHardware) GetBluetoothLowEnergyCapableOk() (*bool, bool)`

GetBluetoothLowEnergyCapableOk returns a tuple with the BluetoothLowEnergyCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBluetoothLowEnergyCapable

`func (o *MobileDeviceHardware) SetBluetoothLowEnergyCapable(v bool)`

SetBluetoothLowEnergyCapable sets BluetoothLowEnergyCapable field to given value.

### HasBluetoothLowEnergyCapable

`func (o *MobileDeviceHardware) HasBluetoothLowEnergyCapable() bool`

HasBluetoothLowEnergyCapable returns a boolean if a field has been set.

### GetDeviceId

`func (o *MobileDeviceHardware) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *MobileDeviceHardware) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *MobileDeviceHardware) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *MobileDeviceHardware) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *MobileDeviceHardware) GetExtensionAttributes() []MobileDeviceExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *MobileDeviceHardware) GetExtensionAttributesOk() (*[]MobileDeviceExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *MobileDeviceHardware) SetExtensionAttributes(v []MobileDeviceExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *MobileDeviceHardware) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


