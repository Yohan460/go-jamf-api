# AzureAdMigrationReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdapServerId** | **int32** |  | 
**AzureServerId** | **int32** |  | 
**AzureMappings** | [**AzureMappings**](AzureMappings.md) |  | 

## Methods

### NewAzureAdMigrationReportRequest

`func NewAzureAdMigrationReportRequest(ldapServerId int32, azureServerId int32, azureMappings AzureMappings, ) *AzureAdMigrationReportRequest`

NewAzureAdMigrationReportRequest instantiates a new AzureAdMigrationReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureAdMigrationReportRequestWithDefaults

`func NewAzureAdMigrationReportRequestWithDefaults() *AzureAdMigrationReportRequest`

NewAzureAdMigrationReportRequestWithDefaults instantiates a new AzureAdMigrationReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdapServerId

`func (o *AzureAdMigrationReportRequest) GetLdapServerId() int32`

GetLdapServerId returns the LdapServerId field if non-nil, zero value otherwise.

### GetLdapServerIdOk

`func (o *AzureAdMigrationReportRequest) GetLdapServerIdOk() (*int32, bool)`

GetLdapServerIdOk returns a tuple with the LdapServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerId

`func (o *AzureAdMigrationReportRequest) SetLdapServerId(v int32)`

SetLdapServerId sets LdapServerId field to given value.


### GetAzureServerId

`func (o *AzureAdMigrationReportRequest) GetAzureServerId() int32`

GetAzureServerId returns the AzureServerId field if non-nil, zero value otherwise.

### GetAzureServerIdOk

`func (o *AzureAdMigrationReportRequest) GetAzureServerIdOk() (*int32, bool)`

GetAzureServerIdOk returns a tuple with the AzureServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureServerId

`func (o *AzureAdMigrationReportRequest) SetAzureServerId(v int32)`

SetAzureServerId sets AzureServerId field to given value.


### GetAzureMappings

`func (o *AzureAdMigrationReportRequest) GetAzureMappings() AzureMappings`

GetAzureMappings returns the AzureMappings field if non-nil, zero value otherwise.

### GetAzureMappingsOk

`func (o *AzureAdMigrationReportRequest) GetAzureMappingsOk() (*AzureMappings, bool)`

GetAzureMappingsOk returns a tuple with the AzureMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureMappings

`func (o *AzureAdMigrationReportRequest) SetAzureMappings(v AzureMappings)`

SetAzureMappings sets AzureMappings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


