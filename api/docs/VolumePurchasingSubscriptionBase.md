# VolumePurchasingSubscriptionBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**Triggers** | Pointer to **[]string** |  | [optional] 
**LocationIds** | Pointer to **[]string** |  | [optional] 
**InternalRecipients** | Pointer to [**[]InternalRecipient**](InternalRecipient.md) |  | [optional] 
**ExternalRecipients** | Pointer to [**[]ExternalRecipient**](ExternalRecipient.md) |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [default to "-1"]

## Methods

### NewVolumePurchasingSubscriptionBase

`func NewVolumePurchasingSubscriptionBase(name string, ) *VolumePurchasingSubscriptionBase`

NewVolumePurchasingSubscriptionBase instantiates a new VolumePurchasingSubscriptionBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumePurchasingSubscriptionBaseWithDefaults

`func NewVolumePurchasingSubscriptionBaseWithDefaults() *VolumePurchasingSubscriptionBase`

NewVolumePurchasingSubscriptionBaseWithDefaults instantiates a new VolumePurchasingSubscriptionBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VolumePurchasingSubscriptionBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumePurchasingSubscriptionBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumePurchasingSubscriptionBase) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *VolumePurchasingSubscriptionBase) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VolumePurchasingSubscriptionBase) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VolumePurchasingSubscriptionBase) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VolumePurchasingSubscriptionBase) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTriggers

`func (o *VolumePurchasingSubscriptionBase) GetTriggers() []string`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *VolumePurchasingSubscriptionBase) GetTriggersOk() (*[]string, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *VolumePurchasingSubscriptionBase) SetTriggers(v []string)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *VolumePurchasingSubscriptionBase) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetLocationIds

`func (o *VolumePurchasingSubscriptionBase) GetLocationIds() []string`

GetLocationIds returns the LocationIds field if non-nil, zero value otherwise.

### GetLocationIdsOk

`func (o *VolumePurchasingSubscriptionBase) GetLocationIdsOk() (*[]string, bool)`

GetLocationIdsOk returns a tuple with the LocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationIds

`func (o *VolumePurchasingSubscriptionBase) SetLocationIds(v []string)`

SetLocationIds sets LocationIds field to given value.

### HasLocationIds

`func (o *VolumePurchasingSubscriptionBase) HasLocationIds() bool`

HasLocationIds returns a boolean if a field has been set.

### GetInternalRecipients

`func (o *VolumePurchasingSubscriptionBase) GetInternalRecipients() []InternalRecipient`

GetInternalRecipients returns the InternalRecipients field if non-nil, zero value otherwise.

### GetInternalRecipientsOk

`func (o *VolumePurchasingSubscriptionBase) GetInternalRecipientsOk() (*[]InternalRecipient, bool)`

GetInternalRecipientsOk returns a tuple with the InternalRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalRecipients

`func (o *VolumePurchasingSubscriptionBase) SetInternalRecipients(v []InternalRecipient)`

SetInternalRecipients sets InternalRecipients field to given value.

### HasInternalRecipients

`func (o *VolumePurchasingSubscriptionBase) HasInternalRecipients() bool`

HasInternalRecipients returns a boolean if a field has been set.

### GetExternalRecipients

`func (o *VolumePurchasingSubscriptionBase) GetExternalRecipients() []ExternalRecipient`

GetExternalRecipients returns the ExternalRecipients field if non-nil, zero value otherwise.

### GetExternalRecipientsOk

`func (o *VolumePurchasingSubscriptionBase) GetExternalRecipientsOk() (*[]ExternalRecipient, bool)`

GetExternalRecipientsOk returns a tuple with the ExternalRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRecipients

`func (o *VolumePurchasingSubscriptionBase) SetExternalRecipients(v []ExternalRecipient)`

SetExternalRecipients sets ExternalRecipients field to given value.

### HasExternalRecipients

`func (o *VolumePurchasingSubscriptionBase) HasExternalRecipients() bool`

HasExternalRecipients returns a boolean if a field has been set.

### GetSiteId

`func (o *VolumePurchasingSubscriptionBase) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *VolumePurchasingSubscriptionBase) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *VolumePurchasingSubscriptionBase) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *VolumePurchasingSubscriptionBase) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


