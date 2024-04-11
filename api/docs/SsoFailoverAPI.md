# \SsoFailoverAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SsoFailoverGeneratePost**](SsoFailoverAPI.md#V1SsoFailoverGeneratePost) | **Post** /v1/sso/failover/generate | Regenerates failover url
[**V1SsoFailoverGet**](SsoFailoverAPI.md#V1SsoFailoverGet) | **Get** /v1/sso/failover | Retrieve the current failover settings



## V1SsoFailoverGeneratePost

> SsoFailoverData V1SsoFailoverGeneratePost(ctx).Execute()

Regenerates failover url



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
	resp, r, err := apiClient.SsoFailoverAPI.V1SsoFailoverGeneratePost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoFailoverAPI.V1SsoFailoverGeneratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SsoFailoverGeneratePost`: SsoFailoverData
	fmt.Fprintf(os.Stdout, "Response from `SsoFailoverAPI.V1SsoFailoverGeneratePost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SsoFailoverGeneratePostRequest struct via the builder pattern


### Return type

[**SsoFailoverData**](SsoFailoverData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SsoFailoverGet

> SsoFailoverData V1SsoFailoverGet(ctx).Execute()

Retrieve the current failover settings



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
	resp, r, err := apiClient.SsoFailoverAPI.V1SsoFailoverGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoFailoverAPI.V1SsoFailoverGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SsoFailoverGet`: SsoFailoverData
	fmt.Fprintf(os.Stdout, "Response from `SsoFailoverAPI.V1SsoFailoverGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SsoFailoverGetRequest struct via the builder pattern


### Return type

[**SsoFailoverData**](SsoFailoverData.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

