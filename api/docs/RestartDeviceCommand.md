# RestartDeviceCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RebuildKernelCache** | Pointer to **bool** |  | [optional] 
**KextPaths** | Pointer to **[]string** | Only used if RebuildKernelCache is true | [optional] 
**NotifyUser** | Pointer to **bool** |  | [optional] 

## Methods

### NewRestartDeviceCommand

`func NewRestartDeviceCommand() *RestartDeviceCommand`

NewRestartDeviceCommand instantiates a new RestartDeviceCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestartDeviceCommandWithDefaults

`func NewRestartDeviceCommandWithDefaults() *RestartDeviceCommand`

NewRestartDeviceCommandWithDefaults instantiates a new RestartDeviceCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRebuildKernelCache

`func (o *RestartDeviceCommand) GetRebuildKernelCache() bool`

GetRebuildKernelCache returns the RebuildKernelCache field if non-nil, zero value otherwise.

### GetRebuildKernelCacheOk

`func (o *RestartDeviceCommand) GetRebuildKernelCacheOk() (*bool, bool)`

GetRebuildKernelCacheOk returns a tuple with the RebuildKernelCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebuildKernelCache

`func (o *RestartDeviceCommand) SetRebuildKernelCache(v bool)`

SetRebuildKernelCache sets RebuildKernelCache field to given value.

### HasRebuildKernelCache

`func (o *RestartDeviceCommand) HasRebuildKernelCache() bool`

HasRebuildKernelCache returns a boolean if a field has been set.

### GetKextPaths

`func (o *RestartDeviceCommand) GetKextPaths() []string`

GetKextPaths returns the KextPaths field if non-nil, zero value otherwise.

### GetKextPathsOk

`func (o *RestartDeviceCommand) GetKextPathsOk() (*[]string, bool)`

GetKextPathsOk returns a tuple with the KextPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKextPaths

`func (o *RestartDeviceCommand) SetKextPaths(v []string)`

SetKextPaths sets KextPaths field to given value.

### HasKextPaths

`func (o *RestartDeviceCommand) HasKextPaths() bool`

HasKextPaths returns a boolean if a field has been set.

### GetNotifyUser

`func (o *RestartDeviceCommand) GetNotifyUser() bool`

GetNotifyUser returns the NotifyUser field if non-nil, zero value otherwise.

### GetNotifyUserOk

`func (o *RestartDeviceCommand) GetNotifyUserOk() (*bool, bool)`

GetNotifyUserOk returns a tuple with the NotifyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUser

`func (o *RestartDeviceCommand) SetNotifyUser(v bool)`

SetNotifyUser sets NotifyUser field to given value.

### HasNotifyUser

`func (o *RestartDeviceCommand) HasNotifyUser() bool`

HasNotifyUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


