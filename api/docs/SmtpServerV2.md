# SmtpServerV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | [default to false]
**AuthenticationType** | **string** |  | 
**ConnectionSettings** | Pointer to [**SmtpConnectionSettings**](SmtpConnectionSettings.md) |  | [optional] 
**SenderSettings** | [**SmtpSenderSettings**](SmtpSenderSettings.md) |  | 
**BasicAuthCredentials** | Pointer to [**SmtpBasicCredentials**](SmtpBasicCredentials.md) |  | [optional] 
**GraphApiCredentials** | Pointer to [**SmtpGraphApiCredentials**](SmtpGraphApiCredentials.md) |  | [optional] 

## Methods

### NewSmtpServerV2

`func NewSmtpServerV2(enabled bool, authenticationType string, senderSettings SmtpSenderSettings, ) *SmtpServerV2`

NewSmtpServerV2 instantiates a new SmtpServerV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpServerV2WithDefaults

`func NewSmtpServerV2WithDefaults() *SmtpServerV2`

NewSmtpServerV2WithDefaults instantiates a new SmtpServerV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SmtpServerV2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SmtpServerV2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SmtpServerV2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAuthenticationType

`func (o *SmtpServerV2) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *SmtpServerV2) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *SmtpServerV2) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetConnectionSettings

`func (o *SmtpServerV2) GetConnectionSettings() SmtpConnectionSettings`

GetConnectionSettings returns the ConnectionSettings field if non-nil, zero value otherwise.

### GetConnectionSettingsOk

`func (o *SmtpServerV2) GetConnectionSettingsOk() (*SmtpConnectionSettings, bool)`

GetConnectionSettingsOk returns a tuple with the ConnectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSettings

`func (o *SmtpServerV2) SetConnectionSettings(v SmtpConnectionSettings)`

SetConnectionSettings sets ConnectionSettings field to given value.

### HasConnectionSettings

`func (o *SmtpServerV2) HasConnectionSettings() bool`

HasConnectionSettings returns a boolean if a field has been set.

### GetSenderSettings

`func (o *SmtpServerV2) GetSenderSettings() SmtpSenderSettings`

GetSenderSettings returns the SenderSettings field if non-nil, zero value otherwise.

### GetSenderSettingsOk

`func (o *SmtpServerV2) GetSenderSettingsOk() (*SmtpSenderSettings, bool)`

GetSenderSettingsOk returns a tuple with the SenderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderSettings

`func (o *SmtpServerV2) SetSenderSettings(v SmtpSenderSettings)`

SetSenderSettings sets SenderSettings field to given value.


### GetBasicAuthCredentials

`func (o *SmtpServerV2) GetBasicAuthCredentials() SmtpBasicCredentials`

GetBasicAuthCredentials returns the BasicAuthCredentials field if non-nil, zero value otherwise.

### GetBasicAuthCredentialsOk

`func (o *SmtpServerV2) GetBasicAuthCredentialsOk() (*SmtpBasicCredentials, bool)`

GetBasicAuthCredentialsOk returns a tuple with the BasicAuthCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthCredentials

`func (o *SmtpServerV2) SetBasicAuthCredentials(v SmtpBasicCredentials)`

SetBasicAuthCredentials sets BasicAuthCredentials field to given value.

### HasBasicAuthCredentials

`func (o *SmtpServerV2) HasBasicAuthCredentials() bool`

HasBasicAuthCredentials returns a boolean if a field has been set.

### GetGraphApiCredentials

`func (o *SmtpServerV2) GetGraphApiCredentials() SmtpGraphApiCredentials`

GetGraphApiCredentials returns the GraphApiCredentials field if non-nil, zero value otherwise.

### GetGraphApiCredentialsOk

`func (o *SmtpServerV2) GetGraphApiCredentialsOk() (*SmtpGraphApiCredentials, bool)`

GetGraphApiCredentialsOk returns a tuple with the GraphApiCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphApiCredentials

`func (o *SmtpServerV2) SetGraphApiCredentials(v SmtpGraphApiCredentials)`

SetGraphApiCredentials sets GraphApiCredentials field to given value.

### HasGraphApiCredentials

`func (o *SmtpServerV2) HasGraphApiCredentials() bool`

HasGraphApiCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


