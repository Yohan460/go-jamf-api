# ComputerInventoryCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Udid** | Pointer to **NullableString** |  | [optional] 
**General** | Pointer to [**ComputerGeneralCreate**](ComputerGeneralCreate.md) |  | [optional] 
**Applications** | Pointer to [**[]ComputerApplicationCreate**](ComputerApplicationCreate.md) |  | [optional] 
**Storage** | Pointer to [**ComputerStorageCreate**](ComputerStorageCreate.md) |  | [optional] 
**ConfigurationProfiles** | Pointer to [**[]ComputerConfigurationProfileCreate**](ComputerConfigurationProfileCreate.md) |  | [optional] 
**Printers** | Pointer to [**[]ComputerPrinterCreate**](ComputerPrinterCreate.md) |  | [optional] 
**Services** | Pointer to [**[]ComputerServiceCreate**](ComputerServiceCreate.md) |  | [optional] 
**LocalUserAccounts** | Pointer to [**[]ComputerLocalUserAccountCreate**](ComputerLocalUserAccountCreate.md) |  | [optional] 
**Certificates** | Pointer to [**[]ComputerCertificateCreate**](ComputerCertificateCreate.md) |  | [optional] 
**Plugins** | Pointer to **[]map[string]interface{}** |  | [optional] 
**PackageReceipts** | Pointer to [**ComputerPackageReceiptsCreate**](ComputerPackageReceiptsCreate.md) |  | [optional] 
**Fonts** | Pointer to [**[]ComputerFontCreate**](ComputerFontCreate.md) |  | [optional] 
**Security** | Pointer to [**ComputerSecurityCreate**](ComputerSecurityCreate.md) |  | [optional] 
**SoftwareUpdates** | Pointer to [**[]ComputerSoftwareUpdateCreate**](ComputerSoftwareUpdateCreate.md) |  | [optional] 
**Purchasing** | Pointer to [**ComputerPurchaseCreate**](ComputerPurchaseCreate.md) |  | [optional] 
**UserAndLocation** | Pointer to [**ComputerUserAndLocationCreate**](ComputerUserAndLocationCreate.md) |  | [optional] 
**Hardware** | Pointer to [**ComputerHardwareCreate**](ComputerHardwareCreate.md) |  | [optional] 
**OperatingSystem** | Pointer to [**ComputerOperatingSystemCreate**](ComputerOperatingSystemCreate.md) |  | [optional] 

## Methods

### NewComputerInventoryCreateRequest

`func NewComputerInventoryCreateRequest() *ComputerInventoryCreateRequest`

NewComputerInventoryCreateRequest instantiates a new ComputerInventoryCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerInventoryCreateRequestWithDefaults

`func NewComputerInventoryCreateRequestWithDefaults() *ComputerInventoryCreateRequest`

NewComputerInventoryCreateRequestWithDefaults instantiates a new ComputerInventoryCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUdid

`func (o *ComputerInventoryCreateRequest) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *ComputerInventoryCreateRequest) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *ComputerInventoryCreateRequest) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *ComputerInventoryCreateRequest) HasUdid() bool`

HasUdid returns a boolean if a field has been set.

### SetUdidNil

`func (o *ComputerInventoryCreateRequest) SetUdidNil(b bool)`

 SetUdidNil sets the value for Udid to be an explicit nil

### UnsetUdid
`func (o *ComputerInventoryCreateRequest) UnsetUdid()`

UnsetUdid ensures that no value is present for Udid, not even an explicit nil
### GetGeneral

`func (o *ComputerInventoryCreateRequest) GetGeneral() ComputerGeneralCreate`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *ComputerInventoryCreateRequest) GetGeneralOk() (*ComputerGeneralCreate, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *ComputerInventoryCreateRequest) SetGeneral(v ComputerGeneralCreate)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *ComputerInventoryCreateRequest) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetApplications

`func (o *ComputerInventoryCreateRequest) GetApplications() []ComputerApplicationCreate`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ComputerInventoryCreateRequest) GetApplicationsOk() (*[]ComputerApplicationCreate, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ComputerInventoryCreateRequest) SetApplications(v []ComputerApplicationCreate)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *ComputerInventoryCreateRequest) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### SetApplicationsNil

`func (o *ComputerInventoryCreateRequest) SetApplicationsNil(b bool)`

 SetApplicationsNil sets the value for Applications to be an explicit nil

### UnsetApplications
`func (o *ComputerInventoryCreateRequest) UnsetApplications()`

UnsetApplications ensures that no value is present for Applications, not even an explicit nil
### GetStorage

`func (o *ComputerInventoryCreateRequest) GetStorage() ComputerStorageCreate`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ComputerInventoryCreateRequest) GetStorageOk() (*ComputerStorageCreate, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ComputerInventoryCreateRequest) SetStorage(v ComputerStorageCreate)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ComputerInventoryCreateRequest) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetConfigurationProfiles

`func (o *ComputerInventoryCreateRequest) GetConfigurationProfiles() []ComputerConfigurationProfileCreate`

GetConfigurationProfiles returns the ConfigurationProfiles field if non-nil, zero value otherwise.

### GetConfigurationProfilesOk

`func (o *ComputerInventoryCreateRequest) GetConfigurationProfilesOk() (*[]ComputerConfigurationProfileCreate, bool)`

GetConfigurationProfilesOk returns a tuple with the ConfigurationProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationProfiles

`func (o *ComputerInventoryCreateRequest) SetConfigurationProfiles(v []ComputerConfigurationProfileCreate)`

SetConfigurationProfiles sets ConfigurationProfiles field to given value.

### HasConfigurationProfiles

`func (o *ComputerInventoryCreateRequest) HasConfigurationProfiles() bool`

HasConfigurationProfiles returns a boolean if a field has been set.

### SetConfigurationProfilesNil

`func (o *ComputerInventoryCreateRequest) SetConfigurationProfilesNil(b bool)`

 SetConfigurationProfilesNil sets the value for ConfigurationProfiles to be an explicit nil

### UnsetConfigurationProfiles
`func (o *ComputerInventoryCreateRequest) UnsetConfigurationProfiles()`

UnsetConfigurationProfiles ensures that no value is present for ConfigurationProfiles, not even an explicit nil
### GetPrinters

`func (o *ComputerInventoryCreateRequest) GetPrinters() []ComputerPrinterCreate`

GetPrinters returns the Printers field if non-nil, zero value otherwise.

### GetPrintersOk

`func (o *ComputerInventoryCreateRequest) GetPrintersOk() (*[]ComputerPrinterCreate, bool)`

GetPrintersOk returns a tuple with the Printers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinters

`func (o *ComputerInventoryCreateRequest) SetPrinters(v []ComputerPrinterCreate)`

SetPrinters sets Printers field to given value.

### HasPrinters

`func (o *ComputerInventoryCreateRequest) HasPrinters() bool`

HasPrinters returns a boolean if a field has been set.

### SetPrintersNil

`func (o *ComputerInventoryCreateRequest) SetPrintersNil(b bool)`

 SetPrintersNil sets the value for Printers to be an explicit nil

### UnsetPrinters
`func (o *ComputerInventoryCreateRequest) UnsetPrinters()`

UnsetPrinters ensures that no value is present for Printers, not even an explicit nil
### GetServices

`func (o *ComputerInventoryCreateRequest) GetServices() []ComputerServiceCreate`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ComputerInventoryCreateRequest) GetServicesOk() (*[]ComputerServiceCreate, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ComputerInventoryCreateRequest) SetServices(v []ComputerServiceCreate)`

SetServices sets Services field to given value.

### HasServices

`func (o *ComputerInventoryCreateRequest) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *ComputerInventoryCreateRequest) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *ComputerInventoryCreateRequest) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetLocalUserAccounts

`func (o *ComputerInventoryCreateRequest) GetLocalUserAccounts() []ComputerLocalUserAccountCreate`

GetLocalUserAccounts returns the LocalUserAccounts field if non-nil, zero value otherwise.

### GetLocalUserAccountsOk

`func (o *ComputerInventoryCreateRequest) GetLocalUserAccountsOk() (*[]ComputerLocalUserAccountCreate, bool)`

GetLocalUserAccountsOk returns a tuple with the LocalUserAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserAccounts

`func (o *ComputerInventoryCreateRequest) SetLocalUserAccounts(v []ComputerLocalUserAccountCreate)`

SetLocalUserAccounts sets LocalUserAccounts field to given value.

### HasLocalUserAccounts

`func (o *ComputerInventoryCreateRequest) HasLocalUserAccounts() bool`

HasLocalUserAccounts returns a boolean if a field has been set.

### SetLocalUserAccountsNil

`func (o *ComputerInventoryCreateRequest) SetLocalUserAccountsNil(b bool)`

 SetLocalUserAccountsNil sets the value for LocalUserAccounts to be an explicit nil

### UnsetLocalUserAccounts
`func (o *ComputerInventoryCreateRequest) UnsetLocalUserAccounts()`

UnsetLocalUserAccounts ensures that no value is present for LocalUserAccounts, not even an explicit nil
### GetCertificates

`func (o *ComputerInventoryCreateRequest) GetCertificates() []ComputerCertificateCreate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *ComputerInventoryCreateRequest) GetCertificatesOk() (*[]ComputerCertificateCreate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *ComputerInventoryCreateRequest) SetCertificates(v []ComputerCertificateCreate)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *ComputerInventoryCreateRequest) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### SetCertificatesNil

`func (o *ComputerInventoryCreateRequest) SetCertificatesNil(b bool)`

 SetCertificatesNil sets the value for Certificates to be an explicit nil

### UnsetCertificates
`func (o *ComputerInventoryCreateRequest) UnsetCertificates()`

UnsetCertificates ensures that no value is present for Certificates, not even an explicit nil
### GetPlugins

`func (o *ComputerInventoryCreateRequest) GetPlugins() []map[string]interface{}`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *ComputerInventoryCreateRequest) GetPluginsOk() (*[]map[string]interface{}, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *ComputerInventoryCreateRequest) SetPlugins(v []map[string]interface{})`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *ComputerInventoryCreateRequest) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### SetPluginsNil

`func (o *ComputerInventoryCreateRequest) SetPluginsNil(b bool)`

 SetPluginsNil sets the value for Plugins to be an explicit nil

### UnsetPlugins
`func (o *ComputerInventoryCreateRequest) UnsetPlugins()`

UnsetPlugins ensures that no value is present for Plugins, not even an explicit nil
### GetPackageReceipts

`func (o *ComputerInventoryCreateRequest) GetPackageReceipts() ComputerPackageReceiptsCreate`

GetPackageReceipts returns the PackageReceipts field if non-nil, zero value otherwise.

### GetPackageReceiptsOk

`func (o *ComputerInventoryCreateRequest) GetPackageReceiptsOk() (*ComputerPackageReceiptsCreate, bool)`

GetPackageReceiptsOk returns a tuple with the PackageReceipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageReceipts

`func (o *ComputerInventoryCreateRequest) SetPackageReceipts(v ComputerPackageReceiptsCreate)`

SetPackageReceipts sets PackageReceipts field to given value.

### HasPackageReceipts

`func (o *ComputerInventoryCreateRequest) HasPackageReceipts() bool`

HasPackageReceipts returns a boolean if a field has been set.

### GetFonts

`func (o *ComputerInventoryCreateRequest) GetFonts() []ComputerFontCreate`

GetFonts returns the Fonts field if non-nil, zero value otherwise.

### GetFontsOk

`func (o *ComputerInventoryCreateRequest) GetFontsOk() (*[]ComputerFontCreate, bool)`

GetFontsOk returns a tuple with the Fonts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFonts

`func (o *ComputerInventoryCreateRequest) SetFonts(v []ComputerFontCreate)`

SetFonts sets Fonts field to given value.

### HasFonts

`func (o *ComputerInventoryCreateRequest) HasFonts() bool`

HasFonts returns a boolean if a field has been set.

### SetFontsNil

`func (o *ComputerInventoryCreateRequest) SetFontsNil(b bool)`

 SetFontsNil sets the value for Fonts to be an explicit nil

### UnsetFonts
`func (o *ComputerInventoryCreateRequest) UnsetFonts()`

UnsetFonts ensures that no value is present for Fonts, not even an explicit nil
### GetSecurity

`func (o *ComputerInventoryCreateRequest) GetSecurity() ComputerSecurityCreate`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *ComputerInventoryCreateRequest) GetSecurityOk() (*ComputerSecurityCreate, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *ComputerInventoryCreateRequest) SetSecurity(v ComputerSecurityCreate)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *ComputerInventoryCreateRequest) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetSoftwareUpdates

`func (o *ComputerInventoryCreateRequest) GetSoftwareUpdates() []ComputerSoftwareUpdateCreate`

GetSoftwareUpdates returns the SoftwareUpdates field if non-nil, zero value otherwise.

### GetSoftwareUpdatesOk

`func (o *ComputerInventoryCreateRequest) GetSoftwareUpdatesOk() (*[]ComputerSoftwareUpdateCreate, bool)`

GetSoftwareUpdatesOk returns a tuple with the SoftwareUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdates

`func (o *ComputerInventoryCreateRequest) SetSoftwareUpdates(v []ComputerSoftwareUpdateCreate)`

SetSoftwareUpdates sets SoftwareUpdates field to given value.

### HasSoftwareUpdates

`func (o *ComputerInventoryCreateRequest) HasSoftwareUpdates() bool`

HasSoftwareUpdates returns a boolean if a field has been set.

### SetSoftwareUpdatesNil

`func (o *ComputerInventoryCreateRequest) SetSoftwareUpdatesNil(b bool)`

 SetSoftwareUpdatesNil sets the value for SoftwareUpdates to be an explicit nil

### UnsetSoftwareUpdates
`func (o *ComputerInventoryCreateRequest) UnsetSoftwareUpdates()`

UnsetSoftwareUpdates ensures that no value is present for SoftwareUpdates, not even an explicit nil
### GetPurchasing

`func (o *ComputerInventoryCreateRequest) GetPurchasing() ComputerPurchaseCreate`

GetPurchasing returns the Purchasing field if non-nil, zero value otherwise.

### GetPurchasingOk

`func (o *ComputerInventoryCreateRequest) GetPurchasingOk() (*ComputerPurchaseCreate, bool)`

GetPurchasingOk returns a tuple with the Purchasing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasing

`func (o *ComputerInventoryCreateRequest) SetPurchasing(v ComputerPurchaseCreate)`

SetPurchasing sets Purchasing field to given value.

### HasPurchasing

`func (o *ComputerInventoryCreateRequest) HasPurchasing() bool`

HasPurchasing returns a boolean if a field has been set.

### GetUserAndLocation

`func (o *ComputerInventoryCreateRequest) GetUserAndLocation() ComputerUserAndLocationCreate`

GetUserAndLocation returns the UserAndLocation field if non-nil, zero value otherwise.

### GetUserAndLocationOk

`func (o *ComputerInventoryCreateRequest) GetUserAndLocationOk() (*ComputerUserAndLocationCreate, bool)`

GetUserAndLocationOk returns a tuple with the UserAndLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAndLocation

`func (o *ComputerInventoryCreateRequest) SetUserAndLocation(v ComputerUserAndLocationCreate)`

SetUserAndLocation sets UserAndLocation field to given value.

### HasUserAndLocation

`func (o *ComputerInventoryCreateRequest) HasUserAndLocation() bool`

HasUserAndLocation returns a boolean if a field has been set.

### GetHardware

`func (o *ComputerInventoryCreateRequest) GetHardware() ComputerHardwareCreate`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *ComputerInventoryCreateRequest) GetHardwareOk() (*ComputerHardwareCreate, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *ComputerInventoryCreateRequest) SetHardware(v ComputerHardwareCreate)`

SetHardware sets Hardware field to given value.

### HasHardware

`func (o *ComputerInventoryCreateRequest) HasHardware() bool`

HasHardware returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *ComputerInventoryCreateRequest) GetOperatingSystem() ComputerOperatingSystemCreate`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *ComputerInventoryCreateRequest) GetOperatingSystemOk() (*ComputerOperatingSystemCreate, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *ComputerInventoryCreateRequest) SetOperatingSystem(v ComputerOperatingSystemCreate)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *ComputerInventoryCreateRequest) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


