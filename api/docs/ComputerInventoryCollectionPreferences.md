# ComputerInventoryCollectionPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorApplicationUsage** | Pointer to **bool** |  | [optional] [default to false]
**IncludeFonts** | Pointer to **bool** |  | [optional] [default to false]
**IncludePlugins** | Pointer to **bool** |  | [optional] [default to false]
**IncludePackages** | Pointer to **bool** |  | [optional] [default to false]
**IncludeSoftwareUpdates** | Pointer to **bool** |  | [optional] [default to false]
**IncludeSoftwareId** | Pointer to **bool** |  | [optional] [default to false]
**IncludeAccounts** | Pointer to **bool** |  | [optional] [default to false]
**CalculateSizes** | Pointer to **bool** |  | [optional] [default to false]
**IncludeHiddenAccounts** | Pointer to **bool** |  | [optional] [default to false]
**IncludePrinters** | Pointer to **bool** |  | [optional] [default to false]
**IncludeServices** | Pointer to **bool** |  | [optional] [default to false]
**CollectSyncedMobileDeviceInfo** | Pointer to **bool** |  | [optional] [default to false]
**UpdateLdapInfoOnComputerInventorySubmissions** | Pointer to **bool** |  | [optional] [default to false]
**MonitorBeacons** | Pointer to **bool** |  | [optional] [default to false]
**AllowChangingUserAndLocation** | Pointer to **bool** |  | [optional] [default to true]
**UseUnixUserPaths** | Pointer to **bool** |  | [optional] [default to true]
**CollectUnmanagedCertificates** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewComputerInventoryCollectionPreferences

`func NewComputerInventoryCollectionPreferences() *ComputerInventoryCollectionPreferences`

NewComputerInventoryCollectionPreferences instantiates a new ComputerInventoryCollectionPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerInventoryCollectionPreferencesWithDefaults

`func NewComputerInventoryCollectionPreferencesWithDefaults() *ComputerInventoryCollectionPreferences`

NewComputerInventoryCollectionPreferencesWithDefaults instantiates a new ComputerInventoryCollectionPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorApplicationUsage

`func (o *ComputerInventoryCollectionPreferences) GetMonitorApplicationUsage() bool`

GetMonitorApplicationUsage returns the MonitorApplicationUsage field if non-nil, zero value otherwise.

### GetMonitorApplicationUsageOk

`func (o *ComputerInventoryCollectionPreferences) GetMonitorApplicationUsageOk() (*bool, bool)`

GetMonitorApplicationUsageOk returns a tuple with the MonitorApplicationUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorApplicationUsage

`func (o *ComputerInventoryCollectionPreferences) SetMonitorApplicationUsage(v bool)`

SetMonitorApplicationUsage sets MonitorApplicationUsage field to given value.

### HasMonitorApplicationUsage

`func (o *ComputerInventoryCollectionPreferences) HasMonitorApplicationUsage() bool`

HasMonitorApplicationUsage returns a boolean if a field has been set.

### GetIncludeFonts

`func (o *ComputerInventoryCollectionPreferences) GetIncludeFonts() bool`

GetIncludeFonts returns the IncludeFonts field if non-nil, zero value otherwise.

### GetIncludeFontsOk

`func (o *ComputerInventoryCollectionPreferences) GetIncludeFontsOk() (*bool, bool)`

GetIncludeFontsOk returns a tuple with the IncludeFonts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFonts

`func (o *ComputerInventoryCollectionPreferences) SetIncludeFonts(v bool)`

SetIncludeFonts sets IncludeFonts field to given value.

### HasIncludeFonts

`func (o *ComputerInventoryCollectionPreferences) HasIncludeFonts() bool`

HasIncludeFonts returns a boolean if a field has been set.

### GetIncludePlugins

`func (o *ComputerInventoryCollectionPreferences) GetIncludePlugins() bool`

GetIncludePlugins returns the IncludePlugins field if non-nil, zero value otherwise.

### GetIncludePluginsOk

`func (o *ComputerInventoryCollectionPreferences) GetIncludePluginsOk() (*bool, bool)`

GetIncludePluginsOk returns a tuple with the IncludePlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePlugins

`func (o *ComputerInventoryCollectionPreferences) SetIncludePlugins(v bool)`

SetIncludePlugins sets IncludePlugins field to given value.

### HasIncludePlugins

`func (o *ComputerInventoryCollectionPreferences) HasIncludePlugins() bool`

HasIncludePlugins returns a boolean if a field has been set.

### GetIncludePackages

`func (o *ComputerInventoryCollectionPreferences) GetIncludePackages() bool`

GetIncludePackages returns the IncludePackages field if non-nil, zero value otherwise.

### GetIncludePackagesOk

`func (o *ComputerInventoryCollectionPreferences) GetIncludePackagesOk() (*bool, bool)`

GetIncludePackagesOk returns a tuple with the IncludePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePackages

`func (o *ComputerInventoryCollectionPreferences) SetIncludePackages(v bool)`

SetIncludePackages sets IncludePackages field to given value.

### HasIncludePackages

`func (o *ComputerInventoryCollectionPreferences) HasIncludePackages() bool`

HasIncludePackages returns a boolean if a field has been set.

### GetIncludeSoftwareUpdates

`func (o *ComputerInventoryCollectionPreferences) GetIncludeSoftwareUpdates() bool`

GetIncludeSoftwareUpdates returns the IncludeSoftwareUpdates field if non-nil, zero value otherwise.

### GetIncludeSoftwareUpdatesOk

`func (o *ComputerInventoryCollectionPreferences) GetIncludeSoftwareUpdatesOk() (*bool, bool)`

GetIncludeSoftwareUpdatesOk returns a tuple with the IncludeSoftwareUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSoftwareUpdates

`func (o *ComputerInventoryCollectionPreferences) SetIncludeSoftwareUpdates(v bool)`

SetIncludeSoftwareUpdates sets IncludeSoftwareUpdates field to given value.

### HasIncludeSoftwareUpdates

`func (o *ComputerInventoryCollectionPreferences) HasIncludeSoftwareUpdates() bool`

HasIncludeSoftwareUpdates returns a boolean if a field has been set.

### GetIncludeSoftwareId

`func (o *ComputerInventoryCollectionPreferences) GetIncludeSoftwareId() bool`

GetIncludeSoftwareId returns the IncludeSoftwareId field if non-nil, zero value otherwise.

### GetIncludeSoftwareIdOk

`func (o *ComputerInventoryCollectionPreferences) GetIncludeSoftwareIdOk() (*bool, bool)`

GetIncludeSoftwareIdOk returns a tuple with the IncludeSoftwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSoftwareId

`func (o *ComputerInventoryCollectionPreferences) SetIncludeSoftwareId(v bool)`

SetIncludeSoftwareId sets IncludeSoftwareId field to given value.

### HasIncludeSoftwareId

`func (o *ComputerInventoryCollectionPreferences) HasIncludeSoftwareId() bool`

HasIncludeSoftwareId returns a boolean if a field has been set.

### GetIncludeAccounts

`func (o *ComputerInventoryCollectionPreferences) GetIncludeAccounts() bool`

GetIncludeAccounts returns the IncludeAccounts field if non-nil, zero value otherwise.

### GetIncludeAccountsOk

`func (o *ComputerInventoryCollectionPreferences) GetIncludeAccountsOk() (*bool, bool)`

GetIncludeAccountsOk returns a tuple with the IncludeAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAccounts

`func (o *ComputerInventoryCollectionPreferences) SetIncludeAccounts(v bool)`

SetIncludeAccounts sets IncludeAccounts field to given value.

### HasIncludeAccounts

`func (o *ComputerInventoryCollectionPreferences) HasIncludeAccounts() bool`

HasIncludeAccounts returns a boolean if a field has been set.

### GetCalculateSizes

`func (o *ComputerInventoryCollectionPreferences) GetCalculateSizes() bool`

GetCalculateSizes returns the CalculateSizes field if non-nil, zero value otherwise.

### GetCalculateSizesOk

`func (o *ComputerInventoryCollectionPreferences) GetCalculateSizesOk() (*bool, bool)`

GetCalculateSizesOk returns a tuple with the CalculateSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculateSizes

`func (o *ComputerInventoryCollectionPreferences) SetCalculateSizes(v bool)`

SetCalculateSizes sets CalculateSizes field to given value.

### HasCalculateSizes

`func (o *ComputerInventoryCollectionPreferences) HasCalculateSizes() bool`

HasCalculateSizes returns a boolean if a field has been set.

### GetIncludeHiddenAccounts

`func (o *ComputerInventoryCollectionPreferences) GetIncludeHiddenAccounts() bool`

GetIncludeHiddenAccounts returns the IncludeHiddenAccounts field if non-nil, zero value otherwise.

### GetIncludeHiddenAccountsOk

`func (o *ComputerInventoryCollectionPreferences) GetIncludeHiddenAccountsOk() (*bool, bool)`

GetIncludeHiddenAccountsOk returns a tuple with the IncludeHiddenAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeHiddenAccounts

`func (o *ComputerInventoryCollectionPreferences) SetIncludeHiddenAccounts(v bool)`

SetIncludeHiddenAccounts sets IncludeHiddenAccounts field to given value.

### HasIncludeHiddenAccounts

`func (o *ComputerInventoryCollectionPreferences) HasIncludeHiddenAccounts() bool`

HasIncludeHiddenAccounts returns a boolean if a field has been set.

### GetIncludePrinters

`func (o *ComputerInventoryCollectionPreferences) GetIncludePrinters() bool`

GetIncludePrinters returns the IncludePrinters field if non-nil, zero value otherwise.

### GetIncludePrintersOk

`func (o *ComputerInventoryCollectionPreferences) GetIncludePrintersOk() (*bool, bool)`

GetIncludePrintersOk returns a tuple with the IncludePrinters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePrinters

`func (o *ComputerInventoryCollectionPreferences) SetIncludePrinters(v bool)`

SetIncludePrinters sets IncludePrinters field to given value.

### HasIncludePrinters

`func (o *ComputerInventoryCollectionPreferences) HasIncludePrinters() bool`

HasIncludePrinters returns a boolean if a field has been set.

### GetIncludeServices

`func (o *ComputerInventoryCollectionPreferences) GetIncludeServices() bool`

GetIncludeServices returns the IncludeServices field if non-nil, zero value otherwise.

### GetIncludeServicesOk

`func (o *ComputerInventoryCollectionPreferences) GetIncludeServicesOk() (*bool, bool)`

GetIncludeServicesOk returns a tuple with the IncludeServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeServices

`func (o *ComputerInventoryCollectionPreferences) SetIncludeServices(v bool)`

SetIncludeServices sets IncludeServices field to given value.

### HasIncludeServices

`func (o *ComputerInventoryCollectionPreferences) HasIncludeServices() bool`

HasIncludeServices returns a boolean if a field has been set.

### GetCollectSyncedMobileDeviceInfo

`func (o *ComputerInventoryCollectionPreferences) GetCollectSyncedMobileDeviceInfo() bool`

GetCollectSyncedMobileDeviceInfo returns the CollectSyncedMobileDeviceInfo field if non-nil, zero value otherwise.

### GetCollectSyncedMobileDeviceInfoOk

`func (o *ComputerInventoryCollectionPreferences) GetCollectSyncedMobileDeviceInfoOk() (*bool, bool)`

GetCollectSyncedMobileDeviceInfoOk returns a tuple with the CollectSyncedMobileDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectSyncedMobileDeviceInfo

`func (o *ComputerInventoryCollectionPreferences) SetCollectSyncedMobileDeviceInfo(v bool)`

SetCollectSyncedMobileDeviceInfo sets CollectSyncedMobileDeviceInfo field to given value.

### HasCollectSyncedMobileDeviceInfo

`func (o *ComputerInventoryCollectionPreferences) HasCollectSyncedMobileDeviceInfo() bool`

HasCollectSyncedMobileDeviceInfo returns a boolean if a field has been set.

### GetUpdateLdapInfoOnComputerInventorySubmissions

`func (o *ComputerInventoryCollectionPreferences) GetUpdateLdapInfoOnComputerInventorySubmissions() bool`

GetUpdateLdapInfoOnComputerInventorySubmissions returns the UpdateLdapInfoOnComputerInventorySubmissions field if non-nil, zero value otherwise.

### GetUpdateLdapInfoOnComputerInventorySubmissionsOk

`func (o *ComputerInventoryCollectionPreferences) GetUpdateLdapInfoOnComputerInventorySubmissionsOk() (*bool, bool)`

GetUpdateLdapInfoOnComputerInventorySubmissionsOk returns a tuple with the UpdateLdapInfoOnComputerInventorySubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateLdapInfoOnComputerInventorySubmissions

`func (o *ComputerInventoryCollectionPreferences) SetUpdateLdapInfoOnComputerInventorySubmissions(v bool)`

SetUpdateLdapInfoOnComputerInventorySubmissions sets UpdateLdapInfoOnComputerInventorySubmissions field to given value.

### HasUpdateLdapInfoOnComputerInventorySubmissions

`func (o *ComputerInventoryCollectionPreferences) HasUpdateLdapInfoOnComputerInventorySubmissions() bool`

HasUpdateLdapInfoOnComputerInventorySubmissions returns a boolean if a field has been set.

### GetMonitorBeacons

`func (o *ComputerInventoryCollectionPreferences) GetMonitorBeacons() bool`

GetMonitorBeacons returns the MonitorBeacons field if non-nil, zero value otherwise.

### GetMonitorBeaconsOk

`func (o *ComputerInventoryCollectionPreferences) GetMonitorBeaconsOk() (*bool, bool)`

GetMonitorBeaconsOk returns a tuple with the MonitorBeacons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorBeacons

`func (o *ComputerInventoryCollectionPreferences) SetMonitorBeacons(v bool)`

SetMonitorBeacons sets MonitorBeacons field to given value.

### HasMonitorBeacons

`func (o *ComputerInventoryCollectionPreferences) HasMonitorBeacons() bool`

HasMonitorBeacons returns a boolean if a field has been set.

### GetAllowChangingUserAndLocation

`func (o *ComputerInventoryCollectionPreferences) GetAllowChangingUserAndLocation() bool`

GetAllowChangingUserAndLocation returns the AllowChangingUserAndLocation field if non-nil, zero value otherwise.

### GetAllowChangingUserAndLocationOk

`func (o *ComputerInventoryCollectionPreferences) GetAllowChangingUserAndLocationOk() (*bool, bool)`

GetAllowChangingUserAndLocationOk returns a tuple with the AllowChangingUserAndLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowChangingUserAndLocation

`func (o *ComputerInventoryCollectionPreferences) SetAllowChangingUserAndLocation(v bool)`

SetAllowChangingUserAndLocation sets AllowChangingUserAndLocation field to given value.

### HasAllowChangingUserAndLocation

`func (o *ComputerInventoryCollectionPreferences) HasAllowChangingUserAndLocation() bool`

HasAllowChangingUserAndLocation returns a boolean if a field has been set.

### GetUseUnixUserPaths

`func (o *ComputerInventoryCollectionPreferences) GetUseUnixUserPaths() bool`

GetUseUnixUserPaths returns the UseUnixUserPaths field if non-nil, zero value otherwise.

### GetUseUnixUserPathsOk

`func (o *ComputerInventoryCollectionPreferences) GetUseUnixUserPathsOk() (*bool, bool)`

GetUseUnixUserPathsOk returns a tuple with the UseUnixUserPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUnixUserPaths

`func (o *ComputerInventoryCollectionPreferences) SetUseUnixUserPaths(v bool)`

SetUseUnixUserPaths sets UseUnixUserPaths field to given value.

### HasUseUnixUserPaths

`func (o *ComputerInventoryCollectionPreferences) HasUseUnixUserPaths() bool`

HasUseUnixUserPaths returns a boolean if a field has been set.

### GetCollectUnmanagedCertificates

`func (o *ComputerInventoryCollectionPreferences) GetCollectUnmanagedCertificates() bool`

GetCollectUnmanagedCertificates returns the CollectUnmanagedCertificates field if non-nil, zero value otherwise.

### GetCollectUnmanagedCertificatesOk

`func (o *ComputerInventoryCollectionPreferences) GetCollectUnmanagedCertificatesOk() (*bool, bool)`

GetCollectUnmanagedCertificatesOk returns a tuple with the CollectUnmanagedCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectUnmanagedCertificates

`func (o *ComputerInventoryCollectionPreferences) SetCollectUnmanagedCertificates(v bool)`

SetCollectUnmanagedCertificates sets CollectUnmanagedCertificates field to given value.

### HasCollectUnmanagedCertificates

`func (o *ComputerInventoryCollectionPreferences) HasCollectUnmanagedCertificates() bool`

HasCollectUnmanagedCertificates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


