# AccountPreferencesV5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | **string** | Language codes supported by Jamf Pro | [default to "en"]
**DateFormat** | **string** |  | 
**Timezone** | **string** |  | 
**ResultsPerPage** | **int64** |  | [default to 100]
**UserInterfaceDisplayTheme** | [**AccountPreferencesUserInterfaceDisplayTheme**](AccountPreferencesUserInterfaceDisplayTheme.md) |  | 
**DisableRelativeDates** | **bool** |  | 
**DisablePageLeaveCheck** | **bool** |  | 
**DisableTablePagination** | **bool** |  | 
**DisableShortcutsTooltips** | **bool** |  | 
**ConfigProfilesSortingMethod** | **string** |  | 
**ComputerSearchMethod** | [**AccountPreferencesSearchType**](AccountPreferencesSearchType.md) |  | 
**ComputerApplicationSearchMethod** | [**AccountPreferencesSearchType**](AccountPreferencesSearchType.md) |  | 
**ComputerApplicationUsageSearchMethod** | [**AccountPreferencesSearchType**](AccountPreferencesSearchType.md) |  | 
**ComputerFontSearchMethod** | Pointer to [**AccountPreferencesSearchType**](AccountPreferencesSearchType.md) |  | [optional] 
**ComputerPluginSearchMethod** | Pointer to [**AccountPreferencesSearchType**](AccountPreferencesSearchType.md) |  | [optional] 
**ComputerSoftwareUpdateSearchMethod** | Pointer to [**AccountPreferencesSearchType**](AccountPreferencesSearchType.md) |  | [optional] 
**ComputerLocalUserAccountSearchMethod** | [**AccountPreferencesSearchType**](AccountPreferencesSearchType.md) |  | 
**ComputerPackageReceiptSearchMethod** | [**AccountPreferencesSearchType**](AccountPreferencesSearchType.md) |  | 
**ComputerPrinterSearchMethod** | [**AccountPreferencesSearchType**](AccountPreferencesSearchType.md) |  | 
**ComputerPeripheralSearchMethod** | Pointer to [**AccountPreferencesSearchType**](AccountPreferencesSearchType.md) |  | [optional] 
**ComputerServiceSearchMethod** | [**AccountPreferencesSearchType**](AccountPreferencesSearchType.md) |  | 
**MobileDeviceSearchMethod** | [**AccountPreferencesSearchType**](AccountPreferencesSearchType.md) |  | 
**MobileDeviceAppSearchMethod** | [**AccountPreferencesSearchType**](AccountPreferencesSearchType.md) |  | 
**UserSearchMethod** | [**AccountPreferencesSearchType**](AccountPreferencesSearchType.md) |  | 
**UserAllContentSearchMethod** | [**AccountPreferencesSearchType**](AccountPreferencesSearchType.md) |  | 
**UserMobileDeviceAppSearchMethod** | [**AccountPreferencesSearchType**](AccountPreferencesSearchType.md) |  | 
**UserMacAppStoreAppSearchMethod** | [**AccountPreferencesSearchType**](AccountPreferencesSearchType.md) |  | 
**UserEbookSearchMethod** | [**AccountPreferencesSearchType**](AccountPreferencesSearchType.md) |  | 

## Methods

### NewAccountPreferencesV5

`func NewAccountPreferencesV5(language string, dateFormat string, timezone string, resultsPerPage int64, userInterfaceDisplayTheme AccountPreferencesUserInterfaceDisplayTheme, disableRelativeDates bool, disablePageLeaveCheck bool, disableTablePagination bool, disableShortcutsTooltips bool, configProfilesSortingMethod string, computerSearchMethod AccountPreferencesSearchType, computerApplicationSearchMethod AccountPreferencesSearchType, computerApplicationUsageSearchMethod AccountPreferencesSearchType, computerLocalUserAccountSearchMethod AccountPreferencesSearchType, computerPackageReceiptSearchMethod AccountPreferencesSearchType, computerPrinterSearchMethod AccountPreferencesSearchType, computerServiceSearchMethod AccountPreferencesSearchType, mobileDeviceSearchMethod AccountPreferencesSearchType, mobileDeviceAppSearchMethod AccountPreferencesSearchType, userSearchMethod AccountPreferencesSearchType, userAllContentSearchMethod AccountPreferencesSearchType, userMobileDeviceAppSearchMethod AccountPreferencesSearchType, userMacAppStoreAppSearchMethod AccountPreferencesSearchType, userEbookSearchMethod AccountPreferencesSearchType, ) *AccountPreferencesV5`

NewAccountPreferencesV5 instantiates a new AccountPreferencesV5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountPreferencesV5WithDefaults

`func NewAccountPreferencesV5WithDefaults() *AccountPreferencesV5`

NewAccountPreferencesV5WithDefaults instantiates a new AccountPreferencesV5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *AccountPreferencesV5) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AccountPreferencesV5) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AccountPreferencesV5) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetDateFormat

`func (o *AccountPreferencesV5) GetDateFormat() string`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *AccountPreferencesV5) GetDateFormatOk() (*string, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *AccountPreferencesV5) SetDateFormat(v string)`

SetDateFormat sets DateFormat field to given value.


### GetTimezone

`func (o *AccountPreferencesV5) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AccountPreferencesV5) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AccountPreferencesV5) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetResultsPerPage

`func (o *AccountPreferencesV5) GetResultsPerPage() int64`

GetResultsPerPage returns the ResultsPerPage field if non-nil, zero value otherwise.

### GetResultsPerPageOk

`func (o *AccountPreferencesV5) GetResultsPerPageOk() (*int64, bool)`

GetResultsPerPageOk returns a tuple with the ResultsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerPage

`func (o *AccountPreferencesV5) SetResultsPerPage(v int64)`

SetResultsPerPage sets ResultsPerPage field to given value.


### GetUserInterfaceDisplayTheme

`func (o *AccountPreferencesV5) GetUserInterfaceDisplayTheme() AccountPreferencesUserInterfaceDisplayTheme`

GetUserInterfaceDisplayTheme returns the UserInterfaceDisplayTheme field if non-nil, zero value otherwise.

### GetUserInterfaceDisplayThemeOk

`func (o *AccountPreferencesV5) GetUserInterfaceDisplayThemeOk() (*AccountPreferencesUserInterfaceDisplayTheme, bool)`

GetUserInterfaceDisplayThemeOk returns a tuple with the UserInterfaceDisplayTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInterfaceDisplayTheme

`func (o *AccountPreferencesV5) SetUserInterfaceDisplayTheme(v AccountPreferencesUserInterfaceDisplayTheme)`

SetUserInterfaceDisplayTheme sets UserInterfaceDisplayTheme field to given value.


### GetDisableRelativeDates

`func (o *AccountPreferencesV5) GetDisableRelativeDates() bool`

GetDisableRelativeDates returns the DisableRelativeDates field if non-nil, zero value otherwise.

### GetDisableRelativeDatesOk

`func (o *AccountPreferencesV5) GetDisableRelativeDatesOk() (*bool, bool)`

GetDisableRelativeDatesOk returns a tuple with the DisableRelativeDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRelativeDates

`func (o *AccountPreferencesV5) SetDisableRelativeDates(v bool)`

SetDisableRelativeDates sets DisableRelativeDates field to given value.


### GetDisablePageLeaveCheck

`func (o *AccountPreferencesV5) GetDisablePageLeaveCheck() bool`

GetDisablePageLeaveCheck returns the DisablePageLeaveCheck field if non-nil, zero value otherwise.

### GetDisablePageLeaveCheckOk

`func (o *AccountPreferencesV5) GetDisablePageLeaveCheckOk() (*bool, bool)`

GetDisablePageLeaveCheckOk returns a tuple with the DisablePageLeaveCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisablePageLeaveCheck

`func (o *AccountPreferencesV5) SetDisablePageLeaveCheck(v bool)`

SetDisablePageLeaveCheck sets DisablePageLeaveCheck field to given value.


### GetDisableTablePagination

`func (o *AccountPreferencesV5) GetDisableTablePagination() bool`

GetDisableTablePagination returns the DisableTablePagination field if non-nil, zero value otherwise.

### GetDisableTablePaginationOk

`func (o *AccountPreferencesV5) GetDisableTablePaginationOk() (*bool, bool)`

GetDisableTablePaginationOk returns a tuple with the DisableTablePagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableTablePagination

`func (o *AccountPreferencesV5) SetDisableTablePagination(v bool)`

SetDisableTablePagination sets DisableTablePagination field to given value.


### GetDisableShortcutsTooltips

`func (o *AccountPreferencesV5) GetDisableShortcutsTooltips() bool`

GetDisableShortcutsTooltips returns the DisableShortcutsTooltips field if non-nil, zero value otherwise.

### GetDisableShortcutsTooltipsOk

`func (o *AccountPreferencesV5) GetDisableShortcutsTooltipsOk() (*bool, bool)`

GetDisableShortcutsTooltipsOk returns a tuple with the DisableShortcutsTooltips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableShortcutsTooltips

`func (o *AccountPreferencesV5) SetDisableShortcutsTooltips(v bool)`

SetDisableShortcutsTooltips sets DisableShortcutsTooltips field to given value.


### GetConfigProfilesSortingMethod

`func (o *AccountPreferencesV5) GetConfigProfilesSortingMethod() string`

GetConfigProfilesSortingMethod returns the ConfigProfilesSortingMethod field if non-nil, zero value otherwise.

### GetConfigProfilesSortingMethodOk

`func (o *AccountPreferencesV5) GetConfigProfilesSortingMethodOk() (*string, bool)`

GetConfigProfilesSortingMethodOk returns a tuple with the ConfigProfilesSortingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigProfilesSortingMethod

`func (o *AccountPreferencesV5) SetConfigProfilesSortingMethod(v string)`

SetConfigProfilesSortingMethod sets ConfigProfilesSortingMethod field to given value.


### GetComputerSearchMethod

`func (o *AccountPreferencesV5) GetComputerSearchMethod() AccountPreferencesSearchType`

GetComputerSearchMethod returns the ComputerSearchMethod field if non-nil, zero value otherwise.

### GetComputerSearchMethodOk

`func (o *AccountPreferencesV5) GetComputerSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetComputerSearchMethodOk returns a tuple with the ComputerSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerSearchMethod

`func (o *AccountPreferencesV5) SetComputerSearchMethod(v AccountPreferencesSearchType)`

SetComputerSearchMethod sets ComputerSearchMethod field to given value.


### GetComputerApplicationSearchMethod

`func (o *AccountPreferencesV5) GetComputerApplicationSearchMethod() AccountPreferencesSearchType`

GetComputerApplicationSearchMethod returns the ComputerApplicationSearchMethod field if non-nil, zero value otherwise.

### GetComputerApplicationSearchMethodOk

`func (o *AccountPreferencesV5) GetComputerApplicationSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetComputerApplicationSearchMethodOk returns a tuple with the ComputerApplicationSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerApplicationSearchMethod

`func (o *AccountPreferencesV5) SetComputerApplicationSearchMethod(v AccountPreferencesSearchType)`

SetComputerApplicationSearchMethod sets ComputerApplicationSearchMethod field to given value.


### GetComputerApplicationUsageSearchMethod

`func (o *AccountPreferencesV5) GetComputerApplicationUsageSearchMethod() AccountPreferencesSearchType`

GetComputerApplicationUsageSearchMethod returns the ComputerApplicationUsageSearchMethod field if non-nil, zero value otherwise.

### GetComputerApplicationUsageSearchMethodOk

`func (o *AccountPreferencesV5) GetComputerApplicationUsageSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetComputerApplicationUsageSearchMethodOk returns a tuple with the ComputerApplicationUsageSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerApplicationUsageSearchMethod

`func (o *AccountPreferencesV5) SetComputerApplicationUsageSearchMethod(v AccountPreferencesSearchType)`

SetComputerApplicationUsageSearchMethod sets ComputerApplicationUsageSearchMethod field to given value.


### GetComputerFontSearchMethod

`func (o *AccountPreferencesV5) GetComputerFontSearchMethod() AccountPreferencesSearchType`

GetComputerFontSearchMethod returns the ComputerFontSearchMethod field if non-nil, zero value otherwise.

### GetComputerFontSearchMethodOk

`func (o *AccountPreferencesV5) GetComputerFontSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetComputerFontSearchMethodOk returns a tuple with the ComputerFontSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerFontSearchMethod

`func (o *AccountPreferencesV5) SetComputerFontSearchMethod(v AccountPreferencesSearchType)`

SetComputerFontSearchMethod sets ComputerFontSearchMethod field to given value.

### HasComputerFontSearchMethod

`func (o *AccountPreferencesV5) HasComputerFontSearchMethod() bool`

HasComputerFontSearchMethod returns a boolean if a field has been set.

### GetComputerPluginSearchMethod

`func (o *AccountPreferencesV5) GetComputerPluginSearchMethod() AccountPreferencesSearchType`

GetComputerPluginSearchMethod returns the ComputerPluginSearchMethod field if non-nil, zero value otherwise.

### GetComputerPluginSearchMethodOk

`func (o *AccountPreferencesV5) GetComputerPluginSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetComputerPluginSearchMethodOk returns a tuple with the ComputerPluginSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerPluginSearchMethod

`func (o *AccountPreferencesV5) SetComputerPluginSearchMethod(v AccountPreferencesSearchType)`

SetComputerPluginSearchMethod sets ComputerPluginSearchMethod field to given value.

### HasComputerPluginSearchMethod

`func (o *AccountPreferencesV5) HasComputerPluginSearchMethod() bool`

HasComputerPluginSearchMethod returns a boolean if a field has been set.

### GetComputerSoftwareUpdateSearchMethod

`func (o *AccountPreferencesV5) GetComputerSoftwareUpdateSearchMethod() AccountPreferencesSearchType`

GetComputerSoftwareUpdateSearchMethod returns the ComputerSoftwareUpdateSearchMethod field if non-nil, zero value otherwise.

### GetComputerSoftwareUpdateSearchMethodOk

`func (o *AccountPreferencesV5) GetComputerSoftwareUpdateSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetComputerSoftwareUpdateSearchMethodOk returns a tuple with the ComputerSoftwareUpdateSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerSoftwareUpdateSearchMethod

`func (o *AccountPreferencesV5) SetComputerSoftwareUpdateSearchMethod(v AccountPreferencesSearchType)`

SetComputerSoftwareUpdateSearchMethod sets ComputerSoftwareUpdateSearchMethod field to given value.

### HasComputerSoftwareUpdateSearchMethod

`func (o *AccountPreferencesV5) HasComputerSoftwareUpdateSearchMethod() bool`

HasComputerSoftwareUpdateSearchMethod returns a boolean if a field has been set.

### GetComputerLocalUserAccountSearchMethod

`func (o *AccountPreferencesV5) GetComputerLocalUserAccountSearchMethod() AccountPreferencesSearchType`

GetComputerLocalUserAccountSearchMethod returns the ComputerLocalUserAccountSearchMethod field if non-nil, zero value otherwise.

### GetComputerLocalUserAccountSearchMethodOk

`func (o *AccountPreferencesV5) GetComputerLocalUserAccountSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetComputerLocalUserAccountSearchMethodOk returns a tuple with the ComputerLocalUserAccountSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerLocalUserAccountSearchMethod

`func (o *AccountPreferencesV5) SetComputerLocalUserAccountSearchMethod(v AccountPreferencesSearchType)`

SetComputerLocalUserAccountSearchMethod sets ComputerLocalUserAccountSearchMethod field to given value.


### GetComputerPackageReceiptSearchMethod

`func (o *AccountPreferencesV5) GetComputerPackageReceiptSearchMethod() AccountPreferencesSearchType`

GetComputerPackageReceiptSearchMethod returns the ComputerPackageReceiptSearchMethod field if non-nil, zero value otherwise.

### GetComputerPackageReceiptSearchMethodOk

`func (o *AccountPreferencesV5) GetComputerPackageReceiptSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetComputerPackageReceiptSearchMethodOk returns a tuple with the ComputerPackageReceiptSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerPackageReceiptSearchMethod

`func (o *AccountPreferencesV5) SetComputerPackageReceiptSearchMethod(v AccountPreferencesSearchType)`

SetComputerPackageReceiptSearchMethod sets ComputerPackageReceiptSearchMethod field to given value.


### GetComputerPrinterSearchMethod

`func (o *AccountPreferencesV5) GetComputerPrinterSearchMethod() AccountPreferencesSearchType`

GetComputerPrinterSearchMethod returns the ComputerPrinterSearchMethod field if non-nil, zero value otherwise.

### GetComputerPrinterSearchMethodOk

`func (o *AccountPreferencesV5) GetComputerPrinterSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetComputerPrinterSearchMethodOk returns a tuple with the ComputerPrinterSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerPrinterSearchMethod

`func (o *AccountPreferencesV5) SetComputerPrinterSearchMethod(v AccountPreferencesSearchType)`

SetComputerPrinterSearchMethod sets ComputerPrinterSearchMethod field to given value.


### GetComputerPeripheralSearchMethod

`func (o *AccountPreferencesV5) GetComputerPeripheralSearchMethod() AccountPreferencesSearchType`

GetComputerPeripheralSearchMethod returns the ComputerPeripheralSearchMethod field if non-nil, zero value otherwise.

### GetComputerPeripheralSearchMethodOk

`func (o *AccountPreferencesV5) GetComputerPeripheralSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetComputerPeripheralSearchMethodOk returns a tuple with the ComputerPeripheralSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerPeripheralSearchMethod

`func (o *AccountPreferencesV5) SetComputerPeripheralSearchMethod(v AccountPreferencesSearchType)`

SetComputerPeripheralSearchMethod sets ComputerPeripheralSearchMethod field to given value.

### HasComputerPeripheralSearchMethod

`func (o *AccountPreferencesV5) HasComputerPeripheralSearchMethod() bool`

HasComputerPeripheralSearchMethod returns a boolean if a field has been set.

### GetComputerServiceSearchMethod

`func (o *AccountPreferencesV5) GetComputerServiceSearchMethod() AccountPreferencesSearchType`

GetComputerServiceSearchMethod returns the ComputerServiceSearchMethod field if non-nil, zero value otherwise.

### GetComputerServiceSearchMethodOk

`func (o *AccountPreferencesV5) GetComputerServiceSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetComputerServiceSearchMethodOk returns a tuple with the ComputerServiceSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerServiceSearchMethod

`func (o *AccountPreferencesV5) SetComputerServiceSearchMethod(v AccountPreferencesSearchType)`

SetComputerServiceSearchMethod sets ComputerServiceSearchMethod field to given value.


### GetMobileDeviceSearchMethod

`func (o *AccountPreferencesV5) GetMobileDeviceSearchMethod() AccountPreferencesSearchType`

GetMobileDeviceSearchMethod returns the MobileDeviceSearchMethod field if non-nil, zero value otherwise.

### GetMobileDeviceSearchMethodOk

`func (o *AccountPreferencesV5) GetMobileDeviceSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetMobileDeviceSearchMethodOk returns a tuple with the MobileDeviceSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileDeviceSearchMethod

`func (o *AccountPreferencesV5) SetMobileDeviceSearchMethod(v AccountPreferencesSearchType)`

SetMobileDeviceSearchMethod sets MobileDeviceSearchMethod field to given value.


### GetMobileDeviceAppSearchMethod

`func (o *AccountPreferencesV5) GetMobileDeviceAppSearchMethod() AccountPreferencesSearchType`

GetMobileDeviceAppSearchMethod returns the MobileDeviceAppSearchMethod field if non-nil, zero value otherwise.

### GetMobileDeviceAppSearchMethodOk

`func (o *AccountPreferencesV5) GetMobileDeviceAppSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetMobileDeviceAppSearchMethodOk returns a tuple with the MobileDeviceAppSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileDeviceAppSearchMethod

`func (o *AccountPreferencesV5) SetMobileDeviceAppSearchMethod(v AccountPreferencesSearchType)`

SetMobileDeviceAppSearchMethod sets MobileDeviceAppSearchMethod field to given value.


### GetUserSearchMethod

`func (o *AccountPreferencesV5) GetUserSearchMethod() AccountPreferencesSearchType`

GetUserSearchMethod returns the UserSearchMethod field if non-nil, zero value otherwise.

### GetUserSearchMethodOk

`func (o *AccountPreferencesV5) GetUserSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetUserSearchMethodOk returns a tuple with the UserSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSearchMethod

`func (o *AccountPreferencesV5) SetUserSearchMethod(v AccountPreferencesSearchType)`

SetUserSearchMethod sets UserSearchMethod field to given value.


### GetUserAllContentSearchMethod

`func (o *AccountPreferencesV5) GetUserAllContentSearchMethod() AccountPreferencesSearchType`

GetUserAllContentSearchMethod returns the UserAllContentSearchMethod field if non-nil, zero value otherwise.

### GetUserAllContentSearchMethodOk

`func (o *AccountPreferencesV5) GetUserAllContentSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetUserAllContentSearchMethodOk returns a tuple with the UserAllContentSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAllContentSearchMethod

`func (o *AccountPreferencesV5) SetUserAllContentSearchMethod(v AccountPreferencesSearchType)`

SetUserAllContentSearchMethod sets UserAllContentSearchMethod field to given value.


### GetUserMobileDeviceAppSearchMethod

`func (o *AccountPreferencesV5) GetUserMobileDeviceAppSearchMethod() AccountPreferencesSearchType`

GetUserMobileDeviceAppSearchMethod returns the UserMobileDeviceAppSearchMethod field if non-nil, zero value otherwise.

### GetUserMobileDeviceAppSearchMethodOk

`func (o *AccountPreferencesV5) GetUserMobileDeviceAppSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetUserMobileDeviceAppSearchMethodOk returns a tuple with the UserMobileDeviceAppSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMobileDeviceAppSearchMethod

`func (o *AccountPreferencesV5) SetUserMobileDeviceAppSearchMethod(v AccountPreferencesSearchType)`

SetUserMobileDeviceAppSearchMethod sets UserMobileDeviceAppSearchMethod field to given value.


### GetUserMacAppStoreAppSearchMethod

`func (o *AccountPreferencesV5) GetUserMacAppStoreAppSearchMethod() AccountPreferencesSearchType`

GetUserMacAppStoreAppSearchMethod returns the UserMacAppStoreAppSearchMethod field if non-nil, zero value otherwise.

### GetUserMacAppStoreAppSearchMethodOk

`func (o *AccountPreferencesV5) GetUserMacAppStoreAppSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetUserMacAppStoreAppSearchMethodOk returns a tuple with the UserMacAppStoreAppSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMacAppStoreAppSearchMethod

`func (o *AccountPreferencesV5) SetUserMacAppStoreAppSearchMethod(v AccountPreferencesSearchType)`

SetUserMacAppStoreAppSearchMethod sets UserMacAppStoreAppSearchMethod field to given value.


### GetUserEbookSearchMethod

`func (o *AccountPreferencesV5) GetUserEbookSearchMethod() AccountPreferencesSearchType`

GetUserEbookSearchMethod returns the UserEbookSearchMethod field if non-nil, zero value otherwise.

### GetUserEbookSearchMethodOk

`func (o *AccountPreferencesV5) GetUserEbookSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetUserEbookSearchMethodOk returns a tuple with the UserEbookSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEbookSearchMethod

`func (o *AccountPreferencesV5) SetUserEbookSearchMethod(v AccountPreferencesSearchType)`

SetUserEbookSearchMethod sets UserEbookSearchMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


