# DeviceEnrollmentPrestage

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
**EnrollmentSiteId** | **int64** |  | 
**IsKeepExistingSiteMembership** | **bool** |  | 
**IsKeepExistingLocationInformation** | **bool** |  | 
**IsRequireAuthentication** | **bool** |  | 
**AuthenticationPrompt** | **string** |  | 
**IsPreventActivationLock** | **bool** |  | 
**IsEnableDeviceBasedActivationLock** | **bool** |  | 
**DeviceEnrollmentProgramInstanceId** | **int64** |  | 
**SkipSetupItems** | Pointer to **map[string]bool** |  | [optional] 
**LocationInformation** | [**LocationInformation**](LocationInformation.md) |  | 
**PurchasingInformation** | [**PrestagePurchasingInformation**](PrestagePurchasingInformation.md) |  | 
**AnchorCertificates** | Pointer to **[]string** | The Base64 encoded PEM Certificate | [optional] 
**EnrollmentCustomizationId** | Pointer to **int64** |  | [optional] 

## Methods

### NewDeviceEnrollmentPrestage

`func NewDeviceEnrollmentPrestage(displayName string, isMandatory bool, isMdmRemovable bool, supportPhoneNumber string, supportEmailAddress string, department string, isDefaultPrestage bool, enrollmentSiteId int64, isKeepExistingSiteMembership bool, isKeepExistingLocationInformation bool, isRequireAuthentication bool, authenticationPrompt string, isPreventActivationLock bool, isEnableDeviceBasedActivationLock bool, deviceEnrollmentProgramInstanceId int64, locationInformation LocationInformation, purchasingInformation PrestagePurchasingInformation, ) *DeviceEnrollmentPrestage`

NewDeviceEnrollmentPrestage instantiates a new DeviceEnrollmentPrestage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceEnrollmentPrestageWithDefaults

`func NewDeviceEnrollmentPrestageWithDefaults() *DeviceEnrollmentPrestage`

NewDeviceEnrollmentPrestageWithDefaults instantiates a new DeviceEnrollmentPrestage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *DeviceEnrollmentPrestage) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DeviceEnrollmentPrestage) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DeviceEnrollmentPrestage) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetIsMandatory

`func (o *DeviceEnrollmentPrestage) GetIsMandatory() bool`

GetIsMandatory returns the IsMandatory field if non-nil, zero value otherwise.

### GetIsMandatoryOk

`func (o *DeviceEnrollmentPrestage) GetIsMandatoryOk() (*bool, bool)`

GetIsMandatoryOk returns a tuple with the IsMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMandatory

`func (o *DeviceEnrollmentPrestage) SetIsMandatory(v bool)`

SetIsMandatory sets IsMandatory field to given value.


### GetIsMdmRemovable

`func (o *DeviceEnrollmentPrestage) GetIsMdmRemovable() bool`

GetIsMdmRemovable returns the IsMdmRemovable field if non-nil, zero value otherwise.

### GetIsMdmRemovableOk

`func (o *DeviceEnrollmentPrestage) GetIsMdmRemovableOk() (*bool, bool)`

GetIsMdmRemovableOk returns a tuple with the IsMdmRemovable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMdmRemovable

`func (o *DeviceEnrollmentPrestage) SetIsMdmRemovable(v bool)`

SetIsMdmRemovable sets IsMdmRemovable field to given value.


### GetSupportPhoneNumber

`func (o *DeviceEnrollmentPrestage) GetSupportPhoneNumber() string`

GetSupportPhoneNumber returns the SupportPhoneNumber field if non-nil, zero value otherwise.

### GetSupportPhoneNumberOk

`func (o *DeviceEnrollmentPrestage) GetSupportPhoneNumberOk() (*string, bool)`

GetSupportPhoneNumberOk returns a tuple with the SupportPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPhoneNumber

`func (o *DeviceEnrollmentPrestage) SetSupportPhoneNumber(v string)`

SetSupportPhoneNumber sets SupportPhoneNumber field to given value.


### GetSupportEmailAddress

`func (o *DeviceEnrollmentPrestage) GetSupportEmailAddress() string`

GetSupportEmailAddress returns the SupportEmailAddress field if non-nil, zero value otherwise.

### GetSupportEmailAddressOk

`func (o *DeviceEnrollmentPrestage) GetSupportEmailAddressOk() (*string, bool)`

GetSupportEmailAddressOk returns a tuple with the SupportEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmailAddress

`func (o *DeviceEnrollmentPrestage) SetSupportEmailAddress(v string)`

SetSupportEmailAddress sets SupportEmailAddress field to given value.


### GetDepartment

`func (o *DeviceEnrollmentPrestage) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *DeviceEnrollmentPrestage) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *DeviceEnrollmentPrestage) SetDepartment(v string)`

SetDepartment sets Department field to given value.


### GetIsDefaultPrestage

`func (o *DeviceEnrollmentPrestage) GetIsDefaultPrestage() bool`

GetIsDefaultPrestage returns the IsDefaultPrestage field if non-nil, zero value otherwise.

### GetIsDefaultPrestageOk

`func (o *DeviceEnrollmentPrestage) GetIsDefaultPrestageOk() (*bool, bool)`

GetIsDefaultPrestageOk returns a tuple with the IsDefaultPrestage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultPrestage

`func (o *DeviceEnrollmentPrestage) SetIsDefaultPrestage(v bool)`

SetIsDefaultPrestage sets IsDefaultPrestage field to given value.


### GetEnrollmentSiteId

`func (o *DeviceEnrollmentPrestage) GetEnrollmentSiteId() int64`

GetEnrollmentSiteId returns the EnrollmentSiteId field if non-nil, zero value otherwise.

### GetEnrollmentSiteIdOk

`func (o *DeviceEnrollmentPrestage) GetEnrollmentSiteIdOk() (*int64, bool)`

GetEnrollmentSiteIdOk returns a tuple with the EnrollmentSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSiteId

`func (o *DeviceEnrollmentPrestage) SetEnrollmentSiteId(v int64)`

SetEnrollmentSiteId sets EnrollmentSiteId field to given value.


### GetIsKeepExistingSiteMembership

`func (o *DeviceEnrollmentPrestage) GetIsKeepExistingSiteMembership() bool`

GetIsKeepExistingSiteMembership returns the IsKeepExistingSiteMembership field if non-nil, zero value otherwise.

### GetIsKeepExistingSiteMembershipOk

`func (o *DeviceEnrollmentPrestage) GetIsKeepExistingSiteMembershipOk() (*bool, bool)`

GetIsKeepExistingSiteMembershipOk returns a tuple with the IsKeepExistingSiteMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKeepExistingSiteMembership

`func (o *DeviceEnrollmentPrestage) SetIsKeepExistingSiteMembership(v bool)`

SetIsKeepExistingSiteMembership sets IsKeepExistingSiteMembership field to given value.


### GetIsKeepExistingLocationInformation

`func (o *DeviceEnrollmentPrestage) GetIsKeepExistingLocationInformation() bool`

GetIsKeepExistingLocationInformation returns the IsKeepExistingLocationInformation field if non-nil, zero value otherwise.

### GetIsKeepExistingLocationInformationOk

`func (o *DeviceEnrollmentPrestage) GetIsKeepExistingLocationInformationOk() (*bool, bool)`

GetIsKeepExistingLocationInformationOk returns a tuple with the IsKeepExistingLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKeepExistingLocationInformation

`func (o *DeviceEnrollmentPrestage) SetIsKeepExistingLocationInformation(v bool)`

SetIsKeepExistingLocationInformation sets IsKeepExistingLocationInformation field to given value.


### GetIsRequireAuthentication

`func (o *DeviceEnrollmentPrestage) GetIsRequireAuthentication() bool`

GetIsRequireAuthentication returns the IsRequireAuthentication field if non-nil, zero value otherwise.

### GetIsRequireAuthenticationOk

`func (o *DeviceEnrollmentPrestage) GetIsRequireAuthenticationOk() (*bool, bool)`

GetIsRequireAuthenticationOk returns a tuple with the IsRequireAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequireAuthentication

`func (o *DeviceEnrollmentPrestage) SetIsRequireAuthentication(v bool)`

SetIsRequireAuthentication sets IsRequireAuthentication field to given value.


### GetAuthenticationPrompt

`func (o *DeviceEnrollmentPrestage) GetAuthenticationPrompt() string`

GetAuthenticationPrompt returns the AuthenticationPrompt field if non-nil, zero value otherwise.

### GetAuthenticationPromptOk

`func (o *DeviceEnrollmentPrestage) GetAuthenticationPromptOk() (*string, bool)`

GetAuthenticationPromptOk returns a tuple with the AuthenticationPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPrompt

`func (o *DeviceEnrollmentPrestage) SetAuthenticationPrompt(v string)`

SetAuthenticationPrompt sets AuthenticationPrompt field to given value.


### GetIsPreventActivationLock

`func (o *DeviceEnrollmentPrestage) GetIsPreventActivationLock() bool`

GetIsPreventActivationLock returns the IsPreventActivationLock field if non-nil, zero value otherwise.

### GetIsPreventActivationLockOk

`func (o *DeviceEnrollmentPrestage) GetIsPreventActivationLockOk() (*bool, bool)`

GetIsPreventActivationLockOk returns a tuple with the IsPreventActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreventActivationLock

`func (o *DeviceEnrollmentPrestage) SetIsPreventActivationLock(v bool)`

SetIsPreventActivationLock sets IsPreventActivationLock field to given value.


### GetIsEnableDeviceBasedActivationLock

`func (o *DeviceEnrollmentPrestage) GetIsEnableDeviceBasedActivationLock() bool`

GetIsEnableDeviceBasedActivationLock returns the IsEnableDeviceBasedActivationLock field if non-nil, zero value otherwise.

### GetIsEnableDeviceBasedActivationLockOk

`func (o *DeviceEnrollmentPrestage) GetIsEnableDeviceBasedActivationLockOk() (*bool, bool)`

GetIsEnableDeviceBasedActivationLockOk returns a tuple with the IsEnableDeviceBasedActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnableDeviceBasedActivationLock

`func (o *DeviceEnrollmentPrestage) SetIsEnableDeviceBasedActivationLock(v bool)`

SetIsEnableDeviceBasedActivationLock sets IsEnableDeviceBasedActivationLock field to given value.


### GetDeviceEnrollmentProgramInstanceId

`func (o *DeviceEnrollmentPrestage) GetDeviceEnrollmentProgramInstanceId() int64`

GetDeviceEnrollmentProgramInstanceId returns the DeviceEnrollmentProgramInstanceId field if non-nil, zero value otherwise.

### GetDeviceEnrollmentProgramInstanceIdOk

`func (o *DeviceEnrollmentPrestage) GetDeviceEnrollmentProgramInstanceIdOk() (*int64, bool)`

GetDeviceEnrollmentProgramInstanceIdOk returns a tuple with the DeviceEnrollmentProgramInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceEnrollmentProgramInstanceId

`func (o *DeviceEnrollmentPrestage) SetDeviceEnrollmentProgramInstanceId(v int64)`

SetDeviceEnrollmentProgramInstanceId sets DeviceEnrollmentProgramInstanceId field to given value.


### GetSkipSetupItems

`func (o *DeviceEnrollmentPrestage) GetSkipSetupItems() map[string]bool`

GetSkipSetupItems returns the SkipSetupItems field if non-nil, zero value otherwise.

### GetSkipSetupItemsOk

`func (o *DeviceEnrollmentPrestage) GetSkipSetupItemsOk() (*map[string]bool, bool)`

GetSkipSetupItemsOk returns a tuple with the SkipSetupItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSetupItems

`func (o *DeviceEnrollmentPrestage) SetSkipSetupItems(v map[string]bool)`

SetSkipSetupItems sets SkipSetupItems field to given value.

### HasSkipSetupItems

`func (o *DeviceEnrollmentPrestage) HasSkipSetupItems() bool`

HasSkipSetupItems returns a boolean if a field has been set.

### GetLocationInformation

`func (o *DeviceEnrollmentPrestage) GetLocationInformation() LocationInformation`

GetLocationInformation returns the LocationInformation field if non-nil, zero value otherwise.

### GetLocationInformationOk

`func (o *DeviceEnrollmentPrestage) GetLocationInformationOk() (*LocationInformation, bool)`

GetLocationInformationOk returns a tuple with the LocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationInformation

`func (o *DeviceEnrollmentPrestage) SetLocationInformation(v LocationInformation)`

SetLocationInformation sets LocationInformation field to given value.


### GetPurchasingInformation

`func (o *DeviceEnrollmentPrestage) GetPurchasingInformation() PrestagePurchasingInformation`

GetPurchasingInformation returns the PurchasingInformation field if non-nil, zero value otherwise.

### GetPurchasingInformationOk

`func (o *DeviceEnrollmentPrestage) GetPurchasingInformationOk() (*PrestagePurchasingInformation, bool)`

GetPurchasingInformationOk returns a tuple with the PurchasingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingInformation

`func (o *DeviceEnrollmentPrestage) SetPurchasingInformation(v PrestagePurchasingInformation)`

SetPurchasingInformation sets PurchasingInformation field to given value.


### GetAnchorCertificates

`func (o *DeviceEnrollmentPrestage) GetAnchorCertificates() []string`

GetAnchorCertificates returns the AnchorCertificates field if non-nil, zero value otherwise.

### GetAnchorCertificatesOk

`func (o *DeviceEnrollmentPrestage) GetAnchorCertificatesOk() (*[]string, bool)`

GetAnchorCertificatesOk returns a tuple with the AnchorCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorCertificates

`func (o *DeviceEnrollmentPrestage) SetAnchorCertificates(v []string)`

SetAnchorCertificates sets AnchorCertificates field to given value.

### HasAnchorCertificates

`func (o *DeviceEnrollmentPrestage) HasAnchorCertificates() bool`

HasAnchorCertificates returns a boolean if a field has been set.

### GetEnrollmentCustomizationId

`func (o *DeviceEnrollmentPrestage) GetEnrollmentCustomizationId() int64`

GetEnrollmentCustomizationId returns the EnrollmentCustomizationId field if non-nil, zero value otherwise.

### GetEnrollmentCustomizationIdOk

`func (o *DeviceEnrollmentPrestage) GetEnrollmentCustomizationIdOk() (*int64, bool)`

GetEnrollmentCustomizationIdOk returns a tuple with the EnrollmentCustomizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCustomizationId

`func (o *DeviceEnrollmentPrestage) SetEnrollmentCustomizationId(v int64)`

SetEnrollmentCustomizationId sets EnrollmentCustomizationId field to given value.

### HasEnrollmentCustomizationId

`func (o *DeviceEnrollmentPrestage) HasEnrollmentCustomizationId() bool`

HasEnrollmentCustomizationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


