# LdapConfigurationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudIdPCommon** | [**CloudIdPCommon**](CloudIdPCommon.md) |  | 
**Server** | [**CloudLdapServerUpdate**](CloudLdapServerUpdate.md) |  | 
**Mappings** | Pointer to [**CloudLdapMappingsRequest**](CloudLdapMappingsRequest.md) |  | [optional] 

## Methods

### NewLdapConfigurationUpdate

`func NewLdapConfigurationUpdate(cloudIdPCommon CloudIdPCommon, server CloudLdapServerUpdate, ) *LdapConfigurationUpdate`

NewLdapConfigurationUpdate instantiates a new LdapConfigurationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapConfigurationUpdateWithDefaults

`func NewLdapConfigurationUpdateWithDefaults() *LdapConfigurationUpdate`

NewLdapConfigurationUpdateWithDefaults instantiates a new LdapConfigurationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudIdPCommon

`func (o *LdapConfigurationUpdate) GetCloudIdPCommon() CloudIdPCommon`

GetCloudIdPCommon returns the CloudIdPCommon field if non-nil, zero value otherwise.

### GetCloudIdPCommonOk

`func (o *LdapConfigurationUpdate) GetCloudIdPCommonOk() (*CloudIdPCommon, bool)`

GetCloudIdPCommonOk returns a tuple with the CloudIdPCommon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudIdPCommon

`func (o *LdapConfigurationUpdate) SetCloudIdPCommon(v CloudIdPCommon)`

SetCloudIdPCommon sets CloudIdPCommon field to given value.


### GetServer

`func (o *LdapConfigurationUpdate) GetServer() CloudLdapServerUpdate`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *LdapConfigurationUpdate) GetServerOk() (*CloudLdapServerUpdate, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *LdapConfigurationUpdate) SetServer(v CloudLdapServerUpdate)`

SetServer sets Server field to given value.


### GetMappings

`func (o *LdapConfigurationUpdate) GetMappings() CloudLdapMappingsRequest`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *LdapConfigurationUpdate) GetMappingsOk() (*CloudLdapMappingsRequest, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *LdapConfigurationUpdate) SetMappings(v CloudLdapMappingsRequest)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *LdapConfigurationUpdate) HasMappings() bool`

HasMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


