# \InventoryInformationAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1InventoryInformationGet**](InventoryInformationAPI.md#V1InventoryInformationGet) | **Get** /v1/inventory-information | Get statistics about managed/unmanaged devices and computers in the inventory 



## V1InventoryInformationGet

> InventoryInformation V1InventoryInformationGet(ctx).Execute()

Get statistics about managed/unmanaged devices and computers in the inventory 



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
	resp, r, err := apiClient.InventoryInformationAPI.V1InventoryInformationGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryInformationAPI.V1InventoryInformationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1InventoryInformationGet`: InventoryInformation
	fmt.Fprintf(os.Stdout, "Response from `InventoryInformationAPI.V1InventoryInformationGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1InventoryInformationGetRequest struct via the builder pattern


### Return type

[**InventoryInformation**](InventoryInformation.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

