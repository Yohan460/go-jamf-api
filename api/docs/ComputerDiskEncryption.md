# ComputerDiskEncryption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootPartitionEncryptionDetails** | Pointer to [**ComputerPartitionEncryption**](ComputerPartitionEncryption.md) |  | [optional] 
**IndividualRecoveryKeyValidityStatus** | Pointer to **string** |  | [optional] 
**InstitutionalRecoveryKeyPresent** | Pointer to **bool** |  | [optional] 
**DiskEncryptionConfigurationName** | Pointer to **string** |  | [optional] 
**FileVault2EnabledUserNames** | Pointer to **[]string** |  | [optional] 
**FileVault2EligibilityMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewComputerDiskEncryption

`func NewComputerDiskEncryption() *ComputerDiskEncryption`

NewComputerDiskEncryption instantiates a new ComputerDiskEncryption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerDiskEncryptionWithDefaults

`func NewComputerDiskEncryptionWithDefaults() *ComputerDiskEncryption`

NewComputerDiskEncryptionWithDefaults instantiates a new ComputerDiskEncryption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootPartitionEncryptionDetails

`func (o *ComputerDiskEncryption) GetBootPartitionEncryptionDetails() ComputerPartitionEncryption`

GetBootPartitionEncryptionDetails returns the BootPartitionEncryptionDetails field if non-nil, zero value otherwise.

### GetBootPartitionEncryptionDetailsOk

`func (o *ComputerDiskEncryption) GetBootPartitionEncryptionDetailsOk() (*ComputerPartitionEncryption, bool)`

GetBootPartitionEncryptionDetailsOk returns a tuple with the BootPartitionEncryptionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootPartitionEncryptionDetails

`func (o *ComputerDiskEncryption) SetBootPartitionEncryptionDetails(v ComputerPartitionEncryption)`

SetBootPartitionEncryptionDetails sets BootPartitionEncryptionDetails field to given value.

### HasBootPartitionEncryptionDetails

`func (o *ComputerDiskEncryption) HasBootPartitionEncryptionDetails() bool`

HasBootPartitionEncryptionDetails returns a boolean if a field has been set.

### GetIndividualRecoveryKeyValidityStatus

`func (o *ComputerDiskEncryption) GetIndividualRecoveryKeyValidityStatus() string`

GetIndividualRecoveryKeyValidityStatus returns the IndividualRecoveryKeyValidityStatus field if non-nil, zero value otherwise.

### GetIndividualRecoveryKeyValidityStatusOk

`func (o *ComputerDiskEncryption) GetIndividualRecoveryKeyValidityStatusOk() (*string, bool)`

GetIndividualRecoveryKeyValidityStatusOk returns a tuple with the IndividualRecoveryKeyValidityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualRecoveryKeyValidityStatus

`func (o *ComputerDiskEncryption) SetIndividualRecoveryKeyValidityStatus(v string)`

SetIndividualRecoveryKeyValidityStatus sets IndividualRecoveryKeyValidityStatus field to given value.

### HasIndividualRecoveryKeyValidityStatus

`func (o *ComputerDiskEncryption) HasIndividualRecoveryKeyValidityStatus() bool`

HasIndividualRecoveryKeyValidityStatus returns a boolean if a field has been set.

### GetInstitutionalRecoveryKeyPresent

`func (o *ComputerDiskEncryption) GetInstitutionalRecoveryKeyPresent() bool`

GetInstitutionalRecoveryKeyPresent returns the InstitutionalRecoveryKeyPresent field if non-nil, zero value otherwise.

### GetInstitutionalRecoveryKeyPresentOk

`func (o *ComputerDiskEncryption) GetInstitutionalRecoveryKeyPresentOk() (*bool, bool)`

GetInstitutionalRecoveryKeyPresentOk returns a tuple with the InstitutionalRecoveryKeyPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionalRecoveryKeyPresent

`func (o *ComputerDiskEncryption) SetInstitutionalRecoveryKeyPresent(v bool)`

SetInstitutionalRecoveryKeyPresent sets InstitutionalRecoveryKeyPresent field to given value.

### HasInstitutionalRecoveryKeyPresent

`func (o *ComputerDiskEncryption) HasInstitutionalRecoveryKeyPresent() bool`

HasInstitutionalRecoveryKeyPresent returns a boolean if a field has been set.

### GetDiskEncryptionConfigurationName

`func (o *ComputerDiskEncryption) GetDiskEncryptionConfigurationName() string`

GetDiskEncryptionConfigurationName returns the DiskEncryptionConfigurationName field if non-nil, zero value otherwise.

### GetDiskEncryptionConfigurationNameOk

`func (o *ComputerDiskEncryption) GetDiskEncryptionConfigurationNameOk() (*string, bool)`

GetDiskEncryptionConfigurationNameOk returns a tuple with the DiskEncryptionConfigurationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryptionConfigurationName

`func (o *ComputerDiskEncryption) SetDiskEncryptionConfigurationName(v string)`

SetDiskEncryptionConfigurationName sets DiskEncryptionConfigurationName field to given value.

### HasDiskEncryptionConfigurationName

`func (o *ComputerDiskEncryption) HasDiskEncryptionConfigurationName() bool`

HasDiskEncryptionConfigurationName returns a boolean if a field has been set.

### GetFileVault2EnabledUserNames

`func (o *ComputerDiskEncryption) GetFileVault2EnabledUserNames() []string`

GetFileVault2EnabledUserNames returns the FileVault2EnabledUserNames field if non-nil, zero value otherwise.

### GetFileVault2EnabledUserNamesOk

`func (o *ComputerDiskEncryption) GetFileVault2EnabledUserNamesOk() (*[]string, bool)`

GetFileVault2EnabledUserNamesOk returns a tuple with the FileVault2EnabledUserNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileVault2EnabledUserNames

`func (o *ComputerDiskEncryption) SetFileVault2EnabledUserNames(v []string)`

SetFileVault2EnabledUserNames sets FileVault2EnabledUserNames field to given value.

### HasFileVault2EnabledUserNames

`func (o *ComputerDiskEncryption) HasFileVault2EnabledUserNames() bool`

HasFileVault2EnabledUserNames returns a boolean if a field has been set.

### GetFileVault2EligibilityMessage

`func (o *ComputerDiskEncryption) GetFileVault2EligibilityMessage() string`

GetFileVault2EligibilityMessage returns the FileVault2EligibilityMessage field if non-nil, zero value otherwise.

### GetFileVault2EligibilityMessageOk

`func (o *ComputerDiskEncryption) GetFileVault2EligibilityMessageOk() (*string, bool)`

GetFileVault2EligibilityMessageOk returns a tuple with the FileVault2EligibilityMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileVault2EligibilityMessage

`func (o *ComputerDiskEncryption) SetFileVault2EligibilityMessage(v string)`

SetFileVault2EligibilityMessage sets FileVault2EligibilityMessage field to given value.

### HasFileVault2EligibilityMessage

`func (o *ComputerDiskEncryption) HasFileVault2EligibilityMessage() bool`

HasFileVault2EligibilityMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


