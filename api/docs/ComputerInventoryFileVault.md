# ComputerInventoryFileVault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputerId** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**PersonalRecoveryKey** | Pointer to **string** |  | [optional] 
**BootPartitionEncryptionDetails** | Pointer to [**ComputerPartitionEncryption**](ComputerPartitionEncryption.md) |  | [optional] 
**IndividualRecoveryKeyValidityStatus** | Pointer to **string** |  | [optional] 
**InstitutionalRecoveryKeyPresent** | Pointer to **bool** |  | [optional] 
**DiskEncryptionConfigurationName** | Pointer to **string** |  | [optional] 

## Methods

### NewComputerInventoryFileVault

`func NewComputerInventoryFileVault() *ComputerInventoryFileVault`

NewComputerInventoryFileVault instantiates a new ComputerInventoryFileVault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerInventoryFileVaultWithDefaults

`func NewComputerInventoryFileVaultWithDefaults() *ComputerInventoryFileVault`

NewComputerInventoryFileVaultWithDefaults instantiates a new ComputerInventoryFileVault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputerId

`func (o *ComputerInventoryFileVault) GetComputerId() string`

GetComputerId returns the ComputerId field if non-nil, zero value otherwise.

### GetComputerIdOk

`func (o *ComputerInventoryFileVault) GetComputerIdOk() (*string, bool)`

GetComputerIdOk returns a tuple with the ComputerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerId

`func (o *ComputerInventoryFileVault) SetComputerId(v string)`

SetComputerId sets ComputerId field to given value.

### HasComputerId

`func (o *ComputerInventoryFileVault) HasComputerId() bool`

HasComputerId returns a boolean if a field has been set.

### GetName

`func (o *ComputerInventoryFileVault) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputerInventoryFileVault) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputerInventoryFileVault) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputerInventoryFileVault) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPersonalRecoveryKey

`func (o *ComputerInventoryFileVault) GetPersonalRecoveryKey() string`

GetPersonalRecoveryKey returns the PersonalRecoveryKey field if non-nil, zero value otherwise.

### GetPersonalRecoveryKeyOk

`func (o *ComputerInventoryFileVault) GetPersonalRecoveryKeyOk() (*string, bool)`

GetPersonalRecoveryKeyOk returns a tuple with the PersonalRecoveryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalRecoveryKey

`func (o *ComputerInventoryFileVault) SetPersonalRecoveryKey(v string)`

SetPersonalRecoveryKey sets PersonalRecoveryKey field to given value.

### HasPersonalRecoveryKey

`func (o *ComputerInventoryFileVault) HasPersonalRecoveryKey() bool`

HasPersonalRecoveryKey returns a boolean if a field has been set.

### GetBootPartitionEncryptionDetails

`func (o *ComputerInventoryFileVault) GetBootPartitionEncryptionDetails() ComputerPartitionEncryption`

GetBootPartitionEncryptionDetails returns the BootPartitionEncryptionDetails field if non-nil, zero value otherwise.

### GetBootPartitionEncryptionDetailsOk

`func (o *ComputerInventoryFileVault) GetBootPartitionEncryptionDetailsOk() (*ComputerPartitionEncryption, bool)`

GetBootPartitionEncryptionDetailsOk returns a tuple with the BootPartitionEncryptionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootPartitionEncryptionDetails

`func (o *ComputerInventoryFileVault) SetBootPartitionEncryptionDetails(v ComputerPartitionEncryption)`

SetBootPartitionEncryptionDetails sets BootPartitionEncryptionDetails field to given value.

### HasBootPartitionEncryptionDetails

`func (o *ComputerInventoryFileVault) HasBootPartitionEncryptionDetails() bool`

HasBootPartitionEncryptionDetails returns a boolean if a field has been set.

### GetIndividualRecoveryKeyValidityStatus

`func (o *ComputerInventoryFileVault) GetIndividualRecoveryKeyValidityStatus() string`

GetIndividualRecoveryKeyValidityStatus returns the IndividualRecoveryKeyValidityStatus field if non-nil, zero value otherwise.

### GetIndividualRecoveryKeyValidityStatusOk

`func (o *ComputerInventoryFileVault) GetIndividualRecoveryKeyValidityStatusOk() (*string, bool)`

GetIndividualRecoveryKeyValidityStatusOk returns a tuple with the IndividualRecoveryKeyValidityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualRecoveryKeyValidityStatus

`func (o *ComputerInventoryFileVault) SetIndividualRecoveryKeyValidityStatus(v string)`

SetIndividualRecoveryKeyValidityStatus sets IndividualRecoveryKeyValidityStatus field to given value.

### HasIndividualRecoveryKeyValidityStatus

`func (o *ComputerInventoryFileVault) HasIndividualRecoveryKeyValidityStatus() bool`

HasIndividualRecoveryKeyValidityStatus returns a boolean if a field has been set.

### GetInstitutionalRecoveryKeyPresent

`func (o *ComputerInventoryFileVault) GetInstitutionalRecoveryKeyPresent() bool`

GetInstitutionalRecoveryKeyPresent returns the InstitutionalRecoveryKeyPresent field if non-nil, zero value otherwise.

### GetInstitutionalRecoveryKeyPresentOk

`func (o *ComputerInventoryFileVault) GetInstitutionalRecoveryKeyPresentOk() (*bool, bool)`

GetInstitutionalRecoveryKeyPresentOk returns a tuple with the InstitutionalRecoveryKeyPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionalRecoveryKeyPresent

`func (o *ComputerInventoryFileVault) SetInstitutionalRecoveryKeyPresent(v bool)`

SetInstitutionalRecoveryKeyPresent sets InstitutionalRecoveryKeyPresent field to given value.

### HasInstitutionalRecoveryKeyPresent

`func (o *ComputerInventoryFileVault) HasInstitutionalRecoveryKeyPresent() bool`

HasInstitutionalRecoveryKeyPresent returns a boolean if a field has been set.

### GetDiskEncryptionConfigurationName

`func (o *ComputerInventoryFileVault) GetDiskEncryptionConfigurationName() string`

GetDiskEncryptionConfigurationName returns the DiskEncryptionConfigurationName field if non-nil, zero value otherwise.

### GetDiskEncryptionConfigurationNameOk

`func (o *ComputerInventoryFileVault) GetDiskEncryptionConfigurationNameOk() (*string, bool)`

GetDiskEncryptionConfigurationNameOk returns a tuple with the DiskEncryptionConfigurationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryptionConfigurationName

`func (o *ComputerInventoryFileVault) SetDiskEncryptionConfigurationName(v string)`

SetDiskEncryptionConfigurationName sets DiskEncryptionConfigurationName field to given value.

### HasDiskEncryptionConfigurationName

`func (o *ComputerInventoryFileVault) HasDiskEncryptionConfigurationName() bool`

HasDiskEncryptionConfigurationName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


