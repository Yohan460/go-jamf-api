# MdmCommandClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagementId** | Pointer to **string** |  | [optional] 
**ClientType** | Pointer to [**MdmClientType**](MdmClientType.md) |  | [optional] 

## Methods

### NewMdmCommandClient

`func NewMdmCommandClient() *MdmCommandClient`

NewMdmCommandClient instantiates a new MdmCommandClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdmCommandClientWithDefaults

`func NewMdmCommandClientWithDefaults() *MdmCommandClient`

NewMdmCommandClientWithDefaults instantiates a new MdmCommandClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagementId

`func (o *MdmCommandClient) GetManagementId() string`

GetManagementId returns the ManagementId field if non-nil, zero value otherwise.

### GetManagementIdOk

`func (o *MdmCommandClient) GetManagementIdOk() (*string, bool)`

GetManagementIdOk returns a tuple with the ManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementId

`func (o *MdmCommandClient) SetManagementId(v string)`

SetManagementId sets ManagementId field to given value.

### HasManagementId

`func (o *MdmCommandClient) HasManagementId() bool`

HasManagementId returns a boolean if a field has been set.

### GetClientType

`func (o *MdmCommandClient) GetClientType() MdmClientType`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *MdmCommandClient) GetClientTypeOk() (*MdmClientType, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *MdmCommandClient) SetClientType(v MdmClientType)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *MdmCommandClient) HasClientType() bool`

HasClientType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


