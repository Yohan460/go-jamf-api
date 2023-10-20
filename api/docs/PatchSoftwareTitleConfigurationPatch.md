# PatchSoftwareTitleConfigurationPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**CategoryId** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**UiNotifications** | Pointer to **bool** |  | [optional] 
**EmailNotifications** | Pointer to **bool** |  | [optional] 
**SoftwareTitleId** | Pointer to **string** |  | [optional] 
**Packages** | Pointer to [**[]PatchSoftwareTitlePackages**](PatchSoftwareTitlePackages.md) |  | [optional] 
**ExtensionAttributes** | Pointer to [**[]PatchSoftwareTitleConfigurationExtensionAttributes**](PatchSoftwareTitleConfigurationExtensionAttributes.md) |  | [optional] 

## Methods

### NewPatchSoftwareTitleConfigurationPatch

`func NewPatchSoftwareTitleConfigurationPatch() *PatchSoftwareTitleConfigurationPatch`

NewPatchSoftwareTitleConfigurationPatch instantiates a new PatchSoftwareTitleConfigurationPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchSoftwareTitleConfigurationPatchWithDefaults

`func NewPatchSoftwareTitleConfigurationPatchWithDefaults() *PatchSoftwareTitleConfigurationPatch`

NewPatchSoftwareTitleConfigurationPatchWithDefaults instantiates a new PatchSoftwareTitleConfigurationPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *PatchSoftwareTitleConfigurationPatch) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PatchSoftwareTitleConfigurationPatch) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PatchSoftwareTitleConfigurationPatch) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PatchSoftwareTitleConfigurationPatch) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCategoryId

`func (o *PatchSoftwareTitleConfigurationPatch) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *PatchSoftwareTitleConfigurationPatch) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *PatchSoftwareTitleConfigurationPatch) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *PatchSoftwareTitleConfigurationPatch) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetSiteId

`func (o *PatchSoftwareTitleConfigurationPatch) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *PatchSoftwareTitleConfigurationPatch) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *PatchSoftwareTitleConfigurationPatch) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *PatchSoftwareTitleConfigurationPatch) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetUiNotifications

`func (o *PatchSoftwareTitleConfigurationPatch) GetUiNotifications() bool`

GetUiNotifications returns the UiNotifications field if non-nil, zero value otherwise.

### GetUiNotificationsOk

`func (o *PatchSoftwareTitleConfigurationPatch) GetUiNotificationsOk() (*bool, bool)`

GetUiNotificationsOk returns a tuple with the UiNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiNotifications

`func (o *PatchSoftwareTitleConfigurationPatch) SetUiNotifications(v bool)`

SetUiNotifications sets UiNotifications field to given value.

### HasUiNotifications

`func (o *PatchSoftwareTitleConfigurationPatch) HasUiNotifications() bool`

HasUiNotifications returns a boolean if a field has been set.

### GetEmailNotifications

`func (o *PatchSoftwareTitleConfigurationPatch) GetEmailNotifications() bool`

GetEmailNotifications returns the EmailNotifications field if non-nil, zero value otherwise.

### GetEmailNotificationsOk

`func (o *PatchSoftwareTitleConfigurationPatch) GetEmailNotificationsOk() (*bool, bool)`

GetEmailNotificationsOk returns a tuple with the EmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotifications

`func (o *PatchSoftwareTitleConfigurationPatch) SetEmailNotifications(v bool)`

SetEmailNotifications sets EmailNotifications field to given value.

### HasEmailNotifications

`func (o *PatchSoftwareTitleConfigurationPatch) HasEmailNotifications() bool`

HasEmailNotifications returns a boolean if a field has been set.

### GetSoftwareTitleId

`func (o *PatchSoftwareTitleConfigurationPatch) GetSoftwareTitleId() string`

GetSoftwareTitleId returns the SoftwareTitleId field if non-nil, zero value otherwise.

### GetSoftwareTitleIdOk

`func (o *PatchSoftwareTitleConfigurationPatch) GetSoftwareTitleIdOk() (*string, bool)`

GetSoftwareTitleIdOk returns a tuple with the SoftwareTitleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTitleId

`func (o *PatchSoftwareTitleConfigurationPatch) SetSoftwareTitleId(v string)`

SetSoftwareTitleId sets SoftwareTitleId field to given value.

### HasSoftwareTitleId

`func (o *PatchSoftwareTitleConfigurationPatch) HasSoftwareTitleId() bool`

HasSoftwareTitleId returns a boolean if a field has been set.

### GetPackages

`func (o *PatchSoftwareTitleConfigurationPatch) GetPackages() []PatchSoftwareTitlePackages`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *PatchSoftwareTitleConfigurationPatch) GetPackagesOk() (*[]PatchSoftwareTitlePackages, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *PatchSoftwareTitleConfigurationPatch) SetPackages(v []PatchSoftwareTitlePackages)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *PatchSoftwareTitleConfigurationPatch) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *PatchSoftwareTitleConfigurationPatch) GetExtensionAttributes() []PatchSoftwareTitleConfigurationExtensionAttributes`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *PatchSoftwareTitleConfigurationPatch) GetExtensionAttributesOk() (*[]PatchSoftwareTitleConfigurationExtensionAttributes, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *PatchSoftwareTitleConfigurationPatch) SetExtensionAttributes(v []PatchSoftwareTitleConfigurationExtensionAttributes)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *PatchSoftwareTitleConfigurationPatch) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


