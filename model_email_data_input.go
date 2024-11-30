/*
Val Town API

Val Town’s public API  OpenAPI JSON endpoint:  https://api.val.town/openapi.json

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package valgo

import (
	"encoding/json"
	"fmt"
)

// EmailDataInput A single email or list of emails for one of the address fields
type EmailDataInput struct {
	EmailData *EmailData
	ArrayOfEmailData *[]EmailData
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EmailDataInput) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into EmailData
	err = json.Unmarshal(data, &dst.EmailData);
	if err == nil {
		jsonEmailData, _ := json.Marshal(dst.EmailData)
		if string(jsonEmailData) == "{}" { // empty struct
			dst.EmailData = nil
		} else {
			return nil // data stored in dst.EmailData, return on the first match
		}
	} else {
		dst.EmailData = nil
	}

	// try to unmarshal JSON data into ArrayOfEmailData
	err = json.Unmarshal(data, &dst.ArrayOfEmailData);
	if err == nil {
		jsonArrayOfEmailData, _ := json.Marshal(dst.ArrayOfEmailData)
		if string(jsonArrayOfEmailData) == "{}" { // empty struct
			dst.ArrayOfEmailData = nil
		} else {
			return nil // data stored in dst.ArrayOfEmailData, return on the first match
		}
	} else {
		dst.ArrayOfEmailData = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(EmailDataInput)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EmailDataInput) MarshalJSON() ([]byte, error) {
	if src.EmailData != nil {
		return json.Marshal(&src.EmailData)
	}

	if src.ArrayOfEmailData != nil {
		return json.Marshal(&src.ArrayOfEmailData)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEmailDataInput struct {
	value *EmailDataInput
	isSet bool
}

func (v NullableEmailDataInput) Get() *EmailDataInput {
	return v.value
}

func (v *NullableEmailDataInput) Set(val *EmailDataInput) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailDataInput) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailDataInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailDataInput(val *EmailDataInput) *NullableEmailDataInput {
	return &NullableEmailDataInput{value: val, isSet: true}
}

func (v NullableEmailDataInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailDataInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

