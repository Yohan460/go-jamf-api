# StaticGroupAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** |  | [optional] [readonly] 
**GroupName** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**Assignments** | Pointer to [**[]Assignment**](Assignment.md) |  | [optional] 

## Methods

### NewStaticGroupAssignment

`func NewStaticGroupAssignment() *StaticGroupAssignment`

NewStaticGroupAssignment instantiates a new StaticGroupAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticGroupAssignmentWithDefaults

`func NewStaticGroupAssignmentWithDefaults() *StaticGroupAssignment`

NewStaticGroupAssignmentWithDefaults instantiates a new StaticGroupAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *StaticGroupAssignment) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *StaticGroupAssignment) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *StaticGroupAssignment) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *StaticGroupAssignment) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *StaticGroupAssignment) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *StaticGroupAssignment) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *StaticGroupAssignment) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *StaticGroupAssignment) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetSiteId

`func (o *StaticGroupAssignment) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *StaticGroupAssignment) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *StaticGroupAssignment) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *StaticGroupAssignment) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetAssignments

`func (o *StaticGroupAssignment) GetAssignments() []Assignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *StaticGroupAssignment) GetAssignmentsOk() (*[]Assignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *StaticGroupAssignment) SetAssignments(v []Assignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *StaticGroupAssignment) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


