# LapsPendingRotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LapsUser** | Pointer to [**LapsUserV2**](LapsUserV2.md) |  | [optional] 
**CreatedDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewLapsPendingRotation

`func NewLapsPendingRotation() *LapsPendingRotation`

NewLapsPendingRotation instantiates a new LapsPendingRotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLapsPendingRotationWithDefaults

`func NewLapsPendingRotationWithDefaults() *LapsPendingRotation`

NewLapsPendingRotationWithDefaults instantiates a new LapsPendingRotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLapsUser

`func (o *LapsPendingRotation) GetLapsUser() LapsUserV2`

GetLapsUser returns the LapsUser field if non-nil, zero value otherwise.

### GetLapsUserOk

`func (o *LapsPendingRotation) GetLapsUserOk() (*LapsUserV2, bool)`

GetLapsUserOk returns a tuple with the LapsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLapsUser

`func (o *LapsPendingRotation) SetLapsUser(v LapsUserV2)`

SetLapsUser sets LapsUser field to given value.

### HasLapsUser

`func (o *LapsPendingRotation) HasLapsUser() bool`

HasLapsUser returns a boolean if a field has been set.

### GetCreatedDate

`func (o *LapsPendingRotation) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *LapsPendingRotation) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *LapsPendingRotation) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *LapsPendingRotation) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


