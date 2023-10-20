# CloudLdapServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerUrl** | **string** |  | 
**Enabled** | **bool** |  | 
**DomainName** | **string** |  | 
**Port** | **int32** |  | 
**Keystore** | [**CloudLdapKeystoreFile**](CloudLdapKeystoreFile.md) |  | 
**ConnectionTimeout** | **int32** |  | 
**SearchTimeout** | **int32** |  | 
**UseWildcards** | **bool** |  | 
**ConnectionType** | **string** |  | 
**MembershipCalculationOptimizationEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCloudLdapServerRequest

`func NewCloudLdapServerRequest(serverUrl string, enabled bool, domainName string, port int32, keystore CloudLdapKeystoreFile, connectionTimeout int32, searchTimeout int32, useWildcards bool, connectionType string, ) *CloudLdapServerRequest`

NewCloudLdapServerRequest instantiates a new CloudLdapServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudLdapServerRequestWithDefaults

`func NewCloudLdapServerRequestWithDefaults() *CloudLdapServerRequest`

NewCloudLdapServerRequestWithDefaults instantiates a new CloudLdapServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerUrl

`func (o *CloudLdapServerRequest) GetServerUrl() string`

GetServerUrl returns the ServerUrl field if non-nil, zero value otherwise.

### GetServerUrlOk

`func (o *CloudLdapServerRequest) GetServerUrlOk() (*string, bool)`

GetServerUrlOk returns a tuple with the ServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUrl

`func (o *CloudLdapServerRequest) SetServerUrl(v string)`

SetServerUrl sets ServerUrl field to given value.


### GetEnabled

`func (o *CloudLdapServerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CloudLdapServerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CloudLdapServerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDomainName

`func (o *CloudLdapServerRequest) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CloudLdapServerRequest) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CloudLdapServerRequest) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetPort

`func (o *CloudLdapServerRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CloudLdapServerRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CloudLdapServerRequest) SetPort(v int32)`

SetPort sets Port field to given value.


### GetKeystore

`func (o *CloudLdapServerRequest) GetKeystore() CloudLdapKeystoreFile`

GetKeystore returns the Keystore field if non-nil, zero value otherwise.

### GetKeystoreOk

`func (o *CloudLdapServerRequest) GetKeystoreOk() (*CloudLdapKeystoreFile, bool)`

GetKeystoreOk returns a tuple with the Keystore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystore

`func (o *CloudLdapServerRequest) SetKeystore(v CloudLdapKeystoreFile)`

SetKeystore sets Keystore field to given value.


### GetConnectionTimeout

`func (o *CloudLdapServerRequest) GetConnectionTimeout() int32`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *CloudLdapServerRequest) GetConnectionTimeoutOk() (*int32, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *CloudLdapServerRequest) SetConnectionTimeout(v int32)`

SetConnectionTimeout sets ConnectionTimeout field to given value.


### GetSearchTimeout

`func (o *CloudLdapServerRequest) GetSearchTimeout() int32`

GetSearchTimeout returns the SearchTimeout field if non-nil, zero value otherwise.

### GetSearchTimeoutOk

`func (o *CloudLdapServerRequest) GetSearchTimeoutOk() (*int32, bool)`

GetSearchTimeoutOk returns a tuple with the SearchTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTimeout

`func (o *CloudLdapServerRequest) SetSearchTimeout(v int32)`

SetSearchTimeout sets SearchTimeout field to given value.


### GetUseWildcards

`func (o *CloudLdapServerRequest) GetUseWildcards() bool`

GetUseWildcards returns the UseWildcards field if non-nil, zero value otherwise.

### GetUseWildcardsOk

`func (o *CloudLdapServerRequest) GetUseWildcardsOk() (*bool, bool)`

GetUseWildcardsOk returns a tuple with the UseWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWildcards

`func (o *CloudLdapServerRequest) SetUseWildcards(v bool)`

SetUseWildcards sets UseWildcards field to given value.


### GetConnectionType

`func (o *CloudLdapServerRequest) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CloudLdapServerRequest) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *CloudLdapServerRequest) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetMembershipCalculationOptimizationEnabled

`func (o *CloudLdapServerRequest) GetMembershipCalculationOptimizationEnabled() bool`

GetMembershipCalculationOptimizationEnabled returns the MembershipCalculationOptimizationEnabled field if non-nil, zero value otherwise.

### GetMembershipCalculationOptimizationEnabledOk

`func (o *CloudLdapServerRequest) GetMembershipCalculationOptimizationEnabledOk() (*bool, bool)`

GetMembershipCalculationOptimizationEnabledOk returns a tuple with the MembershipCalculationOptimizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipCalculationOptimizationEnabled

`func (o *CloudLdapServerRequest) SetMembershipCalculationOptimizationEnabled(v bool)`

SetMembershipCalculationOptimizationEnabled sets MembershipCalculationOptimizationEnabled field to given value.

### HasMembershipCalculationOptimizationEnabled

`func (o *CloudLdapServerRequest) HasMembershipCalculationOptimizationEnabled() bool`

HasMembershipCalculationOptimizationEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


