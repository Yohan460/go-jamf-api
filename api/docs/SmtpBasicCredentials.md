# SmtpBasicCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | [default to ""]
**Password** | **string** |  | [default to ""]

## Methods

### NewSmtpBasicCredentials

`func NewSmtpBasicCredentials(username string, password string, ) *SmtpBasicCredentials`

NewSmtpBasicCredentials instantiates a new SmtpBasicCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpBasicCredentialsWithDefaults

`func NewSmtpBasicCredentialsWithDefaults() *SmtpBasicCredentials`

NewSmtpBasicCredentialsWithDefaults instantiates a new SmtpBasicCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *SmtpBasicCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SmtpBasicCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SmtpBasicCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *SmtpBasicCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SmtpBasicCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SmtpBasicCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


