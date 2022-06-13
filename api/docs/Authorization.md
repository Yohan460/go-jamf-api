# Authorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**AuthAccount**](AuthAccount.md) |  | [optional] 
**AccountGroups** | Pointer to [**[]AccountGroup**](AccountGroup.md) |  | [optional] 
**Sites** | Pointer to [**[]Site**](Site.md) |  | [optional] 
**AuthenticationType** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthorization

`func NewAuthorization() *Authorization`

NewAuthorization instantiates a new Authorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationWithDefaults

`func NewAuthorizationWithDefaults() *Authorization`

NewAuthorizationWithDefaults instantiates a new Authorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *Authorization) GetAccount() AuthAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Authorization) GetAccountOk() (*AuthAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Authorization) SetAccount(v AuthAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Authorization) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccountGroups

`func (o *Authorization) GetAccountGroups() []AccountGroup`

GetAccountGroups returns the AccountGroups field if non-nil, zero value otherwise.

### GetAccountGroupsOk

`func (o *Authorization) GetAccountGroupsOk() (*[]AccountGroup, bool)`

GetAccountGroupsOk returns a tuple with the AccountGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroups

`func (o *Authorization) SetAccountGroups(v []AccountGroup)`

SetAccountGroups sets AccountGroups field to given value.

### HasAccountGroups

`func (o *Authorization) HasAccountGroups() bool`

HasAccountGroups returns a boolean if a field has been set.

### GetSites

`func (o *Authorization) GetSites() []Site`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *Authorization) GetSitesOk() (*[]Site, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *Authorization) SetSites(v []Site)`

SetSites sets Sites field to given value.

### HasSites

`func (o *Authorization) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *Authorization) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *Authorization) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *Authorization) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *Authorization) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


