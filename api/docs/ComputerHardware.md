# ComputerHardware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Make** | Pointer to **string** |  | [optional] [readonly] 
**Model** | Pointer to **string** |  | [optional] [readonly] 
**ModelIdentifier** | Pointer to **string** |  | [optional] [readonly] 
**SerialNumber** | Pointer to **string** |  | [optional] [readonly] 
**ProcessorSpeedMhz** | Pointer to **int64** | Processor Speed in MHz. | [optional] [readonly] 
**ProcessorCount** | Pointer to **int32** |  | [optional] [readonly] 
**CoreCount** | Pointer to **int32** |  | [optional] [readonly] 
**ProcessorType** | Pointer to **string** |  | [optional] [readonly] 
**ProcessorArchitecture** | Pointer to **string** |  | [optional] [readonly] 
**BusSpeedMhz** | Pointer to **int64** |  | [optional] [readonly] 
**CacheSizeKilobytes** | Pointer to **int64** | Cache Size in KB. | [optional] [readonly] 
**NetworkAdapterType** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**AltNetworkAdapterType** | Pointer to **string** |  | [optional] 
**AltMacAddress** | Pointer to **string** |  | [optional] 
**TotalRamMegabytes** | Pointer to **int64** | Total RAM Size in MB. | [optional] [readonly] 
**OpenRamSlots** | Pointer to **int32** | Available RAM slots. | [optional] [readonly] 
**BatteryCapacityPercent** | Pointer to **int32** | Remaining percentage of battery power. | [optional] [readonly] 
**SmcVersion** | Pointer to **string** |  | [optional] [readonly] 
**NicSpeed** | Pointer to **string** |  | [optional] [readonly] 
**OpticalDrive** | Pointer to **string** |  | [optional] [readonly] 
**BootRom** | Pointer to **string** |  | [optional] [readonly] 
**BleCapable** | Pointer to **bool** |  | [optional] [readonly] 
**SupportsIosAppInstalls** | Pointer to **bool** |  | [optional] [readonly] 
**AppleSilicon** | Pointer to **bool** |  | [optional] [readonly] 
**ExtensionAttributes** | Pointer to [**[]ComputerExtensionAttribute**](ComputerExtensionAttribute.md) |  | [optional] 

## Methods

### NewComputerHardware

`func NewComputerHardware() *ComputerHardware`

NewComputerHardware instantiates a new ComputerHardware object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerHardwareWithDefaults

`func NewComputerHardwareWithDefaults() *ComputerHardware`

NewComputerHardwareWithDefaults instantiates a new ComputerHardware object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMake

`func (o *ComputerHardware) GetMake() string`

GetMake returns the Make field if non-nil, zero value otherwise.

### GetMakeOk

`func (o *ComputerHardware) GetMakeOk() (*string, bool)`

GetMakeOk returns a tuple with the Make field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMake

`func (o *ComputerHardware) SetMake(v string)`

SetMake sets Make field to given value.

### HasMake

`func (o *ComputerHardware) HasMake() bool`

HasMake returns a boolean if a field has been set.

### GetModel

`func (o *ComputerHardware) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ComputerHardware) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ComputerHardware) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ComputerHardware) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelIdentifier

`func (o *ComputerHardware) GetModelIdentifier() string`

GetModelIdentifier returns the ModelIdentifier field if non-nil, zero value otherwise.

### GetModelIdentifierOk

`func (o *ComputerHardware) GetModelIdentifierOk() (*string, bool)`

GetModelIdentifierOk returns a tuple with the ModelIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelIdentifier

`func (o *ComputerHardware) SetModelIdentifier(v string)`

SetModelIdentifier sets ModelIdentifier field to given value.

### HasModelIdentifier

`func (o *ComputerHardware) HasModelIdentifier() bool`

HasModelIdentifier returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ComputerHardware) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ComputerHardware) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ComputerHardware) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ComputerHardware) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetProcessorSpeedMhz

`func (o *ComputerHardware) GetProcessorSpeedMhz() int64`

GetProcessorSpeedMhz returns the ProcessorSpeedMhz field if non-nil, zero value otherwise.

### GetProcessorSpeedMhzOk

`func (o *ComputerHardware) GetProcessorSpeedMhzOk() (*int64, bool)`

GetProcessorSpeedMhzOk returns a tuple with the ProcessorSpeedMhz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorSpeedMhz

`func (o *ComputerHardware) SetProcessorSpeedMhz(v int64)`

SetProcessorSpeedMhz sets ProcessorSpeedMhz field to given value.

### HasProcessorSpeedMhz

`func (o *ComputerHardware) HasProcessorSpeedMhz() bool`

HasProcessorSpeedMhz returns a boolean if a field has been set.

### GetProcessorCount

`func (o *ComputerHardware) GetProcessorCount() int32`

GetProcessorCount returns the ProcessorCount field if non-nil, zero value otherwise.

### GetProcessorCountOk

`func (o *ComputerHardware) GetProcessorCountOk() (*int32, bool)`

GetProcessorCountOk returns a tuple with the ProcessorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCount

`func (o *ComputerHardware) SetProcessorCount(v int32)`

SetProcessorCount sets ProcessorCount field to given value.

### HasProcessorCount

`func (o *ComputerHardware) HasProcessorCount() bool`

HasProcessorCount returns a boolean if a field has been set.

### GetCoreCount

`func (o *ComputerHardware) GetCoreCount() int32`

GetCoreCount returns the CoreCount field if non-nil, zero value otherwise.

### GetCoreCountOk

`func (o *ComputerHardware) GetCoreCountOk() (*int32, bool)`

GetCoreCountOk returns a tuple with the CoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreCount

`func (o *ComputerHardware) SetCoreCount(v int32)`

SetCoreCount sets CoreCount field to given value.

### HasCoreCount

`func (o *ComputerHardware) HasCoreCount() bool`

HasCoreCount returns a boolean if a field has been set.

### GetProcessorType

`func (o *ComputerHardware) GetProcessorType() string`

GetProcessorType returns the ProcessorType field if non-nil, zero value otherwise.

### GetProcessorTypeOk

`func (o *ComputerHardware) GetProcessorTypeOk() (*string, bool)`

GetProcessorTypeOk returns a tuple with the ProcessorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorType

`func (o *ComputerHardware) SetProcessorType(v string)`

SetProcessorType sets ProcessorType field to given value.

### HasProcessorType

`func (o *ComputerHardware) HasProcessorType() bool`

HasProcessorType returns a boolean if a field has been set.

### GetProcessorArchitecture

`func (o *ComputerHardware) GetProcessorArchitecture() string`

GetProcessorArchitecture returns the ProcessorArchitecture field if non-nil, zero value otherwise.

### GetProcessorArchitectureOk

`func (o *ComputerHardware) GetProcessorArchitectureOk() (*string, bool)`

GetProcessorArchitectureOk returns a tuple with the ProcessorArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorArchitecture

`func (o *ComputerHardware) SetProcessorArchitecture(v string)`

SetProcessorArchitecture sets ProcessorArchitecture field to given value.

### HasProcessorArchitecture

`func (o *ComputerHardware) HasProcessorArchitecture() bool`

HasProcessorArchitecture returns a boolean if a field has been set.

### GetBusSpeedMhz

`func (o *ComputerHardware) GetBusSpeedMhz() int64`

GetBusSpeedMhz returns the BusSpeedMhz field if non-nil, zero value otherwise.

### GetBusSpeedMhzOk

`func (o *ComputerHardware) GetBusSpeedMhzOk() (*int64, bool)`

GetBusSpeedMhzOk returns a tuple with the BusSpeedMhz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusSpeedMhz

`func (o *ComputerHardware) SetBusSpeedMhz(v int64)`

SetBusSpeedMhz sets BusSpeedMhz field to given value.

### HasBusSpeedMhz

`func (o *ComputerHardware) HasBusSpeedMhz() bool`

HasBusSpeedMhz returns a boolean if a field has been set.

### GetCacheSizeKilobytes

`func (o *ComputerHardware) GetCacheSizeKilobytes() int64`

GetCacheSizeKilobytes returns the CacheSizeKilobytes field if non-nil, zero value otherwise.

### GetCacheSizeKilobytesOk

`func (o *ComputerHardware) GetCacheSizeKilobytesOk() (*int64, bool)`

GetCacheSizeKilobytesOk returns a tuple with the CacheSizeKilobytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheSizeKilobytes

`func (o *ComputerHardware) SetCacheSizeKilobytes(v int64)`

SetCacheSizeKilobytes sets CacheSizeKilobytes field to given value.

### HasCacheSizeKilobytes

`func (o *ComputerHardware) HasCacheSizeKilobytes() bool`

HasCacheSizeKilobytes returns a boolean if a field has been set.

### GetNetworkAdapterType

`func (o *ComputerHardware) GetNetworkAdapterType() string`

GetNetworkAdapterType returns the NetworkAdapterType field if non-nil, zero value otherwise.

### GetNetworkAdapterTypeOk

`func (o *ComputerHardware) GetNetworkAdapterTypeOk() (*string, bool)`

GetNetworkAdapterTypeOk returns a tuple with the NetworkAdapterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAdapterType

`func (o *ComputerHardware) SetNetworkAdapterType(v string)`

SetNetworkAdapterType sets NetworkAdapterType field to given value.

### HasNetworkAdapterType

`func (o *ComputerHardware) HasNetworkAdapterType() bool`

HasNetworkAdapterType returns a boolean if a field has been set.

### GetMacAddress

`func (o *ComputerHardware) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ComputerHardware) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ComputerHardware) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *ComputerHardware) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetAltNetworkAdapterType

`func (o *ComputerHardware) GetAltNetworkAdapterType() string`

GetAltNetworkAdapterType returns the AltNetworkAdapterType field if non-nil, zero value otherwise.

### GetAltNetworkAdapterTypeOk

`func (o *ComputerHardware) GetAltNetworkAdapterTypeOk() (*string, bool)`

GetAltNetworkAdapterTypeOk returns a tuple with the AltNetworkAdapterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNetworkAdapterType

`func (o *ComputerHardware) SetAltNetworkAdapterType(v string)`

SetAltNetworkAdapterType sets AltNetworkAdapterType field to given value.

### HasAltNetworkAdapterType

`func (o *ComputerHardware) HasAltNetworkAdapterType() bool`

HasAltNetworkAdapterType returns a boolean if a field has been set.

### GetAltMacAddress

`func (o *ComputerHardware) GetAltMacAddress() string`

GetAltMacAddress returns the AltMacAddress field if non-nil, zero value otherwise.

### GetAltMacAddressOk

`func (o *ComputerHardware) GetAltMacAddressOk() (*string, bool)`

GetAltMacAddressOk returns a tuple with the AltMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltMacAddress

`func (o *ComputerHardware) SetAltMacAddress(v string)`

SetAltMacAddress sets AltMacAddress field to given value.

### HasAltMacAddress

`func (o *ComputerHardware) HasAltMacAddress() bool`

HasAltMacAddress returns a boolean if a field has been set.

### GetTotalRamMegabytes

`func (o *ComputerHardware) GetTotalRamMegabytes() int64`

GetTotalRamMegabytes returns the TotalRamMegabytes field if non-nil, zero value otherwise.

### GetTotalRamMegabytesOk

`func (o *ComputerHardware) GetTotalRamMegabytesOk() (*int64, bool)`

GetTotalRamMegabytesOk returns a tuple with the TotalRamMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRamMegabytes

`func (o *ComputerHardware) SetTotalRamMegabytes(v int64)`

SetTotalRamMegabytes sets TotalRamMegabytes field to given value.

### HasTotalRamMegabytes

`func (o *ComputerHardware) HasTotalRamMegabytes() bool`

HasTotalRamMegabytes returns a boolean if a field has been set.

### GetOpenRamSlots

`func (o *ComputerHardware) GetOpenRamSlots() int32`

GetOpenRamSlots returns the OpenRamSlots field if non-nil, zero value otherwise.

### GetOpenRamSlotsOk

`func (o *ComputerHardware) GetOpenRamSlotsOk() (*int32, bool)`

GetOpenRamSlotsOk returns a tuple with the OpenRamSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenRamSlots

`func (o *ComputerHardware) SetOpenRamSlots(v int32)`

SetOpenRamSlots sets OpenRamSlots field to given value.

### HasOpenRamSlots

`func (o *ComputerHardware) HasOpenRamSlots() bool`

HasOpenRamSlots returns a boolean if a field has been set.

### GetBatteryCapacityPercent

`func (o *ComputerHardware) GetBatteryCapacityPercent() int32`

GetBatteryCapacityPercent returns the BatteryCapacityPercent field if non-nil, zero value otherwise.

### GetBatteryCapacityPercentOk

`func (o *ComputerHardware) GetBatteryCapacityPercentOk() (*int32, bool)`

GetBatteryCapacityPercentOk returns a tuple with the BatteryCapacityPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryCapacityPercent

`func (o *ComputerHardware) SetBatteryCapacityPercent(v int32)`

SetBatteryCapacityPercent sets BatteryCapacityPercent field to given value.

### HasBatteryCapacityPercent

`func (o *ComputerHardware) HasBatteryCapacityPercent() bool`

HasBatteryCapacityPercent returns a boolean if a field has been set.

### GetSmcVersion

`func (o *ComputerHardware) GetSmcVersion() string`

GetSmcVersion returns the SmcVersion field if non-nil, zero value otherwise.

### GetSmcVersionOk

`func (o *ComputerHardware) GetSmcVersionOk() (*string, bool)`

GetSmcVersionOk returns a tuple with the SmcVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmcVersion

`func (o *ComputerHardware) SetSmcVersion(v string)`

SetSmcVersion sets SmcVersion field to given value.

### HasSmcVersion

`func (o *ComputerHardware) HasSmcVersion() bool`

HasSmcVersion returns a boolean if a field has been set.

### GetNicSpeed

`func (o *ComputerHardware) GetNicSpeed() string`

GetNicSpeed returns the NicSpeed field if non-nil, zero value otherwise.

### GetNicSpeedOk

`func (o *ComputerHardware) GetNicSpeedOk() (*string, bool)`

GetNicSpeedOk returns a tuple with the NicSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicSpeed

`func (o *ComputerHardware) SetNicSpeed(v string)`

SetNicSpeed sets NicSpeed field to given value.

### HasNicSpeed

`func (o *ComputerHardware) HasNicSpeed() bool`

HasNicSpeed returns a boolean if a field has been set.

### GetOpticalDrive

`func (o *ComputerHardware) GetOpticalDrive() string`

GetOpticalDrive returns the OpticalDrive field if non-nil, zero value otherwise.

### GetOpticalDriveOk

`func (o *ComputerHardware) GetOpticalDriveOk() (*string, bool)`

GetOpticalDriveOk returns a tuple with the OpticalDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpticalDrive

`func (o *ComputerHardware) SetOpticalDrive(v string)`

SetOpticalDrive sets OpticalDrive field to given value.

### HasOpticalDrive

`func (o *ComputerHardware) HasOpticalDrive() bool`

HasOpticalDrive returns a boolean if a field has been set.

### GetBootRom

`func (o *ComputerHardware) GetBootRom() string`

GetBootRom returns the BootRom field if non-nil, zero value otherwise.

### GetBootRomOk

`func (o *ComputerHardware) GetBootRomOk() (*string, bool)`

GetBootRomOk returns a tuple with the BootRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootRom

`func (o *ComputerHardware) SetBootRom(v string)`

SetBootRom sets BootRom field to given value.

### HasBootRom

`func (o *ComputerHardware) HasBootRom() bool`

HasBootRom returns a boolean if a field has been set.

### GetBleCapable

`func (o *ComputerHardware) GetBleCapable() bool`

GetBleCapable returns the BleCapable field if non-nil, zero value otherwise.

### GetBleCapableOk

`func (o *ComputerHardware) GetBleCapableOk() (*bool, bool)`

GetBleCapableOk returns a tuple with the BleCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBleCapable

`func (o *ComputerHardware) SetBleCapable(v bool)`

SetBleCapable sets BleCapable field to given value.

### HasBleCapable

`func (o *ComputerHardware) HasBleCapable() bool`

HasBleCapable returns a boolean if a field has been set.

### GetSupportsIosAppInstalls

`func (o *ComputerHardware) GetSupportsIosAppInstalls() bool`

GetSupportsIosAppInstalls returns the SupportsIosAppInstalls field if non-nil, zero value otherwise.

### GetSupportsIosAppInstallsOk

`func (o *ComputerHardware) GetSupportsIosAppInstallsOk() (*bool, bool)`

GetSupportsIosAppInstallsOk returns a tuple with the SupportsIosAppInstalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsIosAppInstalls

`func (o *ComputerHardware) SetSupportsIosAppInstalls(v bool)`

SetSupportsIosAppInstalls sets SupportsIosAppInstalls field to given value.

### HasSupportsIosAppInstalls

`func (o *ComputerHardware) HasSupportsIosAppInstalls() bool`

HasSupportsIosAppInstalls returns a boolean if a field has been set.

### GetAppleSilicon

`func (o *ComputerHardware) GetAppleSilicon() bool`

GetAppleSilicon returns the AppleSilicon field if non-nil, zero value otherwise.

### GetAppleSiliconOk

`func (o *ComputerHardware) GetAppleSiliconOk() (*bool, bool)`

GetAppleSiliconOk returns a tuple with the AppleSilicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleSilicon

`func (o *ComputerHardware) SetAppleSilicon(v bool)`

SetAppleSilicon sets AppleSilicon field to given value.

### HasAppleSilicon

`func (o *ComputerHardware) HasAppleSilicon() bool`

HasAppleSilicon returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *ComputerHardware) GetExtensionAttributes() []ComputerExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *ComputerHardware) GetExtensionAttributesOk() (*[]ComputerExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *ComputerHardware) SetExtensionAttributes(v []ComputerExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *ComputerHardware) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


