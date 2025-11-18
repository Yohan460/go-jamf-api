# ComputerInventoryCollectionPreferencesV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorApplicationUsage** | Pointer to **bool** |  | [optional] [default to false]
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

### NewComputerInventoryCollectionPreferencesV2

`func NewComputerInventoryCollectionPreferencesV2() *ComputerInventoryCollectionPreferencesV2`

NewComputerInventoryCollectionPreferencesV2 instantiates a new ComputerInventoryCollectionPreferencesV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerInventoryCollectionPreferencesV2WithDefaults

`func NewComputerInventoryCollectionPreferencesV2WithDefaults() *ComputerInventoryCollectionPreferencesV2`

NewComputerInventoryCollectionPreferencesV2WithDefaults instantiates a new ComputerInventoryCollectionPreferencesV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorApplicationUsage

`func (o *ComputerInventoryCollectionPreferencesV2) GetMonitorApplicationUsage() bool`

GetMonitorApplicationUsage returns the MonitorApplicationUsage field if non-nil, zero value otherwise.

### GetMonitorApplicationUsageOk

`func (o *ComputerInventoryCollectionPreferencesV2) GetMonitorApplicationUsageOk() (*bool, bool)`

GetMonitorApplicationUsageOk returns a tuple with the MonitorApplicationUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorApplicationUsage

`func (o *ComputerInventoryCollectionPreferencesV2) SetMonitorApplicationUsage(v bool)`

SetMonitorApplicationUsage sets MonitorApplicationUsage field to given value.

### HasMonitorApplicationUsage

`func (o *ComputerInventoryCollectionPreferencesV2) HasMonitorApplicationUsage() bool`

HasMonitorApplicationUsage returns a boolean if a field has been set.

### GetIncludePackages

`func (o *ComputerInventoryCollectionPreferencesV2) GetIncludePackages() bool`

GetIncludePackages returns the IncludePackages field if non-nil, zero value otherwise.

### GetIncludePackagesOk

`func (o *ComputerInventoryCollectionPreferencesV2) GetIncludePackagesOk() (*bool, bool)`

GetIncludePackagesOk returns a tuple with the IncludePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePackages

`func (o *ComputerInventoryCollectionPreferencesV2) SetIncludePackages(v bool)`

SetIncludePackages sets IncludePackages field to given value.

### HasIncludePackages

`func (o *ComputerInventoryCollectionPreferencesV2) HasIncludePackages() bool`

HasIncludePackages returns a boolean if a field has been set.

### GetIncludeSoftwareUpdates

`func (o *ComputerInventoryCollectionPreferencesV2) GetIncludeSoftwareUpdates() bool`

GetIncludeSoftwareUpdates returns the IncludeSoftwareUpdates field if non-nil, zero value otherwise.

### GetIncludeSoftwareUpdatesOk

`func (o *ComputerInventoryCollectionPreferencesV2) GetIncludeSoftwareUpdatesOk() (*bool, bool)`

GetIncludeSoftwareUpdatesOk returns a tuple with the IncludeSoftwareUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSoftwareUpdates

`func (o *ComputerInventoryCollectionPreferencesV2) SetIncludeSoftwareUpdates(v bool)`

SetIncludeSoftwareUpdates sets IncludeSoftwareUpdates field to given value.

### HasIncludeSoftwareUpdates

`func (o *ComputerInventoryCollectionPreferencesV2) HasIncludeSoftwareUpdates() bool`

HasIncludeSoftwareUpdates returns a boolean if a field has been set.

### GetIncludeSoftwareId

`func (o *ComputerInventoryCollectionPreferencesV2) GetIncludeSoftwareId() bool`

GetIncludeSoftwareId returns the IncludeSoftwareId field if non-nil, zero value otherwise.

### GetIncludeSoftwareIdOk

`func (o *ComputerInventoryCollectionPreferencesV2) GetIncludeSoftwareIdOk() (*bool, bool)`

GetIncludeSoftwareIdOk returns a tuple with the IncludeSoftwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSoftwareId

`func (o *ComputerInventoryCollectionPreferencesV2) SetIncludeSoftwareId(v bool)`

SetIncludeSoftwareId sets IncludeSoftwareId field to given value.

### HasIncludeSoftwareId

`func (o *ComputerInventoryCollectionPreferencesV2) HasIncludeSoftwareId() bool`

HasIncludeSoftwareId returns a boolean if a field has been set.

### GetIncludeAccounts

`func (o *ComputerInventoryCollectionPreferencesV2) GetIncludeAccounts() bool`

GetIncludeAccounts returns the IncludeAccounts field if non-nil, zero value otherwise.

### GetIncludeAccountsOk

`func (o *ComputerInventoryCollectionPreferencesV2) GetIncludeAccountsOk() (*bool, bool)`

GetIncludeAccountsOk returns a tuple with the IncludeAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAccounts

`func (o *ComputerInventoryCollectionPreferencesV2) SetIncludeAccounts(v bool)`

SetIncludeAccounts sets IncludeAccounts field to given value.

### HasIncludeAccounts

`func (o *ComputerInventoryCollectionPreferencesV2) HasIncludeAccounts() bool`

HasIncludeAccounts returns a boolean if a field has been set.

### GetCalculateSizes

`func (o *ComputerInventoryCollectionPreferencesV2) GetCalculateSizes() bool`

GetCalculateSizes returns the CalculateSizes field if non-nil, zero value otherwise.

### GetCalculateSizesOk

`func (o *ComputerInventoryCollectionPreferencesV2) GetCalculateSizesOk() (*bool, bool)`

GetCalculateSizesOk returns a tuple with the CalculateSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculateSizes

`func (o *ComputerInventoryCollectionPreferencesV2) SetCalculateSizes(v bool)`

SetCalculateSizes sets CalculateSizes field to given value.

### HasCalculateSizes

`func (o *ComputerInventoryCollectionPreferencesV2) HasCalculateSizes() bool`

HasCalculateSizes returns a boolean if a field has been set.

### GetIncludeHiddenAccounts

`func (o *ComputerInventoryCollectionPreferencesV2) GetIncludeHiddenAccounts() bool`

GetIncludeHiddenAccounts returns the IncludeHiddenAccounts field if non-nil, zero value otherwise.

### GetIncludeHiddenAccountsOk

`func (o *ComputerInventoryCollectionPreferencesV2) GetIncludeHiddenAccountsOk() (*bool, bool)`

GetIncludeHiddenAccountsOk returns a tuple with the IncludeHiddenAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeHiddenAccounts

`func (o *ComputerInventoryCollectionPreferencesV2) SetIncludeHiddenAccounts(v bool)`

SetIncludeHiddenAccounts sets IncludeHiddenAccounts field to given value.

### HasIncludeHiddenAccounts

`func (o *ComputerInventoryCollectionPreferencesV2) HasIncludeHiddenAccounts() bool`

HasIncludeHiddenAccounts returns a boolean if a field has been set.

### GetIncludePrinters

`func (o *ComputerInventoryCollectionPreferencesV2) GetIncludePrinters() bool`

GetIncludePrinters returns the IncludePrinters field if non-nil, zero value otherwise.

### GetIncludePrintersOk

`func (o *ComputerInventoryCollectionPreferencesV2) GetIncludePrintersOk() (*bool, bool)`

GetIncludePrintersOk returns a tuple with the IncludePrinters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePrinters

`func (o *ComputerInventoryCollectionPreferencesV2) SetIncludePrinters(v bool)`

SetIncludePrinters sets IncludePrinters field to given value.

### HasIncludePrinters

`func (o *ComputerInventoryCollectionPreferencesV2) HasIncludePrinters() bool`

HasIncludePrinters returns a boolean if a field has been set.

### GetIncludeServices

`func (o *ComputerInventoryCollectionPreferencesV2) GetIncludeServices() bool`

GetIncludeServices returns the IncludeServices field if non-nil, zero value otherwise.

### GetIncludeServicesOk

`func (o *ComputerInventoryCollectionPreferencesV2) GetIncludeServicesOk() (*bool, bool)`

GetIncludeServicesOk returns a tuple with the IncludeServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeServices

`func (o *ComputerInventoryCollectionPreferencesV2) SetIncludeServices(v bool)`

SetIncludeServices sets IncludeServices field to given value.

### HasIncludeServices

`func (o *ComputerInventoryCollectionPreferencesV2) HasIncludeServices() bool`

HasIncludeServices returns a boolean if a field has been set.

### GetCollectSyncedMobileDeviceInfo

`func (o *ComputerInventoryCollectionPreferencesV2) GetCollectSyncedMobileDeviceInfo() bool`

GetCollectSyncedMobileDeviceInfo returns the CollectSyncedMobileDeviceInfo field if non-nil, zero value otherwise.

### GetCollectSyncedMobileDeviceInfoOk

`func (o *ComputerInventoryCollectionPreferencesV2) GetCollectSyncedMobileDeviceInfoOk() (*bool, bool)`

GetCollectSyncedMobileDeviceInfoOk returns a tuple with the CollectSyncedMobileDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectSyncedMobileDeviceInfo

`func (o *ComputerInventoryCollectionPreferencesV2) SetCollectSyncedMobileDeviceInfo(v bool)`

SetCollectSyncedMobileDeviceInfo sets CollectSyncedMobileDeviceInfo field to given value.

### HasCollectSyncedMobileDeviceInfo

`func (o *ComputerInventoryCollectionPreferencesV2) HasCollectSyncedMobileDeviceInfo() bool`

HasCollectSyncedMobileDeviceInfo returns a boolean if a field has been set.

### GetUpdateLdapInfoOnComputerInventorySubmissions

`func (o *ComputerInventoryCollectionPreferencesV2) GetUpdateLdapInfoOnComputerInventorySubmissions() bool`

GetUpdateLdapInfoOnComputerInventorySubmissions returns the UpdateLdapInfoOnComputerInventorySubmissions field if non-nil, zero value otherwise.

### GetUpdateLdapInfoOnComputerInventorySubmissionsOk

`func (o *ComputerInventoryCollectionPreferencesV2) GetUpdateLdapInfoOnComputerInventorySubmissionsOk() (*bool, bool)`

GetUpdateLdapInfoOnComputerInventorySubmissionsOk returns a tuple with the UpdateLdapInfoOnComputerInventorySubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateLdapInfoOnComputerInventorySubmissions

`func (o *ComputerInventoryCollectionPreferencesV2) SetUpdateLdapInfoOnComputerInventorySubmissions(v bool)`

SetUpdateLdapInfoOnComputerInventorySubmissions sets UpdateLdapInfoOnComputerInventorySubmissions field to given value.

### HasUpdateLdapInfoOnComputerInventorySubmissions

`func (o *ComputerInventoryCollectionPreferencesV2) HasUpdateLdapInfoOnComputerInventorySubmissions() bool`

HasUpdateLdapInfoOnComputerInventorySubmissions returns a boolean if a field has been set.

### GetMonitorBeacons

`func (o *ComputerInventoryCollectionPreferencesV2) GetMonitorBeacons() bool`

GetMonitorBeacons returns the MonitorBeacons field if non-nil, zero value otherwise.

### GetMonitorBeaconsOk

`func (o *ComputerInventoryCollectionPreferencesV2) GetMonitorBeaconsOk() (*bool, bool)`

GetMonitorBeaconsOk returns a tuple with the MonitorBeacons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorBeacons

`func (o *ComputerInventoryCollectionPreferencesV2) SetMonitorBeacons(v bool)`

SetMonitorBeacons sets MonitorBeacons field to given value.

### HasMonitorBeacons

`func (o *ComputerInventoryCollectionPreferencesV2) HasMonitorBeacons() bool`

HasMonitorBeacons returns a boolean if a field has been set.

### GetAllowChangingUserAndLocation

`func (o *ComputerInventoryCollectionPreferencesV2) GetAllowChangingUserAndLocation() bool`

GetAllowChangingUserAndLocation returns the AllowChangingUserAndLocation field if non-nil, zero value otherwise.

### GetAllowChangingUserAndLocationOk

`func (o *ComputerInventoryCollectionPreferencesV2) GetAllowChangingUserAndLocationOk() (*bool, bool)`

GetAllowChangingUserAndLocationOk returns a tuple with the AllowChangingUserAndLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowChangingUserAndLocation

`func (o *ComputerInventoryCollectionPreferencesV2) SetAllowChangingUserAndLocation(v bool)`

SetAllowChangingUserAndLocation sets AllowChangingUserAndLocation field to given value.

### HasAllowChangingUserAndLocation

`func (o *ComputerInventoryCollectionPreferencesV2) HasAllowChangingUserAndLocation() bool`

HasAllowChangingUserAndLocation returns a boolean if a field has been set.

### GetUseUnixUserPaths

`func (o *ComputerInventoryCollectionPreferencesV2) GetUseUnixUserPaths() bool`

GetUseUnixUserPaths returns the UseUnixUserPaths field if non-nil, zero value otherwise.

### GetUseUnixUserPathsOk

`func (o *ComputerInventoryCollectionPreferencesV2) GetUseUnixUserPathsOk() (*bool, bool)`

GetUseUnixUserPathsOk returns a tuple with the UseUnixUserPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUnixUserPaths

`func (o *ComputerInventoryCollectionPreferencesV2) SetUseUnixUserPaths(v bool)`

SetUseUnixUserPaths sets UseUnixUserPaths field to given value.

### HasUseUnixUserPaths

`func (o *ComputerInventoryCollectionPreferencesV2) HasUseUnixUserPaths() bool`

HasUseUnixUserPaths returns a boolean if a field has been set.

### GetCollectUnmanagedCertificates

`func (o *ComputerInventoryCollectionPreferencesV2) GetCollectUnmanagedCertificates() bool`

GetCollectUnmanagedCertificates returns the CollectUnmanagedCertificates field if non-nil, zero value otherwise.

### GetCollectUnmanagedCertificatesOk

`func (o *ComputerInventoryCollectionPreferencesV2) GetCollectUnmanagedCertificatesOk() (*bool, bool)`

GetCollectUnmanagedCertificatesOk returns a tuple with the CollectUnmanagedCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectUnmanagedCertificates

`func (o *ComputerInventoryCollectionPreferencesV2) SetCollectUnmanagedCertificates(v bool)`

SetCollectUnmanagedCertificates sets CollectUnmanagedCertificates field to given value.

### HasCollectUnmanagedCertificates

`func (o *ComputerInventoryCollectionPreferencesV2) HasCollectUnmanagedCertificates() bool`

HasCollectUnmanagedCertificates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


