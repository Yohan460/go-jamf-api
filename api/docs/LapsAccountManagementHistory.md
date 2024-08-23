# LapsAccountManagementHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**EventTime** | Pointer to **NullableTime** |  | [optional] 
**ViewedBy** | Pointer to **NullableString** |  | [optional] 
**UserSource** | Pointer to **string** |  | [optional] 

## Methods

### NewLapsAccountManagementHistory

`func NewLapsAccountManagementHistory() *LapsAccountManagementHistory`

NewLapsAccountManagementHistory instantiates a new LapsAccountManagementHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLapsAccountManagementHistoryWithDefaults

`func NewLapsAccountManagementHistoryWithDefaults() *LapsAccountManagementHistory`

NewLapsAccountManagementHistoryWithDefaults instantiates a new LapsAccountManagementHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *LapsAccountManagementHistory) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *LapsAccountManagementHistory) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *LapsAccountManagementHistory) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *LapsAccountManagementHistory) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEventType

`func (o *LapsAccountManagementHistory) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *LapsAccountManagementHistory) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *LapsAccountManagementHistory) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *LapsAccountManagementHistory) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventTime

`func (o *LapsAccountManagementHistory) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *LapsAccountManagementHistory) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *LapsAccountManagementHistory) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *LapsAccountManagementHistory) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### SetEventTimeNil

`func (o *LapsAccountManagementHistory) SetEventTimeNil(b bool)`

 SetEventTimeNil sets the value for EventTime to be an explicit nil

### UnsetEventTime
`func (o *LapsAccountManagementHistory) UnsetEventTime()`

UnsetEventTime ensures that no value is present for EventTime, not even an explicit nil
### GetViewedBy

`func (o *LapsAccountManagementHistory) GetViewedBy() string`

GetViewedBy returns the ViewedBy field if non-nil, zero value otherwise.

### GetViewedByOk

`func (o *LapsAccountManagementHistory) GetViewedByOk() (*string, bool)`

GetViewedByOk returns a tuple with the ViewedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedBy

`func (o *LapsAccountManagementHistory) SetViewedBy(v string)`

SetViewedBy sets ViewedBy field to given value.

### HasViewedBy

`func (o *LapsAccountManagementHistory) HasViewedBy() bool`

HasViewedBy returns a boolean if a field has been set.

### SetViewedByNil

`func (o *LapsAccountManagementHistory) SetViewedByNil(b bool)`

 SetViewedByNil sets the value for ViewedBy to be an explicit nil

### UnsetViewedBy
`func (o *LapsAccountManagementHistory) UnsetViewedBy()`

UnsetViewedBy ensures that no value is present for ViewedBy, not even an explicit nil
### GetUserSource

`func (o *LapsAccountManagementHistory) GetUserSource() string`

GetUserSource returns the UserSource field if non-nil, zero value otherwise.

### GetUserSourceOk

`func (o *LapsAccountManagementHistory) GetUserSourceOk() (*string, bool)`

GetUserSourceOk returns a tuple with the UserSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSource

`func (o *LapsAccountManagementHistory) SetUserSource(v string)`

SetUserSource sets UserSource field to given value.

### HasUserSource

`func (o *LapsAccountManagementHistory) HasUserSource() bool`

HasUserSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


