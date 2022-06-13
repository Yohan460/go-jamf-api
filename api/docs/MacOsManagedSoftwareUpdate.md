# MacOsManagedSoftwareUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceIds** | **[]string** |  | 
**MaxDeferrals** | Pointer to **int32** | Allow users to defer the update the provided number of times before macOS forces the update. If a value is provided, the Software Update will use the InstallLater install action. | [optional] 
**Version** | Pointer to **string** | If no value is provided, the version will default to latest version based on device eligibility. | [optional] 
**UpdateAction** | Pointer to **string** | MaxDeferral is ignored if using the DownloadOnly install action. | [optional] 

## Methods

### NewMacOsManagedSoftwareUpdate

`func NewMacOsManagedSoftwareUpdate(deviceIds []string, ) *MacOsManagedSoftwareUpdate`

NewMacOsManagedSoftwareUpdate instantiates a new MacOsManagedSoftwareUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacOsManagedSoftwareUpdateWithDefaults

`func NewMacOsManagedSoftwareUpdateWithDefaults() *MacOsManagedSoftwareUpdate`

NewMacOsManagedSoftwareUpdateWithDefaults instantiates a new MacOsManagedSoftwareUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceIds

`func (o *MacOsManagedSoftwareUpdate) GetDeviceIds() []string`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *MacOsManagedSoftwareUpdate) GetDeviceIdsOk() (*[]string, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *MacOsManagedSoftwareUpdate) SetDeviceIds(v []string)`

SetDeviceIds sets DeviceIds field to given value.


### GetMaxDeferrals

`func (o *MacOsManagedSoftwareUpdate) GetMaxDeferrals() int32`

GetMaxDeferrals returns the MaxDeferrals field if non-nil, zero value otherwise.

### GetMaxDeferralsOk

`func (o *MacOsManagedSoftwareUpdate) GetMaxDeferralsOk() (*int32, bool)`

GetMaxDeferralsOk returns a tuple with the MaxDeferrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeferrals

`func (o *MacOsManagedSoftwareUpdate) SetMaxDeferrals(v int32)`

SetMaxDeferrals sets MaxDeferrals field to given value.

### HasMaxDeferrals

`func (o *MacOsManagedSoftwareUpdate) HasMaxDeferrals() bool`

HasMaxDeferrals returns a boolean if a field has been set.

### GetVersion

`func (o *MacOsManagedSoftwareUpdate) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MacOsManagedSoftwareUpdate) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MacOsManagedSoftwareUpdate) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MacOsManagedSoftwareUpdate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUpdateAction

`func (o *MacOsManagedSoftwareUpdate) GetUpdateAction() string`

GetUpdateAction returns the UpdateAction field if non-nil, zero value otherwise.

### GetUpdateActionOk

`func (o *MacOsManagedSoftwareUpdate) GetUpdateActionOk() (*string, bool)`

GetUpdateActionOk returns a tuple with the UpdateAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAction

`func (o *MacOsManagedSoftwareUpdate) SetUpdateAction(v string)`

SetUpdateAction sets UpdateAction field to given value.

### HasUpdateAction

`func (o *MacOsManagedSoftwareUpdate) HasUpdateAction() bool`

HasUpdateAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


