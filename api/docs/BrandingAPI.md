# \BrandingAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1BrandingImagesDownloadIdGet**](BrandingAPI.md#V1BrandingImagesDownloadIdGet) | **Get** /v1/branding-images/download/{id} | Download a self service branding image 



## V1BrandingImagesDownloadIdGet

> *os.File V1BrandingImagesDownloadIdGet(ctx, id).Execute()

Download a self service branding image 



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
    id := "id_example" // string | id of the self service branding image

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingAPI.V1BrandingImagesDownloadIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingAPI.V1BrandingImagesDownloadIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1BrandingImagesDownloadIdGet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `BrandingAPI.V1BrandingImagesDownloadIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the self service branding image | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1BrandingImagesDownloadIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

