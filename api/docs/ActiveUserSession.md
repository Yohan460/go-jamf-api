# ActiveUserSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | Username of the logged in user | [optional] 
**SessionId** | Pointer to **string** | Unique session identifier | [optional] 
**LastAccessedTime** | Pointer to **time.Time** | Timestamp of when the session was last accessed | [optional] 
**CreationTime** | Pointer to **time.Time** | Timestamp of when the session was created | [optional] 
**UserAgent** | Pointer to **NullableString** | User agent string from the session | [optional] 
**IpAddress** | Pointer to **NullableString** | IP address associated with the session | [optional] 

## Methods

### NewActiveUserSession

`func NewActiveUserSession() *ActiveUserSession`

NewActiveUserSession instantiates a new ActiveUserSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveUserSessionWithDefaults

`func NewActiveUserSessionWithDefaults() *ActiveUserSession`

NewActiveUserSessionWithDefaults instantiates a new ActiveUserSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *ActiveUserSession) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ActiveUserSession) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ActiveUserSession) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ActiveUserSession) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetSessionId

`func (o *ActiveUserSession) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ActiveUserSession) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ActiveUserSession) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ActiveUserSession) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetLastAccessedTime

`func (o *ActiveUserSession) GetLastAccessedTime() time.Time`

GetLastAccessedTime returns the LastAccessedTime field if non-nil, zero value otherwise.

### GetLastAccessedTimeOk

`func (o *ActiveUserSession) GetLastAccessedTimeOk() (*time.Time, bool)`

GetLastAccessedTimeOk returns a tuple with the LastAccessedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessedTime

`func (o *ActiveUserSession) SetLastAccessedTime(v time.Time)`

SetLastAccessedTime sets LastAccessedTime field to given value.

### HasLastAccessedTime

`func (o *ActiveUserSession) HasLastAccessedTime() bool`

HasLastAccessedTime returns a boolean if a field has been set.

### GetCreationTime

`func (o *ActiveUserSession) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ActiveUserSession) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ActiveUserSession) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *ActiveUserSession) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetUserAgent

`func (o *ActiveUserSession) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *ActiveUserSession) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *ActiveUserSession) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *ActiveUserSession) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### SetUserAgentNil

`func (o *ActiveUserSession) SetUserAgentNil(b bool)`

 SetUserAgentNil sets the value for UserAgent to be an explicit nil

### UnsetUserAgent
`func (o *ActiveUserSession) UnsetUserAgent()`

UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil
### GetIpAddress

`func (o *ActiveUserSession) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ActiveUserSession) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ActiveUserSession) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ActiveUserSession) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *ActiveUserSession) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *ActiveUserSession) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


