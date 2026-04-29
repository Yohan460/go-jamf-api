# PlatformInitializeV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationCode** | **string** |  | 
**InstitutionName** | **string** |  | 
**EulaAccepted** | **bool** |  | 
**Username** | **string** | Federated user OIDC username to create | 
**Email** | **string** | Federated user OIDC email to create | 
**JssUrl** | **string** |  | 

## Methods

### NewPlatformInitializeV1

`func NewPlatformInitializeV1(activationCode string, institutionName string, eulaAccepted bool, username string, email string, jssUrl string, ) *PlatformInitializeV1`

NewPlatformInitializeV1 instantiates a new PlatformInitializeV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformInitializeV1WithDefaults

`func NewPlatformInitializeV1WithDefaults() *PlatformInitializeV1`

NewPlatformInitializeV1WithDefaults instantiates a new PlatformInitializeV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationCode

`func (o *PlatformInitializeV1) GetActivationCode() string`

GetActivationCode returns the ActivationCode field if non-nil, zero value otherwise.

### GetActivationCodeOk

`func (o *PlatformInitializeV1) GetActivationCodeOk() (*string, bool)`

GetActivationCodeOk returns a tuple with the ActivationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationCode

`func (o *PlatformInitializeV1) SetActivationCode(v string)`

SetActivationCode sets ActivationCode field to given value.


### GetInstitutionName

`func (o *PlatformInitializeV1) GetInstitutionName() string`

GetInstitutionName returns the InstitutionName field if non-nil, zero value otherwise.

### GetInstitutionNameOk

`func (o *PlatformInitializeV1) GetInstitutionNameOk() (*string, bool)`

GetInstitutionNameOk returns a tuple with the InstitutionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionName

`func (o *PlatformInitializeV1) SetInstitutionName(v string)`

SetInstitutionName sets InstitutionName field to given value.


### GetEulaAccepted

`func (o *PlatformInitializeV1) GetEulaAccepted() bool`

GetEulaAccepted returns the EulaAccepted field if non-nil, zero value otherwise.

### GetEulaAcceptedOk

`func (o *PlatformInitializeV1) GetEulaAcceptedOk() (*bool, bool)`

GetEulaAcceptedOk returns a tuple with the EulaAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaAccepted

`func (o *PlatformInitializeV1) SetEulaAccepted(v bool)`

SetEulaAccepted sets EulaAccepted field to given value.


### GetUsername

`func (o *PlatformInitializeV1) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PlatformInitializeV1) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PlatformInitializeV1) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *PlatformInitializeV1) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PlatformInitializeV1) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PlatformInitializeV1) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetJssUrl

`func (o *PlatformInitializeV1) GetJssUrl() string`

GetJssUrl returns the JssUrl field if non-nil, zero value otherwise.

### GetJssUrlOk

`func (o *PlatformInitializeV1) GetJssUrlOk() (*string, bool)`

GetJssUrlOk returns a tuple with the JssUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJssUrl

`func (o *PlatformInitializeV1) SetJssUrl(v string)`

SetJssUrl sets JssUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


