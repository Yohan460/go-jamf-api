# DeprecatedServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**ProviderName** | **string** |  | 
**DisplayName** | **string** |  | 
**ServerUrl** | **string** |  | 
**DomainName** | **string** |  | 
**Port** | **int32** |  | 
**Keystore** | [**CloudLdapKeystoreFile**](CloudLdapKeystoreFile.md) |  | 
**ConnectionTimeout** | **int32** |  | 
**SearchTimeout** | **int32** |  | 
**UseWildcards** | **bool** |  | 
**ConnectionType** | **string** |  | 

## Methods

### NewDeprecatedServerRequest

`func NewDeprecatedServerRequest(enabled bool, providerName string, displayName string, serverUrl string, domainName string, port int32, keystore CloudLdapKeystoreFile, connectionTimeout int32, searchTimeout int32, useWildcards bool, connectionType string, ) *DeprecatedServerRequest`

NewDeprecatedServerRequest instantiates a new DeprecatedServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeprecatedServerRequestWithDefaults

`func NewDeprecatedServerRequestWithDefaults() *DeprecatedServerRequest`

NewDeprecatedServerRequestWithDefaults instantiates a new DeprecatedServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DeprecatedServerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeprecatedServerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeprecatedServerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetProviderName

`func (o *DeprecatedServerRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *DeprecatedServerRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *DeprecatedServerRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.


### GetDisplayName

`func (o *DeprecatedServerRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DeprecatedServerRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DeprecatedServerRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetServerUrl

`func (o *DeprecatedServerRequest) GetServerUrl() string`

GetServerUrl returns the ServerUrl field if non-nil, zero value otherwise.

### GetServerUrlOk

`func (o *DeprecatedServerRequest) GetServerUrlOk() (*string, bool)`

GetServerUrlOk returns a tuple with the ServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUrl

`func (o *DeprecatedServerRequest) SetServerUrl(v string)`

SetServerUrl sets ServerUrl field to given value.


### GetDomainName

`func (o *DeprecatedServerRequest) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DeprecatedServerRequest) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DeprecatedServerRequest) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetPort

`func (o *DeprecatedServerRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DeprecatedServerRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DeprecatedServerRequest) SetPort(v int32)`

SetPort sets Port field to given value.


### GetKeystore

`func (o *DeprecatedServerRequest) GetKeystore() CloudLdapKeystoreFile`

GetKeystore returns the Keystore field if non-nil, zero value otherwise.

### GetKeystoreOk

`func (o *DeprecatedServerRequest) GetKeystoreOk() (*CloudLdapKeystoreFile, bool)`

GetKeystoreOk returns a tuple with the Keystore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystore

`func (o *DeprecatedServerRequest) SetKeystore(v CloudLdapKeystoreFile)`

SetKeystore sets Keystore field to given value.


### GetConnectionTimeout

`func (o *DeprecatedServerRequest) GetConnectionTimeout() int32`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *DeprecatedServerRequest) GetConnectionTimeoutOk() (*int32, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *DeprecatedServerRequest) SetConnectionTimeout(v int32)`

SetConnectionTimeout sets ConnectionTimeout field to given value.


### GetSearchTimeout

`func (o *DeprecatedServerRequest) GetSearchTimeout() int32`

GetSearchTimeout returns the SearchTimeout field if non-nil, zero value otherwise.

### GetSearchTimeoutOk

`func (o *DeprecatedServerRequest) GetSearchTimeoutOk() (*int32, bool)`

GetSearchTimeoutOk returns a tuple with the SearchTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTimeout

`func (o *DeprecatedServerRequest) SetSearchTimeout(v int32)`

SetSearchTimeout sets SearchTimeout field to given value.


### GetUseWildcards

`func (o *DeprecatedServerRequest) GetUseWildcards() bool`

GetUseWildcards returns the UseWildcards field if non-nil, zero value otherwise.

### GetUseWildcardsOk

`func (o *DeprecatedServerRequest) GetUseWildcardsOk() (*bool, bool)`

GetUseWildcardsOk returns a tuple with the UseWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWildcards

`func (o *DeprecatedServerRequest) SetUseWildcards(v bool)`

SetUseWildcards sets UseWildcards field to given value.


### GetConnectionType

`func (o *DeprecatedServerRequest) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *DeprecatedServerRequest) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *DeprecatedServerRequest) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


