# LastLoginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastLogin** | **time.Time** | Timestamp of the last login (ISO 8601 format) | 

## Methods

### NewLastLoginResponse

`func NewLastLoginResponse(lastLogin time.Time, ) *LastLoginResponse`

NewLastLoginResponse instantiates a new LastLoginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastLoginResponseWithDefaults

`func NewLastLoginResponseWithDefaults() *LastLoginResponse`

NewLastLoginResponseWithDefaults instantiates a new LastLoginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastLogin

`func (o *LastLoginResponse) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *LastLoginResponse) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *LastLoginResponse) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


