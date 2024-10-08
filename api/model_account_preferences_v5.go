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

// checks if the AccountPreferencesV5 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountPreferencesV5{}

// AccountPreferencesV5 struct for AccountPreferencesV5
type AccountPreferencesV5 struct {
	// Language codes supported by Jamf Pro
	Language string `json:"language"`
	DateFormat string `json:"dateFormat"`
	Timezone string `json:"timezone"`
	ResultsPerPage int64 `json:"resultsPerPage"`
	UserInterfaceDisplayTheme AccountPreferencesUserInterfaceDisplayTheme `json:"userInterfaceDisplayTheme"`
	DisableRelativeDates bool `json:"disableRelativeDates"`
	DisablePageLeaveCheck bool `json:"disablePageLeaveCheck"`
	DisableTablePagination bool `json:"disableTablePagination"`
	DisableShortcutsTooltips bool `json:"disableShortcutsTooltips"`
	ConfigProfilesSortingMethod string `json:"configProfilesSortingMethod"`
	ComputerSearchMethod AccountPreferencesSearchType `json:"computerSearchMethod"`
	ComputerApplicationSearchMethod AccountPreferencesSearchType `json:"computerApplicationSearchMethod"`
	ComputerApplicationUsageSearchMethod AccountPreferencesSearchType `json:"computerApplicationUsageSearchMethod"`
	ComputerFontSearchMethod *AccountPreferencesSearchType `json:"computerFontSearchMethod,omitempty"`
	ComputerPluginSearchMethod *AccountPreferencesSearchType `json:"computerPluginSearchMethod,omitempty"`
	ComputerSoftwareUpdateSearchMethod *AccountPreferencesSearchType `json:"computerSoftwareUpdateSearchMethod,omitempty"`
	ComputerLocalUserAccountSearchMethod AccountPreferencesSearchType `json:"computerLocalUserAccountSearchMethod"`
	ComputerPackageReceiptSearchMethod AccountPreferencesSearchType `json:"computerPackageReceiptSearchMethod"`
	ComputerPrinterSearchMethod AccountPreferencesSearchType `json:"computerPrinterSearchMethod"`
	ComputerPeripheralSearchMethod *AccountPreferencesSearchType `json:"computerPeripheralSearchMethod,omitempty"`
	ComputerServiceSearchMethod AccountPreferencesSearchType `json:"computerServiceSearchMethod"`
	MobileDeviceSearchMethod AccountPreferencesSearchType `json:"mobileDeviceSearchMethod"`
	MobileDeviceAppSearchMethod AccountPreferencesSearchType `json:"mobileDeviceAppSearchMethod"`
	UserSearchMethod AccountPreferencesSearchType `json:"userSearchMethod"`
	UserAllContentSearchMethod AccountPreferencesSearchType `json:"userAllContentSearchMethod"`
	UserMobileDeviceAppSearchMethod AccountPreferencesSearchType `json:"userMobileDeviceAppSearchMethod"`
	UserMacAppStoreAppSearchMethod AccountPreferencesSearchType `json:"userMacAppStoreAppSearchMethod"`
	UserEbookSearchMethod AccountPreferencesSearchType `json:"userEbookSearchMethod"`
}

type _AccountPreferencesV5 AccountPreferencesV5

// NewAccountPreferencesV5 instantiates a new AccountPreferencesV5 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountPreferencesV5(language string, dateFormat string, timezone string, resultsPerPage int64, userInterfaceDisplayTheme AccountPreferencesUserInterfaceDisplayTheme, disableRelativeDates bool, disablePageLeaveCheck bool, disableTablePagination bool, disableShortcutsTooltips bool, configProfilesSortingMethod string, computerSearchMethod AccountPreferencesSearchType, computerApplicationSearchMethod AccountPreferencesSearchType, computerApplicationUsageSearchMethod AccountPreferencesSearchType, computerLocalUserAccountSearchMethod AccountPreferencesSearchType, computerPackageReceiptSearchMethod AccountPreferencesSearchType, computerPrinterSearchMethod AccountPreferencesSearchType, computerServiceSearchMethod AccountPreferencesSearchType, mobileDeviceSearchMethod AccountPreferencesSearchType, mobileDeviceAppSearchMethod AccountPreferencesSearchType, userSearchMethod AccountPreferencesSearchType, userAllContentSearchMethod AccountPreferencesSearchType, userMobileDeviceAppSearchMethod AccountPreferencesSearchType, userMacAppStoreAppSearchMethod AccountPreferencesSearchType, userEbookSearchMethod AccountPreferencesSearchType) *AccountPreferencesV5 {
	this := AccountPreferencesV5{}
	this.Language = language
	this.DateFormat = dateFormat
	this.Timezone = timezone
	this.ResultsPerPage = resultsPerPage
	this.UserInterfaceDisplayTheme = userInterfaceDisplayTheme
	this.DisableRelativeDates = disableRelativeDates
	this.DisablePageLeaveCheck = disablePageLeaveCheck
	this.DisableTablePagination = disableTablePagination
	this.DisableShortcutsTooltips = disableShortcutsTooltips
	this.ConfigProfilesSortingMethod = configProfilesSortingMethod
	this.ComputerSearchMethod = computerSearchMethod
	this.ComputerApplicationSearchMethod = computerApplicationSearchMethod
	this.ComputerApplicationUsageSearchMethod = computerApplicationUsageSearchMethod
	this.ComputerLocalUserAccountSearchMethod = computerLocalUserAccountSearchMethod
	this.ComputerPackageReceiptSearchMethod = computerPackageReceiptSearchMethod
	this.ComputerPrinterSearchMethod = computerPrinterSearchMethod
	this.ComputerServiceSearchMethod = computerServiceSearchMethod
	this.MobileDeviceSearchMethod = mobileDeviceSearchMethod
	this.MobileDeviceAppSearchMethod = mobileDeviceAppSearchMethod
	this.UserSearchMethod = userSearchMethod
	this.UserAllContentSearchMethod = userAllContentSearchMethod
	this.UserMobileDeviceAppSearchMethod = userMobileDeviceAppSearchMethod
	this.UserMacAppStoreAppSearchMethod = userMacAppStoreAppSearchMethod
	this.UserEbookSearchMethod = userEbookSearchMethod
	return &this
}

// NewAccountPreferencesV5WithDefaults instantiates a new AccountPreferencesV5 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountPreferencesV5WithDefaults() *AccountPreferencesV5 {
	this := AccountPreferencesV5{}
	var language string = "en"
	this.Language = language
	var resultsPerPage int64 = 100
	this.ResultsPerPage = resultsPerPage
	return &this
}

// GetLanguage returns the Language field value
func (o *AccountPreferencesV5) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *AccountPreferencesV5) SetLanguage(v string) {
	o.Language = v
}

// GetDateFormat returns the DateFormat field value
func (o *AccountPreferencesV5) GetDateFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateFormat
}

// GetDateFormatOk returns a tuple with the DateFormat field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetDateFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateFormat, true
}

// SetDateFormat sets field value
func (o *AccountPreferencesV5) SetDateFormat(v string) {
	o.DateFormat = v
}

// GetTimezone returns the Timezone field value
func (o *AccountPreferencesV5) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *AccountPreferencesV5) SetTimezone(v string) {
	o.Timezone = v
}

// GetResultsPerPage returns the ResultsPerPage field value
func (o *AccountPreferencesV5) GetResultsPerPage() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ResultsPerPage
}

// GetResultsPerPageOk returns a tuple with the ResultsPerPage field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetResultsPerPageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultsPerPage, true
}

// SetResultsPerPage sets field value
func (o *AccountPreferencesV5) SetResultsPerPage(v int64) {
	o.ResultsPerPage = v
}

// GetUserInterfaceDisplayTheme returns the UserInterfaceDisplayTheme field value
func (o *AccountPreferencesV5) GetUserInterfaceDisplayTheme() AccountPreferencesUserInterfaceDisplayTheme {
	if o == nil {
		var ret AccountPreferencesUserInterfaceDisplayTheme
		return ret
	}

	return o.UserInterfaceDisplayTheme
}

// GetUserInterfaceDisplayThemeOk returns a tuple with the UserInterfaceDisplayTheme field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetUserInterfaceDisplayThemeOk() (*AccountPreferencesUserInterfaceDisplayTheme, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserInterfaceDisplayTheme, true
}

// SetUserInterfaceDisplayTheme sets field value
func (o *AccountPreferencesV5) SetUserInterfaceDisplayTheme(v AccountPreferencesUserInterfaceDisplayTheme) {
	o.UserInterfaceDisplayTheme = v
}

// GetDisableRelativeDates returns the DisableRelativeDates field value
func (o *AccountPreferencesV5) GetDisableRelativeDates() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableRelativeDates
}

// GetDisableRelativeDatesOk returns a tuple with the DisableRelativeDates field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetDisableRelativeDatesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableRelativeDates, true
}

// SetDisableRelativeDates sets field value
func (o *AccountPreferencesV5) SetDisableRelativeDates(v bool) {
	o.DisableRelativeDates = v
}

// GetDisablePageLeaveCheck returns the DisablePageLeaveCheck field value
func (o *AccountPreferencesV5) GetDisablePageLeaveCheck() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisablePageLeaveCheck
}

// GetDisablePageLeaveCheckOk returns a tuple with the DisablePageLeaveCheck field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetDisablePageLeaveCheckOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisablePageLeaveCheck, true
}

// SetDisablePageLeaveCheck sets field value
func (o *AccountPreferencesV5) SetDisablePageLeaveCheck(v bool) {
	o.DisablePageLeaveCheck = v
}

// GetDisableTablePagination returns the DisableTablePagination field value
func (o *AccountPreferencesV5) GetDisableTablePagination() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableTablePagination
}

// GetDisableTablePaginationOk returns a tuple with the DisableTablePagination field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetDisableTablePaginationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableTablePagination, true
}

// SetDisableTablePagination sets field value
func (o *AccountPreferencesV5) SetDisableTablePagination(v bool) {
	o.DisableTablePagination = v
}

// GetDisableShortcutsTooltips returns the DisableShortcutsTooltips field value
func (o *AccountPreferencesV5) GetDisableShortcutsTooltips() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableShortcutsTooltips
}

// GetDisableShortcutsTooltipsOk returns a tuple with the DisableShortcutsTooltips field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetDisableShortcutsTooltipsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableShortcutsTooltips, true
}

// SetDisableShortcutsTooltips sets field value
func (o *AccountPreferencesV5) SetDisableShortcutsTooltips(v bool) {
	o.DisableShortcutsTooltips = v
}

// GetConfigProfilesSortingMethod returns the ConfigProfilesSortingMethod field value
func (o *AccountPreferencesV5) GetConfigProfilesSortingMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigProfilesSortingMethod
}

// GetConfigProfilesSortingMethodOk returns a tuple with the ConfigProfilesSortingMethod field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetConfigProfilesSortingMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigProfilesSortingMethod, true
}

// SetConfigProfilesSortingMethod sets field value
func (o *AccountPreferencesV5) SetConfigProfilesSortingMethod(v string) {
	o.ConfigProfilesSortingMethod = v
}

// GetComputerSearchMethod returns the ComputerSearchMethod field value
func (o *AccountPreferencesV5) GetComputerSearchMethod() AccountPreferencesSearchType {
	if o == nil {
		var ret AccountPreferencesSearchType
		return ret
	}

	return o.ComputerSearchMethod
}

// GetComputerSearchMethodOk returns a tuple with the ComputerSearchMethod field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetComputerSearchMethodOk() (*AccountPreferencesSearchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComputerSearchMethod, true
}

// SetComputerSearchMethod sets field value
func (o *AccountPreferencesV5) SetComputerSearchMethod(v AccountPreferencesSearchType) {
	o.ComputerSearchMethod = v
}

// GetComputerApplicationSearchMethod returns the ComputerApplicationSearchMethod field value
func (o *AccountPreferencesV5) GetComputerApplicationSearchMethod() AccountPreferencesSearchType {
	if o == nil {
		var ret AccountPreferencesSearchType
		return ret
	}

	return o.ComputerApplicationSearchMethod
}

// GetComputerApplicationSearchMethodOk returns a tuple with the ComputerApplicationSearchMethod field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetComputerApplicationSearchMethodOk() (*AccountPreferencesSearchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComputerApplicationSearchMethod, true
}

// SetComputerApplicationSearchMethod sets field value
func (o *AccountPreferencesV5) SetComputerApplicationSearchMethod(v AccountPreferencesSearchType) {
	o.ComputerApplicationSearchMethod = v
}

// GetComputerApplicationUsageSearchMethod returns the ComputerApplicationUsageSearchMethod field value
func (o *AccountPreferencesV5) GetComputerApplicationUsageSearchMethod() AccountPreferencesSearchType {
	if o == nil {
		var ret AccountPreferencesSearchType
		return ret
	}

	return o.ComputerApplicationUsageSearchMethod
}

// GetComputerApplicationUsageSearchMethodOk returns a tuple with the ComputerApplicationUsageSearchMethod field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetComputerApplicationUsageSearchMethodOk() (*AccountPreferencesSearchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComputerApplicationUsageSearchMethod, true
}

// SetComputerApplicationUsageSearchMethod sets field value
func (o *AccountPreferencesV5) SetComputerApplicationUsageSearchMethod(v AccountPreferencesSearchType) {
	o.ComputerApplicationUsageSearchMethod = v
}

// GetComputerFontSearchMethod returns the ComputerFontSearchMethod field value if set, zero value otherwise.
func (o *AccountPreferencesV5) GetComputerFontSearchMethod() AccountPreferencesSearchType {
	if o == nil || IsNil(o.ComputerFontSearchMethod) {
		var ret AccountPreferencesSearchType
		return ret
	}
	return *o.ComputerFontSearchMethod
}

// GetComputerFontSearchMethodOk returns a tuple with the ComputerFontSearchMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetComputerFontSearchMethodOk() (*AccountPreferencesSearchType, bool) {
	if o == nil || IsNil(o.ComputerFontSearchMethod) {
		return nil, false
	}
	return o.ComputerFontSearchMethod, true
}

// HasComputerFontSearchMethod returns a boolean if a field has been set.
func (o *AccountPreferencesV5) HasComputerFontSearchMethod() bool {
	if o != nil && !IsNil(o.ComputerFontSearchMethod) {
		return true
	}

	return false
}

// SetComputerFontSearchMethod gets a reference to the given AccountPreferencesSearchType and assigns it to the ComputerFontSearchMethod field.
func (o *AccountPreferencesV5) SetComputerFontSearchMethod(v AccountPreferencesSearchType) {
	o.ComputerFontSearchMethod = &v
}

// GetComputerPluginSearchMethod returns the ComputerPluginSearchMethod field value if set, zero value otherwise.
func (o *AccountPreferencesV5) GetComputerPluginSearchMethod() AccountPreferencesSearchType {
	if o == nil || IsNil(o.ComputerPluginSearchMethod) {
		var ret AccountPreferencesSearchType
		return ret
	}
	return *o.ComputerPluginSearchMethod
}

// GetComputerPluginSearchMethodOk returns a tuple with the ComputerPluginSearchMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetComputerPluginSearchMethodOk() (*AccountPreferencesSearchType, bool) {
	if o == nil || IsNil(o.ComputerPluginSearchMethod) {
		return nil, false
	}
	return o.ComputerPluginSearchMethod, true
}

// HasComputerPluginSearchMethod returns a boolean if a field has been set.
func (o *AccountPreferencesV5) HasComputerPluginSearchMethod() bool {
	if o != nil && !IsNil(o.ComputerPluginSearchMethod) {
		return true
	}

	return false
}

// SetComputerPluginSearchMethod gets a reference to the given AccountPreferencesSearchType and assigns it to the ComputerPluginSearchMethod field.
func (o *AccountPreferencesV5) SetComputerPluginSearchMethod(v AccountPreferencesSearchType) {
	o.ComputerPluginSearchMethod = &v
}

// GetComputerSoftwareUpdateSearchMethod returns the ComputerSoftwareUpdateSearchMethod field value if set, zero value otherwise.
func (o *AccountPreferencesV5) GetComputerSoftwareUpdateSearchMethod() AccountPreferencesSearchType {
	if o == nil || IsNil(o.ComputerSoftwareUpdateSearchMethod) {
		var ret AccountPreferencesSearchType
		return ret
	}
	return *o.ComputerSoftwareUpdateSearchMethod
}

// GetComputerSoftwareUpdateSearchMethodOk returns a tuple with the ComputerSoftwareUpdateSearchMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetComputerSoftwareUpdateSearchMethodOk() (*AccountPreferencesSearchType, bool) {
	if o == nil || IsNil(o.ComputerSoftwareUpdateSearchMethod) {
		return nil, false
	}
	return o.ComputerSoftwareUpdateSearchMethod, true
}

// HasComputerSoftwareUpdateSearchMethod returns a boolean if a field has been set.
func (o *AccountPreferencesV5) HasComputerSoftwareUpdateSearchMethod() bool {
	if o != nil && !IsNil(o.ComputerSoftwareUpdateSearchMethod) {
		return true
	}

	return false
}

// SetComputerSoftwareUpdateSearchMethod gets a reference to the given AccountPreferencesSearchType and assigns it to the ComputerSoftwareUpdateSearchMethod field.
func (o *AccountPreferencesV5) SetComputerSoftwareUpdateSearchMethod(v AccountPreferencesSearchType) {
	o.ComputerSoftwareUpdateSearchMethod = &v
}

// GetComputerLocalUserAccountSearchMethod returns the ComputerLocalUserAccountSearchMethod field value
func (o *AccountPreferencesV5) GetComputerLocalUserAccountSearchMethod() AccountPreferencesSearchType {
	if o == nil {
		var ret AccountPreferencesSearchType
		return ret
	}

	return o.ComputerLocalUserAccountSearchMethod
}

// GetComputerLocalUserAccountSearchMethodOk returns a tuple with the ComputerLocalUserAccountSearchMethod field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetComputerLocalUserAccountSearchMethodOk() (*AccountPreferencesSearchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComputerLocalUserAccountSearchMethod, true
}

// SetComputerLocalUserAccountSearchMethod sets field value
func (o *AccountPreferencesV5) SetComputerLocalUserAccountSearchMethod(v AccountPreferencesSearchType) {
	o.ComputerLocalUserAccountSearchMethod = v
}

// GetComputerPackageReceiptSearchMethod returns the ComputerPackageReceiptSearchMethod field value
func (o *AccountPreferencesV5) GetComputerPackageReceiptSearchMethod() AccountPreferencesSearchType {
	if o == nil {
		var ret AccountPreferencesSearchType
		return ret
	}

	return o.ComputerPackageReceiptSearchMethod
}

// GetComputerPackageReceiptSearchMethodOk returns a tuple with the ComputerPackageReceiptSearchMethod field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetComputerPackageReceiptSearchMethodOk() (*AccountPreferencesSearchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComputerPackageReceiptSearchMethod, true
}

// SetComputerPackageReceiptSearchMethod sets field value
func (o *AccountPreferencesV5) SetComputerPackageReceiptSearchMethod(v AccountPreferencesSearchType) {
	o.ComputerPackageReceiptSearchMethod = v
}

// GetComputerPrinterSearchMethod returns the ComputerPrinterSearchMethod field value
func (o *AccountPreferencesV5) GetComputerPrinterSearchMethod() AccountPreferencesSearchType {
	if o == nil {
		var ret AccountPreferencesSearchType
		return ret
	}

	return o.ComputerPrinterSearchMethod
}

// GetComputerPrinterSearchMethodOk returns a tuple with the ComputerPrinterSearchMethod field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetComputerPrinterSearchMethodOk() (*AccountPreferencesSearchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComputerPrinterSearchMethod, true
}

// SetComputerPrinterSearchMethod sets field value
func (o *AccountPreferencesV5) SetComputerPrinterSearchMethod(v AccountPreferencesSearchType) {
	o.ComputerPrinterSearchMethod = v
}

// GetComputerPeripheralSearchMethod returns the ComputerPeripheralSearchMethod field value if set, zero value otherwise.
func (o *AccountPreferencesV5) GetComputerPeripheralSearchMethod() AccountPreferencesSearchType {
	if o == nil || IsNil(o.ComputerPeripheralSearchMethod) {
		var ret AccountPreferencesSearchType
		return ret
	}
	return *o.ComputerPeripheralSearchMethod
}

// GetComputerPeripheralSearchMethodOk returns a tuple with the ComputerPeripheralSearchMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetComputerPeripheralSearchMethodOk() (*AccountPreferencesSearchType, bool) {
	if o == nil || IsNil(o.ComputerPeripheralSearchMethod) {
		return nil, false
	}
	return o.ComputerPeripheralSearchMethod, true
}

// HasComputerPeripheralSearchMethod returns a boolean if a field has been set.
func (o *AccountPreferencesV5) HasComputerPeripheralSearchMethod() bool {
	if o != nil && !IsNil(o.ComputerPeripheralSearchMethod) {
		return true
	}

	return false
}

// SetComputerPeripheralSearchMethod gets a reference to the given AccountPreferencesSearchType and assigns it to the ComputerPeripheralSearchMethod field.
func (o *AccountPreferencesV5) SetComputerPeripheralSearchMethod(v AccountPreferencesSearchType) {
	o.ComputerPeripheralSearchMethod = &v
}

// GetComputerServiceSearchMethod returns the ComputerServiceSearchMethod field value
func (o *AccountPreferencesV5) GetComputerServiceSearchMethod() AccountPreferencesSearchType {
	if o == nil {
		var ret AccountPreferencesSearchType
		return ret
	}

	return o.ComputerServiceSearchMethod
}

// GetComputerServiceSearchMethodOk returns a tuple with the ComputerServiceSearchMethod field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetComputerServiceSearchMethodOk() (*AccountPreferencesSearchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComputerServiceSearchMethod, true
}

// SetComputerServiceSearchMethod sets field value
func (o *AccountPreferencesV5) SetComputerServiceSearchMethod(v AccountPreferencesSearchType) {
	o.ComputerServiceSearchMethod = v
}

// GetMobileDeviceSearchMethod returns the MobileDeviceSearchMethod field value
func (o *AccountPreferencesV5) GetMobileDeviceSearchMethod() AccountPreferencesSearchType {
	if o == nil {
		var ret AccountPreferencesSearchType
		return ret
	}

	return o.MobileDeviceSearchMethod
}

// GetMobileDeviceSearchMethodOk returns a tuple with the MobileDeviceSearchMethod field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetMobileDeviceSearchMethodOk() (*AccountPreferencesSearchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MobileDeviceSearchMethod, true
}

// SetMobileDeviceSearchMethod sets field value
func (o *AccountPreferencesV5) SetMobileDeviceSearchMethod(v AccountPreferencesSearchType) {
	o.MobileDeviceSearchMethod = v
}

// GetMobileDeviceAppSearchMethod returns the MobileDeviceAppSearchMethod field value
func (o *AccountPreferencesV5) GetMobileDeviceAppSearchMethod() AccountPreferencesSearchType {
	if o == nil {
		var ret AccountPreferencesSearchType
		return ret
	}

	return o.MobileDeviceAppSearchMethod
}

// GetMobileDeviceAppSearchMethodOk returns a tuple with the MobileDeviceAppSearchMethod field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetMobileDeviceAppSearchMethodOk() (*AccountPreferencesSearchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MobileDeviceAppSearchMethod, true
}

// SetMobileDeviceAppSearchMethod sets field value
func (o *AccountPreferencesV5) SetMobileDeviceAppSearchMethod(v AccountPreferencesSearchType) {
	o.MobileDeviceAppSearchMethod = v
}

// GetUserSearchMethod returns the UserSearchMethod field value
func (o *AccountPreferencesV5) GetUserSearchMethod() AccountPreferencesSearchType {
	if o == nil {
		var ret AccountPreferencesSearchType
		return ret
	}

	return o.UserSearchMethod
}

// GetUserSearchMethodOk returns a tuple with the UserSearchMethod field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetUserSearchMethodOk() (*AccountPreferencesSearchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserSearchMethod, true
}

// SetUserSearchMethod sets field value
func (o *AccountPreferencesV5) SetUserSearchMethod(v AccountPreferencesSearchType) {
	o.UserSearchMethod = v
}

// GetUserAllContentSearchMethod returns the UserAllContentSearchMethod field value
func (o *AccountPreferencesV5) GetUserAllContentSearchMethod() AccountPreferencesSearchType {
	if o == nil {
		var ret AccountPreferencesSearchType
		return ret
	}

	return o.UserAllContentSearchMethod
}

// GetUserAllContentSearchMethodOk returns a tuple with the UserAllContentSearchMethod field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetUserAllContentSearchMethodOk() (*AccountPreferencesSearchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAllContentSearchMethod, true
}

// SetUserAllContentSearchMethod sets field value
func (o *AccountPreferencesV5) SetUserAllContentSearchMethod(v AccountPreferencesSearchType) {
	o.UserAllContentSearchMethod = v
}

// GetUserMobileDeviceAppSearchMethod returns the UserMobileDeviceAppSearchMethod field value
func (o *AccountPreferencesV5) GetUserMobileDeviceAppSearchMethod() AccountPreferencesSearchType {
	if o == nil {
		var ret AccountPreferencesSearchType
		return ret
	}

	return o.UserMobileDeviceAppSearchMethod
}

// GetUserMobileDeviceAppSearchMethodOk returns a tuple with the UserMobileDeviceAppSearchMethod field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetUserMobileDeviceAppSearchMethodOk() (*AccountPreferencesSearchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserMobileDeviceAppSearchMethod, true
}

// SetUserMobileDeviceAppSearchMethod sets field value
func (o *AccountPreferencesV5) SetUserMobileDeviceAppSearchMethod(v AccountPreferencesSearchType) {
	o.UserMobileDeviceAppSearchMethod = v
}

// GetUserMacAppStoreAppSearchMethod returns the UserMacAppStoreAppSearchMethod field value
func (o *AccountPreferencesV5) GetUserMacAppStoreAppSearchMethod() AccountPreferencesSearchType {
	if o == nil {
		var ret AccountPreferencesSearchType
		return ret
	}

	return o.UserMacAppStoreAppSearchMethod
}

// GetUserMacAppStoreAppSearchMethodOk returns a tuple with the UserMacAppStoreAppSearchMethod field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetUserMacAppStoreAppSearchMethodOk() (*AccountPreferencesSearchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserMacAppStoreAppSearchMethod, true
}

// SetUserMacAppStoreAppSearchMethod sets field value
func (o *AccountPreferencesV5) SetUserMacAppStoreAppSearchMethod(v AccountPreferencesSearchType) {
	o.UserMacAppStoreAppSearchMethod = v
}

// GetUserEbookSearchMethod returns the UserEbookSearchMethod field value
func (o *AccountPreferencesV5) GetUserEbookSearchMethod() AccountPreferencesSearchType {
	if o == nil {
		var ret AccountPreferencesSearchType
		return ret
	}

	return o.UserEbookSearchMethod
}

// GetUserEbookSearchMethodOk returns a tuple with the UserEbookSearchMethod field value
// and a boolean to check if the value has been set.
func (o *AccountPreferencesV5) GetUserEbookSearchMethodOk() (*AccountPreferencesSearchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserEbookSearchMethod, true
}

// SetUserEbookSearchMethod sets field value
func (o *AccountPreferencesV5) SetUserEbookSearchMethod(v AccountPreferencesSearchType) {
	o.UserEbookSearchMethod = v
}

func (o AccountPreferencesV5) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountPreferencesV5) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["language"] = o.Language
	toSerialize["dateFormat"] = o.DateFormat
	toSerialize["timezone"] = o.Timezone
	toSerialize["resultsPerPage"] = o.ResultsPerPage
	toSerialize["userInterfaceDisplayTheme"] = o.UserInterfaceDisplayTheme
	toSerialize["disableRelativeDates"] = o.DisableRelativeDates
	toSerialize["disablePageLeaveCheck"] = o.DisablePageLeaveCheck
	toSerialize["disableTablePagination"] = o.DisableTablePagination
	toSerialize["disableShortcutsTooltips"] = o.DisableShortcutsTooltips
	toSerialize["configProfilesSortingMethod"] = o.ConfigProfilesSortingMethod
	toSerialize["computerSearchMethod"] = o.ComputerSearchMethod
	toSerialize["computerApplicationSearchMethod"] = o.ComputerApplicationSearchMethod
	toSerialize["computerApplicationUsageSearchMethod"] = o.ComputerApplicationUsageSearchMethod
	if !IsNil(o.ComputerFontSearchMethod) {
		toSerialize["computerFontSearchMethod"] = o.ComputerFontSearchMethod
	}
	if !IsNil(o.ComputerPluginSearchMethod) {
		toSerialize["computerPluginSearchMethod"] = o.ComputerPluginSearchMethod
	}
	if !IsNil(o.ComputerSoftwareUpdateSearchMethod) {
		toSerialize["computerSoftwareUpdateSearchMethod"] = o.ComputerSoftwareUpdateSearchMethod
	}
	toSerialize["computerLocalUserAccountSearchMethod"] = o.ComputerLocalUserAccountSearchMethod
	toSerialize["computerPackageReceiptSearchMethod"] = o.ComputerPackageReceiptSearchMethod
	toSerialize["computerPrinterSearchMethod"] = o.ComputerPrinterSearchMethod
	if !IsNil(o.ComputerPeripheralSearchMethod) {
		toSerialize["computerPeripheralSearchMethod"] = o.ComputerPeripheralSearchMethod
	}
	toSerialize["computerServiceSearchMethod"] = o.ComputerServiceSearchMethod
	toSerialize["mobileDeviceSearchMethod"] = o.MobileDeviceSearchMethod
	toSerialize["mobileDeviceAppSearchMethod"] = o.MobileDeviceAppSearchMethod
	toSerialize["userSearchMethod"] = o.UserSearchMethod
	toSerialize["userAllContentSearchMethod"] = o.UserAllContentSearchMethod
	toSerialize["userMobileDeviceAppSearchMethod"] = o.UserMobileDeviceAppSearchMethod
	toSerialize["userMacAppStoreAppSearchMethod"] = o.UserMacAppStoreAppSearchMethod
	toSerialize["userEbookSearchMethod"] = o.UserEbookSearchMethod
	return toSerialize, nil
}

func (o *AccountPreferencesV5) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"language",
		"dateFormat",
		"timezone",
		"resultsPerPage",
		"userInterfaceDisplayTheme",
		"disableRelativeDates",
		"disablePageLeaveCheck",
		"disableTablePagination",
		"disableShortcutsTooltips",
		"configProfilesSortingMethod",
		"computerSearchMethod",
		"computerApplicationSearchMethod",
		"computerApplicationUsageSearchMethod",
		"computerLocalUserAccountSearchMethod",
		"computerPackageReceiptSearchMethod",
		"computerPrinterSearchMethod",
		"computerServiceSearchMethod",
		"mobileDeviceSearchMethod",
		"mobileDeviceAppSearchMethod",
		"userSearchMethod",
		"userAllContentSearchMethod",
		"userMobileDeviceAppSearchMethod",
		"userMacAppStoreAppSearchMethod",
		"userEbookSearchMethod",
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

	varAccountPreferencesV5 := _AccountPreferencesV5{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountPreferencesV5)

	if err != nil {
		return err
	}

	*o = AccountPreferencesV5(varAccountPreferencesV5)

	return err
}

type NullableAccountPreferencesV5 struct {
	value *AccountPreferencesV5
	isSet bool
}

func (v NullableAccountPreferencesV5) Get() *AccountPreferencesV5 {
	return v.value
}

func (v *NullableAccountPreferencesV5) Set(val *AccountPreferencesV5) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountPreferencesV5) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountPreferencesV5) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountPreferencesV5(val *AccountPreferencesV5) *NullableAccountPreferencesV5 {
	return &NullableAccountPreferencesV5{value: val, isSet: true}
}

func (v NullableAccountPreferencesV5) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountPreferencesV5) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


