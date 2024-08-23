# SsoSettingsV2

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
**EnrollmentSsoForAccountDrivenEnrollmentEnabled** | **bool** |  | [default to false]
**EnrollmentSsoConfig** | Pointer to [**EnrollmentSsoConfig**](EnrollmentSsoConfig.md) |  | [optional] 
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
**SessionTimeout** | Pointer to **int64** |  | [optional] [default to 480]

## Methods

### NewSsoSettingsV2

`func NewSsoSettingsV2(ssoForEnrollmentEnabled bool, ssoBypassAllowed bool, ssoEnabled bool, ssoForMacOsSelfServiceEnabled bool, tokenExpirationDisabled bool, userAttributeEnabled bool, userMapping string, enrollmentSsoForAccountDrivenEnrollmentEnabled bool, groupEnrollmentAccessEnabled bool, groupAttributeName string, groupRdnKey string, idpProviderType string, entityId string, metadataSource string, ) *SsoSettingsV2`

NewSsoSettingsV2 instantiates a new SsoSettingsV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoSettingsV2WithDefaults

`func NewSsoSettingsV2WithDefaults() *SsoSettingsV2`

NewSsoSettingsV2WithDefaults instantiates a new SsoSettingsV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsoForEnrollmentEnabled

`func (o *SsoSettingsV2) GetSsoForEnrollmentEnabled() bool`

GetSsoForEnrollmentEnabled returns the SsoForEnrollmentEnabled field if non-nil, zero value otherwise.

### GetSsoForEnrollmentEnabledOk

`func (o *SsoSettingsV2) GetSsoForEnrollmentEnabledOk() (*bool, bool)`

GetSsoForEnrollmentEnabledOk returns a tuple with the SsoForEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoForEnrollmentEnabled

`func (o *SsoSettingsV2) SetSsoForEnrollmentEnabled(v bool)`

SetSsoForEnrollmentEnabled sets SsoForEnrollmentEnabled field to given value.


### GetSsoBypassAllowed

`func (o *SsoSettingsV2) GetSsoBypassAllowed() bool`

GetSsoBypassAllowed returns the SsoBypassAllowed field if non-nil, zero value otherwise.

### GetSsoBypassAllowedOk

`func (o *SsoSettingsV2) GetSsoBypassAllowedOk() (*bool, bool)`

GetSsoBypassAllowedOk returns a tuple with the SsoBypassAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoBypassAllowed

`func (o *SsoSettingsV2) SetSsoBypassAllowed(v bool)`

SetSsoBypassAllowed sets SsoBypassAllowed field to given value.


### GetSsoEnabled

`func (o *SsoSettingsV2) GetSsoEnabled() bool`

GetSsoEnabled returns the SsoEnabled field if non-nil, zero value otherwise.

### GetSsoEnabledOk

`func (o *SsoSettingsV2) GetSsoEnabledOk() (*bool, bool)`

GetSsoEnabledOk returns a tuple with the SsoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoEnabled

`func (o *SsoSettingsV2) SetSsoEnabled(v bool)`

SetSsoEnabled sets SsoEnabled field to given value.


### GetSsoForMacOsSelfServiceEnabled

`func (o *SsoSettingsV2) GetSsoForMacOsSelfServiceEnabled() bool`

GetSsoForMacOsSelfServiceEnabled returns the SsoForMacOsSelfServiceEnabled field if non-nil, zero value otherwise.

### GetSsoForMacOsSelfServiceEnabledOk

`func (o *SsoSettingsV2) GetSsoForMacOsSelfServiceEnabledOk() (*bool, bool)`

GetSsoForMacOsSelfServiceEnabledOk returns a tuple with the SsoForMacOsSelfServiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoForMacOsSelfServiceEnabled

`func (o *SsoSettingsV2) SetSsoForMacOsSelfServiceEnabled(v bool)`

SetSsoForMacOsSelfServiceEnabled sets SsoForMacOsSelfServiceEnabled field to given value.


### GetTokenExpirationDisabled

`func (o *SsoSettingsV2) GetTokenExpirationDisabled() bool`

GetTokenExpirationDisabled returns the TokenExpirationDisabled field if non-nil, zero value otherwise.

### GetTokenExpirationDisabledOk

`func (o *SsoSettingsV2) GetTokenExpirationDisabledOk() (*bool, bool)`

GetTokenExpirationDisabledOk returns a tuple with the TokenExpirationDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpirationDisabled

`func (o *SsoSettingsV2) SetTokenExpirationDisabled(v bool)`

SetTokenExpirationDisabled sets TokenExpirationDisabled field to given value.


### GetUserAttributeEnabled

`func (o *SsoSettingsV2) GetUserAttributeEnabled() bool`

GetUserAttributeEnabled returns the UserAttributeEnabled field if non-nil, zero value otherwise.

### GetUserAttributeEnabledOk

`func (o *SsoSettingsV2) GetUserAttributeEnabledOk() (*bool, bool)`

GetUserAttributeEnabledOk returns a tuple with the UserAttributeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeEnabled

`func (o *SsoSettingsV2) SetUserAttributeEnabled(v bool)`

SetUserAttributeEnabled sets UserAttributeEnabled field to given value.


### GetUserAttributeName

`func (o *SsoSettingsV2) GetUserAttributeName() string`

GetUserAttributeName returns the UserAttributeName field if non-nil, zero value otherwise.

### GetUserAttributeNameOk

`func (o *SsoSettingsV2) GetUserAttributeNameOk() (*string, bool)`

GetUserAttributeNameOk returns a tuple with the UserAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeName

`func (o *SsoSettingsV2) SetUserAttributeName(v string)`

SetUserAttributeName sets UserAttributeName field to given value.

### HasUserAttributeName

`func (o *SsoSettingsV2) HasUserAttributeName() bool`

HasUserAttributeName returns a boolean if a field has been set.

### GetUserMapping

`func (o *SsoSettingsV2) GetUserMapping() string`

GetUserMapping returns the UserMapping field if non-nil, zero value otherwise.

### GetUserMappingOk

`func (o *SsoSettingsV2) GetUserMappingOk() (*string, bool)`

GetUserMappingOk returns a tuple with the UserMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMapping

`func (o *SsoSettingsV2) SetUserMapping(v string)`

SetUserMapping sets UserMapping field to given value.


### GetEnrollmentSsoForAccountDrivenEnrollmentEnabled

`func (o *SsoSettingsV2) GetEnrollmentSsoForAccountDrivenEnrollmentEnabled() bool`

GetEnrollmentSsoForAccountDrivenEnrollmentEnabled returns the EnrollmentSsoForAccountDrivenEnrollmentEnabled field if non-nil, zero value otherwise.

### GetEnrollmentSsoForAccountDrivenEnrollmentEnabledOk

`func (o *SsoSettingsV2) GetEnrollmentSsoForAccountDrivenEnrollmentEnabledOk() (*bool, bool)`

GetEnrollmentSsoForAccountDrivenEnrollmentEnabledOk returns a tuple with the EnrollmentSsoForAccountDrivenEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSsoForAccountDrivenEnrollmentEnabled

`func (o *SsoSettingsV2) SetEnrollmentSsoForAccountDrivenEnrollmentEnabled(v bool)`

SetEnrollmentSsoForAccountDrivenEnrollmentEnabled sets EnrollmentSsoForAccountDrivenEnrollmentEnabled field to given value.


### GetEnrollmentSsoConfig

`func (o *SsoSettingsV2) GetEnrollmentSsoConfig() EnrollmentSsoConfig`

GetEnrollmentSsoConfig returns the EnrollmentSsoConfig field if non-nil, zero value otherwise.

### GetEnrollmentSsoConfigOk

`func (o *SsoSettingsV2) GetEnrollmentSsoConfigOk() (*EnrollmentSsoConfig, bool)`

GetEnrollmentSsoConfigOk returns a tuple with the EnrollmentSsoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSsoConfig

`func (o *SsoSettingsV2) SetEnrollmentSsoConfig(v EnrollmentSsoConfig)`

SetEnrollmentSsoConfig sets EnrollmentSsoConfig field to given value.

### HasEnrollmentSsoConfig

`func (o *SsoSettingsV2) HasEnrollmentSsoConfig() bool`

HasEnrollmentSsoConfig returns a boolean if a field has been set.

### GetGroupEnrollmentAccessEnabled

`func (o *SsoSettingsV2) GetGroupEnrollmentAccessEnabled() bool`

GetGroupEnrollmentAccessEnabled returns the GroupEnrollmentAccessEnabled field if non-nil, zero value otherwise.

### GetGroupEnrollmentAccessEnabledOk

`func (o *SsoSettingsV2) GetGroupEnrollmentAccessEnabledOk() (*bool, bool)`

GetGroupEnrollmentAccessEnabledOk returns a tuple with the GroupEnrollmentAccessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupEnrollmentAccessEnabled

`func (o *SsoSettingsV2) SetGroupEnrollmentAccessEnabled(v bool)`

SetGroupEnrollmentAccessEnabled sets GroupEnrollmentAccessEnabled field to given value.


### GetGroupAttributeName

`func (o *SsoSettingsV2) GetGroupAttributeName() string`

GetGroupAttributeName returns the GroupAttributeName field if non-nil, zero value otherwise.

### GetGroupAttributeNameOk

`func (o *SsoSettingsV2) GetGroupAttributeNameOk() (*string, bool)`

GetGroupAttributeNameOk returns a tuple with the GroupAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAttributeName

`func (o *SsoSettingsV2) SetGroupAttributeName(v string)`

SetGroupAttributeName sets GroupAttributeName field to given value.


### GetGroupRdnKey

`func (o *SsoSettingsV2) GetGroupRdnKey() string`

GetGroupRdnKey returns the GroupRdnKey field if non-nil, zero value otherwise.

### GetGroupRdnKeyOk

`func (o *SsoSettingsV2) GetGroupRdnKeyOk() (*string, bool)`

GetGroupRdnKeyOk returns a tuple with the GroupRdnKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupRdnKey

`func (o *SsoSettingsV2) SetGroupRdnKey(v string)`

SetGroupRdnKey sets GroupRdnKey field to given value.


### GetGroupEnrollmentAccessName

`func (o *SsoSettingsV2) GetGroupEnrollmentAccessName() string`

GetGroupEnrollmentAccessName returns the GroupEnrollmentAccessName field if non-nil, zero value otherwise.

### GetGroupEnrollmentAccessNameOk

`func (o *SsoSettingsV2) GetGroupEnrollmentAccessNameOk() (*string, bool)`

GetGroupEnrollmentAccessNameOk returns a tuple with the GroupEnrollmentAccessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupEnrollmentAccessName

`func (o *SsoSettingsV2) SetGroupEnrollmentAccessName(v string)`

SetGroupEnrollmentAccessName sets GroupEnrollmentAccessName field to given value.

### HasGroupEnrollmentAccessName

`func (o *SsoSettingsV2) HasGroupEnrollmentAccessName() bool`

HasGroupEnrollmentAccessName returns a boolean if a field has been set.

### GetIdpProviderType

`func (o *SsoSettingsV2) GetIdpProviderType() string`

GetIdpProviderType returns the IdpProviderType field if non-nil, zero value otherwise.

### GetIdpProviderTypeOk

`func (o *SsoSettingsV2) GetIdpProviderTypeOk() (*string, bool)`

GetIdpProviderTypeOk returns a tuple with the IdpProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpProviderType

`func (o *SsoSettingsV2) SetIdpProviderType(v string)`

SetIdpProviderType sets IdpProviderType field to given value.


### GetIdpUrl

`func (o *SsoSettingsV2) GetIdpUrl() string`

GetIdpUrl returns the IdpUrl field if non-nil, zero value otherwise.

### GetIdpUrlOk

`func (o *SsoSettingsV2) GetIdpUrlOk() (*string, bool)`

GetIdpUrlOk returns a tuple with the IdpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpUrl

`func (o *SsoSettingsV2) SetIdpUrl(v string)`

SetIdpUrl sets IdpUrl field to given value.

### HasIdpUrl

`func (o *SsoSettingsV2) HasIdpUrl() bool`

HasIdpUrl returns a boolean if a field has been set.

### GetEntityId

`func (o *SsoSettingsV2) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SsoSettingsV2) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SsoSettingsV2) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetMetadataFileName

`func (o *SsoSettingsV2) GetMetadataFileName() string`

GetMetadataFileName returns the MetadataFileName field if non-nil, zero value otherwise.

### GetMetadataFileNameOk

`func (o *SsoSettingsV2) GetMetadataFileNameOk() (*string, bool)`

GetMetadataFileNameOk returns a tuple with the MetadataFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFileName

`func (o *SsoSettingsV2) SetMetadataFileName(v string)`

SetMetadataFileName sets MetadataFileName field to given value.

### HasMetadataFileName

`func (o *SsoSettingsV2) HasMetadataFileName() bool`

HasMetadataFileName returns a boolean if a field has been set.

### GetOtherProviderTypeName

`func (o *SsoSettingsV2) GetOtherProviderTypeName() string`

GetOtherProviderTypeName returns the OtherProviderTypeName field if non-nil, zero value otherwise.

### GetOtherProviderTypeNameOk

`func (o *SsoSettingsV2) GetOtherProviderTypeNameOk() (*string, bool)`

GetOtherProviderTypeNameOk returns a tuple with the OtherProviderTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherProviderTypeName

`func (o *SsoSettingsV2) SetOtherProviderTypeName(v string)`

SetOtherProviderTypeName sets OtherProviderTypeName field to given value.

### HasOtherProviderTypeName

`func (o *SsoSettingsV2) HasOtherProviderTypeName() bool`

HasOtherProviderTypeName returns a boolean if a field has been set.

### GetFederationMetadataFile

`func (o *SsoSettingsV2) GetFederationMetadataFile() string`

GetFederationMetadataFile returns the FederationMetadataFile field if non-nil, zero value otherwise.

### GetFederationMetadataFileOk

`func (o *SsoSettingsV2) GetFederationMetadataFileOk() (*string, bool)`

GetFederationMetadataFileOk returns a tuple with the FederationMetadataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationMetadataFile

`func (o *SsoSettingsV2) SetFederationMetadataFile(v string)`

SetFederationMetadataFile sets FederationMetadataFile field to given value.

### HasFederationMetadataFile

`func (o *SsoSettingsV2) HasFederationMetadataFile() bool`

HasFederationMetadataFile returns a boolean if a field has been set.

### GetMetadataSource

`func (o *SsoSettingsV2) GetMetadataSource() string`

GetMetadataSource returns the MetadataSource field if non-nil, zero value otherwise.

### GetMetadataSourceOk

`func (o *SsoSettingsV2) GetMetadataSourceOk() (*string, bool)`

GetMetadataSourceOk returns a tuple with the MetadataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSource

`func (o *SsoSettingsV2) SetMetadataSource(v string)`

SetMetadataSource sets MetadataSource field to given value.


### GetSessionTimeout

`func (o *SsoSettingsV2) GetSessionTimeout() int64`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *SsoSettingsV2) GetSessionTimeoutOk() (*int64, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *SsoSettingsV2) SetSessionTimeout(v int64)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *SsoSettingsV2) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


