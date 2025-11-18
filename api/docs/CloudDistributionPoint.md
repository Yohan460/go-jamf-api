# CloudDistributionPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasConnectionSucceeded** | **bool** | Indicates whether the connection to the cloud distribution point was successful.  If &#x60;true&#x60;, the connection was successful. If &#x60;false&#x60;, the connection failed.&lt;br/&gt; Possible values are:&lt;br/&gt; false &lt;br/&gt; true  | [readonly] [default to false]
**Message** | **string** | A message detailing the result of the connection test. This could be a success message or an error message if the connection failed.  | [readonly] 
**InventoryId** | Pointer to **string** | The unique identifier (inventoryId) that links the cloud distribution point to its inventory data. By default, its value is 0, and it increments by +1 based on the existing inventory ID present in the table for each new cloud distribution point configuration. If the cdnType is set to NONE in the next configuration, the ID resets and starts from 1. | [optional] [readonly] [default to "0"]
**CdnType** | **string** | Specifies the content delivery network (CDN) used to distribute content for the cloud distribution point. | [default to "NONE"]
**Master** | Pointer to **bool** | Use as principal distribution point. Use as the authoritative source for all files.&lt;br/&gt; Possible values are:&lt;br/&gt; false &lt;br/&gt; true | [optional] [default to false]
**Username** | **string** | The username or access key used for authenticating with the selected content delivery network (CDN).  This field is required when the **cdnType** is set to **Rackspace Cloud Files(RACKSPACE_CLOUD_FILES)**, **Amazon Web Services(AMAZON_S3)**, or **Akamai(AKAMAI)**, as it is used to authenticate and authorize access to the respective cloud services. - For **Rackspace Cloud Files(RACKSPACE_CLOUD_FILES)**, this is typically the username associated with your Rackspace cloud account. - For **Amazon Web Services(AMAZON_S3)**, this corresponds to the **Access Key ID** used to interact with Amazon Web Services(AMAZON_S3) resources. - For **Akamai(AKAMAI)**, this is the username used for API authentication to access Akamai&#39;s content delivery services. If the **cdnType** is **None**, this field is not applicable.  | 
**Password** | **string** | The password or authentication key used for connecting to the selected content delivery network (CDN).  This field is required when the **cdnType** is set to **Rackspace Cloud Files(RACKSPACE_CLOUD_FILES)**, **Amazon Web Services(AMAZON_S3)**, or **Akamai(AKAMAI)**, and is used to authenticate and authorize access to the respective cloud services. - For **Rackspace Cloud Files(RACKSPACE_CLOUD_FILES)**, this refers to the **API Key** that is used in conjunction with the username for authenticating API requests. - For **Amazon Web Services(AMAZON_S3)**, this corresponds to the **Secret Access Key** associated with your AWS account, used to securely sign requests to AWS services. - For **Akamai(AKAMAI)**, this is the **password** used for API authentication to access Akamai&#39;s content delivery services. If the **cdnType** is **None**, this field is not applicable.  | 
**Directory** | Pointer to **string** | The directory or path for content delivery in Akamai. This field is required when the **cdnType** is set to **Akamai(AKAMAI)** and specifies where content is stored within Akamai&#39;s system. | [optional] 
**CdnUrl** | Pointer to **string** | The CDN URL for the cloud distribution point. The URL format varies depending on the selected CDN provider: - **Rackspace Cloud Files(RACKSPACE_CLOUD_FILES)** - **Amazon Web Services(AMAZON_S3)**  - **Akamai(AKAMAI)**  The **cdnUrl** should point to the content distribution location where software or other content is stored and made available for distribution.  | [optional] [readonly] 
**UploadUrl** | Pointer to **string** | The URL used to upload files to Akamai&#39;s NetStorage. This field is required when the **cdnType** is set to **Akamai(AKAMAI)**.  It specifies where content should be uploaded to Akamai’s cloud storage before being distributed via their CDN. The upload typically uses FTP or SFTP.  | [optional] 
**DownloadUrl** | Pointer to **string** | The URL used to access and download content from Akamai&#39;s EdgeSuite. This field is required when the **cdnType** is set to **Akamai(AKAMAI)**.  It specifies the endpoint from which files are retrieved by devices or users.  | [optional] 
**SecondaryAuthRequired** | Pointer to **bool** | Enable Remote Authentication.Authorize requests for files stored on the distribution point. This field is required when the **cdnType** is set to **Akamai(AKAMAI)**.&lt;br/&gt; Possible values are:&lt;br/&gt; false &lt;br/&gt; true | [optional] [default to false]
**SecondaryAuthStatusCode** | Pointer to **int64** | Secondary Auth Status Code. Configure the HTTP response code that will be returned by Jamf Pro during remote authentication. This field is required when the **cdnType** is set to **Akamai(AKAMAI)** and **secondaryAuthRequired** is true. | [optional] [default to 200]
**SecondaryAuthTimeToLive** | Pointer to **int64** | Secondary Auth Time To Live. Number of seconds before the authorization token expires. This field is required when the **cdnType** is set to **Akamai(AKAMAI)** and **secondaryAuthRequired** is true. | [optional] [default to 3600]
**RequireSignedUrls** | Pointer to **bool** | Amazon Sign Url. It restrict access to requests that use a signed URL. This field is required when the **cdnType** is set to **Amazon Web Services(AMAZON_S3)**.&lt;br/&gt; Possible values are:&lt;br/&gt; false &lt;br/&gt; true | [optional] [default to false]
**KeyPairId** | Pointer to **string** | The CloudFront Access Key ID (keyPairId) is part of the credentials used to generate signed URLs for secure access to content in a CloudFront distribution. When using AWS, this key is paired with the CloudFront Secret Access Key to create the signed URL, ensuring that only authorized users can access specific content within a specified timeframe. This field is required when the **cdnType** is set to **Amazon Web Services(AMAZON_S3)** and **requireSignedUrls** is true.  | [optional] 
**ExpirationSeconds** | Pointer to **int64** | Signed URL Expiration. Number of seconds before the signed URL expires, This field is required when the **cdnType** is set to **Amazon Web Services(AMAZON_S3)** and **requireSignedUrls** is true. | [optional] [default to 3600]
**PrivateKey** | Pointer to **string** | The CloudFront Private Key file is required when the **cdnType** is set to **Amazon Web Services(AMAZON_S3)** and **requireSignedUrls** parameter is enabled. This private key is used for signing URLs for restricted access to CloudFront-distributed content.  The private key allows secure URL generation for signed URLs, ensuring that only authorized users can access certain content. The key must be uploaded in one of the following formats: - **.pem**: A Privacy-Enhanced Mail (PEM) file containing the private key in base64 encoded format. - **.der**: A Distinguished Encoding Rules (DER) encoded file, which is a binary format for the private key. The uploaded file should be kept secure, as it provides the ability to generate signed URLs with access to protected content.  | [optional] 

## Methods

### NewCloudDistributionPoint

`func NewCloudDistributionPoint(hasConnectionSucceeded bool, message string, cdnType string, username string, password string, ) *CloudDistributionPoint`

NewCloudDistributionPoint instantiates a new CloudDistributionPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDistributionPointWithDefaults

`func NewCloudDistributionPointWithDefaults() *CloudDistributionPoint`

NewCloudDistributionPointWithDefaults instantiates a new CloudDistributionPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasConnectionSucceeded

`func (o *CloudDistributionPoint) GetHasConnectionSucceeded() bool`

GetHasConnectionSucceeded returns the HasConnectionSucceeded field if non-nil, zero value otherwise.

### GetHasConnectionSucceededOk

`func (o *CloudDistributionPoint) GetHasConnectionSucceededOk() (*bool, bool)`

GetHasConnectionSucceededOk returns a tuple with the HasConnectionSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConnectionSucceeded

`func (o *CloudDistributionPoint) SetHasConnectionSucceeded(v bool)`

SetHasConnectionSucceeded sets HasConnectionSucceeded field to given value.


### GetMessage

`func (o *CloudDistributionPoint) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CloudDistributionPoint) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CloudDistributionPoint) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetInventoryId

`func (o *CloudDistributionPoint) GetInventoryId() string`

GetInventoryId returns the InventoryId field if non-nil, zero value otherwise.

### GetInventoryIdOk

`func (o *CloudDistributionPoint) GetInventoryIdOk() (*string, bool)`

GetInventoryIdOk returns a tuple with the InventoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryId

`func (o *CloudDistributionPoint) SetInventoryId(v string)`

SetInventoryId sets InventoryId field to given value.

### HasInventoryId

`func (o *CloudDistributionPoint) HasInventoryId() bool`

HasInventoryId returns a boolean if a field has been set.

### GetCdnType

`func (o *CloudDistributionPoint) GetCdnType() string`

GetCdnType returns the CdnType field if non-nil, zero value otherwise.

### GetCdnTypeOk

`func (o *CloudDistributionPoint) GetCdnTypeOk() (*string, bool)`

GetCdnTypeOk returns a tuple with the CdnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnType

`func (o *CloudDistributionPoint) SetCdnType(v string)`

SetCdnType sets CdnType field to given value.


### GetMaster

`func (o *CloudDistributionPoint) GetMaster() bool`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *CloudDistributionPoint) GetMasterOk() (*bool, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *CloudDistributionPoint) SetMaster(v bool)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *CloudDistributionPoint) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetUsername

`func (o *CloudDistributionPoint) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CloudDistributionPoint) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CloudDistributionPoint) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CloudDistributionPoint) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CloudDistributionPoint) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CloudDistributionPoint) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDirectory

`func (o *CloudDistributionPoint) GetDirectory() string`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *CloudDistributionPoint) GetDirectoryOk() (*string, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *CloudDistributionPoint) SetDirectory(v string)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *CloudDistributionPoint) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### GetCdnUrl

`func (o *CloudDistributionPoint) GetCdnUrl() string`

GetCdnUrl returns the CdnUrl field if non-nil, zero value otherwise.

### GetCdnUrlOk

`func (o *CloudDistributionPoint) GetCdnUrlOk() (*string, bool)`

GetCdnUrlOk returns a tuple with the CdnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnUrl

`func (o *CloudDistributionPoint) SetCdnUrl(v string)`

SetCdnUrl sets CdnUrl field to given value.

### HasCdnUrl

`func (o *CloudDistributionPoint) HasCdnUrl() bool`

HasCdnUrl returns a boolean if a field has been set.

### GetUploadUrl

`func (o *CloudDistributionPoint) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *CloudDistributionPoint) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *CloudDistributionPoint) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *CloudDistributionPoint) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *CloudDistributionPoint) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *CloudDistributionPoint) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *CloudDistributionPoint) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *CloudDistributionPoint) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetSecondaryAuthRequired

`func (o *CloudDistributionPoint) GetSecondaryAuthRequired() bool`

GetSecondaryAuthRequired returns the SecondaryAuthRequired field if non-nil, zero value otherwise.

### GetSecondaryAuthRequiredOk

`func (o *CloudDistributionPoint) GetSecondaryAuthRequiredOk() (*bool, bool)`

GetSecondaryAuthRequiredOk returns a tuple with the SecondaryAuthRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAuthRequired

`func (o *CloudDistributionPoint) SetSecondaryAuthRequired(v bool)`

SetSecondaryAuthRequired sets SecondaryAuthRequired field to given value.

### HasSecondaryAuthRequired

`func (o *CloudDistributionPoint) HasSecondaryAuthRequired() bool`

HasSecondaryAuthRequired returns a boolean if a field has been set.

### GetSecondaryAuthStatusCode

`func (o *CloudDistributionPoint) GetSecondaryAuthStatusCode() int64`

GetSecondaryAuthStatusCode returns the SecondaryAuthStatusCode field if non-nil, zero value otherwise.

### GetSecondaryAuthStatusCodeOk

`func (o *CloudDistributionPoint) GetSecondaryAuthStatusCodeOk() (*int64, bool)`

GetSecondaryAuthStatusCodeOk returns a tuple with the SecondaryAuthStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAuthStatusCode

`func (o *CloudDistributionPoint) SetSecondaryAuthStatusCode(v int64)`

SetSecondaryAuthStatusCode sets SecondaryAuthStatusCode field to given value.

### HasSecondaryAuthStatusCode

`func (o *CloudDistributionPoint) HasSecondaryAuthStatusCode() bool`

HasSecondaryAuthStatusCode returns a boolean if a field has been set.

### GetSecondaryAuthTimeToLive

`func (o *CloudDistributionPoint) GetSecondaryAuthTimeToLive() int64`

GetSecondaryAuthTimeToLive returns the SecondaryAuthTimeToLive field if non-nil, zero value otherwise.

### GetSecondaryAuthTimeToLiveOk

`func (o *CloudDistributionPoint) GetSecondaryAuthTimeToLiveOk() (*int64, bool)`

GetSecondaryAuthTimeToLiveOk returns a tuple with the SecondaryAuthTimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAuthTimeToLive

`func (o *CloudDistributionPoint) SetSecondaryAuthTimeToLive(v int64)`

SetSecondaryAuthTimeToLive sets SecondaryAuthTimeToLive field to given value.

### HasSecondaryAuthTimeToLive

`func (o *CloudDistributionPoint) HasSecondaryAuthTimeToLive() bool`

HasSecondaryAuthTimeToLive returns a boolean if a field has been set.

### GetRequireSignedUrls

`func (o *CloudDistributionPoint) GetRequireSignedUrls() bool`

GetRequireSignedUrls returns the RequireSignedUrls field if non-nil, zero value otherwise.

### GetRequireSignedUrlsOk

`func (o *CloudDistributionPoint) GetRequireSignedUrlsOk() (*bool, bool)`

GetRequireSignedUrlsOk returns a tuple with the RequireSignedUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSignedUrls

`func (o *CloudDistributionPoint) SetRequireSignedUrls(v bool)`

SetRequireSignedUrls sets RequireSignedUrls field to given value.

### HasRequireSignedUrls

`func (o *CloudDistributionPoint) HasRequireSignedUrls() bool`

HasRequireSignedUrls returns a boolean if a field has been set.

### GetKeyPairId

`func (o *CloudDistributionPoint) GetKeyPairId() string`

GetKeyPairId returns the KeyPairId field if non-nil, zero value otherwise.

### GetKeyPairIdOk

`func (o *CloudDistributionPoint) GetKeyPairIdOk() (*string, bool)`

GetKeyPairIdOk returns a tuple with the KeyPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPairId

`func (o *CloudDistributionPoint) SetKeyPairId(v string)`

SetKeyPairId sets KeyPairId field to given value.

### HasKeyPairId

`func (o *CloudDistributionPoint) HasKeyPairId() bool`

HasKeyPairId returns a boolean if a field has been set.

### GetExpirationSeconds

`func (o *CloudDistributionPoint) GetExpirationSeconds() int64`

GetExpirationSeconds returns the ExpirationSeconds field if non-nil, zero value otherwise.

### GetExpirationSecondsOk

`func (o *CloudDistributionPoint) GetExpirationSecondsOk() (*int64, bool)`

GetExpirationSecondsOk returns a tuple with the ExpirationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationSeconds

`func (o *CloudDistributionPoint) SetExpirationSeconds(v int64)`

SetExpirationSeconds sets ExpirationSeconds field to given value.

### HasExpirationSeconds

`func (o *CloudDistributionPoint) HasExpirationSeconds() bool`

HasExpirationSeconds returns a boolean if a field has been set.

### GetPrivateKey

`func (o *CloudDistributionPoint) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CloudDistributionPoint) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CloudDistributionPoint) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *CloudDistributionPoint) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


