# \BetaAPI

All URIs are relative to *https://api.val.town*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BranchesGet**](BetaAPI.md#BranchesGet) | **Get** /v1/projects/{project_id}/branches/{branch_id} | 
[**BranchesList**](BetaAPI.md#BranchesList) | **Get** /v1/projects/{project_id}/branches | 
[**ProjectsGet**](BetaAPI.md#ProjectsGet) | **Get** /v1/projects/{project_id} | 
[**ProjectsList**](BetaAPI.md#ProjectsList) | **Get** /v1/projects | 



## BranchesGet

> Branch BranchesGet(ctx, projectId, branchId).Execute()





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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of a project
	branchId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of a branch

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAPI.BranchesGet(context.Background(), projectId, branchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAPI.BranchesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BranchesGet`: Branch
	fmt.Fprintf(os.Stdout, "Response from `BetaAPI.BranchesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Id of a project | 
**branchId** | **string** | Id of a branch | 

### Other Parameters

Other parameters are passed through a pointer to a apiBranchesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Branch**](Branch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BranchesList

> BranchesList200Response BranchesList(ctx, projectId).Offset(offset).Limit(limit).Execute()





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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of a project

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAPI.BranchesList(context.Background(), projectId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAPI.BranchesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BranchesList`: BranchesList200Response
	fmt.Fprintf(os.Stdout, "Response from `BetaAPI.BranchesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Id of a project | 

### Other Parameters

Other parameters are passed through a pointer to a apiBranchesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Number of items to skip in order to deliver paginated results | [default to 0]
 **limit** | **int32** | Maximum items to return in each paginated response | [default to 20]


### Return type

[**BranchesList200Response**](BranchesList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsGet

> Project ProjectsGet(ctx, projectId).Execute()





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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of a project

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAPI.ProjectsGet(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAPI.ProjectsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsGet`: Project
	fmt.Fprintf(os.Stdout, "Response from `BetaAPI.ProjectsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Id of a project | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsList

> MeProjects200Response ProjectsList(ctx).Offset(offset).Limit(limit).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaAPI.ProjectsList(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaAPI.ProjectsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsList`: MeProjects200Response
	fmt.Fprintf(os.Stdout, "Response from `BetaAPI.ProjectsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Number of items to skip in order to deliver paginated results | [default to 0]
 **limit** | **int32** | Maximum items to return in each paginated response | [default to 20]

### Return type

[**MeProjects200Response**](MeProjects200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

