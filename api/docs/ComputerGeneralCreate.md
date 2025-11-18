# ComputerGeneralCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**LastIpAddress** | Pointer to **string** |  | [optional] 
**LastReportedIp** | Pointer to **string** |  | [optional] 
**JamfBinaryVersion** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Barcode1** | Pointer to **string** |  | [optional] 
**Barcode2** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **string** |  | [optional] 
**RemoteManagement** | Pointer to [**ComputerRemoteManagementCreate**](ComputerRemoteManagementCreate.md) |  | [optional] 
**Supervised** | Pointer to **bool** |  | [optional] 
**MdmCapable** | Pointer to **bool** |  | [optional] 
**ReportDate** | Pointer to **time.Time** |  | [optional] 
**LastContactTime** | Pointer to **time.Time** |  | [optional] 
**LastCloudBackupDate** | Pointer to **time.Time** |  | [optional] 
**LastEnrolledDate** | Pointer to **time.Time** |  | [optional] 
**DistributionPointId** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**ItunesStoreAccountActive** | Pointer to **bool** |  | [optional] 
**EnrolledViaAutomatedDeviceEnrollment** | Pointer to **bool** |  | [optional] 
**UserApprovedMdm** | Pointer to **bool** |  | [optional] 
**DeclarativeDeviceManagementEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewComputerGeneralCreate

`func NewComputerGeneralCreate(name string, ) *ComputerGeneralCreate`

NewComputerGeneralCreate instantiates a new ComputerGeneralCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerGeneralCreateWithDefaults

`func NewComputerGeneralCreateWithDefaults() *ComputerGeneralCreate`

NewComputerGeneralCreateWithDefaults instantiates a new ComputerGeneralCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ComputerGeneralCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputerGeneralCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputerGeneralCreate) SetName(v string)`

SetName sets Name field to given value.


### GetLastIpAddress

`func (o *ComputerGeneralCreate) GetLastIpAddress() string`

GetLastIpAddress returns the LastIpAddress field if non-nil, zero value otherwise.

### GetLastIpAddressOk

`func (o *ComputerGeneralCreate) GetLastIpAddressOk() (*string, bool)`

GetLastIpAddressOk returns a tuple with the LastIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIpAddress

`func (o *ComputerGeneralCreate) SetLastIpAddress(v string)`

SetLastIpAddress sets LastIpAddress field to given value.

### HasLastIpAddress

`func (o *ComputerGeneralCreate) HasLastIpAddress() bool`

HasLastIpAddress returns a boolean if a field has been set.

### GetLastReportedIp

`func (o *ComputerGeneralCreate) GetLastReportedIp() string`

GetLastReportedIp returns the LastReportedIp field if non-nil, zero value otherwise.

### GetLastReportedIpOk

`func (o *ComputerGeneralCreate) GetLastReportedIpOk() (*string, bool)`

GetLastReportedIpOk returns a tuple with the LastReportedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedIp

`func (o *ComputerGeneralCreate) SetLastReportedIp(v string)`

SetLastReportedIp sets LastReportedIp field to given value.

### HasLastReportedIp

`func (o *ComputerGeneralCreate) HasLastReportedIp() bool`

HasLastReportedIp returns a boolean if a field has been set.

### GetJamfBinaryVersion

`func (o *ComputerGeneralCreate) GetJamfBinaryVersion() string`

GetJamfBinaryVersion returns the JamfBinaryVersion field if non-nil, zero value otherwise.

### GetJamfBinaryVersionOk

`func (o *ComputerGeneralCreate) GetJamfBinaryVersionOk() (*string, bool)`

GetJamfBinaryVersionOk returns a tuple with the JamfBinaryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJamfBinaryVersion

`func (o *ComputerGeneralCreate) SetJamfBinaryVersion(v string)`

SetJamfBinaryVersion sets JamfBinaryVersion field to given value.

### HasJamfBinaryVersion

`func (o *ComputerGeneralCreate) HasJamfBinaryVersion() bool`

HasJamfBinaryVersion returns a boolean if a field has been set.

### GetPlatform

`func (o *ComputerGeneralCreate) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ComputerGeneralCreate) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ComputerGeneralCreate) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ComputerGeneralCreate) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetBarcode1

`func (o *ComputerGeneralCreate) GetBarcode1() string`

GetBarcode1 returns the Barcode1 field if non-nil, zero value otherwise.

### GetBarcode1Ok

`func (o *ComputerGeneralCreate) GetBarcode1Ok() (*string, bool)`

GetBarcode1Ok returns a tuple with the Barcode1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode1

`func (o *ComputerGeneralCreate) SetBarcode1(v string)`

SetBarcode1 sets Barcode1 field to given value.

### HasBarcode1

`func (o *ComputerGeneralCreate) HasBarcode1() bool`

HasBarcode1 returns a boolean if a field has been set.

### GetBarcode2

`func (o *ComputerGeneralCreate) GetBarcode2() string`

GetBarcode2 returns the Barcode2 field if non-nil, zero value otherwise.

### GetBarcode2Ok

`func (o *ComputerGeneralCreate) GetBarcode2Ok() (*string, bool)`

GetBarcode2Ok returns a tuple with the Barcode2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode2

`func (o *ComputerGeneralCreate) SetBarcode2(v string)`

SetBarcode2 sets Barcode2 field to given value.

### HasBarcode2

`func (o *ComputerGeneralCreate) HasBarcode2() bool`

HasBarcode2 returns a boolean if a field has been set.

### GetAssetTag

`func (o *ComputerGeneralCreate) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *ComputerGeneralCreate) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *ComputerGeneralCreate) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *ComputerGeneralCreate) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetRemoteManagement

`func (o *ComputerGeneralCreate) GetRemoteManagement() ComputerRemoteManagementCreate`

GetRemoteManagement returns the RemoteManagement field if non-nil, zero value otherwise.

### GetRemoteManagementOk

`func (o *ComputerGeneralCreate) GetRemoteManagementOk() (*ComputerRemoteManagementCreate, bool)`

GetRemoteManagementOk returns a tuple with the RemoteManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteManagement

`func (o *ComputerGeneralCreate) SetRemoteManagement(v ComputerRemoteManagementCreate)`

SetRemoteManagement sets RemoteManagement field to given value.

### HasRemoteManagement

`func (o *ComputerGeneralCreate) HasRemoteManagement() bool`

HasRemoteManagement returns a boolean if a field has been set.

### GetSupervised

`func (o *ComputerGeneralCreate) GetSupervised() bool`

GetSupervised returns the Supervised field if non-nil, zero value otherwise.

### GetSupervisedOk

`func (o *ComputerGeneralCreate) GetSupervisedOk() (*bool, bool)`

GetSupervisedOk returns a tuple with the Supervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervised

`func (o *ComputerGeneralCreate) SetSupervised(v bool)`

SetSupervised sets Supervised field to given value.

### HasSupervised

`func (o *ComputerGeneralCreate) HasSupervised() bool`

HasSupervised returns a boolean if a field has been set.

### GetMdmCapable

`func (o *ComputerGeneralCreate) GetMdmCapable() bool`

GetMdmCapable returns the MdmCapable field if non-nil, zero value otherwise.

### GetMdmCapableOk

`func (o *ComputerGeneralCreate) GetMdmCapableOk() (*bool, bool)`

GetMdmCapableOk returns a tuple with the MdmCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmCapable

`func (o *ComputerGeneralCreate) SetMdmCapable(v bool)`

SetMdmCapable sets MdmCapable field to given value.

### HasMdmCapable

`func (o *ComputerGeneralCreate) HasMdmCapable() bool`

HasMdmCapable returns a boolean if a field has been set.

### GetReportDate

`func (o *ComputerGeneralCreate) GetReportDate() time.Time`

GetReportDate returns the ReportDate field if non-nil, zero value otherwise.

### GetReportDateOk

`func (o *ComputerGeneralCreate) GetReportDateOk() (*time.Time, bool)`

GetReportDateOk returns a tuple with the ReportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDate

`func (o *ComputerGeneralCreate) SetReportDate(v time.Time)`

SetReportDate sets ReportDate field to given value.

### HasReportDate

`func (o *ComputerGeneralCreate) HasReportDate() bool`

HasReportDate returns a boolean if a field has been set.

### GetLastContactTime

`func (o *ComputerGeneralCreate) GetLastContactTime() time.Time`

GetLastContactTime returns the LastContactTime field if non-nil, zero value otherwise.

### GetLastContactTimeOk

`func (o *ComputerGeneralCreate) GetLastContactTimeOk() (*time.Time, bool)`

GetLastContactTimeOk returns a tuple with the LastContactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastContactTime

`func (o *ComputerGeneralCreate) SetLastContactTime(v time.Time)`

SetLastContactTime sets LastContactTime field to given value.

### HasLastContactTime

`func (o *ComputerGeneralCreate) HasLastContactTime() bool`

HasLastContactTime returns a boolean if a field has been set.

### GetLastCloudBackupDate

`func (o *ComputerGeneralCreate) GetLastCloudBackupDate() time.Time`

GetLastCloudBackupDate returns the LastCloudBackupDate field if non-nil, zero value otherwise.

### GetLastCloudBackupDateOk

`func (o *ComputerGeneralCreate) GetLastCloudBackupDateOk() (*time.Time, bool)`

GetLastCloudBackupDateOk returns a tuple with the LastCloudBackupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCloudBackupDate

`func (o *ComputerGeneralCreate) SetLastCloudBackupDate(v time.Time)`

SetLastCloudBackupDate sets LastCloudBackupDate field to given value.

### HasLastCloudBackupDate

`func (o *ComputerGeneralCreate) HasLastCloudBackupDate() bool`

HasLastCloudBackupDate returns a boolean if a field has been set.

### GetLastEnrolledDate

`func (o *ComputerGeneralCreate) GetLastEnrolledDate() time.Time`

GetLastEnrolledDate returns the LastEnrolledDate field if non-nil, zero value otherwise.

### GetLastEnrolledDateOk

`func (o *ComputerGeneralCreate) GetLastEnrolledDateOk() (*time.Time, bool)`

GetLastEnrolledDateOk returns a tuple with the LastEnrolledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEnrolledDate

`func (o *ComputerGeneralCreate) SetLastEnrolledDate(v time.Time)`

SetLastEnrolledDate sets LastEnrolledDate field to given value.

### HasLastEnrolledDate

`func (o *ComputerGeneralCreate) HasLastEnrolledDate() bool`

HasLastEnrolledDate returns a boolean if a field has been set.

### GetDistributionPointId

`func (o *ComputerGeneralCreate) GetDistributionPointId() string`

GetDistributionPointId returns the DistributionPointId field if non-nil, zero value otherwise.

### GetDistributionPointIdOk

`func (o *ComputerGeneralCreate) GetDistributionPointIdOk() (*string, bool)`

GetDistributionPointIdOk returns a tuple with the DistributionPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionPointId

`func (o *ComputerGeneralCreate) SetDistributionPointId(v string)`

SetDistributionPointId sets DistributionPointId field to given value.

### HasDistributionPointId

`func (o *ComputerGeneralCreate) HasDistributionPointId() bool`

HasDistributionPointId returns a boolean if a field has been set.

### GetSiteId

`func (o *ComputerGeneralCreate) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ComputerGeneralCreate) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ComputerGeneralCreate) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ComputerGeneralCreate) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetItunesStoreAccountActive

`func (o *ComputerGeneralCreate) GetItunesStoreAccountActive() bool`

GetItunesStoreAccountActive returns the ItunesStoreAccountActive field if non-nil, zero value otherwise.

### GetItunesStoreAccountActiveOk

`func (o *ComputerGeneralCreate) GetItunesStoreAccountActiveOk() (*bool, bool)`

GetItunesStoreAccountActiveOk returns a tuple with the ItunesStoreAccountActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItunesStoreAccountActive

`func (o *ComputerGeneralCreate) SetItunesStoreAccountActive(v bool)`

SetItunesStoreAccountActive sets ItunesStoreAccountActive field to given value.

### HasItunesStoreAccountActive

`func (o *ComputerGeneralCreate) HasItunesStoreAccountActive() bool`

HasItunesStoreAccountActive returns a boolean if a field has been set.

### GetEnrolledViaAutomatedDeviceEnrollment

`func (o *ComputerGeneralCreate) GetEnrolledViaAutomatedDeviceEnrollment() bool`

GetEnrolledViaAutomatedDeviceEnrollment returns the EnrolledViaAutomatedDeviceEnrollment field if non-nil, zero value otherwise.

### GetEnrolledViaAutomatedDeviceEnrollmentOk

`func (o *ComputerGeneralCreate) GetEnrolledViaAutomatedDeviceEnrollmentOk() (*bool, bool)`

GetEnrolledViaAutomatedDeviceEnrollmentOk returns a tuple with the EnrolledViaAutomatedDeviceEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolledViaAutomatedDeviceEnrollment

`func (o *ComputerGeneralCreate) SetEnrolledViaAutomatedDeviceEnrollment(v bool)`

SetEnrolledViaAutomatedDeviceEnrollment sets EnrolledViaAutomatedDeviceEnrollment field to given value.

### HasEnrolledViaAutomatedDeviceEnrollment

`func (o *ComputerGeneralCreate) HasEnrolledViaAutomatedDeviceEnrollment() bool`

HasEnrolledViaAutomatedDeviceEnrollment returns a boolean if a field has been set.

### GetUserApprovedMdm

`func (o *ComputerGeneralCreate) GetUserApprovedMdm() bool`

GetUserApprovedMdm returns the UserApprovedMdm field if non-nil, zero value otherwise.

### GetUserApprovedMdmOk

`func (o *ComputerGeneralCreate) GetUserApprovedMdmOk() (*bool, bool)`

GetUserApprovedMdmOk returns a tuple with the UserApprovedMdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserApprovedMdm

`func (o *ComputerGeneralCreate) SetUserApprovedMdm(v bool)`

SetUserApprovedMdm sets UserApprovedMdm field to given value.

### HasUserApprovedMdm

`func (o *ComputerGeneralCreate) HasUserApprovedMdm() bool`

HasUserApprovedMdm returns a boolean if a field has been set.

### GetDeclarativeDeviceManagementEnabled

`func (o *ComputerGeneralCreate) GetDeclarativeDeviceManagementEnabled() bool`

GetDeclarativeDeviceManagementEnabled returns the DeclarativeDeviceManagementEnabled field if non-nil, zero value otherwise.

### GetDeclarativeDeviceManagementEnabledOk

`func (o *ComputerGeneralCreate) GetDeclarativeDeviceManagementEnabledOk() (*bool, bool)`

GetDeclarativeDeviceManagementEnabledOk returns a tuple with the DeclarativeDeviceManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarativeDeviceManagementEnabled

`func (o *ComputerGeneralCreate) SetDeclarativeDeviceManagementEnabled(v bool)`

SetDeclarativeDeviceManagementEnabled sets DeclarativeDeviceManagementEnabled field to given value.

### HasDeclarativeDeviceManagementEnabled

`func (o *ComputerGeneralCreate) HasDeclarativeDeviceManagementEnabled() bool`

HasDeclarativeDeviceManagementEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


