# \SsoOauthSessionTokensAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1Oauth2SessionTokensGet**](SsoOauthSessionTokensAPI.md#V1Oauth2SessionTokensGet) | **Get** /v1/oauth2/session-tokens | Retrieve the access token and user information for the current session. 



## V1Oauth2SessionTokensGet

> OauthTokens V1Oauth2SessionTokensGet(ctx).Execute()

Retrieve the access token and user information for the current session. 



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
	resp, r, err := apiClient.SsoOauthSessionTokensAPI.V1Oauth2SessionTokensGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoOauthSessionTokensAPI.V1Oauth2SessionTokensGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1Oauth2SessionTokensGet`: OauthTokens
	fmt.Fprintf(os.Stdout, "Response from `SsoOauthSessionTokensAPI.V1Oauth2SessionTokensGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1Oauth2SessionTokensGetRequest struct via the builder pattern


### Return type

[**OauthTokens**](OauthTokens.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [ApiClient](../README.md#ApiClient), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

