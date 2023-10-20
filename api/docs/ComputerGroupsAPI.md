# \ComputerGroupsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ComputerGroupsGet**](ComputerGroupsAPI.md#V1ComputerGroupsGet) | **Get** /v1/computer-groups | Returns the list of all computer groups 



## V1ComputerGroupsGet

> []ComputerGroup V1ComputerGroupsGet(ctx).Execute()

Returns the list of all computer groups 



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
    resp, r, err := apiClient.ComputerGroupsAPI.V1ComputerGroupsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComputerGroupsAPI.V1ComputerGroupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ComputerGroupsGet`: []ComputerGroup
    fmt.Fprintf(os.Stdout, "Response from `ComputerGroupsAPI.V1ComputerGroupsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ComputerGroupsGetRequest struct via the builder pattern


### Return type

[**[]ComputerGroup**](ComputerGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

