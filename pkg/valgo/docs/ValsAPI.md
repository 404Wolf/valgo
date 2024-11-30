# \ValsAPI

All URIs are relative to *https://api.val.town*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RunGet**](ValsAPI.md#RunGet) | **Get** /v1/run/{valname} | 
[**RunPost**](ValsAPI.md#RunPost) | **Post** /v1/run/{valname} | 
[**ValsCancel**](ValsAPI.md#ValsCancel) | **Post** /v1/vals/{val_id}/evaluations/{evaluation_id}/cancel | 
[**ValsCreate**](ValsAPI.md#ValsCreate) | **Post** /v1/vals | 
[**ValsCreateVersion**](ValsAPI.md#ValsCreateVersion) | **Post** /v1/vals/{val_id}/versions | 
[**ValsDelete**](ValsAPI.md#ValsDelete) | **Delete** /v1/vals/{val_id} | 
[**ValsDeleteVersion**](ValsAPI.md#ValsDeleteVersion) | **Delete** /v1/vals/{val_id}/versions/{version} | 
[**ValsGet**](ValsAPI.md#ValsGet) | **Get** /v1/vals/{val_id} | 
[**ValsGetVersion**](ValsAPI.md#ValsGetVersion) | **Get** /v1/vals/{val_id}/versions/{version} | 
[**ValsList**](ValsAPI.md#ValsList) | **Get** /v1/vals/{val_id}/versions | 
[**ValsPut**](ValsAPI.md#ValsPut) | **Put** /v1/vals | 
[**ValsUpdate**](ValsAPI.md#ValsUpdate) | **Put** /v1/vals/{val_id} | 



## RunGet

> RunGet(ctx, valname).Args(args).Execute()





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
	valname := "valname_example" // string | Name of the val, in concatenated style, like username.valname
	args := "args_example" // string | An argument that will be passed to the function as a JavaScript value (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ValsAPI.RunGet(context.Background(), valname).Args(args).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValsAPI.RunGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**valname** | **string** | Name of the val, in concatenated style, like username.valname | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **args** | **string** | An argument that will be passed to the function as a JavaScript value | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunPost

> RunPost(ctx, valname).RunPostRequest(runPostRequest).Execute()





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
	valname := "valname_example" // string | Name of the val, in concatenated style, like username.valname
	runPostRequest := *openapiclient.NewRunPostRequest() // RunPostRequest | Arguments to call the method with (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ValsAPI.RunPost(context.Background(), valname).RunPostRequest(runPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValsAPI.RunPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**valname** | **string** | Name of the val, in concatenated style, like username.valname | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **runPostRequest** | [**RunPostRequest**](RunPostRequest.md) | Arguments to call the method with | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValsCancel

> ValsCancel200Response ValsCancel(ctx, valId, evaluationId).Execute()





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
	valId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of a val
	evaluationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the evaluation - the specific time that a val was run - that you want to cancel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValsAPI.ValsCancel(context.Background(), valId, evaluationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValsAPI.ValsCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValsCancel`: ValsCancel200Response
	fmt.Fprintf(os.Stdout, "Response from `ValsAPI.ValsCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**valId** | **string** | Id of a val | 
**evaluationId** | **string** | The ID of the evaluation - the specific time that a val was run - that you want to cancel | 

### Other Parameters

Other parameters are passed through a pointer to a apiValsCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ValsCancel200Response**](ValsCancel200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValsCreate

> ExtendedVal ValsCreate(ctx).ValsCreateRequest(valsCreateRequest).Execute()





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
	valsCreateRequest := *openapiclient.NewValsCreateRequest("Code_example") // ValsCreateRequest | Val information provided to create a new val

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValsAPI.ValsCreate(context.Background()).ValsCreateRequest(valsCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValsAPI.ValsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValsCreate`: ExtendedVal
	fmt.Fprintf(os.Stdout, "Response from `ValsAPI.ValsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **valsCreateRequest** | [**ValsCreateRequest**](ValsCreateRequest.md) | Val information provided to create a new val | 

### Return type

[**ExtendedVal**](ExtendedVal.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValsCreateVersion

> ExtendedVal ValsCreateVersion(ctx, valId).ValsCreateRequest(valsCreateRequest).Execute()





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
	valId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of a val
	valsCreateRequest := *openapiclient.NewValsCreateRequest("Code_example") // ValsCreateRequest | Val information provided to create a new val

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValsAPI.ValsCreateVersion(context.Background(), valId).ValsCreateRequest(valsCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValsAPI.ValsCreateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValsCreateVersion`: ExtendedVal
	fmt.Fprintf(os.Stdout, "Response from `ValsAPI.ValsCreateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**valId** | **string** | Id of a val | 

### Other Parameters

Other parameters are passed through a pointer to a apiValsCreateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **valsCreateRequest** | [**ValsCreateRequest**](ValsCreateRequest.md) | Val information provided to create a new val | 

### Return type

[**ExtendedVal**](ExtendedVal.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValsDelete

> ValsDelete(ctx, valId).Execute()





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
	valId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of a val

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ValsAPI.ValsDelete(context.Background(), valId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValsAPI.ValsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**valId** | **string** | Id of a val | 

### Other Parameters

Other parameters are passed through a pointer to a apiValsDeleteRequest struct via the builder pattern


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


## ValsDeleteVersion

> ValsDeleteVersion(ctx, valId, version).Execute()





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
	valId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of a val
	version := int32(56) // int32 | The val version

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ValsAPI.ValsDeleteVersion(context.Background(), valId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValsAPI.ValsDeleteVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**valId** | **string** | Id of a val | 
**version** | **int32** | The val version | 

### Other Parameters

Other parameters are passed through a pointer to a apiValsDeleteVersionRequest struct via the builder pattern


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


## ValsGet

> ExtendedVal ValsGet(ctx, valId).Execute()





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
	valId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of a val

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValsAPI.ValsGet(context.Background(), valId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValsAPI.ValsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValsGet`: ExtendedVal
	fmt.Fprintf(os.Stdout, "Response from `ValsAPI.ValsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**valId** | **string** | Id of a val | 

### Other Parameters

Other parameters are passed through a pointer to a apiValsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExtendedVal**](ExtendedVal.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValsGetVersion

> ExtendedVal ValsGetVersion(ctx, valId, version).Offset(offset).Limit(limit).Execute()





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
	valId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of a val
	version := int32(56) // int32 | The val version

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValsAPI.ValsGetVersion(context.Background(), valId, version).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValsAPI.ValsGetVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValsGetVersion`: ExtendedVal
	fmt.Fprintf(os.Stdout, "Response from `ValsAPI.ValsGetVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**valId** | **string** | Id of a val | 
**version** | **int32** | The val version | 

### Other Parameters

Other parameters are passed through a pointer to a apiValsGetVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Number of items to skip in order to deliver paginated results | [default to 0]
 **limit** | **int32** | Maximum items to return in each paginated response | [default to 20]



### Return type

[**ExtendedVal**](ExtendedVal.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValsList

> ValsList200Response ValsList(ctx, valId).Offset(offset).Limit(limit).Execute()





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
	valId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of a val

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValsAPI.ValsList(context.Background(), valId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValsAPI.ValsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValsList`: ValsList200Response
	fmt.Fprintf(os.Stdout, "Response from `ValsAPI.ValsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**valId** | **string** | Id of a val | 

### Other Parameters

Other parameters are passed through a pointer to a apiValsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Number of items to skip in order to deliver paginated results | [default to 0]
 **limit** | **int32** | Maximum items to return in each paginated response | [default to 20]


### Return type

[**ValsList200Response**](ValsList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValsPut

> ValsPut(ctx).ValsPutRequest(valsPutRequest).Execute()





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
	valsPutRequest := *openapiclient.NewValsPutRequest("Code_example", "Name_example") // ValsPutRequest | Create a new val, or version of a val, with new code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ValsAPI.ValsPut(context.Background()).ValsPutRequest(valsPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValsAPI.ValsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **valsPutRequest** | [**ValsPutRequest**](ValsPutRequest.md) | Create a new val, or version of a val, with new code | 

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


## ValsUpdate

> ValsUpdate(ctx, valId).ValsUpdateRequest(valsUpdateRequest).Execute()





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
	valId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of a val
	valsUpdateRequest := *openapiclient.NewValsUpdateRequest() // ValsUpdateRequest | Note: you must supply at least one of the keys in this object in order to update a val (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ValsAPI.ValsUpdate(context.Background(), valId).ValsUpdateRequest(valsUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValsAPI.ValsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**valId** | **string** | Id of a val | 

### Other Parameters

Other parameters are passed through a pointer to a apiValsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **valsUpdateRequest** | [**ValsUpdateRequest**](ValsUpdateRequest.md) | Note: you must supply at least one of the keys in this object in order to update a val | 

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

