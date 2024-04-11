# MobileDevicePrestageV2

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
**AllowPairing** | **bool** |  | 
**MultiUser** | **bool** |  | 
**Supervised** | **bool** |  | 
**MaximumSharedAccounts** | **int64** |  | 
**ConfigureDeviceBeforeSetupAssistant** | **bool** |  | 
**Names** | Pointer to [**MobileDevicePrestageNamesV2**](MobileDevicePrestageNamesV2.md) |  | [optional] 
**SendTimezone** | **bool** |  | 
**Timezone** | **string** |  | 
**StorageQuotaSizeMegabytes** | **int64** |  | 
**UseStorageQuotaSize** | **bool** |  | 
**TemporarySessionOnly** | Pointer to **bool** |  | [optional] 
**EnforceTemporarySessionTimeout** | Pointer to **bool** |  | [optional] 
**TemporarySessionTimeout** | Pointer to **int64** |  | [optional] 
**EnforceUserSessionTimeout** | Pointer to **bool** |  | [optional] 
**UserSessionTimeout** | Pointer to **int64** |  | [optional] 

## Methods

### NewMobileDevicePrestageV2

`func NewMobileDevicePrestageV2(displayName string, mandatory bool, mdmRemovable bool, supportPhoneNumber string, supportEmailAddress string, department string, defaultPrestage bool, enrollmentSiteId string, keepExistingSiteMembership bool, keepExistingLocationInformation bool, requireAuthentication bool, authenticationPrompt string, preventActivationLock bool, enableDeviceBasedActivationLock bool, deviceEnrollmentProgramInstanceId string, locationInformation LocationInformationV2, purchasingInformation PrestagePurchasingInformationV2, autoAdvanceSetup bool, allowPairing bool, multiUser bool, supervised bool, maximumSharedAccounts int64, configureDeviceBeforeSetupAssistant bool, sendTimezone bool, timezone string, storageQuotaSizeMegabytes int64, useStorageQuotaSize bool, ) *MobileDevicePrestageV2`

NewMobileDevicePrestageV2 instantiates a new MobileDevicePrestageV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDevicePrestageV2WithDefaults

`func NewMobileDevicePrestageV2WithDefaults() *MobileDevicePrestageV2`

NewMobileDevicePrestageV2WithDefaults instantiates a new MobileDevicePrestageV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *MobileDevicePrestageV2) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MobileDevicePrestageV2) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MobileDevicePrestageV2) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMandatory

`func (o *MobileDevicePrestageV2) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *MobileDevicePrestageV2) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *MobileDevicePrestageV2) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.


### GetMdmRemovable

`func (o *MobileDevicePrestageV2) GetMdmRemovable() bool`

GetMdmRemovable returns the MdmRemovable field if non-nil, zero value otherwise.

### GetMdmRemovableOk

`func (o *MobileDevicePrestageV2) GetMdmRemovableOk() (*bool, bool)`

GetMdmRemovableOk returns a tuple with the MdmRemovable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmRemovable

`func (o *MobileDevicePrestageV2) SetMdmRemovable(v bool)`

SetMdmRemovable sets MdmRemovable field to given value.


### GetSupportPhoneNumber

`func (o *MobileDevicePrestageV2) GetSupportPhoneNumber() string`

GetSupportPhoneNumber returns the SupportPhoneNumber field if non-nil, zero value otherwise.

### GetSupportPhoneNumberOk

`func (o *MobileDevicePrestageV2) GetSupportPhoneNumberOk() (*string, bool)`

GetSupportPhoneNumberOk returns a tuple with the SupportPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPhoneNumber

`func (o *MobileDevicePrestageV2) SetSupportPhoneNumber(v string)`

SetSupportPhoneNumber sets SupportPhoneNumber field to given value.


### GetSupportEmailAddress

`func (o *MobileDevicePrestageV2) GetSupportEmailAddress() string`

GetSupportEmailAddress returns the SupportEmailAddress field if non-nil, zero value otherwise.

### GetSupportEmailAddressOk

`func (o *MobileDevicePrestageV2) GetSupportEmailAddressOk() (*string, bool)`

GetSupportEmailAddressOk returns a tuple with the SupportEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmailAddress

`func (o *MobileDevicePrestageV2) SetSupportEmailAddress(v string)`

SetSupportEmailAddress sets SupportEmailAddress field to given value.


### GetDepartment

`func (o *MobileDevicePrestageV2) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *MobileDevicePrestageV2) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *MobileDevicePrestageV2) SetDepartment(v string)`

SetDepartment sets Department field to given value.


### GetDefaultPrestage

`func (o *MobileDevicePrestageV2) GetDefaultPrestage() bool`

GetDefaultPrestage returns the DefaultPrestage field if non-nil, zero value otherwise.

### GetDefaultPrestageOk

`func (o *MobileDevicePrestageV2) GetDefaultPrestageOk() (*bool, bool)`

GetDefaultPrestageOk returns a tuple with the DefaultPrestage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrestage

`func (o *MobileDevicePrestageV2) SetDefaultPrestage(v bool)`

SetDefaultPrestage sets DefaultPrestage field to given value.


### GetEnrollmentSiteId

`func (o *MobileDevicePrestageV2) GetEnrollmentSiteId() string`

GetEnrollmentSiteId returns the EnrollmentSiteId field if non-nil, zero value otherwise.

### GetEnrollmentSiteIdOk

`func (o *MobileDevicePrestageV2) GetEnrollmentSiteIdOk() (*string, bool)`

GetEnrollmentSiteIdOk returns a tuple with the EnrollmentSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSiteId

`func (o *MobileDevicePrestageV2) SetEnrollmentSiteId(v string)`

SetEnrollmentSiteId sets EnrollmentSiteId field to given value.


### GetKeepExistingSiteMembership

`func (o *MobileDevicePrestageV2) GetKeepExistingSiteMembership() bool`

GetKeepExistingSiteMembership returns the KeepExistingSiteMembership field if non-nil, zero value otherwise.

### GetKeepExistingSiteMembershipOk

`func (o *MobileDevicePrestageV2) GetKeepExistingSiteMembershipOk() (*bool, bool)`

GetKeepExistingSiteMembershipOk returns a tuple with the KeepExistingSiteMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepExistingSiteMembership

`func (o *MobileDevicePrestageV2) SetKeepExistingSiteMembership(v bool)`

SetKeepExistingSiteMembership sets KeepExistingSiteMembership field to given value.


### GetKeepExistingLocationInformation

`func (o *MobileDevicePrestageV2) GetKeepExistingLocationInformation() bool`

GetKeepExistingLocationInformation returns the KeepExistingLocationInformation field if non-nil, zero value otherwise.

### GetKeepExistingLocationInformationOk

`func (o *MobileDevicePrestageV2) GetKeepExistingLocationInformationOk() (*bool, bool)`

GetKeepExistingLocationInformationOk returns a tuple with the KeepExistingLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepExistingLocationInformation

`func (o *MobileDevicePrestageV2) SetKeepExistingLocationInformation(v bool)`

SetKeepExistingLocationInformation sets KeepExistingLocationInformation field to given value.


### GetRequireAuthentication

`func (o *MobileDevicePrestageV2) GetRequireAuthentication() bool`

GetRequireAuthentication returns the RequireAuthentication field if non-nil, zero value otherwise.

### GetRequireAuthenticationOk

`func (o *MobileDevicePrestageV2) GetRequireAuthenticationOk() (*bool, bool)`

GetRequireAuthenticationOk returns a tuple with the RequireAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuthentication

`func (o *MobileDevicePrestageV2) SetRequireAuthentication(v bool)`

SetRequireAuthentication sets RequireAuthentication field to given value.


### GetAuthenticationPrompt

`func (o *MobileDevicePrestageV2) GetAuthenticationPrompt() string`

GetAuthenticationPrompt returns the AuthenticationPrompt field if non-nil, zero value otherwise.

### GetAuthenticationPromptOk

`func (o *MobileDevicePrestageV2) GetAuthenticationPromptOk() (*string, bool)`

GetAuthenticationPromptOk returns a tuple with the AuthenticationPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPrompt

`func (o *MobileDevicePrestageV2) SetAuthenticationPrompt(v string)`

SetAuthenticationPrompt sets AuthenticationPrompt field to given value.


### GetPreventActivationLock

`func (o *MobileDevicePrestageV2) GetPreventActivationLock() bool`

GetPreventActivationLock returns the PreventActivationLock field if non-nil, zero value otherwise.

### GetPreventActivationLockOk

`func (o *MobileDevicePrestageV2) GetPreventActivationLockOk() (*bool, bool)`

GetPreventActivationLockOk returns a tuple with the PreventActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventActivationLock

`func (o *MobileDevicePrestageV2) SetPreventActivationLock(v bool)`

SetPreventActivationLock sets PreventActivationLock field to given value.


### GetEnableDeviceBasedActivationLock

`func (o *MobileDevicePrestageV2) GetEnableDeviceBasedActivationLock() bool`

GetEnableDeviceBasedActivationLock returns the EnableDeviceBasedActivationLock field if non-nil, zero value otherwise.

### GetEnableDeviceBasedActivationLockOk

`func (o *MobileDevicePrestageV2) GetEnableDeviceBasedActivationLockOk() (*bool, bool)`

GetEnableDeviceBasedActivationLockOk returns a tuple with the EnableDeviceBasedActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDeviceBasedActivationLock

`func (o *MobileDevicePrestageV2) SetEnableDeviceBasedActivationLock(v bool)`

SetEnableDeviceBasedActivationLock sets EnableDeviceBasedActivationLock field to given value.


### GetDeviceEnrollmentProgramInstanceId

`func (o *MobileDevicePrestageV2) GetDeviceEnrollmentProgramInstanceId() string`

GetDeviceEnrollmentProgramInstanceId returns the DeviceEnrollmentProgramInstanceId field if non-nil, zero value otherwise.

### GetDeviceEnrollmentProgramInstanceIdOk

`func (o *MobileDevicePrestageV2) GetDeviceEnrollmentProgramInstanceIdOk() (*string, bool)`

GetDeviceEnrollmentProgramInstanceIdOk returns a tuple with the DeviceEnrollmentProgramInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceEnrollmentProgramInstanceId

`func (o *MobileDevicePrestageV2) SetDeviceEnrollmentProgramInstanceId(v string)`

SetDeviceEnrollmentProgramInstanceId sets DeviceEnrollmentProgramInstanceId field to given value.


### GetSkipSetupItems

`func (o *MobileDevicePrestageV2) GetSkipSetupItems() map[string]bool`

GetSkipSetupItems returns the SkipSetupItems field if non-nil, zero value otherwise.

### GetSkipSetupItemsOk

`func (o *MobileDevicePrestageV2) GetSkipSetupItemsOk() (*map[string]bool, bool)`

GetSkipSetupItemsOk returns a tuple with the SkipSetupItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSetupItems

`func (o *MobileDevicePrestageV2) SetSkipSetupItems(v map[string]bool)`

SetSkipSetupItems sets SkipSetupItems field to given value.

### HasSkipSetupItems

`func (o *MobileDevicePrestageV2) HasSkipSetupItems() bool`

HasSkipSetupItems returns a boolean if a field has been set.

### GetLocationInformation

`func (o *MobileDevicePrestageV2) GetLocationInformation() LocationInformationV2`

GetLocationInformation returns the LocationInformation field if non-nil, zero value otherwise.

### GetLocationInformationOk

`func (o *MobileDevicePrestageV2) GetLocationInformationOk() (*LocationInformationV2, bool)`

GetLocationInformationOk returns a tuple with the LocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationInformation

`func (o *MobileDevicePrestageV2) SetLocationInformation(v LocationInformationV2)`

SetLocationInformation sets LocationInformation field to given value.


### GetPurchasingInformation

`func (o *MobileDevicePrestageV2) GetPurchasingInformation() PrestagePurchasingInformationV2`

GetPurchasingInformation returns the PurchasingInformation field if non-nil, zero value otherwise.

### GetPurchasingInformationOk

`func (o *MobileDevicePrestageV2) GetPurchasingInformationOk() (*PrestagePurchasingInformationV2, bool)`

GetPurchasingInformationOk returns a tuple with the PurchasingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingInformation

`func (o *MobileDevicePrestageV2) SetPurchasingInformation(v PrestagePurchasingInformationV2)`

SetPurchasingInformation sets PurchasingInformation field to given value.


### GetAnchorCertificates

`func (o *MobileDevicePrestageV2) GetAnchorCertificates() []string`

GetAnchorCertificates returns the AnchorCertificates field if non-nil, zero value otherwise.

### GetAnchorCertificatesOk

`func (o *MobileDevicePrestageV2) GetAnchorCertificatesOk() (*[]string, bool)`

GetAnchorCertificatesOk returns a tuple with the AnchorCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorCertificates

`func (o *MobileDevicePrestageV2) SetAnchorCertificates(v []string)`

SetAnchorCertificates sets AnchorCertificates field to given value.

### HasAnchorCertificates

`func (o *MobileDevicePrestageV2) HasAnchorCertificates() bool`

HasAnchorCertificates returns a boolean if a field has been set.

### GetEnrollmentCustomizationId

`func (o *MobileDevicePrestageV2) GetEnrollmentCustomizationId() string`

GetEnrollmentCustomizationId returns the EnrollmentCustomizationId field if non-nil, zero value otherwise.

### GetEnrollmentCustomizationIdOk

`func (o *MobileDevicePrestageV2) GetEnrollmentCustomizationIdOk() (*string, bool)`

GetEnrollmentCustomizationIdOk returns a tuple with the EnrollmentCustomizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCustomizationId

`func (o *MobileDevicePrestageV2) SetEnrollmentCustomizationId(v string)`

SetEnrollmentCustomizationId sets EnrollmentCustomizationId field to given value.

### HasEnrollmentCustomizationId

`func (o *MobileDevicePrestageV2) HasEnrollmentCustomizationId() bool`

HasEnrollmentCustomizationId returns a boolean if a field has been set.

### GetLanguage

`func (o *MobileDevicePrestageV2) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MobileDevicePrestageV2) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *MobileDevicePrestageV2) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *MobileDevicePrestageV2) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetRegion

`func (o *MobileDevicePrestageV2) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MobileDevicePrestageV2) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MobileDevicePrestageV2) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MobileDevicePrestageV2) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetAutoAdvanceSetup

`func (o *MobileDevicePrestageV2) GetAutoAdvanceSetup() bool`

GetAutoAdvanceSetup returns the AutoAdvanceSetup field if non-nil, zero value otherwise.

### GetAutoAdvanceSetupOk

`func (o *MobileDevicePrestageV2) GetAutoAdvanceSetupOk() (*bool, bool)`

GetAutoAdvanceSetupOk returns a tuple with the AutoAdvanceSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAdvanceSetup

`func (o *MobileDevicePrestageV2) SetAutoAdvanceSetup(v bool)`

SetAutoAdvanceSetup sets AutoAdvanceSetup field to given value.


### GetAllowPairing

`func (o *MobileDevicePrestageV2) GetAllowPairing() bool`

GetAllowPairing returns the AllowPairing field if non-nil, zero value otherwise.

### GetAllowPairingOk

`func (o *MobileDevicePrestageV2) GetAllowPairingOk() (*bool, bool)`

GetAllowPairingOk returns a tuple with the AllowPairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPairing

`func (o *MobileDevicePrestageV2) SetAllowPairing(v bool)`

SetAllowPairing sets AllowPairing field to given value.


### GetMultiUser

`func (o *MobileDevicePrestageV2) GetMultiUser() bool`

GetMultiUser returns the MultiUser field if non-nil, zero value otherwise.

### GetMultiUserOk

`func (o *MobileDevicePrestageV2) GetMultiUserOk() (*bool, bool)`

GetMultiUserOk returns a tuple with the MultiUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiUser

`func (o *MobileDevicePrestageV2) SetMultiUser(v bool)`

SetMultiUser sets MultiUser field to given value.


### GetSupervised

`func (o *MobileDevicePrestageV2) GetSupervised() bool`

GetSupervised returns the Supervised field if non-nil, zero value otherwise.

### GetSupervisedOk

`func (o *MobileDevicePrestageV2) GetSupervisedOk() (*bool, bool)`

GetSupervisedOk returns a tuple with the Supervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervised

`func (o *MobileDevicePrestageV2) SetSupervised(v bool)`

SetSupervised sets Supervised field to given value.


### GetMaximumSharedAccounts

`func (o *MobileDevicePrestageV2) GetMaximumSharedAccounts() int64`

GetMaximumSharedAccounts returns the MaximumSharedAccounts field if non-nil, zero value otherwise.

### GetMaximumSharedAccountsOk

`func (o *MobileDevicePrestageV2) GetMaximumSharedAccountsOk() (*int64, bool)`

GetMaximumSharedAccountsOk returns a tuple with the MaximumSharedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSharedAccounts

`func (o *MobileDevicePrestageV2) SetMaximumSharedAccounts(v int64)`

SetMaximumSharedAccounts sets MaximumSharedAccounts field to given value.


### GetConfigureDeviceBeforeSetupAssistant

`func (o *MobileDevicePrestageV2) GetConfigureDeviceBeforeSetupAssistant() bool`

GetConfigureDeviceBeforeSetupAssistant returns the ConfigureDeviceBeforeSetupAssistant field if non-nil, zero value otherwise.

### GetConfigureDeviceBeforeSetupAssistantOk

`func (o *MobileDevicePrestageV2) GetConfigureDeviceBeforeSetupAssistantOk() (*bool, bool)`

GetConfigureDeviceBeforeSetupAssistantOk returns a tuple with the ConfigureDeviceBeforeSetupAssistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureDeviceBeforeSetupAssistant

`func (o *MobileDevicePrestageV2) SetConfigureDeviceBeforeSetupAssistant(v bool)`

SetConfigureDeviceBeforeSetupAssistant sets ConfigureDeviceBeforeSetupAssistant field to given value.


### GetNames

`func (o *MobileDevicePrestageV2) GetNames() MobileDevicePrestageNamesV2`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *MobileDevicePrestageV2) GetNamesOk() (*MobileDevicePrestageNamesV2, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *MobileDevicePrestageV2) SetNames(v MobileDevicePrestageNamesV2)`

SetNames sets Names field to given value.

### HasNames

`func (o *MobileDevicePrestageV2) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetSendTimezone

`func (o *MobileDevicePrestageV2) GetSendTimezone() bool`

GetSendTimezone returns the SendTimezone field if non-nil, zero value otherwise.

### GetSendTimezoneOk

`func (o *MobileDevicePrestageV2) GetSendTimezoneOk() (*bool, bool)`

GetSendTimezoneOk returns a tuple with the SendTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendTimezone

`func (o *MobileDevicePrestageV2) SetSendTimezone(v bool)`

SetSendTimezone sets SendTimezone field to given value.


### GetTimezone

`func (o *MobileDevicePrestageV2) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *MobileDevicePrestageV2) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *MobileDevicePrestageV2) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetStorageQuotaSizeMegabytes

`func (o *MobileDevicePrestageV2) GetStorageQuotaSizeMegabytes() int64`

GetStorageQuotaSizeMegabytes returns the StorageQuotaSizeMegabytes field if non-nil, zero value otherwise.

### GetStorageQuotaSizeMegabytesOk

`func (o *MobileDevicePrestageV2) GetStorageQuotaSizeMegabytesOk() (*int64, bool)`

GetStorageQuotaSizeMegabytesOk returns a tuple with the StorageQuotaSizeMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageQuotaSizeMegabytes

`func (o *MobileDevicePrestageV2) SetStorageQuotaSizeMegabytes(v int64)`

SetStorageQuotaSizeMegabytes sets StorageQuotaSizeMegabytes field to given value.


### GetUseStorageQuotaSize

`func (o *MobileDevicePrestageV2) GetUseStorageQuotaSize() bool`

GetUseStorageQuotaSize returns the UseStorageQuotaSize field if non-nil, zero value otherwise.

### GetUseStorageQuotaSizeOk

`func (o *MobileDevicePrestageV2) GetUseStorageQuotaSizeOk() (*bool, bool)`

GetUseStorageQuotaSizeOk returns a tuple with the UseStorageQuotaSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseStorageQuotaSize

`func (o *MobileDevicePrestageV2) SetUseStorageQuotaSize(v bool)`

SetUseStorageQuotaSize sets UseStorageQuotaSize field to given value.


### GetTemporarySessionOnly

`func (o *MobileDevicePrestageV2) GetTemporarySessionOnly() bool`

GetTemporarySessionOnly returns the TemporarySessionOnly field if non-nil, zero value otherwise.

### GetTemporarySessionOnlyOk

`func (o *MobileDevicePrestageV2) GetTemporarySessionOnlyOk() (*bool, bool)`

GetTemporarySessionOnlyOk returns a tuple with the TemporarySessionOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporarySessionOnly

`func (o *MobileDevicePrestageV2) SetTemporarySessionOnly(v bool)`

SetTemporarySessionOnly sets TemporarySessionOnly field to given value.

### HasTemporarySessionOnly

`func (o *MobileDevicePrestageV2) HasTemporarySessionOnly() bool`

HasTemporarySessionOnly returns a boolean if a field has been set.

### GetEnforceTemporarySessionTimeout

`func (o *MobileDevicePrestageV2) GetEnforceTemporarySessionTimeout() bool`

GetEnforceTemporarySessionTimeout returns the EnforceTemporarySessionTimeout field if non-nil, zero value otherwise.

### GetEnforceTemporarySessionTimeoutOk

`func (o *MobileDevicePrestageV2) GetEnforceTemporarySessionTimeoutOk() (*bool, bool)`

GetEnforceTemporarySessionTimeoutOk returns a tuple with the EnforceTemporarySessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceTemporarySessionTimeout

`func (o *MobileDevicePrestageV2) SetEnforceTemporarySessionTimeout(v bool)`

SetEnforceTemporarySessionTimeout sets EnforceTemporarySessionTimeout field to given value.

### HasEnforceTemporarySessionTimeout

`func (o *MobileDevicePrestageV2) HasEnforceTemporarySessionTimeout() bool`

HasEnforceTemporarySessionTimeout returns a boolean if a field has been set.

### GetTemporarySessionTimeout

`func (o *MobileDevicePrestageV2) GetTemporarySessionTimeout() int64`

GetTemporarySessionTimeout returns the TemporarySessionTimeout field if non-nil, zero value otherwise.

### GetTemporarySessionTimeoutOk

`func (o *MobileDevicePrestageV2) GetTemporarySessionTimeoutOk() (*int64, bool)`

GetTemporarySessionTimeoutOk returns a tuple with the TemporarySessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporarySessionTimeout

`func (o *MobileDevicePrestageV2) SetTemporarySessionTimeout(v int64)`

SetTemporarySessionTimeout sets TemporarySessionTimeout field to given value.

### HasTemporarySessionTimeout

`func (o *MobileDevicePrestageV2) HasTemporarySessionTimeout() bool`

HasTemporarySessionTimeout returns a boolean if a field has been set.

### GetEnforceUserSessionTimeout

`func (o *MobileDevicePrestageV2) GetEnforceUserSessionTimeout() bool`

GetEnforceUserSessionTimeout returns the EnforceUserSessionTimeout field if non-nil, zero value otherwise.

### GetEnforceUserSessionTimeoutOk

`func (o *MobileDevicePrestageV2) GetEnforceUserSessionTimeoutOk() (*bool, bool)`

GetEnforceUserSessionTimeoutOk returns a tuple with the EnforceUserSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceUserSessionTimeout

`func (o *MobileDevicePrestageV2) SetEnforceUserSessionTimeout(v bool)`

SetEnforceUserSessionTimeout sets EnforceUserSessionTimeout field to given value.

### HasEnforceUserSessionTimeout

`func (o *MobileDevicePrestageV2) HasEnforceUserSessionTimeout() bool`

HasEnforceUserSessionTimeout returns a boolean if a field has been set.

### GetUserSessionTimeout

`func (o *MobileDevicePrestageV2) GetUserSessionTimeout() int64`

GetUserSessionTimeout returns the UserSessionTimeout field if non-nil, zero value otherwise.

### GetUserSessionTimeoutOk

`func (o *MobileDevicePrestageV2) GetUserSessionTimeoutOk() (*int64, bool)`

GetUserSessionTimeoutOk returns a tuple with the UserSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSessionTimeout

`func (o *MobileDevicePrestageV2) SetUserSessionTimeout(v int64)`

SetUserSessionTimeout sets UserSessionTimeout field to given value.

### HasUserSessionTimeout

`func (o *MobileDevicePrestageV2) HasUserSessionTimeout() bool`

HasUserSessionTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


