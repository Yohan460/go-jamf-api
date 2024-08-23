# FileTransferItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePath** | Pointer to **string** |  | [optional] 
**TransferTimestamp** | Pointer to **time.Time** |  | [optional] 
**FileTransferType** | Pointer to **string** |  | [optional] 

## Methods

### NewFileTransferItem

`func NewFileTransferItem() *FileTransferItem`

NewFileTransferItem instantiates a new FileTransferItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileTransferItemWithDefaults

`func NewFileTransferItemWithDefaults() *FileTransferItem`

NewFileTransferItemWithDefaults instantiates a new FileTransferItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePath

`func (o *FileTransferItem) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *FileTransferItem) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *FileTransferItem) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *FileTransferItem) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetTransferTimestamp

`func (o *FileTransferItem) GetTransferTimestamp() time.Time`

GetTransferTimestamp returns the TransferTimestamp field if non-nil, zero value otherwise.

### GetTransferTimestampOk

`func (o *FileTransferItem) GetTransferTimestampOk() (*time.Time, bool)`

GetTransferTimestampOk returns a tuple with the TransferTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferTimestamp

`func (o *FileTransferItem) SetTransferTimestamp(v time.Time)`

SetTransferTimestamp sets TransferTimestamp field to given value.

### HasTransferTimestamp

`func (o *FileTransferItem) HasTransferTimestamp() bool`

HasTransferTimestamp returns a boolean if a field has been set.

### GetFileTransferType

`func (o *FileTransferItem) GetFileTransferType() string`

GetFileTransferType returns the FileTransferType field if non-nil, zero value otherwise.

### GetFileTransferTypeOk

`func (o *FileTransferItem) GetFileTransferTypeOk() (*string, bool)`

GetFileTransferTypeOk returns a tuple with the FileTransferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTransferType

`func (o *FileTransferItem) SetFileTransferType(v string)`

SetFileTransferType sets FileTransferType field to given value.

### HasFileTransferType

`func (o *FileTransferItem) HasFileTransferType() bool`

HasFileTransferType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


