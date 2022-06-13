# GetEnrollmentCustomizationPanelLdapAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**Rank** | **int32** |  | 
**UsernameLabel** | **string** |  | 
**PasswordLabel** | **string** |  | 
**Title** | **string** |  | 
**BackButtonText** | **string** |  | 
**ContinueButtonText** | **string** |  | 
**LdapGroupAccess** | Pointer to [**[]EnrollmentCustomizationLdapGroupAccess**](EnrollmentCustomizationLdapGroupAccess.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewGetEnrollmentCustomizationPanelLdapAuth

`func NewGetEnrollmentCustomizationPanelLdapAuth(displayName string, rank int32, usernameLabel string, passwordLabel string, title string, backButtonText string, continueButtonText string, ) *GetEnrollmentCustomizationPanelLdapAuth`

NewGetEnrollmentCustomizationPanelLdapAuth instantiates a new GetEnrollmentCustomizationPanelLdapAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEnrollmentCustomizationPanelLdapAuthWithDefaults

`func NewGetEnrollmentCustomizationPanelLdapAuthWithDefaults() *GetEnrollmentCustomizationPanelLdapAuth`

NewGetEnrollmentCustomizationPanelLdapAuthWithDefaults instantiates a new GetEnrollmentCustomizationPanelLdapAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *GetEnrollmentCustomizationPanelLdapAuth) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetEnrollmentCustomizationPanelLdapAuth) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetEnrollmentCustomizationPanelLdapAuth) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetRank

`func (o *GetEnrollmentCustomizationPanelLdapAuth) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *GetEnrollmentCustomizationPanelLdapAuth) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *GetEnrollmentCustomizationPanelLdapAuth) SetRank(v int32)`

SetRank sets Rank field to given value.


### GetUsernameLabel

`func (o *GetEnrollmentCustomizationPanelLdapAuth) GetUsernameLabel() string`

GetUsernameLabel returns the UsernameLabel field if non-nil, zero value otherwise.

### GetUsernameLabelOk

`func (o *GetEnrollmentCustomizationPanelLdapAuth) GetUsernameLabelOk() (*string, bool)`

GetUsernameLabelOk returns a tuple with the UsernameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameLabel

`func (o *GetEnrollmentCustomizationPanelLdapAuth) SetUsernameLabel(v string)`

SetUsernameLabel sets UsernameLabel field to given value.


### GetPasswordLabel

`func (o *GetEnrollmentCustomizationPanelLdapAuth) GetPasswordLabel() string`

GetPasswordLabel returns the PasswordLabel field if non-nil, zero value otherwise.

### GetPasswordLabelOk

`func (o *GetEnrollmentCustomizationPanelLdapAuth) GetPasswordLabelOk() (*string, bool)`

GetPasswordLabelOk returns a tuple with the PasswordLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLabel

`func (o *GetEnrollmentCustomizationPanelLdapAuth) SetPasswordLabel(v string)`

SetPasswordLabel sets PasswordLabel field to given value.


### GetTitle

`func (o *GetEnrollmentCustomizationPanelLdapAuth) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetEnrollmentCustomizationPanelLdapAuth) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetEnrollmentCustomizationPanelLdapAuth) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBackButtonText

`func (o *GetEnrollmentCustomizationPanelLdapAuth) GetBackButtonText() string`

GetBackButtonText returns the BackButtonText field if non-nil, zero value otherwise.

### GetBackButtonTextOk

`func (o *GetEnrollmentCustomizationPanelLdapAuth) GetBackButtonTextOk() (*string, bool)`

GetBackButtonTextOk returns a tuple with the BackButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackButtonText

`func (o *GetEnrollmentCustomizationPanelLdapAuth) SetBackButtonText(v string)`

SetBackButtonText sets BackButtonText field to given value.


### GetContinueButtonText

`func (o *GetEnrollmentCustomizationPanelLdapAuth) GetContinueButtonText() string`

GetContinueButtonText returns the ContinueButtonText field if non-nil, zero value otherwise.

### GetContinueButtonTextOk

`func (o *GetEnrollmentCustomizationPanelLdapAuth) GetContinueButtonTextOk() (*string, bool)`

GetContinueButtonTextOk returns a tuple with the ContinueButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueButtonText

`func (o *GetEnrollmentCustomizationPanelLdapAuth) SetContinueButtonText(v string)`

SetContinueButtonText sets ContinueButtonText field to given value.


### GetLdapGroupAccess

`func (o *GetEnrollmentCustomizationPanelLdapAuth) GetLdapGroupAccess() []EnrollmentCustomizationLdapGroupAccess`

GetLdapGroupAccess returns the LdapGroupAccess field if non-nil, zero value otherwise.

### GetLdapGroupAccessOk

`func (o *GetEnrollmentCustomizationPanelLdapAuth) GetLdapGroupAccessOk() (*[]EnrollmentCustomizationLdapGroupAccess, bool)`

GetLdapGroupAccessOk returns a tuple with the LdapGroupAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupAccess

`func (o *GetEnrollmentCustomizationPanelLdapAuth) SetLdapGroupAccess(v []EnrollmentCustomizationLdapGroupAccess)`

SetLdapGroupAccess sets LdapGroupAccess field to given value.

### HasLdapGroupAccess

`func (o *GetEnrollmentCustomizationPanelLdapAuth) HasLdapGroupAccess() bool`

HasLdapGroupAccess returns a boolean if a field has been set.

### GetId

`func (o *GetEnrollmentCustomizationPanelLdapAuth) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEnrollmentCustomizationPanelLdapAuth) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEnrollmentCustomizationPanelLdapAuth) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetEnrollmentCustomizationPanelLdapAuth) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetEnrollmentCustomizationPanelLdapAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetEnrollmentCustomizationPanelLdapAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetEnrollmentCustomizationPanelLdapAuth) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetEnrollmentCustomizationPanelLdapAuth) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


