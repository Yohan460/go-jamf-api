# SelfServiceInstallSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallAutomatically** | Pointer to **bool** | true if Self Service is insalled automatically, false if not  | [optional] [default to false]
**InstallLocation** | **string** | path at which Self Service is installed  | 

## Methods

### NewSelfServiceInstallSettings

`func NewSelfServiceInstallSettings(installLocation string, ) *SelfServiceInstallSettings`

NewSelfServiceInstallSettings instantiates a new SelfServiceInstallSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfServiceInstallSettingsWithDefaults

`func NewSelfServiceInstallSettingsWithDefaults() *SelfServiceInstallSettings`

NewSelfServiceInstallSettingsWithDefaults instantiates a new SelfServiceInstallSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallAutomatically

`func (o *SelfServiceInstallSettings) GetInstallAutomatically() bool`

GetInstallAutomatically returns the InstallAutomatically field if non-nil, zero value otherwise.

### GetInstallAutomaticallyOk

`func (o *SelfServiceInstallSettings) GetInstallAutomaticallyOk() (*bool, bool)`

GetInstallAutomaticallyOk returns a tuple with the InstallAutomatically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAutomatically

`func (o *SelfServiceInstallSettings) SetInstallAutomatically(v bool)`

SetInstallAutomatically sets InstallAutomatically field to given value.

### HasInstallAutomatically

`func (o *SelfServiceInstallSettings) HasInstallAutomatically() bool`

HasInstallAutomatically returns a boolean if a field has been set.

### GetInstallLocation

`func (o *SelfServiceInstallSettings) GetInstallLocation() string`

GetInstallLocation returns the InstallLocation field if non-nil, zero value otherwise.

### GetInstallLocationOk

`func (o *SelfServiceInstallSettings) GetInstallLocationOk() (*string, bool)`

GetInstallLocationOk returns a tuple with the InstallLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallLocation

`func (o *SelfServiceInstallSettings) SetInstallLocation(v string)`

SetInstallLocation sets InstallLocation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


