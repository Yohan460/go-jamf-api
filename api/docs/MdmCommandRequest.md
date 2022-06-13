# MdmCommandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientData** | Pointer to [**[]MdmCommandClient**](MdmCommandClient.md) |  | [optional] 
**CommandData** | Pointer to [**MdmCommandRequestCommandData**](MdmCommandRequestCommandData.md) |  | [optional] 

## Methods

### NewMdmCommandRequest

`func NewMdmCommandRequest() *MdmCommandRequest`

NewMdmCommandRequest instantiates a new MdmCommandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdmCommandRequestWithDefaults

`func NewMdmCommandRequestWithDefaults() *MdmCommandRequest`

NewMdmCommandRequestWithDefaults instantiates a new MdmCommandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientData

`func (o *MdmCommandRequest) GetClientData() []MdmCommandClient`

GetClientData returns the ClientData field if non-nil, zero value otherwise.

### GetClientDataOk

`func (o *MdmCommandRequest) GetClientDataOk() (*[]MdmCommandClient, bool)`

GetClientDataOk returns a tuple with the ClientData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientData

`func (o *MdmCommandRequest) SetClientData(v []MdmCommandClient)`

SetClientData sets ClientData field to given value.

### HasClientData

`func (o *MdmCommandRequest) HasClientData() bool`

HasClientData returns a boolean if a field has been set.

### GetCommandData

`func (o *MdmCommandRequest) GetCommandData() MdmCommandRequestCommandData`

GetCommandData returns the CommandData field if non-nil, zero value otherwise.

### GetCommandDataOk

`func (o *MdmCommandRequest) GetCommandDataOk() (*MdmCommandRequestCommandData, bool)`

GetCommandDataOk returns a tuple with the CommandData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandData

`func (o *MdmCommandRequest) SetCommandData(v MdmCommandRequestCommandData)`

SetCommandData sets CommandData field to given value.

### HasCommandData

`func (o *MdmCommandRequest) HasCommandData() bool`

HasCommandData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


