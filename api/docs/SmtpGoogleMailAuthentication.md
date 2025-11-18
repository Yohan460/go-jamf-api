# SmtpGoogleMailAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** |  | [default to ""]
**Status** | **string** |  | [readonly] 

## Methods

### NewSmtpGoogleMailAuthentication

`func NewSmtpGoogleMailAuthentication(emailAddress string, status string, ) *SmtpGoogleMailAuthentication`

NewSmtpGoogleMailAuthentication instantiates a new SmtpGoogleMailAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpGoogleMailAuthenticationWithDefaults

`func NewSmtpGoogleMailAuthenticationWithDefaults() *SmtpGoogleMailAuthentication`

NewSmtpGoogleMailAuthenticationWithDefaults instantiates a new SmtpGoogleMailAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *SmtpGoogleMailAuthentication) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *SmtpGoogleMailAuthentication) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *SmtpGoogleMailAuthentication) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetStatus

`func (o *SmtpGoogleMailAuthentication) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SmtpGoogleMailAuthentication) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SmtpGoogleMailAuthentication) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


