# \SqliteAPI

All URIs are relative to *https://api.val.town*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SqliteBatch**](SqliteAPI.md#SqliteBatch) | **Post** /v1/sqlite/batch | 
[**SqliteExecute**](SqliteAPI.md#SqliteExecute) | **Post** /v1/sqlite/execute | 



## SqliteBatch

> []ResultSet SqliteBatch(ctx).SqliteBatchRequest(sqliteBatchRequest).Execute()





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
	sqliteBatchRequest := *openapiclient.NewSqliteBatchRequest([]openapiclient.SqliteExecuteRequestStatement{*openapiclient.NewSqliteExecuteRequestStatement("Sql_example", *openapiclient.NewStatementArg())}) // SqliteBatchRequest | A set of statements to be run in a single batch

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SqliteAPI.SqliteBatch(context.Background()).SqliteBatchRequest(sqliteBatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqliteAPI.SqliteBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SqliteBatch`: []ResultSet
	fmt.Fprintf(os.Stdout, "Response from `SqliteAPI.SqliteBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSqliteBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sqliteBatchRequest** | [**SqliteBatchRequest**](SqliteBatchRequest.md) | A set of statements to be run in a single batch | 

### Return type

[**[]ResultSet**](ResultSet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SqliteExecute

> ResultSet SqliteExecute(ctx).SqliteExecuteRequest(sqliteExecuteRequest).Execute()





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
	sqliteExecuteRequest := *openapiclient.NewSqliteExecuteRequest(*openapiclient.NewSqliteExecuteRequestStatement("Sql_example", *openapiclient.NewStatementArg())) // SqliteExecuteRequest | A single statement to run

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SqliteAPI.SqliteExecute(context.Background()).SqliteExecuteRequest(sqliteExecuteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqliteAPI.SqliteExecute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SqliteExecute`: ResultSet
	fmt.Fprintf(os.Stdout, "Response from `SqliteAPI.SqliteExecute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSqliteExecuteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sqliteExecuteRequest** | [**SqliteExecuteRequest**](SqliteExecuteRequest.md) | A single statement to run | 

### Return type

[**ResultSet**](ResultSet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

