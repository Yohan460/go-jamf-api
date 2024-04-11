# MobileDevicePrestage

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
**EnrollmentSiteId** | **int64** |  | 
**IsKeepExistingSiteMembership** | **bool** |  | 
**IsKeepExistingLocationInformation** | **bool** |  | 
**IsRequireAuthentication** | **bool** |  | 
**AuthenticationPrompt** | **string** |  | 
**IsPreventActivationLock** | **bool** |  | 
**IsEnableDeviceBasedActivationLock** | **bool** |  | 
**DeviceEnrollmentProgramInstanceId** | **int64** |  | 
**SkipSetupItems** | Pointer to **map[string]bool** |  | [optional] 
**LocationInformation** | [**LocationInformation**](LocationInformation.md) |  | 
**PurchasingInformation** | [**PrestagePurchasingInformation**](PrestagePurchasingInformation.md) |  | 
**AnchorCertificates** | Pointer to **[]string** | The Base64 encoded PEM Certificate | [optional] 
**EnrollmentCustomizationId** | Pointer to **int64** |  | [optional] 
**IsAllowPairing** | **bool** |  | 
**IsMultiUser** | **bool** |  | 
**IsSupervised** | **bool** |  | 
**MaximumSharedAccounts** | **int64** |  | 
**IsAutoAdvanceSetup** | **bool** |  | 
**IsConfigureDeviceBeforeSetupAssistant** | **bool** |  | 
**Language** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Names** | Pointer to [**MobileDevicePrestageNames**](MobileDevicePrestageNames.md) |  | [optional] 

## Methods

### NewMobileDevicePrestage

`func NewMobileDevicePrestage(displayName string, isMandatory bool, isMdmRemovable bool, supportPhoneNumber string, supportEmailAddress string, department string, isDefaultPrestage bool, enrollmentSiteId int64, isKeepExistingSiteMembership bool, isKeepExistingLocationInformation bool, isRequireAuthentication bool, authenticationPrompt string, isPreventActivationLock bool, isEnableDeviceBasedActivationLock bool, deviceEnrollmentProgramInstanceId int64, locationInformation LocationInformation, purchasingInformation PrestagePurchasingInformation, isAllowPairing bool, isMultiUser bool, isSupervised bool, maximumSharedAccounts int64, isAutoAdvanceSetup bool, isConfigureDeviceBeforeSetupAssistant bool, ) *MobileDevicePrestage`

NewMobileDevicePrestage instantiates a new MobileDevicePrestage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDevicePrestageWithDefaults

`func NewMobileDevicePrestageWithDefaults() *MobileDevicePrestage`

NewMobileDevicePrestageWithDefaults instantiates a new MobileDevicePrestage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *MobileDevicePrestage) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MobileDevicePrestage) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MobileDevicePrestage) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetIsMandatory

`func (o *MobileDevicePrestage) GetIsMandatory() bool`

GetIsMandatory returns the IsMandatory field if non-nil, zero value otherwise.

### GetIsMandatoryOk

`func (o *MobileDevicePrestage) GetIsMandatoryOk() (*bool, bool)`

GetIsMandatoryOk returns a tuple with the IsMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMandatory

`func (o *MobileDevicePrestage) SetIsMandatory(v bool)`

SetIsMandatory sets IsMandatory field to given value.


### GetIsMdmRemovable

`func (o *MobileDevicePrestage) GetIsMdmRemovable() bool`

GetIsMdmRemovable returns the IsMdmRemovable field if non-nil, zero value otherwise.

### GetIsMdmRemovableOk

`func (o *MobileDevicePrestage) GetIsMdmRemovableOk() (*bool, bool)`

GetIsMdmRemovableOk returns a tuple with the IsMdmRemovable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMdmRemovable

`func (o *MobileDevicePrestage) SetIsMdmRemovable(v bool)`

SetIsMdmRemovable sets IsMdmRemovable field to given value.


### GetSupportPhoneNumber

`func (o *MobileDevicePrestage) GetSupportPhoneNumber() string`

GetSupportPhoneNumber returns the SupportPhoneNumber field if non-nil, zero value otherwise.

### GetSupportPhoneNumberOk

`func (o *MobileDevicePrestage) GetSupportPhoneNumberOk() (*string, bool)`

GetSupportPhoneNumberOk returns a tuple with the SupportPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPhoneNumber

`func (o *MobileDevicePrestage) SetSupportPhoneNumber(v string)`

SetSupportPhoneNumber sets SupportPhoneNumber field to given value.


### GetSupportEmailAddress

`func (o *MobileDevicePrestage) GetSupportEmailAddress() string`

GetSupportEmailAddress returns the SupportEmailAddress field if non-nil, zero value otherwise.

### GetSupportEmailAddressOk

`func (o *MobileDevicePrestage) GetSupportEmailAddressOk() (*string, bool)`

GetSupportEmailAddressOk returns a tuple with the SupportEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmailAddress

`func (o *MobileDevicePrestage) SetSupportEmailAddress(v string)`

SetSupportEmailAddress sets SupportEmailAddress field to given value.


### GetDepartment

`func (o *MobileDevicePrestage) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *MobileDevicePrestage) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *MobileDevicePrestage) SetDepartment(v string)`

SetDepartment sets Department field to given value.


### GetIsDefaultPrestage

`func (o *MobileDevicePrestage) GetIsDefaultPrestage() bool`

GetIsDefaultPrestage returns the IsDefaultPrestage field if non-nil, zero value otherwise.

### GetIsDefaultPrestageOk

`func (o *MobileDevicePrestage) GetIsDefaultPrestageOk() (*bool, bool)`

GetIsDefaultPrestageOk returns a tuple with the IsDefaultPrestage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultPrestage

`func (o *MobileDevicePrestage) SetIsDefaultPrestage(v bool)`

SetIsDefaultPrestage sets IsDefaultPrestage field to given value.


### GetEnrollmentSiteId

`func (o *MobileDevicePrestage) GetEnrollmentSiteId() int64`

GetEnrollmentSiteId returns the EnrollmentSiteId field if non-nil, zero value otherwise.

### GetEnrollmentSiteIdOk

`func (o *MobileDevicePrestage) GetEnrollmentSiteIdOk() (*int64, bool)`

GetEnrollmentSiteIdOk returns a tuple with the EnrollmentSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSiteId

`func (o *MobileDevicePrestage) SetEnrollmentSiteId(v int64)`

SetEnrollmentSiteId sets EnrollmentSiteId field to given value.


### GetIsKeepExistingSiteMembership

`func (o *MobileDevicePrestage) GetIsKeepExistingSiteMembership() bool`

GetIsKeepExistingSiteMembership returns the IsKeepExistingSiteMembership field if non-nil, zero value otherwise.

### GetIsKeepExistingSiteMembershipOk

`func (o *MobileDevicePrestage) GetIsKeepExistingSiteMembershipOk() (*bool, bool)`

GetIsKeepExistingSiteMembershipOk returns a tuple with the IsKeepExistingSiteMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKeepExistingSiteMembership

`func (o *MobileDevicePrestage) SetIsKeepExistingSiteMembership(v bool)`

SetIsKeepExistingSiteMembership sets IsKeepExistingSiteMembership field to given value.


### GetIsKeepExistingLocationInformation

`func (o *MobileDevicePrestage) GetIsKeepExistingLocationInformation() bool`

GetIsKeepExistingLocationInformation returns the IsKeepExistingLocationInformation field if non-nil, zero value otherwise.

### GetIsKeepExistingLocationInformationOk

`func (o *MobileDevicePrestage) GetIsKeepExistingLocationInformationOk() (*bool, bool)`

GetIsKeepExistingLocationInformationOk returns a tuple with the IsKeepExistingLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKeepExistingLocationInformation

`func (o *MobileDevicePrestage) SetIsKeepExistingLocationInformation(v bool)`

SetIsKeepExistingLocationInformation sets IsKeepExistingLocationInformation field to given value.


### GetIsRequireAuthentication

`func (o *MobileDevicePrestage) GetIsRequireAuthentication() bool`

GetIsRequireAuthentication returns the IsRequireAuthentication field if non-nil, zero value otherwise.

### GetIsRequireAuthenticationOk

`func (o *MobileDevicePrestage) GetIsRequireAuthenticationOk() (*bool, bool)`

GetIsRequireAuthenticationOk returns a tuple with the IsRequireAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequireAuthentication

`func (o *MobileDevicePrestage) SetIsRequireAuthentication(v bool)`

SetIsRequireAuthentication sets IsRequireAuthentication field to given value.


### GetAuthenticationPrompt

`func (o *MobileDevicePrestage) GetAuthenticationPrompt() string`

GetAuthenticationPrompt returns the AuthenticationPrompt field if non-nil, zero value otherwise.

### GetAuthenticationPromptOk

`func (o *MobileDevicePrestage) GetAuthenticationPromptOk() (*string, bool)`

GetAuthenticationPromptOk returns a tuple with the AuthenticationPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPrompt

`func (o *MobileDevicePrestage) SetAuthenticationPrompt(v string)`

SetAuthenticationPrompt sets AuthenticationPrompt field to given value.


### GetIsPreventActivationLock

`func (o *MobileDevicePrestage) GetIsPreventActivationLock() bool`

GetIsPreventActivationLock returns the IsPreventActivationLock field if non-nil, zero value otherwise.

### GetIsPreventActivationLockOk

`func (o *MobileDevicePrestage) GetIsPreventActivationLockOk() (*bool, bool)`

GetIsPreventActivationLockOk returns a tuple with the IsPreventActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreventActivationLock

`func (o *MobileDevicePrestage) SetIsPreventActivationLock(v bool)`

SetIsPreventActivationLock sets IsPreventActivationLock field to given value.


### GetIsEnableDeviceBasedActivationLock

`func (o *MobileDevicePrestage) GetIsEnableDeviceBasedActivationLock() bool`

GetIsEnableDeviceBasedActivationLock returns the IsEnableDeviceBasedActivationLock field if non-nil, zero value otherwise.

### GetIsEnableDeviceBasedActivationLockOk

`func (o *MobileDevicePrestage) GetIsEnableDeviceBasedActivationLockOk() (*bool, bool)`

GetIsEnableDeviceBasedActivationLockOk returns a tuple with the IsEnableDeviceBasedActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnableDeviceBasedActivationLock

`func (o *MobileDevicePrestage) SetIsEnableDeviceBasedActivationLock(v bool)`

SetIsEnableDeviceBasedActivationLock sets IsEnableDeviceBasedActivationLock field to given value.


### GetDeviceEnrollmentProgramInstanceId

`func (o *MobileDevicePrestage) GetDeviceEnrollmentProgramInstanceId() int64`

GetDeviceEnrollmentProgramInstanceId returns the DeviceEnrollmentProgramInstanceId field if non-nil, zero value otherwise.

### GetDeviceEnrollmentProgramInstanceIdOk

`func (o *MobileDevicePrestage) GetDeviceEnrollmentProgramInstanceIdOk() (*int64, bool)`

GetDeviceEnrollmentProgramInstanceIdOk returns a tuple with the DeviceEnrollmentProgramInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceEnrollmentProgramInstanceId

`func (o *MobileDevicePrestage) SetDeviceEnrollmentProgramInstanceId(v int64)`

SetDeviceEnrollmentProgramInstanceId sets DeviceEnrollmentProgramInstanceId field to given value.


### GetSkipSetupItems

`func (o *MobileDevicePrestage) GetSkipSetupItems() map[string]bool`

GetSkipSetupItems returns the SkipSetupItems field if non-nil, zero value otherwise.

### GetSkipSetupItemsOk

`func (o *MobileDevicePrestage) GetSkipSetupItemsOk() (*map[string]bool, bool)`

GetSkipSetupItemsOk returns a tuple with the SkipSetupItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSetupItems

`func (o *MobileDevicePrestage) SetSkipSetupItems(v map[string]bool)`

SetSkipSetupItems sets SkipSetupItems field to given value.

### HasSkipSetupItems

`func (o *MobileDevicePrestage) HasSkipSetupItems() bool`

HasSkipSetupItems returns a boolean if a field has been set.

### GetLocationInformation

`func (o *MobileDevicePrestage) GetLocationInformation() LocationInformation`

GetLocationInformation returns the LocationInformation field if non-nil, zero value otherwise.

### GetLocationInformationOk

`func (o *MobileDevicePrestage) GetLocationInformationOk() (*LocationInformation, bool)`

GetLocationInformationOk returns a tuple with the LocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationInformation

`func (o *MobileDevicePrestage) SetLocationInformation(v LocationInformation)`

SetLocationInformation sets LocationInformation field to given value.


### GetPurchasingInformation

`func (o *MobileDevicePrestage) GetPurchasingInformation() PrestagePurchasingInformation`

GetPurchasingInformation returns the PurchasingInformation field if non-nil, zero value otherwise.

### GetPurchasingInformationOk

`func (o *MobileDevicePrestage) GetPurchasingInformationOk() (*PrestagePurchasingInformation, bool)`

GetPurchasingInformationOk returns a tuple with the PurchasingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingInformation

`func (o *MobileDevicePrestage) SetPurchasingInformation(v PrestagePurchasingInformation)`

SetPurchasingInformation sets PurchasingInformation field to given value.


### GetAnchorCertificates

`func (o *MobileDevicePrestage) GetAnchorCertificates() []string`

GetAnchorCertificates returns the AnchorCertificates field if non-nil, zero value otherwise.

### GetAnchorCertificatesOk

`func (o *MobileDevicePrestage) GetAnchorCertificatesOk() (*[]string, bool)`

GetAnchorCertificatesOk returns a tuple with the AnchorCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorCertificates

`func (o *MobileDevicePrestage) SetAnchorCertificates(v []string)`

SetAnchorCertificates sets AnchorCertificates field to given value.

### HasAnchorCertificates

`func (o *MobileDevicePrestage) HasAnchorCertificates() bool`

HasAnchorCertificates returns a boolean if a field has been set.

### GetEnrollmentCustomizationId

`func (o *MobileDevicePrestage) GetEnrollmentCustomizationId() int64`

GetEnrollmentCustomizationId returns the EnrollmentCustomizationId field if non-nil, zero value otherwise.

### GetEnrollmentCustomizationIdOk

`func (o *MobileDevicePrestage) GetEnrollmentCustomizationIdOk() (*int64, bool)`

GetEnrollmentCustomizationIdOk returns a tuple with the EnrollmentCustomizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCustomizationId

`func (o *MobileDevicePrestage) SetEnrollmentCustomizationId(v int64)`

SetEnrollmentCustomizationId sets EnrollmentCustomizationId field to given value.

### HasEnrollmentCustomizationId

`func (o *MobileDevicePrestage) HasEnrollmentCustomizationId() bool`

HasEnrollmentCustomizationId returns a boolean if a field has been set.

### GetIsAllowPairing

`func (o *MobileDevicePrestage) GetIsAllowPairing() bool`

GetIsAllowPairing returns the IsAllowPairing field if non-nil, zero value otherwise.

### GetIsAllowPairingOk

`func (o *MobileDevicePrestage) GetIsAllowPairingOk() (*bool, bool)`

GetIsAllowPairingOk returns a tuple with the IsAllowPairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowPairing

`func (o *MobileDevicePrestage) SetIsAllowPairing(v bool)`

SetIsAllowPairing sets IsAllowPairing field to given value.


### GetIsMultiUser

`func (o *MobileDevicePrestage) GetIsMultiUser() bool`

GetIsMultiUser returns the IsMultiUser field if non-nil, zero value otherwise.

### GetIsMultiUserOk

`func (o *MobileDevicePrestage) GetIsMultiUserOk() (*bool, bool)`

GetIsMultiUserOk returns a tuple with the IsMultiUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiUser

`func (o *MobileDevicePrestage) SetIsMultiUser(v bool)`

SetIsMultiUser sets IsMultiUser field to given value.


### GetIsSupervised

`func (o *MobileDevicePrestage) GetIsSupervised() bool`

GetIsSupervised returns the IsSupervised field if non-nil, zero value otherwise.

### GetIsSupervisedOk

`func (o *MobileDevicePrestage) GetIsSupervisedOk() (*bool, bool)`

GetIsSupervisedOk returns a tuple with the IsSupervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupervised

`func (o *MobileDevicePrestage) SetIsSupervised(v bool)`

SetIsSupervised sets IsSupervised field to given value.


### GetMaximumSharedAccounts

`func (o *MobileDevicePrestage) GetMaximumSharedAccounts() int64`

GetMaximumSharedAccounts returns the MaximumSharedAccounts field if non-nil, zero value otherwise.

### GetMaximumSharedAccountsOk

`func (o *MobileDevicePrestage) GetMaximumSharedAccountsOk() (*int64, bool)`

GetMaximumSharedAccountsOk returns a tuple with the MaximumSharedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSharedAccounts

`func (o *MobileDevicePrestage) SetMaximumSharedAccounts(v int64)`

SetMaximumSharedAccounts sets MaximumSharedAccounts field to given value.


### GetIsAutoAdvanceSetup

`func (o *MobileDevicePrestage) GetIsAutoAdvanceSetup() bool`

GetIsAutoAdvanceSetup returns the IsAutoAdvanceSetup field if non-nil, zero value otherwise.

### GetIsAutoAdvanceSetupOk

`func (o *MobileDevicePrestage) GetIsAutoAdvanceSetupOk() (*bool, bool)`

GetIsAutoAdvanceSetupOk returns a tuple with the IsAutoAdvanceSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoAdvanceSetup

`func (o *MobileDevicePrestage) SetIsAutoAdvanceSetup(v bool)`

SetIsAutoAdvanceSetup sets IsAutoAdvanceSetup field to given value.


### GetIsConfigureDeviceBeforeSetupAssistant

`func (o *MobileDevicePrestage) GetIsConfigureDeviceBeforeSetupAssistant() bool`

GetIsConfigureDeviceBeforeSetupAssistant returns the IsConfigureDeviceBeforeSetupAssistant field if non-nil, zero value otherwise.

### GetIsConfigureDeviceBeforeSetupAssistantOk

`func (o *MobileDevicePrestage) GetIsConfigureDeviceBeforeSetupAssistantOk() (*bool, bool)`

GetIsConfigureDeviceBeforeSetupAssistantOk returns a tuple with the IsConfigureDeviceBeforeSetupAssistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigureDeviceBeforeSetupAssistant

`func (o *MobileDevicePrestage) SetIsConfigureDeviceBeforeSetupAssistant(v bool)`

SetIsConfigureDeviceBeforeSetupAssistant sets IsConfigureDeviceBeforeSetupAssistant field to given value.


### GetLanguage

`func (o *MobileDevicePrestage) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MobileDevicePrestage) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *MobileDevicePrestage) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *MobileDevicePrestage) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetRegion

`func (o *MobileDevicePrestage) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MobileDevicePrestage) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MobileDevicePrestage) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MobileDevicePrestage) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetNames

`func (o *MobileDevicePrestage) GetNames() MobileDevicePrestageNames`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *MobileDevicePrestage) GetNamesOk() (*MobileDevicePrestageNames, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *MobileDevicePrestage) SetNames(v MobileDevicePrestageNames)`

SetNames sets Names field to given value.

### HasNames

`func (o *MobileDevicePrestage) HasNames() bool`

HasNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


