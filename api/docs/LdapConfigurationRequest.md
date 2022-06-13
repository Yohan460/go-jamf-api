# LdapConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudIdPCommon** | [**CloudIdPCommonRequest**](CloudIdPCommonRequest.md) |  | 
**Server** | [**CloudLdapServerRequest**](CloudLdapServerRequest.md) |  | 
**Mappings** | Pointer to [**CloudLdapMappingsRequest**](CloudLdapMappingsRequest.md) |  | [optional] 

## Methods

### NewLdapConfigurationRequest

`func NewLdapConfigurationRequest(cloudIdPCommon CloudIdPCommonRequest, server CloudLdapServerRequest, ) *LdapConfigurationRequest`

NewLdapConfigurationRequest instantiates a new LdapConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapConfigurationRequestWithDefaults

`func NewLdapConfigurationRequestWithDefaults() *LdapConfigurationRequest`

NewLdapConfigurationRequestWithDefaults instantiates a new LdapConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudIdPCommon

`func (o *LdapConfigurationRequest) GetCloudIdPCommon() CloudIdPCommonRequest`

GetCloudIdPCommon returns the CloudIdPCommon field if non-nil, zero value otherwise.

### GetCloudIdPCommonOk

`func (o *LdapConfigurationRequest) GetCloudIdPCommonOk() (*CloudIdPCommonRequest, bool)`

GetCloudIdPCommonOk returns a tuple with the CloudIdPCommon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudIdPCommon

`func (o *LdapConfigurationRequest) SetCloudIdPCommon(v CloudIdPCommonRequest)`

SetCloudIdPCommon sets CloudIdPCommon field to given value.


### GetServer

`func (o *LdapConfigurationRequest) GetServer() CloudLdapServerRequest`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *LdapConfigurationRequest) GetServerOk() (*CloudLdapServerRequest, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *LdapConfigurationRequest) SetServer(v CloudLdapServerRequest)`

SetServer sets Server field to given value.


### GetMappings

`func (o *LdapConfigurationRequest) GetMappings() CloudLdapMappingsRequest`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *LdapConfigurationRequest) GetMappingsOk() (*CloudLdapMappingsRequest, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *LdapConfigurationRequest) SetMappings(v CloudLdapMappingsRequest)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *LdapConfigurationRequest) HasMappings() bool`

HasMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


