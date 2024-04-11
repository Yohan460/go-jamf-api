# InstallPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Manifest** | [**PackageManifest**](PackageManifest.md) |  | 
**InstallAsManaged** | Pointer to **bool** |  | [optional] 
**Devices** | Pointer to **[]int64** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 

## Methods

### NewInstallPackage

`func NewInstallPackage(manifest PackageManifest, ) *InstallPackage`

NewInstallPackage instantiates a new InstallPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallPackageWithDefaults

`func NewInstallPackageWithDefaults() *InstallPackage`

NewInstallPackageWithDefaults instantiates a new InstallPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManifest

`func (o *InstallPackage) GetManifest() PackageManifest`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *InstallPackage) GetManifestOk() (*PackageManifest, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *InstallPackage) SetManifest(v PackageManifest)`

SetManifest sets Manifest field to given value.


### GetInstallAsManaged

`func (o *InstallPackage) GetInstallAsManaged() bool`

GetInstallAsManaged returns the InstallAsManaged field if non-nil, zero value otherwise.

### GetInstallAsManagedOk

`func (o *InstallPackage) GetInstallAsManagedOk() (*bool, bool)`

GetInstallAsManagedOk returns a tuple with the InstallAsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAsManaged

`func (o *InstallPackage) SetInstallAsManaged(v bool)`

SetInstallAsManaged sets InstallAsManaged field to given value.

### HasInstallAsManaged

`func (o *InstallPackage) HasInstallAsManaged() bool`

HasInstallAsManaged returns a boolean if a field has been set.

### GetDevices

`func (o *InstallPackage) GetDevices() []int64`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InstallPackage) GetDevicesOk() (*[]int64, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InstallPackage) SetDevices(v []int64)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InstallPackage) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetGroupId

`func (o *InstallPackage) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *InstallPackage) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *InstallPackage) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *InstallPackage) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


