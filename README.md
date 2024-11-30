# Go API client for valgo

Val Town’s public API

OpenAPI JSON endpoint:

https://api.val.town/openapi.json

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1
- Package version: v1.0.0
- Generator version: 7.7.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import valgo "github.com/404wolf/valgo"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `valgo.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), valgo.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `valgo.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), valgo.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `valgo.ContextOperationServerIndices` and `valgo.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), valgo.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), valgo.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.val.town*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AliasAPI* | [**AliasUsername**](docs/AliasAPI.md#aliasusername) | **Get** /v1/alias/{username} | 
*AliasAPI* | [**AliasVal**](docs/AliasAPI.md#aliasval) | **Get** /v1/alias/{username}/{val_name} | 
*BlobsAPI* | [**BlobsDelete**](docs/BlobsAPI.md#blobsdelete) | **Delete** /v1/blob/{key} | 
*BlobsAPI* | [**BlobsGet**](docs/BlobsAPI.md#blobsget) | **Get** /v1/blob/{key} | 
*BlobsAPI* | [**BlobsList**](docs/BlobsAPI.md#blobslist) | **Get** /v1/blob | 
*BlobsAPI* | [**BlobsStore**](docs/BlobsAPI.md#blobsstore) | **Post** /v1/blob/{key} | 
*EmailsAPI* | [**EmailsSend**](docs/EmailsAPI.md#emailssend) | **Post** /v1/email | 
*MeAPI* | [**MeComments**](docs/MeAPI.md#mecomments) | **Get** /v1/me/comments | 
*MeAPI* | [**MeGet**](docs/MeAPI.md#meget) | **Get** /v1/me | 
*MeAPI* | [**MeLikes**](docs/MeAPI.md#melikes) | **Get** /v1/me/likes | 
*MeAPI* | [**MeReferences**](docs/MeAPI.md#mereferences) | **Get** /v1/me/references | 
*SearchAPI* | [**SearchVals**](docs/SearchAPI.md#searchvals) | **Get** /v1/search/vals | 
*SqliteAPI* | [**SqliteBatch**](docs/SqliteAPI.md#sqlitebatch) | **Post** /v1/sqlite/batch | 
*SqliteAPI* | [**SqliteExecute**](docs/SqliteAPI.md#sqliteexecute) | **Post** /v1/sqlite/execute | 
*UsersAPI* | [**UsersGet**](docs/UsersAPI.md#usersget) | **Get** /v1/users/{user_id} | 
*UsersAPI* | [**UsersVals**](docs/UsersAPI.md#usersvals) | **Get** /v1/users/{user_id}/vals | 
*ValsAPI* | [**RunGet**](docs/ValsAPI.md#runget) | **Get** /v1/run/{valname} | 
*ValsAPI* | [**RunPost**](docs/ValsAPI.md#runpost) | **Post** /v1/run/{valname} | 
*ValsAPI* | [**ValsCancel**](docs/ValsAPI.md#valscancel) | **Post** /v1/vals/{val_id}/evaluations/{evaluation_id}/cancel | 
*ValsAPI* | [**ValsCreate**](docs/ValsAPI.md#valscreate) | **Post** /v1/vals | 
*ValsAPI* | [**ValsCreateVersion**](docs/ValsAPI.md#valscreateversion) | **Post** /v1/vals/{val_id}/versions | 
*ValsAPI* | [**ValsDelete**](docs/ValsAPI.md#valsdelete) | **Delete** /v1/vals/{val_id} | 
*ValsAPI* | [**ValsDeleteVersion**](docs/ValsAPI.md#valsdeleteversion) | **Delete** /v1/vals/{val_id}/versions/{version} | 
*ValsAPI* | [**ValsGet**](docs/ValsAPI.md#valsget) | **Get** /v1/vals/{val_id} | 
*ValsAPI* | [**ValsGetVersion**](docs/ValsAPI.md#valsgetversion) | **Get** /v1/vals/{val_id}/versions/{version} | 
*ValsAPI* | [**ValsList**](docs/ValsAPI.md#valslist) | **Get** /v1/vals/{val_id}/versions | 
*ValsAPI* | [**ValsPut**](docs/ValsAPI.md#valsput) | **Put** /v1/vals | 
*ValsAPI* | [**ValsUpdate**](docs/ValsAPI.md#valsupdate) | **Put** /v1/vals/{val_id} | 


## Documentation For Models

 - [AliasVal200Response](docs/AliasVal200Response.md)
 - [AliasVal200ResponseAuthor](docs/AliasVal200ResponseAuthor.md)
 - [AliasVal200ResponseLinks](docs/AliasVal200ResponseLinks.md)
 - [AttachmentObject](docs/AttachmentObject.md)
 - [BasicVal](docs/BasicVal.md)
 - [BlobListingItem](docs/BlobListingItem.md)
 - [EmailData](docs/EmailData.md)
 - [EmailDataInput](docs/EmailDataInput.md)
 - [EmailNameAndAddress](docs/EmailNameAndAddress.md)
 - [EmailsSend200Response](docs/EmailsSend200Response.md)
 - [EmailsSend500Response](docs/EmailsSend500Response.md)
 - [EmailsSendRequest](docs/EmailsSendRequest.md)
 - [ExtendedVal](docs/ExtendedVal.md)
 - [MeComments200Response](docs/MeComments200Response.md)
 - [MeComments200ResponseDataInner](docs/MeComments200ResponseDataInner.md)
 - [MeComments200ResponseDataInnerAuthor](docs/MeComments200ResponseDataInnerAuthor.md)
 - [MeComments200ResponseDataInnerVal](docs/MeComments200ResponseDataInnerVal.md)
 - [MeGet200Response](docs/MeGet200Response.md)
 - [MeGet200ResponseLinks](docs/MeGet200ResponseLinks.md)
 - [MeReferences200Response](docs/MeReferences200Response.md)
 - [MeReferences200ResponseDataInner](docs/MeReferences200ResponseDataInner.md)
 - [MeReferences200ResponseDataInnerReference](docs/MeReferences200ResponseDataInnerReference.md)
 - [PaginationLinks](docs/PaginationLinks.md)
 - [ParameterizedQuery](docs/ParameterizedQuery.md)
 - [ReplyToList](docs/ReplyToList.md)
 - [ResultSet](docs/ResultSet.md)
 - [ResultSetLastInsertRowid](docs/ResultSetLastInsertRowid.md)
 - [RunPostRequest](docs/RunPostRequest.md)
 - [SearchVals200Response](docs/SearchVals200Response.md)
 - [SqliteBatchRequest](docs/SqliteBatchRequest.md)
 - [SqliteExecuteRequest](docs/SqliteExecuteRequest.md)
 - [SqliteExecuteRequestStatement](docs/SqliteExecuteRequestStatement.md)
 - [StatementArg](docs/StatementArg.md)
 - [User](docs/User.md)
 - [ValsCancel200Response](docs/ValsCancel200Response.md)
 - [ValsCreateRequest](docs/ValsCreateRequest.md)
 - [ValsList200Response](docs/ValsList200Response.md)
 - [ValsList200ResponseDataInner](docs/ValsList200ResponseDataInner.md)
 - [ValsPutRequest](docs/ValsPutRequest.md)
 - [ValsUpdateRequest](docs/ValsUpdateRequest.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), valgo.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



