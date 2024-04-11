# SharedDeviceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuotaSize** | Pointer to **int64** |  | [optional] 
**ResidentUsers** | Pointer to **int64** |  | [optional] 

## Methods

### NewSharedDeviceConfiguration

`func NewSharedDeviceConfiguration() *SharedDeviceConfiguration`

NewSharedDeviceConfiguration instantiates a new SharedDeviceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedDeviceConfigurationWithDefaults

`func NewSharedDeviceConfigurationWithDefaults() *SharedDeviceConfiguration`

NewSharedDeviceConfigurationWithDefaults instantiates a new SharedDeviceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuotaSize

`func (o *SharedDeviceConfiguration) GetQuotaSize() int64`

GetQuotaSize returns the QuotaSize field if non-nil, zero value otherwise.

### GetQuotaSizeOk

`func (o *SharedDeviceConfiguration) GetQuotaSizeOk() (*int64, bool)`

GetQuotaSizeOk returns a tuple with the QuotaSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaSize

`func (o *SharedDeviceConfiguration) SetQuotaSize(v int64)`

SetQuotaSize sets QuotaSize field to given value.

### HasQuotaSize

`func (o *SharedDeviceConfiguration) HasQuotaSize() bool`

HasQuotaSize returns a boolean if a field has been set.

### GetResidentUsers

`func (o *SharedDeviceConfiguration) GetResidentUsers() int64`

GetResidentUsers returns the ResidentUsers field if non-nil, zero value otherwise.

### GetResidentUsersOk

`func (o *SharedDeviceConfiguration) GetResidentUsersOk() (*int64, bool)`

GetResidentUsersOk returns a tuple with the ResidentUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentUsers

`func (o *SharedDeviceConfiguration) SetResidentUsers(v int64)`

SetResidentUsers sets ResidentUsers field to given value.

### HasResidentUsers

`func (o *SharedDeviceConfiguration) HasResidentUsers() bool`

HasResidentUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


