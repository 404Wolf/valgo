/*
Val Town API

Val Town’s public API  OpenAPI JSON endpoint:  https://api.val.town/openapi.json

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RunPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunPostRequest{}

// RunPostRequest Arguments to call the method with
type RunPostRequest struct {
	Args []interface{} `json:"args,omitempty"`
}

// NewRunPostRequest instantiates a new RunPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunPostRequest() *RunPostRequest {
	this := RunPostRequest{}
	return &this
}

// NewRunPostRequestWithDefaults instantiates a new RunPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunPostRequestWithDefaults() *RunPostRequest {
	this := RunPostRequest{}
	return &this
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *RunPostRequest) GetArgs() []interface{} {
	if o == nil || IsNil(o.Args) {
		var ret []interface{}
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunPostRequest) GetArgsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Args) {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *RunPostRequest) HasArgs() bool {
	if o != nil && !IsNil(o.Args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []interface{} and assigns it to the Args field.
func (o *RunPostRequest) SetArgs(v []interface{}) {
	o.Args = v
}

func (o RunPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Args) {
		toSerialize["args"] = o.Args
	}
	return toSerialize, nil
}

type NullableRunPostRequest struct {
	value *RunPostRequest
	isSet bool
}

func (v NullableRunPostRequest) Get() *RunPostRequest {
	return v.value
}

func (v *NullableRunPostRequest) Set(val *RunPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRunPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRunPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunPostRequest(val *RunPostRequest) *NullableRunPostRequest {
	return &NullableRunPostRequest{value: val, isSet: true}
}

func (v NullableRunPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

