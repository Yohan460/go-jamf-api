# DeprecatedConfigurationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | [**DeprecatedServerUpdate**](DeprecatedServerUpdate.md) |  | 
**Mappings** | Pointer to [**CloudLdapMappingsRequest**](CloudLdapMappingsRequest.md) |  | [optional] 

## Methods

### NewDeprecatedConfigurationUpdate

`func NewDeprecatedConfigurationUpdate(server DeprecatedServerUpdate, ) *DeprecatedConfigurationUpdate`

NewDeprecatedConfigurationUpdate instantiates a new DeprecatedConfigurationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeprecatedConfigurationUpdateWithDefaults

`func NewDeprecatedConfigurationUpdateWithDefaults() *DeprecatedConfigurationUpdate`

NewDeprecatedConfigurationUpdateWithDefaults instantiates a new DeprecatedConfigurationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *DeprecatedConfigurationUpdate) GetServer() DeprecatedServerUpdate`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *DeprecatedConfigurationUpdate) GetServerOk() (*DeprecatedServerUpdate, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *DeprecatedConfigurationUpdate) SetServer(v DeprecatedServerUpdate)`

SetServer sets Server field to given value.


### GetMappings

`func (o *DeprecatedConfigurationUpdate) GetMappings() CloudLdapMappingsRequest`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *DeprecatedConfigurationUpdate) GetMappingsOk() (*CloudLdapMappingsRequest, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *DeprecatedConfigurationUpdate) SetMappings(v CloudLdapMappingsRequest)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *DeprecatedConfigurationUpdate) HasMappings() bool`

HasMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


