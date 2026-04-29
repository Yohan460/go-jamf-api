# SmartGroupAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | The unique identifier of the smart group | [optional] [readonly] 
**GroupName** | **string** |  | 
**GroupDescription** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**Criteria** | Pointer to [**[]SmartGroupCriteria**](SmartGroupCriteria.md) |  | [optional] 

## Methods

### NewSmartGroupAssignment

`func NewSmartGroupAssignment(groupName string, ) *SmartGroupAssignment`

NewSmartGroupAssignment instantiates a new SmartGroupAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartGroupAssignmentWithDefaults

`func NewSmartGroupAssignmentWithDefaults() *SmartGroupAssignment`

NewSmartGroupAssignmentWithDefaults instantiates a new SmartGroupAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *SmartGroupAssignment) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SmartGroupAssignment) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SmartGroupAssignment) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *SmartGroupAssignment) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *SmartGroupAssignment) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *SmartGroupAssignment) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *SmartGroupAssignment) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetGroupDescription

`func (o *SmartGroupAssignment) GetGroupDescription() string`

GetGroupDescription returns the GroupDescription field if non-nil, zero value otherwise.

### GetGroupDescriptionOk

`func (o *SmartGroupAssignment) GetGroupDescriptionOk() (*string, bool)`

GetGroupDescriptionOk returns a tuple with the GroupDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDescription

`func (o *SmartGroupAssignment) SetGroupDescription(v string)`

SetGroupDescription sets GroupDescription field to given value.

### HasGroupDescription

`func (o *SmartGroupAssignment) HasGroupDescription() bool`

HasGroupDescription returns a boolean if a field has been set.

### GetSiteId

`func (o *SmartGroupAssignment) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SmartGroupAssignment) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SmartGroupAssignment) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *SmartGroupAssignment) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetCriteria

`func (o *SmartGroupAssignment) GetCriteria() []SmartGroupCriteria`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *SmartGroupAssignment) GetCriteriaOk() (*[]SmartGroupCriteria, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *SmartGroupAssignment) SetCriteria(v []SmartGroupCriteria)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *SmartGroupAssignment) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


