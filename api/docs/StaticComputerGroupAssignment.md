# StaticComputerGroupAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**SiteId** | Pointer to **NullableString** |  | [optional] 
**Assignments** | Pointer to **[]string** | Array of computer IDs to assign to the static group | [optional] 

## Methods

### NewStaticComputerGroupAssignment

`func NewStaticComputerGroupAssignment(name string, ) *StaticComputerGroupAssignment`

NewStaticComputerGroupAssignment instantiates a new StaticComputerGroupAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticComputerGroupAssignmentWithDefaults

`func NewStaticComputerGroupAssignmentWithDefaults() *StaticComputerGroupAssignment`

NewStaticComputerGroupAssignmentWithDefaults instantiates a new StaticComputerGroupAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StaticComputerGroupAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StaticComputerGroupAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StaticComputerGroupAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StaticComputerGroupAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StaticComputerGroupAssignment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StaticComputerGroupAssignment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StaticComputerGroupAssignment) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *StaticComputerGroupAssignment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StaticComputerGroupAssignment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StaticComputerGroupAssignment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StaticComputerGroupAssignment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *StaticComputerGroupAssignment) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *StaticComputerGroupAssignment) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSiteId

`func (o *StaticComputerGroupAssignment) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *StaticComputerGroupAssignment) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *StaticComputerGroupAssignment) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *StaticComputerGroupAssignment) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *StaticComputerGroupAssignment) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *StaticComputerGroupAssignment) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
### GetAssignments

`func (o *StaticComputerGroupAssignment) GetAssignments() []string`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *StaticComputerGroupAssignment) GetAssignmentsOk() (*[]string, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *StaticComputerGroupAssignment) SetAssignments(v []string)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *StaticComputerGroupAssignment) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


