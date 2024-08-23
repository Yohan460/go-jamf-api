# SmtpSenderSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] [default to "Jamf Pro Server"]
**EmailAddress** | **string** |  | [default to ""]

## Methods

### NewSmtpSenderSettings

`func NewSmtpSenderSettings(emailAddress string, ) *SmtpSenderSettings`

NewSmtpSenderSettings instantiates a new SmtpSenderSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpSenderSettingsWithDefaults

`func NewSmtpSenderSettingsWithDefaults() *SmtpSenderSettings`

NewSmtpSenderSettingsWithDefaults instantiates a new SmtpSenderSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *SmtpSenderSettings) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SmtpSenderSettings) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SmtpSenderSettings) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SmtpSenderSettings) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *SmtpSenderSettings) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *SmtpSenderSettings) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *SmtpSenderSettings) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


