# SamlSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenExpirationDisabled** | Pointer to **bool** |  | [optional] [default to false]
**UserAttributeEnabled** | Pointer to **bool** |  | [optional] [default to false]
**UserAttributeName** | Pointer to **string** |  | [optional] [default to " "]
**UserMapping** | Pointer to **string** |  | [optional] 
**GroupAttributeName** | Pointer to **string** |  | [optional] [default to "http://schemas.xmlsoap.org/claims/Group"]
**GroupRdnKey** | Pointer to **string** |  | [optional] [default to " "]
**IdpProviderType** | Pointer to **string** |  | [optional] 
**IdpUrl** | Pointer to **string** |  | [optional] 
**EntityId** | Pointer to **string** |  | [optional] 
**MetadataFileName** | Pointer to **string** |  | [optional] 
**OtherProviderTypeName** | Pointer to **string** |  | [optional] [default to " "]
**FederationMetadataFile** | Pointer to **string** |  | [optional] 
**MetadataSource** | Pointer to **string** |  | [optional] 
**SessionTimeout** | Pointer to **int64** |  | [optional] [default to 480]

## Methods

### NewSamlSettings

`func NewSamlSettings() *SamlSettings`

NewSamlSettings instantiates a new SamlSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlSettingsWithDefaults

`func NewSamlSettingsWithDefaults() *SamlSettings`

NewSamlSettingsWithDefaults instantiates a new SamlSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenExpirationDisabled

`func (o *SamlSettings) GetTokenExpirationDisabled() bool`

GetTokenExpirationDisabled returns the TokenExpirationDisabled field if non-nil, zero value otherwise.

### GetTokenExpirationDisabledOk

`func (o *SamlSettings) GetTokenExpirationDisabledOk() (*bool, bool)`

GetTokenExpirationDisabledOk returns a tuple with the TokenExpirationDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpirationDisabled

`func (o *SamlSettings) SetTokenExpirationDisabled(v bool)`

SetTokenExpirationDisabled sets TokenExpirationDisabled field to given value.

### HasTokenExpirationDisabled

`func (o *SamlSettings) HasTokenExpirationDisabled() bool`

HasTokenExpirationDisabled returns a boolean if a field has been set.

### GetUserAttributeEnabled

`func (o *SamlSettings) GetUserAttributeEnabled() bool`

GetUserAttributeEnabled returns the UserAttributeEnabled field if non-nil, zero value otherwise.

### GetUserAttributeEnabledOk

`func (o *SamlSettings) GetUserAttributeEnabledOk() (*bool, bool)`

GetUserAttributeEnabledOk returns a tuple with the UserAttributeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeEnabled

`func (o *SamlSettings) SetUserAttributeEnabled(v bool)`

SetUserAttributeEnabled sets UserAttributeEnabled field to given value.

### HasUserAttributeEnabled

`func (o *SamlSettings) HasUserAttributeEnabled() bool`

HasUserAttributeEnabled returns a boolean if a field has been set.

### GetUserAttributeName

`func (o *SamlSettings) GetUserAttributeName() string`

GetUserAttributeName returns the UserAttributeName field if non-nil, zero value otherwise.

### GetUserAttributeNameOk

`func (o *SamlSettings) GetUserAttributeNameOk() (*string, bool)`

GetUserAttributeNameOk returns a tuple with the UserAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeName

`func (o *SamlSettings) SetUserAttributeName(v string)`

SetUserAttributeName sets UserAttributeName field to given value.

### HasUserAttributeName

`func (o *SamlSettings) HasUserAttributeName() bool`

HasUserAttributeName returns a boolean if a field has been set.

### GetUserMapping

`func (o *SamlSettings) GetUserMapping() string`

GetUserMapping returns the UserMapping field if non-nil, zero value otherwise.

### GetUserMappingOk

`func (o *SamlSettings) GetUserMappingOk() (*string, bool)`

GetUserMappingOk returns a tuple with the UserMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMapping

`func (o *SamlSettings) SetUserMapping(v string)`

SetUserMapping sets UserMapping field to given value.

### HasUserMapping

`func (o *SamlSettings) HasUserMapping() bool`

HasUserMapping returns a boolean if a field has been set.

### GetGroupAttributeName

`func (o *SamlSettings) GetGroupAttributeName() string`

GetGroupAttributeName returns the GroupAttributeName field if non-nil, zero value otherwise.

### GetGroupAttributeNameOk

`func (o *SamlSettings) GetGroupAttributeNameOk() (*string, bool)`

GetGroupAttributeNameOk returns a tuple with the GroupAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAttributeName

`func (o *SamlSettings) SetGroupAttributeName(v string)`

SetGroupAttributeName sets GroupAttributeName field to given value.

### HasGroupAttributeName

`func (o *SamlSettings) HasGroupAttributeName() bool`

HasGroupAttributeName returns a boolean if a field has been set.

### GetGroupRdnKey

`func (o *SamlSettings) GetGroupRdnKey() string`

GetGroupRdnKey returns the GroupRdnKey field if non-nil, zero value otherwise.

### GetGroupRdnKeyOk

`func (o *SamlSettings) GetGroupRdnKeyOk() (*string, bool)`

GetGroupRdnKeyOk returns a tuple with the GroupRdnKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupRdnKey

`func (o *SamlSettings) SetGroupRdnKey(v string)`

SetGroupRdnKey sets GroupRdnKey field to given value.

### HasGroupRdnKey

`func (o *SamlSettings) HasGroupRdnKey() bool`

HasGroupRdnKey returns a boolean if a field has been set.

### GetIdpProviderType

`func (o *SamlSettings) GetIdpProviderType() string`

GetIdpProviderType returns the IdpProviderType field if non-nil, zero value otherwise.

### GetIdpProviderTypeOk

`func (o *SamlSettings) GetIdpProviderTypeOk() (*string, bool)`

GetIdpProviderTypeOk returns a tuple with the IdpProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpProviderType

`func (o *SamlSettings) SetIdpProviderType(v string)`

SetIdpProviderType sets IdpProviderType field to given value.

### HasIdpProviderType

`func (o *SamlSettings) HasIdpProviderType() bool`

HasIdpProviderType returns a boolean if a field has been set.

### GetIdpUrl

`func (o *SamlSettings) GetIdpUrl() string`

GetIdpUrl returns the IdpUrl field if non-nil, zero value otherwise.

### GetIdpUrlOk

`func (o *SamlSettings) GetIdpUrlOk() (*string, bool)`

GetIdpUrlOk returns a tuple with the IdpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpUrl

`func (o *SamlSettings) SetIdpUrl(v string)`

SetIdpUrl sets IdpUrl field to given value.

### HasIdpUrl

`func (o *SamlSettings) HasIdpUrl() bool`

HasIdpUrl returns a boolean if a field has been set.

### GetEntityId

`func (o *SamlSettings) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SamlSettings) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SamlSettings) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *SamlSettings) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetMetadataFileName

`func (o *SamlSettings) GetMetadataFileName() string`

GetMetadataFileName returns the MetadataFileName field if non-nil, zero value otherwise.

### GetMetadataFileNameOk

`func (o *SamlSettings) GetMetadataFileNameOk() (*string, bool)`

GetMetadataFileNameOk returns a tuple with the MetadataFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFileName

`func (o *SamlSettings) SetMetadataFileName(v string)`

SetMetadataFileName sets MetadataFileName field to given value.

### HasMetadataFileName

`func (o *SamlSettings) HasMetadataFileName() bool`

HasMetadataFileName returns a boolean if a field has been set.

### GetOtherProviderTypeName

`func (o *SamlSettings) GetOtherProviderTypeName() string`

GetOtherProviderTypeName returns the OtherProviderTypeName field if non-nil, zero value otherwise.

### GetOtherProviderTypeNameOk

`func (o *SamlSettings) GetOtherProviderTypeNameOk() (*string, bool)`

GetOtherProviderTypeNameOk returns a tuple with the OtherProviderTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherProviderTypeName

`func (o *SamlSettings) SetOtherProviderTypeName(v string)`

SetOtherProviderTypeName sets OtherProviderTypeName field to given value.

### HasOtherProviderTypeName

`func (o *SamlSettings) HasOtherProviderTypeName() bool`

HasOtherProviderTypeName returns a boolean if a field has been set.

### GetFederationMetadataFile

`func (o *SamlSettings) GetFederationMetadataFile() string`

GetFederationMetadataFile returns the FederationMetadataFile field if non-nil, zero value otherwise.

### GetFederationMetadataFileOk

`func (o *SamlSettings) GetFederationMetadataFileOk() (*string, bool)`

GetFederationMetadataFileOk returns a tuple with the FederationMetadataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationMetadataFile

`func (o *SamlSettings) SetFederationMetadataFile(v string)`

SetFederationMetadataFile sets FederationMetadataFile field to given value.

### HasFederationMetadataFile

`func (o *SamlSettings) HasFederationMetadataFile() bool`

HasFederationMetadataFile returns a boolean if a field has been set.

### GetMetadataSource

`func (o *SamlSettings) GetMetadataSource() string`

GetMetadataSource returns the MetadataSource field if non-nil, zero value otherwise.

### GetMetadataSourceOk

`func (o *SamlSettings) GetMetadataSourceOk() (*string, bool)`

GetMetadataSourceOk returns a tuple with the MetadataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSource

`func (o *SamlSettings) SetMetadataSource(v string)`

SetMetadataSource sets MetadataSource field to given value.

### HasMetadataSource

`func (o *SamlSettings) HasMetadataSource() bool`

HasMetadataSource returns a boolean if a field has been set.

### GetSessionTimeout

`func (o *SamlSettings) GetSessionTimeout() int64`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *SamlSettings) GetSessionTimeoutOk() (*int64, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *SamlSettings) SetSessionTimeout(v int64)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *SamlSettings) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


