# FileData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **string** | The name of the file | [optional] [readonly] 
**Length** | Pointer to **int64** | The length in bytes | [optional] [readonly] 
**Md5** | Pointer to **string** | The MD5 in hex | [optional] [readonly] 
**Region** | Pointer to **string** | The region the file is hosted in | [optional] [readonly] 
**Sha3** | Pointer to **string** | The SHA3_512 in hex | [optional] [readonly] 

## Methods

### NewFileData

`func NewFileData() *FileData`

NewFileData instantiates a new FileData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileDataWithDefaults

`func NewFileDataWithDefaults() *FileData`

NewFileDataWithDefaults instantiates a new FileData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *FileData) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *FileData) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *FileData) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *FileData) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetLength

`func (o *FileData) GetLength() int64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *FileData) GetLengthOk() (*int64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *FileData) SetLength(v int64)`

SetLength sets Length field to given value.

### HasLength

`func (o *FileData) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetMd5

`func (o *FileData) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *FileData) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *FileData) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *FileData) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetRegion

`func (o *FileData) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FileData) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FileData) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *FileData) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSha3

`func (o *FileData) GetSha3() string`

GetSha3 returns the Sha3 field if non-nil, zero value otherwise.

### GetSha3Ok

`func (o *FileData) GetSha3Ok() (*string, bool)`

GetSha3Ok returns a tuple with the Sha3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha3

`func (o *FileData) SetSha3(v string)`

SetSha3 sets Sha3 field to given value.

### HasSha3

`func (o *FileData) HasSha3() bool`

HasSha3 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


