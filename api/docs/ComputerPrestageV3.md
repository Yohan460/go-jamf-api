# ComputerPrestageV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**Mandatory** | **bool** |  | 
**MdmRemovable** | **bool** |  | 
**SupportPhoneNumber** | **string** |  | 
**SupportEmailAddress** | **string** |  | 
**Department** | **string** |  | 
**DefaultPrestage** | **bool** |  | 
**EnrollmentSiteId** | **string** |  | 
**KeepExistingSiteMembership** | **bool** |  | 
**KeepExistingLocationInformation** | **bool** |  | 
**RequireAuthentication** | **bool** |  | 
**AuthenticationPrompt** | **string** |  | 
**PreventActivationLock** | **bool** |  | 
**EnableDeviceBasedActivationLock** | **bool** |  | 
**DeviceEnrollmentProgramInstanceId** | **string** |  | 
**SkipSetupItems** | Pointer to **map[string]bool** |  | [optional] 
**LocationInformation** | [**LocationInformationV2**](LocationInformationV2.md) |  | 
**PurchasingInformation** | [**PrestagePurchasingInformationV2**](PrestagePurchasingInformationV2.md) |  | 
**AnchorCertificates** | Pointer to **[]string** | The Base64 encoded PEM Certificate | [optional] 
**EnrollmentCustomizationId** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**AutoAdvanceSetup** | **bool** |  | 
**InstallProfilesDuringSetup** | **bool** |  | 
**PrestageInstalledProfileIds** | **[]string** |  | 
**CustomPackageIds** | **[]string** |  | 
**CustomPackageDistributionPointId** | **string** |  | 
**EnableRecoveryLock** | Pointer to **bool** |  | [optional] 
**RecoveryLockPasswordType** | Pointer to **string** |  | [optional] 
**RotateRecoveryLockPassword** | Pointer to **bool** |  | [optional] 

## Methods

### NewComputerPrestageV3

`func NewComputerPrestageV3(displayName string, mandatory bool, mdmRemovable bool, supportPhoneNumber string, supportEmailAddress string, department string, defaultPrestage bool, enrollmentSiteId string, keepExistingSiteMembership bool, keepExistingLocationInformation bool, requireAuthentication bool, authenticationPrompt string, preventActivationLock bool, enableDeviceBasedActivationLock bool, deviceEnrollmentProgramInstanceId string, locationInformation LocationInformationV2, purchasingInformation PrestagePurchasingInformationV2, autoAdvanceSetup bool, installProfilesDuringSetup bool, prestageInstalledProfileIds []string, customPackageIds []string, customPackageDistributionPointId string, ) *ComputerPrestageV3`

NewComputerPrestageV3 instantiates a new ComputerPrestageV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerPrestageV3WithDefaults

`func NewComputerPrestageV3WithDefaults() *ComputerPrestageV3`

NewComputerPrestageV3WithDefaults instantiates a new ComputerPrestageV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ComputerPrestageV3) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ComputerPrestageV3) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ComputerPrestageV3) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMandatory

`func (o *ComputerPrestageV3) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *ComputerPrestageV3) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *ComputerPrestageV3) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.


### GetMdmRemovable

`func (o *ComputerPrestageV3) GetMdmRemovable() bool`

GetMdmRemovable returns the MdmRemovable field if non-nil, zero value otherwise.

### GetMdmRemovableOk

`func (o *ComputerPrestageV3) GetMdmRemovableOk() (*bool, bool)`

GetMdmRemovableOk returns a tuple with the MdmRemovable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmRemovable

`func (o *ComputerPrestageV3) SetMdmRemovable(v bool)`

SetMdmRemovable sets MdmRemovable field to given value.


### GetSupportPhoneNumber

`func (o *ComputerPrestageV3) GetSupportPhoneNumber() string`

GetSupportPhoneNumber returns the SupportPhoneNumber field if non-nil, zero value otherwise.

### GetSupportPhoneNumberOk

`func (o *ComputerPrestageV3) GetSupportPhoneNumberOk() (*string, bool)`

GetSupportPhoneNumberOk returns a tuple with the SupportPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPhoneNumber

`func (o *ComputerPrestageV3) SetSupportPhoneNumber(v string)`

SetSupportPhoneNumber sets SupportPhoneNumber field to given value.


### GetSupportEmailAddress

`func (o *ComputerPrestageV3) GetSupportEmailAddress() string`

GetSupportEmailAddress returns the SupportEmailAddress field if non-nil, zero value otherwise.

### GetSupportEmailAddressOk

`func (o *ComputerPrestageV3) GetSupportEmailAddressOk() (*string, bool)`

GetSupportEmailAddressOk returns a tuple with the SupportEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmailAddress

`func (o *ComputerPrestageV3) SetSupportEmailAddress(v string)`

SetSupportEmailAddress sets SupportEmailAddress field to given value.


### GetDepartment

`func (o *ComputerPrestageV3) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *ComputerPrestageV3) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *ComputerPrestageV3) SetDepartment(v string)`

SetDepartment sets Department field to given value.


### GetDefaultPrestage

`func (o *ComputerPrestageV3) GetDefaultPrestage() bool`

GetDefaultPrestage returns the DefaultPrestage field if non-nil, zero value otherwise.

### GetDefaultPrestageOk

`func (o *ComputerPrestageV3) GetDefaultPrestageOk() (*bool, bool)`

GetDefaultPrestageOk returns a tuple with the DefaultPrestage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrestage

`func (o *ComputerPrestageV3) SetDefaultPrestage(v bool)`

SetDefaultPrestage sets DefaultPrestage field to given value.


### GetEnrollmentSiteId

`func (o *ComputerPrestageV3) GetEnrollmentSiteId() string`

GetEnrollmentSiteId returns the EnrollmentSiteId field if non-nil, zero value otherwise.

### GetEnrollmentSiteIdOk

`func (o *ComputerPrestageV3) GetEnrollmentSiteIdOk() (*string, bool)`

GetEnrollmentSiteIdOk returns a tuple with the EnrollmentSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSiteId

`func (o *ComputerPrestageV3) SetEnrollmentSiteId(v string)`

SetEnrollmentSiteId sets EnrollmentSiteId field to given value.


### GetKeepExistingSiteMembership

`func (o *ComputerPrestageV3) GetKeepExistingSiteMembership() bool`

GetKeepExistingSiteMembership returns the KeepExistingSiteMembership field if non-nil, zero value otherwise.

### GetKeepExistingSiteMembershipOk

`func (o *ComputerPrestageV3) GetKeepExistingSiteMembershipOk() (*bool, bool)`

GetKeepExistingSiteMembershipOk returns a tuple with the KeepExistingSiteMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepExistingSiteMembership

`func (o *ComputerPrestageV3) SetKeepExistingSiteMembership(v bool)`

SetKeepExistingSiteMembership sets KeepExistingSiteMembership field to given value.


### GetKeepExistingLocationInformation

`func (o *ComputerPrestageV3) GetKeepExistingLocationInformation() bool`

GetKeepExistingLocationInformation returns the KeepExistingLocationInformation field if non-nil, zero value otherwise.

### GetKeepExistingLocationInformationOk

`func (o *ComputerPrestageV3) GetKeepExistingLocationInformationOk() (*bool, bool)`

GetKeepExistingLocationInformationOk returns a tuple with the KeepExistingLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepExistingLocationInformation

`func (o *ComputerPrestageV3) SetKeepExistingLocationInformation(v bool)`

SetKeepExistingLocationInformation sets KeepExistingLocationInformation field to given value.


### GetRequireAuthentication

`func (o *ComputerPrestageV3) GetRequireAuthentication() bool`

GetRequireAuthentication returns the RequireAuthentication field if non-nil, zero value otherwise.

### GetRequireAuthenticationOk

`func (o *ComputerPrestageV3) GetRequireAuthenticationOk() (*bool, bool)`

GetRequireAuthenticationOk returns a tuple with the RequireAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuthentication

`func (o *ComputerPrestageV3) SetRequireAuthentication(v bool)`

SetRequireAuthentication sets RequireAuthentication field to given value.


### GetAuthenticationPrompt

`func (o *ComputerPrestageV3) GetAuthenticationPrompt() string`

GetAuthenticationPrompt returns the AuthenticationPrompt field if non-nil, zero value otherwise.

### GetAuthenticationPromptOk

`func (o *ComputerPrestageV3) GetAuthenticationPromptOk() (*string, bool)`

GetAuthenticationPromptOk returns a tuple with the AuthenticationPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPrompt

`func (o *ComputerPrestageV3) SetAuthenticationPrompt(v string)`

SetAuthenticationPrompt sets AuthenticationPrompt field to given value.


### GetPreventActivationLock

`func (o *ComputerPrestageV3) GetPreventActivationLock() bool`

GetPreventActivationLock returns the PreventActivationLock field if non-nil, zero value otherwise.

### GetPreventActivationLockOk

`func (o *ComputerPrestageV3) GetPreventActivationLockOk() (*bool, bool)`

GetPreventActivationLockOk returns a tuple with the PreventActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventActivationLock

`func (o *ComputerPrestageV3) SetPreventActivationLock(v bool)`

SetPreventActivationLock sets PreventActivationLock field to given value.


### GetEnableDeviceBasedActivationLock

`func (o *ComputerPrestageV3) GetEnableDeviceBasedActivationLock() bool`

GetEnableDeviceBasedActivationLock returns the EnableDeviceBasedActivationLock field if non-nil, zero value otherwise.

### GetEnableDeviceBasedActivationLockOk

`func (o *ComputerPrestageV3) GetEnableDeviceBasedActivationLockOk() (*bool, bool)`

GetEnableDeviceBasedActivationLockOk returns a tuple with the EnableDeviceBasedActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDeviceBasedActivationLock

`func (o *ComputerPrestageV3) SetEnableDeviceBasedActivationLock(v bool)`

SetEnableDeviceBasedActivationLock sets EnableDeviceBasedActivationLock field to given value.


### GetDeviceEnrollmentProgramInstanceId

`func (o *ComputerPrestageV3) GetDeviceEnrollmentProgramInstanceId() string`

GetDeviceEnrollmentProgramInstanceId returns the DeviceEnrollmentProgramInstanceId field if non-nil, zero value otherwise.

### GetDeviceEnrollmentProgramInstanceIdOk

`func (o *ComputerPrestageV3) GetDeviceEnrollmentProgramInstanceIdOk() (*string, bool)`

GetDeviceEnrollmentProgramInstanceIdOk returns a tuple with the DeviceEnrollmentProgramInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceEnrollmentProgramInstanceId

`func (o *ComputerPrestageV3) SetDeviceEnrollmentProgramInstanceId(v string)`

SetDeviceEnrollmentProgramInstanceId sets DeviceEnrollmentProgramInstanceId field to given value.


### GetSkipSetupItems

`func (o *ComputerPrestageV3) GetSkipSetupItems() map[string]bool`

GetSkipSetupItems returns the SkipSetupItems field if non-nil, zero value otherwise.

### GetSkipSetupItemsOk

`func (o *ComputerPrestageV3) GetSkipSetupItemsOk() (*map[string]bool, bool)`

GetSkipSetupItemsOk returns a tuple with the SkipSetupItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSetupItems

`func (o *ComputerPrestageV3) SetSkipSetupItems(v map[string]bool)`

SetSkipSetupItems sets SkipSetupItems field to given value.

### HasSkipSetupItems

`func (o *ComputerPrestageV3) HasSkipSetupItems() bool`

HasSkipSetupItems returns a boolean if a field has been set.

### GetLocationInformation

`func (o *ComputerPrestageV3) GetLocationInformation() LocationInformationV2`

GetLocationInformation returns the LocationInformation field if non-nil, zero value otherwise.

### GetLocationInformationOk

`func (o *ComputerPrestageV3) GetLocationInformationOk() (*LocationInformationV2, bool)`

GetLocationInformationOk returns a tuple with the LocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationInformation

`func (o *ComputerPrestageV3) SetLocationInformation(v LocationInformationV2)`

SetLocationInformation sets LocationInformation field to given value.


### GetPurchasingInformation

`func (o *ComputerPrestageV3) GetPurchasingInformation() PrestagePurchasingInformationV2`

GetPurchasingInformation returns the PurchasingInformation field if non-nil, zero value otherwise.

### GetPurchasingInformationOk

`func (o *ComputerPrestageV3) GetPurchasingInformationOk() (*PrestagePurchasingInformationV2, bool)`

GetPurchasingInformationOk returns a tuple with the PurchasingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingInformation

`func (o *ComputerPrestageV3) SetPurchasingInformation(v PrestagePurchasingInformationV2)`

SetPurchasingInformation sets PurchasingInformation field to given value.


### GetAnchorCertificates

`func (o *ComputerPrestageV3) GetAnchorCertificates() []string`

GetAnchorCertificates returns the AnchorCertificates field if non-nil, zero value otherwise.

### GetAnchorCertificatesOk

`func (o *ComputerPrestageV3) GetAnchorCertificatesOk() (*[]string, bool)`

GetAnchorCertificatesOk returns a tuple with the AnchorCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorCertificates

`func (o *ComputerPrestageV3) SetAnchorCertificates(v []string)`

SetAnchorCertificates sets AnchorCertificates field to given value.

### HasAnchorCertificates

`func (o *ComputerPrestageV3) HasAnchorCertificates() bool`

HasAnchorCertificates returns a boolean if a field has been set.

### GetEnrollmentCustomizationId

`func (o *ComputerPrestageV3) GetEnrollmentCustomizationId() string`

GetEnrollmentCustomizationId returns the EnrollmentCustomizationId field if non-nil, zero value otherwise.

### GetEnrollmentCustomizationIdOk

`func (o *ComputerPrestageV3) GetEnrollmentCustomizationIdOk() (*string, bool)`

GetEnrollmentCustomizationIdOk returns a tuple with the EnrollmentCustomizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCustomizationId

`func (o *ComputerPrestageV3) SetEnrollmentCustomizationId(v string)`

SetEnrollmentCustomizationId sets EnrollmentCustomizationId field to given value.

### HasEnrollmentCustomizationId

`func (o *ComputerPrestageV3) HasEnrollmentCustomizationId() bool`

HasEnrollmentCustomizationId returns a boolean if a field has been set.

### GetLanguage

`func (o *ComputerPrestageV3) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ComputerPrestageV3) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ComputerPrestageV3) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ComputerPrestageV3) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetRegion

`func (o *ComputerPrestageV3) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ComputerPrestageV3) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ComputerPrestageV3) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ComputerPrestageV3) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetAutoAdvanceSetup

`func (o *ComputerPrestageV3) GetAutoAdvanceSetup() bool`

GetAutoAdvanceSetup returns the AutoAdvanceSetup field if non-nil, zero value otherwise.

### GetAutoAdvanceSetupOk

`func (o *ComputerPrestageV3) GetAutoAdvanceSetupOk() (*bool, bool)`

GetAutoAdvanceSetupOk returns a tuple with the AutoAdvanceSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAdvanceSetup

`func (o *ComputerPrestageV3) SetAutoAdvanceSetup(v bool)`

SetAutoAdvanceSetup sets AutoAdvanceSetup field to given value.


### GetInstallProfilesDuringSetup

`func (o *ComputerPrestageV3) GetInstallProfilesDuringSetup() bool`

GetInstallProfilesDuringSetup returns the InstallProfilesDuringSetup field if non-nil, zero value otherwise.

### GetInstallProfilesDuringSetupOk

`func (o *ComputerPrestageV3) GetInstallProfilesDuringSetupOk() (*bool, bool)`

GetInstallProfilesDuringSetupOk returns a tuple with the InstallProfilesDuringSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallProfilesDuringSetup

`func (o *ComputerPrestageV3) SetInstallProfilesDuringSetup(v bool)`

SetInstallProfilesDuringSetup sets InstallProfilesDuringSetup field to given value.


### GetPrestageInstalledProfileIds

`func (o *ComputerPrestageV3) GetPrestageInstalledProfileIds() []string`

GetPrestageInstalledProfileIds returns the PrestageInstalledProfileIds field if non-nil, zero value otherwise.

### GetPrestageInstalledProfileIdsOk

`func (o *ComputerPrestageV3) GetPrestageInstalledProfileIdsOk() (*[]string, bool)`

GetPrestageInstalledProfileIdsOk returns a tuple with the PrestageInstalledProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrestageInstalledProfileIds

`func (o *ComputerPrestageV3) SetPrestageInstalledProfileIds(v []string)`

SetPrestageInstalledProfileIds sets PrestageInstalledProfileIds field to given value.


### GetCustomPackageIds

`func (o *ComputerPrestageV3) GetCustomPackageIds() []string`

GetCustomPackageIds returns the CustomPackageIds field if non-nil, zero value otherwise.

### GetCustomPackageIdsOk

`func (o *ComputerPrestageV3) GetCustomPackageIdsOk() (*[]string, bool)`

GetCustomPackageIdsOk returns a tuple with the CustomPackageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPackageIds

`func (o *ComputerPrestageV3) SetCustomPackageIds(v []string)`

SetCustomPackageIds sets CustomPackageIds field to given value.


### GetCustomPackageDistributionPointId

`func (o *ComputerPrestageV3) GetCustomPackageDistributionPointId() string`

GetCustomPackageDistributionPointId returns the CustomPackageDistributionPointId field if non-nil, zero value otherwise.

### GetCustomPackageDistributionPointIdOk

`func (o *ComputerPrestageV3) GetCustomPackageDistributionPointIdOk() (*string, bool)`

GetCustomPackageDistributionPointIdOk returns a tuple with the CustomPackageDistributionPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPackageDistributionPointId

`func (o *ComputerPrestageV3) SetCustomPackageDistributionPointId(v string)`

SetCustomPackageDistributionPointId sets CustomPackageDistributionPointId field to given value.


### GetEnableRecoveryLock

`func (o *ComputerPrestageV3) GetEnableRecoveryLock() bool`

GetEnableRecoveryLock returns the EnableRecoveryLock field if non-nil, zero value otherwise.

### GetEnableRecoveryLockOk

`func (o *ComputerPrestageV3) GetEnableRecoveryLockOk() (*bool, bool)`

GetEnableRecoveryLockOk returns a tuple with the EnableRecoveryLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRecoveryLock

`func (o *ComputerPrestageV3) SetEnableRecoveryLock(v bool)`

SetEnableRecoveryLock sets EnableRecoveryLock field to given value.

### HasEnableRecoveryLock

`func (o *ComputerPrestageV3) HasEnableRecoveryLock() bool`

HasEnableRecoveryLock returns a boolean if a field has been set.

### GetRecoveryLockPasswordType

`func (o *ComputerPrestageV3) GetRecoveryLockPasswordType() string`

GetRecoveryLockPasswordType returns the RecoveryLockPasswordType field if non-nil, zero value otherwise.

### GetRecoveryLockPasswordTypeOk

`func (o *ComputerPrestageV3) GetRecoveryLockPasswordTypeOk() (*string, bool)`

GetRecoveryLockPasswordTypeOk returns a tuple with the RecoveryLockPasswordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryLockPasswordType

`func (o *ComputerPrestageV3) SetRecoveryLockPasswordType(v string)`

SetRecoveryLockPasswordType sets RecoveryLockPasswordType field to given value.

### HasRecoveryLockPasswordType

`func (o *ComputerPrestageV3) HasRecoveryLockPasswordType() bool`

HasRecoveryLockPasswordType returns a boolean if a field has been set.

### GetRotateRecoveryLockPassword

`func (o *ComputerPrestageV3) GetRotateRecoveryLockPassword() bool`

GetRotateRecoveryLockPassword returns the RotateRecoveryLockPassword field if non-nil, zero value otherwise.

### GetRotateRecoveryLockPasswordOk

`func (o *ComputerPrestageV3) GetRotateRecoveryLockPasswordOk() (*bool, bool)`

GetRotateRecoveryLockPasswordOk returns a tuple with the RotateRecoveryLockPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateRecoveryLockPassword

`func (o *ComputerPrestageV3) SetRotateRecoveryLockPassword(v bool)`

SetRotateRecoveryLockPassword sets RotateRecoveryLockPassword field to given value.

### HasRotateRecoveryLockPassword

`func (o *ComputerPrestageV3) HasRotateRecoveryLockPassword() bool`

HasRotateRecoveryLockPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


