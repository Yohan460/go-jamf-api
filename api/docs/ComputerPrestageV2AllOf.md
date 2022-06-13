# ComputerPrestageV2AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallProfilesDuringSetup** | **bool** |  | 
**PrestageInstalledProfileIds** | **[]string** |  | 
**CustomPackageIds** | **[]string** |  | 
**CustomPackageDistributionPointId** | **string** |  | 
**EnableRecoveryLock** | Pointer to **bool** |  | [optional] 
**RecoveryLockPasswordType** | Pointer to **string** |  | [optional] 
**RotateRecoveryLockPassword** | Pointer to **bool** |  | [optional] 

## Methods

### NewComputerPrestageV2AllOf

`func NewComputerPrestageV2AllOf(installProfilesDuringSetup bool, prestageInstalledProfileIds []string, customPackageIds []string, customPackageDistributionPointId string, ) *ComputerPrestageV2AllOf`

NewComputerPrestageV2AllOf instantiates a new ComputerPrestageV2AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerPrestageV2AllOfWithDefaults

`func NewComputerPrestageV2AllOfWithDefaults() *ComputerPrestageV2AllOf`

NewComputerPrestageV2AllOfWithDefaults instantiates a new ComputerPrestageV2AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallProfilesDuringSetup

`func (o *ComputerPrestageV2AllOf) GetInstallProfilesDuringSetup() bool`

GetInstallProfilesDuringSetup returns the InstallProfilesDuringSetup field if non-nil, zero value otherwise.

### GetInstallProfilesDuringSetupOk

`func (o *ComputerPrestageV2AllOf) GetInstallProfilesDuringSetupOk() (*bool, bool)`

GetInstallProfilesDuringSetupOk returns a tuple with the InstallProfilesDuringSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallProfilesDuringSetup

`func (o *ComputerPrestageV2AllOf) SetInstallProfilesDuringSetup(v bool)`

SetInstallProfilesDuringSetup sets InstallProfilesDuringSetup field to given value.


### GetPrestageInstalledProfileIds

`func (o *ComputerPrestageV2AllOf) GetPrestageInstalledProfileIds() []string`

GetPrestageInstalledProfileIds returns the PrestageInstalledProfileIds field if non-nil, zero value otherwise.

### GetPrestageInstalledProfileIdsOk

`func (o *ComputerPrestageV2AllOf) GetPrestageInstalledProfileIdsOk() (*[]string, bool)`

GetPrestageInstalledProfileIdsOk returns a tuple with the PrestageInstalledProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrestageInstalledProfileIds

`func (o *ComputerPrestageV2AllOf) SetPrestageInstalledProfileIds(v []string)`

SetPrestageInstalledProfileIds sets PrestageInstalledProfileIds field to given value.


### GetCustomPackageIds

`func (o *ComputerPrestageV2AllOf) GetCustomPackageIds() []string`

GetCustomPackageIds returns the CustomPackageIds field if non-nil, zero value otherwise.

### GetCustomPackageIdsOk

`func (o *ComputerPrestageV2AllOf) GetCustomPackageIdsOk() (*[]string, bool)`

GetCustomPackageIdsOk returns a tuple with the CustomPackageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPackageIds

`func (o *ComputerPrestageV2AllOf) SetCustomPackageIds(v []string)`

SetCustomPackageIds sets CustomPackageIds field to given value.


### GetCustomPackageDistributionPointId

`func (o *ComputerPrestageV2AllOf) GetCustomPackageDistributionPointId() string`

GetCustomPackageDistributionPointId returns the CustomPackageDistributionPointId field if non-nil, zero value otherwise.

### GetCustomPackageDistributionPointIdOk

`func (o *ComputerPrestageV2AllOf) GetCustomPackageDistributionPointIdOk() (*string, bool)`

GetCustomPackageDistributionPointIdOk returns a tuple with the CustomPackageDistributionPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPackageDistributionPointId

`func (o *ComputerPrestageV2AllOf) SetCustomPackageDistributionPointId(v string)`

SetCustomPackageDistributionPointId sets CustomPackageDistributionPointId field to given value.


### GetEnableRecoveryLock

`func (o *ComputerPrestageV2AllOf) GetEnableRecoveryLock() bool`

GetEnableRecoveryLock returns the EnableRecoveryLock field if non-nil, zero value otherwise.

### GetEnableRecoveryLockOk

`func (o *ComputerPrestageV2AllOf) GetEnableRecoveryLockOk() (*bool, bool)`

GetEnableRecoveryLockOk returns a tuple with the EnableRecoveryLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRecoveryLock

`func (o *ComputerPrestageV2AllOf) SetEnableRecoveryLock(v bool)`

SetEnableRecoveryLock sets EnableRecoveryLock field to given value.

### HasEnableRecoveryLock

`func (o *ComputerPrestageV2AllOf) HasEnableRecoveryLock() bool`

HasEnableRecoveryLock returns a boolean if a field has been set.

### GetRecoveryLockPasswordType

`func (o *ComputerPrestageV2AllOf) GetRecoveryLockPasswordType() string`

GetRecoveryLockPasswordType returns the RecoveryLockPasswordType field if non-nil, zero value otherwise.

### GetRecoveryLockPasswordTypeOk

`func (o *ComputerPrestageV2AllOf) GetRecoveryLockPasswordTypeOk() (*string, bool)`

GetRecoveryLockPasswordTypeOk returns a tuple with the RecoveryLockPasswordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryLockPasswordType

`func (o *ComputerPrestageV2AllOf) SetRecoveryLockPasswordType(v string)`

SetRecoveryLockPasswordType sets RecoveryLockPasswordType field to given value.

### HasRecoveryLockPasswordType

`func (o *ComputerPrestageV2AllOf) HasRecoveryLockPasswordType() bool`

HasRecoveryLockPasswordType returns a boolean if a field has been set.

### GetRotateRecoveryLockPassword

`func (o *ComputerPrestageV2AllOf) GetRotateRecoveryLockPassword() bool`

GetRotateRecoveryLockPassword returns the RotateRecoveryLockPassword field if non-nil, zero value otherwise.

### GetRotateRecoveryLockPasswordOk

`func (o *ComputerPrestageV2AllOf) GetRotateRecoveryLockPasswordOk() (*bool, bool)`

GetRotateRecoveryLockPasswordOk returns a tuple with the RotateRecoveryLockPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateRecoveryLockPassword

`func (o *ComputerPrestageV2AllOf) SetRotateRecoveryLockPassword(v bool)`

SetRotateRecoveryLockPassword sets RotateRecoveryLockPassword field to given value.

### HasRotateRecoveryLockPassword

`func (o *ComputerPrestageV2AllOf) HasRotateRecoveryLockPassword() bool`

HasRotateRecoveryLockPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


