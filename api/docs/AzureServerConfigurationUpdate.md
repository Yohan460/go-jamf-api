# AzureServerConfigurationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Enabled** | **bool** |  | 
**Mappings** | [**AzureMappings**](AzureMappings.md) |  | 
**SearchTimeout** | **int32** |  | 
**TransitiveMembershipEnabled** | **bool** | Use this field to enable transitive membership lookup with Single Sign On | 
**TransitiveMembershipUserField** | **string** | Use this field to set user field mapping for transitive membership lookup with Single Sign On | 
**TransitiveDirectoryMembershipEnabled** | **bool** | Use this field to enable transitive membership lookup. This setting would not apply to Single Sign On | 

## Methods

### NewAzureServerConfigurationUpdate

`func NewAzureServerConfigurationUpdate(id string, enabled bool, mappings AzureMappings, searchTimeout int32, transitiveMembershipEnabled bool, transitiveMembershipUserField string, transitiveDirectoryMembershipEnabled bool, ) *AzureServerConfigurationUpdate`

NewAzureServerConfigurationUpdate instantiates a new AzureServerConfigurationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureServerConfigurationUpdateWithDefaults

`func NewAzureServerConfigurationUpdateWithDefaults() *AzureServerConfigurationUpdate`

NewAzureServerConfigurationUpdateWithDefaults instantiates a new AzureServerConfigurationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AzureServerConfigurationUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureServerConfigurationUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureServerConfigurationUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetEnabled

`func (o *AzureServerConfigurationUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AzureServerConfigurationUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AzureServerConfigurationUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMappings

`func (o *AzureServerConfigurationUpdate) GetMappings() AzureMappings`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *AzureServerConfigurationUpdate) GetMappingsOk() (*AzureMappings, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *AzureServerConfigurationUpdate) SetMappings(v AzureMappings)`

SetMappings sets Mappings field to given value.


### GetSearchTimeout

`func (o *AzureServerConfigurationUpdate) GetSearchTimeout() int32`

GetSearchTimeout returns the SearchTimeout field if non-nil, zero value otherwise.

### GetSearchTimeoutOk

`func (o *AzureServerConfigurationUpdate) GetSearchTimeoutOk() (*int32, bool)`

GetSearchTimeoutOk returns a tuple with the SearchTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTimeout

`func (o *AzureServerConfigurationUpdate) SetSearchTimeout(v int32)`

SetSearchTimeout sets SearchTimeout field to given value.


### GetTransitiveMembershipEnabled

`func (o *AzureServerConfigurationUpdate) GetTransitiveMembershipEnabled() bool`

GetTransitiveMembershipEnabled returns the TransitiveMembershipEnabled field if non-nil, zero value otherwise.

### GetTransitiveMembershipEnabledOk

`func (o *AzureServerConfigurationUpdate) GetTransitiveMembershipEnabledOk() (*bool, bool)`

GetTransitiveMembershipEnabledOk returns a tuple with the TransitiveMembershipEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveMembershipEnabled

`func (o *AzureServerConfigurationUpdate) SetTransitiveMembershipEnabled(v bool)`

SetTransitiveMembershipEnabled sets TransitiveMembershipEnabled field to given value.


### GetTransitiveMembershipUserField

`func (o *AzureServerConfigurationUpdate) GetTransitiveMembershipUserField() string`

GetTransitiveMembershipUserField returns the TransitiveMembershipUserField field if non-nil, zero value otherwise.

### GetTransitiveMembershipUserFieldOk

`func (o *AzureServerConfigurationUpdate) GetTransitiveMembershipUserFieldOk() (*string, bool)`

GetTransitiveMembershipUserFieldOk returns a tuple with the TransitiveMembershipUserField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveMembershipUserField

`func (o *AzureServerConfigurationUpdate) SetTransitiveMembershipUserField(v string)`

SetTransitiveMembershipUserField sets TransitiveMembershipUserField field to given value.


### GetTransitiveDirectoryMembershipEnabled

`func (o *AzureServerConfigurationUpdate) GetTransitiveDirectoryMembershipEnabled() bool`

GetTransitiveDirectoryMembershipEnabled returns the TransitiveDirectoryMembershipEnabled field if non-nil, zero value otherwise.

### GetTransitiveDirectoryMembershipEnabledOk

`func (o *AzureServerConfigurationUpdate) GetTransitiveDirectoryMembershipEnabledOk() (*bool, bool)`

GetTransitiveDirectoryMembershipEnabledOk returns a tuple with the TransitiveDirectoryMembershipEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveDirectoryMembershipEnabled

`func (o *AzureServerConfigurationUpdate) SetTransitiveDirectoryMembershipEnabled(v bool)`

SetTransitiveDirectoryMembershipEnabled sets TransitiveDirectoryMembershipEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


