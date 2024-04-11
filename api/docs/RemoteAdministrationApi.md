# \RemoteAdministrationAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreviewRemoteAdministrationConfigurationsGet**](RemoteAdministrationAPI.md#PreviewRemoteAdministrationConfigurationsGet) | **Get** /preview/remote-administration-configurations | Get information about all remote administration configurations.



## PreviewRemoteAdministrationConfigurationsGet

> RemoteAdministrationSearchResults PreviewRemoteAdministrationConfigurationsGet(ctx).Page(page).PageSize(pageSize).Execute()

Get information about all remote administration configurations.



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
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteAdministrationAPI.PreviewRemoteAdministrationConfigurationsGet(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteAdministrationAPI.PreviewRemoteAdministrationConfigurationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreviewRemoteAdministrationConfigurationsGet`: RemoteAdministrationSearchResults
	fmt.Fprintf(os.Stdout, "Response from `RemoteAdministrationAPI.PreviewRemoteAdministrationConfigurationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewRemoteAdministrationConfigurationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]

### Return type

[**RemoteAdministrationSearchResults**](RemoteAdministrationSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

