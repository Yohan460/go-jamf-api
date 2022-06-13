# DeprecatedConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Server** | [**DeprecatedServerResponse**](DeprecatedServerResponse.md) |  | 
**Mappings** | Pointer to [**CloudLdapMappingsResponse**](CloudLdapMappingsResponse.md) |  | [optional] 

## Methods

### NewDeprecatedConfigurationResponse

`func NewDeprecatedConfigurationResponse(id string, server DeprecatedServerResponse, ) *DeprecatedConfigurationResponse`

NewDeprecatedConfigurationResponse instantiates a new DeprecatedConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeprecatedConfigurationResponseWithDefaults

`func NewDeprecatedConfigurationResponseWithDefaults() *DeprecatedConfigurationResponse`

NewDeprecatedConfigurationResponseWithDefaults instantiates a new DeprecatedConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeprecatedConfigurationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeprecatedConfigurationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeprecatedConfigurationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetServer

`func (o *DeprecatedConfigurationResponse) GetServer() DeprecatedServerResponse`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *DeprecatedConfigurationResponse) GetServerOk() (*DeprecatedServerResponse, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *DeprecatedConfigurationResponse) SetServer(v DeprecatedServerResponse)`

SetServer sets Server field to given value.


### GetMappings

`func (o *DeprecatedConfigurationResponse) GetMappings() CloudLdapMappingsResponse`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *DeprecatedConfigurationResponse) GetMappingsOk() (*CloudLdapMappingsResponse, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *DeprecatedConfigurationResponse) SetMappings(v CloudLdapMappingsResponse)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *DeprecatedConfigurationResponse) HasMappings() bool`

HasMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


