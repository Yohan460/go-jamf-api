# ComputerStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootDriveAvailableSpaceMegabytes** | Pointer to **int64** |  | [optional] [readonly] 
**Disks** | Pointer to [**[]ComputerDisk**](ComputerDisk.md) |  | [optional] 

## Methods

### NewComputerStorage

`func NewComputerStorage() *ComputerStorage`

NewComputerStorage instantiates a new ComputerStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerStorageWithDefaults

`func NewComputerStorageWithDefaults() *ComputerStorage`

NewComputerStorageWithDefaults instantiates a new ComputerStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootDriveAvailableSpaceMegabytes

`func (o *ComputerStorage) GetBootDriveAvailableSpaceMegabytes() int64`

GetBootDriveAvailableSpaceMegabytes returns the BootDriveAvailableSpaceMegabytes field if non-nil, zero value otherwise.

### GetBootDriveAvailableSpaceMegabytesOk

`func (o *ComputerStorage) GetBootDriveAvailableSpaceMegabytesOk() (*int64, bool)`

GetBootDriveAvailableSpaceMegabytesOk returns a tuple with the BootDriveAvailableSpaceMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDriveAvailableSpaceMegabytes

`func (o *ComputerStorage) SetBootDriveAvailableSpaceMegabytes(v int64)`

SetBootDriveAvailableSpaceMegabytes sets BootDriveAvailableSpaceMegabytes field to given value.

### HasBootDriveAvailableSpaceMegabytes

`func (o *ComputerStorage) HasBootDriveAvailableSpaceMegabytes() bool`

HasBootDriveAvailableSpaceMegabytes returns a boolean if a field has been set.

### GetDisks

`func (o *ComputerStorage) GetDisks() []ComputerDisk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *ComputerStorage) GetDisksOk() (*[]ComputerDisk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *ComputerStorage) SetDisks(v []ComputerDisk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *ComputerStorage) HasDisks() bool`

HasDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


