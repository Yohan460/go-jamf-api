# SessionHistoryItemWithDetails

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
**Details** | Pointer to [**SessionHistoryItemDetails**](SessionHistoryItemDetails.md) |  | [optional] 

## Methods

### NewSessionHistoryItemWithDetails

`func NewSessionHistoryItemWithDetails() *SessionHistoryItemWithDetails`

NewSessionHistoryItemWithDetails instantiates a new SessionHistoryItemWithDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionHistoryItemWithDetailsWithDefaults

`func NewSessionHistoryItemWithDetailsWithDefaults() *SessionHistoryItemWithDetails`

NewSessionHistoryItemWithDetailsWithDefaults instantiates a new SessionHistoryItemWithDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *SessionHistoryItemWithDetails) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SessionHistoryItemWithDetails) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SessionHistoryItemWithDetails) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SessionHistoryItemWithDetails) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetSessionId

`func (o *SessionHistoryItemWithDetails) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *SessionHistoryItemWithDetails) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *SessionHistoryItemWithDetails) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *SessionHistoryItemWithDetails) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetDeviceId

`func (o *SessionHistoryItemWithDetails) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SessionHistoryItemWithDetails) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SessionHistoryItemWithDetails) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *SessionHistoryItemWithDetails) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetSessionStartedTimestamp

`func (o *SessionHistoryItemWithDetails) GetSessionStartedTimestamp() time.Time`

GetSessionStartedTimestamp returns the SessionStartedTimestamp field if non-nil, zero value otherwise.

### GetSessionStartedTimestampOk

`func (o *SessionHistoryItemWithDetails) GetSessionStartedTimestampOk() (*time.Time, bool)`

GetSessionStartedTimestampOk returns a tuple with the SessionStartedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStartedTimestamp

`func (o *SessionHistoryItemWithDetails) SetSessionStartedTimestamp(v time.Time)`

SetSessionStartedTimestamp sets SessionStartedTimestamp field to given value.

### HasSessionStartedTimestamp

`func (o *SessionHistoryItemWithDetails) HasSessionStartedTimestamp() bool`

HasSessionStartedTimestamp returns a boolean if a field has been set.

### GetSessionEndedTimestamp

`func (o *SessionHistoryItemWithDetails) GetSessionEndedTimestamp() time.Time`

GetSessionEndedTimestamp returns the SessionEndedTimestamp field if non-nil, zero value otherwise.

### GetSessionEndedTimestampOk

`func (o *SessionHistoryItemWithDetails) GetSessionEndedTimestampOk() (*time.Time, bool)`

GetSessionEndedTimestampOk returns a tuple with the SessionEndedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionEndedTimestamp

`func (o *SessionHistoryItemWithDetails) SetSessionEndedTimestamp(v time.Time)`

SetSessionEndedTimestamp sets SessionEndedTimestamp field to given value.

### HasSessionEndedTimestamp

`func (o *SessionHistoryItemWithDetails) HasSessionEndedTimestamp() bool`

HasSessionEndedTimestamp returns a boolean if a field has been set.

### GetSessionType

`func (o *SessionHistoryItemWithDetails) GetSessionType() string`

GetSessionType returns the SessionType field if non-nil, zero value otherwise.

### GetSessionTypeOk

`func (o *SessionHistoryItemWithDetails) GetSessionTypeOk() (*string, bool)`

GetSessionTypeOk returns a tuple with the SessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionType

`func (o *SessionHistoryItemWithDetails) SetSessionType(v string)`

SetSessionType sets SessionType field to given value.

### HasSessionType

`func (o *SessionHistoryItemWithDetails) HasSessionType() bool`

HasSessionType returns a boolean if a field has been set.

### GetStatusType

`func (o *SessionHistoryItemWithDetails) GetStatusType() string`

GetStatusType returns the StatusType field if non-nil, zero value otherwise.

### GetStatusTypeOk

`func (o *SessionHistoryItemWithDetails) GetStatusTypeOk() (*string, bool)`

GetStatusTypeOk returns a tuple with the StatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusType

`func (o *SessionHistoryItemWithDetails) SetStatusType(v string)`

SetStatusType sets StatusType field to given value.

### HasStatusType

`func (o *SessionHistoryItemWithDetails) HasStatusType() bool`

HasStatusType returns a boolean if a field has been set.

### GetSessionAdminId

`func (o *SessionHistoryItemWithDetails) GetSessionAdminId() string`

GetSessionAdminId returns the SessionAdminId field if non-nil, zero value otherwise.

### GetSessionAdminIdOk

`func (o *SessionHistoryItemWithDetails) GetSessionAdminIdOk() (*string, bool)`

GetSessionAdminIdOk returns a tuple with the SessionAdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionAdminId

`func (o *SessionHistoryItemWithDetails) SetSessionAdminId(v string)`

SetSessionAdminId sets SessionAdminId field to given value.

### HasSessionAdminId

`func (o *SessionHistoryItemWithDetails) HasSessionAdminId() bool`

HasSessionAdminId returns a boolean if a field has been set.

### GetComment

`func (o *SessionHistoryItemWithDetails) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SessionHistoryItemWithDetails) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SessionHistoryItemWithDetails) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SessionHistoryItemWithDetails) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDetails

`func (o *SessionHistoryItemWithDetails) GetDetails() SessionHistoryItemDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SessionHistoryItemWithDetails) GetDetailsOk() (*SessionHistoryItemDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SessionHistoryItemWithDetails) SetDetails(v SessionHistoryItemDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *SessionHistoryItemWithDetails) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


