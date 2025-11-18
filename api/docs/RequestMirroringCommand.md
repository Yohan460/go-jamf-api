# RequestMirroringCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandType** | [**MdmCommandType**](MdmCommandType.md) |  | 
**DestinationDeviceId** | Pointer to **string** | The hardware address of the destination device to which the screen will be mirrored. This value isn’t case-sensitive. Required if destinationName is not provided. | [optional] 
**DestinationName** | Pointer to **string** | The name of the destination device to which the screen will be mirrored. Required if destinationDeviceId is not provided. | [optional] 
**Password** | Pointer to **NullableString** | The screen-sharing password used to connect to the destination device. | [optional] 
**ScanTime** | Pointer to **NullableInt64** | The scan time which device spends in seconds to find the destination device. | [optional] 

## Methods

### NewRequestMirroringCommand

`func NewRequestMirroringCommand(commandType MdmCommandType, ) *RequestMirroringCommand`

NewRequestMirroringCommand instantiates a new RequestMirroringCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestMirroringCommandWithDefaults

`func NewRequestMirroringCommandWithDefaults() *RequestMirroringCommand`

NewRequestMirroringCommandWithDefaults instantiates a new RequestMirroringCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandType

`func (o *RequestMirroringCommand) GetCommandType() MdmCommandType`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *RequestMirroringCommand) GetCommandTypeOk() (*MdmCommandType, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *RequestMirroringCommand) SetCommandType(v MdmCommandType)`

SetCommandType sets CommandType field to given value.


### GetDestinationDeviceId

`func (o *RequestMirroringCommand) GetDestinationDeviceId() string`

GetDestinationDeviceId returns the DestinationDeviceId field if non-nil, zero value otherwise.

### GetDestinationDeviceIdOk

`func (o *RequestMirroringCommand) GetDestinationDeviceIdOk() (*string, bool)`

GetDestinationDeviceIdOk returns a tuple with the DestinationDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDeviceId

`func (o *RequestMirroringCommand) SetDestinationDeviceId(v string)`

SetDestinationDeviceId sets DestinationDeviceId field to given value.

### HasDestinationDeviceId

`func (o *RequestMirroringCommand) HasDestinationDeviceId() bool`

HasDestinationDeviceId returns a boolean if a field has been set.

### GetDestinationName

`func (o *RequestMirroringCommand) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *RequestMirroringCommand) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *RequestMirroringCommand) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.

### HasDestinationName

`func (o *RequestMirroringCommand) HasDestinationName() bool`

HasDestinationName returns a boolean if a field has been set.

### GetPassword

`func (o *RequestMirroringCommand) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RequestMirroringCommand) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RequestMirroringCommand) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RequestMirroringCommand) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *RequestMirroringCommand) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *RequestMirroringCommand) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetScanTime

`func (o *RequestMirroringCommand) GetScanTime() int64`

GetScanTime returns the ScanTime field if non-nil, zero value otherwise.

### GetScanTimeOk

`func (o *RequestMirroringCommand) GetScanTimeOk() (*int64, bool)`

GetScanTimeOk returns a tuple with the ScanTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanTime

`func (o *RequestMirroringCommand) SetScanTime(v int64)`

SetScanTime sets ScanTime field to given value.

### HasScanTime

`func (o *RequestMirroringCommand) HasScanTime() bool`

HasScanTime returns a boolean if a field has been set.

### SetScanTimeNil

`func (o *RequestMirroringCommand) SetScanTimeNil(b bool)`

 SetScanTimeNil sets the value for ScanTime to be an explicit nil

### UnsetScanTime
`func (o *RequestMirroringCommand) UnsetScanTime()`

UnsetScanTime ensures that no value is present for ScanTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


