# PatchPolicyListView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**PolicyName** | Pointer to **string** |  | [optional] 
**PolicyEnabled** | Pointer to **bool** |  | [optional] 
**PolicyTargetVersion** | Pointer to **string** |  | [optional] 
**PolicyDeploymentMethod** | Pointer to **string** |  | [optional] 
**SoftwareTitle** | Pointer to **string** |  | [optional] 
**SoftwareTitleConfigurationId** | Pointer to **string** |  | [optional] 
**Pending** | Pointer to **int64** |  | [optional] 
**Completed** | Pointer to **int64** |  | [optional] 
**Deferred** | Pointer to **int64** |  | [optional] 
**Failed** | Pointer to **int64** |  | [optional] 

## Methods

### NewPatchPolicyListView

`func NewPatchPolicyListView() *PatchPolicyListView`

NewPatchPolicyListView instantiates a new PatchPolicyListView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPolicyListViewWithDefaults

`func NewPatchPolicyListViewWithDefaults() *PatchPolicyListView`

NewPatchPolicyListViewWithDefaults instantiates a new PatchPolicyListView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchPolicyListView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchPolicyListView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchPolicyListView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchPolicyListView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPolicyName

`func (o *PatchPolicyListView) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *PatchPolicyListView) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *PatchPolicyListView) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *PatchPolicyListView) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyEnabled

`func (o *PatchPolicyListView) GetPolicyEnabled() bool`

GetPolicyEnabled returns the PolicyEnabled field if non-nil, zero value otherwise.

### GetPolicyEnabledOk

`func (o *PatchPolicyListView) GetPolicyEnabledOk() (*bool, bool)`

GetPolicyEnabledOk returns a tuple with the PolicyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEnabled

`func (o *PatchPolicyListView) SetPolicyEnabled(v bool)`

SetPolicyEnabled sets PolicyEnabled field to given value.

### HasPolicyEnabled

`func (o *PatchPolicyListView) HasPolicyEnabled() bool`

HasPolicyEnabled returns a boolean if a field has been set.

### GetPolicyTargetVersion

`func (o *PatchPolicyListView) GetPolicyTargetVersion() string`

GetPolicyTargetVersion returns the PolicyTargetVersion field if non-nil, zero value otherwise.

### GetPolicyTargetVersionOk

`func (o *PatchPolicyListView) GetPolicyTargetVersionOk() (*string, bool)`

GetPolicyTargetVersionOk returns a tuple with the PolicyTargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTargetVersion

`func (o *PatchPolicyListView) SetPolicyTargetVersion(v string)`

SetPolicyTargetVersion sets PolicyTargetVersion field to given value.

### HasPolicyTargetVersion

`func (o *PatchPolicyListView) HasPolicyTargetVersion() bool`

HasPolicyTargetVersion returns a boolean if a field has been set.

### GetPolicyDeploymentMethod

`func (o *PatchPolicyListView) GetPolicyDeploymentMethod() string`

GetPolicyDeploymentMethod returns the PolicyDeploymentMethod field if non-nil, zero value otherwise.

### GetPolicyDeploymentMethodOk

`func (o *PatchPolicyListView) GetPolicyDeploymentMethodOk() (*string, bool)`

GetPolicyDeploymentMethodOk returns a tuple with the PolicyDeploymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDeploymentMethod

`func (o *PatchPolicyListView) SetPolicyDeploymentMethod(v string)`

SetPolicyDeploymentMethod sets PolicyDeploymentMethod field to given value.

### HasPolicyDeploymentMethod

`func (o *PatchPolicyListView) HasPolicyDeploymentMethod() bool`

HasPolicyDeploymentMethod returns a boolean if a field has been set.

### GetSoftwareTitle

`func (o *PatchPolicyListView) GetSoftwareTitle() string`

GetSoftwareTitle returns the SoftwareTitle field if non-nil, zero value otherwise.

### GetSoftwareTitleOk

`func (o *PatchPolicyListView) GetSoftwareTitleOk() (*string, bool)`

GetSoftwareTitleOk returns a tuple with the SoftwareTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTitle

`func (o *PatchPolicyListView) SetSoftwareTitle(v string)`

SetSoftwareTitle sets SoftwareTitle field to given value.

### HasSoftwareTitle

`func (o *PatchPolicyListView) HasSoftwareTitle() bool`

HasSoftwareTitle returns a boolean if a field has been set.

### GetSoftwareTitleConfigurationId

`func (o *PatchPolicyListView) GetSoftwareTitleConfigurationId() string`

GetSoftwareTitleConfigurationId returns the SoftwareTitleConfigurationId field if non-nil, zero value otherwise.

### GetSoftwareTitleConfigurationIdOk

`func (o *PatchPolicyListView) GetSoftwareTitleConfigurationIdOk() (*string, bool)`

GetSoftwareTitleConfigurationIdOk returns a tuple with the SoftwareTitleConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTitleConfigurationId

`func (o *PatchPolicyListView) SetSoftwareTitleConfigurationId(v string)`

SetSoftwareTitleConfigurationId sets SoftwareTitleConfigurationId field to given value.

### HasSoftwareTitleConfigurationId

`func (o *PatchPolicyListView) HasSoftwareTitleConfigurationId() bool`

HasSoftwareTitleConfigurationId returns a boolean if a field has been set.

### GetPending

`func (o *PatchPolicyListView) GetPending() int64`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *PatchPolicyListView) GetPendingOk() (*int64, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *PatchPolicyListView) SetPending(v int64)`

SetPending sets Pending field to given value.

### HasPending

`func (o *PatchPolicyListView) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetCompleted

`func (o *PatchPolicyListView) GetCompleted() int64`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *PatchPolicyListView) GetCompletedOk() (*int64, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *PatchPolicyListView) SetCompleted(v int64)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *PatchPolicyListView) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetDeferred

`func (o *PatchPolicyListView) GetDeferred() int64`

GetDeferred returns the Deferred field if non-nil, zero value otherwise.

### GetDeferredOk

`func (o *PatchPolicyListView) GetDeferredOk() (*int64, bool)`

GetDeferredOk returns a tuple with the Deferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferred

`func (o *PatchPolicyListView) SetDeferred(v int64)`

SetDeferred sets Deferred field to given value.

### HasDeferred

`func (o *PatchPolicyListView) HasDeferred() bool`

HasDeferred returns a boolean if a field has been set.

### GetFailed

`func (o *PatchPolicyListView) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *PatchPolicyListView) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *PatchPolicyListView) SetFailed(v int64)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *PatchPolicyListView) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


