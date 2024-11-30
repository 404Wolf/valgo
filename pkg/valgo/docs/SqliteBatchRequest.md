# SqliteBatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statements** | [**[]SqliteExecuteRequestStatement**](SqliteExecuteRequestStatement.md) |  | 
**Mode** | Pointer to **string** |  | [optional] 

## Methods

### NewSqliteBatchRequest

`func NewSqliteBatchRequest(statements []SqliteExecuteRequestStatement, ) *SqliteBatchRequest`

NewSqliteBatchRequest instantiates a new SqliteBatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqliteBatchRequestWithDefaults

`func NewSqliteBatchRequestWithDefaults() *SqliteBatchRequest`

NewSqliteBatchRequestWithDefaults instantiates a new SqliteBatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatements

`func (o *SqliteBatchRequest) GetStatements() []SqliteExecuteRequestStatement`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *SqliteBatchRequest) GetStatementsOk() (*[]SqliteExecuteRequestStatement, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *SqliteBatchRequest) SetStatements(v []SqliteExecuteRequestStatement)`

SetStatements sets Statements field to given value.


### GetMode

`func (o *SqliteBatchRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SqliteBatchRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SqliteBatchRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SqliteBatchRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


