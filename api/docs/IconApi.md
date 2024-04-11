# \IconAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1IconDownloadIdGet**](IconAPI.md#V1IconDownloadIdGet) | **Get** /v1/icon/download/{id} | Download a self service icon 
[**V1IconIdGet**](IconAPI.md#V1IconIdGet) | **Get** /v1/icon/{id} | Get an icon 
[**V1IconPost**](IconAPI.md#V1IconPost) | **Post** /v1/icon | Upload an icon 



## V1IconDownloadIdGet

> *os.File V1IconDownloadIdGet(ctx, id).Res(res).Scale(scale).Execute()

Download a self service icon 



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
	id := "id_example" // string | id of the self service icon
	res := "300" // string | request a specific resolution of original, 300, or 512; invalid options will result in original resolution (optional) (default to "original")
	scale := "1" // string | request a scale; 0 results in original image, non-0 results in scaled to 300 (optional) (default to "0")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IconAPI.V1IconDownloadIdGet(context.Background(), id).Res(res).Scale(scale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IconAPI.V1IconDownloadIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1IconDownloadIdGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `IconAPI.V1IconDownloadIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the self service icon | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1IconDownloadIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **res** | **string** | request a specific resolution of original, 300, or 512; invalid options will result in original resolution | [default to &quot;original&quot;]
 **scale** | **string** | request a scale; 0 results in original image, non-0 results in scaled to 300 | [default to &quot;0&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1IconIdGet

> IconResponse V1IconIdGet(ctx, id).Execute()

Get an icon 



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
	id := "id_example" // string | id of the icon

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IconAPI.V1IconIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IconAPI.V1IconIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1IconIdGet`: IconResponse
	fmt.Fprintf(os.Stdout, "Response from `IconAPI.V1IconIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the icon | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1IconIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IconResponse**](IconResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1IconPost

> IconResponse V1IconPost(ctx).File(file).Execute()

Upload an icon 



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
	file := os.NewFile(1234, "some_file") // *os.File | The file to upload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IconAPI.V1IconPost(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IconAPI.V1IconPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1IconPost`: IconResponse
	fmt.Fprintf(os.Stdout, "Response from `IconAPI.V1IconPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1IconPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The file to upload | 

### Return type

[**IconResponse**](IconResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

