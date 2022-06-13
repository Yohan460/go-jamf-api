# CloudLdapMappingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserMappings** | [**UserMappings**](UserMappings.md) |  | 
**GroupMappings** | [**GroupMappings**](GroupMappings.md) |  | 
**MembershipMappings** | [**MembershipMappings**](MembershipMappings.md) |  | 

## Methods

### NewCloudLdapMappingsRequest

`func NewCloudLdapMappingsRequest(userMappings UserMappings, groupMappings GroupMappings, membershipMappings MembershipMappings, ) *CloudLdapMappingsRequest`

NewCloudLdapMappingsRequest instantiates a new CloudLdapMappingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudLdapMappingsRequestWithDefaults

`func NewCloudLdapMappingsRequestWithDefaults() *CloudLdapMappingsRequest`

NewCloudLdapMappingsRequestWithDefaults instantiates a new CloudLdapMappingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserMappings

`func (o *CloudLdapMappingsRequest) GetUserMappings() UserMappings`

GetUserMappings returns the UserMappings field if non-nil, zero value otherwise.

### GetUserMappingsOk

`func (o *CloudLdapMappingsRequest) GetUserMappingsOk() (*UserMappings, bool)`

GetUserMappingsOk returns a tuple with the UserMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMappings

`func (o *CloudLdapMappingsRequest) SetUserMappings(v UserMappings)`

SetUserMappings sets UserMappings field to given value.


### GetGroupMappings

`func (o *CloudLdapMappingsRequest) GetGroupMappings() GroupMappings`

GetGroupMappings returns the GroupMappings field if non-nil, zero value otherwise.

### GetGroupMappingsOk

`func (o *CloudLdapMappingsRequest) GetGroupMappingsOk() (*GroupMappings, bool)`

GetGroupMappingsOk returns a tuple with the GroupMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMappings

`func (o *CloudLdapMappingsRequest) SetGroupMappings(v GroupMappings)`

SetGroupMappings sets GroupMappings field to given value.


### GetMembershipMappings

`func (o *CloudLdapMappingsRequest) GetMembershipMappings() MembershipMappings`

GetMembershipMappings returns the MembershipMappings field if non-nil, zero value otherwise.

### GetMembershipMappingsOk

`func (o *CloudLdapMappingsRequest) GetMembershipMappingsOk() (*MembershipMappings, bool)`

GetMembershipMappingsOk returns a tuple with the MembershipMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipMappings

`func (o *CloudLdapMappingsRequest) SetMembershipMappings(v MembershipMappings)`

SetMembershipMappings sets MembershipMappings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


