# SmtpConnectionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** |  | [default to ""]
**Port** | **int64** |  | [default to 25]
**EncryptionType** | **string** |  | [default to "NONE"]
**ConnectionTimeout** | **int64** |  | [default to 5]

## Methods

### NewSmtpConnectionSettings

`func NewSmtpConnectionSettings(host string, port int64, encryptionType string, connectionTimeout int64, ) *SmtpConnectionSettings`

NewSmtpConnectionSettings instantiates a new SmtpConnectionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpConnectionSettingsWithDefaults

`func NewSmtpConnectionSettingsWithDefaults() *SmtpConnectionSettings`

NewSmtpConnectionSettingsWithDefaults instantiates a new SmtpConnectionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SmtpConnectionSettings) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SmtpConnectionSettings) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SmtpConnectionSettings) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *SmtpConnectionSettings) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SmtpConnectionSettings) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SmtpConnectionSettings) SetPort(v int64)`

SetPort sets Port field to given value.


### GetEncryptionType

`func (o *SmtpConnectionSettings) GetEncryptionType() string`

GetEncryptionType returns the EncryptionType field if non-nil, zero value otherwise.

### GetEncryptionTypeOk

`func (o *SmtpConnectionSettings) GetEncryptionTypeOk() (*string, bool)`

GetEncryptionTypeOk returns a tuple with the EncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionType

`func (o *SmtpConnectionSettings) SetEncryptionType(v string)`

SetEncryptionType sets EncryptionType field to given value.


### GetConnectionTimeout

`func (o *SmtpConnectionSettings) GetConnectionTimeout() int64`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *SmtpConnectionSettings) GetConnectionTimeoutOk() (*int64, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *SmtpConnectionSettings) SetConnectionTimeout(v int64)`

SetConnectionTimeout sets ConnectionTimeout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


