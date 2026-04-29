# \UserSessionsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1UserSessionsActiveGet**](UserSessionsAPI.md#V1UserSessionsActiveGet) | **Get** /v1/user-sessions/active | Get active user sessions. 
[**V1UserSessionsCountGet**](UserSessionsAPI.md#V1UserSessionsCountGet) | **Get** /v1/user-sessions/count | Get count of active user sessions. 



## V1UserSessionsActiveGet

> []ActiveUserSession V1UserSessionsActiveGet(ctx).Execute()

Get active user sessions. 



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
	resp, r, err := apiClient.UserSessionsAPI.V1UserSessionsActiveGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSessionsAPI.V1UserSessionsActiveGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UserSessionsActiveGet`: []ActiveUserSession
	fmt.Fprintf(os.Stdout, "Response from `UserSessionsAPI.V1UserSessionsActiveGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1UserSessionsActiveGetRequest struct via the builder pattern


### Return type

[**[]ActiveUserSession**](ActiveUserSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UserSessionsCountGet

> ActiveUsersCount V1UserSessionsCountGet(ctx).Execute()

Get count of active user sessions. 



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
	resp, r, err := apiClient.UserSessionsAPI.V1UserSessionsCountGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSessionsAPI.V1UserSessionsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UserSessionsCountGet`: ActiveUsersCount
	fmt.Fprintf(os.Stdout, "Response from `UserSessionsAPI.V1UserSessionsCountGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1UserSessionsCountGetRequest struct via the builder pattern


### Return type

[**ActiveUsersCount**](ActiveUsersCount.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

