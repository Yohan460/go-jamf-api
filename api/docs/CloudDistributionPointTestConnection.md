# CloudDistributionPointTestConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasConnectionSucceeded** | **bool** | Indicates whether the connection to the cloud distribution point was successful.  If &#x60;true&#x60;, the connection was successful. If &#x60;false&#x60;, the connection failed.&lt;br/&gt; Possible values are:&lt;br/&gt; false &lt;br/&gt; true  | [readonly] [default to false]
**Message** | **string** | A message detailing the result of the connection test. This could be a success message or an error message if the connection failed.  | [readonly] 

## Methods

### NewCloudDistributionPointTestConnection

`func NewCloudDistributionPointTestConnection(hasConnectionSucceeded bool, message string, ) *CloudDistributionPointTestConnection`

NewCloudDistributionPointTestConnection instantiates a new CloudDistributionPointTestConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDistributionPointTestConnectionWithDefaults

`func NewCloudDistributionPointTestConnectionWithDefaults() *CloudDistributionPointTestConnection`

NewCloudDistributionPointTestConnectionWithDefaults instantiates a new CloudDistributionPointTestConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasConnectionSucceeded

`func (o *CloudDistributionPointTestConnection) GetHasConnectionSucceeded() bool`

GetHasConnectionSucceeded returns the HasConnectionSucceeded field if non-nil, zero value otherwise.

### GetHasConnectionSucceededOk

`func (o *CloudDistributionPointTestConnection) GetHasConnectionSucceededOk() (*bool, bool)`

GetHasConnectionSucceededOk returns a tuple with the HasConnectionSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConnectionSucceeded

`func (o *CloudDistributionPointTestConnection) SetHasConnectionSucceeded(v bool)`

SetHasConnectionSucceeded sets HasConnectionSucceeded field to given value.


### GetMessage

`func (o *CloudDistributionPointTestConnection) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CloudDistributionPointTestConnection) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CloudDistributionPointTestConnection) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


