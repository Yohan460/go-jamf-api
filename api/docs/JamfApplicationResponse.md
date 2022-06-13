# JamfApplicationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**ReleaseHistoryUrl** | Pointer to **string** |  | [optional] 
**Artifacts** | Pointer to [**[]JamfPackageResponse**](JamfPackageResponse.md) |  | [optional] 

## Methods

### NewJamfApplicationResponse

`func NewJamfApplicationResponse() *JamfApplicationResponse`

NewJamfApplicationResponse instantiates a new JamfApplicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJamfApplicationResponseWithDefaults

`func NewJamfApplicationResponseWithDefaults() *JamfApplicationResponse`

NewJamfApplicationResponseWithDefaults instantiates a new JamfApplicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *JamfApplicationResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *JamfApplicationResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *JamfApplicationResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *JamfApplicationResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetReleaseHistoryUrl

`func (o *JamfApplicationResponse) GetReleaseHistoryUrl() string`

GetReleaseHistoryUrl returns the ReleaseHistoryUrl field if non-nil, zero value otherwise.

### GetReleaseHistoryUrlOk

`func (o *JamfApplicationResponse) GetReleaseHistoryUrlOk() (*string, bool)`

GetReleaseHistoryUrlOk returns a tuple with the ReleaseHistoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseHistoryUrl

`func (o *JamfApplicationResponse) SetReleaseHistoryUrl(v string)`

SetReleaseHistoryUrl sets ReleaseHistoryUrl field to given value.

### HasReleaseHistoryUrl

`func (o *JamfApplicationResponse) HasReleaseHistoryUrl() bool`

HasReleaseHistoryUrl returns a boolean if a field has been set.

### GetArtifacts

`func (o *JamfApplicationResponse) GetArtifacts() []JamfPackageResponse`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *JamfApplicationResponse) GetArtifactsOk() (*[]JamfPackageResponse, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *JamfApplicationResponse) SetArtifacts(v []JamfPackageResponse)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *JamfApplicationResponse) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


