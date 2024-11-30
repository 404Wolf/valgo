# \AliasAPI

All URIs are relative to *https://api.val.town*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AliasUsername**](AliasAPI.md#AliasUsername) | **Get** /v1/alias/{username} | 
[**AliasVal**](AliasAPI.md#AliasVal) | **Get** /v1/alias/{username}/{val_name} | 



## AliasUsername

> User AliasUsername(ctx, username).Execute()





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
	username := "username_example" // string | Username of the user who you are looking for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AliasAPI.AliasUsername(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AliasAPI.AliasUsername``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AliasUsername`: User
	fmt.Fprintf(os.Stdout, "Response from `AliasAPI.AliasUsername`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of the user who you are looking for | 

### Other Parameters

Other parameters are passed through a pointer to a apiAliasUsernameRequest struct via the builder pattern


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


## AliasVal

> AliasVal200Response AliasVal(ctx, username, valName).Execute()





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
	username := "username_example" // string | Username of the user whose val you are looking for
	valName := "valName_example" // string | Name of the val you’re looking for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AliasAPI.AliasVal(context.Background(), username, valName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AliasAPI.AliasVal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AliasVal`: AliasVal200Response
	fmt.Fprintf(os.Stdout, "Response from `AliasAPI.AliasVal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of the user whose val you are looking for | 
**valName** | **string** | Name of the val you’re looking for | 

### Other Parameters

Other parameters are passed through a pointer to a apiAliasValRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AliasVal200Response**](AliasVal200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

