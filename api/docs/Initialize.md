# Initialize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationCode** | **string** |  | 
**InstitutionName** | **string** |  | 
**IsEulaAccepted** | **bool** |  | 
**Username** | **string** |  | 
**Password** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**JssUrl** | **string** |  | 

## Methods

### NewInitialize

`func NewInitialize(activationCode string, institutionName string, isEulaAccepted bool, username string, password string, jssUrl string, ) *Initialize`

NewInitialize instantiates a new Initialize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitializeWithDefaults

`func NewInitializeWithDefaults() *Initialize`

NewInitializeWithDefaults instantiates a new Initialize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationCode

`func (o *Initialize) GetActivationCode() string`

GetActivationCode returns the ActivationCode field if non-nil, zero value otherwise.

### GetActivationCodeOk

`func (o *Initialize) GetActivationCodeOk() (*string, bool)`

GetActivationCodeOk returns a tuple with the ActivationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationCode

`func (o *Initialize) SetActivationCode(v string)`

SetActivationCode sets ActivationCode field to given value.


### GetInstitutionName

`func (o *Initialize) GetInstitutionName() string`

GetInstitutionName returns the InstitutionName field if non-nil, zero value otherwise.

### GetInstitutionNameOk

`func (o *Initialize) GetInstitutionNameOk() (*string, bool)`

GetInstitutionNameOk returns a tuple with the InstitutionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionName

`func (o *Initialize) SetInstitutionName(v string)`

SetInstitutionName sets InstitutionName field to given value.


### GetIsEulaAccepted

`func (o *Initialize) GetIsEulaAccepted() bool`

GetIsEulaAccepted returns the IsEulaAccepted field if non-nil, zero value otherwise.

### GetIsEulaAcceptedOk

`func (o *Initialize) GetIsEulaAcceptedOk() (*bool, bool)`

GetIsEulaAcceptedOk returns a tuple with the IsEulaAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEulaAccepted

`func (o *Initialize) SetIsEulaAccepted(v bool)`

SetIsEulaAccepted sets IsEulaAccepted field to given value.


### GetUsername

`func (o *Initialize) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Initialize) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Initialize) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *Initialize) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Initialize) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Initialize) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetEmail

`func (o *Initialize) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Initialize) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Initialize) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Initialize) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetJssUrl

`func (o *Initialize) GetJssUrl() string`

GetJssUrl returns the JssUrl field if non-nil, zero value otherwise.

### GetJssUrlOk

`func (o *Initialize) GetJssUrlOk() (*string, bool)`

GetJssUrlOk returns a tuple with the JssUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJssUrl

`func (o *Initialize) SetJssUrl(v string)`

SetJssUrl sets JssUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


