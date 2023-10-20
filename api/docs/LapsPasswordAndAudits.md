# LapsPasswordAndAudits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** |  | [optional] [readonly] 
**DateLastSeen** | Pointer to **NullableTime** |  | [optional] 
**ExpirationTime** | Pointer to **NullableTime** |  | [optional] 
**Audits** | Pointer to [**[]LapsAudit**](LapsAudit.md) |  | [optional] 

## Methods

### NewLapsPasswordAndAudits

`func NewLapsPasswordAndAudits() *LapsPasswordAndAudits`

NewLapsPasswordAndAudits instantiates a new LapsPasswordAndAudits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLapsPasswordAndAuditsWithDefaults

`func NewLapsPasswordAndAuditsWithDefaults() *LapsPasswordAndAudits`

NewLapsPasswordAndAuditsWithDefaults instantiates a new LapsPasswordAndAudits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *LapsPasswordAndAudits) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LapsPasswordAndAudits) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LapsPasswordAndAudits) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LapsPasswordAndAudits) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetDateLastSeen

`func (o *LapsPasswordAndAudits) GetDateLastSeen() time.Time`

GetDateLastSeen returns the DateLastSeen field if non-nil, zero value otherwise.

### GetDateLastSeenOk

`func (o *LapsPasswordAndAudits) GetDateLastSeenOk() (*time.Time, bool)`

GetDateLastSeenOk returns a tuple with the DateLastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastSeen

`func (o *LapsPasswordAndAudits) SetDateLastSeen(v time.Time)`

SetDateLastSeen sets DateLastSeen field to given value.

### HasDateLastSeen

`func (o *LapsPasswordAndAudits) HasDateLastSeen() bool`

HasDateLastSeen returns a boolean if a field has been set.

### SetDateLastSeenNil

`func (o *LapsPasswordAndAudits) SetDateLastSeenNil(b bool)`

 SetDateLastSeenNil sets the value for DateLastSeen to be an explicit nil

### UnsetDateLastSeen
`func (o *LapsPasswordAndAudits) UnsetDateLastSeen()`

UnsetDateLastSeen ensures that no value is present for DateLastSeen, not even an explicit nil
### GetExpirationTime

`func (o *LapsPasswordAndAudits) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *LapsPasswordAndAudits) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *LapsPasswordAndAudits) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *LapsPasswordAndAudits) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### SetExpirationTimeNil

`func (o *LapsPasswordAndAudits) SetExpirationTimeNil(b bool)`

 SetExpirationTimeNil sets the value for ExpirationTime to be an explicit nil

### UnsetExpirationTime
`func (o *LapsPasswordAndAudits) UnsetExpirationTime()`

UnsetExpirationTime ensures that no value is present for ExpirationTime, not even an explicit nil
### GetAudits

`func (o *LapsPasswordAndAudits) GetAudits() []LapsAudit`

GetAudits returns the Audits field if non-nil, zero value otherwise.

### GetAuditsOk

`func (o *LapsPasswordAndAudits) GetAuditsOk() (*[]LapsAudit, bool)`

GetAuditsOk returns a tuple with the Audits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudits

`func (o *LapsPasswordAndAudits) SetAudits(v []LapsAudit)`

SetAudits sets Audits field to given value.

### HasAudits

`func (o *LapsPasswordAndAudits) HasAudits() bool`

HasAudits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


