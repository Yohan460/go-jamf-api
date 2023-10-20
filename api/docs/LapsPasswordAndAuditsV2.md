# LapsPasswordAndAuditsV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** |  | [optional] [readonly] 
**DateLastSeen** | Pointer to **NullableTime** |  | [optional] 
**ExpirationTime** | Pointer to **NullableTime** |  | [optional] 
**Audits** | Pointer to [**[]LapsAuditV2**](LapsAuditV2.md) |  | [optional] 

## Methods

### NewLapsPasswordAndAuditsV2

`func NewLapsPasswordAndAuditsV2() *LapsPasswordAndAuditsV2`

NewLapsPasswordAndAuditsV2 instantiates a new LapsPasswordAndAuditsV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLapsPasswordAndAuditsV2WithDefaults

`func NewLapsPasswordAndAuditsV2WithDefaults() *LapsPasswordAndAuditsV2`

NewLapsPasswordAndAuditsV2WithDefaults instantiates a new LapsPasswordAndAuditsV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *LapsPasswordAndAuditsV2) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LapsPasswordAndAuditsV2) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LapsPasswordAndAuditsV2) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LapsPasswordAndAuditsV2) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetDateLastSeen

`func (o *LapsPasswordAndAuditsV2) GetDateLastSeen() time.Time`

GetDateLastSeen returns the DateLastSeen field if non-nil, zero value otherwise.

### GetDateLastSeenOk

`func (o *LapsPasswordAndAuditsV2) GetDateLastSeenOk() (*time.Time, bool)`

GetDateLastSeenOk returns a tuple with the DateLastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastSeen

`func (o *LapsPasswordAndAuditsV2) SetDateLastSeen(v time.Time)`

SetDateLastSeen sets DateLastSeen field to given value.

### HasDateLastSeen

`func (o *LapsPasswordAndAuditsV2) HasDateLastSeen() bool`

HasDateLastSeen returns a boolean if a field has been set.

### SetDateLastSeenNil

`func (o *LapsPasswordAndAuditsV2) SetDateLastSeenNil(b bool)`

 SetDateLastSeenNil sets the value for DateLastSeen to be an explicit nil

### UnsetDateLastSeen
`func (o *LapsPasswordAndAuditsV2) UnsetDateLastSeen()`

UnsetDateLastSeen ensures that no value is present for DateLastSeen, not even an explicit nil
### GetExpirationTime

`func (o *LapsPasswordAndAuditsV2) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *LapsPasswordAndAuditsV2) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *LapsPasswordAndAuditsV2) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *LapsPasswordAndAuditsV2) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### SetExpirationTimeNil

`func (o *LapsPasswordAndAuditsV2) SetExpirationTimeNil(b bool)`

 SetExpirationTimeNil sets the value for ExpirationTime to be an explicit nil

### UnsetExpirationTime
`func (o *LapsPasswordAndAuditsV2) UnsetExpirationTime()`

UnsetExpirationTime ensures that no value is present for ExpirationTime, not even an explicit nil
### GetAudits

`func (o *LapsPasswordAndAuditsV2) GetAudits() []LapsAuditV2`

GetAudits returns the Audits field if non-nil, zero value otherwise.

### GetAuditsOk

`func (o *LapsPasswordAndAuditsV2) GetAuditsOk() (*[]LapsAuditV2, bool)`

GetAuditsOk returns a tuple with the Audits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudits

`func (o *LapsPasswordAndAuditsV2) SetAudits(v []LapsAuditV2)`

SetAudits sets Audits field to given value.

### HasAudits

`func (o *LapsPasswordAndAuditsV2) HasAudits() bool`

HasAudits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


