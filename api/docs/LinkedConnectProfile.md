# LinkedConnectProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] [readonly] 
**ProfileId** | Pointer to **string** |  | [optional] [readonly] 
**ProfileName** | Pointer to **string** |  | [optional] [readonly] 
**ProfileScopeDescription** | Pointer to **string** |  | [optional] [readonly] 
**Version** | Pointer to **string** | Must be a valid Jamf Connect version 2.3.0 or higher. Versions are listed here &#x60;https://www.jamf.com/resources/product-documentation/jamf-connect-administrators-guide/&#x60; | [optional] 
**AutoDeploymentType** | Pointer to **string** | Determines how the server will behave regarding application updates and installs on the devices that have the configuration profile installed. * &#x60;PATCH_UPDATES&#x60; - Server handles initial installation of the application and any patch updates. * &#x60;MINOR_AND_PATCH_UPDATES&#x60; - Server handles initial installation of the application and any patch and minor updates. * &#x60;INITIAL_INSTALLATION_ONLY&#x60; - Server only handles initial installation of the application. Updates will have to be done manually. * &#x60;NONE&#x60; - Server does not handle any installations or updates for the application. Version is ignored for this type.  | [optional] [default to "NONE"]

## Methods

### NewLinkedConnectProfile

`func NewLinkedConnectProfile() *LinkedConnectProfile`

NewLinkedConnectProfile instantiates a new LinkedConnectProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedConnectProfileWithDefaults

`func NewLinkedConnectProfileWithDefaults() *LinkedConnectProfile`

NewLinkedConnectProfileWithDefaults instantiates a new LinkedConnectProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *LinkedConnectProfile) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *LinkedConnectProfile) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *LinkedConnectProfile) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *LinkedConnectProfile) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetProfileId

`func (o *LinkedConnectProfile) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *LinkedConnectProfile) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *LinkedConnectProfile) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *LinkedConnectProfile) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileName

`func (o *LinkedConnectProfile) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *LinkedConnectProfile) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *LinkedConnectProfile) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *LinkedConnectProfile) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### GetProfileScopeDescription

`func (o *LinkedConnectProfile) GetProfileScopeDescription() string`

GetProfileScopeDescription returns the ProfileScopeDescription field if non-nil, zero value otherwise.

### GetProfileScopeDescriptionOk

`func (o *LinkedConnectProfile) GetProfileScopeDescriptionOk() (*string, bool)`

GetProfileScopeDescriptionOk returns a tuple with the ProfileScopeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileScopeDescription

`func (o *LinkedConnectProfile) SetProfileScopeDescription(v string)`

SetProfileScopeDescription sets ProfileScopeDescription field to given value.

### HasProfileScopeDescription

`func (o *LinkedConnectProfile) HasProfileScopeDescription() bool`

HasProfileScopeDescription returns a boolean if a field has been set.

### GetVersion

`func (o *LinkedConnectProfile) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LinkedConnectProfile) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LinkedConnectProfile) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LinkedConnectProfile) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAutoDeploymentType

`func (o *LinkedConnectProfile) GetAutoDeploymentType() string`

GetAutoDeploymentType returns the AutoDeploymentType field if non-nil, zero value otherwise.

### GetAutoDeploymentTypeOk

`func (o *LinkedConnectProfile) GetAutoDeploymentTypeOk() (*string, bool)`

GetAutoDeploymentTypeOk returns a tuple with the AutoDeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploymentType

`func (o *LinkedConnectProfile) SetAutoDeploymentType(v string)`

SetAutoDeploymentType sets AutoDeploymentType field to given value.

### HasAutoDeploymentType

`func (o *LinkedConnectProfile) HasAutoDeploymentType() bool`

HasAutoDeploymentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


