# AccountGroupSearchResultsV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int64** | Total number of account groups matching the filter criteria | [optional] 
**Results** | Pointer to [**[]AccountGroupV1**](AccountGroupV1.md) | The collection of account groups for the requested page | [optional] 

## Methods

### NewAccountGroupSearchResultsV1

`func NewAccountGroupSearchResultsV1() *AccountGroupSearchResultsV1`

NewAccountGroupSearchResultsV1 instantiates a new AccountGroupSearchResultsV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountGroupSearchResultsV1WithDefaults

`func NewAccountGroupSearchResultsV1WithDefaults() *AccountGroupSearchResultsV1`

NewAccountGroupSearchResultsV1WithDefaults instantiates a new AccountGroupSearchResultsV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *AccountGroupSearchResultsV1) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AccountGroupSearchResultsV1) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AccountGroupSearchResultsV1) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *AccountGroupSearchResultsV1) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetResults

`func (o *AccountGroupSearchResultsV1) GetResults() []AccountGroupV1`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AccountGroupSearchResultsV1) GetResultsOk() (*[]AccountGroupV1, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AccountGroupSearchResultsV1) SetResults(v []AccountGroupV1)`

SetResults sets Results field to given value.

### HasResults

`func (o *AccountGroupSearchResultsV1) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


