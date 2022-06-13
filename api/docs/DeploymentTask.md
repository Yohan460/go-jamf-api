# DeploymentTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ComputerId** | Pointer to **string** |  | [optional] [readonly] 
**ComputerName** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] [readonly] 
**Updated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Status** | Pointer to **string** | Status of this Jamf Connect deployment task. \&quot;Command\&quot; below refers to an &#x60;InstallEnterpriseApplication&#x60; command. Tasks that are not finished (i.e., &#x60;COMPLETE&#x60; or &#x60;GAVE_UP&#x60;) are evaluated once every thirty minutes, so the status value for a device may lag behind a successful Jamf Connect package install up to thirty minutes. * &#x60;COMMAND_QUEUED&#x60; - command has been queued * &#x60;NO_COMMAND&#x60; - command has not yet been queued * &#x60;PENDING_MANIFEST&#x60; - task is waiting to obtain a valid package manifest before a command can be queued * &#x60;COMPLETE&#x60; - command has been completed successfully * &#x60;GAVE_UP&#x60; - the command failed with an error or the device did not process it in a reasonable amount of time * &#x60;UNKNOWN&#x60; - unknown; tasks in this state will be evaluated  | [optional] 

## Methods

### NewDeploymentTask

`func NewDeploymentTask() *DeploymentTask`

NewDeploymentTask instantiates a new DeploymentTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentTaskWithDefaults

`func NewDeploymentTaskWithDefaults() *DeploymentTask`

NewDeploymentTaskWithDefaults instantiates a new DeploymentTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentTask) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetComputerId

`func (o *DeploymentTask) GetComputerId() string`

GetComputerId returns the ComputerId field if non-nil, zero value otherwise.

### GetComputerIdOk

`func (o *DeploymentTask) GetComputerIdOk() (*string, bool)`

GetComputerIdOk returns a tuple with the ComputerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerId

`func (o *DeploymentTask) SetComputerId(v string)`

SetComputerId sets ComputerId field to given value.

### HasComputerId

`func (o *DeploymentTask) HasComputerId() bool`

HasComputerId returns a boolean if a field has been set.

### GetComputerName

`func (o *DeploymentTask) GetComputerName() string`

GetComputerName returns the ComputerName field if non-nil, zero value otherwise.

### GetComputerNameOk

`func (o *DeploymentTask) GetComputerNameOk() (*string, bool)`

GetComputerNameOk returns a tuple with the ComputerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerName

`func (o *DeploymentTask) SetComputerName(v string)`

SetComputerName sets ComputerName field to given value.

### HasComputerName

`func (o *DeploymentTask) HasComputerName() bool`

HasComputerName returns a boolean if a field has been set.

### GetVersion

`func (o *DeploymentTask) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeploymentTask) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeploymentTask) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeploymentTask) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUpdated

`func (o *DeploymentTask) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *DeploymentTask) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *DeploymentTask) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *DeploymentTask) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


