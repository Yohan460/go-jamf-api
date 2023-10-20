# ApiIntegrationSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | [readonly] 
**Results** | [**[]ApiIntegrationResponse**](ApiIntegrationResponse.md) |  | 

## Methods

### NewApiIntegrationSearchResult

`func NewApiIntegrationSearchResult(totalCount int32, results []ApiIntegrationResponse, ) *ApiIntegrationSearchResult`

NewApiIntegrationSearchResult instantiates a new ApiIntegrationSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiIntegrationSearchResultWithDefaults

`func NewApiIntegrationSearchResultWithDefaults() *ApiIntegrationSearchResult`

NewApiIntegrationSearchResultWithDefaults instantiates a new ApiIntegrationSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ApiIntegrationSearchResult) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ApiIntegrationSearchResult) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ApiIntegrationSearchResult) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetResults

`func (o *ApiIntegrationSearchResult) GetResults() []ApiIntegrationResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ApiIntegrationSearchResult) GetResultsOk() (*[]ApiIntegrationResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ApiIntegrationSearchResult) SetResults(v []ApiIntegrationResponse)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


