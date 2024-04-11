/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MobileDevicePrestage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MobileDevicePrestage{}

// MobileDevicePrestage struct for MobileDevicePrestage
type MobileDevicePrestage struct {
	DisplayName string `json:"displayName"`
	IsMandatory bool `json:"isMandatory"`
	IsMdmRemovable bool `json:"isMdmRemovable"`
	SupportPhoneNumber string `json:"supportPhoneNumber"`
	SupportEmailAddress string `json:"supportEmailAddress"`
	Department string `json:"department"`
	IsDefaultPrestage bool `json:"isDefaultPrestage"`
	EnrollmentSiteId int64 `json:"enrollmentSiteId"`
	IsKeepExistingSiteMembership bool `json:"isKeepExistingSiteMembership"`
	IsKeepExistingLocationInformation bool `json:"isKeepExistingLocationInformation"`
	IsRequireAuthentication bool `json:"isRequireAuthentication"`
	AuthenticationPrompt string `json:"authenticationPrompt"`
	IsPreventActivationLock bool `json:"isPreventActivationLock"`
	IsEnableDeviceBasedActivationLock bool `json:"isEnableDeviceBasedActivationLock"`
	DeviceEnrollmentProgramInstanceId int64 `json:"deviceEnrollmentProgramInstanceId"`
	SkipSetupItems *map[string]bool `json:"skipSetupItems,omitempty"`
	LocationInformation LocationInformation `json:"locationInformation"`
	PurchasingInformation PrestagePurchasingInformation `json:"purchasingInformation"`
	// The Base64 encoded PEM Certificate
	AnchorCertificates []string `json:"anchorCertificates,omitempty"`
	EnrollmentCustomizationId *int64 `json:"enrollmentCustomizationId,omitempty"`
	IsAllowPairing bool `json:"isAllowPairing"`
	IsMultiUser bool `json:"isMultiUser"`
	IsSupervised bool `json:"isSupervised"`
	MaximumSharedAccounts int64 `json:"maximumSharedAccounts"`
	IsAutoAdvanceSetup bool `json:"isAutoAdvanceSetup"`
	IsConfigureDeviceBeforeSetupAssistant bool `json:"isConfigureDeviceBeforeSetupAssistant"`
	Language *string `json:"language,omitempty"`
	Region *string `json:"region,omitempty"`
	Names *MobileDevicePrestageNames `json:"names,omitempty"`
}

type _MobileDevicePrestage MobileDevicePrestage

// NewMobileDevicePrestage instantiates a new MobileDevicePrestage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileDevicePrestage(displayName string, isMandatory bool, isMdmRemovable bool, supportPhoneNumber string, supportEmailAddress string, department string, isDefaultPrestage bool, enrollmentSiteId int64, isKeepExistingSiteMembership bool, isKeepExistingLocationInformation bool, isRequireAuthentication bool, authenticationPrompt string, isPreventActivationLock bool, isEnableDeviceBasedActivationLock bool, deviceEnrollmentProgramInstanceId int64, locationInformation LocationInformation, purchasingInformation PrestagePurchasingInformation, isAllowPairing bool, isMultiUser bool, isSupervised bool, maximumSharedAccounts int64, isAutoAdvanceSetup bool, isConfigureDeviceBeforeSetupAssistant bool) *MobileDevicePrestage {
	this := MobileDevicePrestage{}
	this.DisplayName = displayName
	this.IsMandatory = isMandatory
	this.IsMdmRemovable = isMdmRemovable
	this.SupportPhoneNumber = supportPhoneNumber
	this.SupportEmailAddress = supportEmailAddress
	this.Department = department
	this.IsDefaultPrestage = isDefaultPrestage
	this.EnrollmentSiteId = enrollmentSiteId
	this.IsKeepExistingSiteMembership = isKeepExistingSiteMembership
	this.IsKeepExistingLocationInformation = isKeepExistingLocationInformation
	this.IsRequireAuthentication = isRequireAuthentication
	this.AuthenticationPrompt = authenticationPrompt
	this.IsPreventActivationLock = isPreventActivationLock
	this.IsEnableDeviceBasedActivationLock = isEnableDeviceBasedActivationLock
	this.DeviceEnrollmentProgramInstanceId = deviceEnrollmentProgramInstanceId
	this.LocationInformation = locationInformation
	this.PurchasingInformation = purchasingInformation
	this.IsAllowPairing = isAllowPairing
	this.IsMultiUser = isMultiUser
	this.IsSupervised = isSupervised
	this.MaximumSharedAccounts = maximumSharedAccounts
	this.IsAutoAdvanceSetup = isAutoAdvanceSetup
	this.IsConfigureDeviceBeforeSetupAssistant = isConfigureDeviceBeforeSetupAssistant
	return &this
}

// NewMobileDevicePrestageWithDefaults instantiates a new MobileDevicePrestage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileDevicePrestageWithDefaults() *MobileDevicePrestage {
	this := MobileDevicePrestage{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *MobileDevicePrestage) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *MobileDevicePrestage) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetIsMandatory returns the IsMandatory field value
func (o *MobileDevicePrestage) GetIsMandatory() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMandatory
}

// GetIsMandatoryOk returns a tuple with the IsMandatory field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetIsMandatoryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMandatory, true
}

// SetIsMandatory sets field value
func (o *MobileDevicePrestage) SetIsMandatory(v bool) {
	o.IsMandatory = v
}

// GetIsMdmRemovable returns the IsMdmRemovable field value
func (o *MobileDevicePrestage) GetIsMdmRemovable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMdmRemovable
}

// GetIsMdmRemovableOk returns a tuple with the IsMdmRemovable field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetIsMdmRemovableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMdmRemovable, true
}

// SetIsMdmRemovable sets field value
func (o *MobileDevicePrestage) SetIsMdmRemovable(v bool) {
	o.IsMdmRemovable = v
}

// GetSupportPhoneNumber returns the SupportPhoneNumber field value
func (o *MobileDevicePrestage) GetSupportPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportPhoneNumber
}

// GetSupportPhoneNumberOk returns a tuple with the SupportPhoneNumber field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetSupportPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportPhoneNumber, true
}

// SetSupportPhoneNumber sets field value
func (o *MobileDevicePrestage) SetSupportPhoneNumber(v string) {
	o.SupportPhoneNumber = v
}

// GetSupportEmailAddress returns the SupportEmailAddress field value
func (o *MobileDevicePrestage) GetSupportEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportEmailAddress
}

// GetSupportEmailAddressOk returns a tuple with the SupportEmailAddress field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetSupportEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportEmailAddress, true
}

// SetSupportEmailAddress sets field value
func (o *MobileDevicePrestage) SetSupportEmailAddress(v string) {
	o.SupportEmailAddress = v
}

// GetDepartment returns the Department field value
func (o *MobileDevicePrestage) GetDepartment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Department
}

// GetDepartmentOk returns a tuple with the Department field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetDepartmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Department, true
}

// SetDepartment sets field value
func (o *MobileDevicePrestage) SetDepartment(v string) {
	o.Department = v
}

// GetIsDefaultPrestage returns the IsDefaultPrestage field value
func (o *MobileDevicePrestage) GetIsDefaultPrestage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefaultPrestage
}

// GetIsDefaultPrestageOk returns a tuple with the IsDefaultPrestage field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetIsDefaultPrestageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefaultPrestage, true
}

// SetIsDefaultPrestage sets field value
func (o *MobileDevicePrestage) SetIsDefaultPrestage(v bool) {
	o.IsDefaultPrestage = v
}

// GetEnrollmentSiteId returns the EnrollmentSiteId field value
func (o *MobileDevicePrestage) GetEnrollmentSiteId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EnrollmentSiteId
}

// GetEnrollmentSiteIdOk returns a tuple with the EnrollmentSiteId field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetEnrollmentSiteIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnrollmentSiteId, true
}

// SetEnrollmentSiteId sets field value
func (o *MobileDevicePrestage) SetEnrollmentSiteId(v int64) {
	o.EnrollmentSiteId = v
}

// GetIsKeepExistingSiteMembership returns the IsKeepExistingSiteMembership field value
func (o *MobileDevicePrestage) GetIsKeepExistingSiteMembership() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsKeepExistingSiteMembership
}

// GetIsKeepExistingSiteMembershipOk returns a tuple with the IsKeepExistingSiteMembership field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetIsKeepExistingSiteMembershipOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsKeepExistingSiteMembership, true
}

// SetIsKeepExistingSiteMembership sets field value
func (o *MobileDevicePrestage) SetIsKeepExistingSiteMembership(v bool) {
	o.IsKeepExistingSiteMembership = v
}

// GetIsKeepExistingLocationInformation returns the IsKeepExistingLocationInformation field value
func (o *MobileDevicePrestage) GetIsKeepExistingLocationInformation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsKeepExistingLocationInformation
}

// GetIsKeepExistingLocationInformationOk returns a tuple with the IsKeepExistingLocationInformation field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetIsKeepExistingLocationInformationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsKeepExistingLocationInformation, true
}

// SetIsKeepExistingLocationInformation sets field value
func (o *MobileDevicePrestage) SetIsKeepExistingLocationInformation(v bool) {
	o.IsKeepExistingLocationInformation = v
}

// GetIsRequireAuthentication returns the IsRequireAuthentication field value
func (o *MobileDevicePrestage) GetIsRequireAuthentication() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRequireAuthentication
}

// GetIsRequireAuthenticationOk returns a tuple with the IsRequireAuthentication field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetIsRequireAuthenticationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRequireAuthentication, true
}

// SetIsRequireAuthentication sets field value
func (o *MobileDevicePrestage) SetIsRequireAuthentication(v bool) {
	o.IsRequireAuthentication = v
}

// GetAuthenticationPrompt returns the AuthenticationPrompt field value
func (o *MobileDevicePrestage) GetAuthenticationPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationPrompt
}

// GetAuthenticationPromptOk returns a tuple with the AuthenticationPrompt field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetAuthenticationPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationPrompt, true
}

// SetAuthenticationPrompt sets field value
func (o *MobileDevicePrestage) SetAuthenticationPrompt(v string) {
	o.AuthenticationPrompt = v
}

// GetIsPreventActivationLock returns the IsPreventActivationLock field value
func (o *MobileDevicePrestage) GetIsPreventActivationLock() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPreventActivationLock
}

// GetIsPreventActivationLockOk returns a tuple with the IsPreventActivationLock field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetIsPreventActivationLockOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPreventActivationLock, true
}

// SetIsPreventActivationLock sets field value
func (o *MobileDevicePrestage) SetIsPreventActivationLock(v bool) {
	o.IsPreventActivationLock = v
}

// GetIsEnableDeviceBasedActivationLock returns the IsEnableDeviceBasedActivationLock field value
func (o *MobileDevicePrestage) GetIsEnableDeviceBasedActivationLock() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnableDeviceBasedActivationLock
}

// GetIsEnableDeviceBasedActivationLockOk returns a tuple with the IsEnableDeviceBasedActivationLock field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetIsEnableDeviceBasedActivationLockOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnableDeviceBasedActivationLock, true
}

// SetIsEnableDeviceBasedActivationLock sets field value
func (o *MobileDevicePrestage) SetIsEnableDeviceBasedActivationLock(v bool) {
	o.IsEnableDeviceBasedActivationLock = v
}

// GetDeviceEnrollmentProgramInstanceId returns the DeviceEnrollmentProgramInstanceId field value
func (o *MobileDevicePrestage) GetDeviceEnrollmentProgramInstanceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DeviceEnrollmentProgramInstanceId
}

// GetDeviceEnrollmentProgramInstanceIdOk returns a tuple with the DeviceEnrollmentProgramInstanceId field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetDeviceEnrollmentProgramInstanceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceEnrollmentProgramInstanceId, true
}

// SetDeviceEnrollmentProgramInstanceId sets field value
func (o *MobileDevicePrestage) SetDeviceEnrollmentProgramInstanceId(v int64) {
	o.DeviceEnrollmentProgramInstanceId = v
}

// GetSkipSetupItems returns the SkipSetupItems field value if set, zero value otherwise.
func (o *MobileDevicePrestage) GetSkipSetupItems() map[string]bool {
	if o == nil || IsNil(o.SkipSetupItems) {
		var ret map[string]bool
		return ret
	}
	return *o.SkipSetupItems
}

// GetSkipSetupItemsOk returns a tuple with the SkipSetupItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetSkipSetupItemsOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.SkipSetupItems) {
		return nil, false
	}
	return o.SkipSetupItems, true
}

// HasSkipSetupItems returns a boolean if a field has been set.
func (o *MobileDevicePrestage) HasSkipSetupItems() bool {
	if o != nil && !IsNil(o.SkipSetupItems) {
		return true
	}

	return false
}

// SetSkipSetupItems gets a reference to the given map[string]bool and assigns it to the SkipSetupItems field.
func (o *MobileDevicePrestage) SetSkipSetupItems(v map[string]bool) {
	o.SkipSetupItems = &v
}

// GetLocationInformation returns the LocationInformation field value
func (o *MobileDevicePrestage) GetLocationInformation() LocationInformation {
	if o == nil {
		var ret LocationInformation
		return ret
	}

	return o.LocationInformation
}

// GetLocationInformationOk returns a tuple with the LocationInformation field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetLocationInformationOk() (*LocationInformation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocationInformation, true
}

// SetLocationInformation sets field value
func (o *MobileDevicePrestage) SetLocationInformation(v LocationInformation) {
	o.LocationInformation = v
}

// GetPurchasingInformation returns the PurchasingInformation field value
func (o *MobileDevicePrestage) GetPurchasingInformation() PrestagePurchasingInformation {
	if o == nil {
		var ret PrestagePurchasingInformation
		return ret
	}

	return o.PurchasingInformation
}

// GetPurchasingInformationOk returns a tuple with the PurchasingInformation field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetPurchasingInformationOk() (*PrestagePurchasingInformation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchasingInformation, true
}

// SetPurchasingInformation sets field value
func (o *MobileDevicePrestage) SetPurchasingInformation(v PrestagePurchasingInformation) {
	o.PurchasingInformation = v
}

// GetAnchorCertificates returns the AnchorCertificates field value if set, zero value otherwise.
func (o *MobileDevicePrestage) GetAnchorCertificates() []string {
	if o == nil || IsNil(o.AnchorCertificates) {
		var ret []string
		return ret
	}
	return o.AnchorCertificates
}

// GetAnchorCertificatesOk returns a tuple with the AnchorCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetAnchorCertificatesOk() ([]string, bool) {
	if o == nil || IsNil(o.AnchorCertificates) {
		return nil, false
	}
	return o.AnchorCertificates, true
}

// HasAnchorCertificates returns a boolean if a field has been set.
func (o *MobileDevicePrestage) HasAnchorCertificates() bool {
	if o != nil && !IsNil(o.AnchorCertificates) {
		return true
	}

	return false
}

// SetAnchorCertificates gets a reference to the given []string and assigns it to the AnchorCertificates field.
func (o *MobileDevicePrestage) SetAnchorCertificates(v []string) {
	o.AnchorCertificates = v
}

// GetEnrollmentCustomizationId returns the EnrollmentCustomizationId field value if set, zero value otherwise.
func (o *MobileDevicePrestage) GetEnrollmentCustomizationId() int64 {
	if o == nil || IsNil(o.EnrollmentCustomizationId) {
		var ret int64
		return ret
	}
	return *o.EnrollmentCustomizationId
}

// GetEnrollmentCustomizationIdOk returns a tuple with the EnrollmentCustomizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetEnrollmentCustomizationIdOk() (*int64, bool) {
	if o == nil || IsNil(o.EnrollmentCustomizationId) {
		return nil, false
	}
	return o.EnrollmentCustomizationId, true
}

// HasEnrollmentCustomizationId returns a boolean if a field has been set.
func (o *MobileDevicePrestage) HasEnrollmentCustomizationId() bool {
	if o != nil && !IsNil(o.EnrollmentCustomizationId) {
		return true
	}

	return false
}

// SetEnrollmentCustomizationId gets a reference to the given int64 and assigns it to the EnrollmentCustomizationId field.
func (o *MobileDevicePrestage) SetEnrollmentCustomizationId(v int64) {
	o.EnrollmentCustomizationId = &v
}

// GetIsAllowPairing returns the IsAllowPairing field value
func (o *MobileDevicePrestage) GetIsAllowPairing() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAllowPairing
}

// GetIsAllowPairingOk returns a tuple with the IsAllowPairing field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetIsAllowPairingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAllowPairing, true
}

// SetIsAllowPairing sets field value
func (o *MobileDevicePrestage) SetIsAllowPairing(v bool) {
	o.IsAllowPairing = v
}

// GetIsMultiUser returns the IsMultiUser field value
func (o *MobileDevicePrestage) GetIsMultiUser() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMultiUser
}

// GetIsMultiUserOk returns a tuple with the IsMultiUser field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetIsMultiUserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMultiUser, true
}

// SetIsMultiUser sets field value
func (o *MobileDevicePrestage) SetIsMultiUser(v bool) {
	o.IsMultiUser = v
}

// GetIsSupervised returns the IsSupervised field value
func (o *MobileDevicePrestage) GetIsSupervised() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSupervised
}

// GetIsSupervisedOk returns a tuple with the IsSupervised field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetIsSupervisedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSupervised, true
}

// SetIsSupervised sets field value
func (o *MobileDevicePrestage) SetIsSupervised(v bool) {
	o.IsSupervised = v
}

// GetMaximumSharedAccounts returns the MaximumSharedAccounts field value
func (o *MobileDevicePrestage) GetMaximumSharedAccounts() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaximumSharedAccounts
}

// GetMaximumSharedAccountsOk returns a tuple with the MaximumSharedAccounts field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetMaximumSharedAccountsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaximumSharedAccounts, true
}

// SetMaximumSharedAccounts sets field value
func (o *MobileDevicePrestage) SetMaximumSharedAccounts(v int64) {
	o.MaximumSharedAccounts = v
}

// GetIsAutoAdvanceSetup returns the IsAutoAdvanceSetup field value
func (o *MobileDevicePrestage) GetIsAutoAdvanceSetup() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAutoAdvanceSetup
}

// GetIsAutoAdvanceSetupOk returns a tuple with the IsAutoAdvanceSetup field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetIsAutoAdvanceSetupOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAutoAdvanceSetup, true
}

// SetIsAutoAdvanceSetup sets field value
func (o *MobileDevicePrestage) SetIsAutoAdvanceSetup(v bool) {
	o.IsAutoAdvanceSetup = v
}

// GetIsConfigureDeviceBeforeSetupAssistant returns the IsConfigureDeviceBeforeSetupAssistant field value
func (o *MobileDevicePrestage) GetIsConfigureDeviceBeforeSetupAssistant() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsConfigureDeviceBeforeSetupAssistant
}

// GetIsConfigureDeviceBeforeSetupAssistantOk returns a tuple with the IsConfigureDeviceBeforeSetupAssistant field value
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetIsConfigureDeviceBeforeSetupAssistantOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsConfigureDeviceBeforeSetupAssistant, true
}

// SetIsConfigureDeviceBeforeSetupAssistant sets field value
func (o *MobileDevicePrestage) SetIsConfigureDeviceBeforeSetupAssistant(v bool) {
	o.IsConfigureDeviceBeforeSetupAssistant = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *MobileDevicePrestage) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *MobileDevicePrestage) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *MobileDevicePrestage) SetLanguage(v string) {
	o.Language = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *MobileDevicePrestage) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *MobileDevicePrestage) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *MobileDevicePrestage) SetRegion(v string) {
	o.Region = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *MobileDevicePrestage) GetNames() MobileDevicePrestageNames {
	if o == nil || IsNil(o.Names) {
		var ret MobileDevicePrestageNames
		return ret
	}
	return *o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDevicePrestage) GetNamesOk() (*MobileDevicePrestageNames, bool) {
	if o == nil || IsNil(o.Names) {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *MobileDevicePrestage) HasNames() bool {
	if o != nil && !IsNil(o.Names) {
		return true
	}

	return false
}

// SetNames gets a reference to the given MobileDevicePrestageNames and assigns it to the Names field.
func (o *MobileDevicePrestage) SetNames(v MobileDevicePrestageNames) {
	o.Names = &v
}

func (o MobileDevicePrestage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MobileDevicePrestage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["isMandatory"] = o.IsMandatory
	toSerialize["isMdmRemovable"] = o.IsMdmRemovable
	toSerialize["supportPhoneNumber"] = o.SupportPhoneNumber
	toSerialize["supportEmailAddress"] = o.SupportEmailAddress
	toSerialize["department"] = o.Department
	toSerialize["isDefaultPrestage"] = o.IsDefaultPrestage
	toSerialize["enrollmentSiteId"] = o.EnrollmentSiteId
	toSerialize["isKeepExistingSiteMembership"] = o.IsKeepExistingSiteMembership
	toSerialize["isKeepExistingLocationInformation"] = o.IsKeepExistingLocationInformation
	toSerialize["isRequireAuthentication"] = o.IsRequireAuthentication
	toSerialize["authenticationPrompt"] = o.AuthenticationPrompt
	toSerialize["isPreventActivationLock"] = o.IsPreventActivationLock
	toSerialize["isEnableDeviceBasedActivationLock"] = o.IsEnableDeviceBasedActivationLock
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
	toSerialize["isAllowPairing"] = o.IsAllowPairing
	toSerialize["isMultiUser"] = o.IsMultiUser
	toSerialize["isSupervised"] = o.IsSupervised
	toSerialize["maximumSharedAccounts"] = o.MaximumSharedAccounts
	toSerialize["isAutoAdvanceSetup"] = o.IsAutoAdvanceSetup
	toSerialize["isConfigureDeviceBeforeSetupAssistant"] = o.IsConfigureDeviceBeforeSetupAssistant
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Names) {
		toSerialize["names"] = o.Names
	}
	return toSerialize, nil
}

func (o *MobileDevicePrestage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"displayName",
		"isMandatory",
		"isMdmRemovable",
		"supportPhoneNumber",
		"supportEmailAddress",
		"department",
		"isDefaultPrestage",
		"enrollmentSiteId",
		"isKeepExistingSiteMembership",
		"isKeepExistingLocationInformation",
		"isRequireAuthentication",
		"authenticationPrompt",
		"isPreventActivationLock",
		"isEnableDeviceBasedActivationLock",
		"deviceEnrollmentProgramInstanceId",
		"locationInformation",
		"purchasingInformation",
		"isAllowPairing",
		"isMultiUser",
		"isSupervised",
		"maximumSharedAccounts",
		"isAutoAdvanceSetup",
		"isConfigureDeviceBeforeSetupAssistant",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMobileDevicePrestage := _MobileDevicePrestage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMobileDevicePrestage)

	if err != nil {
		return err
	}

	*o = MobileDevicePrestage(varMobileDevicePrestage)

	return err
}

type NullableMobileDevicePrestage struct {
	value *MobileDevicePrestage
	isSet bool
}

func (v NullableMobileDevicePrestage) Get() *MobileDevicePrestage {
	return v.value
}

func (v *NullableMobileDevicePrestage) Set(val *MobileDevicePrestage) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileDevicePrestage) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileDevicePrestage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileDevicePrestage(val *MobileDevicePrestage) *NullableMobileDevicePrestage {
	return &NullableMobileDevicePrestage{value: val, isSet: true}
}

func (v NullableMobileDevicePrestage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileDevicePrestage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


