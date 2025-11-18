# ComputerStorageCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disks** | Pointer to [**[]ComputerDiskCreate**](ComputerDiskCreate.md) |  | [optional] 

## Methods

### NewComputerStorageCreate

`func NewComputerStorageCreate() *ComputerStorageCreate`

NewComputerStorageCreate instantiates a new ComputerStorageCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerStorageCreateWithDefaults

`func NewComputerStorageCreateWithDefaults() *ComputerStorageCreate`

NewComputerStorageCreateWithDefaults instantiates a new ComputerStorageCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisks

`func (o *ComputerStorageCreate) GetDisks() []ComputerDiskCreate`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *ComputerStorageCreate) GetDisksOk() (*[]ComputerDiskCreate, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *ComputerStorageCreate) SetDisks(v []ComputerDiskCreate)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *ComputerStorageCreate) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### SetDisksNil

`func (o *ComputerStorageCreate) SetDisksNil(b bool)`

 SetDisksNil sets the value for Disks to be an explicit nil

### UnsetDisks
`func (o *ComputerStorageCreate) UnsetDisks()`

UnsetDisks ensures that no value is present for Disks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


