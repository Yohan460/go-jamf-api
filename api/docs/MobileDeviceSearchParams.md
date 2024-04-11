# MobileDeviceSearchParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | Pointer to **int64** |  | [optional] 
**PageSize** | Pointer to **int64** |  | [optional] 
**IsLoadToEnd** | Pointer to **bool** |  | [optional] 
**OrderBy** | Pointer to [**[]OrderBy**](OrderBy.md) |  | [optional] 
**Udid** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**OsType** | Pointer to **string** |  | [optional] 
**IsManaged** | Pointer to **bool** |  | [optional] 
**ExcludedIds** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewMobileDeviceSearchParams

`func NewMobileDeviceSearchParams() *MobileDeviceSearchParams`

NewMobileDeviceSearchParams instantiates a new MobileDeviceSearchParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileDeviceSearchParamsWithDefaults

`func NewMobileDeviceSearchParamsWithDefaults() *MobileDeviceSearchParams`

NewMobileDeviceSearchParamsWithDefaults instantiates a new MobileDeviceSearchParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *MobileDeviceSearchParams) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *MobileDeviceSearchParams) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *MobileDeviceSearchParams) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *MobileDeviceSearchParams) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *MobileDeviceSearchParams) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *MobileDeviceSearchParams) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *MobileDeviceSearchParams) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *MobileDeviceSearchParams) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetIsLoadToEnd

`func (o *MobileDeviceSearchParams) GetIsLoadToEnd() bool`

GetIsLoadToEnd returns the IsLoadToEnd field if non-nil, zero value otherwise.

### GetIsLoadToEndOk

`func (o *MobileDeviceSearchParams) GetIsLoadToEndOk() (*bool, bool)`

GetIsLoadToEndOk returns a tuple with the IsLoadToEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLoadToEnd

`func (o *MobileDeviceSearchParams) SetIsLoadToEnd(v bool)`

SetIsLoadToEnd sets IsLoadToEnd field to given value.

### HasIsLoadToEnd

`func (o *MobileDeviceSearchParams) HasIsLoadToEnd() bool`

HasIsLoadToEnd returns a boolean if a field has been set.

### GetOrderBy

`func (o *MobileDeviceSearchParams) GetOrderBy() []OrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *MobileDeviceSearchParams) GetOrderByOk() (*[]OrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *MobileDeviceSearchParams) SetOrderBy(v []OrderBy)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *MobileDeviceSearchParams) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetUdid

`func (o *MobileDeviceSearchParams) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *MobileDeviceSearchParams) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *MobileDeviceSearchParams) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *MobileDeviceSearchParams) HasUdid() bool`

HasUdid returns a boolean if a field has been set.

### GetMacAddress

`func (o *MobileDeviceSearchParams) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *MobileDeviceSearchParams) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *MobileDeviceSearchParams) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *MobileDeviceSearchParams) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetName

`func (o *MobileDeviceSearchParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MobileDeviceSearchParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MobileDeviceSearchParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MobileDeviceSearchParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerialNumber

`func (o *MobileDeviceSearchParams) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *MobileDeviceSearchParams) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *MobileDeviceSearchParams) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *MobileDeviceSearchParams) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetOsType

`func (o *MobileDeviceSearchParams) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *MobileDeviceSearchParams) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *MobileDeviceSearchParams) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *MobileDeviceSearchParams) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetIsManaged

`func (o *MobileDeviceSearchParams) GetIsManaged() bool`

GetIsManaged returns the IsManaged field if non-nil, zero value otherwise.

### GetIsManagedOk

`func (o *MobileDeviceSearchParams) GetIsManagedOk() (*bool, bool)`

GetIsManagedOk returns a tuple with the IsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManaged

`func (o *MobileDeviceSearchParams) SetIsManaged(v bool)`

SetIsManaged sets IsManaged field to given value.

### HasIsManaged

`func (o *MobileDeviceSearchParams) HasIsManaged() bool`

HasIsManaged returns a boolean if a field has been set.

### GetExcludedIds

`func (o *MobileDeviceSearchParams) GetExcludedIds() []int64`

GetExcludedIds returns the ExcludedIds field if non-nil, zero value otherwise.

### GetExcludedIdsOk

`func (o *MobileDeviceSearchParams) GetExcludedIdsOk() (*[]int64, bool)`

GetExcludedIdsOk returns a tuple with the ExcludedIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedIds

`func (o *MobileDeviceSearchParams) SetExcludedIds(v []int64)`

SetExcludedIds sets ExcludedIds field to given value.

### HasExcludedIds

`func (o *MobileDeviceSearchParams) HasExcludedIds() bool`

HasExcludedIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


