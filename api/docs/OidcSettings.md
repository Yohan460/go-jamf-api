# OidcSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserMapping** | **string** |  | 
**JamfIdAuthenticationEnabled** | Pointer to **bool** |  | [optional] [default to true]
**UsernameAttributeClaimMapping** | Pointer to **string** |  | [optional] 

## Methods

### NewOidcSettings

`func NewOidcSettings(userMapping string, ) *OidcSettings`

NewOidcSettings instantiates a new OidcSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcSettingsWithDefaults

`func NewOidcSettingsWithDefaults() *OidcSettings`

NewOidcSettingsWithDefaults instantiates a new OidcSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserMapping

`func (o *OidcSettings) GetUserMapping() string`

GetUserMapping returns the UserMapping field if non-nil, zero value otherwise.

### GetUserMappingOk

`func (o *OidcSettings) GetUserMappingOk() (*string, bool)`

GetUserMappingOk returns a tuple with the UserMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMapping

`func (o *OidcSettings) SetUserMapping(v string)`

SetUserMapping sets UserMapping field to given value.


### GetJamfIdAuthenticationEnabled

`func (o *OidcSettings) GetJamfIdAuthenticationEnabled() bool`

GetJamfIdAuthenticationEnabled returns the JamfIdAuthenticationEnabled field if non-nil, zero value otherwise.

### GetJamfIdAuthenticationEnabledOk

`func (o *OidcSettings) GetJamfIdAuthenticationEnabledOk() (*bool, bool)`

GetJamfIdAuthenticationEnabledOk returns a tuple with the JamfIdAuthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJamfIdAuthenticationEnabled

`func (o *OidcSettings) SetJamfIdAuthenticationEnabled(v bool)`

SetJamfIdAuthenticationEnabled sets JamfIdAuthenticationEnabled field to given value.

### HasJamfIdAuthenticationEnabled

`func (o *OidcSettings) HasJamfIdAuthenticationEnabled() bool`

HasJamfIdAuthenticationEnabled returns a boolean if a field has been set.

### GetUsernameAttributeClaimMapping

`func (o *OidcSettings) GetUsernameAttributeClaimMapping() string`

GetUsernameAttributeClaimMapping returns the UsernameAttributeClaimMapping field if non-nil, zero value otherwise.

### GetUsernameAttributeClaimMappingOk

`func (o *OidcSettings) GetUsernameAttributeClaimMappingOk() (*string, bool)`

GetUsernameAttributeClaimMappingOk returns a tuple with the UsernameAttributeClaimMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameAttributeClaimMapping

`func (o *OidcSettings) SetUsernameAttributeClaimMapping(v string)`

SetUsernameAttributeClaimMapping sets UsernameAttributeClaimMapping field to given value.

### HasUsernameAttributeClaimMapping

`func (o *OidcSettings) HasUsernameAttributeClaimMapping() bool`

HasUsernameAttributeClaimMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


