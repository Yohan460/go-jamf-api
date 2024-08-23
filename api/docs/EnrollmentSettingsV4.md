# EnrollmentSettingsV4

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
**AccountDrivenDeviceIosEnrollmentEnabled** | Pointer to **bool** |  | [optional] [default to false]
**AccountDrivenDeviceMacosEnrollmentEnabled** | Pointer to **bool** |  | [optional] [default to false]
**AccountDrivenUserVisionosEnrollmentEnabled** | Pointer to **bool** |  | [optional] [default to false]
**AccountDrivenDeviceVisionosEnrollmentEnabled** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewEnrollmentSettingsV4

`func NewEnrollmentSettingsV4(managementUsername string, ) *EnrollmentSettingsV4`

NewEnrollmentSettingsV4 instantiates a new EnrollmentSettingsV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentSettingsV4WithDefaults

`func NewEnrollmentSettingsV4WithDefaults() *EnrollmentSettingsV4`

NewEnrollmentSettingsV4WithDefaults instantiates a new EnrollmentSettingsV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallSingleProfile

`func (o *EnrollmentSettingsV4) GetInstallSingleProfile() bool`

GetInstallSingleProfile returns the InstallSingleProfile field if non-nil, zero value otherwise.

### GetInstallSingleProfileOk

`func (o *EnrollmentSettingsV4) GetInstallSingleProfileOk() (*bool, bool)`

GetInstallSingleProfileOk returns a tuple with the InstallSingleProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallSingleProfile

`func (o *EnrollmentSettingsV4) SetInstallSingleProfile(v bool)`

SetInstallSingleProfile sets InstallSingleProfile field to given value.

### HasInstallSingleProfile

`func (o *EnrollmentSettingsV4) HasInstallSingleProfile() bool`

HasInstallSingleProfile returns a boolean if a field has been set.

### GetSigningMdmProfileEnabled

`func (o *EnrollmentSettingsV4) GetSigningMdmProfileEnabled() bool`

GetSigningMdmProfileEnabled returns the SigningMdmProfileEnabled field if non-nil, zero value otherwise.

### GetSigningMdmProfileEnabledOk

`func (o *EnrollmentSettingsV4) GetSigningMdmProfileEnabledOk() (*bool, bool)`

GetSigningMdmProfileEnabledOk returns a tuple with the SigningMdmProfileEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningMdmProfileEnabled

`func (o *EnrollmentSettingsV4) SetSigningMdmProfileEnabled(v bool)`

SetSigningMdmProfileEnabled sets SigningMdmProfileEnabled field to given value.

### HasSigningMdmProfileEnabled

`func (o *EnrollmentSettingsV4) HasSigningMdmProfileEnabled() bool`

HasSigningMdmProfileEnabled returns a boolean if a field has been set.

### GetMdmSigningCertificate

`func (o *EnrollmentSettingsV4) GetMdmSigningCertificate() CertificateIdentityV2`

GetMdmSigningCertificate returns the MdmSigningCertificate field if non-nil, zero value otherwise.

### GetMdmSigningCertificateOk

`func (o *EnrollmentSettingsV4) GetMdmSigningCertificateOk() (*CertificateIdentityV2, bool)`

GetMdmSigningCertificateOk returns a tuple with the MdmSigningCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmSigningCertificate

`func (o *EnrollmentSettingsV4) SetMdmSigningCertificate(v CertificateIdentityV2)`

SetMdmSigningCertificate sets MdmSigningCertificate field to given value.

### HasMdmSigningCertificate

`func (o *EnrollmentSettingsV4) HasMdmSigningCertificate() bool`

HasMdmSigningCertificate returns a boolean if a field has been set.

### SetMdmSigningCertificateNil

`func (o *EnrollmentSettingsV4) SetMdmSigningCertificateNil(b bool)`

 SetMdmSigningCertificateNil sets the value for MdmSigningCertificate to be an explicit nil

### UnsetMdmSigningCertificate
`func (o *EnrollmentSettingsV4) UnsetMdmSigningCertificate()`

UnsetMdmSigningCertificate ensures that no value is present for MdmSigningCertificate, not even an explicit nil
### GetRestrictReenrollment

`func (o *EnrollmentSettingsV4) GetRestrictReenrollment() bool`

GetRestrictReenrollment returns the RestrictReenrollment field if non-nil, zero value otherwise.

### GetRestrictReenrollmentOk

`func (o *EnrollmentSettingsV4) GetRestrictReenrollmentOk() (*bool, bool)`

GetRestrictReenrollmentOk returns a tuple with the RestrictReenrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictReenrollment

`func (o *EnrollmentSettingsV4) SetRestrictReenrollment(v bool)`

SetRestrictReenrollment sets RestrictReenrollment field to given value.

### HasRestrictReenrollment

`func (o *EnrollmentSettingsV4) HasRestrictReenrollment() bool`

HasRestrictReenrollment returns a boolean if a field has been set.

### GetFlushLocationInformation

`func (o *EnrollmentSettingsV4) GetFlushLocationInformation() bool`

GetFlushLocationInformation returns the FlushLocationInformation field if non-nil, zero value otherwise.

### GetFlushLocationInformationOk

`func (o *EnrollmentSettingsV4) GetFlushLocationInformationOk() (*bool, bool)`

GetFlushLocationInformationOk returns a tuple with the FlushLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushLocationInformation

`func (o *EnrollmentSettingsV4) SetFlushLocationInformation(v bool)`

SetFlushLocationInformation sets FlushLocationInformation field to given value.

### HasFlushLocationInformation

`func (o *EnrollmentSettingsV4) HasFlushLocationInformation() bool`

HasFlushLocationInformation returns a boolean if a field has been set.

### GetFlushLocationHistoryInformation

`func (o *EnrollmentSettingsV4) GetFlushLocationHistoryInformation() bool`

GetFlushLocationHistoryInformation returns the FlushLocationHistoryInformation field if non-nil, zero value otherwise.

### GetFlushLocationHistoryInformationOk

`func (o *EnrollmentSettingsV4) GetFlushLocationHistoryInformationOk() (*bool, bool)`

GetFlushLocationHistoryInformationOk returns a tuple with the FlushLocationHistoryInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushLocationHistoryInformation

`func (o *EnrollmentSettingsV4) SetFlushLocationHistoryInformation(v bool)`

SetFlushLocationHistoryInformation sets FlushLocationHistoryInformation field to given value.

### HasFlushLocationHistoryInformation

`func (o *EnrollmentSettingsV4) HasFlushLocationHistoryInformation() bool`

HasFlushLocationHistoryInformation returns a boolean if a field has been set.

### GetFlushPolicyHistory

`func (o *EnrollmentSettingsV4) GetFlushPolicyHistory() bool`

GetFlushPolicyHistory returns the FlushPolicyHistory field if non-nil, zero value otherwise.

### GetFlushPolicyHistoryOk

`func (o *EnrollmentSettingsV4) GetFlushPolicyHistoryOk() (*bool, bool)`

GetFlushPolicyHistoryOk returns a tuple with the FlushPolicyHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushPolicyHistory

`func (o *EnrollmentSettingsV4) SetFlushPolicyHistory(v bool)`

SetFlushPolicyHistory sets FlushPolicyHistory field to given value.

### HasFlushPolicyHistory

`func (o *EnrollmentSettingsV4) HasFlushPolicyHistory() bool`

HasFlushPolicyHistory returns a boolean if a field has been set.

### GetFlushExtensionAttributes

`func (o *EnrollmentSettingsV4) GetFlushExtensionAttributes() bool`

GetFlushExtensionAttributes returns the FlushExtensionAttributes field if non-nil, zero value otherwise.

### GetFlushExtensionAttributesOk

`func (o *EnrollmentSettingsV4) GetFlushExtensionAttributesOk() (*bool, bool)`

GetFlushExtensionAttributesOk returns a tuple with the FlushExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushExtensionAttributes

`func (o *EnrollmentSettingsV4) SetFlushExtensionAttributes(v bool)`

SetFlushExtensionAttributes sets FlushExtensionAttributes field to given value.

### HasFlushExtensionAttributes

`func (o *EnrollmentSettingsV4) HasFlushExtensionAttributes() bool`

HasFlushExtensionAttributes returns a boolean if a field has been set.

### GetFlushMdmCommandsOnReenroll

`func (o *EnrollmentSettingsV4) GetFlushMdmCommandsOnReenroll() string`

GetFlushMdmCommandsOnReenroll returns the FlushMdmCommandsOnReenroll field if non-nil, zero value otherwise.

### GetFlushMdmCommandsOnReenrollOk

`func (o *EnrollmentSettingsV4) GetFlushMdmCommandsOnReenrollOk() (*string, bool)`

GetFlushMdmCommandsOnReenrollOk returns a tuple with the FlushMdmCommandsOnReenroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushMdmCommandsOnReenroll

`func (o *EnrollmentSettingsV4) SetFlushMdmCommandsOnReenroll(v string)`

SetFlushMdmCommandsOnReenroll sets FlushMdmCommandsOnReenroll field to given value.

### HasFlushMdmCommandsOnReenroll

`func (o *EnrollmentSettingsV4) HasFlushMdmCommandsOnReenroll() bool`

HasFlushMdmCommandsOnReenroll returns a boolean if a field has been set.

### GetMacOsEnterpriseEnrollmentEnabled

`func (o *EnrollmentSettingsV4) GetMacOsEnterpriseEnrollmentEnabled() bool`

GetMacOsEnterpriseEnrollmentEnabled returns the MacOsEnterpriseEnrollmentEnabled field if non-nil, zero value otherwise.

### GetMacOsEnterpriseEnrollmentEnabledOk

`func (o *EnrollmentSettingsV4) GetMacOsEnterpriseEnrollmentEnabledOk() (*bool, bool)`

GetMacOsEnterpriseEnrollmentEnabledOk returns a tuple with the MacOsEnterpriseEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacOsEnterpriseEnrollmentEnabled

`func (o *EnrollmentSettingsV4) SetMacOsEnterpriseEnrollmentEnabled(v bool)`

SetMacOsEnterpriseEnrollmentEnabled sets MacOsEnterpriseEnrollmentEnabled field to given value.

### HasMacOsEnterpriseEnrollmentEnabled

`func (o *EnrollmentSettingsV4) HasMacOsEnterpriseEnrollmentEnabled() bool`

HasMacOsEnterpriseEnrollmentEnabled returns a boolean if a field has been set.

### GetManagementUsername

`func (o *EnrollmentSettingsV4) GetManagementUsername() string`

GetManagementUsername returns the ManagementUsername field if non-nil, zero value otherwise.

### GetManagementUsernameOk

`func (o *EnrollmentSettingsV4) GetManagementUsernameOk() (*string, bool)`

GetManagementUsernameOk returns a tuple with the ManagementUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementUsername

`func (o *EnrollmentSettingsV4) SetManagementUsername(v string)`

SetManagementUsername sets ManagementUsername field to given value.


### GetCreateManagementAccount

`func (o *EnrollmentSettingsV4) GetCreateManagementAccount() bool`

GetCreateManagementAccount returns the CreateManagementAccount field if non-nil, zero value otherwise.

### GetCreateManagementAccountOk

`func (o *EnrollmentSettingsV4) GetCreateManagementAccountOk() (*bool, bool)`

GetCreateManagementAccountOk returns a tuple with the CreateManagementAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateManagementAccount

`func (o *EnrollmentSettingsV4) SetCreateManagementAccount(v bool)`

SetCreateManagementAccount sets CreateManagementAccount field to given value.

### HasCreateManagementAccount

`func (o *EnrollmentSettingsV4) HasCreateManagementAccount() bool`

HasCreateManagementAccount returns a boolean if a field has been set.

### GetHideManagementAccount

`func (o *EnrollmentSettingsV4) GetHideManagementAccount() bool`

GetHideManagementAccount returns the HideManagementAccount field if non-nil, zero value otherwise.

### GetHideManagementAccountOk

`func (o *EnrollmentSettingsV4) GetHideManagementAccountOk() (*bool, bool)`

GetHideManagementAccountOk returns a tuple with the HideManagementAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideManagementAccount

`func (o *EnrollmentSettingsV4) SetHideManagementAccount(v bool)`

SetHideManagementAccount sets HideManagementAccount field to given value.

### HasHideManagementAccount

`func (o *EnrollmentSettingsV4) HasHideManagementAccount() bool`

HasHideManagementAccount returns a boolean if a field has been set.

### GetAllowSshOnlyManagementAccount

`func (o *EnrollmentSettingsV4) GetAllowSshOnlyManagementAccount() bool`

GetAllowSshOnlyManagementAccount returns the AllowSshOnlyManagementAccount field if non-nil, zero value otherwise.

### GetAllowSshOnlyManagementAccountOk

`func (o *EnrollmentSettingsV4) GetAllowSshOnlyManagementAccountOk() (*bool, bool)`

GetAllowSshOnlyManagementAccountOk returns a tuple with the AllowSshOnlyManagementAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSshOnlyManagementAccount

`func (o *EnrollmentSettingsV4) SetAllowSshOnlyManagementAccount(v bool)`

SetAllowSshOnlyManagementAccount sets AllowSshOnlyManagementAccount field to given value.

### HasAllowSshOnlyManagementAccount

`func (o *EnrollmentSettingsV4) HasAllowSshOnlyManagementAccount() bool`

HasAllowSshOnlyManagementAccount returns a boolean if a field has been set.

### GetEnsureSshRunning

`func (o *EnrollmentSettingsV4) GetEnsureSshRunning() bool`

GetEnsureSshRunning returns the EnsureSshRunning field if non-nil, zero value otherwise.

### GetEnsureSshRunningOk

`func (o *EnrollmentSettingsV4) GetEnsureSshRunningOk() (*bool, bool)`

GetEnsureSshRunningOk returns a tuple with the EnsureSshRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsureSshRunning

`func (o *EnrollmentSettingsV4) SetEnsureSshRunning(v bool)`

SetEnsureSshRunning sets EnsureSshRunning field to given value.

### HasEnsureSshRunning

`func (o *EnrollmentSettingsV4) HasEnsureSshRunning() bool`

HasEnsureSshRunning returns a boolean if a field has been set.

### GetLaunchSelfService

`func (o *EnrollmentSettingsV4) GetLaunchSelfService() bool`

GetLaunchSelfService returns the LaunchSelfService field if non-nil, zero value otherwise.

### GetLaunchSelfServiceOk

`func (o *EnrollmentSettingsV4) GetLaunchSelfServiceOk() (*bool, bool)`

GetLaunchSelfServiceOk returns a tuple with the LaunchSelfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSelfService

`func (o *EnrollmentSettingsV4) SetLaunchSelfService(v bool)`

SetLaunchSelfService sets LaunchSelfService field to given value.

### HasLaunchSelfService

`func (o *EnrollmentSettingsV4) HasLaunchSelfService() bool`

HasLaunchSelfService returns a boolean if a field has been set.

### GetSignQuickAdd

`func (o *EnrollmentSettingsV4) GetSignQuickAdd() bool`

GetSignQuickAdd returns the SignQuickAdd field if non-nil, zero value otherwise.

### GetSignQuickAddOk

`func (o *EnrollmentSettingsV4) GetSignQuickAddOk() (*bool, bool)`

GetSignQuickAddOk returns a tuple with the SignQuickAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignQuickAdd

`func (o *EnrollmentSettingsV4) SetSignQuickAdd(v bool)`

SetSignQuickAdd sets SignQuickAdd field to given value.

### HasSignQuickAdd

`func (o *EnrollmentSettingsV4) HasSignQuickAdd() bool`

HasSignQuickAdd returns a boolean if a field has been set.

### GetDeveloperCertificateIdentity

`func (o *EnrollmentSettingsV4) GetDeveloperCertificateIdentity() CertificateIdentityV2`

GetDeveloperCertificateIdentity returns the DeveloperCertificateIdentity field if non-nil, zero value otherwise.

### GetDeveloperCertificateIdentityOk

`func (o *EnrollmentSettingsV4) GetDeveloperCertificateIdentityOk() (*CertificateIdentityV2, bool)`

GetDeveloperCertificateIdentityOk returns a tuple with the DeveloperCertificateIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperCertificateIdentity

`func (o *EnrollmentSettingsV4) SetDeveloperCertificateIdentity(v CertificateIdentityV2)`

SetDeveloperCertificateIdentity sets DeveloperCertificateIdentity field to given value.

### HasDeveloperCertificateIdentity

`func (o *EnrollmentSettingsV4) HasDeveloperCertificateIdentity() bool`

HasDeveloperCertificateIdentity returns a boolean if a field has been set.

### SetDeveloperCertificateIdentityNil

`func (o *EnrollmentSettingsV4) SetDeveloperCertificateIdentityNil(b bool)`

 SetDeveloperCertificateIdentityNil sets the value for DeveloperCertificateIdentity to be an explicit nil

### UnsetDeveloperCertificateIdentity
`func (o *EnrollmentSettingsV4) UnsetDeveloperCertificateIdentity()`

UnsetDeveloperCertificateIdentity ensures that no value is present for DeveloperCertificateIdentity, not even an explicit nil
### GetDeveloperCertificateIdentityDetails

`func (o *EnrollmentSettingsV4) GetDeveloperCertificateIdentityDetails() CertificateDetails`

GetDeveloperCertificateIdentityDetails returns the DeveloperCertificateIdentityDetails field if non-nil, zero value otherwise.

### GetDeveloperCertificateIdentityDetailsOk

`func (o *EnrollmentSettingsV4) GetDeveloperCertificateIdentityDetailsOk() (*CertificateDetails, bool)`

GetDeveloperCertificateIdentityDetailsOk returns a tuple with the DeveloperCertificateIdentityDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperCertificateIdentityDetails

`func (o *EnrollmentSettingsV4) SetDeveloperCertificateIdentityDetails(v CertificateDetails)`

SetDeveloperCertificateIdentityDetails sets DeveloperCertificateIdentityDetails field to given value.

### HasDeveloperCertificateIdentityDetails

`func (o *EnrollmentSettingsV4) HasDeveloperCertificateIdentityDetails() bool`

HasDeveloperCertificateIdentityDetails returns a boolean if a field has been set.

### GetMdmSigningCertificateDetails

`func (o *EnrollmentSettingsV4) GetMdmSigningCertificateDetails() CertificateDetails`

GetMdmSigningCertificateDetails returns the MdmSigningCertificateDetails field if non-nil, zero value otherwise.

### GetMdmSigningCertificateDetailsOk

`func (o *EnrollmentSettingsV4) GetMdmSigningCertificateDetailsOk() (*CertificateDetails, bool)`

GetMdmSigningCertificateDetailsOk returns a tuple with the MdmSigningCertificateDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmSigningCertificateDetails

`func (o *EnrollmentSettingsV4) SetMdmSigningCertificateDetails(v CertificateDetails)`

SetMdmSigningCertificateDetails sets MdmSigningCertificateDetails field to given value.

### HasMdmSigningCertificateDetails

`func (o *EnrollmentSettingsV4) HasMdmSigningCertificateDetails() bool`

HasMdmSigningCertificateDetails returns a boolean if a field has been set.

### GetIosEnterpriseEnrollmentEnabled

`func (o *EnrollmentSettingsV4) GetIosEnterpriseEnrollmentEnabled() bool`

GetIosEnterpriseEnrollmentEnabled returns the IosEnterpriseEnrollmentEnabled field if non-nil, zero value otherwise.

### GetIosEnterpriseEnrollmentEnabledOk

`func (o *EnrollmentSettingsV4) GetIosEnterpriseEnrollmentEnabledOk() (*bool, bool)`

GetIosEnterpriseEnrollmentEnabledOk returns a tuple with the IosEnterpriseEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosEnterpriseEnrollmentEnabled

`func (o *EnrollmentSettingsV4) SetIosEnterpriseEnrollmentEnabled(v bool)`

SetIosEnterpriseEnrollmentEnabled sets IosEnterpriseEnrollmentEnabled field to given value.

### HasIosEnterpriseEnrollmentEnabled

`func (o *EnrollmentSettingsV4) HasIosEnterpriseEnrollmentEnabled() bool`

HasIosEnterpriseEnrollmentEnabled returns a boolean if a field has been set.

### GetIosPersonalEnrollmentEnabled

`func (o *EnrollmentSettingsV4) GetIosPersonalEnrollmentEnabled() bool`

GetIosPersonalEnrollmentEnabled returns the IosPersonalEnrollmentEnabled field if non-nil, zero value otherwise.

### GetIosPersonalEnrollmentEnabledOk

`func (o *EnrollmentSettingsV4) GetIosPersonalEnrollmentEnabledOk() (*bool, bool)`

GetIosPersonalEnrollmentEnabledOk returns a tuple with the IosPersonalEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosPersonalEnrollmentEnabled

`func (o *EnrollmentSettingsV4) SetIosPersonalEnrollmentEnabled(v bool)`

SetIosPersonalEnrollmentEnabled sets IosPersonalEnrollmentEnabled field to given value.

### HasIosPersonalEnrollmentEnabled

`func (o *EnrollmentSettingsV4) HasIosPersonalEnrollmentEnabled() bool`

HasIosPersonalEnrollmentEnabled returns a boolean if a field has been set.

### GetPersonalDeviceEnrollmentType

`func (o *EnrollmentSettingsV4) GetPersonalDeviceEnrollmentType() string`

GetPersonalDeviceEnrollmentType returns the PersonalDeviceEnrollmentType field if non-nil, zero value otherwise.

### GetPersonalDeviceEnrollmentTypeOk

`func (o *EnrollmentSettingsV4) GetPersonalDeviceEnrollmentTypeOk() (*string, bool)`

GetPersonalDeviceEnrollmentTypeOk returns a tuple with the PersonalDeviceEnrollmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalDeviceEnrollmentType

`func (o *EnrollmentSettingsV4) SetPersonalDeviceEnrollmentType(v string)`

SetPersonalDeviceEnrollmentType sets PersonalDeviceEnrollmentType field to given value.

### HasPersonalDeviceEnrollmentType

`func (o *EnrollmentSettingsV4) HasPersonalDeviceEnrollmentType() bool`

HasPersonalDeviceEnrollmentType returns a boolean if a field has been set.

### GetAccountDrivenUserEnrollmentEnabled

`func (o *EnrollmentSettingsV4) GetAccountDrivenUserEnrollmentEnabled() bool`

GetAccountDrivenUserEnrollmentEnabled returns the AccountDrivenUserEnrollmentEnabled field if non-nil, zero value otherwise.

### GetAccountDrivenUserEnrollmentEnabledOk

`func (o *EnrollmentSettingsV4) GetAccountDrivenUserEnrollmentEnabledOk() (*bool, bool)`

GetAccountDrivenUserEnrollmentEnabledOk returns a tuple with the AccountDrivenUserEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDrivenUserEnrollmentEnabled

`func (o *EnrollmentSettingsV4) SetAccountDrivenUserEnrollmentEnabled(v bool)`

SetAccountDrivenUserEnrollmentEnabled sets AccountDrivenUserEnrollmentEnabled field to given value.

### HasAccountDrivenUserEnrollmentEnabled

`func (o *EnrollmentSettingsV4) HasAccountDrivenUserEnrollmentEnabled() bool`

HasAccountDrivenUserEnrollmentEnabled returns a boolean if a field has been set.

### GetAccountDrivenDeviceIosEnrollmentEnabled

`func (o *EnrollmentSettingsV4) GetAccountDrivenDeviceIosEnrollmentEnabled() bool`

GetAccountDrivenDeviceIosEnrollmentEnabled returns the AccountDrivenDeviceIosEnrollmentEnabled field if non-nil, zero value otherwise.

### GetAccountDrivenDeviceIosEnrollmentEnabledOk

`func (o *EnrollmentSettingsV4) GetAccountDrivenDeviceIosEnrollmentEnabledOk() (*bool, bool)`

GetAccountDrivenDeviceIosEnrollmentEnabledOk returns a tuple with the AccountDrivenDeviceIosEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDrivenDeviceIosEnrollmentEnabled

`func (o *EnrollmentSettingsV4) SetAccountDrivenDeviceIosEnrollmentEnabled(v bool)`

SetAccountDrivenDeviceIosEnrollmentEnabled sets AccountDrivenDeviceIosEnrollmentEnabled field to given value.

### HasAccountDrivenDeviceIosEnrollmentEnabled

`func (o *EnrollmentSettingsV4) HasAccountDrivenDeviceIosEnrollmentEnabled() bool`

HasAccountDrivenDeviceIosEnrollmentEnabled returns a boolean if a field has been set.

### GetAccountDrivenDeviceMacosEnrollmentEnabled

`func (o *EnrollmentSettingsV4) GetAccountDrivenDeviceMacosEnrollmentEnabled() bool`

GetAccountDrivenDeviceMacosEnrollmentEnabled returns the AccountDrivenDeviceMacosEnrollmentEnabled field if non-nil, zero value otherwise.

### GetAccountDrivenDeviceMacosEnrollmentEnabledOk

`func (o *EnrollmentSettingsV4) GetAccountDrivenDeviceMacosEnrollmentEnabledOk() (*bool, bool)`

GetAccountDrivenDeviceMacosEnrollmentEnabledOk returns a tuple with the AccountDrivenDeviceMacosEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDrivenDeviceMacosEnrollmentEnabled

`func (o *EnrollmentSettingsV4) SetAccountDrivenDeviceMacosEnrollmentEnabled(v bool)`

SetAccountDrivenDeviceMacosEnrollmentEnabled sets AccountDrivenDeviceMacosEnrollmentEnabled field to given value.

### HasAccountDrivenDeviceMacosEnrollmentEnabled

`func (o *EnrollmentSettingsV4) HasAccountDrivenDeviceMacosEnrollmentEnabled() bool`

HasAccountDrivenDeviceMacosEnrollmentEnabled returns a boolean if a field has been set.

### GetAccountDrivenUserVisionosEnrollmentEnabled

`func (o *EnrollmentSettingsV4) GetAccountDrivenUserVisionosEnrollmentEnabled() bool`

GetAccountDrivenUserVisionosEnrollmentEnabled returns the AccountDrivenUserVisionosEnrollmentEnabled field if non-nil, zero value otherwise.

### GetAccountDrivenUserVisionosEnrollmentEnabledOk

`func (o *EnrollmentSettingsV4) GetAccountDrivenUserVisionosEnrollmentEnabledOk() (*bool, bool)`

GetAccountDrivenUserVisionosEnrollmentEnabledOk returns a tuple with the AccountDrivenUserVisionosEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDrivenUserVisionosEnrollmentEnabled

`func (o *EnrollmentSettingsV4) SetAccountDrivenUserVisionosEnrollmentEnabled(v bool)`

SetAccountDrivenUserVisionosEnrollmentEnabled sets AccountDrivenUserVisionosEnrollmentEnabled field to given value.

### HasAccountDrivenUserVisionosEnrollmentEnabled

`func (o *EnrollmentSettingsV4) HasAccountDrivenUserVisionosEnrollmentEnabled() bool`

HasAccountDrivenUserVisionosEnrollmentEnabled returns a boolean if a field has been set.

### GetAccountDrivenDeviceVisionosEnrollmentEnabled

`func (o *EnrollmentSettingsV4) GetAccountDrivenDeviceVisionosEnrollmentEnabled() bool`

GetAccountDrivenDeviceVisionosEnrollmentEnabled returns the AccountDrivenDeviceVisionosEnrollmentEnabled field if non-nil, zero value otherwise.

### GetAccountDrivenDeviceVisionosEnrollmentEnabledOk

`func (o *EnrollmentSettingsV4) GetAccountDrivenDeviceVisionosEnrollmentEnabledOk() (*bool, bool)`

GetAccountDrivenDeviceVisionosEnrollmentEnabledOk returns a tuple with the AccountDrivenDeviceVisionosEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDrivenDeviceVisionosEnrollmentEnabled

`func (o *EnrollmentSettingsV4) SetAccountDrivenDeviceVisionosEnrollmentEnabled(v bool)`

SetAccountDrivenDeviceVisionosEnrollmentEnabled sets AccountDrivenDeviceVisionosEnrollmentEnabled field to given value.

### HasAccountDrivenDeviceVisionosEnrollmentEnabled

`func (o *EnrollmentSettingsV4) HasAccountDrivenDeviceVisionosEnrollmentEnabled() bool`

HasAccountDrivenDeviceVisionosEnrollmentEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


