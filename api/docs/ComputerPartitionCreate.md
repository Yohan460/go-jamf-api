# ComputerPartitionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**SizeMegabytes** | Pointer to **NullableInt64** | Partition Size in MB. | [optional] 
**AvailableMegabytes** | Pointer to **NullableInt64** | Available space in MB. | [optional] 
**PartitionType** | Pointer to **NullableString** |  | [optional] 
**PercentUsed** | Pointer to **NullableInt64** | Percentage of space used. | [optional] 
**FileVault2State** | Pointer to [**ComputerPartitionFileVault2State**](ComputerPartitionFileVault2State.md) |  | [optional] 
**FileVault2ProgressPercent** | Pointer to **NullableInt64** | Percentage progress of current FileVault 2 operation. | [optional] 
**LvmManaged** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewComputerPartitionCreate

`func NewComputerPartitionCreate() *ComputerPartitionCreate`

NewComputerPartitionCreate instantiates a new ComputerPartitionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerPartitionCreateWithDefaults

`func NewComputerPartitionCreateWithDefaults() *ComputerPartitionCreate`

NewComputerPartitionCreateWithDefaults instantiates a new ComputerPartitionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ComputerPartitionCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputerPartitionCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputerPartitionCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputerPartitionCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ComputerPartitionCreate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ComputerPartitionCreate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSizeMegabytes

`func (o *ComputerPartitionCreate) GetSizeMegabytes() int64`

GetSizeMegabytes returns the SizeMegabytes field if non-nil, zero value otherwise.

### GetSizeMegabytesOk

`func (o *ComputerPartitionCreate) GetSizeMegabytesOk() (*int64, bool)`

GetSizeMegabytesOk returns a tuple with the SizeMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMegabytes

`func (o *ComputerPartitionCreate) SetSizeMegabytes(v int64)`

SetSizeMegabytes sets SizeMegabytes field to given value.

### HasSizeMegabytes

`func (o *ComputerPartitionCreate) HasSizeMegabytes() bool`

HasSizeMegabytes returns a boolean if a field has been set.

### SetSizeMegabytesNil

`func (o *ComputerPartitionCreate) SetSizeMegabytesNil(b bool)`

 SetSizeMegabytesNil sets the value for SizeMegabytes to be an explicit nil

### UnsetSizeMegabytes
`func (o *ComputerPartitionCreate) UnsetSizeMegabytes()`

UnsetSizeMegabytes ensures that no value is present for SizeMegabytes, not even an explicit nil
### GetAvailableMegabytes

`func (o *ComputerPartitionCreate) GetAvailableMegabytes() int64`

GetAvailableMegabytes returns the AvailableMegabytes field if non-nil, zero value otherwise.

### GetAvailableMegabytesOk

`func (o *ComputerPartitionCreate) GetAvailableMegabytesOk() (*int64, bool)`

GetAvailableMegabytesOk returns a tuple with the AvailableMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMegabytes

`func (o *ComputerPartitionCreate) SetAvailableMegabytes(v int64)`

SetAvailableMegabytes sets AvailableMegabytes field to given value.

### HasAvailableMegabytes

`func (o *ComputerPartitionCreate) HasAvailableMegabytes() bool`

HasAvailableMegabytes returns a boolean if a field has been set.

### SetAvailableMegabytesNil

`func (o *ComputerPartitionCreate) SetAvailableMegabytesNil(b bool)`

 SetAvailableMegabytesNil sets the value for AvailableMegabytes to be an explicit nil

### UnsetAvailableMegabytes
`func (o *ComputerPartitionCreate) UnsetAvailableMegabytes()`

UnsetAvailableMegabytes ensures that no value is present for AvailableMegabytes, not even an explicit nil
### GetPartitionType

`func (o *ComputerPartitionCreate) GetPartitionType() string`

GetPartitionType returns the PartitionType field if non-nil, zero value otherwise.

### GetPartitionTypeOk

`func (o *ComputerPartitionCreate) GetPartitionTypeOk() (*string, bool)`

GetPartitionTypeOk returns a tuple with the PartitionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionType

`func (o *ComputerPartitionCreate) SetPartitionType(v string)`

SetPartitionType sets PartitionType field to given value.

### HasPartitionType

`func (o *ComputerPartitionCreate) HasPartitionType() bool`

HasPartitionType returns a boolean if a field has been set.

### SetPartitionTypeNil

`func (o *ComputerPartitionCreate) SetPartitionTypeNil(b bool)`

 SetPartitionTypeNil sets the value for PartitionType to be an explicit nil

### UnsetPartitionType
`func (o *ComputerPartitionCreate) UnsetPartitionType()`

UnsetPartitionType ensures that no value is present for PartitionType, not even an explicit nil
### GetPercentUsed

`func (o *ComputerPartitionCreate) GetPercentUsed() int64`

GetPercentUsed returns the PercentUsed field if non-nil, zero value otherwise.

### GetPercentUsedOk

`func (o *ComputerPartitionCreate) GetPercentUsedOk() (*int64, bool)`

GetPercentUsedOk returns a tuple with the PercentUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentUsed

`func (o *ComputerPartitionCreate) SetPercentUsed(v int64)`

SetPercentUsed sets PercentUsed field to given value.

### HasPercentUsed

`func (o *ComputerPartitionCreate) HasPercentUsed() bool`

HasPercentUsed returns a boolean if a field has been set.

### SetPercentUsedNil

`func (o *ComputerPartitionCreate) SetPercentUsedNil(b bool)`

 SetPercentUsedNil sets the value for PercentUsed to be an explicit nil

### UnsetPercentUsed
`func (o *ComputerPartitionCreate) UnsetPercentUsed()`

UnsetPercentUsed ensures that no value is present for PercentUsed, not even an explicit nil
### GetFileVault2State

`func (o *ComputerPartitionCreate) GetFileVault2State() ComputerPartitionFileVault2State`

GetFileVault2State returns the FileVault2State field if non-nil, zero value otherwise.

### GetFileVault2StateOk

`func (o *ComputerPartitionCreate) GetFileVault2StateOk() (*ComputerPartitionFileVault2State, bool)`

GetFileVault2StateOk returns a tuple with the FileVault2State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileVault2State

`func (o *ComputerPartitionCreate) SetFileVault2State(v ComputerPartitionFileVault2State)`

SetFileVault2State sets FileVault2State field to given value.

### HasFileVault2State

`func (o *ComputerPartitionCreate) HasFileVault2State() bool`

HasFileVault2State returns a boolean if a field has been set.

### GetFileVault2ProgressPercent

`func (o *ComputerPartitionCreate) GetFileVault2ProgressPercent() int64`

GetFileVault2ProgressPercent returns the FileVault2ProgressPercent field if non-nil, zero value otherwise.

### GetFileVault2ProgressPercentOk

`func (o *ComputerPartitionCreate) GetFileVault2ProgressPercentOk() (*int64, bool)`

GetFileVault2ProgressPercentOk returns a tuple with the FileVault2ProgressPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileVault2ProgressPercent

`func (o *ComputerPartitionCreate) SetFileVault2ProgressPercent(v int64)`

SetFileVault2ProgressPercent sets FileVault2ProgressPercent field to given value.

### HasFileVault2ProgressPercent

`func (o *ComputerPartitionCreate) HasFileVault2ProgressPercent() bool`

HasFileVault2ProgressPercent returns a boolean if a field has been set.

### SetFileVault2ProgressPercentNil

`func (o *ComputerPartitionCreate) SetFileVault2ProgressPercentNil(b bool)`

 SetFileVault2ProgressPercentNil sets the value for FileVault2ProgressPercent to be an explicit nil

### UnsetFileVault2ProgressPercent
`func (o *ComputerPartitionCreate) UnsetFileVault2ProgressPercent()`

UnsetFileVault2ProgressPercent ensures that no value is present for FileVault2ProgressPercent, not even an explicit nil
### GetLvmManaged

`func (o *ComputerPartitionCreate) GetLvmManaged() bool`

GetLvmManaged returns the LvmManaged field if non-nil, zero value otherwise.

### GetLvmManagedOk

`func (o *ComputerPartitionCreate) GetLvmManagedOk() (*bool, bool)`

GetLvmManagedOk returns a tuple with the LvmManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmManaged

`func (o *ComputerPartitionCreate) SetLvmManaged(v bool)`

SetLvmManaged sets LvmManaged field to given value.

### HasLvmManaged

`func (o *ComputerPartitionCreate) HasLvmManaged() bool`

HasLvmManaged returns a boolean if a field has been set.

### SetLvmManagedNil

`func (o *ComputerPartitionCreate) SetLvmManagedNil(b bool)`

 SetLvmManagedNil sets the value for LvmManaged to be an explicit nil

### UnsetLvmManaged
`func (o *ComputerPartitionCreate) UnsetLvmManaged()`

UnsetLvmManaged ensures that no value is present for LvmManaged, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


