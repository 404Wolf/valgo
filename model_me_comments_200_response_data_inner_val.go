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

// checks if the MeComments200ResponseDataInnerVal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeComments200ResponseDataInnerVal{}

// MeComments200ResponseDataInnerVal struct for MeComments200ResponseDataInnerVal
type MeComments200ResponseDataInnerVal struct {
	// Name of the val that is being commented on
	Name string `json:"name"`
	Id string `json:"id"`
	Version int32 `json:"version"`
	// This resource's privacy setting. Unlisted resources do not appear on profile pages or elsewhere, but you can link to them.
	Privacy string `json:"privacy"`
	Author NullableAliasVal200ResponseAuthor `json:"author"`
}

type _MeComments200ResponseDataInnerVal MeComments200ResponseDataInnerVal

// NewMeComments200ResponseDataInnerVal instantiates a new MeComments200ResponseDataInnerVal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeComments200ResponseDataInnerVal(name string, id string, version int32, privacy string, author NullableAliasVal200ResponseAuthor) *MeComments200ResponseDataInnerVal {
	this := MeComments200ResponseDataInnerVal{}
	this.Name = name
	this.Id = id
	this.Version = version
	this.Privacy = privacy
	this.Author = author
	return &this
}

// NewMeComments200ResponseDataInnerValWithDefaults instantiates a new MeComments200ResponseDataInnerVal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeComments200ResponseDataInnerValWithDefaults() *MeComments200ResponseDataInnerVal {
	this := MeComments200ResponseDataInnerVal{}
	return &this
}

// GetName returns the Name field value
func (o *MeComments200ResponseDataInnerVal) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MeComments200ResponseDataInnerVal) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MeComments200ResponseDataInnerVal) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value
func (o *MeComments200ResponseDataInnerVal) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MeComments200ResponseDataInnerVal) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MeComments200ResponseDataInnerVal) SetId(v string) {
	o.Id = v
}

// GetVersion returns the Version field value
func (o *MeComments200ResponseDataInnerVal) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *MeComments200ResponseDataInnerVal) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *MeComments200ResponseDataInnerVal) SetVersion(v int32) {
	o.Version = v
}

// GetPrivacy returns the Privacy field value
func (o *MeComments200ResponseDataInnerVal) GetPrivacy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value
// and a boolean to check if the value has been set.
func (o *MeComments200ResponseDataInnerVal) GetPrivacyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privacy, true
}

// SetPrivacy sets field value
func (o *MeComments200ResponseDataInnerVal) SetPrivacy(v string) {
	o.Privacy = v
}

// GetAuthor returns the Author field value
// If the value is explicit nil, the zero value for AliasVal200ResponseAuthor will be returned
func (o *MeComments200ResponseDataInnerVal) GetAuthor() AliasVal200ResponseAuthor {
	if o == nil || o.Author.Get() == nil {
		var ret AliasVal200ResponseAuthor
		return ret
	}

	return *o.Author.Get()
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MeComments200ResponseDataInnerVal) GetAuthorOk() (*AliasVal200ResponseAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return o.Author.Get(), o.Author.IsSet()
}

// SetAuthor sets field value
func (o *MeComments200ResponseDataInnerVal) SetAuthor(v AliasVal200ResponseAuthor) {
	o.Author.Set(&v)
}

func (o MeComments200ResponseDataInnerVal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeComments200ResponseDataInnerVal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["id"] = o.Id
	toSerialize["version"] = o.Version
	toSerialize["privacy"] = o.Privacy
	toSerialize["author"] = o.Author.Get()
	return toSerialize, nil
}

func (o *MeComments200ResponseDataInnerVal) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"id",
		"version",
		"privacy",
		"author",
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

	varMeComments200ResponseDataInnerVal := _MeComments200ResponseDataInnerVal{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMeComments200ResponseDataInnerVal)

	if err != nil {
		return err
	}

	*o = MeComments200ResponseDataInnerVal(varMeComments200ResponseDataInnerVal)

	return err
}

type NullableMeComments200ResponseDataInnerVal struct {
	value *MeComments200ResponseDataInnerVal
	isSet bool
}

func (v NullableMeComments200ResponseDataInnerVal) Get() *MeComments200ResponseDataInnerVal {
	return v.value
}

func (v *NullableMeComments200ResponseDataInnerVal) Set(val *MeComments200ResponseDataInnerVal) {
	v.value = val
	v.isSet = true
}

func (v NullableMeComments200ResponseDataInnerVal) IsSet() bool {
	return v.isSet
}

func (v *NullableMeComments200ResponseDataInnerVal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeComments200ResponseDataInnerVal(val *MeComments200ResponseDataInnerVal) *NullableMeComments200ResponseDataInnerVal {
	return &NullableMeComments200ResponseDataInnerVal{value: val, isSet: true}
}

func (v NullableMeComments200ResponseDataInnerVal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeComments200ResponseDataInnerVal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


