# MobileDeviceGeneral

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

## Methods

### NewMobileDeviceGeneral

`func NewMobileDeviceGeneral() *MobileDeviceGeneral`

NewMobileDeviceGeneral instantiates a new MobileDeviceGeneral object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceGeneralWithDefaults

`func NewMobileDeviceGeneralWithDefaults() *MobileDeviceGeneral`

NewMobileDeviceGeneralWithDefaults instantiates a new MobileDeviceGeneral object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUdid

`func (o *MobileDeviceGeneral) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *MobileDeviceGeneral) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *MobileDeviceGeneral) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *MobileDeviceGeneral) HasUdid() bool`

HasUdid returns a boolean if a field has been set.

### GetDisplayName

`func (o *MobileDeviceGeneral) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MobileDeviceGeneral) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MobileDeviceGeneral) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MobileDeviceGeneral) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAssetTag

`func (o *MobileDeviceGeneral) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *MobileDeviceGeneral) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *MobileDeviceGeneral) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *MobileDeviceGeneral) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetSiteId

`func (o *MobileDeviceGeneral) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *MobileDeviceGeneral) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *MobileDeviceGeneral) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *MobileDeviceGeneral) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetLastInventoryUpdateDate

`func (o *MobileDeviceGeneral) GetLastInventoryUpdateDate() time.Time`

GetLastInventoryUpdateDate returns the LastInventoryUpdateDate field if non-nil, zero value otherwise.

### GetLastInventoryUpdateDateOk

`func (o *MobileDeviceGeneral) GetLastInventoryUpdateDateOk() (*time.Time, bool)`

GetLastInventoryUpdateDateOk returns a tuple with the LastInventoryUpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInventoryUpdateDate

`func (o *MobileDeviceGeneral) SetLastInventoryUpdateDate(v time.Time)`

SetLastInventoryUpdateDate sets LastInventoryUpdateDate field to given value.

### HasLastInventoryUpdateDate

`func (o *MobileDeviceGeneral) HasLastInventoryUpdateDate() bool`

HasLastInventoryUpdateDate returns a boolean if a field has been set.

### GetOsVersion

`func (o *MobileDeviceGeneral) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *MobileDeviceGeneral) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *MobileDeviceGeneral) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *MobileDeviceGeneral) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetOsRapidSecurityResponse

`func (o *MobileDeviceGeneral) GetOsRapidSecurityResponse() string`

GetOsRapidSecurityResponse returns the OsRapidSecurityResponse field if non-nil, zero value otherwise.

### GetOsRapidSecurityResponseOk

`func (o *MobileDeviceGeneral) GetOsRapidSecurityResponseOk() (*string, bool)`

GetOsRapidSecurityResponseOk returns a tuple with the OsRapidSecurityResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsRapidSecurityResponse

`func (o *MobileDeviceGeneral) SetOsRapidSecurityResponse(v string)`

SetOsRapidSecurityResponse sets OsRapidSecurityResponse field to given value.

### HasOsRapidSecurityResponse

`func (o *MobileDeviceGeneral) HasOsRapidSecurityResponse() bool`

HasOsRapidSecurityResponse returns a boolean if a field has been set.

### GetOsBuild

`func (o *MobileDeviceGeneral) GetOsBuild() string`

GetOsBuild returns the OsBuild field if non-nil, zero value otherwise.

### GetOsBuildOk

`func (o *MobileDeviceGeneral) GetOsBuildOk() (*string, bool)`

GetOsBuildOk returns a tuple with the OsBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsBuild

`func (o *MobileDeviceGeneral) SetOsBuild(v string)`

SetOsBuild sets OsBuild field to given value.

### HasOsBuild

`func (o *MobileDeviceGeneral) HasOsBuild() bool`

HasOsBuild returns a boolean if a field has been set.

### GetOsSupplementalBuildVersion

`func (o *MobileDeviceGeneral) GetOsSupplementalBuildVersion() string`

GetOsSupplementalBuildVersion returns the OsSupplementalBuildVersion field if non-nil, zero value otherwise.

### GetOsSupplementalBuildVersionOk

`func (o *MobileDeviceGeneral) GetOsSupplementalBuildVersionOk() (*string, bool)`

GetOsSupplementalBuildVersionOk returns a tuple with the OsSupplementalBuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsSupplementalBuildVersion

`func (o *MobileDeviceGeneral) SetOsSupplementalBuildVersion(v string)`

SetOsSupplementalBuildVersion sets OsSupplementalBuildVersion field to given value.

### HasOsSupplementalBuildVersion

`func (o *MobileDeviceGeneral) HasOsSupplementalBuildVersion() bool`

HasOsSupplementalBuildVersion returns a boolean if a field has been set.

### GetSoftwareUpdateDeviceId

`func (o *MobileDeviceGeneral) GetSoftwareUpdateDeviceId() string`

GetSoftwareUpdateDeviceId returns the SoftwareUpdateDeviceId field if non-nil, zero value otherwise.

### GetSoftwareUpdateDeviceIdOk

`func (o *MobileDeviceGeneral) GetSoftwareUpdateDeviceIdOk() (*string, bool)`

GetSoftwareUpdateDeviceIdOk returns a tuple with the SoftwareUpdateDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdateDeviceId

`func (o *MobileDeviceGeneral) SetSoftwareUpdateDeviceId(v string)`

SetSoftwareUpdateDeviceId sets SoftwareUpdateDeviceId field to given value.

### HasSoftwareUpdateDeviceId

`func (o *MobileDeviceGeneral) HasSoftwareUpdateDeviceId() bool`

HasSoftwareUpdateDeviceId returns a boolean if a field has been set.

### GetIpAddress

`func (o *MobileDeviceGeneral) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *MobileDeviceGeneral) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *MobileDeviceGeneral) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *MobileDeviceGeneral) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetManaged

`func (o *MobileDeviceGeneral) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *MobileDeviceGeneral) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *MobileDeviceGeneral) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *MobileDeviceGeneral) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetSupervised

`func (o *MobileDeviceGeneral) GetSupervised() bool`

GetSupervised returns the Supervised field if non-nil, zero value otherwise.

### GetSupervisedOk

`func (o *MobileDeviceGeneral) GetSupervisedOk() (*bool, bool)`

GetSupervisedOk returns a tuple with the Supervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervised

`func (o *MobileDeviceGeneral) SetSupervised(v bool)`

SetSupervised sets Supervised field to given value.

### HasSupervised

`func (o *MobileDeviceGeneral) HasSupervised() bool`

HasSupervised returns a boolean if a field has been set.

### GetDeviceOwnershipType

`func (o *MobileDeviceGeneral) GetDeviceOwnershipType() string`

GetDeviceOwnershipType returns the DeviceOwnershipType field if non-nil, zero value otherwise.

### GetDeviceOwnershipTypeOk

`func (o *MobileDeviceGeneral) GetDeviceOwnershipTypeOk() (*string, bool)`

GetDeviceOwnershipTypeOk returns a tuple with the DeviceOwnershipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOwnershipType

`func (o *MobileDeviceGeneral) SetDeviceOwnershipType(v string)`

SetDeviceOwnershipType sets DeviceOwnershipType field to given value.

### HasDeviceOwnershipType

`func (o *MobileDeviceGeneral) HasDeviceOwnershipType() bool`

HasDeviceOwnershipType returns a boolean if a field has been set.

### GetEnrollmentMethodPrestage

`func (o *MobileDeviceGeneral) GetEnrollmentMethodPrestage() EnrollmentMethodPrestage`

GetEnrollmentMethodPrestage returns the EnrollmentMethodPrestage field if non-nil, zero value otherwise.

### GetEnrollmentMethodPrestageOk

`func (o *MobileDeviceGeneral) GetEnrollmentMethodPrestageOk() (*EnrollmentMethodPrestage, bool)`

GetEnrollmentMethodPrestageOk returns a tuple with the EnrollmentMethodPrestage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentMethodPrestage

`func (o *MobileDeviceGeneral) SetEnrollmentMethodPrestage(v EnrollmentMethodPrestage)`

SetEnrollmentMethodPrestage sets EnrollmentMethodPrestage field to given value.

### HasEnrollmentMethodPrestage

`func (o *MobileDeviceGeneral) HasEnrollmentMethodPrestage() bool`

HasEnrollmentMethodPrestage returns a boolean if a field has been set.

### GetEnrollmentSessionTokenValid

`func (o *MobileDeviceGeneral) GetEnrollmentSessionTokenValid() bool`

GetEnrollmentSessionTokenValid returns the EnrollmentSessionTokenValid field if non-nil, zero value otherwise.

### GetEnrollmentSessionTokenValidOk

`func (o *MobileDeviceGeneral) GetEnrollmentSessionTokenValidOk() (*bool, bool)`

GetEnrollmentSessionTokenValidOk returns a tuple with the EnrollmentSessionTokenValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSessionTokenValid

`func (o *MobileDeviceGeneral) SetEnrollmentSessionTokenValid(v bool)`

SetEnrollmentSessionTokenValid sets EnrollmentSessionTokenValid field to given value.

### HasEnrollmentSessionTokenValid

`func (o *MobileDeviceGeneral) HasEnrollmentSessionTokenValid() bool`

HasEnrollmentSessionTokenValid returns a boolean if a field has been set.

### GetLastEnrolledDate

`func (o *MobileDeviceGeneral) GetLastEnrolledDate() time.Time`

GetLastEnrolledDate returns the LastEnrolledDate field if non-nil, zero value otherwise.

### GetLastEnrolledDateOk

`func (o *MobileDeviceGeneral) GetLastEnrolledDateOk() (*time.Time, bool)`

GetLastEnrolledDateOk returns a tuple with the LastEnrolledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEnrolledDate

`func (o *MobileDeviceGeneral) SetLastEnrolledDate(v time.Time)`

SetLastEnrolledDate sets LastEnrolledDate field to given value.

### HasLastEnrolledDate

`func (o *MobileDeviceGeneral) HasLastEnrolledDate() bool`

HasLastEnrolledDate returns a boolean if a field has been set.

### GetMdmProfileExpirationDate

`func (o *MobileDeviceGeneral) GetMdmProfileExpirationDate() time.Time`

GetMdmProfileExpirationDate returns the MdmProfileExpirationDate field if non-nil, zero value otherwise.

### GetMdmProfileExpirationDateOk

`func (o *MobileDeviceGeneral) GetMdmProfileExpirationDateOk() (*time.Time, bool)`

GetMdmProfileExpirationDateOk returns a tuple with the MdmProfileExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmProfileExpirationDate

`func (o *MobileDeviceGeneral) SetMdmProfileExpirationDate(v time.Time)`

SetMdmProfileExpirationDate sets MdmProfileExpirationDate field to given value.

### HasMdmProfileExpirationDate

`func (o *MobileDeviceGeneral) HasMdmProfileExpirationDate() bool`

HasMdmProfileExpirationDate returns a boolean if a field has been set.

### GetTimeZone

`func (o *MobileDeviceGeneral) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MobileDeviceGeneral) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MobileDeviceGeneral) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MobileDeviceGeneral) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetDeclarativeDeviceManagementEnabled

`func (o *MobileDeviceGeneral) GetDeclarativeDeviceManagementEnabled() bool`

GetDeclarativeDeviceManagementEnabled returns the DeclarativeDeviceManagementEnabled field if non-nil, zero value otherwise.

### GetDeclarativeDeviceManagementEnabledOk

`func (o *MobileDeviceGeneral) GetDeclarativeDeviceManagementEnabledOk() (*bool, bool)`

GetDeclarativeDeviceManagementEnabledOk returns a tuple with the DeclarativeDeviceManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarativeDeviceManagementEnabled

`func (o *MobileDeviceGeneral) SetDeclarativeDeviceManagementEnabled(v bool)`

SetDeclarativeDeviceManagementEnabled sets DeclarativeDeviceManagementEnabled field to given value.

### HasDeclarativeDeviceManagementEnabled

`func (o *MobileDeviceGeneral) HasDeclarativeDeviceManagementEnabled() bool`

HasDeclarativeDeviceManagementEnabled returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *MobileDeviceGeneral) GetExtensionAttributes() []MobileDeviceExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *MobileDeviceGeneral) GetExtensionAttributesOk() (*[]MobileDeviceExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *MobileDeviceGeneral) SetExtensionAttributes(v []MobileDeviceExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *MobileDeviceGeneral) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


