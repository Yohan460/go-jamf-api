# GetEnrollmentCustomizationPanelSsoAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**Rank** | **int32** |  | 
**IsUseJamfConnect** | **bool** |  | 
**LongNameAttribute** | **string** |  | 
**ShortNameAttribute** | **string** |  | 
**IsGroupEnrollmentAccessEnabled** | **bool** |  | 
**GroupEnrollmentAccessName** | **string** |  | [default to ""]
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewGetEnrollmentCustomizationPanelSsoAuth

`func NewGetEnrollmentCustomizationPanelSsoAuth(displayName string, rank int32, isUseJamfConnect bool, longNameAttribute string, shortNameAttribute string, isGroupEnrollmentAccessEnabled bool, groupEnrollmentAccessName string, ) *GetEnrollmentCustomizationPanelSsoAuth`

NewGetEnrollmentCustomizationPanelSsoAuth instantiates a new GetEnrollmentCustomizationPanelSsoAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEnrollmentCustomizationPanelSsoAuthWithDefaults

`func NewGetEnrollmentCustomizationPanelSsoAuthWithDefaults() *GetEnrollmentCustomizationPanelSsoAuth`

NewGetEnrollmentCustomizationPanelSsoAuthWithDefaults instantiates a new GetEnrollmentCustomizationPanelSsoAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *GetEnrollmentCustomizationPanelSsoAuth) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetEnrollmentCustomizationPanelSsoAuth) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetEnrollmentCustomizationPanelSsoAuth) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetRank

`func (o *GetEnrollmentCustomizationPanelSsoAuth) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *GetEnrollmentCustomizationPanelSsoAuth) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *GetEnrollmentCustomizationPanelSsoAuth) SetRank(v int32)`

SetRank sets Rank field to given value.


### GetIsUseJamfConnect

`func (o *GetEnrollmentCustomizationPanelSsoAuth) GetIsUseJamfConnect() bool`

GetIsUseJamfConnect returns the IsUseJamfConnect field if non-nil, zero value otherwise.

### GetIsUseJamfConnectOk

`func (o *GetEnrollmentCustomizationPanelSsoAuth) GetIsUseJamfConnectOk() (*bool, bool)`

GetIsUseJamfConnectOk returns a tuple with the IsUseJamfConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUseJamfConnect

`func (o *GetEnrollmentCustomizationPanelSsoAuth) SetIsUseJamfConnect(v bool)`

SetIsUseJamfConnect sets IsUseJamfConnect field to given value.


### GetLongNameAttribute

`func (o *GetEnrollmentCustomizationPanelSsoAuth) GetLongNameAttribute() string`

GetLongNameAttribute returns the LongNameAttribute field if non-nil, zero value otherwise.

### GetLongNameAttributeOk

`func (o *GetEnrollmentCustomizationPanelSsoAuth) GetLongNameAttributeOk() (*string, bool)`

GetLongNameAttributeOk returns a tuple with the LongNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongNameAttribute

`func (o *GetEnrollmentCustomizationPanelSsoAuth) SetLongNameAttribute(v string)`

SetLongNameAttribute sets LongNameAttribute field to given value.


### GetShortNameAttribute

`func (o *GetEnrollmentCustomizationPanelSsoAuth) GetShortNameAttribute() string`

GetShortNameAttribute returns the ShortNameAttribute field if non-nil, zero value otherwise.

### GetShortNameAttributeOk

`func (o *GetEnrollmentCustomizationPanelSsoAuth) GetShortNameAttributeOk() (*string, bool)`

GetShortNameAttributeOk returns a tuple with the ShortNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortNameAttribute

`func (o *GetEnrollmentCustomizationPanelSsoAuth) SetShortNameAttribute(v string)`

SetShortNameAttribute sets ShortNameAttribute field to given value.


### GetIsGroupEnrollmentAccessEnabled

`func (o *GetEnrollmentCustomizationPanelSsoAuth) GetIsGroupEnrollmentAccessEnabled() bool`

GetIsGroupEnrollmentAccessEnabled returns the IsGroupEnrollmentAccessEnabled field if non-nil, zero value otherwise.

### GetIsGroupEnrollmentAccessEnabledOk

`func (o *GetEnrollmentCustomizationPanelSsoAuth) GetIsGroupEnrollmentAccessEnabledOk() (*bool, bool)`

GetIsGroupEnrollmentAccessEnabledOk returns a tuple with the IsGroupEnrollmentAccessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGroupEnrollmentAccessEnabled

`func (o *GetEnrollmentCustomizationPanelSsoAuth) SetIsGroupEnrollmentAccessEnabled(v bool)`

SetIsGroupEnrollmentAccessEnabled sets IsGroupEnrollmentAccessEnabled field to given value.


### GetGroupEnrollmentAccessName

`func (o *GetEnrollmentCustomizationPanelSsoAuth) GetGroupEnrollmentAccessName() string`

GetGroupEnrollmentAccessName returns the GroupEnrollmentAccessName field if non-nil, zero value otherwise.

### GetGroupEnrollmentAccessNameOk

`func (o *GetEnrollmentCustomizationPanelSsoAuth) GetGroupEnrollmentAccessNameOk() (*string, bool)`

GetGroupEnrollmentAccessNameOk returns a tuple with the GroupEnrollmentAccessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupEnrollmentAccessName

`func (o *GetEnrollmentCustomizationPanelSsoAuth) SetGroupEnrollmentAccessName(v string)`

SetGroupEnrollmentAccessName sets GroupEnrollmentAccessName field to given value.


### GetId

`func (o *GetEnrollmentCustomizationPanelSsoAuth) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEnrollmentCustomizationPanelSsoAuth) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEnrollmentCustomizationPanelSsoAuth) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetEnrollmentCustomizationPanelSsoAuth) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetEnrollmentCustomizationPanelSsoAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetEnrollmentCustomizationPanelSsoAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetEnrollmentCustomizationPanelSsoAuth) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetEnrollmentCustomizationPanelSsoAuth) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


