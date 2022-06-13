# AzureConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudIdPCommon** | [**CloudIdPCommonRequest**](CloudIdPCommonRequest.md) |  | 
**Server** | [**AzureServerConfigurationRequest**](AzureServerConfigurationRequest.md) |  | 

## Methods

### NewAzureConfigurationRequest

`func NewAzureConfigurationRequest(cloudIdPCommon CloudIdPCommonRequest, server AzureServerConfigurationRequest, ) *AzureConfigurationRequest`

NewAzureConfigurationRequest instantiates a new AzureConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureConfigurationRequestWithDefaults

`func NewAzureConfigurationRequestWithDefaults() *AzureConfigurationRequest`

NewAzureConfigurationRequestWithDefaults instantiates a new AzureConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudIdPCommon

`func (o *AzureConfigurationRequest) GetCloudIdPCommon() CloudIdPCommonRequest`

GetCloudIdPCommon returns the CloudIdPCommon field if non-nil, zero value otherwise.

### GetCloudIdPCommonOk

`func (o *AzureConfigurationRequest) GetCloudIdPCommonOk() (*CloudIdPCommonRequest, bool)`

GetCloudIdPCommonOk returns a tuple with the CloudIdPCommon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudIdPCommon

`func (o *AzureConfigurationRequest) SetCloudIdPCommon(v CloudIdPCommonRequest)`

SetCloudIdPCommon sets CloudIdPCommon field to given value.


### GetServer

`func (o *AzureConfigurationRequest) GetServer() AzureServerConfigurationRequest`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *AzureConfigurationRequest) GetServerOk() (*AzureServerConfigurationRequest, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *AzureConfigurationRequest) SetServer(v AzureServerConfigurationRequest)`

SetServer sets Server field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


