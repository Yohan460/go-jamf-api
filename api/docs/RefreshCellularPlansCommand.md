# RefreshCellularPlansCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandType** | [**MdmCommandType**](MdmCommandType.md) |  | 
**EsimServerUrl** | **string** | The URL of the eSIM server that the device should contact to refresh cellular plans. | 

## Methods

### NewRefreshCellularPlansCommand

`func NewRefreshCellularPlansCommand(commandType MdmCommandType, esimServerUrl string, ) *RefreshCellularPlansCommand`

NewRefreshCellularPlansCommand instantiates a new RefreshCellularPlansCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshCellularPlansCommandWithDefaults

`func NewRefreshCellularPlansCommandWithDefaults() *RefreshCellularPlansCommand`

NewRefreshCellularPlansCommandWithDefaults instantiates a new RefreshCellularPlansCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandType

`func (o *RefreshCellularPlansCommand) GetCommandType() MdmCommandType`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *RefreshCellularPlansCommand) GetCommandTypeOk() (*MdmCommandType, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *RefreshCellularPlansCommand) SetCommandType(v MdmCommandType)`

SetCommandType sets CommandType field to given value.


### GetEsimServerUrl

`func (o *RefreshCellularPlansCommand) GetEsimServerUrl() string`

GetEsimServerUrl returns the EsimServerUrl field if non-nil, zero value otherwise.

### GetEsimServerUrlOk

`func (o *RefreshCellularPlansCommand) GetEsimServerUrlOk() (*string, bool)`

GetEsimServerUrlOk returns a tuple with the EsimServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsimServerUrl

`func (o *RefreshCellularPlansCommand) SetEsimServerUrl(v string)`

SetEsimServerUrl sets EsimServerUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


