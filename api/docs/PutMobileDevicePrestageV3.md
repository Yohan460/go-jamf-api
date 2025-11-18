# PutMobileDevicePrestageV3

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
**LocationInformation** | [**LocationInformationV3**](LocationInformationV3.md) |  | 
**PurchasingInformation** | [**PrestagePurchasingInformationV3**](PrestagePurchasingInformationV3.md) |  | 
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
**Names** | Pointer to [**MobileDevicePrestageNamesV3**](MobileDevicePrestageNamesV3.md) |  | [optional] 
**SendTimezone** | **bool** |  | 
**Timezone** | **string** |  | 
**StorageQuotaSizeMegabytes** | **int64** |  | 
**UseStorageQuotaSize** | **bool** |  | 
**TemporarySessionOnly** | Pointer to **bool** |  | [optional] 
**EnforceTemporarySessionTimeout** | Pointer to **bool** |  | [optional] 
**TemporarySessionTimeout** | Pointer to **int64** |  | [optional] 
**EnforceUserSessionTimeout** | Pointer to **bool** |  | [optional] 
**UserSessionTimeout** | Pointer to **int64** |  | [optional] 
**PrestageMinimumOsTargetVersionTypeIos** | Pointer to **string** |  | [optional] 
**MinimumOsSpecificVersionIos** | Pointer to **string** |  | [optional] 
**PrestageMinimumOsTargetVersionTypeIpad** | Pointer to **string** |  | [optional] 
**MinimumOsSpecificVersionIpad** | Pointer to **string** |  | [optional] 
**RtsEnabled** | Pointer to **bool** |  | [optional] 
**RtsConfigProfileId** | Pointer to **string** |  | [optional] 
**PreserveManagedApps** | Pointer to **bool** | Controls whether managed apps are preserved during Return to Service operations. | [optional] 
**VersionLock** | Pointer to **int64** |  | [optional] 

## Methods

### NewPutMobileDevicePrestageV3

`func NewPutMobileDevicePrestageV3(displayName string, mandatory bool, mdmRemovable bool, supportPhoneNumber string, supportEmailAddress string, department string, defaultPrestage bool, enrollmentSiteId string, keepExistingSiteMembership bool, keepExistingLocationInformation bool, requireAuthentication bool, authenticationPrompt string, preventActivationLock bool, enableDeviceBasedActivationLock bool, deviceEnrollmentProgramInstanceId string, locationInformation LocationInformationV3, purchasingInformation PrestagePurchasingInformationV3, autoAdvanceSetup bool, allowPairing bool, multiUser bool, supervised bool, maximumSharedAccounts int64, configureDeviceBeforeSetupAssistant bool, sendTimezone bool, timezone string, storageQuotaSizeMegabytes int64, useStorageQuotaSize bool, ) *PutMobileDevicePrestageV3`

NewPutMobileDevicePrestageV3 instantiates a new PutMobileDevicePrestageV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutMobileDevicePrestageV3WithDefaults

`func NewPutMobileDevicePrestageV3WithDefaults() *PutMobileDevicePrestageV3`

NewPutMobileDevicePrestageV3WithDefaults instantiates a new PutMobileDevicePrestageV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *PutMobileDevicePrestageV3) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PutMobileDevicePrestageV3) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PutMobileDevicePrestageV3) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMandatory

`func (o *PutMobileDevicePrestageV3) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *PutMobileDevicePrestageV3) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *PutMobileDevicePrestageV3) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.


### GetMdmRemovable

`func (o *PutMobileDevicePrestageV3) GetMdmRemovable() bool`

GetMdmRemovable returns the MdmRemovable field if non-nil, zero value otherwise.

### GetMdmRemovableOk

`func (o *PutMobileDevicePrestageV3) GetMdmRemovableOk() (*bool, bool)`

GetMdmRemovableOk returns a tuple with the MdmRemovable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmRemovable

`func (o *PutMobileDevicePrestageV3) SetMdmRemovable(v bool)`

SetMdmRemovable sets MdmRemovable field to given value.


### GetSupportPhoneNumber

`func (o *PutMobileDevicePrestageV3) GetSupportPhoneNumber() string`

GetSupportPhoneNumber returns the SupportPhoneNumber field if non-nil, zero value otherwise.

### GetSupportPhoneNumberOk

`func (o *PutMobileDevicePrestageV3) GetSupportPhoneNumberOk() (*string, bool)`

GetSupportPhoneNumberOk returns a tuple with the SupportPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPhoneNumber

`func (o *PutMobileDevicePrestageV3) SetSupportPhoneNumber(v string)`

SetSupportPhoneNumber sets SupportPhoneNumber field to given value.


### GetSupportEmailAddress

`func (o *PutMobileDevicePrestageV3) GetSupportEmailAddress() string`

GetSupportEmailAddress returns the SupportEmailAddress field if non-nil, zero value otherwise.

### GetSupportEmailAddressOk

`func (o *PutMobileDevicePrestageV3) GetSupportEmailAddressOk() (*string, bool)`

GetSupportEmailAddressOk returns a tuple with the SupportEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmailAddress

`func (o *PutMobileDevicePrestageV3) SetSupportEmailAddress(v string)`

SetSupportEmailAddress sets SupportEmailAddress field to given value.


### GetDepartment

`func (o *PutMobileDevicePrestageV3) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *PutMobileDevicePrestageV3) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *PutMobileDevicePrestageV3) SetDepartment(v string)`

SetDepartment sets Department field to given value.


### GetDefaultPrestage

`func (o *PutMobileDevicePrestageV3) GetDefaultPrestage() bool`

GetDefaultPrestage returns the DefaultPrestage field if non-nil, zero value otherwise.

### GetDefaultPrestageOk

`func (o *PutMobileDevicePrestageV3) GetDefaultPrestageOk() (*bool, bool)`

GetDefaultPrestageOk returns a tuple with the DefaultPrestage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrestage

`func (o *PutMobileDevicePrestageV3) SetDefaultPrestage(v bool)`

SetDefaultPrestage sets DefaultPrestage field to given value.


### GetEnrollmentSiteId

`func (o *PutMobileDevicePrestageV3) GetEnrollmentSiteId() string`

GetEnrollmentSiteId returns the EnrollmentSiteId field if non-nil, zero value otherwise.

### GetEnrollmentSiteIdOk

`func (o *PutMobileDevicePrestageV3) GetEnrollmentSiteIdOk() (*string, bool)`

GetEnrollmentSiteIdOk returns a tuple with the EnrollmentSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSiteId

`func (o *PutMobileDevicePrestageV3) SetEnrollmentSiteId(v string)`

SetEnrollmentSiteId sets EnrollmentSiteId field to given value.


### GetKeepExistingSiteMembership

`func (o *PutMobileDevicePrestageV3) GetKeepExistingSiteMembership() bool`

GetKeepExistingSiteMembership returns the KeepExistingSiteMembership field if non-nil, zero value otherwise.

### GetKeepExistingSiteMembershipOk

`func (o *PutMobileDevicePrestageV3) GetKeepExistingSiteMembershipOk() (*bool, bool)`

GetKeepExistingSiteMembershipOk returns a tuple with the KeepExistingSiteMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepExistingSiteMembership

`func (o *PutMobileDevicePrestageV3) SetKeepExistingSiteMembership(v bool)`

SetKeepExistingSiteMembership sets KeepExistingSiteMembership field to given value.


### GetKeepExistingLocationInformation

`func (o *PutMobileDevicePrestageV3) GetKeepExistingLocationInformation() bool`

GetKeepExistingLocationInformation returns the KeepExistingLocationInformation field if non-nil, zero value otherwise.

### GetKeepExistingLocationInformationOk

`func (o *PutMobileDevicePrestageV3) GetKeepExistingLocationInformationOk() (*bool, bool)`

GetKeepExistingLocationInformationOk returns a tuple with the KeepExistingLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepExistingLocationInformation

`func (o *PutMobileDevicePrestageV3) SetKeepExistingLocationInformation(v bool)`

SetKeepExistingLocationInformation sets KeepExistingLocationInformation field to given value.


### GetRequireAuthentication

`func (o *PutMobileDevicePrestageV3) GetRequireAuthentication() bool`

GetRequireAuthentication returns the RequireAuthentication field if non-nil, zero value otherwise.

### GetRequireAuthenticationOk

`func (o *PutMobileDevicePrestageV3) GetRequireAuthenticationOk() (*bool, bool)`

GetRequireAuthenticationOk returns a tuple with the RequireAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuthentication

`func (o *PutMobileDevicePrestageV3) SetRequireAuthentication(v bool)`

SetRequireAuthentication sets RequireAuthentication field to given value.


### GetAuthenticationPrompt

`func (o *PutMobileDevicePrestageV3) GetAuthenticationPrompt() string`

GetAuthenticationPrompt returns the AuthenticationPrompt field if non-nil, zero value otherwise.

### GetAuthenticationPromptOk

`func (o *PutMobileDevicePrestageV3) GetAuthenticationPromptOk() (*string, bool)`

GetAuthenticationPromptOk returns a tuple with the AuthenticationPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPrompt

`func (o *PutMobileDevicePrestageV3) SetAuthenticationPrompt(v string)`

SetAuthenticationPrompt sets AuthenticationPrompt field to given value.


### GetPreventActivationLock

`func (o *PutMobileDevicePrestageV3) GetPreventActivationLock() bool`

GetPreventActivationLock returns the PreventActivationLock field if non-nil, zero value otherwise.

### GetPreventActivationLockOk

`func (o *PutMobileDevicePrestageV3) GetPreventActivationLockOk() (*bool, bool)`

GetPreventActivationLockOk returns a tuple with the PreventActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventActivationLock

`func (o *PutMobileDevicePrestageV3) SetPreventActivationLock(v bool)`

SetPreventActivationLock sets PreventActivationLock field to given value.


### GetEnableDeviceBasedActivationLock

`func (o *PutMobileDevicePrestageV3) GetEnableDeviceBasedActivationLock() bool`

GetEnableDeviceBasedActivationLock returns the EnableDeviceBasedActivationLock field if non-nil, zero value otherwise.

### GetEnableDeviceBasedActivationLockOk

`func (o *PutMobileDevicePrestageV3) GetEnableDeviceBasedActivationLockOk() (*bool, bool)`

GetEnableDeviceBasedActivationLockOk returns a tuple with the EnableDeviceBasedActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDeviceBasedActivationLock

`func (o *PutMobileDevicePrestageV3) SetEnableDeviceBasedActivationLock(v bool)`

SetEnableDeviceBasedActivationLock sets EnableDeviceBasedActivationLock field to given value.


### GetDeviceEnrollmentProgramInstanceId

`func (o *PutMobileDevicePrestageV3) GetDeviceEnrollmentProgramInstanceId() string`

GetDeviceEnrollmentProgramInstanceId returns the DeviceEnrollmentProgramInstanceId field if non-nil, zero value otherwise.

### GetDeviceEnrollmentProgramInstanceIdOk

`func (o *PutMobileDevicePrestageV3) GetDeviceEnrollmentProgramInstanceIdOk() (*string, bool)`

GetDeviceEnrollmentProgramInstanceIdOk returns a tuple with the DeviceEnrollmentProgramInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceEnrollmentProgramInstanceId

`func (o *PutMobileDevicePrestageV3) SetDeviceEnrollmentProgramInstanceId(v string)`

SetDeviceEnrollmentProgramInstanceId sets DeviceEnrollmentProgramInstanceId field to given value.


### GetSkipSetupItems

`func (o *PutMobileDevicePrestageV3) GetSkipSetupItems() map[string]bool`

GetSkipSetupItems returns the SkipSetupItems field if non-nil, zero value otherwise.

### GetSkipSetupItemsOk

`func (o *PutMobileDevicePrestageV3) GetSkipSetupItemsOk() (*map[string]bool, bool)`

GetSkipSetupItemsOk returns a tuple with the SkipSetupItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSetupItems

`func (o *PutMobileDevicePrestageV3) SetSkipSetupItems(v map[string]bool)`

SetSkipSetupItems sets SkipSetupItems field to given value.

### HasSkipSetupItems

`func (o *PutMobileDevicePrestageV3) HasSkipSetupItems() bool`

HasSkipSetupItems returns a boolean if a field has been set.

### GetLocationInformation

`func (o *PutMobileDevicePrestageV3) GetLocationInformation() LocationInformationV3`

GetLocationInformation returns the LocationInformation field if non-nil, zero value otherwise.

### GetLocationInformationOk

`func (o *PutMobileDevicePrestageV3) GetLocationInformationOk() (*LocationInformationV3, bool)`

GetLocationInformationOk returns a tuple with the LocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationInformation

`func (o *PutMobileDevicePrestageV3) SetLocationInformation(v LocationInformationV3)`

SetLocationInformation sets LocationInformation field to given value.


### GetPurchasingInformation

`func (o *PutMobileDevicePrestageV3) GetPurchasingInformation() PrestagePurchasingInformationV3`

GetPurchasingInformation returns the PurchasingInformation field if non-nil, zero value otherwise.

### GetPurchasingInformationOk

`func (o *PutMobileDevicePrestageV3) GetPurchasingInformationOk() (*PrestagePurchasingInformationV3, bool)`

GetPurchasingInformationOk returns a tuple with the PurchasingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingInformation

`func (o *PutMobileDevicePrestageV3) SetPurchasingInformation(v PrestagePurchasingInformationV3)`

SetPurchasingInformation sets PurchasingInformation field to given value.


### GetAnchorCertificates

`func (o *PutMobileDevicePrestageV3) GetAnchorCertificates() []string`

GetAnchorCertificates returns the AnchorCertificates field if non-nil, zero value otherwise.

### GetAnchorCertificatesOk

`func (o *PutMobileDevicePrestageV3) GetAnchorCertificatesOk() (*[]string, bool)`

GetAnchorCertificatesOk returns a tuple with the AnchorCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorCertificates

`func (o *PutMobileDevicePrestageV3) SetAnchorCertificates(v []string)`

SetAnchorCertificates sets AnchorCertificates field to given value.

### HasAnchorCertificates

`func (o *PutMobileDevicePrestageV3) HasAnchorCertificates() bool`

HasAnchorCertificates returns a boolean if a field has been set.

### GetEnrollmentCustomizationId

`func (o *PutMobileDevicePrestageV3) GetEnrollmentCustomizationId() string`

GetEnrollmentCustomizationId returns the EnrollmentCustomizationId field if non-nil, zero value otherwise.

### GetEnrollmentCustomizationIdOk

`func (o *PutMobileDevicePrestageV3) GetEnrollmentCustomizationIdOk() (*string, bool)`

GetEnrollmentCustomizationIdOk returns a tuple with the EnrollmentCustomizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCustomizationId

`func (o *PutMobileDevicePrestageV3) SetEnrollmentCustomizationId(v string)`

SetEnrollmentCustomizationId sets EnrollmentCustomizationId field to given value.

### HasEnrollmentCustomizationId

`func (o *PutMobileDevicePrestageV3) HasEnrollmentCustomizationId() bool`

HasEnrollmentCustomizationId returns a boolean if a field has been set.

### GetLanguage

`func (o *PutMobileDevicePrestageV3) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PutMobileDevicePrestageV3) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PutMobileDevicePrestageV3) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PutMobileDevicePrestageV3) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetRegion

`func (o *PutMobileDevicePrestageV3) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PutMobileDevicePrestageV3) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PutMobileDevicePrestageV3) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PutMobileDevicePrestageV3) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetAutoAdvanceSetup

`func (o *PutMobileDevicePrestageV3) GetAutoAdvanceSetup() bool`

GetAutoAdvanceSetup returns the AutoAdvanceSetup field if non-nil, zero value otherwise.

### GetAutoAdvanceSetupOk

`func (o *PutMobileDevicePrestageV3) GetAutoAdvanceSetupOk() (*bool, bool)`

GetAutoAdvanceSetupOk returns a tuple with the AutoAdvanceSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAdvanceSetup

`func (o *PutMobileDevicePrestageV3) SetAutoAdvanceSetup(v bool)`

SetAutoAdvanceSetup sets AutoAdvanceSetup field to given value.


### GetAllowPairing

`func (o *PutMobileDevicePrestageV3) GetAllowPairing() bool`

GetAllowPairing returns the AllowPairing field if non-nil, zero value otherwise.

### GetAllowPairingOk

`func (o *PutMobileDevicePrestageV3) GetAllowPairingOk() (*bool, bool)`

GetAllowPairingOk returns a tuple with the AllowPairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPairing

`func (o *PutMobileDevicePrestageV3) SetAllowPairing(v bool)`

SetAllowPairing sets AllowPairing field to given value.


### GetMultiUser

`func (o *PutMobileDevicePrestageV3) GetMultiUser() bool`

GetMultiUser returns the MultiUser field if non-nil, zero value otherwise.

### GetMultiUserOk

`func (o *PutMobileDevicePrestageV3) GetMultiUserOk() (*bool, bool)`

GetMultiUserOk returns a tuple with the MultiUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiUser

`func (o *PutMobileDevicePrestageV3) SetMultiUser(v bool)`

SetMultiUser sets MultiUser field to given value.


### GetSupervised

`func (o *PutMobileDevicePrestageV3) GetSupervised() bool`

GetSupervised returns the Supervised field if non-nil, zero value otherwise.

### GetSupervisedOk

`func (o *PutMobileDevicePrestageV3) GetSupervisedOk() (*bool, bool)`

GetSupervisedOk returns a tuple with the Supervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervised

`func (o *PutMobileDevicePrestageV3) SetSupervised(v bool)`

SetSupervised sets Supervised field to given value.


### GetMaximumSharedAccounts

`func (o *PutMobileDevicePrestageV3) GetMaximumSharedAccounts() int64`

GetMaximumSharedAccounts returns the MaximumSharedAccounts field if non-nil, zero value otherwise.

### GetMaximumSharedAccountsOk

`func (o *PutMobileDevicePrestageV3) GetMaximumSharedAccountsOk() (*int64, bool)`

GetMaximumSharedAccountsOk returns a tuple with the MaximumSharedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSharedAccounts

`func (o *PutMobileDevicePrestageV3) SetMaximumSharedAccounts(v int64)`

SetMaximumSharedAccounts sets MaximumSharedAccounts field to given value.


### GetConfigureDeviceBeforeSetupAssistant

`func (o *PutMobileDevicePrestageV3) GetConfigureDeviceBeforeSetupAssistant() bool`

GetConfigureDeviceBeforeSetupAssistant returns the ConfigureDeviceBeforeSetupAssistant field if non-nil, zero value otherwise.

### GetConfigureDeviceBeforeSetupAssistantOk

`func (o *PutMobileDevicePrestageV3) GetConfigureDeviceBeforeSetupAssistantOk() (*bool, bool)`

GetConfigureDeviceBeforeSetupAssistantOk returns a tuple with the ConfigureDeviceBeforeSetupAssistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureDeviceBeforeSetupAssistant

`func (o *PutMobileDevicePrestageV3) SetConfigureDeviceBeforeSetupAssistant(v bool)`

SetConfigureDeviceBeforeSetupAssistant sets ConfigureDeviceBeforeSetupAssistant field to given value.


### GetNames

`func (o *PutMobileDevicePrestageV3) GetNames() MobileDevicePrestageNamesV3`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *PutMobileDevicePrestageV3) GetNamesOk() (*MobileDevicePrestageNamesV3, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *PutMobileDevicePrestageV3) SetNames(v MobileDevicePrestageNamesV3)`

SetNames sets Names field to given value.

### HasNames

`func (o *PutMobileDevicePrestageV3) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetSendTimezone

`func (o *PutMobileDevicePrestageV3) GetSendTimezone() bool`

GetSendTimezone returns the SendTimezone field if non-nil, zero value otherwise.

### GetSendTimezoneOk

`func (o *PutMobileDevicePrestageV3) GetSendTimezoneOk() (*bool, bool)`

GetSendTimezoneOk returns a tuple with the SendTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendTimezone

`func (o *PutMobileDevicePrestageV3) SetSendTimezone(v bool)`

SetSendTimezone sets SendTimezone field to given value.


### GetTimezone

`func (o *PutMobileDevicePrestageV3) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *PutMobileDevicePrestageV3) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *PutMobileDevicePrestageV3) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetStorageQuotaSizeMegabytes

`func (o *PutMobileDevicePrestageV3) GetStorageQuotaSizeMegabytes() int64`

GetStorageQuotaSizeMegabytes returns the StorageQuotaSizeMegabytes field if non-nil, zero value otherwise.

### GetStorageQuotaSizeMegabytesOk

`func (o *PutMobileDevicePrestageV3) GetStorageQuotaSizeMegabytesOk() (*int64, bool)`

GetStorageQuotaSizeMegabytesOk returns a tuple with the StorageQuotaSizeMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageQuotaSizeMegabytes

`func (o *PutMobileDevicePrestageV3) SetStorageQuotaSizeMegabytes(v int64)`

SetStorageQuotaSizeMegabytes sets StorageQuotaSizeMegabytes field to given value.


### GetUseStorageQuotaSize

`func (o *PutMobileDevicePrestageV3) GetUseStorageQuotaSize() bool`

GetUseStorageQuotaSize returns the UseStorageQuotaSize field if non-nil, zero value otherwise.

### GetUseStorageQuotaSizeOk

`func (o *PutMobileDevicePrestageV3) GetUseStorageQuotaSizeOk() (*bool, bool)`

GetUseStorageQuotaSizeOk returns a tuple with the UseStorageQuotaSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseStorageQuotaSize

`func (o *PutMobileDevicePrestageV3) SetUseStorageQuotaSize(v bool)`

SetUseStorageQuotaSize sets UseStorageQuotaSize field to given value.


### GetTemporarySessionOnly

`func (o *PutMobileDevicePrestageV3) GetTemporarySessionOnly() bool`

GetTemporarySessionOnly returns the TemporarySessionOnly field if non-nil, zero value otherwise.

### GetTemporarySessionOnlyOk

`func (o *PutMobileDevicePrestageV3) GetTemporarySessionOnlyOk() (*bool, bool)`

GetTemporarySessionOnlyOk returns a tuple with the TemporarySessionOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporarySessionOnly

`func (o *PutMobileDevicePrestageV3) SetTemporarySessionOnly(v bool)`

SetTemporarySessionOnly sets TemporarySessionOnly field to given value.

### HasTemporarySessionOnly

`func (o *PutMobileDevicePrestageV3) HasTemporarySessionOnly() bool`

HasTemporarySessionOnly returns a boolean if a field has been set.

### GetEnforceTemporarySessionTimeout

`func (o *PutMobileDevicePrestageV3) GetEnforceTemporarySessionTimeout() bool`

GetEnforceTemporarySessionTimeout returns the EnforceTemporarySessionTimeout field if non-nil, zero value otherwise.

### GetEnforceTemporarySessionTimeoutOk

`func (o *PutMobileDevicePrestageV3) GetEnforceTemporarySessionTimeoutOk() (*bool, bool)`

GetEnforceTemporarySessionTimeoutOk returns a tuple with the EnforceTemporarySessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceTemporarySessionTimeout

`func (o *PutMobileDevicePrestageV3) SetEnforceTemporarySessionTimeout(v bool)`

SetEnforceTemporarySessionTimeout sets EnforceTemporarySessionTimeout field to given value.

### HasEnforceTemporarySessionTimeout

`func (o *PutMobileDevicePrestageV3) HasEnforceTemporarySessionTimeout() bool`

HasEnforceTemporarySessionTimeout returns a boolean if a field has been set.

### GetTemporarySessionTimeout

`func (o *PutMobileDevicePrestageV3) GetTemporarySessionTimeout() int64`

GetTemporarySessionTimeout returns the TemporarySessionTimeout field if non-nil, zero value otherwise.

### GetTemporarySessionTimeoutOk

`func (o *PutMobileDevicePrestageV3) GetTemporarySessionTimeoutOk() (*int64, bool)`

GetTemporarySessionTimeoutOk returns a tuple with the TemporarySessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporarySessionTimeout

`func (o *PutMobileDevicePrestageV3) SetTemporarySessionTimeout(v int64)`

SetTemporarySessionTimeout sets TemporarySessionTimeout field to given value.

### HasTemporarySessionTimeout

`func (o *PutMobileDevicePrestageV3) HasTemporarySessionTimeout() bool`

HasTemporarySessionTimeout returns a boolean if a field has been set.

### GetEnforceUserSessionTimeout

`func (o *PutMobileDevicePrestageV3) GetEnforceUserSessionTimeout() bool`

GetEnforceUserSessionTimeout returns the EnforceUserSessionTimeout field if non-nil, zero value otherwise.

### GetEnforceUserSessionTimeoutOk

`func (o *PutMobileDevicePrestageV3) GetEnforceUserSessionTimeoutOk() (*bool, bool)`

GetEnforceUserSessionTimeoutOk returns a tuple with the EnforceUserSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceUserSessionTimeout

`func (o *PutMobileDevicePrestageV3) SetEnforceUserSessionTimeout(v bool)`

SetEnforceUserSessionTimeout sets EnforceUserSessionTimeout field to given value.

### HasEnforceUserSessionTimeout

`func (o *PutMobileDevicePrestageV3) HasEnforceUserSessionTimeout() bool`

HasEnforceUserSessionTimeout returns a boolean if a field has been set.

### GetUserSessionTimeout

`func (o *PutMobileDevicePrestageV3) GetUserSessionTimeout() int64`

GetUserSessionTimeout returns the UserSessionTimeout field if non-nil, zero value otherwise.

### GetUserSessionTimeoutOk

`func (o *PutMobileDevicePrestageV3) GetUserSessionTimeoutOk() (*int64, bool)`

GetUserSessionTimeoutOk returns a tuple with the UserSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSessionTimeout

`func (o *PutMobileDevicePrestageV3) SetUserSessionTimeout(v int64)`

SetUserSessionTimeout sets UserSessionTimeout field to given value.

### HasUserSessionTimeout

`func (o *PutMobileDevicePrestageV3) HasUserSessionTimeout() bool`

HasUserSessionTimeout returns a boolean if a field has been set.

### GetPrestageMinimumOsTargetVersionTypeIos

`func (o *PutMobileDevicePrestageV3) GetPrestageMinimumOsTargetVersionTypeIos() string`

GetPrestageMinimumOsTargetVersionTypeIos returns the PrestageMinimumOsTargetVersionTypeIos field if non-nil, zero value otherwise.

### GetPrestageMinimumOsTargetVersionTypeIosOk

`func (o *PutMobileDevicePrestageV3) GetPrestageMinimumOsTargetVersionTypeIosOk() (*string, bool)`

GetPrestageMinimumOsTargetVersionTypeIosOk returns a tuple with the PrestageMinimumOsTargetVersionTypeIos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrestageMinimumOsTargetVersionTypeIos

`func (o *PutMobileDevicePrestageV3) SetPrestageMinimumOsTargetVersionTypeIos(v string)`

SetPrestageMinimumOsTargetVersionTypeIos sets PrestageMinimumOsTargetVersionTypeIos field to given value.

### HasPrestageMinimumOsTargetVersionTypeIos

`func (o *PutMobileDevicePrestageV3) HasPrestageMinimumOsTargetVersionTypeIos() bool`

HasPrestageMinimumOsTargetVersionTypeIos returns a boolean if a field has been set.

### GetMinimumOsSpecificVersionIos

`func (o *PutMobileDevicePrestageV3) GetMinimumOsSpecificVersionIos() string`

GetMinimumOsSpecificVersionIos returns the MinimumOsSpecificVersionIos field if non-nil, zero value otherwise.

### GetMinimumOsSpecificVersionIosOk

`func (o *PutMobileDevicePrestageV3) GetMinimumOsSpecificVersionIosOk() (*string, bool)`

GetMinimumOsSpecificVersionIosOk returns a tuple with the MinimumOsSpecificVersionIos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumOsSpecificVersionIos

`func (o *PutMobileDevicePrestageV3) SetMinimumOsSpecificVersionIos(v string)`

SetMinimumOsSpecificVersionIos sets MinimumOsSpecificVersionIos field to given value.

### HasMinimumOsSpecificVersionIos

`func (o *PutMobileDevicePrestageV3) HasMinimumOsSpecificVersionIos() bool`

HasMinimumOsSpecificVersionIos returns a boolean if a field has been set.

### GetPrestageMinimumOsTargetVersionTypeIpad

`func (o *PutMobileDevicePrestageV3) GetPrestageMinimumOsTargetVersionTypeIpad() string`

GetPrestageMinimumOsTargetVersionTypeIpad returns the PrestageMinimumOsTargetVersionTypeIpad field if non-nil, zero value otherwise.

### GetPrestageMinimumOsTargetVersionTypeIpadOk

`func (o *PutMobileDevicePrestageV3) GetPrestageMinimumOsTargetVersionTypeIpadOk() (*string, bool)`

GetPrestageMinimumOsTargetVersionTypeIpadOk returns a tuple with the PrestageMinimumOsTargetVersionTypeIpad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrestageMinimumOsTargetVersionTypeIpad

`func (o *PutMobileDevicePrestageV3) SetPrestageMinimumOsTargetVersionTypeIpad(v string)`

SetPrestageMinimumOsTargetVersionTypeIpad sets PrestageMinimumOsTargetVersionTypeIpad field to given value.

### HasPrestageMinimumOsTargetVersionTypeIpad

`func (o *PutMobileDevicePrestageV3) HasPrestageMinimumOsTargetVersionTypeIpad() bool`

HasPrestageMinimumOsTargetVersionTypeIpad returns a boolean if a field has been set.

### GetMinimumOsSpecificVersionIpad

`func (o *PutMobileDevicePrestageV3) GetMinimumOsSpecificVersionIpad() string`

GetMinimumOsSpecificVersionIpad returns the MinimumOsSpecificVersionIpad field if non-nil, zero value otherwise.

### GetMinimumOsSpecificVersionIpadOk

`func (o *PutMobileDevicePrestageV3) GetMinimumOsSpecificVersionIpadOk() (*string, bool)`

GetMinimumOsSpecificVersionIpadOk returns a tuple with the MinimumOsSpecificVersionIpad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumOsSpecificVersionIpad

`func (o *PutMobileDevicePrestageV3) SetMinimumOsSpecificVersionIpad(v string)`

SetMinimumOsSpecificVersionIpad sets MinimumOsSpecificVersionIpad field to given value.

### HasMinimumOsSpecificVersionIpad

`func (o *PutMobileDevicePrestageV3) HasMinimumOsSpecificVersionIpad() bool`

HasMinimumOsSpecificVersionIpad returns a boolean if a field has been set.

### GetRtsEnabled

`func (o *PutMobileDevicePrestageV3) GetRtsEnabled() bool`

GetRtsEnabled returns the RtsEnabled field if non-nil, zero value otherwise.

### GetRtsEnabledOk

`func (o *PutMobileDevicePrestageV3) GetRtsEnabledOk() (*bool, bool)`

GetRtsEnabledOk returns a tuple with the RtsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtsEnabled

`func (o *PutMobileDevicePrestageV3) SetRtsEnabled(v bool)`

SetRtsEnabled sets RtsEnabled field to given value.

### HasRtsEnabled

`func (o *PutMobileDevicePrestageV3) HasRtsEnabled() bool`

HasRtsEnabled returns a boolean if a field has been set.

### GetRtsConfigProfileId

`func (o *PutMobileDevicePrestageV3) GetRtsConfigProfileId() string`

GetRtsConfigProfileId returns the RtsConfigProfileId field if non-nil, zero value otherwise.

### GetRtsConfigProfileIdOk

`func (o *PutMobileDevicePrestageV3) GetRtsConfigProfileIdOk() (*string, bool)`

GetRtsConfigProfileIdOk returns a tuple with the RtsConfigProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtsConfigProfileId

`func (o *PutMobileDevicePrestageV3) SetRtsConfigProfileId(v string)`

SetRtsConfigProfileId sets RtsConfigProfileId field to given value.

### HasRtsConfigProfileId

`func (o *PutMobileDevicePrestageV3) HasRtsConfigProfileId() bool`

HasRtsConfigProfileId returns a boolean if a field has been set.

### GetPreserveManagedApps

`func (o *PutMobileDevicePrestageV3) GetPreserveManagedApps() bool`

GetPreserveManagedApps returns the PreserveManagedApps field if non-nil, zero value otherwise.

### GetPreserveManagedAppsOk

`func (o *PutMobileDevicePrestageV3) GetPreserveManagedAppsOk() (*bool, bool)`

GetPreserveManagedAppsOk returns a tuple with the PreserveManagedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveManagedApps

`func (o *PutMobileDevicePrestageV3) SetPreserveManagedApps(v bool)`

SetPreserveManagedApps sets PreserveManagedApps field to given value.

### HasPreserveManagedApps

`func (o *PutMobileDevicePrestageV3) HasPreserveManagedApps() bool`

HasPreserveManagedApps returns a boolean if a field has been set.

### GetVersionLock

`func (o *PutMobileDevicePrestageV3) GetVersionLock() int64`

GetVersionLock returns the VersionLock field if non-nil, zero value otherwise.

### GetVersionLockOk

`func (o *PutMobileDevicePrestageV3) GetVersionLockOk() (*int64, bool)`

GetVersionLockOk returns a tuple with the VersionLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionLock

`func (o *PutMobileDevicePrestageV3) SetVersionLock(v int64)`

SetVersionLock sets VersionLock field to given value.

### HasVersionLock

`func (o *PutMobileDevicePrestageV3) HasVersionLock() bool`

HasVersionLock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


