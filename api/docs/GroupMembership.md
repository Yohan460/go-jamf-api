# GroupMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**SmartGroup** | Pointer to **bool** | Indicates that group is smart group | [optional] 

## Methods

### NewGroupMembership

`func NewGroupMembership() *GroupMembership`

NewGroupMembership instantiates a new GroupMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMembershipWithDefaults

`func NewGroupMembershipWithDefaults() *GroupMembership`

NewGroupMembershipWithDefaults instantiates a new GroupMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GroupMembership) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupMembership) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupMembership) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GroupMembership) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *GroupMembership) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GroupMembership) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GroupMembership) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *GroupMembership) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetSmartGroup

`func (o *GroupMembership) GetSmartGroup() bool`

GetSmartGroup returns the SmartGroup field if non-nil, zero value otherwise.

### GetSmartGroupOk

`func (o *GroupMembership) GetSmartGroupOk() (*bool, bool)`

GetSmartGroupOk returns a tuple with the SmartGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartGroup

`func (o *GroupMembership) SetSmartGroup(v bool)`

SetSmartGroup sets SmartGroup field to given value.

### HasSmartGroup

`func (o *GroupMembership) HasSmartGroup() bool`

HasSmartGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


