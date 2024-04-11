# ObjectHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewObjectHistory

`func NewObjectHistory() *ObjectHistory`

NewObjectHistory instantiates a new ObjectHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectHistoryWithDefaults

`func NewObjectHistoryWithDefaults() *ObjectHistory`

NewObjectHistoryWithDefaults instantiates a new ObjectHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectHistory) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectHistory) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectHistory) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectHistory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *ObjectHistory) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ObjectHistory) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ObjectHistory) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ObjectHistory) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDate

`func (o *ObjectHistory) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ObjectHistory) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ObjectHistory) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *ObjectHistory) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetNote

`func (o *ObjectHistory) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ObjectHistory) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ObjectHistory) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ObjectHistory) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetDetails

`func (o *ObjectHistory) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ObjectHistory) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ObjectHistory) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ObjectHistory) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *ObjectHistory) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *ObjectHistory) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


