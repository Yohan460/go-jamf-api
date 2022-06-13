# InitializeV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationCode** | **string** |  | 
**InstitutionName** | **string** |  | 
**EulaAccepted** | **bool** |  | 
**Username** | **string** |  | 
**Password** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**JssUrl** | **string** |  | 

## Methods

### NewInitializeV1

`func NewInitializeV1(activationCode string, institutionName string, eulaAccepted bool, username string, password string, jssUrl string, ) *InitializeV1`

NewInitializeV1 instantiates a new InitializeV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitializeV1WithDefaults

`func NewInitializeV1WithDefaults() *InitializeV1`

NewInitializeV1WithDefaults instantiates a new InitializeV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationCode

`func (o *InitializeV1) GetActivationCode() string`

GetActivationCode returns the ActivationCode field if non-nil, zero value otherwise.

### GetActivationCodeOk

`func (o *InitializeV1) GetActivationCodeOk() (*string, bool)`

GetActivationCodeOk returns a tuple with the ActivationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationCode

`func (o *InitializeV1) SetActivationCode(v string)`

SetActivationCode sets ActivationCode field to given value.


### GetInstitutionName

`func (o *InitializeV1) GetInstitutionName() string`

GetInstitutionName returns the InstitutionName field if non-nil, zero value otherwise.

### GetInstitutionNameOk

`func (o *InitializeV1) GetInstitutionNameOk() (*string, bool)`

GetInstitutionNameOk returns a tuple with the InstitutionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionName

`func (o *InitializeV1) SetInstitutionName(v string)`

SetInstitutionName sets InstitutionName field to given value.


### GetEulaAccepted

`func (o *InitializeV1) GetEulaAccepted() bool`

GetEulaAccepted returns the EulaAccepted field if non-nil, zero value otherwise.

### GetEulaAcceptedOk

`func (o *InitializeV1) GetEulaAcceptedOk() (*bool, bool)`

GetEulaAcceptedOk returns a tuple with the EulaAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaAccepted

`func (o *InitializeV1) SetEulaAccepted(v bool)`

SetEulaAccepted sets EulaAccepted field to given value.


### GetUsername

`func (o *InitializeV1) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InitializeV1) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InitializeV1) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *InitializeV1) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InitializeV1) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InitializeV1) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetEmail

`func (o *InitializeV1) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InitializeV1) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InitializeV1) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InitializeV1) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetJssUrl

`func (o *InitializeV1) GetJssUrl() string`

GetJssUrl returns the JssUrl field if non-nil, zero value otherwise.

### GetJssUrlOk

`func (o *InitializeV1) GetJssUrlOk() (*string, bool)`

GetJssUrlOk returns a tuple with the JssUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJssUrl

`func (o *InitializeV1) SetJssUrl(v string)`

SetJssUrl sets JssUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


