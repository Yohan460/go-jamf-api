# GetMobileDevicePrestageV2

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
**Id** | Pointer to **string** |  | [optional] 
**ProfileUuid** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**VersionLock** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetMobileDevicePrestageV2

`func NewGetMobileDevicePrestageV2(displayName string, mandatory bool, mdmRemovable bool, supportPhoneNumber string, supportEmailAddress string, department string, defaultPrestage bool, enrollmentSiteId string, keepExistingSiteMembership bool, keepExistingLocationInformation bool, requireAuthentication bool, authenticationPrompt string, preventActivationLock bool, enableDeviceBasedActivationLock bool, deviceEnrollmentProgramInstanceId string, locationInformation LocationInformationV2, purchasingInformation PrestagePurchasingInformationV2, autoAdvanceSetup bool, allowPairing bool, multiUser bool, supervised bool, maximumSharedAccounts int64, configureDeviceBeforeSetupAssistant bool, sendTimezone bool, timezone string, storageQuotaSizeMegabytes int64, useStorageQuotaSize bool, ) *GetMobileDevicePrestageV2`

NewGetMobileDevicePrestageV2 instantiates a new GetMobileDevicePrestageV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMobileDevicePrestageV2WithDefaults

`func NewGetMobileDevicePrestageV2WithDefaults() *GetMobileDevicePrestageV2`

NewGetMobileDevicePrestageV2WithDefaults instantiates a new GetMobileDevicePrestageV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *GetMobileDevicePrestageV2) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetMobileDevicePrestageV2) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetMobileDevicePrestageV2) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMandatory

`func (o *GetMobileDevicePrestageV2) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *GetMobileDevicePrestageV2) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *GetMobileDevicePrestageV2) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.


### GetMdmRemovable

`func (o *GetMobileDevicePrestageV2) GetMdmRemovable() bool`

GetMdmRemovable returns the MdmRemovable field if non-nil, zero value otherwise.

### GetMdmRemovableOk

`func (o *GetMobileDevicePrestageV2) GetMdmRemovableOk() (*bool, bool)`

GetMdmRemovableOk returns a tuple with the MdmRemovable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmRemovable

`func (o *GetMobileDevicePrestageV2) SetMdmRemovable(v bool)`

SetMdmRemovable sets MdmRemovable field to given value.


### GetSupportPhoneNumber

`func (o *GetMobileDevicePrestageV2) GetSupportPhoneNumber() string`

GetSupportPhoneNumber returns the SupportPhoneNumber field if non-nil, zero value otherwise.

### GetSupportPhoneNumberOk

`func (o *GetMobileDevicePrestageV2) GetSupportPhoneNumberOk() (*string, bool)`

GetSupportPhoneNumberOk returns a tuple with the SupportPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPhoneNumber

`func (o *GetMobileDevicePrestageV2) SetSupportPhoneNumber(v string)`

SetSupportPhoneNumber sets SupportPhoneNumber field to given value.


### GetSupportEmailAddress

`func (o *GetMobileDevicePrestageV2) GetSupportEmailAddress() string`

GetSupportEmailAddress returns the SupportEmailAddress field if non-nil, zero value otherwise.

### GetSupportEmailAddressOk

`func (o *GetMobileDevicePrestageV2) GetSupportEmailAddressOk() (*string, bool)`

GetSupportEmailAddressOk returns a tuple with the SupportEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmailAddress

`func (o *GetMobileDevicePrestageV2) SetSupportEmailAddress(v string)`

SetSupportEmailAddress sets SupportEmailAddress field to given value.


### GetDepartment

`func (o *GetMobileDevicePrestageV2) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *GetMobileDevicePrestageV2) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *GetMobileDevicePrestageV2) SetDepartment(v string)`

SetDepartment sets Department field to given value.


### GetDefaultPrestage

`func (o *GetMobileDevicePrestageV2) GetDefaultPrestage() bool`

GetDefaultPrestage returns the DefaultPrestage field if non-nil, zero value otherwise.

### GetDefaultPrestageOk

`func (o *GetMobileDevicePrestageV2) GetDefaultPrestageOk() (*bool, bool)`

GetDefaultPrestageOk returns a tuple with the DefaultPrestage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrestage

`func (o *GetMobileDevicePrestageV2) SetDefaultPrestage(v bool)`

SetDefaultPrestage sets DefaultPrestage field to given value.


### GetEnrollmentSiteId

`func (o *GetMobileDevicePrestageV2) GetEnrollmentSiteId() string`

GetEnrollmentSiteId returns the EnrollmentSiteId field if non-nil, zero value otherwise.

### GetEnrollmentSiteIdOk

`func (o *GetMobileDevicePrestageV2) GetEnrollmentSiteIdOk() (*string, bool)`

GetEnrollmentSiteIdOk returns a tuple with the EnrollmentSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSiteId

`func (o *GetMobileDevicePrestageV2) SetEnrollmentSiteId(v string)`

SetEnrollmentSiteId sets EnrollmentSiteId field to given value.


### GetKeepExistingSiteMembership

`func (o *GetMobileDevicePrestageV2) GetKeepExistingSiteMembership() bool`

GetKeepExistingSiteMembership returns the KeepExistingSiteMembership field if non-nil, zero value otherwise.

### GetKeepExistingSiteMembershipOk

`func (o *GetMobileDevicePrestageV2) GetKeepExistingSiteMembershipOk() (*bool, bool)`

GetKeepExistingSiteMembershipOk returns a tuple with the KeepExistingSiteMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepExistingSiteMembership

`func (o *GetMobileDevicePrestageV2) SetKeepExistingSiteMembership(v bool)`

SetKeepExistingSiteMembership sets KeepExistingSiteMembership field to given value.


### GetKeepExistingLocationInformation

`func (o *GetMobileDevicePrestageV2) GetKeepExistingLocationInformation() bool`

GetKeepExistingLocationInformation returns the KeepExistingLocationInformation field if non-nil, zero value otherwise.

### GetKeepExistingLocationInformationOk

`func (o *GetMobileDevicePrestageV2) GetKeepExistingLocationInformationOk() (*bool, bool)`

GetKeepExistingLocationInformationOk returns a tuple with the KeepExistingLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepExistingLocationInformation

`func (o *GetMobileDevicePrestageV2) SetKeepExistingLocationInformation(v bool)`

SetKeepExistingLocationInformation sets KeepExistingLocationInformation field to given value.


### GetRequireAuthentication

`func (o *GetMobileDevicePrestageV2) GetRequireAuthentication() bool`

GetRequireAuthentication returns the RequireAuthentication field if non-nil, zero value otherwise.

### GetRequireAuthenticationOk

`func (o *GetMobileDevicePrestageV2) GetRequireAuthenticationOk() (*bool, bool)`

GetRequireAuthenticationOk returns a tuple with the RequireAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuthentication

`func (o *GetMobileDevicePrestageV2) SetRequireAuthentication(v bool)`

SetRequireAuthentication sets RequireAuthentication field to given value.


### GetAuthenticationPrompt

`func (o *GetMobileDevicePrestageV2) GetAuthenticationPrompt() string`

GetAuthenticationPrompt returns the AuthenticationPrompt field if non-nil, zero value otherwise.

### GetAuthenticationPromptOk

`func (o *GetMobileDevicePrestageV2) GetAuthenticationPromptOk() (*string, bool)`

GetAuthenticationPromptOk returns a tuple with the AuthenticationPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPrompt

`func (o *GetMobileDevicePrestageV2) SetAuthenticationPrompt(v string)`

SetAuthenticationPrompt sets AuthenticationPrompt field to given value.


### GetPreventActivationLock

`func (o *GetMobileDevicePrestageV2) GetPreventActivationLock() bool`

GetPreventActivationLock returns the PreventActivationLock field if non-nil, zero value otherwise.

### GetPreventActivationLockOk

`func (o *GetMobileDevicePrestageV2) GetPreventActivationLockOk() (*bool, bool)`

GetPreventActivationLockOk returns a tuple with the PreventActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventActivationLock

`func (o *GetMobileDevicePrestageV2) SetPreventActivationLock(v bool)`

SetPreventActivationLock sets PreventActivationLock field to given value.


### GetEnableDeviceBasedActivationLock

`func (o *GetMobileDevicePrestageV2) GetEnableDeviceBasedActivationLock() bool`

GetEnableDeviceBasedActivationLock returns the EnableDeviceBasedActivationLock field if non-nil, zero value otherwise.

### GetEnableDeviceBasedActivationLockOk

`func (o *GetMobileDevicePrestageV2) GetEnableDeviceBasedActivationLockOk() (*bool, bool)`

GetEnableDeviceBasedActivationLockOk returns a tuple with the EnableDeviceBasedActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDeviceBasedActivationLock

`func (o *GetMobileDevicePrestageV2) SetEnableDeviceBasedActivationLock(v bool)`

SetEnableDeviceBasedActivationLock sets EnableDeviceBasedActivationLock field to given value.


### GetDeviceEnrollmentProgramInstanceId

`func (o *GetMobileDevicePrestageV2) GetDeviceEnrollmentProgramInstanceId() string`

GetDeviceEnrollmentProgramInstanceId returns the DeviceEnrollmentProgramInstanceId field if non-nil, zero value otherwise.

### GetDeviceEnrollmentProgramInstanceIdOk

`func (o *GetMobileDevicePrestageV2) GetDeviceEnrollmentProgramInstanceIdOk() (*string, bool)`

GetDeviceEnrollmentProgramInstanceIdOk returns a tuple with the DeviceEnrollmentProgramInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceEnrollmentProgramInstanceId

`func (o *GetMobileDevicePrestageV2) SetDeviceEnrollmentProgramInstanceId(v string)`

SetDeviceEnrollmentProgramInstanceId sets DeviceEnrollmentProgramInstanceId field to given value.


### GetSkipSetupItems

`func (o *GetMobileDevicePrestageV2) GetSkipSetupItems() map[string]bool`

GetSkipSetupItems returns the SkipSetupItems field if non-nil, zero value otherwise.

### GetSkipSetupItemsOk

`func (o *GetMobileDevicePrestageV2) GetSkipSetupItemsOk() (*map[string]bool, bool)`

GetSkipSetupItemsOk returns a tuple with the SkipSetupItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSetupItems

`func (o *GetMobileDevicePrestageV2) SetSkipSetupItems(v map[string]bool)`

SetSkipSetupItems sets SkipSetupItems field to given value.

### HasSkipSetupItems

`func (o *GetMobileDevicePrestageV2) HasSkipSetupItems() bool`

HasSkipSetupItems returns a boolean if a field has been set.

### GetLocationInformation

`func (o *GetMobileDevicePrestageV2) GetLocationInformation() LocationInformationV2`

GetLocationInformation returns the LocationInformation field if non-nil, zero value otherwise.

### GetLocationInformationOk

`func (o *GetMobileDevicePrestageV2) GetLocationInformationOk() (*LocationInformationV2, bool)`

GetLocationInformationOk returns a tuple with the LocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationInformation

`func (o *GetMobileDevicePrestageV2) SetLocationInformation(v LocationInformationV2)`

SetLocationInformation sets LocationInformation field to given value.


### GetPurchasingInformation

`func (o *GetMobileDevicePrestageV2) GetPurchasingInformation() PrestagePurchasingInformationV2`

GetPurchasingInformation returns the PurchasingInformation field if non-nil, zero value otherwise.

### GetPurchasingInformationOk

`func (o *GetMobileDevicePrestageV2) GetPurchasingInformationOk() (*PrestagePurchasingInformationV2, bool)`

GetPurchasingInformationOk returns a tuple with the PurchasingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingInformation

`func (o *GetMobileDevicePrestageV2) SetPurchasingInformation(v PrestagePurchasingInformationV2)`

SetPurchasingInformation sets PurchasingInformation field to given value.


### GetAnchorCertificates

`func (o *GetMobileDevicePrestageV2) GetAnchorCertificates() []string`

GetAnchorCertificates returns the AnchorCertificates field if non-nil, zero value otherwise.

### GetAnchorCertificatesOk

`func (o *GetMobileDevicePrestageV2) GetAnchorCertificatesOk() (*[]string, bool)`

GetAnchorCertificatesOk returns a tuple with the AnchorCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorCertificates

`func (o *GetMobileDevicePrestageV2) SetAnchorCertificates(v []string)`

SetAnchorCertificates sets AnchorCertificates field to given value.

### HasAnchorCertificates

`func (o *GetMobileDevicePrestageV2) HasAnchorCertificates() bool`

HasAnchorCertificates returns a boolean if a field has been set.

### GetEnrollmentCustomizationId

`func (o *GetMobileDevicePrestageV2) GetEnrollmentCustomizationId() string`

GetEnrollmentCustomizationId returns the EnrollmentCustomizationId field if non-nil, zero value otherwise.

### GetEnrollmentCustomizationIdOk

`func (o *GetMobileDevicePrestageV2) GetEnrollmentCustomizationIdOk() (*string, bool)`

GetEnrollmentCustomizationIdOk returns a tuple with the EnrollmentCustomizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCustomizationId

`func (o *GetMobileDevicePrestageV2) SetEnrollmentCustomizationId(v string)`

SetEnrollmentCustomizationId sets EnrollmentCustomizationId field to given value.

### HasEnrollmentCustomizationId

`func (o *GetMobileDevicePrestageV2) HasEnrollmentCustomizationId() bool`

HasEnrollmentCustomizationId returns a boolean if a field has been set.

### GetLanguage

`func (o *GetMobileDevicePrestageV2) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GetMobileDevicePrestageV2) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GetMobileDevicePrestageV2) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *GetMobileDevicePrestageV2) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetRegion

`func (o *GetMobileDevicePrestageV2) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetMobileDevicePrestageV2) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetMobileDevicePrestageV2) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetMobileDevicePrestageV2) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetAutoAdvanceSetup

`func (o *GetMobileDevicePrestageV2) GetAutoAdvanceSetup() bool`

GetAutoAdvanceSetup returns the AutoAdvanceSetup field if non-nil, zero value otherwise.

### GetAutoAdvanceSetupOk

`func (o *GetMobileDevicePrestageV2) GetAutoAdvanceSetupOk() (*bool, bool)`

GetAutoAdvanceSetupOk returns a tuple with the AutoAdvanceSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAdvanceSetup

`func (o *GetMobileDevicePrestageV2) SetAutoAdvanceSetup(v bool)`

SetAutoAdvanceSetup sets AutoAdvanceSetup field to given value.


### GetAllowPairing

`func (o *GetMobileDevicePrestageV2) GetAllowPairing() bool`

GetAllowPairing returns the AllowPairing field if non-nil, zero value otherwise.

### GetAllowPairingOk

`func (o *GetMobileDevicePrestageV2) GetAllowPairingOk() (*bool, bool)`

GetAllowPairingOk returns a tuple with the AllowPairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPairing

`func (o *GetMobileDevicePrestageV2) SetAllowPairing(v bool)`

SetAllowPairing sets AllowPairing field to given value.


### GetMultiUser

`func (o *GetMobileDevicePrestageV2) GetMultiUser() bool`

GetMultiUser returns the MultiUser field if non-nil, zero value otherwise.

### GetMultiUserOk

`func (o *GetMobileDevicePrestageV2) GetMultiUserOk() (*bool, bool)`

GetMultiUserOk returns a tuple with the MultiUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiUser

`func (o *GetMobileDevicePrestageV2) SetMultiUser(v bool)`

SetMultiUser sets MultiUser field to given value.


### GetSupervised

`func (o *GetMobileDevicePrestageV2) GetSupervised() bool`

GetSupervised returns the Supervised field if non-nil, zero value otherwise.

### GetSupervisedOk

`func (o *GetMobileDevicePrestageV2) GetSupervisedOk() (*bool, bool)`

GetSupervisedOk returns a tuple with the Supervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervised

`func (o *GetMobileDevicePrestageV2) SetSupervised(v bool)`

SetSupervised sets Supervised field to given value.


### GetMaximumSharedAccounts

`func (o *GetMobileDevicePrestageV2) GetMaximumSharedAccounts() int64`

GetMaximumSharedAccounts returns the MaximumSharedAccounts field if non-nil, zero value otherwise.

### GetMaximumSharedAccountsOk

`func (o *GetMobileDevicePrestageV2) GetMaximumSharedAccountsOk() (*int64, bool)`

GetMaximumSharedAccountsOk returns a tuple with the MaximumSharedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSharedAccounts

`func (o *GetMobileDevicePrestageV2) SetMaximumSharedAccounts(v int64)`

SetMaximumSharedAccounts sets MaximumSharedAccounts field to given value.


### GetConfigureDeviceBeforeSetupAssistant

`func (o *GetMobileDevicePrestageV2) GetConfigureDeviceBeforeSetupAssistant() bool`

GetConfigureDeviceBeforeSetupAssistant returns the ConfigureDeviceBeforeSetupAssistant field if non-nil, zero value otherwise.

### GetConfigureDeviceBeforeSetupAssistantOk

`func (o *GetMobileDevicePrestageV2) GetConfigureDeviceBeforeSetupAssistantOk() (*bool, bool)`

GetConfigureDeviceBeforeSetupAssistantOk returns a tuple with the ConfigureDeviceBeforeSetupAssistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureDeviceBeforeSetupAssistant

`func (o *GetMobileDevicePrestageV2) SetConfigureDeviceBeforeSetupAssistant(v bool)`

SetConfigureDeviceBeforeSetupAssistant sets ConfigureDeviceBeforeSetupAssistant field to given value.


### GetNames

`func (o *GetMobileDevicePrestageV2) GetNames() MobileDevicePrestageNamesV2`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *GetMobileDevicePrestageV2) GetNamesOk() (*MobileDevicePrestageNamesV2, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *GetMobileDevicePrestageV2) SetNames(v MobileDevicePrestageNamesV2)`

SetNames sets Names field to given value.

### HasNames

`func (o *GetMobileDevicePrestageV2) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetSendTimezone

`func (o *GetMobileDevicePrestageV2) GetSendTimezone() bool`

GetSendTimezone returns the SendTimezone field if non-nil, zero value otherwise.

### GetSendTimezoneOk

`func (o *GetMobileDevicePrestageV2) GetSendTimezoneOk() (*bool, bool)`

GetSendTimezoneOk returns a tuple with the SendTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendTimezone

`func (o *GetMobileDevicePrestageV2) SetSendTimezone(v bool)`

SetSendTimezone sets SendTimezone field to given value.


### GetTimezone

`func (o *GetMobileDevicePrestageV2) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *GetMobileDevicePrestageV2) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *GetMobileDevicePrestageV2) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetStorageQuotaSizeMegabytes

`func (o *GetMobileDevicePrestageV2) GetStorageQuotaSizeMegabytes() int64`

GetStorageQuotaSizeMegabytes returns the StorageQuotaSizeMegabytes field if non-nil, zero value otherwise.

### GetStorageQuotaSizeMegabytesOk

`func (o *GetMobileDevicePrestageV2) GetStorageQuotaSizeMegabytesOk() (*int64, bool)`

GetStorageQuotaSizeMegabytesOk returns a tuple with the StorageQuotaSizeMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageQuotaSizeMegabytes

`func (o *GetMobileDevicePrestageV2) SetStorageQuotaSizeMegabytes(v int64)`

SetStorageQuotaSizeMegabytes sets StorageQuotaSizeMegabytes field to given value.


### GetUseStorageQuotaSize

`func (o *GetMobileDevicePrestageV2) GetUseStorageQuotaSize() bool`

GetUseStorageQuotaSize returns the UseStorageQuotaSize field if non-nil, zero value otherwise.

### GetUseStorageQuotaSizeOk

`func (o *GetMobileDevicePrestageV2) GetUseStorageQuotaSizeOk() (*bool, bool)`

GetUseStorageQuotaSizeOk returns a tuple with the UseStorageQuotaSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseStorageQuotaSize

`func (o *GetMobileDevicePrestageV2) SetUseStorageQuotaSize(v bool)`

SetUseStorageQuotaSize sets UseStorageQuotaSize field to given value.


### GetTemporarySessionOnly

`func (o *GetMobileDevicePrestageV2) GetTemporarySessionOnly() bool`

GetTemporarySessionOnly returns the TemporarySessionOnly field if non-nil, zero value otherwise.

### GetTemporarySessionOnlyOk

`func (o *GetMobileDevicePrestageV2) GetTemporarySessionOnlyOk() (*bool, bool)`

GetTemporarySessionOnlyOk returns a tuple with the TemporarySessionOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporarySessionOnly

`func (o *GetMobileDevicePrestageV2) SetTemporarySessionOnly(v bool)`

SetTemporarySessionOnly sets TemporarySessionOnly field to given value.

### HasTemporarySessionOnly

`func (o *GetMobileDevicePrestageV2) HasTemporarySessionOnly() bool`

HasTemporarySessionOnly returns a boolean if a field has been set.

### GetEnforceTemporarySessionTimeout

`func (o *GetMobileDevicePrestageV2) GetEnforceTemporarySessionTimeout() bool`

GetEnforceTemporarySessionTimeout returns the EnforceTemporarySessionTimeout field if non-nil, zero value otherwise.

### GetEnforceTemporarySessionTimeoutOk

`func (o *GetMobileDevicePrestageV2) GetEnforceTemporarySessionTimeoutOk() (*bool, bool)`

GetEnforceTemporarySessionTimeoutOk returns a tuple with the EnforceTemporarySessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceTemporarySessionTimeout

`func (o *GetMobileDevicePrestageV2) SetEnforceTemporarySessionTimeout(v bool)`

SetEnforceTemporarySessionTimeout sets EnforceTemporarySessionTimeout field to given value.

### HasEnforceTemporarySessionTimeout

`func (o *GetMobileDevicePrestageV2) HasEnforceTemporarySessionTimeout() bool`

HasEnforceTemporarySessionTimeout returns a boolean if a field has been set.

### GetTemporarySessionTimeout

`func (o *GetMobileDevicePrestageV2) GetTemporarySessionTimeout() int64`

GetTemporarySessionTimeout returns the TemporarySessionTimeout field if non-nil, zero value otherwise.

### GetTemporarySessionTimeoutOk

`func (o *GetMobileDevicePrestageV2) GetTemporarySessionTimeoutOk() (*int64, bool)`

GetTemporarySessionTimeoutOk returns a tuple with the TemporarySessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporarySessionTimeout

`func (o *GetMobileDevicePrestageV2) SetTemporarySessionTimeout(v int64)`

SetTemporarySessionTimeout sets TemporarySessionTimeout field to given value.

### HasTemporarySessionTimeout

`func (o *GetMobileDevicePrestageV2) HasTemporarySessionTimeout() bool`

HasTemporarySessionTimeout returns a boolean if a field has been set.

### GetEnforceUserSessionTimeout

`func (o *GetMobileDevicePrestageV2) GetEnforceUserSessionTimeout() bool`

GetEnforceUserSessionTimeout returns the EnforceUserSessionTimeout field if non-nil, zero value otherwise.

### GetEnforceUserSessionTimeoutOk

`func (o *GetMobileDevicePrestageV2) GetEnforceUserSessionTimeoutOk() (*bool, bool)`

GetEnforceUserSessionTimeoutOk returns a tuple with the EnforceUserSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceUserSessionTimeout

`func (o *GetMobileDevicePrestageV2) SetEnforceUserSessionTimeout(v bool)`

SetEnforceUserSessionTimeout sets EnforceUserSessionTimeout field to given value.

### HasEnforceUserSessionTimeout

`func (o *GetMobileDevicePrestageV2) HasEnforceUserSessionTimeout() bool`

HasEnforceUserSessionTimeout returns a boolean if a field has been set.

### GetUserSessionTimeout

`func (o *GetMobileDevicePrestageV2) GetUserSessionTimeout() int64`

GetUserSessionTimeout returns the UserSessionTimeout field if non-nil, zero value otherwise.

### GetUserSessionTimeoutOk

`func (o *GetMobileDevicePrestageV2) GetUserSessionTimeoutOk() (*int64, bool)`

GetUserSessionTimeoutOk returns a tuple with the UserSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSessionTimeout

`func (o *GetMobileDevicePrestageV2) SetUserSessionTimeout(v int64)`

SetUserSessionTimeout sets UserSessionTimeout field to given value.

### HasUserSessionTimeout

`func (o *GetMobileDevicePrestageV2) HasUserSessionTimeout() bool`

HasUserSessionTimeout returns a boolean if a field has been set.

### GetId

`func (o *GetMobileDevicePrestageV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMobileDevicePrestageV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMobileDevicePrestageV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetMobileDevicePrestageV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProfileUuid

`func (o *GetMobileDevicePrestageV2) GetProfileUuid() string`

GetProfileUuid returns the ProfileUuid field if non-nil, zero value otherwise.

### GetProfileUuidOk

`func (o *GetMobileDevicePrestageV2) GetProfileUuidOk() (*string, bool)`

GetProfileUuidOk returns a tuple with the ProfileUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUuid

`func (o *GetMobileDevicePrestageV2) SetProfileUuid(v string)`

SetProfileUuid sets ProfileUuid field to given value.

### HasProfileUuid

`func (o *GetMobileDevicePrestageV2) HasProfileUuid() bool`

HasProfileUuid returns a boolean if a field has been set.

### GetSiteId

`func (o *GetMobileDevicePrestageV2) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *GetMobileDevicePrestageV2) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *GetMobileDevicePrestageV2) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *GetMobileDevicePrestageV2) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetVersionLock

`func (o *GetMobileDevicePrestageV2) GetVersionLock() int64`

GetVersionLock returns the VersionLock field if non-nil, zero value otherwise.

### GetVersionLockOk

`func (o *GetMobileDevicePrestageV2) GetVersionLockOk() (*int64, bool)`

GetVersionLockOk returns a tuple with the VersionLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionLock

`func (o *GetMobileDevicePrestageV2) SetVersionLock(v int64)`

SetVersionLock sets VersionLock field to given value.

### HasVersionLock

`func (o *GetMobileDevicePrestageV2) HasVersionLock() bool`

HasVersionLock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


