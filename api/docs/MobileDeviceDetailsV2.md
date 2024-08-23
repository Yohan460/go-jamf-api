# MobileDeviceDetailsV2

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

## Methods

### NewMobileDeviceDetailsV2

`func NewMobileDeviceDetailsV2() *MobileDeviceDetailsV2`

NewMobileDeviceDetailsV2 instantiates a new MobileDeviceDetailsV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceDetailsV2WithDefaults

`func NewMobileDeviceDetailsV2WithDefaults() *MobileDeviceDetailsV2`

NewMobileDeviceDetailsV2WithDefaults instantiates a new MobileDeviceDetailsV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MobileDeviceDetailsV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MobileDeviceDetailsV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MobileDeviceDetailsV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MobileDeviceDetailsV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MobileDeviceDetailsV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MobileDeviceDetailsV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MobileDeviceDetailsV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MobileDeviceDetailsV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnforceName

`func (o *MobileDeviceDetailsV2) GetEnforceName() bool`

GetEnforceName returns the EnforceName field if non-nil, zero value otherwise.

### GetEnforceNameOk

`func (o *MobileDeviceDetailsV2) GetEnforceNameOk() (*bool, bool)`

GetEnforceNameOk returns a tuple with the EnforceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceName

`func (o *MobileDeviceDetailsV2) SetEnforceName(v bool)`

SetEnforceName sets EnforceName field to given value.

### HasEnforceName

`func (o *MobileDeviceDetailsV2) HasEnforceName() bool`

HasEnforceName returns a boolean if a field has been set.

### GetAssetTag

`func (o *MobileDeviceDetailsV2) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *MobileDeviceDetailsV2) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *MobileDeviceDetailsV2) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *MobileDeviceDetailsV2) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetLastInventoryUpdateTimestamp

`func (o *MobileDeviceDetailsV2) GetLastInventoryUpdateTimestamp() time.Time`

GetLastInventoryUpdateTimestamp returns the LastInventoryUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastInventoryUpdateTimestampOk

`func (o *MobileDeviceDetailsV2) GetLastInventoryUpdateTimestampOk() (*time.Time, bool)`

GetLastInventoryUpdateTimestampOk returns a tuple with the LastInventoryUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInventoryUpdateTimestamp

`func (o *MobileDeviceDetailsV2) SetLastInventoryUpdateTimestamp(v time.Time)`

SetLastInventoryUpdateTimestamp sets LastInventoryUpdateTimestamp field to given value.

### HasLastInventoryUpdateTimestamp

`func (o *MobileDeviceDetailsV2) HasLastInventoryUpdateTimestamp() bool`

HasLastInventoryUpdateTimestamp returns a boolean if a field has been set.

### GetOsVersion

`func (o *MobileDeviceDetailsV2) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *MobileDeviceDetailsV2) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *MobileDeviceDetailsV2) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *MobileDeviceDetailsV2) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetOsBuild

`func (o *MobileDeviceDetailsV2) GetOsBuild() string`

GetOsBuild returns the OsBuild field if non-nil, zero value otherwise.

### GetOsBuildOk

`func (o *MobileDeviceDetailsV2) GetOsBuildOk() (*string, bool)`

GetOsBuildOk returns a tuple with the OsBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsBuild

`func (o *MobileDeviceDetailsV2) SetOsBuild(v string)`

SetOsBuild sets OsBuild field to given value.

### HasOsBuild

`func (o *MobileDeviceDetailsV2) HasOsBuild() bool`

HasOsBuild returns a boolean if a field has been set.

### GetOsSupplementalBuildVersion

`func (o *MobileDeviceDetailsV2) GetOsSupplementalBuildVersion() string`

GetOsSupplementalBuildVersion returns the OsSupplementalBuildVersion field if non-nil, zero value otherwise.

### GetOsSupplementalBuildVersionOk

`func (o *MobileDeviceDetailsV2) GetOsSupplementalBuildVersionOk() (*string, bool)`

GetOsSupplementalBuildVersionOk returns a tuple with the OsSupplementalBuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsSupplementalBuildVersion

`func (o *MobileDeviceDetailsV2) SetOsSupplementalBuildVersion(v string)`

SetOsSupplementalBuildVersion sets OsSupplementalBuildVersion field to given value.

### HasOsSupplementalBuildVersion

`func (o *MobileDeviceDetailsV2) HasOsSupplementalBuildVersion() bool`

HasOsSupplementalBuildVersion returns a boolean if a field has been set.

### GetOsRapidSecurityResponse

`func (o *MobileDeviceDetailsV2) GetOsRapidSecurityResponse() string`

GetOsRapidSecurityResponse returns the OsRapidSecurityResponse field if non-nil, zero value otherwise.

### GetOsRapidSecurityResponseOk

`func (o *MobileDeviceDetailsV2) GetOsRapidSecurityResponseOk() (*string, bool)`

GetOsRapidSecurityResponseOk returns a tuple with the OsRapidSecurityResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsRapidSecurityResponse

`func (o *MobileDeviceDetailsV2) SetOsRapidSecurityResponse(v string)`

SetOsRapidSecurityResponse sets OsRapidSecurityResponse field to given value.

### HasOsRapidSecurityResponse

`func (o *MobileDeviceDetailsV2) HasOsRapidSecurityResponse() bool`

HasOsRapidSecurityResponse returns a boolean if a field has been set.

### GetSoftwareUpdateDeviceId

`func (o *MobileDeviceDetailsV2) GetSoftwareUpdateDeviceId() string`

GetSoftwareUpdateDeviceId returns the SoftwareUpdateDeviceId field if non-nil, zero value otherwise.

### GetSoftwareUpdateDeviceIdOk

`func (o *MobileDeviceDetailsV2) GetSoftwareUpdateDeviceIdOk() (*string, bool)`

GetSoftwareUpdateDeviceIdOk returns a tuple with the SoftwareUpdateDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdateDeviceId

`func (o *MobileDeviceDetailsV2) SetSoftwareUpdateDeviceId(v string)`

SetSoftwareUpdateDeviceId sets SoftwareUpdateDeviceId field to given value.

### HasSoftwareUpdateDeviceId

`func (o *MobileDeviceDetailsV2) HasSoftwareUpdateDeviceId() bool`

HasSoftwareUpdateDeviceId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *MobileDeviceDetailsV2) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *MobileDeviceDetailsV2) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *MobileDeviceDetailsV2) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *MobileDeviceDetailsV2) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetUdid

`func (o *MobileDeviceDetailsV2) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *MobileDeviceDetailsV2) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *MobileDeviceDetailsV2) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *MobileDeviceDetailsV2) HasUdid() bool`

HasUdid returns a boolean if a field has been set.

### GetIpAddress

`func (o *MobileDeviceDetailsV2) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *MobileDeviceDetailsV2) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *MobileDeviceDetailsV2) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *MobileDeviceDetailsV2) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetWifiMacAddress

`func (o *MobileDeviceDetailsV2) GetWifiMacAddress() string`

GetWifiMacAddress returns the WifiMacAddress field if non-nil, zero value otherwise.

### GetWifiMacAddressOk

`func (o *MobileDeviceDetailsV2) GetWifiMacAddressOk() (*string, bool)`

GetWifiMacAddressOk returns a tuple with the WifiMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMacAddress

`func (o *MobileDeviceDetailsV2) SetWifiMacAddress(v string)`

SetWifiMacAddress sets WifiMacAddress field to given value.

### HasWifiMacAddress

`func (o *MobileDeviceDetailsV2) HasWifiMacAddress() bool`

HasWifiMacAddress returns a boolean if a field has been set.

### GetBluetoothMacAddress

`func (o *MobileDeviceDetailsV2) GetBluetoothMacAddress() string`

GetBluetoothMacAddress returns the BluetoothMacAddress field if non-nil, zero value otherwise.

### GetBluetoothMacAddressOk

`func (o *MobileDeviceDetailsV2) GetBluetoothMacAddressOk() (*string, bool)`

GetBluetoothMacAddressOk returns a tuple with the BluetoothMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBluetoothMacAddress

`func (o *MobileDeviceDetailsV2) SetBluetoothMacAddress(v string)`

SetBluetoothMacAddress sets BluetoothMacAddress field to given value.

### HasBluetoothMacAddress

`func (o *MobileDeviceDetailsV2) HasBluetoothMacAddress() bool`

HasBluetoothMacAddress returns a boolean if a field has been set.

### GetManaged

`func (o *MobileDeviceDetailsV2) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *MobileDeviceDetailsV2) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *MobileDeviceDetailsV2) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *MobileDeviceDetailsV2) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetTimeZone

`func (o *MobileDeviceDetailsV2) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MobileDeviceDetailsV2) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MobileDeviceDetailsV2) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MobileDeviceDetailsV2) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetInitialEntryTimestamp

`func (o *MobileDeviceDetailsV2) GetInitialEntryTimestamp() time.Time`

GetInitialEntryTimestamp returns the InitialEntryTimestamp field if non-nil, zero value otherwise.

### GetInitialEntryTimestampOk

`func (o *MobileDeviceDetailsV2) GetInitialEntryTimestampOk() (*time.Time, bool)`

GetInitialEntryTimestampOk returns a tuple with the InitialEntryTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialEntryTimestamp

`func (o *MobileDeviceDetailsV2) SetInitialEntryTimestamp(v time.Time)`

SetInitialEntryTimestamp sets InitialEntryTimestamp field to given value.

### HasInitialEntryTimestamp

`func (o *MobileDeviceDetailsV2) HasInitialEntryTimestamp() bool`

HasInitialEntryTimestamp returns a boolean if a field has been set.

### GetLastEnrollmentTimestamp

`func (o *MobileDeviceDetailsV2) GetLastEnrollmentTimestamp() time.Time`

GetLastEnrollmentTimestamp returns the LastEnrollmentTimestamp field if non-nil, zero value otherwise.

### GetLastEnrollmentTimestampOk

`func (o *MobileDeviceDetailsV2) GetLastEnrollmentTimestampOk() (*time.Time, bool)`

GetLastEnrollmentTimestampOk returns a tuple with the LastEnrollmentTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEnrollmentTimestamp

`func (o *MobileDeviceDetailsV2) SetLastEnrollmentTimestamp(v time.Time)`

SetLastEnrollmentTimestamp sets LastEnrollmentTimestamp field to given value.

### HasLastEnrollmentTimestamp

`func (o *MobileDeviceDetailsV2) HasLastEnrollmentTimestamp() bool`

HasLastEnrollmentTimestamp returns a boolean if a field has been set.

### GetMdmProfileExpirationTimestamp

`func (o *MobileDeviceDetailsV2) GetMdmProfileExpirationTimestamp() time.Time`

GetMdmProfileExpirationTimestamp returns the MdmProfileExpirationTimestamp field if non-nil, zero value otherwise.

### GetMdmProfileExpirationTimestampOk

`func (o *MobileDeviceDetailsV2) GetMdmProfileExpirationTimestampOk() (*time.Time, bool)`

GetMdmProfileExpirationTimestampOk returns a tuple with the MdmProfileExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmProfileExpirationTimestamp

`func (o *MobileDeviceDetailsV2) SetMdmProfileExpirationTimestamp(v time.Time)`

SetMdmProfileExpirationTimestamp sets MdmProfileExpirationTimestamp field to given value.

### HasMdmProfileExpirationTimestamp

`func (o *MobileDeviceDetailsV2) HasMdmProfileExpirationTimestamp() bool`

HasMdmProfileExpirationTimestamp returns a boolean if a field has been set.

### GetDeviceOwnershipLevel

`func (o *MobileDeviceDetailsV2) GetDeviceOwnershipLevel() string`

GetDeviceOwnershipLevel returns the DeviceOwnershipLevel field if non-nil, zero value otherwise.

### GetDeviceOwnershipLevelOk

`func (o *MobileDeviceDetailsV2) GetDeviceOwnershipLevelOk() (*string, bool)`

GetDeviceOwnershipLevelOk returns a tuple with the DeviceOwnershipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOwnershipLevel

`func (o *MobileDeviceDetailsV2) SetDeviceOwnershipLevel(v string)`

SetDeviceOwnershipLevel sets DeviceOwnershipLevel field to given value.

### HasDeviceOwnershipLevel

`func (o *MobileDeviceDetailsV2) HasDeviceOwnershipLevel() bool`

HasDeviceOwnershipLevel returns a boolean if a field has been set.

### GetEnrollmentMethod

`func (o *MobileDeviceDetailsV2) GetEnrollmentMethod() string`

GetEnrollmentMethod returns the EnrollmentMethod field if non-nil, zero value otherwise.

### GetEnrollmentMethodOk

`func (o *MobileDeviceDetailsV2) GetEnrollmentMethodOk() (*string, bool)`

GetEnrollmentMethodOk returns a tuple with the EnrollmentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentMethod

`func (o *MobileDeviceDetailsV2) SetEnrollmentMethod(v string)`

SetEnrollmentMethod sets EnrollmentMethod field to given value.

### HasEnrollmentMethod

`func (o *MobileDeviceDetailsV2) HasEnrollmentMethod() bool`

HasEnrollmentMethod returns a boolean if a field has been set.

### GetEnrollmentSessionTokenValid

`func (o *MobileDeviceDetailsV2) GetEnrollmentSessionTokenValid() bool`

GetEnrollmentSessionTokenValid returns the EnrollmentSessionTokenValid field if non-nil, zero value otherwise.

### GetEnrollmentSessionTokenValidOk

`func (o *MobileDeviceDetailsV2) GetEnrollmentSessionTokenValidOk() (*bool, bool)`

GetEnrollmentSessionTokenValidOk returns a tuple with the EnrollmentSessionTokenValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSessionTokenValid

`func (o *MobileDeviceDetailsV2) SetEnrollmentSessionTokenValid(v bool)`

SetEnrollmentSessionTokenValid sets EnrollmentSessionTokenValid field to given value.

### HasEnrollmentSessionTokenValid

`func (o *MobileDeviceDetailsV2) HasEnrollmentSessionTokenValid() bool`

HasEnrollmentSessionTokenValid returns a boolean if a field has been set.

### GetDeclarativeDeviceManagementEnabled

`func (o *MobileDeviceDetailsV2) GetDeclarativeDeviceManagementEnabled() bool`

GetDeclarativeDeviceManagementEnabled returns the DeclarativeDeviceManagementEnabled field if non-nil, zero value otherwise.

### GetDeclarativeDeviceManagementEnabledOk

`func (o *MobileDeviceDetailsV2) GetDeclarativeDeviceManagementEnabledOk() (*bool, bool)`

GetDeclarativeDeviceManagementEnabledOk returns a tuple with the DeclarativeDeviceManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarativeDeviceManagementEnabled

`func (o *MobileDeviceDetailsV2) SetDeclarativeDeviceManagementEnabled(v bool)`

SetDeclarativeDeviceManagementEnabled sets DeclarativeDeviceManagementEnabled field to given value.

### HasDeclarativeDeviceManagementEnabled

`func (o *MobileDeviceDetailsV2) HasDeclarativeDeviceManagementEnabled() bool`

HasDeclarativeDeviceManagementEnabled returns a boolean if a field has been set.

### GetSite

`func (o *MobileDeviceDetailsV2) GetSite() V1Site`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *MobileDeviceDetailsV2) GetSiteOk() (*V1Site, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *MobileDeviceDetailsV2) SetSite(v V1Site)`

SetSite sets Site field to given value.

### HasSite

`func (o *MobileDeviceDetailsV2) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *MobileDeviceDetailsV2) GetExtensionAttributes() []ExtensionAttributeV2`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *MobileDeviceDetailsV2) GetExtensionAttributesOk() (*[]ExtensionAttributeV2, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *MobileDeviceDetailsV2) SetExtensionAttributes(v []ExtensionAttributeV2)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *MobileDeviceDetailsV2) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.

### GetLocation

`func (o *MobileDeviceDetailsV2) GetLocation() LocationV2`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MobileDeviceDetailsV2) GetLocationOk() (*LocationV2, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MobileDeviceDetailsV2) SetLocation(v LocationV2)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MobileDeviceDetailsV2) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetType

`func (o *MobileDeviceDetailsV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MobileDeviceDetailsV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MobileDeviceDetailsV2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MobileDeviceDetailsV2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIos

`func (o *MobileDeviceDetailsV2) GetIos() DetailsV2`

GetIos returns the Ios field if non-nil, zero value otherwise.

### GetIosOk

`func (o *MobileDeviceDetailsV2) GetIosOk() (*DetailsV2, bool)`

GetIosOk returns a tuple with the Ios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIos

`func (o *MobileDeviceDetailsV2) SetIos(v DetailsV2)`

SetIos sets Ios field to given value.

### HasIos

`func (o *MobileDeviceDetailsV2) HasIos() bool`

HasIos returns a boolean if a field has been set.

### GetTvos

`func (o *MobileDeviceDetailsV2) GetTvos() TvOsDetails`

GetTvos returns the Tvos field if non-nil, zero value otherwise.

### GetTvosOk

`func (o *MobileDeviceDetailsV2) GetTvosOk() (*TvOsDetails, bool)`

GetTvosOk returns a tuple with the Tvos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvos

`func (o *MobileDeviceDetailsV2) SetTvos(v TvOsDetails)`

SetTvos sets Tvos field to given value.

### HasTvos

`func (o *MobileDeviceDetailsV2) HasTvos() bool`

HasTvos returns a boolean if a field has been set.

### GetWatchos

`func (o *MobileDeviceDetailsV2) GetWatchos() WatchOsDetailsV2`

GetWatchos returns the Watchos field if non-nil, zero value otherwise.

### GetWatchosOk

`func (o *MobileDeviceDetailsV2) GetWatchosOk() (*WatchOsDetailsV2, bool)`

GetWatchosOk returns a tuple with the Watchos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchos

`func (o *MobileDeviceDetailsV2) SetWatchos(v WatchOsDetailsV2)`

SetWatchos sets Watchos field to given value.

### HasWatchos

`func (o *MobileDeviceDetailsV2) HasWatchos() bool`

HasWatchos returns a boolean if a field has been set.

### GetVisionos

`func (o *MobileDeviceDetailsV2) GetVisionos() DetailsV2`

GetVisionos returns the Visionos field if non-nil, zero value otherwise.

### GetVisionosOk

`func (o *MobileDeviceDetailsV2) GetVisionosOk() (*DetailsV2, bool)`

GetVisionosOk returns a tuple with the Visionos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisionos

`func (o *MobileDeviceDetailsV2) SetVisionos(v DetailsV2)`

SetVisionos sets Visionos field to given value.

### HasVisionos

`func (o *MobileDeviceDetailsV2) HasVisionos() bool`

HasVisionos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


