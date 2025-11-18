# ComputerHardwareCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Make** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**ModelIdentifier** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**ProcessorSpeedMhz** | Pointer to **int64** | Processor Speed in MHz. | [optional] 
**ProcessorCount** | Pointer to **int64** |  | [optional] 
**CoreCount** | Pointer to **int64** |  | [optional] 
**ProcessorType** | Pointer to **string** |  | [optional] 
**ProcessorArchitecture** | Pointer to **string** |  | [optional] 
**BusSpeedMhz** | Pointer to **int64** |  | [optional] 
**CacheSizeKilobytes** | Pointer to **int64** | Cache Size in KB. | [optional] 
**NetworkAdapterType** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**AltNetworkAdapterType** | Pointer to **string** |  | [optional] 
**AltMacAddress** | Pointer to **string** |  | [optional] 
**TotalRamMegabytes** | Pointer to **int64** | Total RAM Size in MB. | [optional] 
**OpenRamSlots** | Pointer to **int64** | Available RAM slots. | [optional] 
**BatteryCapacityPercent** | Pointer to **int64** | Remaining percentage of battery power. | [optional] 
**BatteryHealth** | Pointer to **string** | - NON_GENUINE: The battery isn’t a genuine Apple battery. - NORMAL: The battery is operating normally. - SERVICE_RECOMMENDED: The system recommends battery service. - UNKNOWN: The system couldn’t determine battery health information. - UNSUPPORTED: The device doesn’t support battery health reporting.  | [optional] [default to "UNKNOWN"]
**SmcVersion** | Pointer to **string** |  | [optional] 
**NicSpeed** | Pointer to **string** |  | [optional] 
**OpticalDrive** | Pointer to **string** |  | [optional] 
**BootRom** | Pointer to **string** |  | [optional] 
**BleCapable** | Pointer to **bool** |  | [optional] 
**SupportsIosAppInstalls** | Pointer to **bool** |  | [optional] 
**AppleSilicon** | Pointer to **bool** |  | [optional] 

## Methods

### NewComputerHardwareCreate

`func NewComputerHardwareCreate() *ComputerHardwareCreate`

NewComputerHardwareCreate instantiates a new ComputerHardwareCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerHardwareCreateWithDefaults

`func NewComputerHardwareCreateWithDefaults() *ComputerHardwareCreate`

NewComputerHardwareCreateWithDefaults instantiates a new ComputerHardwareCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMake

`func (o *ComputerHardwareCreate) GetMake() string`

GetMake returns the Make field if non-nil, zero value otherwise.

### GetMakeOk

`func (o *ComputerHardwareCreate) GetMakeOk() (*string, bool)`

GetMakeOk returns a tuple with the Make field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMake

`func (o *ComputerHardwareCreate) SetMake(v string)`

SetMake sets Make field to given value.

### HasMake

`func (o *ComputerHardwareCreate) HasMake() bool`

HasMake returns a boolean if a field has been set.

### GetModel

`func (o *ComputerHardwareCreate) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ComputerHardwareCreate) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ComputerHardwareCreate) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ComputerHardwareCreate) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelIdentifier

`func (o *ComputerHardwareCreate) GetModelIdentifier() string`

GetModelIdentifier returns the ModelIdentifier field if non-nil, zero value otherwise.

### GetModelIdentifierOk

`func (o *ComputerHardwareCreate) GetModelIdentifierOk() (*string, bool)`

GetModelIdentifierOk returns a tuple with the ModelIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelIdentifier

`func (o *ComputerHardwareCreate) SetModelIdentifier(v string)`

SetModelIdentifier sets ModelIdentifier field to given value.

### HasModelIdentifier

`func (o *ComputerHardwareCreate) HasModelIdentifier() bool`

HasModelIdentifier returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ComputerHardwareCreate) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ComputerHardwareCreate) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ComputerHardwareCreate) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ComputerHardwareCreate) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetProcessorSpeedMhz

`func (o *ComputerHardwareCreate) GetProcessorSpeedMhz() int64`

GetProcessorSpeedMhz returns the ProcessorSpeedMhz field if non-nil, zero value otherwise.

### GetProcessorSpeedMhzOk

`func (o *ComputerHardwareCreate) GetProcessorSpeedMhzOk() (*int64, bool)`

GetProcessorSpeedMhzOk returns a tuple with the ProcessorSpeedMhz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorSpeedMhz

`func (o *ComputerHardwareCreate) SetProcessorSpeedMhz(v int64)`

SetProcessorSpeedMhz sets ProcessorSpeedMhz field to given value.

### HasProcessorSpeedMhz

`func (o *ComputerHardwareCreate) HasProcessorSpeedMhz() bool`

HasProcessorSpeedMhz returns a boolean if a field has been set.

### GetProcessorCount

`func (o *ComputerHardwareCreate) GetProcessorCount() int64`

GetProcessorCount returns the ProcessorCount field if non-nil, zero value otherwise.

### GetProcessorCountOk

`func (o *ComputerHardwareCreate) GetProcessorCountOk() (*int64, bool)`

GetProcessorCountOk returns a tuple with the ProcessorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCount

`func (o *ComputerHardwareCreate) SetProcessorCount(v int64)`

SetProcessorCount sets ProcessorCount field to given value.

### HasProcessorCount

`func (o *ComputerHardwareCreate) HasProcessorCount() bool`

HasProcessorCount returns a boolean if a field has been set.

### GetCoreCount

`func (o *ComputerHardwareCreate) GetCoreCount() int64`

GetCoreCount returns the CoreCount field if non-nil, zero value otherwise.

### GetCoreCountOk

`func (o *ComputerHardwareCreate) GetCoreCountOk() (*int64, bool)`

GetCoreCountOk returns a tuple with the CoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreCount

`func (o *ComputerHardwareCreate) SetCoreCount(v int64)`

SetCoreCount sets CoreCount field to given value.

### HasCoreCount

`func (o *ComputerHardwareCreate) HasCoreCount() bool`

HasCoreCount returns a boolean if a field has been set.

### GetProcessorType

`func (o *ComputerHardwareCreate) GetProcessorType() string`

GetProcessorType returns the ProcessorType field if non-nil, zero value otherwise.

### GetProcessorTypeOk

`func (o *ComputerHardwareCreate) GetProcessorTypeOk() (*string, bool)`

GetProcessorTypeOk returns a tuple with the ProcessorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorType

`func (o *ComputerHardwareCreate) SetProcessorType(v string)`

SetProcessorType sets ProcessorType field to given value.

### HasProcessorType

`func (o *ComputerHardwareCreate) HasProcessorType() bool`

HasProcessorType returns a boolean if a field has been set.

### GetProcessorArchitecture

`func (o *ComputerHardwareCreate) GetProcessorArchitecture() string`

GetProcessorArchitecture returns the ProcessorArchitecture field if non-nil, zero value otherwise.

### GetProcessorArchitectureOk

`func (o *ComputerHardwareCreate) GetProcessorArchitectureOk() (*string, bool)`

GetProcessorArchitectureOk returns a tuple with the ProcessorArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorArchitecture

`func (o *ComputerHardwareCreate) SetProcessorArchitecture(v string)`

SetProcessorArchitecture sets ProcessorArchitecture field to given value.

### HasProcessorArchitecture

`func (o *ComputerHardwareCreate) HasProcessorArchitecture() bool`

HasProcessorArchitecture returns a boolean if a field has been set.

### GetBusSpeedMhz

`func (o *ComputerHardwareCreate) GetBusSpeedMhz() int64`

GetBusSpeedMhz returns the BusSpeedMhz field if non-nil, zero value otherwise.

### GetBusSpeedMhzOk

`func (o *ComputerHardwareCreate) GetBusSpeedMhzOk() (*int64, bool)`

GetBusSpeedMhzOk returns a tuple with the BusSpeedMhz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusSpeedMhz

`func (o *ComputerHardwareCreate) SetBusSpeedMhz(v int64)`

SetBusSpeedMhz sets BusSpeedMhz field to given value.

### HasBusSpeedMhz

`func (o *ComputerHardwareCreate) HasBusSpeedMhz() bool`

HasBusSpeedMhz returns a boolean if a field has been set.

### GetCacheSizeKilobytes

`func (o *ComputerHardwareCreate) GetCacheSizeKilobytes() int64`

GetCacheSizeKilobytes returns the CacheSizeKilobytes field if non-nil, zero value otherwise.

### GetCacheSizeKilobytesOk

`func (o *ComputerHardwareCreate) GetCacheSizeKilobytesOk() (*int64, bool)`

GetCacheSizeKilobytesOk returns a tuple with the CacheSizeKilobytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheSizeKilobytes

`func (o *ComputerHardwareCreate) SetCacheSizeKilobytes(v int64)`

SetCacheSizeKilobytes sets CacheSizeKilobytes field to given value.

### HasCacheSizeKilobytes

`func (o *ComputerHardwareCreate) HasCacheSizeKilobytes() bool`

HasCacheSizeKilobytes returns a boolean if a field has been set.

### GetNetworkAdapterType

`func (o *ComputerHardwareCreate) GetNetworkAdapterType() string`

GetNetworkAdapterType returns the NetworkAdapterType field if non-nil, zero value otherwise.

### GetNetworkAdapterTypeOk

`func (o *ComputerHardwareCreate) GetNetworkAdapterTypeOk() (*string, bool)`

GetNetworkAdapterTypeOk returns a tuple with the NetworkAdapterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAdapterType

`func (o *ComputerHardwareCreate) SetNetworkAdapterType(v string)`

SetNetworkAdapterType sets NetworkAdapterType field to given value.

### HasNetworkAdapterType

`func (o *ComputerHardwareCreate) HasNetworkAdapterType() bool`

HasNetworkAdapterType returns a boolean if a field has been set.

### GetMacAddress

`func (o *ComputerHardwareCreate) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ComputerHardwareCreate) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ComputerHardwareCreate) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *ComputerHardwareCreate) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetAltNetworkAdapterType

`func (o *ComputerHardwareCreate) GetAltNetworkAdapterType() string`

GetAltNetworkAdapterType returns the AltNetworkAdapterType field if non-nil, zero value otherwise.

### GetAltNetworkAdapterTypeOk

`func (o *ComputerHardwareCreate) GetAltNetworkAdapterTypeOk() (*string, bool)`

GetAltNetworkAdapterTypeOk returns a tuple with the AltNetworkAdapterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNetworkAdapterType

`func (o *ComputerHardwareCreate) SetAltNetworkAdapterType(v string)`

SetAltNetworkAdapterType sets AltNetworkAdapterType field to given value.

### HasAltNetworkAdapterType

`func (o *ComputerHardwareCreate) HasAltNetworkAdapterType() bool`

HasAltNetworkAdapterType returns a boolean if a field has been set.

### GetAltMacAddress

`func (o *ComputerHardwareCreate) GetAltMacAddress() string`

GetAltMacAddress returns the AltMacAddress field if non-nil, zero value otherwise.

### GetAltMacAddressOk

`func (o *ComputerHardwareCreate) GetAltMacAddressOk() (*string, bool)`

GetAltMacAddressOk returns a tuple with the AltMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltMacAddress

`func (o *ComputerHardwareCreate) SetAltMacAddress(v string)`

SetAltMacAddress sets AltMacAddress field to given value.

### HasAltMacAddress

`func (o *ComputerHardwareCreate) HasAltMacAddress() bool`

HasAltMacAddress returns a boolean if a field has been set.

### GetTotalRamMegabytes

`func (o *ComputerHardwareCreate) GetTotalRamMegabytes() int64`

GetTotalRamMegabytes returns the TotalRamMegabytes field if non-nil, zero value otherwise.

### GetTotalRamMegabytesOk

`func (o *ComputerHardwareCreate) GetTotalRamMegabytesOk() (*int64, bool)`

GetTotalRamMegabytesOk returns a tuple with the TotalRamMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRamMegabytes

`func (o *ComputerHardwareCreate) SetTotalRamMegabytes(v int64)`

SetTotalRamMegabytes sets TotalRamMegabytes field to given value.

### HasTotalRamMegabytes

`func (o *ComputerHardwareCreate) HasTotalRamMegabytes() bool`

HasTotalRamMegabytes returns a boolean if a field has been set.

### GetOpenRamSlots

`func (o *ComputerHardwareCreate) GetOpenRamSlots() int64`

GetOpenRamSlots returns the OpenRamSlots field if non-nil, zero value otherwise.

### GetOpenRamSlotsOk

`func (o *ComputerHardwareCreate) GetOpenRamSlotsOk() (*int64, bool)`

GetOpenRamSlotsOk returns a tuple with the OpenRamSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenRamSlots

`func (o *ComputerHardwareCreate) SetOpenRamSlots(v int64)`

SetOpenRamSlots sets OpenRamSlots field to given value.

### HasOpenRamSlots

`func (o *ComputerHardwareCreate) HasOpenRamSlots() bool`

HasOpenRamSlots returns a boolean if a field has been set.

### GetBatteryCapacityPercent

`func (o *ComputerHardwareCreate) GetBatteryCapacityPercent() int64`

GetBatteryCapacityPercent returns the BatteryCapacityPercent field if non-nil, zero value otherwise.

### GetBatteryCapacityPercentOk

`func (o *ComputerHardwareCreate) GetBatteryCapacityPercentOk() (*int64, bool)`

GetBatteryCapacityPercentOk returns a tuple with the BatteryCapacityPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryCapacityPercent

`func (o *ComputerHardwareCreate) SetBatteryCapacityPercent(v int64)`

SetBatteryCapacityPercent sets BatteryCapacityPercent field to given value.

### HasBatteryCapacityPercent

`func (o *ComputerHardwareCreate) HasBatteryCapacityPercent() bool`

HasBatteryCapacityPercent returns a boolean if a field has been set.

### GetBatteryHealth

`func (o *ComputerHardwareCreate) GetBatteryHealth() string`

GetBatteryHealth returns the BatteryHealth field if non-nil, zero value otherwise.

### GetBatteryHealthOk

`func (o *ComputerHardwareCreate) GetBatteryHealthOk() (*string, bool)`

GetBatteryHealthOk returns a tuple with the BatteryHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryHealth

`func (o *ComputerHardwareCreate) SetBatteryHealth(v string)`

SetBatteryHealth sets BatteryHealth field to given value.

### HasBatteryHealth

`func (o *ComputerHardwareCreate) HasBatteryHealth() bool`

HasBatteryHealth returns a boolean if a field has been set.

### GetSmcVersion

`func (o *ComputerHardwareCreate) GetSmcVersion() string`

GetSmcVersion returns the SmcVersion field if non-nil, zero value otherwise.

### GetSmcVersionOk

`func (o *ComputerHardwareCreate) GetSmcVersionOk() (*string, bool)`

GetSmcVersionOk returns a tuple with the SmcVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmcVersion

`func (o *ComputerHardwareCreate) SetSmcVersion(v string)`

SetSmcVersion sets SmcVersion field to given value.

### HasSmcVersion

`func (o *ComputerHardwareCreate) HasSmcVersion() bool`

HasSmcVersion returns a boolean if a field has been set.

### GetNicSpeed

`func (o *ComputerHardwareCreate) GetNicSpeed() string`

GetNicSpeed returns the NicSpeed field if non-nil, zero value otherwise.

### GetNicSpeedOk

`func (o *ComputerHardwareCreate) GetNicSpeedOk() (*string, bool)`

GetNicSpeedOk returns a tuple with the NicSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicSpeed

`func (o *ComputerHardwareCreate) SetNicSpeed(v string)`

SetNicSpeed sets NicSpeed field to given value.

### HasNicSpeed

`func (o *ComputerHardwareCreate) HasNicSpeed() bool`

HasNicSpeed returns a boolean if a field has been set.

### GetOpticalDrive

`func (o *ComputerHardwareCreate) GetOpticalDrive() string`

GetOpticalDrive returns the OpticalDrive field if non-nil, zero value otherwise.

### GetOpticalDriveOk

`func (o *ComputerHardwareCreate) GetOpticalDriveOk() (*string, bool)`

GetOpticalDriveOk returns a tuple with the OpticalDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpticalDrive

`func (o *ComputerHardwareCreate) SetOpticalDrive(v string)`

SetOpticalDrive sets OpticalDrive field to given value.

### HasOpticalDrive

`func (o *ComputerHardwareCreate) HasOpticalDrive() bool`

HasOpticalDrive returns a boolean if a field has been set.

### GetBootRom

`func (o *ComputerHardwareCreate) GetBootRom() string`

GetBootRom returns the BootRom field if non-nil, zero value otherwise.

### GetBootRomOk

`func (o *ComputerHardwareCreate) GetBootRomOk() (*string, bool)`

GetBootRomOk returns a tuple with the BootRom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootRom

`func (o *ComputerHardwareCreate) SetBootRom(v string)`

SetBootRom sets BootRom field to given value.

### HasBootRom

`func (o *ComputerHardwareCreate) HasBootRom() bool`

HasBootRom returns a boolean if a field has been set.

### GetBleCapable

`func (o *ComputerHardwareCreate) GetBleCapable() bool`

GetBleCapable returns the BleCapable field if non-nil, zero value otherwise.

### GetBleCapableOk

`func (o *ComputerHardwareCreate) GetBleCapableOk() (*bool, bool)`

GetBleCapableOk returns a tuple with the BleCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBleCapable

`func (o *ComputerHardwareCreate) SetBleCapable(v bool)`

SetBleCapable sets BleCapable field to given value.

### HasBleCapable

`func (o *ComputerHardwareCreate) HasBleCapable() bool`

HasBleCapable returns a boolean if a field has been set.

### GetSupportsIosAppInstalls

`func (o *ComputerHardwareCreate) GetSupportsIosAppInstalls() bool`

GetSupportsIosAppInstalls returns the SupportsIosAppInstalls field if non-nil, zero value otherwise.

### GetSupportsIosAppInstallsOk

`func (o *ComputerHardwareCreate) GetSupportsIosAppInstallsOk() (*bool, bool)`

GetSupportsIosAppInstallsOk returns a tuple with the SupportsIosAppInstalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsIosAppInstalls

`func (o *ComputerHardwareCreate) SetSupportsIosAppInstalls(v bool)`

SetSupportsIosAppInstalls sets SupportsIosAppInstalls field to given value.

### HasSupportsIosAppInstalls

`func (o *ComputerHardwareCreate) HasSupportsIosAppInstalls() bool`

HasSupportsIosAppInstalls returns a boolean if a field has been set.

### GetAppleSilicon

`func (o *ComputerHardwareCreate) GetAppleSilicon() bool`

GetAppleSilicon returns the AppleSilicon field if non-nil, zero value otherwise.

### GetAppleSiliconOk

`func (o *ComputerHardwareCreate) GetAppleSiliconOk() (*bool, bool)`

GetAppleSiliconOk returns a tuple with the AppleSilicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleSilicon

`func (o *ComputerHardwareCreate) SetAppleSilicon(v bool)`

SetAppleSilicon sets AppleSilicon field to given value.

### HasAppleSilicon

`func (o *ComputerHardwareCreate) HasAppleSilicon() bool`

HasAppleSilicon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


