# SsoSettingsV1

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
**EnrollmentSsoForAdueEnabled** | **bool** |  | [default to false]
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

### NewSsoSettingsV1

`func NewSsoSettingsV1(ssoForEnrollmentEnabled bool, ssoBypassAllowed bool, ssoEnabled bool, ssoForMacOsSelfServiceEnabled bool, tokenExpirationDisabled bool, userAttributeEnabled bool, userMapping string, enrollmentSsoForAdueEnabled bool, groupEnrollmentAccessEnabled bool, groupAttributeName string, groupRdnKey string, idpProviderType string, entityId string, metadataSource string, ) *SsoSettingsV1`

NewSsoSettingsV1 instantiates a new SsoSettingsV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoSettingsV1WithDefaults

`func NewSsoSettingsV1WithDefaults() *SsoSettingsV1`

NewSsoSettingsV1WithDefaults instantiates a new SsoSettingsV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsoForEnrollmentEnabled

`func (o *SsoSettingsV1) GetSsoForEnrollmentEnabled() bool`

GetSsoForEnrollmentEnabled returns the SsoForEnrollmentEnabled field if non-nil, zero value otherwise.

### GetSsoForEnrollmentEnabledOk

`func (o *SsoSettingsV1) GetSsoForEnrollmentEnabledOk() (*bool, bool)`

GetSsoForEnrollmentEnabledOk returns a tuple with the SsoForEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoForEnrollmentEnabled

`func (o *SsoSettingsV1) SetSsoForEnrollmentEnabled(v bool)`

SetSsoForEnrollmentEnabled sets SsoForEnrollmentEnabled field to given value.


### GetSsoBypassAllowed

`func (o *SsoSettingsV1) GetSsoBypassAllowed() bool`

GetSsoBypassAllowed returns the SsoBypassAllowed field if non-nil, zero value otherwise.

### GetSsoBypassAllowedOk

`func (o *SsoSettingsV1) GetSsoBypassAllowedOk() (*bool, bool)`

GetSsoBypassAllowedOk returns a tuple with the SsoBypassAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoBypassAllowed

`func (o *SsoSettingsV1) SetSsoBypassAllowed(v bool)`

SetSsoBypassAllowed sets SsoBypassAllowed field to given value.


### GetSsoEnabled

`func (o *SsoSettingsV1) GetSsoEnabled() bool`

GetSsoEnabled returns the SsoEnabled field if non-nil, zero value otherwise.

### GetSsoEnabledOk

`func (o *SsoSettingsV1) GetSsoEnabledOk() (*bool, bool)`

GetSsoEnabledOk returns a tuple with the SsoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoEnabled

`func (o *SsoSettingsV1) SetSsoEnabled(v bool)`

SetSsoEnabled sets SsoEnabled field to given value.


### GetSsoForMacOsSelfServiceEnabled

`func (o *SsoSettingsV1) GetSsoForMacOsSelfServiceEnabled() bool`

GetSsoForMacOsSelfServiceEnabled returns the SsoForMacOsSelfServiceEnabled field if non-nil, zero value otherwise.

### GetSsoForMacOsSelfServiceEnabledOk

`func (o *SsoSettingsV1) GetSsoForMacOsSelfServiceEnabledOk() (*bool, bool)`

GetSsoForMacOsSelfServiceEnabledOk returns a tuple with the SsoForMacOsSelfServiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoForMacOsSelfServiceEnabled

`func (o *SsoSettingsV1) SetSsoForMacOsSelfServiceEnabled(v bool)`

SetSsoForMacOsSelfServiceEnabled sets SsoForMacOsSelfServiceEnabled field to given value.


### GetTokenExpirationDisabled

`func (o *SsoSettingsV1) GetTokenExpirationDisabled() bool`

GetTokenExpirationDisabled returns the TokenExpirationDisabled field if non-nil, zero value otherwise.

### GetTokenExpirationDisabledOk

`func (o *SsoSettingsV1) GetTokenExpirationDisabledOk() (*bool, bool)`

GetTokenExpirationDisabledOk returns a tuple with the TokenExpirationDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpirationDisabled

`func (o *SsoSettingsV1) SetTokenExpirationDisabled(v bool)`

SetTokenExpirationDisabled sets TokenExpirationDisabled field to given value.


### GetUserAttributeEnabled

`func (o *SsoSettingsV1) GetUserAttributeEnabled() bool`

GetUserAttributeEnabled returns the UserAttributeEnabled field if non-nil, zero value otherwise.

### GetUserAttributeEnabledOk

`func (o *SsoSettingsV1) GetUserAttributeEnabledOk() (*bool, bool)`

GetUserAttributeEnabledOk returns a tuple with the UserAttributeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeEnabled

`func (o *SsoSettingsV1) SetUserAttributeEnabled(v bool)`

SetUserAttributeEnabled sets UserAttributeEnabled field to given value.


### GetUserAttributeName

`func (o *SsoSettingsV1) GetUserAttributeName() string`

GetUserAttributeName returns the UserAttributeName field if non-nil, zero value otherwise.

### GetUserAttributeNameOk

`func (o *SsoSettingsV1) GetUserAttributeNameOk() (*string, bool)`

GetUserAttributeNameOk returns a tuple with the UserAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeName

`func (o *SsoSettingsV1) SetUserAttributeName(v string)`

SetUserAttributeName sets UserAttributeName field to given value.

### HasUserAttributeName

`func (o *SsoSettingsV1) HasUserAttributeName() bool`

HasUserAttributeName returns a boolean if a field has been set.

### GetUserMapping

`func (o *SsoSettingsV1) GetUserMapping() string`

GetUserMapping returns the UserMapping field if non-nil, zero value otherwise.

### GetUserMappingOk

`func (o *SsoSettingsV1) GetUserMappingOk() (*string, bool)`

GetUserMappingOk returns a tuple with the UserMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMapping

`func (o *SsoSettingsV1) SetUserMapping(v string)`

SetUserMapping sets UserMapping field to given value.


### GetEnrollmentSsoForAdueEnabled

`func (o *SsoSettingsV1) GetEnrollmentSsoForAdueEnabled() bool`

GetEnrollmentSsoForAdueEnabled returns the EnrollmentSsoForAdueEnabled field if non-nil, zero value otherwise.

### GetEnrollmentSsoForAdueEnabledOk

`func (o *SsoSettingsV1) GetEnrollmentSsoForAdueEnabledOk() (*bool, bool)`

GetEnrollmentSsoForAdueEnabledOk returns a tuple with the EnrollmentSsoForAdueEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSsoForAdueEnabled

`func (o *SsoSettingsV1) SetEnrollmentSsoForAdueEnabled(v bool)`

SetEnrollmentSsoForAdueEnabled sets EnrollmentSsoForAdueEnabled field to given value.


### GetEnrollmentSsoConfig

`func (o *SsoSettingsV1) GetEnrollmentSsoConfig() EnrollmentSsoConfig`

GetEnrollmentSsoConfig returns the EnrollmentSsoConfig field if non-nil, zero value otherwise.

### GetEnrollmentSsoConfigOk

`func (o *SsoSettingsV1) GetEnrollmentSsoConfigOk() (*EnrollmentSsoConfig, bool)`

GetEnrollmentSsoConfigOk returns a tuple with the EnrollmentSsoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSsoConfig

`func (o *SsoSettingsV1) SetEnrollmentSsoConfig(v EnrollmentSsoConfig)`

SetEnrollmentSsoConfig sets EnrollmentSsoConfig field to given value.

### HasEnrollmentSsoConfig

`func (o *SsoSettingsV1) HasEnrollmentSsoConfig() bool`

HasEnrollmentSsoConfig returns a boolean if a field has been set.

### GetGroupEnrollmentAccessEnabled

`func (o *SsoSettingsV1) GetGroupEnrollmentAccessEnabled() bool`

GetGroupEnrollmentAccessEnabled returns the GroupEnrollmentAccessEnabled field if non-nil, zero value otherwise.

### GetGroupEnrollmentAccessEnabledOk

`func (o *SsoSettingsV1) GetGroupEnrollmentAccessEnabledOk() (*bool, bool)`

GetGroupEnrollmentAccessEnabledOk returns a tuple with the GroupEnrollmentAccessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupEnrollmentAccessEnabled

`func (o *SsoSettingsV1) SetGroupEnrollmentAccessEnabled(v bool)`

SetGroupEnrollmentAccessEnabled sets GroupEnrollmentAccessEnabled field to given value.


### GetGroupAttributeName

`func (o *SsoSettingsV1) GetGroupAttributeName() string`

GetGroupAttributeName returns the GroupAttributeName field if non-nil, zero value otherwise.

### GetGroupAttributeNameOk

`func (o *SsoSettingsV1) GetGroupAttributeNameOk() (*string, bool)`

GetGroupAttributeNameOk returns a tuple with the GroupAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAttributeName

`func (o *SsoSettingsV1) SetGroupAttributeName(v string)`

SetGroupAttributeName sets GroupAttributeName field to given value.


### GetGroupRdnKey

`func (o *SsoSettingsV1) GetGroupRdnKey() string`

GetGroupRdnKey returns the GroupRdnKey field if non-nil, zero value otherwise.

### GetGroupRdnKeyOk

`func (o *SsoSettingsV1) GetGroupRdnKeyOk() (*string, bool)`

GetGroupRdnKeyOk returns a tuple with the GroupRdnKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupRdnKey

`func (o *SsoSettingsV1) SetGroupRdnKey(v string)`

SetGroupRdnKey sets GroupRdnKey field to given value.


### GetGroupEnrollmentAccessName

`func (o *SsoSettingsV1) GetGroupEnrollmentAccessName() string`

GetGroupEnrollmentAccessName returns the GroupEnrollmentAccessName field if non-nil, zero value otherwise.

### GetGroupEnrollmentAccessNameOk

`func (o *SsoSettingsV1) GetGroupEnrollmentAccessNameOk() (*string, bool)`

GetGroupEnrollmentAccessNameOk returns a tuple with the GroupEnrollmentAccessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupEnrollmentAccessName

`func (o *SsoSettingsV1) SetGroupEnrollmentAccessName(v string)`

SetGroupEnrollmentAccessName sets GroupEnrollmentAccessName field to given value.

### HasGroupEnrollmentAccessName

`func (o *SsoSettingsV1) HasGroupEnrollmentAccessName() bool`

HasGroupEnrollmentAccessName returns a boolean if a field has been set.

### GetIdpProviderType

`func (o *SsoSettingsV1) GetIdpProviderType() string`

GetIdpProviderType returns the IdpProviderType field if non-nil, zero value otherwise.

### GetIdpProviderTypeOk

`func (o *SsoSettingsV1) GetIdpProviderTypeOk() (*string, bool)`

GetIdpProviderTypeOk returns a tuple with the IdpProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpProviderType

`func (o *SsoSettingsV1) SetIdpProviderType(v string)`

SetIdpProviderType sets IdpProviderType field to given value.


### GetIdpUrl

`func (o *SsoSettingsV1) GetIdpUrl() string`

GetIdpUrl returns the IdpUrl field if non-nil, zero value otherwise.

### GetIdpUrlOk

`func (o *SsoSettingsV1) GetIdpUrlOk() (*string, bool)`

GetIdpUrlOk returns a tuple with the IdpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpUrl

`func (o *SsoSettingsV1) SetIdpUrl(v string)`

SetIdpUrl sets IdpUrl field to given value.

### HasIdpUrl

`func (o *SsoSettingsV1) HasIdpUrl() bool`

HasIdpUrl returns a boolean if a field has been set.

### GetEntityId

`func (o *SsoSettingsV1) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SsoSettingsV1) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SsoSettingsV1) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetMetadataFileName

`func (o *SsoSettingsV1) GetMetadataFileName() string`

GetMetadataFileName returns the MetadataFileName field if non-nil, zero value otherwise.

### GetMetadataFileNameOk

`func (o *SsoSettingsV1) GetMetadataFileNameOk() (*string, bool)`

GetMetadataFileNameOk returns a tuple with the MetadataFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFileName

`func (o *SsoSettingsV1) SetMetadataFileName(v string)`

SetMetadataFileName sets MetadataFileName field to given value.

### HasMetadataFileName

`func (o *SsoSettingsV1) HasMetadataFileName() bool`

HasMetadataFileName returns a boolean if a field has been set.

### GetOtherProviderTypeName

`func (o *SsoSettingsV1) GetOtherProviderTypeName() string`

GetOtherProviderTypeName returns the OtherProviderTypeName field if non-nil, zero value otherwise.

### GetOtherProviderTypeNameOk

`func (o *SsoSettingsV1) GetOtherProviderTypeNameOk() (*string, bool)`

GetOtherProviderTypeNameOk returns a tuple with the OtherProviderTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherProviderTypeName

`func (o *SsoSettingsV1) SetOtherProviderTypeName(v string)`

SetOtherProviderTypeName sets OtherProviderTypeName field to given value.

### HasOtherProviderTypeName

`func (o *SsoSettingsV1) HasOtherProviderTypeName() bool`

HasOtherProviderTypeName returns a boolean if a field has been set.

### GetFederationMetadataFile

`func (o *SsoSettingsV1) GetFederationMetadataFile() string`

GetFederationMetadataFile returns the FederationMetadataFile field if non-nil, zero value otherwise.

### GetFederationMetadataFileOk

`func (o *SsoSettingsV1) GetFederationMetadataFileOk() (*string, bool)`

GetFederationMetadataFileOk returns a tuple with the FederationMetadataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationMetadataFile

`func (o *SsoSettingsV1) SetFederationMetadataFile(v string)`

SetFederationMetadataFile sets FederationMetadataFile field to given value.

### HasFederationMetadataFile

`func (o *SsoSettingsV1) HasFederationMetadataFile() bool`

HasFederationMetadataFile returns a boolean if a field has been set.

### GetMetadataSource

`func (o *SsoSettingsV1) GetMetadataSource() string`

GetMetadataSource returns the MetadataSource field if non-nil, zero value otherwise.

### GetMetadataSourceOk

`func (o *SsoSettingsV1) GetMetadataSourceOk() (*string, bool)`

GetMetadataSourceOk returns a tuple with the MetadataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSource

`func (o *SsoSettingsV1) SetMetadataSource(v string)`

SetMetadataSource sets MetadataSource field to given value.


### GetSessionTimeout

`func (o *SsoSettingsV1) GetSessionTimeout() int64`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *SsoSettingsV1) GetSessionTimeoutOk() (*int64, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *SsoSettingsV1) SetSessionTimeout(v int64)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *SsoSettingsV1) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


