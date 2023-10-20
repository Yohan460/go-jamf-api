# DashboardItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Subtitle** | Pointer to **NullableString** |  | [optional] 
**Info** | Pointer to **NullableString** | Additional information such as identifiers for a specific policy within a software patch | [optional] 
**Enabled** | Pointer to **bool** | Logical to decide whether widget should be enabled or disabled; i.e. Policy | [optional] [default to true]
**Metrics** | Pointer to [**[]DashboardMetric**](DashboardMetric.md) |  | [optional] 
**Details** | Pointer to [**[]DashboardDetail**](DashboardDetail.md) |  | [optional] 
**Error** | Pointer to [**DashboardApiError**](DashboardApiError.md) |  | [optional] 

## Methods

### NewDashboardItem

`func NewDashboardItem() *DashboardItem`

NewDashboardItem instantiates a new DashboardItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardItemWithDefaults

`func NewDashboardItemWithDefaults() *DashboardItem`

NewDashboardItemWithDefaults instantiates a new DashboardItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DashboardItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DashboardItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *DashboardItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DashboardItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DashboardItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DashboardItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *DashboardItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *DashboardItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetSubtitle

`func (o *DashboardItem) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *DashboardItem) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *DashboardItem) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *DashboardItem) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### SetSubtitleNil

`func (o *DashboardItem) SetSubtitleNil(b bool)`

 SetSubtitleNil sets the value for Subtitle to be an explicit nil

### UnsetSubtitle
`func (o *DashboardItem) UnsetSubtitle()`

UnsetSubtitle ensures that no value is present for Subtitle, not even an explicit nil
### GetInfo

`func (o *DashboardItem) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *DashboardItem) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *DashboardItem) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *DashboardItem) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### SetInfoNil

`func (o *DashboardItem) SetInfoNil(b bool)`

 SetInfoNil sets the value for Info to be an explicit nil

### UnsetInfo
`func (o *DashboardItem) UnsetInfo()`

UnsetInfo ensures that no value is present for Info, not even an explicit nil
### GetEnabled

`func (o *DashboardItem) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DashboardItem) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DashboardItem) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DashboardItem) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMetrics

`func (o *DashboardItem) GetMetrics() []DashboardMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *DashboardItem) GetMetricsOk() (*[]DashboardMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *DashboardItem) SetMetrics(v []DashboardMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *DashboardItem) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetDetails

`func (o *DashboardItem) GetDetails() []DashboardDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DashboardItem) GetDetailsOk() (*[]DashboardDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DashboardItem) SetDetails(v []DashboardDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *DashboardItem) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetError

`func (o *DashboardItem) GetError() DashboardApiError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DashboardItem) GetErrorOk() (*DashboardApiError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DashboardItem) SetError(v DashboardApiError)`

SetError sets Error field to given value.

### HasError

`func (o *DashboardItem) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


