# GroupUpdateDtoV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | Pointer to **string** |  | [optional] 
**GroupDescription** | Pointer to **string** |  | [optional] 
**Criteria** | Pointer to [**[]SmartGroupCriteria**](SmartGroupCriteria.md) |  | [optional] 
**Assignments** | Pointer to [**[]AssignmentDtoV1**](AssignmentDtoV1.md) |  | [optional] 

## Methods

### NewGroupUpdateDtoV1

`func NewGroupUpdateDtoV1() *GroupUpdateDtoV1`

NewGroupUpdateDtoV1 instantiates a new GroupUpdateDtoV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUpdateDtoV1WithDefaults

`func NewGroupUpdateDtoV1WithDefaults() *GroupUpdateDtoV1`

NewGroupUpdateDtoV1WithDefaults instantiates a new GroupUpdateDtoV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *GroupUpdateDtoV1) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GroupUpdateDtoV1) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GroupUpdateDtoV1) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *GroupUpdateDtoV1) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetGroupDescription

`func (o *GroupUpdateDtoV1) GetGroupDescription() string`

GetGroupDescription returns the GroupDescription field if non-nil, zero value otherwise.

### GetGroupDescriptionOk

`func (o *GroupUpdateDtoV1) GetGroupDescriptionOk() (*string, bool)`

GetGroupDescriptionOk returns a tuple with the GroupDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDescription

`func (o *GroupUpdateDtoV1) SetGroupDescription(v string)`

SetGroupDescription sets GroupDescription field to given value.

### HasGroupDescription

`func (o *GroupUpdateDtoV1) HasGroupDescription() bool`

HasGroupDescription returns a boolean if a field has been set.

### GetCriteria

`func (o *GroupUpdateDtoV1) GetCriteria() []SmartGroupCriteria`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *GroupUpdateDtoV1) GetCriteriaOk() (*[]SmartGroupCriteria, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *GroupUpdateDtoV1) SetCriteria(v []SmartGroupCriteria)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *GroupUpdateDtoV1) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetAssignments

`func (o *GroupUpdateDtoV1) GetAssignments() []AssignmentDtoV1`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *GroupUpdateDtoV1) GetAssignmentsOk() (*[]AssignmentDtoV1, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *GroupUpdateDtoV1) SetAssignments(v []AssignmentDtoV1)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *GroupUpdateDtoV1) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignmentsNil

`func (o *GroupUpdateDtoV1) SetAssignmentsNil(b bool)`

 SetAssignmentsNil sets the value for Assignments to be an explicit nil

### UnsetAssignments
`func (o *GroupUpdateDtoV1) UnsetAssignments()`

UnsetAssignments ensures that no value is present for Assignments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


