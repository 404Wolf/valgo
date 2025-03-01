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

// checks if the SqliteBatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SqliteBatchRequest{}

// SqliteBatchRequest A set of statements to be run in a single batch
type SqliteBatchRequest struct {
	Statements []SqliteExecuteRequestStatement `json:"statements"`
	Mode *string `json:"mode,omitempty"`
}

type _SqliteBatchRequest SqliteBatchRequest

// NewSqliteBatchRequest instantiates a new SqliteBatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSqliteBatchRequest(statements []SqliteExecuteRequestStatement) *SqliteBatchRequest {
	this := SqliteBatchRequest{}
	this.Statements = statements
	return &this
}

// NewSqliteBatchRequestWithDefaults instantiates a new SqliteBatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSqliteBatchRequestWithDefaults() *SqliteBatchRequest {
	this := SqliteBatchRequest{}
	return &this
}

// GetStatements returns the Statements field value
func (o *SqliteBatchRequest) GetStatements() []SqliteExecuteRequestStatement {
	if o == nil {
		var ret []SqliteExecuteRequestStatement
		return ret
	}

	return o.Statements
}

// GetStatementsOk returns a tuple with the Statements field value
// and a boolean to check if the value has been set.
func (o *SqliteBatchRequest) GetStatementsOk() ([]SqliteExecuteRequestStatement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Statements, true
}

// SetStatements sets field value
func (o *SqliteBatchRequest) SetStatements(v []SqliteExecuteRequestStatement) {
	o.Statements = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *SqliteBatchRequest) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqliteBatchRequest) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *SqliteBatchRequest) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *SqliteBatchRequest) SetMode(v string) {
	o.Mode = &v
}

func (o SqliteBatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SqliteBatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["statements"] = o.Statements
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return toSerialize, nil
}

func (o *SqliteBatchRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"statements",
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

	varSqliteBatchRequest := _SqliteBatchRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSqliteBatchRequest)

	if err != nil {
		return err
	}

	*o = SqliteBatchRequest(varSqliteBatchRequest)

	return err
}

type NullableSqliteBatchRequest struct {
	value *SqliteBatchRequest
	isSet bool
}

func (v NullableSqliteBatchRequest) Get() *SqliteBatchRequest {
	return v.value
}

func (v *NullableSqliteBatchRequest) Set(val *SqliteBatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSqliteBatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSqliteBatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSqliteBatchRequest(val *SqliteBatchRequest) *NullableSqliteBatchRequest {
	return &NullableSqliteBatchRequest{value: val, isSet: true}
}

func (v NullableSqliteBatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSqliteBatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


