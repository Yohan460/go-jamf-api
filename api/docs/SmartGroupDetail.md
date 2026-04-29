# SmartGroupDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**GroupDescription** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int64** | membership count | [optional] 
**Criteria** | Pointer to [**[]SmartGroupCriteria**](SmartGroupCriteria.md) | The criteria used to define the smart group | [optional] 

## Methods

### NewSmartGroupDetail

`func NewSmartGroupDetail() *SmartGroupDetail`

NewSmartGroupDetail instantiates a new SmartGroupDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartGroupDetailWithDefaults

`func NewSmartGroupDetailWithDefaults() *SmartGroupDetail`

NewSmartGroupDetailWithDefaults instantiates a new SmartGroupDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *SmartGroupDetail) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SmartGroupDetail) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SmartGroupDetail) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *SmartGroupDetail) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *SmartGroupDetail) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *SmartGroupDetail) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *SmartGroupDetail) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *SmartGroupDetail) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetGroupDescription

`func (o *SmartGroupDetail) GetGroupDescription() string`

GetGroupDescription returns the GroupDescription field if non-nil, zero value otherwise.

### GetGroupDescriptionOk

`func (o *SmartGroupDetail) GetGroupDescriptionOk() (*string, bool)`

GetGroupDescriptionOk returns a tuple with the GroupDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDescription

`func (o *SmartGroupDetail) SetGroupDescription(v string)`

SetGroupDescription sets GroupDescription field to given value.

### HasGroupDescription

`func (o *SmartGroupDetail) HasGroupDescription() bool`

HasGroupDescription returns a boolean if a field has been set.

### GetSiteId

`func (o *SmartGroupDetail) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SmartGroupDetail) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SmartGroupDetail) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SmartGroupDetail) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetCount

`func (o *SmartGroupDetail) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SmartGroupDetail) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SmartGroupDetail) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *SmartGroupDetail) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCriteria

`func (o *SmartGroupDetail) GetCriteria() []SmartGroupCriteria`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *SmartGroupDetail) GetCriteriaOk() (*[]SmartGroupCriteria, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *SmartGroupDetail) SetCriteria(v []SmartGroupCriteria)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *SmartGroupDetail) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


