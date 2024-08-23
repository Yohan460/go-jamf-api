# \CsaAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CsaTenantIdGet**](CsaAPI.md#V1CsaTenantIdGet) | **Get** /v1/csa/tenant-id | Returns the CSA tenant ID.
[**V1CsaTokenDelete**](CsaAPI.md#V1CsaTokenDelete) | **Delete** /v1/csa/token | Delete the CSA token exchange - This will disable Jamf Pro&#39;s ability to authenticate with cloud-hosted services 
[**V1CsaTokenGet**](CsaAPI.md#V1CsaTokenGet) | **Get** /v1/csa/token | Get details regarding the CSA token exchange 



## V1CsaTenantIdGet

> CsaTenantIdInfo V1CsaTenantIdGet(ctx).Execute()

Returns the CSA tenant ID.



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
	resp, r, err := apiClient.CsaAPI.V1CsaTenantIdGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CsaAPI.V1CsaTenantIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CsaTenantIdGet`: CsaTenantIdInfo
	fmt.Fprintf(os.Stdout, "Response from `CsaAPI.V1CsaTenantIdGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CsaTenantIdGetRequest struct via the builder pattern


### Return type

[**CsaTenantIdInfo**](CsaTenantIdInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CsaTokenDelete

> V1CsaTokenDelete(ctx).Execute()

Delete the CSA token exchange - This will disable Jamf Pro's ability to authenticate with cloud-hosted services 



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
	r, err := apiClient.CsaAPI.V1CsaTokenDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CsaAPI.V1CsaTokenDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CsaTokenDeleteRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CsaTokenGet

> CsaToken V1CsaTokenGet(ctx).Execute()

Get details regarding the CSA token exchange 



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
	resp, r, err := apiClient.CsaAPI.V1CsaTokenGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CsaAPI.V1CsaTokenGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CsaTokenGet`: CsaToken
	fmt.Fprintf(os.Stdout, "Response from `CsaAPI.V1CsaTokenGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1CsaTokenGetRequest struct via the builder pattern


### Return type

[**CsaToken**](CsaToken.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

