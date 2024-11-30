# SqliteExecuteRequestStatement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sql** | **string** | SQL statement, with ? placeholders for arguments | 
**Args** | [**StatementArg**](StatementArg.md) |  | 

## Methods

### NewSqliteExecuteRequestStatement

`func NewSqliteExecuteRequestStatement(sql string, args StatementArg, ) *SqliteExecuteRequestStatement`

NewSqliteExecuteRequestStatement instantiates a new SqliteExecuteRequestStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqliteExecuteRequestStatementWithDefaults

`func NewSqliteExecuteRequestStatementWithDefaults() *SqliteExecuteRequestStatement`

NewSqliteExecuteRequestStatementWithDefaults instantiates a new SqliteExecuteRequestStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSql

`func (o *SqliteExecuteRequestStatement) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *SqliteExecuteRequestStatement) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *SqliteExecuteRequestStatement) SetSql(v string)`

SetSql sets Sql field to given value.


### GetArgs

`func (o *SqliteExecuteRequestStatement) GetArgs() StatementArg`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *SqliteExecuteRequestStatement) GetArgsOk() (*StatementArg, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *SqliteExecuteRequestStatement) SetArgs(v StatementArg)`

SetArgs sets Args field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


