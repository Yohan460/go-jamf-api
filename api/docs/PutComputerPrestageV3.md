# PutComputerPrestageV3

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
**AccountSettings** | Pointer to [**AccountSettingsRequest**](AccountSettingsRequest.md) |  | [optional] 
**RecoveryLockPassword** | Pointer to **string** |  | [optional] 
**VersionLock** | Pointer to **int64** |  | [optional] 

## Methods

### NewPutComputerPrestageV3

`func NewPutComputerPrestageV3(displayName string, mandatory bool, mdmRemovable bool, supportPhoneNumber string, supportEmailAddress string, department string, defaultPrestage bool, enrollmentSiteId string, keepExistingSiteMembership bool, keepExistingLocationInformation bool, requireAuthentication bool, authenticationPrompt string, preventActivationLock bool, enableDeviceBasedActivationLock bool, deviceEnrollmentProgramInstanceId string, locationInformation LocationInformationV2, purchasingInformation PrestagePurchasingInformationV2, autoAdvanceSetup bool, installProfilesDuringSetup bool, prestageInstalledProfileIds []string, customPackageIds []string, customPackageDistributionPointId string, ) *PutComputerPrestageV3`

NewPutComputerPrestageV3 instantiates a new PutComputerPrestageV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutComputerPrestageV3WithDefaults

`func NewPutComputerPrestageV3WithDefaults() *PutComputerPrestageV3`

NewPutComputerPrestageV3WithDefaults instantiates a new PutComputerPrestageV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *PutComputerPrestageV3) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PutComputerPrestageV3) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PutComputerPrestageV3) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMandatory

`func (o *PutComputerPrestageV3) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *PutComputerPrestageV3) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *PutComputerPrestageV3) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.


### GetMdmRemovable

`func (o *PutComputerPrestageV3) GetMdmRemovable() bool`

GetMdmRemovable returns the MdmRemovable field if non-nil, zero value otherwise.

### GetMdmRemovableOk

`func (o *PutComputerPrestageV3) GetMdmRemovableOk() (*bool, bool)`

GetMdmRemovableOk returns a tuple with the MdmRemovable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmRemovable

`func (o *PutComputerPrestageV3) SetMdmRemovable(v bool)`

SetMdmRemovable sets MdmRemovable field to given value.


### GetSupportPhoneNumber

`func (o *PutComputerPrestageV3) GetSupportPhoneNumber() string`

GetSupportPhoneNumber returns the SupportPhoneNumber field if non-nil, zero value otherwise.

### GetSupportPhoneNumberOk

`func (o *PutComputerPrestageV3) GetSupportPhoneNumberOk() (*string, bool)`

GetSupportPhoneNumberOk returns a tuple with the SupportPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPhoneNumber

`func (o *PutComputerPrestageV3) SetSupportPhoneNumber(v string)`

SetSupportPhoneNumber sets SupportPhoneNumber field to given value.


### GetSupportEmailAddress

`func (o *PutComputerPrestageV3) GetSupportEmailAddress() string`

GetSupportEmailAddress returns the SupportEmailAddress field if non-nil, zero value otherwise.

### GetSupportEmailAddressOk

`func (o *PutComputerPrestageV3) GetSupportEmailAddressOk() (*string, bool)`

GetSupportEmailAddressOk returns a tuple with the SupportEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmailAddress

`func (o *PutComputerPrestageV3) SetSupportEmailAddress(v string)`

SetSupportEmailAddress sets SupportEmailAddress field to given value.


### GetDepartment

`func (o *PutComputerPrestageV3) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *PutComputerPrestageV3) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *PutComputerPrestageV3) SetDepartment(v string)`

SetDepartment sets Department field to given value.


### GetDefaultPrestage

`func (o *PutComputerPrestageV3) GetDefaultPrestage() bool`

GetDefaultPrestage returns the DefaultPrestage field if non-nil, zero value otherwise.

### GetDefaultPrestageOk

`func (o *PutComputerPrestageV3) GetDefaultPrestageOk() (*bool, bool)`

GetDefaultPrestageOk returns a tuple with the DefaultPrestage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrestage

`func (o *PutComputerPrestageV3) SetDefaultPrestage(v bool)`

SetDefaultPrestage sets DefaultPrestage field to given value.


### GetEnrollmentSiteId

`func (o *PutComputerPrestageV3) GetEnrollmentSiteId() string`

GetEnrollmentSiteId returns the EnrollmentSiteId field if non-nil, zero value otherwise.

### GetEnrollmentSiteIdOk

`func (o *PutComputerPrestageV3) GetEnrollmentSiteIdOk() (*string, bool)`

GetEnrollmentSiteIdOk returns a tuple with the EnrollmentSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSiteId

`func (o *PutComputerPrestageV3) SetEnrollmentSiteId(v string)`

SetEnrollmentSiteId sets EnrollmentSiteId field to given value.


### GetKeepExistingSiteMembership

`func (o *PutComputerPrestageV3) GetKeepExistingSiteMembership() bool`

GetKeepExistingSiteMembership returns the KeepExistingSiteMembership field if non-nil, zero value otherwise.

### GetKeepExistingSiteMembershipOk

`func (o *PutComputerPrestageV3) GetKeepExistingSiteMembershipOk() (*bool, bool)`

GetKeepExistingSiteMembershipOk returns a tuple with the KeepExistingSiteMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepExistingSiteMembership

`func (o *PutComputerPrestageV3) SetKeepExistingSiteMembership(v bool)`

SetKeepExistingSiteMembership sets KeepExistingSiteMembership field to given value.


### GetKeepExistingLocationInformation

`func (o *PutComputerPrestageV3) GetKeepExistingLocationInformation() bool`

GetKeepExistingLocationInformation returns the KeepExistingLocationInformation field if non-nil, zero value otherwise.

### GetKeepExistingLocationInformationOk

`func (o *PutComputerPrestageV3) GetKeepExistingLocationInformationOk() (*bool, bool)`

GetKeepExistingLocationInformationOk returns a tuple with the KeepExistingLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepExistingLocationInformation

`func (o *PutComputerPrestageV3) SetKeepExistingLocationInformation(v bool)`

SetKeepExistingLocationInformation sets KeepExistingLocationInformation field to given value.


### GetRequireAuthentication

`func (o *PutComputerPrestageV3) GetRequireAuthentication() bool`

GetRequireAuthentication returns the RequireAuthentication field if non-nil, zero value otherwise.

### GetRequireAuthenticationOk

`func (o *PutComputerPrestageV3) GetRequireAuthenticationOk() (*bool, bool)`

GetRequireAuthenticationOk returns a tuple with the RequireAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuthentication

`func (o *PutComputerPrestageV3) SetRequireAuthentication(v bool)`

SetRequireAuthentication sets RequireAuthentication field to given value.


### GetAuthenticationPrompt

`func (o *PutComputerPrestageV3) GetAuthenticationPrompt() string`

GetAuthenticationPrompt returns the AuthenticationPrompt field if non-nil, zero value otherwise.

### GetAuthenticationPromptOk

`func (o *PutComputerPrestageV3) GetAuthenticationPromptOk() (*string, bool)`

GetAuthenticationPromptOk returns a tuple with the AuthenticationPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPrompt

`func (o *PutComputerPrestageV3) SetAuthenticationPrompt(v string)`

SetAuthenticationPrompt sets AuthenticationPrompt field to given value.


### GetPreventActivationLock

`func (o *PutComputerPrestageV3) GetPreventActivationLock() bool`

GetPreventActivationLock returns the PreventActivationLock field if non-nil, zero value otherwise.

### GetPreventActivationLockOk

`func (o *PutComputerPrestageV3) GetPreventActivationLockOk() (*bool, bool)`

GetPreventActivationLockOk returns a tuple with the PreventActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventActivationLock

`func (o *PutComputerPrestageV3) SetPreventActivationLock(v bool)`

SetPreventActivationLock sets PreventActivationLock field to given value.


### GetEnableDeviceBasedActivationLock

`func (o *PutComputerPrestageV3) GetEnableDeviceBasedActivationLock() bool`

GetEnableDeviceBasedActivationLock returns the EnableDeviceBasedActivationLock field if non-nil, zero value otherwise.

### GetEnableDeviceBasedActivationLockOk

`func (o *PutComputerPrestageV3) GetEnableDeviceBasedActivationLockOk() (*bool, bool)`

GetEnableDeviceBasedActivationLockOk returns a tuple with the EnableDeviceBasedActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDeviceBasedActivationLock

`func (o *PutComputerPrestageV3) SetEnableDeviceBasedActivationLock(v bool)`

SetEnableDeviceBasedActivationLock sets EnableDeviceBasedActivationLock field to given value.


### GetDeviceEnrollmentProgramInstanceId

`func (o *PutComputerPrestageV3) GetDeviceEnrollmentProgramInstanceId() string`

GetDeviceEnrollmentProgramInstanceId returns the DeviceEnrollmentProgramInstanceId field if non-nil, zero value otherwise.

### GetDeviceEnrollmentProgramInstanceIdOk

`func (o *PutComputerPrestageV3) GetDeviceEnrollmentProgramInstanceIdOk() (*string, bool)`

GetDeviceEnrollmentProgramInstanceIdOk returns a tuple with the DeviceEnrollmentProgramInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceEnrollmentProgramInstanceId

`func (o *PutComputerPrestageV3) SetDeviceEnrollmentProgramInstanceId(v string)`

SetDeviceEnrollmentProgramInstanceId sets DeviceEnrollmentProgramInstanceId field to given value.


### GetSkipSetupItems

`func (o *PutComputerPrestageV3) GetSkipSetupItems() map[string]bool`

GetSkipSetupItems returns the SkipSetupItems field if non-nil, zero value otherwise.

### GetSkipSetupItemsOk

`func (o *PutComputerPrestageV3) GetSkipSetupItemsOk() (*map[string]bool, bool)`

GetSkipSetupItemsOk returns a tuple with the SkipSetupItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSetupItems

`func (o *PutComputerPrestageV3) SetSkipSetupItems(v map[string]bool)`

SetSkipSetupItems sets SkipSetupItems field to given value.

### HasSkipSetupItems

`func (o *PutComputerPrestageV3) HasSkipSetupItems() bool`

HasSkipSetupItems returns a boolean if a field has been set.

### GetLocationInformation

`func (o *PutComputerPrestageV3) GetLocationInformation() LocationInformationV2`

GetLocationInformation returns the LocationInformation field if non-nil, zero value otherwise.

### GetLocationInformationOk

`func (o *PutComputerPrestageV3) GetLocationInformationOk() (*LocationInformationV2, bool)`

GetLocationInformationOk returns a tuple with the LocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationInformation

`func (o *PutComputerPrestageV3) SetLocationInformation(v LocationInformationV2)`

SetLocationInformation sets LocationInformation field to given value.


### GetPurchasingInformation

`func (o *PutComputerPrestageV3) GetPurchasingInformation() PrestagePurchasingInformationV2`

GetPurchasingInformation returns the PurchasingInformation field if non-nil, zero value otherwise.

### GetPurchasingInformationOk

`func (o *PutComputerPrestageV3) GetPurchasingInformationOk() (*PrestagePurchasingInformationV2, bool)`

GetPurchasingInformationOk returns a tuple with the PurchasingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingInformation

`func (o *PutComputerPrestageV3) SetPurchasingInformation(v PrestagePurchasingInformationV2)`

SetPurchasingInformation sets PurchasingInformation field to given value.


### GetAnchorCertificates

`func (o *PutComputerPrestageV3) GetAnchorCertificates() []string`

GetAnchorCertificates returns the AnchorCertificates field if non-nil, zero value otherwise.

### GetAnchorCertificatesOk

`func (o *PutComputerPrestageV3) GetAnchorCertificatesOk() (*[]string, bool)`

GetAnchorCertificatesOk returns a tuple with the AnchorCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorCertificates

`func (o *PutComputerPrestageV3) SetAnchorCertificates(v []string)`

SetAnchorCertificates sets AnchorCertificates field to given value.

### HasAnchorCertificates

`func (o *PutComputerPrestageV3) HasAnchorCertificates() bool`

HasAnchorCertificates returns a boolean if a field has been set.

### GetEnrollmentCustomizationId

`func (o *PutComputerPrestageV3) GetEnrollmentCustomizationId() string`

GetEnrollmentCustomizationId returns the EnrollmentCustomizationId field if non-nil, zero value otherwise.

### GetEnrollmentCustomizationIdOk

`func (o *PutComputerPrestageV3) GetEnrollmentCustomizationIdOk() (*string, bool)`

GetEnrollmentCustomizationIdOk returns a tuple with the EnrollmentCustomizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCustomizationId

`func (o *PutComputerPrestageV3) SetEnrollmentCustomizationId(v string)`

SetEnrollmentCustomizationId sets EnrollmentCustomizationId field to given value.

### HasEnrollmentCustomizationId

`func (o *PutComputerPrestageV3) HasEnrollmentCustomizationId() bool`

HasEnrollmentCustomizationId returns a boolean if a field has been set.

### GetLanguage

`func (o *PutComputerPrestageV3) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PutComputerPrestageV3) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PutComputerPrestageV3) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PutComputerPrestageV3) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetRegion

`func (o *PutComputerPrestageV3) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PutComputerPrestageV3) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PutComputerPrestageV3) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PutComputerPrestageV3) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetAutoAdvanceSetup

`func (o *PutComputerPrestageV3) GetAutoAdvanceSetup() bool`

GetAutoAdvanceSetup returns the AutoAdvanceSetup field if non-nil, zero value otherwise.

### GetAutoAdvanceSetupOk

`func (o *PutComputerPrestageV3) GetAutoAdvanceSetupOk() (*bool, bool)`

GetAutoAdvanceSetupOk returns a tuple with the AutoAdvanceSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAdvanceSetup

`func (o *PutComputerPrestageV3) SetAutoAdvanceSetup(v bool)`

SetAutoAdvanceSetup sets AutoAdvanceSetup field to given value.


### GetInstallProfilesDuringSetup

`func (o *PutComputerPrestageV3) GetInstallProfilesDuringSetup() bool`

GetInstallProfilesDuringSetup returns the InstallProfilesDuringSetup field if non-nil, zero value otherwise.

### GetInstallProfilesDuringSetupOk

`func (o *PutComputerPrestageV3) GetInstallProfilesDuringSetupOk() (*bool, bool)`

GetInstallProfilesDuringSetupOk returns a tuple with the InstallProfilesDuringSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallProfilesDuringSetup

`func (o *PutComputerPrestageV3) SetInstallProfilesDuringSetup(v bool)`

SetInstallProfilesDuringSetup sets InstallProfilesDuringSetup field to given value.


### GetPrestageInstalledProfileIds

`func (o *PutComputerPrestageV3) GetPrestageInstalledProfileIds() []string`

GetPrestageInstalledProfileIds returns the PrestageInstalledProfileIds field if non-nil, zero value otherwise.

### GetPrestageInstalledProfileIdsOk

`func (o *PutComputerPrestageV3) GetPrestageInstalledProfileIdsOk() (*[]string, bool)`

GetPrestageInstalledProfileIdsOk returns a tuple with the PrestageInstalledProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrestageInstalledProfileIds

`func (o *PutComputerPrestageV3) SetPrestageInstalledProfileIds(v []string)`

SetPrestageInstalledProfileIds sets PrestageInstalledProfileIds field to given value.


### GetCustomPackageIds

`func (o *PutComputerPrestageV3) GetCustomPackageIds() []string`

GetCustomPackageIds returns the CustomPackageIds field if non-nil, zero value otherwise.

### GetCustomPackageIdsOk

`func (o *PutComputerPrestageV3) GetCustomPackageIdsOk() (*[]string, bool)`

GetCustomPackageIdsOk returns a tuple with the CustomPackageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPackageIds

`func (o *PutComputerPrestageV3) SetCustomPackageIds(v []string)`

SetCustomPackageIds sets CustomPackageIds field to given value.


### GetCustomPackageDistributionPointId

`func (o *PutComputerPrestageV3) GetCustomPackageDistributionPointId() string`

GetCustomPackageDistributionPointId returns the CustomPackageDistributionPointId field if non-nil, zero value otherwise.

### GetCustomPackageDistributionPointIdOk

`func (o *PutComputerPrestageV3) GetCustomPackageDistributionPointIdOk() (*string, bool)`

GetCustomPackageDistributionPointIdOk returns a tuple with the CustomPackageDistributionPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPackageDistributionPointId

`func (o *PutComputerPrestageV3) SetCustomPackageDistributionPointId(v string)`

SetCustomPackageDistributionPointId sets CustomPackageDistributionPointId field to given value.


### GetEnableRecoveryLock

`func (o *PutComputerPrestageV3) GetEnableRecoveryLock() bool`

GetEnableRecoveryLock returns the EnableRecoveryLock field if non-nil, zero value otherwise.

### GetEnableRecoveryLockOk

`func (o *PutComputerPrestageV3) GetEnableRecoveryLockOk() (*bool, bool)`

GetEnableRecoveryLockOk returns a tuple with the EnableRecoveryLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRecoveryLock

`func (o *PutComputerPrestageV3) SetEnableRecoveryLock(v bool)`

SetEnableRecoveryLock sets EnableRecoveryLock field to given value.

### HasEnableRecoveryLock

`func (o *PutComputerPrestageV3) HasEnableRecoveryLock() bool`

HasEnableRecoveryLock returns a boolean if a field has been set.

### GetRecoveryLockPasswordType

`func (o *PutComputerPrestageV3) GetRecoveryLockPasswordType() string`

GetRecoveryLockPasswordType returns the RecoveryLockPasswordType field if non-nil, zero value otherwise.

### GetRecoveryLockPasswordTypeOk

`func (o *PutComputerPrestageV3) GetRecoveryLockPasswordTypeOk() (*string, bool)`

GetRecoveryLockPasswordTypeOk returns a tuple with the RecoveryLockPasswordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryLockPasswordType

`func (o *PutComputerPrestageV3) SetRecoveryLockPasswordType(v string)`

SetRecoveryLockPasswordType sets RecoveryLockPasswordType field to given value.

### HasRecoveryLockPasswordType

`func (o *PutComputerPrestageV3) HasRecoveryLockPasswordType() bool`

HasRecoveryLockPasswordType returns a boolean if a field has been set.

### GetRotateRecoveryLockPassword

`func (o *PutComputerPrestageV3) GetRotateRecoveryLockPassword() bool`

GetRotateRecoveryLockPassword returns the RotateRecoveryLockPassword field if non-nil, zero value otherwise.

### GetRotateRecoveryLockPasswordOk

`func (o *PutComputerPrestageV3) GetRotateRecoveryLockPasswordOk() (*bool, bool)`

GetRotateRecoveryLockPasswordOk returns a tuple with the RotateRecoveryLockPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateRecoveryLockPassword

`func (o *PutComputerPrestageV3) SetRotateRecoveryLockPassword(v bool)`

SetRotateRecoveryLockPassword sets RotateRecoveryLockPassword field to given value.

### HasRotateRecoveryLockPassword

`func (o *PutComputerPrestageV3) HasRotateRecoveryLockPassword() bool`

HasRotateRecoveryLockPassword returns a boolean if a field has been set.

### GetAccountSettings

`func (o *PutComputerPrestageV3) GetAccountSettings() AccountSettingsRequest`

GetAccountSettings returns the AccountSettings field if non-nil, zero value otherwise.

### GetAccountSettingsOk

`func (o *PutComputerPrestageV3) GetAccountSettingsOk() (*AccountSettingsRequest, bool)`

GetAccountSettingsOk returns a tuple with the AccountSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSettings

`func (o *PutComputerPrestageV3) SetAccountSettings(v AccountSettingsRequest)`

SetAccountSettings sets AccountSettings field to given value.

### HasAccountSettings

`func (o *PutComputerPrestageV3) HasAccountSettings() bool`

HasAccountSettings returns a boolean if a field has been set.

### GetRecoveryLockPassword

`func (o *PutComputerPrestageV3) GetRecoveryLockPassword() string`

GetRecoveryLockPassword returns the RecoveryLockPassword field if non-nil, zero value otherwise.

### GetRecoveryLockPasswordOk

`func (o *PutComputerPrestageV3) GetRecoveryLockPasswordOk() (*string, bool)`

GetRecoveryLockPasswordOk returns a tuple with the RecoveryLockPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryLockPassword

`func (o *PutComputerPrestageV3) SetRecoveryLockPassword(v string)`

SetRecoveryLockPassword sets RecoveryLockPassword field to given value.

### HasRecoveryLockPassword

`func (o *PutComputerPrestageV3) HasRecoveryLockPassword() bool`

HasRecoveryLockPassword returns a boolean if a field has been set.

### GetVersionLock

`func (o *PutComputerPrestageV3) GetVersionLock() int64`

GetVersionLock returns the VersionLock field if non-nil, zero value otherwise.

### GetVersionLockOk

`func (o *PutComputerPrestageV3) GetVersionLockOk() (*int64, bool)`

GetVersionLockOk returns a tuple with the VersionLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionLock

`func (o *PutComputerPrestageV3) SetVersionLock(v int64)`

SetVersionLock sets VersionLock field to given value.

### HasVersionLock

`func (o *PutComputerPrestageV3) HasVersionLock() bool`

HasVersionLock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


