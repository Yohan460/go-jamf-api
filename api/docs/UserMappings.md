# UserMappings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectClassLimitation** | **string** |  | 
**ObjectClasses** | **string** |  | 
**SearchBase** | **string** |  | 
**SearchScope** | **string** |  | 
**AdditionalSearchBase** | Pointer to **string** |  | [optional] [default to ""]
**UserID** | **string** |  | 
**Username** | **string** |  | 
**RealName** | **string** |  | 
**EmailAddress** | **string** |  | 
**Department** | **string** |  | 
**Building** | **string** |  | [default to ""]
**Room** | **string** |  | [default to ""]
**Phone** | **string** |  | [default to ""]
**Position** | **string** |  | 
**UserUuid** | **string** |  | 

## Methods

### NewUserMappings

`func NewUserMappings(objectClassLimitation string, objectClasses string, searchBase string, searchScope string, userID string, username string, realName string, emailAddress string, department string, building string, room string, phone string, position string, userUuid string, ) *UserMappings`

NewUserMappings instantiates a new UserMappings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserMappingsWithDefaults

`func NewUserMappingsWithDefaults() *UserMappings`

NewUserMappingsWithDefaults instantiates a new UserMappings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectClassLimitation

`func (o *UserMappings) GetObjectClassLimitation() string`

GetObjectClassLimitation returns the ObjectClassLimitation field if non-nil, zero value otherwise.

### GetObjectClassLimitationOk

`func (o *UserMappings) GetObjectClassLimitationOk() (*string, bool)`

GetObjectClassLimitationOk returns a tuple with the ObjectClassLimitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClassLimitation

`func (o *UserMappings) SetObjectClassLimitation(v string)`

SetObjectClassLimitation sets ObjectClassLimitation field to given value.


### GetObjectClasses

`func (o *UserMappings) GetObjectClasses() string`

GetObjectClasses returns the ObjectClasses field if non-nil, zero value otherwise.

### GetObjectClassesOk

`func (o *UserMappings) GetObjectClassesOk() (*string, bool)`

GetObjectClassesOk returns a tuple with the ObjectClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClasses

`func (o *UserMappings) SetObjectClasses(v string)`

SetObjectClasses sets ObjectClasses field to given value.


### GetSearchBase

`func (o *UserMappings) GetSearchBase() string`

GetSearchBase returns the SearchBase field if non-nil, zero value otherwise.

### GetSearchBaseOk

`func (o *UserMappings) GetSearchBaseOk() (*string, bool)`

GetSearchBaseOk returns a tuple with the SearchBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBase

`func (o *UserMappings) SetSearchBase(v string)`

SetSearchBase sets SearchBase field to given value.


### GetSearchScope

`func (o *UserMappings) GetSearchScope() string`

GetSearchScope returns the SearchScope field if non-nil, zero value otherwise.

### GetSearchScopeOk

`func (o *UserMappings) GetSearchScopeOk() (*string, bool)`

GetSearchScopeOk returns a tuple with the SearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchScope

`func (o *UserMappings) SetSearchScope(v string)`

SetSearchScope sets SearchScope field to given value.


### GetAdditionalSearchBase

`func (o *UserMappings) GetAdditionalSearchBase() string`

GetAdditionalSearchBase returns the AdditionalSearchBase field if non-nil, zero value otherwise.

### GetAdditionalSearchBaseOk

`func (o *UserMappings) GetAdditionalSearchBaseOk() (*string, bool)`

GetAdditionalSearchBaseOk returns a tuple with the AdditionalSearchBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSearchBase

`func (o *UserMappings) SetAdditionalSearchBase(v string)`

SetAdditionalSearchBase sets AdditionalSearchBase field to given value.

### HasAdditionalSearchBase

`func (o *UserMappings) HasAdditionalSearchBase() bool`

HasAdditionalSearchBase returns a boolean if a field has been set.

### GetUserID

`func (o *UserMappings) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *UserMappings) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *UserMappings) SetUserID(v string)`

SetUserID sets UserID field to given value.


### GetUsername

`func (o *UserMappings) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserMappings) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserMappings) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetRealName

`func (o *UserMappings) GetRealName() string`

GetRealName returns the RealName field if non-nil, zero value otherwise.

### GetRealNameOk

`func (o *UserMappings) GetRealNameOk() (*string, bool)`

GetRealNameOk returns a tuple with the RealName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealName

`func (o *UserMappings) SetRealName(v string)`

SetRealName sets RealName field to given value.


### GetEmailAddress

`func (o *UserMappings) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *UserMappings) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *UserMappings) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetDepartment

`func (o *UserMappings) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *UserMappings) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *UserMappings) SetDepartment(v string)`

SetDepartment sets Department field to given value.


### GetBuilding

`func (o *UserMappings) GetBuilding() string`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *UserMappings) GetBuildingOk() (*string, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *UserMappings) SetBuilding(v string)`

SetBuilding sets Building field to given value.


### GetRoom

`func (o *UserMappings) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *UserMappings) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *UserMappings) SetRoom(v string)`

SetRoom sets Room field to given value.


### GetPhone

`func (o *UserMappings) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UserMappings) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UserMappings) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetPosition

`func (o *UserMappings) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *UserMappings) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *UserMappings) SetPosition(v string)`

SetPosition sets Position field to given value.


### GetUserUuid

`func (o *UserMappings) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *UserMappings) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *UserMappings) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


