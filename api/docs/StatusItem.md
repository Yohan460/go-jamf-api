# StatusItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The StatusItem key | [optional] 
**Value** | Pointer to **string** | The StatusItem value | [optional] 
**LastUpdateTime** | Pointer to **string** | The local server time when the StatusItem was last updated | [optional] 

## Methods

### NewStatusItem

`func NewStatusItem() *StatusItem`

NewStatusItem instantiates a new StatusItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusItemWithDefaults

`func NewStatusItemWithDefaults() *StatusItem`

NewStatusItemWithDefaults instantiates a new StatusItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *StatusItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StatusItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StatusItem) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *StatusItem) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *StatusItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StatusItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StatusItem) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *StatusItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *StatusItem) GetLastUpdateTime() string`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *StatusItem) GetLastUpdateTimeOk() (*string, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *StatusItem) SetLastUpdateTime(v string)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *StatusItem) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


