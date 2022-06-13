# CloudLdapServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ServerUrl** | Pointer to **string** |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Keystore** | Pointer to [**CloudLdapKeystore**](CloudLdapKeystore.md) |  | [optional] 
**ConnectionTimeout** | Pointer to **int32** |  | [optional] 
**SearchTimeout** | Pointer to **int32** |  | [optional] 
**UseWildcards** | Pointer to **bool** |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudLdapServerResponse

`func NewCloudLdapServerResponse() *CloudLdapServerResponse`

NewCloudLdapServerResponse instantiates a new CloudLdapServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudLdapServerResponseWithDefaults

`func NewCloudLdapServerResponseWithDefaults() *CloudLdapServerResponse`

NewCloudLdapServerResponseWithDefaults instantiates a new CloudLdapServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudLdapServerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudLdapServerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudLdapServerResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudLdapServerResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnabled

`func (o *CloudLdapServerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CloudLdapServerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CloudLdapServerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CloudLdapServerResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServerUrl

`func (o *CloudLdapServerResponse) GetServerUrl() string`

GetServerUrl returns the ServerUrl field if non-nil, zero value otherwise.

### GetServerUrlOk

`func (o *CloudLdapServerResponse) GetServerUrlOk() (*string, bool)`

GetServerUrlOk returns a tuple with the ServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUrl

`func (o *CloudLdapServerResponse) SetServerUrl(v string)`

SetServerUrl sets ServerUrl field to given value.

### HasServerUrl

`func (o *CloudLdapServerResponse) HasServerUrl() bool`

HasServerUrl returns a boolean if a field has been set.

### GetDomainName

`func (o *CloudLdapServerResponse) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CloudLdapServerResponse) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CloudLdapServerResponse) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *CloudLdapServerResponse) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetPort

`func (o *CloudLdapServerResponse) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CloudLdapServerResponse) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CloudLdapServerResponse) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CloudLdapServerResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetKeystore

`func (o *CloudLdapServerResponse) GetKeystore() CloudLdapKeystore`

GetKeystore returns the Keystore field if non-nil, zero value otherwise.

### GetKeystoreOk

`func (o *CloudLdapServerResponse) GetKeystoreOk() (*CloudLdapKeystore, bool)`

GetKeystoreOk returns a tuple with the Keystore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystore

`func (o *CloudLdapServerResponse) SetKeystore(v CloudLdapKeystore)`

SetKeystore sets Keystore field to given value.

### HasKeystore

`func (o *CloudLdapServerResponse) HasKeystore() bool`

HasKeystore returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *CloudLdapServerResponse) GetConnectionTimeout() int32`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *CloudLdapServerResponse) GetConnectionTimeoutOk() (*int32, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *CloudLdapServerResponse) SetConnectionTimeout(v int32)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *CloudLdapServerResponse) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetSearchTimeout

`func (o *CloudLdapServerResponse) GetSearchTimeout() int32`

GetSearchTimeout returns the SearchTimeout field if non-nil, zero value otherwise.

### GetSearchTimeoutOk

`func (o *CloudLdapServerResponse) GetSearchTimeoutOk() (*int32, bool)`

GetSearchTimeoutOk returns a tuple with the SearchTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTimeout

`func (o *CloudLdapServerResponse) SetSearchTimeout(v int32)`

SetSearchTimeout sets SearchTimeout field to given value.

### HasSearchTimeout

`func (o *CloudLdapServerResponse) HasSearchTimeout() bool`

HasSearchTimeout returns a boolean if a field has been set.

### GetUseWildcards

`func (o *CloudLdapServerResponse) GetUseWildcards() bool`

GetUseWildcards returns the UseWildcards field if non-nil, zero value otherwise.

### GetUseWildcardsOk

`func (o *CloudLdapServerResponse) GetUseWildcardsOk() (*bool, bool)`

GetUseWildcardsOk returns a tuple with the UseWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWildcards

`func (o *CloudLdapServerResponse) SetUseWildcards(v bool)`

SetUseWildcards sets UseWildcards field to given value.

### HasUseWildcards

`func (o *CloudLdapServerResponse) HasUseWildcards() bool`

HasUseWildcards returns a boolean if a field has been set.

### GetConnectionType

`func (o *CloudLdapServerResponse) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CloudLdapServerResponse) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *CloudLdapServerResponse) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *CloudLdapServerResponse) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


