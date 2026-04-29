# GroupWithCriteriaDtoV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupPlatformId** | Pointer to **string** |  | [optional] [readonly] 
**GroupJamfProId** | Pointer to **string** |  | [optional] [readonly] 
**GroupName** | Pointer to **string** |  | [optional] [readonly] 
**GroupDescription** | Pointer to **string** |  | [optional] [readonly] 
**GroupType** | Pointer to **string** |  | [optional] [readonly] 
**Smart** | Pointer to **bool** |  | [optional] [readonly] 
**MembershipCount** | Pointer to **int64** |  | [optional] [readonly] 
**Criteria** | Pointer to [**[]SmartGroupCriteria**](SmartGroupCriteria.md) |  | [optional] 

## Methods

### NewGroupWithCriteriaDtoV1

`func NewGroupWithCriteriaDtoV1() *GroupWithCriteriaDtoV1`

NewGroupWithCriteriaDtoV1 instantiates a new GroupWithCriteriaDtoV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithCriteriaDtoV1WithDefaults

`func NewGroupWithCriteriaDtoV1WithDefaults() *GroupWithCriteriaDtoV1`

NewGroupWithCriteriaDtoV1WithDefaults instantiates a new GroupWithCriteriaDtoV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupPlatformId

`func (o *GroupWithCriteriaDtoV1) GetGroupPlatformId() string`

GetGroupPlatformId returns the GroupPlatformId field if non-nil, zero value otherwise.

### GetGroupPlatformIdOk

`func (o *GroupWithCriteriaDtoV1) GetGroupPlatformIdOk() (*string, bool)`

GetGroupPlatformIdOk returns a tuple with the GroupPlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPlatformId

`func (o *GroupWithCriteriaDtoV1) SetGroupPlatformId(v string)`

SetGroupPlatformId sets GroupPlatformId field to given value.

### HasGroupPlatformId

`func (o *GroupWithCriteriaDtoV1) HasGroupPlatformId() bool`

HasGroupPlatformId returns a boolean if a field has been set.

### GetGroupJamfProId

`func (o *GroupWithCriteriaDtoV1) GetGroupJamfProId() string`

GetGroupJamfProId returns the GroupJamfProId field if non-nil, zero value otherwise.

### GetGroupJamfProIdOk

`func (o *GroupWithCriteriaDtoV1) GetGroupJamfProIdOk() (*string, bool)`

GetGroupJamfProIdOk returns a tuple with the GroupJamfProId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupJamfProId

`func (o *GroupWithCriteriaDtoV1) SetGroupJamfProId(v string)`

SetGroupJamfProId sets GroupJamfProId field to given value.

### HasGroupJamfProId

`func (o *GroupWithCriteriaDtoV1) HasGroupJamfProId() bool`

HasGroupJamfProId returns a boolean if a field has been set.

### GetGroupName

`func (o *GroupWithCriteriaDtoV1) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GroupWithCriteriaDtoV1) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GroupWithCriteriaDtoV1) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *GroupWithCriteriaDtoV1) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetGroupDescription

`func (o *GroupWithCriteriaDtoV1) GetGroupDescription() string`

GetGroupDescription returns the GroupDescription field if non-nil, zero value otherwise.

### GetGroupDescriptionOk

`func (o *GroupWithCriteriaDtoV1) GetGroupDescriptionOk() (*string, bool)`

GetGroupDescriptionOk returns a tuple with the GroupDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDescription

`func (o *GroupWithCriteriaDtoV1) SetGroupDescription(v string)`

SetGroupDescription sets GroupDescription field to given value.

### HasGroupDescription

`func (o *GroupWithCriteriaDtoV1) HasGroupDescription() bool`

HasGroupDescription returns a boolean if a field has been set.

### GetGroupType

`func (o *GroupWithCriteriaDtoV1) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *GroupWithCriteriaDtoV1) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *GroupWithCriteriaDtoV1) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *GroupWithCriteriaDtoV1) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetSmart

`func (o *GroupWithCriteriaDtoV1) GetSmart() bool`

GetSmart returns the Smart field if non-nil, zero value otherwise.

### GetSmartOk

`func (o *GroupWithCriteriaDtoV1) GetSmartOk() (*bool, bool)`

GetSmartOk returns a tuple with the Smart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmart

`func (o *GroupWithCriteriaDtoV1) SetSmart(v bool)`

SetSmart sets Smart field to given value.

### HasSmart

`func (o *GroupWithCriteriaDtoV1) HasSmart() bool`

HasSmart returns a boolean if a field has been set.

### GetMembershipCount

`func (o *GroupWithCriteriaDtoV1) GetMembershipCount() int64`

GetMembershipCount returns the MembershipCount field if non-nil, zero value otherwise.

### GetMembershipCountOk

`func (o *GroupWithCriteriaDtoV1) GetMembershipCountOk() (*int64, bool)`

GetMembershipCountOk returns a tuple with the MembershipCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipCount

`func (o *GroupWithCriteriaDtoV1) SetMembershipCount(v int64)`

SetMembershipCount sets MembershipCount field to given value.

### HasMembershipCount

`func (o *GroupWithCriteriaDtoV1) HasMembershipCount() bool`

HasMembershipCount returns a boolean if a field has been set.

### GetCriteria

`func (o *GroupWithCriteriaDtoV1) GetCriteria() []SmartGroupCriteria`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *GroupWithCriteriaDtoV1) GetCriteriaOk() (*[]SmartGroupCriteria, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *GroupWithCriteriaDtoV1) SetCriteria(v []SmartGroupCriteria)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *GroupWithCriteriaDtoV1) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### SetCriteriaNil

`func (o *GroupWithCriteriaDtoV1) SetCriteriaNil(b bool)`

 SetCriteriaNil sets the value for Criteria to be an explicit nil

### UnsetCriteria
`func (o *GroupWithCriteriaDtoV1) UnsetCriteria()`

UnsetCriteria ensures that no value is present for Criteria, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


