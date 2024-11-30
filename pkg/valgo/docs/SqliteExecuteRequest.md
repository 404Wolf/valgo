# SqliteExecuteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statement** | [**SqliteExecuteRequestStatement**](SqliteExecuteRequestStatement.md) |  | 

## Methods

### NewSqliteExecuteRequest

`func NewSqliteExecuteRequest(statement SqliteExecuteRequestStatement, ) *SqliteExecuteRequest`

NewSqliteExecuteRequest instantiates a new SqliteExecuteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqliteExecuteRequestWithDefaults

`func NewSqliteExecuteRequestWithDefaults() *SqliteExecuteRequest`

NewSqliteExecuteRequestWithDefaults instantiates a new SqliteExecuteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatement

`func (o *SqliteExecuteRequest) GetStatement() SqliteExecuteRequestStatement`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *SqliteExecuteRequest) GetStatementOk() (*SqliteExecuteRequestStatement, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *SqliteExecuteRequest) SetStatement(v SqliteExecuteRequestStatement)`

SetStatement sets Statement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


