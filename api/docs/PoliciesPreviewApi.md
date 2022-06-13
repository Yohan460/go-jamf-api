# \PoliciesPreviewApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SettingsObjPolicyPropertiesGet**](PoliciesPreviewApi.md#SettingsObjPolicyPropertiesGet) | **Get** /settings/obj/policyProperties | Get Policy Properties object 
[**SettingsObjPolicyPropertiesPut**](PoliciesPreviewApi.md#SettingsObjPolicyPropertiesPut) | **Put** /settings/obj/policyProperties | Update Policy Properties object 
[**V1PolicyPropertiesGet**](PoliciesPreviewApi.md#V1PolicyPropertiesGet) | **Get** /v1/policy-properties | Get Policy Properties object 
[**V1PolicyPropertiesPut**](PoliciesPreviewApi.md#V1PolicyPropertiesPut) | **Put** /v1/policy-properties | Update Policy Properties object 



## SettingsObjPolicyPropertiesGet

> PolicyProperties SettingsObjPolicyPropertiesGet(ctx).Execute()

Get Policy Properties object 



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
    resp, r, err := apiClient.PoliciesPreviewApi.SettingsObjPolicyPropertiesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesPreviewApi.SettingsObjPolicyPropertiesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SettingsObjPolicyPropertiesGet`: PolicyProperties
    fmt.Fprintf(os.Stdout, "Response from `PoliciesPreviewApi.SettingsObjPolicyPropertiesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsObjPolicyPropertiesGetRequest struct via the builder pattern


### Return type

[**PolicyProperties**](PolicyProperties.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SettingsObjPolicyPropertiesPut

> PolicyProperties SettingsObjPolicyPropertiesPut(ctx).PolicyProperties(policyProperties).Execute()

Update Policy Properties object 



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
    policyProperties := *openapiclient.NewPolicyProperties() // PolicyProperties | Policy Properties object to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesPreviewApi.SettingsObjPolicyPropertiesPut(context.Background()).PolicyProperties(policyProperties).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesPreviewApi.SettingsObjPolicyPropertiesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SettingsObjPolicyPropertiesPut`: PolicyProperties
    fmt.Fprintf(os.Stdout, "Response from `PoliciesPreviewApi.SettingsObjPolicyPropertiesPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSettingsObjPolicyPropertiesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyProperties** | [**PolicyProperties**](PolicyProperties.md) | Policy Properties object to update | 

### Return type

[**PolicyProperties**](PolicyProperties.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PolicyPropertiesGet

> PolicyPropertiesV1 V1PolicyPropertiesGet(ctx).Execute()

Get Policy Properties object 



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
    resp, r, err := apiClient.PoliciesPreviewApi.V1PolicyPropertiesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesPreviewApi.V1PolicyPropertiesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PolicyPropertiesGet`: PolicyPropertiesV1
    fmt.Fprintf(os.Stdout, "Response from `PoliciesPreviewApi.V1PolicyPropertiesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1PolicyPropertiesGetRequest struct via the builder pattern


### Return type

[**PolicyPropertiesV1**](PolicyPropertiesV1.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PolicyPropertiesPut

> PolicyPropertiesV1 V1PolicyPropertiesPut(ctx).PolicyPropertiesV1(policyPropertiesV1).Execute()

Update Policy Properties object 



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
    policyPropertiesV1 := *openapiclient.NewPolicyPropertiesV1() // PolicyPropertiesV1 | Policy Properties object to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesPreviewApi.V1PolicyPropertiesPut(context.Background()).PolicyPropertiesV1(policyPropertiesV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesPreviewApi.V1PolicyPropertiesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PolicyPropertiesPut`: PolicyPropertiesV1
    fmt.Fprintf(os.Stdout, "Response from `PoliciesPreviewApi.V1PolicyPropertiesPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1PolicyPropertiesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyPropertiesV1** | [**PolicyPropertiesV1**](PolicyPropertiesV1.md) | Policy Properties object to update | 

### Return type

[**PolicyPropertiesV1**](PolicyPropertiesV1.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

