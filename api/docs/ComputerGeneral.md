# ComputerGeneral

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**LastIpAddress** | Pointer to **string** |  | [optional] 
**LastReportedIp** | Pointer to **string** |  | [optional] 
**JamfBinaryVersion** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Barcode1** | Pointer to **string** |  | [optional] 
**Barcode2** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **string** |  | [optional] 
**RemoteManagement** | Pointer to [**ComputerRemoteManagement**](ComputerRemoteManagement.md) |  | [optional] 
**Supervised** | Pointer to **bool** |  | [optional] 
**MdmCapable** | Pointer to [**ComputerMdmCapability**](ComputerMdmCapability.md) |  | [optional] 
**ReportDate** | Pointer to **time.Time** |  | [optional] 
**LastContactTime** | Pointer to **time.Time** |  | [optional] 
**LastCloudBackupDate** | Pointer to **time.Time** |  | [optional] 
**LastEnrolledDate** | Pointer to **time.Time** |  | [optional] 
**MdmProfileExpiration** | Pointer to **time.Time** |  | [optional] 
**InitialEntryDate** | Pointer to **string** |  | [optional] 
**DistributionPoint** | Pointer to **string** |  | [optional] 
**EnrollmentMethod** | Pointer to [**EnrollmentMethod**](EnrollmentMethod.md) |  | [optional] 
**Site** | Pointer to [**V1Site**](V1Site.md) |  | [optional] 
**ItunesStoreAccountActive** | Pointer to **bool** |  | [optional] 
**EnrolledViaAutomatedDeviceEnrollment** | Pointer to **bool** |  | [optional] 
**UserApprovedMdm** | Pointer to **bool** |  | [optional] 
**DeclarativeDeviceManagementEnabled** | Pointer to **bool** |  | [optional] 
**ExtensionAttributes** | Pointer to [**[]ComputerExtensionAttribute**](ComputerExtensionAttribute.md) |  | [optional] 
**ManagementId** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewComputerGeneral

`func NewComputerGeneral() *ComputerGeneral`

NewComputerGeneral instantiates a new ComputerGeneral object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerGeneralWithDefaults

`func NewComputerGeneralWithDefaults() *ComputerGeneral`

NewComputerGeneralWithDefaults instantiates a new ComputerGeneral object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ComputerGeneral) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputerGeneral) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputerGeneral) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputerGeneral) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLastIpAddress

`func (o *ComputerGeneral) GetLastIpAddress() string`

GetLastIpAddress returns the LastIpAddress field if non-nil, zero value otherwise.

### GetLastIpAddressOk

`func (o *ComputerGeneral) GetLastIpAddressOk() (*string, bool)`

GetLastIpAddressOk returns a tuple with the LastIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIpAddress

`func (o *ComputerGeneral) SetLastIpAddress(v string)`

SetLastIpAddress sets LastIpAddress field to given value.

### HasLastIpAddress

`func (o *ComputerGeneral) HasLastIpAddress() bool`

HasLastIpAddress returns a boolean if a field has been set.

### GetLastReportedIp

`func (o *ComputerGeneral) GetLastReportedIp() string`

GetLastReportedIp returns the LastReportedIp field if non-nil, zero value otherwise.

### GetLastReportedIpOk

`func (o *ComputerGeneral) GetLastReportedIpOk() (*string, bool)`

GetLastReportedIpOk returns a tuple with the LastReportedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedIp

`func (o *ComputerGeneral) SetLastReportedIp(v string)`

SetLastReportedIp sets LastReportedIp field to given value.

### HasLastReportedIp

`func (o *ComputerGeneral) HasLastReportedIp() bool`

HasLastReportedIp returns a boolean if a field has been set.

### GetJamfBinaryVersion

`func (o *ComputerGeneral) GetJamfBinaryVersion() string`

GetJamfBinaryVersion returns the JamfBinaryVersion field if non-nil, zero value otherwise.

### GetJamfBinaryVersionOk

`func (o *ComputerGeneral) GetJamfBinaryVersionOk() (*string, bool)`

GetJamfBinaryVersionOk returns a tuple with the JamfBinaryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJamfBinaryVersion

`func (o *ComputerGeneral) SetJamfBinaryVersion(v string)`

SetJamfBinaryVersion sets JamfBinaryVersion field to given value.

### HasJamfBinaryVersion

`func (o *ComputerGeneral) HasJamfBinaryVersion() bool`

HasJamfBinaryVersion returns a boolean if a field has been set.

### GetPlatform

`func (o *ComputerGeneral) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ComputerGeneral) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ComputerGeneral) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ComputerGeneral) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetBarcode1

`func (o *ComputerGeneral) GetBarcode1() string`

GetBarcode1 returns the Barcode1 field if non-nil, zero value otherwise.

### GetBarcode1Ok

`func (o *ComputerGeneral) GetBarcode1Ok() (*string, bool)`

GetBarcode1Ok returns a tuple with the Barcode1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode1

`func (o *ComputerGeneral) SetBarcode1(v string)`

SetBarcode1 sets Barcode1 field to given value.

### HasBarcode1

`func (o *ComputerGeneral) HasBarcode1() bool`

HasBarcode1 returns a boolean if a field has been set.

### GetBarcode2

`func (o *ComputerGeneral) GetBarcode2() string`

GetBarcode2 returns the Barcode2 field if non-nil, zero value otherwise.

### GetBarcode2Ok

`func (o *ComputerGeneral) GetBarcode2Ok() (*string, bool)`

GetBarcode2Ok returns a tuple with the Barcode2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode2

`func (o *ComputerGeneral) SetBarcode2(v string)`

SetBarcode2 sets Barcode2 field to given value.

### HasBarcode2

`func (o *ComputerGeneral) HasBarcode2() bool`

HasBarcode2 returns a boolean if a field has been set.

### GetAssetTag

`func (o *ComputerGeneral) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *ComputerGeneral) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *ComputerGeneral) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *ComputerGeneral) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetRemoteManagement

`func (o *ComputerGeneral) GetRemoteManagement() ComputerRemoteManagement`

GetRemoteManagement returns the RemoteManagement field if non-nil, zero value otherwise.

### GetRemoteManagementOk

`func (o *ComputerGeneral) GetRemoteManagementOk() (*ComputerRemoteManagement, bool)`

GetRemoteManagementOk returns a tuple with the RemoteManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteManagement

`func (o *ComputerGeneral) SetRemoteManagement(v ComputerRemoteManagement)`

SetRemoteManagement sets RemoteManagement field to given value.

### HasRemoteManagement

`func (o *ComputerGeneral) HasRemoteManagement() bool`

HasRemoteManagement returns a boolean if a field has been set.

### GetSupervised

`func (o *ComputerGeneral) GetSupervised() bool`

GetSupervised returns the Supervised field if non-nil, zero value otherwise.

### GetSupervisedOk

`func (o *ComputerGeneral) GetSupervisedOk() (*bool, bool)`

GetSupervisedOk returns a tuple with the Supervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervised

`func (o *ComputerGeneral) SetSupervised(v bool)`

SetSupervised sets Supervised field to given value.

### HasSupervised

`func (o *ComputerGeneral) HasSupervised() bool`

HasSupervised returns a boolean if a field has been set.

### GetMdmCapable

`func (o *ComputerGeneral) GetMdmCapable() ComputerMdmCapability`

GetMdmCapable returns the MdmCapable field if non-nil, zero value otherwise.

### GetMdmCapableOk

`func (o *ComputerGeneral) GetMdmCapableOk() (*ComputerMdmCapability, bool)`

GetMdmCapableOk returns a tuple with the MdmCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmCapable

`func (o *ComputerGeneral) SetMdmCapable(v ComputerMdmCapability)`

SetMdmCapable sets MdmCapable field to given value.

### HasMdmCapable

`func (o *ComputerGeneral) HasMdmCapable() bool`

HasMdmCapable returns a boolean if a field has been set.

### GetReportDate

`func (o *ComputerGeneral) GetReportDate() time.Time`

GetReportDate returns the ReportDate field if non-nil, zero value otherwise.

### GetReportDateOk

`func (o *ComputerGeneral) GetReportDateOk() (*time.Time, bool)`

GetReportDateOk returns a tuple with the ReportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDate

`func (o *ComputerGeneral) SetReportDate(v time.Time)`

SetReportDate sets ReportDate field to given value.

### HasReportDate

`func (o *ComputerGeneral) HasReportDate() bool`

HasReportDate returns a boolean if a field has been set.

### GetLastContactTime

`func (o *ComputerGeneral) GetLastContactTime() time.Time`

GetLastContactTime returns the LastContactTime field if non-nil, zero value otherwise.

### GetLastContactTimeOk

`func (o *ComputerGeneral) GetLastContactTimeOk() (*time.Time, bool)`

GetLastContactTimeOk returns a tuple with the LastContactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastContactTime

`func (o *ComputerGeneral) SetLastContactTime(v time.Time)`

SetLastContactTime sets LastContactTime field to given value.

### HasLastContactTime

`func (o *ComputerGeneral) HasLastContactTime() bool`

HasLastContactTime returns a boolean if a field has been set.

### GetLastCloudBackupDate

`func (o *ComputerGeneral) GetLastCloudBackupDate() time.Time`

GetLastCloudBackupDate returns the LastCloudBackupDate field if non-nil, zero value otherwise.

### GetLastCloudBackupDateOk

`func (o *ComputerGeneral) GetLastCloudBackupDateOk() (*time.Time, bool)`

GetLastCloudBackupDateOk returns a tuple with the LastCloudBackupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCloudBackupDate

`func (o *ComputerGeneral) SetLastCloudBackupDate(v time.Time)`

SetLastCloudBackupDate sets LastCloudBackupDate field to given value.

### HasLastCloudBackupDate

`func (o *ComputerGeneral) HasLastCloudBackupDate() bool`

HasLastCloudBackupDate returns a boolean if a field has been set.

### GetLastEnrolledDate

`func (o *ComputerGeneral) GetLastEnrolledDate() time.Time`

GetLastEnrolledDate returns the LastEnrolledDate field if non-nil, zero value otherwise.

### GetLastEnrolledDateOk

`func (o *ComputerGeneral) GetLastEnrolledDateOk() (*time.Time, bool)`

GetLastEnrolledDateOk returns a tuple with the LastEnrolledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEnrolledDate

`func (o *ComputerGeneral) SetLastEnrolledDate(v time.Time)`

SetLastEnrolledDate sets LastEnrolledDate field to given value.

### HasLastEnrolledDate

`func (o *ComputerGeneral) HasLastEnrolledDate() bool`

HasLastEnrolledDate returns a boolean if a field has been set.

### GetMdmProfileExpiration

`func (o *ComputerGeneral) GetMdmProfileExpiration() time.Time`

GetMdmProfileExpiration returns the MdmProfileExpiration field if non-nil, zero value otherwise.

### GetMdmProfileExpirationOk

`func (o *ComputerGeneral) GetMdmProfileExpirationOk() (*time.Time, bool)`

GetMdmProfileExpirationOk returns a tuple with the MdmProfileExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmProfileExpiration

`func (o *ComputerGeneral) SetMdmProfileExpiration(v time.Time)`

SetMdmProfileExpiration sets MdmProfileExpiration field to given value.

### HasMdmProfileExpiration

`func (o *ComputerGeneral) HasMdmProfileExpiration() bool`

HasMdmProfileExpiration returns a boolean if a field has been set.

### GetInitialEntryDate

`func (o *ComputerGeneral) GetInitialEntryDate() string`

GetInitialEntryDate returns the InitialEntryDate field if non-nil, zero value otherwise.

### GetInitialEntryDateOk

`func (o *ComputerGeneral) GetInitialEntryDateOk() (*string, bool)`

GetInitialEntryDateOk returns a tuple with the InitialEntryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialEntryDate

`func (o *ComputerGeneral) SetInitialEntryDate(v string)`

SetInitialEntryDate sets InitialEntryDate field to given value.

### HasInitialEntryDate

`func (o *ComputerGeneral) HasInitialEntryDate() bool`

HasInitialEntryDate returns a boolean if a field has been set.

### GetDistributionPoint

`func (o *ComputerGeneral) GetDistributionPoint() string`

GetDistributionPoint returns the DistributionPoint field if non-nil, zero value otherwise.

### GetDistributionPointOk

`func (o *ComputerGeneral) GetDistributionPointOk() (*string, bool)`

GetDistributionPointOk returns a tuple with the DistributionPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionPoint

`func (o *ComputerGeneral) SetDistributionPoint(v string)`

SetDistributionPoint sets DistributionPoint field to given value.

### HasDistributionPoint

`func (o *ComputerGeneral) HasDistributionPoint() bool`

HasDistributionPoint returns a boolean if a field has been set.

### GetEnrollmentMethod

`func (o *ComputerGeneral) GetEnrollmentMethod() EnrollmentMethod`

GetEnrollmentMethod returns the EnrollmentMethod field if non-nil, zero value otherwise.

### GetEnrollmentMethodOk

`func (o *ComputerGeneral) GetEnrollmentMethodOk() (*EnrollmentMethod, bool)`

GetEnrollmentMethodOk returns a tuple with the EnrollmentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentMethod

`func (o *ComputerGeneral) SetEnrollmentMethod(v EnrollmentMethod)`

SetEnrollmentMethod sets EnrollmentMethod field to given value.

### HasEnrollmentMethod

`func (o *ComputerGeneral) HasEnrollmentMethod() bool`

HasEnrollmentMethod returns a boolean if a field has been set.

### GetSite

`func (o *ComputerGeneral) GetSite() V1Site`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ComputerGeneral) GetSiteOk() (*V1Site, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ComputerGeneral) SetSite(v V1Site)`

SetSite sets Site field to given value.

### HasSite

`func (o *ComputerGeneral) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetItunesStoreAccountActive

`func (o *ComputerGeneral) GetItunesStoreAccountActive() bool`

GetItunesStoreAccountActive returns the ItunesStoreAccountActive field if non-nil, zero value otherwise.

### GetItunesStoreAccountActiveOk

`func (o *ComputerGeneral) GetItunesStoreAccountActiveOk() (*bool, bool)`

GetItunesStoreAccountActiveOk returns a tuple with the ItunesStoreAccountActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItunesStoreAccountActive

`func (o *ComputerGeneral) SetItunesStoreAccountActive(v bool)`

SetItunesStoreAccountActive sets ItunesStoreAccountActive field to given value.

### HasItunesStoreAccountActive

`func (o *ComputerGeneral) HasItunesStoreAccountActive() bool`

HasItunesStoreAccountActive returns a boolean if a field has been set.

### GetEnrolledViaAutomatedDeviceEnrollment

`func (o *ComputerGeneral) GetEnrolledViaAutomatedDeviceEnrollment() bool`

GetEnrolledViaAutomatedDeviceEnrollment returns the EnrolledViaAutomatedDeviceEnrollment field if non-nil, zero value otherwise.

### GetEnrolledViaAutomatedDeviceEnrollmentOk

`func (o *ComputerGeneral) GetEnrolledViaAutomatedDeviceEnrollmentOk() (*bool, bool)`

GetEnrolledViaAutomatedDeviceEnrollmentOk returns a tuple with the EnrolledViaAutomatedDeviceEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolledViaAutomatedDeviceEnrollment

`func (o *ComputerGeneral) SetEnrolledViaAutomatedDeviceEnrollment(v bool)`

SetEnrolledViaAutomatedDeviceEnrollment sets EnrolledViaAutomatedDeviceEnrollment field to given value.

### HasEnrolledViaAutomatedDeviceEnrollment

`func (o *ComputerGeneral) HasEnrolledViaAutomatedDeviceEnrollment() bool`

HasEnrolledViaAutomatedDeviceEnrollment returns a boolean if a field has been set.

### GetUserApprovedMdm

`func (o *ComputerGeneral) GetUserApprovedMdm() bool`

GetUserApprovedMdm returns the UserApprovedMdm field if non-nil, zero value otherwise.

### GetUserApprovedMdmOk

`func (o *ComputerGeneral) GetUserApprovedMdmOk() (*bool, bool)`

GetUserApprovedMdmOk returns a tuple with the UserApprovedMdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserApprovedMdm

`func (o *ComputerGeneral) SetUserApprovedMdm(v bool)`

SetUserApprovedMdm sets UserApprovedMdm field to given value.

### HasUserApprovedMdm

`func (o *ComputerGeneral) HasUserApprovedMdm() bool`

HasUserApprovedMdm returns a boolean if a field has been set.

### GetDeclarativeDeviceManagementEnabled

`func (o *ComputerGeneral) GetDeclarativeDeviceManagementEnabled() bool`

GetDeclarativeDeviceManagementEnabled returns the DeclarativeDeviceManagementEnabled field if non-nil, zero value otherwise.

### GetDeclarativeDeviceManagementEnabledOk

`func (o *ComputerGeneral) GetDeclarativeDeviceManagementEnabledOk() (*bool, bool)`

GetDeclarativeDeviceManagementEnabledOk returns a tuple with the DeclarativeDeviceManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarativeDeviceManagementEnabled

`func (o *ComputerGeneral) SetDeclarativeDeviceManagementEnabled(v bool)`

SetDeclarativeDeviceManagementEnabled sets DeclarativeDeviceManagementEnabled field to given value.

### HasDeclarativeDeviceManagementEnabled

`func (o *ComputerGeneral) HasDeclarativeDeviceManagementEnabled() bool`

HasDeclarativeDeviceManagementEnabled returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *ComputerGeneral) GetExtensionAttributes() []ComputerExtensionAttribute`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *ComputerGeneral) GetExtensionAttributesOk() (*[]ComputerExtensionAttribute, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *ComputerGeneral) SetExtensionAttributes(v []ComputerExtensionAttribute)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *ComputerGeneral) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.

### GetManagementId

`func (o *ComputerGeneral) GetManagementId() string`

GetManagementId returns the ManagementId field if non-nil, zero value otherwise.

### GetManagementIdOk

`func (o *ComputerGeneral) GetManagementIdOk() (*string, bool)`

GetManagementIdOk returns a tuple with the ManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementId

`func (o *ComputerGeneral) SetManagementId(v string)`

SetManagementId sets ManagementId field to given value.

### HasManagementId

`func (o *ComputerGeneral) HasManagementId() bool`

HasManagementId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


