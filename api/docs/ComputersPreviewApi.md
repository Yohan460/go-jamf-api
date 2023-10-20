# \ComputersPreviewAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreviewComputersGet**](ComputersPreviewAPI.md#PreviewComputersGet) | **Get** /preview/computers | Return a list of Computers 



## PreviewComputersGet

> ComputersSearchResults PreviewComputersGet(ctx).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()

Return a list of Computers 



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
    page := int32(56) // int32 |  (optional) (default to 0)
    size := int32(56) // int32 |  (optional) (default to 100)
    pagesize := int32(56) // int32 |  (optional) (default to 100)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    sort := "sort_example" // string | Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc (optional) (default to "name:asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComputersPreviewAPI.PreviewComputersGet(context.Background()).Page(page).Size(size).Pagesize(pagesize).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputersPreviewAPI.PreviewComputersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewComputersGet`: ComputersSearchResults
    fmt.Fprintf(os.Stdout, "Response from `ComputersPreviewAPI.PreviewComputersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewComputersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 100]
 **pagesize** | **int32** |  | [default to 100]
 **pageSize** | **int32** |  | [default to 100]
 **sort** | **string** | Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc | [default to &quot;name:asc&quot;]

### Return type

[**ComputersSearchResults**](ComputersSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

