# MacOsManagedSoftwareUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessManagerUuids** | Pointer to **[]string** |  | [optional] 
**Errors** | Pointer to [**[]ApiErrorCause**](ApiErrorCause.md) |  | [optional] 

## Methods

### NewMacOsManagedSoftwareUpdateResponse

`func NewMacOsManagedSoftwareUpdateResponse() *MacOsManagedSoftwareUpdateResponse`

NewMacOsManagedSoftwareUpdateResponse instantiates a new MacOsManagedSoftwareUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacOsManagedSoftwareUpdateResponseWithDefaults

`func NewMacOsManagedSoftwareUpdateResponseWithDefaults() *MacOsManagedSoftwareUpdateResponse`

NewMacOsManagedSoftwareUpdateResponseWithDefaults instantiates a new MacOsManagedSoftwareUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessManagerUuids

`func (o *MacOsManagedSoftwareUpdateResponse) GetProcessManagerUuids() []string`

GetProcessManagerUuids returns the ProcessManagerUuids field if non-nil, zero value otherwise.

### GetProcessManagerUuidsOk

`func (o *MacOsManagedSoftwareUpdateResponse) GetProcessManagerUuidsOk() (*[]string, bool)`

GetProcessManagerUuidsOk returns a tuple with the ProcessManagerUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessManagerUuids

`func (o *MacOsManagedSoftwareUpdateResponse) SetProcessManagerUuids(v []string)`

SetProcessManagerUuids sets ProcessManagerUuids field to given value.

### HasProcessManagerUuids

`func (o *MacOsManagedSoftwareUpdateResponse) HasProcessManagerUuids() bool`

HasProcessManagerUuids returns a boolean if a field has been set.

### GetErrors

`func (o *MacOsManagedSoftwareUpdateResponse) GetErrors() []ApiErrorCause`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MacOsManagedSoftwareUpdateResponse) GetErrorsOk() (*[]ApiErrorCause, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MacOsManagedSoftwareUpdateResponse) SetErrors(v []ApiErrorCause)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *MacOsManagedSoftwareUpdateResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


