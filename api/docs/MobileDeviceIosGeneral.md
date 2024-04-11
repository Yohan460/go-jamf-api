# MobileDeviceIosGeneral

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
**SharedIpad** | Pointer to **bool** |  | [optional] 
**DiagnosticAndUsageReportingEnabled** | Pointer to **bool** |  | [optional] 
**AppAnalyticsEnabled** | Pointer to **bool** |  | [optional] 
**ResidentUsers** | Pointer to **int64** |  | [optional] 
**QuotaSize** | Pointer to **int64** |  | [optional] 
**TemporarySessionOnly** | Pointer to **bool** |  | [optional] 
**TemporarySessionTimeout** | Pointer to **int64** |  | [optional] 
**UserSessionTimeout** | Pointer to **int64** |  | [optional] 
**SyncedToComputer** | Pointer to **int64** |  | [optional] 
**MaximumSharediPadUsersStored** | Pointer to **int64** |  | [optional] 
**LastBackupDate** | Pointer to **time.Time** |  | [optional] 
**DeviceLocatorServiceEnabled** | Pointer to **bool** |  | [optional] 
**DoNotDisturbEnabled** | Pointer to **bool** |  | [optional] 
**CloudBackupEnabled** | Pointer to **bool** |  | [optional] 
**LastCloudBackupDate** | Pointer to **time.Time** |  | [optional] 
**LocationServicesForSelfServiceMobileEnabled** | Pointer to **bool** |  | [optional] 
**ItunesStoreAccountActive** | Pointer to **bool** |  | [optional] 
**ExchangeDeviceId** | Pointer to **string** |  | [optional] 
**Tethered** | Pointer to **bool** |  | [optional] 

## Methods

### NewMobileDeviceIosGeneral

`func NewMobileDeviceIosGeneral() *MobileDeviceIosGeneral`

NewMobileDeviceIosGeneral instantiates a new MobileDeviceIosGeneral object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceIosGeneralWithDefaults

`func NewMobileDeviceIosGeneralWithDefaults() *MobileDeviceIosGeneral`

NewMobileDeviceIosGeneralWithDefaults instantiates a new MobileDeviceIosGeneral object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUdid

`func (o *MobileDeviceIosGeneral) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *MobileDeviceIosGeneral) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *MobileDeviceIosGeneral) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *MobileDeviceIosGeneral) HasUdid() bool`

HasUdid returns a boolean if a field has been set.

### GetDisplayName

`func (o *MobileDeviceIosGeneral) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MobileDeviceIosGeneral) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MobileDeviceIosGeneral) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MobileDeviceIosGeneral) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAssetTag

`func (o *MobileDeviceIosGeneral) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *MobileDeviceIosGeneral) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *MobileDeviceIosGeneral) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *MobileDeviceIosGeneral) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetSiteId

`func (o *MobileDeviceIosGeneral) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *MobileDeviceIosGeneral) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *MobileDeviceIosGeneral) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *MobileDeviceIosGeneral) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetLastInventoryUpdateDate

`func (o *MobileDeviceIosGeneral) GetLastInventoryUpdateDate() time.Time`

GetLastInventoryUpdateDate returns the LastInventoryUpdateDate field if non-nil, zero value otherwise.

### GetLastInventoryUpdateDateOk

`func (o *MobileDeviceIosGeneral) GetLastInventoryUpdateDateOk() (*time.Time, bool)`

GetLastInventoryUpdateDateOk returns a tuple with the LastInventoryUpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInventoryUpdateDate

`func (o *MobileDeviceIosGeneral) SetLastInventoryUpdateDate(v time.Time)`

SetLastInventoryUpdateDate sets LastInventoryUpdateDate field to given value.

### HasLastInventoryUpdateDate

`func (o *MobileDeviceIosGeneral) HasLastInventoryUpdateDate() bool`

HasLastInventoryUpdateDate returns a boolean if a field has been set.

### GetOsVersion

`func (o *MobileDeviceIosGeneral) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *MobileDeviceIosGeneral) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *MobileDeviceIosGeneral) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *MobileDeviceIosGeneral) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetOsRapidSecurityResponse

`func (o *MobileDeviceIosGeneral) GetOsRapidSecurityResponse() string`

GetOsRapidSecurityResponse returns the OsRapidSecurityResponse field if non-nil, zero value otherwise.

### GetOsRapidSecurityResponseOk

`func (o *MobileDeviceIosGeneral) GetOsRapidSecurityResponseOk() (*string, bool)`

GetOsRapidSecurityResponseOk returns a tuple with the OsRapidSecurityResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsRapidSecurityResponse

`func (o *MobileDeviceIosGeneral) SetOsRapidSecurityResponse(v string)`

SetOsRapidSecurityResponse sets OsRapidSecurityResponse field to given value.

### HasOsRapidSecurityResponse

`func (o *MobileDeviceIosGeneral) HasOsRapidSecurityResponse() bool`

HasOsRapidSecurityResponse returns a boolean if a field has been set.

### GetOsBuild

`func (o *MobileDeviceIosGeneral) GetOsBuild() string`

GetOsBuild returns the OsBuild field if non-nil, zero value otherwise.

### GetOsBuildOk

`func (o *MobileDeviceIosGeneral) GetOsBuildOk() (*string, bool)`

GetOsBuildOk returns a tuple with the OsBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsBuild

`func (o *MobileDeviceIosGeneral) SetOsBuild(v string)`

SetOsBuild sets OsBuild field to given value.

### HasOsBuild

`func (o *MobileDeviceIosGeneral) HasOsBuild() bool`

HasOsBuild returns a boolean if a field has been set.

### GetOsSupplementalBuildVersion

`func (o *MobileDeviceIosGeneral) GetOsSupplementalBuildVersion() string`

GetOsSupplementalBuildVersion returns the OsSupplementalBuildVersion field if non-nil, zero value otherwise.

### GetOsSupplementalBuildVersionOk

`func (o *MobileDeviceIosGeneral) GetOsSupplementalBuildVersionOk() (*string, bool)`

GetOsSupplementalBuildVersionOk returns a tuple with the OsSupplementalBuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsSupplementalBuildVersion

`func (o *MobileDeviceIosGeneral) SetOsSupplementalBuildVersion(v string)`

SetOsSupplementalBuildVersion sets OsSupplementalBuildVersion field to given value.

### HasOsSupplementalBuildVersion

`func (o *MobileDeviceIosGeneral) HasOsSupplementalBuildVersion() bool`

HasOsSupplementalBuildVersion returns a boolean if a field has been set.

### GetSoftwareUpdateDeviceId

`func (o *MobileDeviceIosGeneral) GetSoftwareUpdateDeviceId() string`

GetSoftwareUpdateDeviceId returns the SoftwareUpdateDeviceId field if non-nil, zero value otherwise.

### GetSoftwareUpdateDeviceIdOk

`func (o *MobileDeviceIosGeneral) GetSoftwareUpdateDeviceIdOk() (*string, bool)`

GetSoftwareUpdateDeviceIdOk returns a tuple with the SoftwareUpdateDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdateDeviceId

`func (o *MobileDeviceIosGeneral) SetSoftwareUpdateDeviceId(v string)`

SetSoftwareUpdateDeviceId sets SoftwareUpdateDeviceId field to given value.

### HasSoftwareUpdateDeviceId

`func (o *MobileDeviceIosGeneral) HasSoftwareUpdateDeviceId() bool`

HasSoftwareUpdateDeviceId returns a boolean if a field has been set.

### GetIpAddress

`func (o *MobileDeviceIosGeneral) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *MobileDeviceIosGeneral) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *MobileDeviceIosGeneral) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *MobileDeviceIosGeneral) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetManaged

`func (o *MobileDeviceIosGeneral) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *MobileDeviceIosGeneral) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *MobileDeviceIosGeneral) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *MobileDeviceIosGeneral) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetSupervised

`func (o *MobileDeviceIosGeneral) GetSupervised() bool`

GetSupervised returns the Supervised field if non-nil, zero value otherwise.

### GetSupervisedOk

`func (o *MobileDeviceIosGeneral) GetSupervisedOk() (*bool, bool)`

GetSupervisedOk returns a tuple with the Supervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervised

`func (o *MobileDeviceIosGeneral) SetSupervised(v bool)`

SetSupervised sets Supervised field to given value.

### HasSupervised

`func (o *MobileDeviceIosGeneral) HasSupervised() bool`

HasSupervised returns a boolean if a field has been set.

### GetDeviceOwnershipType

`func (o *MobileDeviceIosGeneral) GetDeviceOwnershipType() string`

GetDeviceOwnershipType returns the DeviceOwnershipType field if non-nil, zero value otherwise.

### GetDeviceOwnershipTypeOk

`func (o *MobileDeviceIosGeneral) GetDeviceOwnershipTypeOk() (*string, bool)`

GetDeviceOwnershipTypeOk returns a tuple with the DeviceOwnershipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOwnershipType

`func (o *MobileDeviceIosGeneral) SetDeviceOwnershipType(v string)`

SetDeviceOwnershipType sets DeviceOwnershipType field to given value.

### HasDeviceOwnershipType

`func (o *MobileDeviceIosGeneral) HasDeviceOwnershipType() bool`

HasDeviceOwnershipType returns a boolean if a field has been set.

### GetEnrollmentMethodPrestage

`func (o *MobileDeviceIosGeneral) GetEnrollmentMethodPrestage() EnrollmentMethodPrestage`

GetEnrollmentMethodPrestage returns the EnrollmentMethodPrestage field if non-nil, zero value otherwise.

### GetEnrollmentMethodPrestageOk

`func (o *MobileDeviceIosGeneral) GetEnrollmentMethodPrestageOk() (*EnrollmentMethodPrestage, bool)`

GetEnrollmentMethodPrestageOk returns a tuple with the EnrollmentMethodPrestage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentMethodPrestage

`func (o *MobileDeviceIosGeneral) SetEnrollmentMethodPrestage(v EnrollmentMethodPrestage)`

SetEnrollmentMethodPrestage sets EnrollmentMethodPrestage field to given value.

### HasEnrollmentMethodPrestage

`func (o *MobileDeviceIosGeneral) HasEnrollmentMethodPrestage() bool`

HasEnrollmentMethodPrestage returns a boolean if a field has been set.

### GetEnrollmentSessionTokenValid

`func (o *MobileDeviceIosGeneral) GetEnrollmentSessionTokenValid() bool`

GetEnrollmentSessionTokenValid returns the EnrollmentSessionTokenValid field if non-nil, zero value otherwise.

### GetEnrollmentSessionTokenValidOk

`func (o *MobileDeviceIosGeneral) GetEnrollmentSessionTokenValidOk() (*bool, bool)`

GetEnrollmentSessionTokenValidOk returns a tuple with the EnrollmentSessionTokenValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSessionTokenValid

`func (o *MobileDeviceIosGeneral) SetEnrollmentSessionTokenValid(v bool)`

SetEnrollmentSessionTokenValid sets EnrollmentSessionTokenValid field to given value.

### HasEnrollmentSessionTokenValid

`func (o *MobileDeviceIosGeneral) HasEnrollmentSessionTokenValid() bool`

HasEnrollmentSessionTokenValid returns a boolean if a field has been set.

### GetLastEnrolledDate

`func (o *MobileDeviceIosGeneral) GetLastEnrolledDate() time.Time`

GetLastEnrolledDate returns the LastEnrolledDate field if non-nil, zero value otherwise.

### GetLastEnrolledDateOk

`func (o *MobileDeviceIosGeneral) GetLastEnrolledDateOk() (*time.Time, bool)`

GetLastEnrolledDateOk returns a tuple with the LastEnrolledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEnrolledDate

`func (o *MobileDeviceIosGeneral) SetLastEnrolledDate(v time.Time)`

SetLastEnrolledDate sets LastEnrolledDate field to given value.

### HasLastEnrolledDate

`func (o *MobileDeviceIosGeneral) HasLastEnrolledDate() bool`

HasLastEnrolledDate returns a boolean if a field has been set.

### GetMdmProfileExpirationDate

`func (o *MobileDeviceIosGeneral) GetMdmProfileExpirationDate() time.Time`

GetMdmProfileExpirationDate returns the MdmProfileExpirationDate field if non-nil, zero value otherwise.

### GetMdmProfileExpirationDateOk

`func (o *MobileDeviceIosGeneral) GetMdmProfileExpirationDateOk() (*time.Time, bool)`

GetMdmProfileExpirationDateOk returns a tuple with the MdmProfileExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmProfileExpirationDate

`func (o *MobileDeviceIosGeneral) SetMdmProfileExpirationDate(v time.Time)`

SetMdmProfileExpirationDate sets MdmProfileExpirationDate field to given value.

### HasMdmProfileExpirationDate

`func (o *MobileDeviceIosGeneral) HasMdmProfileExpirationDate() bool`

HasMdmProfileExpirationDate returns a boolean if a field has been set.

### GetTimeZone

`func (o *MobileDeviceIosGeneral) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MobileDeviceIosGeneral) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MobileDeviceIosGeneral) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MobileDeviceIosGeneral) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetDeclarativeDeviceManagementEnabled

`func (o *MobileDeviceIosGeneral) GetDeclarativeDeviceManagementEnabled() bool`

GetDeclarativeDeviceManagementEnabled returns the DeclarativeDeviceManagementEnabled field if non-nil, zero value otherwise.

### GetDeclarativeDeviceManagementEnabledOk

`func (o *MobileDeviceIosGeneral) GetDeclarativeDeviceManagementEnabledOk() (*bool, bool)`

GetDeclarativeDeviceManagementEnabledOk returns a tuple with the DeclarativeDeviceManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarativeDeviceManagementEnabled

`func (o *MobileDeviceIosGeneral) SetDeclarativeDeviceManagementEnabled(v bool)`

SetDeclarativeDeviceManagementEnabled sets DeclarativeDeviceManagementEnabled field to given value.

### HasDeclarativeDeviceManagementEnabled

`func (o *MobileDeviceIosGeneral) HasDeclarativeDeviceManagementEnabled() bool`

HasDeclarativeDeviceManagementEnabled returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *MobileDeviceIosGeneral) GetExtensionAttributes() []MobileDeviceExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *MobileDeviceIosGeneral) GetExtensionAttributesOk() (*[]MobileDeviceExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *MobileDeviceIosGeneral) SetExtensionAttributes(v []MobileDeviceExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *MobileDeviceIosGeneral) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.

### GetSharedIpad

`func (o *MobileDeviceIosGeneral) GetSharedIpad() bool`

GetSharedIpad returns the SharedIpad field if non-nil, zero value otherwise.

### GetSharedIpadOk

`func (o *MobileDeviceIosGeneral) GetSharedIpadOk() (*bool, bool)`

GetSharedIpadOk returns a tuple with the SharedIpad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedIpad

`func (o *MobileDeviceIosGeneral) SetSharedIpad(v bool)`

SetSharedIpad sets SharedIpad field to given value.

### HasSharedIpad

`func (o *MobileDeviceIosGeneral) HasSharedIpad() bool`

HasSharedIpad returns a boolean if a field has been set.

### GetDiagnosticAndUsageReportingEnabled

`func (o *MobileDeviceIosGeneral) GetDiagnosticAndUsageReportingEnabled() bool`

GetDiagnosticAndUsageReportingEnabled returns the DiagnosticAndUsageReportingEnabled field if non-nil, zero value otherwise.

### GetDiagnosticAndUsageReportingEnabledOk

`func (o *MobileDeviceIosGeneral) GetDiagnosticAndUsageReportingEnabledOk() (*bool, bool)`

GetDiagnosticAndUsageReportingEnabledOk returns a tuple with the DiagnosticAndUsageReportingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticAndUsageReportingEnabled

`func (o *MobileDeviceIosGeneral) SetDiagnosticAndUsageReportingEnabled(v bool)`

SetDiagnosticAndUsageReportingEnabled sets DiagnosticAndUsageReportingEnabled field to given value.

### HasDiagnosticAndUsageReportingEnabled

`func (o *MobileDeviceIosGeneral) HasDiagnosticAndUsageReportingEnabled() bool`

HasDiagnosticAndUsageReportingEnabled returns a boolean if a field has been set.

### GetAppAnalyticsEnabled

`func (o *MobileDeviceIosGeneral) GetAppAnalyticsEnabled() bool`

GetAppAnalyticsEnabled returns the AppAnalyticsEnabled field if non-nil, zero value otherwise.

### GetAppAnalyticsEnabledOk

`func (o *MobileDeviceIosGeneral) GetAppAnalyticsEnabledOk() (*bool, bool)`

GetAppAnalyticsEnabledOk returns a tuple with the AppAnalyticsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAnalyticsEnabled

`func (o *MobileDeviceIosGeneral) SetAppAnalyticsEnabled(v bool)`

SetAppAnalyticsEnabled sets AppAnalyticsEnabled field to given value.

### HasAppAnalyticsEnabled

`func (o *MobileDeviceIosGeneral) HasAppAnalyticsEnabled() bool`

HasAppAnalyticsEnabled returns a boolean if a field has been set.

### GetResidentUsers

`func (o *MobileDeviceIosGeneral) GetResidentUsers() int64`

GetResidentUsers returns the ResidentUsers field if non-nil, zero value otherwise.

### GetResidentUsersOk

`func (o *MobileDeviceIosGeneral) GetResidentUsersOk() (*int64, bool)`

GetResidentUsersOk returns a tuple with the ResidentUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentUsers

`func (o *MobileDeviceIosGeneral) SetResidentUsers(v int64)`

SetResidentUsers sets ResidentUsers field to given value.

### HasResidentUsers

`func (o *MobileDeviceIosGeneral) HasResidentUsers() bool`

HasResidentUsers returns a boolean if a field has been set.

### GetQuotaSize

`func (o *MobileDeviceIosGeneral) GetQuotaSize() int64`

GetQuotaSize returns the QuotaSize field if non-nil, zero value otherwise.

### GetQuotaSizeOk

`func (o *MobileDeviceIosGeneral) GetQuotaSizeOk() (*int64, bool)`

GetQuotaSizeOk returns a tuple with the QuotaSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaSize

`func (o *MobileDeviceIosGeneral) SetQuotaSize(v int64)`

SetQuotaSize sets QuotaSize field to given value.

### HasQuotaSize

`func (o *MobileDeviceIosGeneral) HasQuotaSize() bool`

HasQuotaSize returns a boolean if a field has been set.

### GetTemporarySessionOnly

`func (o *MobileDeviceIosGeneral) GetTemporarySessionOnly() bool`

GetTemporarySessionOnly returns the TemporarySessionOnly field if non-nil, zero value otherwise.

### GetTemporarySessionOnlyOk

`func (o *MobileDeviceIosGeneral) GetTemporarySessionOnlyOk() (*bool, bool)`

GetTemporarySessionOnlyOk returns a tuple with the TemporarySessionOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporarySessionOnly

`func (o *MobileDeviceIosGeneral) SetTemporarySessionOnly(v bool)`

SetTemporarySessionOnly sets TemporarySessionOnly field to given value.

### HasTemporarySessionOnly

`func (o *MobileDeviceIosGeneral) HasTemporarySessionOnly() bool`

HasTemporarySessionOnly returns a boolean if a field has been set.

### GetTemporarySessionTimeout

`func (o *MobileDeviceIosGeneral) GetTemporarySessionTimeout() int64`

GetTemporarySessionTimeout returns the TemporarySessionTimeout field if non-nil, zero value otherwise.

### GetTemporarySessionTimeoutOk

`func (o *MobileDeviceIosGeneral) GetTemporarySessionTimeoutOk() (*int64, bool)`

GetTemporarySessionTimeoutOk returns a tuple with the TemporarySessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporarySessionTimeout

`func (o *MobileDeviceIosGeneral) SetTemporarySessionTimeout(v int64)`

SetTemporarySessionTimeout sets TemporarySessionTimeout field to given value.

### HasTemporarySessionTimeout

`func (o *MobileDeviceIosGeneral) HasTemporarySessionTimeout() bool`

HasTemporarySessionTimeout returns a boolean if a field has been set.

### GetUserSessionTimeout

`func (o *MobileDeviceIosGeneral) GetUserSessionTimeout() int64`

GetUserSessionTimeout returns the UserSessionTimeout field if non-nil, zero value otherwise.

### GetUserSessionTimeoutOk

`func (o *MobileDeviceIosGeneral) GetUserSessionTimeoutOk() (*int64, bool)`

GetUserSessionTimeoutOk returns a tuple with the UserSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSessionTimeout

`func (o *MobileDeviceIosGeneral) SetUserSessionTimeout(v int64)`

SetUserSessionTimeout sets UserSessionTimeout field to given value.

### HasUserSessionTimeout

`func (o *MobileDeviceIosGeneral) HasUserSessionTimeout() bool`

HasUserSessionTimeout returns a boolean if a field has been set.

### GetSyncedToComputer

`func (o *MobileDeviceIosGeneral) GetSyncedToComputer() int64`

GetSyncedToComputer returns the SyncedToComputer field if non-nil, zero value otherwise.

### GetSyncedToComputerOk

`func (o *MobileDeviceIosGeneral) GetSyncedToComputerOk() (*int64, bool)`

GetSyncedToComputerOk returns a tuple with the SyncedToComputer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedToComputer

`func (o *MobileDeviceIosGeneral) SetSyncedToComputer(v int64)`

SetSyncedToComputer sets SyncedToComputer field to given value.

### HasSyncedToComputer

`func (o *MobileDeviceIosGeneral) HasSyncedToComputer() bool`

HasSyncedToComputer returns a boolean if a field has been set.

### GetMaximumSharediPadUsersStored

`func (o *MobileDeviceIosGeneral) GetMaximumSharediPadUsersStored() int64`

GetMaximumSharediPadUsersStored returns the MaximumSharediPadUsersStored field if non-nil, zero value otherwise.

### GetMaximumSharediPadUsersStoredOk

`func (o *MobileDeviceIosGeneral) GetMaximumSharediPadUsersStoredOk() (*int64, bool)`

GetMaximumSharediPadUsersStoredOk returns a tuple with the MaximumSharediPadUsersStored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSharediPadUsersStored

`func (o *MobileDeviceIosGeneral) SetMaximumSharediPadUsersStored(v int64)`

SetMaximumSharediPadUsersStored sets MaximumSharediPadUsersStored field to given value.

### HasMaximumSharediPadUsersStored

`func (o *MobileDeviceIosGeneral) HasMaximumSharediPadUsersStored() bool`

HasMaximumSharediPadUsersStored returns a boolean if a field has been set.

### GetLastBackupDate

`func (o *MobileDeviceIosGeneral) GetLastBackupDate() time.Time`

GetLastBackupDate returns the LastBackupDate field if non-nil, zero value otherwise.

### GetLastBackupDateOk

`func (o *MobileDeviceIosGeneral) GetLastBackupDateOk() (*time.Time, bool)`

GetLastBackupDateOk returns a tuple with the LastBackupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackupDate

`func (o *MobileDeviceIosGeneral) SetLastBackupDate(v time.Time)`

SetLastBackupDate sets LastBackupDate field to given value.

### HasLastBackupDate

`func (o *MobileDeviceIosGeneral) HasLastBackupDate() bool`

HasLastBackupDate returns a boolean if a field has been set.

### GetDeviceLocatorServiceEnabled

`func (o *MobileDeviceIosGeneral) GetDeviceLocatorServiceEnabled() bool`

GetDeviceLocatorServiceEnabled returns the DeviceLocatorServiceEnabled field if non-nil, zero value otherwise.

### GetDeviceLocatorServiceEnabledOk

`func (o *MobileDeviceIosGeneral) GetDeviceLocatorServiceEnabledOk() (*bool, bool)`

GetDeviceLocatorServiceEnabledOk returns a tuple with the DeviceLocatorServiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLocatorServiceEnabled

`func (o *MobileDeviceIosGeneral) SetDeviceLocatorServiceEnabled(v bool)`

SetDeviceLocatorServiceEnabled sets DeviceLocatorServiceEnabled field to given value.

### HasDeviceLocatorServiceEnabled

`func (o *MobileDeviceIosGeneral) HasDeviceLocatorServiceEnabled() bool`

HasDeviceLocatorServiceEnabled returns a boolean if a field has been set.

### GetDoNotDisturbEnabled

`func (o *MobileDeviceIosGeneral) GetDoNotDisturbEnabled() bool`

GetDoNotDisturbEnabled returns the DoNotDisturbEnabled field if non-nil, zero value otherwise.

### GetDoNotDisturbEnabledOk

`func (o *MobileDeviceIosGeneral) GetDoNotDisturbEnabledOk() (*bool, bool)`

GetDoNotDisturbEnabledOk returns a tuple with the DoNotDisturbEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDisturbEnabled

`func (o *MobileDeviceIosGeneral) SetDoNotDisturbEnabled(v bool)`

SetDoNotDisturbEnabled sets DoNotDisturbEnabled field to given value.

### HasDoNotDisturbEnabled

`func (o *MobileDeviceIosGeneral) HasDoNotDisturbEnabled() bool`

HasDoNotDisturbEnabled returns a boolean if a field has been set.

### GetCloudBackupEnabled

`func (o *MobileDeviceIosGeneral) GetCloudBackupEnabled() bool`

GetCloudBackupEnabled returns the CloudBackupEnabled field if non-nil, zero value otherwise.

### GetCloudBackupEnabledOk

`func (o *MobileDeviceIosGeneral) GetCloudBackupEnabledOk() (*bool, bool)`

GetCloudBackupEnabledOk returns a tuple with the CloudBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudBackupEnabled

`func (o *MobileDeviceIosGeneral) SetCloudBackupEnabled(v bool)`

SetCloudBackupEnabled sets CloudBackupEnabled field to given value.

### HasCloudBackupEnabled

`func (o *MobileDeviceIosGeneral) HasCloudBackupEnabled() bool`

HasCloudBackupEnabled returns a boolean if a field has been set.

### GetLastCloudBackupDate

`func (o *MobileDeviceIosGeneral) GetLastCloudBackupDate() time.Time`

GetLastCloudBackupDate returns the LastCloudBackupDate field if non-nil, zero value otherwise.

### GetLastCloudBackupDateOk

`func (o *MobileDeviceIosGeneral) GetLastCloudBackupDateOk() (*time.Time, bool)`

GetLastCloudBackupDateOk returns a tuple with the LastCloudBackupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCloudBackupDate

`func (o *MobileDeviceIosGeneral) SetLastCloudBackupDate(v time.Time)`

SetLastCloudBackupDate sets LastCloudBackupDate field to given value.

### HasLastCloudBackupDate

`func (o *MobileDeviceIosGeneral) HasLastCloudBackupDate() bool`

HasLastCloudBackupDate returns a boolean if a field has been set.

### GetLocationServicesForSelfServiceMobileEnabled

`func (o *MobileDeviceIosGeneral) GetLocationServicesForSelfServiceMobileEnabled() bool`

GetLocationServicesForSelfServiceMobileEnabled returns the LocationServicesForSelfServiceMobileEnabled field if non-nil, zero value otherwise.

### GetLocationServicesForSelfServiceMobileEnabledOk

`func (o *MobileDeviceIosGeneral) GetLocationServicesForSelfServiceMobileEnabledOk() (*bool, bool)`

GetLocationServicesForSelfServiceMobileEnabledOk returns a tuple with the LocationServicesForSelfServiceMobileEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationServicesForSelfServiceMobileEnabled

`func (o *MobileDeviceIosGeneral) SetLocationServicesForSelfServiceMobileEnabled(v bool)`

SetLocationServicesForSelfServiceMobileEnabled sets LocationServicesForSelfServiceMobileEnabled field to given value.

### HasLocationServicesForSelfServiceMobileEnabled

`func (o *MobileDeviceIosGeneral) HasLocationServicesForSelfServiceMobileEnabled() bool`

HasLocationServicesForSelfServiceMobileEnabled returns a boolean if a field has been set.

### GetItunesStoreAccountActive

`func (o *MobileDeviceIosGeneral) GetItunesStoreAccountActive() bool`

GetItunesStoreAccountActive returns the ItunesStoreAccountActive field if non-nil, zero value otherwise.

### GetItunesStoreAccountActiveOk

`func (o *MobileDeviceIosGeneral) GetItunesStoreAccountActiveOk() (*bool, bool)`

GetItunesStoreAccountActiveOk returns a tuple with the ItunesStoreAccountActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItunesStoreAccountActive

`func (o *MobileDeviceIosGeneral) SetItunesStoreAccountActive(v bool)`

SetItunesStoreAccountActive sets ItunesStoreAccountActive field to given value.

### HasItunesStoreAccountActive

`func (o *MobileDeviceIosGeneral) HasItunesStoreAccountActive() bool`

HasItunesStoreAccountActive returns a boolean if a field has been set.

### GetExchangeDeviceId

`func (o *MobileDeviceIosGeneral) GetExchangeDeviceId() string`

GetExchangeDeviceId returns the ExchangeDeviceId field if non-nil, zero value otherwise.

### GetExchangeDeviceIdOk

`func (o *MobileDeviceIosGeneral) GetExchangeDeviceIdOk() (*string, bool)`

GetExchangeDeviceIdOk returns a tuple with the ExchangeDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeDeviceId

`func (o *MobileDeviceIosGeneral) SetExchangeDeviceId(v string)`

SetExchangeDeviceId sets ExchangeDeviceId field to given value.

### HasExchangeDeviceId

`func (o *MobileDeviceIosGeneral) HasExchangeDeviceId() bool`

HasExchangeDeviceId returns a boolean if a field has been set.

### GetTethered

`func (o *MobileDeviceIosGeneral) GetTethered() bool`

GetTethered returns the Tethered field if non-nil, zero value otherwise.

### GetTetheredOk

`func (o *MobileDeviceIosGeneral) GetTetheredOk() (*bool, bool)`

GetTetheredOk returns a tuple with the Tethered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTethered

`func (o *MobileDeviceIosGeneral) SetTethered(v bool)`

SetTethered sets Tethered field to given value.

### HasTethered

`func (o *MobileDeviceIosGeneral) HasTethered() bool`

HasTethered returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


