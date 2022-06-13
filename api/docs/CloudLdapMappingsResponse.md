# CloudLdapMappingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserMappings** | Pointer to [**UserMappings**](UserMappings.md) |  | [optional] 
**GroupMappings** | Pointer to [**GroupMappings**](GroupMappings.md) |  | [optional] 
**MembershipMappings** | Pointer to [**MembershipMappings**](MembershipMappings.md) |  | [optional] 

## Methods

### NewCloudLdapMappingsResponse

`func NewCloudLdapMappingsResponse() *CloudLdapMappingsResponse`

NewCloudLdapMappingsResponse instantiates a new CloudLdapMappingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudLdapMappingsResponseWithDefaults

`func NewCloudLdapMappingsResponseWithDefaults() *CloudLdapMappingsResponse`

NewCloudLdapMappingsResponseWithDefaults instantiates a new CloudLdapMappingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserMappings

`func (o *CloudLdapMappingsResponse) GetUserMappings() UserMappings`

GetUserMappings returns the UserMappings field if non-nil, zero value otherwise.

### GetUserMappingsOk

`func (o *CloudLdapMappingsResponse) GetUserMappingsOk() (*UserMappings, bool)`

GetUserMappingsOk returns a tuple with the UserMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMappings

`func (o *CloudLdapMappingsResponse) SetUserMappings(v UserMappings)`

SetUserMappings sets UserMappings field to given value.

### HasUserMappings

`func (o *CloudLdapMappingsResponse) HasUserMappings() bool`

HasUserMappings returns a boolean if a field has been set.

### GetGroupMappings

`func (o *CloudLdapMappingsResponse) GetGroupMappings() GroupMappings`

GetGroupMappings returns the GroupMappings field if non-nil, zero value otherwise.

### GetGroupMappingsOk

`func (o *CloudLdapMappingsResponse) GetGroupMappingsOk() (*GroupMappings, bool)`

GetGroupMappingsOk returns a tuple with the GroupMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMappings

`func (o *CloudLdapMappingsResponse) SetGroupMappings(v GroupMappings)`

SetGroupMappings sets GroupMappings field to given value.

### HasGroupMappings

`func (o *CloudLdapMappingsResponse) HasGroupMappings() bool`

HasGroupMappings returns a boolean if a field has been set.

### GetMembershipMappings

`func (o *CloudLdapMappingsResponse) GetMembershipMappings() MembershipMappings`

GetMembershipMappings returns the MembershipMappings field if non-nil, zero value otherwise.

### GetMembershipMappingsOk

`func (o *CloudLdapMappingsResponse) GetMembershipMappingsOk() (*MembershipMappings, bool)`

GetMembershipMappingsOk returns a tuple with the MembershipMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipMappings

`func (o *CloudLdapMappingsResponse) SetMembershipMappings(v MembershipMappings)`

SetMembershipMappings sets MembershipMappings field to given value.

### HasMembershipMappings

`func (o *CloudLdapMappingsResponse) HasMembershipMappings() bool`

HasMembershipMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


