# \LastLoginAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1LastLoginGet**](LastLoginAPI.md#V1LastLoginGet) | **Get** /v1/last-login | Get the date of the last login event



## V1LastLoginGet

> LastLoginResponse V1LastLoginGet(ctx).Execute()

Get the date of the last login event



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
	resp, r, err := apiClient.LastLoginAPI.V1LastLoginGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LastLoginAPI.V1LastLoginGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LastLoginGet`: LastLoginResponse
	fmt.Fprintf(os.Stdout, "Response from `LastLoginAPI.V1LastLoginGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1LastLoginGetRequest struct via the builder pattern


### Return type

[**LastLoginResponse**](LastLoginResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

