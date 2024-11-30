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

// checks if the SearchVals200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchVals200Response{}

// SearchVals200Response A paginated result set
type SearchVals200Response struct {
	Data []BasicVal `json:"data"`
	Links PaginationLinks `json:"links"`
}

type _SearchVals200Response SearchVals200Response

// NewSearchVals200Response instantiates a new SearchVals200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchVals200Response(data []BasicVal, links PaginationLinks) *SearchVals200Response {
	this := SearchVals200Response{}
	this.Data = data
	this.Links = links
	return &this
}

// NewSearchVals200ResponseWithDefaults instantiates a new SearchVals200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchVals200ResponseWithDefaults() *SearchVals200Response {
	this := SearchVals200Response{}
	return &this
}

// GetData returns the Data field value
func (o *SearchVals200Response) GetData() []BasicVal {
	if o == nil {
		var ret []BasicVal
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SearchVals200Response) GetDataOk() ([]BasicVal, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SearchVals200Response) SetData(v []BasicVal) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *SearchVals200Response) GetLinks() PaginationLinks {
	if o == nil {
		var ret PaginationLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SearchVals200Response) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SearchVals200Response) SetLinks(v PaginationLinks) {
	o.Links = v
}

func (o SearchVals200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchVals200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *SearchVals200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"links",
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

	varSearchVals200Response := _SearchVals200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchVals200Response)

	if err != nil {
		return err
	}

	*o = SearchVals200Response(varSearchVals200Response)

	return err
}

type NullableSearchVals200Response struct {
	value *SearchVals200Response
	isSet bool
}

func (v NullableSearchVals200Response) Get() *SearchVals200Response {
	return v.value
}

func (v *NullableSearchVals200Response) Set(val *SearchVals200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchVals200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchVals200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchVals200Response(val *SearchVals200Response) *NullableSearchVals200Response {
	return &NullableSearchVals200Response{value: val, isSet: true}
}

func (v NullableSearchVals200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchVals200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

