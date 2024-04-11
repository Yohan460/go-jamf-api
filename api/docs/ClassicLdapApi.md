# \ClassicLdapAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ClassicLdapIdGet**](ClassicLdapAPI.md#V1ClassicLdapIdGet) | **Get** /v1/classic-ldap/{id} | Get mappings for OnPrem Ldap configuration with given id.



## V1ClassicLdapIdGet

> ClassicLdapMappings V1ClassicLdapIdGet(ctx, id).Execute()

Get mappings for OnPrem Ldap configuration with given id.



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
	id := "id_example" // string | OnPrem Ldap identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClassicLdapAPI.V1ClassicLdapIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassicLdapAPI.V1ClassicLdapIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ClassicLdapIdGet`: ClassicLdapMappings
	fmt.Fprintf(os.Stdout, "Response from `ClassicLdapAPI.V1ClassicLdapIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | OnPrem Ldap identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ClassicLdapIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClassicLdapMappings**](ClassicLdapMappings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

