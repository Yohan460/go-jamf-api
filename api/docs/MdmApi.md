# \MdmAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreviewMdmCommandsPost**](MdmAPI.md#PreviewMdmCommandsPost) | **Post** /preview/mdm/commands | Post a command for creation and queuing 
[**V1DeployPackagePost**](MdmAPI.md#V1DeployPackagePost) | **Post** /v1/deploy-package | Deploy packages using MDM
[**V1MdmCommandsGet**](MdmAPI.md#V1MdmCommandsGet) | **Get** /v1/mdm/commands | Get information about mdm commands made by Jamf Pro.
[**V1MdmRenewProfilePost**](MdmAPI.md#V1MdmRenewProfilePost) | **Post** /v1/mdm/renew-profile | Renew MDM Profile 
[**V2MdmCommandsGet**](MdmAPI.md#V2MdmCommandsGet) | **Get** /v2/mdm/commands | Get information about mdm commands made by Jamf Pro. 
[**V2MdmCommandsPost**](MdmAPI.md#V2MdmCommandsPost) | **Post** /v2/mdm/commands | Post a command for creation and queuing 



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
	openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {
	mdmCommandRequest := *openapiclient.NewMdmCommandRequest() // MdmCommandRequest | The mdm command object to create and queue (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MdmAPI.PreviewMdmCommandsPost(context.Background()).MdmCommandRequest(mdmCommandRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MdmAPI.PreviewMdmCommandsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreviewMdmCommandsPost`: []HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `MdmAPI.PreviewMdmCommandsPost`: %v\n", resp)
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
	openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {
	installPackage := *openapiclient.NewInstallPackage(*openapiclient.NewPackageManifest("https://example.jamf.com/this/package", "dcb02a41cd6d842943459a88c96a5f72", "MD5", "com.jamf.example", "0.1.0", "Title", int64(12345))) // InstallPackage | 
	verbose := true // bool | Enables the 'verbose' response, which includes information about the commands queued as well as information about commands that failed to queue. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MdmAPI.V1DeployPackagePost(context.Background()).InstallPackage(installPackage).Verbose(verbose).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MdmAPI.V1DeployPackagePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1DeployPackagePost`: VerbosePackageDeploymentResponse
	fmt.Fprintf(os.Stdout, "Response from `MdmAPI.V1DeployPackagePost`: %v\n", resp)
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

> []MdmCommand V1MdmCommandsGet(ctx).Uuids(uuids).ClientManagementId(clientManagementId).Execute()

Get information about mdm commands made by Jamf Pro.



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
	uuids := []string{"Inner_example"} // []string | A list of the UUIDs of the commands being searched for.  Limited to 40 UUIDs in length. Choose one of two parameters, but not both. (optional)
	clientManagementId := "fd68c371-5921-436e-b16b-8a3c1bf90ee5" // string | The client management id used to search for a list of commands. Choose one of two parameters, but not both. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MdmAPI.V1MdmCommandsGet(context.Background()).Uuids(uuids).ClientManagementId(clientManagementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MdmAPI.V1MdmCommandsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MdmCommandsGet`: []MdmCommand
	fmt.Fprintf(os.Stdout, "Response from `MdmAPI.V1MdmCommandsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MdmCommandsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuids** | **[]string** | A list of the UUIDs of the commands being searched for.  Limited to 40 UUIDs in length. Choose one of two parameters, but not both. | 
 **clientManagementId** | **string** | The client management id used to search for a list of commands. Choose one of two parameters, but not both. | 

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
	openapiclient "github.com/yohan460/go-jamf-api/api"
)

func main() {
	udids := *openapiclient.NewUdids() // Udids | List of devices' UDIDs to perform MDM profile renewal

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MdmAPI.V1MdmRenewProfilePost(context.Background()).Udids(udids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MdmAPI.V1MdmRenewProfilePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MdmRenewProfilePost`: RenewMdmProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `MdmAPI.V1MdmRenewProfilePost`: %v\n", resp)
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


## V2MdmCommandsGet

> MdmCommandResults V2MdmCommandsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get information about mdm commands made by Jamf Pro. 



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
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Default sort is dateSent:asc. Multiple sort criteria are supported and must be separated with a comma. (optional) (default to ["dateSent:asc"])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter, for a list of commands. All url must contain minimum one filter field. Fields allowed in the query: uuid, clientManagementId, command, status, clientType, dateSent, validAfter, dateCompleted, profileIdentifier, and active. This param can be combined with paging. Please note that any date filters must be used with gt, lt, ge, le Example: clientManagementId==fb511aae-c557-474f-a9c1-5dc845b90d0f;status==Pending;command==INSTALL_PROFILE;uuid==9e18f849-e689-4f2d-b616-a99d3da7db42;clientType==COMPUTER_USER;profileIdentifier==18cc61c2-01fc-11ed-b939-0242ac120002;dateCompleted=ge=2021-08-04T14:25:18.26Z;dateCompleted=le=2021-08-04T14:25:18.26Z;validAfter=ge=2021-08-05T14:25:18.26Z;active==true (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MdmAPI.V2MdmCommandsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MdmAPI.V2MdmCommandsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2MdmCommandsGet`: MdmCommandResults
	fmt.Fprintf(os.Stdout, "Response from `MdmAPI.V2MdmCommandsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2MdmCommandsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Default sort is dateSent:asc. Multiple sort criteria are supported and must be separated with a comma. | [default to [&quot;dateSent:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter, for a list of commands. All url must contain minimum one filter field. Fields allowed in the query: uuid, clientManagementId, command, status, clientType, dateSent, validAfter, dateCompleted, profileIdentifier, and active. This param can be combined with paging. Please note that any date filters must be used with gt, lt, ge, le Example: clientManagementId&#x3D;&#x3D;fb511aae-c557-474f-a9c1-5dc845b90d0f;status&#x3D;&#x3D;Pending;command&#x3D;&#x3D;INSTALL_PROFILE;uuid&#x3D;&#x3D;9e18f849-e689-4f2d-b616-a99d3da7db42;clientType&#x3D;&#x3D;COMPUTER_USER;profileIdentifier&#x3D;&#x3D;18cc61c2-01fc-11ed-b939-0242ac120002;dateCompleted&#x3D;ge&#x3D;2021-08-04T14:25:18.26Z;dateCompleted&#x3D;le&#x3D;2021-08-04T14:25:18.26Z;validAfter&#x3D;ge&#x3D;2021-08-05T14:25:18.26Z;active&#x3D;&#x3D;true | [default to &quot;&quot;]

### Return type

[**MdmCommandResults**](MdmCommandResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2MdmCommandsPost

> []HrefResponse V2MdmCommandsPost(ctx).MdmCommandRequest(mdmCommandRequest).Execute()

Post a command for creation and queuing 



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
	mdmCommandRequest := *openapiclient.NewMdmCommandRequest() // MdmCommandRequest | The mdm command object to create and queue (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MdmAPI.V2MdmCommandsPost(context.Background()).MdmCommandRequest(mdmCommandRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MdmAPI.V2MdmCommandsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2MdmCommandsPost`: []HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `MdmAPI.V2MdmCommandsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2MdmCommandsPostRequest struct via the builder pattern


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

