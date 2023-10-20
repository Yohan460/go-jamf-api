# DashboardApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpStatusCode** | Pointer to **int32** |  | [optional] [default to 500]
**Id** | Pointer to **string** |  | [optional] [default to ""]
**Description** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewDashboardApiError

`func NewDashboardApiError() *DashboardApiError`

NewDashboardApiError instantiates a new DashboardApiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardApiErrorWithDefaults

`func NewDashboardApiErrorWithDefaults() *DashboardApiError`

NewDashboardApiErrorWithDefaults instantiates a new DashboardApiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpStatusCode

`func (o *DashboardApiError) GetHttpStatusCode() int32`

GetHttpStatusCode returns the HttpStatusCode field if non-nil, zero value otherwise.

### GetHttpStatusCodeOk

`func (o *DashboardApiError) GetHttpStatusCodeOk() (*int32, bool)`

GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatusCode

`func (o *DashboardApiError) SetHttpStatusCode(v int32)`

SetHttpStatusCode sets HttpStatusCode field to given value.

### HasHttpStatusCode

`func (o *DashboardApiError) HasHttpStatusCode() bool`

HasHttpStatusCode returns a boolean if a field has been set.

### GetId

`func (o *DashboardApiError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardApiError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardApiError) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DashboardApiError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *DashboardApiError) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardApiError) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DashboardApiError) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DashboardApiError) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


