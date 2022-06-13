# PutMobileDevicePrestage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**IsMandatory** | **bool** |  | 
**IsMdmRemovable** | **bool** |  | 
**SupportPhoneNumber** | **string** |  | 
**SupportEmailAddress** | **string** |  | 
**Department** | **string** |  | 
**IsDefaultPrestage** | **bool** |  | 
**EnrollmentSiteId** | **int32** |  | 
**IsKeepExistingSiteMembership** | **bool** |  | 
**IsKeepExistingLocationInformation** | **bool** |  | 
**IsRequireAuthentication** | **bool** |  | 
**AuthenticationPrompt** | **string** |  | 
**IsPreventActivationLock** | **bool** |  | 
**IsEnableDeviceBasedActivationLock** | **bool** |  | 
**DeviceEnrollmentProgramInstanceId** | **int32** |  | 
**SkipSetupItems** | Pointer to **map[string]bool** |  | [optional] 
**LocationInformation** | [**LocationInformation**](LocationInformation.md) |  | 
**PurchasingInformation** | [**PrestagePurchasingInformation**](PrestagePurchasingInformation.md) |  | 
**AnchorCertificates** | Pointer to **[]string** | The Base64 encoded PEM Certificate | [optional] 
**EnrollmentCustomizationId** | Pointer to **int32** |  | [optional] 
**IsAllowPairing** | **bool** |  | 
**IsMultiUser** | **bool** |  | 
**IsSupervised** | **bool** |  | 
**MaximumSharedAccounts** | **int32** |  | 
**IsAutoAdvanceSetup** | **bool** |  | 
**IsConfigureDeviceBeforeSetupAssistant** | **bool** |  | 
**Language** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Names** | Pointer to [**MobileDevicePrestageNames**](MobileDevicePrestageNames.md) |  | [optional] 
**VersionLock** | Pointer to **int32** |  | [optional] 

## Methods

### NewPutMobileDevicePrestage

`func NewPutMobileDevicePrestage(displayName string, isMandatory bool, isMdmRemovable bool, supportPhoneNumber string, supportEmailAddress string, department string, isDefaultPrestage bool, enrollmentSiteId int32, isKeepExistingSiteMembership bool, isKeepExistingLocationInformation bool, isRequireAuthentication bool, authenticationPrompt string, isPreventActivationLock bool, isEnableDeviceBasedActivationLock bool, deviceEnrollmentProgramInstanceId int32, locationInformation LocationInformation, purchasingInformation PrestagePurchasingInformation, isAllowPairing bool, isMultiUser bool, isSupervised bool, maximumSharedAccounts int32, isAutoAdvanceSetup bool, isConfigureDeviceBeforeSetupAssistant bool, ) *PutMobileDevicePrestage`

NewPutMobileDevicePrestage instantiates a new PutMobileDevicePrestage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutMobileDevicePrestageWithDefaults

`func NewPutMobileDevicePrestageWithDefaults() *PutMobileDevicePrestage`

NewPutMobileDevicePrestageWithDefaults instantiates a new PutMobileDevicePrestage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *PutMobileDevicePrestage) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PutMobileDevicePrestage) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PutMobileDevicePrestage) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetIsMandatory

`func (o *PutMobileDevicePrestage) GetIsMandatory() bool`

GetIsMandatory returns the IsMandatory field if non-nil, zero value otherwise.

### GetIsMandatoryOk

`func (o *PutMobileDevicePrestage) GetIsMandatoryOk() (*bool, bool)`

GetIsMandatoryOk returns a tuple with the IsMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMandatory

`func (o *PutMobileDevicePrestage) SetIsMandatory(v bool)`

SetIsMandatory sets IsMandatory field to given value.


### GetIsMdmRemovable

`func (o *PutMobileDevicePrestage) GetIsMdmRemovable() bool`

GetIsMdmRemovable returns the IsMdmRemovable field if non-nil, zero value otherwise.

### GetIsMdmRemovableOk

`func (o *PutMobileDevicePrestage) GetIsMdmRemovableOk() (*bool, bool)`

GetIsMdmRemovableOk returns a tuple with the IsMdmRemovable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMdmRemovable

`func (o *PutMobileDevicePrestage) SetIsMdmRemovable(v bool)`

SetIsMdmRemovable sets IsMdmRemovable field to given value.


### GetSupportPhoneNumber

`func (o *PutMobileDevicePrestage) GetSupportPhoneNumber() string`

GetSupportPhoneNumber returns the SupportPhoneNumber field if non-nil, zero value otherwise.

### GetSupportPhoneNumberOk

`func (o *PutMobileDevicePrestage) GetSupportPhoneNumberOk() (*string, bool)`

GetSupportPhoneNumberOk returns a tuple with the SupportPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPhoneNumber

`func (o *PutMobileDevicePrestage) SetSupportPhoneNumber(v string)`

SetSupportPhoneNumber sets SupportPhoneNumber field to given value.


### GetSupportEmailAddress

`func (o *PutMobileDevicePrestage) GetSupportEmailAddress() string`

GetSupportEmailAddress returns the SupportEmailAddress field if non-nil, zero value otherwise.

### GetSupportEmailAddressOk

`func (o *PutMobileDevicePrestage) GetSupportEmailAddressOk() (*string, bool)`

GetSupportEmailAddressOk returns a tuple with the SupportEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmailAddress

`func (o *PutMobileDevicePrestage) SetSupportEmailAddress(v string)`

SetSupportEmailAddress sets SupportEmailAddress field to given value.


### GetDepartment

`func (o *PutMobileDevicePrestage) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *PutMobileDevicePrestage) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *PutMobileDevicePrestage) SetDepartment(v string)`

SetDepartment sets Department field to given value.


### GetIsDefaultPrestage

`func (o *PutMobileDevicePrestage) GetIsDefaultPrestage() bool`

GetIsDefaultPrestage returns the IsDefaultPrestage field if non-nil, zero value otherwise.

### GetIsDefaultPrestageOk

`func (o *PutMobileDevicePrestage) GetIsDefaultPrestageOk() (*bool, bool)`

GetIsDefaultPrestageOk returns a tuple with the IsDefaultPrestage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultPrestage

`func (o *PutMobileDevicePrestage) SetIsDefaultPrestage(v bool)`

SetIsDefaultPrestage sets IsDefaultPrestage field to given value.


### GetEnrollmentSiteId

`func (o *PutMobileDevicePrestage) GetEnrollmentSiteId() int32`

GetEnrollmentSiteId returns the EnrollmentSiteId field if non-nil, zero value otherwise.

### GetEnrollmentSiteIdOk

`func (o *PutMobileDevicePrestage) GetEnrollmentSiteIdOk() (*int32, bool)`

GetEnrollmentSiteIdOk returns a tuple with the EnrollmentSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSiteId

`func (o *PutMobileDevicePrestage) SetEnrollmentSiteId(v int32)`

SetEnrollmentSiteId sets EnrollmentSiteId field to given value.


### GetIsKeepExistingSiteMembership

`func (o *PutMobileDevicePrestage) GetIsKeepExistingSiteMembership() bool`

GetIsKeepExistingSiteMembership returns the IsKeepExistingSiteMembership field if non-nil, zero value otherwise.

### GetIsKeepExistingSiteMembershipOk

`func (o *PutMobileDevicePrestage) GetIsKeepExistingSiteMembershipOk() (*bool, bool)`

GetIsKeepExistingSiteMembershipOk returns a tuple with the IsKeepExistingSiteMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKeepExistingSiteMembership

`func (o *PutMobileDevicePrestage) SetIsKeepExistingSiteMembership(v bool)`

SetIsKeepExistingSiteMembership sets IsKeepExistingSiteMembership field to given value.


### GetIsKeepExistingLocationInformation

`func (o *PutMobileDevicePrestage) GetIsKeepExistingLocationInformation() bool`

GetIsKeepExistingLocationInformation returns the IsKeepExistingLocationInformation field if non-nil, zero value otherwise.

### GetIsKeepExistingLocationInformationOk

`func (o *PutMobileDevicePrestage) GetIsKeepExistingLocationInformationOk() (*bool, bool)`

GetIsKeepExistingLocationInformationOk returns a tuple with the IsKeepExistingLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKeepExistingLocationInformation

`func (o *PutMobileDevicePrestage) SetIsKeepExistingLocationInformation(v bool)`

SetIsKeepExistingLocationInformation sets IsKeepExistingLocationInformation field to given value.


### GetIsRequireAuthentication

`func (o *PutMobileDevicePrestage) GetIsRequireAuthentication() bool`

GetIsRequireAuthentication returns the IsRequireAuthentication field if non-nil, zero value otherwise.

### GetIsRequireAuthenticationOk

`func (o *PutMobileDevicePrestage) GetIsRequireAuthenticationOk() (*bool, bool)`

GetIsRequireAuthenticationOk returns a tuple with the IsRequireAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequireAuthentication

`func (o *PutMobileDevicePrestage) SetIsRequireAuthentication(v bool)`

SetIsRequireAuthentication sets IsRequireAuthentication field to given value.


### GetAuthenticationPrompt

`func (o *PutMobileDevicePrestage) GetAuthenticationPrompt() string`

GetAuthenticationPrompt returns the AuthenticationPrompt field if non-nil, zero value otherwise.

### GetAuthenticationPromptOk

`func (o *PutMobileDevicePrestage) GetAuthenticationPromptOk() (*string, bool)`

GetAuthenticationPromptOk returns a tuple with the AuthenticationPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPrompt

`func (o *PutMobileDevicePrestage) SetAuthenticationPrompt(v string)`

SetAuthenticationPrompt sets AuthenticationPrompt field to given value.


### GetIsPreventActivationLock

`func (o *PutMobileDevicePrestage) GetIsPreventActivationLock() bool`

GetIsPreventActivationLock returns the IsPreventActivationLock field if non-nil, zero value otherwise.

### GetIsPreventActivationLockOk

`func (o *PutMobileDevicePrestage) GetIsPreventActivationLockOk() (*bool, bool)`

GetIsPreventActivationLockOk returns a tuple with the IsPreventActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreventActivationLock

`func (o *PutMobileDevicePrestage) SetIsPreventActivationLock(v bool)`

SetIsPreventActivationLock sets IsPreventActivationLock field to given value.


### GetIsEnableDeviceBasedActivationLock

`func (o *PutMobileDevicePrestage) GetIsEnableDeviceBasedActivationLock() bool`

GetIsEnableDeviceBasedActivationLock returns the IsEnableDeviceBasedActivationLock field if non-nil, zero value otherwise.

### GetIsEnableDeviceBasedActivationLockOk

`func (o *PutMobileDevicePrestage) GetIsEnableDeviceBasedActivationLockOk() (*bool, bool)`

GetIsEnableDeviceBasedActivationLockOk returns a tuple with the IsEnableDeviceBasedActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnableDeviceBasedActivationLock

`func (o *PutMobileDevicePrestage) SetIsEnableDeviceBasedActivationLock(v bool)`

SetIsEnableDeviceBasedActivationLock sets IsEnableDeviceBasedActivationLock field to given value.


### GetDeviceEnrollmentProgramInstanceId

`func (o *PutMobileDevicePrestage) GetDeviceEnrollmentProgramInstanceId() int32`

GetDeviceEnrollmentProgramInstanceId returns the DeviceEnrollmentProgramInstanceId field if non-nil, zero value otherwise.

### GetDeviceEnrollmentProgramInstanceIdOk

`func (o *PutMobileDevicePrestage) GetDeviceEnrollmentProgramInstanceIdOk() (*int32, bool)`

GetDeviceEnrollmentProgramInstanceIdOk returns a tuple with the DeviceEnrollmentProgramInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceEnrollmentProgramInstanceId

`func (o *PutMobileDevicePrestage) SetDeviceEnrollmentProgramInstanceId(v int32)`

SetDeviceEnrollmentProgramInstanceId sets DeviceEnrollmentProgramInstanceId field to given value.


### GetSkipSetupItems

`func (o *PutMobileDevicePrestage) GetSkipSetupItems() map[string]bool`

GetSkipSetupItems returns the SkipSetupItems field if non-nil, zero value otherwise.

### GetSkipSetupItemsOk

`func (o *PutMobileDevicePrestage) GetSkipSetupItemsOk() (*map[string]bool, bool)`

GetSkipSetupItemsOk returns a tuple with the SkipSetupItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSetupItems

`func (o *PutMobileDevicePrestage) SetSkipSetupItems(v map[string]bool)`

SetSkipSetupItems sets SkipSetupItems field to given value.

### HasSkipSetupItems

`func (o *PutMobileDevicePrestage) HasSkipSetupItems() bool`

HasSkipSetupItems returns a boolean if a field has been set.

### GetLocationInformation

`func (o *PutMobileDevicePrestage) GetLocationInformation() LocationInformation`

GetLocationInformation returns the LocationInformation field if non-nil, zero value otherwise.

### GetLocationInformationOk

`func (o *PutMobileDevicePrestage) GetLocationInformationOk() (*LocationInformation, bool)`

GetLocationInformationOk returns a tuple with the LocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationInformation

`func (o *PutMobileDevicePrestage) SetLocationInformation(v LocationInformation)`

SetLocationInformation sets LocationInformation field to given value.


### GetPurchasingInformation

`func (o *PutMobileDevicePrestage) GetPurchasingInformation() PrestagePurchasingInformation`

GetPurchasingInformation returns the PurchasingInformation field if non-nil, zero value otherwise.

### GetPurchasingInformationOk

`func (o *PutMobileDevicePrestage) GetPurchasingInformationOk() (*PrestagePurchasingInformation, bool)`

GetPurchasingInformationOk returns a tuple with the PurchasingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingInformation

`func (o *PutMobileDevicePrestage) SetPurchasingInformation(v PrestagePurchasingInformation)`

SetPurchasingInformation sets PurchasingInformation field to given value.


### GetAnchorCertificates

`func (o *PutMobileDevicePrestage) GetAnchorCertificates() []string`

GetAnchorCertificates returns the AnchorCertificates field if non-nil, zero value otherwise.

### GetAnchorCertificatesOk

`func (o *PutMobileDevicePrestage) GetAnchorCertificatesOk() (*[]string, bool)`

GetAnchorCertificatesOk returns a tuple with the AnchorCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorCertificates

`func (o *PutMobileDevicePrestage) SetAnchorCertificates(v []string)`

SetAnchorCertificates sets AnchorCertificates field to given value.

### HasAnchorCertificates

`func (o *PutMobileDevicePrestage) HasAnchorCertificates() bool`

HasAnchorCertificates returns a boolean if a field has been set.

### GetEnrollmentCustomizationId

`func (o *PutMobileDevicePrestage) GetEnrollmentCustomizationId() int32`

GetEnrollmentCustomizationId returns the EnrollmentCustomizationId field if non-nil, zero value otherwise.

### GetEnrollmentCustomizationIdOk

`func (o *PutMobileDevicePrestage) GetEnrollmentCustomizationIdOk() (*int32, bool)`

GetEnrollmentCustomizationIdOk returns a tuple with the EnrollmentCustomizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCustomizationId

`func (o *PutMobileDevicePrestage) SetEnrollmentCustomizationId(v int32)`

SetEnrollmentCustomizationId sets EnrollmentCustomizationId field to given value.

### HasEnrollmentCustomizationId

`func (o *PutMobileDevicePrestage) HasEnrollmentCustomizationId() bool`

HasEnrollmentCustomizationId returns a boolean if a field has been set.

### GetIsAllowPairing

`func (o *PutMobileDevicePrestage) GetIsAllowPairing() bool`

GetIsAllowPairing returns the IsAllowPairing field if non-nil, zero value otherwise.

### GetIsAllowPairingOk

`func (o *PutMobileDevicePrestage) GetIsAllowPairingOk() (*bool, bool)`

GetIsAllowPairingOk returns a tuple with the IsAllowPairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowPairing

`func (o *PutMobileDevicePrestage) SetIsAllowPairing(v bool)`

SetIsAllowPairing sets IsAllowPairing field to given value.


### GetIsMultiUser

`func (o *PutMobileDevicePrestage) GetIsMultiUser() bool`

GetIsMultiUser returns the IsMultiUser field if non-nil, zero value otherwise.

### GetIsMultiUserOk

`func (o *PutMobileDevicePrestage) GetIsMultiUserOk() (*bool, bool)`

GetIsMultiUserOk returns a tuple with the IsMultiUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiUser

`func (o *PutMobileDevicePrestage) SetIsMultiUser(v bool)`

SetIsMultiUser sets IsMultiUser field to given value.


### GetIsSupervised

`func (o *PutMobileDevicePrestage) GetIsSupervised() bool`

GetIsSupervised returns the IsSupervised field if non-nil, zero value otherwise.

### GetIsSupervisedOk

`func (o *PutMobileDevicePrestage) GetIsSupervisedOk() (*bool, bool)`

GetIsSupervisedOk returns a tuple with the IsSupervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupervised

`func (o *PutMobileDevicePrestage) SetIsSupervised(v bool)`

SetIsSupervised sets IsSupervised field to given value.


### GetMaximumSharedAccounts

`func (o *PutMobileDevicePrestage) GetMaximumSharedAccounts() int32`

GetMaximumSharedAccounts returns the MaximumSharedAccounts field if non-nil, zero value otherwise.

### GetMaximumSharedAccountsOk

`func (o *PutMobileDevicePrestage) GetMaximumSharedAccountsOk() (*int32, bool)`

GetMaximumSharedAccountsOk returns a tuple with the MaximumSharedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSharedAccounts

`func (o *PutMobileDevicePrestage) SetMaximumSharedAccounts(v int32)`

SetMaximumSharedAccounts sets MaximumSharedAccounts field to given value.


### GetIsAutoAdvanceSetup

`func (o *PutMobileDevicePrestage) GetIsAutoAdvanceSetup() bool`

GetIsAutoAdvanceSetup returns the IsAutoAdvanceSetup field if non-nil, zero value otherwise.

### GetIsAutoAdvanceSetupOk

`func (o *PutMobileDevicePrestage) GetIsAutoAdvanceSetupOk() (*bool, bool)`

GetIsAutoAdvanceSetupOk returns a tuple with the IsAutoAdvanceSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoAdvanceSetup

`func (o *PutMobileDevicePrestage) SetIsAutoAdvanceSetup(v bool)`

SetIsAutoAdvanceSetup sets IsAutoAdvanceSetup field to given value.


### GetIsConfigureDeviceBeforeSetupAssistant

`func (o *PutMobileDevicePrestage) GetIsConfigureDeviceBeforeSetupAssistant() bool`

GetIsConfigureDeviceBeforeSetupAssistant returns the IsConfigureDeviceBeforeSetupAssistant field if non-nil, zero value otherwise.

### GetIsConfigureDeviceBeforeSetupAssistantOk

`func (o *PutMobileDevicePrestage) GetIsConfigureDeviceBeforeSetupAssistantOk() (*bool, bool)`

GetIsConfigureDeviceBeforeSetupAssistantOk returns a tuple with the IsConfigureDeviceBeforeSetupAssistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigureDeviceBeforeSetupAssistant

`func (o *PutMobileDevicePrestage) SetIsConfigureDeviceBeforeSetupAssistant(v bool)`

SetIsConfigureDeviceBeforeSetupAssistant sets IsConfigureDeviceBeforeSetupAssistant field to given value.


### GetLanguage

`func (o *PutMobileDevicePrestage) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PutMobileDevicePrestage) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PutMobileDevicePrestage) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PutMobileDevicePrestage) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetRegion

`func (o *PutMobileDevicePrestage) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PutMobileDevicePrestage) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PutMobileDevicePrestage) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PutMobileDevicePrestage) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetNames

`func (o *PutMobileDevicePrestage) GetNames() MobileDevicePrestageNames`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *PutMobileDevicePrestage) GetNamesOk() (*MobileDevicePrestageNames, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *PutMobileDevicePrestage) SetNames(v MobileDevicePrestageNames)`

SetNames sets Names field to given value.

### HasNames

`func (o *PutMobileDevicePrestage) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetVersionLock

`func (o *PutMobileDevicePrestage) GetVersionLock() int32`

GetVersionLock returns the VersionLock field if non-nil, zero value otherwise.

### GetVersionLockOk

`func (o *PutMobileDevicePrestage) GetVersionLockOk() (*int32, bool)`

GetVersionLockOk returns a tuple with the VersionLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionLock

`func (o *PutMobileDevicePrestage) SetVersionLock(v int32)`

SetVersionLock sets VersionLock field to given value.

### HasVersionLock

`func (o *PutMobileDevicePrestage) HasVersionLock() bool`

HasVersionLock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


