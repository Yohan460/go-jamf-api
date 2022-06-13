# VolumePurchasingLocationPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | If no value is provided when creating a VolumePurchasingLocation object, the &#39;name&#39; will default to the &#39;locationName&#39; value | [optional] 
**AutomaticallyPopulatePurchasedContent** | Pointer to **bool** |  | [optional] [default to false]
**SendNotificationWhenNoLongerAssigned** | Pointer to **bool** |  | [optional] [default to false]
**AutoRegisterManagedUsers** | Pointer to **bool** |  | [optional] [default to false]
**SiteId** | Pointer to **string** |  | [optional] [default to "-1"]
**ServiceToken** | **string** |  | 

## Methods

### NewVolumePurchasingLocationPost

`func NewVolumePurchasingLocationPost(serviceToken string, ) *VolumePurchasingLocationPost`

NewVolumePurchasingLocationPost instantiates a new VolumePurchasingLocationPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumePurchasingLocationPostWithDefaults

`func NewVolumePurchasingLocationPostWithDefaults() *VolumePurchasingLocationPost`

NewVolumePurchasingLocationPostWithDefaults instantiates a new VolumePurchasingLocationPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VolumePurchasingLocationPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumePurchasingLocationPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumePurchasingLocationPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumePurchasingLocationPost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAutomaticallyPopulatePurchasedContent

`func (o *VolumePurchasingLocationPost) GetAutomaticallyPopulatePurchasedContent() bool`

GetAutomaticallyPopulatePurchasedContent returns the AutomaticallyPopulatePurchasedContent field if non-nil, zero value otherwise.

### GetAutomaticallyPopulatePurchasedContentOk

`func (o *VolumePurchasingLocationPost) GetAutomaticallyPopulatePurchasedContentOk() (*bool, bool)`

GetAutomaticallyPopulatePurchasedContentOk returns a tuple with the AutomaticallyPopulatePurchasedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyPopulatePurchasedContent

`func (o *VolumePurchasingLocationPost) SetAutomaticallyPopulatePurchasedContent(v bool)`

SetAutomaticallyPopulatePurchasedContent sets AutomaticallyPopulatePurchasedContent field to given value.

### HasAutomaticallyPopulatePurchasedContent

`func (o *VolumePurchasingLocationPost) HasAutomaticallyPopulatePurchasedContent() bool`

HasAutomaticallyPopulatePurchasedContent returns a boolean if a field has been set.

### GetSendNotificationWhenNoLongerAssigned

`func (o *VolumePurchasingLocationPost) GetSendNotificationWhenNoLongerAssigned() bool`

GetSendNotificationWhenNoLongerAssigned returns the SendNotificationWhenNoLongerAssigned field if non-nil, zero value otherwise.

### GetSendNotificationWhenNoLongerAssignedOk

`func (o *VolumePurchasingLocationPost) GetSendNotificationWhenNoLongerAssignedOk() (*bool, bool)`

GetSendNotificationWhenNoLongerAssignedOk returns a tuple with the SendNotificationWhenNoLongerAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNotificationWhenNoLongerAssigned

`func (o *VolumePurchasingLocationPost) SetSendNotificationWhenNoLongerAssigned(v bool)`

SetSendNotificationWhenNoLongerAssigned sets SendNotificationWhenNoLongerAssigned field to given value.

### HasSendNotificationWhenNoLongerAssigned

`func (o *VolumePurchasingLocationPost) HasSendNotificationWhenNoLongerAssigned() bool`

HasSendNotificationWhenNoLongerAssigned returns a boolean if a field has been set.

### GetAutoRegisterManagedUsers

`func (o *VolumePurchasingLocationPost) GetAutoRegisterManagedUsers() bool`

GetAutoRegisterManagedUsers returns the AutoRegisterManagedUsers field if non-nil, zero value otherwise.

### GetAutoRegisterManagedUsersOk

`func (o *VolumePurchasingLocationPost) GetAutoRegisterManagedUsersOk() (*bool, bool)`

GetAutoRegisterManagedUsersOk returns a tuple with the AutoRegisterManagedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRegisterManagedUsers

`func (o *VolumePurchasingLocationPost) SetAutoRegisterManagedUsers(v bool)`

SetAutoRegisterManagedUsers sets AutoRegisterManagedUsers field to given value.

### HasAutoRegisterManagedUsers

`func (o *VolumePurchasingLocationPost) HasAutoRegisterManagedUsers() bool`

HasAutoRegisterManagedUsers returns a boolean if a field has been set.

### GetSiteId

`func (o *VolumePurchasingLocationPost) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *VolumePurchasingLocationPost) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *VolumePurchasingLocationPost) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *VolumePurchasingLocationPost) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetServiceToken

`func (o *VolumePurchasingLocationPost) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *VolumePurchasingLocationPost) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *VolumePurchasingLocationPost) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


