# MobileDeviceDetailsGetV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** | Mobile device name. | [optional] 
**EnforceName** | Pointer to **bool** | Enforce the mobile device name. Device must be supervised. If set to true, Jamf Pro will revert the Mobile Device Name to the ���name��� value each time the device checks in. | [optional] 
**AssetTag** | Pointer to **string** |  | [optional] 
**LastInventoryUpdateTimestamp** | Pointer to **time.Time** |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**OsBuild** | Pointer to **string** |  | [optional] 
**OsSupplementalBuildVersion** | Pointer to **string** | Collected for iOS 16 and iPadOS 16.1 or later | [optional] 
**OsRapidSecurityResponse** | Pointer to **string** | Collected for iOS 16 and iPadOS 16.1 or later | [optional] 
**SoftwareUpdateDeviceId** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Udid** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**WifiMacAddress** | Pointer to **string** |  | [optional] 
**BluetoothMacAddress** | Pointer to **string** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 
**InitialEntryTimestamp** | Pointer to **time.Time** |  | [optional] 
**LastEnrollmentTimestamp** | Pointer to **time.Time** |  | [optional] 
**MdmProfileExpirationTimestamp** | Pointer to **time.Time** |  | [optional] 
**DeviceOwnershipLevel** | Pointer to **string** |  | [optional] 
**EnrollmentMethod** | Pointer to **string** |  | [optional] 
**EnrollmentSessionTokenValid** | Pointer to **bool** |  | [optional] 
**DeclarativeDeviceManagementEnabled** | Pointer to **bool** |  | [optional] 
**Site** | Pointer to [**V1Site**](V1Site.md) |  | [optional] 
**ExtensionAttributes** | Pointer to [**[]ExtensionAttributeV2**](ExtensionAttributeV2.md) |  | [optional] 
**Location** | Pointer to [**LocationV2**](LocationV2.md) |  | [optional] 
**Type** | Pointer to **string** | Based on the value of this either iOS, tvOS, watch or visionOS objects will be populated. | [optional] 
**Ios** | Pointer to [**DetailsV2**](DetailsV2.md) |  | [optional] 
**Tvos** | Pointer to [**TvOsDetails**](TvOsDetails.md) |  | [optional] 
**Watchos** | Pointer to [**WatchOsDetailsV2**](WatchOsDetailsV2.md) |  | [optional] 
**Visionos** | Pointer to [**DetailsV2**](DetailsV2.md) |  | [optional] 
**ManagementId** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewMobileDeviceDetailsGetV2

`func NewMobileDeviceDetailsGetV2() *MobileDeviceDetailsGetV2`

NewMobileDeviceDetailsGetV2 instantiates a new MobileDeviceDetailsGetV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceDetailsGetV2WithDefaults

`func NewMobileDeviceDetailsGetV2WithDefaults() *MobileDeviceDetailsGetV2`

NewMobileDeviceDetailsGetV2WithDefaults instantiates a new MobileDeviceDetailsGetV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MobileDeviceDetailsGetV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MobileDeviceDetailsGetV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MobileDeviceDetailsGetV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MobileDeviceDetailsGetV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MobileDeviceDetailsGetV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MobileDeviceDetailsGetV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MobileDeviceDetailsGetV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MobileDeviceDetailsGetV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnforceName

`func (o *MobileDeviceDetailsGetV2) GetEnforceName() bool`

GetEnforceName returns the EnforceName field if non-nil, zero value otherwise.

### GetEnforceNameOk

`func (o *MobileDeviceDetailsGetV2) GetEnforceNameOk() (*bool, bool)`

GetEnforceNameOk returns a tuple with the EnforceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceName

`func (o *MobileDeviceDetailsGetV2) SetEnforceName(v bool)`

SetEnforceName sets EnforceName field to given value.

### HasEnforceName

`func (o *MobileDeviceDetailsGetV2) HasEnforceName() bool`

HasEnforceName returns a boolean if a field has been set.

### GetAssetTag

`func (o *MobileDeviceDetailsGetV2) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *MobileDeviceDetailsGetV2) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *MobileDeviceDetailsGetV2) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *MobileDeviceDetailsGetV2) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetLastInventoryUpdateTimestamp

`func (o *MobileDeviceDetailsGetV2) GetLastInventoryUpdateTimestamp() time.Time`

GetLastInventoryUpdateTimestamp returns the LastInventoryUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastInventoryUpdateTimestampOk

`func (o *MobileDeviceDetailsGetV2) GetLastInventoryUpdateTimestampOk() (*time.Time, bool)`

GetLastInventoryUpdateTimestampOk returns a tuple with the LastInventoryUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInventoryUpdateTimestamp

`func (o *MobileDeviceDetailsGetV2) SetLastInventoryUpdateTimestamp(v time.Time)`

SetLastInventoryUpdateTimestamp sets LastInventoryUpdateTimestamp field to given value.

### HasLastInventoryUpdateTimestamp

`func (o *MobileDeviceDetailsGetV2) HasLastInventoryUpdateTimestamp() bool`

HasLastInventoryUpdateTimestamp returns a boolean if a field has been set.

### GetOsVersion

`func (o *MobileDeviceDetailsGetV2) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *MobileDeviceDetailsGetV2) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *MobileDeviceDetailsGetV2) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *MobileDeviceDetailsGetV2) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetOsBuild

`func (o *MobileDeviceDetailsGetV2) GetOsBuild() string`

GetOsBuild returns the OsBuild field if non-nil, zero value otherwise.

### GetOsBuildOk

`func (o *MobileDeviceDetailsGetV2) GetOsBuildOk() (*string, bool)`

GetOsBuildOk returns a tuple with the OsBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsBuild

`func (o *MobileDeviceDetailsGetV2) SetOsBuild(v string)`

SetOsBuild sets OsBuild field to given value.

### HasOsBuild

`func (o *MobileDeviceDetailsGetV2) HasOsBuild() bool`

HasOsBuild returns a boolean if a field has been set.

### GetOsSupplementalBuildVersion

`func (o *MobileDeviceDetailsGetV2) GetOsSupplementalBuildVersion() string`

GetOsSupplementalBuildVersion returns the OsSupplementalBuildVersion field if non-nil, zero value otherwise.

### GetOsSupplementalBuildVersionOk

`func (o *MobileDeviceDetailsGetV2) GetOsSupplementalBuildVersionOk() (*string, bool)`

GetOsSupplementalBuildVersionOk returns a tuple with the OsSupplementalBuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsSupplementalBuildVersion

`func (o *MobileDeviceDetailsGetV2) SetOsSupplementalBuildVersion(v string)`

SetOsSupplementalBuildVersion sets OsSupplementalBuildVersion field to given value.

### HasOsSupplementalBuildVersion

`func (o *MobileDeviceDetailsGetV2) HasOsSupplementalBuildVersion() bool`

HasOsSupplementalBuildVersion returns a boolean if a field has been set.

### GetOsRapidSecurityResponse

`func (o *MobileDeviceDetailsGetV2) GetOsRapidSecurityResponse() string`

GetOsRapidSecurityResponse returns the OsRapidSecurityResponse field if non-nil, zero value otherwise.

### GetOsRapidSecurityResponseOk

`func (o *MobileDeviceDetailsGetV2) GetOsRapidSecurityResponseOk() (*string, bool)`

GetOsRapidSecurityResponseOk returns a tuple with the OsRapidSecurityResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsRapidSecurityResponse

`func (o *MobileDeviceDetailsGetV2) SetOsRapidSecurityResponse(v string)`

SetOsRapidSecurityResponse sets OsRapidSecurityResponse field to given value.

### HasOsRapidSecurityResponse

`func (o *MobileDeviceDetailsGetV2) HasOsRapidSecurityResponse() bool`

HasOsRapidSecurityResponse returns a boolean if a field has been set.

### GetSoftwareUpdateDeviceId

`func (o *MobileDeviceDetailsGetV2) GetSoftwareUpdateDeviceId() string`

GetSoftwareUpdateDeviceId returns the SoftwareUpdateDeviceId field if non-nil, zero value otherwise.

### GetSoftwareUpdateDeviceIdOk

`func (o *MobileDeviceDetailsGetV2) GetSoftwareUpdateDeviceIdOk() (*string, bool)`

GetSoftwareUpdateDeviceIdOk returns a tuple with the SoftwareUpdateDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdateDeviceId

`func (o *MobileDeviceDetailsGetV2) SetSoftwareUpdateDeviceId(v string)`

SetSoftwareUpdateDeviceId sets SoftwareUpdateDeviceId field to given value.

### HasSoftwareUpdateDeviceId

`func (o *MobileDeviceDetailsGetV2) HasSoftwareUpdateDeviceId() bool`

HasSoftwareUpdateDeviceId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *MobileDeviceDetailsGetV2) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *MobileDeviceDetailsGetV2) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *MobileDeviceDetailsGetV2) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *MobileDeviceDetailsGetV2) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetUdid

`func (o *MobileDeviceDetailsGetV2) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *MobileDeviceDetailsGetV2) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *MobileDeviceDetailsGetV2) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *MobileDeviceDetailsGetV2) HasUdid() bool`

HasUdid returns a boolean if a field has been set.

### GetIpAddress

`func (o *MobileDeviceDetailsGetV2) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *MobileDeviceDetailsGetV2) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *MobileDeviceDetailsGetV2) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *MobileDeviceDetailsGetV2) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetWifiMacAddress

`func (o *MobileDeviceDetailsGetV2) GetWifiMacAddress() string`

GetWifiMacAddress returns the WifiMacAddress field if non-nil, zero value otherwise.

### GetWifiMacAddressOk

`func (o *MobileDeviceDetailsGetV2) GetWifiMacAddressOk() (*string, bool)`

GetWifiMacAddressOk returns a tuple with the WifiMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMacAddress

`func (o *MobileDeviceDetailsGetV2) SetWifiMacAddress(v string)`

SetWifiMacAddress sets WifiMacAddress field to given value.

### HasWifiMacAddress

`func (o *MobileDeviceDetailsGetV2) HasWifiMacAddress() bool`

HasWifiMacAddress returns a boolean if a field has been set.

### GetBluetoothMacAddress

`func (o *MobileDeviceDetailsGetV2) GetBluetoothMacAddress() string`

GetBluetoothMacAddress returns the BluetoothMacAddress field if non-nil, zero value otherwise.

### GetBluetoothMacAddressOk

`func (o *MobileDeviceDetailsGetV2) GetBluetoothMacAddressOk() (*string, bool)`

GetBluetoothMacAddressOk returns a tuple with the BluetoothMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBluetoothMacAddress

`func (o *MobileDeviceDetailsGetV2) SetBluetoothMacAddress(v string)`

SetBluetoothMacAddress sets BluetoothMacAddress field to given value.

### HasBluetoothMacAddress

`func (o *MobileDeviceDetailsGetV2) HasBluetoothMacAddress() bool`

HasBluetoothMacAddress returns a boolean if a field has been set.

### GetManaged

`func (o *MobileDeviceDetailsGetV2) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *MobileDeviceDetailsGetV2) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *MobileDeviceDetailsGetV2) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *MobileDeviceDetailsGetV2) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetTimeZone

`func (o *MobileDeviceDetailsGetV2) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MobileDeviceDetailsGetV2) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MobileDeviceDetailsGetV2) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MobileDeviceDetailsGetV2) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetInitialEntryTimestamp

`func (o *MobileDeviceDetailsGetV2) GetInitialEntryTimestamp() time.Time`

GetInitialEntryTimestamp returns the InitialEntryTimestamp field if non-nil, zero value otherwise.

### GetInitialEntryTimestampOk

`func (o *MobileDeviceDetailsGetV2) GetInitialEntryTimestampOk() (*time.Time, bool)`

GetInitialEntryTimestampOk returns a tuple with the InitialEntryTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialEntryTimestamp

`func (o *MobileDeviceDetailsGetV2) SetInitialEntryTimestamp(v time.Time)`

SetInitialEntryTimestamp sets InitialEntryTimestamp field to given value.

### HasInitialEntryTimestamp

`func (o *MobileDeviceDetailsGetV2) HasInitialEntryTimestamp() bool`

HasInitialEntryTimestamp returns a boolean if a field has been set.

### GetLastEnrollmentTimestamp

`func (o *MobileDeviceDetailsGetV2) GetLastEnrollmentTimestamp() time.Time`

GetLastEnrollmentTimestamp returns the LastEnrollmentTimestamp field if non-nil, zero value otherwise.

### GetLastEnrollmentTimestampOk

`func (o *MobileDeviceDetailsGetV2) GetLastEnrollmentTimestampOk() (*time.Time, bool)`

GetLastEnrollmentTimestampOk returns a tuple with the LastEnrollmentTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEnrollmentTimestamp

`func (o *MobileDeviceDetailsGetV2) SetLastEnrollmentTimestamp(v time.Time)`

SetLastEnrollmentTimestamp sets LastEnrollmentTimestamp field to given value.

### HasLastEnrollmentTimestamp

`func (o *MobileDeviceDetailsGetV2) HasLastEnrollmentTimestamp() bool`

HasLastEnrollmentTimestamp returns a boolean if a field has been set.

### GetMdmProfileExpirationTimestamp

`func (o *MobileDeviceDetailsGetV2) GetMdmProfileExpirationTimestamp() time.Time`

GetMdmProfileExpirationTimestamp returns the MdmProfileExpirationTimestamp field if non-nil, zero value otherwise.

### GetMdmProfileExpirationTimestampOk

`func (o *MobileDeviceDetailsGetV2) GetMdmProfileExpirationTimestampOk() (*time.Time, bool)`

GetMdmProfileExpirationTimestampOk returns a tuple with the MdmProfileExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmProfileExpirationTimestamp

`func (o *MobileDeviceDetailsGetV2) SetMdmProfileExpirationTimestamp(v time.Time)`

SetMdmProfileExpirationTimestamp sets MdmProfileExpirationTimestamp field to given value.

### HasMdmProfileExpirationTimestamp

`func (o *MobileDeviceDetailsGetV2) HasMdmProfileExpirationTimestamp() bool`

HasMdmProfileExpirationTimestamp returns a boolean if a field has been set.

### GetDeviceOwnershipLevel

`func (o *MobileDeviceDetailsGetV2) GetDeviceOwnershipLevel() string`

GetDeviceOwnershipLevel returns the DeviceOwnershipLevel field if non-nil, zero value otherwise.

### GetDeviceOwnershipLevelOk

`func (o *MobileDeviceDetailsGetV2) GetDeviceOwnershipLevelOk() (*string, bool)`

GetDeviceOwnershipLevelOk returns a tuple with the DeviceOwnershipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOwnershipLevel

`func (o *MobileDeviceDetailsGetV2) SetDeviceOwnershipLevel(v string)`

SetDeviceOwnershipLevel sets DeviceOwnershipLevel field to given value.

### HasDeviceOwnershipLevel

`func (o *MobileDeviceDetailsGetV2) HasDeviceOwnershipLevel() bool`

HasDeviceOwnershipLevel returns a boolean if a field has been set.

### GetEnrollmentMethod

`func (o *MobileDeviceDetailsGetV2) GetEnrollmentMethod() string`

GetEnrollmentMethod returns the EnrollmentMethod field if non-nil, zero value otherwise.

### GetEnrollmentMethodOk

`func (o *MobileDeviceDetailsGetV2) GetEnrollmentMethodOk() (*string, bool)`

GetEnrollmentMethodOk returns a tuple with the EnrollmentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentMethod

`func (o *MobileDeviceDetailsGetV2) SetEnrollmentMethod(v string)`

SetEnrollmentMethod sets EnrollmentMethod field to given value.

### HasEnrollmentMethod

`func (o *MobileDeviceDetailsGetV2) HasEnrollmentMethod() bool`

HasEnrollmentMethod returns a boolean if a field has been set.

### GetEnrollmentSessionTokenValid

`func (o *MobileDeviceDetailsGetV2) GetEnrollmentSessionTokenValid() bool`

GetEnrollmentSessionTokenValid returns the EnrollmentSessionTokenValid field if non-nil, zero value otherwise.

### GetEnrollmentSessionTokenValidOk

`func (o *MobileDeviceDetailsGetV2) GetEnrollmentSessionTokenValidOk() (*bool, bool)`

GetEnrollmentSessionTokenValidOk returns a tuple with the EnrollmentSessionTokenValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSessionTokenValid

`func (o *MobileDeviceDetailsGetV2) SetEnrollmentSessionTokenValid(v bool)`

SetEnrollmentSessionTokenValid sets EnrollmentSessionTokenValid field to given value.

### HasEnrollmentSessionTokenValid

`func (o *MobileDeviceDetailsGetV2) HasEnrollmentSessionTokenValid() bool`

HasEnrollmentSessionTokenValid returns a boolean if a field has been set.

### GetDeclarativeDeviceManagementEnabled

`func (o *MobileDeviceDetailsGetV2) GetDeclarativeDeviceManagementEnabled() bool`

GetDeclarativeDeviceManagementEnabled returns the DeclarativeDeviceManagementEnabled field if non-nil, zero value otherwise.

### GetDeclarativeDeviceManagementEnabledOk

`func (o *MobileDeviceDetailsGetV2) GetDeclarativeDeviceManagementEnabledOk() (*bool, bool)`

GetDeclarativeDeviceManagementEnabledOk returns a tuple with the DeclarativeDeviceManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarativeDeviceManagementEnabled

`func (o *MobileDeviceDetailsGetV2) SetDeclarativeDeviceManagementEnabled(v bool)`

SetDeclarativeDeviceManagementEnabled sets DeclarativeDeviceManagementEnabled field to given value.

### HasDeclarativeDeviceManagementEnabled

`func (o *MobileDeviceDetailsGetV2) HasDeclarativeDeviceManagementEnabled() bool`

HasDeclarativeDeviceManagementEnabled returns a boolean if a field has been set.

### GetSite

`func (o *MobileDeviceDetailsGetV2) GetSite() V1Site`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *MobileDeviceDetailsGetV2) GetSiteOk() (*V1Site, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *MobileDeviceDetailsGetV2) SetSite(v V1Site)`

SetSite sets Site field to given value.

### HasSite

`func (o *MobileDeviceDetailsGetV2) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *MobileDeviceDetailsGetV2) GetExtensionAttributes() []ExtensionAttributeV2`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *MobileDeviceDetailsGetV2) GetExtensionAttributesOk() (*[]ExtensionAttributeV2, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *MobileDeviceDetailsGetV2) SetExtensionAttributes(v []ExtensionAttributeV2)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *MobileDeviceDetailsGetV2) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.

### GetLocation

`func (o *MobileDeviceDetailsGetV2) GetLocation() LocationV2`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MobileDeviceDetailsGetV2) GetLocationOk() (*LocationV2, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MobileDeviceDetailsGetV2) SetLocation(v LocationV2)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MobileDeviceDetailsGetV2) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetType

`func (o *MobileDeviceDetailsGetV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MobileDeviceDetailsGetV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MobileDeviceDetailsGetV2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MobileDeviceDetailsGetV2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIos

`func (o *MobileDeviceDetailsGetV2) GetIos() DetailsV2`

GetIos returns the Ios field if non-nil, zero value otherwise.

### GetIosOk

`func (o *MobileDeviceDetailsGetV2) GetIosOk() (*DetailsV2, bool)`

GetIosOk returns a tuple with the Ios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIos

`func (o *MobileDeviceDetailsGetV2) SetIos(v DetailsV2)`

SetIos sets Ios field to given value.

### HasIos

`func (o *MobileDeviceDetailsGetV2) HasIos() bool`

HasIos returns a boolean if a field has been set.

### GetTvos

`func (o *MobileDeviceDetailsGetV2) GetTvos() TvOsDetails`

GetTvos returns the Tvos field if non-nil, zero value otherwise.

### GetTvosOk

`func (o *MobileDeviceDetailsGetV2) GetTvosOk() (*TvOsDetails, bool)`

GetTvosOk returns a tuple with the Tvos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvos

`func (o *MobileDeviceDetailsGetV2) SetTvos(v TvOsDetails)`

SetTvos sets Tvos field to given value.

### HasTvos

`func (o *MobileDeviceDetailsGetV2) HasTvos() bool`

HasTvos returns a boolean if a field has been set.

### GetWatchos

`func (o *MobileDeviceDetailsGetV2) GetWatchos() WatchOsDetailsV2`

GetWatchos returns the Watchos field if non-nil, zero value otherwise.

### GetWatchosOk

`func (o *MobileDeviceDetailsGetV2) GetWatchosOk() (*WatchOsDetailsV2, bool)`

GetWatchosOk returns a tuple with the Watchos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchos

`func (o *MobileDeviceDetailsGetV2) SetWatchos(v WatchOsDetailsV2)`

SetWatchos sets Watchos field to given value.

### HasWatchos

`func (o *MobileDeviceDetailsGetV2) HasWatchos() bool`

HasWatchos returns a boolean if a field has been set.

### GetVisionos

`func (o *MobileDeviceDetailsGetV2) GetVisionos() DetailsV2`

GetVisionos returns the Visionos field if non-nil, zero value otherwise.

### GetVisionosOk

`func (o *MobileDeviceDetailsGetV2) GetVisionosOk() (*DetailsV2, bool)`

GetVisionosOk returns a tuple with the Visionos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisionos

`func (o *MobileDeviceDetailsGetV2) SetVisionos(v DetailsV2)`

SetVisionos sets Visionos field to given value.

### HasVisionos

`func (o *MobileDeviceDetailsGetV2) HasVisionos() bool`

HasVisionos returns a boolean if a field has been set.

### GetManagementId

`func (o *MobileDeviceDetailsGetV2) GetManagementId() string`

GetManagementId returns the ManagementId field if non-nil, zero value otherwise.

### GetManagementIdOk

`func (o *MobileDeviceDetailsGetV2) GetManagementIdOk() (*string, bool)`

GetManagementIdOk returns a tuple with the ManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementId

`func (o *MobileDeviceDetailsGetV2) SetManagementId(v string)`

SetManagementId sets ManagementId field to given value.

### HasManagementId

`func (o *MobileDeviceDetailsGetV2) HasManagementId() bool`

HasManagementId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


