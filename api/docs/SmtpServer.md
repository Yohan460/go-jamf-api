# SmtpServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | [default to false]
**Server** | **string** |  | [default to ""]
**Port** | **int64** |  | [default to 25]
**EncryptionType** | **string** |  | [default to "NONE"]
**ConnectionTimeout** | **int64** |  | [default to 5]
**SenderDisplayName** | **string** |  | [default to "Jamf Pro Server"]
**SenderEmailAddress** | **string** |  | [default to ""]
**RequiresAuthentication** | **bool** |  | [default to false]
**Username** | Pointer to **string** |  | [optional] [default to ""]
**Password** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewSmtpServer

`func NewSmtpServer(enabled bool, server string, port int64, encryptionType string, connectionTimeout int64, senderDisplayName string, senderEmailAddress string, requiresAuthentication bool, ) *SmtpServer`

NewSmtpServer instantiates a new SmtpServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpServerWithDefaults

`func NewSmtpServerWithDefaults() *SmtpServer`

NewSmtpServerWithDefaults instantiates a new SmtpServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SmtpServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SmtpServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SmtpServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetServer

`func (o *SmtpServer) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *SmtpServer) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *SmtpServer) SetServer(v string)`

SetServer sets Server field to given value.


### GetPort

`func (o *SmtpServer) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SmtpServer) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SmtpServer) SetPort(v int64)`

SetPort sets Port field to given value.


### GetEncryptionType

`func (o *SmtpServer) GetEncryptionType() string`

GetEncryptionType returns the EncryptionType field if non-nil, zero value otherwise.

### GetEncryptionTypeOk

`func (o *SmtpServer) GetEncryptionTypeOk() (*string, bool)`

GetEncryptionTypeOk returns a tuple with the EncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionType

`func (o *SmtpServer) SetEncryptionType(v string)`

SetEncryptionType sets EncryptionType field to given value.


### GetConnectionTimeout

`func (o *SmtpServer) GetConnectionTimeout() int64`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *SmtpServer) GetConnectionTimeoutOk() (*int64, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *SmtpServer) SetConnectionTimeout(v int64)`

SetConnectionTimeout sets ConnectionTimeout field to given value.


### GetSenderDisplayName

`func (o *SmtpServer) GetSenderDisplayName() string`

GetSenderDisplayName returns the SenderDisplayName field if non-nil, zero value otherwise.

### GetSenderDisplayNameOk

`func (o *SmtpServer) GetSenderDisplayNameOk() (*string, bool)`

GetSenderDisplayNameOk returns a tuple with the SenderDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderDisplayName

`func (o *SmtpServer) SetSenderDisplayName(v string)`

SetSenderDisplayName sets SenderDisplayName field to given value.


### GetSenderEmailAddress

`func (o *SmtpServer) GetSenderEmailAddress() string`

GetSenderEmailAddress returns the SenderEmailAddress field if non-nil, zero value otherwise.

### GetSenderEmailAddressOk

`func (o *SmtpServer) GetSenderEmailAddressOk() (*string, bool)`

GetSenderEmailAddressOk returns a tuple with the SenderEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderEmailAddress

`func (o *SmtpServer) SetSenderEmailAddress(v string)`

SetSenderEmailAddress sets SenderEmailAddress field to given value.


### GetRequiresAuthentication

`func (o *SmtpServer) GetRequiresAuthentication() bool`

GetRequiresAuthentication returns the RequiresAuthentication field if non-nil, zero value otherwise.

### GetRequiresAuthenticationOk

`func (o *SmtpServer) GetRequiresAuthenticationOk() (*bool, bool)`

GetRequiresAuthenticationOk returns a tuple with the RequiresAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresAuthentication

`func (o *SmtpServer) SetRequiresAuthentication(v bool)`

SetRequiresAuthentication sets RequiresAuthentication field to given value.


### GetUsername

`func (o *SmtpServer) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SmtpServer) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SmtpServer) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SmtpServer) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *SmtpServer) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SmtpServer) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SmtpServer) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SmtpServer) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


