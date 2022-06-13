# EnrollmentCustomizationPanelLdapAuth

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

## Methods

### NewEnrollmentCustomizationPanelLdapAuth

`func NewEnrollmentCustomizationPanelLdapAuth(displayName string, rank int32, usernameLabel string, passwordLabel string, title string, backButtonText string, continueButtonText string, ) *EnrollmentCustomizationPanelLdapAuth`

NewEnrollmentCustomizationPanelLdapAuth instantiates a new EnrollmentCustomizationPanelLdapAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentCustomizationPanelLdapAuthWithDefaults

`func NewEnrollmentCustomizationPanelLdapAuthWithDefaults() *EnrollmentCustomizationPanelLdapAuth`

NewEnrollmentCustomizationPanelLdapAuthWithDefaults instantiates a new EnrollmentCustomizationPanelLdapAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *EnrollmentCustomizationPanelLdapAuth) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnrollmentCustomizationPanelLdapAuth) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnrollmentCustomizationPanelLdapAuth) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetRank

`func (o *EnrollmentCustomizationPanelLdapAuth) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *EnrollmentCustomizationPanelLdapAuth) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *EnrollmentCustomizationPanelLdapAuth) SetRank(v int32)`

SetRank sets Rank field to given value.


### GetUsernameLabel

`func (o *EnrollmentCustomizationPanelLdapAuth) GetUsernameLabel() string`

GetUsernameLabel returns the UsernameLabel field if non-nil, zero value otherwise.

### GetUsernameLabelOk

`func (o *EnrollmentCustomizationPanelLdapAuth) GetUsernameLabelOk() (*string, bool)`

GetUsernameLabelOk returns a tuple with the UsernameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameLabel

`func (o *EnrollmentCustomizationPanelLdapAuth) SetUsernameLabel(v string)`

SetUsernameLabel sets UsernameLabel field to given value.


### GetPasswordLabel

`func (o *EnrollmentCustomizationPanelLdapAuth) GetPasswordLabel() string`

GetPasswordLabel returns the PasswordLabel field if non-nil, zero value otherwise.

### GetPasswordLabelOk

`func (o *EnrollmentCustomizationPanelLdapAuth) GetPasswordLabelOk() (*string, bool)`

GetPasswordLabelOk returns a tuple with the PasswordLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLabel

`func (o *EnrollmentCustomizationPanelLdapAuth) SetPasswordLabel(v string)`

SetPasswordLabel sets PasswordLabel field to given value.


### GetTitle

`func (o *EnrollmentCustomizationPanelLdapAuth) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EnrollmentCustomizationPanelLdapAuth) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EnrollmentCustomizationPanelLdapAuth) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBackButtonText

`func (o *EnrollmentCustomizationPanelLdapAuth) GetBackButtonText() string`

GetBackButtonText returns the BackButtonText field if non-nil, zero value otherwise.

### GetBackButtonTextOk

`func (o *EnrollmentCustomizationPanelLdapAuth) GetBackButtonTextOk() (*string, bool)`

GetBackButtonTextOk returns a tuple with the BackButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackButtonText

`func (o *EnrollmentCustomizationPanelLdapAuth) SetBackButtonText(v string)`

SetBackButtonText sets BackButtonText field to given value.


### GetContinueButtonText

`func (o *EnrollmentCustomizationPanelLdapAuth) GetContinueButtonText() string`

GetContinueButtonText returns the ContinueButtonText field if non-nil, zero value otherwise.

### GetContinueButtonTextOk

`func (o *EnrollmentCustomizationPanelLdapAuth) GetContinueButtonTextOk() (*string, bool)`

GetContinueButtonTextOk returns a tuple with the ContinueButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueButtonText

`func (o *EnrollmentCustomizationPanelLdapAuth) SetContinueButtonText(v string)`

SetContinueButtonText sets ContinueButtonText field to given value.


### GetLdapGroupAccess

`func (o *EnrollmentCustomizationPanelLdapAuth) GetLdapGroupAccess() []EnrollmentCustomizationLdapGroupAccess`

GetLdapGroupAccess returns the LdapGroupAccess field if non-nil, zero value otherwise.

### GetLdapGroupAccessOk

`func (o *EnrollmentCustomizationPanelLdapAuth) GetLdapGroupAccessOk() (*[]EnrollmentCustomizationLdapGroupAccess, bool)`

GetLdapGroupAccessOk returns a tuple with the LdapGroupAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupAccess

`func (o *EnrollmentCustomizationPanelLdapAuth) SetLdapGroupAccess(v []EnrollmentCustomizationLdapGroupAccess)`

SetLdapGroupAccess sets LdapGroupAccess field to given value.

### HasLdapGroupAccess

`func (o *EnrollmentCustomizationPanelLdapAuth) HasLdapGroupAccess() bool`

HasLdapGroupAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


