# PatchPolicySummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyId** | Pointer to **int32** |  | [optional] 
**PolicyName** | Pointer to **string** |  | [optional] 
**IsPolicyEnabled** | Pointer to **bool** |  | [optional] 
**PolicyTargetVersion** | Pointer to **string** |  | [optional] 
**PolicyDeploymentMethod** | Pointer to **string** |  | [optional] 
**SoftwareTitle** | Pointer to **string** |  | [optional] 
**SoftwareTitleConfigurationId** | Pointer to **int32** |  | [optional] 
**Pending** | Pointer to **int32** |  | [optional] 
**Completed** | Pointer to **int32** |  | [optional] 
**Deferred** | Pointer to **int32** |  | [optional] 
**Failed** | Pointer to **int32** |  | [optional] 

## Methods

### NewPatchPolicySummary

`func NewPatchPolicySummary() *PatchPolicySummary`

NewPatchPolicySummary instantiates a new PatchPolicySummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPolicySummaryWithDefaults

`func NewPatchPolicySummaryWithDefaults() *PatchPolicySummary`

NewPatchPolicySummaryWithDefaults instantiates a new PatchPolicySummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyId

`func (o *PatchPolicySummary) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *PatchPolicySummary) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *PatchPolicySummary) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *PatchPolicySummary) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *PatchPolicySummary) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *PatchPolicySummary) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *PatchPolicySummary) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *PatchPolicySummary) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetIsPolicyEnabled

`func (o *PatchPolicySummary) GetIsPolicyEnabled() bool`

GetIsPolicyEnabled returns the IsPolicyEnabled field if non-nil, zero value otherwise.

### GetIsPolicyEnabledOk

`func (o *PatchPolicySummary) GetIsPolicyEnabledOk() (*bool, bool)`

GetIsPolicyEnabledOk returns a tuple with the IsPolicyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPolicyEnabled

`func (o *PatchPolicySummary) SetIsPolicyEnabled(v bool)`

SetIsPolicyEnabled sets IsPolicyEnabled field to given value.

### HasIsPolicyEnabled

`func (o *PatchPolicySummary) HasIsPolicyEnabled() bool`

HasIsPolicyEnabled returns a boolean if a field has been set.

### GetPolicyTargetVersion

`func (o *PatchPolicySummary) GetPolicyTargetVersion() string`

GetPolicyTargetVersion returns the PolicyTargetVersion field if non-nil, zero value otherwise.

### GetPolicyTargetVersionOk

`func (o *PatchPolicySummary) GetPolicyTargetVersionOk() (*string, bool)`

GetPolicyTargetVersionOk returns a tuple with the PolicyTargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTargetVersion

`func (o *PatchPolicySummary) SetPolicyTargetVersion(v string)`

SetPolicyTargetVersion sets PolicyTargetVersion field to given value.

### HasPolicyTargetVersion

`func (o *PatchPolicySummary) HasPolicyTargetVersion() bool`

HasPolicyTargetVersion returns a boolean if a field has been set.

### GetPolicyDeploymentMethod

`func (o *PatchPolicySummary) GetPolicyDeploymentMethod() string`

GetPolicyDeploymentMethod returns the PolicyDeploymentMethod field if non-nil, zero value otherwise.

### GetPolicyDeploymentMethodOk

`func (o *PatchPolicySummary) GetPolicyDeploymentMethodOk() (*string, bool)`

GetPolicyDeploymentMethodOk returns a tuple with the PolicyDeploymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDeploymentMethod

`func (o *PatchPolicySummary) SetPolicyDeploymentMethod(v string)`

SetPolicyDeploymentMethod sets PolicyDeploymentMethod field to given value.

### HasPolicyDeploymentMethod

`func (o *PatchPolicySummary) HasPolicyDeploymentMethod() bool`

HasPolicyDeploymentMethod returns a boolean if a field has been set.

### GetSoftwareTitle

`func (o *PatchPolicySummary) GetSoftwareTitle() string`

GetSoftwareTitle returns the SoftwareTitle field if non-nil, zero value otherwise.

### GetSoftwareTitleOk

`func (o *PatchPolicySummary) GetSoftwareTitleOk() (*string, bool)`

GetSoftwareTitleOk returns a tuple with the SoftwareTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTitle

`func (o *PatchPolicySummary) SetSoftwareTitle(v string)`

SetSoftwareTitle sets SoftwareTitle field to given value.

### HasSoftwareTitle

`func (o *PatchPolicySummary) HasSoftwareTitle() bool`

HasSoftwareTitle returns a boolean if a field has been set.

### GetSoftwareTitleConfigurationId

`func (o *PatchPolicySummary) GetSoftwareTitleConfigurationId() int32`

GetSoftwareTitleConfigurationId returns the SoftwareTitleConfigurationId field if non-nil, zero value otherwise.

### GetSoftwareTitleConfigurationIdOk

`func (o *PatchPolicySummary) GetSoftwareTitleConfigurationIdOk() (*int32, bool)`

GetSoftwareTitleConfigurationIdOk returns a tuple with the SoftwareTitleConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTitleConfigurationId

`func (o *PatchPolicySummary) SetSoftwareTitleConfigurationId(v int32)`

SetSoftwareTitleConfigurationId sets SoftwareTitleConfigurationId field to given value.

### HasSoftwareTitleConfigurationId

`func (o *PatchPolicySummary) HasSoftwareTitleConfigurationId() bool`

HasSoftwareTitleConfigurationId returns a boolean if a field has been set.

### GetPending

`func (o *PatchPolicySummary) GetPending() int32`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *PatchPolicySummary) GetPendingOk() (*int32, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *PatchPolicySummary) SetPending(v int32)`

SetPending sets Pending field to given value.

### HasPending

`func (o *PatchPolicySummary) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetCompleted

`func (o *PatchPolicySummary) GetCompleted() int32`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *PatchPolicySummary) GetCompletedOk() (*int32, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *PatchPolicySummary) SetCompleted(v int32)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *PatchPolicySummary) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetDeferred

`func (o *PatchPolicySummary) GetDeferred() int32`

GetDeferred returns the Deferred field if non-nil, zero value otherwise.

### GetDeferredOk

`func (o *PatchPolicySummary) GetDeferredOk() (*int32, bool)`

GetDeferredOk returns a tuple with the Deferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferred

`func (o *PatchPolicySummary) SetDeferred(v int32)`

SetDeferred sets Deferred field to given value.

### HasDeferred

`func (o *PatchPolicySummary) HasDeferred() bool`

HasDeferred returns a boolean if a field has been set.

### GetFailed

`func (o *PatchPolicySummary) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *PatchPolicySummary) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *PatchPolicySummary) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *PatchPolicySummary) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


