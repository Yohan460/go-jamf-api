# \TomcatSettingsPreviewAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SettingsIssueTomcatSslCertificatePost**](TomcatSettingsPreviewAPI.md#SettingsIssueTomcatSslCertificatePost) | **Post** /settings/issueTomcatSslCertificate | Generate a SSL Certificate using Jamf Certificate Authority 



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
	openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TomcatSettingsPreviewAPI.SettingsIssueTomcatSslCertificatePost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TomcatSettingsPreviewAPI.SettingsIssueTomcatSslCertificatePost``: %v\n", err)
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

