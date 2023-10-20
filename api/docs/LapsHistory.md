# LapsHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDate** | Pointer to **NullableTime** |  | [optional] 
**DateLastSeen** | Pointer to **NullableTime** |  | [optional] 
**ExpirationTime** | Pointer to **NullableTime** |  | [optional] 
**RotationStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewLapsHistory

`func NewLapsHistory() *LapsHistory`

NewLapsHistory instantiates a new LapsHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLapsHistoryWithDefaults

`func NewLapsHistoryWithDefaults() *LapsHistory`

NewLapsHistoryWithDefaults instantiates a new LapsHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *LapsHistory) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *LapsHistory) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *LapsHistory) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *LapsHistory) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *LapsHistory) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *LapsHistory) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetDateLastSeen

`func (o *LapsHistory) GetDateLastSeen() time.Time`

GetDateLastSeen returns the DateLastSeen field if non-nil, zero value otherwise.

### GetDateLastSeenOk

`func (o *LapsHistory) GetDateLastSeenOk() (*time.Time, bool)`

GetDateLastSeenOk returns a tuple with the DateLastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastSeen

`func (o *LapsHistory) SetDateLastSeen(v time.Time)`

SetDateLastSeen sets DateLastSeen field to given value.

### HasDateLastSeen

`func (o *LapsHistory) HasDateLastSeen() bool`

HasDateLastSeen returns a boolean if a field has been set.

### SetDateLastSeenNil

`func (o *LapsHistory) SetDateLastSeenNil(b bool)`

 SetDateLastSeenNil sets the value for DateLastSeen to be an explicit nil

### UnsetDateLastSeen
`func (o *LapsHistory) UnsetDateLastSeen()`

UnsetDateLastSeen ensures that no value is present for DateLastSeen, not even an explicit nil
### GetExpirationTime

`func (o *LapsHistory) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *LapsHistory) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *LapsHistory) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *LapsHistory) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### SetExpirationTimeNil

`func (o *LapsHistory) SetExpirationTimeNil(b bool)`

 SetExpirationTimeNil sets the value for ExpirationTime to be an explicit nil

### UnsetExpirationTime
`func (o *LapsHistory) UnsetExpirationTime()`

UnsetExpirationTime ensures that no value is present for ExpirationTime, not even an explicit nil
### GetRotationStatus

`func (o *LapsHistory) GetRotationStatus() string`

GetRotationStatus returns the RotationStatus field if non-nil, zero value otherwise.

### GetRotationStatusOk

`func (o *LapsHistory) GetRotationStatusOk() (*string, bool)`

GetRotationStatusOk returns a tuple with the RotationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationStatus

`func (o *LapsHistory) SetRotationStatus(v string)`

SetRotationStatus sets RotationStatus field to given value.

### HasRotationStatus

`func (o *LapsHistory) HasRotationStatus() bool`

HasRotationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


