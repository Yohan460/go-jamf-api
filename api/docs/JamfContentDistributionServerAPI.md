# \JamfContentDistributionServerAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1JcdsFilesFileNameDelete**](JamfContentDistributionServerAPI.md#V1JcdsFilesFileNameDelete) | **Delete** /v1/jcds/files/{fileName} | Delete a file from the Jamf Content Distribution Server
[**V1JcdsFilesFileNameGet**](JamfContentDistributionServerAPI.md#V1JcdsFilesFileNameGet) | **Get** /v1/jcds/files/{fileName} | Retrieve a download URL for a specific file from the Jamf Content Distribution Server
[**V1JcdsFilesGet**](JamfContentDistributionServerAPI.md#V1JcdsFilesGet) | **Get** /v1/jcds/files | Retrieve a list of files and file metadata from the Jamf Content Distribution Server
[**V1JcdsFilesPost**](JamfContentDistributionServerAPI.md#V1JcdsFilesPost) | **Post** /v1/jcds/files | Initiate an upload to the Jamf Content Distribution Server
[**V1JcdsPropertiesGet**](JamfContentDistributionServerAPI.md#V1JcdsPropertiesGet) | **Get** /v1/jcds/properties | Gets information about JCDS distribution points. 
[**V1JcdsRenewCredentialsPost**](JamfContentDistributionServerAPI.md#V1JcdsRenewCredentialsPost) | **Post** /v1/jcds/renew-credentials | Renew credentials for an upload to the Jamf Content Distribution Server



## V1JcdsFilesFileNameDelete

> V1JcdsFilesFileNameDelete(ctx, fileName).Execute()

Delete a file from the Jamf Content Distribution Server



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
	fileName := "fileName_example" // string | Name of the file that will be deleted from the Jamf Content Distribution Server.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JamfContentDistributionServerAPI.V1JcdsFilesFileNameDelete(context.Background(), fileName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfContentDistributionServerAPI.V1JcdsFilesFileNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileName** | **string** | Name of the file that will be deleted from the Jamf Content Distribution Server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1JcdsFilesFileNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JcdsFilesFileNameGet

> DownloadUrl V1JcdsFilesFileNameGet(ctx, fileName).Execute()

Retrieve a download URL for a specific file from the Jamf Content Distribution Server



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
	fileName := "fileName_example" // string | Name of the file stored in the Jamf Content Distribution Server.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamfContentDistributionServerAPI.V1JcdsFilesFileNameGet(context.Background(), fileName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfContentDistributionServerAPI.V1JcdsFilesFileNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1JcdsFilesFileNameGet`: DownloadUrl
	fmt.Fprintf(os.Stdout, "Response from `JamfContentDistributionServerAPI.V1JcdsFilesFileNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileName** | **string** | Name of the file stored in the Jamf Content Distribution Server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1JcdsFilesFileNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DownloadUrl**](DownloadUrl.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JcdsFilesGet

> []FileData V1JcdsFilesGet(ctx).Execute()

Retrieve a list of files and file metadata from the Jamf Content Distribution Server



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
	resp, r, err := apiClient.JamfContentDistributionServerAPI.V1JcdsFilesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfContentDistributionServerAPI.V1JcdsFilesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1JcdsFilesGet`: []FileData
	fmt.Fprintf(os.Stdout, "Response from `JamfContentDistributionServerAPI.V1JcdsFilesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1JcdsFilesGetRequest struct via the builder pattern


### Return type

[**[]FileData**](FileData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JcdsFilesPost

> Credentials V1JcdsFilesPost(ctx).Execute()

Initiate an upload to the Jamf Content Distribution Server



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
	resp, r, err := apiClient.JamfContentDistributionServerAPI.V1JcdsFilesPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfContentDistributionServerAPI.V1JcdsFilesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1JcdsFilesPost`: Credentials
	fmt.Fprintf(os.Stdout, "Response from `JamfContentDistributionServerAPI.V1JcdsFilesPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1JcdsFilesPostRequest struct via the builder pattern


### Return type

[**Credentials**](Credentials.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JcdsPropertiesGet

> JcdsProperties V1JcdsPropertiesGet(ctx).Execute()

Gets information about JCDS distribution points. 



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
	resp, r, err := apiClient.JamfContentDistributionServerAPI.V1JcdsPropertiesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfContentDistributionServerAPI.V1JcdsPropertiesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1JcdsPropertiesGet`: JcdsProperties
	fmt.Fprintf(os.Stdout, "Response from `JamfContentDistributionServerAPI.V1JcdsPropertiesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1JcdsPropertiesGetRequest struct via the builder pattern


### Return type

[**JcdsProperties**](JcdsProperties.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1JcdsRenewCredentialsPost

> Credentials V1JcdsRenewCredentialsPost(ctx).Execute()

Renew credentials for an upload to the Jamf Content Distribution Server



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
	resp, r, err := apiClient.JamfContentDistributionServerAPI.V1JcdsRenewCredentialsPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamfContentDistributionServerAPI.V1JcdsRenewCredentialsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1JcdsRenewCredentialsPost`: Credentials
	fmt.Fprintf(os.Stdout, "Response from `JamfContentDistributionServerAPI.V1JcdsRenewCredentialsPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1JcdsRenewCredentialsPostRequest struct via the builder pattern


### Return type

[**Credentials**](Credentials.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

