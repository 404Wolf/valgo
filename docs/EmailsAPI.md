# \EmailsAPI

All URIs are relative to *https://api.val.town*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmailsSend**](EmailsAPI.md#EmailsSend) | **Post** /v1/email | 



## EmailsSend

> EmailsSend200Response EmailsSend(ctx).EmailsSendRequest(emailsSendRequest).Execute()





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
	emailsSendRequest := *openapiclient.NewEmailsSendRequest() // EmailsSendRequest | Fields for an email to be sent (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailsAPI.EmailsSend(context.Background()).EmailsSendRequest(emailsSendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.EmailsSend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmailsSend`: EmailsSend200Response
	fmt.Fprintf(os.Stdout, "Response from `EmailsAPI.EmailsSend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailsSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailsSendRequest** | [**EmailsSendRequest**](EmailsSendRequest.md) | Fields for an email to be sent | 

### Return type

[**EmailsSend200Response**](EmailsSend200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

