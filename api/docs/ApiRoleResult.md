# ApiRoleResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int64** |  | [readonly] 
**Results** | [**[]ApiRole**](ApiRole.md) |  | 

## Methods

### NewApiRoleResult

`func NewApiRoleResult(totalCount int64, results []ApiRole, ) *ApiRoleResult`

NewApiRoleResult instantiates a new ApiRoleResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRoleResultWithDefaults

`func NewApiRoleResultWithDefaults() *ApiRoleResult`

NewApiRoleResultWithDefaults instantiates a new ApiRoleResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ApiRoleResult) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ApiRoleResult) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ApiRoleResult) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetResults

`func (o *ApiRoleResult) GetResults() []ApiRole`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ApiRoleResult) GetResultsOk() (*[]ApiRole, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ApiRoleResult) SetResults(v []ApiRole)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


