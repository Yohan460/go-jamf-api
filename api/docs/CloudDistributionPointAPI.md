# \CloudDistributionPointAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CloudDistributionPointUploadCapabilityGet**](CloudDistributionPointAPI.md#V1CloudDistributionPointUploadCapabilityGet) | **Get** /v1/cloud-distribution-point/upload-capability | Finds specific information for the currently configured Cloud Distribution Point. 



## V1CloudDistributionPointUploadCapabilityGet

> CloudDistributionPointUploadCapability V1CloudDistributionPointUploadCapabilityGet(ctx).Execute()

Finds specific information for the currently configured Cloud Distribution Point. 



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
	resp, r, err := apiClient.CloudDistributionPointAPI.V1CloudDistributionPointUploadCapabilityGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudDistributionPointAPI.V1CloudDistributionPointUploadCapabilityGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CloudDistributionPointUploadCapabilityGet`: CloudDistributionPointUploadCapability
	fmt.Fprintf(os.Stdout, "Response from `CloudDistributionPointAPI.V1CloudDistributionPointUploadCapabilityGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CloudDistributionPointUploadCapabilityGetRequest struct via the builder pattern


### Return type

[**CloudDistributionPointUploadCapability**](CloudDistributionPointUploadCapability.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

