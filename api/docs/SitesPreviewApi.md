# \SitesPreviewApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SettingsSitesGet**](SitesPreviewApi.md#SettingsSitesGet) | **Get** /settings/sites | Find all sites 



## SettingsSitesGet

> []Site SettingsSitesGet(ctx).Execute()

Find all sites 



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
    resp, r, err := apiClient.SitesPreviewApi.SettingsSitesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesPreviewApi.SettingsSitesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SettingsSitesGet`: []Site
    fmt.Fprintf(os.Stdout, "Response from `SitesPreviewApi.SettingsSitesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsSitesGetRequest struct via the builder pattern


### Return type

[**[]Site**](Site.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

