# PatchSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SoftwareTitleId** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**LatestVersion** | Pointer to **string** |  | [optional] 
**ReleaseDate** | Pointer to **time.Time** |  | [optional] 
**UpToDate** | Pointer to **int64** |  | [optional] 
**OutOfDate** | Pointer to **int64** |  | [optional] 
**OnDashboard** | Pointer to **bool** |  | [optional] 
**SoftwareTitleConfigurationId** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchSummary

`func NewPatchSummary() *PatchSummary`

NewPatchSummary instantiates a new PatchSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchSummaryWithDefaults

`func NewPatchSummaryWithDefaults() *PatchSummary`

NewPatchSummaryWithDefaults instantiates a new PatchSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSoftwareTitleId

`func (o *PatchSummary) GetSoftwareTitleId() string`

GetSoftwareTitleId returns the SoftwareTitleId field if non-nil, zero value otherwise.

### GetSoftwareTitleIdOk

`func (o *PatchSummary) GetSoftwareTitleIdOk() (*string, bool)`

GetSoftwareTitleIdOk returns a tuple with the SoftwareTitleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTitleId

`func (o *PatchSummary) SetSoftwareTitleId(v string)`

SetSoftwareTitleId sets SoftwareTitleId field to given value.

### HasSoftwareTitleId

`func (o *PatchSummary) HasSoftwareTitleId() bool`

HasSoftwareTitleId returns a boolean if a field has been set.

### GetTitle

`func (o *PatchSummary) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PatchSummary) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PatchSummary) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PatchSummary) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetLatestVersion

`func (o *PatchSummary) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *PatchSummary) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *PatchSummary) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *PatchSummary) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetReleaseDate

`func (o *PatchSummary) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *PatchSummary) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *PatchSummary) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *PatchSummary) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetUpToDate

`func (o *PatchSummary) GetUpToDate() int64`

GetUpToDate returns the UpToDate field if non-nil, zero value otherwise.

### GetUpToDateOk

`func (o *PatchSummary) GetUpToDateOk() (*int64, bool)`

GetUpToDateOk returns a tuple with the UpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToDate

`func (o *PatchSummary) SetUpToDate(v int64)`

SetUpToDate sets UpToDate field to given value.

### HasUpToDate

`func (o *PatchSummary) HasUpToDate() bool`

HasUpToDate returns a boolean if a field has been set.

### GetOutOfDate

`func (o *PatchSummary) GetOutOfDate() int64`

GetOutOfDate returns the OutOfDate field if non-nil, zero value otherwise.

### GetOutOfDateOk

`func (o *PatchSummary) GetOutOfDateOk() (*int64, bool)`

GetOutOfDateOk returns a tuple with the OutOfDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfDate

`func (o *PatchSummary) SetOutOfDate(v int64)`

SetOutOfDate sets OutOfDate field to given value.

### HasOutOfDate

`func (o *PatchSummary) HasOutOfDate() bool`

HasOutOfDate returns a boolean if a field has been set.

### GetOnDashboard

`func (o *PatchSummary) GetOnDashboard() bool`

GetOnDashboard returns the OnDashboard field if non-nil, zero value otherwise.

### GetOnDashboardOk

`func (o *PatchSummary) GetOnDashboardOk() (*bool, bool)`

GetOnDashboardOk returns a tuple with the OnDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDashboard

`func (o *PatchSummary) SetOnDashboard(v bool)`

SetOnDashboard sets OnDashboard field to given value.

### HasOnDashboard

`func (o *PatchSummary) HasOnDashboard() bool`

HasOnDashboard returns a boolean if a field has been set.

### GetSoftwareTitleConfigurationId

`func (o *PatchSummary) GetSoftwareTitleConfigurationId() string`

GetSoftwareTitleConfigurationId returns the SoftwareTitleConfigurationId field if non-nil, zero value otherwise.

### GetSoftwareTitleConfigurationIdOk

`func (o *PatchSummary) GetSoftwareTitleConfigurationIdOk() (*string, bool)`

GetSoftwareTitleConfigurationIdOk returns a tuple with the SoftwareTitleConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareTitleConfigurationId

`func (o *PatchSummary) SetSoftwareTitleConfigurationId(v string)`

SetSoftwareTitleConfigurationId sets SoftwareTitleConfigurationId field to given value.

### HasSoftwareTitleConfigurationId

`func (o *PatchSummary) HasSoftwareTitleConfigurationId() bool`

HasSoftwareTitleConfigurationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


