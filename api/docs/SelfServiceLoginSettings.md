# SelfServiceLoginSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserLoginLevel** | **string** | login setting to tell clients how to let users log in  | 
**AllowRememberMe** | Pointer to **bool** | true if remember me functionality is allowed, false if not  | [optional] [default to false]
**AuthType** | **string** | login type to be used when asking users to log in  | 

## Methods

### NewSelfServiceLoginSettings

`func NewSelfServiceLoginSettings(userLoginLevel string, authType string, ) *SelfServiceLoginSettings`

NewSelfServiceLoginSettings instantiates a new SelfServiceLoginSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfServiceLoginSettingsWithDefaults

`func NewSelfServiceLoginSettingsWithDefaults() *SelfServiceLoginSettings`

NewSelfServiceLoginSettingsWithDefaults instantiates a new SelfServiceLoginSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserLoginLevel

`func (o *SelfServiceLoginSettings) GetUserLoginLevel() string`

GetUserLoginLevel returns the UserLoginLevel field if non-nil, zero value otherwise.

### GetUserLoginLevelOk

`func (o *SelfServiceLoginSettings) GetUserLoginLevelOk() (*string, bool)`

GetUserLoginLevelOk returns a tuple with the UserLoginLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLoginLevel

`func (o *SelfServiceLoginSettings) SetUserLoginLevel(v string)`

SetUserLoginLevel sets UserLoginLevel field to given value.


### GetAllowRememberMe

`func (o *SelfServiceLoginSettings) GetAllowRememberMe() bool`

GetAllowRememberMe returns the AllowRememberMe field if non-nil, zero value otherwise.

### GetAllowRememberMeOk

`func (o *SelfServiceLoginSettings) GetAllowRememberMeOk() (*bool, bool)`

GetAllowRememberMeOk returns a tuple with the AllowRememberMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRememberMe

`func (o *SelfServiceLoginSettings) SetAllowRememberMe(v bool)`

SetAllowRememberMe sets AllowRememberMe field to given value.

### HasAllowRememberMe

`func (o *SelfServiceLoginSettings) HasAllowRememberMe() bool`

HasAllowRememberMe returns a boolean if a field has been set.

### GetAuthType

`func (o *SelfServiceLoginSettings) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *SelfServiceLoginSettings) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *SelfServiceLoginSettings) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


