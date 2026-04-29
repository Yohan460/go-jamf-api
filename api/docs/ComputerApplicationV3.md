# ComputerApplicationV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**CfBundleShortVersionString** | Pointer to **string** |  | [optional] 
**CfBundleVersion** | Pointer to **string** |  | [optional] 
**MacAppStore** | Pointer to **bool** |  | [optional] 
**SizeMegabytes** | Pointer to **int64** |  | [optional] 
**BundleId** | Pointer to **string** |  | [optional] 
**UpdateAvailable** | Pointer to **bool** |  | [optional] 
**ExternalVersionId** | Pointer to **string** | The app&#39;s external version ID. It can be used in the iTunes Search API to decide if the app needs to be updated | [optional] 

## Methods

### NewComputerApplicationV3

`func NewComputerApplicationV3() *ComputerApplicationV3`

NewComputerApplicationV3 instantiates a new ComputerApplicationV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerApplicationV3WithDefaults

`func NewComputerApplicationV3WithDefaults() *ComputerApplicationV3`

NewComputerApplicationV3WithDefaults instantiates a new ComputerApplicationV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ComputerApplicationV3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputerApplicationV3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputerApplicationV3) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputerApplicationV3) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *ComputerApplicationV3) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ComputerApplicationV3) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ComputerApplicationV3) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ComputerApplicationV3) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetVersion

`func (o *ComputerApplicationV3) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ComputerApplicationV3) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ComputerApplicationV3) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ComputerApplicationV3) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCfBundleShortVersionString

`func (o *ComputerApplicationV3) GetCfBundleShortVersionString() string`

GetCfBundleShortVersionString returns the CfBundleShortVersionString field if non-nil, zero value otherwise.

### GetCfBundleShortVersionStringOk

`func (o *ComputerApplicationV3) GetCfBundleShortVersionStringOk() (*string, bool)`

GetCfBundleShortVersionStringOk returns a tuple with the CfBundleShortVersionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfBundleShortVersionString

`func (o *ComputerApplicationV3) SetCfBundleShortVersionString(v string)`

SetCfBundleShortVersionString sets CfBundleShortVersionString field to given value.

### HasCfBundleShortVersionString

`func (o *ComputerApplicationV3) HasCfBundleShortVersionString() bool`

HasCfBundleShortVersionString returns a boolean if a field has been set.

### GetCfBundleVersion

`func (o *ComputerApplicationV3) GetCfBundleVersion() string`

GetCfBundleVersion returns the CfBundleVersion field if non-nil, zero value otherwise.

### GetCfBundleVersionOk

`func (o *ComputerApplicationV3) GetCfBundleVersionOk() (*string, bool)`

GetCfBundleVersionOk returns a tuple with the CfBundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfBundleVersion

`func (o *ComputerApplicationV3) SetCfBundleVersion(v string)`

SetCfBundleVersion sets CfBundleVersion field to given value.

### HasCfBundleVersion

`func (o *ComputerApplicationV3) HasCfBundleVersion() bool`

HasCfBundleVersion returns a boolean if a field has been set.

### GetMacAppStore

`func (o *ComputerApplicationV3) GetMacAppStore() bool`

GetMacAppStore returns the MacAppStore field if non-nil, zero value otherwise.

### GetMacAppStoreOk

`func (o *ComputerApplicationV3) GetMacAppStoreOk() (*bool, bool)`

GetMacAppStoreOk returns a tuple with the MacAppStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAppStore

`func (o *ComputerApplicationV3) SetMacAppStore(v bool)`

SetMacAppStore sets MacAppStore field to given value.

### HasMacAppStore

`func (o *ComputerApplicationV3) HasMacAppStore() bool`

HasMacAppStore returns a boolean if a field has been set.

### GetSizeMegabytes

`func (o *ComputerApplicationV3) GetSizeMegabytes() int64`

GetSizeMegabytes returns the SizeMegabytes field if non-nil, zero value otherwise.

### GetSizeMegabytesOk

`func (o *ComputerApplicationV3) GetSizeMegabytesOk() (*int64, bool)`

GetSizeMegabytesOk returns a tuple with the SizeMegabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMegabytes

`func (o *ComputerApplicationV3) SetSizeMegabytes(v int64)`

SetSizeMegabytes sets SizeMegabytes field to given value.

### HasSizeMegabytes

`func (o *ComputerApplicationV3) HasSizeMegabytes() bool`

HasSizeMegabytes returns a boolean if a field has been set.

### GetBundleId

`func (o *ComputerApplicationV3) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *ComputerApplicationV3) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *ComputerApplicationV3) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *ComputerApplicationV3) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### GetUpdateAvailable

`func (o *ComputerApplicationV3) GetUpdateAvailable() bool`

GetUpdateAvailable returns the UpdateAvailable field if non-nil, zero value otherwise.

### GetUpdateAvailableOk

`func (o *ComputerApplicationV3) GetUpdateAvailableOk() (*bool, bool)`

GetUpdateAvailableOk returns a tuple with the UpdateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAvailable

`func (o *ComputerApplicationV3) SetUpdateAvailable(v bool)`

SetUpdateAvailable sets UpdateAvailable field to given value.

### HasUpdateAvailable

`func (o *ComputerApplicationV3) HasUpdateAvailable() bool`

HasUpdateAvailable returns a boolean if a field has been set.

### GetExternalVersionId

`func (o *ComputerApplicationV3) GetExternalVersionId() string`

GetExternalVersionId returns the ExternalVersionId field if non-nil, zero value otherwise.

### GetExternalVersionIdOk

`func (o *ComputerApplicationV3) GetExternalVersionIdOk() (*string, bool)`

GetExternalVersionIdOk returns a tuple with the ExternalVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalVersionId

`func (o *ComputerApplicationV3) SetExternalVersionId(v string)`

SetExternalVersionId sets ExternalVersionId field to given value.

### HasExternalVersionId

`func (o *ComputerApplicationV3) HasExternalVersionId() bool`

HasExternalVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


