# ComputerPartition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] [readonly] 
**SizeMegabytes** | Pointer to **int64** | Partition Size in MB. | [optional] [readonly] 
**AvailableMegabytes** | Pointer to **int64** | Available space in MB. | [optional] [readonly] 
**PartitionType** | Pointer to **string** |  | [optional] [readonly] 
**PercentUsed** | Pointer to **int64** | Percentage of space used. | [optional] [readonly] 
**FileVault2State** | Pointer to [**ComputerPartitionFileVault2State**](ComputerPartitionFileVault2State.md) |  | [optional] 
**FileVault2ProgressPercent** | Pointer to **NullableInt64** | Percentage progress of current FileVault 2 operation. | [optional] 
**LvmManaged** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewComputerPartition

`func NewComputerPartition() *ComputerPartition`

NewComputerPartition instantiates a new ComputerPartition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerPartitionWithDefaults

`func NewComputerPartitionWithDefaults() *ComputerPartition`

NewComputerPartitionWithDefaults instantiates a new ComputerPartition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ComputerPartition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputerPartition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputerPartition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputerPartition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSizeMegabytes

`func (o *ComputerPartition) GetSizeMegabytes() int64`

GetSizeMegabytes returns the SizeMegabytes field if non-nil, zero value otherwise.

### GetSizeMegabytesOk

`func (o *ComputerPartition) GetSizeMegabytesOk() (*int64, bool)`

GetSizeMegabytesOk returns a tuple with the SizeMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMegabytes

`func (o *ComputerPartition) SetSizeMegabytes(v int64)`

SetSizeMegabytes sets SizeMegabytes field to given value.

### HasSizeMegabytes

`func (o *ComputerPartition) HasSizeMegabytes() bool`

HasSizeMegabytes returns a boolean if a field has been set.

### GetAvailableMegabytes

`func (o *ComputerPartition) GetAvailableMegabytes() int64`

GetAvailableMegabytes returns the AvailableMegabytes field if non-nil, zero value otherwise.

### GetAvailableMegabytesOk

`func (o *ComputerPartition) GetAvailableMegabytesOk() (*int64, bool)`

GetAvailableMegabytesOk returns a tuple with the AvailableMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMegabytes

`func (o *ComputerPartition) SetAvailableMegabytes(v int64)`

SetAvailableMegabytes sets AvailableMegabytes field to given value.

### HasAvailableMegabytes

`func (o *ComputerPartition) HasAvailableMegabytes() bool`

HasAvailableMegabytes returns a boolean if a field has been set.

### GetPartitionType

`func (o *ComputerPartition) GetPartitionType() string`

GetPartitionType returns the PartitionType field if non-nil, zero value otherwise.

### GetPartitionTypeOk

`func (o *ComputerPartition) GetPartitionTypeOk() (*string, bool)`

GetPartitionTypeOk returns a tuple with the PartitionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionType

`func (o *ComputerPartition) SetPartitionType(v string)`

SetPartitionType sets PartitionType field to given value.

### HasPartitionType

`func (o *ComputerPartition) HasPartitionType() bool`

HasPartitionType returns a boolean if a field has been set.

### GetPercentUsed

`func (o *ComputerPartition) GetPercentUsed() int64`

GetPercentUsed returns the PercentUsed field if non-nil, zero value otherwise.

### GetPercentUsedOk

`func (o *ComputerPartition) GetPercentUsedOk() (*int64, bool)`

GetPercentUsedOk returns a tuple with the PercentUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentUsed

`func (o *ComputerPartition) SetPercentUsed(v int64)`

SetPercentUsed sets PercentUsed field to given value.

### HasPercentUsed

`func (o *ComputerPartition) HasPercentUsed() bool`

HasPercentUsed returns a boolean if a field has been set.

### GetFileVault2State

`func (o *ComputerPartition) GetFileVault2State() ComputerPartitionFileVault2State`

GetFileVault2State returns the FileVault2State field if non-nil, zero value otherwise.

### GetFileVault2StateOk

`func (o *ComputerPartition) GetFileVault2StateOk() (*ComputerPartitionFileVault2State, bool)`

GetFileVault2StateOk returns a tuple with the FileVault2State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileVault2State

`func (o *ComputerPartition) SetFileVault2State(v ComputerPartitionFileVault2State)`

SetFileVault2State sets FileVault2State field to given value.

### HasFileVault2State

`func (o *ComputerPartition) HasFileVault2State() bool`

HasFileVault2State returns a boolean if a field has been set.

### GetFileVault2ProgressPercent

`func (o *ComputerPartition) GetFileVault2ProgressPercent() int64`

GetFileVault2ProgressPercent returns the FileVault2ProgressPercent field if non-nil, zero value otherwise.

### GetFileVault2ProgressPercentOk

`func (o *ComputerPartition) GetFileVault2ProgressPercentOk() (*int64, bool)`

GetFileVault2ProgressPercentOk returns a tuple with the FileVault2ProgressPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileVault2ProgressPercent

`func (o *ComputerPartition) SetFileVault2ProgressPercent(v int64)`

SetFileVault2ProgressPercent sets FileVault2ProgressPercent field to given value.

### HasFileVault2ProgressPercent

`func (o *ComputerPartition) HasFileVault2ProgressPercent() bool`

HasFileVault2ProgressPercent returns a boolean if a field has been set.

### SetFileVault2ProgressPercentNil

`func (o *ComputerPartition) SetFileVault2ProgressPercentNil(b bool)`

 SetFileVault2ProgressPercentNil sets the value for FileVault2ProgressPercent to be an explicit nil

### UnsetFileVault2ProgressPercent
`func (o *ComputerPartition) UnsetFileVault2ProgressPercent()`

UnsetFileVault2ProgressPercent ensures that no value is present for FileVault2ProgressPercent, not even an explicit nil
### GetLvmManaged

`func (o *ComputerPartition) GetLvmManaged() bool`

GetLvmManaged returns the LvmManaged field if non-nil, zero value otherwise.

### GetLvmManagedOk

`func (o *ComputerPartition) GetLvmManagedOk() (*bool, bool)`

GetLvmManagedOk returns a tuple with the LvmManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmManaged

`func (o *ComputerPartition) SetLvmManaged(v bool)`

SetLvmManaged sets LvmManaged field to given value.

### HasLvmManaged

`func (o *ComputerPartition) HasLvmManaged() bool`

HasLvmManaged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


