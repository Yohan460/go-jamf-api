# AndroidDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OsName** | Pointer to **string** |  | [optional] 
**Manufacturer** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**InternalCapacityMb** | Pointer to **int32** |  | [optional] 
**InternalAvailableMb** | Pointer to **int32** |  | [optional] 
**InternalPercentUsed** | Pointer to **int32** |  | [optional] 
**ExternalCapacityMb** | Pointer to **int32** |  | [optional] 
**ExternalAvailableMb** | Pointer to **int32** |  | [optional] 
**ExternalPercentUsed** | Pointer to **int32** |  | [optional] 
**BatteryLevel** | Pointer to **int32** |  | [optional] 
**LastBackupTimestamp** | Pointer to **time.Time** |  | [optional] 
**ApiVersion** | Pointer to **int32** |  | [optional] 
**Computer** | Pointer to [**IdAndName**](IdAndName.md) |  | [optional] 
**Security** | Pointer to [**Security**](Security.md) |  | [optional] 

## Methods

### NewAndroidDetails

`func NewAndroidDetails() *AndroidDetails`

NewAndroidDetails instantiates a new AndroidDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAndroidDetailsWithDefaults

`func NewAndroidDetailsWithDefaults() *AndroidDetails`

NewAndroidDetailsWithDefaults instantiates a new AndroidDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOsName

`func (o *AndroidDetails) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *AndroidDetails) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *AndroidDetails) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *AndroidDetails) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetManufacturer

`func (o *AndroidDetails) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *AndroidDetails) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *AndroidDetails) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *AndroidDetails) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetModel

`func (o *AndroidDetails) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AndroidDetails) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AndroidDetails) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AndroidDetails) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetInternalCapacityMb

`func (o *AndroidDetails) GetInternalCapacityMb() int32`

GetInternalCapacityMb returns the InternalCapacityMb field if non-nil, zero value otherwise.

### GetInternalCapacityMbOk

`func (o *AndroidDetails) GetInternalCapacityMbOk() (*int32, bool)`

GetInternalCapacityMbOk returns a tuple with the InternalCapacityMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalCapacityMb

`func (o *AndroidDetails) SetInternalCapacityMb(v int32)`

SetInternalCapacityMb sets InternalCapacityMb field to given value.

### HasInternalCapacityMb

`func (o *AndroidDetails) HasInternalCapacityMb() bool`

HasInternalCapacityMb returns a boolean if a field has been set.

### GetInternalAvailableMb

`func (o *AndroidDetails) GetInternalAvailableMb() int32`

GetInternalAvailableMb returns the InternalAvailableMb field if non-nil, zero value otherwise.

### GetInternalAvailableMbOk

`func (o *AndroidDetails) GetInternalAvailableMbOk() (*int32, bool)`

GetInternalAvailableMbOk returns a tuple with the InternalAvailableMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAvailableMb

`func (o *AndroidDetails) SetInternalAvailableMb(v int32)`

SetInternalAvailableMb sets InternalAvailableMb field to given value.

### HasInternalAvailableMb

`func (o *AndroidDetails) HasInternalAvailableMb() bool`

HasInternalAvailableMb returns a boolean if a field has been set.

### GetInternalPercentUsed

`func (o *AndroidDetails) GetInternalPercentUsed() int32`

GetInternalPercentUsed returns the InternalPercentUsed field if non-nil, zero value otherwise.

### GetInternalPercentUsedOk

`func (o *AndroidDetails) GetInternalPercentUsedOk() (*int32, bool)`

GetInternalPercentUsedOk returns a tuple with the InternalPercentUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalPercentUsed

`func (o *AndroidDetails) SetInternalPercentUsed(v int32)`

SetInternalPercentUsed sets InternalPercentUsed field to given value.

### HasInternalPercentUsed

`func (o *AndroidDetails) HasInternalPercentUsed() bool`

HasInternalPercentUsed returns a boolean if a field has been set.

### GetExternalCapacityMb

`func (o *AndroidDetails) GetExternalCapacityMb() int32`

GetExternalCapacityMb returns the ExternalCapacityMb field if non-nil, zero value otherwise.

### GetExternalCapacityMbOk

`func (o *AndroidDetails) GetExternalCapacityMbOk() (*int32, bool)`

GetExternalCapacityMbOk returns a tuple with the ExternalCapacityMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCapacityMb

`func (o *AndroidDetails) SetExternalCapacityMb(v int32)`

SetExternalCapacityMb sets ExternalCapacityMb field to given value.

### HasExternalCapacityMb

`func (o *AndroidDetails) HasExternalCapacityMb() bool`

HasExternalCapacityMb returns a boolean if a field has been set.

### GetExternalAvailableMb

`func (o *AndroidDetails) GetExternalAvailableMb() int32`

GetExternalAvailableMb returns the ExternalAvailableMb field if non-nil, zero value otherwise.

### GetExternalAvailableMbOk

`func (o *AndroidDetails) GetExternalAvailableMbOk() (*int32, bool)`

GetExternalAvailableMbOk returns a tuple with the ExternalAvailableMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAvailableMb

`func (o *AndroidDetails) SetExternalAvailableMb(v int32)`

SetExternalAvailableMb sets ExternalAvailableMb field to given value.

### HasExternalAvailableMb

`func (o *AndroidDetails) HasExternalAvailableMb() bool`

HasExternalAvailableMb returns a boolean if a field has been set.

### GetExternalPercentUsed

`func (o *AndroidDetails) GetExternalPercentUsed() int32`

GetExternalPercentUsed returns the ExternalPercentUsed field if non-nil, zero value otherwise.

### GetExternalPercentUsedOk

`func (o *AndroidDetails) GetExternalPercentUsedOk() (*int32, bool)`

GetExternalPercentUsedOk returns a tuple with the ExternalPercentUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPercentUsed

`func (o *AndroidDetails) SetExternalPercentUsed(v int32)`

SetExternalPercentUsed sets ExternalPercentUsed field to given value.

### HasExternalPercentUsed

`func (o *AndroidDetails) HasExternalPercentUsed() bool`

HasExternalPercentUsed returns a boolean if a field has been set.

### GetBatteryLevel

`func (o *AndroidDetails) GetBatteryLevel() int32`

GetBatteryLevel returns the BatteryLevel field if non-nil, zero value otherwise.

### GetBatteryLevelOk

`func (o *AndroidDetails) GetBatteryLevelOk() (*int32, bool)`

GetBatteryLevelOk returns a tuple with the BatteryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryLevel

`func (o *AndroidDetails) SetBatteryLevel(v int32)`

SetBatteryLevel sets BatteryLevel field to given value.

### HasBatteryLevel

`func (o *AndroidDetails) HasBatteryLevel() bool`

HasBatteryLevel returns a boolean if a field has been set.

### GetLastBackupTimestamp

`func (o *AndroidDetails) GetLastBackupTimestamp() time.Time`

GetLastBackupTimestamp returns the LastBackupTimestamp field if non-nil, zero value otherwise.

### GetLastBackupTimestampOk

`func (o *AndroidDetails) GetLastBackupTimestampOk() (*time.Time, bool)`

GetLastBackupTimestampOk returns a tuple with the LastBackupTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackupTimestamp

`func (o *AndroidDetails) SetLastBackupTimestamp(v time.Time)`

SetLastBackupTimestamp sets LastBackupTimestamp field to given value.

### HasLastBackupTimestamp

`func (o *AndroidDetails) HasLastBackupTimestamp() bool`

HasLastBackupTimestamp returns a boolean if a field has been set.

### GetApiVersion

`func (o *AndroidDetails) GetApiVersion() int32`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AndroidDetails) GetApiVersionOk() (*int32, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AndroidDetails) SetApiVersion(v int32)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AndroidDetails) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetComputer

`func (o *AndroidDetails) GetComputer() IdAndName`

GetComputer returns the Computer field if non-nil, zero value otherwise.

### GetComputerOk

`func (o *AndroidDetails) GetComputerOk() (*IdAndName, bool)`

GetComputerOk returns a tuple with the Computer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputer

`func (o *AndroidDetails) SetComputer(v IdAndName)`

SetComputer sets Computer field to given value.

### HasComputer

`func (o *AndroidDetails) HasComputer() bool`

HasComputer returns a boolean if a field has been set.

### GetSecurity

`func (o *AndroidDetails) GetSecurity() Security`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *AndroidDetails) GetSecurityOk() (*Security, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *AndroidDetails) SetSecurity(v Security)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *AndroidDetails) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


