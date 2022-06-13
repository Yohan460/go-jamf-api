/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// EnrollmentSettings struct for EnrollmentSettings
type EnrollmentSettings struct {
	IsInstallSingleProfile *bool `json:"isInstallSingleProfile,omitempty"`
	IsSigningMdmProfileEnabled *bool `json:"isSigningMdmProfileEnabled,omitempty"`
	MdmSigningCertificate *MdmSigningCertificate `json:"mdmSigningCertificate,omitempty"`
	IsRestrictReenrollment *bool `json:"isRestrictReenrollment,omitempty"`
	IsFlushLocationInformation *bool `json:"isFlushLocationInformation,omitempty"`
	IsFlushLocationHistoryInformation *bool `json:"isFlushLocationHistoryInformation,omitempty"`
	IsFlushPolicyHistory *bool `json:"isFlushPolicyHistory,omitempty"`
	IsFlushExtensionAttributes *bool `json:"isFlushExtensionAttributes,omitempty"`
	FlushMdmCommandsOnReenroll string `json:"flushMdmCommandsOnReenroll"`
	IsEnabledMacosEnterpriseEnrollment bool `json:"isEnabledMacosEnterpriseEnrollment"`
	ManagementUsername string `json:"managementUsername"`
	ManagementPassword *string `json:"managementPassword,omitempty"`
	PasswordType string `json:"passwordType"`
	RandomPasswordLength *int32 `json:"randomPasswordLength,omitempty"`
	IsCreateManagementAccount *bool `json:"isCreateManagementAccount,omitempty"`
	IsHideManagementAccount *bool `json:"isHideManagementAccount,omitempty"`
	IsAllowSshOnlyManagementAccount *bool `json:"isAllowSshOnlyManagementAccount,omitempty"`
	IsEnsureSshRunning *bool `json:"isEnsureSshRunning,omitempty"`
	IsLaunchSelfService *bool `json:"isLaunchSelfService,omitempty"`
	IsSignQuickAdd *bool `json:"isSignQuickAdd,omitempty"`
	DeveloperCertificateIdentity *CertificateIdentityV1 `json:"developerCertificateIdentity,omitempty"`
	DeveloperCertificateIdentityDetails *CertificateDetails `json:"developerCertificateIdentityDetails,omitempty"`
	MdmSigningCertificateDetails *CertificateDetails `json:"mdmSigningCertificateDetails,omitempty"`
	IsEnableIosEnterpriseEnrollment *bool `json:"isEnableIosEnterpriseEnrollment,omitempty"`
	IsEnableIosPersonalEnrollment *bool `json:"isEnableIosPersonalEnrollment,omitempty"`
	PersonalDeviceEnrollmentType string `json:"personalDeviceEnrollmentType"`
}

// NewEnrollmentSettings instantiates a new EnrollmentSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrollmentSettings(flushMdmCommandsOnReenroll string, isEnabledMacosEnterpriseEnrollment bool, managementUsername string, passwordType string, personalDeviceEnrollmentType string) *EnrollmentSettings {
	this := EnrollmentSettings{}
	var isInstallSingleProfile bool = false
	this.IsInstallSingleProfile = &isInstallSingleProfile
	var isSigningMdmProfileEnabled bool = false
	this.IsSigningMdmProfileEnabled = &isSigningMdmProfileEnabled
	var isRestrictReenrollment bool = false
	this.IsRestrictReenrollment = &isRestrictReenrollment
	var isFlushLocationInformation bool = false
	this.IsFlushLocationInformation = &isFlushLocationInformation
	var isFlushLocationHistoryInformation bool = false
	this.IsFlushLocationHistoryInformation = &isFlushLocationHistoryInformation
	var isFlushPolicyHistory bool = false
	this.IsFlushPolicyHistory = &isFlushPolicyHistory
	var isFlushExtensionAttributes bool = false
	this.IsFlushExtensionAttributes = &isFlushExtensionAttributes
	this.FlushMdmCommandsOnReenroll = flushMdmCommandsOnReenroll
	this.IsEnabledMacosEnterpriseEnrollment = isEnabledMacosEnterpriseEnrollment
	this.ManagementUsername = managementUsername
	var managementPassword string = "null"
	this.ManagementPassword = &managementPassword
	this.PasswordType = passwordType
	var randomPasswordLength int32 = 8
	this.RandomPasswordLength = &randomPasswordLength
	var isCreateManagementAccount bool = true
	this.IsCreateManagementAccount = &isCreateManagementAccount
	var isHideManagementAccount bool = false
	this.IsHideManagementAccount = &isHideManagementAccount
	var isAllowSshOnlyManagementAccount bool = false
	this.IsAllowSshOnlyManagementAccount = &isAllowSshOnlyManagementAccount
	var isEnsureSshRunning bool = true
	this.IsEnsureSshRunning = &isEnsureSshRunning
	var isLaunchSelfService bool = false
	this.IsLaunchSelfService = &isLaunchSelfService
	var isSignQuickAdd bool = false
	this.IsSignQuickAdd = &isSignQuickAdd
	var isEnableIosEnterpriseEnrollment bool = true
	this.IsEnableIosEnterpriseEnrollment = &isEnableIosEnterpriseEnrollment
	var isEnableIosPersonalEnrollment bool = false
	this.IsEnableIosPersonalEnrollment = &isEnableIosPersonalEnrollment
	this.PersonalDeviceEnrollmentType = personalDeviceEnrollmentType
	return &this
}

// NewEnrollmentSettingsWithDefaults instantiates a new EnrollmentSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrollmentSettingsWithDefaults() *EnrollmentSettings {
	this := EnrollmentSettings{}
	var isInstallSingleProfile bool = false
	this.IsInstallSingleProfile = &isInstallSingleProfile
	var isSigningMdmProfileEnabled bool = false
	this.IsSigningMdmProfileEnabled = &isSigningMdmProfileEnabled
	var isRestrictReenrollment bool = false
	this.IsRestrictReenrollment = &isRestrictReenrollment
	var isFlushLocationInformation bool = false
	this.IsFlushLocationInformation = &isFlushLocationInformation
	var isFlushLocationHistoryInformation bool = false
	this.IsFlushLocationHistoryInformation = &isFlushLocationHistoryInformation
	var isFlushPolicyHistory bool = false
	this.IsFlushPolicyHistory = &isFlushPolicyHistory
	var isFlushExtensionAttributes bool = false
	this.IsFlushExtensionAttributes = &isFlushExtensionAttributes
	var flushMdmCommandsOnReenroll string = "DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED"
	this.FlushMdmCommandsOnReenroll = flushMdmCommandsOnReenroll
	var isEnabledMacosEnterpriseEnrollment bool = false
	this.IsEnabledMacosEnterpriseEnrollment = isEnabledMacosEnterpriseEnrollment
	var managementUsername string = ""
	this.ManagementUsername = managementUsername
	var managementPassword string = "null"
	this.ManagementPassword = &managementPassword
	var passwordType string = "STATIC"
	this.PasswordType = passwordType
	var randomPasswordLength int32 = 8
	this.RandomPasswordLength = &randomPasswordLength
	var isCreateManagementAccount bool = true
	this.IsCreateManagementAccount = &isCreateManagementAccount
	var isHideManagementAccount bool = false
	this.IsHideManagementAccount = &isHideManagementAccount
	var isAllowSshOnlyManagementAccount bool = false
	this.IsAllowSshOnlyManagementAccount = &isAllowSshOnlyManagementAccount
	var isEnsureSshRunning bool = true
	this.IsEnsureSshRunning = &isEnsureSshRunning
	var isLaunchSelfService bool = false
	this.IsLaunchSelfService = &isLaunchSelfService
	var isSignQuickAdd bool = false
	this.IsSignQuickAdd = &isSignQuickAdd
	var isEnableIosEnterpriseEnrollment bool = true
	this.IsEnableIosEnterpriseEnrollment = &isEnableIosEnterpriseEnrollment
	var isEnableIosPersonalEnrollment bool = false
	this.IsEnableIosPersonalEnrollment = &isEnableIosPersonalEnrollment
	var personalDeviceEnrollmentType string = "PERSONALDEVICEPROFILES"
	this.PersonalDeviceEnrollmentType = personalDeviceEnrollmentType
	return &this
}

// GetIsInstallSingleProfile returns the IsInstallSingleProfile field value if set, zero value otherwise.
func (o *EnrollmentSettings) GetIsInstallSingleProfile() bool {
	if o == nil || o.IsInstallSingleProfile == nil {
		var ret bool
		return ret
	}
	return *o.IsInstallSingleProfile
}

// GetIsInstallSingleProfileOk returns a tuple with the IsInstallSingleProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetIsInstallSingleProfileOk() (*bool, bool) {
	if o == nil || o.IsInstallSingleProfile == nil {
		return nil, false
	}
	return o.IsInstallSingleProfile, true
}

// HasIsInstallSingleProfile returns a boolean if a field has been set.
func (o *EnrollmentSettings) HasIsInstallSingleProfile() bool {
	if o != nil && o.IsInstallSingleProfile != nil {
		return true
	}

	return false
}

// SetIsInstallSingleProfile gets a reference to the given bool and assigns it to the IsInstallSingleProfile field.
func (o *EnrollmentSettings) SetIsInstallSingleProfile(v bool) {
	o.IsInstallSingleProfile = &v
}

// GetIsSigningMdmProfileEnabled returns the IsSigningMdmProfileEnabled field value if set, zero value otherwise.
func (o *EnrollmentSettings) GetIsSigningMdmProfileEnabled() bool {
	if o == nil || o.IsSigningMdmProfileEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsSigningMdmProfileEnabled
}

// GetIsSigningMdmProfileEnabledOk returns a tuple with the IsSigningMdmProfileEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetIsSigningMdmProfileEnabledOk() (*bool, bool) {
	if o == nil || o.IsSigningMdmProfileEnabled == nil {
		return nil, false
	}
	return o.IsSigningMdmProfileEnabled, true
}

// HasIsSigningMdmProfileEnabled returns a boolean if a field has been set.
func (o *EnrollmentSettings) HasIsSigningMdmProfileEnabled() bool {
	if o != nil && o.IsSigningMdmProfileEnabled != nil {
		return true
	}

	return false
}

// SetIsSigningMdmProfileEnabled gets a reference to the given bool and assigns it to the IsSigningMdmProfileEnabled field.
func (o *EnrollmentSettings) SetIsSigningMdmProfileEnabled(v bool) {
	o.IsSigningMdmProfileEnabled = &v
}

// GetMdmSigningCertificate returns the MdmSigningCertificate field value if set, zero value otherwise.
func (o *EnrollmentSettings) GetMdmSigningCertificate() MdmSigningCertificate {
	if o == nil || o.MdmSigningCertificate == nil {
		var ret MdmSigningCertificate
		return ret
	}
	return *o.MdmSigningCertificate
}

// GetMdmSigningCertificateOk returns a tuple with the MdmSigningCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetMdmSigningCertificateOk() (*MdmSigningCertificate, bool) {
	if o == nil || o.MdmSigningCertificate == nil {
		return nil, false
	}
	return o.MdmSigningCertificate, true
}

// HasMdmSigningCertificate returns a boolean if a field has been set.
func (o *EnrollmentSettings) HasMdmSigningCertificate() bool {
	if o != nil && o.MdmSigningCertificate != nil {
		return true
	}

	return false
}

// SetMdmSigningCertificate gets a reference to the given MdmSigningCertificate and assigns it to the MdmSigningCertificate field.
func (o *EnrollmentSettings) SetMdmSigningCertificate(v MdmSigningCertificate) {
	o.MdmSigningCertificate = &v
}

// GetIsRestrictReenrollment returns the IsRestrictReenrollment field value if set, zero value otherwise.
func (o *EnrollmentSettings) GetIsRestrictReenrollment() bool {
	if o == nil || o.IsRestrictReenrollment == nil {
		var ret bool
		return ret
	}
	return *o.IsRestrictReenrollment
}

// GetIsRestrictReenrollmentOk returns a tuple with the IsRestrictReenrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetIsRestrictReenrollmentOk() (*bool, bool) {
	if o == nil || o.IsRestrictReenrollment == nil {
		return nil, false
	}
	return o.IsRestrictReenrollment, true
}

// HasIsRestrictReenrollment returns a boolean if a field has been set.
func (o *EnrollmentSettings) HasIsRestrictReenrollment() bool {
	if o != nil && o.IsRestrictReenrollment != nil {
		return true
	}

	return false
}

// SetIsRestrictReenrollment gets a reference to the given bool and assigns it to the IsRestrictReenrollment field.
func (o *EnrollmentSettings) SetIsRestrictReenrollment(v bool) {
	o.IsRestrictReenrollment = &v
}

// GetIsFlushLocationInformation returns the IsFlushLocationInformation field value if set, zero value otherwise.
func (o *EnrollmentSettings) GetIsFlushLocationInformation() bool {
	if o == nil || o.IsFlushLocationInformation == nil {
		var ret bool
		return ret
	}
	return *o.IsFlushLocationInformation
}

// GetIsFlushLocationInformationOk returns a tuple with the IsFlushLocationInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetIsFlushLocationInformationOk() (*bool, bool) {
	if o == nil || o.IsFlushLocationInformation == nil {
		return nil, false
	}
	return o.IsFlushLocationInformation, true
}

// HasIsFlushLocationInformation returns a boolean if a field has been set.
func (o *EnrollmentSettings) HasIsFlushLocationInformation() bool {
	if o != nil && o.IsFlushLocationInformation != nil {
		return true
	}

	return false
}

// SetIsFlushLocationInformation gets a reference to the given bool and assigns it to the IsFlushLocationInformation field.
func (o *EnrollmentSettings) SetIsFlushLocationInformation(v bool) {
	o.IsFlushLocationInformation = &v
}

// GetIsFlushLocationHistoryInformation returns the IsFlushLocationHistoryInformation field value if set, zero value otherwise.
func (o *EnrollmentSettings) GetIsFlushLocationHistoryInformation() bool {
	if o == nil || o.IsFlushLocationHistoryInformation == nil {
		var ret bool
		return ret
	}
	return *o.IsFlushLocationHistoryInformation
}

// GetIsFlushLocationHistoryInformationOk returns a tuple with the IsFlushLocationHistoryInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetIsFlushLocationHistoryInformationOk() (*bool, bool) {
	if o == nil || o.IsFlushLocationHistoryInformation == nil {
		return nil, false
	}
	return o.IsFlushLocationHistoryInformation, true
}

// HasIsFlushLocationHistoryInformation returns a boolean if a field has been set.
func (o *EnrollmentSettings) HasIsFlushLocationHistoryInformation() bool {
	if o != nil && o.IsFlushLocationHistoryInformation != nil {
		return true
	}

	return false
}

// SetIsFlushLocationHistoryInformation gets a reference to the given bool and assigns it to the IsFlushLocationHistoryInformation field.
func (o *EnrollmentSettings) SetIsFlushLocationHistoryInformation(v bool) {
	o.IsFlushLocationHistoryInformation = &v
}

// GetIsFlushPolicyHistory returns the IsFlushPolicyHistory field value if set, zero value otherwise.
func (o *EnrollmentSettings) GetIsFlushPolicyHistory() bool {
	if o == nil || o.IsFlushPolicyHistory == nil {
		var ret bool
		return ret
	}
	return *o.IsFlushPolicyHistory
}

// GetIsFlushPolicyHistoryOk returns a tuple with the IsFlushPolicyHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetIsFlushPolicyHistoryOk() (*bool, bool) {
	if o == nil || o.IsFlushPolicyHistory == nil {
		return nil, false
	}
	return o.IsFlushPolicyHistory, true
}

// HasIsFlushPolicyHistory returns a boolean if a field has been set.
func (o *EnrollmentSettings) HasIsFlushPolicyHistory() bool {
	if o != nil && o.IsFlushPolicyHistory != nil {
		return true
	}

	return false
}

// SetIsFlushPolicyHistory gets a reference to the given bool and assigns it to the IsFlushPolicyHistory field.
func (o *EnrollmentSettings) SetIsFlushPolicyHistory(v bool) {
	o.IsFlushPolicyHistory = &v
}

// GetIsFlushExtensionAttributes returns the IsFlushExtensionAttributes field value if set, zero value otherwise.
func (o *EnrollmentSettings) GetIsFlushExtensionAttributes() bool {
	if o == nil || o.IsFlushExtensionAttributes == nil {
		var ret bool
		return ret
	}
	return *o.IsFlushExtensionAttributes
}

// GetIsFlushExtensionAttributesOk returns a tuple with the IsFlushExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetIsFlushExtensionAttributesOk() (*bool, bool) {
	if o == nil || o.IsFlushExtensionAttributes == nil {
		return nil, false
	}
	return o.IsFlushExtensionAttributes, true
}

// HasIsFlushExtensionAttributes returns a boolean if a field has been set.
func (o *EnrollmentSettings) HasIsFlushExtensionAttributes() bool {
	if o != nil && o.IsFlushExtensionAttributes != nil {
		return true
	}

	return false
}

// SetIsFlushExtensionAttributes gets a reference to the given bool and assigns it to the IsFlushExtensionAttributes field.
func (o *EnrollmentSettings) SetIsFlushExtensionAttributes(v bool) {
	o.IsFlushExtensionAttributes = &v
}

// GetFlushMdmCommandsOnReenroll returns the FlushMdmCommandsOnReenroll field value
func (o *EnrollmentSettings) GetFlushMdmCommandsOnReenroll() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FlushMdmCommandsOnReenroll
}

// GetFlushMdmCommandsOnReenrollOk returns a tuple with the FlushMdmCommandsOnReenroll field value
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetFlushMdmCommandsOnReenrollOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlushMdmCommandsOnReenroll, true
}

// SetFlushMdmCommandsOnReenroll sets field value
func (o *EnrollmentSettings) SetFlushMdmCommandsOnReenroll(v string) {
	o.FlushMdmCommandsOnReenroll = v
}

// GetIsEnabledMacosEnterpriseEnrollment returns the IsEnabledMacosEnterpriseEnrollment field value
func (o *EnrollmentSettings) GetIsEnabledMacosEnterpriseEnrollment() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabledMacosEnterpriseEnrollment
}

// GetIsEnabledMacosEnterpriseEnrollmentOk returns a tuple with the IsEnabledMacosEnterpriseEnrollment field value
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetIsEnabledMacosEnterpriseEnrollmentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabledMacosEnterpriseEnrollment, true
}

// SetIsEnabledMacosEnterpriseEnrollment sets field value
func (o *EnrollmentSettings) SetIsEnabledMacosEnterpriseEnrollment(v bool) {
	o.IsEnabledMacosEnterpriseEnrollment = v
}

// GetManagementUsername returns the ManagementUsername field value
func (o *EnrollmentSettings) GetManagementUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ManagementUsername
}

// GetManagementUsernameOk returns a tuple with the ManagementUsername field value
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetManagementUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManagementUsername, true
}

// SetManagementUsername sets field value
func (o *EnrollmentSettings) SetManagementUsername(v string) {
	o.ManagementUsername = v
}

// GetManagementPassword returns the ManagementPassword field value if set, zero value otherwise.
func (o *EnrollmentSettings) GetManagementPassword() string {
	if o == nil || o.ManagementPassword == nil {
		var ret string
		return ret
	}
	return *o.ManagementPassword
}

// GetManagementPasswordOk returns a tuple with the ManagementPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetManagementPasswordOk() (*string, bool) {
	if o == nil || o.ManagementPassword == nil {
		return nil, false
	}
	return o.ManagementPassword, true
}

// HasManagementPassword returns a boolean if a field has been set.
func (o *EnrollmentSettings) HasManagementPassword() bool {
	if o != nil && o.ManagementPassword != nil {
		return true
	}

	return false
}

// SetManagementPassword gets a reference to the given string and assigns it to the ManagementPassword field.
func (o *EnrollmentSettings) SetManagementPassword(v string) {
	o.ManagementPassword = &v
}

// GetPasswordType returns the PasswordType field value
func (o *EnrollmentSettings) GetPasswordType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PasswordType
}

// GetPasswordTypeOk returns a tuple with the PasswordType field value
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetPasswordTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordType, true
}

// SetPasswordType sets field value
func (o *EnrollmentSettings) SetPasswordType(v string) {
	o.PasswordType = v
}

// GetRandomPasswordLength returns the RandomPasswordLength field value if set, zero value otherwise.
func (o *EnrollmentSettings) GetRandomPasswordLength() int32 {
	if o == nil || o.RandomPasswordLength == nil {
		var ret int32
		return ret
	}
	return *o.RandomPasswordLength
}

// GetRandomPasswordLengthOk returns a tuple with the RandomPasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetRandomPasswordLengthOk() (*int32, bool) {
	if o == nil || o.RandomPasswordLength == nil {
		return nil, false
	}
	return o.RandomPasswordLength, true
}

// HasRandomPasswordLength returns a boolean if a field has been set.
func (o *EnrollmentSettings) HasRandomPasswordLength() bool {
	if o != nil && o.RandomPasswordLength != nil {
		return true
	}

	return false
}

// SetRandomPasswordLength gets a reference to the given int32 and assigns it to the RandomPasswordLength field.
func (o *EnrollmentSettings) SetRandomPasswordLength(v int32) {
	o.RandomPasswordLength = &v
}

// GetIsCreateManagementAccount returns the IsCreateManagementAccount field value if set, zero value otherwise.
func (o *EnrollmentSettings) GetIsCreateManagementAccount() bool {
	if o == nil || o.IsCreateManagementAccount == nil {
		var ret bool
		return ret
	}
	return *o.IsCreateManagementAccount
}

// GetIsCreateManagementAccountOk returns a tuple with the IsCreateManagementAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetIsCreateManagementAccountOk() (*bool, bool) {
	if o == nil || o.IsCreateManagementAccount == nil {
		return nil, false
	}
	return o.IsCreateManagementAccount, true
}

// HasIsCreateManagementAccount returns a boolean if a field has been set.
func (o *EnrollmentSettings) HasIsCreateManagementAccount() bool {
	if o != nil && o.IsCreateManagementAccount != nil {
		return true
	}

	return false
}

// SetIsCreateManagementAccount gets a reference to the given bool and assigns it to the IsCreateManagementAccount field.
func (o *EnrollmentSettings) SetIsCreateManagementAccount(v bool) {
	o.IsCreateManagementAccount = &v
}

// GetIsHideManagementAccount returns the IsHideManagementAccount field value if set, zero value otherwise.
func (o *EnrollmentSettings) GetIsHideManagementAccount() bool {
	if o == nil || o.IsHideManagementAccount == nil {
		var ret bool
		return ret
	}
	return *o.IsHideManagementAccount
}

// GetIsHideManagementAccountOk returns a tuple with the IsHideManagementAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetIsHideManagementAccountOk() (*bool, bool) {
	if o == nil || o.IsHideManagementAccount == nil {
		return nil, false
	}
	return o.IsHideManagementAccount, true
}

// HasIsHideManagementAccount returns a boolean if a field has been set.
func (o *EnrollmentSettings) HasIsHideManagementAccount() bool {
	if o != nil && o.IsHideManagementAccount != nil {
		return true
	}

	return false
}

// SetIsHideManagementAccount gets a reference to the given bool and assigns it to the IsHideManagementAccount field.
func (o *EnrollmentSettings) SetIsHideManagementAccount(v bool) {
	o.IsHideManagementAccount = &v
}

// GetIsAllowSshOnlyManagementAccount returns the IsAllowSshOnlyManagementAccount field value if set, zero value otherwise.
func (o *EnrollmentSettings) GetIsAllowSshOnlyManagementAccount() bool {
	if o == nil || o.IsAllowSshOnlyManagementAccount == nil {
		var ret bool
		return ret
	}
	return *o.IsAllowSshOnlyManagementAccount
}

// GetIsAllowSshOnlyManagementAccountOk returns a tuple with the IsAllowSshOnlyManagementAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetIsAllowSshOnlyManagementAccountOk() (*bool, bool) {
	if o == nil || o.IsAllowSshOnlyManagementAccount == nil {
		return nil, false
	}
	return o.IsAllowSshOnlyManagementAccount, true
}

// HasIsAllowSshOnlyManagementAccount returns a boolean if a field has been set.
func (o *EnrollmentSettings) HasIsAllowSshOnlyManagementAccount() bool {
	if o != nil && o.IsAllowSshOnlyManagementAccount != nil {
		return true
	}

	return false
}

// SetIsAllowSshOnlyManagementAccount gets a reference to the given bool and assigns it to the IsAllowSshOnlyManagementAccount field.
func (o *EnrollmentSettings) SetIsAllowSshOnlyManagementAccount(v bool) {
	o.IsAllowSshOnlyManagementAccount = &v
}

// GetIsEnsureSshRunning returns the IsEnsureSshRunning field value if set, zero value otherwise.
func (o *EnrollmentSettings) GetIsEnsureSshRunning() bool {
	if o == nil || o.IsEnsureSshRunning == nil {
		var ret bool
		return ret
	}
	return *o.IsEnsureSshRunning
}

// GetIsEnsureSshRunningOk returns a tuple with the IsEnsureSshRunning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetIsEnsureSshRunningOk() (*bool, bool) {
	if o == nil || o.IsEnsureSshRunning == nil {
		return nil, false
	}
	return o.IsEnsureSshRunning, true
}

// HasIsEnsureSshRunning returns a boolean if a field has been set.
func (o *EnrollmentSettings) HasIsEnsureSshRunning() bool {
	if o != nil && o.IsEnsureSshRunning != nil {
		return true
	}

	return false
}

// SetIsEnsureSshRunning gets a reference to the given bool and assigns it to the IsEnsureSshRunning field.
func (o *EnrollmentSettings) SetIsEnsureSshRunning(v bool) {
	o.IsEnsureSshRunning = &v
}

// GetIsLaunchSelfService returns the IsLaunchSelfService field value if set, zero value otherwise.
func (o *EnrollmentSettings) GetIsLaunchSelfService() bool {
	if o == nil || o.IsLaunchSelfService == nil {
		var ret bool
		return ret
	}
	return *o.IsLaunchSelfService
}

// GetIsLaunchSelfServiceOk returns a tuple with the IsLaunchSelfService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetIsLaunchSelfServiceOk() (*bool, bool) {
	if o == nil || o.IsLaunchSelfService == nil {
		return nil, false
	}
	return o.IsLaunchSelfService, true
}

// HasIsLaunchSelfService returns a boolean if a field has been set.
func (o *EnrollmentSettings) HasIsLaunchSelfService() bool {
	if o != nil && o.IsLaunchSelfService != nil {
		return true
	}

	return false
}

// SetIsLaunchSelfService gets a reference to the given bool and assigns it to the IsLaunchSelfService field.
func (o *EnrollmentSettings) SetIsLaunchSelfService(v bool) {
	o.IsLaunchSelfService = &v
}

// GetIsSignQuickAdd returns the IsSignQuickAdd field value if set, zero value otherwise.
func (o *EnrollmentSettings) GetIsSignQuickAdd() bool {
	if o == nil || o.IsSignQuickAdd == nil {
		var ret bool
		return ret
	}
	return *o.IsSignQuickAdd
}

// GetIsSignQuickAddOk returns a tuple with the IsSignQuickAdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetIsSignQuickAddOk() (*bool, bool) {
	if o == nil || o.IsSignQuickAdd == nil {
		return nil, false
	}
	return o.IsSignQuickAdd, true
}

// HasIsSignQuickAdd returns a boolean if a field has been set.
func (o *EnrollmentSettings) HasIsSignQuickAdd() bool {
	if o != nil && o.IsSignQuickAdd != nil {
		return true
	}

	return false
}

// SetIsSignQuickAdd gets a reference to the given bool and assigns it to the IsSignQuickAdd field.
func (o *EnrollmentSettings) SetIsSignQuickAdd(v bool) {
	o.IsSignQuickAdd = &v
}

// GetDeveloperCertificateIdentity returns the DeveloperCertificateIdentity field value if set, zero value otherwise.
func (o *EnrollmentSettings) GetDeveloperCertificateIdentity() CertificateIdentityV1 {
	if o == nil || o.DeveloperCertificateIdentity == nil {
		var ret CertificateIdentityV1
		return ret
	}
	return *o.DeveloperCertificateIdentity
}

// GetDeveloperCertificateIdentityOk returns a tuple with the DeveloperCertificateIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetDeveloperCertificateIdentityOk() (*CertificateIdentityV1, bool) {
	if o == nil || o.DeveloperCertificateIdentity == nil {
		return nil, false
	}
	return o.DeveloperCertificateIdentity, true
}

// HasDeveloperCertificateIdentity returns a boolean if a field has been set.
func (o *EnrollmentSettings) HasDeveloperCertificateIdentity() bool {
	if o != nil && o.DeveloperCertificateIdentity != nil {
		return true
	}

	return false
}

// SetDeveloperCertificateIdentity gets a reference to the given CertificateIdentityV1 and assigns it to the DeveloperCertificateIdentity field.
func (o *EnrollmentSettings) SetDeveloperCertificateIdentity(v CertificateIdentityV1) {
	o.DeveloperCertificateIdentity = &v
}

// GetDeveloperCertificateIdentityDetails returns the DeveloperCertificateIdentityDetails field value if set, zero value otherwise.
func (o *EnrollmentSettings) GetDeveloperCertificateIdentityDetails() CertificateDetails {
	if o == nil || o.DeveloperCertificateIdentityDetails == nil {
		var ret CertificateDetails
		return ret
	}
	return *o.DeveloperCertificateIdentityDetails
}

// GetDeveloperCertificateIdentityDetailsOk returns a tuple with the DeveloperCertificateIdentityDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetDeveloperCertificateIdentityDetailsOk() (*CertificateDetails, bool) {
	if o == nil || o.DeveloperCertificateIdentityDetails == nil {
		return nil, false
	}
	return o.DeveloperCertificateIdentityDetails, true
}

// HasDeveloperCertificateIdentityDetails returns a boolean if a field has been set.
func (o *EnrollmentSettings) HasDeveloperCertificateIdentityDetails() bool {
	if o != nil && o.DeveloperCertificateIdentityDetails != nil {
		return true
	}

	return false
}

// SetDeveloperCertificateIdentityDetails gets a reference to the given CertificateDetails and assigns it to the DeveloperCertificateIdentityDetails field.
func (o *EnrollmentSettings) SetDeveloperCertificateIdentityDetails(v CertificateDetails) {
	o.DeveloperCertificateIdentityDetails = &v
}

// GetMdmSigningCertificateDetails returns the MdmSigningCertificateDetails field value if set, zero value otherwise.
func (o *EnrollmentSettings) GetMdmSigningCertificateDetails() CertificateDetails {
	if o == nil || o.MdmSigningCertificateDetails == nil {
		var ret CertificateDetails
		return ret
	}
	return *o.MdmSigningCertificateDetails
}

// GetMdmSigningCertificateDetailsOk returns a tuple with the MdmSigningCertificateDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetMdmSigningCertificateDetailsOk() (*CertificateDetails, bool) {
	if o == nil || o.MdmSigningCertificateDetails == nil {
		return nil, false
	}
	return o.MdmSigningCertificateDetails, true
}

// HasMdmSigningCertificateDetails returns a boolean if a field has been set.
func (o *EnrollmentSettings) HasMdmSigningCertificateDetails() bool {
	if o != nil && o.MdmSigningCertificateDetails != nil {
		return true
	}

	return false
}

// SetMdmSigningCertificateDetails gets a reference to the given CertificateDetails and assigns it to the MdmSigningCertificateDetails field.
func (o *EnrollmentSettings) SetMdmSigningCertificateDetails(v CertificateDetails) {
	o.MdmSigningCertificateDetails = &v
}

// GetIsEnableIosEnterpriseEnrollment returns the IsEnableIosEnterpriseEnrollment field value if set, zero value otherwise.
func (o *EnrollmentSettings) GetIsEnableIosEnterpriseEnrollment() bool {
	if o == nil || o.IsEnableIosEnterpriseEnrollment == nil {
		var ret bool
		return ret
	}
	return *o.IsEnableIosEnterpriseEnrollment
}

// GetIsEnableIosEnterpriseEnrollmentOk returns a tuple with the IsEnableIosEnterpriseEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetIsEnableIosEnterpriseEnrollmentOk() (*bool, bool) {
	if o == nil || o.IsEnableIosEnterpriseEnrollment == nil {
		return nil, false
	}
	return o.IsEnableIosEnterpriseEnrollment, true
}

// HasIsEnableIosEnterpriseEnrollment returns a boolean if a field has been set.
func (o *EnrollmentSettings) HasIsEnableIosEnterpriseEnrollment() bool {
	if o != nil && o.IsEnableIosEnterpriseEnrollment != nil {
		return true
	}

	return false
}

// SetIsEnableIosEnterpriseEnrollment gets a reference to the given bool and assigns it to the IsEnableIosEnterpriseEnrollment field.
func (o *EnrollmentSettings) SetIsEnableIosEnterpriseEnrollment(v bool) {
	o.IsEnableIosEnterpriseEnrollment = &v
}

// GetIsEnableIosPersonalEnrollment returns the IsEnableIosPersonalEnrollment field value if set, zero value otherwise.
func (o *EnrollmentSettings) GetIsEnableIosPersonalEnrollment() bool {
	if o == nil || o.IsEnableIosPersonalEnrollment == nil {
		var ret bool
		return ret
	}
	return *o.IsEnableIosPersonalEnrollment
}

// GetIsEnableIosPersonalEnrollmentOk returns a tuple with the IsEnableIosPersonalEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetIsEnableIosPersonalEnrollmentOk() (*bool, bool) {
	if o == nil || o.IsEnableIosPersonalEnrollment == nil {
		return nil, false
	}
	return o.IsEnableIosPersonalEnrollment, true
}

// HasIsEnableIosPersonalEnrollment returns a boolean if a field has been set.
func (o *EnrollmentSettings) HasIsEnableIosPersonalEnrollment() bool {
	if o != nil && o.IsEnableIosPersonalEnrollment != nil {
		return true
	}

	return false
}

// SetIsEnableIosPersonalEnrollment gets a reference to the given bool and assigns it to the IsEnableIosPersonalEnrollment field.
func (o *EnrollmentSettings) SetIsEnableIosPersonalEnrollment(v bool) {
	o.IsEnableIosPersonalEnrollment = &v
}

// GetPersonalDeviceEnrollmentType returns the PersonalDeviceEnrollmentType field value
func (o *EnrollmentSettings) GetPersonalDeviceEnrollmentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PersonalDeviceEnrollmentType
}

// GetPersonalDeviceEnrollmentTypeOk returns a tuple with the PersonalDeviceEnrollmentType field value
// and a boolean to check if the value has been set.
func (o *EnrollmentSettings) GetPersonalDeviceEnrollmentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PersonalDeviceEnrollmentType, true
}

// SetPersonalDeviceEnrollmentType sets field value
func (o *EnrollmentSettings) SetPersonalDeviceEnrollmentType(v string) {
	o.PersonalDeviceEnrollmentType = v
}

func (o EnrollmentSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsInstallSingleProfile != nil {
		toSerialize["isInstallSingleProfile"] = o.IsInstallSingleProfile
	}
	if o.IsSigningMdmProfileEnabled != nil {
		toSerialize["isSigningMdmProfileEnabled"] = o.IsSigningMdmProfileEnabled
	}
	if o.MdmSigningCertificate != nil {
		toSerialize["mdmSigningCertificate"] = o.MdmSigningCertificate
	}
	if o.IsRestrictReenrollment != nil {
		toSerialize["isRestrictReenrollment"] = o.IsRestrictReenrollment
	}
	if o.IsFlushLocationInformation != nil {
		toSerialize["isFlushLocationInformation"] = o.IsFlushLocationInformation
	}
	if o.IsFlushLocationHistoryInformation != nil {
		toSerialize["isFlushLocationHistoryInformation"] = o.IsFlushLocationHistoryInformation
	}
	if o.IsFlushPolicyHistory != nil {
		toSerialize["isFlushPolicyHistory"] = o.IsFlushPolicyHistory
	}
	if o.IsFlushExtensionAttributes != nil {
		toSerialize["isFlushExtensionAttributes"] = o.IsFlushExtensionAttributes
	}
	if true {
		toSerialize["flushMdmCommandsOnReenroll"] = o.FlushMdmCommandsOnReenroll
	}
	if true {
		toSerialize["isEnabledMacosEnterpriseEnrollment"] = o.IsEnabledMacosEnterpriseEnrollment
	}
	if true {
		toSerialize["managementUsername"] = o.ManagementUsername
	}
	if o.ManagementPassword != nil {
		toSerialize["managementPassword"] = o.ManagementPassword
	}
	if true {
		toSerialize["passwordType"] = o.PasswordType
	}
	if o.RandomPasswordLength != nil {
		toSerialize["randomPasswordLength"] = o.RandomPasswordLength
	}
	if o.IsCreateManagementAccount != nil {
		toSerialize["isCreateManagementAccount"] = o.IsCreateManagementAccount
	}
	if o.IsHideManagementAccount != nil {
		toSerialize["isHideManagementAccount"] = o.IsHideManagementAccount
	}
	if o.IsAllowSshOnlyManagementAccount != nil {
		toSerialize["isAllowSshOnlyManagementAccount"] = o.IsAllowSshOnlyManagementAccount
	}
	if o.IsEnsureSshRunning != nil {
		toSerialize["isEnsureSshRunning"] = o.IsEnsureSshRunning
	}
	if o.IsLaunchSelfService != nil {
		toSerialize["isLaunchSelfService"] = o.IsLaunchSelfService
	}
	if o.IsSignQuickAdd != nil {
		toSerialize["isSignQuickAdd"] = o.IsSignQuickAdd
	}
	if o.DeveloperCertificateIdentity != nil {
		toSerialize["developerCertificateIdentity"] = o.DeveloperCertificateIdentity
	}
	if o.DeveloperCertificateIdentityDetails != nil {
		toSerialize["developerCertificateIdentityDetails"] = o.DeveloperCertificateIdentityDetails
	}
	if o.MdmSigningCertificateDetails != nil {
		toSerialize["mdmSigningCertificateDetails"] = o.MdmSigningCertificateDetails
	}
	if o.IsEnableIosEnterpriseEnrollment != nil {
		toSerialize["isEnableIosEnterpriseEnrollment"] = o.IsEnableIosEnterpriseEnrollment
	}
	if o.IsEnableIosPersonalEnrollment != nil {
		toSerialize["isEnableIosPersonalEnrollment"] = o.IsEnableIosPersonalEnrollment
	}
	if true {
		toSerialize["personalDeviceEnrollmentType"] = o.PersonalDeviceEnrollmentType
	}
	return json.Marshal(toSerialize)
}

type NullableEnrollmentSettings struct {
	value *EnrollmentSettings
	isSet bool
}

func (v NullableEnrollmentSettings) Get() *EnrollmentSettings {
	return v.value
}

func (v *NullableEnrollmentSettings) Set(val *EnrollmentSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollmentSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollmentSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollmentSettings(val *EnrollmentSettings) *NullableEnrollmentSettings {
	return &NullableEnrollmentSettings{value: val, isSet: true}
}

func (v NullableEnrollmentSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrollmentSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


