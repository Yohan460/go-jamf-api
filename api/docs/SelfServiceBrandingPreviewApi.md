# \SelfServiceBrandingPreviewApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SelfServiceBrandingImagesPost**](SelfServiceBrandingPreviewApi.md#SelfServiceBrandingImagesPost) | **Post** /self-service/branding/images | Upload an image 



## SelfServiceBrandingImagesPost

> BrandingImageUrl SelfServiceBrandingImagesPost(ctx).File(file).Execute()

Upload an image 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    file := os.NewFile(1234, "some_file") // *os.File | The file to upload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SelfServiceBrandingPreviewApi.SelfServiceBrandingImagesPost(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceBrandingPreviewApi.SelfServiceBrandingImagesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SelfServiceBrandingImagesPost`: BrandingImageUrl
    fmt.Fprintf(os.Stdout, "Response from `SelfServiceBrandingPreviewApi.SelfServiceBrandingImagesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSelfServiceBrandingImagesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The file to upload | 

### Return type

[**BrandingImageUrl**](BrandingImageUrl.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

