/*
Val Town API

Val Town’s public API  OpenAPI JSON endpoint:  https://api.val.town/openapi.json

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MeReferences200ResponseDataInnerReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeReferences200ResponseDataInnerReference{}

// MeReferences200ResponseDataInnerReference A val in a dependency relationship
type MeReferences200ResponseDataInnerReference struct {
	// The id of this val
	Id string `json:"id"`
	// The username of the person who authored this val
	Username NullableString `json:"username"`
	// The ID of the person who authored this val
	AuthorId NullableString `json:"author_id"`
	// The name of this val
	Name string `json:"name"`
}

type _MeReferences200ResponseDataInnerReference MeReferences200ResponseDataInnerReference

// NewMeReferences200ResponseDataInnerReference instantiates a new MeReferences200ResponseDataInnerReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeReferences200ResponseDataInnerReference(id string, username NullableString, authorId NullableString, name string) *MeReferences200ResponseDataInnerReference {
	this := MeReferences200ResponseDataInnerReference{}
	this.Id = id
	this.Username = username
	this.AuthorId = authorId
	this.Name = name
	return &this
}

// NewMeReferences200ResponseDataInnerReferenceWithDefaults instantiates a new MeReferences200ResponseDataInnerReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeReferences200ResponseDataInnerReferenceWithDefaults() *MeReferences200ResponseDataInnerReference {
	this := MeReferences200ResponseDataInnerReference{}
	return &this
}

// GetId returns the Id field value
func (o *MeReferences200ResponseDataInnerReference) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MeReferences200ResponseDataInnerReference) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MeReferences200ResponseDataInnerReference) SetId(v string) {
	o.Id = v
}

// GetUsername returns the Username field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MeReferences200ResponseDataInnerReference) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}

	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MeReferences200ResponseDataInnerReference) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// SetUsername sets field value
func (o *MeReferences200ResponseDataInnerReference) SetUsername(v string) {
	o.Username.Set(&v)
}

// GetAuthorId returns the AuthorId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MeReferences200ResponseDataInnerReference) GetAuthorId() string {
	if o == nil || o.AuthorId.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuthorId.Get()
}

// GetAuthorIdOk returns a tuple with the AuthorId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MeReferences200ResponseDataInnerReference) GetAuthorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorId.Get(), o.AuthorId.IsSet()
}

// SetAuthorId sets field value
func (o *MeReferences200ResponseDataInnerReference) SetAuthorId(v string) {
	o.AuthorId.Set(&v)
}

// GetName returns the Name field value
func (o *MeReferences200ResponseDataInnerReference) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MeReferences200ResponseDataInnerReference) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MeReferences200ResponseDataInnerReference) SetName(v string) {
	o.Name = v
}

func (o MeReferences200ResponseDataInnerReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeReferences200ResponseDataInnerReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["username"] = o.Username.Get()
	toSerialize["author_id"] = o.AuthorId.Get()
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *MeReferences200ResponseDataInnerReference) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"username",
		"author_id",
		"name",
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

	varMeReferences200ResponseDataInnerReference := _MeReferences200ResponseDataInnerReference{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMeReferences200ResponseDataInnerReference)

	if err != nil {
		return err
	}

	*o = MeReferences200ResponseDataInnerReference(varMeReferences200ResponseDataInnerReference)

	return err
}

type NullableMeReferences200ResponseDataInnerReference struct {
	value *MeReferences200ResponseDataInnerReference
	isSet bool
}

func (v NullableMeReferences200ResponseDataInnerReference) Get() *MeReferences200ResponseDataInnerReference {
	return v.value
}

func (v *NullableMeReferences200ResponseDataInnerReference) Set(val *MeReferences200ResponseDataInnerReference) {
	v.value = val
	v.isSet = true
}

func (v NullableMeReferences200ResponseDataInnerReference) IsSet() bool {
	return v.isSet
}

func (v *NullableMeReferences200ResponseDataInnerReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeReferences200ResponseDataInnerReference(val *MeReferences200ResponseDataInnerReference) *NullableMeReferences200ResponseDataInnerReference {
	return &NullableMeReferences200ResponseDataInnerReference{value: val, isSet: true}
}

func (v NullableMeReferences200ResponseDataInnerReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeReferences200ResponseDataInnerReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

