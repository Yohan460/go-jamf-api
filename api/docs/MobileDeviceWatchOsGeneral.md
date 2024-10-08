# MobileDeviceWatchOsGeneral

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Udid** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**LastInventoryUpdateDate** | Pointer to **time.Time** |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**OsRapidSecurityResponse** | Pointer to **string** |  | [optional] 
**OsBuild** | Pointer to **string** |  | [optional] 
**OsSupplementalBuildVersion** | Pointer to **string** |  | [optional] 
**SoftwareUpdateDeviceId** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 
**Supervised** | Pointer to **bool** |  | [optional] 
**DeviceOwnershipType** | Pointer to **string** |  | [optional] 
**EnrollmentMethodPrestage** | Pointer to [**EnrollmentMethodPrestage**](EnrollmentMethodPrestage.md) |  | [optional] 
**EnrollmentSessionTokenValid** | Pointer to **bool** |  | [optional] 
**LastEnrolledDate** | Pointer to **time.Time** |  | [optional] 
**MdmProfileExpirationDate** | Pointer to **time.Time** |  | [optional] 
**TimeZone** | Pointer to **string** | IANA time zone database name | [optional] 
**DeclarativeDeviceManagementEnabled** | Pointer to **bool** |  | [optional] 
**ExtensionAttributes** | Pointer to [**[]MobileDeviceExtensionAttribute**](MobileDeviceExtensionAttribute.md) |  | [optional] 
**DiagnosticAndUsageReportingEnabled** | Pointer to **bool** |  | [optional] 
**AppAnalyticsEnabled** | Pointer to **bool** |  | [optional] 
**DeviceLocatorServiceEnabled** | Pointer to **bool** |  | [optional] 
**DoNotDisturbEnabled** | Pointer to **bool** |  | [optional] 
**LastCloudBackupDate** | Pointer to **time.Time** |  | [optional] 
**ItunesStoreAccountActive** | Pointer to **bool** |  | [optional] 

## Methods

### NewMobileDeviceWatchOsGeneral

`func NewMobileDeviceWatchOsGeneral() *MobileDeviceWatchOsGeneral`

NewMobileDeviceWatchOsGeneral instantiates a new MobileDeviceWatchOsGeneral object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceWatchOsGeneralWithDefaults

`func NewMobileDeviceWatchOsGeneralWithDefaults() *MobileDeviceWatchOsGeneral`

NewMobileDeviceWatchOsGeneralWithDefaults instantiates a new MobileDeviceWatchOsGeneral object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUdid

`func (o *MobileDeviceWatchOsGeneral) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *MobileDeviceWatchOsGeneral) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *MobileDeviceWatchOsGeneral) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *MobileDeviceWatchOsGeneral) HasUdid() bool`

HasUdid returns a boolean if a field has been set.

### GetDisplayName

`func (o *MobileDeviceWatchOsGeneral) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MobileDeviceWatchOsGeneral) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MobileDeviceWatchOsGeneral) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MobileDeviceWatchOsGeneral) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAssetTag

`func (o *MobileDeviceWatchOsGeneral) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *MobileDeviceWatchOsGeneral) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *MobileDeviceWatchOsGeneral) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *MobileDeviceWatchOsGeneral) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetSiteId

`func (o *MobileDeviceWatchOsGeneral) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *MobileDeviceWatchOsGeneral) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *MobileDeviceWatchOsGeneral) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *MobileDeviceWatchOsGeneral) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetLastInventoryUpdateDate

`func (o *MobileDeviceWatchOsGeneral) GetLastInventoryUpdateDate() time.Time`

GetLastInventoryUpdateDate returns the LastInventoryUpdateDate field if non-nil, zero value otherwise.

### GetLastInventoryUpdateDateOk

`func (o *MobileDeviceWatchOsGeneral) GetLastInventoryUpdateDateOk() (*time.Time, bool)`

GetLastInventoryUpdateDateOk returns a tuple with the LastInventoryUpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInventoryUpdateDate

`func (o *MobileDeviceWatchOsGeneral) SetLastInventoryUpdateDate(v time.Time)`

SetLastInventoryUpdateDate sets LastInventoryUpdateDate field to given value.

### HasLastInventoryUpdateDate

`func (o *MobileDeviceWatchOsGeneral) HasLastInventoryUpdateDate() bool`

HasLastInventoryUpdateDate returns a boolean if a field has been set.

### GetOsVersion

`func (o *MobileDeviceWatchOsGeneral) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *MobileDeviceWatchOsGeneral) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *MobileDeviceWatchOsGeneral) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *MobileDeviceWatchOsGeneral) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetOsRapidSecurityResponse

`func (o *MobileDeviceWatchOsGeneral) GetOsRapidSecurityResponse() string`

GetOsRapidSecurityResponse returns the OsRapidSecurityResponse field if non-nil, zero value otherwise.

### GetOsRapidSecurityResponseOk

`func (o *MobileDeviceWatchOsGeneral) GetOsRapidSecurityResponseOk() (*string, bool)`

GetOsRapidSecurityResponseOk returns a tuple with the OsRapidSecurityResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsRapidSecurityResponse

`func (o *MobileDeviceWatchOsGeneral) SetOsRapidSecurityResponse(v string)`

SetOsRapidSecurityResponse sets OsRapidSecurityResponse field to given value.

### HasOsRapidSecurityResponse

`func (o *MobileDeviceWatchOsGeneral) HasOsRapidSecurityResponse() bool`

HasOsRapidSecurityResponse returns a boolean if a field has been set.

### GetOsBuild

`func (o *MobileDeviceWatchOsGeneral) GetOsBuild() string`

GetOsBuild returns the OsBuild field if non-nil, zero value otherwise.

### GetOsBuildOk

`func (o *MobileDeviceWatchOsGeneral) GetOsBuildOk() (*string, bool)`

GetOsBuildOk returns a tuple with the OsBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsBuild

`func (o *MobileDeviceWatchOsGeneral) SetOsBuild(v string)`

SetOsBuild sets OsBuild field to given value.

### HasOsBuild

`func (o *MobileDeviceWatchOsGeneral) HasOsBuild() bool`

HasOsBuild returns a boolean if a field has been set.

### GetOsSupplementalBuildVersion

`func (o *MobileDeviceWatchOsGeneral) GetOsSupplementalBuildVersion() string`

GetOsSupplementalBuildVersion returns the OsSupplementalBuildVersion field if non-nil, zero value otherwise.

### GetOsSupplementalBuildVersionOk

`func (o *MobileDeviceWatchOsGeneral) GetOsSupplementalBuildVersionOk() (*string, bool)`

GetOsSupplementalBuildVersionOk returns a tuple with the OsSupplementalBuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsSupplementalBuildVersion

`func (o *MobileDeviceWatchOsGeneral) SetOsSupplementalBuildVersion(v string)`

SetOsSupplementalBuildVersion sets OsSupplementalBuildVersion field to given value.

### HasOsSupplementalBuildVersion

`func (o *MobileDeviceWatchOsGeneral) HasOsSupplementalBuildVersion() bool`

HasOsSupplementalBuildVersion returns a boolean if a field has been set.

### GetSoftwareUpdateDeviceId

`func (o *MobileDeviceWatchOsGeneral) GetSoftwareUpdateDeviceId() string`

GetSoftwareUpdateDeviceId returns the SoftwareUpdateDeviceId field if non-nil, zero value otherwise.

### GetSoftwareUpdateDeviceIdOk

`func (o *MobileDeviceWatchOsGeneral) GetSoftwareUpdateDeviceIdOk() (*string, bool)`

GetSoftwareUpdateDeviceIdOk returns a tuple with the SoftwareUpdateDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdateDeviceId

`func (o *MobileDeviceWatchOsGeneral) SetSoftwareUpdateDeviceId(v string)`

SetSoftwareUpdateDeviceId sets SoftwareUpdateDeviceId field to given value.

### HasSoftwareUpdateDeviceId

`func (o *MobileDeviceWatchOsGeneral) HasSoftwareUpdateDeviceId() bool`

HasSoftwareUpdateDeviceId returns a boolean if a field has been set.

### GetIpAddress

`func (o *MobileDeviceWatchOsGeneral) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *MobileDeviceWatchOsGeneral) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *MobileDeviceWatchOsGeneral) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *MobileDeviceWatchOsGeneral) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetManaged

`func (o *MobileDeviceWatchOsGeneral) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *MobileDeviceWatchOsGeneral) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *MobileDeviceWatchOsGeneral) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *MobileDeviceWatchOsGeneral) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetSupervised

`func (o *MobileDeviceWatchOsGeneral) GetSupervised() bool`

GetSupervised returns the Supervised field if non-nil, zero value otherwise.

### GetSupervisedOk

`func (o *MobileDeviceWatchOsGeneral) GetSupervisedOk() (*bool, bool)`

GetSupervisedOk returns a tuple with the Supervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervised

`func (o *MobileDeviceWatchOsGeneral) SetSupervised(v bool)`

SetSupervised sets Supervised field to given value.

### HasSupervised

`func (o *MobileDeviceWatchOsGeneral) HasSupervised() bool`

HasSupervised returns a boolean if a field has been set.

### GetDeviceOwnershipType

`func (o *MobileDeviceWatchOsGeneral) GetDeviceOwnershipType() string`

GetDeviceOwnershipType returns the DeviceOwnershipType field if non-nil, zero value otherwise.

### GetDeviceOwnershipTypeOk

`func (o *MobileDeviceWatchOsGeneral) GetDeviceOwnershipTypeOk() (*string, bool)`

GetDeviceOwnershipTypeOk returns a tuple with the DeviceOwnershipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOwnershipType

`func (o *MobileDeviceWatchOsGeneral) SetDeviceOwnershipType(v string)`

SetDeviceOwnershipType sets DeviceOwnershipType field to given value.

### HasDeviceOwnershipType

`func (o *MobileDeviceWatchOsGeneral) HasDeviceOwnershipType() bool`

HasDeviceOwnershipType returns a boolean if a field has been set.

### GetEnrollmentMethodPrestage

`func (o *MobileDeviceWatchOsGeneral) GetEnrollmentMethodPrestage() EnrollmentMethodPrestage`

GetEnrollmentMethodPrestage returns the EnrollmentMethodPrestage field if non-nil, zero value otherwise.

### GetEnrollmentMethodPrestageOk

`func (o *MobileDeviceWatchOsGeneral) GetEnrollmentMethodPrestageOk() (*EnrollmentMethodPrestage, bool)`

GetEnrollmentMethodPrestageOk returns a tuple with the EnrollmentMethodPrestage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentMethodPrestage

`func (o *MobileDeviceWatchOsGeneral) SetEnrollmentMethodPrestage(v EnrollmentMethodPrestage)`

SetEnrollmentMethodPrestage sets EnrollmentMethodPrestage field to given value.

### HasEnrollmentMethodPrestage

`func (o *MobileDeviceWatchOsGeneral) HasEnrollmentMethodPrestage() bool`

HasEnrollmentMethodPrestage returns a boolean if a field has been set.

### GetEnrollmentSessionTokenValid

`func (o *MobileDeviceWatchOsGeneral) GetEnrollmentSessionTokenValid() bool`

GetEnrollmentSessionTokenValid returns the EnrollmentSessionTokenValid field if non-nil, zero value otherwise.

### GetEnrollmentSessionTokenValidOk

`func (o *MobileDeviceWatchOsGeneral) GetEnrollmentSessionTokenValidOk() (*bool, bool)`

GetEnrollmentSessionTokenValidOk returns a tuple with the EnrollmentSessionTokenValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSessionTokenValid

`func (o *MobileDeviceWatchOsGeneral) SetEnrollmentSessionTokenValid(v bool)`

SetEnrollmentSessionTokenValid sets EnrollmentSessionTokenValid field to given value.

### HasEnrollmentSessionTokenValid

`func (o *MobileDeviceWatchOsGeneral) HasEnrollmentSessionTokenValid() bool`

HasEnrollmentSessionTokenValid returns a boolean if a field has been set.

### GetLastEnrolledDate

`func (o *MobileDeviceWatchOsGeneral) GetLastEnrolledDate() time.Time`

GetLastEnrolledDate returns the LastEnrolledDate field if non-nil, zero value otherwise.

### GetLastEnrolledDateOk

`func (o *MobileDeviceWatchOsGeneral) GetLastEnrolledDateOk() (*time.Time, bool)`

GetLastEnrolledDateOk returns a tuple with the LastEnrolledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEnrolledDate

`func (o *MobileDeviceWatchOsGeneral) SetLastEnrolledDate(v time.Time)`

SetLastEnrolledDate sets LastEnrolledDate field to given value.

### HasLastEnrolledDate

`func (o *MobileDeviceWatchOsGeneral) HasLastEnrolledDate() bool`

HasLastEnrolledDate returns a boolean if a field has been set.

### GetMdmProfileExpirationDate

`func (o *MobileDeviceWatchOsGeneral) GetMdmProfileExpirationDate() time.Time`

GetMdmProfileExpirationDate returns the MdmProfileExpirationDate field if non-nil, zero value otherwise.

### GetMdmProfileExpirationDateOk

`func (o *MobileDeviceWatchOsGeneral) GetMdmProfileExpirationDateOk() (*time.Time, bool)`

GetMdmProfileExpirationDateOk returns a tuple with the MdmProfileExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmProfileExpirationDate

`func (o *MobileDeviceWatchOsGeneral) SetMdmProfileExpirationDate(v time.Time)`

SetMdmProfileExpirationDate sets MdmProfileExpirationDate field to given value.

### HasMdmProfileExpirationDate

`func (o *MobileDeviceWatchOsGeneral) HasMdmProfileExpirationDate() bool`

HasMdmProfileExpirationDate returns a boolean if a field has been set.

### GetTimeZone

`func (o *MobileDeviceWatchOsGeneral) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MobileDeviceWatchOsGeneral) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MobileDeviceWatchOsGeneral) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MobileDeviceWatchOsGeneral) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetDeclarativeDeviceManagementEnabled

`func (o *MobileDeviceWatchOsGeneral) GetDeclarativeDeviceManagementEnabled() bool`

GetDeclarativeDeviceManagementEnabled returns the DeclarativeDeviceManagementEnabled field if non-nil, zero value otherwise.

### GetDeclarativeDeviceManagementEnabledOk

`func (o *MobileDeviceWatchOsGeneral) GetDeclarativeDeviceManagementEnabledOk() (*bool, bool)`

GetDeclarativeDeviceManagementEnabledOk returns a tuple with the DeclarativeDeviceManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarativeDeviceManagementEnabled

`func (o *MobileDeviceWatchOsGeneral) SetDeclarativeDeviceManagementEnabled(v bool)`

SetDeclarativeDeviceManagementEnabled sets DeclarativeDeviceManagementEnabled field to given value.

### HasDeclarativeDeviceManagementEnabled

`func (o *MobileDeviceWatchOsGeneral) HasDeclarativeDeviceManagementEnabled() bool`

HasDeclarativeDeviceManagementEnabled returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *MobileDeviceWatchOsGeneral) GetExtensionAttributes() []MobileDeviceExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *MobileDeviceWatchOsGeneral) GetExtensionAttributesOk() (*[]MobileDeviceExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *MobileDeviceWatchOsGeneral) SetExtensionAttributes(v []MobileDeviceExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *MobileDeviceWatchOsGeneral) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.

### GetDiagnosticAndUsageReportingEnabled

`func (o *MobileDeviceWatchOsGeneral) GetDiagnosticAndUsageReportingEnabled() bool`

GetDiagnosticAndUsageReportingEnabled returns the DiagnosticAndUsageReportingEnabled field if non-nil, zero value otherwise.

### GetDiagnosticAndUsageReportingEnabledOk

`func (o *MobileDeviceWatchOsGeneral) GetDiagnosticAndUsageReportingEnabledOk() (*bool, bool)`

GetDiagnosticAndUsageReportingEnabledOk returns a tuple with the DiagnosticAndUsageReportingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticAndUsageReportingEnabled

`func (o *MobileDeviceWatchOsGeneral) SetDiagnosticAndUsageReportingEnabled(v bool)`

SetDiagnosticAndUsageReportingEnabled sets DiagnosticAndUsageReportingEnabled field to given value.

### HasDiagnosticAndUsageReportingEnabled

`func (o *MobileDeviceWatchOsGeneral) HasDiagnosticAndUsageReportingEnabled() bool`

HasDiagnosticAndUsageReportingEnabled returns a boolean if a field has been set.

### GetAppAnalyticsEnabled

`func (o *MobileDeviceWatchOsGeneral) GetAppAnalyticsEnabled() bool`

GetAppAnalyticsEnabled returns the AppAnalyticsEnabled field if non-nil, zero value otherwise.

### GetAppAnalyticsEnabledOk

`func (o *MobileDeviceWatchOsGeneral) GetAppAnalyticsEnabledOk() (*bool, bool)`

GetAppAnalyticsEnabledOk returns a tuple with the AppAnalyticsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAnalyticsEnabled

`func (o *MobileDeviceWatchOsGeneral) SetAppAnalyticsEnabled(v bool)`

SetAppAnalyticsEnabled sets AppAnalyticsEnabled field to given value.

### HasAppAnalyticsEnabled

`func (o *MobileDeviceWatchOsGeneral) HasAppAnalyticsEnabled() bool`

HasAppAnalyticsEnabled returns a boolean if a field has been set.

### GetDeviceLocatorServiceEnabled

`func (o *MobileDeviceWatchOsGeneral) GetDeviceLocatorServiceEnabled() bool`

GetDeviceLocatorServiceEnabled returns the DeviceLocatorServiceEnabled field if non-nil, zero value otherwise.

### GetDeviceLocatorServiceEnabledOk

`func (o *MobileDeviceWatchOsGeneral) GetDeviceLocatorServiceEnabledOk() (*bool, bool)`

GetDeviceLocatorServiceEnabledOk returns a tuple with the DeviceLocatorServiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLocatorServiceEnabled

`func (o *MobileDeviceWatchOsGeneral) SetDeviceLocatorServiceEnabled(v bool)`

SetDeviceLocatorServiceEnabled sets DeviceLocatorServiceEnabled field to given value.

### HasDeviceLocatorServiceEnabled

`func (o *MobileDeviceWatchOsGeneral) HasDeviceLocatorServiceEnabled() bool`

HasDeviceLocatorServiceEnabled returns a boolean if a field has been set.

### GetDoNotDisturbEnabled

`func (o *MobileDeviceWatchOsGeneral) GetDoNotDisturbEnabled() bool`

GetDoNotDisturbEnabled returns the DoNotDisturbEnabled field if non-nil, zero value otherwise.

### GetDoNotDisturbEnabledOk

`func (o *MobileDeviceWatchOsGeneral) GetDoNotDisturbEnabledOk() (*bool, bool)`

GetDoNotDisturbEnabledOk returns a tuple with the DoNotDisturbEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDisturbEnabled

`func (o *MobileDeviceWatchOsGeneral) SetDoNotDisturbEnabled(v bool)`

SetDoNotDisturbEnabled sets DoNotDisturbEnabled field to given value.

### HasDoNotDisturbEnabled

`func (o *MobileDeviceWatchOsGeneral) HasDoNotDisturbEnabled() bool`

HasDoNotDisturbEnabled returns a boolean if a field has been set.

### GetLastCloudBackupDate

`func (o *MobileDeviceWatchOsGeneral) GetLastCloudBackupDate() time.Time`

GetLastCloudBackupDate returns the LastCloudBackupDate field if non-nil, zero value otherwise.

### GetLastCloudBackupDateOk

`func (o *MobileDeviceWatchOsGeneral) GetLastCloudBackupDateOk() (*time.Time, bool)`

GetLastCloudBackupDateOk returns a tuple with the LastCloudBackupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCloudBackupDate

`func (o *MobileDeviceWatchOsGeneral) SetLastCloudBackupDate(v time.Time)`

SetLastCloudBackupDate sets LastCloudBackupDate field to given value.

### HasLastCloudBackupDate

`func (o *MobileDeviceWatchOsGeneral) HasLastCloudBackupDate() bool`

HasLastCloudBackupDate returns a boolean if a field has been set.

### GetItunesStoreAccountActive

`func (o *MobileDeviceWatchOsGeneral) GetItunesStoreAccountActive() bool`

GetItunesStoreAccountActive returns the ItunesStoreAccountActive field if non-nil, zero value otherwise.

### GetItunesStoreAccountActiveOk

`func (o *MobileDeviceWatchOsGeneral) GetItunesStoreAccountActiveOk() (*bool, bool)`

GetItunesStoreAccountActiveOk returns a tuple with the ItunesStoreAccountActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItunesStoreAccountActive

`func (o *MobileDeviceWatchOsGeneral) SetItunesStoreAccountActive(v bool)`

SetItunesStoreAccountActive sets ItunesStoreAccountActive field to given value.

### HasItunesStoreAccountActive

`func (o *MobileDeviceWatchOsGeneral) HasItunesStoreAccountActive() bool`

HasItunesStoreAccountActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


