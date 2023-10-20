# PatchSoftwareTitleConfigurationBase

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

## Methods

### NewPatchSoftwareTitleConfigurationBase

`func NewPatchSoftwareTitleConfigurationBase(displayName string, softwareTitleId string, ) *PatchSoftwareTitleConfigurationBase`

NewPatchSoftwareTitleConfigurationBase instantiates a new PatchSoftwareTitleConfigurationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchSoftwareTitleConfigurationBaseWithDefaults

`func NewPatchSoftwareTitleConfigurationBaseWithDefaults() *PatchSoftwareTitleConfigurationBase`

NewPatchSoftwareTitleConfigurationBaseWithDefaults instantiates a new PatchSoftwareTitleConfigurationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *PatchSoftwareTitleConfigurationBase) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PatchSoftwareTitleConfigurationBase) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PatchSoftwareTitleConfigurationBase) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetCategoryId

`func (o *PatchSoftwareTitleConfigurationBase) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *PatchSoftwareTitleConfigurationBase) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *PatchSoftwareTitleConfigurationBase) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *PatchSoftwareTitleConfigurationBase) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetSiteId

`func (o *PatchSoftwareTitleConfigurationBase) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *PatchSoftwareTitleConfigurationBase) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *PatchSoftwareTitleConfigurationBase) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *PatchSoftwareTitleConfigurationBase) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetUiNotifications

`func (o *PatchSoftwareTitleConfigurationBase) GetUiNotifications() bool`

GetUiNotifications returns the UiNotifications field if non-nil, zero value otherwise.

### GetUiNotificationsOk

`func (o *PatchSoftwareTitleConfigurationBase) GetUiNotificationsOk() (*bool, bool)`

GetUiNotificationsOk returns a tuple with the UiNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiNotifications

`func (o *PatchSoftwareTitleConfigurationBase) SetUiNotifications(v bool)`

SetUiNotifications sets UiNotifications field to given value.

### HasUiNotifications

`func (o *PatchSoftwareTitleConfigurationBase) HasUiNotifications() bool`

HasUiNotifications returns a boolean if a field has been set.

### GetEmailNotifications

`func (o *PatchSoftwareTitleConfigurationBase) GetEmailNotifications() bool`

GetEmailNotifications returns the EmailNotifications field if non-nil, zero value otherwise.

### GetEmailNotificationsOk

`func (o *PatchSoftwareTitleConfigurationBase) GetEmailNotificationsOk() (*bool, bool)`

GetEmailNotificationsOk returns a tuple with the EmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotifications

`func (o *PatchSoftwareTitleConfigurationBase) SetEmailNotifications(v bool)`

SetEmailNotifications sets EmailNotifications field to given value.

### HasEmailNotifications

`func (o *PatchSoftwareTitleConfigurationBase) HasEmailNotifications() bool`

HasEmailNotifications returns a boolean if a field has been set.

### GetSoftwareTitleId

`func (o *PatchSoftwareTitleConfigurationBase) GetSoftwareTitleId() string`

GetSoftwareTitleId returns the SoftwareTitleId field if non-nil, zero value otherwise.

### GetSoftwareTitleIdOk

`func (o *PatchSoftwareTitleConfigurationBase) GetSoftwareTitleIdOk() (*string, bool)`

GetSoftwareTitleIdOk returns a tuple with the SoftwareTitleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTitleId

`func (o *PatchSoftwareTitleConfigurationBase) SetSoftwareTitleId(v string)`

SetSoftwareTitleId sets SoftwareTitleId field to given value.


### GetJamfOfficial

`func (o *PatchSoftwareTitleConfigurationBase) GetJamfOfficial() bool`

GetJamfOfficial returns the JamfOfficial field if non-nil, zero value otherwise.

### GetJamfOfficialOk

`func (o *PatchSoftwareTitleConfigurationBase) GetJamfOfficialOk() (*bool, bool)`

GetJamfOfficialOk returns a tuple with the JamfOfficial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJamfOfficial

`func (o *PatchSoftwareTitleConfigurationBase) SetJamfOfficial(v bool)`

SetJamfOfficial sets JamfOfficial field to given value.

### HasJamfOfficial

`func (o *PatchSoftwareTitleConfigurationBase) HasJamfOfficial() bool`

HasJamfOfficial returns a boolean if a field has been set.

### GetExtensionAttributes

`func (o *PatchSoftwareTitleConfigurationBase) GetExtensionAttributes() []PatchSoftwareTitleConfigurationExtensionAttributes`

GetExtensionAttributes returns the ExtensionAttributes field if non-nil, zero value otherwise.

### GetExtensionAttributesOk

`func (o *PatchSoftwareTitleConfigurationBase) GetExtensionAttributesOk() (*[]PatchSoftwareTitleConfigurationExtensionAttributes, bool)`

GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionAttributes

`func (o *PatchSoftwareTitleConfigurationBase) SetExtensionAttributes(v []PatchSoftwareTitleConfigurationExtensionAttributes)`

SetExtensionAttributes sets ExtensionAttributes field to given value.

### HasExtensionAttributes

`func (o *PatchSoftwareTitleConfigurationBase) HasExtensionAttributes() bool`

HasExtensionAttributes returns a boolean if a field has been set.

### GetSoftwareTitleName

`func (o *PatchSoftwareTitleConfigurationBase) GetSoftwareTitleName() string`

GetSoftwareTitleName returns the SoftwareTitleName field if non-nil, zero value otherwise.

### GetSoftwareTitleNameOk

`func (o *PatchSoftwareTitleConfigurationBase) GetSoftwareTitleNameOk() (*string, bool)`

GetSoftwareTitleNameOk returns a tuple with the SoftwareTitleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTitleName

`func (o *PatchSoftwareTitleConfigurationBase) SetSoftwareTitleName(v string)`

SetSoftwareTitleName sets SoftwareTitleName field to given value.

### HasSoftwareTitleName

`func (o *PatchSoftwareTitleConfigurationBase) HasSoftwareTitleName() bool`

HasSoftwareTitleName returns a boolean if a field has been set.

### GetSoftwareTitleNameId

`func (o *PatchSoftwareTitleConfigurationBase) GetSoftwareTitleNameId() string`

GetSoftwareTitleNameId returns the SoftwareTitleNameId field if non-nil, zero value otherwise.

### GetSoftwareTitleNameIdOk

`func (o *PatchSoftwareTitleConfigurationBase) GetSoftwareTitleNameIdOk() (*string, bool)`

GetSoftwareTitleNameIdOk returns a tuple with the SoftwareTitleNameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTitleNameId

`func (o *PatchSoftwareTitleConfigurationBase) SetSoftwareTitleNameId(v string)`

SetSoftwareTitleNameId sets SoftwareTitleNameId field to given value.

### HasSoftwareTitleNameId

`func (o *PatchSoftwareTitleConfigurationBase) HasSoftwareTitleNameId() bool`

HasSoftwareTitleNameId returns a boolean if a field has been set.

### GetSoftwareTitlePublisher

`func (o *PatchSoftwareTitleConfigurationBase) GetSoftwareTitlePublisher() string`

GetSoftwareTitlePublisher returns the SoftwareTitlePublisher field if non-nil, zero value otherwise.

### GetSoftwareTitlePublisherOk

`func (o *PatchSoftwareTitleConfigurationBase) GetSoftwareTitlePublisherOk() (*string, bool)`

GetSoftwareTitlePublisherOk returns a tuple with the SoftwareTitlePublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTitlePublisher

`func (o *PatchSoftwareTitleConfigurationBase) SetSoftwareTitlePublisher(v string)`

SetSoftwareTitlePublisher sets SoftwareTitlePublisher field to given value.

### HasSoftwareTitlePublisher

`func (o *PatchSoftwareTitleConfigurationBase) HasSoftwareTitlePublisher() bool`

HasSoftwareTitlePublisher returns a boolean if a field has been set.

### GetPatchSourceName

`func (o *PatchSoftwareTitleConfigurationBase) GetPatchSourceName() string`

GetPatchSourceName returns the PatchSourceName field if non-nil, zero value otherwise.

### GetPatchSourceNameOk

`func (o *PatchSoftwareTitleConfigurationBase) GetPatchSourceNameOk() (*string, bool)`

GetPatchSourceNameOk returns a tuple with the PatchSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchSourceName

`func (o *PatchSoftwareTitleConfigurationBase) SetPatchSourceName(v string)`

SetPatchSourceName sets PatchSourceName field to given value.

### HasPatchSourceName

`func (o *PatchSoftwareTitleConfigurationBase) HasPatchSourceName() bool`

HasPatchSourceName returns a boolean if a field has been set.

### GetPatchSourceEnabled

`func (o *PatchSoftwareTitleConfigurationBase) GetPatchSourceEnabled() bool`

GetPatchSourceEnabled returns the PatchSourceEnabled field if non-nil, zero value otherwise.

### GetPatchSourceEnabledOk

`func (o *PatchSoftwareTitleConfigurationBase) GetPatchSourceEnabledOk() (*bool, bool)`

GetPatchSourceEnabledOk returns a tuple with the PatchSourceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchSourceEnabled

`func (o *PatchSoftwareTitleConfigurationBase) SetPatchSourceEnabled(v bool)`

SetPatchSourceEnabled sets PatchSourceEnabled field to given value.

### HasPatchSourceEnabled

`func (o *PatchSoftwareTitleConfigurationBase) HasPatchSourceEnabled() bool`

HasPatchSourceEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


