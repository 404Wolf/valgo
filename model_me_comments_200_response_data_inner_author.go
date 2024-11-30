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

// checks if the MeComments200ResponseDataInnerAuthor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeComments200ResponseDataInnerAuthor{}

// MeComments200ResponseDataInnerAuthor struct for MeComments200ResponseDataInnerAuthor
type MeComments200ResponseDataInnerAuthor struct {
	Id string `json:"id"`
	Username NullableString `json:"username"`
}

type _MeComments200ResponseDataInnerAuthor MeComments200ResponseDataInnerAuthor

// NewMeComments200ResponseDataInnerAuthor instantiates a new MeComments200ResponseDataInnerAuthor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeComments200ResponseDataInnerAuthor(id string, username NullableString) *MeComments200ResponseDataInnerAuthor {
	this := MeComments200ResponseDataInnerAuthor{}
	this.Id = id
	this.Username = username
	return &this
}

// NewMeComments200ResponseDataInnerAuthorWithDefaults instantiates a new MeComments200ResponseDataInnerAuthor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeComments200ResponseDataInnerAuthorWithDefaults() *MeComments200ResponseDataInnerAuthor {
	this := MeComments200ResponseDataInnerAuthor{}
	return &this
}

// GetId returns the Id field value
func (o *MeComments200ResponseDataInnerAuthor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MeComments200ResponseDataInnerAuthor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MeComments200ResponseDataInnerAuthor) SetId(v string) {
	o.Id = v
}

// GetUsername returns the Username field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MeComments200ResponseDataInnerAuthor) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}

	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MeComments200ResponseDataInnerAuthor) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// SetUsername sets field value
func (o *MeComments200ResponseDataInnerAuthor) SetUsername(v string) {
	o.Username.Set(&v)
}

func (o MeComments200ResponseDataInnerAuthor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeComments200ResponseDataInnerAuthor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["username"] = o.Username.Get()
	return toSerialize, nil
}

func (o *MeComments200ResponseDataInnerAuthor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"username",
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

	varMeComments200ResponseDataInnerAuthor := _MeComments200ResponseDataInnerAuthor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMeComments200ResponseDataInnerAuthor)

	if err != nil {
		return err
	}

	*o = MeComments200ResponseDataInnerAuthor(varMeComments200ResponseDataInnerAuthor)

	return err
}

type NullableMeComments200ResponseDataInnerAuthor struct {
	value *MeComments200ResponseDataInnerAuthor
	isSet bool
}

func (v NullableMeComments200ResponseDataInnerAuthor) Get() *MeComments200ResponseDataInnerAuthor {
	return v.value
}

func (v *NullableMeComments200ResponseDataInnerAuthor) Set(val *MeComments200ResponseDataInnerAuthor) {
	v.value = val
	v.isSet = true
}

func (v NullableMeComments200ResponseDataInnerAuthor) IsSet() bool {
	return v.isSet
}

func (v *NullableMeComments200ResponseDataInnerAuthor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeComments200ResponseDataInnerAuthor(val *MeComments200ResponseDataInnerAuthor) *NullableMeComments200ResponseDataInnerAuthor {
	return &NullableMeComments200ResponseDataInnerAuthor{value: val, isSet: true}
}

func (v NullableMeComments200ResponseDataInnerAuthor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeComments200ResponseDataInnerAuthor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

