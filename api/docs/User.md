# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**Realname** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**Position** | Pointer to **NullableString** |  | [optional] 
**EnableCustomPhotoUrl** | Pointer to **bool** |  | [optional] 
**CustomPhotoUrl** | Pointer to **string** |  | [optional] 
**ManagedAppleId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetRealname

`func (o *User) GetRealname() string`

GetRealname returns the Realname field if non-nil, zero value otherwise.

### GetRealnameOk

`func (o *User) GetRealnameOk() (*string, bool)`

GetRealnameOk returns a tuple with the Realname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealname

`func (o *User) SetRealname(v string)`

SetRealname sets Realname field to given value.

### HasRealname

`func (o *User) HasRealname() bool`

HasRealname returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *User) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *User) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *User) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *User) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *User) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *User) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetPosition

`func (o *User) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *User) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *User) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *User) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *User) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *User) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetEnableCustomPhotoUrl

`func (o *User) GetEnableCustomPhotoUrl() bool`

GetEnableCustomPhotoUrl returns the EnableCustomPhotoUrl field if non-nil, zero value otherwise.

### GetEnableCustomPhotoUrlOk

`func (o *User) GetEnableCustomPhotoUrlOk() (*bool, bool)`

GetEnableCustomPhotoUrlOk returns a tuple with the EnableCustomPhotoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCustomPhotoUrl

`func (o *User) SetEnableCustomPhotoUrl(v bool)`

SetEnableCustomPhotoUrl sets EnableCustomPhotoUrl field to given value.

### HasEnableCustomPhotoUrl

`func (o *User) HasEnableCustomPhotoUrl() bool`

HasEnableCustomPhotoUrl returns a boolean if a field has been set.

### GetCustomPhotoUrl

`func (o *User) GetCustomPhotoUrl() string`

GetCustomPhotoUrl returns the CustomPhotoUrl field if non-nil, zero value otherwise.

### GetCustomPhotoUrlOk

`func (o *User) GetCustomPhotoUrlOk() (*string, bool)`

GetCustomPhotoUrlOk returns a tuple with the CustomPhotoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPhotoUrl

`func (o *User) SetCustomPhotoUrl(v string)`

SetCustomPhotoUrl sets CustomPhotoUrl field to given value.

### HasCustomPhotoUrl

`func (o *User) HasCustomPhotoUrl() bool`

HasCustomPhotoUrl returns a boolean if a field has been set.

### GetManagedAppleId

`func (o *User) GetManagedAppleId() string`

GetManagedAppleId returns the ManagedAppleId field if non-nil, zero value otherwise.

### GetManagedAppleIdOk

`func (o *User) GetManagedAppleIdOk() (*string, bool)`

GetManagedAppleIdOk returns a tuple with the ManagedAppleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedAppleId

`func (o *User) SetManagedAppleId(v string)`

SetManagedAppleId sets ManagedAppleId field to given value.

### HasManagedAppleId

`func (o *User) HasManagedAppleId() bool`

HasManagedAppleId returns a boolean if a field has been set.

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


