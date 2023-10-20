# AuthorizationV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**AuthAccountV1**](AuthAccountV1.md) |  | [optional] 
**AccountGroups** | Pointer to [**[]AccountGroup**](AccountGroup.md) |  | [optional] 
**Sites** | Pointer to [**[]V1Site**](V1Site.md) |  | [optional] 
**AuthenticationType** | Pointer to [**AuthenticationType**](AuthenticationType.md) |  | [optional] 

## Methods

### NewAuthorizationV1

`func NewAuthorizationV1() *AuthorizationV1`

NewAuthorizationV1 instantiates a new AuthorizationV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationV1WithDefaults

`func NewAuthorizationV1WithDefaults() *AuthorizationV1`

NewAuthorizationV1WithDefaults instantiates a new AuthorizationV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AuthorizationV1) GetAccount() AuthAccountV1`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AuthorizationV1) GetAccountOk() (*AuthAccountV1, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AuthorizationV1) SetAccount(v AuthAccountV1)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AuthorizationV1) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccountGroups

`func (o *AuthorizationV1) GetAccountGroups() []AccountGroup`

GetAccountGroups returns the AccountGroups field if non-nil, zero value otherwise.

### GetAccountGroupsOk

`func (o *AuthorizationV1) GetAccountGroupsOk() (*[]AccountGroup, bool)`

GetAccountGroupsOk returns a tuple with the AccountGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroups

`func (o *AuthorizationV1) SetAccountGroups(v []AccountGroup)`

SetAccountGroups sets AccountGroups field to given value.

### HasAccountGroups

`func (o *AuthorizationV1) HasAccountGroups() bool`

HasAccountGroups returns a boolean if a field has been set.

### GetSites

`func (o *AuthorizationV1) GetSites() []V1Site`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *AuthorizationV1) GetSitesOk() (*[]V1Site, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *AuthorizationV1) SetSites(v []V1Site)`

SetSites sets Sites field to given value.

### HasSites

`func (o *AuthorizationV1) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *AuthorizationV1) GetAuthenticationType() AuthenticationType`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *AuthorizationV1) GetAuthenticationTypeOk() (*AuthenticationType, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *AuthorizationV1) SetAuthenticationType(v AuthenticationType)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *AuthorizationV1) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


