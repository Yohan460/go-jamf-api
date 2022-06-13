# ActivePatchSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SoftwareTitleID** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**LatestVersion** | Pointer to **string** |  | [optional] 
**ReleaseDate** | Pointer to **time.Time** |  | [optional] 
**UpToDate** | Pointer to **int32** |  | [optional] 
**OutOfDate** | Pointer to **int32** |  | [optional] 
**IsOnDashboard** | Pointer to **bool** |  | [optional] 
**SoftwareTitleConfigurationID** | Pointer to **int32** |  | [optional] 

## Methods

### NewActivePatchSummary

`func NewActivePatchSummary() *ActivePatchSummary`

NewActivePatchSummary instantiates a new ActivePatchSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivePatchSummaryWithDefaults

`func NewActivePatchSummaryWithDefaults() *ActivePatchSummary`

NewActivePatchSummaryWithDefaults instantiates a new ActivePatchSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSoftwareTitleID

`func (o *ActivePatchSummary) GetSoftwareTitleID() int32`

GetSoftwareTitleID returns the SoftwareTitleID field if non-nil, zero value otherwise.

### GetSoftwareTitleIDOk

`func (o *ActivePatchSummary) GetSoftwareTitleIDOk() (*int32, bool)`

GetSoftwareTitleIDOk returns a tuple with the SoftwareTitleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTitleID

`func (o *ActivePatchSummary) SetSoftwareTitleID(v int32)`

SetSoftwareTitleID sets SoftwareTitleID field to given value.

### HasSoftwareTitleID

`func (o *ActivePatchSummary) HasSoftwareTitleID() bool`

HasSoftwareTitleID returns a boolean if a field has been set.

### GetTitle

`func (o *ActivePatchSummary) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ActivePatchSummary) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ActivePatchSummary) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ActivePatchSummary) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetLatestVersion

`func (o *ActivePatchSummary) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *ActivePatchSummary) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *ActivePatchSummary) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *ActivePatchSummary) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetReleaseDate

`func (o *ActivePatchSummary) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *ActivePatchSummary) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *ActivePatchSummary) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *ActivePatchSummary) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetUpToDate

`func (o *ActivePatchSummary) GetUpToDate() int32`

GetUpToDate returns the UpToDate field if non-nil, zero value otherwise.

### GetUpToDateOk

`func (o *ActivePatchSummary) GetUpToDateOk() (*int32, bool)`

GetUpToDateOk returns a tuple with the UpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDate

`func (o *ActivePatchSummary) SetUpToDate(v int32)`

SetUpToDate sets UpToDate field to given value.

### HasUpToDate

`func (o *ActivePatchSummary) HasUpToDate() bool`

HasUpToDate returns a boolean if a field has been set.

### GetOutOfDate

`func (o *ActivePatchSummary) GetOutOfDate() int32`

GetOutOfDate returns the OutOfDate field if non-nil, zero value otherwise.

### GetOutOfDateOk

`func (o *ActivePatchSummary) GetOutOfDateOk() (*int32, bool)`

GetOutOfDateOk returns a tuple with the OutOfDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfDate

`func (o *ActivePatchSummary) SetOutOfDate(v int32)`

SetOutOfDate sets OutOfDate field to given value.

### HasOutOfDate

`func (o *ActivePatchSummary) HasOutOfDate() bool`

HasOutOfDate returns a boolean if a field has been set.

### GetIsOnDashboard

`func (o *ActivePatchSummary) GetIsOnDashboard() bool`

GetIsOnDashboard returns the IsOnDashboard field if non-nil, zero value otherwise.

### GetIsOnDashboardOk

`func (o *ActivePatchSummary) GetIsOnDashboardOk() (*bool, bool)`

GetIsOnDashboardOk returns a tuple with the IsOnDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnDashboard

`func (o *ActivePatchSummary) SetIsOnDashboard(v bool)`

SetIsOnDashboard sets IsOnDashboard field to given value.

### HasIsOnDashboard

`func (o *ActivePatchSummary) HasIsOnDashboard() bool`

HasIsOnDashboard returns a boolean if a field has been set.

### GetSoftwareTitleConfigurationID

`func (o *ActivePatchSummary) GetSoftwareTitleConfigurationID() int32`

GetSoftwareTitleConfigurationID returns the SoftwareTitleConfigurationID field if non-nil, zero value otherwise.

### GetSoftwareTitleConfigurationIDOk

`func (o *ActivePatchSummary) GetSoftwareTitleConfigurationIDOk() (*int32, bool)`

GetSoftwareTitleConfigurationIDOk returns a tuple with the SoftwareTitleConfigurationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTitleConfigurationID

`func (o *ActivePatchSummary) SetSoftwareTitleConfigurationID(v int32)`

SetSoftwareTitleConfigurationID sets SoftwareTitleConfigurationID field to given value.

### HasSoftwareTitleConfigurationID

`func (o *ActivePatchSummary) HasSoftwareTitleConfigurationID() bool`

HasSoftwareTitleConfigurationID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


