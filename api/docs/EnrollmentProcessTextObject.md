# EnrollmentProcessTextObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LanguageCode** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**LoginDescription** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**LoginButton** | Pointer to **string** |  | [optional] 
**DeviceClassDescription** | Pointer to **string** |  | [optional] 
**DeviceClassPersonal** | Pointer to **string** |  | [optional] 
**DeviceClassPersonalDescription** | Pointer to **string** |  | [optional] 
**DeviceClassEnterprise** | Pointer to **string** |  | [optional] 
**DeviceClassEnterpriseDescription** | Pointer to **string** |  | [optional] 
**DeviceClassButton** | Pointer to **string** |  | [optional] 
**PersonalEula** | Pointer to **string** |  | [optional] 
**EnterpriseEula** | Pointer to **string** |  | [optional] 
**EulaButton** | Pointer to **string** |  | [optional] 
**SiteDescription** | Pointer to **string** |  | [optional] 
**CertificateText** | Pointer to **string** |  | [optional] 
**CertificateButton** | Pointer to **string** |  | [optional] 
**CertificateProfileName** | Pointer to **string** |  | [optional] 
**CertificateProfileDescription** | Pointer to **string** |  | [optional] 
**PersonalText** | Pointer to **string** |  | [optional] 
**PersonalButton** | Pointer to **string** |  | [optional] 
**PersonalProfileName** | Pointer to **string** |  | [optional] 
**PersonalProfileDescription** | Pointer to **string** |  | [optional] 
**UserEnrollmentText** | Pointer to **string** |  | [optional] [default to "Enter your Managed Apple id to install the MDM Profile."]
**UserEnrollmentButton** | Pointer to **string** |  | [optional] [default to "Continue"]
**UserEnrollmentProfileName** | Pointer to **string** |  | [optional] [default to "MDM Profile"]
**UserEnrollmentProfileDescription** | Pointer to **string** |  | [optional] [default to "MDM Profile for mobile device management"]
**EnterpriseText** | Pointer to **string** |  | [optional] 
**EnterpriseButton** | Pointer to **string** |  | [optional] 
**EnterpriseProfileName** | Pointer to **string** |  | [optional] 
**EnterpriseProfileDescription** | Pointer to **string** |  | [optional] 
**EnterprisePending** | Pointer to **string** |  | [optional] 
**QuickAddText** | Pointer to **string** |  | [optional] 
**QuickAddButton** | Pointer to **string** |  | [optional] 
**QuickAddName** | Pointer to **string** |  | [optional] [default to "QuickAdd.pkg"]
**QuickAddPending** | Pointer to **string** |  | [optional] 
**CompleteMessage** | Pointer to **string** |  | [optional] 
**FailedMessage** | Pointer to **string** |  | [optional] 
**TryAgainButton** | Pointer to **string** |  | [optional] 
**CheckNowButton** | Pointer to **string** |  | [optional] 
**CheckEnrollmentMessage** | Pointer to **string** |  | [optional] 
**LogoutButton** | Pointer to **string** |  | [optional] 

## Methods

### NewEnrollmentProcessTextObject

`func NewEnrollmentProcessTextObject() *EnrollmentProcessTextObject`

NewEnrollmentProcessTextObject instantiates a new EnrollmentProcessTextObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentProcessTextObjectWithDefaults

`func NewEnrollmentProcessTextObjectWithDefaults() *EnrollmentProcessTextObject`

NewEnrollmentProcessTextObjectWithDefaults instantiates a new EnrollmentProcessTextObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguageCode

`func (o *EnrollmentProcessTextObject) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *EnrollmentProcessTextObject) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *EnrollmentProcessTextObject) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *EnrollmentProcessTextObject) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### GetName

`func (o *EnrollmentProcessTextObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnrollmentProcessTextObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnrollmentProcessTextObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnrollmentProcessTextObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *EnrollmentProcessTextObject) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EnrollmentProcessTextObject) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EnrollmentProcessTextObject) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EnrollmentProcessTextObject) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetLoginDescription

`func (o *EnrollmentProcessTextObject) GetLoginDescription() string`

GetLoginDescription returns the LoginDescription field if non-nil, zero value otherwise.

### GetLoginDescriptionOk

`func (o *EnrollmentProcessTextObject) GetLoginDescriptionOk() (*string, bool)`

GetLoginDescriptionOk returns a tuple with the LoginDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginDescription

`func (o *EnrollmentProcessTextObject) SetLoginDescription(v string)`

SetLoginDescription sets LoginDescription field to given value.

### HasLoginDescription

`func (o *EnrollmentProcessTextObject) HasLoginDescription() bool`

HasLoginDescription returns a boolean if a field has been set.

### GetUsername

`func (o *EnrollmentProcessTextObject) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EnrollmentProcessTextObject) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EnrollmentProcessTextObject) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EnrollmentProcessTextObject) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *EnrollmentProcessTextObject) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EnrollmentProcessTextObject) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EnrollmentProcessTextObject) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EnrollmentProcessTextObject) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetLoginButton

`func (o *EnrollmentProcessTextObject) GetLoginButton() string`

GetLoginButton returns the LoginButton field if non-nil, zero value otherwise.

### GetLoginButtonOk

`func (o *EnrollmentProcessTextObject) GetLoginButtonOk() (*string, bool)`

GetLoginButtonOk returns a tuple with the LoginButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginButton

`func (o *EnrollmentProcessTextObject) SetLoginButton(v string)`

SetLoginButton sets LoginButton field to given value.

### HasLoginButton

`func (o *EnrollmentProcessTextObject) HasLoginButton() bool`

HasLoginButton returns a boolean if a field has been set.

### GetDeviceClassDescription

`func (o *EnrollmentProcessTextObject) GetDeviceClassDescription() string`

GetDeviceClassDescription returns the DeviceClassDescription field if non-nil, zero value otherwise.

### GetDeviceClassDescriptionOk

`func (o *EnrollmentProcessTextObject) GetDeviceClassDescriptionOk() (*string, bool)`

GetDeviceClassDescriptionOk returns a tuple with the DeviceClassDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClassDescription

`func (o *EnrollmentProcessTextObject) SetDeviceClassDescription(v string)`

SetDeviceClassDescription sets DeviceClassDescription field to given value.

### HasDeviceClassDescription

`func (o *EnrollmentProcessTextObject) HasDeviceClassDescription() bool`

HasDeviceClassDescription returns a boolean if a field has been set.

### GetDeviceClassPersonal

`func (o *EnrollmentProcessTextObject) GetDeviceClassPersonal() string`

GetDeviceClassPersonal returns the DeviceClassPersonal field if non-nil, zero value otherwise.

### GetDeviceClassPersonalOk

`func (o *EnrollmentProcessTextObject) GetDeviceClassPersonalOk() (*string, bool)`

GetDeviceClassPersonalOk returns a tuple with the DeviceClassPersonal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClassPersonal

`func (o *EnrollmentProcessTextObject) SetDeviceClassPersonal(v string)`

SetDeviceClassPersonal sets DeviceClassPersonal field to given value.

### HasDeviceClassPersonal

`func (o *EnrollmentProcessTextObject) HasDeviceClassPersonal() bool`

HasDeviceClassPersonal returns a boolean if a field has been set.

### GetDeviceClassPersonalDescription

`func (o *EnrollmentProcessTextObject) GetDeviceClassPersonalDescription() string`

GetDeviceClassPersonalDescription returns the DeviceClassPersonalDescription field if non-nil, zero value otherwise.

### GetDeviceClassPersonalDescriptionOk

`func (o *EnrollmentProcessTextObject) GetDeviceClassPersonalDescriptionOk() (*string, bool)`

GetDeviceClassPersonalDescriptionOk returns a tuple with the DeviceClassPersonalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClassPersonalDescription

`func (o *EnrollmentProcessTextObject) SetDeviceClassPersonalDescription(v string)`

SetDeviceClassPersonalDescription sets DeviceClassPersonalDescription field to given value.

### HasDeviceClassPersonalDescription

`func (o *EnrollmentProcessTextObject) HasDeviceClassPersonalDescription() bool`

HasDeviceClassPersonalDescription returns a boolean if a field has been set.

### GetDeviceClassEnterprise

`func (o *EnrollmentProcessTextObject) GetDeviceClassEnterprise() string`

GetDeviceClassEnterprise returns the DeviceClassEnterprise field if non-nil, zero value otherwise.

### GetDeviceClassEnterpriseOk

`func (o *EnrollmentProcessTextObject) GetDeviceClassEnterpriseOk() (*string, bool)`

GetDeviceClassEnterpriseOk returns a tuple with the DeviceClassEnterprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClassEnterprise

`func (o *EnrollmentProcessTextObject) SetDeviceClassEnterprise(v string)`

SetDeviceClassEnterprise sets DeviceClassEnterprise field to given value.

### HasDeviceClassEnterprise

`func (o *EnrollmentProcessTextObject) HasDeviceClassEnterprise() bool`

HasDeviceClassEnterprise returns a boolean if a field has been set.

### GetDeviceClassEnterpriseDescription

`func (o *EnrollmentProcessTextObject) GetDeviceClassEnterpriseDescription() string`

GetDeviceClassEnterpriseDescription returns the DeviceClassEnterpriseDescription field if non-nil, zero value otherwise.

### GetDeviceClassEnterpriseDescriptionOk

`func (o *EnrollmentProcessTextObject) GetDeviceClassEnterpriseDescriptionOk() (*string, bool)`

GetDeviceClassEnterpriseDescriptionOk returns a tuple with the DeviceClassEnterpriseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClassEnterpriseDescription

`func (o *EnrollmentProcessTextObject) SetDeviceClassEnterpriseDescription(v string)`

SetDeviceClassEnterpriseDescription sets DeviceClassEnterpriseDescription field to given value.

### HasDeviceClassEnterpriseDescription

`func (o *EnrollmentProcessTextObject) HasDeviceClassEnterpriseDescription() bool`

HasDeviceClassEnterpriseDescription returns a boolean if a field has been set.

### GetDeviceClassButton

`func (o *EnrollmentProcessTextObject) GetDeviceClassButton() string`

GetDeviceClassButton returns the DeviceClassButton field if non-nil, zero value otherwise.

### GetDeviceClassButtonOk

`func (o *EnrollmentProcessTextObject) GetDeviceClassButtonOk() (*string, bool)`

GetDeviceClassButtonOk returns a tuple with the DeviceClassButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClassButton

`func (o *EnrollmentProcessTextObject) SetDeviceClassButton(v string)`

SetDeviceClassButton sets DeviceClassButton field to given value.

### HasDeviceClassButton

`func (o *EnrollmentProcessTextObject) HasDeviceClassButton() bool`

HasDeviceClassButton returns a boolean if a field has been set.

### GetPersonalEula

`func (o *EnrollmentProcessTextObject) GetPersonalEula() string`

GetPersonalEula returns the PersonalEula field if non-nil, zero value otherwise.

### GetPersonalEulaOk

`func (o *EnrollmentProcessTextObject) GetPersonalEulaOk() (*string, bool)`

GetPersonalEulaOk returns a tuple with the PersonalEula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalEula

`func (o *EnrollmentProcessTextObject) SetPersonalEula(v string)`

SetPersonalEula sets PersonalEula field to given value.

### HasPersonalEula

`func (o *EnrollmentProcessTextObject) HasPersonalEula() bool`

HasPersonalEula returns a boolean if a field has been set.

### GetEnterpriseEula

`func (o *EnrollmentProcessTextObject) GetEnterpriseEula() string`

GetEnterpriseEula returns the EnterpriseEula field if non-nil, zero value otherwise.

### GetEnterpriseEulaOk

`func (o *EnrollmentProcessTextObject) GetEnterpriseEulaOk() (*string, bool)`

GetEnterpriseEulaOk returns a tuple with the EnterpriseEula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseEula

`func (o *EnrollmentProcessTextObject) SetEnterpriseEula(v string)`

SetEnterpriseEula sets EnterpriseEula field to given value.

### HasEnterpriseEula

`func (o *EnrollmentProcessTextObject) HasEnterpriseEula() bool`

HasEnterpriseEula returns a boolean if a field has been set.

### GetEulaButton

`func (o *EnrollmentProcessTextObject) GetEulaButton() string`

GetEulaButton returns the EulaButton field if non-nil, zero value otherwise.

### GetEulaButtonOk

`func (o *EnrollmentProcessTextObject) GetEulaButtonOk() (*string, bool)`

GetEulaButtonOk returns a tuple with the EulaButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaButton

`func (o *EnrollmentProcessTextObject) SetEulaButton(v string)`

SetEulaButton sets EulaButton field to given value.

### HasEulaButton

`func (o *EnrollmentProcessTextObject) HasEulaButton() bool`

HasEulaButton returns a boolean if a field has been set.

### GetSiteDescription

`func (o *EnrollmentProcessTextObject) GetSiteDescription() string`

GetSiteDescription returns the SiteDescription field if non-nil, zero value otherwise.

### GetSiteDescriptionOk

`func (o *EnrollmentProcessTextObject) GetSiteDescriptionOk() (*string, bool)`

GetSiteDescriptionOk returns a tuple with the SiteDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteDescription

`func (o *EnrollmentProcessTextObject) SetSiteDescription(v string)`

SetSiteDescription sets SiteDescription field to given value.

### HasSiteDescription

`func (o *EnrollmentProcessTextObject) HasSiteDescription() bool`

HasSiteDescription returns a boolean if a field has been set.

### GetCertificateText

`func (o *EnrollmentProcessTextObject) GetCertificateText() string`

GetCertificateText returns the CertificateText field if non-nil, zero value otherwise.

### GetCertificateTextOk

`func (o *EnrollmentProcessTextObject) GetCertificateTextOk() (*string, bool)`

GetCertificateTextOk returns a tuple with the CertificateText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateText

`func (o *EnrollmentProcessTextObject) SetCertificateText(v string)`

SetCertificateText sets CertificateText field to given value.

### HasCertificateText

`func (o *EnrollmentProcessTextObject) HasCertificateText() bool`

HasCertificateText returns a boolean if a field has been set.

### GetCertificateButton

`func (o *EnrollmentProcessTextObject) GetCertificateButton() string`

GetCertificateButton returns the CertificateButton field if non-nil, zero value otherwise.

### GetCertificateButtonOk

`func (o *EnrollmentProcessTextObject) GetCertificateButtonOk() (*string, bool)`

GetCertificateButtonOk returns a tuple with the CertificateButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateButton

`func (o *EnrollmentProcessTextObject) SetCertificateButton(v string)`

SetCertificateButton sets CertificateButton field to given value.

### HasCertificateButton

`func (o *EnrollmentProcessTextObject) HasCertificateButton() bool`

HasCertificateButton returns a boolean if a field has been set.

### GetCertificateProfileName

`func (o *EnrollmentProcessTextObject) GetCertificateProfileName() string`

GetCertificateProfileName returns the CertificateProfileName field if non-nil, zero value otherwise.

### GetCertificateProfileNameOk

`func (o *EnrollmentProcessTextObject) GetCertificateProfileNameOk() (*string, bool)`

GetCertificateProfileNameOk returns a tuple with the CertificateProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfileName

`func (o *EnrollmentProcessTextObject) SetCertificateProfileName(v string)`

SetCertificateProfileName sets CertificateProfileName field to given value.

### HasCertificateProfileName

`func (o *EnrollmentProcessTextObject) HasCertificateProfileName() bool`

HasCertificateProfileName returns a boolean if a field has been set.

### GetCertificateProfileDescription

`func (o *EnrollmentProcessTextObject) GetCertificateProfileDescription() string`

GetCertificateProfileDescription returns the CertificateProfileDescription field if non-nil, zero value otherwise.

### GetCertificateProfileDescriptionOk

`func (o *EnrollmentProcessTextObject) GetCertificateProfileDescriptionOk() (*string, bool)`

GetCertificateProfileDescriptionOk returns a tuple with the CertificateProfileDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfileDescription

`func (o *EnrollmentProcessTextObject) SetCertificateProfileDescription(v string)`

SetCertificateProfileDescription sets CertificateProfileDescription field to given value.

### HasCertificateProfileDescription

`func (o *EnrollmentProcessTextObject) HasCertificateProfileDescription() bool`

HasCertificateProfileDescription returns a boolean if a field has been set.

### GetPersonalText

`func (o *EnrollmentProcessTextObject) GetPersonalText() string`

GetPersonalText returns the PersonalText field if non-nil, zero value otherwise.

### GetPersonalTextOk

`func (o *EnrollmentProcessTextObject) GetPersonalTextOk() (*string, bool)`

GetPersonalTextOk returns a tuple with the PersonalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalText

`func (o *EnrollmentProcessTextObject) SetPersonalText(v string)`

SetPersonalText sets PersonalText field to given value.

### HasPersonalText

`func (o *EnrollmentProcessTextObject) HasPersonalText() bool`

HasPersonalText returns a boolean if a field has been set.

### GetPersonalButton

`func (o *EnrollmentProcessTextObject) GetPersonalButton() string`

GetPersonalButton returns the PersonalButton field if non-nil, zero value otherwise.

### GetPersonalButtonOk

`func (o *EnrollmentProcessTextObject) GetPersonalButtonOk() (*string, bool)`

GetPersonalButtonOk returns a tuple with the PersonalButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalButton

`func (o *EnrollmentProcessTextObject) SetPersonalButton(v string)`

SetPersonalButton sets PersonalButton field to given value.

### HasPersonalButton

`func (o *EnrollmentProcessTextObject) HasPersonalButton() bool`

HasPersonalButton returns a boolean if a field has been set.

### GetPersonalProfileName

`func (o *EnrollmentProcessTextObject) GetPersonalProfileName() string`

GetPersonalProfileName returns the PersonalProfileName field if non-nil, zero value otherwise.

### GetPersonalProfileNameOk

`func (o *EnrollmentProcessTextObject) GetPersonalProfileNameOk() (*string, bool)`

GetPersonalProfileNameOk returns a tuple with the PersonalProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalProfileName

`func (o *EnrollmentProcessTextObject) SetPersonalProfileName(v string)`

SetPersonalProfileName sets PersonalProfileName field to given value.

### HasPersonalProfileName

`func (o *EnrollmentProcessTextObject) HasPersonalProfileName() bool`

HasPersonalProfileName returns a boolean if a field has been set.

### GetPersonalProfileDescription

`func (o *EnrollmentProcessTextObject) GetPersonalProfileDescription() string`

GetPersonalProfileDescription returns the PersonalProfileDescription field if non-nil, zero value otherwise.

### GetPersonalProfileDescriptionOk

`func (o *EnrollmentProcessTextObject) GetPersonalProfileDescriptionOk() (*string, bool)`

GetPersonalProfileDescriptionOk returns a tuple with the PersonalProfileDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalProfileDescription

`func (o *EnrollmentProcessTextObject) SetPersonalProfileDescription(v string)`

SetPersonalProfileDescription sets PersonalProfileDescription field to given value.

### HasPersonalProfileDescription

`func (o *EnrollmentProcessTextObject) HasPersonalProfileDescription() bool`

HasPersonalProfileDescription returns a boolean if a field has been set.

### GetUserEnrollmentText

`func (o *EnrollmentProcessTextObject) GetUserEnrollmentText() string`

GetUserEnrollmentText returns the UserEnrollmentText field if non-nil, zero value otherwise.

### GetUserEnrollmentTextOk

`func (o *EnrollmentProcessTextObject) GetUserEnrollmentTextOk() (*string, bool)`

GetUserEnrollmentTextOk returns a tuple with the UserEnrollmentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEnrollmentText

`func (o *EnrollmentProcessTextObject) SetUserEnrollmentText(v string)`

SetUserEnrollmentText sets UserEnrollmentText field to given value.

### HasUserEnrollmentText

`func (o *EnrollmentProcessTextObject) HasUserEnrollmentText() bool`

HasUserEnrollmentText returns a boolean if a field has been set.

### GetUserEnrollmentButton

`func (o *EnrollmentProcessTextObject) GetUserEnrollmentButton() string`

GetUserEnrollmentButton returns the UserEnrollmentButton field if non-nil, zero value otherwise.

### GetUserEnrollmentButtonOk

`func (o *EnrollmentProcessTextObject) GetUserEnrollmentButtonOk() (*string, bool)`

GetUserEnrollmentButtonOk returns a tuple with the UserEnrollmentButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEnrollmentButton

`func (o *EnrollmentProcessTextObject) SetUserEnrollmentButton(v string)`

SetUserEnrollmentButton sets UserEnrollmentButton field to given value.

### HasUserEnrollmentButton

`func (o *EnrollmentProcessTextObject) HasUserEnrollmentButton() bool`

HasUserEnrollmentButton returns a boolean if a field has been set.

### GetUserEnrollmentProfileName

`func (o *EnrollmentProcessTextObject) GetUserEnrollmentProfileName() string`

GetUserEnrollmentProfileName returns the UserEnrollmentProfileName field if non-nil, zero value otherwise.

### GetUserEnrollmentProfileNameOk

`func (o *EnrollmentProcessTextObject) GetUserEnrollmentProfileNameOk() (*string, bool)`

GetUserEnrollmentProfileNameOk returns a tuple with the UserEnrollmentProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEnrollmentProfileName

`func (o *EnrollmentProcessTextObject) SetUserEnrollmentProfileName(v string)`

SetUserEnrollmentProfileName sets UserEnrollmentProfileName field to given value.

### HasUserEnrollmentProfileName

`func (o *EnrollmentProcessTextObject) HasUserEnrollmentProfileName() bool`

HasUserEnrollmentProfileName returns a boolean if a field has been set.

### GetUserEnrollmentProfileDescription

`func (o *EnrollmentProcessTextObject) GetUserEnrollmentProfileDescription() string`

GetUserEnrollmentProfileDescription returns the UserEnrollmentProfileDescription field if non-nil, zero value otherwise.

### GetUserEnrollmentProfileDescriptionOk

`func (o *EnrollmentProcessTextObject) GetUserEnrollmentProfileDescriptionOk() (*string, bool)`

GetUserEnrollmentProfileDescriptionOk returns a tuple with the UserEnrollmentProfileDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEnrollmentProfileDescription

`func (o *EnrollmentProcessTextObject) SetUserEnrollmentProfileDescription(v string)`

SetUserEnrollmentProfileDescription sets UserEnrollmentProfileDescription field to given value.

### HasUserEnrollmentProfileDescription

`func (o *EnrollmentProcessTextObject) HasUserEnrollmentProfileDescription() bool`

HasUserEnrollmentProfileDescription returns a boolean if a field has been set.

### GetEnterpriseText

`func (o *EnrollmentProcessTextObject) GetEnterpriseText() string`

GetEnterpriseText returns the EnterpriseText field if non-nil, zero value otherwise.

### GetEnterpriseTextOk

`func (o *EnrollmentProcessTextObject) GetEnterpriseTextOk() (*string, bool)`

GetEnterpriseTextOk returns a tuple with the EnterpriseText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseText

`func (o *EnrollmentProcessTextObject) SetEnterpriseText(v string)`

SetEnterpriseText sets EnterpriseText field to given value.

### HasEnterpriseText

`func (o *EnrollmentProcessTextObject) HasEnterpriseText() bool`

HasEnterpriseText returns a boolean if a field has been set.

### GetEnterpriseButton

`func (o *EnrollmentProcessTextObject) GetEnterpriseButton() string`

GetEnterpriseButton returns the EnterpriseButton field if non-nil, zero value otherwise.

### GetEnterpriseButtonOk

`func (o *EnrollmentProcessTextObject) GetEnterpriseButtonOk() (*string, bool)`

GetEnterpriseButtonOk returns a tuple with the EnterpriseButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseButton

`func (o *EnrollmentProcessTextObject) SetEnterpriseButton(v string)`

SetEnterpriseButton sets EnterpriseButton field to given value.

### HasEnterpriseButton

`func (o *EnrollmentProcessTextObject) HasEnterpriseButton() bool`

HasEnterpriseButton returns a boolean if a field has been set.

### GetEnterpriseProfileName

`func (o *EnrollmentProcessTextObject) GetEnterpriseProfileName() string`

GetEnterpriseProfileName returns the EnterpriseProfileName field if non-nil, zero value otherwise.

### GetEnterpriseProfileNameOk

`func (o *EnrollmentProcessTextObject) GetEnterpriseProfileNameOk() (*string, bool)`

GetEnterpriseProfileNameOk returns a tuple with the EnterpriseProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseProfileName

`func (o *EnrollmentProcessTextObject) SetEnterpriseProfileName(v string)`

SetEnterpriseProfileName sets EnterpriseProfileName field to given value.

### HasEnterpriseProfileName

`func (o *EnrollmentProcessTextObject) HasEnterpriseProfileName() bool`

HasEnterpriseProfileName returns a boolean if a field has been set.

### GetEnterpriseProfileDescription

`func (o *EnrollmentProcessTextObject) GetEnterpriseProfileDescription() string`

GetEnterpriseProfileDescription returns the EnterpriseProfileDescription field if non-nil, zero value otherwise.

### GetEnterpriseProfileDescriptionOk

`func (o *EnrollmentProcessTextObject) GetEnterpriseProfileDescriptionOk() (*string, bool)`

GetEnterpriseProfileDescriptionOk returns a tuple with the EnterpriseProfileDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseProfileDescription

`func (o *EnrollmentProcessTextObject) SetEnterpriseProfileDescription(v string)`

SetEnterpriseProfileDescription sets EnterpriseProfileDescription field to given value.

### HasEnterpriseProfileDescription

`func (o *EnrollmentProcessTextObject) HasEnterpriseProfileDescription() bool`

HasEnterpriseProfileDescription returns a boolean if a field has been set.

### GetEnterprisePending

`func (o *EnrollmentProcessTextObject) GetEnterprisePending() string`

GetEnterprisePending returns the EnterprisePending field if non-nil, zero value otherwise.

### GetEnterprisePendingOk

`func (o *EnrollmentProcessTextObject) GetEnterprisePendingOk() (*string, bool)`

GetEnterprisePendingOk returns a tuple with the EnterprisePending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterprisePending

`func (o *EnrollmentProcessTextObject) SetEnterprisePending(v string)`

SetEnterprisePending sets EnterprisePending field to given value.

### HasEnterprisePending

`func (o *EnrollmentProcessTextObject) HasEnterprisePending() bool`

HasEnterprisePending returns a boolean if a field has been set.

### GetQuickAddText

`func (o *EnrollmentProcessTextObject) GetQuickAddText() string`

GetQuickAddText returns the QuickAddText field if non-nil, zero value otherwise.

### GetQuickAddTextOk

`func (o *EnrollmentProcessTextObject) GetQuickAddTextOk() (*string, bool)`

GetQuickAddTextOk returns a tuple with the QuickAddText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickAddText

`func (o *EnrollmentProcessTextObject) SetQuickAddText(v string)`

SetQuickAddText sets QuickAddText field to given value.

### HasQuickAddText

`func (o *EnrollmentProcessTextObject) HasQuickAddText() bool`

HasQuickAddText returns a boolean if a field has been set.

### GetQuickAddButton

`func (o *EnrollmentProcessTextObject) GetQuickAddButton() string`

GetQuickAddButton returns the QuickAddButton field if non-nil, zero value otherwise.

### GetQuickAddButtonOk

`func (o *EnrollmentProcessTextObject) GetQuickAddButtonOk() (*string, bool)`

GetQuickAddButtonOk returns a tuple with the QuickAddButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickAddButton

`func (o *EnrollmentProcessTextObject) SetQuickAddButton(v string)`

SetQuickAddButton sets QuickAddButton field to given value.

### HasQuickAddButton

`func (o *EnrollmentProcessTextObject) HasQuickAddButton() bool`

HasQuickAddButton returns a boolean if a field has been set.

### GetQuickAddName

`func (o *EnrollmentProcessTextObject) GetQuickAddName() string`

GetQuickAddName returns the QuickAddName field if non-nil, zero value otherwise.

### GetQuickAddNameOk

`func (o *EnrollmentProcessTextObject) GetQuickAddNameOk() (*string, bool)`

GetQuickAddNameOk returns a tuple with the QuickAddName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickAddName

`func (o *EnrollmentProcessTextObject) SetQuickAddName(v string)`

SetQuickAddName sets QuickAddName field to given value.

### HasQuickAddName

`func (o *EnrollmentProcessTextObject) HasQuickAddName() bool`

HasQuickAddName returns a boolean if a field has been set.

### GetQuickAddPending

`func (o *EnrollmentProcessTextObject) GetQuickAddPending() string`

GetQuickAddPending returns the QuickAddPending field if non-nil, zero value otherwise.

### GetQuickAddPendingOk

`func (o *EnrollmentProcessTextObject) GetQuickAddPendingOk() (*string, bool)`

GetQuickAddPendingOk returns a tuple with the QuickAddPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickAddPending

`func (o *EnrollmentProcessTextObject) SetQuickAddPending(v string)`

SetQuickAddPending sets QuickAddPending field to given value.

### HasQuickAddPending

`func (o *EnrollmentProcessTextObject) HasQuickAddPending() bool`

HasQuickAddPending returns a boolean if a field has been set.

### GetCompleteMessage

`func (o *EnrollmentProcessTextObject) GetCompleteMessage() string`

GetCompleteMessage returns the CompleteMessage field if non-nil, zero value otherwise.

### GetCompleteMessageOk

`func (o *EnrollmentProcessTextObject) GetCompleteMessageOk() (*string, bool)`

GetCompleteMessageOk returns a tuple with the CompleteMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteMessage

`func (o *EnrollmentProcessTextObject) SetCompleteMessage(v string)`

SetCompleteMessage sets CompleteMessage field to given value.

### HasCompleteMessage

`func (o *EnrollmentProcessTextObject) HasCompleteMessage() bool`

HasCompleteMessage returns a boolean if a field has been set.

### GetFailedMessage

`func (o *EnrollmentProcessTextObject) GetFailedMessage() string`

GetFailedMessage returns the FailedMessage field if non-nil, zero value otherwise.

### GetFailedMessageOk

`func (o *EnrollmentProcessTextObject) GetFailedMessageOk() (*string, bool)`

GetFailedMessageOk returns a tuple with the FailedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedMessage

`func (o *EnrollmentProcessTextObject) SetFailedMessage(v string)`

SetFailedMessage sets FailedMessage field to given value.

### HasFailedMessage

`func (o *EnrollmentProcessTextObject) HasFailedMessage() bool`

HasFailedMessage returns a boolean if a field has been set.

### GetTryAgainButton

`func (o *EnrollmentProcessTextObject) GetTryAgainButton() string`

GetTryAgainButton returns the TryAgainButton field if non-nil, zero value otherwise.

### GetTryAgainButtonOk

`func (o *EnrollmentProcessTextObject) GetTryAgainButtonOk() (*string, bool)`

GetTryAgainButtonOk returns a tuple with the TryAgainButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryAgainButton

`func (o *EnrollmentProcessTextObject) SetTryAgainButton(v string)`

SetTryAgainButton sets TryAgainButton field to given value.

### HasTryAgainButton

`func (o *EnrollmentProcessTextObject) HasTryAgainButton() bool`

HasTryAgainButton returns a boolean if a field has been set.

### GetCheckNowButton

`func (o *EnrollmentProcessTextObject) GetCheckNowButton() string`

GetCheckNowButton returns the CheckNowButton field if non-nil, zero value otherwise.

### GetCheckNowButtonOk

`func (o *EnrollmentProcessTextObject) GetCheckNowButtonOk() (*string, bool)`

GetCheckNowButtonOk returns a tuple with the CheckNowButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNowButton

`func (o *EnrollmentProcessTextObject) SetCheckNowButton(v string)`

SetCheckNowButton sets CheckNowButton field to given value.

### HasCheckNowButton

`func (o *EnrollmentProcessTextObject) HasCheckNowButton() bool`

HasCheckNowButton returns a boolean if a field has been set.

### GetCheckEnrollmentMessage

`func (o *EnrollmentProcessTextObject) GetCheckEnrollmentMessage() string`

GetCheckEnrollmentMessage returns the CheckEnrollmentMessage field if non-nil, zero value otherwise.

### GetCheckEnrollmentMessageOk

`func (o *EnrollmentProcessTextObject) GetCheckEnrollmentMessageOk() (*string, bool)`

GetCheckEnrollmentMessageOk returns a tuple with the CheckEnrollmentMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckEnrollmentMessage

`func (o *EnrollmentProcessTextObject) SetCheckEnrollmentMessage(v string)`

SetCheckEnrollmentMessage sets CheckEnrollmentMessage field to given value.

### HasCheckEnrollmentMessage

`func (o *EnrollmentProcessTextObject) HasCheckEnrollmentMessage() bool`

HasCheckEnrollmentMessage returns a boolean if a field has been set.

### GetLogoutButton

`func (o *EnrollmentProcessTextObject) GetLogoutButton() string`

GetLogoutButton returns the LogoutButton field if non-nil, zero value otherwise.

### GetLogoutButtonOk

`func (o *EnrollmentProcessTextObject) GetLogoutButtonOk() (*string, bool)`

GetLogoutButtonOk returns a tuple with the LogoutButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutButton

`func (o *EnrollmentProcessTextObject) SetLogoutButton(v string)`

SetLogoutButton sets LogoutButton field to given value.

### HasLogoutButton

`func (o *EnrollmentProcessTextObject) HasLogoutButton() bool`

HasLogoutButton returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


