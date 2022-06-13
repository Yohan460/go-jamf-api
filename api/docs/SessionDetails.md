# SessionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Session identifier | [optional] 
**Code** | Pointer to **string** | Sessions code | [optional] 
**Description** | Pointer to **string** | Session description. To be used for additional context on the reason of the session | [optional] 
**SupporterLink** | Pointer to **string** | Supporter session URL | [optional] 
**EndUserLink** | Pointer to **string** | End user session URL | [optional] 
**DeviceId** | Pointer to **string** | Device identifier | [optional] 
**DeviceName** | Pointer to **string** | Device name if found - null otherwise | [optional] 
**DeviceType** | Pointer to **string** | Device type | [optional] 
**State** | Pointer to **string** | Session state | [optional] 
**CreatorId** | Pointer to **string** | ID of session creator if session created by Jamf Pro local user, null otherwise | [optional] 
**CreatorName** | Pointer to **string** | Username of the session creator | [optional] 
**CreatedAt** | Pointer to **time.Time** | Session creation time | [optional] 

## Methods

### NewSessionDetails

`func NewSessionDetails() *SessionDetails`

NewSessionDetails instantiates a new SessionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionDetailsWithDefaults

`func NewSessionDetailsWithDefaults() *SessionDetails`

NewSessionDetailsWithDefaults instantiates a new SessionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SessionDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SessionDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *SessionDetails) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SessionDetails) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SessionDetails) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SessionDetails) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *SessionDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SessionDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SessionDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SessionDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSupporterLink

`func (o *SessionDetails) GetSupporterLink() string`

GetSupporterLink returns the SupporterLink field if non-nil, zero value otherwise.

### GetSupporterLinkOk

`func (o *SessionDetails) GetSupporterLinkOk() (*string, bool)`

GetSupporterLinkOk returns a tuple with the SupporterLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupporterLink

`func (o *SessionDetails) SetSupporterLink(v string)`

SetSupporterLink sets SupporterLink field to given value.

### HasSupporterLink

`func (o *SessionDetails) HasSupporterLink() bool`

HasSupporterLink returns a boolean if a field has been set.

### GetEndUserLink

`func (o *SessionDetails) GetEndUserLink() string`

GetEndUserLink returns the EndUserLink field if non-nil, zero value otherwise.

### GetEndUserLinkOk

`func (o *SessionDetails) GetEndUserLinkOk() (*string, bool)`

GetEndUserLinkOk returns a tuple with the EndUserLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserLink

`func (o *SessionDetails) SetEndUserLink(v string)`

SetEndUserLink sets EndUserLink field to given value.

### HasEndUserLink

`func (o *SessionDetails) HasEndUserLink() bool`

HasEndUserLink returns a boolean if a field has been set.

### GetDeviceId

`func (o *SessionDetails) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SessionDetails) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SessionDetails) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *SessionDetails) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *SessionDetails) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *SessionDetails) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *SessionDetails) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *SessionDetails) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceType

`func (o *SessionDetails) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *SessionDetails) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *SessionDetails) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *SessionDetails) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetState

`func (o *SessionDetails) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SessionDetails) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SessionDetails) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SessionDetails) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCreatorId

`func (o *SessionDetails) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *SessionDetails) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *SessionDetails) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *SessionDetails) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetCreatorName

`func (o *SessionDetails) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *SessionDetails) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *SessionDetails) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *SessionDetails) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SessionDetails) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SessionDetails) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SessionDetails) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SessionDetails) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


