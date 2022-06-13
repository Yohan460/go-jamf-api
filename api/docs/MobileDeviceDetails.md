# MobileDeviceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **string** |  | [optional] 
**LastInventoryUpdateTimestamp** | Pointer to **time.Time** |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**OsBuild** | Pointer to **string** |  | [optional] 
**SoftwareUpdateDeviceId** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Udid** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**WifiMacAddress** | Pointer to **string** |  | [optional] 
**BluetoothMacAddress** | Pointer to **string** |  | [optional] 
**IsManaged** | Pointer to **bool** |  | [optional] 
**InitialEntryTimestamp** | Pointer to **time.Time** |  | [optional] 
**LastEnrollmentTimestamp** | Pointer to **time.Time** |  | [optional] 
**DeviceOwnershipLevel** | Pointer to **string** |  | [optional] 
**Site** | Pointer to [**IdAndName**](IdAndName.md) |  | [optional] 
**ExtensionAttributes** | Pointer to [**[]ExtensionAttribute**](ExtensionAttribute.md) |  | [optional] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**Type** | Pointer to **string** | Based on the value of this either ios, appleTv, android objects will be populated. | [optional] 
**Ios** | Pointer to [**IosDetails**](IosDetails.md) |  | [optional] 
**AppleTv** | Pointer to [**AppleTvDetails**](AppleTvDetails.md) |  | [optional] 
**Android** | Pointer to [**AndroidDetails**](AndroidDetails.md) |  | [optional] 

## Methods

### NewMobileDeviceDetails

`func NewMobileDeviceDetails() *MobileDeviceDetails`

NewMobileDeviceDetails instantiates a new MobileDeviceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceDetailsWithDefaults

`func NewMobileDeviceDetailsWithDefaults() *MobileDeviceDetails`

NewMobileDeviceDetailsWithDefaults instantiates a new MobileDeviceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MobileDeviceDetails) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MobileDeviceDetails) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MobileDeviceDetails) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MobileDeviceDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MobileDeviceDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MobileDeviceDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MobileDeviceDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MobileDeviceDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAssetTag

`func (o *MobileDeviceDetails) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *MobileDeviceDetails) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *MobileDeviceDetails) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *MobileDeviceDetails) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetLastInventoryUpdateTimestamp

`func (o *MobileDeviceDetails) GetLastInventoryUpdateTimestamp() time.Time`

GetLastInventoryUpdateTimestamp returns the LastInventoryUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastInventoryUpdateTimestampOk

`func (o *MobileDeviceDetails) GetLastInventoryUpdateTimestampOk() (*time.Time, bool)`

GetLastInventoryUpdateTimestampOk returns a tuple with the LastInventoryUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInventoryUpdateTimestamp

`func (o *MobileDeviceDetails) SetLastInventoryUpdateTimestamp(v time.Time)`

SetLastInventoryUpdateTimestamp sets LastInventoryUpdateTimestamp field to given value.

### HasLastInventoryUpdateTimestamp

`func (o *MobileDeviceDetails) HasLastInventoryUpdateTimestamp() bool`

HasLastInventoryUpdateTimestamp returns a boolean if a field has been set.

### GetOsVersion

`func (o *MobileDeviceDetails) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *MobileDeviceDetails) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *MobileDeviceDetails) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *MobileDeviceDetails) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetOsBuild

`func (o *MobileDeviceDetails) GetOsBuild() string`

GetOsBuild returns the OsBuild field if non-nil, zero value otherwise.

### GetOsBuildOk

`func (o *MobileDeviceDetails) GetOsBuildOk() (*string, bool)`

GetOsBuildOk returns a tuple with the OsBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsBuild

`func (o *MobileDeviceDetails) SetOsBuild(v string)`

SetOsBuild sets OsBuild field to given value.

### HasOsBuild

`func (o *MobileDeviceDetails) HasOsBuild() bool`

HasOsBuild returns a boolean if a field has been set.

### GetSoftwareUpdateDeviceId

`func (o *MobileDeviceDetails) GetSoftwareUpdateDeviceId() string`

GetSoftwareUpdateDeviceId returns the SoftwareUpdateDeviceId field if non-nil, zero value otherwise.

### GetSoftwareUpdateDeviceIdOk

`func (o *MobileDeviceDetails) GetSoftwareUpdateDeviceIdOk() (*string, bool)`

GetSoftwareUpdateDeviceIdOk returns a tuple with the SoftwareUpdateDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdateDeviceId

`func (o *MobileDeviceDetails) SetSoftwareUpdateDeviceId(v string)`

SetSoftwareUpdateDeviceId sets SoftwareUpdateDeviceId field to given value.

### HasSoftwareUpdateDeviceId

`func (o *MobileDeviceDetails) HasSoftwareUpdateDeviceId() bool`

HasSoftwareUpdateDeviceId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *MobileDeviceDetails) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *MobileDeviceDetails) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *MobileDeviceDetails) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *MobileDeviceDetails) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetUdid

`func (o *MobileDeviceDetails) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *MobileDeviceDetails) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *MobileDeviceDetails) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *MobileDeviceDetails) HasUdid() bool`

HasUdid returns a boolean if a field has been set.

### GetIpAddress

`func (o *MobileDeviceDetails) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *MobileDeviceDetails) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *MobileDeviceDetails) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *MobileDeviceDetails) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetWifiMacAddress

`func (o *MobileDeviceDetails) GetWifiMacAddress() string`

GetWifiMacAddress returns the WifiMacAddress field if non-nil, zero value otherwise.

### GetWifiMacAddressOk

`func (o *MobileDeviceDetails) GetWifiMacAddressOk() (*string, bool)`

GetWifiMacAddressOk returns a tuple with the WifiMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMacAddress

`func (o *MobileDeviceDetails) SetWifiMacAddress(v string)`

SetWifiMacAddress sets WifiMacAddress field to given value.

### HasWifiMacAddress

`func (o *MobileDeviceDetails) HasWifiMacAddress() bool`

HasWifiMacAddress returns a boolean if a field has been set.

### GetBluetoothMacAddress

`func (o *MobileDeviceDetails) GetBluetoothMacAddress() string`

GetBluetoothMacAddress returns the BluetoothMacAddress field if non-nil, zero value otherwise.

### GetBluetoothMacAddressOk

`func (o *MobileDeviceDetails) GetBluetoothMacAddressOk() (*string, bool)`

GetBluetoothMacAddressOk returns a tuple with the BluetoothMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBluetoothMacAddress

`func (o *MobileDeviceDetails) SetBluetoothMacAddress(v string)`

SetBluetoothMacAddress sets BluetoothMacAddress field to given value.

### HasBluetoothMacAddress

`func (o *MobileDeviceDetails) HasBluetoothMacAddress() bool`

HasBluetoothMacAddress returns a boolean if a field has been set.

### GetIsManaged

`func (o *MobileDeviceDetails) GetIsManaged() bool`

GetIsManaged returns the IsManaged field if non-nil, zero value otherwise.

### GetIsManagedOk

`func (o *MobileDeviceDetails) GetIsManagedOk() (*bool, bool)`

GetIsManagedOk returns a tuple with the IsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManaged

`func (o *MobileDeviceDetails) SetIsManaged(v bool)`

SetIsManaged sets IsManaged field to given value.

### HasIsManaged

`func (o *MobileDeviceDetails) HasIsManaged() bool`

HasIsManaged returns a boolean if a field has been set.

### GetInitialEntryTimestamp

`func (o *MobileDeviceDetails) GetInitialEntryTimestamp() time.Time`

GetInitialEntryTimestamp returns the InitialEntryTimestamp field if non-nil, zero value otherwise.

### GetInitialEntryTimestampOk

`func (o *MobileDeviceDetails) GetInitialEntryTimestampOk() (*time.Time, bool)`

GetInitialEntryTimestampOk returns a tuple with the InitialEntryTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialEntryTimestamp

`func (o *MobileDeviceDetails) SetInitialEntryTimestamp(v time.Time)`

SetInitialEntryTimestamp sets InitialEntryTimestamp field to given value.

### HasInitialEntryTimestamp

`func (o *MobileDeviceDetails) HasInitialEntryTimestamp() bool`

HasInitialEntryTimestamp returns a boolean if a field has been set.

### GetLastEnrollmentTimestamp

`func (o *MobileDeviceDetails) GetLastEnrollmentTimestamp() time.Time`

GetLastEnrollmentTimestamp returns the LastEnrollmentTimestamp field if non-nil, zero value otherwise.

### GetLastEnrollmentTimestampOk

`func (o *MobileDeviceDetails) GetLastEnrollmentTimestampOk() (*time.Time, bool)`

GetLastEnrollmentTimestampOk returns a tuple with the LastEnrollmentTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEnrollmentTimestamp

`func (o *MobileDeviceDetails) SetLastEnrollmentTimestamp(v time.Time)`

SetLastEnrollmentTimestamp sets LastEnrollmentTimestamp field to given value.

### HasLastEnrollmentTimestamp

`func (o *MobileDeviceDetails) HasLastEnrollmentTimestamp() bool`

HasLastEnrollmentTimestamp returns a boolean if a field has been set.

### GetDeviceOwnershipLevel

`func (o *MobileDeviceDetails) GetDeviceOwnershipLevel() string`

GetDeviceOwnershipLevel returns the DeviceOwnershipLevel field if non-nil, zero value otherwise.

### GetDeviceOwnershipLevelOk

`func (o *MobileDeviceDetails) GetDeviceOwnershipLevelOk() (*string, bool)`

GetDeviceOwnershipLevelOk returns a tuple with the DeviceOwnershipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOwnershipLevel

`func (o *MobileDeviceDetails) SetDeviceOwnershipLevel(v string)`

SetDeviceOwnershipLevel sets DeviceOwnershipLevel field to given value.

### HasDeviceOwnershipLevel

`func (o *MobileDeviceDetails) HasDeviceOwnershipLevel() bool`

HasDeviceOwnershipLevel returns a boolean if a field has been set.

### GetSite

`func (o *MobileDeviceDetails) GetSite() IdAndName`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *MobileDeviceDetails) GetSiteOk() (*IdAndName, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *MobileDeviceDetails) SetSite(v IdAndName)`

SetSite sets Site field to given value.

### HasSite

`func (o *MobileDeviceDetails) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *MobileDeviceDetails) GetExtensionAttributes() []ExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *MobileDeviceDetails) GetExtensionAttributesOk() (*[]ExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *MobileDeviceDetails) SetExtensionAttributes(v []ExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *MobileDeviceDetails) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.

### GetLocation

`func (o *MobileDeviceDetails) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MobileDeviceDetails) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MobileDeviceDetails) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MobileDeviceDetails) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetType

`func (o *MobileDeviceDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MobileDeviceDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MobileDeviceDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MobileDeviceDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIos

`func (o *MobileDeviceDetails) GetIos() IosDetails`

GetIos returns the Ios field if non-nil, zero value otherwise.

### GetIosOk

`func (o *MobileDeviceDetails) GetIosOk() (*IosDetails, bool)`

GetIosOk returns a tuple with the Ios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIos

`func (o *MobileDeviceDetails) SetIos(v IosDetails)`

SetIos sets Ios field to given value.

### HasIos

`func (o *MobileDeviceDetails) HasIos() bool`

HasIos returns a boolean if a field has been set.

### GetAppleTv

`func (o *MobileDeviceDetails) GetAppleTv() AppleTvDetails`

GetAppleTv returns the AppleTv field if non-nil, zero value otherwise.

### GetAppleTvOk

`func (o *MobileDeviceDetails) GetAppleTvOk() (*AppleTvDetails, bool)`

GetAppleTvOk returns a tuple with the AppleTv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleTv

`func (o *MobileDeviceDetails) SetAppleTv(v AppleTvDetails)`

SetAppleTv sets AppleTv field to given value.

### HasAppleTv

`func (o *MobileDeviceDetails) HasAppleTv() bool`

HasAppleTv returns a boolean if a field has been set.

### GetAndroid

`func (o *MobileDeviceDetails) GetAndroid() AndroidDetails`

GetAndroid returns the Android field if non-nil, zero value otherwise.

### GetAndroidOk

`func (o *MobileDeviceDetails) GetAndroidOk() (*AndroidDetails, bool)`

GetAndroidOk returns a tuple with the Android field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroid

`func (o *MobileDeviceDetails) SetAndroid(v AndroidDetails)`

SetAndroid sets Android field to given value.

### HasAndroid

`func (o *MobileDeviceDetails) HasAndroid() bool`

HasAndroid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


