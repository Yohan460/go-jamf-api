# \LocalesPreviewAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1LocalesGet**](LocalesPreviewAPI.md#V1LocalesGet) | **Get** /v1/locales | Return locales that can be used in other features 



## V1LocalesGet

> []Locale V1LocalesGet(ctx).Execute()

Return locales that can be used in other features 



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
	resp, r, err := apiClient.LocalesPreviewAPI.V1LocalesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalesPreviewAPI.V1LocalesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1LocalesGet`: []Locale
	fmt.Fprintf(os.Stdout, "Response from `LocalesPreviewAPI.V1LocalesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1LocalesGetRequest struct via the builder pattern


### Return type

[**[]Locale**](Locale.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

