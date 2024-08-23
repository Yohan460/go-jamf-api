# \SmtpServerAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SmtpServerGet**](SmtpServerAPI.md#V1SmtpServerGet) | **Get** /v1/smtp-server | Finds the Jamf Pro SMTP Server information 
[**V1SmtpServerHistoryGet**](SmtpServerAPI.md#V1SmtpServerHistoryGet) | **Get** /v1/smtp-server/history | Get specified SMTP Server history object 
[**V1SmtpServerHistoryPost**](SmtpServerAPI.md#V1SmtpServerHistoryPost) | **Post** /v1/smtp-server/history | Add SMTP Server history object notes 
[**V1SmtpServerPut**](SmtpServerAPI.md#V1SmtpServerPut) | **Put** /v1/smtp-server | Updates Jamf Pro SMTP Server information 
[**V1SmtpServerTestPost**](SmtpServerAPI.md#V1SmtpServerTestPost) | **Post** /v1/smtp-server/test | Test functionality of an SMTP Server
[**V2SmtpServerGet**](SmtpServerAPI.md#V2SmtpServerGet) | **Get** /v2/smtp-server | Finds the Jamf Pro SMTP Server information 
[**V2SmtpServerPut**](SmtpServerAPI.md#V2SmtpServerPut) | **Put** /v2/smtp-server | Updates Jamf Pro SMTP Server information 



## V1SmtpServerGet

> SmtpServer V1SmtpServerGet(ctx).Execute()

Finds the Jamf Pro SMTP Server information 



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
	resp, r, err := apiClient.SmtpServerAPI.V1SmtpServerGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmtpServerAPI.V1SmtpServerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SmtpServerGet`: SmtpServer
	fmt.Fprintf(os.Stdout, "Response from `SmtpServerAPI.V1SmtpServerGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1SmtpServerGetRequest struct via the builder pattern


### Return type

[**SmtpServer**](SmtpServer.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SmtpServerHistoryGet

> HistorySearchResultsV1 V1SmtpServerHistoryGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get specified SMTP Server history object 



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,username:asc  (optional) (default to ["date:desc"])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmtpServerAPI.V1SmtpServerHistoryGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmtpServerAPI.V1SmtpServerHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SmtpServerHistoryGet`: HistorySearchResultsV1
	fmt.Fprintf(os.Stdout, "Response from `SmtpServerAPI.V1SmtpServerHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SmtpServerHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,username:asc  | [default to [&quot;date:desc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15 | [default to &quot;&quot;]

### Return type

[**HistorySearchResultsV1**](HistorySearchResultsV1.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SmtpServerHistoryPost

> HrefResponse V1SmtpServerHistoryPost(ctx).ObjectHistoryNote(objectHistoryNote).Execute()

Add SMTP Server history object notes 



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
	objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | History notes to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmtpServerAPI.V1SmtpServerHistoryPost(context.Background()).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmtpServerAPI.V1SmtpServerHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SmtpServerHistoryPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `SmtpServerAPI.V1SmtpServerHistoryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SmtpServerHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectHistoryNote** | [**ObjectHistoryNote**](ObjectHistoryNote.md) | History notes to create | 

### Return type

[**HrefResponse**](HrefResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SmtpServerPut

> SmtpServer V1SmtpServerPut(ctx).SmtpServer(smtpServer).Execute()

Updates Jamf Pro SMTP Server information 



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
	smtpServer := *openapiclient.NewSmtpServer(true, "abcd.server.com", int64(25), "SSL", int64(5), "Jamf Pro Server", "exampleEmail@example.com", true) // SmtpServer | SMTP Server to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmtpServerAPI.V1SmtpServerPut(context.Background()).SmtpServer(smtpServer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmtpServerAPI.V1SmtpServerPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SmtpServerPut`: SmtpServer
	fmt.Fprintf(os.Stdout, "Response from `SmtpServerAPI.V1SmtpServerPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SmtpServerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smtpServer** | [**SmtpServer**](SmtpServer.md) | SMTP Server to update | 

### Return type

[**SmtpServer**](SmtpServer.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SmtpServerTestPost

> V1SmtpServerTestPost(ctx).SmtpServerTest(smtpServerTest).Execute()

Test functionality of an SMTP Server



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
	smtpServerTest := *openapiclient.NewSmtpServerTest("exampleEmail@example.com") // SmtpServerTest | Recipient email to test SMTP Server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SmtpServerAPI.V1SmtpServerTestPost(context.Background()).SmtpServerTest(smtpServerTest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmtpServerAPI.V1SmtpServerTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SmtpServerTestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smtpServerTest** | [**SmtpServerTest**](SmtpServerTest.md) | Recipient email to test SMTP Server | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SmtpServerGet

> SmtpServerV2 V2SmtpServerGet(ctx).Execute()

Finds the Jamf Pro SMTP Server information 



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
	resp, r, err := apiClient.SmtpServerAPI.V2SmtpServerGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmtpServerAPI.V2SmtpServerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2SmtpServerGet`: SmtpServerV2
	fmt.Fprintf(os.Stdout, "Response from `SmtpServerAPI.V2SmtpServerGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2SmtpServerGetRequest struct via the builder pattern


### Return type

[**SmtpServerV2**](SmtpServerV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2SmtpServerPut

> SmtpServerV2 V2SmtpServerPut(ctx).SmtpServerV2(smtpServerV2).Execute()

Updates Jamf Pro SMTP Server information 



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
	smtpServerV2 := *openapiclient.NewSmtpServerV2(true, "NONE", *openapiclient.NewSmtpSenderSettings("exampleEmail@example.com")) // SmtpServerV2 | SMTP Server to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmtpServerAPI.V2SmtpServerPut(context.Background()).SmtpServerV2(smtpServerV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmtpServerAPI.V2SmtpServerPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V2SmtpServerPut`: SmtpServerV2
	fmt.Fprintf(os.Stdout, "Response from `SmtpServerAPI.V2SmtpServerPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2SmtpServerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smtpServerV2** | [**SmtpServerV2**](SmtpServerV2.md) | SMTP Server to update | 

### Return type

[**SmtpServerV2**](SmtpServerV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

