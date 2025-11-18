# AdcsCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** | Server certificate filename should extend .cer or .pem, and client certificate filename should extend .p12 or .pfx. | 
**Data** | **[]string** | Must be base-64 encoded data obtainable by &#x60;openssl base64 &lt; /file/path/filename.pfx | tr -d &#39;\\n&#39; | pbcopy&#x60; in linux terminal, or similar parsing methods. | 
**Password** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAdcsCertificate

`func NewAdcsCertificate(filename string, data []string, ) *AdcsCertificate`

NewAdcsCertificate instantiates a new AdcsCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdcsCertificateWithDefaults

`func NewAdcsCertificateWithDefaults() *AdcsCertificate`

NewAdcsCertificateWithDefaults instantiates a new AdcsCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *AdcsCertificate) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *AdcsCertificate) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *AdcsCertificate) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetData

`func (o *AdcsCertificate) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AdcsCertificate) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AdcsCertificate) SetData(v []string)`

SetData sets Data field to given value.


### GetPassword

`func (o *AdcsCertificate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AdcsCertificate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AdcsCertificate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AdcsCertificate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *AdcsCertificate) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *AdcsCertificate) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


