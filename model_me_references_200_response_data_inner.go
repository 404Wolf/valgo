/*
Val Town API

Val Town’s public API  OpenAPI JSON endpoint:  https://api.val.town/openapi.json

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package valgo

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the MeReferences200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeReferences200ResponseDataInner{}

// MeReferences200ResponseDataInner A description of a dependency from one val (reference) to another (dependsOn) that was introduced at a particular time.
type MeReferences200ResponseDataInner struct {
	Reference MeReferences200ResponseDataInnerReference `json:"reference"`
	DependsOn MeReferences200ResponseDataInnerReference `json:"dependsOn"`
	ReferencedAt time.Time `json:"referencedAt"`
}

type _MeReferences200ResponseDataInner MeReferences200ResponseDataInner

// NewMeReferences200ResponseDataInner instantiates a new MeReferences200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeReferences200ResponseDataInner(reference MeReferences200ResponseDataInnerReference, dependsOn MeReferences200ResponseDataInnerReference, referencedAt time.Time) *MeReferences200ResponseDataInner {
	this := MeReferences200ResponseDataInner{}
	this.Reference = reference
	this.DependsOn = dependsOn
	this.ReferencedAt = referencedAt
	return &this
}

// NewMeReferences200ResponseDataInnerWithDefaults instantiates a new MeReferences200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeReferences200ResponseDataInnerWithDefaults() *MeReferences200ResponseDataInner {
	this := MeReferences200ResponseDataInner{}
	return &this
}

// GetReference returns the Reference field value
func (o *MeReferences200ResponseDataInner) GetReference() MeReferences200ResponseDataInnerReference {
	if o == nil {
		var ret MeReferences200ResponseDataInnerReference
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *MeReferences200ResponseDataInner) GetReferenceOk() (*MeReferences200ResponseDataInnerReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *MeReferences200ResponseDataInner) SetReference(v MeReferences200ResponseDataInnerReference) {
	o.Reference = v
}

// GetDependsOn returns the DependsOn field value
func (o *MeReferences200ResponseDataInner) GetDependsOn() MeReferences200ResponseDataInnerReference {
	if o == nil {
		var ret MeReferences200ResponseDataInnerReference
		return ret
	}

	return o.DependsOn
}

// GetDependsOnOk returns a tuple with the DependsOn field value
// and a boolean to check if the value has been set.
func (o *MeReferences200ResponseDataInner) GetDependsOnOk() (*MeReferences200ResponseDataInnerReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DependsOn, true
}

// SetDependsOn sets field value
func (o *MeReferences200ResponseDataInner) SetDependsOn(v MeReferences200ResponseDataInnerReference) {
	o.DependsOn = v
}

// GetReferencedAt returns the ReferencedAt field value
func (o *MeReferences200ResponseDataInner) GetReferencedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ReferencedAt
}

// GetReferencedAtOk returns a tuple with the ReferencedAt field value
// and a boolean to check if the value has been set.
func (o *MeReferences200ResponseDataInner) GetReferencedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferencedAt, true
}

// SetReferencedAt sets field value
func (o *MeReferences200ResponseDataInner) SetReferencedAt(v time.Time) {
	o.ReferencedAt = v
}

func (o MeReferences200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeReferences200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reference"] = o.Reference
	toSerialize["dependsOn"] = o.DependsOn
	toSerialize["referencedAt"] = o.ReferencedAt
	return toSerialize, nil
}

func (o *MeReferences200ResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reference",
		"dependsOn",
		"referencedAt",
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

	varMeReferences200ResponseDataInner := _MeReferences200ResponseDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMeReferences200ResponseDataInner)

	if err != nil {
		return err
	}

	*o = MeReferences200ResponseDataInner(varMeReferences200ResponseDataInner)

	return err
}

type NullableMeReferences200ResponseDataInner struct {
	value *MeReferences200ResponseDataInner
	isSet bool
}

func (v NullableMeReferences200ResponseDataInner) Get() *MeReferences200ResponseDataInner {
	return v.value
}

func (v *NullableMeReferences200ResponseDataInner) Set(val *MeReferences200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMeReferences200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMeReferences200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeReferences200ResponseDataInner(val *MeReferences200ResponseDataInner) *NullableMeReferences200ResponseDataInner {
	return &NullableMeReferences200ResponseDataInner{value: val, isSet: true}
}

func (v NullableMeReferences200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeReferences200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

