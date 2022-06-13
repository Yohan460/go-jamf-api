# SessionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionState** | Pointer to **string** | Session state | [optional] 
**Online** | Pointer to **bool** | Defines if the end user is online | [optional] 

## Methods

### NewSessionStatus

`func NewSessionStatus() *SessionStatus`

NewSessionStatus instantiates a new SessionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionStatusWithDefaults

`func NewSessionStatusWithDefaults() *SessionStatus`

NewSessionStatusWithDefaults instantiates a new SessionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionState

`func (o *SessionStatus) GetSessionState() string`

GetSessionState returns the SessionState field if non-nil, zero value otherwise.

### GetSessionStateOk

`func (o *SessionStatus) GetSessionStateOk() (*string, bool)`

GetSessionStateOk returns a tuple with the SessionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionState

`func (o *SessionStatus) SetSessionState(v string)`

SetSessionState sets SessionState field to given value.

### HasSessionState

`func (o *SessionStatus) HasSessionState() bool`

HasSessionState returns a boolean if a field has been set.

### GetOnline

`func (o *SessionStatus) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *SessionStatus) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *SessionStatus) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *SessionStatus) HasOnline() bool`

HasOnline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


