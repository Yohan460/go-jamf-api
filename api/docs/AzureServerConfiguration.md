# AzureServerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**Enabled** | **bool** |  | 
**Migrated** | **bool** |  | 
**Mappings** | [**AzureMappings**](AzureMappings.md) |  | 
**SearchTimeout** | **int64** |  | 
**TransitiveMembershipEnabled** | **bool** | Use this field to enable transitive membership lookup with Single Sign On | 
**TransitiveMembershipUserField** | **string** | Use this field to set user field mapping for transitive membership lookup with Single Sign On | 
**TransitiveDirectoryMembershipEnabled** | **bool** | Use this field to enable transitive membership lookup. This setting would not apply to Single Sign On | 
**MembershipCalculationOptimizationEnabled** | Pointer to **bool** | Use this field to enable membership calculation optimization. This setting would not apply to Single Sign On | [optional] 

## Methods

### NewAzureServerConfiguration

`func NewAzureServerConfiguration(id string, tenantId string, enabled bool, migrated bool, mappings AzureMappings, searchTimeout int64, transitiveMembershipEnabled bool, transitiveMembershipUserField string, transitiveDirectoryMembershipEnabled bool, ) *AzureServerConfiguration`

NewAzureServerConfiguration instantiates a new AzureServerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureServerConfigurationWithDefaults

`func NewAzureServerConfigurationWithDefaults() *AzureServerConfiguration`

NewAzureServerConfigurationWithDefaults instantiates a new AzureServerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AzureServerConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureServerConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureServerConfiguration) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *AzureServerConfiguration) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureServerConfiguration) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureServerConfiguration) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetEnabled

`func (o *AzureServerConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AzureServerConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AzureServerConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMigrated

`func (o *AzureServerConfiguration) GetMigrated() bool`

GetMigrated returns the Migrated field if non-nil, zero value otherwise.

### GetMigratedOk

`func (o *AzureServerConfiguration) GetMigratedOk() (*bool, bool)`

GetMigratedOk returns a tuple with the Migrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrated

`func (o *AzureServerConfiguration) SetMigrated(v bool)`

SetMigrated sets Migrated field to given value.


### GetMappings

`func (o *AzureServerConfiguration) GetMappings() AzureMappings`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *AzureServerConfiguration) GetMappingsOk() (*AzureMappings, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *AzureServerConfiguration) SetMappings(v AzureMappings)`

SetMappings sets Mappings field to given value.


### GetSearchTimeout

`func (o *AzureServerConfiguration) GetSearchTimeout() int64`

GetSearchTimeout returns the SearchTimeout field if non-nil, zero value otherwise.

### GetSearchTimeoutOk

`func (o *AzureServerConfiguration) GetSearchTimeoutOk() (*int64, bool)`

GetSearchTimeoutOk returns a tuple with the SearchTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTimeout

`func (o *AzureServerConfiguration) SetSearchTimeout(v int64)`

SetSearchTimeout sets SearchTimeout field to given value.


### GetTransitiveMembershipEnabled

`func (o *AzureServerConfiguration) GetTransitiveMembershipEnabled() bool`

GetTransitiveMembershipEnabled returns the TransitiveMembershipEnabled field if non-nil, zero value otherwise.

### GetTransitiveMembershipEnabledOk

`func (o *AzureServerConfiguration) GetTransitiveMembershipEnabledOk() (*bool, bool)`

GetTransitiveMembershipEnabledOk returns a tuple with the TransitiveMembershipEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveMembershipEnabled

`func (o *AzureServerConfiguration) SetTransitiveMembershipEnabled(v bool)`

SetTransitiveMembershipEnabled sets TransitiveMembershipEnabled field to given value.


### GetTransitiveMembershipUserField

`func (o *AzureServerConfiguration) GetTransitiveMembershipUserField() string`

GetTransitiveMembershipUserField returns the TransitiveMembershipUserField field if non-nil, zero value otherwise.

### GetTransitiveMembershipUserFieldOk

`func (o *AzureServerConfiguration) GetTransitiveMembershipUserFieldOk() (*string, bool)`

GetTransitiveMembershipUserFieldOk returns a tuple with the TransitiveMembershipUserField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveMembershipUserField

`func (o *AzureServerConfiguration) SetTransitiveMembershipUserField(v string)`

SetTransitiveMembershipUserField sets TransitiveMembershipUserField field to given value.


### GetTransitiveDirectoryMembershipEnabled

`func (o *AzureServerConfiguration) GetTransitiveDirectoryMembershipEnabled() bool`

GetTransitiveDirectoryMembershipEnabled returns the TransitiveDirectoryMembershipEnabled field if non-nil, zero value otherwise.

### GetTransitiveDirectoryMembershipEnabledOk

`func (o *AzureServerConfiguration) GetTransitiveDirectoryMembershipEnabledOk() (*bool, bool)`

GetTransitiveDirectoryMembershipEnabledOk returns a tuple with the TransitiveDirectoryMembershipEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveDirectoryMembershipEnabled

`func (o *AzureServerConfiguration) SetTransitiveDirectoryMembershipEnabled(v bool)`

SetTransitiveDirectoryMembershipEnabled sets TransitiveDirectoryMembershipEnabled field to given value.


### GetMembershipCalculationOptimizationEnabled

`func (o *AzureServerConfiguration) GetMembershipCalculationOptimizationEnabled() bool`

GetMembershipCalculationOptimizationEnabled returns the MembershipCalculationOptimizationEnabled field if non-nil, zero value otherwise.

### GetMembershipCalculationOptimizationEnabledOk

`func (o *AzureServerConfiguration) GetMembershipCalculationOptimizationEnabledOk() (*bool, bool)`

GetMembershipCalculationOptimizationEnabledOk returns a tuple with the MembershipCalculationOptimizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipCalculationOptimizationEnabled

`func (o *AzureServerConfiguration) SetMembershipCalculationOptimizationEnabled(v bool)`

SetMembershipCalculationOptimizationEnabled sets MembershipCalculationOptimizationEnabled field to given value.

### HasMembershipCalculationOptimizationEnabled

`func (o *AzureServerConfiguration) HasMembershipCalculationOptimizationEnabled() bool`

HasMembershipCalculationOptimizationEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


