# \MdmApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreviewMdmCommandsPost**](MdmApi.md#PreviewMdmCommandsPost) | **Post** /preview/mdm/commands | Post a command for creation and queuing 
[**V1DeployPackagePost**](MdmApi.md#V1DeployPackagePost) | **Post** /v1/deploy-package | Deploy packages using MDM
[**V1MdmCommandsGet**](MdmApi.md#V1MdmCommandsGet) | **Get** /v1/mdm/commands | Get information about mdm commands made by Jamf Pro. 
[**V1MdmRenewProfilePost**](MdmApi.md#V1MdmRenewProfilePost) | **Post** /v1/mdm/renew-profile | Renew MDM Profile 



## PreviewMdmCommandsPost

> []HrefResponse PreviewMdmCommandsPost(ctx).MdmCommandRequest(mdmCommandRequest).Execute()

Post a command for creation and queuing 



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
    mdmCommandRequest := *openapiclient.NewMdmCommandRequest() // MdmCommandRequest | The mdm command object to create and queue (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MdmApi.PreviewMdmCommandsPost(context.Background()).MdmCommandRequest(mdmCommandRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MdmApi.PreviewMdmCommandsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewMdmCommandsPost`: []HrefResponse
    fmt.Fprintf(os.Stdout, "Response from `MdmApi.PreviewMdmCommandsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewMdmCommandsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mdmCommandRequest** | [**MdmCommandRequest**](MdmCommandRequest.md) | The mdm command object to create and queue | 

### Return type

[**[]HrefResponse**](HrefResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeployPackagePost

> VerbosePackageDeploymentResponse V1DeployPackagePost(ctx).InstallPackage(installPackage).Verbose(verbose).Execute()

Deploy packages using MDM



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
    installPackage := *openapiclient.NewInstallPackage(*openapiclient.NewPackageManifest("https://example.jamf.com/this/package", "dcb02a41cd6d842943459a88c96a5f72", "MD5", "com.jamf.example", "0.1.0", "Title", int32(12345))) // InstallPackage | 
    verbose := true // bool | Enables the 'verbose' response, which includes information about the commands queued as well as information about commands that failed to queue. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MdmApi.V1DeployPackagePost(context.Background()).InstallPackage(installPackage).Verbose(verbose).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MdmApi.V1DeployPackagePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1DeployPackagePost`: VerbosePackageDeploymentResponse
    fmt.Fprintf(os.Stdout, "Response from `MdmApi.V1DeployPackagePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1DeployPackagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **installPackage** | [**InstallPackage**](InstallPackage.md) |  | 
 **verbose** | **bool** | Enables the &#39;verbose&#39; response, which includes information about the commands queued as well as information about commands that failed to queue. | [default to false]

### Return type

[**VerbosePackageDeploymentResponse**](VerbosePackageDeploymentResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MdmCommandsGet

> []MdmCommand V1MdmCommandsGet(ctx).Uuids(uuids).Execute()

Get information about mdm commands made by Jamf Pro. 



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
    uuids := []string{"Inner_example"} // []string | A list of the UUIDs of the commands being searched for.  Limited to 40 UUIDs in length.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MdmApi.V1MdmCommandsGet(context.Background()).Uuids(uuids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MdmApi.V1MdmCommandsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1MdmCommandsGet`: []MdmCommand
    fmt.Fprintf(os.Stdout, "Response from `MdmApi.V1MdmCommandsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MdmCommandsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuids** | **[]string** | A list of the UUIDs of the commands being searched for.  Limited to 40 UUIDs in length. | 

### Return type

[**[]MdmCommand**](MdmCommand.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MdmRenewProfilePost

> RenewMdmProfileResponse V1MdmRenewProfilePost(ctx).Udids(udids).Execute()

Renew MDM Profile 



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
    udids := *openapiclient.NewUdids() // Udids | List of devices' UDIDs to perform MDM profile renewal

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MdmApi.V1MdmRenewProfilePost(context.Background()).Udids(udids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MdmApi.V1MdmRenewProfilePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1MdmRenewProfilePost`: RenewMdmProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `MdmApi.V1MdmRenewProfilePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MdmRenewProfilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **udids** | [**Udids**](Udids.md) | List of devices&#39; UDIDs to perform MDM profile renewal | 

### Return type

[**RenewMdmProfileResponse**](RenewMdmProfileResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

