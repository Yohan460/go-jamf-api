# EnrollmentSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsInstallSingleProfile** | Pointer to **bool** |  | [optional] [default to false]
**IsSigningMdmProfileEnabled** | Pointer to **bool** |  | [optional] [default to false]
**MdmSigningCertificate** | Pointer to [**MdmSigningCertificate**](MdmSigningCertificate.md) |  | [optional] 
**IsRestrictReenrollment** | Pointer to **bool** |  | [optional] [default to false]
**IsFlushLocationInformation** | Pointer to **bool** |  | [optional] [default to false]
**IsFlushLocationHistoryInformation** | Pointer to **bool** |  | [optional] [default to false]
**IsFlushPolicyHistory** | Pointer to **bool** |  | [optional] [default to false]
**IsFlushExtensionAttributes** | Pointer to **bool** |  | [optional] [default to false]
**FlushMdmCommandsOnReenroll** | **string** |  | [default to "DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED"]
**IsEnabledMacosEnterpriseEnrollment** | **bool** |  | [default to false]
**ManagementUsername** | **string** |  | [default to ""]
**ManagementPassword** | Pointer to **string** |  | [optional] [default to "null"]
**PasswordType** | **string** |  | [default to "STATIC"]
**RandomPasswordLength** | Pointer to **int32** |  | [optional] [default to 8]
**IsCreateManagementAccount** | Pointer to **bool** |  | [optional] [default to true]
**IsHideManagementAccount** | Pointer to **bool** |  | [optional] [default to false]
**IsAllowSshOnlyManagementAccount** | Pointer to **bool** |  | [optional] [default to false]
**IsEnsureSshRunning** | Pointer to **bool** |  | [optional] [default to true]
**IsLaunchSelfService** | Pointer to **bool** |  | [optional] [default to false]
**IsSignQuickAdd** | Pointer to **bool** |  | [optional] [default to false]
**DeveloperCertificateIdentity** | Pointer to [**CertificateIdentityV1**](CertificateIdentityV1.md) |  | [optional] 
**DeveloperCertificateIdentityDetails** | Pointer to [**CertificateDetails**](CertificateDetails.md) |  | [optional] 
**MdmSigningCertificateDetails** | Pointer to [**CertificateDetails**](CertificateDetails.md) |  | [optional] 
**IsEnableIosEnterpriseEnrollment** | Pointer to **bool** |  | [optional] [default to true]
**IsEnableIosPersonalEnrollment** | Pointer to **bool** |  | [optional] [default to false]
**PersonalDeviceEnrollmentType** | **string** |  | [default to "PERSONALDEVICEPROFILES"]

## Methods

### NewEnrollmentSettings

`func NewEnrollmentSettings(flushMdmCommandsOnReenroll string, isEnabledMacosEnterpriseEnrollment bool, managementUsername string, passwordType string, personalDeviceEnrollmentType string, ) *EnrollmentSettings`

NewEnrollmentSettings instantiates a new EnrollmentSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentSettingsWithDefaults

`func NewEnrollmentSettingsWithDefaults() *EnrollmentSettings`

NewEnrollmentSettingsWithDefaults instantiates a new EnrollmentSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsInstallSingleProfile

`func (o *EnrollmentSettings) GetIsInstallSingleProfile() bool`

GetIsInstallSingleProfile returns the IsInstallSingleProfile field if non-nil, zero value otherwise.

### GetIsInstallSingleProfileOk

`func (o *EnrollmentSettings) GetIsInstallSingleProfileOk() (*bool, bool)`

GetIsInstallSingleProfileOk returns a tuple with the IsInstallSingleProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInstallSingleProfile

`func (o *EnrollmentSettings) SetIsInstallSingleProfile(v bool)`

SetIsInstallSingleProfile sets IsInstallSingleProfile field to given value.

### HasIsInstallSingleProfile

`func (o *EnrollmentSettings) HasIsInstallSingleProfile() bool`

HasIsInstallSingleProfile returns a boolean if a field has been set.

### GetIsSigningMdmProfileEnabled

`func (o *EnrollmentSettings) GetIsSigningMdmProfileEnabled() bool`

GetIsSigningMdmProfileEnabled returns the IsSigningMdmProfileEnabled field if non-nil, zero value otherwise.

### GetIsSigningMdmProfileEnabledOk

`func (o *EnrollmentSettings) GetIsSigningMdmProfileEnabledOk() (*bool, bool)`

GetIsSigningMdmProfileEnabledOk returns a tuple with the IsSigningMdmProfileEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSigningMdmProfileEnabled

`func (o *EnrollmentSettings) SetIsSigningMdmProfileEnabled(v bool)`

SetIsSigningMdmProfileEnabled sets IsSigningMdmProfileEnabled field to given value.

### HasIsSigningMdmProfileEnabled

`func (o *EnrollmentSettings) HasIsSigningMdmProfileEnabled() bool`

HasIsSigningMdmProfileEnabled returns a boolean if a field has been set.

### GetMdmSigningCertificate

`func (o *EnrollmentSettings) GetMdmSigningCertificate() MdmSigningCertificate`

GetMdmSigningCertificate returns the MdmSigningCertificate field if non-nil, zero value otherwise.

### GetMdmSigningCertificateOk

`func (o *EnrollmentSettings) GetMdmSigningCertificateOk() (*MdmSigningCertificate, bool)`

GetMdmSigningCertificateOk returns a tuple with the MdmSigningCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmSigningCertificate

`func (o *EnrollmentSettings) SetMdmSigningCertificate(v MdmSigningCertificate)`

SetMdmSigningCertificate sets MdmSigningCertificate field to given value.

### HasMdmSigningCertificate

`func (o *EnrollmentSettings) HasMdmSigningCertificate() bool`

HasMdmSigningCertificate returns a boolean if a field has been set.

### GetIsRestrictReenrollment

`func (o *EnrollmentSettings) GetIsRestrictReenrollment() bool`

GetIsRestrictReenrollment returns the IsRestrictReenrollment field if non-nil, zero value otherwise.

### GetIsRestrictReenrollmentOk

`func (o *EnrollmentSettings) GetIsRestrictReenrollmentOk() (*bool, bool)`

GetIsRestrictReenrollmentOk returns a tuple with the IsRestrictReenrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRestrictReenrollment

`func (o *EnrollmentSettings) SetIsRestrictReenrollment(v bool)`

SetIsRestrictReenrollment sets IsRestrictReenrollment field to given value.

### HasIsRestrictReenrollment

`func (o *EnrollmentSettings) HasIsRestrictReenrollment() bool`

HasIsRestrictReenrollment returns a boolean if a field has been set.

### GetIsFlushLocationInformation

`func (o *EnrollmentSettings) GetIsFlushLocationInformation() bool`

GetIsFlushLocationInformation returns the IsFlushLocationInformation field if non-nil, zero value otherwise.

### GetIsFlushLocationInformationOk

`func (o *EnrollmentSettings) GetIsFlushLocationInformationOk() (*bool, bool)`

GetIsFlushLocationInformationOk returns a tuple with the IsFlushLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlushLocationInformation

`func (o *EnrollmentSettings) SetIsFlushLocationInformation(v bool)`

SetIsFlushLocationInformation sets IsFlushLocationInformation field to given value.

### HasIsFlushLocationInformation

`func (o *EnrollmentSettings) HasIsFlushLocationInformation() bool`

HasIsFlushLocationInformation returns a boolean if a field has been set.

### GetIsFlushLocationHistoryInformation

`func (o *EnrollmentSettings) GetIsFlushLocationHistoryInformation() bool`

GetIsFlushLocationHistoryInformation returns the IsFlushLocationHistoryInformation field if non-nil, zero value otherwise.

### GetIsFlushLocationHistoryInformationOk

`func (o *EnrollmentSettings) GetIsFlushLocationHistoryInformationOk() (*bool, bool)`

GetIsFlushLocationHistoryInformationOk returns a tuple with the IsFlushLocationHistoryInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlushLocationHistoryInformation

`func (o *EnrollmentSettings) SetIsFlushLocationHistoryInformation(v bool)`

SetIsFlushLocationHistoryInformation sets IsFlushLocationHistoryInformation field to given value.

### HasIsFlushLocationHistoryInformation

`func (o *EnrollmentSettings) HasIsFlushLocationHistoryInformation() bool`

HasIsFlushLocationHistoryInformation returns a boolean if a field has been set.

### GetIsFlushPolicyHistory

`func (o *EnrollmentSettings) GetIsFlushPolicyHistory() bool`

GetIsFlushPolicyHistory returns the IsFlushPolicyHistory field if non-nil, zero value otherwise.

### GetIsFlushPolicyHistoryOk

`func (o *EnrollmentSettings) GetIsFlushPolicyHistoryOk() (*bool, bool)`

GetIsFlushPolicyHistoryOk returns a tuple with the IsFlushPolicyHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlushPolicyHistory

`func (o *EnrollmentSettings) SetIsFlushPolicyHistory(v bool)`

SetIsFlushPolicyHistory sets IsFlushPolicyHistory field to given value.

### HasIsFlushPolicyHistory

`func (o *EnrollmentSettings) HasIsFlushPolicyHistory() bool`

HasIsFlushPolicyHistory returns a boolean if a field has been set.

### GetIsFlushExtensionAttributes

`func (o *EnrollmentSettings) GetIsFlushExtensionAttributes() bool`

GetIsFlushExtensionAttributes returns the IsFlushExtensionAttributes field if non-nil, zero value otherwise.

### GetIsFlushExtensionAttributesOk

`func (o *EnrollmentSettings) GetIsFlushExtensionAttributesOk() (*bool, bool)`

GetIsFlushExtensionAttributesOk returns a tuple with the IsFlushExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlushExtensionAttributes

`func (o *EnrollmentSettings) SetIsFlushExtensionAttributes(v bool)`

SetIsFlushExtensionAttributes sets IsFlushExtensionAttributes field to given value.

### HasIsFlushExtensionAttributes

`func (o *EnrollmentSettings) HasIsFlushExtensionAttributes() bool`

HasIsFlushExtensionAttributes returns a boolean if a field has been set.

### GetFlushMdmCommandsOnReenroll

`func (o *EnrollmentSettings) GetFlushMdmCommandsOnReenroll() string`

GetFlushMdmCommandsOnReenroll returns the FlushMdmCommandsOnReenroll field if non-nil, zero value otherwise.

### GetFlushMdmCommandsOnReenrollOk

`func (o *EnrollmentSettings) GetFlushMdmCommandsOnReenrollOk() (*string, bool)`

GetFlushMdmCommandsOnReenrollOk returns a tuple with the FlushMdmCommandsOnReenroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushMdmCommandsOnReenroll

`func (o *EnrollmentSettings) SetFlushMdmCommandsOnReenroll(v string)`

SetFlushMdmCommandsOnReenroll sets FlushMdmCommandsOnReenroll field to given value.


### GetIsEnabledMacosEnterpriseEnrollment

`func (o *EnrollmentSettings) GetIsEnabledMacosEnterpriseEnrollment() bool`

GetIsEnabledMacosEnterpriseEnrollment returns the IsEnabledMacosEnterpriseEnrollment field if non-nil, zero value otherwise.

### GetIsEnabledMacosEnterpriseEnrollmentOk

`func (o *EnrollmentSettings) GetIsEnabledMacosEnterpriseEnrollmentOk() (*bool, bool)`

GetIsEnabledMacosEnterpriseEnrollmentOk returns a tuple with the IsEnabledMacosEnterpriseEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledMacosEnterpriseEnrollment

`func (o *EnrollmentSettings) SetIsEnabledMacosEnterpriseEnrollment(v bool)`

SetIsEnabledMacosEnterpriseEnrollment sets IsEnabledMacosEnterpriseEnrollment field to given value.


### GetManagementUsername

`func (o *EnrollmentSettings) GetManagementUsername() string`

GetManagementUsername returns the ManagementUsername field if non-nil, zero value otherwise.

### GetManagementUsernameOk

`func (o *EnrollmentSettings) GetManagementUsernameOk() (*string, bool)`

GetManagementUsernameOk returns a tuple with the ManagementUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementUsername

`func (o *EnrollmentSettings) SetManagementUsername(v string)`

SetManagementUsername sets ManagementUsername field to given value.


### GetManagementPassword

`func (o *EnrollmentSettings) GetManagementPassword() string`

GetManagementPassword returns the ManagementPassword field if non-nil, zero value otherwise.

### GetManagementPasswordOk

`func (o *EnrollmentSettings) GetManagementPasswordOk() (*string, bool)`

GetManagementPasswordOk returns a tuple with the ManagementPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPassword

`func (o *EnrollmentSettings) SetManagementPassword(v string)`

SetManagementPassword sets ManagementPassword field to given value.

### HasManagementPassword

`func (o *EnrollmentSettings) HasManagementPassword() bool`

HasManagementPassword returns a boolean if a field has been set.

### GetPasswordType

`func (o *EnrollmentSettings) GetPasswordType() string`

GetPasswordType returns the PasswordType field if non-nil, zero value otherwise.

### GetPasswordTypeOk

`func (o *EnrollmentSettings) GetPasswordTypeOk() (*string, bool)`

GetPasswordTypeOk returns a tuple with the PasswordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordType

`func (o *EnrollmentSettings) SetPasswordType(v string)`

SetPasswordType sets PasswordType field to given value.


### GetRandomPasswordLength

`func (o *EnrollmentSettings) GetRandomPasswordLength() int32`

GetRandomPasswordLength returns the RandomPasswordLength field if non-nil, zero value otherwise.

### GetRandomPasswordLengthOk

`func (o *EnrollmentSettings) GetRandomPasswordLengthOk() (*int32, bool)`

GetRandomPasswordLengthOk returns a tuple with the RandomPasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomPasswordLength

`func (o *EnrollmentSettings) SetRandomPasswordLength(v int32)`

SetRandomPasswordLength sets RandomPasswordLength field to given value.

### HasRandomPasswordLength

`func (o *EnrollmentSettings) HasRandomPasswordLength() bool`

HasRandomPasswordLength returns a boolean if a field has been set.

### GetIsCreateManagementAccount

`func (o *EnrollmentSettings) GetIsCreateManagementAccount() bool`

GetIsCreateManagementAccount returns the IsCreateManagementAccount field if non-nil, zero value otherwise.

### GetIsCreateManagementAccountOk

`func (o *EnrollmentSettings) GetIsCreateManagementAccountOk() (*bool, bool)`

GetIsCreateManagementAccountOk returns a tuple with the IsCreateManagementAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCreateManagementAccount

`func (o *EnrollmentSettings) SetIsCreateManagementAccount(v bool)`

SetIsCreateManagementAccount sets IsCreateManagementAccount field to given value.

### HasIsCreateManagementAccount

`func (o *EnrollmentSettings) HasIsCreateManagementAccount() bool`

HasIsCreateManagementAccount returns a boolean if a field has been set.

### GetIsHideManagementAccount

`func (o *EnrollmentSettings) GetIsHideManagementAccount() bool`

GetIsHideManagementAccount returns the IsHideManagementAccount field if non-nil, zero value otherwise.

### GetIsHideManagementAccountOk

`func (o *EnrollmentSettings) GetIsHideManagementAccountOk() (*bool, bool)`

GetIsHideManagementAccountOk returns a tuple with the IsHideManagementAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHideManagementAccount

`func (o *EnrollmentSettings) SetIsHideManagementAccount(v bool)`

SetIsHideManagementAccount sets IsHideManagementAccount field to given value.

### HasIsHideManagementAccount

`func (o *EnrollmentSettings) HasIsHideManagementAccount() bool`

HasIsHideManagementAccount returns a boolean if a field has been set.

### GetIsAllowSshOnlyManagementAccount

`func (o *EnrollmentSettings) GetIsAllowSshOnlyManagementAccount() bool`

GetIsAllowSshOnlyManagementAccount returns the IsAllowSshOnlyManagementAccount field if non-nil, zero value otherwise.

### GetIsAllowSshOnlyManagementAccountOk

`func (o *EnrollmentSettings) GetIsAllowSshOnlyManagementAccountOk() (*bool, bool)`

GetIsAllowSshOnlyManagementAccountOk returns a tuple with the IsAllowSshOnlyManagementAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowSshOnlyManagementAccount

`func (o *EnrollmentSettings) SetIsAllowSshOnlyManagementAccount(v bool)`

SetIsAllowSshOnlyManagementAccount sets IsAllowSshOnlyManagementAccount field to given value.

### HasIsAllowSshOnlyManagementAccount

`func (o *EnrollmentSettings) HasIsAllowSshOnlyManagementAccount() bool`

HasIsAllowSshOnlyManagementAccount returns a boolean if a field has been set.

### GetIsEnsureSshRunning

`func (o *EnrollmentSettings) GetIsEnsureSshRunning() bool`

GetIsEnsureSshRunning returns the IsEnsureSshRunning field if non-nil, zero value otherwise.

### GetIsEnsureSshRunningOk

`func (o *EnrollmentSettings) GetIsEnsureSshRunningOk() (*bool, bool)`

GetIsEnsureSshRunningOk returns a tuple with the IsEnsureSshRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnsureSshRunning

`func (o *EnrollmentSettings) SetIsEnsureSshRunning(v bool)`

SetIsEnsureSshRunning sets IsEnsureSshRunning field to given value.

### HasIsEnsureSshRunning

`func (o *EnrollmentSettings) HasIsEnsureSshRunning() bool`

HasIsEnsureSshRunning returns a boolean if a field has been set.

### GetIsLaunchSelfService

`func (o *EnrollmentSettings) GetIsLaunchSelfService() bool`

GetIsLaunchSelfService returns the IsLaunchSelfService field if non-nil, zero value otherwise.

### GetIsLaunchSelfServiceOk

`func (o *EnrollmentSettings) GetIsLaunchSelfServiceOk() (*bool, bool)`

GetIsLaunchSelfServiceOk returns a tuple with the IsLaunchSelfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLaunchSelfService

`func (o *EnrollmentSettings) SetIsLaunchSelfService(v bool)`

SetIsLaunchSelfService sets IsLaunchSelfService field to given value.

### HasIsLaunchSelfService

`func (o *EnrollmentSettings) HasIsLaunchSelfService() bool`

HasIsLaunchSelfService returns a boolean if a field has been set.

### GetIsSignQuickAdd

`func (o *EnrollmentSettings) GetIsSignQuickAdd() bool`

GetIsSignQuickAdd returns the IsSignQuickAdd field if non-nil, zero value otherwise.

### GetIsSignQuickAddOk

`func (o *EnrollmentSettings) GetIsSignQuickAddOk() (*bool, bool)`

GetIsSignQuickAddOk returns a tuple with the IsSignQuickAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSignQuickAdd

`func (o *EnrollmentSettings) SetIsSignQuickAdd(v bool)`

SetIsSignQuickAdd sets IsSignQuickAdd field to given value.

### HasIsSignQuickAdd

`func (o *EnrollmentSettings) HasIsSignQuickAdd() bool`

HasIsSignQuickAdd returns a boolean if a field has been set.

### GetDeveloperCertificateIdentity

`func (o *EnrollmentSettings) GetDeveloperCertificateIdentity() CertificateIdentityV1`

GetDeveloperCertificateIdentity returns the DeveloperCertificateIdentity field if non-nil, zero value otherwise.

### GetDeveloperCertificateIdentityOk

`func (o *EnrollmentSettings) GetDeveloperCertificateIdentityOk() (*CertificateIdentityV1, bool)`

GetDeveloperCertificateIdentityOk returns a tuple with the DeveloperCertificateIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperCertificateIdentity

`func (o *EnrollmentSettings) SetDeveloperCertificateIdentity(v CertificateIdentityV1)`

SetDeveloperCertificateIdentity sets DeveloperCertificateIdentity field to given value.

### HasDeveloperCertificateIdentity

`func (o *EnrollmentSettings) HasDeveloperCertificateIdentity() bool`

HasDeveloperCertificateIdentity returns a boolean if a field has been set.

### GetDeveloperCertificateIdentityDetails

`func (o *EnrollmentSettings) GetDeveloperCertificateIdentityDetails() CertificateDetails`

GetDeveloperCertificateIdentityDetails returns the DeveloperCertificateIdentityDetails field if non-nil, zero value otherwise.

### GetDeveloperCertificateIdentityDetailsOk

`func (o *EnrollmentSettings) GetDeveloperCertificateIdentityDetailsOk() (*CertificateDetails, bool)`

GetDeveloperCertificateIdentityDetailsOk returns a tuple with the DeveloperCertificateIdentityDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperCertificateIdentityDetails

`func (o *EnrollmentSettings) SetDeveloperCertificateIdentityDetails(v CertificateDetails)`

SetDeveloperCertificateIdentityDetails sets DeveloperCertificateIdentityDetails field to given value.

### HasDeveloperCertificateIdentityDetails

`func (o *EnrollmentSettings) HasDeveloperCertificateIdentityDetails() bool`

HasDeveloperCertificateIdentityDetails returns a boolean if a field has been set.

### GetMdmSigningCertificateDetails

`func (o *EnrollmentSettings) GetMdmSigningCertificateDetails() CertificateDetails`

GetMdmSigningCertificateDetails returns the MdmSigningCertificateDetails field if non-nil, zero value otherwise.

### GetMdmSigningCertificateDetailsOk

`func (o *EnrollmentSettings) GetMdmSigningCertificateDetailsOk() (*CertificateDetails, bool)`

GetMdmSigningCertificateDetailsOk returns a tuple with the MdmSigningCertificateDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmSigningCertificateDetails

`func (o *EnrollmentSettings) SetMdmSigningCertificateDetails(v CertificateDetails)`

SetMdmSigningCertificateDetails sets MdmSigningCertificateDetails field to given value.

### HasMdmSigningCertificateDetails

`func (o *EnrollmentSettings) HasMdmSigningCertificateDetails() bool`

HasMdmSigningCertificateDetails returns a boolean if a field has been set.

### GetIsEnableIosEnterpriseEnrollment

`func (o *EnrollmentSettings) GetIsEnableIosEnterpriseEnrollment() bool`

GetIsEnableIosEnterpriseEnrollment returns the IsEnableIosEnterpriseEnrollment field if non-nil, zero value otherwise.

### GetIsEnableIosEnterpriseEnrollmentOk

`func (o *EnrollmentSettings) GetIsEnableIosEnterpriseEnrollmentOk() (*bool, bool)`

GetIsEnableIosEnterpriseEnrollmentOk returns a tuple with the IsEnableIosEnterpriseEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnableIosEnterpriseEnrollment

`func (o *EnrollmentSettings) SetIsEnableIosEnterpriseEnrollment(v bool)`

SetIsEnableIosEnterpriseEnrollment sets IsEnableIosEnterpriseEnrollment field to given value.

### HasIsEnableIosEnterpriseEnrollment

`func (o *EnrollmentSettings) HasIsEnableIosEnterpriseEnrollment() bool`

HasIsEnableIosEnterpriseEnrollment returns a boolean if a field has been set.

### GetIsEnableIosPersonalEnrollment

`func (o *EnrollmentSettings) GetIsEnableIosPersonalEnrollment() bool`

GetIsEnableIosPersonalEnrollment returns the IsEnableIosPersonalEnrollment field if non-nil, zero value otherwise.

### GetIsEnableIosPersonalEnrollmentOk

`func (o *EnrollmentSettings) GetIsEnableIosPersonalEnrollmentOk() (*bool, bool)`

GetIsEnableIosPersonalEnrollmentOk returns a tuple with the IsEnableIosPersonalEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnableIosPersonalEnrollment

`func (o *EnrollmentSettings) SetIsEnableIosPersonalEnrollment(v bool)`

SetIsEnableIosPersonalEnrollment sets IsEnableIosPersonalEnrollment field to given value.

### HasIsEnableIosPersonalEnrollment

`func (o *EnrollmentSettings) HasIsEnableIosPersonalEnrollment() bool`

HasIsEnableIosPersonalEnrollment returns a boolean if a field has been set.

### GetPersonalDeviceEnrollmentType

`func (o *EnrollmentSettings) GetPersonalDeviceEnrollmentType() string`

GetPersonalDeviceEnrollmentType returns the PersonalDeviceEnrollmentType field if non-nil, zero value otherwise.

### GetPersonalDeviceEnrollmentTypeOk

`func (o *EnrollmentSettings) GetPersonalDeviceEnrollmentTypeOk() (*string, bool)`

GetPersonalDeviceEnrollmentTypeOk returns a tuple with the PersonalDeviceEnrollmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalDeviceEnrollmentType

`func (o *EnrollmentSettings) SetPersonalDeviceEnrollmentType(v string)`

SetPersonalDeviceEnrollmentType sets PersonalDeviceEnrollmentType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


