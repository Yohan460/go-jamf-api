# PatchSoftwareTitleConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**CategoryId** | Pointer to **string** |  | [optional] [default to "-1"]
**SiteId** | Pointer to **string** |  | [optional] [default to "-1"]
**UiNotifications** | Pointer to **bool** |  | [optional] [default to false]
**EmailNotifications** | Pointer to **bool** |  | [optional] [default to false]
**SoftwareTitleId** | **string** |  | 
**JamfOfficial** | Pointer to **bool** |  | [optional] [readonly] 
**ExtensionAttributes** | Pointer to [**[]PatchSoftwareTitleConfigurationExtensionAttributes**](PatchSoftwareTitleConfigurationExtensionAttributes.md) |  | [optional] 
**SoftwareTitleName** | Pointer to **string** |  | [optional] [readonly] 
**SoftwareTitleNameId** | Pointer to **string** |  | [optional] [readonly] 
**SoftwareTitlePublisher** | Pointer to **string** |  | [optional] [readonly] 
**PatchSourceName** | Pointer to **string** |  | [optional] [readonly] 
**PatchSourceEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] 
**Packages** | Pointer to [**[]PatchSoftwareTitlePackages**](PatchSoftwareTitlePackages.md) |  | [optional] 

## Methods

### NewPatchSoftwareTitleConfiguration

`func NewPatchSoftwareTitleConfiguration(displayName string, softwareTitleId string, ) *PatchSoftwareTitleConfiguration`

NewPatchSoftwareTitleConfiguration instantiates a new PatchSoftwareTitleConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchSoftwareTitleConfigurationWithDefaults

`func NewPatchSoftwareTitleConfigurationWithDefaults() *PatchSoftwareTitleConfiguration`

NewPatchSoftwareTitleConfigurationWithDefaults instantiates a new PatchSoftwareTitleConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *PatchSoftwareTitleConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PatchSoftwareTitleConfiguration) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PatchSoftwareTitleConfiguration) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetCategoryId

`func (o *PatchSoftwareTitleConfiguration) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *PatchSoftwareTitleConfiguration) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *PatchSoftwareTitleConfiguration) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *PatchSoftwareTitleConfiguration) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetSiteId

`func (o *PatchSoftwareTitleConfiguration) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *PatchSoftwareTitleConfiguration) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *PatchSoftwareTitleConfiguration) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *PatchSoftwareTitleConfiguration) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetUiNotifications

`func (o *PatchSoftwareTitleConfiguration) GetUiNotifications() bool`

GetUiNotifications returns the UiNotifications field if non-nil, zero value otherwise.

### GetUiNotificationsOk

`func (o *PatchSoftwareTitleConfiguration) GetUiNotificationsOk() (*bool, bool)`

GetUiNotificationsOk returns a tuple with the UiNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiNotifications

`func (o *PatchSoftwareTitleConfiguration) SetUiNotifications(v bool)`

SetUiNotifications sets UiNotifications field to given value.

### HasUiNotifications

`func (o *PatchSoftwareTitleConfiguration) HasUiNotifications() bool`

HasUiNotifications returns a boolean if a field has been set.

### GetEmailNotifications

`func (o *PatchSoftwareTitleConfiguration) GetEmailNotifications() bool`

GetEmailNotifications returns the EmailNotifications field if non-nil, zero value otherwise.

### GetEmailNotificationsOk

`func (o *PatchSoftwareTitleConfiguration) GetEmailNotificationsOk() (*bool, bool)`

GetEmailNotificationsOk returns a tuple with the EmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotifications

`func (o *PatchSoftwareTitleConfiguration) SetEmailNotifications(v bool)`

SetEmailNotifications sets EmailNotifications field to given value.

### HasEmailNotifications

`func (o *PatchSoftwareTitleConfiguration) HasEmailNotifications() bool`

HasEmailNotifications returns a boolean if a field has been set.

### GetSoftwareTitleId

`func (o *PatchSoftwareTitleConfiguration) GetSoftwareTitleId() string`

GetSoftwareTitleId returns the SoftwareTitleId field if non-nil, zero value otherwise.

### GetSoftwareTitleIdOk

`func (o *PatchSoftwareTitleConfiguration) GetSoftwareTitleIdOk() (*string, bool)`

GetSoftwareTitleIdOk returns a tuple with the SoftwareTitleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTitleId

`func (o *PatchSoftwareTitleConfiguration) SetSoftwareTitleId(v string)`

SetSoftwareTitleId sets SoftwareTitleId field to given value.


### GetJamfOfficial

`func (o *PatchSoftwareTitleConfiguration) GetJamfOfficial() bool`

GetJamfOfficial returns the JamfOfficial field if non-nil, zero value otherwise.

### GetJamfOfficialOk

`func (o *PatchSoftwareTitleConfiguration) GetJamfOfficialOk() (*bool, bool)`

GetJamfOfficialOk returns a tuple with the JamfOfficial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJamfOfficial

`func (o *PatchSoftwareTitleConfiguration) SetJamfOfficial(v bool)`

SetJamfOfficial sets JamfOfficial field to given value.

### HasJamfOfficial

`func (o *PatchSoftwareTitleConfiguration) HasJamfOfficial() bool`

HasJamfOfficial returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *PatchSoftwareTitleConfiguration) GetExtensionAttributes() []PatchSoftwareTitleConfigurationExtensionAttributes`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *PatchSoftwareTitleConfiguration) GetExtensionAttributesOk() (*[]PatchSoftwareTitleConfigurationExtensionAttributes, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *PatchSoftwareTitleConfiguration) SetExtensionAttributes(v []PatchSoftwareTitleConfigurationExtensionAttributes)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *PatchSoftwareTitleConfiguration) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.

### GetSoftwareTitleName

`func (o *PatchSoftwareTitleConfiguration) GetSoftwareTitleName() string`

GetSoftwareTitleName returns the SoftwareTitleName field if non-nil, zero value otherwise.

### GetSoftwareTitleNameOk

`func (o *PatchSoftwareTitleConfiguration) GetSoftwareTitleNameOk() (*string, bool)`

GetSoftwareTitleNameOk returns a tuple with the SoftwareTitleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTitleName

`func (o *PatchSoftwareTitleConfiguration) SetSoftwareTitleName(v string)`

SetSoftwareTitleName sets SoftwareTitleName field to given value.

### HasSoftwareTitleName

`func (o *PatchSoftwareTitleConfiguration) HasSoftwareTitleName() bool`

HasSoftwareTitleName returns a boolean if a field has been set.

### GetSoftwareTitleNameId

`func (o *PatchSoftwareTitleConfiguration) GetSoftwareTitleNameId() string`

GetSoftwareTitleNameId returns the SoftwareTitleNameId field if non-nil, zero value otherwise.

### GetSoftwareTitleNameIdOk

`func (o *PatchSoftwareTitleConfiguration) GetSoftwareTitleNameIdOk() (*string, bool)`

GetSoftwareTitleNameIdOk returns a tuple with the SoftwareTitleNameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTitleNameId

`func (o *PatchSoftwareTitleConfiguration) SetSoftwareTitleNameId(v string)`

SetSoftwareTitleNameId sets SoftwareTitleNameId field to given value.

### HasSoftwareTitleNameId

`func (o *PatchSoftwareTitleConfiguration) HasSoftwareTitleNameId() bool`

HasSoftwareTitleNameId returns a boolean if a field has been set.

### GetSoftwareTitlePublisher

`func (o *PatchSoftwareTitleConfiguration) GetSoftwareTitlePublisher() string`

GetSoftwareTitlePublisher returns the SoftwareTitlePublisher field if non-nil, zero value otherwise.

### GetSoftwareTitlePublisherOk

`func (o *PatchSoftwareTitleConfiguration) GetSoftwareTitlePublisherOk() (*string, bool)`

GetSoftwareTitlePublisherOk returns a tuple with the SoftwareTitlePublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTitlePublisher

`func (o *PatchSoftwareTitleConfiguration) SetSoftwareTitlePublisher(v string)`

SetSoftwareTitlePublisher sets SoftwareTitlePublisher field to given value.

### HasSoftwareTitlePublisher

`func (o *PatchSoftwareTitleConfiguration) HasSoftwareTitlePublisher() bool`

HasSoftwareTitlePublisher returns a boolean if a field has been set.

### GetPatchSourceName

`func (o *PatchSoftwareTitleConfiguration) GetPatchSourceName() string`

GetPatchSourceName returns the PatchSourceName field if non-nil, zero value otherwise.

### GetPatchSourceNameOk

`func (o *PatchSoftwareTitleConfiguration) GetPatchSourceNameOk() (*string, bool)`

GetPatchSourceNameOk returns a tuple with the PatchSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchSourceName

`func (o *PatchSoftwareTitleConfiguration) SetPatchSourceName(v string)`

SetPatchSourceName sets PatchSourceName field to given value.

### HasPatchSourceName

`func (o *PatchSoftwareTitleConfiguration) HasPatchSourceName() bool`

HasPatchSourceName returns a boolean if a field has been set.

### GetPatchSourceEnabled

`func (o *PatchSoftwareTitleConfiguration) GetPatchSourceEnabled() bool`

GetPatchSourceEnabled returns the PatchSourceEnabled field if non-nil, zero value otherwise.

### GetPatchSourceEnabledOk

`func (o *PatchSoftwareTitleConfiguration) GetPatchSourceEnabledOk() (*bool, bool)`

GetPatchSourceEnabledOk returns a tuple with the PatchSourceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchSourceEnabled

`func (o *PatchSoftwareTitleConfiguration) SetPatchSourceEnabled(v bool)`

SetPatchSourceEnabled sets PatchSourceEnabled field to given value.

### HasPatchSourceEnabled

`func (o *PatchSoftwareTitleConfiguration) HasPatchSourceEnabled() bool`

HasPatchSourceEnabled returns a boolean if a field has been set.

### GetId

`func (o *PatchSoftwareTitleConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchSoftwareTitleConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchSoftwareTitleConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchSoftwareTitleConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPackages

`func (o *PatchSoftwareTitleConfiguration) GetPackages() []PatchSoftwareTitlePackages`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *PatchSoftwareTitleConfiguration) GetPackagesOk() (*[]PatchSoftwareTitlePackages, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *PatchSoftwareTitleConfiguration) SetPackages(v []PatchSoftwareTitlePackages)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *PatchSoftwareTitleConfiguration) HasPackages() bool`

HasPackages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


