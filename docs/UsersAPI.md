# \UsersAPI

All URIs are relative to *https://api.val.town*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersGet**](UsersAPI.md#UsersGet) | **Get** /v1/users/{user_id} | 
[**UsersVals**](UsersAPI.md#UsersVals) | **Get** /v1/users/{user_id}/vals | 



## UsersGet

> User UsersGet(ctx, userId).Execute()





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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersGet(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersGet`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersVals

> SearchVals200Response UsersVals(ctx, userId).Offset(offset).Limit(limit).Execute()





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
	offset := int32(56) // int32 | Number of items to skip in order to deliver paginated results (default to 0)
	limit := int32(56) // int32 | Maximum items to return in each paginated response (default to 20)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersVals(context.Background(), userId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersVals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersVals`: SearchVals200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersVals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersValsRequest struct via the builder pattern


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

