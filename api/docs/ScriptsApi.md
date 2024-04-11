# \ScriptsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ScriptsGet**](ScriptsAPI.md#V1ScriptsGet) | **Get** /v1/scripts | Search for sorted and paged Scripts 
[**V1ScriptsIdDelete**](ScriptsAPI.md#V1ScriptsIdDelete) | **Delete** /v1/scripts/{id} | Delete a Script at the specified id 
[**V1ScriptsIdDownloadGet**](ScriptsAPI.md#V1ScriptsIdDownloadGet) | **Get** /v1/scripts/{id}/download | Download a text file of the Script contents 
[**V1ScriptsIdGet**](ScriptsAPI.md#V1ScriptsIdGet) | **Get** /v1/scripts/{id} | Retrieve a full script object 
[**V1ScriptsIdHistoryGet**](ScriptsAPI.md#V1ScriptsIdHistoryGet) | **Get** /v1/scripts/{id}/history | Get specified Script history object 
[**V1ScriptsIdHistoryPost**](ScriptsAPI.md#V1ScriptsIdHistoryPost) | **Post** /v1/scripts/{id}/history | Add specified Script history object notes 
[**V1ScriptsIdPut**](ScriptsAPI.md#V1ScriptsIdPut) | **Put** /v1/scripts/{id} | Replace the script at the id with the supplied information 
[**V1ScriptsPost**](ScriptsAPI.md#V1ScriptsPost) | **Post** /v1/scripts | Create a Script 



## V1ScriptsGet

> ScriptsSearchResults V1ScriptsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Search for sorted and paged Scripts 



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
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Fields allowed in the query: `id`, `name`, `info`, `notes`, `priority`, `categoryId`, `categoryName`, `parameter4` up to `parameter11`, `osRequirements`, `scriptContents`. Example: sort=date:desc,name:asc  (optional) (default to ["name:asc"])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter scripts collection. Default search is empty query - returning all results for the requested page. Fields allowed in the query: `id`, `name`, `info`, `notes`, `priority`, `categoryId`, `categoryName`, `parameter4` up to `parameter11`, `osRequirements`, `scriptContents`. This param can be combined with paging and sorting. Example: filter=categoryName==\"Category\" and name==\"*script name*\" (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScriptsAPI.V1ScriptsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScriptsAPI.V1ScriptsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ScriptsGet`: ScriptsSearchResults
	fmt.Fprintf(os.Stdout, "Response from `ScriptsAPI.V1ScriptsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ScriptsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Fields allowed in the query: &#x60;id&#x60;, &#x60;name&#x60;, &#x60;info&#x60;, &#x60;notes&#x60;, &#x60;priority&#x60;, &#x60;categoryId&#x60;, &#x60;categoryName&#x60;, &#x60;parameter4&#x60; up to &#x60;parameter11&#x60;, &#x60;osRequirements&#x60;, &#x60;scriptContents&#x60;. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;name:asc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter scripts collection. Default search is empty query - returning all results for the requested page. Fields allowed in the query: &#x60;id&#x60;, &#x60;name&#x60;, &#x60;info&#x60;, &#x60;notes&#x60;, &#x60;priority&#x60;, &#x60;categoryId&#x60;, &#x60;categoryName&#x60;, &#x60;parameter4&#x60; up to &#x60;parameter11&#x60;, &#x60;osRequirements&#x60;, &#x60;scriptContents&#x60;. This param can be combined with paging and sorting. Example: filter&#x3D;categoryName&#x3D;&#x3D;\&quot;Category\&quot; and name&#x3D;&#x3D;\&quot;*script name*\&quot; | [default to &quot;&quot;]

### Return type

[**ScriptsSearchResults**](ScriptsSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ScriptsIdDelete

> V1ScriptsIdDelete(ctx, id).Execute()

Delete a Script at the specified id 



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
	id := "id_example" // string | Script object identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScriptsAPI.V1ScriptsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScriptsAPI.V1ScriptsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Script object identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ScriptsIdDeleteRequest struct via the builder pattern


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


## V1ScriptsIdDownloadGet

> *os.File V1ScriptsIdDownloadGet(ctx, id).Execute()

Download a text file of the Script contents 



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
	id := "id_example" // string | id of the script to be downloaded

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScriptsAPI.V1ScriptsIdDownloadGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScriptsAPI.V1ScriptsIdDownloadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ScriptsIdDownloadGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ScriptsAPI.V1ScriptsIdDownloadGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of the script to be downloaded | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ScriptsIdDownloadGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ScriptsIdGet

> Script V1ScriptsIdGet(ctx, id).Execute()

Retrieve a full script object 



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
	id := "id_example" // string | Script object identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScriptsAPI.V1ScriptsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScriptsAPI.V1ScriptsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ScriptsIdGet`: Script
	fmt.Fprintf(os.Stdout, "Response from `ScriptsAPI.V1ScriptsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Script object identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ScriptsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Script**](Script.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ScriptsIdHistoryGet

> HistorySearchResults V1ScriptsIdHistoryGet(ctx, id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()

Get specified Script history object 



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
	id := "id_example" // string | id of script history record
	page := int64(56) // int64 |  (optional) (default to 0)
	pageSize := int64(56) // int64 |  (optional) (default to 100)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc  (optional) (default to ["date:desc"])
	filter := "filter_example" // string | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==*disabled* and date<2019-12-15 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScriptsAPI.V1ScriptsIdHistoryGet(context.Background(), id).Page(page).PageSize(pageSize).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScriptsAPI.V1ScriptsIdHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ScriptsIdHistoryGet`: HistorySearchResults
	fmt.Fprintf(os.Stdout, "Response from `ScriptsAPI.V1ScriptsIdHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of script history record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ScriptsIdHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** |  | [default to 0]
 **pageSize** | **int64** |  | [default to 100]
 **sort** | **[]string** | Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc  | [default to [&quot;date:desc&quot;]]
 **filter** | **string** | Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15 | [default to &quot;&quot;]

### Return type

[**HistorySearchResults**](HistorySearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ScriptsIdHistoryPost

> ObjectHistory V1ScriptsIdHistoryPost(ctx, id).ObjectHistoryNote(objectHistoryNote).Execute()

Add specified Script history object notes 



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
	id := "id_example" // string | instance id of script history record
	objectHistoryNote := *openapiclient.NewObjectHistoryNote("A generic note can sometimes be useful, but generally not.") // ObjectHistoryNote | history notes to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScriptsAPI.V1ScriptsIdHistoryPost(context.Background(), id).ObjectHistoryNote(objectHistoryNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScriptsAPI.V1ScriptsIdHistoryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ScriptsIdHistoryPost`: ObjectHistory
	fmt.Fprintf(os.Stdout, "Response from `ScriptsAPI.V1ScriptsIdHistoryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | instance id of script history record | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ScriptsIdHistoryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectHistoryNote** | [**ObjectHistoryNote**](ObjectHistoryNote.md) | history notes to create | 

### Return type

[**ObjectHistory**](ObjectHistory.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ScriptsIdPut

> Script V1ScriptsIdPut(ctx, id).Script(script).Execute()

Replace the script at the id with the supplied information 



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
	id := "id_example" // string | Script object identifier
	script := *openapiclient.NewScript("Install Developer Utils Script") // Script | new script to upload to existing id. ids defined in this body will be ignored

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScriptsAPI.V1ScriptsIdPut(context.Background(), id).Script(script).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScriptsAPI.V1ScriptsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ScriptsIdPut`: Script
	fmt.Fprintf(os.Stdout, "Response from `ScriptsAPI.V1ScriptsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Script object identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ScriptsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **script** | [**Script**](Script.md) | new script to upload to existing id. ids defined in this body will be ignored | 

### Return type

[**Script**](Script.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ScriptsPost

> HrefResponse V1ScriptsPost(ctx).Script(script).Execute()

Create a Script 



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
	script := *openapiclient.NewScript("Install Developer Utils Script") // Script | new script to create. ids defined in this body will be ignored

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScriptsAPI.V1ScriptsPost(context.Background()).Script(script).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScriptsAPI.V1ScriptsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ScriptsPost`: HrefResponse
	fmt.Fprintf(os.Stdout, "Response from `ScriptsAPI.V1ScriptsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ScriptsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **script** | [**Script**](Script.md) | new script to create. ids defined in this body will be ignored | 

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

