# \ApiAuthenticationApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthCurrentPost**](ApiAuthenticationApi.md#AuthCurrentPost) | **Post** /auth/current | Get the authorization details associated with the current API token 
[**AuthGet**](ApiAuthenticationApi.md#AuthGet) | **Get** /auth | Get all the Authorization details associated with the current api 
[**AuthInvalidateTokenPost**](ApiAuthenticationApi.md#AuthInvalidateTokenPost) | **Post** /auth/invalidateToken | Invalidate current token 
[**AuthKeepAlivePost**](ApiAuthenticationApi.md#AuthKeepAlivePost) | **Post** /auth/keepAlive | Invalidate existing token and generates new token 
[**AuthTokensPost**](ApiAuthenticationApi.md#AuthTokensPost) | **Post** /auth/tokens | Create a token based on other authentication details (basic, etc.) 
[**V1AuthGet**](ApiAuthenticationApi.md#V1AuthGet) | **Get** /v1/auth | Get all the Authorization details associated with the current api 
[**V1AuthInvalidateTokenPost**](ApiAuthenticationApi.md#V1AuthInvalidateTokenPost) | **Post** /v1/auth/invalidate-token | Invalidate current token 
[**V1AuthKeepAlivePost**](ApiAuthenticationApi.md#V1AuthKeepAlivePost) | **Post** /v1/auth/keep-alive | Invalidate existing token and generates new token 
[**V1AuthTokenPost**](ApiAuthenticationApi.md#V1AuthTokenPost) | **Post** /v1/auth/token | Create a token based on other authentication details (basic, etc.) 



## AuthCurrentPost

> CurrentAuthorization AuthCurrentPost(ctx).Execute()

Get the authorization details associated with the current API token 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiAuthenticationApi.AuthCurrentPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiAuthenticationApi.AuthCurrentPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthCurrentPost`: CurrentAuthorization
    fmt.Fprintf(os.Stdout, "Response from `ApiAuthenticationApi.AuthCurrentPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthCurrentPostRequest struct via the builder pattern


### Return type

[**CurrentAuthorization**](CurrentAuthorization.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthGet

> Authorization AuthGet(ctx).Execute()

Get all the Authorization details associated with the current api 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiAuthenticationApi.AuthGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiAuthenticationApi.AuthGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthGet`: Authorization
    fmt.Fprintf(os.Stdout, "Response from `ApiAuthenticationApi.AuthGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthGetRequest struct via the builder pattern


### Return type

[**Authorization**](Authorization.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthInvalidateTokenPost

> AuthInvalidateTokenPost(ctx).Execute()

Invalidate current token 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiAuthenticationApi.AuthInvalidateTokenPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiAuthenticationApi.AuthInvalidateTokenPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthInvalidateTokenPostRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthKeepAlivePost

> AuthToken AuthKeepAlivePost(ctx).Execute()

Invalidate existing token and generates new token 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiAuthenticationApi.AuthKeepAlivePost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiAuthenticationApi.AuthKeepAlivePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthKeepAlivePost`: AuthToken
    fmt.Fprintf(os.Stdout, "Response from `ApiAuthenticationApi.AuthKeepAlivePost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthKeepAlivePostRequest struct via the builder pattern


### Return type

[**AuthToken**](AuthToken.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokensPost

> AuthToken AuthTokensPost(ctx).Execute()

Create a token based on other authentication details (basic, etc.) 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiAuthenticationApi.AuthTokensPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiAuthenticationApi.AuthTokensPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokensPost`: AuthToken
    fmt.Fprintf(os.Stdout, "Response from `ApiAuthenticationApi.AuthTokensPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokensPostRequest struct via the builder pattern


### Return type

[**AuthToken**](AuthToken.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AuthGet

> AuthorizationV1 V1AuthGet(ctx).Execute()

Get all the Authorization details associated with the current api 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiAuthenticationApi.V1AuthGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiAuthenticationApi.V1AuthGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AuthGet`: AuthorizationV1
    fmt.Fprintf(os.Stdout, "Response from `ApiAuthenticationApi.V1AuthGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AuthGetRequest struct via the builder pattern


### Return type

[**AuthorizationV1**](AuthorizationV1.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AuthInvalidateTokenPost

> V1AuthInvalidateTokenPost(ctx).Execute()

Invalidate current token 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiAuthenticationApi.V1AuthInvalidateTokenPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiAuthenticationApi.V1AuthInvalidateTokenPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AuthInvalidateTokenPostRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AuthKeepAlivePost

> AuthTokenV1 V1AuthKeepAlivePost(ctx).Execute()

Invalidate existing token and generates new token 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiAuthenticationApi.V1AuthKeepAlivePost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiAuthenticationApi.V1AuthKeepAlivePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AuthKeepAlivePost`: AuthTokenV1
    fmt.Fprintf(os.Stdout, "Response from `ApiAuthenticationApi.V1AuthKeepAlivePost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AuthKeepAlivePostRequest struct via the builder pattern


### Return type

[**AuthTokenV1**](AuthTokenV1.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1AuthTokenPost

> AuthTokenV1 V1AuthTokenPost(ctx).Execute()

Create a token based on other authentication details (basic, etc.) 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiAuthenticationApi.V1AuthTokenPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiAuthenticationApi.V1AuthTokenPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1AuthTokenPost`: AuthTokenV1
    fmt.Fprintf(os.Stdout, "Response from `ApiAuthenticationApi.V1AuthTokenPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1AuthTokenPostRequest struct via the builder pattern


### Return type

[**AuthTokenV1**](AuthTokenV1.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

