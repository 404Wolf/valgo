# ResultSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | **[]string** | Names of columns.  Names of columns can be defined using the &#x60;AS&#x60; keyword in SQL:  &#x60;&#x60;&#x60;sql SELECT author AS author, COUNT(*) AS count FROM books GROUP BY author &#x60;&#x60;&#x60; | 
**ColumnTypes** | **[]string** | Types of columns.  The types are currently shown for types declared in a SQL table. For column types of function calls, for example, an empty string is returned. | 
**Rows** | **[][]interface{}** | Rows produced by the statement. | 
**RowsAffected** | **float32** | Number of rows that were affected by an UPDATE, INSERT or DELETE operation.  This value is not specified for other SQL statements. | 
**LastInsertRowid** | Pointer to [**NullableResultSetLastInsertRowid**](ResultSetLastInsertRowid.md) |  | [optional] 

## Methods

### NewResultSet

`func NewResultSet(columns []string, columnTypes []string, rows [][]interface{}, rowsAffected float32, ) *ResultSet`

NewResultSet instantiates a new ResultSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultSetWithDefaults

`func NewResultSetWithDefaults() *ResultSet`

NewResultSetWithDefaults instantiates a new ResultSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *ResultSet) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *ResultSet) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *ResultSet) SetColumns(v []string)`

SetColumns sets Columns field to given value.


### GetColumnTypes

`func (o *ResultSet) GetColumnTypes() []string`

GetColumnTypes returns the ColumnTypes field if non-nil, zero value otherwise.

### GetColumnTypesOk

`func (o *ResultSet) GetColumnTypesOk() (*[]string, bool)`

GetColumnTypesOk returns a tuple with the ColumnTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnTypes

`func (o *ResultSet) SetColumnTypes(v []string)`

SetColumnTypes sets ColumnTypes field to given value.


### GetRows

`func (o *ResultSet) GetRows() [][]interface{}`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ResultSet) GetRowsOk() (*[][]interface{}, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ResultSet) SetRows(v [][]interface{})`

SetRows sets Rows field to given value.


### GetRowsAffected

`func (o *ResultSet) GetRowsAffected() float32`

GetRowsAffected returns the RowsAffected field if non-nil, zero value otherwise.

### GetRowsAffectedOk

`func (o *ResultSet) GetRowsAffectedOk() (*float32, bool)`

GetRowsAffectedOk returns a tuple with the RowsAffected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowsAffected

`func (o *ResultSet) SetRowsAffected(v float32)`

SetRowsAffected sets RowsAffected field to given value.


### GetLastInsertRowid

`func (o *ResultSet) GetLastInsertRowid() ResultSetLastInsertRowid`

GetLastInsertRowid returns the LastInsertRowid field if non-nil, zero value otherwise.

### GetLastInsertRowidOk

`func (o *ResultSet) GetLastInsertRowidOk() (*ResultSetLastInsertRowid, bool)`

GetLastInsertRowidOk returns a tuple with the LastInsertRowid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInsertRowid

`func (o *ResultSet) SetLastInsertRowid(v ResultSetLastInsertRowid)`

SetLastInsertRowid sets LastInsertRowid field to given value.

### HasLastInsertRowid

`func (o *ResultSet) HasLastInsertRowid() bool`

HasLastInsertRowid returns a boolean if a field has been set.

### SetLastInsertRowidNil

`func (o *ResultSet) SetLastInsertRowidNil(b bool)`

 SetLastInsertRowidNil sets the value for LastInsertRowid to be an explicit nil

### UnsetLastInsertRowid
`func (o *ResultSet) UnsetLastInsertRowid()`

UnsetLastInsertRowid ensures that no value is present for LastInsertRowid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


