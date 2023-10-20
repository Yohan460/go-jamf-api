# \PatchManagementAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2PatchManagementAcceptDisclaimerPost**](PatchManagementAPI.md#V2PatchManagementAcceptDisclaimerPost) | **Post** /v2/patch-management-accept-disclaimer | Accept Patch Management disclaimer 



## V2PatchManagementAcceptDisclaimerPost

> V2PatchManagementAcceptDisclaimerPost(ctx).Execute()

Accept Patch Management disclaimer 



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
    r, err := apiClient.PatchManagementAPI.V2PatchManagementAcceptDisclaimerPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PatchManagementAPI.V2PatchManagementAcceptDisclaimerPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2PatchManagementAcceptDisclaimerPostRequest struct via the builder pattern


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

