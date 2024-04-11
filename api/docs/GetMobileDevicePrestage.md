# GetMobileDevicePrestage

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
**Id** | Pointer to **int64** |  | [optional] 
**ProfileUUID** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**VersionLock** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetMobileDevicePrestage

`func NewGetMobileDevicePrestage(displayName string, isMandatory bool, isMdmRemovable bool, supportPhoneNumber string, supportEmailAddress string, department string, isDefaultPrestage bool, enrollmentSiteId int64, isKeepExistingSiteMembership bool, isKeepExistingLocationInformation bool, isRequireAuthentication bool, authenticationPrompt string, isPreventActivationLock bool, isEnableDeviceBasedActivationLock bool, deviceEnrollmentProgramInstanceId int64, locationInformation LocationInformation, purchasingInformation PrestagePurchasingInformation, isAllowPairing bool, isMultiUser bool, isSupervised bool, maximumSharedAccounts int64, isAutoAdvanceSetup bool, isConfigureDeviceBeforeSetupAssistant bool, ) *GetMobileDevicePrestage`

NewGetMobileDevicePrestage instantiates a new GetMobileDevicePrestage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMobileDevicePrestageWithDefaults

`func NewGetMobileDevicePrestageWithDefaults() *GetMobileDevicePrestage`

NewGetMobileDevicePrestageWithDefaults instantiates a new GetMobileDevicePrestage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *GetMobileDevicePrestage) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetMobileDevicePrestage) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetMobileDevicePrestage) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetIsMandatory

`func (o *GetMobileDevicePrestage) GetIsMandatory() bool`

GetIsMandatory returns the IsMandatory field if non-nil, zero value otherwise.

### GetIsMandatoryOk

`func (o *GetMobileDevicePrestage) GetIsMandatoryOk() (*bool, bool)`

GetIsMandatoryOk returns a tuple with the IsMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMandatory

`func (o *GetMobileDevicePrestage) SetIsMandatory(v bool)`

SetIsMandatory sets IsMandatory field to given value.


### GetIsMdmRemovable

`func (o *GetMobileDevicePrestage) GetIsMdmRemovable() bool`

GetIsMdmRemovable returns the IsMdmRemovable field if non-nil, zero value otherwise.

### GetIsMdmRemovableOk

`func (o *GetMobileDevicePrestage) GetIsMdmRemovableOk() (*bool, bool)`

GetIsMdmRemovableOk returns a tuple with the IsMdmRemovable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMdmRemovable

`func (o *GetMobileDevicePrestage) SetIsMdmRemovable(v bool)`

SetIsMdmRemovable sets IsMdmRemovable field to given value.


### GetSupportPhoneNumber

`func (o *GetMobileDevicePrestage) GetSupportPhoneNumber() string`

GetSupportPhoneNumber returns the SupportPhoneNumber field if non-nil, zero value otherwise.

### GetSupportPhoneNumberOk

`func (o *GetMobileDevicePrestage) GetSupportPhoneNumberOk() (*string, bool)`

GetSupportPhoneNumberOk returns a tuple with the SupportPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPhoneNumber

`func (o *GetMobileDevicePrestage) SetSupportPhoneNumber(v string)`

SetSupportPhoneNumber sets SupportPhoneNumber field to given value.


### GetSupportEmailAddress

`func (o *GetMobileDevicePrestage) GetSupportEmailAddress() string`

GetSupportEmailAddress returns the SupportEmailAddress field if non-nil, zero value otherwise.

### GetSupportEmailAddressOk

`func (o *GetMobileDevicePrestage) GetSupportEmailAddressOk() (*string, bool)`

GetSupportEmailAddressOk returns a tuple with the SupportEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmailAddress

`func (o *GetMobileDevicePrestage) SetSupportEmailAddress(v string)`

SetSupportEmailAddress sets SupportEmailAddress field to given value.


### GetDepartment

`func (o *GetMobileDevicePrestage) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *GetMobileDevicePrestage) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *GetMobileDevicePrestage) SetDepartment(v string)`

SetDepartment sets Department field to given value.


### GetIsDefaultPrestage

`func (o *GetMobileDevicePrestage) GetIsDefaultPrestage() bool`

GetIsDefaultPrestage returns the IsDefaultPrestage field if non-nil, zero value otherwise.

### GetIsDefaultPrestageOk

`func (o *GetMobileDevicePrestage) GetIsDefaultPrestageOk() (*bool, bool)`

GetIsDefaultPrestageOk returns a tuple with the IsDefaultPrestage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultPrestage

`func (o *GetMobileDevicePrestage) SetIsDefaultPrestage(v bool)`

SetIsDefaultPrestage sets IsDefaultPrestage field to given value.


### GetEnrollmentSiteId

`func (o *GetMobileDevicePrestage) GetEnrollmentSiteId() int64`

GetEnrollmentSiteId returns the EnrollmentSiteId field if non-nil, zero value otherwise.

### GetEnrollmentSiteIdOk

`func (o *GetMobileDevicePrestage) GetEnrollmentSiteIdOk() (*int64, bool)`

GetEnrollmentSiteIdOk returns a tuple with the EnrollmentSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSiteId

`func (o *GetMobileDevicePrestage) SetEnrollmentSiteId(v int64)`

SetEnrollmentSiteId sets EnrollmentSiteId field to given value.


### GetIsKeepExistingSiteMembership

`func (o *GetMobileDevicePrestage) GetIsKeepExistingSiteMembership() bool`

GetIsKeepExistingSiteMembership returns the IsKeepExistingSiteMembership field if non-nil, zero value otherwise.

### GetIsKeepExistingSiteMembershipOk

`func (o *GetMobileDevicePrestage) GetIsKeepExistingSiteMembershipOk() (*bool, bool)`

GetIsKeepExistingSiteMembershipOk returns a tuple with the IsKeepExistingSiteMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKeepExistingSiteMembership

`func (o *GetMobileDevicePrestage) SetIsKeepExistingSiteMembership(v bool)`

SetIsKeepExistingSiteMembership sets IsKeepExistingSiteMembership field to given value.


### GetIsKeepExistingLocationInformation

`func (o *GetMobileDevicePrestage) GetIsKeepExistingLocationInformation() bool`

GetIsKeepExistingLocationInformation returns the IsKeepExistingLocationInformation field if non-nil, zero value otherwise.

### GetIsKeepExistingLocationInformationOk

`func (o *GetMobileDevicePrestage) GetIsKeepExistingLocationInformationOk() (*bool, bool)`

GetIsKeepExistingLocationInformationOk returns a tuple with the IsKeepExistingLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKeepExistingLocationInformation

`func (o *GetMobileDevicePrestage) SetIsKeepExistingLocationInformation(v bool)`

SetIsKeepExistingLocationInformation sets IsKeepExistingLocationInformation field to given value.


### GetIsRequireAuthentication

`func (o *GetMobileDevicePrestage) GetIsRequireAuthentication() bool`

GetIsRequireAuthentication returns the IsRequireAuthentication field if non-nil, zero value otherwise.

### GetIsRequireAuthenticationOk

`func (o *GetMobileDevicePrestage) GetIsRequireAuthenticationOk() (*bool, bool)`

GetIsRequireAuthenticationOk returns a tuple with the IsRequireAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequireAuthentication

`func (o *GetMobileDevicePrestage) SetIsRequireAuthentication(v bool)`

SetIsRequireAuthentication sets IsRequireAuthentication field to given value.


### GetAuthenticationPrompt

`func (o *GetMobileDevicePrestage) GetAuthenticationPrompt() string`

GetAuthenticationPrompt returns the AuthenticationPrompt field if non-nil, zero value otherwise.

### GetAuthenticationPromptOk

`func (o *GetMobileDevicePrestage) GetAuthenticationPromptOk() (*string, bool)`

GetAuthenticationPromptOk returns a tuple with the AuthenticationPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPrompt

`func (o *GetMobileDevicePrestage) SetAuthenticationPrompt(v string)`

SetAuthenticationPrompt sets AuthenticationPrompt field to given value.


### GetIsPreventActivationLock

`func (o *GetMobileDevicePrestage) GetIsPreventActivationLock() bool`

GetIsPreventActivationLock returns the IsPreventActivationLock field if non-nil, zero value otherwise.

### GetIsPreventActivationLockOk

`func (o *GetMobileDevicePrestage) GetIsPreventActivationLockOk() (*bool, bool)`

GetIsPreventActivationLockOk returns a tuple with the IsPreventActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreventActivationLock

`func (o *GetMobileDevicePrestage) SetIsPreventActivationLock(v bool)`

SetIsPreventActivationLock sets IsPreventActivationLock field to given value.


### GetIsEnableDeviceBasedActivationLock

`func (o *GetMobileDevicePrestage) GetIsEnableDeviceBasedActivationLock() bool`

GetIsEnableDeviceBasedActivationLock returns the IsEnableDeviceBasedActivationLock field if non-nil, zero value otherwise.

### GetIsEnableDeviceBasedActivationLockOk

`func (o *GetMobileDevicePrestage) GetIsEnableDeviceBasedActivationLockOk() (*bool, bool)`

GetIsEnableDeviceBasedActivationLockOk returns a tuple with the IsEnableDeviceBasedActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnableDeviceBasedActivationLock

`func (o *GetMobileDevicePrestage) SetIsEnableDeviceBasedActivationLock(v bool)`

SetIsEnableDeviceBasedActivationLock sets IsEnableDeviceBasedActivationLock field to given value.


### GetDeviceEnrollmentProgramInstanceId

`func (o *GetMobileDevicePrestage) GetDeviceEnrollmentProgramInstanceId() int64`

GetDeviceEnrollmentProgramInstanceId returns the DeviceEnrollmentProgramInstanceId field if non-nil, zero value otherwise.

### GetDeviceEnrollmentProgramInstanceIdOk

`func (o *GetMobileDevicePrestage) GetDeviceEnrollmentProgramInstanceIdOk() (*int64, bool)`

GetDeviceEnrollmentProgramInstanceIdOk returns a tuple with the DeviceEnrollmentProgramInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceEnrollmentProgramInstanceId

`func (o *GetMobileDevicePrestage) SetDeviceEnrollmentProgramInstanceId(v int64)`

SetDeviceEnrollmentProgramInstanceId sets DeviceEnrollmentProgramInstanceId field to given value.


### GetSkipSetupItems

`func (o *GetMobileDevicePrestage) GetSkipSetupItems() map[string]bool`

GetSkipSetupItems returns the SkipSetupItems field if non-nil, zero value otherwise.

### GetSkipSetupItemsOk

`func (o *GetMobileDevicePrestage) GetSkipSetupItemsOk() (*map[string]bool, bool)`

GetSkipSetupItemsOk returns a tuple with the SkipSetupItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSetupItems

`func (o *GetMobileDevicePrestage) SetSkipSetupItems(v map[string]bool)`

SetSkipSetupItems sets SkipSetupItems field to given value.

### HasSkipSetupItems

`func (o *GetMobileDevicePrestage) HasSkipSetupItems() bool`

HasSkipSetupItems returns a boolean if a field has been set.

### GetLocationInformation

`func (o *GetMobileDevicePrestage) GetLocationInformation() LocationInformation`

GetLocationInformation returns the LocationInformation field if non-nil, zero value otherwise.

### GetLocationInformationOk

`func (o *GetMobileDevicePrestage) GetLocationInformationOk() (*LocationInformation, bool)`

GetLocationInformationOk returns a tuple with the LocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationInformation

`func (o *GetMobileDevicePrestage) SetLocationInformation(v LocationInformation)`

SetLocationInformation sets LocationInformation field to given value.


### GetPurchasingInformation

`func (o *GetMobileDevicePrestage) GetPurchasingInformation() PrestagePurchasingInformation`

GetPurchasingInformation returns the PurchasingInformation field if non-nil, zero value otherwise.

### GetPurchasingInformationOk

`func (o *GetMobileDevicePrestage) GetPurchasingInformationOk() (*PrestagePurchasingInformation, bool)`

GetPurchasingInformationOk returns a tuple with the PurchasingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingInformation

`func (o *GetMobileDevicePrestage) SetPurchasingInformation(v PrestagePurchasingInformation)`

SetPurchasingInformation sets PurchasingInformation field to given value.


### GetAnchorCertificates

`func (o *GetMobileDevicePrestage) GetAnchorCertificates() []string`

GetAnchorCertificates returns the AnchorCertificates field if non-nil, zero value otherwise.

### GetAnchorCertificatesOk

`func (o *GetMobileDevicePrestage) GetAnchorCertificatesOk() (*[]string, bool)`

GetAnchorCertificatesOk returns a tuple with the AnchorCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorCertificates

`func (o *GetMobileDevicePrestage) SetAnchorCertificates(v []string)`

SetAnchorCertificates sets AnchorCertificates field to given value.

### HasAnchorCertificates

`func (o *GetMobileDevicePrestage) HasAnchorCertificates() bool`

HasAnchorCertificates returns a boolean if a field has been set.

### GetEnrollmentCustomizationId

`func (o *GetMobileDevicePrestage) GetEnrollmentCustomizationId() int64`

GetEnrollmentCustomizationId returns the EnrollmentCustomizationId field if non-nil, zero value otherwise.

### GetEnrollmentCustomizationIdOk

`func (o *GetMobileDevicePrestage) GetEnrollmentCustomizationIdOk() (*int64, bool)`

GetEnrollmentCustomizationIdOk returns a tuple with the EnrollmentCustomizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCustomizationId

`func (o *GetMobileDevicePrestage) SetEnrollmentCustomizationId(v int64)`

SetEnrollmentCustomizationId sets EnrollmentCustomizationId field to given value.

### HasEnrollmentCustomizationId

`func (o *GetMobileDevicePrestage) HasEnrollmentCustomizationId() bool`

HasEnrollmentCustomizationId returns a boolean if a field has been set.

### GetIsAllowPairing

`func (o *GetMobileDevicePrestage) GetIsAllowPairing() bool`

GetIsAllowPairing returns the IsAllowPairing field if non-nil, zero value otherwise.

### GetIsAllowPairingOk

`func (o *GetMobileDevicePrestage) GetIsAllowPairingOk() (*bool, bool)`

GetIsAllowPairingOk returns a tuple with the IsAllowPairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowPairing

`func (o *GetMobileDevicePrestage) SetIsAllowPairing(v bool)`

SetIsAllowPairing sets IsAllowPairing field to given value.


### GetIsMultiUser

`func (o *GetMobileDevicePrestage) GetIsMultiUser() bool`

GetIsMultiUser returns the IsMultiUser field if non-nil, zero value otherwise.

### GetIsMultiUserOk

`func (o *GetMobileDevicePrestage) GetIsMultiUserOk() (*bool, bool)`

GetIsMultiUserOk returns a tuple with the IsMultiUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiUser

`func (o *GetMobileDevicePrestage) SetIsMultiUser(v bool)`

SetIsMultiUser sets IsMultiUser field to given value.


### GetIsSupervised

`func (o *GetMobileDevicePrestage) GetIsSupervised() bool`

GetIsSupervised returns the IsSupervised field if non-nil, zero value otherwise.

### GetIsSupervisedOk

`func (o *GetMobileDevicePrestage) GetIsSupervisedOk() (*bool, bool)`

GetIsSupervisedOk returns a tuple with the IsSupervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupervised

`func (o *GetMobileDevicePrestage) SetIsSupervised(v bool)`

SetIsSupervised sets IsSupervised field to given value.


### GetMaximumSharedAccounts

`func (o *GetMobileDevicePrestage) GetMaximumSharedAccounts() int64`

GetMaximumSharedAccounts returns the MaximumSharedAccounts field if non-nil, zero value otherwise.

### GetMaximumSharedAccountsOk

`func (o *GetMobileDevicePrestage) GetMaximumSharedAccountsOk() (*int64, bool)`

GetMaximumSharedAccountsOk returns a tuple with the MaximumSharedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSharedAccounts

`func (o *GetMobileDevicePrestage) SetMaximumSharedAccounts(v int64)`

SetMaximumSharedAccounts sets MaximumSharedAccounts field to given value.


### GetIsAutoAdvanceSetup

`func (o *GetMobileDevicePrestage) GetIsAutoAdvanceSetup() bool`

GetIsAutoAdvanceSetup returns the IsAutoAdvanceSetup field if non-nil, zero value otherwise.

### GetIsAutoAdvanceSetupOk

`func (o *GetMobileDevicePrestage) GetIsAutoAdvanceSetupOk() (*bool, bool)`

GetIsAutoAdvanceSetupOk returns a tuple with the IsAutoAdvanceSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoAdvanceSetup

`func (o *GetMobileDevicePrestage) SetIsAutoAdvanceSetup(v bool)`

SetIsAutoAdvanceSetup sets IsAutoAdvanceSetup field to given value.


### GetIsConfigureDeviceBeforeSetupAssistant

`func (o *GetMobileDevicePrestage) GetIsConfigureDeviceBeforeSetupAssistant() bool`

GetIsConfigureDeviceBeforeSetupAssistant returns the IsConfigureDeviceBeforeSetupAssistant field if non-nil, zero value otherwise.

### GetIsConfigureDeviceBeforeSetupAssistantOk

`func (o *GetMobileDevicePrestage) GetIsConfigureDeviceBeforeSetupAssistantOk() (*bool, bool)`

GetIsConfigureDeviceBeforeSetupAssistantOk returns a tuple with the IsConfigureDeviceBeforeSetupAssistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigureDeviceBeforeSetupAssistant

`func (o *GetMobileDevicePrestage) SetIsConfigureDeviceBeforeSetupAssistant(v bool)`

SetIsConfigureDeviceBeforeSetupAssistant sets IsConfigureDeviceBeforeSetupAssistant field to given value.


### GetLanguage

`func (o *GetMobileDevicePrestage) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GetMobileDevicePrestage) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GetMobileDevicePrestage) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *GetMobileDevicePrestage) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetRegion

`func (o *GetMobileDevicePrestage) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetMobileDevicePrestage) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetMobileDevicePrestage) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetMobileDevicePrestage) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetNames

`func (o *GetMobileDevicePrestage) GetNames() MobileDevicePrestageNames`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *GetMobileDevicePrestage) GetNamesOk() (*MobileDevicePrestageNames, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *GetMobileDevicePrestage) SetNames(v MobileDevicePrestageNames)`

SetNames sets Names field to given value.

### HasNames

`func (o *GetMobileDevicePrestage) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetId

`func (o *GetMobileDevicePrestage) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMobileDevicePrestage) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMobileDevicePrestage) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetMobileDevicePrestage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProfileUUID

`func (o *GetMobileDevicePrestage) GetProfileUUID() string`

GetProfileUUID returns the ProfileUUID field if non-nil, zero value otherwise.

### GetProfileUUIDOk

`func (o *GetMobileDevicePrestage) GetProfileUUIDOk() (*string, bool)`

GetProfileUUIDOk returns a tuple with the ProfileUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUUID

`func (o *GetMobileDevicePrestage) SetProfileUUID(v string)`

SetProfileUUID sets ProfileUUID field to given value.

### HasProfileUUID

`func (o *GetMobileDevicePrestage) HasProfileUUID() bool`

HasProfileUUID returns a boolean if a field has been set.

### GetSiteId

`func (o *GetMobileDevicePrestage) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *GetMobileDevicePrestage) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *GetMobileDevicePrestage) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *GetMobileDevicePrestage) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetVersionLock

`func (o *GetMobileDevicePrestage) GetVersionLock() int64`

GetVersionLock returns the VersionLock field if non-nil, zero value otherwise.

### GetVersionLockOk

`func (o *GetMobileDevicePrestage) GetVersionLockOk() (*int64, bool)`

GetVersionLockOk returns a tuple with the VersionLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionLock

`func (o *GetMobileDevicePrestage) SetVersionLock(v int64)`

SetVersionLock sets VersionLock field to given value.

### HasVersionLock

`func (o *GetMobileDevicePrestage) HasVersionLock() bool`

HasVersionLock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


