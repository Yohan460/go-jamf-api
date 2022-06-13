# AzureConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudIdPCommon** | [**CloudIdPCommon**](CloudIdPCommon.md) |  | 
**Server** | [**AzureServerConfiguration**](AzureServerConfiguration.md) |  | 

## Methods

### NewAzureConfiguration

`func NewAzureConfiguration(cloudIdPCommon CloudIdPCommon, server AzureServerConfiguration, ) *AzureConfiguration`

NewAzureConfiguration instantiates a new AzureConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureConfigurationWithDefaults

`func NewAzureConfigurationWithDefaults() *AzureConfiguration`

NewAzureConfigurationWithDefaults instantiates a new AzureConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudIdPCommon

`func (o *AzureConfiguration) GetCloudIdPCommon() CloudIdPCommon`

GetCloudIdPCommon returns the CloudIdPCommon field if non-nil, zero value otherwise.

### GetCloudIdPCommonOk

`func (o *AzureConfiguration) GetCloudIdPCommonOk() (*CloudIdPCommon, bool)`

GetCloudIdPCommonOk returns a tuple with the CloudIdPCommon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudIdPCommon

`func (o *AzureConfiguration) SetCloudIdPCommon(v CloudIdPCommon)`

SetCloudIdPCommon sets CloudIdPCommon field to given value.


### GetServer

`func (o *AzureConfiguration) GetServer() AzureServerConfiguration`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *AzureConfiguration) GetServerOk() (*AzureServerConfiguration, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *AzureConfiguration) SetServer(v AzureServerConfiguration)`

SetServer sets Server field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


