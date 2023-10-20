# \VppAdminAccountsPreviewAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VppAdminAccountsGet**](VppAdminAccountsPreviewAPI.md#VppAdminAccountsGet) | **Get** /vpp/admin-accounts | Found all VPP Admin Accounts 



## VppAdminAccountsGet

> []VppAdminAccount VppAdminAccountsGet(ctx).Execute()

Found all VPP Admin Accounts 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VppAdminAccountsPreviewAPI.VppAdminAccountsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VppAdminAccountsPreviewAPI.VppAdminAccountsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VppAdminAccountsGet`: []VppAdminAccount
    fmt.Fprintf(os.Stdout, "Response from `VppAdminAccountsPreviewAPI.VppAdminAccountsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVppAdminAccountsGetRequest struct via the builder pattern


### Return type

[**[]VppAdminAccount**](VppAdminAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

