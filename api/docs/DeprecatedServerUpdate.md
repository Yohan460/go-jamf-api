# DeprecatedServerUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Enabled** | **bool** |  | 
**ProviderName** | **string** |  | 
**DisplayName** | **string** |  | 
**ServerUrl** | **string** |  | 
**DomainName** | **string** |  | 
**Port** | **int32** |  | 
**Keystore** | Pointer to [**CloudLdapKeystoreFile**](CloudLdapKeystoreFile.md) |  | [optional] 
**ConnectionTimeout** | **int32** |  | 
**SearchTimeout** | **int32** |  | 
**UseWildcards** | **bool** |  | 
**ConnectionType** | **string** |  | 

## Methods

### NewDeprecatedServerUpdate

`func NewDeprecatedServerUpdate(id string, enabled bool, providerName string, displayName string, serverUrl string, domainName string, port int32, connectionTimeout int32, searchTimeout int32, useWildcards bool, connectionType string, ) *DeprecatedServerUpdate`

NewDeprecatedServerUpdate instantiates a new DeprecatedServerUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeprecatedServerUpdateWithDefaults

`func NewDeprecatedServerUpdateWithDefaults() *DeprecatedServerUpdate`

NewDeprecatedServerUpdateWithDefaults instantiates a new DeprecatedServerUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeprecatedServerUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeprecatedServerUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeprecatedServerUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetEnabled

`func (o *DeprecatedServerUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeprecatedServerUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeprecatedServerUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetProviderName

`func (o *DeprecatedServerUpdate) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *DeprecatedServerUpdate) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *DeprecatedServerUpdate) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetDisplayName

`func (o *DeprecatedServerUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DeprecatedServerUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DeprecatedServerUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetServerUrl

`func (o *DeprecatedServerUpdate) GetServerUrl() string`

GetServerUrl returns the ServerUrl field if non-nil, zero value otherwise.

### GetServerUrlOk

`func (o *DeprecatedServerUpdate) GetServerUrlOk() (*string, bool)`

GetServerUrlOk returns a tuple with the ServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUrl

`func (o *DeprecatedServerUpdate) SetServerUrl(v string)`

SetServerUrl sets ServerUrl field to given value.


### GetDomainName

`func (o *DeprecatedServerUpdate) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DeprecatedServerUpdate) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DeprecatedServerUpdate) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetPort

`func (o *DeprecatedServerUpdate) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DeprecatedServerUpdate) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DeprecatedServerUpdate) SetPort(v int32)`

SetPort sets Port field to given value.


### GetKeystore

`func (o *DeprecatedServerUpdate) GetKeystore() CloudLdapKeystoreFile`

GetKeystore returns the Keystore field if non-nil, zero value otherwise.

### GetKeystoreOk

`func (o *DeprecatedServerUpdate) GetKeystoreOk() (*CloudLdapKeystoreFile, bool)`

GetKeystoreOk returns a tuple with the Keystore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystore

`func (o *DeprecatedServerUpdate) SetKeystore(v CloudLdapKeystoreFile)`

SetKeystore sets Keystore field to given value.

### HasKeystore

`func (o *DeprecatedServerUpdate) HasKeystore() bool`

HasKeystore returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *DeprecatedServerUpdate) GetConnectionTimeout() int32`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *DeprecatedServerUpdate) GetConnectionTimeoutOk() (*int32, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *DeprecatedServerUpdate) SetConnectionTimeout(v int32)`

SetConnectionTimeout sets ConnectionTimeout field to given value.


### GetSearchTimeout

`func (o *DeprecatedServerUpdate) GetSearchTimeout() int32`

GetSearchTimeout returns the SearchTimeout field if non-nil, zero value otherwise.

### GetSearchTimeoutOk

`func (o *DeprecatedServerUpdate) GetSearchTimeoutOk() (*int32, bool)`

GetSearchTimeoutOk returns a tuple with the SearchTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTimeout

`func (o *DeprecatedServerUpdate) SetSearchTimeout(v int32)`

SetSearchTimeout sets SearchTimeout field to given value.


### GetUseWildcards

`func (o *DeprecatedServerUpdate) GetUseWildcards() bool`

GetUseWildcards returns the UseWildcards field if non-nil, zero value otherwise.

### GetUseWildcardsOk

`func (o *DeprecatedServerUpdate) GetUseWildcardsOk() (*bool, bool)`

GetUseWildcardsOk returns a tuple with the UseWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWildcards

`func (o *DeprecatedServerUpdate) SetUseWildcards(v bool)`

SetUseWildcards sets UseWildcards field to given value.


### GetConnectionType

`func (o *DeprecatedServerUpdate) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *DeprecatedServerUpdate) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *DeprecatedServerUpdate) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


