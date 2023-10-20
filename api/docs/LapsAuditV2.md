# LapsAuditV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewedBy** | Pointer to **NullableString** |  | [optional] 
**DateSeen** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewLapsAuditV2

`func NewLapsAuditV2() *LapsAuditV2`

NewLapsAuditV2 instantiates a new LapsAuditV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLapsAuditV2WithDefaults

`func NewLapsAuditV2WithDefaults() *LapsAuditV2`

NewLapsAuditV2WithDefaults instantiates a new LapsAuditV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewedBy

`func (o *LapsAuditV2) GetViewedBy() string`

GetViewedBy returns the ViewedBy field if non-nil, zero value otherwise.

### GetViewedByOk

`func (o *LapsAuditV2) GetViewedByOk() (*string, bool)`

GetViewedByOk returns a tuple with the ViewedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedBy

`func (o *LapsAuditV2) SetViewedBy(v string)`

SetViewedBy sets ViewedBy field to given value.

### HasViewedBy

`func (o *LapsAuditV2) HasViewedBy() bool`

HasViewedBy returns a boolean if a field has been set.

### SetViewedByNil

`func (o *LapsAuditV2) SetViewedByNil(b bool)`

 SetViewedByNil sets the value for ViewedBy to be an explicit nil

### UnsetViewedBy
`func (o *LapsAuditV2) UnsetViewedBy()`

UnsetViewedBy ensures that no value is present for ViewedBy, not even an explicit nil
### GetDateSeen

`func (o *LapsAuditV2) GetDateSeen() time.Time`

GetDateSeen returns the DateSeen field if non-nil, zero value otherwise.

### GetDateSeenOk

`func (o *LapsAuditV2) GetDateSeenOk() (*time.Time, bool)`

GetDateSeenOk returns a tuple with the DateSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateSeen

`func (o *LapsAuditV2) SetDateSeen(v time.Time)`

SetDateSeen sets DateSeen field to given value.

### HasDateSeen

`func (o *LapsAuditV2) HasDateSeen() bool`

HasDateSeen returns a boolean if a field has been set.

### SetDateSeenNil

`func (o *LapsAuditV2) SetDateSeenNil(b bool)`

 SetDateSeenNil sets the value for DateSeen to be an explicit nil

### UnsetDateSeen
`func (o *LapsAuditV2) UnsetDateSeen()`

UnsetDateSeen ensures that no value is present for DateSeen, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


