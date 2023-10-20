# \AppStoreCountryCodesPreviewAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AppStoreCountryCodesGet**](AppStoreCountryCodesPreviewAPI.md#V1AppStoreCountryCodesGet) | **Get** /v1/app-store-country-codes | Return a list of Countries and the associated Codes 



## V1AppStoreCountryCodesGet

> CountryCodes V1AppStoreCountryCodesGet(ctx).Execute()

Return a list of Countries and the associated Codes 



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
    resp, r, err := apiClient.AppStoreCountryCodesPreviewAPI.V1AppStoreCountryCodesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppStoreCountryCodesPreviewAPI.V1AppStoreCountryCodesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AppStoreCountryCodesGet`: CountryCodes
    fmt.Fprintf(os.Stdout, "Response from `AppStoreCountryCodesPreviewAPI.V1AppStoreCountryCodesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AppStoreCountryCodesGetRequest struct via the builder pattern


### Return type

[**CountryCodes**](CountryCodes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

