# SessionDetailsSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**[]SessionDetails**](SessionDetails.md) |  | [optional] 

## Methods

### NewSessionDetailsSearchResults

`func NewSessionDetailsSearchResults() *SessionDetailsSearchResults`

NewSessionDetailsSearchResults instantiates a new SessionDetailsSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionDetailsSearchResultsWithDefaults

`func NewSessionDetailsSearchResultsWithDefaults() *SessionDetailsSearchResults`

NewSessionDetailsSearchResultsWithDefaults instantiates a new SessionDetailsSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *SessionDetailsSearchResults) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SessionDetailsSearchResults) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SessionDetailsSearchResults) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *SessionDetailsSearchResults) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetResults

`func (o *SessionDetailsSearchResults) GetResults() []SessionDetails`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SessionDetailsSearchResults) GetResultsOk() (*[]SessionDetails, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SessionDetailsSearchResults) SetResults(v []SessionDetails)`

SetResults sets Results field to given value.

### HasResults

`func (o *SessionDetailsSearchResults) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


