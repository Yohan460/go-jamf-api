# ProtectUpdatableSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoInstall** | Pointer to **bool** | determines whether the Jamf Protect agent will be automatically installed on client computers | [optional] 

## Methods

### NewProtectUpdatableSettingsRequest

`func NewProtectUpdatableSettingsRequest() *ProtectUpdatableSettingsRequest`

NewProtectUpdatableSettingsRequest instantiates a new ProtectUpdatableSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectUpdatableSettingsRequestWithDefaults

`func NewProtectUpdatableSettingsRequestWithDefaults() *ProtectUpdatableSettingsRequest`

NewProtectUpdatableSettingsRequestWithDefaults instantiates a new ProtectUpdatableSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoInstall

`func (o *ProtectUpdatableSettingsRequest) GetAutoInstall() bool`

GetAutoInstall returns the AutoInstall field if non-nil, zero value otherwise.

### GetAutoInstallOk

`func (o *ProtectUpdatableSettingsRequest) GetAutoInstallOk() (*bool, bool)`

GetAutoInstallOk returns a tuple with the AutoInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoInstall

`func (o *ProtectUpdatableSettingsRequest) SetAutoInstall(v bool)`

SetAutoInstall sets AutoInstall field to given value.

### HasAutoInstall

`func (o *ProtectUpdatableSettingsRequest) HasAutoInstall() bool`

HasAutoInstall returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


