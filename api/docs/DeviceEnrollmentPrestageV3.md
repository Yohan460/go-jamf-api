# DeviceEnrollmentPrestageV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**Mandatory** | **bool** |  | 
**MdmRemovable** | **bool** |  | 
**SupportPhoneNumber** | **string** |  | 
**SupportEmailAddress** | **string** |  | 
**Department** | **string** |  | 
**DefaultPrestage** | **bool** |  | 
**EnrollmentSiteId** | **string** |  | 
**KeepExistingSiteMembership** | **bool** |  | 
**KeepExistingLocationInformation** | **bool** |  | 
**RequireAuthentication** | **bool** |  | 
**AuthenticationPrompt** | **string** |  | 
**PreventActivationLock** | **bool** |  | 
**EnableDeviceBasedActivationLock** | **bool** |  | 
**DeviceEnrollmentProgramInstanceId** | **string** |  | 
**SkipSetupItems** | Pointer to **map[string]bool** |  | [optional] 
**LocationInformation** | [**LocationInformationV3**](LocationInformationV3.md) |  | 
**PurchasingInformation** | [**PrestagePurchasingInformationV3**](PrestagePurchasingInformationV3.md) |  | 
**AnchorCertificates** | Pointer to **[]string** | The Base64 encoded PEM Certificate | [optional] 
**EnrollmentCustomizationId** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**AutoAdvanceSetup** | **bool** |  | 

## Methods

### NewDeviceEnrollmentPrestageV3

`func NewDeviceEnrollmentPrestageV3(displayName string, mandatory bool, mdmRemovable bool, supportPhoneNumber string, supportEmailAddress string, department string, defaultPrestage bool, enrollmentSiteId string, keepExistingSiteMembership bool, keepExistingLocationInformation bool, requireAuthentication bool, authenticationPrompt string, preventActivationLock bool, enableDeviceBasedActivationLock bool, deviceEnrollmentProgramInstanceId string, locationInformation LocationInformationV3, purchasingInformation PrestagePurchasingInformationV3, autoAdvanceSetup bool, ) *DeviceEnrollmentPrestageV3`

NewDeviceEnrollmentPrestageV3 instantiates a new DeviceEnrollmentPrestageV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceEnrollmentPrestageV3WithDefaults

`func NewDeviceEnrollmentPrestageV3WithDefaults() *DeviceEnrollmentPrestageV3`

NewDeviceEnrollmentPrestageV3WithDefaults instantiates a new DeviceEnrollmentPrestageV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *DeviceEnrollmentPrestageV3) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DeviceEnrollmentPrestageV3) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DeviceEnrollmentPrestageV3) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMandatory

`func (o *DeviceEnrollmentPrestageV3) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *DeviceEnrollmentPrestageV3) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *DeviceEnrollmentPrestageV3) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.


### GetMdmRemovable

`func (o *DeviceEnrollmentPrestageV3) GetMdmRemovable() bool`

GetMdmRemovable returns the MdmRemovable field if non-nil, zero value otherwise.

### GetMdmRemovableOk

`func (o *DeviceEnrollmentPrestageV3) GetMdmRemovableOk() (*bool, bool)`

GetMdmRemovableOk returns a tuple with the MdmRemovable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmRemovable

`func (o *DeviceEnrollmentPrestageV3) SetMdmRemovable(v bool)`

SetMdmRemovable sets MdmRemovable field to given value.


### GetSupportPhoneNumber

`func (o *DeviceEnrollmentPrestageV3) GetSupportPhoneNumber() string`

GetSupportPhoneNumber returns the SupportPhoneNumber field if non-nil, zero value otherwise.

### GetSupportPhoneNumberOk

`func (o *DeviceEnrollmentPrestageV3) GetSupportPhoneNumberOk() (*string, bool)`

GetSupportPhoneNumberOk returns a tuple with the SupportPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPhoneNumber

`func (o *DeviceEnrollmentPrestageV3) SetSupportPhoneNumber(v string)`

SetSupportPhoneNumber sets SupportPhoneNumber field to given value.


### GetSupportEmailAddress

`func (o *DeviceEnrollmentPrestageV3) GetSupportEmailAddress() string`

GetSupportEmailAddress returns the SupportEmailAddress field if non-nil, zero value otherwise.

### GetSupportEmailAddressOk

`func (o *DeviceEnrollmentPrestageV3) GetSupportEmailAddressOk() (*string, bool)`

GetSupportEmailAddressOk returns a tuple with the SupportEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmailAddress

`func (o *DeviceEnrollmentPrestageV3) SetSupportEmailAddress(v string)`

SetSupportEmailAddress sets SupportEmailAddress field to given value.


### GetDepartment

`func (o *DeviceEnrollmentPrestageV3) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *DeviceEnrollmentPrestageV3) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *DeviceEnrollmentPrestageV3) SetDepartment(v string)`

SetDepartment sets Department field to given value.


### GetDefaultPrestage

`func (o *DeviceEnrollmentPrestageV3) GetDefaultPrestage() bool`

GetDefaultPrestage returns the DefaultPrestage field if non-nil, zero value otherwise.

### GetDefaultPrestageOk

`func (o *DeviceEnrollmentPrestageV3) GetDefaultPrestageOk() (*bool, bool)`

GetDefaultPrestageOk returns a tuple with the DefaultPrestage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrestage

`func (o *DeviceEnrollmentPrestageV3) SetDefaultPrestage(v bool)`

SetDefaultPrestage sets DefaultPrestage field to given value.


### GetEnrollmentSiteId

`func (o *DeviceEnrollmentPrestageV3) GetEnrollmentSiteId() string`

GetEnrollmentSiteId returns the EnrollmentSiteId field if non-nil, zero value otherwise.

### GetEnrollmentSiteIdOk

`func (o *DeviceEnrollmentPrestageV3) GetEnrollmentSiteIdOk() (*string, bool)`

GetEnrollmentSiteIdOk returns a tuple with the EnrollmentSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSiteId

`func (o *DeviceEnrollmentPrestageV3) SetEnrollmentSiteId(v string)`

SetEnrollmentSiteId sets EnrollmentSiteId field to given value.


### GetKeepExistingSiteMembership

`func (o *DeviceEnrollmentPrestageV3) GetKeepExistingSiteMembership() bool`

GetKeepExistingSiteMembership returns the KeepExistingSiteMembership field if non-nil, zero value otherwise.

### GetKeepExistingSiteMembershipOk

`func (o *DeviceEnrollmentPrestageV3) GetKeepExistingSiteMembershipOk() (*bool, bool)`

GetKeepExistingSiteMembershipOk returns a tuple with the KeepExistingSiteMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepExistingSiteMembership

`func (o *DeviceEnrollmentPrestageV3) SetKeepExistingSiteMembership(v bool)`

SetKeepExistingSiteMembership sets KeepExistingSiteMembership field to given value.


### GetKeepExistingLocationInformation

`func (o *DeviceEnrollmentPrestageV3) GetKeepExistingLocationInformation() bool`

GetKeepExistingLocationInformation returns the KeepExistingLocationInformation field if non-nil, zero value otherwise.

### GetKeepExistingLocationInformationOk

`func (o *DeviceEnrollmentPrestageV3) GetKeepExistingLocationInformationOk() (*bool, bool)`

GetKeepExistingLocationInformationOk returns a tuple with the KeepExistingLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepExistingLocationInformation

`func (o *DeviceEnrollmentPrestageV3) SetKeepExistingLocationInformation(v bool)`

SetKeepExistingLocationInformation sets KeepExistingLocationInformation field to given value.


### GetRequireAuthentication

`func (o *DeviceEnrollmentPrestageV3) GetRequireAuthentication() bool`

GetRequireAuthentication returns the RequireAuthentication field if non-nil, zero value otherwise.

### GetRequireAuthenticationOk

`func (o *DeviceEnrollmentPrestageV3) GetRequireAuthenticationOk() (*bool, bool)`

GetRequireAuthenticationOk returns a tuple with the RequireAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuthentication

`func (o *DeviceEnrollmentPrestageV3) SetRequireAuthentication(v bool)`

SetRequireAuthentication sets RequireAuthentication field to given value.


### GetAuthenticationPrompt

`func (o *DeviceEnrollmentPrestageV3) GetAuthenticationPrompt() string`

GetAuthenticationPrompt returns the AuthenticationPrompt field if non-nil, zero value otherwise.

### GetAuthenticationPromptOk

`func (o *DeviceEnrollmentPrestageV3) GetAuthenticationPromptOk() (*string, bool)`

GetAuthenticationPromptOk returns a tuple with the AuthenticationPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPrompt

`func (o *DeviceEnrollmentPrestageV3) SetAuthenticationPrompt(v string)`

SetAuthenticationPrompt sets AuthenticationPrompt field to given value.


### GetPreventActivationLock

`func (o *DeviceEnrollmentPrestageV3) GetPreventActivationLock() bool`

GetPreventActivationLock returns the PreventActivationLock field if non-nil, zero value otherwise.

### GetPreventActivationLockOk

`func (o *DeviceEnrollmentPrestageV3) GetPreventActivationLockOk() (*bool, bool)`

GetPreventActivationLockOk returns a tuple with the PreventActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventActivationLock

`func (o *DeviceEnrollmentPrestageV3) SetPreventActivationLock(v bool)`

SetPreventActivationLock sets PreventActivationLock field to given value.


### GetEnableDeviceBasedActivationLock

`func (o *DeviceEnrollmentPrestageV3) GetEnableDeviceBasedActivationLock() bool`

GetEnableDeviceBasedActivationLock returns the EnableDeviceBasedActivationLock field if non-nil, zero value otherwise.

### GetEnableDeviceBasedActivationLockOk

`func (o *DeviceEnrollmentPrestageV3) GetEnableDeviceBasedActivationLockOk() (*bool, bool)`

GetEnableDeviceBasedActivationLockOk returns a tuple with the EnableDeviceBasedActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDeviceBasedActivationLock

`func (o *DeviceEnrollmentPrestageV3) SetEnableDeviceBasedActivationLock(v bool)`

SetEnableDeviceBasedActivationLock sets EnableDeviceBasedActivationLock field to given value.


### GetDeviceEnrollmentProgramInstanceId

`func (o *DeviceEnrollmentPrestageV3) GetDeviceEnrollmentProgramInstanceId() string`

GetDeviceEnrollmentProgramInstanceId returns the DeviceEnrollmentProgramInstanceId field if non-nil, zero value otherwise.

### GetDeviceEnrollmentProgramInstanceIdOk

`func (o *DeviceEnrollmentPrestageV3) GetDeviceEnrollmentProgramInstanceIdOk() (*string, bool)`

GetDeviceEnrollmentProgramInstanceIdOk returns a tuple with the DeviceEnrollmentProgramInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceEnrollmentProgramInstanceId

`func (o *DeviceEnrollmentPrestageV3) SetDeviceEnrollmentProgramInstanceId(v string)`

SetDeviceEnrollmentProgramInstanceId sets DeviceEnrollmentProgramInstanceId field to given value.


### GetSkipSetupItems

`func (o *DeviceEnrollmentPrestageV3) GetSkipSetupItems() map[string]bool`

GetSkipSetupItems returns the SkipSetupItems field if non-nil, zero value otherwise.

### GetSkipSetupItemsOk

`func (o *DeviceEnrollmentPrestageV3) GetSkipSetupItemsOk() (*map[string]bool, bool)`

GetSkipSetupItemsOk returns a tuple with the SkipSetupItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSetupItems

`func (o *DeviceEnrollmentPrestageV3) SetSkipSetupItems(v map[string]bool)`

SetSkipSetupItems sets SkipSetupItems field to given value.

### HasSkipSetupItems

`func (o *DeviceEnrollmentPrestageV3) HasSkipSetupItems() bool`

HasSkipSetupItems returns a boolean if a field has been set.

### GetLocationInformation

`func (o *DeviceEnrollmentPrestageV3) GetLocationInformation() LocationInformationV3`

GetLocationInformation returns the LocationInformation field if non-nil, zero value otherwise.

### GetLocationInformationOk

`func (o *DeviceEnrollmentPrestageV3) GetLocationInformationOk() (*LocationInformationV3, bool)`

GetLocationInformationOk returns a tuple with the LocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationInformation

`func (o *DeviceEnrollmentPrestageV3) SetLocationInformation(v LocationInformationV3)`

SetLocationInformation sets LocationInformation field to given value.


### GetPurchasingInformation

`func (o *DeviceEnrollmentPrestageV3) GetPurchasingInformation() PrestagePurchasingInformationV3`

GetPurchasingInformation returns the PurchasingInformation field if non-nil, zero value otherwise.

### GetPurchasingInformationOk

`func (o *DeviceEnrollmentPrestageV3) GetPurchasingInformationOk() (*PrestagePurchasingInformationV3, bool)`

GetPurchasingInformationOk returns a tuple with the PurchasingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasingInformation

`func (o *DeviceEnrollmentPrestageV3) SetPurchasingInformation(v PrestagePurchasingInformationV3)`

SetPurchasingInformation sets PurchasingInformation field to given value.


### GetAnchorCertificates

`func (o *DeviceEnrollmentPrestageV3) GetAnchorCertificates() []string`

GetAnchorCertificates returns the AnchorCertificates field if non-nil, zero value otherwise.

### GetAnchorCertificatesOk

`func (o *DeviceEnrollmentPrestageV3) GetAnchorCertificatesOk() (*[]string, bool)`

GetAnchorCertificatesOk returns a tuple with the AnchorCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorCertificates

`func (o *DeviceEnrollmentPrestageV3) SetAnchorCertificates(v []string)`

SetAnchorCertificates sets AnchorCertificates field to given value.

### HasAnchorCertificates

`func (o *DeviceEnrollmentPrestageV3) HasAnchorCertificates() bool`

HasAnchorCertificates returns a boolean if a field has been set.

### GetEnrollmentCustomizationId

`func (o *DeviceEnrollmentPrestageV3) GetEnrollmentCustomizationId() string`

GetEnrollmentCustomizationId returns the EnrollmentCustomizationId field if non-nil, zero value otherwise.

### GetEnrollmentCustomizationIdOk

`func (o *DeviceEnrollmentPrestageV3) GetEnrollmentCustomizationIdOk() (*string, bool)`

GetEnrollmentCustomizationIdOk returns a tuple with the EnrollmentCustomizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCustomizationId

`func (o *DeviceEnrollmentPrestageV3) SetEnrollmentCustomizationId(v string)`

SetEnrollmentCustomizationId sets EnrollmentCustomizationId field to given value.

### HasEnrollmentCustomizationId

`func (o *DeviceEnrollmentPrestageV3) HasEnrollmentCustomizationId() bool`

HasEnrollmentCustomizationId returns a boolean if a field has been set.

### GetLanguage

`func (o *DeviceEnrollmentPrestageV3) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *DeviceEnrollmentPrestageV3) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *DeviceEnrollmentPrestageV3) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *DeviceEnrollmentPrestageV3) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetRegion

`func (o *DeviceEnrollmentPrestageV3) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DeviceEnrollmentPrestageV3) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DeviceEnrollmentPrestageV3) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DeviceEnrollmentPrestageV3) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetAutoAdvanceSetup

`func (o *DeviceEnrollmentPrestageV3) GetAutoAdvanceSetup() bool`

GetAutoAdvanceSetup returns the AutoAdvanceSetup field if non-nil, zero value otherwise.

### GetAutoAdvanceSetupOk

`func (o *DeviceEnrollmentPrestageV3) GetAutoAdvanceSetupOk() (*bool, bool)`

GetAutoAdvanceSetupOk returns a tuple with the AutoAdvanceSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAdvanceSetup

`func (o *DeviceEnrollmentPrestageV3) SetAutoAdvanceSetup(v bool)`

SetAutoAdvanceSetup sets AutoAdvanceSetup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


