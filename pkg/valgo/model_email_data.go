/*
Val Town API

Val Town’s public API  OpenAPI JSON endpoint:  https://api.val.town/openapi.json

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// EmailData struct for EmailData
type EmailData struct {
	EmailNameAndAddress *EmailNameAndAddress
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EmailData) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into EmailNameAndAddress
	err = json.Unmarshal(data, &dst.EmailNameAndAddress);
	if err == nil {
		jsonEmailNameAndAddress, _ := json.Marshal(dst.EmailNameAndAddress)
		if string(jsonEmailNameAndAddress) == "{}" { // empty struct
			dst.EmailNameAndAddress = nil
		} else {
			return nil // data stored in dst.EmailNameAndAddress, return on the first match
		}
	} else {
		dst.EmailNameAndAddress = nil
	}

	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(EmailData)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EmailData) MarshalJSON() ([]byte, error) {
	if src.EmailNameAndAddress != nil {
		return json.Marshal(&src.EmailNameAndAddress)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEmailData struct {
	value *EmailData
	isSet bool
}

func (v NullableEmailData) Get() *EmailData {
	return v.value
}

func (v *NullableEmailData) Set(val *EmailData) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailData) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailData(val *EmailData) *NullableEmailData {
	return &NullableEmailData{value: val, isSet: true}
}

func (v NullableEmailData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


