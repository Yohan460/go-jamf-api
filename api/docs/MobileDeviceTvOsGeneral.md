# MobileDeviceTvOsGeneral

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
**AirPlayPassword** | Pointer to **string** |  | [optional] 
**Locales** | Pointer to **string** |  | [optional] 
**Languages** | Pointer to **string** |  | [optional] 

## Methods

### NewMobileDeviceTvOsGeneral

`func NewMobileDeviceTvOsGeneral() *MobileDeviceTvOsGeneral`

NewMobileDeviceTvOsGeneral instantiates a new MobileDeviceTvOsGeneral object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceTvOsGeneralWithDefaults

`func NewMobileDeviceTvOsGeneralWithDefaults() *MobileDeviceTvOsGeneral`

NewMobileDeviceTvOsGeneralWithDefaults instantiates a new MobileDeviceTvOsGeneral object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUdid

`func (o *MobileDeviceTvOsGeneral) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *MobileDeviceTvOsGeneral) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *MobileDeviceTvOsGeneral) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *MobileDeviceTvOsGeneral) HasUdid() bool`

HasUdid returns a boolean if a field has been set.

### GetDisplayName

`func (o *MobileDeviceTvOsGeneral) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MobileDeviceTvOsGeneral) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MobileDeviceTvOsGeneral) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MobileDeviceTvOsGeneral) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAssetTag

`func (o *MobileDeviceTvOsGeneral) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *MobileDeviceTvOsGeneral) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *MobileDeviceTvOsGeneral) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *MobileDeviceTvOsGeneral) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetSiteId

`func (o *MobileDeviceTvOsGeneral) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *MobileDeviceTvOsGeneral) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *MobileDeviceTvOsGeneral) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *MobileDeviceTvOsGeneral) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetLastInventoryUpdateDate

`func (o *MobileDeviceTvOsGeneral) GetLastInventoryUpdateDate() time.Time`

GetLastInventoryUpdateDate returns the LastInventoryUpdateDate field if non-nil, zero value otherwise.

### GetLastInventoryUpdateDateOk

`func (o *MobileDeviceTvOsGeneral) GetLastInventoryUpdateDateOk() (*time.Time, bool)`

GetLastInventoryUpdateDateOk returns a tuple with the LastInventoryUpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInventoryUpdateDate

`func (o *MobileDeviceTvOsGeneral) SetLastInventoryUpdateDate(v time.Time)`

SetLastInventoryUpdateDate sets LastInventoryUpdateDate field to given value.

### HasLastInventoryUpdateDate

`func (o *MobileDeviceTvOsGeneral) HasLastInventoryUpdateDate() bool`

HasLastInventoryUpdateDate returns a boolean if a field has been set.

### GetOsVersion

`func (o *MobileDeviceTvOsGeneral) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *MobileDeviceTvOsGeneral) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *MobileDeviceTvOsGeneral) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *MobileDeviceTvOsGeneral) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetOsRapidSecurityResponse

`func (o *MobileDeviceTvOsGeneral) GetOsRapidSecurityResponse() string`

GetOsRapidSecurityResponse returns the OsRapidSecurityResponse field if non-nil, zero value otherwise.

### GetOsRapidSecurityResponseOk

`func (o *MobileDeviceTvOsGeneral) GetOsRapidSecurityResponseOk() (*string, bool)`

GetOsRapidSecurityResponseOk returns a tuple with the OsRapidSecurityResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsRapidSecurityResponse

`func (o *MobileDeviceTvOsGeneral) SetOsRapidSecurityResponse(v string)`

SetOsRapidSecurityResponse sets OsRapidSecurityResponse field to given value.

### HasOsRapidSecurityResponse

`func (o *MobileDeviceTvOsGeneral) HasOsRapidSecurityResponse() bool`

HasOsRapidSecurityResponse returns a boolean if a field has been set.

### GetOsBuild

`func (o *MobileDeviceTvOsGeneral) GetOsBuild() string`

GetOsBuild returns the OsBuild field if non-nil, zero value otherwise.

### GetOsBuildOk

`func (o *MobileDeviceTvOsGeneral) GetOsBuildOk() (*string, bool)`

GetOsBuildOk returns a tuple with the OsBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsBuild

`func (o *MobileDeviceTvOsGeneral) SetOsBuild(v string)`

SetOsBuild sets OsBuild field to given value.

### HasOsBuild

`func (o *MobileDeviceTvOsGeneral) HasOsBuild() bool`

HasOsBuild returns a boolean if a field has been set.

### GetOsSupplementalBuildVersion

`func (o *MobileDeviceTvOsGeneral) GetOsSupplementalBuildVersion() string`

GetOsSupplementalBuildVersion returns the OsSupplementalBuildVersion field if non-nil, zero value otherwise.

### GetOsSupplementalBuildVersionOk

`func (o *MobileDeviceTvOsGeneral) GetOsSupplementalBuildVersionOk() (*string, bool)`

GetOsSupplementalBuildVersionOk returns a tuple with the OsSupplementalBuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsSupplementalBuildVersion

`func (o *MobileDeviceTvOsGeneral) SetOsSupplementalBuildVersion(v string)`

SetOsSupplementalBuildVersion sets OsSupplementalBuildVersion field to given value.

### HasOsSupplementalBuildVersion

`func (o *MobileDeviceTvOsGeneral) HasOsSupplementalBuildVersion() bool`

HasOsSupplementalBuildVersion returns a boolean if a field has been set.

### GetSoftwareUpdateDeviceId

`func (o *MobileDeviceTvOsGeneral) GetSoftwareUpdateDeviceId() string`

GetSoftwareUpdateDeviceId returns the SoftwareUpdateDeviceId field if non-nil, zero value otherwise.

### GetSoftwareUpdateDeviceIdOk

`func (o *MobileDeviceTvOsGeneral) GetSoftwareUpdateDeviceIdOk() (*string, bool)`

GetSoftwareUpdateDeviceIdOk returns a tuple with the SoftwareUpdateDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdateDeviceId

`func (o *MobileDeviceTvOsGeneral) SetSoftwareUpdateDeviceId(v string)`

SetSoftwareUpdateDeviceId sets SoftwareUpdateDeviceId field to given value.

### HasSoftwareUpdateDeviceId

`func (o *MobileDeviceTvOsGeneral) HasSoftwareUpdateDeviceId() bool`

HasSoftwareUpdateDeviceId returns a boolean if a field has been set.

### GetIpAddress

`func (o *MobileDeviceTvOsGeneral) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *MobileDeviceTvOsGeneral) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *MobileDeviceTvOsGeneral) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *MobileDeviceTvOsGeneral) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetManaged

`func (o *MobileDeviceTvOsGeneral) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *MobileDeviceTvOsGeneral) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *MobileDeviceTvOsGeneral) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *MobileDeviceTvOsGeneral) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetSupervised

`func (o *MobileDeviceTvOsGeneral) GetSupervised() bool`

GetSupervised returns the Supervised field if non-nil, zero value otherwise.

### GetSupervisedOk

`func (o *MobileDeviceTvOsGeneral) GetSupervisedOk() (*bool, bool)`

GetSupervisedOk returns a tuple with the Supervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervised

`func (o *MobileDeviceTvOsGeneral) SetSupervised(v bool)`

SetSupervised sets Supervised field to given value.

### HasSupervised

`func (o *MobileDeviceTvOsGeneral) HasSupervised() bool`

HasSupervised returns a boolean if a field has been set.

### GetDeviceOwnershipType

`func (o *MobileDeviceTvOsGeneral) GetDeviceOwnershipType() string`

GetDeviceOwnershipType returns the DeviceOwnershipType field if non-nil, zero value otherwise.

### GetDeviceOwnershipTypeOk

`func (o *MobileDeviceTvOsGeneral) GetDeviceOwnershipTypeOk() (*string, bool)`

GetDeviceOwnershipTypeOk returns a tuple with the DeviceOwnershipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOwnershipType

`func (o *MobileDeviceTvOsGeneral) SetDeviceOwnershipType(v string)`

SetDeviceOwnershipType sets DeviceOwnershipType field to given value.

### HasDeviceOwnershipType

`func (o *MobileDeviceTvOsGeneral) HasDeviceOwnershipType() bool`

HasDeviceOwnershipType returns a boolean if a field has been set.

### GetEnrollmentMethodPrestage

`func (o *MobileDeviceTvOsGeneral) GetEnrollmentMethodPrestage() EnrollmentMethodPrestage`

GetEnrollmentMethodPrestage returns the EnrollmentMethodPrestage field if non-nil, zero value otherwise.

### GetEnrollmentMethodPrestageOk

`func (o *MobileDeviceTvOsGeneral) GetEnrollmentMethodPrestageOk() (*EnrollmentMethodPrestage, bool)`

GetEnrollmentMethodPrestageOk returns a tuple with the EnrollmentMethodPrestage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentMethodPrestage

`func (o *MobileDeviceTvOsGeneral) SetEnrollmentMethodPrestage(v EnrollmentMethodPrestage)`

SetEnrollmentMethodPrestage sets EnrollmentMethodPrestage field to given value.

### HasEnrollmentMethodPrestage

`func (o *MobileDeviceTvOsGeneral) HasEnrollmentMethodPrestage() bool`

HasEnrollmentMethodPrestage returns a boolean if a field has been set.

### GetEnrollmentSessionTokenValid

`func (o *MobileDeviceTvOsGeneral) GetEnrollmentSessionTokenValid() bool`

GetEnrollmentSessionTokenValid returns the EnrollmentSessionTokenValid field if non-nil, zero value otherwise.

### GetEnrollmentSessionTokenValidOk

`func (o *MobileDeviceTvOsGeneral) GetEnrollmentSessionTokenValidOk() (*bool, bool)`

GetEnrollmentSessionTokenValidOk returns a tuple with the EnrollmentSessionTokenValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSessionTokenValid

`func (o *MobileDeviceTvOsGeneral) SetEnrollmentSessionTokenValid(v bool)`

SetEnrollmentSessionTokenValid sets EnrollmentSessionTokenValid field to given value.

### HasEnrollmentSessionTokenValid

`func (o *MobileDeviceTvOsGeneral) HasEnrollmentSessionTokenValid() bool`

HasEnrollmentSessionTokenValid returns a boolean if a field has been set.

### GetLastEnrolledDate

`func (o *MobileDeviceTvOsGeneral) GetLastEnrolledDate() time.Time`

GetLastEnrolledDate returns the LastEnrolledDate field if non-nil, zero value otherwise.

### GetLastEnrolledDateOk

`func (o *MobileDeviceTvOsGeneral) GetLastEnrolledDateOk() (*time.Time, bool)`

GetLastEnrolledDateOk returns a tuple with the LastEnrolledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEnrolledDate

`func (o *MobileDeviceTvOsGeneral) SetLastEnrolledDate(v time.Time)`

SetLastEnrolledDate sets LastEnrolledDate field to given value.

### HasLastEnrolledDate

`func (o *MobileDeviceTvOsGeneral) HasLastEnrolledDate() bool`

HasLastEnrolledDate returns a boolean if a field has been set.

### GetMdmProfileExpirationDate

`func (o *MobileDeviceTvOsGeneral) GetMdmProfileExpirationDate() time.Time`

GetMdmProfileExpirationDate returns the MdmProfileExpirationDate field if non-nil, zero value otherwise.

### GetMdmProfileExpirationDateOk

`func (o *MobileDeviceTvOsGeneral) GetMdmProfileExpirationDateOk() (*time.Time, bool)`

GetMdmProfileExpirationDateOk returns a tuple with the MdmProfileExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmProfileExpirationDate

`func (o *MobileDeviceTvOsGeneral) SetMdmProfileExpirationDate(v time.Time)`

SetMdmProfileExpirationDate sets MdmProfileExpirationDate field to given value.

### HasMdmProfileExpirationDate

`func (o *MobileDeviceTvOsGeneral) HasMdmProfileExpirationDate() bool`

HasMdmProfileExpirationDate returns a boolean if a field has been set.

### GetTimeZone

`func (o *MobileDeviceTvOsGeneral) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MobileDeviceTvOsGeneral) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MobileDeviceTvOsGeneral) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MobileDeviceTvOsGeneral) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetDeclarativeDeviceManagementEnabled

`func (o *MobileDeviceTvOsGeneral) GetDeclarativeDeviceManagementEnabled() bool`

GetDeclarativeDeviceManagementEnabled returns the DeclarativeDeviceManagementEnabled field if non-nil, zero value otherwise.

### GetDeclarativeDeviceManagementEnabledOk

`func (o *MobileDeviceTvOsGeneral) GetDeclarativeDeviceManagementEnabledOk() (*bool, bool)`

GetDeclarativeDeviceManagementEnabledOk returns a tuple with the DeclarativeDeviceManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarativeDeviceManagementEnabled

`func (o *MobileDeviceTvOsGeneral) SetDeclarativeDeviceManagementEnabled(v bool)`

SetDeclarativeDeviceManagementEnabled sets DeclarativeDeviceManagementEnabled field to given value.

### HasDeclarativeDeviceManagementEnabled

`func (o *MobileDeviceTvOsGeneral) HasDeclarativeDeviceManagementEnabled() bool`

HasDeclarativeDeviceManagementEnabled returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *MobileDeviceTvOsGeneral) GetExtensionAttributes() []MobileDeviceExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *MobileDeviceTvOsGeneral) GetExtensionAttributesOk() (*[]MobileDeviceExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *MobileDeviceTvOsGeneral) SetExtensionAttributes(v []MobileDeviceExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *MobileDeviceTvOsGeneral) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.

### GetAirPlayPassword

`func (o *MobileDeviceTvOsGeneral) GetAirPlayPassword() string`

GetAirPlayPassword returns the AirPlayPassword field if non-nil, zero value otherwise.

### GetAirPlayPasswordOk

`func (o *MobileDeviceTvOsGeneral) GetAirPlayPasswordOk() (*string, bool)`

GetAirPlayPasswordOk returns a tuple with the AirPlayPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirPlayPassword

`func (o *MobileDeviceTvOsGeneral) SetAirPlayPassword(v string)`

SetAirPlayPassword sets AirPlayPassword field to given value.

### HasAirPlayPassword

`func (o *MobileDeviceTvOsGeneral) HasAirPlayPassword() bool`

HasAirPlayPassword returns a boolean if a field has been set.

### GetLocales

`func (o *MobileDeviceTvOsGeneral) GetLocales() string`

GetLocales returns the Locales field if non-nil, zero value otherwise.

### GetLocalesOk

`func (o *MobileDeviceTvOsGeneral) GetLocalesOk() (*string, bool)`

GetLocalesOk returns a tuple with the Locales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocales

`func (o *MobileDeviceTvOsGeneral) SetLocales(v string)`

SetLocales sets Locales field to given value.

### HasLocales

`func (o *MobileDeviceTvOsGeneral) HasLocales() bool`

HasLocales returns a boolean if a field has been set.

### GetLanguages

`func (o *MobileDeviceTvOsGeneral) GetLanguages() string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *MobileDeviceTvOsGeneral) GetLanguagesOk() (*string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *MobileDeviceTvOsGeneral) SetLanguages(v string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *MobileDeviceTvOsGeneral) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


