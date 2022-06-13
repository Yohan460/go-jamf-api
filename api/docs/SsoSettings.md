# SsoSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsoForEnrollmentEnabled** | **bool** |  | [default to false]
**SsoBypassAllowed** | **bool** |  | [default to false]
**SsoEnabled** | **bool** |  | [default to false]
**SsoForMacOsSelfServiceEnabled** | **bool** |  | [default to false]
**TokenExpirationDisabled** | **bool** |  | [default to false]
**UserAttributeEnabled** | **bool** |  | [default to false]
**UserAttributeName** | Pointer to **string** |  | [optional] [default to " "]
**UserMapping** | **string** |  | 
**GroupEnrollmentAccessEnabled** | **bool** |  | [default to false]
**GroupAttributeName** | **string** |  | [default to "http://schemas.xmlsoap.org/claims/Group"]
**GroupRdnKey** | **string** |  | [default to " "]
**GroupEnrollmentAccessName** | Pointer to **string** |  | [optional] [default to " "]
**IdpProviderType** | **string** |  | 
**IdpUrl** | Pointer to **string** |  | [optional] 
**EntityId** | **string** |  | 
**MetadataFileName** | Pointer to **string** |  | [optional] 
**OtherProviderTypeName** | Pointer to **string** |  | [optional] [default to " "]
**FederationMetadataFile** | Pointer to **string** |  | [optional] 
**MetadataSource** | **string** |  | 
**SessionTimeout** | Pointer to **int32** |  | [optional] [default to 480]

## Methods

### NewSsoSettings

`func NewSsoSettings(ssoForEnrollmentEnabled bool, ssoBypassAllowed bool, ssoEnabled bool, ssoForMacOsSelfServiceEnabled bool, tokenExpirationDisabled bool, userAttributeEnabled bool, userMapping string, groupEnrollmentAccessEnabled bool, groupAttributeName string, groupRdnKey string, idpProviderType string, entityId string, metadataSource string, ) *SsoSettings`

NewSsoSettings instantiates a new SsoSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoSettingsWithDefaults

`func NewSsoSettingsWithDefaults() *SsoSettings`

NewSsoSettingsWithDefaults instantiates a new SsoSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsoForEnrollmentEnabled

`func (o *SsoSettings) GetSsoForEnrollmentEnabled() bool`

GetSsoForEnrollmentEnabled returns the SsoForEnrollmentEnabled field if non-nil, zero value otherwise.

### GetSsoForEnrollmentEnabledOk

`func (o *SsoSettings) GetSsoForEnrollmentEnabledOk() (*bool, bool)`

GetSsoForEnrollmentEnabledOk returns a tuple with the SsoForEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoForEnrollmentEnabled

`func (o *SsoSettings) SetSsoForEnrollmentEnabled(v bool)`

SetSsoForEnrollmentEnabled sets SsoForEnrollmentEnabled field to given value.


### GetSsoBypassAllowed

`func (o *SsoSettings) GetSsoBypassAllowed() bool`

GetSsoBypassAllowed returns the SsoBypassAllowed field if non-nil, zero value otherwise.

### GetSsoBypassAllowedOk

`func (o *SsoSettings) GetSsoBypassAllowedOk() (*bool, bool)`

GetSsoBypassAllowedOk returns a tuple with the SsoBypassAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoBypassAllowed

`func (o *SsoSettings) SetSsoBypassAllowed(v bool)`

SetSsoBypassAllowed sets SsoBypassAllowed field to given value.


### GetSsoEnabled

`func (o *SsoSettings) GetSsoEnabled() bool`

GetSsoEnabled returns the SsoEnabled field if non-nil, zero value otherwise.

### GetSsoEnabledOk

`func (o *SsoSettings) GetSsoEnabledOk() (*bool, bool)`

GetSsoEnabledOk returns a tuple with the SsoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoEnabled

`func (o *SsoSettings) SetSsoEnabled(v bool)`

SetSsoEnabled sets SsoEnabled field to given value.


### GetSsoForMacOsSelfServiceEnabled

`func (o *SsoSettings) GetSsoForMacOsSelfServiceEnabled() bool`

GetSsoForMacOsSelfServiceEnabled returns the SsoForMacOsSelfServiceEnabled field if non-nil, zero value otherwise.

### GetSsoForMacOsSelfServiceEnabledOk

`func (o *SsoSettings) GetSsoForMacOsSelfServiceEnabledOk() (*bool, bool)`

GetSsoForMacOsSelfServiceEnabledOk returns a tuple with the SsoForMacOsSelfServiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoForMacOsSelfServiceEnabled

`func (o *SsoSettings) SetSsoForMacOsSelfServiceEnabled(v bool)`

SetSsoForMacOsSelfServiceEnabled sets SsoForMacOsSelfServiceEnabled field to given value.


### GetTokenExpirationDisabled

`func (o *SsoSettings) GetTokenExpirationDisabled() bool`

GetTokenExpirationDisabled returns the TokenExpirationDisabled field if non-nil, zero value otherwise.

### GetTokenExpirationDisabledOk

`func (o *SsoSettings) GetTokenExpirationDisabledOk() (*bool, bool)`

GetTokenExpirationDisabledOk returns a tuple with the TokenExpirationDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpirationDisabled

`func (o *SsoSettings) SetTokenExpirationDisabled(v bool)`

SetTokenExpirationDisabled sets TokenExpirationDisabled field to given value.


### GetUserAttributeEnabled

`func (o *SsoSettings) GetUserAttributeEnabled() bool`

GetUserAttributeEnabled returns the UserAttributeEnabled field if non-nil, zero value otherwise.

### GetUserAttributeEnabledOk

`func (o *SsoSettings) GetUserAttributeEnabledOk() (*bool, bool)`

GetUserAttributeEnabledOk returns a tuple with the UserAttributeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeEnabled

`func (o *SsoSettings) SetUserAttributeEnabled(v bool)`

SetUserAttributeEnabled sets UserAttributeEnabled field to given value.


### GetUserAttributeName

`func (o *SsoSettings) GetUserAttributeName() string`

GetUserAttributeName returns the UserAttributeName field if non-nil, zero value otherwise.

### GetUserAttributeNameOk

`func (o *SsoSettings) GetUserAttributeNameOk() (*string, bool)`

GetUserAttributeNameOk returns a tuple with the UserAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeName

`func (o *SsoSettings) SetUserAttributeName(v string)`

SetUserAttributeName sets UserAttributeName field to given value.

### HasUserAttributeName

`func (o *SsoSettings) HasUserAttributeName() bool`

HasUserAttributeName returns a boolean if a field has been set.

### GetUserMapping

`func (o *SsoSettings) GetUserMapping() string`

GetUserMapping returns the UserMapping field if non-nil, zero value otherwise.

### GetUserMappingOk

`func (o *SsoSettings) GetUserMappingOk() (*string, bool)`

GetUserMappingOk returns a tuple with the UserMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMapping

`func (o *SsoSettings) SetUserMapping(v string)`

SetUserMapping sets UserMapping field to given value.


### GetGroupEnrollmentAccessEnabled

`func (o *SsoSettings) GetGroupEnrollmentAccessEnabled() bool`

GetGroupEnrollmentAccessEnabled returns the GroupEnrollmentAccessEnabled field if non-nil, zero value otherwise.

### GetGroupEnrollmentAccessEnabledOk

`func (o *SsoSettings) GetGroupEnrollmentAccessEnabledOk() (*bool, bool)`

GetGroupEnrollmentAccessEnabledOk returns a tuple with the GroupEnrollmentAccessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupEnrollmentAccessEnabled

`func (o *SsoSettings) SetGroupEnrollmentAccessEnabled(v bool)`

SetGroupEnrollmentAccessEnabled sets GroupEnrollmentAccessEnabled field to given value.


### GetGroupAttributeName

`func (o *SsoSettings) GetGroupAttributeName() string`

GetGroupAttributeName returns the GroupAttributeName field if non-nil, zero value otherwise.

### GetGroupAttributeNameOk

`func (o *SsoSettings) GetGroupAttributeNameOk() (*string, bool)`

GetGroupAttributeNameOk returns a tuple with the GroupAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAttributeName

`func (o *SsoSettings) SetGroupAttributeName(v string)`

SetGroupAttributeName sets GroupAttributeName field to given value.


### GetGroupRdnKey

`func (o *SsoSettings) GetGroupRdnKey() string`

GetGroupRdnKey returns the GroupRdnKey field if non-nil, zero value otherwise.

### GetGroupRdnKeyOk

`func (o *SsoSettings) GetGroupRdnKeyOk() (*string, bool)`

GetGroupRdnKeyOk returns a tuple with the GroupRdnKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupRdnKey

`func (o *SsoSettings) SetGroupRdnKey(v string)`

SetGroupRdnKey sets GroupRdnKey field to given value.


### GetGroupEnrollmentAccessName

`func (o *SsoSettings) GetGroupEnrollmentAccessName() string`

GetGroupEnrollmentAccessName returns the GroupEnrollmentAccessName field if non-nil, zero value otherwise.

### GetGroupEnrollmentAccessNameOk

`func (o *SsoSettings) GetGroupEnrollmentAccessNameOk() (*string, bool)`

GetGroupEnrollmentAccessNameOk returns a tuple with the GroupEnrollmentAccessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupEnrollmentAccessName

`func (o *SsoSettings) SetGroupEnrollmentAccessName(v string)`

SetGroupEnrollmentAccessName sets GroupEnrollmentAccessName field to given value.

### HasGroupEnrollmentAccessName

`func (o *SsoSettings) HasGroupEnrollmentAccessName() bool`

HasGroupEnrollmentAccessName returns a boolean if a field has been set.

### GetIdpProviderType

`func (o *SsoSettings) GetIdpProviderType() string`

GetIdpProviderType returns the IdpProviderType field if non-nil, zero value otherwise.

### GetIdpProviderTypeOk

`func (o *SsoSettings) GetIdpProviderTypeOk() (*string, bool)`

GetIdpProviderTypeOk returns a tuple with the IdpProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpProviderType

`func (o *SsoSettings) SetIdpProviderType(v string)`

SetIdpProviderType sets IdpProviderType field to given value.


### GetIdpUrl

`func (o *SsoSettings) GetIdpUrl() string`

GetIdpUrl returns the IdpUrl field if non-nil, zero value otherwise.

### GetIdpUrlOk

`func (o *SsoSettings) GetIdpUrlOk() (*string, bool)`

GetIdpUrlOk returns a tuple with the IdpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpUrl

`func (o *SsoSettings) SetIdpUrl(v string)`

SetIdpUrl sets IdpUrl field to given value.

### HasIdpUrl

`func (o *SsoSettings) HasIdpUrl() bool`

HasIdpUrl returns a boolean if a field has been set.

### GetEntityId

`func (o *SsoSettings) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SsoSettings) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SsoSettings) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetMetadataFileName

`func (o *SsoSettings) GetMetadataFileName() string`

GetMetadataFileName returns the MetadataFileName field if non-nil, zero value otherwise.

### GetMetadataFileNameOk

`func (o *SsoSettings) GetMetadataFileNameOk() (*string, bool)`

GetMetadataFileNameOk returns a tuple with the MetadataFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFileName

`func (o *SsoSettings) SetMetadataFileName(v string)`

SetMetadataFileName sets MetadataFileName field to given value.

### HasMetadataFileName

`func (o *SsoSettings) HasMetadataFileName() bool`

HasMetadataFileName returns a boolean if a field has been set.

### GetOtherProviderTypeName

`func (o *SsoSettings) GetOtherProviderTypeName() string`

GetOtherProviderTypeName returns the OtherProviderTypeName field if non-nil, zero value otherwise.

### GetOtherProviderTypeNameOk

`func (o *SsoSettings) GetOtherProviderTypeNameOk() (*string, bool)`

GetOtherProviderTypeNameOk returns a tuple with the OtherProviderTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherProviderTypeName

`func (o *SsoSettings) SetOtherProviderTypeName(v string)`

SetOtherProviderTypeName sets OtherProviderTypeName field to given value.

### HasOtherProviderTypeName

`func (o *SsoSettings) HasOtherProviderTypeName() bool`

HasOtherProviderTypeName returns a boolean if a field has been set.

### GetFederationMetadataFile

`func (o *SsoSettings) GetFederationMetadataFile() string`

GetFederationMetadataFile returns the FederationMetadataFile field if non-nil, zero value otherwise.

### GetFederationMetadataFileOk

`func (o *SsoSettings) GetFederationMetadataFileOk() (*string, bool)`

GetFederationMetadataFileOk returns a tuple with the FederationMetadataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationMetadataFile

`func (o *SsoSettings) SetFederationMetadataFile(v string)`

SetFederationMetadataFile sets FederationMetadataFile field to given value.

### HasFederationMetadataFile

`func (o *SsoSettings) HasFederationMetadataFile() bool`

HasFederationMetadataFile returns a boolean if a field has been set.

### GetMetadataSource

`func (o *SsoSettings) GetMetadataSource() string`

GetMetadataSource returns the MetadataSource field if non-nil, zero value otherwise.

### GetMetadataSourceOk

`func (o *SsoSettings) GetMetadataSourceOk() (*string, bool)`

GetMetadataSourceOk returns a tuple with the MetadataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSource

`func (o *SsoSettings) SetMetadataSource(v string)`

SetMetadataSource sets MetadataSource field to given value.


### GetSessionTimeout

`func (o *SsoSettings) GetSessionTimeout() int32`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *SsoSettings) GetSessionTimeoutOk() (*int32, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *SsoSettings) SetSessionTimeout(v int32)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *SsoSettings) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


