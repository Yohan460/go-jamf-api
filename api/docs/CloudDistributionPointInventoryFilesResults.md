# CloudDistributionPointInventoryFilesResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int64** |  | 
**Results** | Pointer to [**[]CloudDistributionPointInventoryFileInfo**](CloudDistributionPointInventoryFileInfo.md) |  | [optional] 

## Methods

### NewCloudDistributionPointInventoryFilesResults

`func NewCloudDistributionPointInventoryFilesResults(totalCount int64, ) *CloudDistributionPointInventoryFilesResults`

NewCloudDistributionPointInventoryFilesResults instantiates a new CloudDistributionPointInventoryFilesResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDistributionPointInventoryFilesResultsWithDefaults

`func NewCloudDistributionPointInventoryFilesResultsWithDefaults() *CloudDistributionPointInventoryFilesResults`

NewCloudDistributionPointInventoryFilesResultsWithDefaults instantiates a new CloudDistributionPointInventoryFilesResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *CloudDistributionPointInventoryFilesResults) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CloudDistributionPointInventoryFilesResults) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CloudDistributionPointInventoryFilesResults) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetResults

`func (o *CloudDistributionPointInventoryFilesResults) GetResults() []CloudDistributionPointInventoryFileInfo`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CloudDistributionPointInventoryFilesResults) GetResultsOk() (*[]CloudDistributionPointInventoryFileInfo, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CloudDistributionPointInventoryFilesResults) SetResults(v []CloudDistributionPointInventoryFileInfo)`

SetResults sets Results field to given value.

### HasResults

`func (o *CloudDistributionPointInventoryFilesResults) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


