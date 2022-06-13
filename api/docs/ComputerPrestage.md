# ComputerPrestage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**IsMandatory** | **bool** |  | 
**IsMdmRemovable** | **bool** |  | 
**SupportPhoneNumber** | **string** |  | 
**SupportEmailAddress** | **string** |  | 
**Department** | **string** |  | 
**IsDefaultPrestage** | **bool** |  | 
**EnrollmentSiteId** | **int32** |  | 
**IsKeepExistingSiteMembership** | **bool** |  | 
**IsKeepExistingLocationInformation** | **bool** |  | 
**IsRequireAuthentication** | **bool** |  | 
**AuthenticationPrompt** | **string** |  | 
**IsPreventActivationLock** | **bool** |  | 
**IsEnableDeviceBasedActivationLock** | **bool** |  | 
**DeviceEnrollmentProgramInstanceId** | **int32** |  | 
**SkipSetupItems** | Pointer to **map[string]bool** |  | [optional] 
**LocationInformation** | [**LocationInformation**](LocationInformation.md) |  | 
**PurchasingInformation** | [**PrestagePurchasingInformation**](PrestagePurchasingInformation.md) |  | 
**AnchorCertificates** | Pointer to **[]string** | The Base64 encoded PEM Certificate | [optional] 
**EnrollmentCustomizationId** | Pointer to **int32** |  | [optional] 
**IsInstallProfilesDuringSetup** | **bool** |  | 
**PrestageInstalledProfileIds** | **[]int32** |  | 
**CustomPackageIds** | **[]int32** |  | 
**CustomPackageDistributionPointId** | **int32** |  | 

## Methods

### NewComputerPrestage

`func NewComputerPrestage(displayName string, isMandatory bool, isMdmRemovable bool, supportPhoneNumber string, supportEmailAddress string, department string, isDefaultPrestage bool, enrollmentSiteId int32, isKeepExistingSiteMembership bool, isKeepExistingLocationInformation bool, isRequireAuthentication bool, authenticationPrompt string, isPreventActivationLock bool, isEnableDeviceBasedActivationLock bool, deviceEnrollmentProgramInstanceId int32, locationInformation LocationInformation, purchasingInformation PrestagePurchasingInformation, isInstallProfilesDuringSetup bool, prestageInstalledProfileIds []int32, customPackageIds []int32, customPackageDistributionPointId int32, ) *ComputerPrestage`

NewComputerPrestage instantiates a new ComputerPrestage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerPrestageWithDefaults

`func NewComputerPrestageWithDefaults() *ComputerPrestage`

NewComputerPrestageWithDefaults instantiates a new ComputerPrestage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ComputerPrestage) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ComputerPrestage) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ComputerPrestage) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetIsMandatory

`func (o *ComputerPrestage) GetIsMandatory() bool`

GetIsMandatory returns the IsMandatory field if non-nil, zero value otherwise.

### GetIsMandatoryOk

`func (o *ComputerPrestage) GetIsMandatoryOk() (*bool, bool)`

GetIsMandatoryOk returns a tuple with the IsMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMandatory

`func (o *ComputerPrestage) SetIsMandatory(v bool)`

SetIsMandatory sets IsMandatory field to given value.


### GetIsMdmRemovable

`func (o *ComputerPrestage) GetIsMdmRemovable() bool`

GetIsMdmRemovable returns the IsMdmRemovable field if non-nil, zero value otherwise.

### GetIsMdmRemovableOk

`func (o *ComputerPrestage) GetIsMdmRemovableOk() (*bool, bool)`

GetIsMdmRemovableOk returns a tuple with the IsMdmRemovable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMdmRemovable

`func (o *ComputerPrestage) SetIsMdmRemovable(v bool)`

SetIsMdmRemovable sets IsMdmRemovable field to given value.


### GetSupportPhoneNumber

`func (o *ComputerPrestage) GetSupportPhoneNumber() string`

GetSupportPhoneNumber returns the SupportPhoneNumber field if non-nil, zero value otherwise.

### GetSupportPhoneNumberOk

`func (o *ComputerPrestage) GetSupportPhoneNumberOk() (*string, bool)`

GetSupportPhoneNumberOk returns a tuple with the SupportPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPhoneNumber

`func (o *ComputerPrestage) SetSupportPhoneNumber(v string)`

SetSupportPhoneNumber sets SupportPhoneNumber field to given value.


### GetSupportEmailAddress

`func (o *ComputerPrestage) GetSupportEmailAddress() string`

GetSupportEmailAddress returns the SupportEmailAddress field if non-nil, zero value otherwise.

### GetSupportEmailAddressOk

`func (o *ComputerPrestage) GetSupportEmailAddressOk() (*string, bool)`

GetSupportEmailAddressOk returns a tuple with the SupportEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmailAddress

`func (o *ComputerPrestage) SetSupportEmailAddress(v string)`

SetSupportEmailAddress sets SupportEmailAddress field to given value.


### GetDepartment

`func (o *ComputerPrestage) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *ComputerPrestage) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *ComputerPrestage) SetDepartment(v string)`

SetDepartment sets Department field to given value.


### GetIsDefaultPrestage

`func (o *ComputerPrestage) GetIsDefaultPrestage() bool`

GetIsDefaultPrestage returns the IsDefaultPrestage field if non-nil, zero value otherwise.

### GetIsDefaultPrestageOk

`func (o *ComputerPrestage) GetIsDefaultPrestageOk() (*bool, bool)`

GetIsDefaultPrestageOk returns a tuple with the IsDefaultPrestage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultPrestage

`func (o *ComputerPrestage) SetIsDefaultPrestage(v bool)`

SetIsDefaultPrestage sets IsDefaultPrestage field to given value.


### GetEnrollmentSiteId

`func (o *ComputerPrestage) GetEnrollmentSiteId() int32`

GetEnrollmentSiteId returns the EnrollmentSiteId field if non-nil, zero value otherwise.

### GetEnrollmentSiteIdOk

`func (o *ComputerPrestage) GetEnrollmentSiteIdOk() (*int32, bool)`

GetEnrollmentSiteIdOk returns a tuple with the EnrollmentSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSiteId

`func (o *ComputerPrestage) SetEnrollmentSiteId(v int32)`

SetEnrollmentSiteId sets EnrollmentSiteId field to given value.


### GetIsKeepExistingSiteMembership

`func (o *ComputerPrestage) GetIsKeepExistingSiteMembership() bool`

GetIsKeepExistingSiteMembership returns the IsKeepExistingSiteMembership field if non-nil, zero value otherwise.

### GetIsKeepExistingSiteMembershipOk

`func (o *ComputerPrestage) GetIsKeepExistingSiteMembershipOk() (*bool, bool)`

GetIsKeepExistingSiteMembershipOk returns a tuple with the IsKeepExistingSiteMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKeepExistingSiteMembership

`func (o *ComputerPrestage) SetIsKeepExistingSiteMembership(v bool)`

SetIsKeepExistingSiteMembership sets IsKeepExistingSiteMembership field to given value.


### GetIsKeepExistingLocationInformation

`func (o *ComputerPrestage) GetIsKeepExistingLocationInformation() bool`

GetIsKeepExistingLocationInformation returns the IsKeepExistingLocationInformation field if non-nil, zero value otherwise.

### GetIsKeepExistingLocationInformationOk

`func (o *ComputerPrestage) GetIsKeepExistingLocationInformationOk() (*bool, bool)`

GetIsKeepExistingLocationInformationOk returns a tuple with the IsKeepExistingLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKeepExistingLocationInformation

`func (o *ComputerPrestage) SetIsKeepExistingLocationInformation(v bool)`

SetIsKeepExistingLocationInformation sets IsKeepExistingLocationInformation field to given value.


### GetIsRequireAuthentication

`func (o *ComputerPrestage) GetIsRequireAuthentication() bool`

GetIsRequireAuthentication returns the IsRequireAuthentication field if non-nil, zero value otherwise.

### GetIsRequireAuthenticationOk

`func (o *ComputerPrestage) GetIsRequireAuthenticationOk() (*bool, bool)`

GetIsRequireAuthenticationOk returns a tuple with the IsRequireAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequireAuthentication

`func (o *ComputerPrestage) SetIsRequireAuthentication(v bool)`

SetIsRequireAuthentication sets IsRequireAuthentication field to given value.


### GetAuthenticationPrompt

`func (o *ComputerPrestage) GetAuthenticationPrompt() string`

GetAuthenticationPrompt returns the AuthenticationPrompt field if non-nil, zero value otherwise.

### GetAuthenticationPromptOk

`func (o *ComputerPrestage) GetAuthenticationPromptOk() (*string, bool)`

GetAuthenticationPromptOk returns a tuple with the AuthenticationPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPrompt

`func (o *ComputerPrestage) SetAuthenticationPrompt(v string)`

SetAuthenticationPrompt sets AuthenticationPrompt field to given value.


### GetIsPreventActivationLock

`func (o *ComputerPrestage) GetIsPreventActivationLock() bool`

GetIsPreventActivationLock returns the IsPreventActivationLock field if non-nil, zero value otherwise.

### GetIsPreventActivationLockOk

`func (o *ComputerPrestage) GetIsPreventActivationLockOk() (*bool, bool)`

GetIsPreventActivationLockOk returns a tuple with the IsPreventActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreventActivationLock

`func (o *ComputerPrestage) SetIsPreventActivationLock(v bool)`

SetIsPreventActivationLock sets IsPreventActivationLock field to given value.


### GetIsEnableDeviceBasedActivationLock

`func (o *ComputerPrestage) GetIsEnableDeviceBasedActivationLock() bool`

GetIsEnableDeviceBasedActivationLock returns the IsEnableDeviceBasedActivationLock field if non-nil, zero value otherwise.

### GetIsEnableDeviceBasedActivationLockOk

`func (o *ComputerPrestage) GetIsEnableDeviceBasedActivationLockOk() (*bool, bool)`

GetIsEnableDeviceBasedActivationLockOk returns a tuple with the IsEnableDeviceBasedActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnableDeviceBasedActivationLock

`func (o *ComputerPrestage) SetIsEnableDeviceBasedActivationLock(v bool)`

SetIsEnableDeviceBasedActivationLock sets IsEnableDeviceBasedActivationLock field to given value.


### GetDeviceEnrollmentProgramInstanceId

`func (o *ComputerPrestage) GetDeviceEnrollmentProgramInstanceId() int32`

GetDeviceEnrollmentProgramInstanceId returns the DeviceEnrollmentProgramInstanceId field if non-nil, zero value otherwise.

### GetDeviceEnrollmentProgramInstanceIdOk

`func (o *ComputerPrestage) GetDeviceEnrollmentProgramInstanceIdOk() (*int32, bool)`

GetDeviceEnrollmentProgramInstanceIdOk returns a tuple with the DeviceEnrollmentProgramInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceEnrollmentProgramInstanceId

`func (o *ComputerPrestage) SetDeviceEnrollmentProgramInstanceId(v int32)`

SetDeviceEnrollmentProgramInstanceId sets DeviceEnrollmentProgramInstanceId field to given value.


### GetSkipSetupItems

`func (o *ComputerPrestage) GetSkipSetupItems() map[string]bool`

GetSkipSetupItems returns the SkipSetupItems field if non-nil, zero value otherwise.

### GetSkipSetupItemsOk

`func (o *ComputerPrestage) GetSkipSetupItemsOk() (*map[string]bool, bool)`

GetSkipSetupItemsOk returns a tuple with the SkipSetupItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSetupItems

`func (o *ComputerPrestage) SetSkipSetupItems(v map[string]bool)`

SetSkipSetupItems sets SkipSetupItems field to given value.

### HasSkipSetupItems

`func (o *ComputerPrestage) HasSkipSetupItems() bool`

HasSkipSetupItems returns a boolean if a field has been set.

### GetLocationInformation

`func (o *ComputerPrestage) GetLocationInformation() LocationInformation`

GetLocationInformation returns the LocationInformation field if non-nil, zero value otherwise.

### GetLocationInformationOk

`func (o *ComputerPrestage) GetLocationInformationOk() (*LocationInformation, bool)`

GetLocationInformationOk returns a tuple with the LocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationInformation

`func (o *ComputerPrestage) SetLocationInformation(v LocationInformation)`

SetLocationInformation sets LocationInformation field to given value.


### GetPurchasingInformation

`func (o *ComputerPrestage) GetPurchasingInformation() PrestagePurchasingInformation`

GetPurchasingInformation returns the PurchasingInformation field if non-nil, zero value otherwise.

### GetPurchasingInformationOk

`func (o *ComputerPrestage) GetPurchasingInformationOk() (*PrestagePurchasingInformation, bool)`

GetPurchasingInformationOk returns a tuple with the PurchasingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingInformation

`func (o *ComputerPrestage) SetPurchasingInformation(v PrestagePurchasingInformation)`

SetPurchasingInformation sets PurchasingInformation field to given value.


### GetAnchorCertificates

`func (o *ComputerPrestage) GetAnchorCertificates() []string`

GetAnchorCertificates returns the AnchorCertificates field if non-nil, zero value otherwise.

### GetAnchorCertificatesOk

`func (o *ComputerPrestage) GetAnchorCertificatesOk() (*[]string, bool)`

GetAnchorCertificatesOk returns a tuple with the AnchorCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorCertificates

`func (o *ComputerPrestage) SetAnchorCertificates(v []string)`

SetAnchorCertificates sets AnchorCertificates field to given value.

### HasAnchorCertificates

`func (o *ComputerPrestage) HasAnchorCertificates() bool`

HasAnchorCertificates returns a boolean if a field has been set.

### GetEnrollmentCustomizationId

`func (o *ComputerPrestage) GetEnrollmentCustomizationId() int32`

GetEnrollmentCustomizationId returns the EnrollmentCustomizationId field if non-nil, zero value otherwise.

### GetEnrollmentCustomizationIdOk

`func (o *ComputerPrestage) GetEnrollmentCustomizationIdOk() (*int32, bool)`

GetEnrollmentCustomizationIdOk returns a tuple with the EnrollmentCustomizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCustomizationId

`func (o *ComputerPrestage) SetEnrollmentCustomizationId(v int32)`

SetEnrollmentCustomizationId sets EnrollmentCustomizationId field to given value.

### HasEnrollmentCustomizationId

`func (o *ComputerPrestage) HasEnrollmentCustomizationId() bool`

HasEnrollmentCustomizationId returns a boolean if a field has been set.

### GetIsInstallProfilesDuringSetup

`func (o *ComputerPrestage) GetIsInstallProfilesDuringSetup() bool`

GetIsInstallProfilesDuringSetup returns the IsInstallProfilesDuringSetup field if non-nil, zero value otherwise.

### GetIsInstallProfilesDuringSetupOk

`func (o *ComputerPrestage) GetIsInstallProfilesDuringSetupOk() (*bool, bool)`

GetIsInstallProfilesDuringSetupOk returns a tuple with the IsInstallProfilesDuringSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInstallProfilesDuringSetup

`func (o *ComputerPrestage) SetIsInstallProfilesDuringSetup(v bool)`

SetIsInstallProfilesDuringSetup sets IsInstallProfilesDuringSetup field to given value.


### GetPrestageInstalledProfileIds

`func (o *ComputerPrestage) GetPrestageInstalledProfileIds() []int32`

GetPrestageInstalledProfileIds returns the PrestageInstalledProfileIds field if non-nil, zero value otherwise.

### GetPrestageInstalledProfileIdsOk

`func (o *ComputerPrestage) GetPrestageInstalledProfileIdsOk() (*[]int32, bool)`

GetPrestageInstalledProfileIdsOk returns a tuple with the PrestageInstalledProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrestageInstalledProfileIds

`func (o *ComputerPrestage) SetPrestageInstalledProfileIds(v []int32)`

SetPrestageInstalledProfileIds sets PrestageInstalledProfileIds field to given value.


### GetCustomPackageIds

`func (o *ComputerPrestage) GetCustomPackageIds() []int32`

GetCustomPackageIds returns the CustomPackageIds field if non-nil, zero value otherwise.

### GetCustomPackageIdsOk

`func (o *ComputerPrestage) GetCustomPackageIdsOk() (*[]int32, bool)`

GetCustomPackageIdsOk returns a tuple with the CustomPackageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPackageIds

`func (o *ComputerPrestage) SetCustomPackageIds(v []int32)`

SetCustomPackageIds sets CustomPackageIds field to given value.


### GetCustomPackageDistributionPointId

`func (o *ComputerPrestage) GetCustomPackageDistributionPointId() int32`

GetCustomPackageDistributionPointId returns the CustomPackageDistributionPointId field if non-nil, zero value otherwise.

### GetCustomPackageDistributionPointIdOk

`func (o *ComputerPrestage) GetCustomPackageDistributionPointIdOk() (*int32, bool)`

GetCustomPackageDistributionPointIdOk returns a tuple with the CustomPackageDistributionPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPackageDistributionPointId

`func (o *ComputerPrestage) SetCustomPackageDistributionPointId(v int32)`

SetCustomPackageDistributionPointId sets CustomPackageDistributionPointId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


