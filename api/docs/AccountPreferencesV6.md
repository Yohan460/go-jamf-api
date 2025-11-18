# AccountPreferencesV6

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

### NewAccountPreferencesV6

`func NewAccountPreferencesV6(language string, dateFormat string, timezone string, resultsPerPage int64, userInterfaceDisplayTheme AccountPreferencesUserInterfaceDisplayTheme, disableRelativeDates bool, disablePageLeaveCheck bool, disableTablePagination bool, disableShortcutsTooltips bool, configProfilesSortingMethod string, computerSearchMethod AccountPreferencesSearchType, computerApplicationSearchMethod AccountPreferencesSearchType, computerApplicationUsageSearchMethod AccountPreferencesSearchType, computerLocalUserAccountSearchMethod AccountPreferencesSearchType, computerPackageReceiptSearchMethod AccountPreferencesSearchType, computerPrinterSearchMethod AccountPreferencesSearchType, computerServiceSearchMethod AccountPreferencesSearchType, mobileDeviceSearchMethod AccountPreferencesSearchType, mobileDeviceAppSearchMethod AccountPreferencesSearchType, userSearchMethod AccountPreferencesSearchType, userAllContentSearchMethod AccountPreferencesSearchType, userMobileDeviceAppSearchMethod AccountPreferencesSearchType, userMacAppStoreAppSearchMethod AccountPreferencesSearchType, userEbookSearchMethod AccountPreferencesSearchType, ) *AccountPreferencesV6`

NewAccountPreferencesV6 instantiates a new AccountPreferencesV6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountPreferencesV6WithDefaults

`func NewAccountPreferencesV6WithDefaults() *AccountPreferencesV6`

NewAccountPreferencesV6WithDefaults instantiates a new AccountPreferencesV6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *AccountPreferencesV6) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AccountPreferencesV6) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AccountPreferencesV6) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetDateFormat

`func (o *AccountPreferencesV6) GetDateFormat() string`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *AccountPreferencesV6) GetDateFormatOk() (*string, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *AccountPreferencesV6) SetDateFormat(v string)`

SetDateFormat sets DateFormat field to given value.


### GetTimezone

`func (o *AccountPreferencesV6) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AccountPreferencesV6) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AccountPreferencesV6) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetResultsPerPage

`func (o *AccountPreferencesV6) GetResultsPerPage() int64`

GetResultsPerPage returns the ResultsPerPage field if non-nil, zero value otherwise.

### GetResultsPerPageOk

`func (o *AccountPreferencesV6) GetResultsPerPageOk() (*int64, bool)`

GetResultsPerPageOk returns a tuple with the ResultsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsPerPage

`func (o *AccountPreferencesV6) SetResultsPerPage(v int64)`

SetResultsPerPage sets ResultsPerPage field to given value.


### GetUserInterfaceDisplayTheme

`func (o *AccountPreferencesV6) GetUserInterfaceDisplayTheme() AccountPreferencesUserInterfaceDisplayTheme`

GetUserInterfaceDisplayTheme returns the UserInterfaceDisplayTheme field if non-nil, zero value otherwise.

### GetUserInterfaceDisplayThemeOk

`func (o *AccountPreferencesV6) GetUserInterfaceDisplayThemeOk() (*AccountPreferencesUserInterfaceDisplayTheme, bool)`

GetUserInterfaceDisplayThemeOk returns a tuple with the UserInterfaceDisplayTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInterfaceDisplayTheme

`func (o *AccountPreferencesV6) SetUserInterfaceDisplayTheme(v AccountPreferencesUserInterfaceDisplayTheme)`

SetUserInterfaceDisplayTheme sets UserInterfaceDisplayTheme field to given value.


### GetDisableRelativeDates

`func (o *AccountPreferencesV6) GetDisableRelativeDates() bool`

GetDisableRelativeDates returns the DisableRelativeDates field if non-nil, zero value otherwise.

### GetDisableRelativeDatesOk

`func (o *AccountPreferencesV6) GetDisableRelativeDatesOk() (*bool, bool)`

GetDisableRelativeDatesOk returns a tuple with the DisableRelativeDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRelativeDates

`func (o *AccountPreferencesV6) SetDisableRelativeDates(v bool)`

SetDisableRelativeDates sets DisableRelativeDates field to given value.


### GetDisablePageLeaveCheck

`func (o *AccountPreferencesV6) GetDisablePageLeaveCheck() bool`

GetDisablePageLeaveCheck returns the DisablePageLeaveCheck field if non-nil, zero value otherwise.

### GetDisablePageLeaveCheckOk

`func (o *AccountPreferencesV6) GetDisablePageLeaveCheckOk() (*bool, bool)`

GetDisablePageLeaveCheckOk returns a tuple with the DisablePageLeaveCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisablePageLeaveCheck

`func (o *AccountPreferencesV6) SetDisablePageLeaveCheck(v bool)`

SetDisablePageLeaveCheck sets DisablePageLeaveCheck field to given value.


### GetDisableTablePagination

`func (o *AccountPreferencesV6) GetDisableTablePagination() bool`

GetDisableTablePagination returns the DisableTablePagination field if non-nil, zero value otherwise.

### GetDisableTablePaginationOk

`func (o *AccountPreferencesV6) GetDisableTablePaginationOk() (*bool, bool)`

GetDisableTablePaginationOk returns a tuple with the DisableTablePagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableTablePagination

`func (o *AccountPreferencesV6) SetDisableTablePagination(v bool)`

SetDisableTablePagination sets DisableTablePagination field to given value.


### GetDisableShortcutsTooltips

`func (o *AccountPreferencesV6) GetDisableShortcutsTooltips() bool`

GetDisableShortcutsTooltips returns the DisableShortcutsTooltips field if non-nil, zero value otherwise.

### GetDisableShortcutsTooltipsOk

`func (o *AccountPreferencesV6) GetDisableShortcutsTooltipsOk() (*bool, bool)`

GetDisableShortcutsTooltipsOk returns a tuple with the DisableShortcutsTooltips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableShortcutsTooltips

`func (o *AccountPreferencesV6) SetDisableShortcutsTooltips(v bool)`

SetDisableShortcutsTooltips sets DisableShortcutsTooltips field to given value.


### GetConfigProfilesSortingMethod

`func (o *AccountPreferencesV6) GetConfigProfilesSortingMethod() string`

GetConfigProfilesSortingMethod returns the ConfigProfilesSortingMethod field if non-nil, zero value otherwise.

### GetConfigProfilesSortingMethodOk

`func (o *AccountPreferencesV6) GetConfigProfilesSortingMethodOk() (*string, bool)`

GetConfigProfilesSortingMethodOk returns a tuple with the ConfigProfilesSortingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigProfilesSortingMethod

`func (o *AccountPreferencesV6) SetConfigProfilesSortingMethod(v string)`

SetConfigProfilesSortingMethod sets ConfigProfilesSortingMethod field to given value.


### GetComputerSearchMethod

`func (o *AccountPreferencesV6) GetComputerSearchMethod() AccountPreferencesSearchType`

GetComputerSearchMethod returns the ComputerSearchMethod field if non-nil, zero value otherwise.

### GetComputerSearchMethodOk

`func (o *AccountPreferencesV6) GetComputerSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetComputerSearchMethodOk returns a tuple with the ComputerSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerSearchMethod

`func (o *AccountPreferencesV6) SetComputerSearchMethod(v AccountPreferencesSearchType)`

SetComputerSearchMethod sets ComputerSearchMethod field to given value.


### GetComputerApplicationSearchMethod

`func (o *AccountPreferencesV6) GetComputerApplicationSearchMethod() AccountPreferencesSearchType`

GetComputerApplicationSearchMethod returns the ComputerApplicationSearchMethod field if non-nil, zero value otherwise.

### GetComputerApplicationSearchMethodOk

`func (o *AccountPreferencesV6) GetComputerApplicationSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetComputerApplicationSearchMethodOk returns a tuple with the ComputerApplicationSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerApplicationSearchMethod

`func (o *AccountPreferencesV6) SetComputerApplicationSearchMethod(v AccountPreferencesSearchType)`

SetComputerApplicationSearchMethod sets ComputerApplicationSearchMethod field to given value.


### GetComputerApplicationUsageSearchMethod

`func (o *AccountPreferencesV6) GetComputerApplicationUsageSearchMethod() AccountPreferencesSearchType`

GetComputerApplicationUsageSearchMethod returns the ComputerApplicationUsageSearchMethod field if non-nil, zero value otherwise.

### GetComputerApplicationUsageSearchMethodOk

`func (o *AccountPreferencesV6) GetComputerApplicationUsageSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetComputerApplicationUsageSearchMethodOk returns a tuple with the ComputerApplicationUsageSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerApplicationUsageSearchMethod

`func (o *AccountPreferencesV6) SetComputerApplicationUsageSearchMethod(v AccountPreferencesSearchType)`

SetComputerApplicationUsageSearchMethod sets ComputerApplicationUsageSearchMethod field to given value.


### GetComputerSoftwareUpdateSearchMethod

`func (o *AccountPreferencesV6) GetComputerSoftwareUpdateSearchMethod() AccountPreferencesSearchType`

GetComputerSoftwareUpdateSearchMethod returns the ComputerSoftwareUpdateSearchMethod field if non-nil, zero value otherwise.

### GetComputerSoftwareUpdateSearchMethodOk

`func (o *AccountPreferencesV6) GetComputerSoftwareUpdateSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetComputerSoftwareUpdateSearchMethodOk returns a tuple with the ComputerSoftwareUpdateSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerSoftwareUpdateSearchMethod

`func (o *AccountPreferencesV6) SetComputerSoftwareUpdateSearchMethod(v AccountPreferencesSearchType)`

SetComputerSoftwareUpdateSearchMethod sets ComputerSoftwareUpdateSearchMethod field to given value.

### HasComputerSoftwareUpdateSearchMethod

`func (o *AccountPreferencesV6) HasComputerSoftwareUpdateSearchMethod() bool`

HasComputerSoftwareUpdateSearchMethod returns a boolean if a field has been set.

### GetComputerLocalUserAccountSearchMethod

`func (o *AccountPreferencesV6) GetComputerLocalUserAccountSearchMethod() AccountPreferencesSearchType`

GetComputerLocalUserAccountSearchMethod returns the ComputerLocalUserAccountSearchMethod field if non-nil, zero value otherwise.

### GetComputerLocalUserAccountSearchMethodOk

`func (o *AccountPreferencesV6) GetComputerLocalUserAccountSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetComputerLocalUserAccountSearchMethodOk returns a tuple with the ComputerLocalUserAccountSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerLocalUserAccountSearchMethod

`func (o *AccountPreferencesV6) SetComputerLocalUserAccountSearchMethod(v AccountPreferencesSearchType)`

SetComputerLocalUserAccountSearchMethod sets ComputerLocalUserAccountSearchMethod field to given value.


### GetComputerPackageReceiptSearchMethod

`func (o *AccountPreferencesV6) GetComputerPackageReceiptSearchMethod() AccountPreferencesSearchType`

GetComputerPackageReceiptSearchMethod returns the ComputerPackageReceiptSearchMethod field if non-nil, zero value otherwise.

### GetComputerPackageReceiptSearchMethodOk

`func (o *AccountPreferencesV6) GetComputerPackageReceiptSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetComputerPackageReceiptSearchMethodOk returns a tuple with the ComputerPackageReceiptSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerPackageReceiptSearchMethod

`func (o *AccountPreferencesV6) SetComputerPackageReceiptSearchMethod(v AccountPreferencesSearchType)`

SetComputerPackageReceiptSearchMethod sets ComputerPackageReceiptSearchMethod field to given value.


### GetComputerPrinterSearchMethod

`func (o *AccountPreferencesV6) GetComputerPrinterSearchMethod() AccountPreferencesSearchType`

GetComputerPrinterSearchMethod returns the ComputerPrinterSearchMethod field if non-nil, zero value otherwise.

### GetComputerPrinterSearchMethodOk

`func (o *AccountPreferencesV6) GetComputerPrinterSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetComputerPrinterSearchMethodOk returns a tuple with the ComputerPrinterSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerPrinterSearchMethod

`func (o *AccountPreferencesV6) SetComputerPrinterSearchMethod(v AccountPreferencesSearchType)`

SetComputerPrinterSearchMethod sets ComputerPrinterSearchMethod field to given value.


### GetComputerPeripheralSearchMethod

`func (o *AccountPreferencesV6) GetComputerPeripheralSearchMethod() AccountPreferencesSearchType`

GetComputerPeripheralSearchMethod returns the ComputerPeripheralSearchMethod field if non-nil, zero value otherwise.

### GetComputerPeripheralSearchMethodOk

`func (o *AccountPreferencesV6) GetComputerPeripheralSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetComputerPeripheralSearchMethodOk returns a tuple with the ComputerPeripheralSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerPeripheralSearchMethod

`func (o *AccountPreferencesV6) SetComputerPeripheralSearchMethod(v AccountPreferencesSearchType)`

SetComputerPeripheralSearchMethod sets ComputerPeripheralSearchMethod field to given value.

### HasComputerPeripheralSearchMethod

`func (o *AccountPreferencesV6) HasComputerPeripheralSearchMethod() bool`

HasComputerPeripheralSearchMethod returns a boolean if a field has been set.

### GetComputerServiceSearchMethod

`func (o *AccountPreferencesV6) GetComputerServiceSearchMethod() AccountPreferencesSearchType`

GetComputerServiceSearchMethod returns the ComputerServiceSearchMethod field if non-nil, zero value otherwise.

### GetComputerServiceSearchMethodOk

`func (o *AccountPreferencesV6) GetComputerServiceSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetComputerServiceSearchMethodOk returns a tuple with the ComputerServiceSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerServiceSearchMethod

`func (o *AccountPreferencesV6) SetComputerServiceSearchMethod(v AccountPreferencesSearchType)`

SetComputerServiceSearchMethod sets ComputerServiceSearchMethod field to given value.


### GetMobileDeviceSearchMethod

`func (o *AccountPreferencesV6) GetMobileDeviceSearchMethod() AccountPreferencesSearchType`

GetMobileDeviceSearchMethod returns the MobileDeviceSearchMethod field if non-nil, zero value otherwise.

### GetMobileDeviceSearchMethodOk

`func (o *AccountPreferencesV6) GetMobileDeviceSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetMobileDeviceSearchMethodOk returns a tuple with the MobileDeviceSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileDeviceSearchMethod

`func (o *AccountPreferencesV6) SetMobileDeviceSearchMethod(v AccountPreferencesSearchType)`

SetMobileDeviceSearchMethod sets MobileDeviceSearchMethod field to given value.


### GetMobileDeviceAppSearchMethod

`func (o *AccountPreferencesV6) GetMobileDeviceAppSearchMethod() AccountPreferencesSearchType`

GetMobileDeviceAppSearchMethod returns the MobileDeviceAppSearchMethod field if non-nil, zero value otherwise.

### GetMobileDeviceAppSearchMethodOk

`func (o *AccountPreferencesV6) GetMobileDeviceAppSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetMobileDeviceAppSearchMethodOk returns a tuple with the MobileDeviceAppSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileDeviceAppSearchMethod

`func (o *AccountPreferencesV6) SetMobileDeviceAppSearchMethod(v AccountPreferencesSearchType)`

SetMobileDeviceAppSearchMethod sets MobileDeviceAppSearchMethod field to given value.


### GetUserSearchMethod

`func (o *AccountPreferencesV6) GetUserSearchMethod() AccountPreferencesSearchType`

GetUserSearchMethod returns the UserSearchMethod field if non-nil, zero value otherwise.

### GetUserSearchMethodOk

`func (o *AccountPreferencesV6) GetUserSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetUserSearchMethodOk returns a tuple with the UserSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSearchMethod

`func (o *AccountPreferencesV6) SetUserSearchMethod(v AccountPreferencesSearchType)`

SetUserSearchMethod sets UserSearchMethod field to given value.


### GetUserAllContentSearchMethod

`func (o *AccountPreferencesV6) GetUserAllContentSearchMethod() AccountPreferencesSearchType`

GetUserAllContentSearchMethod returns the UserAllContentSearchMethod field if non-nil, zero value otherwise.

### GetUserAllContentSearchMethodOk

`func (o *AccountPreferencesV6) GetUserAllContentSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetUserAllContentSearchMethodOk returns a tuple with the UserAllContentSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAllContentSearchMethod

`func (o *AccountPreferencesV6) SetUserAllContentSearchMethod(v AccountPreferencesSearchType)`

SetUserAllContentSearchMethod sets UserAllContentSearchMethod field to given value.


### GetUserMobileDeviceAppSearchMethod

`func (o *AccountPreferencesV6) GetUserMobileDeviceAppSearchMethod() AccountPreferencesSearchType`

GetUserMobileDeviceAppSearchMethod returns the UserMobileDeviceAppSearchMethod field if non-nil, zero value otherwise.

### GetUserMobileDeviceAppSearchMethodOk

`func (o *AccountPreferencesV6) GetUserMobileDeviceAppSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetUserMobileDeviceAppSearchMethodOk returns a tuple with the UserMobileDeviceAppSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMobileDeviceAppSearchMethod

`func (o *AccountPreferencesV6) SetUserMobileDeviceAppSearchMethod(v AccountPreferencesSearchType)`

SetUserMobileDeviceAppSearchMethod sets UserMobileDeviceAppSearchMethod field to given value.


### GetUserMacAppStoreAppSearchMethod

`func (o *AccountPreferencesV6) GetUserMacAppStoreAppSearchMethod() AccountPreferencesSearchType`

GetUserMacAppStoreAppSearchMethod returns the UserMacAppStoreAppSearchMethod field if non-nil, zero value otherwise.

### GetUserMacAppStoreAppSearchMethodOk

`func (o *AccountPreferencesV6) GetUserMacAppStoreAppSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetUserMacAppStoreAppSearchMethodOk returns a tuple with the UserMacAppStoreAppSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMacAppStoreAppSearchMethod

`func (o *AccountPreferencesV6) SetUserMacAppStoreAppSearchMethod(v AccountPreferencesSearchType)`

SetUserMacAppStoreAppSearchMethod sets UserMacAppStoreAppSearchMethod field to given value.


### GetUserEbookSearchMethod

`func (o *AccountPreferencesV6) GetUserEbookSearchMethod() AccountPreferencesSearchType`

GetUserEbookSearchMethod returns the UserEbookSearchMethod field if non-nil, zero value otherwise.

### GetUserEbookSearchMethodOk

`func (o *AccountPreferencesV6) GetUserEbookSearchMethodOk() (*AccountPreferencesSearchType, bool)`

GetUserEbookSearchMethodOk returns a tuple with the UserEbookSearchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEbookSearchMethod

`func (o *AccountPreferencesV6) SetUserEbookSearchMethod(v AccountPreferencesSearchType)`

SetUserEbookSearchMethod sets UserEbookSearchMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


