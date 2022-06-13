# DeviceEnrollmentToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenFileName** | Pointer to **string** | Optional name of the token to be saved, if no name is provided one will be auto-generated | [optional] 
**EncodedToken** | Pointer to **string** | The base 64 encoded token | [optional] 

## Methods

### NewDeviceEnrollmentToken

`func NewDeviceEnrollmentToken() *DeviceEnrollmentToken`

NewDeviceEnrollmentToken instantiates a new DeviceEnrollmentToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceEnrollmentTokenWithDefaults

`func NewDeviceEnrollmentTokenWithDefaults() *DeviceEnrollmentToken`

NewDeviceEnrollmentTokenWithDefaults instantiates a new DeviceEnrollmentToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenFileName

`func (o *DeviceEnrollmentToken) GetTokenFileName() string`

GetTokenFileName returns the TokenFileName field if non-nil, zero value otherwise.

### GetTokenFileNameOk

`func (o *DeviceEnrollmentToken) GetTokenFileNameOk() (*string, bool)`

GetTokenFileNameOk returns a tuple with the TokenFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenFileName

`func (o *DeviceEnrollmentToken) SetTokenFileName(v string)`

SetTokenFileName sets TokenFileName field to given value.

### HasTokenFileName

`func (o *DeviceEnrollmentToken) HasTokenFileName() bool`

HasTokenFileName returns a boolean if a field has been set.

### GetEncodedToken

`func (o *DeviceEnrollmentToken) GetEncodedToken() string`

GetEncodedToken returns the EncodedToken field if non-nil, zero value otherwise.

### GetEncodedTokenOk

`func (o *DeviceEnrollmentToken) GetEncodedTokenOk() (*string, bool)`

GetEncodedTokenOk returns a tuple with the EncodedToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedToken

`func (o *DeviceEnrollmentToken) SetEncodedToken(v string)`

SetEncodedToken sets EncodedToken field to given value.

### HasEncodedToken

`func (o *DeviceEnrollmentToken) HasEncodedToken() bool`

HasEncodedToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


