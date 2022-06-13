# \TomcatSettingsPreviewApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SettingsIssueTomcatSslCertificatePost**](TomcatSettingsPreviewApi.md#SettingsIssueTomcatSslCertificatePost) | **Post** /settings/issueTomcatSslCertificate | Generate a SSL Certificate using Jamf Certificate Authority 



## SettingsIssueTomcatSslCertificatePost

> SettingsIssueTomcatSslCertificatePost(ctx).Execute()

Generate a SSL Certificate using Jamf Certificate Authority 



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TomcatSettingsPreviewApi.SettingsIssueTomcatSslCertificatePost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TomcatSettingsPreviewApi.SettingsIssueTomcatSslCertificatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsIssueTomcatSslCertificatePostRequest struct via the builder pattern


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

