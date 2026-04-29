# ApnsPushEnableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedTime** | Pointer to **time.Time** | Timestamp when the request was created (ISO-8601 format) | [optional] 
**Status** | Pointer to **string** | Current status of the request | [optional] 
**ProcessedTime** | Pointer to **NullableTime** | Timestamp when the request was processed (ISO-8601 format), null if not yet processed | [optional] 

## Methods

### NewApnsPushEnableRequest

`func NewApnsPushEnableRequest() *ApnsPushEnableRequest`

NewApnsPushEnableRequest instantiates a new ApnsPushEnableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApnsPushEnableRequestWithDefaults

`func NewApnsPushEnableRequestWithDefaults() *ApnsPushEnableRequest`

NewApnsPushEnableRequestWithDefaults instantiates a new ApnsPushEnableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedTime

`func (o *ApnsPushEnableRequest) GetRequestedTime() time.Time`

GetRequestedTime returns the RequestedTime field if non-nil, zero value otherwise.

### GetRequestedTimeOk

`func (o *ApnsPushEnableRequest) GetRequestedTimeOk() (*time.Time, bool)`

GetRequestedTimeOk returns a tuple with the RequestedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTime

`func (o *ApnsPushEnableRequest) SetRequestedTime(v time.Time)`

SetRequestedTime sets RequestedTime field to given value.

### HasRequestedTime

`func (o *ApnsPushEnableRequest) HasRequestedTime() bool`

HasRequestedTime returns a boolean if a field has been set.

### GetStatus

`func (o *ApnsPushEnableRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApnsPushEnableRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApnsPushEnableRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApnsPushEnableRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProcessedTime

`func (o *ApnsPushEnableRequest) GetProcessedTime() time.Time`

GetProcessedTime returns the ProcessedTime field if non-nil, zero value otherwise.

### GetProcessedTimeOk

`func (o *ApnsPushEnableRequest) GetProcessedTimeOk() (*time.Time, bool)`

GetProcessedTimeOk returns a tuple with the ProcessedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedTime

`func (o *ApnsPushEnableRequest) SetProcessedTime(v time.Time)`

SetProcessedTime sets ProcessedTime field to given value.

### HasProcessedTime

`func (o *ApnsPushEnableRequest) HasProcessedTime() bool`

HasProcessedTime returns a boolean if a field has been set.

### SetProcessedTimeNil

`func (o *ApnsPushEnableRequest) SetProcessedTimeNil(b bool)`

 SetProcessedTimeNil sets the value for ProcessedTime to be an explicit nil

### UnsetProcessedTime
`func (o *ApnsPushEnableRequest) UnsetProcessedTime()`

UnsetProcessedTime ensures that no value is present for ProcessedTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


