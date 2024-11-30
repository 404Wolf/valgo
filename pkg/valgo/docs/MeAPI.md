# \MeAPI

All URIs are relative to *https://api.val.town*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeComments**](MeAPI.md#MeComments) | **Get** /v1/me/comments | 
[**MeGet**](MeAPI.md#MeGet) | **Get** /v1/me | 
[**MeLikes**](MeAPI.md#MeLikes) | **Get** /v1/me/likes | 
[**MeReferences**](MeAPI.md#MeReferences) | **Get** /v1/me/references | 



## MeComments

> MeComments200Response MeComments(ctx).Offset(offset).Limit(limit).Relationship(relationship).Since(since).Until(until).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	offset := int32(56) // int32 | Number of items to skip in order to deliver paginated results (default to 0)
	limit := int32(56) // int32 | Maximum items to return in each paginated response (default to 20)
	relationship := "relationship_example" // string | Whether to get comments you have received, given, or both (default to "any")
	since := time.Now() // time.Time | Include items created after this date (optional)
	until := time.Now() // time.Time | Include items created before this date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.MeComments(context.Background()).Offset(offset).Limit(limit).Relationship(relationship).Since(since).Until(until).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.MeComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MeComments`: MeComments200Response
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.MeComments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Number of items to skip in order to deliver paginated results | [default to 0]
 **limit** | **int32** | Maximum items to return in each paginated response | [default to 20]
 **relationship** | **string** | Whether to get comments you have received, given, or both | [default to &quot;any&quot;]
 **since** | **time.Time** | Include items created after this date | 
 **until** | **time.Time** | Include items created before this date | 

### Return type

[**MeComments200Response**](MeComments200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeGet

> MeGet200Response MeGet(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.MeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.MeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MeGet`: MeGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.MeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMeGetRequest struct via the builder pattern


### Return type

[**MeGet200Response**](MeGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeLikes

> SearchVals200Response MeLikes(ctx).Offset(offset).Limit(limit).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.MeLikes(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.MeLikes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MeLikes`: SearchVals200Response
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.MeLikes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeLikesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Number of items to skip in order to deliver paginated results | [default to 0]
 **limit** | **int32** | Maximum items to return in each paginated response | [default to 20]

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


## MeReferences

> MeReferences200Response MeReferences(ctx).Offset(offset).Limit(limit).Since(since).Until(until).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	offset := int32(56) // int32 | Number of items to skip in order to deliver paginated results (default to 0)
	limit := int32(56) // int32 | Maximum items to return in each paginated response (default to 20)
	since := time.Now() // time.Time | Include items created after this date (optional)
	until := time.Now() // time.Time | Include items created before this date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.MeReferences(context.Background()).Offset(offset).Limit(limit).Since(since).Until(until).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.MeReferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MeReferences`: MeReferences200Response
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.MeReferences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Number of items to skip in order to deliver paginated results | [default to 0]
 **limit** | **int32** | Maximum items to return in each paginated response | [default to 20]
 **since** | **time.Time** | Include items created after this date | 
 **until** | **time.Time** | Include items created before this date | 

### Return type

[**MeReferences200Response**](MeReferences200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

