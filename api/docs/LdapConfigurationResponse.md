# LdapConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudIdPCommon** | [**CloudIdPCommon**](CloudIdPCommon.md) |  | 
**Server** | [**CloudLdapServerResponse**](CloudLdapServerResponse.md) |  | 
**Mappings** | Pointer to [**CloudLdapMappingsResponse**](CloudLdapMappingsResponse.md) |  | [optional] 

## Methods

### NewLdapConfigurationResponse

`func NewLdapConfigurationResponse(cloudIdPCommon CloudIdPCommon, server CloudLdapServerResponse, ) *LdapConfigurationResponse`

NewLdapConfigurationResponse instantiates a new LdapConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapConfigurationResponseWithDefaults

`func NewLdapConfigurationResponseWithDefaults() *LdapConfigurationResponse`

NewLdapConfigurationResponseWithDefaults instantiates a new LdapConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudIdPCommon

`func (o *LdapConfigurationResponse) GetCloudIdPCommon() CloudIdPCommon`

GetCloudIdPCommon returns the CloudIdPCommon field if non-nil, zero value otherwise.

### GetCloudIdPCommonOk

`func (o *LdapConfigurationResponse) GetCloudIdPCommonOk() (*CloudIdPCommon, bool)`

GetCloudIdPCommonOk returns a tuple with the CloudIdPCommon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudIdPCommon

`func (o *LdapConfigurationResponse) SetCloudIdPCommon(v CloudIdPCommon)`

SetCloudIdPCommon sets CloudIdPCommon field to given value.


### GetServer

`func (o *LdapConfigurationResponse) GetServer() CloudLdapServerResponse`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *LdapConfigurationResponse) GetServerOk() (*CloudLdapServerResponse, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *LdapConfigurationResponse) SetServer(v CloudLdapServerResponse)`

SetServer sets Server field to given value.


### GetMappings

`func (o *LdapConfigurationResponse) GetMappings() CloudLdapMappingsResponse`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *LdapConfigurationResponse) GetMappingsOk() (*CloudLdapMappingsResponse, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *LdapConfigurationResponse) SetMappings(v CloudLdapMappingsResponse)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *LdapConfigurationResponse) HasMappings() bool`

HasMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


