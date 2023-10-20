# \EnrollmentCustomizationPreviewAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnrollmentCustomizationIdAllGet**](EnrollmentCustomizationPreviewAPI.md#V1EnrollmentCustomizationIdAllGet) | **Get** /v1/enrollment-customization/{id}/all | Get all Panels for single Enrollment Customization 
[**V1EnrollmentCustomizationIdAllPanelIdDelete**](EnrollmentCustomizationPreviewAPI.md#V1EnrollmentCustomizationIdAllPanelIdDelete) | **Delete** /v1/enrollment-customization/{id}/all/{panel-id} | Delete a single Panel from an Enrollment Customization 
[**V1EnrollmentCustomizationIdAllPanelIdGet**](EnrollmentCustomizationPreviewAPI.md#V1EnrollmentCustomizationIdAllPanelIdGet) | **Get** /v1/enrollment-customization/{id}/all/{panel-id} | Get a single Panel for a single Enrollment Customization 
[**V1EnrollmentCustomizationIdLdapPanelIdDelete**](EnrollmentCustomizationPreviewAPI.md#V1EnrollmentCustomizationIdLdapPanelIdDelete) | **Delete** /v1/enrollment-customization/{id}/ldap/{panel-id} | Delete an LDAP single panel from an Enrollment Customization 
[**V1EnrollmentCustomizationIdLdapPanelIdGet**](EnrollmentCustomizationPreviewAPI.md#V1EnrollmentCustomizationIdLdapPanelIdGet) | **Get** /v1/enrollment-customization/{id}/ldap/{panel-id} | Get a single LDAP panel for a single Enrollment Customization 
[**V1EnrollmentCustomizationIdLdapPanelIdPut**](EnrollmentCustomizationPreviewAPI.md#V1EnrollmentCustomizationIdLdapPanelIdPut) | **Put** /v1/enrollment-customization/{id}/ldap/{panel-id} | Update a single LDAP Panel for a single Enrollment Customization 
[**V1EnrollmentCustomizationIdLdapPost**](EnrollmentCustomizationPreviewAPI.md#V1EnrollmentCustomizationIdLdapPost) | **Post** /v1/enrollment-customization/{id}/ldap | Create an LDAP Panel for a single Enrollment Customization 
[**V1EnrollmentCustomizationIdSsoPanelIdDelete**](EnrollmentCustomizationPreviewAPI.md#V1EnrollmentCustomizationIdSsoPanelIdDelete) | **Delete** /v1/enrollment-customization/{id}/sso/{panel-id} | Delete a single SSO Panel from an Enrollment Customization 
[**V1EnrollmentCustomizationIdSsoPanelIdGet**](EnrollmentCustomizationPreviewAPI.md#V1EnrollmentCustomizationIdSsoPanelIdGet) | **Get** /v1/enrollment-customization/{id}/sso/{panel-id} | Get a single SSO Panel for a single Enrollment Customization 
[**V1EnrollmentCustomizationIdSsoPanelIdPut**](EnrollmentCustomizationPreviewAPI.md#V1EnrollmentCustomizationIdSsoPanelIdPut) | **Put** /v1/enrollment-customization/{id}/sso/{panel-id} | Update a single SSO Panel for a single Enrollment Customization 
[**V1EnrollmentCustomizationIdSsoPost**](EnrollmentCustomizationPreviewAPI.md#V1EnrollmentCustomizationIdSsoPost) | **Post** /v1/enrollment-customization/{id}/sso | Create an SSO Panel for a single Enrollment Customization 
[**V1EnrollmentCustomizationIdTextPanelIdDelete**](EnrollmentCustomizationPreviewAPI.md#V1EnrollmentCustomizationIdTextPanelIdDelete) | **Delete** /v1/enrollment-customization/{id}/text/{panel-id} | Delete a Text single Panel from an Enrollment Customization 
[**V1EnrollmentCustomizationIdTextPanelIdGet**](EnrollmentCustomizationPreviewAPI.md#V1EnrollmentCustomizationIdTextPanelIdGet) | **Get** /v1/enrollment-customization/{id}/text/{panel-id} | Get a single Text Panel for a single Enrollment Customization 
[**V1EnrollmentCustomizationIdTextPanelIdMarkdownGet**](EnrollmentCustomizationPreviewAPI.md#V1EnrollmentCustomizationIdTextPanelIdMarkdownGet) | **Get** /v1/enrollment-customization/{id}/text/{panel-id}/markdown | Get the markdown output of a single Text Panel for a single Enrollment 
[**V1EnrollmentCustomizationIdTextPanelIdPut**](EnrollmentCustomizationPreviewAPI.md#V1EnrollmentCustomizationIdTextPanelIdPut) | **Put** /v1/enrollment-customization/{id}/text/{panel-id} | Update a single Text Panel for a single Enrollment Customization 
[**V1EnrollmentCustomizationIdTextPost**](EnrollmentCustomizationPreviewAPI.md#V1EnrollmentCustomizationIdTextPost) | **Post** /v1/enrollment-customization/{id}/text | Create a Text Panel for a single Enrollment Customization 
[**V1EnrollmentCustomizationParseMarkdownPost**](EnrollmentCustomizationPreviewAPI.md#V1EnrollmentCustomizationParseMarkdownPost) | **Post** /v1/enrollment-customization/parse-markdown | Parse the given string as markdown text and return Html output 



## V1EnrollmentCustomizationIdAllGet

> EnrollmentCustomizationPanelList V1EnrollmentCustomizationIdAllGet(ctx, id).Execute()

Get all Panels for single Enrollment Customization 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    id := int32(56) // int32 | Enrollment Customization identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdAllGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdAllGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentCustomizationIdAllGet`: EnrollmentCustomizationPanelList
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdAllGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Enrollment Customization identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdAllGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnrollmentCustomizationPanelList**](EnrollmentCustomizationPanelList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentCustomizationIdAllPanelIdDelete

> V1EnrollmentCustomizationIdAllPanelIdDelete(ctx, id, panelId).Execute()

Delete a single Panel from an Enrollment Customization 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    id := int32(56) // int32 | Enrollment Customization identifier
    panelId := int32(56) // int32 | Panel object identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdAllPanelIdDelete(context.Background(), id, panelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdAllPanelIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Enrollment Customization identifier | 
**panelId** | **int32** | Panel object identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdAllPanelIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## V1EnrollmentCustomizationIdAllPanelIdGet

> GetEnrollmentCustomizationPanel V1EnrollmentCustomizationIdAllPanelIdGet(ctx, id, panelId).Execute()

Get a single Panel for a single Enrollment Customization 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    id := int32(56) // int32 | Enrollment Customization identifier
    panelId := int32(56) // int32 | Panel object identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdAllPanelIdGet(context.Background(), id, panelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdAllPanelIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentCustomizationIdAllPanelIdGet`: GetEnrollmentCustomizationPanel
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdAllPanelIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Enrollment Customization identifier | 
**panelId** | **int32** | Panel object identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdAllPanelIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetEnrollmentCustomizationPanel**](GetEnrollmentCustomizationPanel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentCustomizationIdLdapPanelIdDelete

> V1EnrollmentCustomizationIdLdapPanelIdDelete(ctx, id, panelId).Execute()

Delete an LDAP single panel from an Enrollment Customization 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    id := int32(56) // int32 | Enrollment Customization identifier
    panelId := int32(56) // int32 | Panel object identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdLdapPanelIdDelete(context.Background(), id, panelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdLdapPanelIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Enrollment Customization identifier | 
**panelId** | **int32** | Panel object identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdLdapPanelIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## V1EnrollmentCustomizationIdLdapPanelIdGet

> GetEnrollmentCustomizationPanelLdapAuth V1EnrollmentCustomizationIdLdapPanelIdGet(ctx, id, panelId).Execute()

Get a single LDAP panel for a single Enrollment Customization 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    id := int32(56) // int32 | Enrollment Customization identifier
    panelId := int32(56) // int32 | Panel object identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdLdapPanelIdGet(context.Background(), id, panelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdLdapPanelIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentCustomizationIdLdapPanelIdGet`: GetEnrollmentCustomizationPanelLdapAuth
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdLdapPanelIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Enrollment Customization identifier | 
**panelId** | **int32** | Panel object identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdLdapPanelIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetEnrollmentCustomizationPanelLdapAuth**](GetEnrollmentCustomizationPanelLdapAuth.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentCustomizationIdLdapPanelIdPut

> GetEnrollmentCustomizationPanelLdapAuth V1EnrollmentCustomizationIdLdapPanelIdPut(ctx, id, panelId).EnrollmentCustomizationPanelLdapAuth(enrollmentCustomizationPanelLdapAuth).Execute()

Update a single LDAP Panel for a single Enrollment Customization 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    id := int32(56) // int32 | Enrollment Customization identifier
    panelId := int32(56) // int32 | Panel object identifier
    enrollmentCustomizationPanelLdapAuth := *openapiclient.NewEnrollmentCustomizationPanelLdapAuth("A Panel", int32(0), "Username", "Password", "My Ldap Panel", "Back", "Continue") // EnrollmentCustomizationPanelLdapAuth | Enrollment Customization Panel to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdLdapPanelIdPut(context.Background(), id, panelId).EnrollmentCustomizationPanelLdapAuth(enrollmentCustomizationPanelLdapAuth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdLdapPanelIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentCustomizationIdLdapPanelIdPut`: GetEnrollmentCustomizationPanelLdapAuth
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdLdapPanelIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Enrollment Customization identifier | 
**panelId** | **int32** | Panel object identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdLdapPanelIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enrollmentCustomizationPanelLdapAuth** | [**EnrollmentCustomizationPanelLdapAuth**](EnrollmentCustomizationPanelLdapAuth.md) | Enrollment Customization Panel to update | 

### Return type

[**GetEnrollmentCustomizationPanelLdapAuth**](GetEnrollmentCustomizationPanelLdapAuth.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentCustomizationIdLdapPost

> GetEnrollmentCustomizationPanelLdapAuth V1EnrollmentCustomizationIdLdapPost(ctx, id).EnrollmentCustomizationPanelLdapAuth(enrollmentCustomizationPanelLdapAuth).Execute()

Create an LDAP Panel for a single Enrollment Customization 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    id := int32(56) // int32 | Enrollment Customization identifier
    enrollmentCustomizationPanelLdapAuth := *openapiclient.NewEnrollmentCustomizationPanelLdapAuth("A Panel", int32(0), "Username", "Password", "My Ldap Panel", "Back", "Continue") // EnrollmentCustomizationPanelLdapAuth | Enrollment Customization Panel to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdLdapPost(context.Background(), id).EnrollmentCustomizationPanelLdapAuth(enrollmentCustomizationPanelLdapAuth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdLdapPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentCustomizationIdLdapPost`: GetEnrollmentCustomizationPanelLdapAuth
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdLdapPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Enrollment Customization identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdLdapPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enrollmentCustomizationPanelLdapAuth** | [**EnrollmentCustomizationPanelLdapAuth**](EnrollmentCustomizationPanelLdapAuth.md) | Enrollment Customization Panel to create | 

### Return type

[**GetEnrollmentCustomizationPanelLdapAuth**](GetEnrollmentCustomizationPanelLdapAuth.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentCustomizationIdSsoPanelIdDelete

> V1EnrollmentCustomizationIdSsoPanelIdDelete(ctx, id, panelId).Execute()

Delete a single SSO Panel from an Enrollment Customization 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    id := int32(56) // int32 | Enrollment Customization identifier
    panelId := int32(56) // int32 | Panel object identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdSsoPanelIdDelete(context.Background(), id, panelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdSsoPanelIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Enrollment Customization identifier | 
**panelId** | **int32** | Panel object identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdSsoPanelIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## V1EnrollmentCustomizationIdSsoPanelIdGet

> GetEnrollmentCustomizationPanelSsoAuth V1EnrollmentCustomizationIdSsoPanelIdGet(ctx, id, panelId).Execute()

Get a single SSO Panel for a single Enrollment Customization 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    id := int32(56) // int32 | Enrollment Customization identifier
    panelId := int32(56) // int32 | Panel object identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdSsoPanelIdGet(context.Background(), id, panelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdSsoPanelIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentCustomizationIdSsoPanelIdGet`: GetEnrollmentCustomizationPanelSsoAuth
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdSsoPanelIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Enrollment Customization identifier | 
**panelId** | **int32** | Panel object identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdSsoPanelIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetEnrollmentCustomizationPanelSsoAuth**](GetEnrollmentCustomizationPanelSsoAuth.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentCustomizationIdSsoPanelIdPut

> GetEnrollmentCustomizationPanelSsoAuth V1EnrollmentCustomizationIdSsoPanelIdPut(ctx, id, panelId).EnrollmentCustomizationPanelSsoAuth(enrollmentCustomizationPanelSsoAuth).Execute()

Update a single SSO Panel for a single Enrollment Customization 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    id := int32(56) // int32 | Enrollment Customization identifier
    panelId := int32(56) // int32 | Panel object identifier
    enrollmentCustomizationPanelSsoAuth := *openapiclient.NewEnrollmentCustomizationPanelSsoAuth("A Panel", int32(0), false, "long name", "name", false, "GroupNameA") // EnrollmentCustomizationPanelSsoAuth | Enrollment Customization Panel to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdSsoPanelIdPut(context.Background(), id, panelId).EnrollmentCustomizationPanelSsoAuth(enrollmentCustomizationPanelSsoAuth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdSsoPanelIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentCustomizationIdSsoPanelIdPut`: GetEnrollmentCustomizationPanelSsoAuth
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdSsoPanelIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Enrollment Customization identifier | 
**panelId** | **int32** | Panel object identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdSsoPanelIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enrollmentCustomizationPanelSsoAuth** | [**EnrollmentCustomizationPanelSsoAuth**](EnrollmentCustomizationPanelSsoAuth.md) | Enrollment Customization Panel to update | 

### Return type

[**GetEnrollmentCustomizationPanelSsoAuth**](GetEnrollmentCustomizationPanelSsoAuth.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentCustomizationIdSsoPost

> GetEnrollmentCustomizationPanelSsoAuth V1EnrollmentCustomizationIdSsoPost(ctx, id).EnrollmentCustomizationPanelSsoAuth(enrollmentCustomizationPanelSsoAuth).Execute()

Create an SSO Panel for a single Enrollment Customization 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    id := int32(56) // int32 | Enrollment Customization identifier
    enrollmentCustomizationPanelSsoAuth := *openapiclient.NewEnrollmentCustomizationPanelSsoAuth("A Panel", int32(0), false, "long name", "name", false, "GroupNameA") // EnrollmentCustomizationPanelSsoAuth | Enrollment Customization Panel to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdSsoPost(context.Background(), id).EnrollmentCustomizationPanelSsoAuth(enrollmentCustomizationPanelSsoAuth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdSsoPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentCustomizationIdSsoPost`: GetEnrollmentCustomizationPanelSsoAuth
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdSsoPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Enrollment Customization identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdSsoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enrollmentCustomizationPanelSsoAuth** | [**EnrollmentCustomizationPanelSsoAuth**](EnrollmentCustomizationPanelSsoAuth.md) | Enrollment Customization Panel to create | 

### Return type

[**GetEnrollmentCustomizationPanelSsoAuth**](GetEnrollmentCustomizationPanelSsoAuth.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentCustomizationIdTextPanelIdDelete

> V1EnrollmentCustomizationIdTextPanelIdDelete(ctx, id, panelId).Execute()

Delete a Text single Panel from an Enrollment Customization 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    id := int32(56) // int32 | Enrollment Customization identifier
    panelId := int32(56) // int32 | Panel object identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdTextPanelIdDelete(context.Background(), id, panelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdTextPanelIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Enrollment Customization identifier | 
**panelId** | **int32** | Panel object identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdTextPanelIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## V1EnrollmentCustomizationIdTextPanelIdGet

> GetEnrollmentCustomizationPanelText V1EnrollmentCustomizationIdTextPanelIdGet(ctx, id, panelId).Execute()

Get a single Text Panel for a single Enrollment Customization 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    id := int32(56) // int32 | Enrollment Customization identifier
    panelId := int32(56) // int32 | Panel object identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdTextPanelIdGet(context.Background(), id, panelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdTextPanelIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentCustomizationIdTextPanelIdGet`: GetEnrollmentCustomizationPanelText
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdTextPanelIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Enrollment Customization identifier | 
**panelId** | **int32** | Panel object identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdTextPanelIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetEnrollmentCustomizationPanelText**](GetEnrollmentCustomizationPanelText.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentCustomizationIdTextPanelIdMarkdownGet

> Markdown V1EnrollmentCustomizationIdTextPanelIdMarkdownGet(ctx, id, panelId).Execute()

Get the markdown output of a single Text Panel for a single Enrollment 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    id := int32(56) // int32 | Enrollment Customization identifier
    panelId := int32(56) // int32 | Panel object identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdTextPanelIdMarkdownGet(context.Background(), id, panelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdTextPanelIdMarkdownGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentCustomizationIdTextPanelIdMarkdownGet`: Markdown
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdTextPanelIdMarkdownGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Enrollment Customization identifier | 
**panelId** | **int32** | Panel object identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdTextPanelIdMarkdownGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Markdown**](Markdown.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentCustomizationIdTextPanelIdPut

> GetEnrollmentCustomizationPanelText V1EnrollmentCustomizationIdTextPanelIdPut(ctx, id, panelId).EnrollmentCustomizationPanelText(enrollmentCustomizationPanelText).Execute()

Update a single Text Panel for a single Enrollment Customization 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    id := int32(56) // int32 | Enrollment Customization identifier
    panelId := int32(56) // int32 | Panel object identifier
    enrollmentCustomizationPanelText := *openapiclient.NewEnrollmentCustomizationPanelText("A Panel", int32(0), "Welcome!", "My text panel", "Back", "Continue") // EnrollmentCustomizationPanelText | Enrollment Customization Panel to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdTextPanelIdPut(context.Background(), id, panelId).EnrollmentCustomizationPanelText(enrollmentCustomizationPanelText).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdTextPanelIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentCustomizationIdTextPanelIdPut`: GetEnrollmentCustomizationPanelText
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdTextPanelIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Enrollment Customization identifier | 
**panelId** | **int32** | Panel object identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdTextPanelIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enrollmentCustomizationPanelText** | [**EnrollmentCustomizationPanelText**](EnrollmentCustomizationPanelText.md) | Enrollment Customization Panel to update | 

### Return type

[**GetEnrollmentCustomizationPanelText**](GetEnrollmentCustomizationPanelText.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentCustomizationIdTextPost

> GetEnrollmentCustomizationPanelText V1EnrollmentCustomizationIdTextPost(ctx, id).EnrollmentCustomizationPanelText(enrollmentCustomizationPanelText).Execute()

Create a Text Panel for a single Enrollment Customization 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    id := int32(56) // int32 | Enrollment Customization identifier
    enrollmentCustomizationPanelText := *openapiclient.NewEnrollmentCustomizationPanelText("A Panel", int32(0), "Welcome!", "My text panel", "Back", "Continue") // EnrollmentCustomizationPanelText | Enrollment Customization Panel to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdTextPost(context.Background(), id).EnrollmentCustomizationPanelText(enrollmentCustomizationPanelText).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdTextPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentCustomizationIdTextPost`: GetEnrollmentCustomizationPanelText
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationIdTextPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Enrollment Customization identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationIdTextPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enrollmentCustomizationPanelText** | [**EnrollmentCustomizationPanelText**](EnrollmentCustomizationPanelText.md) | Enrollment Customization Panel to create | 

### Return type

[**GetEnrollmentCustomizationPanelText**](GetEnrollmentCustomizationPanelText.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnrollmentCustomizationParseMarkdownPost

> Markdown V1EnrollmentCustomizationParseMarkdownPost(ctx).Markdown(markdown).Execute()

Parse the given string as markdown text and return Html output 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yohan460/go-jamf-api"
)

func main() {
    markdown := *openapiclient.NewMarkdown() // Markdown | Enrollment Customization Panel to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationParseMarkdownPost(context.Background()).Markdown(markdown).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationParseMarkdownPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrollmentCustomizationParseMarkdownPost`: Markdown
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentCustomizationPreviewAPI.V1EnrollmentCustomizationParseMarkdownPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollmentCustomizationParseMarkdownPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **markdown** | [**Markdown**](Markdown.md) | Enrollment Customization Panel to create | 

### Return type

[**Markdown**](Markdown.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

