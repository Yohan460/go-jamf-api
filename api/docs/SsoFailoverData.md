# SsoFailoverData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailoverUrl** | Pointer to **string** |  | [optional] 
**GenerationTime** | Pointer to **int64** | Generation time of failover key | [optional] 

## Methods

### NewSsoFailoverData

`func NewSsoFailoverData() *SsoFailoverData`

NewSsoFailoverData instantiates a new SsoFailoverData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoFailoverDataWithDefaults

`func NewSsoFailoverDataWithDefaults() *SsoFailoverData`

NewSsoFailoverDataWithDefaults instantiates a new SsoFailoverData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailoverUrl

`func (o *SsoFailoverData) GetFailoverUrl() string`

GetFailoverUrl returns the FailoverUrl field if non-nil, zero value otherwise.

### GetFailoverUrlOk

`func (o *SsoFailoverData) GetFailoverUrlOk() (*string, bool)`

GetFailoverUrlOk returns a tuple with the FailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverUrl

`func (o *SsoFailoverData) SetFailoverUrl(v string)`

SetFailoverUrl sets FailoverUrl field to given value.

### HasFailoverUrl

`func (o *SsoFailoverData) HasFailoverUrl() bool`

HasFailoverUrl returns a boolean if a field has been set.

### GetGenerationTime

`func (o *SsoFailoverData) GetGenerationTime() int64`

GetGenerationTime returns the GenerationTime field if non-nil, zero value otherwise.

### GetGenerationTimeOk

`func (o *SsoFailoverData) GetGenerationTimeOk() (*int64, bool)`

GetGenerationTimeOk returns a tuple with the GenerationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationTime

`func (o *SsoFailoverData) SetGenerationTime(v int64)`

SetGenerationTime sets GenerationTime field to given value.

### HasGenerationTime

`func (o *SsoFailoverData) HasGenerationTime() bool`

HasGenerationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


