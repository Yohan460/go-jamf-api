# OidcPublicFeaturesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JamfIdAuthenticationEnabled** | **bool** | Indicates whether Jamf ID authentication is enabled for this instance. When true, users can authenticate using Jamf ID credentials. When false, Jamf ID login option is not available. | 

## Methods

### NewOidcPublicFeaturesResponse

`func NewOidcPublicFeaturesResponse(jamfIdAuthenticationEnabled bool, ) *OidcPublicFeaturesResponse`

NewOidcPublicFeaturesResponse instantiates a new OidcPublicFeaturesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcPublicFeaturesResponseWithDefaults

`func NewOidcPublicFeaturesResponseWithDefaults() *OidcPublicFeaturesResponse`

NewOidcPublicFeaturesResponseWithDefaults instantiates a new OidcPublicFeaturesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJamfIdAuthenticationEnabled

`func (o *OidcPublicFeaturesResponse) GetJamfIdAuthenticationEnabled() bool`

GetJamfIdAuthenticationEnabled returns the JamfIdAuthenticationEnabled field if non-nil, zero value otherwise.

### GetJamfIdAuthenticationEnabledOk

`func (o *OidcPublicFeaturesResponse) GetJamfIdAuthenticationEnabledOk() (*bool, bool)`

GetJamfIdAuthenticationEnabledOk returns a tuple with the JamfIdAuthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJamfIdAuthenticationEnabled

`func (o *OidcPublicFeaturesResponse) SetJamfIdAuthenticationEnabled(v bool)`

SetJamfIdAuthenticationEnabled sets JamfIdAuthenticationEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


