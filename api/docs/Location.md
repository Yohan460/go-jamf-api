# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**RealName** | Pointer to **string** |  | [optional] 
**EmailAddress** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**Department** | Pointer to [**IdAndName**](IdAndName.md) |  | [optional] 
**Building** | Pointer to [**IdAndName**](IdAndName.md) |  | [optional] 
**Room** | Pointer to **string** |  | [optional] 

## Methods

### NewLocation

`func NewLocation() *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *Location) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Location) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Location) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Location) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetRealName

`func (o *Location) GetRealName() string`

GetRealName returns the RealName field if non-nil, zero value otherwise.

### GetRealNameOk

`func (o *Location) GetRealNameOk() (*string, bool)`

GetRealNameOk returns a tuple with the RealName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealName

`func (o *Location) SetRealName(v string)`

SetRealName sets RealName field to given value.

### HasRealName

`func (o *Location) HasRealName() bool`

HasRealName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *Location) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *Location) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *Location) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *Location) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetPosition

`func (o *Location) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Location) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Location) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Location) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Location) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Location) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Location) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Location) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetDepartment

`func (o *Location) GetDepartment() IdAndName`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Location) GetDepartmentOk() (*IdAndName, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Location) SetDepartment(v IdAndName)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *Location) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetBuilding

`func (o *Location) GetBuilding() IdAndName`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *Location) GetBuildingOk() (*IdAndName, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *Location) SetBuilding(v IdAndName)`

SetBuilding sets Building field to given value.

### HasBuilding

`func (o *Location) HasBuilding() bool`

HasBuilding returns a boolean if a field has been set.

### GetRoom

`func (o *Location) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *Location) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *Location) SetRoom(v string)`

SetRoom sets Room field to given value.

### HasRoom

`func (o *Location) HasRoom() bool`

HasRoom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


