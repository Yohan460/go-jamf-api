# EnrollmentSettingsV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallSingleProfile** | Pointer to **bool** |  | [optional] [default to false]
**SigningMdmProfileEnabled** | Pointer to **bool** |  | [optional] [default to false]
**MdmSigningCertificate** | Pointer to [**NullableCertificateIdentityV2**](CertificateIdentityV2.md) |  | [optional] 
**RestrictReenrollment** | Pointer to **bool** |  | [optional] [default to false]
**FlushLocationInformation** | Pointer to **bool** |  | [optional] [default to false]
**FlushLocationHistoryInformation** | Pointer to **bool** |  | [optional] [default to false]
**FlushPolicyHistory** | Pointer to **bool** |  | [optional] [default to false]
**FlushExtensionAttributes** | Pointer to **bool** |  | [optional] [default to false]
**FlushMdmCommandsOnReenroll** | Pointer to **string** |  | [optional] [default to "DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED"]
**MacOsEnterpriseEnrollmentEnabled** | Pointer to **bool** |  | [optional] [default to false]
**ManagementUsername** | **string** |  | [default to ""]
**ManagementPassword** | Pointer to **string** | managementPassword is no longer in use. Input value for managementPassword will be ignored. | [optional] [default to "null"]
**ManagementPasswordSet** | Pointer to **bool** | managementPasswordSet is no longer in use. Input value for managementPasswordSet will be ignored. | [optional] [readonly] 
**PasswordType** | Pointer to **string** | passwordType is no longer in use. Input value for passwordType will be ignored. | [optional] [default to "STATIC"]
**RandomPasswordLength** | Pointer to **int64** | randomPasswordLength is no longer in use. Input value for randomPasswordLength will be ignored. | [optional] [default to 8]
**CreateManagementAccount** | Pointer to **bool** |  | [optional] [default to true]
**HideManagementAccount** | Pointer to **bool** |  | [optional] [default to false]
**AllowSshOnlyManagementAccount** | Pointer to **bool** |  | [optional] [default to false]
**EnsureSshRunning** | Pointer to **bool** |  | [optional] [default to true]
**LaunchSelfService** | Pointer to **bool** |  | [optional] [default to false]
**SignQuickAdd** | Pointer to **bool** |  | [optional] [default to false]
**DeveloperCertificateIdentity** | Pointer to [**NullableCertificateIdentityV2**](CertificateIdentityV2.md) |  | [optional] 
**DeveloperCertificateIdentityDetails** | Pointer to [**CertificateDetails**](CertificateDetails.md) |  | [optional] 
**MdmSigningCertificateDetails** | Pointer to [**CertificateDetails**](CertificateDetails.md) |  | [optional] 
**IosEnterpriseEnrollmentEnabled** | Pointer to **bool** |  | [optional] [default to true]
**IosPersonalEnrollmentEnabled** | Pointer to **bool** |  | [optional] [default to false]
**PersonalDeviceEnrollmentType** | Pointer to **string** |  | [optional] [default to "PERSONALDEVICEPROFILES"]
**AccountDrivenUserEnrollmentEnabled** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewEnrollmentSettingsV2

`func NewEnrollmentSettingsV2(managementUsername string, ) *EnrollmentSettingsV2`

NewEnrollmentSettingsV2 instantiates a new EnrollmentSettingsV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentSettingsV2WithDefaults

`func NewEnrollmentSettingsV2WithDefaults() *EnrollmentSettingsV2`

NewEnrollmentSettingsV2WithDefaults instantiates a new EnrollmentSettingsV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallSingleProfile

`func (o *EnrollmentSettingsV2) GetInstallSingleProfile() bool`

GetInstallSingleProfile returns the InstallSingleProfile field if non-nil, zero value otherwise.

### GetInstallSingleProfileOk

`func (o *EnrollmentSettingsV2) GetInstallSingleProfileOk() (*bool, bool)`

GetInstallSingleProfileOk returns a tuple with the InstallSingleProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallSingleProfile

`func (o *EnrollmentSettingsV2) SetInstallSingleProfile(v bool)`

SetInstallSingleProfile sets InstallSingleProfile field to given value.

### HasInstallSingleProfile

`func (o *EnrollmentSettingsV2) HasInstallSingleProfile() bool`

HasInstallSingleProfile returns a boolean if a field has been set.

### GetSigningMdmProfileEnabled

`func (o *EnrollmentSettingsV2) GetSigningMdmProfileEnabled() bool`

GetSigningMdmProfileEnabled returns the SigningMdmProfileEnabled field if non-nil, zero value otherwise.

### GetSigningMdmProfileEnabledOk

`func (o *EnrollmentSettingsV2) GetSigningMdmProfileEnabledOk() (*bool, bool)`

GetSigningMdmProfileEnabledOk returns a tuple with the SigningMdmProfileEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningMdmProfileEnabled

`func (o *EnrollmentSettingsV2) SetSigningMdmProfileEnabled(v bool)`

SetSigningMdmProfileEnabled sets SigningMdmProfileEnabled field to given value.

### HasSigningMdmProfileEnabled

`func (o *EnrollmentSettingsV2) HasSigningMdmProfileEnabled() bool`

HasSigningMdmProfileEnabled returns a boolean if a field has been set.

### GetMdmSigningCertificate

`func (o *EnrollmentSettingsV2) GetMdmSigningCertificate() CertificateIdentityV2`

GetMdmSigningCertificate returns the MdmSigningCertificate field if non-nil, zero value otherwise.

### GetMdmSigningCertificateOk

`func (o *EnrollmentSettingsV2) GetMdmSigningCertificateOk() (*CertificateIdentityV2, bool)`

GetMdmSigningCertificateOk returns a tuple with the MdmSigningCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmSigningCertificate

`func (o *EnrollmentSettingsV2) SetMdmSigningCertificate(v CertificateIdentityV2)`

SetMdmSigningCertificate sets MdmSigningCertificate field to given value.

### HasMdmSigningCertificate

`func (o *EnrollmentSettingsV2) HasMdmSigningCertificate() bool`

HasMdmSigningCertificate returns a boolean if a field has been set.

### SetMdmSigningCertificateNil

`func (o *EnrollmentSettingsV2) SetMdmSigningCertificateNil(b bool)`

 SetMdmSigningCertificateNil sets the value for MdmSigningCertificate to be an explicit nil

### UnsetMdmSigningCertificate
`func (o *EnrollmentSettingsV2) UnsetMdmSigningCertificate()`

UnsetMdmSigningCertificate ensures that no value is present for MdmSigningCertificate, not even an explicit nil
### GetRestrictReenrollment

`func (o *EnrollmentSettingsV2) GetRestrictReenrollment() bool`

GetRestrictReenrollment returns the RestrictReenrollment field if non-nil, zero value otherwise.

### GetRestrictReenrollmentOk

`func (o *EnrollmentSettingsV2) GetRestrictReenrollmentOk() (*bool, bool)`

GetRestrictReenrollmentOk returns a tuple with the RestrictReenrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictReenrollment

`func (o *EnrollmentSettingsV2) SetRestrictReenrollment(v bool)`

SetRestrictReenrollment sets RestrictReenrollment field to given value.

### HasRestrictReenrollment

`func (o *EnrollmentSettingsV2) HasRestrictReenrollment() bool`

HasRestrictReenrollment returns a boolean if a field has been set.

### GetFlushLocationInformation

`func (o *EnrollmentSettingsV2) GetFlushLocationInformation() bool`

GetFlushLocationInformation returns the FlushLocationInformation field if non-nil, zero value otherwise.

### GetFlushLocationInformationOk

`func (o *EnrollmentSettingsV2) GetFlushLocationInformationOk() (*bool, bool)`

GetFlushLocationInformationOk returns a tuple with the FlushLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushLocationInformation

`func (o *EnrollmentSettingsV2) SetFlushLocationInformation(v bool)`

SetFlushLocationInformation sets FlushLocationInformation field to given value.

### HasFlushLocationInformation

`func (o *EnrollmentSettingsV2) HasFlushLocationInformation() bool`

HasFlushLocationInformation returns a boolean if a field has been set.

### GetFlushLocationHistoryInformation

`func (o *EnrollmentSettingsV2) GetFlushLocationHistoryInformation() bool`

GetFlushLocationHistoryInformation returns the FlushLocationHistoryInformation field if non-nil, zero value otherwise.

### GetFlushLocationHistoryInformationOk

`func (o *EnrollmentSettingsV2) GetFlushLocationHistoryInformationOk() (*bool, bool)`

GetFlushLocationHistoryInformationOk returns a tuple with the FlushLocationHistoryInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushLocationHistoryInformation

`func (o *EnrollmentSettingsV2) SetFlushLocationHistoryInformation(v bool)`

SetFlushLocationHistoryInformation sets FlushLocationHistoryInformation field to given value.

### HasFlushLocationHistoryInformation

`func (o *EnrollmentSettingsV2) HasFlushLocationHistoryInformation() bool`

HasFlushLocationHistoryInformation returns a boolean if a field has been set.

### GetFlushPolicyHistory

`func (o *EnrollmentSettingsV2) GetFlushPolicyHistory() bool`

GetFlushPolicyHistory returns the FlushPolicyHistory field if non-nil, zero value otherwise.

### GetFlushPolicyHistoryOk

`func (o *EnrollmentSettingsV2) GetFlushPolicyHistoryOk() (*bool, bool)`

GetFlushPolicyHistoryOk returns a tuple with the FlushPolicyHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushPolicyHistory

`func (o *EnrollmentSettingsV2) SetFlushPolicyHistory(v bool)`

SetFlushPolicyHistory sets FlushPolicyHistory field to given value.

### HasFlushPolicyHistory

`func (o *EnrollmentSettingsV2) HasFlushPolicyHistory() bool`

HasFlushPolicyHistory returns a boolean if a field has been set.

### GetFlushExtensionAttributes

`func (o *EnrollmentSettingsV2) GetFlushExtensionAttributes() bool`

GetFlushExtensionAttributes returns the FlushExtensionAttributes field if non-nil, zero value otherwise.

### GetFlushExtensionAttributesOk

`func (o *EnrollmentSettingsV2) GetFlushExtensionAttributesOk() (*bool, bool)`

GetFlushExtensionAttributesOk returns a tuple with the FlushExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushExtensionAttributes

`func (o *EnrollmentSettingsV2) SetFlushExtensionAttributes(v bool)`

SetFlushExtensionAttributes sets FlushExtensionAttributes field to given value.

### HasFlushExtensionAttributes

`func (o *EnrollmentSettingsV2) HasFlushExtensionAttributes() bool`

HasFlushExtensionAttributes returns a boolean if a field has been set.

### GetFlushMdmCommandsOnReenroll

`func (o *EnrollmentSettingsV2) GetFlushMdmCommandsOnReenroll() string`

GetFlushMdmCommandsOnReenroll returns the FlushMdmCommandsOnReenroll field if non-nil, zero value otherwise.

### GetFlushMdmCommandsOnReenrollOk

`func (o *EnrollmentSettingsV2) GetFlushMdmCommandsOnReenrollOk() (*string, bool)`

GetFlushMdmCommandsOnReenrollOk returns a tuple with the FlushMdmCommandsOnReenroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushMdmCommandsOnReenroll

`func (o *EnrollmentSettingsV2) SetFlushMdmCommandsOnReenroll(v string)`

SetFlushMdmCommandsOnReenroll sets FlushMdmCommandsOnReenroll field to given value.

### HasFlushMdmCommandsOnReenroll

`func (o *EnrollmentSettingsV2) HasFlushMdmCommandsOnReenroll() bool`

HasFlushMdmCommandsOnReenroll returns a boolean if a field has been set.

### GetMacOsEnterpriseEnrollmentEnabled

`func (o *EnrollmentSettingsV2) GetMacOsEnterpriseEnrollmentEnabled() bool`

GetMacOsEnterpriseEnrollmentEnabled returns the MacOsEnterpriseEnrollmentEnabled field if non-nil, zero value otherwise.

### GetMacOsEnterpriseEnrollmentEnabledOk

`func (o *EnrollmentSettingsV2) GetMacOsEnterpriseEnrollmentEnabledOk() (*bool, bool)`

GetMacOsEnterpriseEnrollmentEnabledOk returns a tuple with the MacOsEnterpriseEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacOsEnterpriseEnrollmentEnabled

`func (o *EnrollmentSettingsV2) SetMacOsEnterpriseEnrollmentEnabled(v bool)`

SetMacOsEnterpriseEnrollmentEnabled sets MacOsEnterpriseEnrollmentEnabled field to given value.

### HasMacOsEnterpriseEnrollmentEnabled

`func (o *EnrollmentSettingsV2) HasMacOsEnterpriseEnrollmentEnabled() bool`

HasMacOsEnterpriseEnrollmentEnabled returns a boolean if a field has been set.

### GetManagementUsername

`func (o *EnrollmentSettingsV2) GetManagementUsername() string`

GetManagementUsername returns the ManagementUsername field if non-nil, zero value otherwise.

### GetManagementUsernameOk

`func (o *EnrollmentSettingsV2) GetManagementUsernameOk() (*string, bool)`

GetManagementUsernameOk returns a tuple with the ManagementUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementUsername

`func (o *EnrollmentSettingsV2) SetManagementUsername(v string)`

SetManagementUsername sets ManagementUsername field to given value.


### GetManagementPassword

`func (o *EnrollmentSettingsV2) GetManagementPassword() string`

GetManagementPassword returns the ManagementPassword field if non-nil, zero value otherwise.

### GetManagementPasswordOk

`func (o *EnrollmentSettingsV2) GetManagementPasswordOk() (*string, bool)`

GetManagementPasswordOk returns a tuple with the ManagementPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPassword

`func (o *EnrollmentSettingsV2) SetManagementPassword(v string)`

SetManagementPassword sets ManagementPassword field to given value.

### HasManagementPassword

`func (o *EnrollmentSettingsV2) HasManagementPassword() bool`

HasManagementPassword returns a boolean if a field has been set.

### GetManagementPasswordSet

`func (o *EnrollmentSettingsV2) GetManagementPasswordSet() bool`

GetManagementPasswordSet returns the ManagementPasswordSet field if non-nil, zero value otherwise.

### GetManagementPasswordSetOk

`func (o *EnrollmentSettingsV2) GetManagementPasswordSetOk() (*bool, bool)`

GetManagementPasswordSetOk returns a tuple with the ManagementPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPasswordSet

`func (o *EnrollmentSettingsV2) SetManagementPasswordSet(v bool)`

SetManagementPasswordSet sets ManagementPasswordSet field to given value.

### HasManagementPasswordSet

`func (o *EnrollmentSettingsV2) HasManagementPasswordSet() bool`

HasManagementPasswordSet returns a boolean if a field has been set.

### GetPasswordType

`func (o *EnrollmentSettingsV2) GetPasswordType() string`

GetPasswordType returns the PasswordType field if non-nil, zero value otherwise.

### GetPasswordTypeOk

`func (o *EnrollmentSettingsV2) GetPasswordTypeOk() (*string, bool)`

GetPasswordTypeOk returns a tuple with the PasswordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordType

`func (o *EnrollmentSettingsV2) SetPasswordType(v string)`

SetPasswordType sets PasswordType field to given value.

### HasPasswordType

`func (o *EnrollmentSettingsV2) HasPasswordType() bool`

HasPasswordType returns a boolean if a field has been set.

### GetRandomPasswordLength

`func (o *EnrollmentSettingsV2) GetRandomPasswordLength() int64`

GetRandomPasswordLength returns the RandomPasswordLength field if non-nil, zero value otherwise.

### GetRandomPasswordLengthOk

`func (o *EnrollmentSettingsV2) GetRandomPasswordLengthOk() (*int64, bool)`

GetRandomPasswordLengthOk returns a tuple with the RandomPasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomPasswordLength

`func (o *EnrollmentSettingsV2) SetRandomPasswordLength(v int64)`

SetRandomPasswordLength sets RandomPasswordLength field to given value.

### HasRandomPasswordLength

`func (o *EnrollmentSettingsV2) HasRandomPasswordLength() bool`

HasRandomPasswordLength returns a boolean if a field has been set.

### GetCreateManagementAccount

`func (o *EnrollmentSettingsV2) GetCreateManagementAccount() bool`

GetCreateManagementAccount returns the CreateManagementAccount field if non-nil, zero value otherwise.

### GetCreateManagementAccountOk

`func (o *EnrollmentSettingsV2) GetCreateManagementAccountOk() (*bool, bool)`

GetCreateManagementAccountOk returns a tuple with the CreateManagementAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateManagementAccount

`func (o *EnrollmentSettingsV2) SetCreateManagementAccount(v bool)`

SetCreateManagementAccount sets CreateManagementAccount field to given value.

### HasCreateManagementAccount

`func (o *EnrollmentSettingsV2) HasCreateManagementAccount() bool`

HasCreateManagementAccount returns a boolean if a field has been set.

### GetHideManagementAccount

`func (o *EnrollmentSettingsV2) GetHideManagementAccount() bool`

GetHideManagementAccount returns the HideManagementAccount field if non-nil, zero value otherwise.

### GetHideManagementAccountOk

`func (o *EnrollmentSettingsV2) GetHideManagementAccountOk() (*bool, bool)`

GetHideManagementAccountOk returns a tuple with the HideManagementAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideManagementAccount

`func (o *EnrollmentSettingsV2) SetHideManagementAccount(v bool)`

SetHideManagementAccount sets HideManagementAccount field to given value.

### HasHideManagementAccount

`func (o *EnrollmentSettingsV2) HasHideManagementAccount() bool`

HasHideManagementAccount returns a boolean if a field has been set.

### GetAllowSshOnlyManagementAccount

`func (o *EnrollmentSettingsV2) GetAllowSshOnlyManagementAccount() bool`

GetAllowSshOnlyManagementAccount returns the AllowSshOnlyManagementAccount field if non-nil, zero value otherwise.

### GetAllowSshOnlyManagementAccountOk

`func (o *EnrollmentSettingsV2) GetAllowSshOnlyManagementAccountOk() (*bool, bool)`

GetAllowSshOnlyManagementAccountOk returns a tuple with the AllowSshOnlyManagementAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSshOnlyManagementAccount

`func (o *EnrollmentSettingsV2) SetAllowSshOnlyManagementAccount(v bool)`

SetAllowSshOnlyManagementAccount sets AllowSshOnlyManagementAccount field to given value.

### HasAllowSshOnlyManagementAccount

`func (o *EnrollmentSettingsV2) HasAllowSshOnlyManagementAccount() bool`

HasAllowSshOnlyManagementAccount returns a boolean if a field has been set.

### GetEnsureSshRunning

`func (o *EnrollmentSettingsV2) GetEnsureSshRunning() bool`

GetEnsureSshRunning returns the EnsureSshRunning field if non-nil, zero value otherwise.

### GetEnsureSshRunningOk

`func (o *EnrollmentSettingsV2) GetEnsureSshRunningOk() (*bool, bool)`

GetEnsureSshRunningOk returns a tuple with the EnsureSshRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsureSshRunning

`func (o *EnrollmentSettingsV2) SetEnsureSshRunning(v bool)`

SetEnsureSshRunning sets EnsureSshRunning field to given value.

### HasEnsureSshRunning

`func (o *EnrollmentSettingsV2) HasEnsureSshRunning() bool`

HasEnsureSshRunning returns a boolean if a field has been set.

### GetLaunchSelfService

`func (o *EnrollmentSettingsV2) GetLaunchSelfService() bool`

GetLaunchSelfService returns the LaunchSelfService field if non-nil, zero value otherwise.

### GetLaunchSelfServiceOk

`func (o *EnrollmentSettingsV2) GetLaunchSelfServiceOk() (*bool, bool)`

GetLaunchSelfServiceOk returns a tuple with the LaunchSelfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSelfService

`func (o *EnrollmentSettingsV2) SetLaunchSelfService(v bool)`

SetLaunchSelfService sets LaunchSelfService field to given value.

### HasLaunchSelfService

`func (o *EnrollmentSettingsV2) HasLaunchSelfService() bool`

HasLaunchSelfService returns a boolean if a field has been set.

### GetSignQuickAdd

`func (o *EnrollmentSettingsV2) GetSignQuickAdd() bool`

GetSignQuickAdd returns the SignQuickAdd field if non-nil, zero value otherwise.

### GetSignQuickAddOk

`func (o *EnrollmentSettingsV2) GetSignQuickAddOk() (*bool, bool)`

GetSignQuickAddOk returns a tuple with the SignQuickAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignQuickAdd

`func (o *EnrollmentSettingsV2) SetSignQuickAdd(v bool)`

SetSignQuickAdd sets SignQuickAdd field to given value.

### HasSignQuickAdd

`func (o *EnrollmentSettingsV2) HasSignQuickAdd() bool`

HasSignQuickAdd returns a boolean if a field has been set.

### GetDeveloperCertificateIdentity

`func (o *EnrollmentSettingsV2) GetDeveloperCertificateIdentity() CertificateIdentityV2`

GetDeveloperCertificateIdentity returns the DeveloperCertificateIdentity field if non-nil, zero value otherwise.

### GetDeveloperCertificateIdentityOk

`func (o *EnrollmentSettingsV2) GetDeveloperCertificateIdentityOk() (*CertificateIdentityV2, bool)`

GetDeveloperCertificateIdentityOk returns a tuple with the DeveloperCertificateIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperCertificateIdentity

`func (o *EnrollmentSettingsV2) SetDeveloperCertificateIdentity(v CertificateIdentityV2)`

SetDeveloperCertificateIdentity sets DeveloperCertificateIdentity field to given value.

### HasDeveloperCertificateIdentity

`func (o *EnrollmentSettingsV2) HasDeveloperCertificateIdentity() bool`

HasDeveloperCertificateIdentity returns a boolean if a field has been set.

### SetDeveloperCertificateIdentityNil

`func (o *EnrollmentSettingsV2) SetDeveloperCertificateIdentityNil(b bool)`

 SetDeveloperCertificateIdentityNil sets the value for DeveloperCertificateIdentity to be an explicit nil

### UnsetDeveloperCertificateIdentity
`func (o *EnrollmentSettingsV2) UnsetDeveloperCertificateIdentity()`

UnsetDeveloperCertificateIdentity ensures that no value is present for DeveloperCertificateIdentity, not even an explicit nil
### GetDeveloperCertificateIdentityDetails

`func (o *EnrollmentSettingsV2) GetDeveloperCertificateIdentityDetails() CertificateDetails`

GetDeveloperCertificateIdentityDetails returns the DeveloperCertificateIdentityDetails field if non-nil, zero value otherwise.

### GetDeveloperCertificateIdentityDetailsOk

`func (o *EnrollmentSettingsV2) GetDeveloperCertificateIdentityDetailsOk() (*CertificateDetails, bool)`

GetDeveloperCertificateIdentityDetailsOk returns a tuple with the DeveloperCertificateIdentityDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperCertificateIdentityDetails

`func (o *EnrollmentSettingsV2) SetDeveloperCertificateIdentityDetails(v CertificateDetails)`

SetDeveloperCertificateIdentityDetails sets DeveloperCertificateIdentityDetails field to given value.

### HasDeveloperCertificateIdentityDetails

`func (o *EnrollmentSettingsV2) HasDeveloperCertificateIdentityDetails() bool`

HasDeveloperCertificateIdentityDetails returns a boolean if a field has been set.

### GetMdmSigningCertificateDetails

`func (o *EnrollmentSettingsV2) GetMdmSigningCertificateDetails() CertificateDetails`

GetMdmSigningCertificateDetails returns the MdmSigningCertificateDetails field if non-nil, zero value otherwise.

### GetMdmSigningCertificateDetailsOk

`func (o *EnrollmentSettingsV2) GetMdmSigningCertificateDetailsOk() (*CertificateDetails, bool)`

GetMdmSigningCertificateDetailsOk returns a tuple with the MdmSigningCertificateDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmSigningCertificateDetails

`func (o *EnrollmentSettingsV2) SetMdmSigningCertificateDetails(v CertificateDetails)`

SetMdmSigningCertificateDetails sets MdmSigningCertificateDetails field to given value.

### HasMdmSigningCertificateDetails

`func (o *EnrollmentSettingsV2) HasMdmSigningCertificateDetails() bool`

HasMdmSigningCertificateDetails returns a boolean if a field has been set.

### GetIosEnterpriseEnrollmentEnabled

`func (o *EnrollmentSettingsV2) GetIosEnterpriseEnrollmentEnabled() bool`

GetIosEnterpriseEnrollmentEnabled returns the IosEnterpriseEnrollmentEnabled field if non-nil, zero value otherwise.

### GetIosEnterpriseEnrollmentEnabledOk

`func (o *EnrollmentSettingsV2) GetIosEnterpriseEnrollmentEnabledOk() (*bool, bool)`

GetIosEnterpriseEnrollmentEnabledOk returns a tuple with the IosEnterpriseEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosEnterpriseEnrollmentEnabled

`func (o *EnrollmentSettingsV2) SetIosEnterpriseEnrollmentEnabled(v bool)`

SetIosEnterpriseEnrollmentEnabled sets IosEnterpriseEnrollmentEnabled field to given value.

### HasIosEnterpriseEnrollmentEnabled

`func (o *EnrollmentSettingsV2) HasIosEnterpriseEnrollmentEnabled() bool`

HasIosEnterpriseEnrollmentEnabled returns a boolean if a field has been set.

### GetIosPersonalEnrollmentEnabled

`func (o *EnrollmentSettingsV2) GetIosPersonalEnrollmentEnabled() bool`

GetIosPersonalEnrollmentEnabled returns the IosPersonalEnrollmentEnabled field if non-nil, zero value otherwise.

### GetIosPersonalEnrollmentEnabledOk

`func (o *EnrollmentSettingsV2) GetIosPersonalEnrollmentEnabledOk() (*bool, bool)`

GetIosPersonalEnrollmentEnabledOk returns a tuple with the IosPersonalEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosPersonalEnrollmentEnabled

`func (o *EnrollmentSettingsV2) SetIosPersonalEnrollmentEnabled(v bool)`

SetIosPersonalEnrollmentEnabled sets IosPersonalEnrollmentEnabled field to given value.

### HasIosPersonalEnrollmentEnabled

`func (o *EnrollmentSettingsV2) HasIosPersonalEnrollmentEnabled() bool`

HasIosPersonalEnrollmentEnabled returns a boolean if a field has been set.

### GetPersonalDeviceEnrollmentType

`func (o *EnrollmentSettingsV2) GetPersonalDeviceEnrollmentType() string`

GetPersonalDeviceEnrollmentType returns the PersonalDeviceEnrollmentType field if non-nil, zero value otherwise.

### GetPersonalDeviceEnrollmentTypeOk

`func (o *EnrollmentSettingsV2) GetPersonalDeviceEnrollmentTypeOk() (*string, bool)`

GetPersonalDeviceEnrollmentTypeOk returns a tuple with the PersonalDeviceEnrollmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalDeviceEnrollmentType

`func (o *EnrollmentSettingsV2) SetPersonalDeviceEnrollmentType(v string)`

SetPersonalDeviceEnrollmentType sets PersonalDeviceEnrollmentType field to given value.

### HasPersonalDeviceEnrollmentType

`func (o *EnrollmentSettingsV2) HasPersonalDeviceEnrollmentType() bool`

HasPersonalDeviceEnrollmentType returns a boolean if a field has been set.

### GetAccountDrivenUserEnrollmentEnabled

`func (o *EnrollmentSettingsV2) GetAccountDrivenUserEnrollmentEnabled() bool`

GetAccountDrivenUserEnrollmentEnabled returns the AccountDrivenUserEnrollmentEnabled field if non-nil, zero value otherwise.

### GetAccountDrivenUserEnrollmentEnabledOk

`func (o *EnrollmentSettingsV2) GetAccountDrivenUserEnrollmentEnabledOk() (*bool, bool)`

GetAccountDrivenUserEnrollmentEnabledOk returns a tuple with the AccountDrivenUserEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDrivenUserEnrollmentEnabled

`func (o *EnrollmentSettingsV2) SetAccountDrivenUserEnrollmentEnabled(v bool)`

SetAccountDrivenUserEnrollmentEnabled sets AccountDrivenUserEnrollmentEnabled field to given value.

### HasAccountDrivenUserEnrollmentEnabled

`func (o *EnrollmentSettingsV2) HasAccountDrivenUserEnrollmentEnabled() bool`

HasAccountDrivenUserEnrollmentEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


