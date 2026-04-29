# MdmRenewalError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MdmRenewalErrorId** | **string** | Unique identifier for the MDM renewal error | 
**ClientManagementId** | **string** | The client management ID associated with this error | 
**MdmRenewalErrorType** | [**MdmRenewalErrorType**](MdmRenewalErrorType.md) |  | 
**ErrorTimeStamp** | Pointer to **time.Time** | Timestamp when the error occurred (ISO 8601 format) | [optional] 
**FailureCount** | Pointer to **int64** | Number of times this error has occurred | [optional] 

## Methods

### NewMdmRenewalError

`func NewMdmRenewalError(mdmRenewalErrorId string, clientManagementId string, mdmRenewalErrorType MdmRenewalErrorType, ) *MdmRenewalError`

NewMdmRenewalError instantiates a new MdmRenewalError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdmRenewalErrorWithDefaults

`func NewMdmRenewalErrorWithDefaults() *MdmRenewalError`

NewMdmRenewalErrorWithDefaults instantiates a new MdmRenewalError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMdmRenewalErrorId

`func (o *MdmRenewalError) GetMdmRenewalErrorId() string`

GetMdmRenewalErrorId returns the MdmRenewalErrorId field if non-nil, zero value otherwise.

### GetMdmRenewalErrorIdOk

`func (o *MdmRenewalError) GetMdmRenewalErrorIdOk() (*string, bool)`

GetMdmRenewalErrorIdOk returns a tuple with the MdmRenewalErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmRenewalErrorId

`func (o *MdmRenewalError) SetMdmRenewalErrorId(v string)`

SetMdmRenewalErrorId sets MdmRenewalErrorId field to given value.


### GetClientManagementId

`func (o *MdmRenewalError) GetClientManagementId() string`

GetClientManagementId returns the ClientManagementId field if non-nil, zero value otherwise.

### GetClientManagementIdOk

`func (o *MdmRenewalError) GetClientManagementIdOk() (*string, bool)`

GetClientManagementIdOk returns a tuple with the ClientManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientManagementId

`func (o *MdmRenewalError) SetClientManagementId(v string)`

SetClientManagementId sets ClientManagementId field to given value.


### GetMdmRenewalErrorType

`func (o *MdmRenewalError) GetMdmRenewalErrorType() MdmRenewalErrorType`

GetMdmRenewalErrorType returns the MdmRenewalErrorType field if non-nil, zero value otherwise.

### GetMdmRenewalErrorTypeOk

`func (o *MdmRenewalError) GetMdmRenewalErrorTypeOk() (*MdmRenewalErrorType, bool)`

GetMdmRenewalErrorTypeOk returns a tuple with the MdmRenewalErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmRenewalErrorType

`func (o *MdmRenewalError) SetMdmRenewalErrorType(v MdmRenewalErrorType)`

SetMdmRenewalErrorType sets MdmRenewalErrorType field to given value.


### GetErrorTimeStamp

`func (o *MdmRenewalError) GetErrorTimeStamp() time.Time`

GetErrorTimeStamp returns the ErrorTimeStamp field if non-nil, zero value otherwise.

### GetErrorTimeStampOk

`func (o *MdmRenewalError) GetErrorTimeStampOk() (*time.Time, bool)`

GetErrorTimeStampOk returns a tuple with the ErrorTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorTimeStamp

`func (o *MdmRenewalError) SetErrorTimeStamp(v time.Time)`

SetErrorTimeStamp sets ErrorTimeStamp field to given value.

### HasErrorTimeStamp

`func (o *MdmRenewalError) HasErrorTimeStamp() bool`

HasErrorTimeStamp returns a boolean if a field has been set.

### GetFailureCount

`func (o *MdmRenewalError) GetFailureCount() int64`

GetFailureCount returns the FailureCount field if non-nil, zero value otherwise.

### GetFailureCountOk

`func (o *MdmRenewalError) GetFailureCountOk() (*int64, bool)`

GetFailureCountOk returns a tuple with the FailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCount

`func (o *MdmRenewalError) SetFailureCount(v int64)`

SetFailureCount sets FailureCount field to given value.

### HasFailureCount

`func (o *MdmRenewalError) HasFailureCount() bool`

HasFailureCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


