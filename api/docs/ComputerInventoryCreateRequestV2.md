# ComputerInventoryCreateRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Udid** | Pointer to **NullableString** |  | [optional] 
**General** | Pointer to [**ComputerGeneralCreate**](ComputerGeneralCreate.md) |  | [optional] 
**Applications** | Pointer to [**[]ComputerApplicationCreate**](ComputerApplicationCreate.md) |  | [optional] 
**Storage** | Pointer to [**ComputerStorageCreate**](ComputerStorageCreate.md) |  | [optional] 
**Security** | Pointer to [**ComputerSecurityCreate**](ComputerSecurityCreate.md) |  | [optional] 
**ConfigurationProfiles** | Pointer to [**[]ComputerConfigurationProfileCreate**](ComputerConfigurationProfileCreate.md) |  | [optional] 
**Printers** | Pointer to [**[]ComputerPrinterCreate**](ComputerPrinterCreate.md) |  | [optional] 
**Services** | Pointer to [**[]ComputerServiceCreate**](ComputerServiceCreate.md) |  | [optional] 
**LocalUserAccounts** | Pointer to [**[]ComputerLocalUserAccountCreate**](ComputerLocalUserAccountCreate.md) |  | [optional] 
**Certificates** | Pointer to [**[]ComputerCertificateCreate**](ComputerCertificateCreate.md) |  | [optional] 
**PackageReceipts** | Pointer to [**ComputerPackageReceiptsCreate**](ComputerPackageReceiptsCreate.md) |  | [optional] 
**SoftwareUpdates** | Pointer to [**[]ComputerSoftwareUpdateCreate**](ComputerSoftwareUpdateCreate.md) |  | [optional] 
**Purchasing** | Pointer to [**ComputerPurchaseCreate**](ComputerPurchaseCreate.md) |  | [optional] 
**UserAndLocation** | Pointer to [**ComputerUserAndLocationCreate**](ComputerUserAndLocationCreate.md) |  | [optional] 
**Hardware** | Pointer to [**ComputerHardwareCreate**](ComputerHardwareCreate.md) |  | [optional] 
**OperatingSystem** | Pointer to [**ComputerOperatingSystemCreate**](ComputerOperatingSystemCreate.md) |  | [optional] 

## Methods

### NewComputerInventoryCreateRequestV2

`func NewComputerInventoryCreateRequestV2() *ComputerInventoryCreateRequestV2`

NewComputerInventoryCreateRequestV2 instantiates a new ComputerInventoryCreateRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerInventoryCreateRequestV2WithDefaults

`func NewComputerInventoryCreateRequestV2WithDefaults() *ComputerInventoryCreateRequestV2`

NewComputerInventoryCreateRequestV2WithDefaults instantiates a new ComputerInventoryCreateRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUdid

`func (o *ComputerInventoryCreateRequestV2) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *ComputerInventoryCreateRequestV2) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *ComputerInventoryCreateRequestV2) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *ComputerInventoryCreateRequestV2) HasUdid() bool`

HasUdid returns a boolean if a field has been set.

### SetUdidNil

`func (o *ComputerInventoryCreateRequestV2) SetUdidNil(b bool)`

 SetUdidNil sets the value for Udid to be an explicit nil

### UnsetUdid
`func (o *ComputerInventoryCreateRequestV2) UnsetUdid()`

UnsetUdid ensures that no value is present for Udid, not even an explicit nil
### GetGeneral

`func (o *ComputerInventoryCreateRequestV2) GetGeneral() ComputerGeneralCreate`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *ComputerInventoryCreateRequestV2) GetGeneralOk() (*ComputerGeneralCreate, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *ComputerInventoryCreateRequestV2) SetGeneral(v ComputerGeneralCreate)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *ComputerInventoryCreateRequestV2) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetApplications

`func (o *ComputerInventoryCreateRequestV2) GetApplications() []ComputerApplicationCreate`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ComputerInventoryCreateRequestV2) GetApplicationsOk() (*[]ComputerApplicationCreate, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ComputerInventoryCreateRequestV2) SetApplications(v []ComputerApplicationCreate)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *ComputerInventoryCreateRequestV2) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### SetApplicationsNil

`func (o *ComputerInventoryCreateRequestV2) SetApplicationsNil(b bool)`

 SetApplicationsNil sets the value for Applications to be an explicit nil

### UnsetApplications
`func (o *ComputerInventoryCreateRequestV2) UnsetApplications()`

UnsetApplications ensures that no value is present for Applications, not even an explicit nil
### GetStorage

`func (o *ComputerInventoryCreateRequestV2) GetStorage() ComputerStorageCreate`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ComputerInventoryCreateRequestV2) GetStorageOk() (*ComputerStorageCreate, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ComputerInventoryCreateRequestV2) SetStorage(v ComputerStorageCreate)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ComputerInventoryCreateRequestV2) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetSecurity

`func (o *ComputerInventoryCreateRequestV2) GetSecurity() ComputerSecurityCreate`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *ComputerInventoryCreateRequestV2) GetSecurityOk() (*ComputerSecurityCreate, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *ComputerInventoryCreateRequestV2) SetSecurity(v ComputerSecurityCreate)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *ComputerInventoryCreateRequestV2) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetConfigurationProfiles

`func (o *ComputerInventoryCreateRequestV2) GetConfigurationProfiles() []ComputerConfigurationProfileCreate`

GetConfigurationProfiles returns the ConfigurationProfiles field if non-nil, zero value otherwise.

### GetConfigurationProfilesOk

`func (o *ComputerInventoryCreateRequestV2) GetConfigurationProfilesOk() (*[]ComputerConfigurationProfileCreate, bool)`

GetConfigurationProfilesOk returns a tuple with the ConfigurationProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationProfiles

`func (o *ComputerInventoryCreateRequestV2) SetConfigurationProfiles(v []ComputerConfigurationProfileCreate)`

SetConfigurationProfiles sets ConfigurationProfiles field to given value.

### HasConfigurationProfiles

`func (o *ComputerInventoryCreateRequestV2) HasConfigurationProfiles() bool`

HasConfigurationProfiles returns a boolean if a field has been set.

### SetConfigurationProfilesNil

`func (o *ComputerInventoryCreateRequestV2) SetConfigurationProfilesNil(b bool)`

 SetConfigurationProfilesNil sets the value for ConfigurationProfiles to be an explicit nil

### UnsetConfigurationProfiles
`func (o *ComputerInventoryCreateRequestV2) UnsetConfigurationProfiles()`

UnsetConfigurationProfiles ensures that no value is present for ConfigurationProfiles, not even an explicit nil
### GetPrinters

`func (o *ComputerInventoryCreateRequestV2) GetPrinters() []ComputerPrinterCreate`

GetPrinters returns the Printers field if non-nil, zero value otherwise.

### GetPrintersOk

`func (o *ComputerInventoryCreateRequestV2) GetPrintersOk() (*[]ComputerPrinterCreate, bool)`

GetPrintersOk returns a tuple with the Printers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinters

`func (o *ComputerInventoryCreateRequestV2) SetPrinters(v []ComputerPrinterCreate)`

SetPrinters sets Printers field to given value.

### HasPrinters

`func (o *ComputerInventoryCreateRequestV2) HasPrinters() bool`

HasPrinters returns a boolean if a field has been set.

### SetPrintersNil

`func (o *ComputerInventoryCreateRequestV2) SetPrintersNil(b bool)`

 SetPrintersNil sets the value for Printers to be an explicit nil

### UnsetPrinters
`func (o *ComputerInventoryCreateRequestV2) UnsetPrinters()`

UnsetPrinters ensures that no value is present for Printers, not even an explicit nil
### GetServices

`func (o *ComputerInventoryCreateRequestV2) GetServices() []ComputerServiceCreate`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ComputerInventoryCreateRequestV2) GetServicesOk() (*[]ComputerServiceCreate, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ComputerInventoryCreateRequestV2) SetServices(v []ComputerServiceCreate)`

SetServices sets Services field to given value.

### HasServices

`func (o *ComputerInventoryCreateRequestV2) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *ComputerInventoryCreateRequestV2) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *ComputerInventoryCreateRequestV2) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetLocalUserAccounts

`func (o *ComputerInventoryCreateRequestV2) GetLocalUserAccounts() []ComputerLocalUserAccountCreate`

GetLocalUserAccounts returns the LocalUserAccounts field if non-nil, zero value otherwise.

### GetLocalUserAccountsOk

`func (o *ComputerInventoryCreateRequestV2) GetLocalUserAccountsOk() (*[]ComputerLocalUserAccountCreate, bool)`

GetLocalUserAccountsOk returns a tuple with the LocalUserAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserAccounts

`func (o *ComputerInventoryCreateRequestV2) SetLocalUserAccounts(v []ComputerLocalUserAccountCreate)`

SetLocalUserAccounts sets LocalUserAccounts field to given value.

### HasLocalUserAccounts

`func (o *ComputerInventoryCreateRequestV2) HasLocalUserAccounts() bool`

HasLocalUserAccounts returns a boolean if a field has been set.

### SetLocalUserAccountsNil

`func (o *ComputerInventoryCreateRequestV2) SetLocalUserAccountsNil(b bool)`

 SetLocalUserAccountsNil sets the value for LocalUserAccounts to be an explicit nil

### UnsetLocalUserAccounts
`func (o *ComputerInventoryCreateRequestV2) UnsetLocalUserAccounts()`

UnsetLocalUserAccounts ensures that no value is present for LocalUserAccounts, not even an explicit nil
### GetCertificates

`func (o *ComputerInventoryCreateRequestV2) GetCertificates() []ComputerCertificateCreate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *ComputerInventoryCreateRequestV2) GetCertificatesOk() (*[]ComputerCertificateCreate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *ComputerInventoryCreateRequestV2) SetCertificates(v []ComputerCertificateCreate)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *ComputerInventoryCreateRequestV2) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### SetCertificatesNil

`func (o *ComputerInventoryCreateRequestV2) SetCertificatesNil(b bool)`

 SetCertificatesNil sets the value for Certificates to be an explicit nil

### UnsetCertificates
`func (o *ComputerInventoryCreateRequestV2) UnsetCertificates()`

UnsetCertificates ensures that no value is present for Certificates, not even an explicit nil
### GetPackageReceipts

`func (o *ComputerInventoryCreateRequestV2) GetPackageReceipts() ComputerPackageReceiptsCreate`

GetPackageReceipts returns the PackageReceipts field if non-nil, zero value otherwise.

### GetPackageReceiptsOk

`func (o *ComputerInventoryCreateRequestV2) GetPackageReceiptsOk() (*ComputerPackageReceiptsCreate, bool)`

GetPackageReceiptsOk returns a tuple with the PackageReceipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageReceipts

`func (o *ComputerInventoryCreateRequestV2) SetPackageReceipts(v ComputerPackageReceiptsCreate)`

SetPackageReceipts sets PackageReceipts field to given value.

### HasPackageReceipts

`func (o *ComputerInventoryCreateRequestV2) HasPackageReceipts() bool`

HasPackageReceipts returns a boolean if a field has been set.

### GetSoftwareUpdates

`func (o *ComputerInventoryCreateRequestV2) GetSoftwareUpdates() []ComputerSoftwareUpdateCreate`

GetSoftwareUpdates returns the SoftwareUpdates field if non-nil, zero value otherwise.

### GetSoftwareUpdatesOk

`func (o *ComputerInventoryCreateRequestV2) GetSoftwareUpdatesOk() (*[]ComputerSoftwareUpdateCreate, bool)`

GetSoftwareUpdatesOk returns a tuple with the SoftwareUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdates

`func (o *ComputerInventoryCreateRequestV2) SetSoftwareUpdates(v []ComputerSoftwareUpdateCreate)`

SetSoftwareUpdates sets SoftwareUpdates field to given value.

### HasSoftwareUpdates

`func (o *ComputerInventoryCreateRequestV2) HasSoftwareUpdates() bool`

HasSoftwareUpdates returns a boolean if a field has been set.

### SetSoftwareUpdatesNil

`func (o *ComputerInventoryCreateRequestV2) SetSoftwareUpdatesNil(b bool)`

 SetSoftwareUpdatesNil sets the value for SoftwareUpdates to be an explicit nil

### UnsetSoftwareUpdates
`func (o *ComputerInventoryCreateRequestV2) UnsetSoftwareUpdates()`

UnsetSoftwareUpdates ensures that no value is present for SoftwareUpdates, not even an explicit nil
### GetPurchasing

`func (o *ComputerInventoryCreateRequestV2) GetPurchasing() ComputerPurchaseCreate`

GetPurchasing returns the Purchasing field if non-nil, zero value otherwise.

### GetPurchasingOk

`func (o *ComputerInventoryCreateRequestV2) GetPurchasingOk() (*ComputerPurchaseCreate, bool)`

GetPurchasingOk returns a tuple with the Purchasing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasing

`func (o *ComputerInventoryCreateRequestV2) SetPurchasing(v ComputerPurchaseCreate)`

SetPurchasing sets Purchasing field to given value.

### HasPurchasing

`func (o *ComputerInventoryCreateRequestV2) HasPurchasing() bool`

HasPurchasing returns a boolean if a field has been set.

### GetUserAndLocation

`func (o *ComputerInventoryCreateRequestV2) GetUserAndLocation() ComputerUserAndLocationCreate`

GetUserAndLocation returns the UserAndLocation field if non-nil, zero value otherwise.

### GetUserAndLocationOk

`func (o *ComputerInventoryCreateRequestV2) GetUserAndLocationOk() (*ComputerUserAndLocationCreate, bool)`

GetUserAndLocationOk returns a tuple with the UserAndLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAndLocation

`func (o *ComputerInventoryCreateRequestV2) SetUserAndLocation(v ComputerUserAndLocationCreate)`

SetUserAndLocation sets UserAndLocation field to given value.

### HasUserAndLocation

`func (o *ComputerInventoryCreateRequestV2) HasUserAndLocation() bool`

HasUserAndLocation returns a boolean if a field has been set.

### GetHardware

`func (o *ComputerInventoryCreateRequestV2) GetHardware() ComputerHardwareCreate`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *ComputerInventoryCreateRequestV2) GetHardwareOk() (*ComputerHardwareCreate, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *ComputerInventoryCreateRequestV2) SetHardware(v ComputerHardwareCreate)`

SetHardware sets Hardware field to given value.

### HasHardware

`func (o *ComputerInventoryCreateRequestV2) HasHardware() bool`

HasHardware returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *ComputerInventoryCreateRequestV2) GetOperatingSystem() ComputerOperatingSystemCreate`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *ComputerInventoryCreateRequestV2) GetOperatingSystemOk() (*ComputerOperatingSystemCreate, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *ComputerInventoryCreateRequestV2) SetOperatingSystem(v ComputerOperatingSystemCreate)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *ComputerInventoryCreateRequestV2) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


