# \ActivationCodeAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ActivationCodePut**](ActivationCodeAPI.md#V1ActivationCodePut) | **Put** /v1/activation-code | Updates Activation Code



## V1ActivationCodePut

> V1ActivationCodePut(ctx).ActivationCode(activationCode).Execute()

Updates Activation Code



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
	activationCode := *openapiclient.NewActivationCode("A1A1-B2B2-C3C3-D4D4-E5E5-F6F6-G7G7-H8H8") // ActivationCode |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ActivationCodeAPI.V1ActivationCodePut(context.Background()).ActivationCode(activationCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivationCodeAPI.V1ActivationCodePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ActivationCodePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activationCode** | [**ActivationCode**](ActivationCode.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

