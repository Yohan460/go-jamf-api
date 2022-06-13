# DeprecatedServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ProviderName** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ServerUrl** | Pointer to **string** |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Keystore** | Pointer to [**CloudLdapKeystore**](CloudLdapKeystore.md) |  | [optional] 
**ConnectionTimeout** | Pointer to **int32** |  | [optional] 
**SearchTimeout** | Pointer to **int32** |  | [optional] 
**UseWildcards** | Pointer to **bool** |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 

## Methods

### NewDeprecatedServerResponse

`func NewDeprecatedServerResponse() *DeprecatedServerResponse`

NewDeprecatedServerResponse instantiates a new DeprecatedServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeprecatedServerResponseWithDefaults

`func NewDeprecatedServerResponseWithDefaults() *DeprecatedServerResponse`

NewDeprecatedServerResponseWithDefaults instantiates a new DeprecatedServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeprecatedServerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeprecatedServerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeprecatedServerResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeprecatedServerResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnabled

`func (o *DeprecatedServerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeprecatedServerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeprecatedServerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DeprecatedServerResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProviderName

`func (o *DeprecatedServerResponse) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *DeprecatedServerResponse) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *DeprecatedServerResponse) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *DeprecatedServerResponse) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetDisplayName

`func (o *DeprecatedServerResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DeprecatedServerResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DeprecatedServerResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DeprecatedServerResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetServerUrl

`func (o *DeprecatedServerResponse) GetServerUrl() string`

GetServerUrl returns the ServerUrl field if non-nil, zero value otherwise.

### GetServerUrlOk

`func (o *DeprecatedServerResponse) GetServerUrlOk() (*string, bool)`

GetServerUrlOk returns a tuple with the ServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUrl

`func (o *DeprecatedServerResponse) SetServerUrl(v string)`

SetServerUrl sets ServerUrl field to given value.

### HasServerUrl

`func (o *DeprecatedServerResponse) HasServerUrl() bool`

HasServerUrl returns a boolean if a field has been set.

### GetDomainName

`func (o *DeprecatedServerResponse) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DeprecatedServerResponse) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DeprecatedServerResponse) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *DeprecatedServerResponse) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetPort

`func (o *DeprecatedServerResponse) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DeprecatedServerResponse) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DeprecatedServerResponse) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *DeprecatedServerResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetKeystore

`func (o *DeprecatedServerResponse) GetKeystore() CloudLdapKeystore`

GetKeystore returns the Keystore field if non-nil, zero value otherwise.

### GetKeystoreOk

`func (o *DeprecatedServerResponse) GetKeystoreOk() (*CloudLdapKeystore, bool)`

GetKeystoreOk returns a tuple with the Keystore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystore

`func (o *DeprecatedServerResponse) SetKeystore(v CloudLdapKeystore)`

SetKeystore sets Keystore field to given value.

### HasKeystore

`func (o *DeprecatedServerResponse) HasKeystore() bool`

HasKeystore returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *DeprecatedServerResponse) GetConnectionTimeout() int32`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *DeprecatedServerResponse) GetConnectionTimeoutOk() (*int32, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *DeprecatedServerResponse) SetConnectionTimeout(v int32)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *DeprecatedServerResponse) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetSearchTimeout

`func (o *DeprecatedServerResponse) GetSearchTimeout() int32`

GetSearchTimeout returns the SearchTimeout field if non-nil, zero value otherwise.

### GetSearchTimeoutOk

`func (o *DeprecatedServerResponse) GetSearchTimeoutOk() (*int32, bool)`

GetSearchTimeoutOk returns a tuple with the SearchTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTimeout

`func (o *DeprecatedServerResponse) SetSearchTimeout(v int32)`

SetSearchTimeout sets SearchTimeout field to given value.

### HasSearchTimeout

`func (o *DeprecatedServerResponse) HasSearchTimeout() bool`

HasSearchTimeout returns a boolean if a field has been set.

### GetUseWildcards

`func (o *DeprecatedServerResponse) GetUseWildcards() bool`

GetUseWildcards returns the UseWildcards field if non-nil, zero value otherwise.

### GetUseWildcardsOk

`func (o *DeprecatedServerResponse) GetUseWildcardsOk() (*bool, bool)`

GetUseWildcardsOk returns a tuple with the UseWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWildcards

`func (o *DeprecatedServerResponse) SetUseWildcards(v bool)`

SetUseWildcards sets UseWildcards field to given value.

### HasUseWildcards

`func (o *DeprecatedServerResponse) HasUseWildcards() bool`

HasUseWildcards returns a boolean if a field has been set.

### GetConnectionType

`func (o *DeprecatedServerResponse) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *DeprecatedServerResponse) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *DeprecatedServerResponse) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *DeprecatedServerResponse) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


