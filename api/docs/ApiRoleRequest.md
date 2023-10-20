# ApiRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**Privileges** | **[]string** |  | 

## Methods

### NewApiRoleRequest

`func NewApiRoleRequest(displayName string, privileges []string, ) *ApiRoleRequest`

NewApiRoleRequest instantiates a new ApiRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRoleRequestWithDefaults

`func NewApiRoleRequestWithDefaults() *ApiRoleRequest`

NewApiRoleRequestWithDefaults instantiates a new ApiRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ApiRoleRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiRoleRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiRoleRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetPrivileges

`func (o *ApiRoleRequest) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *ApiRoleRequest) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *ApiRoleRequest) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


