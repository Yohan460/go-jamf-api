# ComputerDiskCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to **NullableString** |  | [optional] 
**Model** | Pointer to **NullableString** |  | [optional] 
**Revision** | Pointer to **NullableString** |  | [optional] 
**SerialNumber** | Pointer to **NullableString** |  | [optional] 
**SizeMegabytes** | Pointer to **NullableInt64** | Disk Size in MB. | [optional] 
**SmartStatus** | Pointer to **NullableString** | S.M.A.R.T Status | [optional] 
**Type** | Pointer to **NullableString** | Connection type attribute. | [optional] 
**Partitions** | Pointer to [**[]ComputerPartitionCreate**](ComputerPartitionCreate.md) |  | [optional] 

## Methods

### NewComputerDiskCreate

`func NewComputerDiskCreate() *ComputerDiskCreate`

NewComputerDiskCreate instantiates a new ComputerDiskCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerDiskCreateWithDefaults

`func NewComputerDiskCreateWithDefaults() *ComputerDiskCreate`

NewComputerDiskCreateWithDefaults instantiates a new ComputerDiskCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *ComputerDiskCreate) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ComputerDiskCreate) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ComputerDiskCreate) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *ComputerDiskCreate) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *ComputerDiskCreate) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *ComputerDiskCreate) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetModel

`func (o *ComputerDiskCreate) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ComputerDiskCreate) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ComputerDiskCreate) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ComputerDiskCreate) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *ComputerDiskCreate) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *ComputerDiskCreate) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetRevision

`func (o *ComputerDiskCreate) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ComputerDiskCreate) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ComputerDiskCreate) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ComputerDiskCreate) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### SetRevisionNil

`func (o *ComputerDiskCreate) SetRevisionNil(b bool)`

 SetRevisionNil sets the value for Revision to be an explicit nil

### UnsetRevision
`func (o *ComputerDiskCreate) UnsetRevision()`

UnsetRevision ensures that no value is present for Revision, not even an explicit nil
### GetSerialNumber

`func (o *ComputerDiskCreate) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ComputerDiskCreate) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ComputerDiskCreate) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ComputerDiskCreate) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *ComputerDiskCreate) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *ComputerDiskCreate) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetSizeMegabytes

`func (o *ComputerDiskCreate) GetSizeMegabytes() int64`

GetSizeMegabytes returns the SizeMegabytes field if non-nil, zero value otherwise.

### GetSizeMegabytesOk

`func (o *ComputerDiskCreate) GetSizeMegabytesOk() (*int64, bool)`

GetSizeMegabytesOk returns a tuple with the SizeMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMegabytes

`func (o *ComputerDiskCreate) SetSizeMegabytes(v int64)`

SetSizeMegabytes sets SizeMegabytes field to given value.

### HasSizeMegabytes

`func (o *ComputerDiskCreate) HasSizeMegabytes() bool`

HasSizeMegabytes returns a boolean if a field has been set.

### SetSizeMegabytesNil

`func (o *ComputerDiskCreate) SetSizeMegabytesNil(b bool)`

 SetSizeMegabytesNil sets the value for SizeMegabytes to be an explicit nil

### UnsetSizeMegabytes
`func (o *ComputerDiskCreate) UnsetSizeMegabytes()`

UnsetSizeMegabytes ensures that no value is present for SizeMegabytes, not even an explicit nil
### GetSmartStatus

`func (o *ComputerDiskCreate) GetSmartStatus() string`

GetSmartStatus returns the SmartStatus field if non-nil, zero value otherwise.

### GetSmartStatusOk

`func (o *ComputerDiskCreate) GetSmartStatusOk() (*string, bool)`

GetSmartStatusOk returns a tuple with the SmartStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartStatus

`func (o *ComputerDiskCreate) SetSmartStatus(v string)`

SetSmartStatus sets SmartStatus field to given value.

### HasSmartStatus

`func (o *ComputerDiskCreate) HasSmartStatus() bool`

HasSmartStatus returns a boolean if a field has been set.

### SetSmartStatusNil

`func (o *ComputerDiskCreate) SetSmartStatusNil(b bool)`

 SetSmartStatusNil sets the value for SmartStatus to be an explicit nil

### UnsetSmartStatus
`func (o *ComputerDiskCreate) UnsetSmartStatus()`

UnsetSmartStatus ensures that no value is present for SmartStatus, not even an explicit nil
### GetType

`func (o *ComputerDiskCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComputerDiskCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComputerDiskCreate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ComputerDiskCreate) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ComputerDiskCreate) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ComputerDiskCreate) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetPartitions

`func (o *ComputerDiskCreate) GetPartitions() []ComputerPartitionCreate`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *ComputerDiskCreate) GetPartitionsOk() (*[]ComputerPartitionCreate, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *ComputerDiskCreate) SetPartitions(v []ComputerPartitionCreate)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *ComputerDiskCreate) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### SetPartitionsNil

`func (o *ComputerDiskCreate) SetPartitionsNil(b bool)`

 SetPartitionsNil sets the value for Partitions to be an explicit nil

### UnsetPartitions
`func (o *ComputerDiskCreate) UnsetPartitions()`

UnsetPartitions ensures that no value is present for Partitions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


