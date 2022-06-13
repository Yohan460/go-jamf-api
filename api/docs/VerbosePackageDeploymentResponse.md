# VerbosePackageDeploymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueuedCommands** | Pointer to [**[]VerbosePackageDeploymentResponseQueuedCommandsInner**](VerbosePackageDeploymentResponseQueuedCommandsInner.md) |  | [optional] 
**Errors** | Pointer to [**[]VerbosePackageDeploymentResponseErrorsInner**](VerbosePackageDeploymentResponseErrorsInner.md) |  | [optional] 

## Methods

### NewVerbosePackageDeploymentResponse

`func NewVerbosePackageDeploymentResponse() *VerbosePackageDeploymentResponse`

NewVerbosePackageDeploymentResponse instantiates a new VerbosePackageDeploymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerbosePackageDeploymentResponseWithDefaults

`func NewVerbosePackageDeploymentResponseWithDefaults() *VerbosePackageDeploymentResponse`

NewVerbosePackageDeploymentResponseWithDefaults instantiates a new VerbosePackageDeploymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueuedCommands

`func (o *VerbosePackageDeploymentResponse) GetQueuedCommands() []VerbosePackageDeploymentResponseQueuedCommandsInner`

GetQueuedCommands returns the QueuedCommands field if non-nil, zero value otherwise.

### GetQueuedCommandsOk

`func (o *VerbosePackageDeploymentResponse) GetQueuedCommandsOk() (*[]VerbosePackageDeploymentResponseQueuedCommandsInner, bool)`

GetQueuedCommandsOk returns a tuple with the QueuedCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedCommands

`func (o *VerbosePackageDeploymentResponse) SetQueuedCommands(v []VerbosePackageDeploymentResponseQueuedCommandsInner)`

SetQueuedCommands sets QueuedCommands field to given value.

### HasQueuedCommands

`func (o *VerbosePackageDeploymentResponse) HasQueuedCommands() bool`

HasQueuedCommands returns a boolean if a field has been set.

### GetErrors

`func (o *VerbosePackageDeploymentResponse) GetErrors() []VerbosePackageDeploymentResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *VerbosePackageDeploymentResponse) GetErrorsOk() (*[]VerbosePackageDeploymentResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *VerbosePackageDeploymentResponse) SetErrors(v []VerbosePackageDeploymentResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *VerbosePackageDeploymentResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


