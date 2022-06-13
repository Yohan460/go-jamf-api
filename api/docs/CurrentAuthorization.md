# CurrentAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**CurrentAccount**](CurrentAccount.md) |  | [optional] 
**Sites** | Pointer to [**[]Site**](Site.md) |  | [optional] 
**AuthenticationType** | Pointer to **string** |  | [optional] 

## Methods

### NewCurrentAuthorization

`func NewCurrentAuthorization() *CurrentAuthorization`

NewCurrentAuthorization instantiates a new CurrentAuthorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentAuthorizationWithDefaults

`func NewCurrentAuthorizationWithDefaults() *CurrentAuthorization`

NewCurrentAuthorizationWithDefaults instantiates a new CurrentAuthorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CurrentAuthorization) GetAccount() CurrentAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CurrentAuthorization) GetAccountOk() (*CurrentAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CurrentAuthorization) SetAccount(v CurrentAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CurrentAuthorization) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *CurrentAuthorization) GetSites() []Site`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *CurrentAuthorization) GetSitesOk() (*[]Site, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *CurrentAuthorization) SetSites(v []Site)`

SetSites sets Sites field to given value.

### HasSites

`func (o *CurrentAuthorization) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *CurrentAuthorization) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *CurrentAuthorization) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *CurrentAuthorization) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *CurrentAuthorization) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


