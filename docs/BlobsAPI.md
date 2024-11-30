# \BlobsAPI

All URIs are relative to *https://api.val.town*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlobsDelete**](BlobsAPI.md#BlobsDelete) | **Delete** /v1/blob/{key} | 
[**BlobsGet**](BlobsAPI.md#BlobsGet) | **Get** /v1/blob/{key} | 
[**BlobsList**](BlobsAPI.md#BlobsList) | **Get** /v1/blob | 
[**BlobsStore**](BlobsAPI.md#BlobsStore) | **Post** /v1/blob/{key} | 



## BlobsDelete

> BlobsDelete(ctx, key).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/404wolf/valgo"
)

func main() {
	key := "key_example" // string | Key that uniquely identifies this blob

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlobsAPI.BlobsDelete(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobsAPI.BlobsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key that uniquely identifies this blob | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlobsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlobsGet

> *os.File BlobsGet(ctx, key).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/404wolf/valgo"
)

func main() {
	key := "key_example" // string | Key that uniquely identifies this blob

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlobsAPI.BlobsGet(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobsAPI.BlobsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlobsGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `BlobsAPI.BlobsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key that uniquely identifies this blob | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlobsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlobsList

> []BlobListingItem BlobsList(ctx).Prefix(prefix).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/404wolf/valgo"
)

func main() {
	prefix := "prefix_example" // string | If specified, only include blobs that start with this string (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlobsAPI.BlobsList(context.Background()).Prefix(prefix).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobsAPI.BlobsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlobsList`: []BlobListingItem
	fmt.Fprintf(os.Stdout, "Response from `BlobsAPI.BlobsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlobsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **string** | If specified, only include blobs that start with this string | 

### Return type

[**[]BlobListingItem**](BlobListingItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlobsStore

> BlobsStore(ctx, key).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/404wolf/valgo"
)

func main() {
	key := "key_example" // string | Key that uniquely identifies this blob
	body := os.NewFile(1234, "some_file") // *os.File | Binary input data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlobsAPI.BlobsStore(context.Background(), key).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobsAPI.BlobsStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key that uniquely identifies this blob | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlobsStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** | Binary input data | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

