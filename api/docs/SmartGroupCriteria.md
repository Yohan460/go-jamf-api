# SmartGroupCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The field to search on (e.g., Model, OS Version, etc.) | 
**Priority** | **int64** | The priority order of this criterion | 
**AndOr** | **string** | Whether this criterion should be ANDed or ORed with the previous criterion | 
**SearchType** | **string** | The type of search to perform (e.g., is, is not, like, etc.) | 
**Value** | **string** | The value to search for | 
**OpeningParen** | Pointer to **bool** | Whether to add an opening parenthesis before this criterion | [optional] 
**ClosingParen** | Pointer to **bool** | Whether to add a closing parenthesis after this criterion | [optional] 

## Methods

### NewSmartGroupCriteria

`func NewSmartGroupCriteria(name string, priority int64, andOr string, searchType string, value string, ) *SmartGroupCriteria`

NewSmartGroupCriteria instantiates a new SmartGroupCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartGroupCriteriaWithDefaults

`func NewSmartGroupCriteriaWithDefaults() *SmartGroupCriteria`

NewSmartGroupCriteriaWithDefaults instantiates a new SmartGroupCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SmartGroupCriteria) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SmartGroupCriteria) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SmartGroupCriteria) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *SmartGroupCriteria) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SmartGroupCriteria) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SmartGroupCriteria) SetPriority(v int64)`

SetPriority sets Priority field to given value.


### GetAndOr

`func (o *SmartGroupCriteria) GetAndOr() string`

GetAndOr returns the AndOr field if non-nil, zero value otherwise.

### GetAndOrOk

`func (o *SmartGroupCriteria) GetAndOrOk() (*string, bool)`

GetAndOrOk returns a tuple with the AndOr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndOr

`func (o *SmartGroupCriteria) SetAndOr(v string)`

SetAndOr sets AndOr field to given value.


### GetSearchType

`func (o *SmartGroupCriteria) GetSearchType() string`

GetSearchType returns the SearchType field if non-nil, zero value otherwise.

### GetSearchTypeOk

`func (o *SmartGroupCriteria) GetSearchTypeOk() (*string, bool)`

GetSearchTypeOk returns a tuple with the SearchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchType

`func (o *SmartGroupCriteria) SetSearchType(v string)`

SetSearchType sets SearchType field to given value.


### GetValue

`func (o *SmartGroupCriteria) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SmartGroupCriteria) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SmartGroupCriteria) SetValue(v string)`

SetValue sets Value field to given value.


### GetOpeningParen

`func (o *SmartGroupCriteria) GetOpeningParen() bool`

GetOpeningParen returns the OpeningParen field if non-nil, zero value otherwise.

### GetOpeningParenOk

`func (o *SmartGroupCriteria) GetOpeningParenOk() (*bool, bool)`

GetOpeningParenOk returns a tuple with the OpeningParen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningParen

`func (o *SmartGroupCriteria) SetOpeningParen(v bool)`

SetOpeningParen sets OpeningParen field to given value.

### HasOpeningParen

`func (o *SmartGroupCriteria) HasOpeningParen() bool`

HasOpeningParen returns a boolean if a field has been set.

### GetClosingParen

`func (o *SmartGroupCriteria) GetClosingParen() bool`

GetClosingParen returns the ClosingParen field if non-nil, zero value otherwise.

### GetClosingParenOk

`func (o *SmartGroupCriteria) GetClosingParenOk() (*bool, bool)`

GetClosingParenOk returns a tuple with the ClosingParen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingParen

`func (o *SmartGroupCriteria) SetClosingParen(v bool)`

SetClosingParen sets ClosingParen field to given value.

### HasClosingParen

`func (o *SmartGroupCriteria) HasClosingParen() bool`

HasClosingParen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


