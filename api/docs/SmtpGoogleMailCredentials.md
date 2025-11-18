# SmtpGoogleMailCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | [default to ""]
**ClientSecret** | **string** |  | [default to ""]
**Authentications** | Pointer to [**[]SmtpGoogleMailAuthentication**](SmtpGoogleMailAuthentication.md) |  | [optional] 

## Methods

### NewSmtpGoogleMailCredentials

`func NewSmtpGoogleMailCredentials(clientId string, clientSecret string, ) *SmtpGoogleMailCredentials`

NewSmtpGoogleMailCredentials instantiates a new SmtpGoogleMailCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpGoogleMailCredentialsWithDefaults

`func NewSmtpGoogleMailCredentialsWithDefaults() *SmtpGoogleMailCredentials`

NewSmtpGoogleMailCredentialsWithDefaults instantiates a new SmtpGoogleMailCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *SmtpGoogleMailCredentials) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SmtpGoogleMailCredentials) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SmtpGoogleMailCredentials) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *SmtpGoogleMailCredentials) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *SmtpGoogleMailCredentials) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *SmtpGoogleMailCredentials) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetAuthentications

`func (o *SmtpGoogleMailCredentials) GetAuthentications() []SmtpGoogleMailAuthentication`

GetAuthentications returns the Authentications field if non-nil, zero value otherwise.

### GetAuthenticationsOk

`func (o *SmtpGoogleMailCredentials) GetAuthenticationsOk() (*[]SmtpGoogleMailAuthentication, bool)`

GetAuthenticationsOk returns a tuple with the Authentications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentications

`func (o *SmtpGoogleMailCredentials) SetAuthentications(v []SmtpGoogleMailAuthentication)`

SetAuthentications sets Authentications field to given value.

### HasAuthentications

`func (o *SmtpGoogleMailCredentials) HasAuthentications() bool`

HasAuthentications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


