# AzureServerConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**TenantId** | **string** |  | 
**Enabled** | **bool** |  | 
**Mappings** | [**AzureMappings**](AzureMappings.md) |  | 
**SearchTimeout** | **int64** |  | 
**TransitiveMembershipEnabled** | **bool** | Use this field to enable transitive membership lookup with Single Sign On | 
**TransitiveMembershipUserField** | **string** | Use this field to set user field mapping for transitive membership lookup with Single Sign On | 
**TransitiveDirectoryMembershipEnabled** | **bool** | Use this field to enable transitive membership lookup. This setting would not apply to Single Sign On | 
**MembershipCalculationOptimizationEnabled** | Pointer to **bool** | Use this field to enable membership calculation optimization. This setting would not apply to Single Sign On | [optional] 
**Code** | **string** |  | 

## Methods

### NewAzureServerConfigurationRequest

`func NewAzureServerConfigurationRequest(tenantId string, enabled bool, mappings AzureMappings, searchTimeout int64, transitiveMembershipEnabled bool, transitiveMembershipUserField string, transitiveDirectoryMembershipEnabled bool, code string, ) *AzureServerConfigurationRequest`

NewAzureServerConfigurationRequest instantiates a new AzureServerConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureServerConfigurationRequestWithDefaults

`func NewAzureServerConfigurationRequestWithDefaults() *AzureServerConfigurationRequest`

NewAzureServerConfigurationRequestWithDefaults instantiates a new AzureServerConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AzureServerConfigurationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureServerConfigurationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureServerConfigurationRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureServerConfigurationRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTenantId

`func (o *AzureServerConfigurationRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureServerConfigurationRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureServerConfigurationRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetEnabled

`func (o *AzureServerConfigurationRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AzureServerConfigurationRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AzureServerConfigurationRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMappings

`func (o *AzureServerConfigurationRequest) GetMappings() AzureMappings`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *AzureServerConfigurationRequest) GetMappingsOk() (*AzureMappings, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *AzureServerConfigurationRequest) SetMappings(v AzureMappings)`

SetMappings sets Mappings field to given value.


### GetSearchTimeout

`func (o *AzureServerConfigurationRequest) GetSearchTimeout() int64`

GetSearchTimeout returns the SearchTimeout field if non-nil, zero value otherwise.

### GetSearchTimeoutOk

`func (o *AzureServerConfigurationRequest) GetSearchTimeoutOk() (*int64, bool)`

GetSearchTimeoutOk returns a tuple with the SearchTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTimeout

`func (o *AzureServerConfigurationRequest) SetSearchTimeout(v int64)`

SetSearchTimeout sets SearchTimeout field to given value.


### GetTransitiveMembershipEnabled

`func (o *AzureServerConfigurationRequest) GetTransitiveMembershipEnabled() bool`

GetTransitiveMembershipEnabled returns the TransitiveMembershipEnabled field if non-nil, zero value otherwise.

### GetTransitiveMembershipEnabledOk

`func (o *AzureServerConfigurationRequest) GetTransitiveMembershipEnabledOk() (*bool, bool)`

GetTransitiveMembershipEnabledOk returns a tuple with the TransitiveMembershipEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveMembershipEnabled

`func (o *AzureServerConfigurationRequest) SetTransitiveMembershipEnabled(v bool)`

SetTransitiveMembershipEnabled sets TransitiveMembershipEnabled field to given value.


### GetTransitiveMembershipUserField

`func (o *AzureServerConfigurationRequest) GetTransitiveMembershipUserField() string`

GetTransitiveMembershipUserField returns the TransitiveMembershipUserField field if non-nil, zero value otherwise.

### GetTransitiveMembershipUserFieldOk

`func (o *AzureServerConfigurationRequest) GetTransitiveMembershipUserFieldOk() (*string, bool)`

GetTransitiveMembershipUserFieldOk returns a tuple with the TransitiveMembershipUserField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveMembershipUserField

`func (o *AzureServerConfigurationRequest) SetTransitiveMembershipUserField(v string)`

SetTransitiveMembershipUserField sets TransitiveMembershipUserField field to given value.


### GetTransitiveDirectoryMembershipEnabled

`func (o *AzureServerConfigurationRequest) GetTransitiveDirectoryMembershipEnabled() bool`

GetTransitiveDirectoryMembershipEnabled returns the TransitiveDirectoryMembershipEnabled field if non-nil, zero value otherwise.

### GetTransitiveDirectoryMembershipEnabledOk

`func (o *AzureServerConfigurationRequest) GetTransitiveDirectoryMembershipEnabledOk() (*bool, bool)`

GetTransitiveDirectoryMembershipEnabledOk returns a tuple with the TransitiveDirectoryMembershipEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveDirectoryMembershipEnabled

`func (o *AzureServerConfigurationRequest) SetTransitiveDirectoryMembershipEnabled(v bool)`

SetTransitiveDirectoryMembershipEnabled sets TransitiveDirectoryMembershipEnabled field to given value.


### GetMembershipCalculationOptimizationEnabled

`func (o *AzureServerConfigurationRequest) GetMembershipCalculationOptimizationEnabled() bool`

GetMembershipCalculationOptimizationEnabled returns the MembershipCalculationOptimizationEnabled field if non-nil, zero value otherwise.

### GetMembershipCalculationOptimizationEnabledOk

`func (o *AzureServerConfigurationRequest) GetMembershipCalculationOptimizationEnabledOk() (*bool, bool)`

GetMembershipCalculationOptimizationEnabledOk returns a tuple with the MembershipCalculationOptimizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipCalculationOptimizationEnabled

`func (o *AzureServerConfigurationRequest) SetMembershipCalculationOptimizationEnabled(v bool)`

SetMembershipCalculationOptimizationEnabled sets MembershipCalculationOptimizationEnabled field to given value.

### HasMembershipCalculationOptimizationEnabled

`func (o *AzureServerConfigurationRequest) HasMembershipCalculationOptimizationEnabled() bool`

HasMembershipCalculationOptimizationEnabled returns a boolean if a field has been set.

### GetCode

`func (o *AzureServerConfigurationRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AzureServerConfigurationRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AzureServerConfigurationRequest) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


