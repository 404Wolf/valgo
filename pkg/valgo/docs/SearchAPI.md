# \SearchAPI

All URIs are relative to *https://api.val.town*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchVals**](SearchAPI.md#SearchVals) | **Get** /v1/search/vals | 



## SearchVals

> SearchVals200Response SearchVals(ctx).Offset(offset).Limit(limit).Query(query).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	offset := int32(56) // int32 | Number of items to skip in order to deliver paginated results (default to 0)
	limit := int32(56) // int32 | Maximum items to return in each paginated response (default to 20)
	query := "query_example" // string | Search query

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchVals(context.Background()).Offset(offset).Limit(limit).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchVals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchVals`: SearchVals200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchVals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchValsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Number of items to skip in order to deliver paginated results | [default to 0]
 **limit** | **int32** | Maximum items to return in each paginated response | [default to 20]
 **query** | **string** | Search query | 

### Return type

[**SearchVals200Response**](SearchVals200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

