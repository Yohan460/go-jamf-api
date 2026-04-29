# UserInventory

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

## Methods

### NewUserInventory

`func NewUserInventory() *UserInventory`

NewUserInventory instantiates a new UserInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInventoryWithDefaults

`func NewUserInventoryWithDefaults() *UserInventory`

NewUserInventoryWithDefaults instantiates a new UserInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserInventory) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserInventory) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserInventory) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserInventory) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetRealname

`func (o *UserInventory) GetRealname() string`

GetRealname returns the Realname field if non-nil, zero value otherwise.

### GetRealnameOk

`func (o *UserInventory) GetRealnameOk() (*string, bool)`

GetRealnameOk returns a tuple with the Realname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealname

`func (o *UserInventory) SetRealname(v string)`

SetRealname sets Realname field to given value.

### HasRealname

`func (o *UserInventory) HasRealname() bool`

HasRealname returns a boolean if a field has been set.

### GetEmail

`func (o *UserInventory) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserInventory) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserInventory) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserInventory) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *UserInventory) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UserInventory) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UserInventory) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UserInventory) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *UserInventory) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *UserInventory) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetPosition

`func (o *UserInventory) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *UserInventory) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *UserInventory) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *UserInventory) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *UserInventory) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *UserInventory) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetEnableCustomPhotoUrl

`func (o *UserInventory) GetEnableCustomPhotoUrl() bool`

GetEnableCustomPhotoUrl returns the EnableCustomPhotoUrl field if non-nil, zero value otherwise.

### GetEnableCustomPhotoUrlOk

`func (o *UserInventory) GetEnableCustomPhotoUrlOk() (*bool, bool)`

GetEnableCustomPhotoUrlOk returns a tuple with the EnableCustomPhotoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCustomPhotoUrl

`func (o *UserInventory) SetEnableCustomPhotoUrl(v bool)`

SetEnableCustomPhotoUrl sets EnableCustomPhotoUrl field to given value.

### HasEnableCustomPhotoUrl

`func (o *UserInventory) HasEnableCustomPhotoUrl() bool`

HasEnableCustomPhotoUrl returns a boolean if a field has been set.

### GetCustomPhotoUrl

`func (o *UserInventory) GetCustomPhotoUrl() string`

GetCustomPhotoUrl returns the CustomPhotoUrl field if non-nil, zero value otherwise.

### GetCustomPhotoUrlOk

`func (o *UserInventory) GetCustomPhotoUrlOk() (*string, bool)`

GetCustomPhotoUrlOk returns a tuple with the CustomPhotoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPhotoUrl

`func (o *UserInventory) SetCustomPhotoUrl(v string)`

SetCustomPhotoUrl sets CustomPhotoUrl field to given value.

### HasCustomPhotoUrl

`func (o *UserInventory) HasCustomPhotoUrl() bool`

HasCustomPhotoUrl returns a boolean if a field has been set.

### GetManagedAppleId

`func (o *UserInventory) GetManagedAppleId() string`

GetManagedAppleId returns the ManagedAppleId field if non-nil, zero value otherwise.

### GetManagedAppleIdOk

`func (o *UserInventory) GetManagedAppleIdOk() (*string, bool)`

GetManagedAppleIdOk returns a tuple with the ManagedAppleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedAppleId

`func (o *UserInventory) SetManagedAppleId(v string)`

SetManagedAppleId sets ManagedAppleId field to given value.

### HasManagedAppleId

`func (o *UserInventory) HasManagedAppleId() bool`

HasManagedAppleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


