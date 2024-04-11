# VppTokenSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Recipients** | Pointer to [**Recipients**](Recipients.md) |  | [optional] 
**AdminAccounts** | Pointer to [**[]AdminAccount**](AdminAccount.md) |  | [optional] 
**SiteID** | Pointer to **int64** |  | [optional] 
**SiteName** | Pointer to **string** |  | [optional] 

## Methods

### NewVppTokenSubscription

`func NewVppTokenSubscription() *VppTokenSubscription`

NewVppTokenSubscription instantiates a new VppTokenSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVppTokenSubscriptionWithDefaults

`func NewVppTokenSubscriptionWithDefaults() *VppTokenSubscription`

NewVppTokenSubscriptionWithDefaults instantiates a new VppTokenSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VppTokenSubscription) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VppTokenSubscription) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VppTokenSubscription) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VppTokenSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VppTokenSubscription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VppTokenSubscription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VppTokenSubscription) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VppTokenSubscription) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *VppTokenSubscription) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VppTokenSubscription) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VppTokenSubscription) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VppTokenSubscription) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRecipients

`func (o *VppTokenSubscription) GetRecipients() Recipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *VppTokenSubscription) GetRecipientsOk() (*Recipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *VppTokenSubscription) SetRecipients(v Recipients)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *VppTokenSubscription) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetAdminAccounts

`func (o *VppTokenSubscription) GetAdminAccounts() []AdminAccount`

GetAdminAccounts returns the AdminAccounts field if non-nil, zero value otherwise.

### GetAdminAccountsOk

`func (o *VppTokenSubscription) GetAdminAccountsOk() (*[]AdminAccount, bool)`

GetAdminAccountsOk returns a tuple with the AdminAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAccounts

`func (o *VppTokenSubscription) SetAdminAccounts(v []AdminAccount)`

SetAdminAccounts sets AdminAccounts field to given value.

### HasAdminAccounts

`func (o *VppTokenSubscription) HasAdminAccounts() bool`

HasAdminAccounts returns a boolean if a field has been set.

### GetSiteID

`func (o *VppTokenSubscription) GetSiteID() int64`

GetSiteID returns the SiteID field if non-nil, zero value otherwise.

### GetSiteIDOk

`func (o *VppTokenSubscription) GetSiteIDOk() (*int64, bool)`

GetSiteIDOk returns a tuple with the SiteID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteID

`func (o *VppTokenSubscription) SetSiteID(v int64)`

SetSiteID sets SiteID field to given value.

### HasSiteID

`func (o *VppTokenSubscription) HasSiteID() bool`

HasSiteID returns a boolean if a field has been set.

### GetSiteName

`func (o *VppTokenSubscription) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *VppTokenSubscription) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *VppTokenSubscription) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *VppTokenSubscription) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


