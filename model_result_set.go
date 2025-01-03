/*
Val Town API

Val Town’s public API  OpenAPI JSON endpoint:  https://api.val.town/openapi.json

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package valgo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ResultSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultSet{}

// ResultSet Result of executing an SQL statement.
type ResultSet struct {
	// Names of columns.  Names of columns can be defined using the `AS` keyword in SQL:  ```sql SELECT author AS author, COUNT(*) AS count FROM books GROUP BY author ```
	Columns []string `json:"columns"`
	// Types of columns.  The types are currently shown for types declared in a SQL table. For column types of function calls, for example, an empty string is returned.
	ColumnTypes []string `json:"columnTypes"`
	// Rows produced by the statement.
	Rows [][]interface{} `json:"rows"`
	// Number of rows that were affected by an UPDATE, INSERT or DELETE operation.  This value is not specified for other SQL statements.
	RowsAffected float32 `json:"rowsAffected"`
	LastInsertRowid NullableResultSetLastInsertRowid `json:"lastInsertRowid,omitempty"`
}

type _ResultSet ResultSet

// NewResultSet instantiates a new ResultSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultSet(columns []string, columnTypes []string, rows [][]interface{}, rowsAffected float32) *ResultSet {
	this := ResultSet{}
	this.Columns = columns
	this.ColumnTypes = columnTypes
	this.Rows = rows
	this.RowsAffected = rowsAffected
	return &this
}

// NewResultSetWithDefaults instantiates a new ResultSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultSetWithDefaults() *ResultSet {
	this := ResultSet{}
	return &this
}

// GetColumns returns the Columns field value
func (o *ResultSet) GetColumns() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *ResultSet) GetColumnsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Columns, true
}

// SetColumns sets field value
func (o *ResultSet) SetColumns(v []string) {
	o.Columns = v
}

// GetColumnTypes returns the ColumnTypes field value
func (o *ResultSet) GetColumnTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ColumnTypes
}

// GetColumnTypesOk returns a tuple with the ColumnTypes field value
// and a boolean to check if the value has been set.
func (o *ResultSet) GetColumnTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ColumnTypes, true
}

// SetColumnTypes sets field value
func (o *ResultSet) SetColumnTypes(v []string) {
	o.ColumnTypes = v
}

// GetRows returns the Rows field value
func (o *ResultSet) GetRows() [][]interface{} {
	if o == nil {
		var ret [][]interface{}
		return ret
	}

	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value
// and a boolean to check if the value has been set.
func (o *ResultSet) GetRowsOk() ([][]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rows, true
}

// SetRows sets field value
func (o *ResultSet) SetRows(v [][]interface{}) {
	o.Rows = v
}

// GetRowsAffected returns the RowsAffected field value
func (o *ResultSet) GetRowsAffected() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RowsAffected
}

// GetRowsAffectedOk returns a tuple with the RowsAffected field value
// and a boolean to check if the value has been set.
func (o *ResultSet) GetRowsAffectedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RowsAffected, true
}

// SetRowsAffected sets field value
func (o *ResultSet) SetRowsAffected(v float32) {
	o.RowsAffected = v
}

// GetLastInsertRowid returns the LastInsertRowid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultSet) GetLastInsertRowid() ResultSetLastInsertRowid {
	if o == nil || IsNil(o.LastInsertRowid.Get()) {
		var ret ResultSetLastInsertRowid
		return ret
	}
	return *o.LastInsertRowid.Get()
}

// GetLastInsertRowidOk returns a tuple with the LastInsertRowid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultSet) GetLastInsertRowidOk() (*ResultSetLastInsertRowid, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastInsertRowid.Get(), o.LastInsertRowid.IsSet()
}

// HasLastInsertRowid returns a boolean if a field has been set.
func (o *ResultSet) HasLastInsertRowid() bool {
	if o != nil && o.LastInsertRowid.IsSet() {
		return true
	}

	return false
}

// SetLastInsertRowid gets a reference to the given NullableResultSetLastInsertRowid and assigns it to the LastInsertRowid field.
func (o *ResultSet) SetLastInsertRowid(v ResultSetLastInsertRowid) {
	o.LastInsertRowid.Set(&v)
}
// SetLastInsertRowidNil sets the value for LastInsertRowid to be an explicit nil
func (o *ResultSet) SetLastInsertRowidNil() {
	o.LastInsertRowid.Set(nil)
}

// UnsetLastInsertRowid ensures that no value is present for LastInsertRowid, not even an explicit nil
func (o *ResultSet) UnsetLastInsertRowid() {
	o.LastInsertRowid.Unset()
}

func (o ResultSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["columns"] = o.Columns
	toSerialize["columnTypes"] = o.ColumnTypes
	toSerialize["rows"] = o.Rows
	toSerialize["rowsAffected"] = o.RowsAffected
	if o.LastInsertRowid.IsSet() {
		toSerialize["lastInsertRowid"] = o.LastInsertRowid.Get()
	}
	return toSerialize, nil
}

func (o *ResultSet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"columns",
		"columnTypes",
		"rows",
		"rowsAffected",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varResultSet := _ResultSet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResultSet)

	if err != nil {
		return err
	}

	*o = ResultSet(varResultSet)

	return err
}

type NullableResultSet struct {
	value *ResultSet
	isSet bool
}

func (v NullableResultSet) Get() *ResultSet {
	return v.value
}

func (v *NullableResultSet) Set(val *ResultSet) {
	v.value = val
	v.isSet = true
}

func (v NullableResultSet) IsSet() bool {
	return v.isSet
}

func (v *NullableResultSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultSet(val *ResultSet) *NullableResultSet {
	return &NullableResultSet{value: val, isSet: true}
}

func (v NullableResultSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


