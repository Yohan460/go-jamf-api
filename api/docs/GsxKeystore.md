# GsxKeystore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [default to " "]
**ExpirationEpoch** | Pointer to **int64** |  | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** |  | [optional] [readonly] 
**KeystoreBytes** | Pointer to **string** | The base 64 encoded of the GSX Connection keystore. | [optional] 
**KeystorePassword** | **string** |  | 

## Methods

### NewGsxKeystore

`func NewGsxKeystore(name string, keystorePassword string, ) *GsxKeystore`

NewGsxKeystore instantiates a new GsxKeystore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGsxKeystoreWithDefaults

`func NewGsxKeystoreWithDefaults() *GsxKeystore`

NewGsxKeystoreWithDefaults instantiates a new GsxKeystore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GsxKeystore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GsxKeystore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GsxKeystore) SetName(v string)`

SetName sets Name field to given value.


### GetExpirationEpoch

`func (o *GsxKeystore) GetExpirationEpoch() int64`

GetExpirationEpoch returns the ExpirationEpoch field if non-nil, zero value otherwise.

### GetExpirationEpochOk

`func (o *GsxKeystore) GetExpirationEpochOk() (*int64, bool)`

GetExpirationEpochOk returns a tuple with the ExpirationEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationEpoch

`func (o *GsxKeystore) SetExpirationEpoch(v int64)`

SetExpirationEpoch sets ExpirationEpoch field to given value.

### HasExpirationEpoch

`func (o *GsxKeystore) HasExpirationEpoch() bool`

HasExpirationEpoch returns a boolean if a field has been set.

### GetErrorMessage

`func (o *GsxKeystore) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GsxKeystore) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GsxKeystore) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GsxKeystore) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetKeystoreBytes

`func (o *GsxKeystore) GetKeystoreBytes() string`

GetKeystoreBytes returns the KeystoreBytes field if non-nil, zero value otherwise.

### GetKeystoreBytesOk

`func (o *GsxKeystore) GetKeystoreBytesOk() (*string, bool)`

GetKeystoreBytesOk returns a tuple with the KeystoreBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoreBytes

`func (o *GsxKeystore) SetKeystoreBytes(v string)`

SetKeystoreBytes sets KeystoreBytes field to given value.

### HasKeystoreBytes

`func (o *GsxKeystore) HasKeystoreBytes() bool`

HasKeystoreBytes returns a boolean if a field has been set.

### GetKeystorePassword

`func (o *GsxKeystore) GetKeystorePassword() string`

GetKeystorePassword returns the KeystorePassword field if non-nil, zero value otherwise.

### GetKeystorePasswordOk

`func (o *GsxKeystore) GetKeystorePasswordOk() (*string, bool)`

GetKeystorePasswordOk returns a tuple with the KeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystorePassword

`func (o *GsxKeystore) SetKeystorePassword(v string)`

SetKeystorePassword sets KeystorePassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


