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

// checks if the EmailNameAndAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailNameAndAddress{}

// EmailNameAndAddress An email address and name
type EmailNameAndAddress struct {
	Name *string `json:"name,omitempty"`
	Email string `json:"email"`
}

type _EmailNameAndAddress EmailNameAndAddress

// NewEmailNameAndAddress instantiates a new EmailNameAndAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailNameAndAddress(email string) *EmailNameAndAddress {
	this := EmailNameAndAddress{}
	this.Email = email
	return &this
}

// NewEmailNameAndAddressWithDefaults instantiates a new EmailNameAndAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailNameAndAddressWithDefaults() *EmailNameAndAddress {
	this := EmailNameAndAddress{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EmailNameAndAddress) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNameAndAddress) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EmailNameAndAddress) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EmailNameAndAddress) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value
func (o *EmailNameAndAddress) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *EmailNameAndAddress) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *EmailNameAndAddress) SetEmail(v string) {
	o.Email = v
}

func (o EmailNameAndAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailNameAndAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

func (o *EmailNameAndAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varEmailNameAndAddress := _EmailNameAndAddress{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmailNameAndAddress)

	if err != nil {
		return err
	}

	*o = EmailNameAndAddress(varEmailNameAndAddress)

	return err
}

type NullableEmailNameAndAddress struct {
	value *EmailNameAndAddress
	isSet bool
}

func (v NullableEmailNameAndAddress) Get() *EmailNameAndAddress {
	return v.value
}

func (v *NullableEmailNameAndAddress) Set(val *EmailNameAndAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailNameAndAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailNameAndAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailNameAndAddress(val *EmailNameAndAddress) *NullableEmailNameAndAddress {
	return &NullableEmailNameAndAddress{value: val, isSet: true}
}

func (v NullableEmailNameAndAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailNameAndAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


