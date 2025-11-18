# \JamfCloudDistributionServiceAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1JcdsFilesFileNameDelete**](JamfCloudDistributionServiceAPI.md#V1JcdsFilesFileNameDelete) | **Delete** /v1/jcds/files/{fileName} | Delete a file from the Jamf Cloud Distribution Service
[**V1JcdsFilesFileNameGet**](JamfCloudDistributionServiceAPI.md#V1JcdsFilesFileNameGet) | **Get** /v1/jcds/files/{fileName} | Retrieve a download URL for a specific file from the Jamf Cloud Distribution Service
[**V1JcdsFilesGet**](JamfCloudDistributionServiceAPI.md#V1JcdsFilesGet) | **Get** /v1/jcds/files | Retrieve a list of files and file metadata from the Jamf Cloud Distribution Service
[**V1JcdsFilesPost**](JamfCloudDistributionServiceAPI.md#V1JcdsFilesPost) | **Post** /v1/jcds/files | Initiate an upload to the Jamf Cloud Distribution Service
[**V1JcdsRefreshInventoryPost**](JamfCloudDistributionServiceAPI.md#V1JcdsRefreshInventoryPost) | **Post** /v1/jcds/refresh-inventory | Refreshes the inventory and status of uploads in Jamf Pro. This will update the  status of uploads in the Jamf Pro database and allow the uploads to be deployed. 
[**V1JcdsRenewCredentialsPost**](JamfCloudDistributionServiceAPI.md#V1JcdsRenewCredentialsPost) | **Post** /v1/jcds/renew-credentials | Renew credentials for an upload to the Jamf Cloud Distribution Service



## V1JcdsFilesFileNameDelete

> V1JcdsFilesFileNameDelete(ctx, fileName).Execute()

Delete a file from the Jamf Cloud Distribution Service



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {
	fileName := "fileName_example" // string | Name of the file that will be deleted from the Jamf Cloud Distribution Service.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JamfCloudDistributionServiceAPI.V1JcdsFilesFileNameDelete(context.Background(), fileName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfCloudDistributionServiceAPI.V1JcdsFilesFileNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileName** | **string** | Name of the file that will be deleted from the Jamf Cloud Distribution Service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1JcdsFilesFileNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JcdsFilesFileNameGet

> DownloadUrl V1JcdsFilesFileNameGet(ctx, fileName).Execute()

Retrieve a download URL for a specific file from the Jamf Cloud Distribution Service



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {
	fileName := "fileName_example" // string | Name of the file stored in the Jamf Cloud Distribution Service.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfCloudDistributionServiceAPI.V1JcdsFilesFileNameGet(context.Background(), fileName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfCloudDistributionServiceAPI.V1JcdsFilesFileNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1JcdsFilesFileNameGet`: DownloadUrl
	fmt.Fprintf(os.Stdout, "Response from `JamfCloudDistributionServiceAPI.V1JcdsFilesFileNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileName** | **string** | Name of the file stored in the Jamf Cloud Distribution Service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1JcdsFilesFileNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DownloadUrl**](DownloadUrl.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JcdsFilesGet

> []FileData V1JcdsFilesGet(ctx).Execute()

Retrieve a list of files and file metadata from the Jamf Cloud Distribution Service



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfCloudDistributionServiceAPI.V1JcdsFilesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfCloudDistributionServiceAPI.V1JcdsFilesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1JcdsFilesGet`: []FileData
	fmt.Fprintf(os.Stdout, "Response from `JamfCloudDistributionServiceAPI.V1JcdsFilesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1JcdsFilesGetRequest struct via the builder pattern


### Return type

[**[]FileData**](FileData.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JcdsFilesPost

> Credentials V1JcdsFilesPost(ctx).Execute()

Initiate an upload to the Jamf Cloud Distribution Service



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfCloudDistributionServiceAPI.V1JcdsFilesPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfCloudDistributionServiceAPI.V1JcdsFilesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1JcdsFilesPost`: Credentials
	fmt.Fprintf(os.Stdout, "Response from `JamfCloudDistributionServiceAPI.V1JcdsFilesPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1JcdsFilesPostRequest struct via the builder pattern


### Return type

[**Credentials**](Credentials.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JcdsRefreshInventoryPost

> V1JcdsRefreshInventoryPost(ctx).FileName(fileName).Execute()

Refreshes the inventory and status of uploads in Jamf Pro. This will update the  status of uploads in the Jamf Pro database and allow the uploads to be deployed. 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {
	fileName := "filename.pkg" // string | Name of the file to check the availability of in JCDS. If available, the inventory and status will be updated in Jamf Pro. If no file is specified, it will force an immediate inventory refresh at a rate-limit of once every 15 seconds. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JamfCloudDistributionServiceAPI.V1JcdsRefreshInventoryPost(context.Background()).FileName(fileName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfCloudDistributionServiceAPI.V1JcdsRefreshInventoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1JcdsRefreshInventoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileName** | **string** | Name of the file to check the availability of in JCDS. If available, the inventory and status will be updated in Jamf Pro. If no file is specified, it will force an immediate inventory refresh at a rate-limit of once every 15 seconds. | [default to &quot;&quot;]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JcdsRenewCredentialsPost

> Credentials V1JcdsRenewCredentialsPost(ctx).Execute()

Renew credentials for an upload to the Jamf Cloud Distribution Service



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfCloudDistributionServiceAPI.V1JcdsRenewCredentialsPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfCloudDistributionServiceAPI.V1JcdsRenewCredentialsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1JcdsRenewCredentialsPost`: Credentials
	fmt.Fprintf(os.Stdout, "Response from `JamfCloudDistributionServiceAPI.V1JcdsRenewCredentialsPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1JcdsRenewCredentialsPostRequest struct via the builder pattern


### Return type

[**Credentials**](Credentials.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

