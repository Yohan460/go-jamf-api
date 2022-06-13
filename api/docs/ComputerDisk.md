# ComputerDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Device** | Pointer to **string** |  | [optional] [readonly] 
**Model** | Pointer to **string** |  | [optional] [readonly] 
**Revision** | Pointer to **string** |  | [optional] [readonly] 
**SerialNumber** | Pointer to **string** |  | [optional] [readonly] 
**SizeMegabytes** | Pointer to **int64** | Disk Size in MB. | [optional] [readonly] 
**SmartStatus** | Pointer to **string** | S.M.A.R.T Status | [optional] [readonly] 
**Type** | Pointer to **string** | Connection type attribute. | [optional] [readonly] 
**Partitions** | Pointer to [**[]ComputerPartition**](ComputerPartition.md) |  | [optional] 

## Methods

### NewComputerDisk

`func NewComputerDisk() *ComputerDisk`

NewComputerDisk instantiates a new ComputerDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerDiskWithDefaults

`func NewComputerDiskWithDefaults() *ComputerDisk`

NewComputerDiskWithDefaults instantiates a new ComputerDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComputerDisk) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputerDisk) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputerDisk) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComputerDisk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDevice

`func (o *ComputerDisk) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ComputerDisk) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ComputerDisk) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *ComputerDisk) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetModel

`func (o *ComputerDisk) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ComputerDisk) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ComputerDisk) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ComputerDisk) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRevision

`func (o *ComputerDisk) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ComputerDisk) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ComputerDisk) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ComputerDisk) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ComputerDisk) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ComputerDisk) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ComputerDisk) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ComputerDisk) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSizeMegabytes

`func (o *ComputerDisk) GetSizeMegabytes() int64`

GetSizeMegabytes returns the SizeMegabytes field if non-nil, zero value otherwise.

### GetSizeMegabytesOk

`func (o *ComputerDisk) GetSizeMegabytesOk() (*int64, bool)`

GetSizeMegabytesOk returns a tuple with the SizeMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMegabytes

`func (o *ComputerDisk) SetSizeMegabytes(v int64)`

SetSizeMegabytes sets SizeMegabytes field to given value.

### HasSizeMegabytes

`func (o *ComputerDisk) HasSizeMegabytes() bool`

HasSizeMegabytes returns a boolean if a field has been set.

### GetSmartStatus

`func (o *ComputerDisk) GetSmartStatus() string`

GetSmartStatus returns the SmartStatus field if non-nil, zero value otherwise.

### GetSmartStatusOk

`func (o *ComputerDisk) GetSmartStatusOk() (*string, bool)`

GetSmartStatusOk returns a tuple with the SmartStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartStatus

`func (o *ComputerDisk) SetSmartStatus(v string)`

SetSmartStatus sets SmartStatus field to given value.

### HasSmartStatus

`func (o *ComputerDisk) HasSmartStatus() bool`

HasSmartStatus returns a boolean if a field has been set.

### GetType

`func (o *ComputerDisk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComputerDisk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComputerDisk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ComputerDisk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPartitions

`func (o *ComputerDisk) GetPartitions() []ComputerPartition`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *ComputerDisk) GetPartitionsOk() (*[]ComputerPartition, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *ComputerDisk) SetPartitions(v []ComputerPartition)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *ComputerDisk) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


