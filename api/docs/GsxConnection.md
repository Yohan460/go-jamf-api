# GsxConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | [default to false]
**Username** | **string** |  | [default to " "]
**ServiceAccountNo** | **string** |  | 
**ShipToNo** | Pointer to **string** |  | [optional] 
**Token** | **string** |  | 
**GsxKeystore** | [**GsxKeystore**](GsxKeystore.md) |  | 

## Methods

### NewGsxConnection

`func NewGsxConnection(enabled bool, username string, serviceAccountNo string, token string, gsxKeystore GsxKeystore, ) *GsxConnection`

NewGsxConnection instantiates a new GsxConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGsxConnectionWithDefaults

`func NewGsxConnectionWithDefaults() *GsxConnection`

NewGsxConnectionWithDefaults instantiates a new GsxConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GsxConnection) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GsxConnection) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GsxConnection) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetUsername

`func (o *GsxConnection) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GsxConnection) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GsxConnection) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetServiceAccountNo

`func (o *GsxConnection) GetServiceAccountNo() string`

GetServiceAccountNo returns the ServiceAccountNo field if non-nil, zero value otherwise.

### GetServiceAccountNoOk

`func (o *GsxConnection) GetServiceAccountNoOk() (*string, bool)`

GetServiceAccountNoOk returns a tuple with the ServiceAccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountNo

`func (o *GsxConnection) SetServiceAccountNo(v string)`

SetServiceAccountNo sets ServiceAccountNo field to given value.


### GetShipToNo

`func (o *GsxConnection) GetShipToNo() string`

GetShipToNo returns the ShipToNo field if non-nil, zero value otherwise.

### GetShipToNoOk

`func (o *GsxConnection) GetShipToNoOk() (*string, bool)`

GetShipToNoOk returns a tuple with the ShipToNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToNo

`func (o *GsxConnection) SetShipToNo(v string)`

SetShipToNo sets ShipToNo field to given value.

### HasShipToNo

`func (o *GsxConnection) HasShipToNo() bool`

HasShipToNo returns a boolean if a field has been set.

### GetToken

`func (o *GsxConnection) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GsxConnection) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GsxConnection) SetToken(v string)`

SetToken sets Token field to given value.


### GetGsxKeystore

`func (o *GsxConnection) GetGsxKeystore() GsxKeystore`

GetGsxKeystore returns the GsxKeystore field if non-nil, zero value otherwise.

### GetGsxKeystoreOk

`func (o *GsxConnection) GetGsxKeystoreOk() (*GsxKeystore, bool)`

GetGsxKeystoreOk returns a tuple with the GsxKeystore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGsxKeystore

`func (o *GsxConnection) SetGsxKeystore(v GsxKeystore)`

SetGsxKeystore sets GsxKeystore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


