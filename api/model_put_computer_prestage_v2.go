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

// checks if the PutComputerPrestageV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutComputerPrestageV2{}

// PutComputerPrestageV2 struct for PutComputerPrestageV2
type PutComputerPrestageV2 struct {
	DisplayName string `json:"displayName"`
	Mandatory bool `json:"mandatory"`
	MdmRemovable bool `json:"mdmRemovable"`
	SupportPhoneNumber string `json:"supportPhoneNumber"`
	SupportEmailAddress string `json:"supportEmailAddress"`
	Department string `json:"department"`
	DefaultPrestage bool `json:"defaultPrestage"`
	EnrollmentSiteId string `json:"enrollmentSiteId"`
	KeepExistingSiteMembership bool `json:"keepExistingSiteMembership"`
	KeepExistingLocationInformation bool `json:"keepExistingLocationInformation"`
	RequireAuthentication bool `json:"requireAuthentication"`
	AuthenticationPrompt string `json:"authenticationPrompt"`
	PreventActivationLock bool `json:"preventActivationLock"`
	EnableDeviceBasedActivationLock bool `json:"enableDeviceBasedActivationLock"`
	DeviceEnrollmentProgramInstanceId string `json:"deviceEnrollmentProgramInstanceId"`
	SkipSetupItems *map[string]bool `json:"skipSetupItems,omitempty"`
	LocationInformation LocationInformationV2 `json:"locationInformation"`
	PurchasingInformation PrestagePurchasingInformationV2 `json:"purchasingInformation"`
	// The Base64 encoded PEM Certificate
	AnchorCertificates []string `json:"anchorCertificates,omitempty"`
	EnrollmentCustomizationId *string `json:"enrollmentCustomizationId,omitempty"`
	Language *string `json:"language,omitempty"`
	Region *string `json:"region,omitempty"`
	AutoAdvanceSetup bool `json:"autoAdvanceSetup"`
	InstallProfilesDuringSetup bool `json:"installProfilesDuringSetup"`
	PrestageInstalledProfileIds []string `json:"prestageInstalledProfileIds"`
	CustomPackageIds []string `json:"customPackageIds"`
	CustomPackageDistributionPointId string `json:"customPackageDistributionPointId"`
	EnableRecoveryLock *bool `json:"enableRecoveryLock,omitempty"`
	RecoveryLockPasswordType *string `json:"recoveryLockPasswordType,omitempty"`
	RotateRecoveryLockPassword *bool `json:"rotateRecoveryLockPassword,omitempty"`
	RecoveryLockPassword *string `json:"recoveryLockPassword,omitempty"`
	VersionLock *int32 `json:"versionLock,omitempty"`
}

// NewPutComputerPrestageV2 instantiates a new PutComputerPrestageV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutComputerPrestageV2(displayName string, mandatory bool, mdmRemovable bool, supportPhoneNumber string, supportEmailAddress string, department string, defaultPrestage bool, enrollmentSiteId string, keepExistingSiteMembership bool, keepExistingLocationInformation bool, requireAuthentication bool, authenticationPrompt string, preventActivationLock bool, enableDeviceBasedActivationLock bool, deviceEnrollmentProgramInstanceId string, locationInformation LocationInformationV2, purchasingInformation PrestagePurchasingInformationV2, autoAdvanceSetup bool, installProfilesDuringSetup bool, prestageInstalledProfileIds []string, customPackageIds []string, customPackageDistributionPointId string) *PutComputerPrestageV2 {
	this := PutComputerPrestageV2{}
	this.DisplayName = displayName
	this.Mandatory = mandatory
	this.MdmRemovable = mdmRemovable
	this.SupportPhoneNumber = supportPhoneNumber
	this.SupportEmailAddress = supportEmailAddress
	this.Department = department
	this.DefaultPrestage = defaultPrestage
	this.EnrollmentSiteId = enrollmentSiteId
	this.KeepExistingSiteMembership = keepExistingSiteMembership
	this.KeepExistingLocationInformation = keepExistingLocationInformation
	this.RequireAuthentication = requireAuthentication
	this.AuthenticationPrompt = authenticationPrompt
	this.PreventActivationLock = preventActivationLock
	this.EnableDeviceBasedActivationLock = enableDeviceBasedActivationLock
	this.DeviceEnrollmentProgramInstanceId = deviceEnrollmentProgramInstanceId
	this.LocationInformation = locationInformation
	this.PurchasingInformation = purchasingInformation
	this.AutoAdvanceSetup = autoAdvanceSetup
	this.InstallProfilesDuringSetup = installProfilesDuringSetup
	this.PrestageInstalledProfileIds = prestageInstalledProfileIds
	this.CustomPackageIds = customPackageIds
	this.CustomPackageDistributionPointId = customPackageDistributionPointId
	return &this
}

// NewPutComputerPrestageV2WithDefaults instantiates a new PutComputerPrestageV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutComputerPrestageV2WithDefaults() *PutComputerPrestageV2 {
	this := PutComputerPrestageV2{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *PutComputerPrestageV2) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *PutComputerPrestageV2) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetMandatory returns the Mandatory field value
func (o *PutComputerPrestageV2) GetMandatory() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Mandatory
}

// GetMandatoryOk returns a tuple with the Mandatory field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetMandatoryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mandatory, true
}

// SetMandatory sets field value
func (o *PutComputerPrestageV2) SetMandatory(v bool) {
	o.Mandatory = v
}

// GetMdmRemovable returns the MdmRemovable field value
func (o *PutComputerPrestageV2) GetMdmRemovable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MdmRemovable
}

// GetMdmRemovableOk returns a tuple with the MdmRemovable field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetMdmRemovableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MdmRemovable, true
}

// SetMdmRemovable sets field value
func (o *PutComputerPrestageV2) SetMdmRemovable(v bool) {
	o.MdmRemovable = v
}

// GetSupportPhoneNumber returns the SupportPhoneNumber field value
func (o *PutComputerPrestageV2) GetSupportPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportPhoneNumber
}

// GetSupportPhoneNumberOk returns a tuple with the SupportPhoneNumber field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetSupportPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportPhoneNumber, true
}

// SetSupportPhoneNumber sets field value
func (o *PutComputerPrestageV2) SetSupportPhoneNumber(v string) {
	o.SupportPhoneNumber = v
}

// GetSupportEmailAddress returns the SupportEmailAddress field value
func (o *PutComputerPrestageV2) GetSupportEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportEmailAddress
}

// GetSupportEmailAddressOk returns a tuple with the SupportEmailAddress field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetSupportEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportEmailAddress, true
}

// SetSupportEmailAddress sets field value
func (o *PutComputerPrestageV2) SetSupportEmailAddress(v string) {
	o.SupportEmailAddress = v
}

// GetDepartment returns the Department field value
func (o *PutComputerPrestageV2) GetDepartment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Department
}

// GetDepartmentOk returns a tuple with the Department field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetDepartmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Department, true
}

// SetDepartment sets field value
func (o *PutComputerPrestageV2) SetDepartment(v string) {
	o.Department = v
}

// GetDefaultPrestage returns the DefaultPrestage field value
func (o *PutComputerPrestageV2) GetDefaultPrestage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DefaultPrestage
}

// GetDefaultPrestageOk returns a tuple with the DefaultPrestage field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetDefaultPrestageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultPrestage, true
}

// SetDefaultPrestage sets field value
func (o *PutComputerPrestageV2) SetDefaultPrestage(v bool) {
	o.DefaultPrestage = v
}

// GetEnrollmentSiteId returns the EnrollmentSiteId field value
func (o *PutComputerPrestageV2) GetEnrollmentSiteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnrollmentSiteId
}

// GetEnrollmentSiteIdOk returns a tuple with the EnrollmentSiteId field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetEnrollmentSiteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnrollmentSiteId, true
}

// SetEnrollmentSiteId sets field value
func (o *PutComputerPrestageV2) SetEnrollmentSiteId(v string) {
	o.EnrollmentSiteId = v
}

// GetKeepExistingSiteMembership returns the KeepExistingSiteMembership field value
func (o *PutComputerPrestageV2) GetKeepExistingSiteMembership() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.KeepExistingSiteMembership
}

// GetKeepExistingSiteMembershipOk returns a tuple with the KeepExistingSiteMembership field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetKeepExistingSiteMembershipOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeepExistingSiteMembership, true
}

// SetKeepExistingSiteMembership sets field value
func (o *PutComputerPrestageV2) SetKeepExistingSiteMembership(v bool) {
	o.KeepExistingSiteMembership = v
}

// GetKeepExistingLocationInformation returns the KeepExistingLocationInformation field value
func (o *PutComputerPrestageV2) GetKeepExistingLocationInformation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.KeepExistingLocationInformation
}

// GetKeepExistingLocationInformationOk returns a tuple with the KeepExistingLocationInformation field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetKeepExistingLocationInformationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeepExistingLocationInformation, true
}

// SetKeepExistingLocationInformation sets field value
func (o *PutComputerPrestageV2) SetKeepExistingLocationInformation(v bool) {
	o.KeepExistingLocationInformation = v
}

// GetRequireAuthentication returns the RequireAuthentication field value
func (o *PutComputerPrestageV2) GetRequireAuthentication() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequireAuthentication
}

// GetRequireAuthenticationOk returns a tuple with the RequireAuthentication field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetRequireAuthenticationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireAuthentication, true
}

// SetRequireAuthentication sets field value
func (o *PutComputerPrestageV2) SetRequireAuthentication(v bool) {
	o.RequireAuthentication = v
}

// GetAuthenticationPrompt returns the AuthenticationPrompt field value
func (o *PutComputerPrestageV2) GetAuthenticationPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationPrompt
}

// GetAuthenticationPromptOk returns a tuple with the AuthenticationPrompt field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetAuthenticationPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationPrompt, true
}

// SetAuthenticationPrompt sets field value
func (o *PutComputerPrestageV2) SetAuthenticationPrompt(v string) {
	o.AuthenticationPrompt = v
}

// GetPreventActivationLock returns the PreventActivationLock field value
func (o *PutComputerPrestageV2) GetPreventActivationLock() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PreventActivationLock
}

// GetPreventActivationLockOk returns a tuple with the PreventActivationLock field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetPreventActivationLockOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreventActivationLock, true
}

// SetPreventActivationLock sets field value
func (o *PutComputerPrestageV2) SetPreventActivationLock(v bool) {
	o.PreventActivationLock = v
}

// GetEnableDeviceBasedActivationLock returns the EnableDeviceBasedActivationLock field value
func (o *PutComputerPrestageV2) GetEnableDeviceBasedActivationLock() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableDeviceBasedActivationLock
}

// GetEnableDeviceBasedActivationLockOk returns a tuple with the EnableDeviceBasedActivationLock field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetEnableDeviceBasedActivationLockOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableDeviceBasedActivationLock, true
}

// SetEnableDeviceBasedActivationLock sets field value
func (o *PutComputerPrestageV2) SetEnableDeviceBasedActivationLock(v bool) {
	o.EnableDeviceBasedActivationLock = v
}

// GetDeviceEnrollmentProgramInstanceId returns the DeviceEnrollmentProgramInstanceId field value
func (o *PutComputerPrestageV2) GetDeviceEnrollmentProgramInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceEnrollmentProgramInstanceId
}

// GetDeviceEnrollmentProgramInstanceIdOk returns a tuple with the DeviceEnrollmentProgramInstanceId field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetDeviceEnrollmentProgramInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceEnrollmentProgramInstanceId, true
}

// SetDeviceEnrollmentProgramInstanceId sets field value
func (o *PutComputerPrestageV2) SetDeviceEnrollmentProgramInstanceId(v string) {
	o.DeviceEnrollmentProgramInstanceId = v
}

// GetSkipSetupItems returns the SkipSetupItems field value if set, zero value otherwise.
func (o *PutComputerPrestageV2) GetSkipSetupItems() map[string]bool {
	if o == nil || IsNil(o.SkipSetupItems) {
		var ret map[string]bool
		return ret
	}
	return *o.SkipSetupItems
}

// GetSkipSetupItemsOk returns a tuple with the SkipSetupItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetSkipSetupItemsOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.SkipSetupItems) {
		return nil, false
	}
	return o.SkipSetupItems, true
}

// HasSkipSetupItems returns a boolean if a field has been set.
func (o *PutComputerPrestageV2) HasSkipSetupItems() bool {
	if o != nil && !IsNil(o.SkipSetupItems) {
		return true
	}

	return false
}

// SetSkipSetupItems gets a reference to the given map[string]bool and assigns it to the SkipSetupItems field.
func (o *PutComputerPrestageV2) SetSkipSetupItems(v map[string]bool) {
	o.SkipSetupItems = &v
}

// GetLocationInformation returns the LocationInformation field value
func (o *PutComputerPrestageV2) GetLocationInformation() LocationInformationV2 {
	if o == nil {
		var ret LocationInformationV2
		return ret
	}

	return o.LocationInformation
}

// GetLocationInformationOk returns a tuple with the LocationInformation field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetLocationInformationOk() (*LocationInformationV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocationInformation, true
}

// SetLocationInformation sets field value
func (o *PutComputerPrestageV2) SetLocationInformation(v LocationInformationV2) {
	o.LocationInformation = v
}

// GetPurchasingInformation returns the PurchasingInformation field value
func (o *PutComputerPrestageV2) GetPurchasingInformation() PrestagePurchasingInformationV2 {
	if o == nil {
		var ret PrestagePurchasingInformationV2
		return ret
	}

	return o.PurchasingInformation
}

// GetPurchasingInformationOk returns a tuple with the PurchasingInformation field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetPurchasingInformationOk() (*PrestagePurchasingInformationV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchasingInformation, true
}

// SetPurchasingInformation sets field value
func (o *PutComputerPrestageV2) SetPurchasingInformation(v PrestagePurchasingInformationV2) {
	o.PurchasingInformation = v
}

// GetAnchorCertificates returns the AnchorCertificates field value if set, zero value otherwise.
func (o *PutComputerPrestageV2) GetAnchorCertificates() []string {
	if o == nil || IsNil(o.AnchorCertificates) {
		var ret []string
		return ret
	}
	return o.AnchorCertificates
}

// GetAnchorCertificatesOk returns a tuple with the AnchorCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetAnchorCertificatesOk() ([]string, bool) {
	if o == nil || IsNil(o.AnchorCertificates) {
		return nil, false
	}
	return o.AnchorCertificates, true
}

// HasAnchorCertificates returns a boolean if a field has been set.
func (o *PutComputerPrestageV2) HasAnchorCertificates() bool {
	if o != nil && !IsNil(o.AnchorCertificates) {
		return true
	}

	return false
}

// SetAnchorCertificates gets a reference to the given []string and assigns it to the AnchorCertificates field.
func (o *PutComputerPrestageV2) SetAnchorCertificates(v []string) {
	o.AnchorCertificates = v
}

// GetEnrollmentCustomizationId returns the EnrollmentCustomizationId field value if set, zero value otherwise.
func (o *PutComputerPrestageV2) GetEnrollmentCustomizationId() string {
	if o == nil || IsNil(o.EnrollmentCustomizationId) {
		var ret string
		return ret
	}
	return *o.EnrollmentCustomizationId
}

// GetEnrollmentCustomizationIdOk returns a tuple with the EnrollmentCustomizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetEnrollmentCustomizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnrollmentCustomizationId) {
		return nil, false
	}
	return o.EnrollmentCustomizationId, true
}

// HasEnrollmentCustomizationId returns a boolean if a field has been set.
func (o *PutComputerPrestageV2) HasEnrollmentCustomizationId() bool {
	if o != nil && !IsNil(o.EnrollmentCustomizationId) {
		return true
	}

	return false
}

// SetEnrollmentCustomizationId gets a reference to the given string and assigns it to the EnrollmentCustomizationId field.
func (o *PutComputerPrestageV2) SetEnrollmentCustomizationId(v string) {
	o.EnrollmentCustomizationId = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *PutComputerPrestageV2) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *PutComputerPrestageV2) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *PutComputerPrestageV2) SetLanguage(v string) {
	o.Language = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *PutComputerPrestageV2) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *PutComputerPrestageV2) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *PutComputerPrestageV2) SetRegion(v string) {
	o.Region = &v
}

// GetAutoAdvanceSetup returns the AutoAdvanceSetup field value
func (o *PutComputerPrestageV2) GetAutoAdvanceSetup() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoAdvanceSetup
}

// GetAutoAdvanceSetupOk returns a tuple with the AutoAdvanceSetup field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetAutoAdvanceSetupOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoAdvanceSetup, true
}

// SetAutoAdvanceSetup sets field value
func (o *PutComputerPrestageV2) SetAutoAdvanceSetup(v bool) {
	o.AutoAdvanceSetup = v
}

// GetInstallProfilesDuringSetup returns the InstallProfilesDuringSetup field value
func (o *PutComputerPrestageV2) GetInstallProfilesDuringSetup() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InstallProfilesDuringSetup
}

// GetInstallProfilesDuringSetupOk returns a tuple with the InstallProfilesDuringSetup field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetInstallProfilesDuringSetupOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstallProfilesDuringSetup, true
}

// SetInstallProfilesDuringSetup sets field value
func (o *PutComputerPrestageV2) SetInstallProfilesDuringSetup(v bool) {
	o.InstallProfilesDuringSetup = v
}

// GetPrestageInstalledProfileIds returns the PrestageInstalledProfileIds field value
func (o *PutComputerPrestageV2) GetPrestageInstalledProfileIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PrestageInstalledProfileIds
}

// GetPrestageInstalledProfileIdsOk returns a tuple with the PrestageInstalledProfileIds field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetPrestageInstalledProfileIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrestageInstalledProfileIds, true
}

// SetPrestageInstalledProfileIds sets field value
func (o *PutComputerPrestageV2) SetPrestageInstalledProfileIds(v []string) {
	o.PrestageInstalledProfileIds = v
}

// GetCustomPackageIds returns the CustomPackageIds field value
func (o *PutComputerPrestageV2) GetCustomPackageIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CustomPackageIds
}

// GetCustomPackageIdsOk returns a tuple with the CustomPackageIds field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetCustomPackageIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomPackageIds, true
}

// SetCustomPackageIds sets field value
func (o *PutComputerPrestageV2) SetCustomPackageIds(v []string) {
	o.CustomPackageIds = v
}

// GetCustomPackageDistributionPointId returns the CustomPackageDistributionPointId field value
func (o *PutComputerPrestageV2) GetCustomPackageDistributionPointId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomPackageDistributionPointId
}

// GetCustomPackageDistributionPointIdOk returns a tuple with the CustomPackageDistributionPointId field value
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetCustomPackageDistributionPointIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomPackageDistributionPointId, true
}

// SetCustomPackageDistributionPointId sets field value
func (o *PutComputerPrestageV2) SetCustomPackageDistributionPointId(v string) {
	o.CustomPackageDistributionPointId = v
}

// GetEnableRecoveryLock returns the EnableRecoveryLock field value if set, zero value otherwise.
func (o *PutComputerPrestageV2) GetEnableRecoveryLock() bool {
	if o == nil || IsNil(o.EnableRecoveryLock) {
		var ret bool
		return ret
	}
	return *o.EnableRecoveryLock
}

// GetEnableRecoveryLockOk returns a tuple with the EnableRecoveryLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetEnableRecoveryLockOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableRecoveryLock) {
		return nil, false
	}
	return o.EnableRecoveryLock, true
}

// HasEnableRecoveryLock returns a boolean if a field has been set.
func (o *PutComputerPrestageV2) HasEnableRecoveryLock() bool {
	if o != nil && !IsNil(o.EnableRecoveryLock) {
		return true
	}

	return false
}

// SetEnableRecoveryLock gets a reference to the given bool and assigns it to the EnableRecoveryLock field.
func (o *PutComputerPrestageV2) SetEnableRecoveryLock(v bool) {
	o.EnableRecoveryLock = &v
}

// GetRecoveryLockPasswordType returns the RecoveryLockPasswordType field value if set, zero value otherwise.
func (o *PutComputerPrestageV2) GetRecoveryLockPasswordType() string {
	if o == nil || IsNil(o.RecoveryLockPasswordType) {
		var ret string
		return ret
	}
	return *o.RecoveryLockPasswordType
}

// GetRecoveryLockPasswordTypeOk returns a tuple with the RecoveryLockPasswordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetRecoveryLockPasswordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecoveryLockPasswordType) {
		return nil, false
	}
	return o.RecoveryLockPasswordType, true
}

// HasRecoveryLockPasswordType returns a boolean if a field has been set.
func (o *PutComputerPrestageV2) HasRecoveryLockPasswordType() bool {
	if o != nil && !IsNil(o.RecoveryLockPasswordType) {
		return true
	}

	return false
}

// SetRecoveryLockPasswordType gets a reference to the given string and assigns it to the RecoveryLockPasswordType field.
func (o *PutComputerPrestageV2) SetRecoveryLockPasswordType(v string) {
	o.RecoveryLockPasswordType = &v
}

// GetRotateRecoveryLockPassword returns the RotateRecoveryLockPassword field value if set, zero value otherwise.
func (o *PutComputerPrestageV2) GetRotateRecoveryLockPassword() bool {
	if o == nil || IsNil(o.RotateRecoveryLockPassword) {
		var ret bool
		return ret
	}
	return *o.RotateRecoveryLockPassword
}

// GetRotateRecoveryLockPasswordOk returns a tuple with the RotateRecoveryLockPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetRotateRecoveryLockPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.RotateRecoveryLockPassword) {
		return nil, false
	}
	return o.RotateRecoveryLockPassword, true
}

// HasRotateRecoveryLockPassword returns a boolean if a field has been set.
func (o *PutComputerPrestageV2) HasRotateRecoveryLockPassword() bool {
	if o != nil && !IsNil(o.RotateRecoveryLockPassword) {
		return true
	}

	return false
}

// SetRotateRecoveryLockPassword gets a reference to the given bool and assigns it to the RotateRecoveryLockPassword field.
func (o *PutComputerPrestageV2) SetRotateRecoveryLockPassword(v bool) {
	o.RotateRecoveryLockPassword = &v
}

// GetRecoveryLockPassword returns the RecoveryLockPassword field value if set, zero value otherwise.
func (o *PutComputerPrestageV2) GetRecoveryLockPassword() string {
	if o == nil || IsNil(o.RecoveryLockPassword) {
		var ret string
		return ret
	}
	return *o.RecoveryLockPassword
}

// GetRecoveryLockPasswordOk returns a tuple with the RecoveryLockPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetRecoveryLockPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.RecoveryLockPassword) {
		return nil, false
	}
	return o.RecoveryLockPassword, true
}

// HasRecoveryLockPassword returns a boolean if a field has been set.
func (o *PutComputerPrestageV2) HasRecoveryLockPassword() bool {
	if o != nil && !IsNil(o.RecoveryLockPassword) {
		return true
	}

	return false
}

// SetRecoveryLockPassword gets a reference to the given string and assigns it to the RecoveryLockPassword field.
func (o *PutComputerPrestageV2) SetRecoveryLockPassword(v string) {
	o.RecoveryLockPassword = &v
}

// GetVersionLock returns the VersionLock field value if set, zero value otherwise.
func (o *PutComputerPrestageV2) GetVersionLock() int32 {
	if o == nil || IsNil(o.VersionLock) {
		var ret int32
		return ret
	}
	return *o.VersionLock
}

// GetVersionLockOk returns a tuple with the VersionLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutComputerPrestageV2) GetVersionLockOk() (*int32, bool) {
	if o == nil || IsNil(o.VersionLock) {
		return nil, false
	}
	return o.VersionLock, true
}

// HasVersionLock returns a boolean if a field has been set.
func (o *PutComputerPrestageV2) HasVersionLock() bool {
	if o != nil && !IsNil(o.VersionLock) {
		return true
	}

	return false
}

// SetVersionLock gets a reference to the given int32 and assigns it to the VersionLock field.
func (o *PutComputerPrestageV2) SetVersionLock(v int32) {
	o.VersionLock = &v
}

func (o PutComputerPrestageV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutComputerPrestageV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["mandatory"] = o.Mandatory
	toSerialize["mdmRemovable"] = o.MdmRemovable
	toSerialize["supportPhoneNumber"] = o.SupportPhoneNumber
	toSerialize["supportEmailAddress"] = o.SupportEmailAddress
	toSerialize["department"] = o.Department
	toSerialize["defaultPrestage"] = o.DefaultPrestage
	toSerialize["enrollmentSiteId"] = o.EnrollmentSiteId
	toSerialize["keepExistingSiteMembership"] = o.KeepExistingSiteMembership
	toSerialize["keepExistingLocationInformation"] = o.KeepExistingLocationInformation
	toSerialize["requireAuthentication"] = o.RequireAuthentication
	toSerialize["authenticationPrompt"] = o.AuthenticationPrompt
	toSerialize["preventActivationLock"] = o.PreventActivationLock
	toSerialize["enableDeviceBasedActivationLock"] = o.EnableDeviceBasedActivationLock
	toSerialize["deviceEnrollmentProgramInstanceId"] = o.DeviceEnrollmentProgramInstanceId
	if !IsNil(o.SkipSetupItems) {
		toSerialize["skipSetupItems"] = o.SkipSetupItems
	}
	toSerialize["locationInformation"] = o.LocationInformation
	toSerialize["purchasingInformation"] = o.PurchasingInformation
	if !IsNil(o.AnchorCertificates) {
		toSerialize["anchorCertificates"] = o.AnchorCertificates
	}
	if !IsNil(o.EnrollmentCustomizationId) {
		toSerialize["enrollmentCustomizationId"] = o.EnrollmentCustomizationId
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	toSerialize["autoAdvanceSetup"] = o.AutoAdvanceSetup
	toSerialize["installProfilesDuringSetup"] = o.InstallProfilesDuringSetup
	toSerialize["prestageInstalledProfileIds"] = o.PrestageInstalledProfileIds
	toSerialize["customPackageIds"] = o.CustomPackageIds
	toSerialize["customPackageDistributionPointId"] = o.CustomPackageDistributionPointId
	if !IsNil(o.EnableRecoveryLock) {
		toSerialize["enableRecoveryLock"] = o.EnableRecoveryLock
	}
	if !IsNil(o.RecoveryLockPasswordType) {
		toSerialize["recoveryLockPasswordType"] = o.RecoveryLockPasswordType
	}
	if !IsNil(o.RotateRecoveryLockPassword) {
		toSerialize["rotateRecoveryLockPassword"] = o.RotateRecoveryLockPassword
	}
	if !IsNil(o.RecoveryLockPassword) {
		toSerialize["recoveryLockPassword"] = o.RecoveryLockPassword
	}
	if !IsNil(o.VersionLock) {
		toSerialize["versionLock"] = o.VersionLock
	}
	return toSerialize, nil
}

type NullablePutComputerPrestageV2 struct {
	value *PutComputerPrestageV2
	isSet bool
}

func (v NullablePutComputerPrestageV2) Get() *PutComputerPrestageV2 {
	return v.value
}

func (v *NullablePutComputerPrestageV2) Set(val *PutComputerPrestageV2) {
	v.value = val
	v.isSet = true
}

func (v NullablePutComputerPrestageV2) IsSet() bool {
	return v.isSet
}

func (v *NullablePutComputerPrestageV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutComputerPrestageV2(val *PutComputerPrestageV2) *NullablePutComputerPrestageV2 {
	return &NullablePutComputerPrestageV2{value: val, isSet: true}
}

func (v NullablePutComputerPrestageV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutComputerPrestageV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


