# CloudLdapServerUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerUrl** | **string** |  | 
**Enabled** | **bool** |  | 
**DomainName** | **string** |  | 
**Port** | **int32** |  | 
**Keystore** | Pointer to [**CloudLdapKeystoreFile**](CloudLdapKeystoreFile.md) |  | [optional] 
**ConnectionTimeout** | **int32** |  | 
**SearchTimeout** | **int32** |  | 
**UseWildcards** | **bool** |  | 
**ConnectionType** | **string** |  | 
**MembershipCalculationOptimizationEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCloudLdapServerUpdate

`func NewCloudLdapServerUpdate(serverUrl string, enabled bool, domainName string, port int32, connectionTimeout int32, searchTimeout int32, useWildcards bool, connectionType string, ) *CloudLdapServerUpdate`

NewCloudLdapServerUpdate instantiates a new CloudLdapServerUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudLdapServerUpdateWithDefaults

`func NewCloudLdapServerUpdateWithDefaults() *CloudLdapServerUpdate`

NewCloudLdapServerUpdateWithDefaults instantiates a new CloudLdapServerUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerUrl

`func (o *CloudLdapServerUpdate) GetServerUrl() string`

GetServerUrl returns the ServerUrl field if non-nil, zero value otherwise.

### GetServerUrlOk

`func (o *CloudLdapServerUpdate) GetServerUrlOk() (*string, bool)`

GetServerUrlOk returns a tuple with the ServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUrl

`func (o *CloudLdapServerUpdate) SetServerUrl(v string)`

SetServerUrl sets ServerUrl field to given value.


### GetEnabled

`func (o *CloudLdapServerUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CloudLdapServerUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CloudLdapServerUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDomainName

`func (o *CloudLdapServerUpdate) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CloudLdapServerUpdate) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CloudLdapServerUpdate) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetPort

`func (o *CloudLdapServerUpdate) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CloudLdapServerUpdate) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CloudLdapServerUpdate) SetPort(v int32)`

SetPort sets Port field to given value.


### GetKeystore

`func (o *CloudLdapServerUpdate) GetKeystore() CloudLdapKeystoreFile`

GetKeystore returns the Keystore field if non-nil, zero value otherwise.

### GetKeystoreOk

`func (o *CloudLdapServerUpdate) GetKeystoreOk() (*CloudLdapKeystoreFile, bool)`

GetKeystoreOk returns a tuple with the Keystore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystore

`func (o *CloudLdapServerUpdate) SetKeystore(v CloudLdapKeystoreFile)`

SetKeystore sets Keystore field to given value.

### HasKeystore

`func (o *CloudLdapServerUpdate) HasKeystore() bool`

HasKeystore returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *CloudLdapServerUpdate) GetConnectionTimeout() int32`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *CloudLdapServerUpdate) GetConnectionTimeoutOk() (*int32, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *CloudLdapServerUpdate) SetConnectionTimeout(v int32)`

SetConnectionTimeout sets ConnectionTimeout field to given value.


### GetSearchTimeout

`func (o *CloudLdapServerUpdate) GetSearchTimeout() int32`

GetSearchTimeout returns the SearchTimeout field if non-nil, zero value otherwise.

### GetSearchTimeoutOk

`func (o *CloudLdapServerUpdate) GetSearchTimeoutOk() (*int32, bool)`

GetSearchTimeoutOk returns a tuple with the SearchTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTimeout

`func (o *CloudLdapServerUpdate) SetSearchTimeout(v int32)`

SetSearchTimeout sets SearchTimeout field to given value.


### GetUseWildcards

`func (o *CloudLdapServerUpdate) GetUseWildcards() bool`

GetUseWildcards returns the UseWildcards field if non-nil, zero value otherwise.

### GetUseWildcardsOk

`func (o *CloudLdapServerUpdate) GetUseWildcardsOk() (*bool, bool)`

GetUseWildcardsOk returns a tuple with the UseWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWildcards

`func (o *CloudLdapServerUpdate) SetUseWildcards(v bool)`

SetUseWildcards sets UseWildcards field to given value.


### GetConnectionType

`func (o *CloudLdapServerUpdate) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CloudLdapServerUpdate) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *CloudLdapServerUpdate) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetMembershipCalculationOptimizationEnabled

`func (o *CloudLdapServerUpdate) GetMembershipCalculationOptimizationEnabled() bool`

GetMembershipCalculationOptimizationEnabled returns the MembershipCalculationOptimizationEnabled field if non-nil, zero value otherwise.

### GetMembershipCalculationOptimizationEnabledOk

`func (o *CloudLdapServerUpdate) GetMembershipCalculationOptimizationEnabledOk() (*bool, bool)`

GetMembershipCalculationOptimizationEnabledOk returns a tuple with the MembershipCalculationOptimizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipCalculationOptimizationEnabled

`func (o *CloudLdapServerUpdate) SetMembershipCalculationOptimizationEnabled(v bool)`

SetMembershipCalculationOptimizationEnabled sets MembershipCalculationOptimizationEnabled field to given value.

### HasMembershipCalculationOptimizationEnabled

`func (o *CloudLdapServerUpdate) HasMembershipCalculationOptimizationEnabled() bool`

HasMembershipCalculationOptimizationEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


