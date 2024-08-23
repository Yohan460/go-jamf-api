# \DssDeclarationsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DssDeclarationsIdGet**](DssDeclarationsAPI.md#V1DssDeclarationsIdGet) | **Get** /v1/dss-declarations/{id} | Retrieve a declaration from DSS



## V1DssDeclarationsIdGet

> DssDeclarations V1DssDeclarationsIdGet(ctx, id).Execute()

Retrieve a declaration from DSS



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
	id := "id_example" // string | Declaration Uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DssDeclarationsAPI.V1DssDeclarationsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DssDeclarationsAPI.V1DssDeclarationsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DssDeclarationsIdGet`: DssDeclarations
	fmt.Fprintf(os.Stdout, "Response from `DssDeclarationsAPI.V1DssDeclarationsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Declaration Uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1DssDeclarationsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DssDeclarations**](DssDeclarations.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

