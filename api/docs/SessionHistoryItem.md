# SessionHistoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **string** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**SessionStartedTimestamp** | Pointer to **time.Time** |  | [optional] 
**SessionEndedTimestamp** | Pointer to **time.Time** |  | [optional] 
**SessionType** | Pointer to **string** |  | [optional] 
**StatusType** | Pointer to **string** |  | [optional] 
**SessionAdminId** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewSessionHistoryItem

`func NewSessionHistoryItem() *SessionHistoryItem`

NewSessionHistoryItem instantiates a new SessionHistoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionHistoryItemWithDefaults

`func NewSessionHistoryItemWithDefaults() *SessionHistoryItem`

NewSessionHistoryItemWithDefaults instantiates a new SessionHistoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *SessionHistoryItem) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SessionHistoryItem) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SessionHistoryItem) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SessionHistoryItem) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetSessionId

`func (o *SessionHistoryItem) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *SessionHistoryItem) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *SessionHistoryItem) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *SessionHistoryItem) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetDeviceId

`func (o *SessionHistoryItem) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SessionHistoryItem) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SessionHistoryItem) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *SessionHistoryItem) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetSessionStartedTimestamp

`func (o *SessionHistoryItem) GetSessionStartedTimestamp() time.Time`

GetSessionStartedTimestamp returns the SessionStartedTimestamp field if non-nil, zero value otherwise.

### GetSessionStartedTimestampOk

`func (o *SessionHistoryItem) GetSessionStartedTimestampOk() (*time.Time, bool)`

GetSessionStartedTimestampOk returns a tuple with the SessionStartedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStartedTimestamp

`func (o *SessionHistoryItem) SetSessionStartedTimestamp(v time.Time)`

SetSessionStartedTimestamp sets SessionStartedTimestamp field to given value.

### HasSessionStartedTimestamp

`func (o *SessionHistoryItem) HasSessionStartedTimestamp() bool`

HasSessionStartedTimestamp returns a boolean if a field has been set.

### GetSessionEndedTimestamp

`func (o *SessionHistoryItem) GetSessionEndedTimestamp() time.Time`

GetSessionEndedTimestamp returns the SessionEndedTimestamp field if non-nil, zero value otherwise.

### GetSessionEndedTimestampOk

`func (o *SessionHistoryItem) GetSessionEndedTimestampOk() (*time.Time, bool)`

GetSessionEndedTimestampOk returns a tuple with the SessionEndedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionEndedTimestamp

`func (o *SessionHistoryItem) SetSessionEndedTimestamp(v time.Time)`

SetSessionEndedTimestamp sets SessionEndedTimestamp field to given value.

### HasSessionEndedTimestamp

`func (o *SessionHistoryItem) HasSessionEndedTimestamp() bool`

HasSessionEndedTimestamp returns a boolean if a field has been set.

### GetSessionType

`func (o *SessionHistoryItem) GetSessionType() string`

GetSessionType returns the SessionType field if non-nil, zero value otherwise.

### GetSessionTypeOk

`func (o *SessionHistoryItem) GetSessionTypeOk() (*string, bool)`

GetSessionTypeOk returns a tuple with the SessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionType

`func (o *SessionHistoryItem) SetSessionType(v string)`

SetSessionType sets SessionType field to given value.

### HasSessionType

`func (o *SessionHistoryItem) HasSessionType() bool`

HasSessionType returns a boolean if a field has been set.

### GetStatusType

`func (o *SessionHistoryItem) GetStatusType() string`

GetStatusType returns the StatusType field if non-nil, zero value otherwise.

### GetStatusTypeOk

`func (o *SessionHistoryItem) GetStatusTypeOk() (*string, bool)`

GetStatusTypeOk returns a tuple with the StatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusType

`func (o *SessionHistoryItem) SetStatusType(v string)`

SetStatusType sets StatusType field to given value.

### HasStatusType

`func (o *SessionHistoryItem) HasStatusType() bool`

HasStatusType returns a boolean if a field has been set.

### GetSessionAdminId

`func (o *SessionHistoryItem) GetSessionAdminId() string`

GetSessionAdminId returns the SessionAdminId field if non-nil, zero value otherwise.

### GetSessionAdminIdOk

`func (o *SessionHistoryItem) GetSessionAdminIdOk() (*string, bool)`

GetSessionAdminIdOk returns a tuple with the SessionAdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionAdminId

`func (o *SessionHistoryItem) SetSessionAdminId(v string)`

SetSessionAdminId sets SessionAdminId field to given value.

### HasSessionAdminId

`func (o *SessionHistoryItem) HasSessionAdminId() bool`

HasSessionAdminId returns a boolean if a field has been set.

### GetComment

`func (o *SessionHistoryItem) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SessionHistoryItem) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SessionHistoryItem) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SessionHistoryItem) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


