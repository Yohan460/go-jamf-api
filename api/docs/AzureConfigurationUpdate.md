# AzureConfigurationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudIdPCommon** | [**CloudIdPCommon**](CloudIdPCommon.md) |  | 
**Server** | [**AzureServerConfigurationUpdate**](AzureServerConfigurationUpdate.md) |  | 

## Methods

### NewAzureConfigurationUpdate

`func NewAzureConfigurationUpdate(cloudIdPCommon CloudIdPCommon, server AzureServerConfigurationUpdate, ) *AzureConfigurationUpdate`

NewAzureConfigurationUpdate instantiates a new AzureConfigurationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureConfigurationUpdateWithDefaults

`func NewAzureConfigurationUpdateWithDefaults() *AzureConfigurationUpdate`

NewAzureConfigurationUpdateWithDefaults instantiates a new AzureConfigurationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudIdPCommon

`func (o *AzureConfigurationUpdate) GetCloudIdPCommon() CloudIdPCommon`

GetCloudIdPCommon returns the CloudIdPCommon field if non-nil, zero value otherwise.

### GetCloudIdPCommonOk

`func (o *AzureConfigurationUpdate) GetCloudIdPCommonOk() (*CloudIdPCommon, bool)`

GetCloudIdPCommonOk returns a tuple with the CloudIdPCommon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudIdPCommon

`func (o *AzureConfigurationUpdate) SetCloudIdPCommon(v CloudIdPCommon)`

SetCloudIdPCommon sets CloudIdPCommon field to given value.


### GetServer

`func (o *AzureConfigurationUpdate) GetServer() AzureServerConfigurationUpdate`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *AzureConfigurationUpdate) GetServerOk() (*AzureServerConfigurationUpdate, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *AzureConfigurationUpdate) SetServer(v AzureServerConfigurationUpdate)`

SetServer sets Server field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


