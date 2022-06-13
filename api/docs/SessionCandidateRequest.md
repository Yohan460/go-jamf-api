# SessionCandidateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **string** | Device identifier | 
**DeviceType** | **string** | Device type | 
**Description** | **string** | Session description. To be used for additional context on the reason of the session | 

## Methods

### NewSessionCandidateRequest

`func NewSessionCandidateRequest(deviceId string, deviceType string, description string, ) *SessionCandidateRequest`

NewSessionCandidateRequest instantiates a new SessionCandidateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionCandidateRequestWithDefaults

`func NewSessionCandidateRequestWithDefaults() *SessionCandidateRequest`

NewSessionCandidateRequestWithDefaults instantiates a new SessionCandidateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *SessionCandidateRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SessionCandidateRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SessionCandidateRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetDeviceType

`func (o *SessionCandidateRequest) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *SessionCandidateRequest) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *SessionCandidateRequest) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.


### GetDescription

`func (o *SessionCandidateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SessionCandidateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SessionCandidateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


